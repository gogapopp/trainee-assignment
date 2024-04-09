oapi-gen:
	@oapi-codegen -package=handler -generate="chi-server,types,spec" api.yaml > internal/handler/api.gen.go

docker-compose:
	@docker-compose up -d

run: oapi-gen docker-compose
	@go run cmd/main.go