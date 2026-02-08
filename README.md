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
- Transaction processing & sales reporting

## 🏗️ Project Structure

```
kasir-api/
├── main.go                          # Application entry point
├── database/
│   └── database.go                  # Database connection setup
├── models/
│   ├── product.go                   # Product model
│   ├── category.go                  # Category model
│   └── transaction.go               # Transaction & report models
├── repositories/
│   ├── product_repository.go        # Product data access layer
│   ├── category_repository.go       # Category data access layer
│   └── transaction_repository.go    # Transaction data access layer
├── services/
│   ├── product_service.go           # Product business logic
│   ├── category_service.go          # Category business logic
│   └── transaction_service.go       # Transaction business logic
├── handlers/
│   ├── product_handler.go           # Product HTTP handlers
│   ├── category_handler.go          # Category HTTP handlers
│   └── transaction_handler.go       # Transaction HTTP handlers
├── docs/
│   └── swagger.json                 # OpenAPI 3.0 Swagger documentation
├── go.mod
├── go.sum
└── .env                             # Environment variables
```

## 🛠️ Tech Stack

- **Go** 1.24
- **PostgreSQL** - Relational database
- **pgx** - PostgreSQL driver for Go
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

-- Transaction table
CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    total_amount INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Transaction detail table
CREATE TABLE transaction_detail (
    id SERIAL PRIMARY KEY,
    transaction_id INTEGER REFERENCES transaction(id),
    product_id INTEGER REFERENCES product(id),
    product_name VARCHAR(255),
    quantity INTEGER NOT NULL,
    subtotal INTEGER NOT NULL
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
| GET | `/` | API information & endpoint list |
| GET | `/health` | Health check |

### Products
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/produk` | Get all products |
| GET | `/api/produk?name={keyword}` | Search products by name |
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

### Transactions
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/checkout` | Process checkout (multiple items) |

### Reports
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/report/hari-ini` | Today's sales summary |
| GET | `/api/report?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD` | Sales report by date range |

## 📖 API Documentation (Swagger)

Full OpenAPI 3.0 documentation is available at [`docs/swagger.json`](docs/swagger.json).

You can view it interactively using:
- [Swagger Editor](https://editor.swagger.io/) — paste or import the JSON file
- [Swagger UI](https://petstore.swagger.io/) — point to the raw URL of your `swagger.json`

## 📝 Example Requests

### Create Product
```bash
curl -X POST http://localhost:8080/api/produk \
  -H "Content-Type: application/json" \
  -d '{"name": "Nasi Goreng", "price": 15000, "stock": 100}'
```

### Get All Products
```bash
curl http://localhost:8080/api/produk
```

### Search Products by Name
```bash
curl "http://localhost:8080/api/produk?name=nasi"
```

### Create Category
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Makanan", "description": "Kategori produk makanan"}'
```

### Get All Categories
```bash
curl http://localhost:8080/api/categories
```

### Checkout (Create Transaction)
```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -d '{"items": [{"product_id": 1, "quantity": 2}, {"product_id": 2, "quantity": 1}]}'
```

### Today's Sales Report
```bash
curl http://localhost:8080/api/report/hari-ini
```

### Sales Report by Date Range
```bash
curl "http://localhost:8080/api/report?start_date=2026-01-01&end_date=2026-02-08"
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
- [x] Transaction / Checkout system
- [x] Sales report (today & date range)
- [x] Product search by name
- [x] API documentation (Swagger / OpenAPI 3.0)
- [ ] Input validation
- [ ] Error handling middleware
- [ ] Authentication & Authorization
- [ ] Unit tests
- [ ] Docker support
- [ ] CI/CD pipeline

## 🌐 Live Demo

Production API: [https://kasir-api-production-ecd5.up.railway.app](https://kasir-api-production-ecd5.up.railway.app)

## 📖 Key Learnings

- Using Go's standard `net/http` library without external frameworks
- Manual dependency injection pattern (Repository → Service → Handler)
- Configuration management with Viper for `.env` and environment variables
- PostgreSQL integration with `pgx` driver
- Transaction processing with stock management
- Sales reporting with date filtering

## 📄 License

This project is created for learning purposes.
