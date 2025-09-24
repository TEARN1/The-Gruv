# Authentication API Documentation

## Overview

The authentication system provides secure user registration, login, and profile management with JWT-based session handling.

## Endpoints

### 1. Sign Up
**POST** `/auth/signup`

Register a new user account.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "gender": "male"
}
```

**Response (201 Created):**
```json
{
  "message": "User registered successfully",
  "user": {
    "id": "68d3421c4a57926acedc0cbc",
    "name": "John Doe",
    "email": "john@example.com",
    "gender": "male",
    "createdAt": "2025-01-15T10:30:00Z"
  }
}
```

**Error Responses:**
- `400 Bad Request` - Invalid input data
- `409 Conflict` - Email already registered

### 2. Login
**POST** `/auth/login`

Authenticate user and receive JWT token.

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (200 OK):**
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "68d3421c4a57926acedc0cbc",
    "name": "John Doe",
    "email": "john@example.com",
    "gender": "male",
    "createdAt": "2025-01-15T10:30:00Z"
  }
}
```

**Error Responses:**
- `400 Bad Request` - Invalid input data
- `401 Unauthorized` - Invalid email or password

### 3. Get Profile
**GET** `/auth/profile`

Get current user's profile (requires authentication).

**Headers:**
```
Authorization: Bearer <jwt_token>
```

**Response (200 OK):**
```json
{
  "user": {
    "id": "68d3421c4a57926acedc0cbc",
    "name": "John Doe",
    "email": "john@example.com",
    "gender": "male",
    "createdAt": "2025-01-15T10:30:00Z",
    "updatedAt": "2025-01-15T10:30:00Z"
  }
}
```

**Error Responses:**
- `401 Unauthorized` - Missing, invalid, or expired token
- `404 Not Found` - User not found

## Security Features

- **Password Hashing**: bcrypt with default cost
- **JWT Tokens**: 24-hour expiration, HS256 algorithm
- **Email Validation**: Built-in email format validation
- **Input Validation**: Required field validation and constraints
- **Error Handling**: Consistent error messages without information leakage

## Database Schema

### User Model
```json
{
  "id": "ObjectId",
  "name": "string",
  "email": "string (unique, lowercase)",
  "password": "string (hashed)",
  "gender": "string",
  "createdAt": "timestamp",
  "updatedAt": "timestamp"
}
```

## Authentication Flow

1. **Registration**: User provides name, email, password, and gender
2. **Password Hashing**: System hashes password using bcrypt
3. **Storage**: User data stored in MongoDB (or in-memory for development)
4. **Login**: User provides email and password
5. **Verification**: System verifies password against hash
6. **Token Generation**: JWT token issued with user information
7. **Protected Access**: Token required for profile endpoint

## Development Setup

The system supports both MongoDB and in-memory storage:

- **MongoDB**: Set `MONGODB_URI` environment variable
- **In-Memory**: Leave `MONGODB_URI` unset for development testing

## Example Usage

```bash
# Register new user
curl -X POST http://localhost:8081/auth/signup \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"password123","gender":"male"}'

# Login and get token  
TOKEN=$(curl -X POST http://localhost:8081/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}' \
  | jq -r '.token')

# Access protected profile
curl -X GET http://localhost:8081/auth/profile \
  -H "Authorization: Bearer $TOKEN"
```