import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../styles/AdminDashboard.css';

const AdminDashboard = () => {
  const navigate = useNavigate();
  const [bookings, setBookings] = useState([]);
  const [filteredBookings, setFilteredBookings] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [selectedBooking, setSelectedBooking] = useState(null);
  const [notes, setNotes] = useState('');
  const [showModal, setShowModal] = useState(false);
  const [actionType, setActionType] = useState('');
  const [adminInfo, setAdminInfo] = useState(null);
  const [filterStatus, setFilterStatus] = useState('pending'); // pending, approved, rejected, all
  const [filterService, setFilterService] = useState('all');
  const [services, setServices] = useState([]);

  useEffect(() => {
    // Get admin info from localStorage
    const token = localStorage.getItem('token');
    const user = localStorage.getItem('user');
    
    if (!token || !user) {
      navigate('/login');
      return;
    }

    const userData = JSON.parse(user);
    if (userData.role !== 'admin') {
      navigate('/');
      return;
    }

    setAdminInfo(userData);
    fetchAllBookings(userData.id);
    fetchServices();
  }, [navigate]);

  const fetchServices = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/services');
      setServices(Array.isArray(response.data) ? response.data : []);
    } catch (err) {
      console.error('Failed to load services:', err);
    }
  };

  const fetchAllBookings = async (adminId) => {
    try {
      setLoading(true);
      const token = localStorage.getItem('token');
      const response = await axios.get(
        `http://localhost:8080/api/approval/admin/${adminId}/all`,
        {
          headers: { Authorization: `Bearer ${token}` }
        }
      );
      const bookingsData = Array.isArray(response.data) ? response.data : (response.data.data || []);
      setBookings(bookingsData);
      applyFilters(bookingsData, filterStatus, filterService);
      setError('');
    } catch (err) {
      setError('Failed to load bookings');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const normalizeStatus = (status) => {
    if (status === 'completed') {
      return 'approved';
    }
    return status;
  };

  const applyFilters = (bookingsData, status, service) => {
    let filtered = bookingsData;

    if (status !== 'all') {
      filtered = filtered.filter(b => normalizeStatus(b.status) === status);
    }

    if (service !== 'all') {
      // Filter by service category - match service names that contain the category
      const categoryMap = {
        'library': 'Library',
        'dining': 'Dining',
        'transport': 'Transportation',
        'healthcare': 'Health'
      };
      const selectedCategory = categoryMap[service];
      filtered = filtered.filter(b => b.service?.name.includes(selectedCategory));
    }

    setFilteredBookings(filtered);
  };

  const handleStatusChange = (newStatus) => {
    setFilterStatus(newStatus);
    applyFilters(bookings, newStatus, filterService);
  };

  const handleServiceChange = (newService) => {
    setFilterService(newService);
    applyFilters(bookings, filterStatus, newService);
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
        ? `/api/approval/admin/${adminInfo.id}/bookings/${selectedBooking.id}/approve`
        : `/api/approval/admin/${adminInfo.id}/bookings/${selectedBooking.id}/reject`;

      await axios.put(
        `http://localhost:8080${endpoint}`,
        { 
          status: actionType === 'approve' ? 'approved' : 'rejected',
          approvalNotes: notes 
        },
        {
          headers: { Authorization: `Bearer ${token}` }
        }
      );

      setShowModal(false);
      setSelectedBooking(null);
      setNotes('');
      setError('');
      setTimeout(() => fetchAllBookings(adminInfo.id), 500);
    } catch (err) {
      const errorMsg = err.response?.data?.error || err.message;
      setError(`Failed to ${actionType} booking: ${errorMsg}`);
      console.error('Approval error:', err);
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

  const getStatusColor = (status) => {
    const normalizedStatus = normalizeStatus(status);
    switch(normalizedStatus) {
      case 'pending': return '#ffc107';
      case 'approved': return '#28a745';
      case 'rejected': return '#dc3545';
      default: return '#6c757d';
    }
  };

  const stats = {
    total: bookings.length,
    pending: bookings.filter(b => normalizeStatus(b.status) === 'pending').length,
    approved: bookings.filter(b => normalizeStatus(b.status) === 'approved').length,
    rejected: bookings.filter(b => normalizeStatus(b.status) === 'rejected').length,
  };

  return (
    <div className="admin-dashboard">
      <div className="dashboard-header">
        <div>
          <h1>Admin Dashboard</h1>
          <p className="welcome-text">Welcome, {adminInfo?.firstName} {adminInfo?.lastName}</p>
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

      {/* Statistics Cards */}
      <div className="stats-container">
        <div className="stat-card total">
          <div className="stat-number">{stats.total}</div>
          <div className="stat-label">Total Bookings</div>
        </div>
        <div className="stat-card pending">
          <div className="stat-number">{stats.pending}</div>
          <div className="stat-label">Pending</div>
        </div>
        <div className="stat-card approved">
          <div className="stat-number">{stats.approved}</div>
          <div className="stat-label">Approved</div>
        </div>
        <div className="stat-card rejected">
          <div className="stat-number">{stats.rejected}</div>
          <div className="stat-label">Rejected</div>
        </div>
      </div>

      {/* Filters */}
      <div className="filters-section">
        <div className="filter-group">
          <label>Status:</label>
          <div className="filter-buttons">
            <button 
              className={`filter-btn ${filterStatus === 'pending' ? 'active' : ''}`}
              onClick={() => handleStatusChange('pending')}
            >
              Pending
            </button>
            <button 
              className={`filter-btn ${filterStatus === 'approved' ? 'active' : ''}`}
              onClick={() => handleStatusChange('approved')}
            >
              Approved
            </button>
            <button 
              className={`filter-btn ${filterStatus === 'rejected' ? 'active' : ''}`}
              onClick={() => handleStatusChange('rejected')}
            >
              Rejected
            </button>
            <button 
              className={`filter-btn ${filterStatus === 'all' ? 'active' : ''}`}
              onClick={() => handleStatusChange('all')}
            >
              All
            </button>
          </div>
        </div>

        <div className="filter-group">
          <label>Service:</label>
          <select 
            value={filterService}
            onChange={(e) => handleServiceChange(e.target.value)}
            className="filter-select"
          >
            <option value="all">All Services</option>
            <option value="library">Library Services</option>
            <option value="dining">Dining Services</option>
            <option value="transport">Transportation Services</option>
            <option value="healthcare">Healthcare Services</option>
          </select>
        </div>
      </div>

      {/* Bookings Table */}
      <div className="bookings-section">
        <div className="section-header">
          <h2>All Bookings</h2>
          <span className="badge">{filteredBookings.length}</span>
        </div>

        {loading ? (
          <div className="loading">Loading bookings...</div>
        ) : filteredBookings.length === 0 ? (
          <div className="empty-state">
            <p>No bookings found</p>
            <p className="secondary">Try adjusting your filters</p>
          </div>
        ) : (
          <div className="bookings-table-container">
            <table className="bookings-table">
              <thead>
                <tr>
                  <th>Service</th>
                  <th>Student</th>
                  <th>Email</th>
                  <th>Start Time</th>
                  <th>End Time</th>
                  <th>Status</th>
                  <th>Created</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {filteredBookings.map((booking) => (
                  <tr key={booking.id} className={`status-${normalizeStatus(booking.status)}`}>
                    <td className="service-name">{booking.service?.name}</td>
                    <td>{booking.user?.firstName} {booking.user?.lastName}</td>
                    <td>{booking.user?.email}</td>
                    <td className="time-cell">{formatDate(booking.startTime)}</td>
                    <td className="time-cell">{formatDate(booking.endTime)}</td>
                    <td>
                      <span 
                        className="status-badge"
                        style={{ backgroundColor: getStatusColor(booking.status) }}
                      >
                        {normalizeStatus(booking.status)?.charAt(0).toUpperCase() + normalizeStatus(booking.status)?.slice(1)}
                      </span>
                    </td>
                    <td className="time-cell">{formatDate(booking.createdAt)}</td>
                    <td className="actions-cell">
                      {booking.status === 'pending' && (
                        <>
                          <button 
                            className="action-btn approve"
                            onClick={() => handleApprove(booking)}
                            title="Approve booking"
                          >
                            ✓
                          </button>
                          <button 
                            className="action-btn reject"
                            onClick={() => handleReject(booking)}
                            title="Reject booking"
                          >
                            ✕
                          </button>
                        </>
                      )}
                      {booking.status !== 'pending' && (
                        <span className="action-text">
                          {normalizeStatus(booking.status) === 'approved' ? 'Approved' : 'Rejected'}
                        </span>
                      )}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
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

export default AdminDashboard;
