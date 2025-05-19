# Go JWT Authentication & CRUD API

## Overview

This project implements a RESTful API using **Go (Golang)**, **Gin framework**, **GORM**, and **PostgreSQL**, with **JWT authentication**.

---

## 🚀 Setup Requirements

- Go 1.16+
- PostgreSQL
- Git

---

## ⚙️ Environment Configuration

Create a `.env` file in the root directory with the following content:

```env
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=gin_auth_crud
DB_PORT=5432
PORT=3000
JWT_SECRET=your_jwt_secret
```

---

## 🗄️ Database Migration

Run the migration to create necessary tables:

```bash
go run db/migrate/migrate.go
```

---

## 🔐 Authentication Endpoints

| Method | Endpoint    | Description       |
| ------ | ----------- | ----------------- |
| POST   | /api/signup | Register new user |
| POST   | /api/login  | Login user        |
| POST   | /api/logout | Logout user       |

---

## 📂 Category Endpoints

| Method | Endpoint                             | Description             |
| ------ | ------------------------------------ | ----------------------- |
| POST   | /api/categories/create               | Create category         |
| GET    | /api/categories                      | List all categories     |
| GET    | /api/categories/:id/edit             | Get category for edit   |
| PUT    | /api/categories/:id/update           | Update category         |
| DELETE | /api/categories/:id/delete           | Soft delete category    |
| GET    | /api/categories/all-trash            | List trashed categories |
| DELETE | /api/categories/delete-permanent/:id | Permanent delete        |

---

## 📝 Post Endpoints

| Method | Endpoint                        | Description        |
| ------ | ------------------------------- | ------------------ |
| POST   | /api/posts/create               | Create post        |
| GET    | /api/posts                      | List all posts     |
| GET    | /api/posts/:id/show             | Show single post   |
| GET    | /api/posts/:id/edit             | Get post for edit  |
| PUT    | /api/posts/:id/update           | Update post        |
| DELETE | /api/posts/:id/delete           | Soft delete post   |
| GET    | /api/posts/all-trash            | List trashed posts |
| DELETE | /api/posts/delete-permanent/:id | Permanent delete   |

---

## 💬 Comment Endpoints

| Method | Endpoint                          | Description          |
| ------ | --------------------------------- | -------------------- |
| POST   | /api/posts/:id/comment/store      | Create comment       |
| GET    | /api/posts/:id/comment/:id/edit   | Get comment for edit |
| PUT    | /api/posts/:id/comment/:id/update | Update comment       |
| DELETE | /api/posts/:id/comment/:id/delete | Delete comment       |

---

## 📦 Request Examples

### User Registration

```json
POST /api/signup
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "123456"
}
```

### User Login

```json
POST /api/login
{
  "email": "john@example.com",
  "password": "123456"
}
```

### Create Category

```json
POST /api/categories/create
{
  "name": "Technology"
}
```

### Create Post

```json
POST /api/posts/create
{
  "title": "My First Post",
  "body": "Post content goes here",
  "categoryId": 1
}
```

### Create Comment

```json
POST /api/posts/1/comment/store
{
  "postId": 1,
  "body": "Great post!"
}
```

---

## 🔐 Authentication Notes

All endpoints **except** `/api/signup` and `/api/login` require authentication.
Use the **JWT token** received from login in the **Cookie** header:

```http
Cookie: Authorization=<token>
```

---

## ⚠️ Error Handling

The API returns appropriate HTTP status codes:

- `200`: Success
- `201`: Created
- `400`: Bad Request
- `401`: Unauthorized
- `404`: Not Found
- `422`: Unprocessable Entity
- `500`: Internal Server Error

---

## 🛠 Development

```bash
# Run the application
go run main.go

# Build the application
go build -o app

# Run tests
go test ./...
```

---

This documentation provides a comprehensive guide for developers to understand and use the API.
