import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';
import Bookings from '../pages/Bookings';

jest.mock('axios');

describe('Bookings Page Component', () => {
  const mockBookings = [
    {
      id: 'booking-1',
      serviceId: 'service-1',
      service: {
        name: 'Main Library',
        category: 'library',
        location: 'North Campus',
      },
      startTime: '2024-02-15T10:00:00Z',
      endTime: '2024-02-15T12:00:00Z',
      status: 'approved',
      notes: 'Study session',
    },
    {
      id: 'booking-2',
      serviceId: 'service-2',
      service: {
        name: 'Student Dining',
        category: 'dining',
        location: 'Food Court',
      },
      startTime: '2024-02-15T12:00:00Z',
      endTime: '2024-02-15T13:00:00Z',
      status: 'completed',
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
    localStorage.setItem('user', JSON.stringify({ id: 'user-1' }));
    axios.get.mockResolvedValue({ data: mockBookings });
  });

  afterEach(() => {
    localStorage.clear();
  });

  test('renders bookings page', () => {
    renderBookings();
    expect(screen.getByText(/My Bookings/i)).toBeInTheDocument();
  });

  test('loads user bookings on mount without a status filter', async () => {
    renderBookings();

    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith('/bookings/user/user-1', undefined);
    });
  });

  test('displays booking items and status labels', async () => {
    renderBookings();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
      expect(screen.getByText('Student Dining')).toBeInTheDocument();
      expect(screen.getAllByText('Accepted').length).toBeGreaterThan(0);
      expect(screen.getAllByText('Completed').length).toBeGreaterThan(0);
    });
  });

  test('requests filtered bookings from the backend when a status is selected', async () => {
    renderBookings();

    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith('/bookings/user/user-1', undefined);
    });

    fireEvent.click(screen.getByRole('button', { name: /Completed/i }));

    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith('/bookings/user/user-1', {
        params: { status: 'completed' },
      });
    });
  });

  test('shows an empty state when the backend returns no bookings for a filter', async () => {
    axios.get
      .mockResolvedValueOnce({ data: mockBookings })
      .mockResolvedValueOnce({ data: [] });

    renderBookings();

    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith('/bookings/user/user-1', undefined);
    });

    fireEvent.click(screen.getByRole('button', { name: /Rejected/i }));

    await waitFor(() => {
      expect(screen.getByText(/No bookings found for this status/i)).toBeInTheDocument();
    });
  });
});
