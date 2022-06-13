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

`protoc \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
app/internal/transport/grpc/proto/models.proto
`
