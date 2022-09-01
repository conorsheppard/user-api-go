### Description

Small microservice to manage access to Users.

### Build & Run With Docker Compose

```console
$ make up
```

To bring the service down and remove resources:

```console
$ make down
```

The service currently lets you:
- Add a new User
- Return a paginated list of Users, allowing for filtering by certain criteria (e.g. all Users with the country "UK")
- Modify a User
- Remove a User

**Add a new User**

```console
$ curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"first_name":"john","last_name":"smith","nickname":"smithy","password":"asdfas","email":"js@gmail.com","country":"UK"}' \
    http://localhost:8080/user/create
```

**Return a paginated list of Users**

```console
$ curl 'http://localhost:8080/user/all?country=UK&limit=10&page=1'
```

**Update a User**

```console
$ curl --header "Content-Type: application/json" \
    --request PUT \
    --data '{"first_name":"john","last_name":"smith","nickname":"smithy","email":"js@gmail.com","country":"UK"}' \
    http://localhost:8080/user/update/:id
```

**Delete a User**

```console
$ curl --request DELETE http://localhost:8080/user/delete/:id
```

### Testing
To run tests, execute
```console
$ make test
```

### Storage

This project uses Postgres with [SQLC](https://sqlc.dev/) which compiles SQL to type-safe code.

You write SQL queries and run the following command to generate code that presents type-safe interfaces to those queries:
```console
$ make sqlc
```

### Improvements
With more time, more tests could be added to improve test coverage