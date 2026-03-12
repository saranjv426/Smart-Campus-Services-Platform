import React, { useState, useEffect } from 'react';
import { Link, useSearchParams } from 'react-router-dom';
import { serviceAPI } from '../services/api';
import '../styles/Services.css';

function Services() {
  const [searchParams] = useSearchParams();
  const [services, setServices] = useState([]);
  const [filteredServices, setFilteredServices] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState('all');
  const [searchTerm, setSearchTerm] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const categories = [
    'all',
    'library',
    'dining',
    'transportation',
    'health',
  ];

  useEffect(() => {
    // Check if category is passed via URL query params
    const categoryParam = searchParams.get('category');
    if (categoryParam && categories.includes(categoryParam)) {
      setSelectedCategory(categoryParam);
    }
  }, [searchParams]);

  useEffect(() => {
    fetchServices();
  }, []);

  useEffect(() => {
    filterServices();
  }, [services, selectedCategory, searchTerm]);

  const fetchServices = async () => {
    try {
      setLoading(true);
      const response = await serviceAPI.getServices();
      setServices(response.data || []);
      setError(null);
    } catch (err) {
      setError('Failed to load services');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const filterServices = () => {
    let filtered = services;
    const normalizedSearchTerm = searchTerm.trim().toLowerCase();

    if (selectedCategory !== 'all') {
      filtered = filtered.filter(s => s.category === selectedCategory);
    }

    if (normalizedSearchTerm) {
      filtered = filtered.filter((service) => {
        const name = (service.name || '').toLowerCase();
        const description = (service.description || '').toLowerCase();
        return name.includes(normalizedSearchTerm) || description.includes(normalizedSearchTerm);
      });
    }

    setFilteredServices(filtered);
  };

  return (
    <div className="services-page">
      <h1>Campus Services</h1>
      
      <div className="search-filter-container">
        <input
          type="text"
          placeholder="Search by name or description..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="search-input"
        />
        
        <div className="category-filter">
          {categories.map(cat => (
            <button
              key={cat}
              className={`category-btn ${selectedCategory === cat ? 'active' : ''}`}
              onClick={() => setSelectedCategory(cat)}
            >
              {cat.charAt(0).toUpperCase() + cat.slice(1)}
            </button>
          ))}
        </div>
      </div>

      {loading && <div className="loading">Loading services...</div>}
      {error && <div className="error">{error}</div>}

      {!loading && !error && (
        <>
          <p className="results-count">
            Showing {filteredServices.length} service{filteredServices.length !== 1 ? 's' : ''}
          </p>

          <div className="services-grid">
            {filteredServices.length > 0 ? (
              filteredServices.map(service => (
                <div key={service.id} className="service-card">
                  <img 
                    src={service.imageUrl || 'https://via.placeholder.com/300x200?text=' + service.name}
                    alt={service.name}
                    className="service-card-image"
                  />
                  <div className="service-card-content">
                    <h3>{service.name}</h3>
                    <p className="category-badge">{service.category}</p>
                    <p className="description">{service.description}</p>
                    <div className="service-footer">
                      <span className="rating">⭐ {service.rating.toFixed(1)}</span>
                      <Link to={`/services/${service.id}`} className="view-btn">
                        View Details
                      </Link>
                    </div>
                  </div>
                </div>
              ))
            ) : (
              <div className="no-results">No services found</div>
            )}
          </div>
        </>
      )}
    </div>
  );
}

export default Services;
