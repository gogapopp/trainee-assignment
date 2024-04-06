api-gen:
	@oapi-codegen -package=handler -generate="chi-server,types,spec" api/api.yaml > internal/handler/api.gen.go

run: api-gen
	@go run cmd/main.go