package server

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-openapi/spec"

	"github.com/Clever/wag/swagger"
	"github.com/Clever/wag/templates"
)

// Generate server package for a swagger spec.
func Generate(packageName string, s spec.Swagger) error {

	if err := generateRouter(packageName, s, s.Paths); err != nil {
		return err
	}
	if err := generateInterface(packageName, &s, s.Info.InfoProps.Title, s.Paths); err != nil {
		return err
	}
	if err := generateHandlers(packageName, &s, s.Paths); err != nil {
		return err
	}
	return nil
}

type routerFunction struct {
	Method      string
	Path        string
	HandlerName string
	OpID        string
}

type routerTemplate struct {
	Title     string
	Functions []routerFunction
}

var routerTemplateStr = `
package server

// Code auto-generated. Do not edit.

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/Clever/kayvee-go.v5/logger"
	"gopkg.in/tylerb/graceful.v1"
	"github.com/Clever/go-process-metrics/metrics"
)

type contextKey struct{}

// Server defines a HTTP server that implements the Controller interface.
type Server struct {
	// Handler should generally not be changed. It exposed to make testing easier.
	Handler http.Handler
	addr string
	l *logger.Logger
}

// Serve starts the server. It will return if an error occurs.
func (s *Server) Serve() error {

	go func() {
		metrics.Log("{{.Title}}", 1*time.Minute)
	}()

	go func() {
		// This should never return. Listen on the pprof port
		log.Printf("PProf server crashed: %%s", http.ListenAndServe(":6060", nil))
	}()

	s.l.Counter("server-started")

	// Give the sever 30 seconds to shut down
	return graceful.RunWithErr(s.addr,30*time.Second,s.Handler)
}

type handler struct {
	Controller
}

// New returns a Server that implements the Controller interface. It will start when "Serve" is called.
func New(c Controller, addr string) *Server {
	r := mux.NewRouter()
	h := handler{Controller: c}

	l := logger.New("{{.Title}}")

	{{range $index, $val := .Functions}}
	r.Methods("{{$val.Method}}").Path("{{$val.Path}}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.FromContext(r.Context()).AddContext("op", "{{$val.OpID}}")
		h.{{$val.HandlerName}}Handler(r.Context(), w, r)
	})
	{{end}}

	handler := withMiddleware("{{.Title}}", r)
	return &Server{Handler: handler, addr: addr, l: l}
}
`

func generateRouter(packageName string, s spec.Swagger, paths *spec.Paths) error {

	var template routerTemplate
	template.Title = s.Info.Title
	for _, path := range swagger.SortedPathItemKeys(paths.Paths) {
		pathItem := paths.Paths[path]
		pathItemOps := swagger.PathItemOperations(pathItem)
		for _, method := range swagger.SortedOperationsKeys(pathItemOps) {
			op := pathItemOps[method]

			template.Functions = append(template.Functions, routerFunction{
				Method:      method,
				Path:        s.BasePath + path,
				HandlerName: swagger.Capitalize(op.ID),
				OpID:        op.ID,
			})
		}
	}

	routerCode, err := templates.WriteTemplate(routerTemplateStr, template)
	if err != nil {
		return err
	}
	g := swagger.Generator{PackageName: packageName}
	g.Printf(routerCode)
	return g.WriteFile("server/router.go")
}

type interfaceTemplate struct {
	Comment    string
	Definition string
}

type interfaceFileTemplate struct {
	ImportStatements string
	ServiceName      string
	Interfaces       []interfaceTemplate
}

var interfaceTemplateStr = `
package server

{{.ImportStatements}}

//go:generate $GOPATH/bin/mockgen -source=$GOFILE -destination=mock_controller.go -package=server

// Controller defines the interface for the {{.ServiceName}} service.
type Controller interface {

	{{range $interface := .Interfaces}}
		{{$interface.Comment}}
		{{$interface.Definition}}
	{{end}}
}
`

func generateInterface(packageName string, s *spec.Swagger, serviceName string, paths *spec.Paths) error {

	tmpl := interfaceFileTemplate{
		ImportStatements: swagger.ImportStatements([]string{"context", packageName + "/models"}),
		ServiceName:      serviceName,
	}

	for _, pathKey := range swagger.SortedPathItemKeys(paths.Paths) {
		path := paths.Paths[pathKey]
		pathItemOps := swagger.PathItemOperations(path)
		for _, method := range swagger.SortedOperationsKeys(pathItemOps) {
			interfaceComment, err := swagger.InterfaceComment(method, pathKey, s, pathItemOps[method])
			if err != nil {
				return err
			}
			tmpl.Interfaces = append(tmpl.Interfaces, interfaceTemplate{
				Comment:    interfaceComment,
				Definition: swagger.Interface(s, pathItemOps[method]),
			})
		}
	}

	interfaceCode, err := templates.WriteTemplate(interfaceTemplateStr, tmpl)
	if err != nil {
		return err
	}
	g := swagger.Generator{PackageName: packageName}
	g.Printf(interfaceCode)
	return g.WriteFile("server/interface.go")
}

