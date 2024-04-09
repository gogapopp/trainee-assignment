oapi-gen:
	@oapi-codegen -package=handler -generate="chi-server,types,spec" api.yaml > internal/handler/api.gen.go

docker-compose-up:
	@docker-compose up -d

docker-compose-down:
	@docker-compose down

run: oapi-gen docker-compose-up
	@go run cmd/main.go

stop: docker-compose-down