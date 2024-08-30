# D&D Character Sheet API

```shell
#  Build: 
go build;

#  Test:
go test ./...

#  Run:
go run dungeons-and-dragons;

#  Visit
http://localhost:8080
```

### Routes:

GET /characters to see a list of all characters stored in the database. POST /characters will make a new character with JSON from the request body.

GET /characters/:id to see a specific character's information. PUT /characters/:id to update that character with JSON, DELETE /characters/:id to delete them from the storage.

PUT /characters/:id/level-up will level up the character specified.

GET /races and /races/:id to see information about races.

GET /classes and /classes/:id to see information about races.

