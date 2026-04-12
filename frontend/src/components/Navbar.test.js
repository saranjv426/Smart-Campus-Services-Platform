import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import Navbar from './Navbar';

const mockNavigate = jest.fn();

jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockNavigate,
}));

describe('Navbar Component', () => {
  const renderNavbar = () => {
    render(
      <BrowserRouter>
        <Navbar />
      </BrowserRouter>
    );
  };

  const setLoggedOutState = () => {
    localStorage.clear();
  };

  const setLoggedInUser = (user) => {
    localStorage.setItem('token', 'mock-token');
    localStorage.setItem('user', JSON.stringify(user));
  };

  beforeEach(() => {
    localStorage.clear();
    jest.clearAllMocks();
    mockNavigate.mockClear();
  });

  test('renders navigation bar', () => {
    renderNavbar();
    expect(screen.getByRole('navigation')).toBeInTheDocument();
  });

  test('shows guest links when user is not logged in', async () => {
    setLoggedOutState();
    renderNavbar();

    await waitFor(() => {
      expect(screen.getByText('Home')).toBeInTheDocument();
      expect(screen.getByText('Services')).toBeInTheDocument();
      expect(screen.getByText('Login')).toBeInTheDocument();
      expect(screen.getByText('Register')).toBeInTheDocument();
    });

    expect(screen.queryByText('My Bookings')).not.toBeInTheDocument();
    expect(screen.queryByText('Staff Dashboard')).not.toBeInTheDocument();
    expect(screen.queryByText('Admin Dashboard')).not.toBeInTheDocument();
    expect(screen.queryByText('Logout')).not.toBeInTheDocument();
  });

  test('shows student links when logged in as student', async () => {
    setLoggedInUser({
      id: 1,
      firstName: 'Keerthi',
      role: 'student',
    });

    renderNavbar();

    await waitFor(() => {
      expect(screen.getByText('My Bookings')).toBeInTheDocument();
      expect(screen.getByText('Profile')).toBeInTheDocument();
      expect(screen.getByText('Logout')).toBeInTheDocument();
      expect(screen.getByText(/Welcome, Keerthi/i)).toBeInTheDocument();
    });

    expect(screen.queryByText('Staff Dashboard')).not.toBeInTheDocument();
    expect(screen.queryByText('Admin Dashboard')).not.toBeInTheDocument();
    expect(screen.queryByText('Login')).not.toBeInTheDocument();
    expect(screen.queryByText('Register')).not.toBeInTheDocument();
  });

  test('shows staff links when logged in as staff', async () => {
    setLoggedInUser({
      id: 2,
      firstName: 'Alex',
      role: 'staff',
    });

    renderNavbar();

    await waitFor(() => {
      expect(screen.getByText('Staff Dashboard')).toBeInTheDocument();
      expect(screen.getByText('Logout')).toBeInTheDocument();
      expect(screen.getByText(/Welcome, Alex/i)).toBeInTheDocument();
    });

    expect(screen.queryByText('My Bookings')).not.toBeInTheDocument();
    expect(screen.queryByText('Profile')).not.toBeInTheDocument();
    expect(screen.queryByText('Admin Dashboard')).not.toBeInTheDocument();
  });

  test('shows admin links when logged in as admin', async () => {
    setLoggedInUser({
      id: 3,
      firstName: 'Jordan',
      role: 'admin',
    });

    renderNavbar();

    await waitFor(() => {
      expect(screen.getByText('Admin Dashboard')).toBeInTheDocument();
      expect(screen.getByText('Logout')).toBeInTheDocument();
      expect(screen.getByText(/Welcome, Jordan/i)).toBeInTheDocument();
    });

    expect(screen.queryByText('My Bookings')).not.toBeInTheDocument();
    expect(screen.queryByText('Profile')).not.toBeInTheDocument();
    expect(screen.queryByText('Staff Dashboard')).not.toBeInTheDocument();
  });

  test('logout clears storage and navigates home', async () => {
    setLoggedInUser({
      id: 1,
      firstName: 'Keerthi',
      role: 'student',
    });

    renderNavbar();

    const logoutButton = await screen.findByText('Logout');
    fireEvent.click(logoutButton);

    expect(localStorage.getItem('token')).toBeNull();
    expect(localStorage.getItem('user')).toBeNull();
    expect(mockNavigate).toHaveBeenCalledWith('/');
  });

  test('home link points to root route', async () => {
    renderNavbar();
    const homeLink = await screen.findByRole('link', { name: 'Home' });
    expect(homeLink.getAttribute('href')).toBe('/');
  });
});