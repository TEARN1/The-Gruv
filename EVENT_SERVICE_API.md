# The Gruv - Event Service API Documentation

## Overview

The Event Service provides comprehensive event management capabilities including RSVP functionality and location support. This service is part of The Gruv microservices architecture.

## Base URL

```
http://localhost:8082/api/events
```

When accessed through the API Gateway:
```
http://localhost:8080/api/events
```

## Authentication

All endpoints (except health check) require user authentication via the `X-User-ID` header:

```
X-User-ID: {user_id}
```

## Location Types

Events support three types of locations:

1. **Online**: Virtual events with URL links
2. **Venue**: Physical locations with names and addresses
3. **Address**: Specific addresses with optional GPS coordinates

## Event Model

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "creatorId": "string",
  "startTime": "2024-12-15T10:00:00Z",
  "endTime": "2024-12-15T11:00:00Z",
  "location": {
    "type": "venue|online|address",
    "name": "string",
    "address": "string",
    "coordinates": {
      "latitude": 37.7749,
      "longitude": -122.4194
    },
    "url": "string"
  },
  "rsvps": {
    "user123": {
      "status": "going|not_going|maybe",
      "timestamp": "2024-12-15T09:00:00Z",
      "userId": "user123"
    }
  },
  "createdAt": "2024-12-15T08:00:00Z",
  "updatedAt": "2024-12-15T09:00:00Z"
}
```

## API Endpoints

### Health Check

**GET /health**

Returns the service health status.

**Response:**
```json
{
  "service": "Event Service",
  "status": "UP"
}
```

### Create Event

**POST /api/events**

Creates a new event.

**Headers:**
- `X-User-ID: {user_id}` (required)
- `Content-Type: application/json`

**Request Body:**
```json
{
  "title": "Team Meeting",
  "description": "Weekly team sync meeting",
  "startTime": "2024-12-15T10:00:00Z",
  "endTime": "2024-12-15T11:00:00Z",
  "location": {
    "type": "venue",
    "name": "Conference Room A",
    "address": "123 Main St, San Francisco, CA",
    "coordinates": {
      "latitude": 37.7749,
      "longitude": -122.4194
    }
  }
}
```

**Response (201 Created):**
```json
{
  "message": "Event created successfully",
  "event": {
    "id": "event-uuid",
    "title": "Team Meeting",
    // ... full event object
  }
}
```

### List Events

**GET /api/events**

Returns all events.

**Response (200 OK):**
```json
{
  "events": [
    {
      "id": "event-uuid",
      "title": "Team Meeting",
      // ... full event objects
    }
  ],
  "total": 1
}
```

### Get Event

**GET /api/events/{id}**

Returns a specific event by ID.

**Response (200 OK):**
```json
{
  "event": {
    "id": "event-uuid",
    "title": "Team Meeting",
    // ... full event object
  }
}
```

### Update Event

**PUT /api/events/{id}**

Updates an existing event. Only the event creator can update the event.

**Headers:**
- `X-User-ID: {user_id}` (required)
- `Content-Type: application/json`

**Request Body (partial updates allowed):**
```json
{
  "title": "Updated Team Meeting",
  "location": {
    "type": "online",
    "name": "Zoom Meeting",
    "url": "https://zoom.us/j/1234567890"
  }
}
```

**Response (200 OK):**
```json
{
  "message": "Event updated successfully",
  "event": {
    // ... updated event object
  }
}
```

### Delete Event

**DELETE /api/events/{id}**

Deletes an event. Only the event creator can delete the event.

**Headers:**
- `X-User-ID: {user_id}` (required)

**Response (200 OK):**
```json
{
  "message": "Event deleted successfully"
}
```

## RSVP Endpoints

### RSVP to Event

**POST /api/events/{id}/rsvp**

Create or update an RSVP for an event.

**Headers:**
- `X-User-ID: {user_id}` (required)
- `Content-Type: application/json`

**Request Body:**
```json
{
  "status": "going|not_going|maybe"
}
```

**Response (200 OK):**
```json
{
  "message": "RSVP updated successfully",
  "rsvp": {
    "status": "going",
    "timestamp": "2024-12-15T09:00:00Z",
    "userId": "user123"
  }
}
```

### Get Event RSVPs

**GET /api/events/{id}/rsvps**

Get all RSVPs for an event with summary counts.

**Response (200 OK):**
```json
{
  "eventId": "event-uuid",
  "rsvps": [
    {
      "status": "going",
      "timestamp": "2024-12-15T09:00:00Z",
      "userId": "user123"
    }
  ],
  "counts": {
    "going": 1,
    "not_going": 0,
    "maybe": 0
  },
  "total": 1
}
```

### Remove RSVP

**DELETE /api/events/{id}/rsvp**

Remove an RSVP from an event.

**Headers:**
- `X-User-ID: {user_id}` (required)

**Response (200 OK):**
```json
{
  "message": "RSVP removed successfully"
}
```

## Location Features

### Physical Locations (Venue/Address)

For events with physical locations, the system supports:

- **Address Information**: Full addresses for directions
- **GPS Coordinates**: Latitude/longitude for mapping integration
- **Location Icon**: UI can display location icons for non-online events
- **Directions**: Coordinates enable integration with mapping services

**Example Physical Location:**
```json
{
  "location": {
    "type": "venue",
    "name": "Conference Room A",
    "address": "123 Main St, San Francisco, CA",
    "coordinates": {
      "latitude": 37.7749,
      "longitude": -122.4194
    }
  }
}
```

### Online Locations

For virtual events:

**Example Online Location:**
```json
{
  "location": {
    "type": "online",
    "name": "Zoom Meeting",
    "url": "https://zoom.us/j/1234567890"
  }
}
```

## Error Responses

All endpoints return consistent error responses:

**400 Bad Request:**
```json
{
  "error": "Invalid input: validation error message"
}
```

**401 Unauthorized:**
```json
{
  "error": "User ID required"
}
```

**403 Forbidden:**
```json
{
  "error": "Only event creator can update the event"
}
```

**404 Not Found:**
```json
{
  "error": "Event not found"
}
```

## Business Rules

1. **Event Times**: End time must be after start time
2. **RSVP Restrictions**: Cannot RSVP to past events (after end time)
3. **Event Permissions**: Only event creators can update/delete events
4. **RSVP Management**: Users can only manage their own RSVPs
5. **Location Requirements**: All events must have a location specified

## Integration Notes

### Frontend Integration

- Use location type to determine whether to show location icon
- For non-online events, provide directions link using coordinates
- Display RSVP counts and allow users to change their RSVP status
- Show different UI for past vs. future events

### Mapping Integration

Coordinates can be used with mapping services:
```javascript
// Google Maps example
const mapsUrl = `https://www.google.com/maps?q=${latitude},${longitude}`;

// Apple Maps example  
const appleMapsUrl = `https://maps.apple.com/?q=${latitude},${longitude}`;
```