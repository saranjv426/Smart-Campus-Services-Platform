import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import '../styles/Home.css';

function Home() {
  const navigate = useNavigate();
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    setIsLoggedIn(!!token);
  }, []);

  const features = [
    {
      icon: '📚',
      title: 'Library Services',
      description: 'Book study rooms, access library resources, and manage your reading materials',
      category: 'library'
    },
    {
      icon: '🍽️',
      title: 'Dining Services',
      description: 'Check menus, make dining reservations, and explore campus cafeterias',
      category: 'dining'
    },
    {
      icon: '🚌',
      title: 'Transportation',
      description: 'Book campus shuttles and view public transportation schedules',
      category: 'transportation'
    },
    {
      icon: '⚕️',
      title: 'Health Services',
      description: 'Schedule appointments with health center and access wellness resources',
      category: 'health'
    }
  ];

  const handleFeatureClick = (category) => {
    navigate(`/services?category=${category}`);
  };

  const testimonials = [
    {
      name: 'Sarah Johnson',
      role: 'Senior, Computer Science',
      text: 'Smart Campus Services made managing my bookings so easy. I love how I can do everything in one place!'
    },
    {
      name: 'Mike Chen',
      role: 'Junior, Business Administration',
      text: 'The notification system keeps me updated on all my reservations. Highly recommended!'
    },
    {
      name: 'Emma Davis',
      role: 'Staff, Library Services',
      text: 'This platform has streamlined our service delivery and improved student satisfaction significantly.'
    }
  ];

  return (
    <div className="home-page">
      <section className="hero">
        <div className="hero-content">
          <h1>Smart Campus Services Platform</h1>
          <p>Your One-Stop Solution for University of Florida Campus Services</p>
          <p className="subtitle">Easily discover, book, and manage all your campus needs in one convenient platform</p>
          <div className="hero-actions">
            <Link to="/services" className="btn btn-primary">Browse Services</Link>
            {!isLoggedIn && <Link to="/register" className="btn btn-secondary">Get Started</Link>}
          </div>
        </div>
      </section>

      <section className="features-section">
        <h2>Our Services</h2>
        <p className="section-subtitle">Everything you need for a better campus experience</p>
        <div className="features-grid">
          {features.map((feature, index) => (
            <div 
              key={index} 
              className="feature-card"
              onClick={() => handleFeatureClick(feature.category)}
              style={{ cursor: 'pointer' }}
              role="button"
              tabIndex={0}
              onKeyPress={(e) => {
                if (e.key === 'Enter') handleFeatureClick(feature.category);
              }}
            >
              <div className="feature-icon">{feature.icon}</div>
              <h3>{feature.title}</h3>
              <p>{feature.description}</p>
            </div>
          ))}
        </div>
      </section>

      <section className="benefits-section">
        <h2>Why Choose Smart Campus Services?</h2>
        <div className="benefits-list">
          <div className="benefit-item">
            <span className="benefit-number">1</span>
            <h3>Easy Booking System</h3>
            <p>Intuitive interface to book services with just a few clicks</p>
          </div>
          <div className="benefit-item">
            <span className="benefit-number">2</span>
            <h3>Real-Time Notifications</h3>
            <p>Get instant updates about your bookings and service availability</p>
          </div>
          <div className="benefit-item">
            <span className="benefit-number">3</span>
            <h3>User Reviews & Ratings</h3>
            <p>Read reviews from other students to make informed decisions</p>
          </div>
          <div className="benefit-item">
            <span className="benefit-number">4</span>
            <h3>Comprehensive Services</h3>
            <p>Access a wide range of campus services from one platform</p>
          </div>
        </div>
      </section>

      <section className="testimonials-section">
        <h2>What Students Say</h2>
        <div className="testimonials-grid">
          {testimonials.map((testimonial, index) => (
            <div key={index} className="testimonial-card">
              <p className="testimonial-text">"{testimonial.text}"</p>
              <div className="testimonial-author">
                <strong>{testimonial.name}</strong>
                <small>{testimonial.role}</small>
              </div>
            </div>
          ))}
        </div>
      </section>

      <section className="cta-section">
        <h2>Ready to Experience Smart Campus Services?</h2>
        <p>Join thousands of UF students making their campus life easier</p>
        <div className="cta-actions">
          <Link to="/register" className="btn btn-primary">Register Now</Link>
          <Link to="/services" className="btn btn-outline">Explore Services</Link>
        </div>
      </section>
    </div>
  );
}

export default Home;
