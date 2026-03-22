import { render, screen, waitFor } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';
import Home from '../pages/Home';

jest.mock('axios');

describe('Home Page Component', () => {
  const mockServices = [
    {
      id: 1,
      name: 'Main Library',
      category: 'library',
      description: 'Central library',
      imageUrl: 'https://images.pexels.com/photos/1234/pexels.jpeg',
    },
    {
      id: 2,
      name: 'Student Dining',
      category: 'dining',
      description: 'Main dining hall',
      imageUrl: 'https://images.pexels.com/photos/5678/pexels.jpeg',
    },
  ];

  const renderHome = () => {
    render(
      <BrowserRouter>
        <Home />
      </BrowserRouter>
    );
  };

  beforeEach(() => {
    jest.clearAllMocks();
    axios.get.mockResolvedValue({ data: mockServices });
  });

  test('renders home page', () => {
    renderHome();
    expect(screen.getByText(/Welcome/i)).toBeInTheDocument();
  });

  test('displays welcome message', () => {
    renderHome();
    const heading = screen.getByRole('heading', { level: 1 });
    expect(heading).toBeInTheDocument();
  });

  test('renders category buttons', () => {
    renderHome();
    expect(screen.getByText(/Library/i)).toBeInTheDocument();
    expect(screen.getByText(/Dining/i)).toBeInTheDocument();
  });

  test('loads services on mount', async () => {
    renderHome();
    await waitFor(() => {
      expect(axios.get).toHaveBeenCalledWith('http://localhost:8080/api/services');
    });
  });
});
