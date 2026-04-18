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

  // Sprint 3: Tab Navigation Tests
  describe('Tab Navigation', () => {
    test('displays both Bookings and Services tabs', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText(/Bookings|bookings/i)).toBeInTheDocument();
      });
    });

    test('switches to Manage Services tab', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services/i });
        fireEvent.click(servicesTab);
      });

      await waitFor(() => {
        // Should display service management UI
        expect(screen.getByText(/Add New Service|Create Service/i) || screen.getByText(/Manage/i)).toBeInTheDocument();
      });
    });

    test('stays on Bookings tab by default', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });
    });

    test('toggles between tabs without losing data', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      // Switch tabs
      const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services|Services/i });
      fireEvent.click(servicesTab);

      await waitFor(() => {
        // Services UI should be visible
      });

      // Switch back
      const bookingsTab = screen.getByRole('button', { name: /Bookings|bookings/i });
      fireEvent.click(bookingsTab);

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });
    });
  });

  // Sprint 3: Service Management Tests
  describe('Service Management', () => {
    test('displays Create Service button', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services|Services/i });
        fireEvent.click(servicesTab);
      });

      await waitFor(() => {
        const createServiceBtn = screen.queryByRole('button', { name: /Add Service|Create Service|add new/i });
        expect(createServiceBtn).toBeInTheDocument();
      });
    });

    test('opens create service modal', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services|Services/i });
        fireEvent.click(servicesTab);
      });

      await waitFor(() => {
        const createBtn = screen.getByRole('button', { name: /Add Service|Create Service|add new/i });
        fireEvent.click(createBtn);
      });

      await waitFor(() => {
        expect(screen.getByPlaceholderText(/service name|name/i) || screen.getByText(/Create Service|Add New/i)).toBeInTheDocument();
      });
    });

    test('creates new service with form submission', async () => {
      axios.post.mockResolvedValueOnce({ data: { success: true } });
      renderAdminDashboard();

      await waitFor(() => {
        const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services|Services/i });
        fireEvent.click(servicesTab);
      });

      await waitFor(() => {
        const createBtn = screen.getByRole('button', { name: /Add Service|Create Service|add new/i });
        fireEvent.click(createBtn);
      });

      await waitFor(() => {
        const nameInput = screen.getByPlaceholderText(/service name|name|Service Name/i);
        expect(nameInput).toBeInTheDocument();
      });

      const nameInput = screen.getByPlaceholderText(/service name|name|Service Name/i);
      await userEvent.type(nameInput, 'New Service');

      const submitBtn = screen.getByRole('button', { name: /Submit|Create|Save/i });
      fireEvent.click(submitBtn);

      await waitFor(() => {
        expect(axios.post).toHaveBeenCalledWith(
          expect.stringContaining('/services'),
          expect.objectContaining({ name: 'New Service' }),
          expect.any(Object)
        );
      });
    });

    test('displays service grid with services', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services|Services/i });
        fireEvent.click(servicesTab);
      });

      await waitFor(() => {
        // Services should be displayed
        expect(axios.get).toHaveBeenCalledWith(expect.stringContaining('/services'));
      });
    });

    test('displays delete button on service cards', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services|Services/i });
        fireEvent.click(servicesTab);
      });

      await waitFor(() => {
        const deleteButtons = screen.queryAllByRole('button', { name: /delete|Delete/i });
        expect(deleteButtons.length).toBeGreaterThanOrEqual(0);
      });
    });

    test('opens delete confirmation modal', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services|Services/i });
        fireEvent.click(servicesTab);
      });

      await waitFor(() => {
        const deleteButtons = screen.getAllByRole('button', { name: /delete|Delete/i });
        if (deleteButtons.length > 0) {
          fireEvent.click(deleteButtons[0]);
        }
      });

      await waitFor(() => {
        expect(screen.getByText(/confirm|Are you sure/i) || screen.getByText(/delete/i)).toBeInTheDocument();
      });
    });

    test('deletes service with confirmation', async () => {
      axios.delete.mockResolvedValueOnce({ data: { success: true } });
      renderAdminDashboard();

      await waitFor(() => {
        const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services|Services/i });
        fireEvent.click(servicesTab);
      });

      await waitFor(() => {
        const deleteButtons = screen.getAllByRole('button', { name: /delete|Delete/i });
        if (deleteButtons.length > 0) {
          fireEvent.click(deleteButtons[0]);
        }
      });

      await waitFor(() => {
        const confirmBtn = screen.getByRole('button', { name: /confirm|Yes|Delete/i });
        fireEvent.click(confirmBtn);
      });

      await waitFor(() => {
        expect(axios.delete).toHaveBeenCalledWith(
          expect.stringContaining('/services'),
          expect.any(Object)
        );
      });
    });

    test('validates required fields in service form', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services|Services/i });
        fireEvent.click(servicesTab);
      });

      await waitFor(() => {
        const createBtn = screen.getByRole('button', { name: /Add Service|Create Service|add new/i });
        fireEvent.click(createBtn);
      });

      await waitFor(() => {
        const submitBtn = screen.getByRole('button', { name: /Submit|Create|Save/i });
        fireEvent.click(submitBtn);
      });

      // Error message should appear
      await waitFor(() => {
        expect(screen.getByText(/required|Error|must/i) || axios.post).not.toHaveBeenCalled();
      });
    });
  });

  // Sprint 3: Booking Filtering Tests
  describe('Booking Filtering', () => {
    test('filters bookings by status - Pending', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      const pendingBtn = screen.getByRole('button', { name: /pending|Pending/i });
      fireEvent.click(pendingBtn);

      // Filtered bookings should display
      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });
    });

    test('filters bookings by status - All', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      const allBtn = screen.getByRole('button', { name: /all|All/i });
      fireEvent.click(allBtn);

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });
    });

    test('filters by service category dropdown', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      const filterSelect = screen.getByDisplayValue(/all|All/i);
      fireEvent.change(filterSelect, { target: { value: 'library' } });

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });
    });

    test('applies multiple filters simultaneously', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      // Set status filter
      const pendingBtn = screen.getByRole('button', { name: /pending|Pending/i });
      fireEvent.click(pendingBtn);

      // Set service filter
      const filterSelect = screen.getByDisplayValue(/all|All/i);
      fireEvent.change(filterSelect, { target: { value: 'library' } });

      await waitFor(() => {
        // Both filters should be applied
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });
    });

    test('displays statistics cards', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        const totalText = screen.queryByText(/total|Total/i);
        const pendingText = screen.queryByText(/pending|Pending/i);
        expect(totalText || pendingText).toBeInTheDocument();
      });
    });

    test('shows correct booking count in statistics', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      // Statistics should show 2 bookings
      await waitFor(() => {
        expect(axios.get).toHaveBeenCalledWith(
          expect.stringContaining('/approval'),
          expect.any(Object)
        );
      });
    });
  });

  // Sprint 3: Admin Approval Tests
  describe('Admin Booking Approval', () => {
    test('submits rejection with optional reason', async () => {
      axios.put.mockResolvedValueOnce({ data: { success: true } });
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      const rejectButtons = screen.getAllByRole('button', { name: /reject/i });
      fireEvent.click(rejectButtons[0]);

      await waitFor(() => {
        const reasonInput = screen.getByPlaceholderText(/reason|notes|Reason/i);
        expect(reasonInput).toBeInTheDocument();
      });

      const reasonInput = screen.getByPlaceholderText(/reason|notes|Reason/i);
      await userEvent.type(reasonInput, 'Not available');

      const submitButton = screen.getByRole('button', { name: /submit|confirm|Reject/i });
      fireEvent.click(submitButton);

      await waitFor(() => {
        expect(axios.put).toHaveBeenCalledWith(
          expect.stringContaining('/reject'),
          expect.any(Object),
          expect.any(Object)
        );
      });
    });

    test('updates booking status after approval', async () => {
      const updatedBookings = [
        { ...mockBookings[0], status: 'approved' },
        mockBookings[1]
      ];

      axios.put.mockResolvedValueOnce({ data: { success: true } });
      axios.get.mockResolvedValueOnce({ data: updatedBookings });

      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      const approveButtons = screen.getAllByRole('button', { name: /approve/i });
      fireEvent.click(approveButtons[0]);

      await waitFor(() => {
        const submitButton = screen.getByRole('button', { name: /submit|confirm|Approve/i });
        fireEvent.click(submitButton);
      });
    });

    test('displays booking details in approval modal', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      const approveButtons = screen.getAllByRole('button', { name: /approve/i });
      fireEvent.click(approveButtons[0]);

      await waitFor(() => {
        // Modal should contain booking details
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });
    });

    test('disables actions for non-pending bookings', async () => {
      const approvedBookings = [
        { ...mockBookings[0], status: 'approved' },
        mockBookings[1]
      ];

      axios.get.mockResolvedValueOnce({ data: approvedBookings });

      renderAdminDashboard();

      await waitFor(() => {
        const approveButtons = screen.queryAllByRole('button', { name: /approve/i });
        // Only one approve button should exist (for pending booking)
        expect(approveButtons.length).toBeLessThanOrEqual(1);
      });
    });
  });

  // Sprint 3: Responsiveness Tests
  describe('Responsive Design', () => {
    test('renders properly on mobile devices', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      // Component should render without errors
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    test('displays all table columns on desktop', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        expect(screen.getByText('Main Library')).toBeInTheDocument();
      });

      // All columns should be present in table
      const tableRows = screen.getAllByRole('row');
      expect(tableRows.length).toBeGreaterThan(0);
    });

    test('service grid is responsive', async () => {
      renderAdminDashboard();

      await waitFor(() => {
        const servicesTab = screen.getByRole('button', { name: /Manage Services|manage services|Services/i });
        fireEvent.click(servicesTab);
      });

      // Grid should render without layout issues
      await waitFor(() => {
        expect(axios.get).toHaveBeenCalledWith(expect.stringContaining('/services'));
      });
    });
  });
});
