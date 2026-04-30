# Sprint 4

## Backend Work Completed

- Resolved unresolved merge conflict markers in backend authentication and booking tests.
- Rebuilt authentication handler coverage for registration, duplicate email handling, invalid role validation, login, logout, and token refresh.
- Rebuilt booking handler coverage for booking creation, required end time validation, notification creation, lookup, user filtering, update, cancellation, invalid cancellation status, and missing booking handling.
- Added backend validation so booking `endTime` must be after `startTime` for create and update requests.
- Added notification API behavior so marking a missing notification as read returns `404 Not Found`.
- Added service list tests for search, category filtering, active-only filtering, pagination, and sorting.
- Added review validation coverage for ratings outside the allowed 1-5 range.
- Confirmed booking cancellation uses `PATCH /api/bookings/:id/status` with `{"status":"cancelled"}`.
- Updated the root README with backend/frontend setup and test commands.
- Updated backend API documentation with service query parameters, booking status request body, and backend test coverage.

## Backend Unit Tests

Run from `backend/`:

```bash
go test ./...
```

Current backend test files:

- `backend/main_test.go`
- `backend/models/models_test.go`
- `backend/router/routes_test.go`
- `backend/handlers/auth_test.go`
- `backend/handlers/booking_test.go`
- `backend/handlers/service_test.go`
- `backend/handlers/review_test.go`
- `backend/handlers/notification_test.go`
- `backend/handlers/approval_test.go`
- `backend/handlers/user_test.go`

Latest result:

```text
ok smart-campus-services
ok smart-campus-services/handlers
ok smart-campus-services/models
ok smart-campus-services/router
```

## Frontend Unit And Cypress Tests

Frontend unit tests:

- `frontend/src/App.test.js`
- `frontend/src/services/api.test.js`
- `frontend/src/components/Navbar.test.js`
- `frontend/src/components/Footer.test.js`
- `frontend/src/pages/Login.test.js`
- `frontend/src/pages/Register.test.js`
- `frontend/src/pages/Home.test.js`
- `frontend/src/pages/Services.test.js`
- `frontend/src/pages/ServiceDetail.test.js`
- `frontend/src/pages/Bookings.test.js`
- `frontend/src/pages/Profile.test.js`
- `frontend/src/pages/ProfileEdit.test.js`
- `frontend/src/pages/AdminDashboard.test.js`
- `frontend/src/pages/StaffDashboard.test.js`

Cypress test:

- `frontend/cypress/e2e/spec.cy.js`

## Updated Backend API Documentation

The backend API documentation was updated in `backend/README.md`.

Important Sprint 4 notes:

- `GET /api/services` supports `q`, `category`, `activeOnly`, `limit`, `offset`, `sortBy`, and `sortOrder`.
- `PATCH /api/bookings/:id/status` accepts only `{"status":"cancelled"}` for user cancellation.
- Booking create/update requests reject time ranges where `endTime` is not after `startTime`.
- Creating a booking automatically creates an unread booking notification for the user.
- `PUT /api/notifications/:id/read` returns `404` when the notification ID does not exist.
- Protected write routes require the `Authorization: Bearer <token>` header.

## Video Presentation Notes

- Demonstrate registration and login.
- Demonstrate browsing/searching services.
- Demonstrate creating a booking and show that a notification is generated.
- Demonstrate invalid booking time validation with a `400 Bad Request` response.
- Demonstrate booking cancellation through the backend API.
- Show `go test ./...` passing from the backend directory.
- Explain the backend API groups: auth, users, services, bookings, approvals, reviews, and notifications.
