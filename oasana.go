package oasana

//go:generate oapi-codegen -generate client -o client.gen.go -package asana openapi.yaml
//go:generate oapi-codegen -generate types -o schema.gen.go -package oasana openapi.yaml
