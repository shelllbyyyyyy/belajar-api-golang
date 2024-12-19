# Learn API Development with Go and Fiber

This repository is a learning project to understand how to build RESTful APIs using [Fiber](https://gofiber.io), a high-performance web framework for Go.

## Features

- Basic CRUD operations
- Request validation
- Middleware implementation
- Structuring a Go project for scalability
- Integration with a database (PostgreSQL)
- API documentation using Swagger

---

## Prerequisites

- Go 1.20 or later
- A PostgreSQL database
- `go` package manager installed (`go mod`)

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/shelllbyyyyyy/learn-fiber-api.git
cd learn-fiber-api
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Set Up Environment Variables

Create a .env file in the project root with the following variables:

```bash
PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
```

### 4. Run the Application

```bash
go run main.go
```

The API server will be running at http://localhost:3000.

## Project Structure

```arduino
.
├── apps                    // Apps module such as auth, products, payment
├── cmd
│   └── first_api
│       └── main.go         // Entry point
├── common                  // Common middleware, error, & response
├── configs                 // App configurations (database, environment)
├── docs                    // Swagger documentation
├── scripts                 // Migration
└── util                    // Utility

```

## API Endpoints

### Health Check

### GET /health

Response:

```json
{
  "success": "true",
  "message": "API is running"
}
```

## CRUD Example: Users

### GET /users - Get all users

Response:

```json
{
  "success": true,
  "message": "User found",
  "data": {
    "id": "test-12h9e8",
    "username": "test",
    "email": "test@email.com",
    "created_at": "2024-12-19T09:11:50.674361Z",
    "updated_at": "2024-12-19T09:11:50.674361Z"
  }
}
```

### POST /users - Create a new user

Request Body:

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com"
}
```

Response:

```json
{
  "success": true,
  "message": "Register successfully"
}
```

### GET /users/:id - Get user by ID

Response:

```json
{
  "success": true,
  "message": "User found",
  "data": {
    "id": "test-12h9e8",
    "username": "test",
    "email": "test@email.com",
    "created_at": "2024-12-19T09:11:50.674361Z",
    "updated_at": "2024-12-19T09:11:50.674361Z"
  }
}
```

### PATCH /users - Update user base on their token

Request Body:

```json
{
  "username": "John Doe",
  "email": "john.doe@newexample.com"
}
```

Response:

```json
{
  "success": true,
  "message": "User updated",
  "data": true
}
```

### DELETE /users - Delete user base on their token

Response:

```json
{
  "success": true,
  "message": "User deleted",
  "data": true
}
```

## Features

- Middleware
- Logger: Logs HTTP requests
- Error Handling: Global error handling
- CORS: Enable cross-origin requests
- Database Integration

This project uses PostgreSQL.
Make sure to update the .env file with your database credentials. Use the following command to run migrations:

```bash
go run database/migrate.go
```

## API Documentation

This project uses Swagger for API documentation. Access the docs at http://localhost:3000/swagger/index.html after starting the server.

Learning Resources
Fiber Documentation
Go by Example
PostgreSQL Documentation
Contributing
Feel free to submit issues or pull requests to improve this learning repository.

License
This project is licensed under the MIT License.

vbnet

This `README.md` provides a clear overview of the project, including setup instructions, project s
