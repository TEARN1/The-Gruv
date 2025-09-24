import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { useTheme } from '../contexts/ThemeContext';
import { authAPI } from '../utils/api';

const SignUp = ({ onLogin }) => {
  const { theme, updateTheme } = useTheme();
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    confirmPassword: '',
    gender: '',
    email: ''
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value
    });
    
    // Update theme preview when gender changes
    if (name === 'gender' && value) {
      updateTheme(value);
    }
    
    // Clear error when user starts typing
    if (error) setError('');
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');

    // Validate passwords match
    if (formData.password !== formData.confirmPassword) {
      setError('Passwords do not match');
      setIsLoading(false);
      return;
    }

    try {
      // Register user
      const registerResponse = await authAPI.register({
        username: formData.username,
        password: formData.password,
        gender: formData.gender,
        email: formData.email
      });

      // Auto-login after successful registration
      const loginResponse = await authAPI.login({
        username: formData.username,
        password: formData.password
      });

      onLogin(loginResponse);
    } catch (err) {
      setError(err.error || 'Registration failed. Please try again.');
      // Reset theme to default on error
      if (!formData.gender) {
        updateTheme('default');
      }
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div style={{ 
      minHeight: '100vh', 
      display: 'flex', 
      alignItems: 'center', 
      justifyContent: 'center',
      padding: '1rem'
    }}>
      <div className="card" style={{ width: '100%', maxWidth: '500px' }}>
        <div style={{ textAlign: 'center', marginBottom: '2rem' }}>
          <h1 style={{ 
            fontSize: 'var(--font-size-3xl)', 
            fontWeight: 'bold', 
            color: 'var(--color-primary)',
            marginBottom: '0.5rem'
          }}>
            Join The Gruv
          </h1>
          <p style={{ color: 'var(--color-textSecondary)' }}>
            Create your account and experience personalized themes
          </p>
        </div>

        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="username" className="form-label">
              Username *
            </label>
            <input
              type="text"
              id="username"
              name="username"
              className="form-input"
              value={formData.username}
              onChange={handleChange}
              required
              placeholder="Choose a username"
            />
          </div>

          <div className="form-group">
            <label htmlFor="email" className="form-label">
              Email (optional)
            </label>
            <input
              type="email"
              id="email"
              name="email"
              className="form-input"
              value={formData.email}
              onChange={handleChange}
              placeholder="your@email.com"
            />
          </div>

          <div className="form-group">
            <label htmlFor="gender" className="form-label">
              Gender Preference (affects theme)
            </label>
            <select
              id="gender"
              name="gender"
              className="form-select"
              value={formData.gender}
              onChange={handleChange}
            >
              <option value="">Select gender preference</option>
              <option value="male">Male (Blue theme)</option>
              <option value="female">Female (Pink theme)</option>
              <option value="other">Other (Purple theme)</option>
            </select>
            <p style={{ 
              fontSize: 'var(--font-size-sm)', 
              color: 'var(--color-textSecondary)',
              marginTop: '0.25rem'
            }}>
              This customizes your app's color scheme. You can change it later.
            </p>
          </div>

          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem' }}>
            <div className="form-group">
              <label htmlFor="password" className="form-label">
                Password *
              </label>
              <input
                type="password"
                id="password"
                name="password"
                className="form-input"
                value={formData.password}
                onChange={handleChange}
                required
                placeholder="Choose a password"
              />
            </div>

            <div className="form-group">
              <label htmlFor="confirmPassword" className="form-label">
                Confirm Password *
              </label>
              <input
                type="password"
                id="confirmPassword"
                name="confirmPassword"
                className="form-input"
                value={formData.confirmPassword}
                onChange={handleChange}
                required
                placeholder="Confirm password"
              />
            </div>
          </div>

          {error && (
            <div className="error-message" style={{ marginBottom: '1rem' }}>
              {error}
            </div>
          )}

          <button 
            type="submit" 
            className="btn"
            disabled={isLoading}
            style={{ 
              width: '100%', 
              marginBottom: '1rem',
              opacity: isLoading ? 0.7 : 1,
              cursor: isLoading ? 'not-allowed' : 'pointer'
            }}
          >
            {isLoading ? 'Creating Account...' : 'Create Account'}
          </button>
        </form>

        <div style={{ 
          textAlign: 'center', 
          paddingTop: '1rem', 
          borderTop: '1px solid var(--color-border)' 
        }}>
          <p style={{ color: 'var(--color-textSecondary)' }}>
            Already have an account?{' '}
            <Link 
              to="/login" 
              style={{ 
                color: 'var(--color-primary)', 
                textDecoration: 'none',
                fontWeight: '500'
              }}
            >
              Sign in
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
};

export default SignUp;