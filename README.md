# go-dictionary

A collection of Go HTTP routing examples demonstrating different web frameworks and patterns. This repository showcases practical implementations of HTTP servers using popular Go routing libraries and techniques.

## 📋 Overview

This project contains several examples of HTTP server implementations and RPC (Remote Procedure Call) services in Go:

1. **Gorilla Mux Router** - Example using the Gorilla Mux framework
2. **HttpRouter** - Example using the lightweight httprouter package
3. **Stateful API** - A complete RESTful API with in-memory user management
4. **Learning Middlewares** - Example demonstrating HTTP middleware chaining patterns
5. **RPC Server** - Standard Go RPC server providing time service
6. **RPC Client** - Client for connecting to the RPC server
7. **Gorilla RPC Server** - JSON-RPC server for book details lookup

These examples are designed for learning and reference, demonstrating best practices for building HTTP servers and RPC services in Go.

## 📂 Project Structure

```
.
├── learningMiddlewares/
│   └── learningMiddlewares.go  # HTTP middleware chaining example
├── otherMux/
│   ├── gorillaMux.go           # Example using Gorilla Mux router
│   └── httpRouter.go           # Example using HttpRouter
├── testingStatefulApi/
│   └── statefulApi.go          # Stateful REST API with user CRUD operations
├── rpcServer/
│   └── rpcServer.go            # Standard Go RPC server (time service)
├── rpcClient/
│   └── rpcClient.go            # RPC client with server/client modes
└── gorillaRPCServer/
    └── gorillarpcserver.go     # JSON-RPC server using Gorilla RPC
```

## 🚀 Features

### Learning Middlewares Example (`learningMiddlewares/learningMiddlewares.go`)
- HTTP middleware chaining demonstration
- Content-Type validation middleware
- Server timestamp cookie middleware
- RESTful city management API
- Thread-safe operations with mutex
- JSON request/response handling

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

### RPC Server Example (`rpcServer/rpcServer.go`)
- Standard Go RPC server implementation
- TCP-based communication on port 1234
- Time service providing Unix timestamps
- HTTP RPC protocol support
- Simple RPC method registration

### RPC Client Example (`rpcClient/rpcClient.go`)
- RPC client for connecting to the RPC server
- Supports both server and client modes
- Command-line argument handling
- Integration with Air for live reloading
- Remote procedure call demonstration

### Gorilla RPC Server Example (`gorillaRPCServer/gorillarpcserver.go`)
- JSON-RPC server using Gorilla RPC framework
- HTTP/JSON-based communication
- Book details lookup service
- File-based data storage (book.json)
- RESTful-style RPC endpoint

## 📦 Prerequisites

- Go 1.16 or higher
- Required dependencies:
  - `github.com/gorilla/mux` - For Gorilla Mux example
  - `github.com/julienschmidt/httprouter` - For HttpRouter example
  - `github.com/gorilla/rpc` - For Gorilla RPC server example
  - `github.com/gorilla/handlers` - For HTTP handlers
  - `github.com/justinas/alice` - For middleware chaining

## 🔧 Installation

1. Clone the repository:
```bash
git clone https://github.com/Dav16Akin/go-dictionary.git
cd go-dictionary
```

2. Install dependencies:
```bash
go mod download
```

Or install dependencies individually:
```bash
go get github.com/gorilla/mux
go get github.com/julienschmidt/httprouter
go get github.com/gorilla/rpc
```

## 💻 Usage

### Running the Learning Middlewares Example

The Learning Middlewares example demonstrates how to chain multiple HTTP middlewares together:

```bash
cd learningMiddlewares
go run learningMiddlewares.go
```

Server will start on `http://localhost:8080`

**API Endpoints:**

- `POST /city` - Create a new city

**Example Request:**

Create a city:
```bash
curl -X POST http://localhost:8080/city \
  -H "Content-Type: application/json" \
  -d '{"name":"New York","area":783800000}'
```

**Middleware Features:**
- **Content-Type Middleware**: Validates that requests have `Content-Type: application/json` header
- **Server Time Middleware**: Adds a cookie with the current server timestamp (UTC)

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

**Note:** Before running, you need to update the static file path in `otherMux/httpRouter.go` (line 44):
```go
router.ServeFiles("/static/*filepath", http.Dir("/Users/apple/Documents/static"))
```
Change this path to a directory on your local machine, or comment out this line if you don't need static file serving.

```bash
cd otherMux
go run httpRouter.go
```

Server will start on `http://localhost:8000`

**Example Routes:**
- `GET /api/v1/go-version` - Returns Go version
- `GET /api/v1/show-file` - Returns file content
- `GET /static/*filepath` - Serves static files

### Running the RPC Server and Client

The RPC server provides a time service using standard Go RPC over TCP:

**Start the RPC Server:**
```bash
cd rpcServer
go run rpcServer.go
```

Server will start on `tcp://localhost:1234`

**Run the RPC Client:**

The `rpcClient` package provides both server and client modes:

```bash
# Run in server mode
go run rpcClient/rpcClient.go server

# Run in client mode (requires server to be running)
go run rpcClient/rpcClient.go client
```

The client will connect to the RPC server and retrieve the current Unix timestamp.

### Running the Gorilla RPC Server

The Gorilla RPC server provides a JSON-RPC service for book lookups:

**Note:** Before running, you need to create a `book.json` file in the project root with book data:

```json
[
  {
    "id": "1",
    "name": "The Go Programming Language",
    "author": "Alan A. A. Donovan"
  },
  {
    "id": "2",
    "name": "Learning Go",
    "author": "Jon Bodner"
  }
]
```

```bash
cd gorillaRPCServer
go run gorillarpcserver.go
```

Server will start on `http://localhost:1234`

**Example JSON-RPC Request:**

```bash
curl -X POST http://localhost:1234/rpc \
  -H "Content-Type: application/json" \
  -d '{"method":"JSONServer.GiveBookDetail","params":[{"Id":"1"}],"id":1}'
```

## 📚 Learning Resources

These examples demonstrate:
- HTTP server setup and configuration
- Route handling and parameter extraction
- RESTful API design
- Thread-safe concurrent operations
- JSON encoding/decoding
- Error handling in HTTP handlers
- Custom timeouts and server configuration
- RPC (Remote Procedure Call) implementations
- Standard Go RPC over TCP
- JSON-RPC with Gorilla RPC
- Client-server communication patterns

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
- [Gorilla RPC Documentation](https://github.com/gorilla/rpc)
- [Go HTTP Server Documentation](https://pkg.go.dev/net/http)
- [Go RPC Package Documentation](https://pkg.go.dev/net/rpc)

---

Feel free to open issues or pull requests for improvements, bug fixes, or new examples!
