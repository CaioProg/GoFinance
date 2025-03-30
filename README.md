# GoFinance API 🚧 (Under Construction)

GoFinance is a RESTful API built with Go and Fiber to manage financial transactions, including users, expenses, incomes, and categories.

## 🚀 Features
- User registration and authentication
- Expense management (CRUD operations)
- Income management (CRUD operations)
- Category management (CRUD operations)
- Database persistence using GORM
- Docker support

## 🛠 Installation
### Prerequisites
- Go 1.20+
- Docker & Docker Compose
- PostgreSQL (if not using Docker)

### Clone the repository
```sh
git clone https://github.com/CaioProg/GoFinance.git
cd GoFinance
```

### Run with Docker
```sh
docker-compose up --build
```

### Run locally
1. Install dependencies:
```sh
go mod tidy
```
2. Start the application:
```sh
go run cmd/main.go
```

## 📌 API Endpoints
### User Routes
| Method | Endpoint     | Description       |
|--------|--------------|-------------------|
| POST   | /user        | Create a new user |
| GET    | /user/:id    | Get user by Id    |
| GET    | /users       | Get all users     |
| PATCH  | /user/:id    | Update user       |
| DELETE | /user/:id    | Delete user       |

### Expense Routes
| Method | Endpoint       | Description |
|--------|-------------|-------------|
| POST   | /expense    | Create an expense |
| GET    | /expense/:id| Get expense by ID |
| GET    | /expense/:id| Get all expenses|
| PATCH  | /expense/:id| Update expense |
| DELETE | /expense/:id| Delete expense |

## 🏗 Project Structure
```
GoFinance/
│── cmd/
│   ├── main.go     # Main application entry point
│── internal/
│   ├── db/            # Database connection and migrations
│   ├── handlers/      # Request handlers
│   ├── models/        # Data models
│   ├── repositories/  # Database interactions
│   ├── services/      # Business logic
│   ├── routes/        # API routes
│── docker-compose.yml # Docker configuration
