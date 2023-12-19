# AKASIA

Tech Stack:
- Golang 1.19
- Echo 3.3.10
- MySql

## Architecture
```
├── Config
|   ├── databaseEngine.go
|   ├── engine.go
|   ├── interface.go
|   └── model.go
├── Controller
|   ├── Dto
|   |   ├── Request
|   |   |   └── product.go
|   |   └── Response
|   |   |   ├── product.go
|   |   |   └── response.go
|   ├── controllerProduct.go
|   └── interface.go
├── Environment
|   └── Local.yml
├── Repository
|   ├── Product
|   |   ├── model.go
|   |   └── product.go
|   └── interface.go
├── Routes
|   └── route.go
├── Service
|   ├── Migration
|   |   └── 20231219database_migration.sql
|   └── service.go
├── .gitignore
├── go.mod
|   └── go.sum
├── main.go
└── README.md

```

## How to Run
- [Scripts](#scripts)
    - [Terminal](#terminal)

### Scripts
Script go running on terminal, before that you must direct to the path project after clone this project from GIT

### Terminal
```bash
# Running project
go run main.go
```
