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

# Routes:

GET /characters to see a list of all characters stored in the database. POST /characters will make a new character with JSON from the request body.

GET /characters/:id to see a specific character's information. PUT /characters/:id to update that character with JSON, DELETE /characters/:id to delete them from the storage.

PUT /characters/:id/level-up will level up the character specified.

GET /races and /races/:id to see information about races.

GET /classes and /classes/:id to see information about races.


# Database:

MySQL server running locally.
`dungeons_and_dragons` for production database.
`dungeons_and_dragons_test` for testing database.

## Tables

### Key:
- PK - Primary Key
- FK (related column) - Foreign Key
- AI - Auto Increment
- NN - Not Null

### Tables

| Table Name    | classes  |          |
|---------------|----------|----------|
| Column Name   | Type     | Notes    |
| id            | int      | PK, AI   |
| name          | string   | NN       |
| description   | string   |          |

| Table Name    | races    |          |
|---------------|----------|----------|
| Column Name   | Type     | Notes    |
| id            | int      | PK, AI   |
| name          | string   | NN       |

| Table Name    | characters |                     |
|---------------|-----------|----------------------|
| Column Name   | Type      | Notes                |
| id            | int       | PK, AI               |
| name          | string    | NN                   |
| level         | int       | NN, Default: 1       |
| class_id      | int       | NN, FK (classes.id)  |
| race_id       | int       | NN, FK (races.id)    |

