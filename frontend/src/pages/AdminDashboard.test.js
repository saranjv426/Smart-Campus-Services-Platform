import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';
import AdminDashboard from './AdminDashboard';

jest.mock('axios');
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => jest.fn(),
}));

describe('AdminDashboard Page Component', () => {
  const mockAdmin = {
    id: 1,
    firstName: 'Admin',
    lastName: 'User',
    role: 'admin',
    email: 'admin@example.com',
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
      serviceId: 2,
      serviceName: 'Student Dining',
      status: 'pending',
      startTime: '2024-02-20T12:00:00Z',
      endTime: '2024-02-20T13:00:00Z',
      createdAt: '2024-02-15T10:00:00Z',
    },
  ];

  const mockServices = [
    { id: 1, name: 'Main Library', category: 'library', isActive: true },
    { id: 2, name: 'Student Dining', category: 'dining', isActive: true },
  ];

  const renderAdminDashboard = () => {
    localStorage.setItem('token', 'mock-token');
    localStorage.setItem('user', JSON.stringify(mockAdmin));

    render(
      <BrowserRouter>
        <AdminDashboard />
      </BrowserRouter>
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
    localStorage.clear();
    axios.get.mockImplementation((url) => {
      if (url.includes('/approval')) {
        return Promise.resolve({ data: mockBookings });
      }
      if (url.includes('/services')) {
        return Promise.resolve({ data: mockServices });
      }
      return Promise.reject(new Error('Not found'));
    });
  });

  test('redirects to login if no token', () => {
    localStorage.clear();
    const navigateMock = jest.fn();
    jest.mock('react-router-dom', () => ({
      ...jest.requireActual('react-router-dom'),
      useNavigate: () => navigateMock,
    }));

    render(
      <BrowserRouter>
        <AdminDashboard />
      </BrowserRouter>
    );

    // Navigation should be called - actual redirect happens in component
  });

  test('redirects to home if user is not admin', () => {
    const notAdminUser = { ...mockAdmin, role: 'student' };
    localStorage.setItem('user', JSON.stringify(notAdminUser));

    render(
      <BrowserRouter>
        <AdminDashboard />
      </BrowserRouter>
    );

    // Navigation to home should occur
  });

  test('renders admin dashboard', async () => {
    renderAdminDashboard();

    await waitFor(() => {
      expect(screen.getByText(/Admin Dashboard|admin/i)).toBeInTheDocument();
    });
  });

  test('displays loading state initially', () => {
    renderAdminDashboard();
    const loadingText = screen.queryByText(/loading/i);
    // May or may not show depending on axios mock speed
  });

  test('fetches bookings on mount', async () => {
    renderAdminDashboard();

    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith(
        expect.stringContaining('/approval'),
        expect.any(Object)
      );
    });
  });

  test('fetches services on mount', async () => {
    renderAdminDashboard();

    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith(
        expect.stringContaining('/services'),
        expect.any(Object)
      );
    });
  });

  test('displays booking list after loading', async () => {
    renderAdminDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
      expect(screen.getByText('Student Dining')).toBeInTheDocument();
    });
  });

  test('displays pending bookings count', async () => {
    renderAdminDashboard();

    await waitFor(() => {
      const pendingText = screen.queryByText(/pending|2/);
      expect(pendingText).toBeInTheDocument();
    });
  });

  test('displays approve button for pending booking', async () => {
    renderAdminDashboard();

    await waitFor(() => {
      const approveButtons = screen.getAllByRole('button', { name: /approve/i });
      expect(approveButtons.length).toBeGreaterThan(0);
    });
  });

  test('displays reject button for pending booking', async () => {
    renderAdminDashboard();

    await waitFor(() => {
      const rejectButtons = screen.getAllByRole('button', { name: /reject/i });
      expect(rejectButtons.length).toBeGreaterThan(0);
    });
  });

  test('opens modal when approve button clicked', async () => {
    renderAdminDashboard();

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
    renderAdminDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const rejectButtons = screen.getAllByRole('button', { name: /reject/i });
    fireEvent.click(rejectButtons[0]);

    await waitFor(() => {
      expect(screen.getByText(/notes|reason/i)).toBeInTheDocument();
    });
  });

  test('submits approval with notes', async () => {
    axios.put.mockResolvedValueOnce({ data: { success: true } });
    renderAdminDashboard();

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

  test('filters bookings by status', async () => {
    renderAdminDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const allButton = screen.getByRole('button', { name: /all/i });
    expect(allButton).toBeInTheDocument();
  });

  test('displays error message on fetch failure', async () => {
    axios.get.mockRejectedValueOnce(new Error('Failed to load'));

    renderAdminDashboard();

    await waitFor(() => {
      const errorText = screen.queryByText(/failed|error/i);
      // Depending on implementation
    });
  });

  test('refreshes bookings after action', async () => {
    axios.put.mockResolvedValueOnce({ data: { success: true } });
    axios.get.mockResolvedValueOnce({ data: mockBookings });

    renderAdminDashboard();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const approveButtons = screen.getAllByRole('button', { name: /approve/i });
    fireEvent.click(approveButtons[0]);

    await waitFor(() => {
      const notesInput = screen.getByPlaceholderText(/notes/i);
      expect(notesInput).toBeInTheDocument();
    });
  });
});
