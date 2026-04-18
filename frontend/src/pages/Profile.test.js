import React from 'react';
import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import Profile from './Profile';
import { userAPI, reviewAPI } from '../services/api';
import { MemoryRouter, Route, Routes } from 'react-router-dom';

// Mock APIs
jest.mock('../services/api');

const mockUser = {
  id: 1,
  firstName: 'John',
  lastName: 'Doe',
  role: 'user',
  email: 'john@example.com',
  phone: '1234567890',
  department: 'IT',
  bio: 'Test bio',
  bookings: [
    { id: 1, service: { name: 'Haircut' }, status: 'Completed' }
  ]
};

const mockReviews = [
  {
    id: 1,
    service: { name: 'Haircut' },
    rating: 5,
    comment: 'Great service!',
    createdAt: new Date().toISOString()
  }
];

// Helper to render with router
const renderComponent = () => {
  return render(
    <MemoryRouter initialEntries={['/profile/1']}>
      <Routes>
        <Route path="/profile/:id" element={<Profile />} />
      </Routes>
    </MemoryRouter>
  );
};

describe('Profile Page', () => {

  beforeEach(() => {
    jest.clearAllMocks();
  });

  test('shows loading initially', () => {
    userAPI.getProfile.mockResolvedValueOnce({ data: mockUser });

    renderComponent();

    expect(screen.getByText(/loading profile/i)).toBeInTheDocument();
  });

  test('renders user profile data', async () => {
    userAPI.getProfile.mockResolvedValueOnce({ data: mockUser });
    reviewAPI.getUserReviews.mockResolvedValueOnce({ data: mockReviews });

    renderComponent();

    await waitFor(() => {
      expect(screen.getByText(/john doe/i)).toBeInTheDocument();
    });

    expect(screen.getByText(/john@example.com/i)).toBeInTheDocument();
    expect(screen.getByText(/user/i)).toBeInTheDocument();
  });

  test('renders booking history', async () => {
    userAPI.getProfile.mockResolvedValueOnce({ data: mockUser });
    reviewAPI.getUserReviews.mockResolvedValueOnce({ data: [] });

    renderComponent();

    await waitFor(() => {
      expect(screen.getByText(/recent bookings/i)).toBeInTheDocument();
    });

    expect(screen.getByText(/haircut/i)).toBeInTheDocument();
    expect(screen.getByText(/completed/i)).toBeInTheDocument();
  });

  test('renders reviews', async () => {
    userAPI.getProfile.mockResolvedValueOnce({ data: mockUser });
    reviewAPI.getUserReviews.mockResolvedValueOnce({ data: mockReviews });

    renderComponent();

    await waitFor(() => {
      expect(screen.getByText(/my reviews/i)).toBeInTheDocument();
    });

    expect(screen.getByText(/great service/i)).toBeInTheDocument();
  });

  test('handles empty reviews', async () => {
    userAPI.getProfile.mockResolvedValueOnce({ data: mockUser });
    reviewAPI.getUserReviews.mockResolvedValueOnce({ data: [] });

    renderComponent();

    await waitFor(() => {
      expect(screen.getByText(/you haven't written any reviews yet/i)).toBeInTheDocument();
    });
  });

  test('handles API error', async () => {
    userAPI.getProfile.mockRejectedValueOnce(new Error('API error'));

    renderComponent();

    await waitFor(() => {
      expect(screen.getByText(/failed to load profile/i)).toBeInTheDocument();
    });
  });

  test('edit profile button works', async () => {
    userAPI.getProfile.mockResolvedValueOnce({ data: mockUser });
    reviewAPI.getUserReviews.mockResolvedValueOnce({ data: [] });

    renderComponent();

    await waitFor(() => {
      expect(screen.getByText(/edit profile/i)).toBeInTheDocument();
    });

    fireEvent.click(screen.getByText(/edit profile/i));

    expect(screen.getByPlaceholderText(/first name/i)).toBeInTheDocument();
  });

});