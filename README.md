# 🛒 ECom_API – Go E-Commerce Backend

A production-style **E-commerce REST API** built with **Go**, **MySQL**, **JWT authentication**, and **Docker**.  
This project demonstrates clean backend architecture, authentication, and containerized deployment.

---

## 🚀 Features

- User registration & login
- JWT-based authentication (HS256)
- Protected routes with middleware
- Modular service-based architecture
- MySQL database integration
- Docker & Docker Compose setup
- Unit tests for core services

---

## 🧱 Tech Stack

- **Language:** Go
- **Router:** Gorilla Mux
- **Database:** MySQL
- **Authentication:** JWT
- **Password Hashing:** bcrypt
- **Containerization:** Docker, Docker Compose

---

## 🐳 Run with Docker (Recommended)

### 1️⃣ Clone the repository
```bash
git clone https://github.com/Tejasvi1206/ECom_API.git
cd ECom_API

2️⃣ Create .env file
APP_PORT=8080

DB_HOST=mysql
DB_PORT=3306
DB_USER=ecom_user
DB_PASSWORD=ecom_pass
DB_NAME=ecom

JWT_SECRET=a-string-secret-at-least-256-bits-long

3️⃣ Start the application
docker compose up --build


📍 API will be available at:

http://localhost:8080

🔐 Authentication Flow
Register User
POST /api/v1/register

{
  "firstName": "Tejasvi",
  "lastName": "User",
  "email": "tejasvi@example.com",
  "password": "password123"
}

Login
POST /api/v1/login

{
  "email": "tejasvi@example.com",
  "password": "password123"
}


Response:

{
  "token": "<JWT_TOKEN>"
}

🔒 Protected Routes
Get User by ID
GET /api/v1/users/{id}


Headers:

Authorization: Bearer <JWT_TOKEN>

🧪 Run Locally (Without Docker)
go mod tidy
go build ./...
go run cmd/main.go

📌 Status
Dockerized and working
JWT authentication verified
Clean build (go build ./...)
Ready for deployment and extension
```

## 🔮 Future Improvements
- Role-based access control
- Pagination & filtering
- Order checkout flow
- Redis caching
- CI/CD pipeline

## 👨‍💻 Author

- Tejasvi
- Backend / Go Developer

---