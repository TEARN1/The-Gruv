import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { useTheme } from './contexts/ThemeContext';
import { storage } from './utils/api';

// Import components
import Login from './components/Login';
import SignUp from './components/SignUp';
import Profile from './components/Profile';
import Feed from './components/Feed';
import Navbar from './components/Navbar';

function App() {
  const { updateTheme } = useTheme();
  const [user, setUser] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Check for stored user on app load
    const storedUser = storage.getUser();
    if (storedUser) {
      setUser(storedUser);
      // Update theme based on user's gender preference
      if (storedUser.user && storedUser.user.gender) {
        updateTheme(storedUser.user.gender);
      }
    }
    setIsLoading(false);
  }, [updateTheme]);

  const handleLogin = (userData) => {
    setUser(userData);
    storage.setUser(userData);
    // Update theme based on user's gender preference
    if (userData.user && userData.user.gender) {
      updateTheme(userData.user.gender);
    }
  };

  const handleLogout = () => {
    setUser(null);
    storage.removeUser();
    updateTheme('default');
  };

  const handleProfileUpdate = (updatedUserData) => {
    const newUserData = { ...user, user: updatedUserData.user };
    setUser(newUserData);
    storage.setUser(newUserData);
    // Update theme if gender changed
    if (updatedUserData.user && updatedUserData.user.gender) {
      updateTheme(updatedUserData.user.gender);
    }
  };

  if (isLoading) {
    return (
      <div style={{ 
        display: 'flex', 
        justifyContent: 'center', 
        alignItems: 'center', 
        height: '100vh' 
      }}>
        Loading...
      </div>
    );
  }

  return (
    <Router>
      <div className="App">
        {user && <Navbar user={user} onLogout={handleLogout} />}
        <main>
          <Routes>
            <Route 
              path="/" 
              element={user ? <Feed user={user} /> : <Navigate to="/login" />} 
            />
            <Route 
              path="/login" 
              element={user ? <Navigate to="/" /> : <Login onLogin={handleLogin} />} 
            />
            <Route 
              path="/signup" 
              element={user ? <Navigate to="/" /> : <SignUp onLogin={handleLogin} />} 
            />
            <Route 
              path="/profile" 
              element={user ? <Profile user={user} onUpdate={handleProfileUpdate} /> : <Navigate to="/login" />} 
            />
            <Route path="*" element={<Navigate to="/" />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;