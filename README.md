# Go clean architecture template

### gRPC API

Category instance

id: UUID  
name: str  


---

Item instance

id: UUID  
name: str  
category: fk  

URL - 0.0.0.0:8080/api/v1/items/all


---

Develop

```sh
make help
```

`protoc -I app/internal/transport/grpc/proto \
--go_out=plugins=grp:app/internal/transport/grpc/proto \
app/internal/transport/grpc/proto/models.proto
`


1. Импортировать crud_main.proto в postman
2. экспорт переменных окружения
3. запуск локальной бд
4. компилляция
5. запуск