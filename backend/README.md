# Smart Campus Services Platform - Backend

## Setup Instructions

### Prerequisites
- Go 1.21 or higher
- SQLite (embedded, no server required)
- Git

### Installation

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Download Go dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the backend directory:
   ```bash
   DB_PATH=data/smart_campus.db
   PORT=8080
   GIN_MODE=debug
   ```

### Database Setup

1. No external database setup is required.
2. The application will automatically create the SQLite database file and run migrations when started.

### Running the Application

```bash
go run main.go
```

The API will be available at `http://localhost:8080`

### Building for Production

```bash
go build -o smart-campus-backend
./smart-campus-backend
```

## Project Structure

```
backend/
├── config/
│   └── config.go
├── models/
│   └── models.go
├── handlers/
│   ├── auth.go
│   ├── service.go
│   ├── booking.go
│   ├── notification.go
│   ├── review.go
│   └── user.go
├── main.go
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

## API Endpoints

### Authentication
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login user
- `POST /api/auth/logout` - Logout user
- `POST /api/auth/refresh` - Refresh authentication token

### Users
- `GET /api/users/:id` - Get user by ID
- `PUT /api/users/:id` - Update user profile
- `GET /api/users/:id/profile` - Get user profile with bookings and reviews

### Services
- `GET /api/services` - Get services with optional search/filter/pagination
- `GET /api/services/:id` - Get service by ID
- `POST /api/services` - Create new service (Admin)
- `PUT /api/services/:id` - Update service (Admin)
- `DELETE /api/services/:id` - Delete service (Admin)
- `GET /api/services/category/:category` - Get services by category

Query parameters for `GET /api/services`:
- `q` (string): case-insensitive partial match on service name and description
- `category` (string): case-insensitive category filter
- `activeOnly` (`true`/`false`): include only active services when `true`
- `limit` (int): page size, min `1`, max `200`, default `100`
- `offset` (int): number of rows to skip, default `0`
- `sortBy` (string): one of `name`, `category`, `rating`, `createdAt`, `updatedAt`
- `sortOrder` (string): `asc` (default) or `desc`

### Bookings
- `POST /api/bookings` - Create new booking
- `GET /api/bookings/:id` - Get booking by ID
- `GET /api/bookings/user/:userId` - Get user's bookings
- `PUT /api/bookings/:id` - Update booking
- `DELETE /api/bookings/:id` - Cancel booking

### Reviews
- `POST /api/reviews` - Create review
- `GET /api/reviews/service/:serviceId` - Get service reviews
- `GET /api/reviews/:id` - Get review by ID
- `DELETE /api/reviews/:id` - Delete review

### Notifications
- `GET /api/notifications/:userId` - Get user notifications
- `POST /api/notifications` - Create notification
- `PUT /api/notifications/:id/read` - Mark notification as read

## Database Models

### User
- id (UUID, Primary Key)
- email (String, Unique)
- password (String, Hashed)
- firstName (String)
- lastName (String)
- phone (String)
- role (String) - student, staff, admin
- department (String)
- avatarUrl (String)
- bio (String)
- createdAt (Timestamp)
- updatedAt (Timestamp)

### Service
- id (UUID, Primary Key)
- name (String)
- description (String)
- category (String)
- imageUrl (String)
- location (String)
- phone (String)
- email (String)
- hours (String)
- rating (Float)
- isActive (Boolean)
- createdAt (Timestamp)
- updatedAt (Timestamp)

### Booking
- id (UUID, Primary Key)
- userId (UUID, Foreign Key)
- serviceId (UUID, Foreign Key)
- status (String) - pending, confirmed, completed, cancelled
- startTime (Timestamp)
- endTime (Timestamp)
- notes (String)
- createdAt (Timestamp)
- updatedAt (Timestamp)

### Review
- id (UUID, Primary Key)
- userId (UUID, Foreign Key)
- serviceId (UUID, Foreign Key)
- rating (Integer) - 1-5
- comment (String)
- createdAt (Timestamp)
- updatedAt (Timestamp)

### Notification
- id (UUID, Primary Key)
- userId (UUID, Foreign Key)
- title (String)
- message (String)
- type (String) - booking, reminder, announcement
- isRead (Boolean)
- createdAt (Timestamp)

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_PATH | SQLite database file path | data/smart_campus.db |
| PORT | Server port | 8080 |
| GIN_MODE | Gin mode | debug |

## Dependencies

- `gin-gonic/gin` - HTTP web framework
- `gorm/gorm` - ORM for database operations
- `gorm/driver/sqlite` - SQLite driver
- `joho/godotenv` - Environment variable loader
- `google/uuid` - UUID generator

## Testing

To test the API endpoints, you can use:
- Postman
- curl
- Thunder Client
- Insomnia

Example:
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "student@uf.edu",
    "password": "password123",
    "firstName": "John",
    "lastName": "Doe",
    "phone": "+1234567890",
    "role": "student"
  }'
```

## Error Handling

The API returns appropriate HTTP status codes:
- 200: Success
- 201: Created
- 400: Bad Request
- 401: Unauthorized
- 404: Not Found
- 500: Internal Server Error

## Security Notes

- Implement JWT authentication (currently using basic token mechanism)
- Hash passwords using bcrypt before storing
- Add rate limiting
- Implement request validation
- Add CORS policies
- Use HTTPS in production
- Validate and sanitize all inputs

## Future Improvements

- Add user role-based access control
- Implement email notifications
- Add payment integration
- Create admin dashboard
- Add real-time notifications using WebSockets
- Implement request pagination
- Add API documentation with Swagger
