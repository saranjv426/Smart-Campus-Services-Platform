
// Issue #19: User Profile Page with Booking History
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { userAPI } from '../services/api';
import '../styles/Profile.css';

function Profile() {
  const { id } = useParams();
  const [user, setUser] = useState(null);
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchUserProfile();
  }, [id]);

  const fetchUserProfile = async () => {
    try {
      const response = await userAPI.getProfile(id);
      setUser(response.data);
      setFormData(response.data);
      setError(null);
    } catch (err) {
      setError('Failed to load profile');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSaveProfile = async () => {
    try {
      await userAPI.updateUser(id, formData);
      setUser(formData);
      setIsEditing(false);
      alert('Profile updated successfully!');
    } catch (err) {
      alert('Failed to update profile');
    }
  };

  if (loading) return <div className="loading">Loading profile...</div>;
  if (error) return <div className="error">{error}</div>;
  if (!user) return <div className="not-found">User not found</div>;

  return (
    <div className="profile-page">
      <div className="profile-container">
        <div className="profile-header">
          <img 
            src={user.avatarUrl || 'https://via.placeholder.com/150?text=' + user.firstName}
            alt={user.firstName}
            className="profile-avatar"
          />
          <div className="profile-intro">
            {!isEditing ? (
              <>
                <h1>{user.firstName} {user.lastName}</h1>
                <p className="role">{user.role.toUpperCase()}</p>
                <p className="email">{user.email}</p>
              </>
            ) : (
              <>
                <input
                  type="text"
                  name="firstName"
                  value={formData.firstName}
                  onChange={handleChange}
                  placeholder="First Name"
                />
                <input
                  type="text"
                  name="lastName"
                  value={formData.lastName}
                  onChange={handleChange}
                  placeholder="Last Name"
                />
              </>
            )}
          </div>
          {!isEditing ? (
            <button 
              className="edit-btn"
              onClick={() => setIsEditing(true)}
            >
              Edit Profile
            </button>
          ) : (
            <div className="edit-actions">
              <button className="save-btn" onClick={handleSaveProfile}>Save</button>
              <button className="cancel-btn" onClick={() => setIsEditing(false)}>Cancel</button>
            </div>
          )}
        </div>

        <div className="profile-content">
          <div className="profile-section">
            <h2>Contact Information</h2>
            {!isEditing ? (
              <div className="info-list">
                <div className="info-item">
                  <span className="label">Phone:</span>
                  <span className="value">{user.phone}</span>
                </div>
                <div className="info-item">
                  <span className="label">Email:</span>
                  <span className="value">{user.email}</span>
                </div>
                <div className="info-item">
                  <span className="label">Department:</span>
                  <span className="value">{user.department || 'Not specified'}</span>
                </div>
              </div>
            ) : (
              <div className="edit-form">
                <div className="form-group">
                  <label>Phone</label>
                  <input
                    type="tel"
                    name="phone"
                    value={formData.phone}
                    onChange={handleChange}
                  />
                </div>
                <div className="form-group">
                  <label>Department</label>
                  <input
                    type="text"
                    name="department"
                    value={formData.department || ''}
                    onChange={handleChange}
                  />
                </div>
              </div>
            )}
          </div>

          <div className="profile-section">
            <h2>Bio</h2>
            {!isEditing ? (
              <p className="bio">{user.bio || 'No bio added yet'}</p>
            ) : (
              <textarea
                name="bio"
                value={formData.bio || ''}
                onChange={handleChange}
                placeholder="Tell us about yourself..."
              />
            )}
          </div>

          {user.bookings && user.bookings.length > 0 && (
            <div className="profile-section">
              <h2>Recent Bookings</h2>
              <div className="bookings-summary">
                {user.bookings.slice(0, 5).map(booking => (
                  <div key={booking.id} className="booking-summary-item">
                    <span className="service-name">{booking.service?.name}</span>
                    <span className="booking-status">{booking.status}</span>
                  </div>
                ))}
              </div>
            </div>
          )}

          {user.reviews && user.reviews.length > 0 && (
            <div className="profile-section">
              <h2>My Reviews</h2>
              <div className="reviews-summary">
                <p>Total Reviews: {user.reviews.length}</p>
                <p>Average Rating: {(user.reviews.reduce((sum, r) => sum + r.rating, 0) / user.reviews.length).toFixed(1)} ⭐</p>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default Profile;
