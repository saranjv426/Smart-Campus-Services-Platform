import { render, screen, waitFor } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';
import StaffDashboard from './StaffDashboard';

jest.mock('axios');

const mockNavigate = jest.fn();

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockNavigate,
}));

describe('StaffDashboard Page Component', () => {
  const mockStaff = {
    id: 1,
    firstName: 'Staff',
    lastName: 'Member',
    role: 'staff',
    email: 'staff@example.com',
  };

  const mockBookings = [
    {
      id: 1,
      status: 'pending',
      startTime: '2024-02-20T10:00:00Z',
      endTime: '2024-02-20T12:00:00Z',
      createdAt: '2024-02-15T10:00:00Z',
      notes: 'Need projector access',
      service: {
        name: 'Main Library',
      },
      user: {
        firstName: 'John',
        lastName: 'Doe',
        email: 'john@example.com',
        phone: '123-456-7890',
      },
    },
    {
      id: 2,
      status: 'pending',
      startTime: '2024-02-21T14:00:00Z',
      endTime: '2024-02-21T16:00:00Z',
      createdAt: '2024-02-16T09:30:00Z',
      notes: 'Group study session',
      service: {
        name: 'Study Room',
      },
      user: {
        firstName: 'Jane',
        lastName: 'Smith',
        email: 'jane@example.com',
        phone: '555-111-2222',
      },
    },
  ];

  const renderStaffDashboard = () => {
    render(
      <BrowserRouter>
        <StaffDashboard />
      </BrowserRouter>
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
    localStorage.clear();
    mockNavigate.mockClear();

    localStorage.setItem('token', 'mock-token');
    localStorage.setItem('user', JSON.stringify(mockStaff));

    axios.get.mockResolvedValue({ data: mockBookings });
  });

  test('redirects to login if no token', async () => {
    localStorage.clear();

    renderStaffDashboard();

    await waitFor(() => {
      expect(mockNavigate).toHaveBeenCalledWith('/login');
    });
  });

  test('redirects to home if user is not staff', async () => {
    const notStaffUser = { ...mockStaff, role: 'student' };
    localStorage.setItem('token', 'mock-token');
    localStorage.setItem('user', JSON.stringify(notStaffUser));

    renderStaffDashboard();

    await waitFor(() => {
      expect(mockNavigate).toHaveBeenCalledWith('/');
    });
  });

  test('renders staff dashboard heading and welcome text', async () => {
    renderStaffDashboard();

    expect(screen.getByText('Staff Dashboard')).toBeInTheDocument();

    await waitFor(() => {
      expect(screen.getByText(/Welcome, Staff Member/i)).toBeInTheDocument();
    });
  });

  test('displays loading state initially', () => {
    renderStaffDashboard();

    expect(screen.getByText(/Loading pending bookings.../i)).toBeInTheDocument();
  });

  test('fetches pending bookings on mount', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith(
        'http://localhost:8080/api/approval/staff/1/pending',
        {
          headers: {
            Authorization: 'Bearer mock-token',
          },
        }
      );
    });
  });

  test('displays pending bookings after loading', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
      expect(screen.getByText('Study Room')).toBeInTheDocument();
      expect(screen.getByText('John Doe')).toBeInTheDocument();
      expect(screen.getByText('Jane Smith')).toBeInTheDocument();
    });
  });

  test('shows pending bookings count badge', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('2')).toBeInTheDocument();
    });
  });

  test('displays booking detail fields in cards', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getAllByText('Student Name:').length).toBeGreaterThan(0);
      expect(screen.getAllByText('Email:').length).toBeGreaterThan(0);
      expect(screen.getAllByText('Phone:').length).toBeGreaterThan(0);
      expect(screen.getAllByText('Start Time:').length).toBeGreaterThan(0);
      expect(screen.getAllByText('End Time:').length).toBeGreaterThan(0);
      expect(screen.getAllByText('Reason:').length).toBeGreaterThan(0);
      expect(screen.getAllByText('Requested:').length).toBeGreaterThan(0);
    });
  });

  test('displays error message on fetch failure', async () => {
    axios.get.mockRejectedValueOnce({
      response: {
        data: {
          error: 'Failed to load pending bookings',
        },
      },
    });

    renderStaffDashboard();

    await waitFor(() => {
      expect(
        screen.getByText('Failed to load pending bookings')
      ).toBeInTheDocument();
    });
  });

  test('displays empty state when there are no pending bookings', async () => {
    axios.get.mockResolvedValueOnce({ data: [] });

    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('No pending bookings')).toBeInTheDocument();
      expect(
        screen.getByText('There are currently no bookings awaiting review.')
      ).toBeInTheDocument();
    });
  });

  test('handles API response wrapped in data property', async () => {
    axios.get.mockResolvedValueOnce({
      data: {
        data: mockBookings,
      },
    });

    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
      expect(screen.getByText('Study Room')).toBeInTheDocument();
    });
  });

  test('falls back gracefully when booking fields are missing', async () => {
    axios.get.mockResolvedValueOnce({
      data: [
        {
          id: 99,
          startTime: null,
          endTime: null,
          createdAt: null,
          notes: '',
        },
      ],
    });

    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Service Not Available')).toBeInTheDocument();
      expect(screen.getAllByText('N/A').length).toBeGreaterThan(0);
    });
  });
});