# Smart Campus Services Platform

A comprehensive full-stack web application for managing campus services at the University of Florida. This platform allows students and staff to discover, book, and manage various campus services including library services, dining, transportation, maintenance, and more.

## Features

### For Students and Staff
- **Service Discovery**: Browse and search campus services by category
- **Booking System**: Reserve services with intuitive date/time selection
- **Service Reviews**: Read and write reviews with ratings for services
- **Personal Bookings**: Manage your current and past bookings
- **User Profile**: Update profile information and view booking history
- **Notifications**: Get updates about your bookings and service availability

### For Administrators
- **Service Management**: Add, update, and delete campus services
- **Booking Management**: Track and manage service bookings
- **Analytics**: View service ratings and user reviews
- **System Monitoring**: Track service availability and usage

## Technology Stack

### Frontend
- **Framework**: React 18
- **Routing**: React Router v6
- **HTTP Client**: Axios
- **Styling**: CSS3 with CSS Variables
- **State Management**: React Hooks

### Backend
- **Language**: Go 1.21+
- **Framework**: Gin (Web Framework)
- **ORM**: GORM
- **Database**: SQLite
- **Authentication**: Token-based (JWT-ready)

### Database
- **SQLite**: Embedded relational database for data persistence
- **Models**: User, Service, Booking, Review, Notification

## Quick Start

### Prerequisites
- Node.js 14+ (for frontend)
- Go 1.21+ (for backend)
- SQLite (embedded, no server required)
- Git

### Setup Backend

1. Navigate to backend directory:
   ```bash
   cd backend
   ```

2. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your SQLite database path
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the backend:
   ```bash
   go run main.go
   ```

   Backend will be available at `http://localhost:8080`

### Setup Frontend

1. Navigate to frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Create environment file:
   ```bash
   echo "REACT_APP_API_URL=http://localhost:8080/api" > .env
   ```

4. Start the development server:
   ```bash
   npm start
   ```

   Frontend will be available at `http://localhost:3000`

## API Endpoints

### Authentication
```
POST   /api/auth/register    - Register new user
POST   /api/auth/login       - Login user
POST   /api/auth/logout      - Logout user
POST   /api/auth/refresh     - Refresh token
```

### Services
```
GET    /api/services         - List all services
GET    /api/services/:id     - Get service details
POST   /api/services         - Create service (Admin)
PUT    /api/services/:id     - Update service (Admin)
DELETE /api/services/:id     - Delete service (Admin)
GET    /api/services/category/:category - Get by category
```

### Bookings
```
POST   /api/bookings         - Create booking
GET    /api/bookings/:id     - Get booking details
GET    /api/bookings/user/:userId - Get user bookings
PUT    /api/bookings/:id     - Update booking
DELETE /api/bookings/:id     - Cancel booking
```

### Reviews
```
POST   /api/reviews          - Create review
GET    /api/reviews/service/:serviceId - Get service reviews
GET    /api/reviews/:id      - Get review details
DELETE /api/reviews/:id      - Delete review
```

### Users
```
GET    /api/users/:id        - Get user details
PUT    /api/users/:id        - Update user profile
GET    /api/users/:id/profile - Get full profile
```

### Notifications
```
GET    /api/notifications/:userId - Get user notifications
POST   /api/notifications    - Create notification
PUT    /api/notifications/:id/read - Mark as read
```

## Service Categories

- 📚 **Library Services** - Study rooms, resource access
- 🍽️ **Dining Services** - Cafeterias, meal plans
- 🚌 **Transportation** - Campus shuttles, parking
- 🔧 **Maintenance** - Facility repairs, support
- ⚕️ **Health Services** - Medical appointments, wellness
- 🎓 **Academic Support** - Tutoring, advising
- 🏋️ **Recreation** - Gym, sports facilities
- 🎭 **Events** - Campus activities, workshops

## Configuration

### Environment Variables

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

## Security Considerations

- [ ] Implement JWT authentication
- [ ] Hash passwords with bcrypt
- [ ] Add rate limiting
- [ ] Implement CORS properly
- [ ] Validate all inputs
- [ ] Use HTTPS in production

## Future Enhancements

- Real-time notifications with WebSockets
- Email notifications
- Payment integration
- Admin dashboard
- Advanced reporting
- Mobile app
- Service provider portal

---

**Last Updated**: February 1, 2026
**Version**: 1.0.0

# Project Members
- Venkata Sai Saran Jonnalagadda
- Srikar Panuganti
- Keerthi Reddy Gudibandi
- Vishnu Sai Padyala
