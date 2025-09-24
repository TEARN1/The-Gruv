import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { useTheme } from '../contexts/ThemeContext';

const Navbar = ({ user, onLogout }) => {
  const { theme } = useTheme();
  const location = useLocation();

  const isActive = (path) => location.pathname === path;

  return (
    <nav style={{
      backgroundColor: 'var(--color-surface)',
      borderBottom: '1px solid var(--color-border)',
      padding: '0',
      boxShadow: 'var(--shadow-sm)'
    }}>
      <div className="container">
        <div style={{
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'center',
          height: '64px'
        }}>
          {/* Logo */}
          <Link 
            to="/" 
            style={{
              fontSize: 'var(--font-size-xl)',
              fontWeight: 'bold',
              color: 'var(--color-primary)',
              textDecoration: 'none',
              display: 'flex',
              alignItems: 'center',
              gap: '0.5rem'
            }}
          >
            <div style={{
              width: '32px',
              height: '32px',
              borderRadius: 'var(--radius-lg)',
              backgroundColor: 'var(--color-primary)',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              color: 'white',
              fontWeight: 'bold'
            }}>
              G
            </div>
            The Gruv
          </Link>

          {/* Navigation Links */}
          <div style={{ display: 'flex', alignItems: 'center', gap: '1rem' }}>
            <Link
              to="/"
              style={{
                padding: '0.5rem 1rem',
                borderRadius: 'var(--radius-md)',
                textDecoration: 'none',
                color: isActive('/') ? 'var(--color-primary)' : 'var(--color-textSecondary)',
                fontWeight: isActive('/') ? '600' : '400',
                backgroundColor: isActive('/') ? 'var(--color-background)' : 'transparent',
                transition: 'all 0.2s ease'
              }}
            >
              Feed
            </Link>

            <Link
              to="/profile"
              style={{
                padding: '0.5rem 1rem',
                borderRadius: 'var(--radius-md)',
                textDecoration: 'none',
                color: isActive('/profile') ? 'var(--color-primary)' : 'var(--color-textSecondary)',
                fontWeight: isActive('/profile') ? '600' : '400',
                backgroundColor: isActive('/profile') ? 'var(--color-background)' : 'transparent',
                transition: 'all 0.2s ease'
              }}
            >
              Profile
            </Link>

            {/* User Menu */}
            <div style={{ 
              display: 'flex', 
              alignItems: 'center', 
              gap: '1rem',
              paddingLeft: '1rem',
              marginLeft: '1rem',
              borderLeft: '1px solid var(--color-border)'
            }}>
              <span style={{ 
                color: 'var(--color-textSecondary)',
                fontSize: 'var(--font-size-sm)'
              }}>
                Welcome, {user?.user?.username || 'User'}
              </span>

              {/* Theme indicator */}
              <div style={{
                width: '20px',
                height: '20px',
                borderRadius: '50%',
                backgroundColor: 'var(--color-primary)',
                border: '2px solid var(--color-border)',
                title: `Theme: ${user?.user?.gender || 'default'}`
              }} />

              <button
                onClick={onLogout}
                className="btn btn-outline"
                style={{ 
                  fontSize: 'var(--font-size-sm)',
                  padding: '0.25rem 0.75rem'
                }}
              >
                Logout
              </button>
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;