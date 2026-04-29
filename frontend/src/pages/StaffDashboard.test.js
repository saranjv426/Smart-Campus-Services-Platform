
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { BrowserRouter } from 'react-router-dom';
import StaffDashboard from './StaffDashboard';
import { approvalAPI } from '../services/approvalApi';

jest.mock('../services/approvalApi', () => ({
  approvalAPI: {
    getPendingBookingsByStaff: jest.fn(),
    approveBooking: jest.fn(),
    rejectBooking: jest.fn(),
  },
}));

const mockNavigate = jest.fn();

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockNavigate,
}));

describe('StaffDashboard Page Component', () => {
  let consoleErrorSpy;

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

    consoleErrorSpy = jest.spyOn(console, 'error').mockImplementation(() => {});

    localStorage.setItem('token', 'mock-token');
    localStorage.setItem('user', JSON.stringify(mockStaff));

    approvalAPI.getPendingBookingsByStaff.mockResolvedValue(mockBookings);
    approvalAPI.approveBooking.mockResolvedValue({ data: { success: true } });
    approvalAPI.rejectBooking.mockResolvedValue({ data: { success: true } });
  });

  afterEach(() => {
    consoleErrorSpy.mockRestore();
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

  test('fetches pending bookings on mount', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(approvalAPI.getPendingBookingsByStaff).toHaveBeenCalledWith(1);
    });
  });

  test('shows approve and reject buttons for each booking', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getAllByRole('button', { name: 'Approve' })).toHaveLength(2);
      expect(screen.getAllByRole('button', { name: 'Reject' })).toHaveLength(2);
    });
  });

  test('opens approve modal when approve button is clicked', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    await userEvent.click(screen.getAllByRole('button', { name: 'Approve' })[0]);

    expect(
      await screen.findByRole('heading', { name: 'Approve Booking' })
    ).toBeInTheDocument();

    expect(
      screen.getByPlaceholderText(/Add any notes about this approval or rejection/i)
    ).toBeInTheDocument();
  });

  test('opens reject modal when reject button is clicked', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    await userEvent.click(screen.getAllByRole('button', { name: 'Reject' })[0]);

    expect(
      await screen.findByRole('heading', { name: 'Reject Booking' })
    ).toBeInTheDocument();
  });

  test('allows user to type notes in modal', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    await userEvent.click(screen.getAllByRole('button', { name: 'Approve' })[0]);

    const textarea = await screen.findByPlaceholderText(
      /Add any notes about this approval or rejection/i
    );

    await userEvent.type(textarea, 'Approved by staff');

    expect(textarea).toHaveValue('Approved by staff');
  });

  test('submits approval action successfully', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    await userEvent.click(screen.getAllByRole('button', { name: 'Approve' })[0]);

    const textarea = await screen.findByPlaceholderText(
      /Add any notes about this approval or rejection/i
    );
    await userEvent.type(textarea, 'Looks good');

    const approveSubmitButton = await screen.findByRole('button', {
      name: 'Approve Booking',
    });
    await userEvent.click(approveSubmitButton);

    await waitFor(() => {
      expect(approvalAPI.approveBooking).toHaveBeenCalledWith(1, {
        status: 'approved',
        approvalNotes: 'Looks good',
        staffId: 1,
      });
    });

    await waitFor(() => {
      expect(approvalAPI.getPendingBookingsByStaff).toHaveBeenCalledTimes(2);
    });
  });

  test('submits rejection action successfully', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    await userEvent.click(screen.getAllByRole('button', { name: 'Reject' })[0]);

    const textarea = await screen.findByPlaceholderText(
      /Add any notes about this approval or rejection/i
    );
    await userEvent.type(textarea, 'Conflicting schedule');

    const rejectSubmitButton = await screen.findByRole('button', {
      name: 'Reject Booking',
    });
    await userEvent.click(rejectSubmitButton);

    await waitFor(() => {
      expect(approvalAPI.rejectBooking).toHaveBeenCalledWith(1, {
        status: 'rejected',
        approvalNotes: 'Conflicting schedule',
        staffId: 1,
      });
    });
  });

  test('closes modal when cancel button is clicked', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    await userEvent.click(screen.getAllByRole('button', { name: 'Approve' })[0]);

    expect(
      await screen.findByRole('heading', { name: 'Approve Booking' })
    ).toBeInTheDocument();

    await userEvent.click(screen.getByRole('button', { name: 'Cancel' }));

    await waitFor(() => {
      expect(
        screen.queryByRole('heading', { name: 'Approve Booking' })
      ).not.toBeInTheDocument();
    });
  });

  test('closes modal when overlay is clicked', async () => {
    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    await userEvent.click(screen.getAllByRole('button', { name: 'Approve' })[0]);

    expect(
      await screen.findByRole('heading', { name: 'Approve Booking' })
    ).toBeInTheDocument();

    await userEvent.click(screen.getByTestId('booking-action-modal-overlay'));

    await waitFor(() => {
      expect(
        screen.queryByRole('heading', { name: 'Approve Booking' })
      ).not.toBeInTheDocument();
    });
  });

  test('shows error message if approval submission fails', async () => {
    approvalAPI.approveBooking.mockRejectedValueOnce({
      response: {
        data: {
          error: 'Failed to approve booking',
        },
      },
    });

    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    await userEvent.click(screen.getAllByRole('button', { name: 'Approve' })[0]);

    const approveSubmitButton = await screen.findByRole('button', {
      name: 'Approve Booking',
    });
    await userEvent.click(approveSubmitButton);

    await waitFor(() => {
      expect(screen.getByText(/Failed to approve booking/i)).toBeInTheDocument();
    });
  });

  test('shows error message if rejection submission fails', async () => {
    approvalAPI.rejectBooking.mockRejectedValueOnce({
      response: {
        data: {
          error: 'Failed to reject booking',
        },
      },
    });

    renderStaffDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    await userEvent.click(screen.getAllByRole('button', { name: 'Reject' })[0]);

    const rejectSubmitButton = await screen.findByRole('button', {
      name: 'Reject Booking',
    });
    await userEvent.click(rejectSubmitButton);

    await waitFor(() => {
      expect(screen.getByText(/Failed to reject booking/i)).toBeInTheDocument();
    });
  });
});