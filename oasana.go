package oasana

//go:generate oapi-codegen -generate client -o client.gen.go -package oasana openapi.yaml
//go:generate oapi-codegen -generate types -templates=./templates -o schema.gen.go -package oasana openapi.yaml
