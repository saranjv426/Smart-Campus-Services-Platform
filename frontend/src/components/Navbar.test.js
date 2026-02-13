import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import Navbar from './components/Navbar';

describe('Navbar Component', () => {
  const renderNavbar = () => {
    render(
      <BrowserRouter>
        <Navbar />
      </BrowserRouter>
    );
  };

  test('renders navigation bar', () => {
    renderNavbar();
    const nav = screen.getByRole('navigation');
    expect(nav).toBeInTheDocument();
  });

  test('displays Home link', () => {
    renderNavbar();
    expect(screen.getByText(/Home/i)).toBeInTheDocument();
  });

  test('displays Services link', () => {
    renderNavbar();
    expect(screen.getByText(/Services/i)).toBeInTheDocument();
  });

  test('displays Bookings link', () => {
    renderNavbar();
    expect(screen.getByText(/Bookings/i)).toBeInTheDocument();
  });

  test('navigation links are clickable', () => {
    renderNavbar();
    const homeLink = screen.getByRole('link', { name: /Home/i });
    expect(homeLink).toBeInTheDocument();
    expect(homeLink.href).toContain('/');
  });
});
