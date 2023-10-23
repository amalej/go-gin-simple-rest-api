# go-gin-simple-rest-api

## Routes

[GET] localhost:8080/items

- Get a list of items. `curl localhost:8080/items`

[POST] localhost:8080/items

- Create a new item. `curl localhost:8080/items --request "POST" -d @postPayload.json --header "Content-Type: application/json"`

[GET] localhost:8080/items/:id

- Get a specific item. `curl localhost:8080/items/[id]`

[GET] localhost:8080/items/search?type=[type]

- Get a list of items of a specific type. `curl localhost:8080/items/search?type=[type]`
