import React, { useState, useEffect } from 'react';
import { bookingAPI } from '../services/api';
import '../styles/Bookings.css';

function Bookings() {
  const [bookings, setBookings] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [filter, setFilter] = useState('all');
  const [refreshKey, setRefreshKey] = useState(0);

  useEffect(() => {
    fetchBookings();
  }, [refreshKey]);

  // Auto-refresh every 30 seconds
  useEffect(() => {
    const interval = setInterval(() => {
      setRefreshKey(prev => prev + 1);
    }, 30000);
    return () => clearInterval(interval);
  }, []);

  const fetchBookings = async () => {
    try {
      const currentUser = JSON.parse(localStorage.getItem('user'));
      if (!currentUser) {
        setError('Please login to view your bookings');
        setLoading(false);
        return;
      }

      const response = await bookingAPI.getUserBookings(currentUser.id);
      setBookings(response.data || []);
      setError(null);
    } catch (err) {
      setError('Failed to load bookings');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleCancelBooking = async (bookingId) => {
    if (window.confirm('Are you sure you want to cancel this booking?')) {
      try {
        await bookingAPI.cancelBooking(bookingId);
        alert('Booking cancelled successfully');
        fetchBookings();
      } catch (err) {
        alert('Failed to cancel booking');
      }
    }
  };

  const normalizeStatus = (status) => {
    if (status === 'completed') {
      return 'approved';
    }
    return status;
  };

  const getStatusLabel = (status) => {
    if (status === 'approved') {
      return 'Accepted';
    }
    return status.charAt(0).toUpperCase() + status.slice(1);
  };

  const getFilterLabel = (status) => {
    if (status === 'all') {
      return 'All';
    }
    return getStatusLabel(status);
  };

  const filteredBookings = bookings.filter((b) => {
    const normalizedStatus = normalizeStatus(b.status);
    return filter === 'all' ? true : normalizedStatus === filter;
  });

  const statuses = ['all', 'pending', 'approved', 'rejected', 'cancelled'];

  return (
    <div className="bookings-page">
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '20px' }}>
        <h1>My Bookings</h1>
        <button 
          onClick={() => setRefreshKey(prev => prev + 1)}
          style={{ 
            padding: '8px 16px', 
            backgroundColor: '#667eea', 
            color: 'white', 
            border: 'none', 
            borderRadius: '6px',
            cursor: 'pointer',
            fontSize: '14px',
            fontWeight: '500'
          }}
        >
          🔄 Refresh
        </button>
      </div>

      <div className="status-filter">
        {statuses.map(status => (
          <button
            key={status}
            className={`filter-btn ${filter === status ? 'active' : ''}`}
            onClick={() => setFilter(status)}
          >
            {getFilterLabel(status)}
            {filter === status && ` (${filteredBookings.length})`}
          </button>
        ))}
      </div>

      {loading && <div className="loading">Loading your bookings...</div>}
      {error && <div className="error">{error}</div>}

      {!loading && !error && (
        <div className="bookings-list">
          {filteredBookings.length > 0 ? (
            filteredBookings.map(booking => {
              const normalizedStatus = normalizeStatus(booking.status);
              return (
              <div key={booking.id} className={`booking-card ${normalizedStatus}`}>
                <div className="booking-header">
                  <h3>{booking.service?.name}</h3>
                  <span className={`status-badge ${normalizedStatus}`}>
                    {getStatusLabel(normalizedStatus)}
                  </span>
                </div>
                <div className="booking-details">
                  <p><strong>Service:</strong> {booking.service?.category}</p>
                  <p><strong>Location:</strong> {booking.service?.location}</p>
                  <p><strong>Start Time:</strong> {new Date(booking.startTime).toLocaleString()}</p>
                  <p><strong>End Time:</strong> {new Date(booking.endTime).toLocaleString()}</p>
                  {booking.notes && <p><strong>Notes:</strong> {booking.notes}</p>}
                </div>
                {booking.status === 'pending' && (
                  <button 
                    className="cancel-btn"
                    onClick={() => handleCancelBooking(booking.id)}
                  >
                    Cancel Booking
                  </button>
                )}
              </div>
            )})
          ) : (
            <div className="no-bookings">
              <p>No bookings found in this category</p>
            </div>
          )}
        </div>
      )}
    </div>
  );
}

export default Bookings;
