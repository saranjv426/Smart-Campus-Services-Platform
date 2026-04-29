import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';
import Login from './Login';

jest.mock('axios');
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => jest.fn(),
}));

describe('Login Page Component', () => {
  const mockNavigate = jest.fn();
  const mockResponse = {
    data: {
      token: 'mock-token-123',
      id: 1,
      email: 'test@example.com',
      firstName: 'Test',
      role: 'student',
    },
  };

  const renderLogin = () => {
    render(
      <BrowserRouter>
        <Login />
      </BrowserRouter>
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
    axios.post.mockResolvedValue(mockResponse);
    localStorage.clear();
  });

  test('renders login form', () => {
    renderLogin();
    expect(screen.getByText('Login')).toBeInTheDocument();
  });

  test('displays email input field', () => {
    renderLogin();
    const emailInput = screen.getByPlaceholderText(/email/i);
    expect(emailInput).toBeInTheDocument();
  });

  test('displays password input field', () => {
    renderLogin();
    const passwordInput = screen.getByPlaceholderText(/password/i);
    expect(passwordInput).toBeInTheDocument();
  });

  test('updates email input value on change', async () => {
    renderLogin();
    const emailInput = screen.getByPlaceholderText(/email/i);
    await userEvent.type(emailInput, 'test@example.com');
    expect(emailInput).toHaveValue('test@example.com');
  });

  test('updates password input value on change', async () => {
    renderLogin();
    const passwordInput = screen.getByPlaceholderText(/password/i);
    await userEvent.type(passwordInput, 'password123');
    expect(passwordInput).toHaveValue('password123');
  });

  test('submits login form with valid credentials', async () => {
    renderLogin();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    await userEvent.type(passwordInput, 'password123');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(axios.post).toHaveBeenCalledWith(
        expect.stringContaining('/auth/login'),
        {
          email: 'test@example.com',
          password: 'password123',
        }
      );
    });
  });

  test('stores token in localStorage on successful login', async () => {
    renderLogin();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    await userEvent.type(passwordInput, 'password123');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(localStorage.getItem('token')).toBe('mock-token-123');
    });
  });

  test('stores user data in localStorage on successful login', async () => {
    renderLogin();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    await userEvent.type(passwordInput, 'password123');
    fireEvent.click(submitButton);

    await waitFor(() => {
      const userData = JSON.parse(localStorage.getItem('user'));
      expect(userData.email).toBe('test@example.com');
    });
  });

  test('displays error message on login failure', async () => {
    axios.post.mockRejectedValueOnce(new Error('Invalid credentials'));
    renderLogin();

    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    await userEvent.type(passwordInput, 'wrongpassword');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(screen.getByText(/Invalid email or password/i)).toBeInTheDocument();
    });
  });

  test('displays register link', () => {
    renderLogin();
    const registerLink = screen.getByRole('link', { name: /register/i });
    expect(registerLink).toBeInTheDocument();
  });

  test('dispatches authChange event on successful login', async () => {
    const dispatchEventSpy = jest.spyOn(window, 'dispatchEvent');
    renderLogin();

    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    await userEvent.type(passwordInput, 'password123');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(dispatchEventSpy).toHaveBeenCalledWith(
        expect.objectContaining({ type: 'authChange' })
      );
    });

    dispatchEventSpy.mockRestore();
  });

  test('prevents login submission with empty email', async () => {
    renderLogin();
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(passwordInput, 'password123');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(axios.post).not.toHaveBeenCalled();
    });
  });

  test('prevents login submission with empty password', async () => {
    renderLogin();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(axios.post).not.toHaveBeenCalled();
    });
  });

  test('validates email format before submission', async () => {
    renderLogin();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'invalid-email');
    await userEvent.type(passwordInput, 'password123');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(axios.post).not.toHaveBeenCalled();
    });
  });

  test('displays error message on 400 API error', async () => {
    const error = new Error('Bad Request');
    error.response = { status: 400 };
    axios.post.mockRejectedValueOnce(error);
    renderLogin();

    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    await userEvent.type(passwordInput, 'wrongpassword');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(screen.getByText(/Invalid email or password/i)).toBeInTheDocument();
    });
  });

  test('displays error message on 500 API error', async () => {
    const error = new Error('Server Error');
    error.response = { status: 500 };
    axios.post.mockRejectedValueOnce(error);
    renderLogin();

    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    await userEvent.type(passwordInput, 'password123');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(screen.getByText(/Server error|something went wrong/i)).toBeInTheDocument();
    });
  });

  test('disables submit button during API call', async () => {
    axios.post.mockImplementation(() =>
      new Promise(resolve => setTimeout(() => resolve(mockResponse), 200))
    );

    renderLogin();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    await userEvent.type(passwordInput, 'password123');
    fireEvent.click(submitButton);

    // Button should be disabled immediately after click
    expect(submitButton).toBeDisabled();

    await waitFor(() => {
      expect(axios.post).toHaveBeenCalled();
    });
  });

  test('clears error message when user modifies input', async () => {
    axios.post.mockRejectedValueOnce(new Error('Invalid credentials'));
    renderLogin();

    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    await userEvent.type(passwordInput, 'wrongpassword');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(screen.getByText(/Invalid email or password/i)).toBeInTheDocument();
    });

    // Modify input
    fireEvent.change(passwordInput, { target: { value: 'password123' } });

    await waitFor(() => {
      expect(screen.queryByText(/Invalid email or password/i)).not.toBeInTheDocument();
    });
  });

  test('stores user role in localStorage on successful login', async () => {
    const staffResponse = {
      data: {
        token: 'mock-token-123',
        id: 2,
        email: 'staff@example.com',
        firstName: 'Staff',
        role: 'staff',
      },
    };
    axios.post.mockResolvedValue(staffResponse);

    renderLogin();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'staff@example.com');
    await userEvent.type(passwordInput, 'password123');
    fireEvent.click(submitButton);

    await waitFor(() => {
      const userData = JSON.parse(localStorage.getItem('user'));
      expect(userData.role).toBe('staff');
    });
  });

  test('handles network error gracefully', async () => {
    axios.post.mockRejectedValueOnce(new Error('Network Error'));
    renderLogin();

    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /login/i });

    await userEvent.type(emailInput, 'test@example.com');
    await userEvent.type(passwordInput, 'password123');
    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(screen.getByText(/Network|connection|error/i)).toBeInTheDocument();
    });
  });
});
