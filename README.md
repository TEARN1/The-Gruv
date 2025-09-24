# The Gruv

Welcome to **The Gruv**, a next-generation, AI-native platform designed for creating low-latency, highly-responsive user experiences and powering high-throughput semantic search.

## About

This repository contains the source code for all microservices and components that make up The Gruv platform. The architecture is designed to be modular, scalable, and event-driven, allowing for independent development and deployment of each service.

## Architecture Overview

The platform is built on a microservices architecture, with each service communicating asynchronously through a message broker. The key components include:

- **API Gateway**: The single entry point for all client requests.
- **User Service**: Manages user authentication, profiles, and sessions.
- **Frontend**: React-based web application with dynamic theming
- **Real-time Service**: Handles WebSocket connections for real-time communication.
- **Vector Database Service**: Provides an interface for semantic search and vector embeddings.
- **Content Ingestion Service**: Processes and indexes content for search.

## Features

### Frontend Application
- **Dynamic Themes**: Gender-based color customization (blue, pink, purple, default)
- **User Authentication**: Sign-up and login with profile management
- **Social Feed**: Community posts and interactions
- **Responsive Design**: Mobile-friendly interface

### Backend Services
- **User Management**: Registration, authentication, and profile management
- **API Gateway**: Reverse proxy and request routing
- **Microservices Architecture**: Scalable and maintainable service separation

## Getting Started

### Prerequisites
- Go 1.19+ (for backend services)
- Node.js 16+ (for frontend)
- npm or yarn

### Quick Start

1. **Start Backend Services:**
```bash
# Terminal 1 - User Service
cd user-service
go mod tidy
go build -o user-service .
./user-service

# Terminal 2 - API Gateway  
cd api-gateway
go mod tidy
go build -o api-gateway .
./api-gateway
```

2. **Start Frontend:**
```bash
# Terminal 3 - Frontend
cd frontend
npm install
npm run dev
```

3. **Access the Application:**
   - Frontend: http://localhost:3000
   - API Gateway: http://localhost:8080
   - User Service: http://localhost:8081

### Development Workflow

1. Register a new user account
2. Choose a gender preference to see theme changes
3. Explore profile management and social feed
4. Test theme switching in profile settings

## Project Structure

```
├── api-gateway/          # API Gateway service
├── user-service/         # User management service
├── frontend/             # React frontend application
├── README.md            # This file
└── .gitignore          # Git ignore rules
```

## API Endpoints

### User Service (via API Gateway)
- `POST /api/users/register` - User registration
- `POST /api/users/login` - User authentication  
- `GET /api/users/profile/:id` - Get user profile
- `PUT /api/users/profile/:id` - Update user profile

---

*This project is being developed with the assistance of GitHub Copilot.*
