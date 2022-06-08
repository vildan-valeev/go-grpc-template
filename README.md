# Go clean architecture template

### REST API

Category instance

id: UUID  
name: str  

GET /categories -- list of Categories -- 200, 404, 500  
GET /categories/:id -- Category by id -- 200, 404, 500  
POST /categories/:id -- create Category -- 204, 4xx, Header Location: url  
PUT /categories/:id -- fully update Category -- 204/200, 404, 400, 500  
PATCH /categories/:id -- partially update Category -- 204/200, 404, 400, 500  
DELETE /categories/:id -- delete Category by id -- 204, 404, 400  

---

Item instance

id: UUID  
name: str  
category: fk  

GET /items -- list of Items -- 200, 404, 500  
GET /items/:id -- Item by id -- 200, 404, 500  
POST /items/:id -- create Item -- 204, 4xx, Header Location: url  
PUT /items/:id -- fully update Item -- 204/200, 404, 400, 500  
PATCH /items/:id -- partially update Item -- 204/200, 404, 400, 500  
DELETE /items/:id -- delete Item by id -- 204, 404, 400  

---

Develop

```sh
make help
```
