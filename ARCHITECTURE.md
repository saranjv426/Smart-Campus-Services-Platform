# Smart Campus Services Platform - Architecture Documentation

## System Architecture Overview

This document describes the architecture of the Smart Campus Services Platform.

## Architecture Pattern: MVC + RESTful API

### Backend Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                      Frontend (React)                        │
├─────────────────────────────────────────────────────────────┤
│                                                               │
│  ┌──────────────────────────────────────────────────────┐  │
│  │  HTTP Client (Axios)                                 │  │
│  └───────────────────────┬────────────────────────────┬─┘  │
│                          │                            │     │
└──────────────────────────┼────────────────────────────┼─────┘
                           │ REST API                   │
        ┌──────────────────▼──────────────┬────────────▼────────┐
        │                                  │                     │
┌───────▼──────────────────────────────────▼──────────────────┐
│            Go Backend (Gin Framework)                      │
├──────────────────────────────────────────────────────────────┤
│                                                               │
│  ┌─────────────────────────────────────────────────────┐   │
│  │  Middleware Layer                                    │   │
│  │  - CORS Handling                                     │   │
│  │  - Error Handling                                    │   │
│  │  - Logging                                           │   │
│  └────────────────┬────────────────────────────────────┘   │
│                   │                                          │
│  ┌────────────────▼────────────────────────────────────┐   │
│  │  Routes Layer                                        │   │
│  │  - /api/auth/*                                       │   │
│  │  - /api/services/*                                   │   │
│  │  - /api/bookings/*                                   │   │
│  │  - /api/reviews/*                                    │   │
│  │  - /api/users/*                                      │   │
│  │  - /api/notifications/*                              │   │
│  └────────────────┬────────────────────────────────────┘   │
│                   │                                          │
│  ┌────────────────▼────────────────────────────────────┐   │
│  │  Handlers Layer                                      │   │
│  │  - AuthHandler                                       │   │
│  │  - ServiceHandler                                    │   │
│  │  - BookingHandler                                    │   │
│  │  - ReviewHandler                                     │   │
│  │  - UserHandler                                       │   │
│  │  - NotificationHandler                               │   │
│  └────────────────┬────────────────────────────────────┘   │
│                   │                                          │
│  ┌────────────────▼────────────────────────────────────┐   │
│  │  Models Layer (GORM)                                 │   │
│  │  - User                                              │   │
│  │  - Service                                           │   │
│  │  - Booking                                           │   │
│  │  - Review                                            │   │
│  │  - Notification                                      │   │
│  └────────────────┬────────────────────────────────────┘   │
│                   │                                          │
└───────────────────┼──────────────────────────────────────────┘
                    │ SQL Queries
        ┌───────────▼──────────────┐
        │                          │
┌───────▼──────────────────────────▼───────┐
│      SQLite Database                    │
│  ├─ users table                         │
│  ├─ services table                      │
│  ├─ bookings table                      │
│  ├─ reviews table                       │
│  └─ notifications table                 │
└────────────────────────────────────────┘
```

## Frontend Architecture

### Component Structure

```
App (Root)
├── Router Setup
├── Navbar
│   ├── Logo & Links
│   └── Auth Menu
├── Main Content
│   ├── Home Page
│   │   ├── Hero Section
│   │   ├── Features Grid
│   │   ├── Benefits Section
│   │   └── Testimonials
│   ├── Services Page
│   │   ├── Search Bar
│   │   ├── Category Filter
│   │   └── Service Grid
│   ├── Service Detail Page
│   │   ├── Service Info
│   │   ├── Booking Form
│   │   └── Reviews Section
│   ├── Bookings Page
│   │   ├── Status Filter
│   │   └── Booking Cards
│   ├── Login Page
│   │   └── Login Form
│   ├── Register Page
│   │   └── Registration Form
│   └── Profile Page
│       ├── Profile Header
│       ├── Contact Info
│       ├── Booking History
│       └── Reviews Summary
└── Footer
    ├── Links
    ├── Contact Info
    └── Social Links
```

### Data Flow

```
User Interaction
       │
       ▼
Component State Update (React Hooks)
       │
       ▼
API Call (axios via api.js)
       │
       ▼
Backend Endpoint
       │
       ▼
Handler Function
       │
       ▼
GORM Database Query
       │
       ▼
SQLite Database
       │
       ▼
Response JSON
       │
       ▼
Frontend State Update
       │
       ▼
Component Re-render
       │
       ▼
User Sees Update
```

## Database Schema

### Users Table
```sql
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  first_name VARCHAR(100),
  last_name VARCHAR(100),
  phone VARCHAR(20),
  role VARCHAR(50),
  department VARCHAR(100),
  avatar_url TEXT,
  bio TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
```

### Services Table
```sql
CREATE TABLE services (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  category VARCHAR(100),
  image_url TEXT,
  location VARCHAR(255),
  phone VARCHAR(20),
  email VARCHAR(255),
  hours VARCHAR(255),
  rating FLOAT,
  is_active BOOLEAN,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
```

### Bookings Table
```sql
CREATE TABLE bookings (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  service_id UUID REFERENCES services(id),
  status VARCHAR(50),
  start_time TIMESTAMP,
  end_time TIMESTAMP,
  notes TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
```

### Reviews Table
```sql
CREATE TABLE reviews (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  service_id UUID REFERENCES services(id),
  rating INTEGER,
  comment TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
```

### Notifications Table
```sql
CREATE TABLE notifications (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  title VARCHAR(255),
  message TEXT,
  type VARCHAR(50),
  is_read BOOLEAN,
  created_at TIMESTAMP
);
```

## API Response Format

### Success Response
```json
{
  "data": { /* resource data */ },
  "status": 200,
  "message": "Success"
}
```

### Error Response
```json
{
  "error": "Error message",
  "status": 400,
  "message": "Bad Request"
}
```

## Authentication Flow

```
User Registration/Login
       │
       ▼
