package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"wag/hardcoded"
	"wag/models"
	"wag/server"
	"wag/swagger"
	"wag/validation"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/swag"
)

// config contains the configuration of command line flags and configuration derived from command line flags
type config struct {
	outputPath    *string
	versionFlag   *bool
	swaggerFile   *string
	goPackageName *string

	goAbsolutePackagePath string
}

var version string

func main() {
	conf := config{
		swaggerFile:   flag.String("file", "swagger.yml", "the spec file to use"),
		goPackageName: flag.String("go-package", "", "package of the generated go code"),
		outputPath:    flag.String("output-path", "", "relative output path of the generated go code"),
		versionFlag:   flag.Bool("version", false, "print the wag version and exit"),
	}
	flag.Parse()
	if *conf.versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	if err := conf.parse(); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	loads.AddLoader(fmts.YAMLMatcher, fmts.YAMLDoc)
	doc, err := loads.Spec(*conf.swaggerFile)
	if err != nil {
		log.Fatalf("Error loading swagger file: %v", err)
	}
	swaggerSpec := *doc.Spec()

	injectDefaultDefinitions(&swaggerSpec)

	if err := validation.Validate(*doc, false); err != nil {
		log.Fatalf("swagger file not valid: %v", err)
	}

	err = swagger.ValidateResponses(swaggerSpec)
	if err != nil {
		log.Fatalf("failed processing the swagger spec: %v", err)
	}

	if err := generateGoModels(*conf.goPackageName, conf.goAbsolutePackagePath, *conf.outputPath, swaggerSpec); err != nil {
		log.Fatalf("error generating go models: %v", err)
	}

	if err := generateServer(*conf.goPackageName, conf.goAbsolutePackagePath, *conf.outputPath, swaggerSpec); err != nil {
		log.Fatalf("error generating server: %v", err)
	}
}

func generateGoModels(packageName, basePath, outputPath string, swaggerSpec spec.Swagger) error {
	if err := prepareDir(filepath.Join(basePath, "models")); err != nil {
		return err
	}
	if err := models.Generate(packageName, basePath, outputPath, swaggerSpec); err != nil {
		return fmt.Errorf("error generating models: %v", err)
	}
	return nil
}

func generateServer(goPackageName, basePath, outputPath string, swaggerSpec spec.Swagger) error {
	if err := prepareDir(filepath.Join(basePath, "server")); err != nil {
		return err
	}
	if err := server.Generate(goPackageName, basePath, outputPath, swaggerSpec); err != nil {
		return fmt.Errorf("Failed to generate server: %s", err)
	}
	middlewareGenerator := swagger.Generator{BasePath: basePath}
	middlewareGenerator.Write(hardcoded.MustAsset("../_hardcoded/middleware.go"))
	if err := middlewareGenerator.WriteFile("server/middleware.go"); err != nil {
		return fmt.Errorf("Failed to copy middleware.go: %s", err)
	}

	return nil
}

func prepareDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("could not remove directory: %s, error :%v", dir, err)
	}

	if err := os.MkdirAll(dir, 0o700); err != nil {
		if !os.IsExist(err.(*os.PathError)) {
			return fmt.Errorf("could not create directory: %s, error: %v", dir, err)
		}
	}
	return nil
}

// parseCmdConfig determines the which code is generated and the location of the generated code
// from the command line arguments
func (c *config) parse() error {
	if err := c.setGoPaths(swag.StringValue(c.outputPath), swag.StringValue(c.goPackageName)); err != nil {
		return err
	}

	return nil
}

// setGoPaths sets the golang package path and package name. Go repos using modules may have a
// package name differing from its package path.
func (c *config) setGoPaths(outputPath, goPackageName string) error {
	// check if the repo uses modules
	if modFile, err := os.Open("go.mod"); err != nil {
		if _, ok := err.(*os.PathError); !ok {
			return fmt.Errorf("error checking if the repo contains a go.mod file: %v", err)
		}
		if goPackageName == "" {
			return fmt.Errorf("go-package is required")
		}
		// if the repo does not use modules, the package name is equivalent to the package path
		c.goAbsolutePackagePath = filepath.Join(os.Getenv("GOPATH"), "src", goPackageName)
	} else {
		defer modFile.Close()
		if "outputPath" == "" {
			return fmt.Errorf("output-path is required")
		}

		absolutePath, err := filepath.Abs(outputPath)
		if err != nil {
			return fmt.Errorf("converting output-path to absolute path: %v", err)
		}
		c.goAbsolutePackagePath = absolutePath

		*c.goPackageName = getModulePackageName(modFile, path.Clean(outputPath))

	}
	return nil
}

func getModulePackageName(modFile *os.File, outputPath string) string {
	// read first line of module file
	r := bufio.NewReader(modFile)
	b, _, err := r.ReadLine()
	if err != nil {
		log.Fatalf("Error checking module name: %s", err.Error())
	}

	// parse module path
	moduleName := strings.TrimPrefix(string(b), "module")
	moduleName = strings.TrimSpace(moduleName)

	return fmt.Sprintf("%v/%v", moduleName, outputPath)
}

// injectDefaultDefinitions injects default definitions
func injectDefaultDefinitions(swaggerSpec *spec.Swagger) {
	swaggerSpec.Definitions["UnknownResponse"] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type: spec.StringOrArray{"object"},
			Properties: map[string]spec.Schema{
				"statusCode": {
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{"integer"},
					},
				},
				"body": {
					SchemaProps: spec.SchemaProps{
						Type: spec.StringOrArray{"string"},
					},
				},
			},
		},
	}
}
