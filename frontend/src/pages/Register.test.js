import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';
import Register from './Register';

jest.mock('axios');
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => jest.fn(),
}));

describe('Register Page Component', () => {
  const mockNavigate = jest.fn();
  const mockResponse = {
    data: {
      token: 'mock-token-123',
      id: 1,
      email: 'newuser@example.com',
      firstName: 'New',
      lastName: 'User',
      role: 'student',
    },
  };

  const renderRegister = () => {
    render(
      <BrowserRouter>
        <Register />
      </BrowserRouter>
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
    axios.post.mockResolvedValue(mockResponse);
    localStorage.clear();
  });

  test('renders registration form', () => {
    renderRegister();
    expect(screen.getByText(/Register/i)).toBeInTheDocument();
  });

  test('displays email input field', () => {
    renderRegister();
    const emailInput = screen.getByPlaceholderText(/email/i);
    expect(emailInput).toBeInTheDocument();
  });

  test('displays password input field', () => {
    renderRegister();
    const passwordInput = screen.getByPlaceholderText(/password/i);
    expect(passwordInput).toBeInTheDocument();
  });

  test('displays first name input field', () => {
    renderRegister();
    const firstNameInput = screen.getByPlaceholderText(/first name/i);
    expect(firstNameInput).toBeInTheDocument();
  });

  test('displays last name input field', () => {
    renderRegister();
    const lastNameInput = screen.getByPlaceholderText(/last name/i);
    expect(lastNameInput).toBeInTheDocument();
  });

  test('displays phone input field', () => {
    renderRegister();
    const phoneInput = screen.getByPlaceholderText(/phone/i);
    expect(phoneInput).toBeInTheDocument();
  });

  test('displays role select dropdown', () => {
    renderRegister();
    const roleSelect = screen.getByDisplayValue('student');
    expect(roleSelect).toBeInTheDocument();
  });

  test('updates form fields on user input', async () => {
    renderRegister();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const firstNameInput = screen.getByPlaceholderText(/first name/i);
    const lastNameInput = screen.getByPlaceholderText(/last name/i);

    await userEvent.type(emailInput, 'newuser@example.com');
    await userEvent.type(firstNameInput, 'New');
    await userEvent.type(lastNameInput, 'User');

    expect(emailInput).toHaveValue('newuser@example.com');
    expect(firstNameInput).toHaveValue('New');
    expect(lastNameInput).toHaveValue('User');
  });

  test('submits registration form with valid data', async () => {
    renderRegister();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const firstNameInput = screen.getByPlaceholderText(/first name/i);
    const lastNameInput = screen.getByPlaceholderText(/last name/i);
    const phoneInput = screen.getByPlaceholderText(/phone/i);
    const submitButton = screen.getByRole('button', { name: /register/i });

    await userEvent.type(emailInput, 'newuser@example.com');
    await userEvent.type(passwordInput, 'password123');
    await userEvent.type(firstNameInput, 'New');
    await userEvent.type(lastNameInput, 'User');
    await userEvent.type(phoneInput, '1234567890');

    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(axios.post).toHaveBeenCalledWith(
        expect.stringContaining('/auth/register'),
        expect.objectContaining({
          email: 'newuser@example.com',
          password: 'password123',
          firstName: 'New',
          lastName: 'User',
          phone: '1234567890',
          role: 'student',
        })
      );
    });
  });

  test('stores token in localStorage on successful registration', async () => {
    renderRegister();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const firstNameInput = screen.getByPlaceholderText(/first name/i);
    const lastNameInput = screen.getByPlaceholderText(/last name/i);
    const submitButton = screen.getByRole('button', { name: /register/i });

    await userEvent.type(emailInput, 'newuser@example.com');
    await userEvent.type(passwordInput, 'password123');
    await userEvent.type(firstNameInput, 'New');
    await userEvent.type(lastNameInput, 'User');

    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(localStorage.getItem('token')).toBe('mock-token-123');
    });
  });

  test('stores user data in localStorage on successful registration', async () => {
    renderRegister();
    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const firstNameInput = screen.getByPlaceholderText(/first name/i);
    const lastNameInput = screen.getByPlaceholderText(/last name/i);
    const submitButton = screen.getByRole('button', { name: /register/i });

    await userEvent.type(emailInput, 'newuser@example.com');
    await userEvent.type(passwordInput, 'password123');
    await userEvent.type(firstNameInput, 'New');
    await userEvent.type(lastNameInput, 'User');

    fireEvent.click(submitButton);

    await waitFor(() => {
      const userData = JSON.parse(localStorage.getItem('user'));
      expect(userData.email).toBe('newuser@example.com');
    });
  });

  test('displays error message on registration failure', async () => {
    axios.post.mockRejectedValueOnce(new Error('Registration failed'));
    renderRegister();

    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const submitButton = screen.getByRole('button', { name: /register/i });

    await userEvent.type(emailInput, 'newuser@example.com');
    await userEvent.type(passwordInput, 'password123');

    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(screen.getByText(/Failed to register/i)).toBeInTheDocument();
    });
  });

  test('displays login link', () => {
    renderRegister();
    const loginLink = screen.getByRole('link', { name: /login/i });
    expect(loginLink).toBeInTheDocument();
  });

  test('dispatches authChange event on successful registration', async () => {
    const dispatchEventSpy = jest.spyOn(window, 'dispatchEvent');
    renderRegister();

    const emailInput = screen.getByPlaceholderText(/email/i);
    const passwordInput = screen.getByPlaceholderText(/password/i);
    const firstNameInput = screen.getByPlaceholderText(/first name/i);
    const lastNameInput = screen.getByPlaceholderText(/last name/i);
    const submitButton = screen.getByRole('button', { name: /register/i });

    await userEvent.type(emailInput, 'newuser@example.com');
    await userEvent.type(passwordInput, 'password123');
    await userEvent.type(firstNameInput, 'New');
    await userEvent.type(lastNameInput, 'User');

    fireEvent.click(submitButton);

    await waitFor(() => {
      expect(dispatchEventSpy).toHaveBeenCalledWith(
        expect.objectContaining({ type: 'authChange' })
      );
    });

    dispatchEventSpy.mockRestore();
  });
});
