# Sprint 2 Report

## Team Members:
Venkata Sai Saran Jonnalagadda - 11114995 
Srikar Panuganti - 38909216 
Keerthi Reddy Gudibandi - 13652831 
Vishnu Sai Padyala - 32712860

## Issues planned for sprint-2
- Seed Data Expansion for All 8 Categories - Completed
- Review List & Delete for Authenticated User - Completed
- Service Search by Name & Description - Completed
- Notification Auto-Trigger on Booking Create - Completed
- Auth Middleware Integration on Protected Routes - Completed
- Booking Cancel vs Update Endpoint Separation - Completed
- Review Submission Form in Service Detail - Completed
- Profile Edit Form & Update Handler - Completed
- User Profile Page with Booking History - Completed
- Service Detail Page with Booking Form - Completed
- Register Page & Role Selection - Completed
- Login Page & Authentication Flow #16 - Completed

## Back End API endpoints

### Health and setup
#### `GET /health`
- Purpose: confirms the service is running.
- Response:
```json
{
  "status": "healthy",
  "message": "Smart Campus Services Platform is running"
}
```
#### `POST /api/seed`
- Purpose: seeds a small local development dataset if the services table is empty.
- Success response:
```json
{
  "message": "Database seeded successfully"
}
```

### Authentication
#### `POST /api/auth/register`
- Purpose: create a new user account.
- Request body:
```json
{
  "email": "student@campus.edu",
  "password": "secret123",
  "firstName": "Sam",
  "lastName": "Student",
  "phone": "555-1000",
  "role": "student"
}
```
- Validation rules:
  - `email` required and must be a valid email
  - `password` required and minimum 6 characters
  - `firstName`, `lastName`, `phone`, `role` required
- Success response: `201 Created`
```json
{
  "id": "uuid",
  "email": "student@campus.edu",
  "firstName": "Sam",
  "lastName": "Student",
  "phone": "555-1000",
  "role": "student",
  "serviceId": "",
  "token": "token-here"
}
```
- Error responses:
  - `409` if email already exists
  - `400` if payload validation fails

#### `POST /api/auth/login`
- Purpose: authenticate a user.
- Request body:
```json
{
  "email": "student@campus.edu",
  "password": "secret123"
}
```
- Success response: `200 OK`
```json
{
  "id": "uuid",
  "email": "student@campus.edu",
  "firstName": "Sam",
  "lastName": "Student",
  "phone": "555-1000",
  "role": "student",
  "serviceId": "",
  "token": "token-here"
}
```
- Error responses:
  - `401` if the user does not exist or password is wrong
  - `400` if payload validation fails

#### `POST /api/auth/logout`
- Purpose: placeholder logout endpoint.
- Success response:
```json
{
  "message": "Logged out successfully"
}
```
#### `POST /api/auth/refresh`
- Purpose: placeholder token refresh endpoint.
- Success response:

```json
{
  "token": "new-token-here"
}
```

### Users
#### `GET /api/users/:id`
- Purpose: fetch a single user record.
- Success response: full user JSON excluding password.
- Error responses:
  - `404` if user is missing
#### `PUT /api/users/:id`
- Purpose: update profile-facing user fields.
- Request body:
```json
{
  "firstName": "Updated",
  "lastName": "User",
  "phone": "555-9999",
  "department": "Engineering",
  "avatarUrl": "https://example.com/avatar.png",
  "bio": "Updated bio"
}
```
- Notes:
  - The handler updates the record in the database.
  - The current implementation returns `200 OK` with an empty user object rather than a refreshed user payload.
#### `GET /api/users/:id/profile`
- Purpose: fetch a user with preloaded `bookings` and `reviews`.
- Success response: user JSON with `bookings` and `reviews` arrays.
- Error responses:
  - `404` if user is missing

### Services
#### `GET /api/services`
- Purpose: list all services.
- Success response: array of service objects.
#### `GET /api/services/:id`
- Purpose: fetch one service with preloaded `reviews`.
- Error responses:
  - `404` if service is missing
