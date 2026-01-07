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
8. **Go-Restful Fundamentals** - REST API example using the go-restful framework
9. **SQLite Fundamentals** - Database operations example with SQLite demonstrating CRUD operations
10. **Gin Fundamentals** - REST API example using the Gin web framework with SQLite database
11. **Rail API** - Comprehensive railway management REST API using go-restful framework with database-backed train, station, and schedule management

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
├── gorillaRPCServer/
│   └── gorillarpcserver.go     # JSON-RPC server using Gorilla RPC
├── goRestfulFundamemtals/
│   └── gorestfulfundamentals.go # REST API using go-restful framework
├── sqliteFundamentals/
│   └── sqliteFundamentals.go   # SQLite database CRUD operations example
├── ginFundamentals/
│   └── ginFundamentals.go      # REST API using Gin web framework with SQLite
├── railAPI/
│   ├── railAPI.go              # Railway management REST API with go-restful
│   └── dbUtils/
│       ├── init-tables.go      # Database table initialization
│       └── models.go           # Database schema models
└── .air.toml                    # Air live reload configuration
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

### Go-Restful Fundamentals Example (`goRestfulFundamemtals/gorestfulfundamentals.go`)
- REST API using the go-restful framework
- Simple ping time service
- WebService and Route registration
- HTTP GET endpoint demonstration
- Lightweight REST API pattern

### SQLite Fundamentals Example (`sqliteFundamentals/sqliteFundamentals.go`)
- SQLite database integration with Go
- Complete CRUD operations (Create, Read, Update, Delete)
- Table creation and management
- Prepared statements for SQL queries
- Database connection handling
- Book inventory management example

### Gin Fundamentals Example (`ginFundamentals/ginFundamentals.go`)
- REST API using the Gin web framework
- SQLite database integration
- Railway station management API
- Full CRUD operations for stations
- JSON request/response handling with Gin
- Database-backed persistent storage
- Error handling and validation
- HTTP status code management

### Rail API Example (`railAPI/railAPI.go`)
- Comprehensive railway management system
- REST API using go-restful framework
- SQLite database with multiple related tables
- Train management with driver information
- Station management with operating hours
- Schedule management for train-station relationships
- Database utilities for table initialization
- Modular database schema design
- Foreign key relationships between entities

## 📦 Prerequisites

- Go 1.16 or higher
- Required dependencies:
  - `github.com/gorilla/mux` - For Gorilla Mux example
  - `github.com/julienschmidt/httprouter` - For HttpRouter example
  - `github.com/gorilla/rpc` - For Gorilla RPC server example
  - `github.com/gorilla/handlers` - For HTTP handlers
  - `github.com/justinas/alice` - For middleware chaining
  - `github.com/emicklei/go-restful` - For go-restful fundamentals example
  - `github.com/mattn/go-sqlite3` - For SQLite database example
  - `github.com/gin-gonic/gin` - For Gin web framework example
