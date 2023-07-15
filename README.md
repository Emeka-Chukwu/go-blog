# go-blog


## How to run this project

## System Requirements

- Golang
- Docker
- Postgres (included in docker)
- Migrations

## Running

Setting up all containers

```console
$ make postgress
```


## Migration

Migrating sql file into db

```console
$ make migrationup
```

Dropping tables 

```console
$ make migrationdown
```
Note: Run migration when docker is running

## Run Test

Run all the test in the project with

```console
$ make test
```


## Sections 

- ### Users
- ### category
- ### Tag
- ### Post
- ### comment


## Endpoints
The endpoints for this blog api

### Authentication
register

*/api/v1/auths/signup*
#### request
{
    "username": user.Username,
    "password": password,
    "role":     user.Role,
    "email":    user.Email
}

register

*/api/v1/auths/login*
#### request
{
    "password": password,
    "email":    user.Email
}

```console

```

### Category
create

*/api/v1/category/create*
#### request
{
    "name": name   }

fetch by id

*/api/v1/category/:id*

### Completed Tasks

### Tasks
- [x] Implement api for comments
- [x] write the test coverage for it
- [x] complete the api documentation
- [o] complete the api documentation