#### `POST /api/services`
- Purpose: create a new service.
- Request body:
```json
{
  "name": "Health Center",
  "description": "Clinic support",
  "category": "health",
  "imageUrl": "",
  "location": "Building A",
  "phone": "555-2000",
  "email": "health@campus.edu",
  "hours": "Weekdays"
}
```

- Required fields:
  - `name`
  - `description`
  - `category`
  - `location`
- Success response: created service object
#### `PUT /api/services/:id`
- Purpose: update an existing service.
- Request body: same schema as create.
- Notes:
  - The database record is updated successfully.
  - The current implementation returns `200 OK` with an empty service object instead of the refreshed service.
#### `DELETE /api/services/:id`
- Purpose: delete a service.
- Success response:
```json
{
  "message": "Service deleted successfully"
}
```
#### `GET /api/services/category/:category`
- Purpose: list services for a single category.
- Success response: array of services matching the category string exactly.

### Bookings
#### `POST /api/bookings`
- Purpose: create a booking.
- Request body:
```json
{
  "userId": "user-uuid",
  "serviceId": "service-uuid",
  "startTime": "2026-03-21T14:00:00Z",
  "endTime": "2026-03-21T15:00:00Z",
  "notes": "Need transportation"
}
```
- Behavior:
  - New bookings are created with `status: "pending"`.
- Success response: created booking object.
#### `GET /api/bookings/:id`
- Purpose: fetch one booking with preloaded `user` and `service`.
- Error responses:
  - `404` if booking is missing
#### `GET /api/bookings/user/:userId`
- Purpose: list all bookings for one user.
- Success response: array of bookings with preloaded `service`.
#### `PUT /api/bookings/:id`
- Purpose: update booking fields.
- Request body: same schema as create.
- Notes:
  - The record is updated in the database.
  - The current implementation returns `200 OK` with an empty booking object instead of the refreshed row.
#### `DELETE /api/bookings/:id`
- Purpose: cancel a booking.
- Behavior:
  - Updates the booking `status` to `cancelled`.
- Success response:
```json
{
  "message": "Booking cancelled successfully"
}
```

### Approval workflow
#### `GET /api/approval/staff/:staffId/pending`
- Purpose: staff view of pending bookings for the service they manage.
- Rules:
  - `staffId` must belong to a user with role `staff`
  - the staff user must have `serviceId` assigned
- Error responses:
  - `404` if staff member is not found
  - `400` if staff member has no assigned service
#### `GET /api/approval/staff/:staffId/all`
- Purpose: staff view of all bookings for the service they manage.
#### `PUT /api/approval/bookings/:id/approve`
- Purpose: approve a booking as staff.
- Request body:
```json
{
  "status": "approved",
  "approvalNotes": "Approved for your requested slot",
  "staffId": "staff-uuid"
}
```
- Rules:
  - `staffId` can come from request JSON or `staffID` Gin context
  - the staff member must manage the same service as the booking
- Behavior:
  - booking status is updated to `approved`
  - `approvedBy` and `approvalNotes` are stored
  - a notification is created for the student
- Error responses:
  - `400` if no staff ID is provided
  - `403` if the staff member does not own that service
  - `404` if booking is missing
#### `PUT /api/approval/bookings/:id/reject`
- Purpose: reject a booking as staff.
- Request body:
```json
{
  "status": "rejected",
  "approvalNotes": "Requested slot unavailable",
  "staffId": "staff-uuid"
}
```
- Behavior mirrors approval, but booking status becomes `rejected`.
#### `GET /api/approval/admin/:userId/pending`
- Purpose: admin view of all pending bookings across services.
- Rule:
  - `userId` must belong to an `admin`
- Error responses:
  - `403` for non-admin users
#### `GET /api/approval/admin/:userId/all`
- Purpose: admin view of all bookings across services.
#### `PUT /api/approval/admin/:userId/bookings/:id/approve`
- Purpose: approve any booking as admin.
- Request body:
```json
{
  "status": "approved",
  "approvalNotes": "Approved by admin"
}
```
- Behavior:
  - sets status to `approved`
  - stores admin ID in `approvedBy`
  - creates an approval notification
