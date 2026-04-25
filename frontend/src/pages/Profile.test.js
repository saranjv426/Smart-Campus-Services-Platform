
import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import Profile from './Profile';
import { userAPI, reviewAPI } from '../services/api';

jest.mock('../services/api', () => ({
  userAPI: {
    getProfile: jest.fn(),
    updateUser: jest.fn(),
  },
  reviewAPI: {
    getUserReviews: jest.fn(),
    deleteReview: jest.fn(),
  },
}));

describe('Profile Page Component', () => {
  const mockAlert = jest.fn();

  const mockUser = {
    id: 1,
    firstName: 'John',
    lastName: 'Doe',
    email: 'john@example.com',
    phone: '123-456-7890',
    role: 'student',
    department: 'Computer Science',
    bio: 'CS student at UF',
    avatarUrl: '',
    bookings: [
      {
        id: 1,
        status: 'approved',
        service: {
          name: 'Main Library',
        },
      },
      {
        id: 2,
        status: 'pending',
        service: {
          name: 'Study Room',
        },
      },
    ],
  };

  const mockReviews = [
    {
      id: 101,
      rating: 5,
      comment: 'Great library experience',
      createdAt: '2024-02-20T10:00:00Z',
      service: {
        name: 'Main Library',
      },
    },
    {
      id: 102,
      rating: 4,
      comment: 'Helpful staff and good resources',
      createdAt: '2024-02-21T10:00:00Z',
      service: {
        name: 'Study Room',
      },
    },
  ];

  const renderProfile = (route = '/profile/1') => {
    render(
      <MemoryRouter initialEntries={[route]}>
        <Routes>
          <Route path="/profile/:id" element={<Profile />} />
        </Routes>
      </MemoryRouter>
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();

    global.alert = mockAlert;
    window.alert = mockAlert;

    userAPI.getProfile.mockResolvedValue({ data: mockUser });
    userAPI.updateUser.mockResolvedValue({ data: { success: true } });
    reviewAPI.getUserReviews.mockResolvedValue({ data: mockReviews });
    reviewAPI.deleteReview.mockResolvedValue({ data: { success: true } });
  });

  test('renders loading state initially', () => {
    renderProfile();
    expect(screen.getByText(/loading profile/i)).toBeInTheDocument();
  });

  test('fetches profile on mount with route id', async () => {
    renderProfile();

    await waitFor(() => {
      expect(userAPI.getProfile).toHaveBeenCalledWith('1');
    });
  });

  test('displays user basic information', async () => {
    renderProfile();

    expect(await screen.findByText(/john doe/i)).toBeInTheDocument();
    expect(screen.getByText('STUDENT', { selector: 'p.role' })).toBeInTheDocument();
    expect(screen.getByText('john@example.com', { selector: 'p.email' })).toBeInTheDocument();
  });

  test('displays contact information', async () => {
    renderProfile();

    expect(await screen.findByText('123-456-7890')).toBeInTheDocument();
    expect(screen.getByText(/computer science/i)).toBeInTheDocument();
  });

  test('displays bio section', async () => {
    renderProfile();

    expect(await screen.findByText(/cs student at uf/i)).toBeInTheDocument();
  });

  test('displays recent bookings summary', async () => {
    renderProfile();

    expect(await screen.findByText(/recent bookings/i)).toBeInTheDocument();
    expect(screen.getByText(/main library/i)).toBeInTheDocument();
    expect(screen.getByText(/study room/i)).toBeInTheDocument();
    expect(screen.getByText(/approved/i)).toBeInTheDocument();
    expect(screen.getByText(/pending/i)).toBeInTheDocument();
  });

  test('fetches and displays user reviews', async () => {
    renderProfile();

    await waitFor(() => {
      expect(reviewAPI.getUserReviews).toHaveBeenCalledWith(1);
    });

    expect(await screen.findByText(/my reviews \(2\)/i)).toBeInTheDocument();
    expect(screen.getByText(/great library experience/i)).toBeInTheDocument();
    expect(screen.getByText(/helpful staff and good resources/i)).toBeInTheDocument();
  });

  test('handles empty reviews', async () => {
    reviewAPI.getUserReviews.mockResolvedValueOnce({ data: [] });

    renderProfile();

    expect(await screen.findByText(/you haven't written any reviews yet/i)).toBeInTheDocument();
  });

  test('opens edit mode when edit profile is clicked', async () => {
    renderProfile();

    expect(await screen.findByText(/john doe/i)).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button', { name: /edit profile/i }));

    expect(screen.getByDisplayValue('John')).toBeInTheDocument();
    expect(screen.getByDisplayValue('Doe')).toBeInTheDocument();
    expect(screen.getByDisplayValue('123-456-7890')).toBeInTheDocument();
    expect(screen.getByDisplayValue('Computer Science')).toBeInTheDocument();
    expect(screen.getByDisplayValue('CS student at UF')).toBeInTheDocument();
  });

  test('updates form fields in edit mode', async () => {
    renderProfile();

    expect(await screen.findByText(/john doe/i)).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button', { name: /edit profile/i }));

    const firstNameInput = screen.getByDisplayValue('John');
    fireEvent.change(firstNameInput, { target: { name: 'firstName', value: 'Johnny' } });

    expect(screen.getByDisplayValue('Johnny')).toBeInTheDocument();
  });

  test('saves updated profile successfully', async () => {
    renderProfile();

    expect(await screen.findByText(/john doe/i)).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button', { name: /edit profile/i }));

    fireEvent.change(screen.getByDisplayValue('John'), {
      target: { name: 'firstName', value: 'Johnny' },
    });

    fireEvent.click(screen.getByRole('button', { name: /save/i }));

    await waitFor(() => {
      expect(userAPI.updateUser).toHaveBeenCalledWith(
        '1',
        expect.objectContaining({
          firstName: 'Johnny',
        })
      );
    });

    await waitFor(() => {
      expect(mockAlert).toHaveBeenCalledWith('Profile updated successfully!');
    });
  });

  test('shows alert if name fields are empty', async () => {
    renderProfile();

    expect(await screen.findByText(/john doe/i)).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button', { name: /edit profile/i }));

    fireEvent.change(screen.getByDisplayValue('John'), {
      target: { name: 'firstName', value: '' },
    });

    fireEvent.click(screen.getByRole('button', { name: /save/i }));

    expect(mockAlert).toHaveBeenCalledWith('Name fields cannot be empty');
    expect(userAPI.updateUser).not.toHaveBeenCalled();
  });

  test('cancels edit mode', async () => {
    renderProfile();

    expect(await screen.findByText(/john doe/i)).toBeInTheDocument();

    fireEvent.click(screen.getByRole('button', { name: /edit profile/i }));
    fireEvent.click(screen.getByRole('button', { name: /cancel/i }));

    await waitFor(() => {
      expect(screen.getByRole('button', { name: /edit profile/i })).toBeInTheDocument();
    });
  });

  test('shows error message when profile fetch fails', async () => {
    userAPI.getProfile.mockRejectedValueOnce(new Error('Failed to load'));

    renderProfile();

    expect(await screen.findByText(/failed to load profile/i)).toBeInTheDocument();
  });

  test('opens delete confirmation modal for a review', async () => {
    renderProfile();

    expect(await screen.findByText(/great library experience/i)).toBeInTheDocument();

    fireEvent.click(screen.getAllByRole('button', { name: /delete review/i })[0]);

    expect(screen.getByText(/delete review\?/i)).toBeInTheDocument();
    expect(
      screen.getByText(/are you sure you want to delete this review/i)
    ).toBeInTheDocument();
  });

  test('deletes a review successfully', async () => {
    renderProfile();

    expect(await screen.findByText(/great library experience/i)).toBeInTheDocument();

    fireEvent.click(screen.getAllByRole('button', { name: /delete review/i })[0]);
    fireEvent.click(screen.getByRole('button', { name: /^delete$/i }));

    await waitFor(() => {
      expect(reviewAPI.deleteReview).toHaveBeenCalledWith(101);
    });

    await waitFor(() => {
      expect(mockAlert).toHaveBeenCalledWith('Review deleted successfully!');
    });
  });
});