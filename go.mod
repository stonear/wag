module github.com/Clever/wag/v8

go 1.16

require (
	github.com/Clever/go-utils v0.0.0-20150501165843-abc25366fa8e
	github.com/Clever/wag/v8/samples v0.0.0-internal
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/awslabs/goformation/v2 v2.3.1
	github.com/go-openapi/errors v0.19.4
	github.com/go-openapi/jsonreference v0.19.3
	github.com/go-openapi/loads v0.19.5
	github.com/go-openapi/spec v0.19.7
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.14
	github.com/go-openapi/validate v0.19.7
	github.com/go-swagger/go-swagger v0.23.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/kevinburke/go-bindata v3.22.0+incompatible
	github.com/pelletier/go-toml v1.9.1 // indirect
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/otel v0.15.0
	go.opentelemetry.io/otel/sdk v0.15.0
	google.golang.org/grpc v1.36.0 // indirect
	gopkg.in/Clever/kayvee-go.v6 v6.27.0
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)

// For validate, something seems to have broken or changed between this pinned version and 0.16.0  (commit 7c1911976134d3a24d0c03127505163c9f16aa3b)
// That version suddenly stops accepting almost anything inline inside the `paths` block.
// TODO remove this pin.
replace github.com/go-openapi/validate => github.com/go-openapi/validate v0.0.0-20180703152151-9a6e517cddf1 // pre-modules tag 0.15.0

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.0.0

replace github.com/Clever/wag/v8/samples v0.0.0-internal => ./samples