#### `PUT /api/approval/admin/:userId/bookings/:id/reject`
- Purpose: reject any booking as admin.
- Request body:
```json
{
  "status": "rejected",
  "approvalNotes": "Rejected by admin"
}
```
- Behavior:
  - sets status to `rejected`
  - stores admin ID in `approvedBy`
  - creates a rejection notification
### Notifications
#### `GET /api/notifications/:userId`
- Purpose: list notifications for one user ordered by newest first.
#### `POST /api/notifications`
- Purpose: create a notification.
- Request body:
```json
{
  "userId": "user-uuid",
  "title": "Approval Update",
  "message": "Your booking was reviewed",
  "type": "booking"
}
```
- Behavior:
  - `isRead` is always initialized to `false`
#### `PUT /api/notifications/:id/read`
- Purpose: mark a notification as read.
- Behavior:
  - sets `isRead` to `true`
### Reviews
#### `POST /api/reviews`
- Purpose: create a service review.
- Request body:
```json
{
  "userId": "user-uuid",
  "serviceId": "service-uuid",
  "rating": 5,
  "comment": "Excellent support"
}
```
- Validation rules:
  - `userId`, `serviceId`, `comment` required
  - `rating` required and must be between 1 and 5
- Behavior:
  - creates the review
  - recalculates the parent service rating using average review score
#### `GET /api/reviews/service/:serviceId`
- Purpose: list reviews for one service, newest first.
#### `GET /api/reviews/user/:userId`
- Purpose: list reviews written by one user, newest first.
#### `GET /api/reviews/:id`
- Purpose: fetch one review with preloaded `user` and `service`.
- Error responses:
  - `404` if review is missing
#### `DELETE /api/reviews/:id`
- Purpose: delete a review.
- Rules:
  - if `userId` exists in Gin context, it must match the review owner
- Behavior:
  - deletes the review
  - recalculates the service rating
- Error responses:
  - `403` if a different authenticated user tries to delete the review
  - `404` if review is missing

## Unit test implementation
### Framework-specific test approach
- Gin handlers are tested through real HTTP requests using `httptest`
- GORM behavior is tested against in-memory SQLite databases
- Model hooks are tested directly for UUID generation
- The tests stay inside the Go backend and match the framework actually used by the project

### Test files added
- `backend/models/models_test.go`
- `backend/handlers/test_helpers_test.go`
- `backend/handlers/auth_test.go`
- `backend/handlers/service_test.go`
- `backend/handlers/booking_test.go`
- `backend/handlers/notification_test.go`
- `backend/handlers/user_test.go`
- `backend/handlers/review_test.go`
- `backend/handlers/approval_test.go`
- `backend/main_test.go`

### Coverage summary
- Model hooks:
  - user UUID creation
  - service UUID creation
  - booking UUID creation
  - notification UUID creation
  - review UUID creation
- Auth handlers:
  - register success
  - duplicate registration rejection
  - login success
  - invalid password rejection
  - logout response
  - refresh token response
- Service handlers:
  - list services
  - get missing service
  - create service
  - update service persistence
  - delete service
  - filter by category
- Booking handlers:
  - create booking with pending status
  - get booking with relations
  - list user bookings
  - update booking persistence
  - cancel booking
- Notification handlers:
  - list user notifications
  - create unread notification
  - mark notification as read
- User handlers:
  - get user
  - update user persistence
  - get profile with bookings and reviews
- Review handlers:
  - create review and update average rating
  - get service reviews
  - get single review
  - get user reviews
  - delete review and recalculate rating
- Approval handlers:
  - staff pending bookings view
  - staff all bookings view
  - staff approve booking
  - reject booking forbidden for wrong staff
  - admin-only pending view guard
  - admin all bookings view
  - admin approve booking
  - admin reject booking
- Main package helpers:
  - `databasePath()` env override
  - `databasePath()` default path
  - CORS middleware `OPTIONS` handling

### Running the tests
From the `backend` directory:

```bash
GOCACHE=$(pwd)/.gocache go test ./...
```

This local `GOCACHE` path avoids sandbox permission issues with the default Go build cache location.
