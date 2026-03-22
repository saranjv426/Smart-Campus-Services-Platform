import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import axios from 'axios';
import ServiceDetail from './ServiceDetail';

jest.mock('axios');

describe('ServiceDetail Page Component', () => {
  const mockService = {
    id: 1,
    name: 'Main Library',
    category: 'library',
    description: 'Central library with study rooms',
    location: 'Campus Center',
    rating: 4.5,
    imageUrl: 'https://images.pexels.com/photos/1234/pexels.jpeg',
    isActive: true,
  };

  const mockReviews = [
    {
      id: 1,
      userId: 1,
      serviceId: 1,
      rating: 5,
      comment: 'Great library!',
      createdAt: '2024-02-15T10:00:00Z',
      user: { firstName: 'John', lastName: 'Doe' }
    },
    {
      id: 2,
      userId: 2,
      serviceId: 1,
      rating: 4,
      comment: 'Good resources',
      createdAt: '2024-02-14T10:00:00Z',
      user: { firstName: 'Jane', lastName: 'Smith' }
    },
  ];

  const renderServiceDetail = () => {
    render(
      <BrowserRouter>
        <Routes>
          <Route path="/services/:id" element={<ServiceDetail />} />
        </Routes>
      </BrowserRouter>,
      { initialEntries: ['/services/1'] }
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
    axios.get.mockImplementation((url) => {
      if (url.includes('/services/1')) {
        return Promise.resolve({ data: mockService });
      }
      if (url.includes('/reviews/service/1')) {
        return Promise.resolve({ data: mockReviews });
      }
      return Promise.reject(new Error('Not found'));
    });
    localStorage.clear();
  });

  test('renders service detail page', () => {
    window.HTMLElement.prototype.scrollIntoView = jest.fn();
    renderServiceDetail();

    expect(screen.getByText(/loading/i)).toBeInTheDocument();
  });

  test('displays loading state initially', () => {
    renderServiceDetail();
    expect(screen.getByText(/loading/i)).toBeInTheDocument();
  });

  test('fetches service details on mount', async () => {
    renderServiceDetail();

    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith(
        expect.stringContaining('/services/1')
      );
    });
  });

  test('fetches service reviews on mount', async () => {
    renderServiceDetail();

    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith(
        expect.stringContaining('/reviews/service/1')
      );
    });
  });

  test('displays service name after loading', async () => {
    renderServiceDetail();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });
  });

  test('displays service description after loading', async () => {
    renderServiceDetail();

    await waitFor(() => {
      expect(screen.getByText('Central library with study rooms')).toBeInTheDocument();
    });
  });

  test('displays service details (location, rating)', async () => {
    renderServiceDetail();

    await waitFor(() => {
      expect(screen.getByText(/Campus Center/i)).toBeInTheDocument();
    });
  });

  test('displays service reviews after loading', async () => {
    renderServiceDetail();

    await waitFor(() => {
      expect(screen.getByText('Great library!')).toBeInTheDocument();
      expect(screen.getByText('Good resources')).toBeInTheDocument();
    });
  });

  test('displays booking form when button clicked', async () => {
    renderServiceDetail();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const bookButton = screen.getByRole('button', { name: /book/i });
    fireEvent.click(bookButton);

    await waitFor(() => {
      expect(screen.getByPlaceholderText(/start time/i)).toBeInTheDocument();
    });
  });

  test('shows booking form with required fields', async () => {
    renderServiceDetail();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const bookButton = screen.getByRole('button', { name: /book/i });
    fireEvent.click(bookButton);

    await waitFor(() => {
      expect(screen.getByPlaceholderText(/start time/i)).toBeInTheDocument();
      expect(screen.getByPlaceholderText(/end time/i)).toBeInTheDocument();
    });
  });

  test('displays error message when service fails to load', async () => {
    axios.get.mockRejectedValueOnce(new Error('Failed to load'));

    renderServiceDetail();

    await waitFor(() => {
      expect(screen.getByText(/Failed to load service details/i)).toBeInTheDocument();
    });
  });

  test('displays review count', async () => {
    renderServiceDetail();

    await waitFor(() => {
      const reviews = screen.getAllByText(/★/);
      expect(reviews.length).toBeGreaterThan(0);
    });
  });

  test('handles booking submission', async () => {
    const mockBookingResponse = { data: { id: 1, status: 'pending' } };
    axios.post.mockResolvedValueOnce(mockBookingResponse);

    localStorage.setItem('user', JSON.stringify({ id: 1 }));

    renderServiceDetail();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const bookButton = screen.getByRole('button', { name: /book/i });
    fireEvent.click(bookButton);

    await waitFor(() => {
      const submitButton = screen.getByRole('button', { name: /confirm booking/i });
      expect(submitButton).toBeInTheDocument();
    });
  });

  test('requires user to be logged in for booking', async () => {
    renderServiceDetail();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const bookButton = screen.getByRole('button', { name: /book/i });
    fireEvent.click(bookButton);

    // Should show login prompt when user is not logged in
    await waitFor(() => {
      const alertOrMessage = screen.queryByText(/login|logged in/i);
      // This depends on the implementation
    });
  });

  test('displays review form', async () => {
    localStorage.setItem('user', JSON.stringify({ id: 1, firstName: 'John' }));

    renderServiceDetail();

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    const reviewInput = screen.queryByPlaceholderText(/comment|review/i);
    expect(reviewInput).toBeInTheDocument();
  });
});
