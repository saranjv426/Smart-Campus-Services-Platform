import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import BookingActionModal from '../components/BookingActionModal';
import { approvalAPI } from '../services/approvalApi';
import '../styles/StaffDashboard.css';

const StaffDashboard = () => {
  const navigate = useNavigate();
  const [bookings, setBookings] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [staffInfo, setStaffInfo] = useState(null);

  const [selectedBooking, setSelectedBooking] = useState(null);
  const [notes, setNotes] = useState('');
  const [showModal, setShowModal] = useState(false);
  const [actionType, setActionType] = useState('');
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    const user = localStorage.getItem('user');

    if (!token || !user) {
      navigate('/login');
      return;
    }

    try {
      const userData = JSON.parse(user);

      if (userData.role !== 'staff') {
        navigate('/');
        return;
      }

      setStaffInfo(userData);
      fetchPendingBookings(userData.id);
    } catch (err) {
      console.error('Failed to parse stored user data:', err);
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      navigate('/login');
    }
  }, [navigate]);

  const fetchPendingBookings = async (staffId) => {
    try {
      setLoading(true);
      const bookingsData = await approvalAPI.getPendingBookingsByStaff(staffId);
      setBookings(bookingsData);
      setError('');
    } catch (err) {
      console.error('Failed to load pending bookings:', err);
      const errorMsg =
        err.response?.data?.error || 'Failed to load pending bookings';
      setError(errorMsg);
      setBookings([]);
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    navigate('/login');
  };

  const formatDate = (dateString) => {
    if (!dateString) return 'N/A';

    const date = new Date(dateString);
    if (Number.isNaN(date.getTime())) return 'N/A';

    return date.toLocaleString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  const getServiceName = (booking) => {
    return booking.service?.name || booking.serviceName || 'Service Not Available';
  };

  const getStudentName = (booking) => {
    if (booking.user?.firstName || booking.user?.lastName) {
      return `${booking.user?.firstName || ''} ${booking.user?.lastName || ''}`.trim();
    }
    return booking.studentName || 'N/A';
  };

  const getStudentEmail = (booking) => {
    return booking.user?.email || booking.email || 'N/A';
  };

  const getStudentPhone = (booking) => {
    return booking.user?.phone || booking.phone || 'N/A';
  };

  const openActionModal = (booking, type) => {
    setSelectedBooking(booking);
    setActionType(type);
    setNotes('');
    setShowModal(true);
    setError('');
  };

  const closeModal = () => {
    if (submitting) return;
    setShowModal(false);
    setSelectedBooking(null);
    setNotes('');
    setActionType('');
  };

  const submitAction = async () => {
    if (!selectedBooking || !staffInfo) return;

    try {
      setSubmitting(true);
      setError('');

      const payload = {
        status: actionType === 'approve' ? 'approved' : 'rejected',
        approvalNotes: notes,
        staffId: staffInfo.id,
      };

      if (actionType === 'approve') {
        await approvalAPI.approveBooking(selectedBooking.id, payload);
      } else {
        await approvalAPI.rejectBooking(selectedBooking.id, payload);
      }

      closeModal();
      await fetchPendingBookings(staffInfo.id);
    } catch (err) {
      console.error(`Failed to ${actionType} booking:`, err);
      const errorMsg =
        err.response?.data?.error || `Failed to ${actionType} booking`;
      setError(errorMsg);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="staff-dashboard">
      <div className="dashboard-header">
        <div>
          <h1>Staff Dashboard</h1>
          <p className="welcome-text">
            Welcome, {staffInfo?.firstName} {staffInfo?.lastName}
          </p>
        </div>

        <button className="logout-btn" onClick={handleLogout}>
          Logout
        </button>
      </div>

      {error && <div className="error-message">{error}</div>}

      <div className="bookings-section">
        <div className="section-header">
          <h2>Pending Bookings</h2>
          <span className="badge">{bookings.length}</span>
        </div>

        {loading ? (
          <div className="loading">Loading pending bookings...</div>
        ) : bookings.length === 0 ? (
          <div className="empty-state">
            <p>No pending bookings</p>
            <p className="secondary">There are currently no bookings awaiting review.</p>
          </div>
        ) : (
          <div className="bookings-grid">
            {bookings.map((booking) => (
              <div key={booking.id} className="booking-card">
                <div className="booking-header">
                  <h3 className="service-name">{getServiceName(booking)}</h3>
                  <span className="status-badge pending">Pending</span>
                </div>

                <div className="booking-details">
                  <div className="detail-item">
                    <span className="label">Student Name:</span>
                    <span className="value">{getStudentName(booking)}</span>
                  </div>

                  <div className="detail-item">
                    <span className="label">Email:</span>
                    <span className="value">{getStudentEmail(booking)}</span>
                  </div>

                  <div className="detail-item">
                    <span className="label">Phone:</span>
                    <span className="value">{getStudentPhone(booking)}</span>
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
                    onClick={() => openActionModal(booking, 'reject')}
                  >
                    Reject
                  </button>

                  <button
                    className="btn-approve"
                    onClick={() => openActionModal(booking, 'approve')}
                  >
                    Approve
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      <BookingActionModal
        show={showModal}
        actionType={actionType}
        selectedBooking={selectedBooking}
        notes={notes}
        onNotesChange={setNotes}
        onClose={closeModal}
        onSubmit={submitAction}
        submitting={submitting}
        getServiceName={getServiceName}
        getStudentName={getStudentName}
      />
    </div>
  );
};

export default StaffDashboard;
