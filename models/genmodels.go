package models

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/go-openapi/spec"

	"github.com/Clever/go-utils/stringset"
	"github.com/Clever/wag/swagger"
	"github.com/Clever/wag/templates"
	"github.com/Clever/wag/utils"

	"github.com/go-swagger/go-swagger/generator"
)

// Generate writes the files to the client directories
func Generate(packageName string, s spec.Swagger) error {

	tmpFile, err := swagger.WriteToFile(&s)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile)

	// generate models with go-swagger
	if err := generator.GenerateServer("", []string{}, []string{}, &generator.GenOpts{
		Spec:           tmpFile,
		ModelPackage:   "models",
		Target:         fmt.Sprintf("%s/src/%s/", os.Getenv("GOPATH"), packageName),
		IncludeModel:   true,
		IncludeHandler: false,
		IncludeSupport: false,
	}); err != nil {
		return fmt.Errorf("error generating go-swagger models: %s", err)
	}

	if err := generateOutputs(packageName, s); err != nil {
		return fmt.Errorf("error generating outputs: %s", err)
	}
	if err := generateInputs(packageName, s.Paths); err != nil {
		return fmt.Errorf("error generating inputs: %s", err)
	}
	return nil
}

func generateInputs(packageName string, paths *spec.Paths) error {

	g := swagger.Generator{PackageName: packageName}

	g.Printf(`
package models

import(
		"encoding/json"
		"strconv"

		"github.com/go-openapi/validate"
		"github.com/go-openapi/strfmt"
)

// These imports may not be used depending on the input parameters
var _ = json.Marshal
var _ = strconv.FormatInt
var _ = validate.Maximum
var _ = strfmt.NewFormats
`)

	for _, pathKey := range swagger.SortedPathItemKeys(paths.Paths) {
		path := paths.Paths[pathKey]
		pathItemOps := swagger.PathItemOperations(path)
		for _, opKey := range swagger.SortedOperationsKeys(pathItemOps) {
			op := pathItemOps[opKey]
			// Do not generate an input struct + validation for an
			// operation that has a single, schema'd input.
			// The input to these will be the model generated for
			// the schema.
			if singleSchemaedBodyParameter, _ := swagger.SingleSchemaedBodyParameter(op); singleSchemaedBodyParameter {
				continue
			}
			if err := printInputStruct(&g, op); err != nil {
				return err
			}
			if err := printInputValidation(&g, op); err != nil {
				return err
			}
		}
	}

	return g.WriteFile("models/inputs.go")
}

func printInputStruct(g *swagger.Generator, op *spec.Operation) error {
	capOpID := swagger.Capitalize(op.ID)
	g.Printf("// %sInput holds the input parameters for a %s operation.\n", capOpID, op.ID)
	g.Printf("type %sInput struct {\n", capOpID)

	for _, param := range op.Parameters {
		if param.In == "formData" {
			return fmt.Errorf("input parameters with 'In' formData are not supported")
		}

		var typeName string
		var err error
		if param.In != "body" {
			typeName, err = swagger.ParamToType(param, true)
			if err != nil {
				return err
			}
		} else {
			typeName, err = swagger.TypeFromSchema(param.Schema, false)
			if err != nil {
				return err
			}
			// All schema types are pointers
			typeName = "*" + typeName
		}

		g.Printf("\t%s %s\n", swagger.StructParamName(param), typeName)
	}
	g.Printf("}\n\n")

	return nil
}

func printInputValidation(g *swagger.Generator, op *spec.Operation) error {
	singleStringPathParameter, paramName := swagger.SingleStringPathParameter(op)
	if singleStringPathParameter {
		capOpID := swagger.Capitalize(op.ID)
		g.Printf("// Validate%sInput returns an error if the input parameter doesn't\n",
			capOpID)
		g.Printf("// satisfy the requirements in the swagger yml file.\n")
		g.Printf("func Validate%sInput(%s string) error{\n", capOpID, paramName)
	} else {
		capOpID := swagger.Capitalize(op.ID)
		g.Printf("// Validate returns an error if any of the %sInput parameters don't satisfy the\n",
			capOpID)
		g.Printf("// requirements from the swagger yml file.\n")
		g.Printf("func (i %sInput) Validate() error{\n", capOpID)
	}

	for _, param := range op.Parameters {
		if param.In == "body" {
			g.Printf("\tif err := i.%s.Validate(nil); err != nil {\n", swagger.StructParamName(param))
			g.Printf("\t\treturn err\n")
			g.Printf("\t}\n\n")
		}

		validations, err := swagger.ParamToValidationCode(param)
		if err != nil {
			return err
		}
		for _, validation := range validations {
			if param.Required {
				if singleStringPathParameter {
					// replace i.<Param> with <param>
					validation = strings.Replace(validation,
						fmt.Sprintf("i.%s", utils.CamelCase(param.Name, true)),
						paramName, -1)
				}
				g.Printf(errCheck(validation))
			} else {
				g.Printf("\tif i.%s != nil {\n", swagger.StructParamName(param))
				g.Printf(errCheck(validation))
				g.Printf("\t}\n")
			}
		}
	}
	g.Printf("\treturn nil\n")
	g.Printf("}\n\n")

	return nil
}

