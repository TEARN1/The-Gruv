# The Gruv Frontend

A modern React frontend application for The Gruv platform featuring dynamic themes based on gender preferences.

## Features

- **User Authentication**: Sign-up and login functionality
- **Dynamic Themes**: Gender-based color schemes that automatically customize the UI
  - Male: Blue theme
  - Female: Pink theme  
  - Other: Purple theme
  - Default: Gray theme
- **Profile Management**: Users can update their profile information and theme preferences
- **Social Feed**: Community feed for sharing posts and connecting with others
- **Responsive Design**: Mobile-friendly interface

## Tech Stack

- **React 19**: Modern frontend library
- **Vite**: Fast build tool and dev server
- **React Router**: Client-side routing
- **Axios**: HTTP client for API communication
- **CSS Custom Properties**: For dynamic theming
- **Styled Components**: Component-based styling

## Getting Started

### Prerequisites

- Node.js (v16 or higher)
- npm or yarn

### Installation

1. Install dependencies:
```bash
npm install
```

2. Start the development server:
```bash
npm run dev
```

3. Open your browser to `http://localhost:3000`

### Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build

## Architecture

### Components

- **Login**: User authentication form
- **SignUp**: User registration with theme selection
- **Profile**: User profile management
- **Feed**: Social feed with posts
- **Navbar**: Navigation and user menu

### Context

- **ThemeContext**: Manages dynamic themes based on user preferences

### API Integration

The frontend communicates with the backend API through:
- `/api/users/register` - User registration
- `/api/users/login` - User authentication
- `/api/users/profile/:id` - Profile management

## Theme System

The application uses a dynamic theme system that changes colors based on user gender preferences:

```javascript
// Theme updates automatically when user changes gender preference
updateTheme(gender); // 'male', 'female', 'other', or 'default'
```

Themes are applied using CSS custom properties that can be easily customized.

## Development

The app is configured to proxy API requests to `http://localhost:8080` during development.

### Backend Services Required

- API Gateway: `localhost:8080`
- User Service: `localhost:8081`

Make sure both backend services are running before starting the frontend.