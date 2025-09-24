import React, { createContext, useContext, useState, useEffect } from 'react';

// Gender-based theme configurations
const themes = {
  male: {
    primary: '#2563eb',      // Blue
    secondary: '#1e40af',    // Darker blue
    accent: '#3b82f6',       // Light blue
    background: '#f8fafc',   // Light gray
    surface: '#ffffff',      // White
    text: '#1e293b',         // Dark gray
    textSecondary: '#64748b', // Medium gray
    border: '#e2e8f0',       // Light border
    success: '#10b981',      // Green
    warning: '#f59e0b',      // Amber
    error: '#ef4444',        // Red
  },
  female: {
    primary: '#ec4899',      // Pink
    secondary: '#be185d',    // Darker pink
    accent: '#f472b6',       // Light pink
    background: '#fdf2f8',   // Very light pink
    surface: '#ffffff',      // White
    text: '#1e293b',         // Dark gray
    textSecondary: '#64748b', // Medium gray
    border: '#fce7f3',       // Light pink border
    success: '#10b981',      // Green
    warning: '#f59e0b',      // Amber
    error: '#ef4444',        // Red
  },
  other: {
    primary: '#7c3aed',      // Purple
    secondary: '#5b21b6',    // Darker purple
    accent: '#a855f7',       // Light purple
    background: '#faf5ff',   // Very light purple
    surface: '#ffffff',      // White
    text: '#1e293b',         // Dark gray
    textSecondary: '#64748b', // Medium gray
    border: '#e9d5ff',       // Light purple border
    success: '#10b981',      // Green
    warning: '#f59e0b',      // Amber
    error: '#ef4444',        // Red
  },
  default: {
    primary: '#374151',      // Gray
    secondary: '#111827',    // Darker gray
    accent: '#6b7280',       // Medium gray
    background: '#f9fafb',   // Very light gray
    surface: '#ffffff',      // White
    text: '#1e293b',         // Dark gray
    textSecondary: '#64748b', // Medium gray
    border: '#e5e7eb',       // Light gray border
    success: '#10b981',      // Green
    warning: '#f59e0b',      // Amber
    error: '#ef4444',        // Red
  }
};

const ThemeContext = createContext();

export const useTheme = () => {
  const context = useContext(ThemeContext);
  if (!context) {
    throw new Error('useTheme must be used within a ThemeProvider');
  }
  return context;
};

export const ThemeProvider = ({ children }) => {
  const [currentGender, setCurrentGender] = useState('default');
  const [currentTheme, setCurrentTheme] = useState(themes.default);

  const updateTheme = (gender) => {
    const theme = themes[gender] || themes.default;
    setCurrentGender(gender);
    setCurrentTheme(theme);
    
    // Apply theme to CSS custom properties
    const root = document.documentElement;
    Object.entries(theme).forEach(([key, value]) => {
      root.style.setProperty(`--color-${key}`, value);
    });
  };

  useEffect(() => {
    // Initialize theme on mount
    updateTheme(currentGender);
  }, []);

  return (
    <ThemeContext.Provider 
      value={{
        theme: currentTheme,
        gender: currentGender,
        updateTheme,
        themes
      }}
    >
      {children}
    </ThemeContext.Provider>
  );
};