// errCheck returns an if err := ifCondition; err != nil { return err } function
func errCheck(ifCondition string) string {
	return fmt.Sprintf(
		`	if err := %s; err != nil {
		return err
	}
`, ifCondition)
}

func generateOutputs(packageName string, s spec.Swagger) error {
	g := swagger.Generator{PackageName: packageName}

	g.Printf("package models\n\n")

	// It's a bit wonky that we're writing these into output.go instead of the file
	// defining each of the types, but I think that's okay for now. We can clean this
	// up if it becomes confusing.
	errorMethodCode, err := generateErrorMethods(&s)
	if err != nil {
		return err
	}
	g.Printf(errorMethodCode)

	for _, pathKey := range swagger.SortedPathItemKeys(s.Paths.Paths) {
		path := s.Paths.Paths[pathKey]
		pathItemOps := swagger.PathItemOperations(path)
		for _, opKey := range swagger.SortedOperationsKeys(pathItemOps) {
			op := pathItemOps[opKey]
			capOpID := swagger.Capitalize(op.ID)

			// We classify response keys into three types:
			// 1. 200-399 - these are "success" responses and implement the Output interface
			// 	defined above
			// 2. 400-599 - these are "failure" responses and implement the error interface
			// 3. Default - this is defined as a 500
			successTypes, err := generateSuccessTypes(capOpID, op.Responses.StatusCodeResponses)
			if err != nil {
				return err
			}
			g.Printf(successTypes)

		}
	}
	return g.WriteFile("models/outputs.go")
}

func generateSuccessTypes(capOpID string, responses map[int]spec.Response) (string, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("// %sOutput defines the success output interface for %s.\n",
		capOpID, capOpID))
	buf.WriteString(fmt.Sprintf("type %sOutput interface {\n", capOpID))
	buf.WriteString(fmt.Sprintf("\t%sStatusCode() int\n", capOpID))
	buf.WriteString(fmt.Sprintf("}\n\n"))

	var successStatusCodes []int
	for _, statusCode := range swagger.SortedStatusCodeKeys(responses) {
		if statusCode >= 400 {
			continue
		}
		successStatusCodes = append(successStatusCodes, statusCode)
	}

	// We don't need to generate any success types if there is one or less success responses. In that
	// case we can just use the raw type
	if len(successStatusCodes) < 2 {
		return "", nil
	}

	for _, statusCode := range successStatusCodes {
		typeString, err := generateType(capOpID, statusCode, responses[statusCode])
		if err != nil {
			return "", err
		}
		buf.WriteString(typeString)
	}
	return buf.String(), nil
}

func generateType(capOpID string, statusCode int, response spec.Response) (string, error) {
	outputName := fmt.Sprintf("%s%dOutput", capOpID, statusCode)
	typeName, err := swagger.TypeFromSchema(response.Schema, false)
	if err != nil {
		return "", err
	}

	fields := typeTemplateFields{
		Output:     outputName,
		StatusCode: statusCode,
		OpName:     capOpID,
		Type:       typeName,
		ErrorType:  statusCode >= 400,
	}
	return templates.WriteTemplate(typeTemplate, fields)
}

type typeTemplateFields struct {
	Output     string
	StatusCode int
	OpName     string
	Type       string
	ErrorType  bool
}

var typeTemplate = `
	// {{.Output}} defines the {{.StatusCode}} status code response for {{.OpName}}.
	type {{.Output}} {{.Type}}

	// {{.OpName}}StatusCode returns the status code for the operation.
	func (o {{.Output}}) {{.OpName}}StatusCode() int {
		return {{.StatusCode}}
	}
`

// generateErrorMethods finds all responses all error responses and generates an error
// method for them.
func generateErrorMethods(s *spec.Swagger) (string, error) {
	errorTypes := stringset.New()

	for _, pathKey := range swagger.SortedPathItemKeys(s.Paths.Paths) {
		path := s.Paths.Paths[pathKey]
		pathItemOps := swagger.PathItemOperations(path)
		for _, opKey := range swagger.SortedOperationsKeys(pathItemOps) {
			op := pathItemOps[opKey]
			for _, statusCode := range swagger.SortedStatusCodeKeys(op.Responses.StatusCodeResponses) {

				if statusCode < 400 {
					continue
				}

				typeName, _ := swagger.OutputType(s, op, statusCode)
				if strings.HasPrefix(typeName, "models.") {
					typeName = typeName[7:]
				}
				errorTypes.Add(typeName)
			}
		}
	}

	sortedErrors := errorTypes.ToList()
	sort.Strings(sortedErrors)

	var buf bytes.Buffer
	for _, errorType := range sortedErrors {
		buf.WriteString(fmt.Sprintf(`
func (o %s) Error() string {
	return o.Msg
}

`, errorType))
	}

	return buf.String(), nil

}
