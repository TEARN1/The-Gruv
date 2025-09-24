# The Gruv Authentication API

This document describes the authentication system implemented for The Gruv platform.

## Overview

The authentication system provides JWT-based authentication with the following features:
- User registration and login
- JWT token generation and validation 
- Protected route middleware
- User profile management

## API Endpoints

All endpoints are available through the API Gateway at `http://localhost:8080/api/users/` or directly through the User Service at `http://localhost:8081/`.

### Public Endpoints (No Authentication Required)

#### Register User
- **POST** `/register`
- **Description**: Create a new user account
- **Request Body**:
  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```
- **Response** (201 Created):
  ```json
  {
    "message": "User registered successfully",
    "userId": "uuid-here",
    "token": "jwt_token_here"
  }
  ```

#### Login User  
- **POST** `/login`
- **Description**: Authenticate existing user
- **Request Body**:
  ```json
  {
    "username": "your_username", 
    "password": "your_password"
  }
  ```
- **Response** (200 OK):
  ```json
  {
    "message": "Login successful",
    "userId": "uuid-here", 
    "token": "jwt_token_here"
  }
  ```

### Protected Endpoints (Authentication Required)

These endpoints require a valid JWT token in the Authorization header:
`Authorization: Bearer <jwt_token>`

#### Get User Profile
- **GET** `/profile`
- **Description**: Get current user's profile information
- **Headers**: `Authorization: Bearer <jwt_token>`
- **Response** (200 OK):
  ```json
  {
    "id": "uuid-here",
    "username": "your_username"
  }
  ```

### Health Check
- **GET** `/health`
- **Description**: Service health check
- **Response** (200 OK):
  ```json
  {
    "service": "User Service",
    "status": "UP"
  }
  ```

## JWT Token Details

- **Algorithm**: HS256
- **Expiration**: 24 hours
- **Claims**: Contains user ID and username
- **Header Format**: `Authorization: Bearer <token>`

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid input: <validation_error>"
}
```

### 401 Unauthorized
```json
{
  "error": "Invalid username or password"
}
```
or
```json
{
  "error": "Authorization header required"
}
```
or
```json
{
  "error": "Invalid token"
}
```

### 404 Not Found
```json
{
  "error": "User not found"
}
```

### 409 Conflict
```json
{
  "error": "Username already taken"
}
```

### 500 Internal Server Error
```json
{
  "error": "Failed to generate token"
}
```

## Usage Examples

### Register a new user
```bash
curl -X POST http://localhost:8080/api/users/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'
```

### Login
```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}'
```

### Access protected profile endpoint
```bash
curl -X GET http://localhost:8080/api/users/profile \
  -H "Authorization: Bearer <your_jwt_token>"
```

## Security Notes

- Passwords are hashed using bcrypt
- JWT tokens expire after 24 hours
- The JWT secret should be changed in production (currently hardcoded for development)
- All sensitive endpoints require valid JWT authentication
- API Gateway properly forwards authentication headers to backend services