import { render, screen, waitFor, fireEvent } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';
import Services from '../pages/Services';

jest.mock('axios');

describe('Services Page Component', () => {
  const mockServices = [
    {
      id: 1,
      name: 'Main Library',
      category: 'library',
      description: 'Central library',
      location: 'Campus Center',
      rating: 4.5,
      imageUrl: 'https://images.pexels.com/photos/1234/pexels.jpeg',
      isActive: true,
    },
    {
      id: 2,
      name: 'Student Dining',
      category: 'dining',
      description: 'Main dining hall',
      location: 'Student Center',
      rating: 4.0,
      imageUrl: 'https://images.pexels.com/photos/5678/pexels.jpeg',
      isActive: true,
    },
    {
      id: 3,
      name: 'Campus Bus',
      category: 'transportation',
      description: 'Bus service',
      location: 'Main Gate',
      rating: 4.2,
      imageUrl: 'https://images.pexels.com/photos/9101/pexels.jpeg',
      isActive: true,
    },
  ];

  const renderServices = () => {
    render(
      <BrowserRouter>
        <Services />
      </BrowserRouter>
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
    axios.get.mockResolvedValue({ data: mockServices });
  });

  test('renders services page', async () => {
    renderServices();
    await waitFor(() => {
      expect(screen.getByText(/Services/i)).toBeInTheDocument();
    });
  });

  test('loads services on component mount', async () => {
    renderServices();
    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith('http://localhost:8080/api/services');
    });
  });

  test('displays service cards after loading', async () => {
    renderServices();
    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
      expect(screen.getByText('Student Dining')).toBeInTheDocument();
    });
  });

  test('displays category filter buttons', async () => {
    renderServices();
    await waitFor(() => {
      expect(screen.getByText(/All/i)).toBeInTheDocument();
      expect(screen.getByText(/Library/i)).toBeInTheDocument();
      expect(screen.getByText(/Dining/i)).toBeInTheDocument();
    });
  });

  test('filters services by category', async () => {
    renderServices();
    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });

    // Click library filter
    const libraryButton = screen.getByRole('button', { name: /Library/i });
    fireEvent.click(libraryButton);

    await waitFor(() => {
      expect(screen.getByText('Main Library')).toBeInTheDocument();
    });
  });

  test('displays service details', async () => {
    renderServices();
    await waitFor(() => {
      expect(screen.getByText('Central library')).toBeInTheDocument();
      expect(screen.getByText('Campus Center')).toBeInTheDocument();
    });
  });
});
