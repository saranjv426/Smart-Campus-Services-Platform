# Testing Guide

This guide explains how to run and maintain tests for the Smart Campus Services Platform. The current automated backend tests cover handlers, routing, middleware, validation, models, configuration, and shared test utilities.

## Prerequisites

Install the project requirements before running tests:

- Go 1.21 or newer
- Node.js 14 or newer
- npm
- Git

The backend tests use in-memory SQLite databases and do not require the development database file to exist.

## Backend Test Commands

Run all backend tests:

```bash
cd backend
go test ./...
```

Run all backend tests with an explicit temporary Go build cache:

```bash
cd backend
GOCACHE=/tmp/go-build-cache go test ./...
```

Run a specific backend package:

```bash
cd backend
go test ./handlers
go test ./router
go test ./validation
```

Run one backend test by name:

```bash
cd backend
go test ./handlers -run TestCreateBookingSetsPendingStatus
```

Run backend tests with verbose output:

```bash
cd backend
go test -v ./...
```

## Frontend Test Commands

Install frontend dependencies first:

```bash
cd frontend
npm install
```

Run React tests:

```bash
cd frontend
npm test
```

Run a production build check:

```bash
cd frontend
npm run build
```

## Current Backend Test Structure

```text
backend/
|-- main_test.go
|-- config/config_test.go
|-- handlers/
|   |-- approval_test.go
|   |-- auth_test.go
|   |-- booking_test.go
|   |-- notification_test.go
|   |-- review_test.go
|   |-- service_test.go
|   |-- test_helpers_test.go
|   `-- user_test.go
|-- middleware/
|   |-- middleware_test.go
|   `-- test_helpers_test.go
|-- models/models_test.go
|-- router/routes_test.go
|-- testutil/
|   |-- http_test.go
|   `-- test_db_test.go
`-- validation/validator_test.go
```

## Backend Testing Expectations

Handler tests should verify:

- Successful HTTP status codes, such as `200 OK` and `201 Created`.
- Error status codes, such as `400 Bad Request`, `403 Forbidden`, and `404 Not Found`.
- Response bodies where the payload matters to the caller.
- Database persistence after create, update, delete, and state-change operations.
- Missing records and invalid request payloads.
- Malformed JSON when request binding is involved.
- Authorization-sensitive behavior, such as review ownership or staff service ownership.

Use the existing helper functions in `backend/handlers/test_helpers_test.go` for handler tests:

- `setupTestDB(t)`
- `performRequest(t, router, method, path, body)`
- `performRawRequest(router, req)`
- `decodeJSON[T](t, rec)`
- Fixture helpers such as `createUserFixture`, `createServiceFixture`, `createBookingFixture`, `createReviewFixture`, and `createNotificationFixture`

## Common Backend Test Workflow

Before opening a pull request or pushing issue work:

```bash
cd backend
GOCACHE=/tmp/go-build-cache go test ./...
```

After changing only one handler, run the narrow package first:

```bash
cd backend
GOCACHE=/tmp/go-build-cache go test ./handlers
```

Then run the full backend suite:

```bash
cd backend
GOCACHE=/tmp/go-build-cache go test ./...
```

## Manual Local API Checks

Start the backend:

```bash
cd backend
go run main.go
```

Check health:

```bash
curl http://localhost:8080/health
```

Seed local development data:

```bash
curl -X POST http://localhost:8080/api/seed
```

Register a student:

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "student@ufl.edu",
    "password": "student123",
    "firstName": "John",
    "lastName": "Doe",
    "phone": "352-555-1234",
    "role": "student"
  }'
```

List services:

```bash
curl "http://localhost:8080/api/services?activeOnly=true&limit=10"
```

Create a booking:

```bash
curl -X POST http://localhost:8080/api/bookings \
  -H "Content-Type: application/json" \
  -d '{
    "userId": "replace-with-user-id",
    "serviceId": "replace-with-service-id",
    "startTime": "2026-05-01T10:00:00Z",
    "endTime": "2026-05-01T11:00:00Z",
    "notes": "Study room request"
  }'
```

Approve a booking as staff:

```bash
curl -X PUT "http://localhost:8080/api/approval/bookings/replace-with-booking-id/approve" \
  -H "Content-Type: application/json" \
  -d '{
    "staffId": "replace-with-staff-id",
    "approvalNotes": "Approved"
  }'
```

Reject a booking as staff:

```bash
curl -X PUT "http://localhost:8080/api/approval/bookings/replace-with-booking-id/reject" \
  -H "Content-Type: application/json" \
  -d '{
    "staffId": "replace-with-staff-id",
    "approvalNotes": "Time slot unavailable"
  }'
```

## Manual UI Smoke Test

Run the backend and frontend in separate terminals:

```bash
cd backend
go run main.go
```

```bash
cd frontend
npm start
```

Then verify these flows at `http://localhost:3000`:

- Register or log in as a user.
- Browse the services list.
- Open a service detail page.
- Create a booking.
- View the booking in the user's bookings page.
- Log in as staff or admin and approve or reject a booking.
- Confirm the booking status changes in the student view.

## Troubleshooting Tests

Go cannot write to the default build cache:

```bash
cd backend
GOCACHE=/tmp/go-build-cache go test ./...
```

Backend tests fail because of stale local data:

- Backend tests should use in-memory databases, so failures usually come from code changes rather than `data/smart_campus.db`.
- If manual API checks behave unexpectedly, stop the backend, remove `backend/data/smart_campus.db`, restart the backend, and seed again.

Frontend tests cannot find dependencies:

```bash
cd frontend
npm install
npm test
```

Frontend cannot call the backend during manual checks:

- Confirm the backend is running.
- Confirm `frontend/.env` has `REACT_APP_API_URL=http://localhost:8080/api`.
- Restart `npm start` after editing `.env`.

## Documentation Maintenance Checklist

When backend routes, request bodies, or test files change:

- Update this guide with any new test command or package.
- Update the root `README.md` if setup or run instructions change.
- Remove stale endpoints or request fields from examples.
- Keep examples aligned with the current handler request structs.
- Run `go test ./...` before marking backend testing documentation complete.
