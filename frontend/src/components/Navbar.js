import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import '../styles/Navbar.css';
import { getRoleBasedLinks } from '../utils/navbarLinks';

function Navbar() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [currentUser, setCurrentUser] = useState(null);
  const navigate = useNavigate();

  const checkAuthStatus = () => {
    const token = localStorage.getItem('token');
    const user = localStorage.getItem('user');
    if (token && user) {
      try {
        const parsedUser = JSON.parse(user);
        setIsLoggedIn(true);
        setCurrentUser(parsedUser);
      } catch (error) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      setIsLoggedIn(false);
      setCurrentUser(null);
    }
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
  const { navLinks, authLinks } = getRoleBasedLinks(isLoggedIn, currentUser);

  return (
    <nav className="navbar">
      <div className="navbar-container">
        <Link to="/" className="navbar-logo">
          🏫 Smart Campus Services
        </Link>
        <div className="navbar-menu">
  {navLinks.map((link) => (
    <Link key={link.label} to={link.path} className="nav-link">
      {link.label}
    </Link>
  ))}

  <div className="navbar-auth">
    {isLoggedIn ? (
      <>
        <span className="welcome-text">Welcome, {currentUser?.firstName}</span>

        {authLinks.map((link) => (
          <Link
            key={link.label}
            to={link.path}
            className={`nav-link ${link.className || ''}`.trim()}
          >
            {link.label}
          </Link>
        ))}

        <button onClick={handleLogout} className="logout-btn">
          Logout
        </button>
      </>
    ) : (
      <>
        {authLinks.map((link) => (
          <Link
            key={link.label}
            to={link.path}
            className={`nav-link ${link.className || ''}`.trim()}
          >
            {link.label}
          </Link>
        ))}
      </>
    )}
  </div>
</div>
      </div>
    </nav>
  );
}

export default Navbar;
