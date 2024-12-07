package goclient

import (
	"path"

	"github.com/go-openapi/spec"
)

// IsBinaryBody returns true if the format of the body of the operation is binary
func IsBinaryBody(op *spec.Operation, definitions map[string]spec.Schema) bool {
	for _, param := range op.Parameters {
		if param.In == "body" {
			return IsBinaryParam(param, definitions)
		}
	}
	return false
}

// IsBinaryParam returns true of the format of the parameter is binary
func IsBinaryParam(param spec.Parameter, definitions map[string]spec.Schema) bool {
	definitionName := path.Base(param.Schema.Ref.Ref.GetURL().String())
	return definitions[definitionName].Format == "binary"
}
