oapi-codegen -package=handler -generate="echo,types,spec" api/api.yaml > internal/handler/api.gen.go

oapi-codegen -package=handler -generate="chi-server,types,spec" api/api.yaml > internal/handler/api.gen.go