POST /api/auth/register or /api/auth/login
       │
       ▼
Validate Credentials
       │
       ▼
Create/Verify User
       │
       ▼
Generate Token
       │
       ▼
Return Token to Frontend
       │
       ▼
Store Token in localStorage
       │
       ▼
Include Token in API Headers
       │
       ▼
Backend Validates Token
       │
       ▼
Allow/Deny Request
```

## Deployment Architecture

### Development Environment
```
Localhost:3000 (Frontend)
       │
       └─────────► Localhost:8080 (Backend)
                        │
                        └─────────► SQLite (data/smart_campus.db)
```

### Production Environment (Suggested)
```
CDN / Static Host (Frontend)
       │
       └─────────► Cloud Server/Container (Backend)
                        │
                        └─────────► Managed Database (SQLite)
```

## Scalability Considerations

### Database
- Use connection pooling (GORM supports this)
- Add database indexes on frequently queried columns
- Consider read replicas for read-heavy operations

### Backend
- Implement horizontal scaling with load balancer
- Use Docker containers for easy deployment
- Add caching layer (Redis) for frequently accessed data
- Implement pagination for large datasets

### Frontend
- Deploy to CDN for global distribution
- Use code splitting and lazy loading
- Implement service worker for offline capability
- Optimize images and assets

## Security Measures

### Implemented
- CORS middleware
- Environment variables for sensitive data
- UUID for resource IDs (not sequential)

### Recommended for Production
- JWT tokens with expiration
- Password hashing (bcrypt)
- HTTPS/TLS encryption
- Rate limiting
- Input validation and sanitization
- SQL injection prevention (GORM handles this)
- XSS protection
- CSRF tokens
- Database encryption at rest

## Error Handling Strategy

```
API Request
    │
    ├─► Validation Error → 400 Bad Request
    │
    ├─► Authentication Error → 401 Unauthorized
    │
    ├─► Authorization Error → 403 Forbidden
    │
    ├─► Not Found → 404 Not Found
    │
    ├─► Server Error → 500 Internal Server Error
    │
    └─► Success → 200/201 with data
```

## Performance Optimizations

1. **Frontend**
   - Component memoization
   - Lazy loading routes
   - Image optimization
   - CSS minification

2. **Backend**
   - Database query optimization
   - Connection pooling
   - Response caching
   - Efficient JSON serialization

3. **Database**
   - Indexes on foreign keys
   - Denormalization where needed
   - Query optimization

## Monitoring & Logging

### Suggested Implementation
- Application logs (structured logging)
- Error tracking (Sentry, etc.)
- Performance monitoring (APM)
- Database query logging
- API usage analytics

## Configuration Management

### Environment Variables
- Database credentials
- API keys
- Server port
- Frontend URL
- Debug mode

### Configuration Files
- `.env` for local development
- `.env.production` for production
- `config/` directory for constants

## Data Consistency

### Transaction Safety
- GORM transactions for related operations
- Atomic operations for critical updates
- Proper error handling and rollback

### Data Validation
- Input validation in handlers
- Model-level constraints
- Database constraints

## Caching Strategy

### Current Implementation
- localStorage for user tokens and state

### Recommended Additions
- Redis for session management
- API response caching
- Browser caching headers
- Database query caching

---

This architecture provides a solid foundation for a scalable, maintainable campus services platform.
