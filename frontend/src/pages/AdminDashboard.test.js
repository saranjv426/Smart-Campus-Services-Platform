import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';
import AdminDashboard from './AdminDashboard';

jest.mock('axios');

const mockNavigate = jest.fn();

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockNavigate,
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
      status: 'pending',
      startTime: '2024-02-20T10:00:00Z',
      endTime: '2024-02-20T12:00:00Z',
      createdAt: '2024-02-15T10:00:00Z',
      approvalNotes: '',
      service: {
        id: 1,
        name: 'Main Library',
        category: 'library',
      },
      user: {
        id: 2,
        firstName: 'John',
        lastName: 'Doe',
        email: 'john@example.com',
      },
    },
    {
      id: 2,
      userId: 3,
      serviceId: 2,
      status: 'pending',
      startTime: '2024-02-20T12:00:00Z',
      endTime: '2024-02-20T13:00:00Z',
      createdAt: '2024-02-15T10:00:00Z',
      approvalNotes: '',
      service: {
        id: 2,
        name: 'Student Dining',
        category: 'dining',
      },
      user: {
        id: 3,
        firstName: 'Jane',
        lastName: 'Smith',
        email: 'jane@example.com',
      },
    },
  ];

  const mockServices = [
    {
      id: 1,
      name: 'Main Library',
      category: 'library',
      isActive: true,
      description: '',
      location: '',
    },
    {
      id: 2,
      name: 'Student Dining',
      category: 'dining',
      isActive: true,
      description: '',
      location: '',
    },
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
    mockNavigate.mockClear();

    axios.get.mockImplementation((url) => {
      if (url.includes('/approval')) {
        return Promise.resolve({ data: mockBookings });
      }
      if (url.includes('/services')) {
        return Promise.resolve({ data: mockServices });
      }
      return Promise.reject(new Error('Not found'));
    });

    axios.put.mockResolvedValue({ data: { success: true } });
    axios.post.mockResolvedValue({ data: { success: true } });
    axios.delete.mockResolvedValue({ data: { success: true } });
  });

  test('redirects to login if no token', async () => {
    localStorage.clear();

    render(
      <BrowserRouter>
        <AdminDashboard />
      </BrowserRouter>
    );

    await waitFor(() => {
      expect(mockNavigate).toHaveBeenCalled();
    });
  });

  test('redirects to home if user is not admin', async () => {
    localStorage.setItem('token', 'mock-token');
    localStorage.setItem(
      'user',
      JSON.stringify({ ...mockAdmin, role: 'student' })
    );

    render(
      <BrowserRouter>
        <AdminDashboard />
      </BrowserRouter>
    );

    await waitFor(() => {
      expect(mockNavigate).toHaveBeenCalled();
    });
  });

  test('renders admin dashboard and welcome message', async () => {
    renderAdminDashboard();

    expect(await screen.findByText(/Admin Dashboard/i)).toBeInTheDocument();
    expect(screen.getByText(/Welcome,\s*Admin\s*User/i)).toBeInTheDocument();
  });

  test('fetches bookings and services on mount', async () => {
  renderAdminDashboard();

  await waitFor(() => {
    expect(axios.get).toHaveBeenCalledWith(
      expect.stringContaining('/approval'),
      expect.any(Object)
    );
    expect(axios.get).toHaveBeenCalledWith(
      expect.stringContaining('/services')
    );
  });
});

  test('displays statistics cards', async () => {
  renderAdminDashboard();

  expect(await screen.findByText(/Total Bookings/i)).toBeInTheDocument();
  expect(screen.getByText('Pending', { selector: '.stat-label' })).toBeInTheDocument();
  expect(screen.getByText('Approved', { selector: '.stat-label' })).toBeInTheDocument();
  expect(screen.getByText('Rejected', { selector: '.stat-label' })).toBeInTheDocument();
});

  test('displays booking list after loading', async () => {
    renderAdminDashboard();

    expect(await screen.findByText('Main Library')).toBeInTheDocument();
    expect(screen.getByText('Student Dining')).toBeInTheDocument();
    expect(screen.getByText('john@example.com')).toBeInTheDocument();
    expect(screen.getByText('jane@example.com')).toBeInTheDocument();
  });

  test('displays approve and reject action buttons', async () => {
    renderAdminDashboard();

    await waitFor(() => {
      expect(screen.getAllByTitle(/Approve booking/i).length).toBeGreaterThan(0);
      expect(screen.getAllByTitle(/Reject booking/i).length).toBeGreaterThan(0);
    });
  });

  test('opens modal when approve button is clicked', async () => {
    renderAdminDashboard();

    const approveButtons = await screen.findAllByTitle(/Approve booking/i);
    fireEvent.click(approveButtons[0]);

    await waitFor(() => {
      expect(
        screen.getByRole('heading', { name: /Approve Booking/i })
      ).toBeInTheDocument();
    });
  });

  test('opens modal when reject button is clicked', async () => {
    renderAdminDashboard();

    const rejectButtons = await screen.findAllByTitle(/Reject booking/i);
    fireEvent.click(rejectButtons[0]);

    await waitFor(() => {
      expect(
        screen.getByRole('heading', { name: /Reject Booking/i })
      ).toBeInTheDocument();
    });
  });

  test('submits approval with notes', async () => {
    renderAdminDashboard();

    const approveButtons = await screen.findAllByTitle(/Approve booking/i);
    fireEvent.click(approveButtons[0]);

    const notesInput = await screen.findByPlaceholderText(
      /Add any notes about this approval\/rejection/i
    );
    await userEvent.type(notesInput, 'Approved by admin');

    const submitButton = screen.getByRole('button', {
      name: /^Approve Booking$/i,
    });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(axios.put).toHaveBeenCalledWith(
        expect.stringContaining('/approve'),
        expect.any(Object),
        expect.any(Object)
      );
    });
  });

  test('submits rejection with notes', async () => {
    renderAdminDashboard();

    const rejectButtons = await screen.findAllByTitle(/Reject booking/i);
    fireEvent.click(rejectButtons[0]);

    const notesInput = await screen.findByPlaceholderText(
      /Add any notes about this approval\/rejection/i
    );
    await userEvent.type(notesInput, 'Rejected by admin');

    const submitButton = screen.getByRole('button', {
      name: /^Reject Booking$/i,
    });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(axios.put).toHaveBeenCalledWith(
        expect.stringContaining('/reject'),
        expect.any(Object),
        expect.any(Object)
      );
    });
  });

  test('displays both Bookings Overview and Manage Services tabs', async () => {
    renderAdminDashboard();

    expect(
      await screen.findByRole('button', { name: /bookings overview/i })
    ).toBeInTheDocument();
    expect(
      screen.getByRole('button', { name: /manage services/i })
    ).toBeInTheDocument();
  });

  test('switches to Manage Services tab', async () => {
    renderAdminDashboard();

    const servicesTab = await screen.findByRole('button', {
      name: /manage services/i,
    });
    fireEvent.click(servicesTab);

    await waitFor(() => {
      expect(screen.getByText(/Service Management/i)).toBeInTheDocument();
      expect(
        screen.getByRole('button', { name: /Create New Service|Create Service/i })
      ).toBeInTheDocument();
    });
  });

  test('filters bookings by status buttons', async () => {
    renderAdminDashboard();

    expect(await screen.findByText('Main Library')).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button', { name: /^Pending$/i }));
    fireEvent.click(screen.getByRole('button', { name: /^All$/i }));

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });
  });

  test('filters by service category dropdown', async () => {
    renderAdminDashboard();

    expect(await screen.findByText('Main Library')).toBeInTheDocument();

    const filterSelect = screen.getByRole('combobox');
    fireEvent.change(filterSelect, { target: { value: 'library' } });

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });
  });

  test('refreshes bookings after approval action', async () => {
    renderAdminDashboard();

    const approveButtons = await screen.findAllByTitle(/Approve booking/i);
    fireEvent.click(approveButtons[0]);

    const submitButton = await screen.findByRole('button', {
      name: /^Approve Booking$/i,
    });
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(axios.put).toHaveBeenCalled();
      expect(axios.get).toHaveBeenCalled();
    });
  });
});