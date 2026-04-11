import axios from 'axios';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

const getAuthHeaders = () => {
  const token = localStorage.getItem('token');
  return {
    headers: {
      Authorization: `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
  };
};

export const approvalAPI = {
  getPendingBookingsByStaff: async (staffId) => {
    const response = await axios.get(
      `${API_BASE_URL}/approval/staff/${staffId}/pending`,
      getAuthHeaders()
    );

    return Array.isArray(response.data) ? response.data : response.data?.data || [];
  },

  approveBooking: async (bookingId, payload) => {
    return axios.put(
      `${API_BASE_URL}/approval/bookings/${bookingId}/approve`,
      payload,
      getAuthHeaders()
    );
  },

  rejectBooking: async (bookingId, payload) => {
    return axios.put(
      `${API_BASE_URL}/approval/bookings/${bookingId}/reject`,
      payload,
      getAuthHeaders()
    );
  },
};

export default approvalAPI;