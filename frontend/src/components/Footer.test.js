import { render, screen } from '@testing-library/react';
import Footer from './Footer';

describe('Footer Component', () => {
  test('renders footer element', () => {
    render(<Footer />);
    const footer = screen.getByRole('contentinfo');
    expect(footer).toBeInTheDocument();
  });

  test('displays Smart Campus Services heading', () => {
    render(<Footer />);
    expect(screen.getByText('Smart Campus Services')).toBeInTheDocument();
  });

  test('displays mission statement', () => {
    render(<Footer />);
    expect(screen.getByText(/Making campus life easier/i)).toBeInTheDocument();
  });

  test('displays Quick Links section', () => {
    render(<Footer />);
    expect(screen.getByText('Quick Links')).toBeInTheDocument();
  });

  test('displays quick links with correct hrefs', () => {
    render(<Footer />);
    const homeLink = screen.getByRole('link', { name: /Home/i });
    const servicesLink = screen.getByRole('link', { name: /Services/i });
    const bookingsLink = screen.getByRole('link', { name: /Bookings/i });
    
    expect(homeLink).toHaveAttribute('href', '/');
    expect(servicesLink).toHaveAttribute('href', '/services');
    expect(bookingsLink).toHaveAttribute('href', '/bookings');
  });

  test('displays Contact section', () => {
    render(<Footer />);
    expect(screen.getByText('Contact')).toBeInTheDocument();
  });

  test('displays contact email and phone', () => {
    render(<Footer />);
    expect(screen.getByText('support@uf.edu')).toBeInTheDocument();
    expect(screen.getByText('(352) 392-3261')).toBeInTheDocument();
  });

  test('displays Follow Us section with social links', () => {
    render(<Footer />);
    expect(screen.getByText('Follow Us')).toBeInTheDocument();
    expect(screen.getByRole('link', { name: /Facebook/i })).toBeInTheDocument();
    expect(screen.getByRole('link', { name: /Twitter/i })).toBeInTheDocument();
    expect(screen.getByRole('link', { name: /Instagram/i })).toBeInTheDocument();
  });

  test('displays copyright text', () => {
    render(<Footer />);
    expect(screen.getByText(/2026 Smart Campus Services Platform/i)).toBeInTheDocument();
  });
});
