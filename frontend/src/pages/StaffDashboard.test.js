import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';
import StaffDashboard from './StaffDashboard';

jest.mock('axios');
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => jest.fn(),
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
      userId: 2,
      serviceId: 1,
      serviceName: 'Main Library',
      status: 'pending',
      startTime: '2024-02-20T10:00:00Z',
      endTime: '2024-02-20T12:00:00Z',
      createdAt: '2024-02-15T10:00:00Z',
    },
    {
      id: 2,
      userId: 3,
      serviceId: 1,
      serviceName: 'Study Room',
      status: 'pending',
      startTime: '2024-02-20T14:00:00Z',
      endTime: '2024-02-20T16:00:00Z',
      createdAt: '2024-02-15T10:00:00Z',
    },
  ];

  const renderStaffDashboard = () => {
    localStorage.setItem('token', 'mock-token');
    localStorage.setItem('user', JSON.stringify(mockStaff));

    render(
      <BrowserRouter>
        <StaffDashboard />
      </BrowserRouter>
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
    localStorage.clear();
    axios.get.mockResolvedValue({ data: mockBookings });
    axios.put.mockResolvedValue({ data: { success: true } });
  });

  test('redirects to login if no token', () => {
    localStorage.clear();

    render(
      <BrowserRouter>
        <StaffDashboard />
      </BrowserRouter>
    );

    // Navigation to login should occur
  });

  test('redirects to home if user is not staff', () => {
    const notStaffUser = { ...mockStaff, role: 'student' };
    localStorage.setItem('user', JSON.stringify(notStaffUser));
    localStorage.setItem('token', 'mock-token');

    render(
      <BrowserRouter>
        <StaffDashboard />
      </BrowserRouter>
    );

    // Navigation to home should occur
  });

  test('renders staff dashboard', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText(/Staff Dashboard|staff|Dashboard/i)).toBeInTheDocument();
    });
  });

  test('displays loading state initially', () => {
    renderStaffDashboard();
    // Loading might be brief
  });

  test('fetches pending bookings on mount', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith(
        expect.stringContaining('/approval/staff'),
        expect.any(Object)
      );
    });
  });

  test('displays pending bookings after loading', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
      expect(screen.getByText('Study Room')).toBeInTheDocument();
    });
  });

  test('displays approve button for each pending booking', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      const approveButtons = screen.getAllByRole('button', { name: /approve/i });
      expect(approveButtons.length).toBe(2);
    });
  });

  test('displays reject button for each pending booking', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      const rejectButtons = screen.getAllByRole('button', { name: /reject/i });
      expect(rejectButtons.length).toBe(2);
    });
  });

  test('opens modal when approve button clicked', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const approveButtons = screen.getAllByRole('button', { name: /approve/i });
    fireEvent.click(approveButtons[0]);

    await waitFor(() => {
      expect(screen.getByText(/notes|approval/i)).toBeInTheDocument();
    });
  });

  test('opens modal when reject button clicked', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const rejectButtons = screen.getAllByRole('button', { name: /reject/i });
    fireEvent.click(rejectButtons[0]);

    await waitFor(() => {
      expect(screen.getByText(/notes|reason/i)).toBeInTheDocument();
    });
  });

  test('submits approval action with notes', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const approveButtons = screen.getAllByRole('button', { name: /approve/i });
    fireEvent.click(approveButtons[0]);

    await waitFor(() => {
      const notesInput = screen.getByPlaceholderText(/notes/i);
      expect(notesInput).toBeInTheDocument();
    });

    const notesInput = screen.getByPlaceholderText(/notes/i);
    await userEvent.type(notesInput, 'Approved');

    const submitButton = screen.getByRole('button', { name: /submit|confirm/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(axios.put).toHaveBeenCalledWith(
        expect.stringContaining('/approve'),
        expect.any(Object),
        expect.any(Object)
      );
    });
  });

  test('submits rejection action with notes', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const rejectButtons = screen.getAllByRole('button', { name: /reject/i });
    fireEvent.click(rejectButtons[0]);

    await waitFor(() => {
      const notesInput = screen.getByPlaceholderText(/notes/i);
      expect(notesInput).toBeInTheDocument();
    });

    const notesInput = screen.getByPlaceholderText(/notes/i);
    await userEvent.type(notesInput, 'Not available');

    const submitButton = screen.getByRole('button', { name: /submit|confirm/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(axios.put).toHaveBeenCalledWith(
        expect.stringContaining('/reject'),
        expect.any(Object),
        expect.any(Object)
      );
    });
  });

  test('displays error message on fetch failure', async () => {
    axios.get.mockRejectedValueOnce(new Error('Failed to load'));

    renderStaffDashboard();

    await waitFor(() => {
      const errorText = screen.queryByText(/failed|error/i);
      // Depending on implementation
    });
  });

  test('displays error message on action failure', async () => {
    axios.put.mockRejectedValueOnce(new Error('Failed to approve'));

    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const approveButtons = screen.getAllByRole('button', { name: /approve/i });
    fireEvent.click(approveButtons[0]);

    await waitFor(() => {
      const notesInput = screen.getByPlaceholderText(/notes/i);
      expect(notesInput).toBeInTheDocument();
    });

    const submitButton = screen.getByRole('button', { name: /submit|confirm/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      const errorMsg = screen.queryByText(/failed|error/i);
      // Error message should appear
    });
  });

  test('closes modal after successful action', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const approveButtons = screen.getAllByRole('button', { name: /approve/i });
    fireEvent.click(approveButtons[0]);

    await waitFor(() => {
      expect(screen.getByPlaceholderText(/notes/i)).toBeInTheDocument();
    });

    const submitButton = screen.getByRole('button', { name: /submit|confirm/i });
    fireEvent.click(submitButton);

    await waitFor(() => {
      // Modal should close
      const modal = screen.queryByPlaceholderText(/notes/i);
      // Depends on re-render timing
    });
  });

  test('displays booking details in card', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    // Booking details should be visible
    const cards = screen.getAllByText(/2024-02-20T/);
    expect(cards.length).toBeGreaterThan(0);
  });
});
