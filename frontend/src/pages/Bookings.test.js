import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';
import Bookings from '../pages/Bookings';

jest.mock('axios');

describe('Bookings Page Component', () => {
  const mockBookings = [
    {
      id: 1,
      serviceId: 1,
      serviceName: 'Main Library',
      category: 'library',
      startTime: '2024-02-15T10:00:00Z',
      endTime: '2024-02-15T12:00:00Z',
      status: 'confirmed',
      notes: 'Study session',
    },
    {
      id: 2,
      serviceId: 2,
      serviceName: 'Student Dining',
      category: 'dining',
      startTime: '2024-02-15T12:00:00Z',
      endTime: '2024-02-15T13:00:00Z',
      status: 'confirmed',
      notes: 'Lunch',
    },
  ];

  const renderBookings = () => {
    render(
      <BrowserRouter>
        <Bookings />
      </BrowserRouter>
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
    axios.get.mockResolvedValue({ data: mockBookings });
  });

  test('renders bookings page', () => {
    renderBookings();
    expect(screen.getByText(/My Bookings/i)).toBeInTheDocument();
  });

  test('loads user bookings on mount', async () => {
    renderBookings();
    await waitFor(() => {
      expect(axios.get).toHaveBeenCalled();
    });
  });

  test('displays booking items', async () => {
    renderBookings();
    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
      expect(screen.getByText('Student Dining')).toBeInTheDocument();
    });
  });

  test('displays booking status', async () => {
    renderBookings();
    await waitFor(() => {
      const statusElements = screen.getAllByText(/confirmed/i);
      expect(statusElements.length).toBeGreaterThan(0);
    });
  });

  test('displays booking details', async () => {
    renderBookings();
    await waitFor(() => {
      expect(screen.getByText('Study session')).toBeInTheDocument();
      expect(screen.getByText('Lunch')).toBeInTheDocument();
    });
  });
});
