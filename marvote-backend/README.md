# Marvote API

This is a backend service providing endpoints to serve the Frontend module of Marvote.  MongoDB is used as the backend datastore of this API.


## Build and Test

To build it in linux

```shell

go build .

```

To run test

```
go test -v -p=1 -coverpkg=./... -coverprofile=coverage.txt  ./...
go tool cover -html=coverage.txt
```

Starting the application

```shell
./marvote serve --config config/config.yaml
```

## TODO

- [X] Proper logging framework, using zap
- [ ] `UPDATE` a character
- [X] `DELETE` a character
- [X] Vote a character
- [X] Tally the votes against the characters, based on the get all endpoint

## List of endpoints

Character endpoints

| HTTP Action | Endpoint    | Description|
|-------------|-------------|------------|
| `GET`       | `/api/v1/characters/all` | Retrieves all Marvel characters available in the database |
| `GET`       | `/api/v1/character/:id` | Retrieve one Marvel character given the `id` associated to it. |
| `POST`      | `/api/v1/character/` | Add a new Marvel character into the database |
| `PUT`       | `/api/v1/character/:id` | Update a character given the `id`|
| `DELETE`    | `/api/v1/character/:id` | Delete a character given the `id`|
