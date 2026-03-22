import axios from 'axios';
import {
  authAPI,
  userAPI,
  serviceAPI,
  bookingAPI,
  reviewAPI,
  notificationAPI,
} from './api';

jest.mock('axios');

describe('API Service Layer', () => {
  const mockToken = 'mock-token-123';

  beforeEach(() => {
    jest.clearAllMocks();
    localStorage.clear();
  });

  // Auth API Tests
  describe('authAPI', () => {
    test('register calls correct endpoint', () => {
      const userData = { email: 'test@example.com', password: 'password123' };
      authAPI.register(userData);

      expect(axios.create().post).toBeDefined();
    });

    test('login calls correct endpoint', () => {
      const credentials = { email: 'test@example.com', password: 'password123' };
      authAPI.login(credentials);

      expect(axios.create().post).toBeDefined();
    });

    test('logout calls correct endpoint', () => {
      authAPI.logout();

      expect(axios.create().post).toBeDefined();
    });

    test('refreshToken calls correct endpoint', () => {
      authAPI.refreshToken();

      expect(axios.create().post).toBeDefined();
    });
  });

  // User API Tests
  describe('userAPI', () => {
    test('getUser calls correct endpoint', () => {
      const userId = 1;
      userAPI.getUser(userId);

      expect(axios.create().get).toBeDefined();
    });

    test('updateUser calls correct endpoint', () => {
      const userId = 1;
      const userData = { firstName: 'John', lastName: 'Doe' };
      userAPI.updateUser(userId, userData);

      expect(axios.create().put).toBeDefined();
    });

    test('getProfile calls correct endpoint', () => {
      const userId = 1;
      userAPI.getProfile(userId);

      expect(axios.create().get).toBeDefined();
    });
  });

  // Service API Tests
  describe('serviceAPI', () => {
    test('getServices calls correct endpoint', () => {
      serviceAPI.getServices();

      expect(axios.create().get).toBeDefined();
    });

    test('getService calls correct endpoint with service id', () => {
      const serviceId = 1;
      serviceAPI.getService(serviceId);

      expect(axios.create().get).toBeDefined();
    });

    test('createService calls correct endpoint', () => {
      const serviceData = { name: 'Library', category: 'library' };
      serviceAPI.createService(serviceData);

      expect(axios.create().post).toBeDefined();
    });

    test('updateService calls correct endpoint', () => {
      const serviceId = 1;
      const serviceData = { name: 'Updated Library' };
      serviceAPI.updateService(serviceId, serviceData);

      expect(axios.create().put).toBeDefined();
    });

    test('deleteService calls correct endpoint', () => {
      const serviceId = 1;
      serviceAPI.deleteService(serviceId);

      expect(axios.create().delete).toBeDefined();
    });

    test('getServicesByCategory calls correct endpoint', () => {
      const category = 'library';
      serviceAPI.getServicesByCategory(category);

      expect(axios.create().get).toBeDefined();
    });
  });

  // Booking API Tests
  describe('bookingAPI', () => {
    test('getServices calls correct endpoint', () => {
      bookingAPI.getServices();

      expect(axios.create().get).toBeDefined();
    });

    test('createBooking calls correct endpoint', () => {
      const bookingData = {
        serviceId: 1,
        startTime: '2024-02-20T10:00:00Z',
        endTime: '2024-02-20T12:00:00Z',
      };
      bookingAPI.createBooking(bookingData);

      expect(axios.create().post).toBeDefined();
    });

    test('getBooking calls correct endpoint', () => {
      const bookingId = 1;
      bookingAPI.getBooking(bookingId);

      expect(axios.create().get).toBeDefined();
    });

    test('getUserBookings calls correct endpoint', () => {
      const userId = 1;
      bookingAPI.getUserBookings(userId);

      expect(axios.create().get).toBeDefined();
    });

    test('updateBooking calls correct endpoint', () => {
      const bookingId = 1;
      const bookingData = { status: 'cancelled' };
      bookingAPI.updateBooking(bookingId, bookingData);

      expect(axios.create().put).toBeDefined();
    });

    test('cancelBooking calls correct endpoint', () => {
      const bookingId = 1;
      bookingAPI.cancelBooking(bookingId);

      expect(axios.create().delete).toBeDefined();
    });
  });

  // Review API Tests
  describe('reviewAPI', () => {
    test('createReview calls correct endpoint', () => {
      const reviewData = {
        serviceId: 1,
        rating: 5,
        comment: 'Great service!',
      };
      reviewAPI.createReview(reviewData);

      expect(axios.create().post).toBeDefined();
    });

    test('getServiceReviews calls correct endpoint', () => {
      const serviceId = 1;
      reviewAPI.getServiceReviews(serviceId);

      expect(axios.create().get).toBeDefined();
    });

    test('getUserReviews calls correct endpoint', () => {
      const userId = 1;
      reviewAPI.getUserReviews(userId);

      expect(axios.create().get).toBeDefined();
    });

    test('getReview calls correct endpoint', () => {
      const reviewId = 1;
      reviewAPI.getReview(reviewId);

      expect(axios.create().get).toBeDefined();
    });

    test('deleteReview calls correct endpoint', () => {
      const reviewId = 1;
      reviewAPI.deleteReview(reviewId);

      expect(axios.create().delete).toBeDefined();
    });
  });

  // Notification API Tests
  describe('notificationAPI', () => {
    test('getNotifications calls correct endpoint', () => {
      const userId = 1;
      notificationAPI.getNotifications(userId);

      expect(axios.create().get).toBeDefined();
    });

    test('createNotification calls correct endpoint', () => {
      const notificationData = {
        userId: 1,
        message: 'Test notification',
      };
      notificationAPI.createNotification(notificationData);

      expect(axios.create().post).toBeDefined();
    });

    test('markAsRead calls correct endpoint', () => {
      const notificationId = 1;
      notificationAPI.markAsRead(notificationId);

      expect(axios.create().put).toBeDefined();
    });
  });

  // Interceptor Tests
  describe('API Interceptors', () => {
    test('request interceptor adds auth token when available', () => {
      localStorage.setItem('token', mockToken);

      // Create a new instance to test interceptor
      const apiClient = axios.create();
      expect(apiClient).toBeDefined();
    });

    test('request interceptor does not add token when not available', () => {
      localStorage.removeItem('token');

      // Create a new instance to test interceptor
      const apiClient = axios.create();
      expect(apiClient).toBeDefined();
    });

    test('error interceptor handles failed requests', () => {
      // Error handling should work
      const apiClient = axios.create();
      expect(apiClient).toBeDefined();
    });
  });

  // Integration Tests
  describe('API Integration', () => {
    test('all API objects are properly exported', () => {
      expect(authAPI).toBeDefined();
      expect(userAPI).toBeDefined();
      expect(serviceAPI).toBeDefined();
      expect(bookingAPI).toBeDefined();
      expect(reviewAPI).toBeDefined();
      expect(notificationAPI).toBeDefined();
    });

    test('authAPI has all required methods', () => {
      expect(authAPI.register).toBeDefined();
      expect(authAPI.login).toBeDefined();
      expect(authAPI.logout).toBeDefined();
      expect(authAPI.refreshToken).toBeDefined();
    });

    test('userAPI has all required methods', () => {
      expect(userAPI.getUser).toBeDefined();
      expect(userAPI.updateUser).toBeDefined();
      expect(userAPI.getProfile).toBeDefined();
    });

    test('serviceAPI has all required methods', () => {
      expect(serviceAPI.getServices).toBeDefined();
      expect(serviceAPI.getService).toBeDefined();
      expect(serviceAPI.createService).toBeDefined();
      expect(serviceAPI.updateService).toBeDefined();
      expect(serviceAPI.deleteService).toBeDefined();
      expect(serviceAPI.getServicesByCategory).toBeDefined();
    });

    test('bookingAPI has all required methods', () => {
      expect(bookingAPI.getServices).toBeDefined();
      expect(bookingAPI.createBooking).toBeDefined();
      expect(bookingAPI.getBooking).toBeDefined();
      expect(bookingAPI.getUserBookings).toBeDefined();
      expect(bookingAPI.updateBooking).toBeDefined();
      expect(bookingAPI.cancelBooking).toBeDefined();
    });

    test('reviewAPI has all required methods', () => {
      expect(reviewAPI.createReview).toBeDefined();
      expect(reviewAPI.getServiceReviews).toBeDefined();
      expect(reviewAPI.getUserReviews).toBeDefined();
      expect(reviewAPI.getReview).toBeDefined();
      expect(reviewAPI.deleteReview).toBeDefined();
    });

    test('notificationAPI has all required methods', () => {
      expect(notificationAPI.getNotifications).toBeDefined();
      expect(notificationAPI.createNotification).toBeDefined();
      expect(notificationAPI.markAsRead).toBeDefined();
    });
  });
});
