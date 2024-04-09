oapi-gen:
	@oapi-codegen -package=handler -generate="chi-server,types,spec" api.yaml > internal/handler/api.gen.go

run: oapi-gen
	@go run cmd/main.go