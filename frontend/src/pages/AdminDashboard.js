import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import '../styles/AdminDashboard.css';

const AdminDashboard = () => {
  const navigate = useNavigate();
  
  // Tab state
  const [activeTab, setActiveTab] = useState('bookings'); // bookings or services
  
  // Bookings state
  const [bookings, setBookings] = useState([]);
  const [filteredBookings, setFilteredBookings] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [selectedBooking, setSelectedBooking] = useState(null);
  const [notes, setNotes] = useState('');
  const [showModal, setShowModal] = useState(false);
  const [actionType, setActionType] = useState('');
  const [adminInfo, setAdminInfo] = useState(null);
  const [filterStatus, setFilterStatus] = useState('pending');
  const [filterService, setFilterService] = useState('all');
  const [services, setServices] = useState([]);
  
  // Service management state
  const [showServiceModal, setShowServiceModal] = useState(false);
  const [serviceModalType, setServiceModalType] = useState('create'); // create or delete
  const [selectedService, setSelectedService] = useState(null);
  const [serviceForm, setServiceForm] = useState({
    name: '',
    description: '',
    category: 'library',
    location: '',
    phone: '',
    email: '',
    hours: '',
    imageUrl: ''
  });
  const [serviceError, setServiceError] = useState('');
  const [serviceLoading, setServiceLoading] = useState(false);

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

  // Service Management Functions
  const handleCreateServiceClick = () => {
    setServiceModalType('create');
    setServiceForm({
      name: '',
      description: '',
      category: 'library',
      location: '',
      phone: '',
      email: '',
      hours: '',
      imageUrl: ''
    });
    setServiceError('');
    setShowServiceModal(true);
  };

  const handleDeleteServiceClick = (service) => {
    setServiceModalType('delete');
    setSelectedService(service);
    setServiceError('');
    setShowServiceModal(true);
  };

  const handleServiceFormChange = (e) => {
    const { name, value } = e.target;
    setServiceForm(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const submitServiceAction = async () => {
    try {
      setServiceLoading(true);
      const token = localStorage.getItem('token');

      if (serviceModalType === 'create') {
        // Validate required fields
        if (!serviceForm.name || !serviceForm.description || !serviceForm.location) {
          setServiceError('Name, description, and location are required');
          setServiceLoading(false);
          return;
        }

        await axios.post(
          'http://localhost:8080/api/services',
          {
            name: serviceForm.name,
            description: serviceForm.description,
            category: serviceForm.category,
            location: serviceForm.location,
            phone: serviceForm.phone,
            email: serviceForm.email,
            hours: serviceForm.hours,
            imageUrl: serviceForm.imageUrl
          },
          {
            headers: { Authorization: `Bearer ${token}` }
          }
        );

        setShowServiceModal(false);
        setServiceForm({
          name: '',
          description: '',
          category: 'library',
          location: '',
          phone: '',
          email: '',
          hours: '',
          imageUrl: ''
        });
        setError('Service created successfully!');
        setTimeout(() => setError(''), 3000);
        fetchServices();
      } else if (serviceModalType === 'delete') {
        // Delete service
        await axios.delete(
          `http://localhost:8080/api/services/${selectedService.id}`,
          {
            headers: { Authorization: `Bearer ${token}` }
          }
        );

        setShowServiceModal(false);
        setSelectedService(null);
        setError('Service deleted successfully!');
        setTimeout(() => setError(''), 3000);
        fetchServices();
      }
    } catch (err) {
      const errorMsg = err.response?.data?.error || err.message;
      setServiceError(`Failed to ${serviceModalType} service: ${errorMsg}`);
      console.error('Service error:', err);
    } finally {
      setServiceLoading(false);
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

      {/* Tab Navigation */}
      <div className="tab-navigation">
        <button 
          className={`tab-btn ${activeTab === 'bookings' ? 'active' : ''}`}
          onClick={() => setActiveTab('bookings')}
        >
          📋 Bookings Overview
        </button>
        <button 
          className={`tab-btn ${activeTab === 'services' ? 'active' : ''}`}
          onClick={() => setActiveTab('services')}
        >
          🏢 Manage Services
        </button>
      </div>

      {error && <div className="error-message">{error}</div>}

      {/* Bookings Tab */}
      {activeTab === 'bookings' && (
        <>
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
        </>
      )}

      {/* Services Management Tab */}
      {activeTab === 'services' && (
        <div className="services-section">
          <div className="section-header">
            <h2>Service Management</h2>
            <button 
              className="btn-create-service"
              onClick={handleCreateServiceClick}
            >
              + Create New Service
            </button>
          </div>

          {serviceError && <div className="error-message">{serviceError}</div>}

          {services.length === 0 ? (
            <div className="empty-state">
              <p>No services available</p>
              <p className="secondary">Create your first service to get started</p>
            </div>
          ) : (
            <div className="services-grid">
              {services.map((service) => (
                <div key={service.id} className="service-card">
                  <div className="service-image">
                    {service.imageUrl ? (
                      <img src={service.imageUrl} alt={service.name} />
                    ) : (
                      <div className="image-placeholder">📷</div>
                    )}
                  </div>
                  <div className="service-content">
                    <h3>{service.name}</h3>
                    <p className="category-badge">{service.category}</p>
                    <p className="description">{service.description}</p>
                    <div className="service-meta">
                      <p><strong>Location:</strong> {service.location}</p>
                      {service.phone && <p><strong>Phone:</strong> {service.phone}</p>}
                      {service.email && <p><strong>Email:</strong> {service.email}</p>}
                      {service.hours && <p><strong>Hours:</strong> {service.hours}</p>}
                    </div>
                  </div>
                  <div className="service-actions">
                    <button 
                      className="btn-delete-service"
                      onClick={() => handleDeleteServiceClick(service)}
                    >
                      Delete
                    </button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      )}

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

      {/* Service Modal for create/delete */}
      {showServiceModal && (
        <div className="modal-overlay" onClick={() => setShowServiceModal(false)}>
          <div className="modal service-modal" onClick={(e) => e.stopPropagation()}>
            <div className="modal-header">
              <h3>
                {serviceModalType === 'create' ? 'Create New Service' : 'Delete Service'}
              </h3>
              <button 
                className="close-btn"
                onClick={() => setShowServiceModal(false)}
              >
                ✕
              </button>
            </div>

            <div className="modal-body">
              {serviceModalType === 'create' ? (
                <>
                  {serviceError && <div className="form-error">{serviceError}</div>}
                  <div className="form-group">
                    <label>Service Name *</label>
                    <input
                      type="text"
                      name="name"
                      value={serviceForm.name}
                      onChange={handleServiceFormChange}
                      placeholder="e.g., Library Services"
                    />
                  </div>

                  <div className="form-group">
                    <label>Description *</label>
                    <textarea
                      name="description"
                      value={serviceForm.description}
                      onChange={handleServiceFormChange}
                      placeholder="Describe the service..."
                      rows="3"
                    />
                  </div>

                  <div className="form-row">
                    <div className="form-group">
                      <label>Category</label>
                      <select
                        name="category"
                        value={serviceForm.category}
                        onChange={handleServiceFormChange}
                      >
                        <option value="library">Library</option>
                        <option value="dining">Dining</option>
                        <option value="transportation">Transportation</option>
                        <option value="healthcare">Healthcare</option>
                        <option value="other">Other</option>
                      </select>
                    </div>

                    <div className="form-group">
                      <label>Location *</label>
                      <input
                        type="text"
                        name="location"
                        value={serviceForm.location}
                        onChange={handleServiceFormChange}
                        placeholder="Building or location"
                      />
                    </div>
                  </div>

                  <div className="form-row">
                    <div className="form-group">
                      <label>Phone</label>
                      <input
                        type="tel"
                        name="phone"
                        value={serviceForm.phone}
                        onChange={handleServiceFormChange}
                        placeholder="(XXX) XXX-XXXX"
                      />
                    </div>

                    <div className="form-group">
                      <label>Email</label>
                      <input
                        type="email"
                        name="email"
                        value={serviceForm.email}
                        onChange={handleServiceFormChange}
                        placeholder="service@university.edu"
                      />
                    </div>
                  </div>

                  <div className="form-group">
                    <label>Working Hours</label>
                    <input
                      type="text"
                      name="hours"
                      value={serviceForm.hours}
                      onChange={handleServiceFormChange}
                      placeholder="e.g., Mon-Fri 9AM-5PM, Sat 10AM-2PM"
                    />
                  </div>

                  <div className="form-group">
                    <label>Image URL</label>
                    <input
                      type="url"
                      name="imageUrl"
                      value={serviceForm.imageUrl}
                      onChange={handleServiceFormChange}
                      placeholder="https://example.com/image.jpg"
                    />
                  </div>
                </>
              ) : (
                <>
                  <div className="delete-confirmation">
                    <p className="warning-text">⚠️ Warning: This action cannot be undone!</p>
                    <p>Are you sure you want to delete this service?</p>
                    {selectedService && (
                      <div className="service-to-delete">
                        <strong>{selectedService.name}</strong>
                        <p>{selectedService.description}</p>
                      </div>
                    )}
                  </div>
                </>
              )}

              <div className="modal-actions">
                <button 
                  className="btn-cancel"
                  onClick={() => setShowServiceModal(false)}
                  disabled={serviceLoading}
                >
                  Cancel
                </button>
                <button 
                  className={serviceModalType === 'create' ? 'btn-create' : 'btn-delete'}
                  onClick={submitServiceAction}
                  disabled={serviceLoading}
                >
                  {serviceLoading ? 'Processing...' : (serviceModalType === 'create' ? 'Create Service' : 'Delete Service')}
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
