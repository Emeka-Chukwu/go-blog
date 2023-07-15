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

### Tasks
- [x] Users api implementaion
- [x] Users test coverage for the implemented api
- [x] Category api implementaion
- [x] Category test coverage for the implemented api
- [x] Tag api implementaion
- [x] Tag test coverage for the implemented api
- [x] Post api implementaion
- [x] Post test coverage for the implemented api
- [x] Test coverage for all the db calls 
- [ ] Implement api for comments
- [ ] write the test coverage for it
- [ ] complete the api documentation

