# rest_api_project

## Introduction

`rest_api_project` is a simple yet powerful RESTful API developed in Go using the Gin web framework. It serves as an events manager, allowing users to perform various operations related to event management. The API uses SQLite for data storage, providing a lightweight and efficient database solution. It incorporates JWT (JSON Web Tokens) for securing certain endpoints, ensuring that only authenticated users can access specific functionalities.

The server will start on `localhost:8080`. You can access the API through `http://localhost:8080`.

## API Endpoints

### Health Check

- `GET /` - Welcome message
- `GET /api/health` - Check API health

### Event Routes

- `GET /events` - Retrieve all events
- `GET /events/:id` - Retrieve a single event by ID

### Authenticated Event Routes

These routes require a JWT token for access.

- `POST /events` - Create a new event
- `PUT /events/:id` - Update an existing event by ID
- `DELETE /events/:id` - Delete an event by ID
- `POST /events/:id/register` - Register for an event by ID
- `DELETE /events/:id/register` - Unregister from an event by ID

### User Routes

- `POST /signup` - Register a new user
- `POST /login` - Authenticate a user and return a JWT token

## Authentication

This API uses JWT tokens for securing certain routes. To access protected routes, a valid JWT token must be included in the `Authorization` header of the request.

## Testing endpoints

There is a api-test directory that contains files for testing endpoints usingg the vs-code REST client extension. The files are named according to the endpoint they are testing. The extension information is as follows:

Name: REST Client
Id: humao.rest-client
Description: REST Client for Visual Studio Code
Version: 0.25.1
Publisher: Huachao Mao
VS Marketplace Link: https://marketplace.visualstudio.com/items?itemName=humao.rest-client

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
