# Smart Campus Services Platform - Frontend

## Setup Instructions

### Prerequisites
- Node.js (v14 or higher)
- npm or yarn

### Installation

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Create a `.env` file in the frontend directory:
   ```bash
   REACT_APP_API_URL=http://localhost:8080/api
   ```

### Running the Application

Start the development server:
```bash
npm start
```

The application will open at `http://localhost:3000`

### Building for Production

```bash
npm run build
```

This creates an optimized production build in the `build` folder.

## Project Structure

```
frontend/
├── public/
│   └── index.html
├── src/
│   ├── components/
│   │   ├── Navbar.js
│   │   └── Footer.js
│   ├── pages/
│   │   ├── Home.js
│   │   ├── Services.js
│   │   ├── ServiceDetail.js
│   │   ├── Bookings.js
│   │   ├── Profile.js
│   │   ├── Login.js
│   │   └── Register.js
│   ├── services/
│   │   └── api.js
│   ├── styles/
│   │   ├── index.css
│   │   ├── App.css
│   │   ├── Navbar.css
│   │   ├── Footer.css
│   │   ├── Home.css
│   │   ├── Services.css
│   │   ├── ServiceDetail.css
│   │   ├── Auth.css
│   │   ├── Bookings.css
│   │   └── Profile.css
│   ├── App.js
│   └── index.js
├── package.json
└── .env.example
```

## Features

- **Browse Services**: View all campus services with search and filtering
- **Service Details**: See detailed information, reviews, and ratings
- **Bookings**: Create, manage, and cancel service bookings
- **User Profile**: Manage your profile information and view booking history
- **Reviews**: Leave reviews and ratings for services
- **Authentication**: Register and login to your account
- **Real-time Updates**: Get notified about your bookings and services

## API Integration

The application connects to the backend API running on `http://localhost:8080`. Make sure the backend is running before starting the frontend.

## Available Scripts

- `npm start`: Runs the app in development mode
- `npm run build`: Builds the app for production
- `npm test`: Launches the test runner
- `npm run eject`: Ejects from create-react-app (irreversible)

## Technologies Used

- React 18
- React Router v6
- Axios
- CSS3 with CSS Variables
- Responsive Design

## Notes

- Token-based authentication is implemented using localStorage
- All forms include validation and error handling
- The app is fully responsive for mobile and desktop
- UF Gators color scheme is used throughout the application
