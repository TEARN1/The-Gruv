import axios from 'axios';

// Create axios instance with base configuration
const api = axios.create({
  baseURL: '/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

// User authentication API calls
export const authAPI = {
  // Register a new user
  register: async (userData) => {
    try {
      const response = await api.post('/users/register', userData);
      return response.data;
    } catch (error) {
      throw error.response?.data || { error: 'Registration failed' };
    }
  },

  // Login user
  login: async (credentials) => {
    try {
      const response = await api.post('/users/login', credentials);
      return response.data;
    } catch (error) {
      throw error.response?.data || { error: 'Login failed' };
    }
  },

  // Get user profile
  getProfile: async (userId) => {
    try {
      const response = await api.get(`/users/profile/${userId}`);
      return response.data;
    } catch (error) {
      throw error.response?.data || { error: 'Failed to fetch profile' };
    }
  },

  // Update user profile
  updateProfile: async (userId, updates) => {
    try {
      const response = await api.put(`/users/profile/${userId}`, updates);
      return response.data;
    } catch (error) {
      throw error.response?.data || { error: 'Failed to update profile' };
    }
  }
};

// Local storage utilities for user session management
export const storage = {
  // Store user session
  setUser: (userData) => {
    localStorage.setItem('gruv_user', JSON.stringify(userData));
  },

  // Get stored user
  getUser: () => {
    try {
      const userData = localStorage.getItem('gruv_user');
      return userData ? JSON.parse(userData) : null;
    } catch {
      return null;
    }
  },

  // Remove user session
  removeUser: () => {
    localStorage.removeItem('gruv_user');
  },

  // Check if user is logged in
  isLoggedIn: () => {
    return !!storage.getUser();
  }
};

export default api;