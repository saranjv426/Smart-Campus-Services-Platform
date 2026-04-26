import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import Bookings from './Bookings';
import { bookingAPI } from '../services/api';

jest.mock('../services/api', () => ({
  bookingAPI: {
    getUserBookings: jest.fn(),
    cancelBooking: jest.fn(),
  },
}));

const mockConfirm = jest.fn();
const mockAlert = jest.fn();

describe('Bookings Page Component', () => {
  const mockUser = {
    id: 1,
    firstName: 'John',
    lastName: 'Doe',
    role: 'student',
    email: 'john@example.com',
  };

  const mockBookings = [
    {
      id: 1,
      serviceId: 1,
      status: 'accepted',
      startTime: '2024-02-15T10:00:00Z',
      endTime: '2024-02-15T12:00:00Z',
      notes: 'Study session',
      service: {
        id: 1,
        name: 'Main Library',
        location: 'Campus Center',
        category: 'library',
      },
    },
    {
      id: 2,
      serviceId: 2,
      status: 'pending',
      startTime: '2024-02-15T12:00:00Z',
      endTime: '2024-02-15T13:00:00Z',
      notes: 'Lunch',
      service: {
        id: 2,
        name: 'Student Dining',
        location: 'Food Court',
        category: 'dining',
      },
    },
  ];

  const renderBookings = () => {
    localStorage.setItem('token', 'mock-token');
    localStorage.setItem('user', JSON.stringify(mockUser));

    render(
      <BrowserRouter>
        <Bookings />
      </BrowserRouter>
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
    localStorage.clear();

    window.confirm = mockConfirm;
    window.alert = mockAlert;

    mockConfirm.mockReturnValue(true);

    bookingAPI.getUserBookings.mockResolvedValue({ data: mockBookings });
    bookingAPI.cancelBooking.mockResolvedValue({ data: { success: true } });
  });

  test('renders bookings page', () => {
    renderBookings();
    expect(screen.getByText(/My Bookings/i)).toBeInTheDocument();
  });

  test('loads user bookings on mount', async () => {
    renderBookings();

    await waitFor(() => {
      expect(bookingAPI.getUserBookings).toHaveBeenCalledWith(1);
    });
  });

  test('displays booking items', async () => {
    renderBookings();

    expect(await screen.findByText('Main Library')).toBeInTheDocument();
    expect(screen.getByText('Student Dining')).toBeInTheDocument();
  });

  test('displays booking status', async () => {
    renderBookings();

    await waitFor(() => {
      expect(screen.getByText(/Accepted/i)).toBeInTheDocument();
      expect(screen.getByText(/Pending/i)).toBeInTheDocument();
    });
  });

  test('displays booking details', async () => {
    renderBookings();

    expect(await screen.findByText('Study session')).toBeInTheDocument();
    expect(screen.getByText('Lunch')).toBeInTheDocument();
    expect(screen.getByText('Campus Center')).toBeInTheDocument();
    expect(screen.getByText('Food Court')).toBeInTheDocument();
  });

  test('filters bookings by status', async () => {
    renderBookings();

    expect(await screen.findByText('Main Library')).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button', { name: /Pending/i }));

    await waitFor(() => {
      expect(screen.queryByText('Main Library')).not.toBeInTheDocument();
      expect(screen.getByText('Student Dining')).toBeInTheDocument();
    });
  });

  test('resets filtered bookings when All is clicked', async () => {
    renderBookings();

    expect(await screen.findByText('Main Library')).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button', { name: /Pending/i }));

    await waitFor(() => {
      expect(screen.queryByText('Main Library')).not.toBeInTheDocument();
      expect(screen.getByText('Student Dining')).toBeInTheDocument();
    });

    fireEvent.click(screen.getByRole('button', { name: /All/i }));

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
      expect(screen.getByText('Student Dining')).toBeInTheDocument();
    });
  });

  test('shows cancel button only for pending booking', async () => {
    renderBookings();

    expect(await screen.findByText('Student Dining')).toBeInTheDocument();

    const cancelButtons = screen.getAllByRole('button', { name: /Cancel Booking/i });
    expect(cancelButtons).toHaveLength(1);
  });

  test('cancels a pending booking and refreshes the list', async () => {
    renderBookings();

    expect(await screen.findByText('Student Dining')).toBeInTheDocument();

    const cancelButton = screen.getByRole('button', { name: /Cancel Booking/i });
    fireEvent.click(cancelButton);

    await waitFor(() => {
      expect(window.confirm).toHaveBeenCalled();
      expect(bookingAPI.cancelBooking).toHaveBeenCalledWith(2);
    });

    await waitFor(() => {
      expect(window.alert).toHaveBeenCalled();
      expect(bookingAPI.getUserBookings).toHaveBeenCalledTimes(2);
    });
  });

  test('refresh button reloads bookings', async () => {
    renderBookings();

    expect(await screen.findByText('Main Library')).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button', { name: /Refresh/i }));

    await waitFor(() => {
      expect(bookingAPI.getUserBookings).toHaveBeenCalledTimes(2);
    });
  });

  test('shows empty state when there are no bookings', async () => {
    bookingAPI.getUserBookings.mockResolvedValueOnce({ data: [] });

    renderBookings();

    await waitFor(() => {
      expect(
        screen.getByText(/No bookings found|You have no bookings/i)
      ).toBeInTheDocument();
    });
  });

  test('shows login error when user is not authenticated', async () => {
    localStorage.clear();

    render(
      <BrowserRouter>
        <Bookings />
      </BrowserRouter>
    );

    await waitFor(() => {
      expect(
        screen.getByText(/Please login to view your bookings/i)
      ).toBeInTheDocument();
    });
  });
});