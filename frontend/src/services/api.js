import axios from 'axios';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add token to requests if available
apiClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

export const authAPI = {
  register: (userData) => apiClient.post('/auth/register', userData),
  login: (credentials) => apiClient.post('/auth/login', credentials),
  logout: () => apiClient.post('/auth/logout'),
  refreshToken: () => apiClient.post('/auth/refresh'),
};

export const userAPI = {
  getUser: (id) => apiClient.get(`/users/${id}`),
  updateUser: (id, userData) => apiClient.put(`/users/${id}`, userData),
  getProfile: (id) => apiClient.get(`/users/${id}/profile`),
};

export const serviceAPI = {
  getServices: () => apiClient.get('/services'),
  getService: (id) => apiClient.get(`/services/${id}`),
  createService: (serviceData) => apiClient.post('/services', serviceData),
  updateService: (id, serviceData) => apiClient.put(`/services/${id}`, serviceData),
  deleteService: (id) => apiClient.delete(`/services/${id}`),
  getServicesByCategory: (category) => apiClient.get(`/services/category/${category}`),
};

export const bookingAPI = {
  getServices: () => apiClient.get('/services'),
  createBooking: (bookingData) => apiClient.post('/bookings', bookingData),
  getBooking: (id) => apiClient.get(`/bookings/${id}`),
  getUserBookings: (userId) => apiClient.get(`/bookings/user/${userId}`),
  updateBooking: (id, bookingData) => apiClient.put(`/bookings/${id}`, bookingData),
  cancelBooking: (id) => apiClient.delete(`/bookings/${id}`),
};

export const reviewAPI = {
  createReview: (reviewData) => apiClient.post('/reviews', reviewData),
  getServiceReviews: (serviceId) => apiClient.get(`/reviews/service/${serviceId}`),
  getReview: (id) => apiClient.get(`/reviews/${id}`),
  deleteReview: (id) => apiClient.delete(`/reviews/${id}`),
};

export const notificationAPI = {
  getNotifications: (userId) => apiClient.get(`/notifications/${userId}`),
  createNotification: (notificationData) => apiClient.post('/notifications', notificationData),
  markAsRead: (id) => apiClient.put(`/notifications/${id}/read`),
};

export default apiClient;