func lowercase(input string) string {
	return strings.ToLower(input[0:1]) + input[1:]
}

type handlerFileTemplate struct {
	ImportStatements string
	// TODO: Think about possibly factoring this out...
	BaseStringToTypeCode string
	Handlers             []string
}

var handlerFileTemplateStr = `
package server

{{.ImportStatements}}

var _ = strconv.ParseInt
var _ = strfmt.Default
var _ = swag.ConvertInt32
var _ = errors.New
var _ = mux.Vars
var _ = bytes.Compare
var _ = ioutil.ReadAll

{{.BaseStringToTypeCode}}

func jsonMarshalNoError(i interface{}) string {
	bytes, err := json.Marshal(i)
	if err != nil {
		// This should never happen
		return ""
	}
	return string(bytes)
}

{{ range $handler := .Handlers }}
	{{ $handler }}
{{end}}
`

func generateHandlers(packageName string, s *spec.Swagger, paths *spec.Paths) error {

	tmpl := handlerFileTemplate{
		ImportStatements: swagger.ImportStatements([]string{"context", "github.com/gorilla/mux", "gopkg.in/Clever/kayvee-go.v5/logger",
			"net/http", "strconv", "encoding/json", "strconv", packageName + "/models",
			"github.com/go-openapi/strfmt", "github.com/go-openapi/swag", "io/ioutil", "bytes",
			"github.com/go-errors/errors"}),
		BaseStringToTypeCode: swagger.BaseStringToTypeCode(),
	}

	for _, pathKey := range swagger.SortedPathItemKeys(paths.Paths) {
		path := paths.Paths[pathKey]
		pathItemOps := swagger.PathItemOperations(path)
		for _, opKey := range swagger.SortedOperationsKeys(pathItemOps) {
			op := pathItemOps[opKey]

			operationHandler, err := generateOperationHandler(s, op)
			if err != nil {
				return err
			}
			tmpl.Handlers = append(tmpl.Handlers, operationHandler)
		}
	}

	handlerCode, err := templates.WriteTemplate(handlerFileTemplateStr, tmpl)
	if err != nil {
		return err
	}
	g := swagger.Generator{PackageName: packageName}
	g.Printf(handlerCode)
	return g.WriteFile("server/handlers.go")
}

var jsonMarshalString = `

`

// generateOperationHandler generates the handler code for a single handler
func generateOperationHandler(s *spec.Swagger, op *spec.Operation) (string, error) {
	typeToCode := make(map[string]int)
	emptyResponseCode := 200
	codeToType := swagger.CodeToTypeMap(s, op, false)
	typeToCode, err := swagger.TypeToCodeMap(s, op)
	if err != nil {
		return "", err
	}
	if empty, ok := typeToCode[""]; ok {
		emptyResponseCode = empty
		delete(typeToCode, "")
	}

	singleSchemaedBodyParameter, _ := swagger.SingleSchemaedBodyParameter(op)
	singleStringPathParameter, singleStringPathParameterVarName := swagger.SingleStringPathParameter(op)
	handlerOp := handlerOp{
		Op:                               swagger.Capitalize(op.ID),
		SuccessReturnType:                !swagger.NoSuccessType(op),
		HasParams:                        len(op.Parameters) != 0,
		SingleSchemaedBodyParameter:      singleSchemaedBodyParameter,
		EmptyStatusCode:                  emptyResponseCode,
		TypesToStatusCodes:               typeToCode,
		SingleStringPathParameter:        singleStringPathParameter,
		SingleStringPathParameterVarName: singleStringPathParameterVarName,
		StatusCodeToType:                 codeToType,
	}
	handlerCode, err := templates.WriteTemplate(handlerTemplate, handlerOp)
	if err != nil {
		return "", err
	}

	newInputCode, err := generateNewInput(op)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	buf.WriteString(handlerCode)
	buf.WriteString(newInputCode)

	return buf.String(), nil
}

