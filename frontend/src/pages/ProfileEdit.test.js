import React from 'react';
import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import Profile from './Profile';
import { userAPI, reviewAPI } from '../services/api';
import { MemoryRouter, Routes, Route } from 'react-router-dom';

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
  bookings: []
};

const renderComponent = () => {
  return render(
    <MemoryRouter initialEntries={['/profile/1']}>
      <Routes>
        <Route path="/profile/:id" element={<Profile />} />
      </Routes>
    </MemoryRouter>
  );
};

describe('Profile Edit Functionality', () => {

  beforeEach(() => {
    jest.clearAllMocks();
  });

  test('allows user to enter edit mode', async () => {
    userAPI.getProfile.mockResolvedValueOnce({ data: mockUser });
    reviewAPI.getUserReviews.mockResolvedValueOnce({ data: [] });

    renderComponent();

    await waitFor(() => {
      expect(screen.getByText(/edit profile/i)).toBeInTheDocument();
    });

    fireEvent.click(screen.getByText(/edit profile/i));

    expect(screen.getByDisplayValue(/john/i)).toBeInTheDocument();
  });

  test('updates form input values', async () => {
    userAPI.getProfile.mockResolvedValueOnce({ data: mockUser });
    reviewAPI.getUserReviews.mockResolvedValueOnce({ data: [] });

    renderComponent();

    await waitFor(() => {
      fireEvent.click(screen.getByText(/edit profile/i));
    });

    const input = screen.getByDisplayValue('John');

    fireEvent.change(input, { target: { value: 'Jane' } });

    expect(input.value).toBe('Jane');
  });

  test('calls API on save', async () => {
    userAPI.getProfile.mockResolvedValueOnce({ data: mockUser });
    reviewAPI.getUserReviews.mockResolvedValueOnce({ data: [] });
    userAPI.updateUser.mockResolvedValueOnce({});

    renderComponent();

    await waitFor(() => {
      fireEvent.click(screen.getByText(/edit profile/i));
    });

    fireEvent.click(screen.getByText(/save/i));

    await waitFor(() => {
      expect(userAPI.updateUser).toHaveBeenCalled();
    });
  });

});