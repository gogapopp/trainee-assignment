oapi-gen:
	@oapi-codegen -package=handler -generate="chi-server,types,spec" api.yaml > internal/handler/api.gen.go

docker-compose-up:
	@docker-compose up --build

docker-compose-down:
	@docker-compose down

run: docker-compose-up

stop: docker-compose-down