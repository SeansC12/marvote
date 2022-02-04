# Marvote API

This is a backend service providing endpoints to serve the Frontend module of Marvote.

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
## TODO

- [ ] Proper logging framework, using zap
- [ ] `UPDATE` a record
- [ ] allow voting a character
- [ ] Tally the number of votes against the characters

## List of endpoints

Character endpoints

| HTTP Action | Endpoint    | Description|
|-------------|-------------|------------|
| `GET`       | `/api/v1/characters/all` | Retrieves all Marvel characters available in the database |
| `GET`       | `/api/v1/character/:id` | Retrieve one Marvel character given the `id` associated to it. |
| `POST`      | `/api/v1/character/` | Add a new Marvel character into the database |
