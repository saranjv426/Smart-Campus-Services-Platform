# Smart Campus Services Platform - Setup Guide

## Project Overview

A full-stack web application for managing university campus services. Built with:
- **Frontend**: React 18 with React Router and Axios
- **Backend**: Go with Gin framework and GORM ORM
- **Database**: SQLite

## File Structure Created

### Backend (`/backend`)
```
├── main.go                  # Entry point with routes and middleware
├── go.mod                   # Go modules file
├── config/
│   └── config.go           # Configuration constants
├── models/
│   └── models.go           # Database models (User, Service, Booking, Review, Notification)
├── handlers/
│   ├── auth.go             # Authentication handlers (Register, Login, Logout)
│   ├── service.go          # Service handlers (CRUD operations)
│   ├── booking.go          # Booking handlers (Create, List, Update, Cancel)
│   ├── notification.go     # Notification handlers
│   ├── review.go           # Review handlers with rating system
│   └── user.go             # User profile handlers
├── .env.example            # Environment variables template
├── .gitignore              # Git ignore rules
└── README.md               # Backend documentation
```

### Frontend (`/frontend`)
```
├── public/
│   └── index.html          # HTML template
├── src/
│   ├── components/
│   │   ├── Navbar.js       # Navigation bar with auth links
│   │   └── Footer.js       # Footer component
│   ├── pages/
│   │   ├── Home.js         # Landing page with features showcase
│   │   ├── Services.js     # Service listing with search & filters
│   │   ├── ServiceDetail.js # Service details with booking & reviews
│   │   ├── Bookings.js     # User bookings management
│   │   ├── Profile.js      # User profile page
│   │   ├── Login.js        # Login page
│   │   └── Register.js     # Registration page
│   ├── services/
│   │   └── api.js          # API client with all endpoints
│   ├── styles/
│   │   ├── index.css       # Global styles & variables
│   │   ├── App.css         # App layout styles
│   │   ├── Navbar.css      # Navigation styles
│   │   ├── Footer.css      # Footer styles
│   │   ├── Home.css        # Landing page styles
│   │   ├── Services.css    # Services page styles
│   │   ├── ServiceDetail.css # Service detail styles
│   │   ├── Auth.css        # Auth pages styles
│   │   ├── Bookings.css    # Bookings page styles
│   │   └── Profile.css     # Profile page styles
│   ├── App.js              # Main app component with routing
│   └── index.js            # React entry point
├── package.json            # Dependencies and scripts
├── .env.example            # Environment template
├── .gitignore              # Git ignore rules
└── README.md               # Frontend documentation
```

### Root Files
```
├── README.md               # Project overview and setup guide
├── backend/                # Backend directory
└── frontend/               # Frontend directory
```

## Key Features Implemented

### Authentication & Users
- User registration with role selection (student/staff/faculty)
- Login with token-based authentication
- User profile management
- Profile viewing with booking/review history

### Services
- List all services with pagination
- Search services by name/description
- Filter by category (library, dining, transportation, etc.)
- View detailed service information
- Admin service management (create, update, delete)

### Bookings
- Create service bookings with date/time selection
- View user's booking history
- Filter bookings by status (pending, confirmed, completed, cancelled)
- Cancel pending bookings
- Automatic status management

### Reviews & Ratings
- Leave reviews with 1-5 star ratings
- View all reviews for a service
- Automatic average rating calculation
- User reviews display with names and dates

### Notifications
- Get user-specific notifications
- Mark notifications as read
- Support for multiple notification types

### UI/UX Features
- Responsive design (mobile-friendly)
- Modern color scheme (UF Gators blue & orange)
- Intuitive navigation
- Real-time form validation
- Loading and error states
- Empty state messages
- Smooth transitions and hover effects

## Database Models

### User
- UUID, email, password, name, phone, role, department, avatar, bio
- Relationships: bookings, reviews, notifications

### Service
- UUID, name, description, category, image, location, contact info, hours
- Rating calculated from reviews
- Relationships: bookings, reviews

### Booking
- UUID, user ID, service ID, status, start/end times, notes
- Status: pending, confirmed, completed, cancelled
- Relationships: user, service

### Review
- UUID, user ID, service ID, rating (1-5), comment
- Relationships: user, service

### Notification
- UUID, user ID, title, message, type, read status
- Types: booking, reminder, announcement
- Relationships: user

## API Endpoints Summary

**Auth**: `/api/auth/register`, `/api/auth/login`, `/api/auth/logout`, `/api/auth/refresh`

**Services**: `/api/services` (GET/POST), `/api/services/:id` (GET/PUT/DELETE), `/api/services/category/:category`

**Bookings**: `/api/bookings` (POST), `/api/bookings/:id` (GET/PUT/DELETE), `/api/bookings/user/:userId`

**Reviews**: `/api/reviews` (POST), `/api/reviews/:id` (GET/DELETE), `/api/reviews/service/:serviceId`

**Users**: `/api/users/:id` (GET/PUT), `/api/users/:id/profile`

**Notifications**: `/api/notifications/:userId` (GET), `/api/notifications` (POST), `/api/notifications/:id/read` (PUT)

## How to Run

### Start Backend
```bash
cd backend
go mod download
go run .

```
Note: Use `go run .` (not `go run main.go`) so all Go files in the package are compiled, including `seed.go`.
Backend runs on `http://localhost:8080`

### Start Frontend
```bash
cd frontend
npm install
npm start
```
Frontend runs on `http://localhost:3000`

### Setup Database
1. No external database install is required.
2. The SQLite database file is created automatically when the backend starts.
3. Migrations run automatically on startup.

## Environment Configuration

**Backend (.env)**
```
DB_PATH=data/smart_campus.db
PORT=8080
GIN_MODE=debug
```

**Frontend (.env)**
```
REACT_APP_API_URL=http://localhost:8080/api
```

## Next Steps

1. Install dependencies for both frontend and backend
2. Confirm the SQLite database path
3. Configure environment variables
4. Start the backend server
5. Start the frontend development server
6. Access the application at http://localhost:3000

## Features to Enhance

- JWT authentication with secure tokens
- Password hashing (bcrypt)
- Rate limiting
- Email notifications
- Real-time WebSocket updates
- Admin dashboard
- Payment integration
- Analytics and reporting
- Mobile app

## Color Scheme

- Primary: #0033a0 (UF Blue)
- Secondary: #ff6600 (UF Orange)
- Accent: #00ccff (Light Blue)
- Success: #4caf50 (Green)
- Warning: #ff9800 (Orange)
- Error: #f44336 (Red)

---

The complete application is now ready for development and deployment!
