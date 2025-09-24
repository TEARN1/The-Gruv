# The Gruv

Welcome to **The Gruv**, a next-generation, AI-native platform designed for creating low-latency, highly-responsive user experiences and powering high-throughput semantic search.

## About

This repository contains the source code for all microservices and components that make up The Gruv platform. The architecture is designed to be modular, scalable, and event-driven, allowing for independent development and deployment of each service.

## Architecture Overview

The platform is built on a microservices architecture, with each service communicating asynchronously through a message broker. The key components include:

- **API Gateway**: The single entry point for all client requests.
- **User Service**: Manages user authentication, profiles, and sessions.
- **Event Service**: Handles event management, RSVP functionality, and location services.
- **Real-time Service**: Handles WebSocket connections for real-time communication.
- **Vector Database Service**: Provides an interface for semantic search and vector embeddings.
- **Content Ingestion Service**: Processes and indexes content for search.

## Services

### User Service (Port 8081)
Handles user registration, authentication, and profile management.

### Event Service (Port 8082)
Manages events with comprehensive RSVP and location features:
- Event creation and management (CRUD operations)
- RSVP functionality (going/not_going/maybe)
- Location support (online, venue, address with GPS coordinates)
- Permission-based access control

See [EVENT_SERVICE_API.md](./EVENT_SERVICE_API.md) for detailed API documentation.

### API Gateway (Port 8080)
Routes requests to appropriate microservices:
- `/api/users/*` → User Service
- `/api/events/*` → Event Service
- `/api/collaboration/*` → Collaboration Service (planned)

## Getting Started

### Prerequisites
- Go 1.21 or later
- Git

### Running the Services

1. **Start the User Service:**
   ```bash
   cd user-service
   go mod tidy
   go run .
   ```

2. **Start the Event Service:**
   ```bash
   cd event-service  
   go mod tidy
   go run .
   ```

3. **Start the API Gateway:**
   ```bash
   cd api-gateway
   go mod tidy
   go run .
   ```

The API Gateway will be available at `http://localhost:8080` and will route requests to the appropriate services.

## Features

### Event Management
- Create events with rich location information
- Support for online and physical event locations
- GPS coordinates for directions to physical venues
- Event CRUD operations with proper permissions

### RSVP System  
- Users can RSVP with status: going, not going, maybe
- RSVP tracking and summary counts
- Prevents RSVPs to past events
- RSVP removal functionality

### Location Services
- Three location types: online, venue, address
- GPS coordinates for mapping integration
- Location icons for UI differentiation
- Directions support for physical locations

## Getting Started

*(This section will be updated with instructions on how to build, test, and run the project locally.)*

---

*This project is being developed with the assistance of GitHub Copilot.*
