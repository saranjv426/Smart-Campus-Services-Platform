import { render, screen } from '@testing-library/react';
import App from './App';

describe('App Component', () => {
  test('renders without crashing', () => {
    render(<App />);
    const appElement = screen.getByRole('navigation');
    expect(appElement).toBeInTheDocument();
  });

  test('renders main navigation', () => {
    render(<App />);
    expect(screen.getByText(/Home/i)).toBeInTheDocument();
  });

  test('has router setup', () => {
    render(<App />);
    // Check that main app structure is present
    const mainContent = screen.getByRole('main', { hidden: true });
    expect(mainContent).toBeInTheDocument();
  });
});