// handlerOp contains the template variables for the handlerTemplate
type handlerOp struct {
	Op                               string
	SuccessReturnType                bool
	HasParams                        bool
	SingleSchemaedBodyParameter      bool
	EmptyStatusCode                  int
	TypesToStatusCodes               map[string]int
	SingleStringPathParameter        bool
	SingleStringPathParameterVarName string
	StatusCodeToType                 map[int]string
}

var handlerTemplate = `
// statusCodeFor{{.Op}} returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeFor{{.Op}}(obj interface{}) int {

	switch obj.(type) {
	{{ range $type, $code := .TypesToStatusCodes }}
   	case {{$type}}:
   		return {{$code}}
	{{ end }}
	default:
		return -1
	}
}

func (h handler) {{.Op}}Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
{{if .HasParams}}
{{if .SingleStringPathParameter}}
	{{.SingleStringPathParameterVarName}}, err := new{{.Op}}Input(r)
{{else}}
	input, err := new{{.Op}}Input(r)
{{end}}
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError({{index .StatusCodeToType 400}}{Message: err.Error()}), http.StatusBadRequest)
		return
	}

{{if .SingleStringPathParameter}}
	err = models.Validate{{.Op}}Input({{.SingleStringPathParameterVarName}})
{{else}}
	err = input.Validate({{if .SingleSchemaedBodyParameter}}nil{{end}})
{{end}}
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError({{index .StatusCodeToType 400}}{Message: err.Error()}), http.StatusBadRequest)
		return
	}

{{if .SuccessReturnType}}
	resp, err := h.{{.Op}}(ctx, {{if .SingleStringPathParameter}}{{.SingleStringPathParameterVarName}}{{else}}input{{end}})
{{else}}
	err = h.{{.Op}}(ctx, {{if .SingleStringPathParameter}}{{.SingleStringPathParameterVarName}}{{else}}input{{end}})
{{end}}
{{else}}
{{if .SuccessReturnType}}
	resp, err := h.{{.Op}}(ctx)
{{else}}
	err := h.{{.Op}}(ctx)
{{end}}
{{end}}
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeFor{{.Op}}(err)
		if statusCode == -1 {
			err = {{index .StatusCodeToType 500}}{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

{{if .SuccessReturnType}}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError({{index .StatusCodeToType 500}}{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeFor{{.Op}}(resp))
	w.Write(respBytes)
{{else}}
	w.WriteHeader({{.EmptyStatusCode}})
	w.Write([]byte(""))
{{end}}
}
`

type singleStringPathParameterTemplateData struct {
	Op           string
	ParamName    string
	ParamVarName string
}

var singleStringPathParameterTemplate = `
// new{{.Op}}Input takes in an http.Request an returns the {{.ParamName}} parameter
// that it contains. It returns an error if the request doesn't contain the parameter.
func new{{.Op}}Input(r *http.Request) (string, error) {
	{{.ParamVarName}} := mux.Vars(r)["{{.ParamName}}"]
	if len({{.ParamVarName}}) == 0 {
		return "", errors.New("Parameter {{.ParamName}} must be specified")
	}
	return {{.ParamVarName}}, nil
}
`

