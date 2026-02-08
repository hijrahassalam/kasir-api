# Kasir API

> 🚧 **Work in Progress** - This project is actively being developed and improved.

A simple POS (Point of Sale) REST API built with Go, implementing clean Layered Architecture pattern. This project serves as a hands-on exploration of Go programming language and backend development best practices.

## 🎯 About This Project

This is a personal learning project to explore and demonstrate proficiency in:
- Go programming language fundamentals
- Building REST APIs using Go's standard `net/http` library
- Implementing Layered Architecture (Handler → Service → Repository)
- PostgreSQL database integration
- Environment-based configuration management

## 🏗️ Project Structure

```
kasir-api/
├── main.go                 # Application entry point
├── database/
│   └── database.go         # Database connection setup
├── models/
│   ├── product.go          # Product model
│   └── category.go         # Category model
├── repositories/
│   ├── product_repository.go   # Product data access layer
│   └── category_repository.go  # Category data access layer
├── services/
│   ├── product_service.go      # Product business logic
│   └── category_service.go     # Category business logic
├── handlers/
│   ├── product_handler.go      # Product HTTP handlers
│   └── category_handler.go     # Category HTTP handlers
├── go.mod
├── go.sum
└── .env                    # Environment variables
```

## 🛠️ Tech Stack

- **Go** 1.24
- **PostgreSQL** - Relational database
- **Viper** - Configuration management
- **net/http** - HTTP server (Go standard library)

## 📋 Prerequisites

- Go 1.24+
- PostgreSQL

## ⚙️ Configuration

Create a `.env` file in the project root:

```env
PORT=8080
DB_CONN=postgres://username:password@localhost:5432/kasir_db?sslmode=disable
```

## 🗄️ Database Setup

Run the following SQL to create the required tables:

```sql
-- Product table
CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    stock INTEGER NOT NULL
);

-- Category table
CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);
```

## 🚀 Getting Started

```bash
# Clone the repository
git clone <repository-url>
cd kasir-api

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

The server will start at `http://localhost:8080`

## 📡 API Endpoints

### General
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | API information |
| GET | `/health` | Health check |

### Products
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/produk` | Get all products |
| POST | `/api/produk` | Create a new product |
| GET | `/api/produk/{id}` | Get product by ID |
| PUT | `/api/produk/{id}` | Update product |
| DELETE | `/api/produk/{id}` | Delete product |

### Categories
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/categories` | Get all categories |
| POST | `/api/categories` | Create a new category |
| GET | `/api/categories/{id}` | Get category by ID |
| PUT | `/api/categories/{id}` | Update category |
| DELETE | `/api/categories/{id}` | Delete category |

## 📝 Example Requests

### Create Product
```bash
curl -X POST http://localhost:8080/api/produk \
  -H "Content-Type: application/json" \
  -d '{"name": "Fried Rice", "price": 15000, "stock": 100}'
```

### Get All Products
```bash
curl http://localhost:8080/api/produk
```

### Create Category
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Food", "description": "Food products category"}'
```

### Get All Categories
```bash
curl http://localhost:8080/api/categories
```

## 📚 Architecture

This project follows the Layered Architecture pattern:

```
┌─────────────────────────────────────┐
│           HTTP Request              │
└─────────────────┬───────────────────┘
                  ▼
┌─────────────────────────────────────┐
│            Handlers                 │  ← HTTP routing & request/response
└─────────────────┬───────────────────┘
                  ▼
┌─────────────────────────────────────┐
│            Services                 │  ← Business logic
└─────────────────┬───────────────────┘
                  ▼
┌─────────────────────────────────────┐
│          Repositories               │  ← Data access (SQL queries)
└─────────────────┬───────────────────┘
                  ▼
┌─────────────────────────────────────┐
│           Database                  │
└─────────────────────────────────────┘
```

## 🗺️ Roadmap

- [x] Basic CRUD for Products
- [x] Basic CRUD for Categories
- [x] Layered Architecture implementation
- [ ] Input validation
- [ ] Error handling middleware
- [ ] Authentication & Authorization
- [ ] Unit tests
- [ ] API documentation (Swagger)
- [ ] Docker support
- [ ] CI/CD pipeline

## 📖 Key Learnings

- Using Go's standard `net/http` library without external frameworks
- Manual dependency injection pattern (Repository → Service → Handler)
- Configuration management with Viper for `.env` and environment variables
- PostgreSQL integration with `lib/pq` driver

## 📄 License

This project is created for learning purposes.
