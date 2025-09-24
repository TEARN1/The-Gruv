# Comment Service API Documentation

## Overview
The Comment Service provides functionality for creating, managing, and moderating comments with support for video uploads, threaded replies, and reporting features.

## Endpoints

### Health Check
- **GET** `/health`
- Returns service status

### Comments

#### Create Comment
- **POST** `/comments`
- **Body**:
  ```json
  {
    "userId": "string",
    "content": "string", 
    "videoUrl": "string (optional)",
    "videoDuration": "number (optional, in seconds)",
    "parentId": "string (optional, for threaded replies)"
  }
  ```
- **Response**: 201 Created with comment object
- **Validation**: Video uploads are validated based on account age (see Video Validation section)

#### Get Comments
- **GET** `/comments`
- **Query Parameters**:
  - `parentId`: Filter by parent comment (use "null" or empty for top-level comments)
- **Response**: Array of comment objects

#### Get Single Comment  
- **GET** `/comments/:id`
- **Response**: Single comment object

#### Delete Comment
- **DELETE** `/comments/:id` 
- **Headers**:
  - `X-User-ID`: Required, user making the request
  - `X-Is-Admin`: Set to "true" for admin privileges
- **Authorization**: User must own the comment or be an admin

### Moderation

#### Report Comment
- **POST** `/comments/:id/report`
- **Body**:
  ```json
  {
    "reporterId": "string",
    "reason": "string"
  }
  ```
- **Response**: 201 Created with report object

#### Get Reports (Admin Only)
- **GET** `/reports`
- **Headers**:
  - `X-Is-Admin`: Must be "true"
- **Response**: Array of report objects

### Video Validation

#### Get Video Validation Rules
- **GET** `/users/:userId/video-validation`
- **Response**:
  ```json
  {
    "canUpload": "boolean",
    "maxDuration": "number (seconds)", 
    "accountAge": "number (months)"
  }
  ```

## Video Upload Rules

1. **Account Age Requirement**: Accounts must be at least 18 months old to upload video comments
2. **Duration Limits**:
   - 18 months: 5 seconds maximum
   - 24 months (18 + 6): 10 seconds maximum  
   - 30 months (18 + 12): 15 seconds maximum
   - 36+ months: 20 seconds maximum (cap)
3. **Calculation**: Duration increases by 5 seconds every 6 months after the initial 18-month period

## Data Models

### Comment
```json
{
  "id": "string",
  "userId": "string", 
  "content": "string",
  "videoUrl": "string (optional)",
  "parentId": "string (optional)",
  "createdAt": "timestamp",
  "updatedAt": "timestamp"
}
```

### Report  
```json
{
  "id": "string",
  "commentId": "string",
  "reporterId": "string", 
  "reason": "string",
  "createdAt": "timestamp"
}
```

## Threading Structure

Comments support hierarchical threading through the `parentId` field:
- Top-level comments have `parentId: null`
- Replies reference their parent comment's ID
- Use the `parentId` query parameter to fetch replies to a specific comment

## Error Responses

All endpoints return appropriate HTTP status codes:
- `400 Bad Request`: Invalid input or validation errors
- `401 Unauthorized`: Missing authentication
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource conflicts (e.g., duplicate usernames)
- `500 Internal Server Error`: Server-side errors

Error response format:
```json
{
  "error": "Error message description"
}
```