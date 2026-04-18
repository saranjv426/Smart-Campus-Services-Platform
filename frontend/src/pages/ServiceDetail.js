import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { serviceAPI, reviewAPI, bookingAPI } from '../services/api';
import '../styles/ServiceDetail.css';

function ServiceDetail() {
  const { id } = useParams();
  const [service, setService] = useState(null);
  const [reviews, setReviews] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showBookingForm, setShowBookingForm] = useState(false);
  const [bookingData, setBookingData] = useState({
    startTime: '',
    endTime: '',
    notes: '',
  });
  const [newReview, setNewReview] = useState({
    rating: 5,
    comment: '',
  });

  useEffect(() => {
    fetchServiceDetails();
  }, [id]);

  // Fetch service details and associated reviews from backend APIs
  const fetchServiceDetails = async () => {
    try {
      setLoading(true);
      const serviceResponse = await serviceAPI.getService(id);
      const reviewsResponse = await reviewAPI.getServiceReviews(id);
      
      setService(serviceResponse.data);
      setReviews(reviewsResponse.data || []);
      setError(null);
    } catch (err) {
      setError('Failed to load service details');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  // Handles booking form submission including validation and API integration
  const handleBookingSubmit = async (e) => {
    e.preventDefault();
    const currentUser = JSON.parse(localStorage.getItem('user'));
    
    if (!currentUser) {
      alert('Please login to make a booking');
      return;
    }

    // Validate start and end times
    if (!bookingData.startTime || !bookingData.endTime) {
      alert('Please select both start and end times');
      return;
    }

    // Convert to ISO format for backend
    const startTime = new Date(bookingData.startTime).toISOString();
    const endTime = new Date(bookingData.endTime).toISOString();

    try {
      await bookingAPI.createBooking({
        userId: currentUser.id,
        serviceId: id,
        startTime: startTime,
        endTime: endTime,
        notes: bookingData.notes,
      });
      alert('Booking created successfully!');
      setShowBookingForm(false);
      setBookingData({ startTime: '', endTime: '', notes: '' });
    } catch (err) {
      console.error('Booking error:', err);
      alert('Failed to create booking: ' + (err.response?.data?.error || err.message));
    }
  };

  // Handles review submission and sends user rating & comment to backend
  const handleReviewSubmit = async (e) => {
    e.preventDefault();
    const currentUser = JSON.parse(localStorage.getItem('user'));
    
    if (!currentUser) {
      alert('Please login to leave a review');
      return;
    }

    try {
      await reviewAPI.createReview({
        userId: currentUser.id,
        serviceId: id,
        rating: parseInt(newReview.rating),
        comment: newReview.comment,
      });
      alert('Review submitted successfully!');
      setNewReview({ rating: 5, comment: '' });
      fetchServiceDetails();
    } catch (err) {
      alert('Failed to submit review');
    }
  };

  if (loading) return <div className="loading">Loading...</div>;
  if (error) return <div className="error">{error}</div>;
  if (!service) return <div className="not-found">Service not found</div>;

  // Calculate average rating from all reviews for display
  const averageRating = reviews.length > 0
    ? (reviews.reduce((sum, r) => sum + r.rating, 0) / reviews.length).toFixed(1)
    : 'No ratings';

  return (
    <div className="service-detail">
      <div className="service-header">
        <img src={service.imageUrl} alt={service.name} className="service-image" />
        <div className="service-info">
          <h1>{service.name}</h1>
          <p className="category">Category: {service.category}</p>
          <p className="description">{service.description}</p>
          
          <div className="service-meta">
            <div className="meta-item">
              <span className="label">Location:</span>
              <span>{service.location}</span>
            </div>
            <div className="meta-item">
              <span className="label">Phone:</span>
              <span>{service.phone}</span>
            </div>
            <div className="meta-item">
              <span className="label">Email:</span>
              <span>{service.email}</span>
            </div>
            <div className="meta-item">
              <span className="label">Hours:</span>
              <span>{service.hours}</span>
            </div>
            <div className="meta-item">
              <span className="label">Rating:</span>
              <span className="rating">⭐ {averageRating}</span>
            </div>
          </div>

          <button 
            className="booking-btn"
            onClick={() => setShowBookingForm(!showBookingForm)}
          >
            {showBookingForm ? 'Cancel' : 'Book This Service'}
          </button>
        </div>
      </div>

      {showBookingForm && (
        <div className="booking-form-container">
          <h2>Create a Booking</h2>
          <form onSubmit={handleBookingSubmit} className="booking-form">
            <div className="form-group">
              <label htmlFor="startTime">Start Time</label>
              <input
                type="datetime-local"
                id="startTime"
                value={bookingData.startTime}
                onChange={(e) => setBookingData({ ...bookingData, startTime: e.target.value })}
                required
              />
            </div>
            <div className="form-group">
              <label htmlFor="endTime">End Time</label>
              <input
                type="datetime-local"
                id="endTime"
                value={bookingData.endTime}
                onChange={(e) => setBookingData({ ...bookingData, endTime: e.target.value })}
                required
              />
            </div>
            <div className="form-group">
              <label htmlFor="notes">Notes (Optional)</label>
              <textarea
                id="notes"
                value={bookingData.notes}
                onChange={(e) => setBookingData({ ...bookingData, notes: e.target.value })}
                placeholder="Add any special requests..."
              />
            </div>
            <button type="submit" className="submit-btn">Confirm Booking</button>
          </form>
        </div>
      )}

      
      <div className="reviews-section">
        <h2>Reviews ({reviews.length})</h2>
        
        <div className="review-form-container">
          <h3>Leave a Review</h3>
          <form onSubmit={handleReviewSubmit} className="review-form">
            <div className="form-group">
              <label htmlFor="rating">Rating</label>
              <select
                id="rating"
                value={newReview.rating}
                onChange={(e) => setNewReview({ ...newReview, rating: e.target.value })}
              >
                <option value="5">⭐⭐⭐⭐⭐ Excellent</option>
                <option value="4">⭐⭐⭐⭐ Good</option>
                <option value="3">⭐⭐⭐ Average</option>
                <option value="2">⭐⭐ Poor</option>
                <option value="1">⭐ Terrible</option>
              </select>
            </div>
            <div className="form-group">
              <label htmlFor="comment">Comment</label>
              <textarea
                id="comment"
                value={newReview.comment}
                onChange={(e) => setNewReview({ ...newReview, comment: e.target.value })}
                placeholder="Share your experience..."
                required
              />
            </div>
            <button type="submit" className="submit-btn">Submit Review</button>
          </form>
        </div>

        <div className="reviews-list">
          {reviews.length > 0 ? (
            reviews.map(review => (
              <div key={review.id} className="review-item">
                <div className="review-header">
                  <span className="reviewer-name">{review.user?.firstName} {review.user?.lastName}</span>
                  <span className="review-rating">{'⭐'.repeat(review.rating)}</span>
                </div>
                <p className="review-comment">{review.comment}</p>
                <small className="review-date">
                  {new Date(review.createdAt).toLocaleDateString()}
                </small>
              </div>
            ))
          ) : (
            <p className="no-reviews">No reviews yet. Be the first to review!</p>
          )}
        </div>
      </div>
    </div>
  );
}

export default ServiceDetail;
