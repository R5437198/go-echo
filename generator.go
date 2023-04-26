// health check
//
//go:generate oapi-codegen -include-tags helth_check -o ./internal/presentation/controller/health_check/gen/types.go -generate types -package gen api.yaml
//go:generate oapi-codegen -include-tags helth_check -o ./internal/presentation/controller/health_check/gen/server.go -generate server -package gen api.yaml
package main
