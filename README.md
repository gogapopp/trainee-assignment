# banner-service

## Protocol

Described in the file [api.yml](api.yml).

## Generating code from a specification

Install [oapi-codegen](https://github.com/deepmap/oapi-codegen/) and generate:

```bash
oapi-codegen -package=handler -generate="chi-server,types,spec" api.yaml > internal/handler/api.gen.go
```

## How to compile

```bash
make run
```

## Additional tasks
- [ ] Адаптировать систему для значительного увеличения количества тегов и фичей, при котором допускается увеличение времени исполнения по редко запрашиваемым тегам и фичам.
- [ ] Провести нагрузочное тестирование полученного решения и приложить результаты тестирования к решению.
- [ ] Иногда получается так, что необходимо вернуться к одной из трех предыдущих версий баннера в связи с найденной ошибкой в логике, тексте и т.д. Измените API таким образом, чтобы можно было просмотреть существующие версии баннера и выбрать подходящую версию.
- [x] Добавить метод удаления баннеров по фиче или тегу, время ответа которого не должно превышать 100 мс, независимо от количества баннеров. В связи с небольшим временем ответа метода, рекомендуется ознакомиться с механизмом выполнения отложенных действий.
- [ ] Реализовать интеграционное или E2E-тестирование для остальных сценариев.
- [x] Описать конфигурацию линтера.

## Postman doc
The /signup method creates a user with the role  
The /signin method returns a jwt token for user authorization and authentication in the service  
...  
other: [documenter.getpostman.com](https://documenter.getpostman.com/view/26679053/2sA3Bj7DpC)  