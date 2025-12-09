# go-dictionary

A collection of Go HTTP routing examples demonstrating different web frameworks and patterns. This repository showcases practical implementations of HTTP servers using popular Go routing libraries and techniques.

## 📋 Overview

This project contains several examples of HTTP server implementations in Go:

1. **Gorilla Mux Router** - Example using the Gorilla Mux framework
2. **HttpRouter** - Example using the lightweight httprouter package
3. **Stateful API** - A complete RESTful API with in-memory user management

These examples are designed for learning and reference, demonstrating best practices for building HTTP servers in Go.

## 📂 Project Structure

```
.
├── otherMux/
│   ├── gorillaMux.go      # Example using Gorilla Mux router
│   └── httpRouter.go      # Example using HttpRouter
└── testingStatefulApi/
    └── statefulApi.go     # Stateful REST API with user CRUD operations
```

## 🚀 Features

### Gorilla Mux Example (`otherMux/gorillaMux.go`)
- Route parameters extraction
- Pattern matching with regex
- Clean route handling
- Server with custom timeouts

### HttpRouter Example (`otherMux/httpRouter.go`)
- Lightweight and fast routing
- Static file serving
- Command execution endpoint
- File reading endpoint

### Stateful API Example (`testingStatefulApi/statefulApi.go`)
- RESTful user management API
- In-memory data storage
- Thread-safe operations with mutex
- Full CRUD operations (Create, Read, Update, Delete)
- JSON request/response handling

## 📦 Prerequisites

- Go 1.16 or higher
- Required dependencies:
  - `github.com/gorilla/mux` - For Gorilla Mux example
  - `github.com/julienschmidt/httprouter` - For HttpRouter example

## 🔧 Installation

1. Clone the repository:
```bash
git clone https://github.com/Dav16Akin/go-dictionary.git
cd go-dictionary
```

2. Install dependencies:
```bash
go get github.com/gorilla/mux
go get github.com/julienschmidt/httprouter
```

Or use Go modules (recommended):
```bash
go mod init github.com/Dav16Akin/go-dictionary
go mod tidy
```

## 💻 Usage

### Running the Stateful API Example

The Stateful API provides a complete user management system:

```bash
cd testingStatefulApi
go run statefulApi.go
```

Server will start on `http://localhost:8080`

**API Endpoints:**

- `GET /users` - List all users
- `POST /users` - Create a new user
- `GET /users/{id}` - Get a specific user
- `PUT /users/{id}` - Update a user
- `DELETE /users/{id}` - Delete a user

**Example Requests:**

Create a user:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

Get all users:
```bash
curl http://localhost:8080/users
```

Get a specific user:
```bash
curl http://localhost:8080/users/1
```

Update a user:
```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","email":"jane@example.com"}'
```

Delete a user:
```bash
curl -X DELETE http://localhost:8080/users/1
```

### Running the Gorilla Mux Example

```bash
cd otherMux
go run gorillaMux.go
```

Server will start on `http://localhost:8000`

**Example Routes:**
- `GET /` - Home endpoint
- `GET /articles/{category}/{id}` - Article endpoint with path parameters

### Running the HttpRouter Example

Note: Update the static file path in the code to match your local setup.

```bash
cd otherMux
go run httpRouter.go
```

Server will start on `http://localhost:8000`

**Example Routes:**
- `GET /api/v1/go-version` - Returns Go version
- `GET /api/v1/show-file` - Returns file content
- `GET /static/*filepath` - Serves static files

## 📚 Learning Resources

These examples demonstrate:
- HTTP server setup and configuration
- Route handling and parameter extraction
- RESTful API design
- Thread-safe concurrent operations
- JSON encoding/decoding
- Error handling in HTTP handlers
- Custom timeouts and server configuration

## 🤝 Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a branch for your feature: `git checkout -b feature/new-example`
3. Make your changes following Go best practices
4. Test your code thoroughly
5. Open a pull request with a clear description

## 📝 License

This project is open source and available under the MIT License.

## 👤 Author

Created by **Dav16Akin**

- GitHub: [@Dav16Akin](https://github.com/Dav16Akin)
- Repository: [go-dictionary](https://github.com/Dav16Akin/go-dictionary)

## 🔗 Related Resources

- [Gorilla Mux Documentation](https://github.com/gorilla/mux)
- [HttpRouter Documentation](https://github.com/julienschmidt/httprouter)
- [Go HTTP Server Documentation](https://pkg.go.dev/net/http)

---

Feel free to open issues or pull requests for improvements, bug fixes, or new examples!
