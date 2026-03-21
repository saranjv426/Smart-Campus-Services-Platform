import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../styles/StaffDashboard.css';

const StaffDashboard = () => {
  const navigate = useNavigate();
  const [bookings, setBookings] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [selectedBooking, setSelectedBooking] = useState(null);
  const [notes, setNotes] = useState('');
  const [showModal, setShowModal] = useState(false);
  const [actionType, setActionType] = useState('');
  const [staffInfo, setStaffInfo] = useState(null);

  useEffect(() => {
    // Get staff info from localStorage
    const token = localStorage.getItem('token');
    const user = localStorage.getItem('user');
    
    if (!token || !user) {
      navigate('/login');
      return;
    }

    const userData = JSON.parse(user);
    if (userData.role !== 'staff') {
      navigate('/');
      return;
    }

    setStaffInfo(userData);
    fetchPendingBookings(userData.id);
  }, [navigate]);

  const fetchPendingBookings = async (staffId) => {
    try {
      setLoading(true);
      const token = localStorage.getItem('token');
      const response = await axios.get(
        `http://localhost:8080/api/approval/staff/${staffId}/pending`,
        {
          headers: { Authorization: `Bearer ${token}` }
        }
      );
      // Handle both array and {data: array} response formats
      const bookingsData = Array.isArray(response.data) ? response.data : (response.data.data || []);
      setBookings(bookingsData);
      setError('');
    } catch (err) {
      const errorMsg = err.response?.data?.error || 'Failed to load pending bookings';
      setError(errorMsg);
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleApprove = (booking) => {
    setSelectedBooking(booking);
    setActionType('approve');
    setNotes('');
    setShowModal(true);
  };

  const handleReject = (booking) => {
    setSelectedBooking(booking);
    setActionType('reject');
    setNotes('');
    setShowModal(true);
  };

  const submitAction = async () => {
    try {
      const token = localStorage.getItem('token');
      const endpoint = actionType === 'approve' 
        ? `/api/approval/bookings/${selectedBooking.id}/approve`
        : `/api/approval/bookings/${selectedBooking.id}/reject`;

      await axios.put(
        `http://localhost:8080${endpoint}`,
        { 
          status: actionType === 'approve' ? 'approved' : 'rejected',
          approvalNotes: notes,
          staffId: staffInfo.id
        },
        {
          headers: { Authorization: `Bearer ${token}` }
        }
      );

      setShowModal(false);
      setSelectedBooking(null);
      fetchPendingBookings(staffInfo.id);
    } catch (err) {
      setError(`Failed to ${actionType} booking`);
      console.error(err);
    }
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  };

  return (
    <div className="staff-dashboard">
      <div className="dashboard-header">
        <div>
          <h1>Staff Dashboard</h1>
          <p className="welcome-text">Welcome, {staffInfo?.firstName} {staffInfo?.lastName}</p>
        </div>
        <button 
          className="logout-btn"
          onClick={() => {
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            navigate('/login');
          }}
        >
          Logout
        </button>
      </div>

      {error && <div className="error-message">{error}</div>}

      <div className="bookings-section">
        <div className="section-header">
          <h2>Pending Approvals</h2>
          <span className="badge">{bookings.length}</span>
        </div>

        {loading ? (
          <div className="loading">Loading pending bookings...</div>
        ) : bookings.length === 0 ? (
          <div className="empty-state">
            <p>No pending bookings</p>
            <p className="secondary">All bookings have been reviewed</p>
          </div>
        ) : (
          <div className="bookings-grid">
            {bookings.map((booking) => (
              <div key={booking.id} className="booking-card">
                <div className="booking-header">
                  <h3 className="service-name">{booking.service?.name}</h3>
                  <span className="status-badge pending">Pending</span>
                </div>

                <div className="booking-details">
                  <div className="detail-item">
                    <span className="label">Student Name:</span>
                    <span className="value">{booking.user?.firstName} {booking.user?.lastName}</span>
                  </div>
                  <div className="detail-item">
                    <span className="label">Email:</span>
                    <span className="value">{booking.user?.email}</span>
                  </div>
                  <div className="detail-item">
                    <span className="label">Phone:</span>
                    <span className="value">{booking.user?.phone || 'N/A'}</span>
                  </div>
                  <div className="detail-item">
                    <span className="label">Start Time:</span>
                    <span className="value">{formatDate(booking.startTime)}</span>
                  </div>
                  <div className="detail-item">
                    <span className="label">End Time:</span>
                    <span className="value">{formatDate(booking.endTime)}</span>
                  </div>
                  <div className="detail-item">
                    <span className="label">Reason:</span>
                    <span className="value">{booking.notes || 'N/A'}</span>
                  </div>
                  <div className="detail-item">
                    <span className="label">Requested:</span>
                    <span className="value">{formatDate(booking.createdAt)}</span>
                  </div>
                </div>

                <div className="booking-actions">
                  <button 
                    className="btn-reject"
                    onClick={() => handleReject(booking)}
                  >
                    Reject
                  </button>
                  <button 
                    className="btn-approve"
                    onClick={() => handleApprove(booking)}
                  >
                    Approve
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Modal for approval/rejection */}
      {showModal && selectedBooking && (
        <div className="modal-overlay" onClick={() => setShowModal(false)}>
          <div className="modal" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header">
              <h3>
                {actionType === 'approve' ? 'Approve' : 'Reject'} Booking
              </h3>
              <button 
                className="close-btn"
                onClick={() => setShowModal(false)}
              >
                ✕
              </button>
            </div>

            <div className="modal-body">
              <p className="booking-info">
                <strong>{selectedBooking.service?.name}</strong> - {selectedBooking.user?.firstName} {selectedBooking.user?.lastName}
              </p>
              
              <div className="form-group">
                <label>Additional Notes (Optional)</label>
                <textarea
                  value={notes}
                  onChange={(e) => setNotes(e.target.value)}
                  placeholder="Add any notes about this approval/rejection..."
                  rows="4"
                />
              </div>

              <div className="modal-actions">
                <button 
                  className="btn-cancel"
                  onClick={() => setShowModal(false)}
                >
                  Cancel
                </button>
                <button 
                  className={actionType === 'approve' ? 'btn-approve' : 'btn-reject'}
                  onClick={submitAction}
                >
                  {actionType === 'approve' ? 'Approve Booking' : 'Reject Booking'}
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default StaffDashboard;