func generateNewInput(op *spec.Operation) (string, error) {
	capOpID := swagger.Capitalize(op.ID)

	singleStringPathParameter, paramVarName := swagger.SingleStringPathParameter(op)
	if singleStringPathParameter {
		return templates.WriteTemplate(singleStringPathParameterTemplate,
			singleStringPathParameterTemplateData{
				Op:           capOpID,
				ParamName:    op.Parameters[0].Name,
				ParamVarName: paramVarName,
			})
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("// new%sInput takes in an http.Request an returns the input struct.\n", capOpID))
	singleSchemaedBodyParameter, opModel := swagger.SingleSchemaedBodyParameter(op)
	if singleSchemaedBodyParameter {
		buf.WriteString(fmt.Sprintf("func new%sInput(r *http.Request) (*models.%s, error) {\n",
			capOpID, opModel))
		buf.WriteString(fmt.Sprintf("\tvar input models.%s\n\n", opModel))
	} else {
		buf.WriteString(fmt.Sprintf("func new%sInput(r *http.Request) (*models.%sInput, error) {\n",
			capOpID, capOpID))
		buf.WriteString(fmt.Sprintf("\tvar input models.%sInput\n\n", capOpID))
	}

	buf.WriteString(fmt.Sprintf("\tvar err error\n"))
	buf.WriteString(fmt.Sprintf("\t_ = err\n\n"))

	for _, param := range op.Parameters {

		camelParamName := swagger.StructParamName(param)
		paramVarName := lowercase(camelParamName)

		if param.In != "body" {
			if param.Type == "array" && param.In == "query" {
				buf.WriteString(fmt.Sprintf("\tif %s, ok := r.URL.Query()[\"%s\"]; ok {\n\t\tinput.%s = %s\n\t}\n", paramVarName, param.Name, camelParamName, paramVarName))
			} else {
				extractCode := ""
				switch param.In {
				case "query":
					extractCode = fmt.Sprintf("r.URL.Query().Get(\"%s\")", param.Name)
				case "path":
					extractCode = fmt.Sprintf("mux.Vars(r)[\"%s\"]", param.Name)
				case "header":
					extractCode = fmt.Sprintf("r.Header.Get(\"%s\")", param.Name)
				}
				buf.WriteString(fmt.Sprintf("\t%sStr := %s\n", paramVarName, extractCode))

				if param.Required {
					buf.WriteString(fmt.Sprintf("\tif len(%sStr) == 0{\n", paramVarName))
					buf.WriteString(fmt.Sprintf("\t\treturn nil, errors.New(\"Parameter must be specified\")\n"))
					buf.WriteString(fmt.Sprintf("\t}\n"))
				} else if param.Default != nil {
					buf.WriteString(fmt.Sprintf("\tif len(%sStr) == 0 {\n", paramVarName))
					buf.WriteString(fmt.Sprintf("\t\t// Use the default value\n"))
					buf.WriteString(fmt.Sprintf("\t\t%sStr = \"%s\"\n", paramVarName, swagger.DefaultAsString(param)))
					buf.WriteString(fmt.Sprintf("\t}\n"))
				}

				buf.WriteString(fmt.Sprintf("\tif len(%sStr) != 0 {\n", paramVarName))

				typeName, err := swagger.ParamToType(param, false)
				if err != nil {
					return "", err
				}
				typeCode, err := swagger.StringToTypeCode(fmt.Sprintf("%sStr", paramVarName), param)
				if err != nil {
					return "", err
				}
				buf.WriteString(fmt.Sprintf("\t\tvar %sTmp %s\n", paramVarName, typeName))
				buf.WriteString(fmt.Sprintf("\t\t%sTmp, err = %s\n", paramVarName, typeCode))
				buf.WriteString(fmt.Sprintf("\t\tif err != nil {\n"))
				buf.WriteString(fmt.Sprintf("\t\t\treturn nil, err\n"))
				buf.WriteString(fmt.Sprintf("\t\t}\n"))

				// TODO: Factor this out...
				if param.Required || param.Type == "array" {
					buf.WriteString(fmt.Sprintf("\t\tinput.%s = %sTmp\n\n", camelParamName, paramVarName))
				} else {
					buf.WriteString(fmt.Sprintf("\t\tinput.%s = &%sTmp\n\n", camelParamName, paramVarName))
				}

				buf.WriteString(fmt.Sprintf("\t}\n"))
			}
		} else {
			if param.Schema == nil {
				return "", fmt.Errorf("body parameters must have a schema defined")
			}
			typeName, err := swagger.TypeFromSchema(param.Schema, true)
			if err != nil {
				return "", err
			}

			buf.WriteString(fmt.Sprintf("\tdata, err := ioutil.ReadAll(r.Body)\n"))

			if param.Required {
				buf.WriteString(fmt.Sprintf("\tif len(data) == 0 {\n"))
				buf.WriteString(fmt.Sprintf("\t\treturn nil, errors.New(\"Parameter must be specified\")\n"))
				buf.WriteString(fmt.Sprintf("\t}\n"))
			}

			buf.WriteString(fmt.Sprintf("\tif len(data) > 0 {"))

			if singleSchemaedBodyParameter {
				buf.WriteString(fmt.Sprintf("\t\tif err := json.NewDecoder(bytes.NewReader(data)).Decode(&input); err != nil {\n"))
			} else {
				// Initialize the pointer in the object
				buf.WriteString(fmt.Sprintf("\t\tinput.%s = &%s{}\n", camelParamName, typeName))
				buf.WriteString(fmt.Sprintf("\t\tif err := json.NewDecoder(bytes.NewReader(data)).Decode(input.%s); err != nil {\n", camelParamName))
			}
			buf.WriteString(fmt.Sprintf("\t\t\treturn nil, err\n"))
			buf.WriteString(fmt.Sprintf("\t\t}\n"))
			buf.WriteString(fmt.Sprintf("\t}\n"))

		}
	}
	buf.WriteString(fmt.Sprintf("\n"))

	buf.WriteString(fmt.Sprintf("\treturn &input, nil\n"))
	buf.WriteString(fmt.Sprintf("}\n\n"))

	return buf.String(), nil
}
