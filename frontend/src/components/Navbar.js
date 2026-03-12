import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import '../styles/Navbar.css';

function Navbar() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [currentUser, setCurrentUser] = useState(null);
  const navigate = useNavigate();

  const checkAuthStatus = () => {
    const token = localStorage.getItem('token');
    const user = localStorage.getItem('user');
    if (token && user) {
      setIsLoggedIn(true);
      setCurrentUser(JSON.parse(user));
    } else {
      setIsLoggedIn(false);
      setCurrentUser(null);
    }
  };

  useEffect(() => {
    // Check auth status on mount
    checkAuthStatus();

    // Listen for storage changes (from login/logout in other components)
    const handleStorageChange = () => {
      checkAuthStatus();
    };

    // Listen for custom auth change event
    const handleAuthChange = () => {
      checkAuthStatus();
    };
    
    window.addEventListener('storage', handleStorageChange);
    window.addEventListener('authChange', handleAuthChange);
    
    return () => {
      window.removeEventListener('storage', handleStorageChange);
      window.removeEventListener('authChange', handleAuthChange);
    };
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    setIsLoggedIn(false);
    setCurrentUser(null);
    
    // Dispatch custom event to notify other components of logout
    window.dispatchEvent(new Event('authChange'));
    
    navigate('/');
  };

  return (
    <nav className="navbar">
      <div className="navbar-container">
        <Link to="/" className="navbar-logo">
          🏫 Smart Campus Services
        </Link>
        <div className="navbar-menu">
          <Link to="/" className="nav-link">Home</Link>
          <Link to="/services" className="nav-link">Services</Link>
          {isLoggedIn && currentUser?.role === 'student' && <Link to="/bookings" className="nav-link">My Bookings</Link>}
          {isLoggedIn && currentUser?.role === 'staff' && <Link to="/dashboard/staff" className="nav-link">Staff Dashboard</Link>}
          {isLoggedIn && currentUser?.role === 'admin' && <Link to="/dashboard/admin" className="nav-link">Admin Dashboard</Link>}
          
          <div className="navbar-auth">
            {isLoggedIn ? (
              <>
                <span className="welcome-text">Welcome, {currentUser?.firstName}</span>
                {currentUser?.role === 'student' && <Link to={`/profile/${currentUser?.id}`} className="nav-link">Profile</Link>}
                <button onClick={handleLogout} className="logout-btn">Logout</button>
              </>
            ) : (
              <>
                <Link to="/login" className="nav-link">Login</Link>
                <Link to="/register" className="nav-link register-btn">Register</Link>
              </>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
