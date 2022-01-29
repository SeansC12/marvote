# Marvote API

This is a backend service providing endpoints to serve the Frontend module of Marvote.

## List of endpoints

Base endpoint:


`/api/v1/characters` - This is the base endpoint for any actions related to creating and updating Marvel Characters into the database.

| HTTP Action | Endpoint    | Description|
|-------------|-------------|------------|
| `GET`       | `/api/v1/characters/all` | Retrieves all characters available in the database
| `GET`       | `/api/v1/characters/${id}` | Retrieve one character given the `id` associated to it.
| `POST`      | `/api/v1/characters/` | Add a new character into the database


## Responses