- Optional tools:
  - [Air](https://github.com/air-verse/air) - Live reload for Go apps (configured via `.air.toml`)

## 🔧 Installation

1. Clone the repository:
```bash
git clone https://github.com/Dav16Akin/go-dictionary.git
cd go-dictionary
```

2. Install dependencies:
```bash
go mod tidy
```

This will download all required dependencies listed in `go.mod`.

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
cd rpcClient
go run rpcClient.go server

# Run in client mode (requires server to be running)
cd rpcClient
go run rpcClient.go client
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

### Running the Go-Restful Fundamentals Example

The Go-Restful Fundamentals example demonstrates building REST APIs using the go-restful framework:

```bash
cd goRestfulFundamemtals
go run gorestfulfundamentals.go
```

Server will start on `http://localhost:8000`

**API Endpoints:**

- `GET /ping` - Returns the current server time

**Example Request:**

```bash
curl http://localhost:8000/ping
```

This will return the current server timestamp.

### Running the SQLite Fundamentals Example

The SQLite Fundamentals example demonstrates database operations with SQLite:

```bash
cd sqliteFundamentals
go run sqliteFundamentals.go
```

This example demonstrates:
- Creating a database and table (`books.db`)
- Inserting records into the database
- Reading all records from the database
- Updating existing records
- Deleting records

The example will create a `books.db` file in the project root directory and perform all CRUD operations, logging each step to the console.

**Note:** The program performs automated CRUD operations and doesn't expose HTTP endpoints. Check the console output to see the results of each database operation.

### Running the Gin Fundamentals Example

The Gin Fundamentals example demonstrates building a REST API using the Gin web framework with SQLite database:

```bash
cd ginFundamentals
go run ginFundamentals.go
```

Server will start on `http://localhost:8000`

**API Endpoints:**

- `GET /v1/stations` - List all railway stations
- `GET /v1/stations/:station_id` - Get a specific station by ID
- `POST /v1/stations` - Create a new station
- `DELETE /v1/stations/:station_id` - Delete a station

**Example Requests:**

Create a station:
```bash
curl -X POST http://localhost:8000/v1/stations \
  -H "Content-Type: application/json" \
  -d '{"name":"Grand Central","opening_time":"08:00:00","closing_time":"22:00:00"}'
```

Get all stations:
```bash
curl http://localhost:8000/v1/stations
```

Get a specific station:
```bash
curl http://localhost:8000/v1/stations/1
```

Delete a station:
```bash
curl -X DELETE http://localhost:8000/v1/stations/1
```

This example demonstrates:
- Building REST APIs with the Gin framework
- Database integration with SQLite
- CRUD operations with database persistence
- JSON request/response handling
- Error handling and HTTP status codes
- Database connection management

### Running the Rail API Example

The Rail API example demonstrates a comprehensive railway management system using go-restful framework:

```bash
cd railAPI
go run railAPI.go
```

Server will start on `http://localhost:8000`

**API Endpoints:**

- `GET /v1/trains/{train-id}` - Get train details by ID
- `POST /v1/trains` - Create a new train
- `DELETE /v1/trains/{train-id}` - Delete a train

**Example Requests:**

Create a train:
```bash
curl -X POST http://localhost:8000/v1/trains \
  -H "Content-Type: application/json" \
  -d '{"driver_name":"John Smith","operating_status":true}'
```

Get a train:
```bash
curl http://localhost:8000/v1/trains/1
```

Delete a train:
```bash
curl -X DELETE http://localhost:8000/v1/trains/1
```

This example demonstrates:
- Building REST APIs with go-restful framework
- SQLite database with multiple related tables (trains, stations, schedules)
- Database schema design with foreign key relationships
- Modular code organization with database utilities
- Input validation and error handling
- RESTful API design patterns
- Database transaction management

**Note:** The Rail API uses a shared database (`railapi.db`) that includes tables for trains, stations, and schedules. The database is automatically initialized on first run.

### Using Air for Live Reload

This project includes an `.air.toml` configuration file for the [Air](https://github.com/air-verse/air) live reload tool, which automatically rebuilds and restarts your Go application when you make changes.

**Install Air:**

```bash
go install github.com/air-verse/air@latest
```

**Run with Air:**

```bash
# From the project root
air
```

Air will watch for changes in your `.go` files and automatically rebuild and restart your application.

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
- REST API development with go-restful framework
- REST API development with Gin web framework
- Database integration with SQLite
- CRUD operations with SQL databases
- Prepared statements and SQL query execution
- Database schema design and relationships
- Live reload development workflow with Air

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
- [Go-Restful Documentation](https://github.com/emicklei/go-restful)
- [Gin Web Framework Documentation](https://github.com/gin-gonic/gin)
- [Go SQLite3 Driver Documentation](https://github.com/mattn/go-sqlite3)
- [Air Live Reload Tool](https://github.com/air-verse/air)

---

Feel free to open issues or pull requests for improvements, bug fixes, or new examples!
