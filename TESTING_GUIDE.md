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
# Quick Testing Guide - Service-Specific Approval System

## 🎯 Complete End-to-End Testing Flow

### Prerequisites
- Backend running on `http://localhost:8080`
- Frontend running on `http://localhost:3000`
- SQLite database with seeded data

---

## 📋 Test Scenario: Library Booking Approval

### Step 1: Student Registers & Logs In

**Register new student:**
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

**Login as student:**
- Go to: `http://localhost:3000/login`
- Email: `student@ufl.edu`
- Password: `student123`
- Click Login

**Result**: Redirected to home page with "My Bookings" link visible

---

### Step 2: Student Books a Service

**From home page:**
1. Click on "Library Services" feature card
2. You're taken to `/services?category=library`
3. Click on "Main Library" service card
4. You're taken to `/services/{serviceId}` detail page

**Book the service:**
1. Fill in booking details:
   - Start Date/Time: Tomorrow at 10:00 AM
   - End Date/Time: Tomorrow at 12:00 PM
   - Reason: "Study session for exams"
2. Click "Book Service"

**Expected Result:**
- Booking created with status: "pending"
- Notification: "Booking created successfully"
- Booking ID returned in response

---

### Step 3: Staff Logs In & Views Pending Approvals

**Login as library staff:**
- Go to: `http://localhost:3000/login`
- Email: `library-admin@ufl.edu`
- Password: `library123`
- Click Login

**Expected Result:**
- Redirected to `/dashboard/staff`
- Page shows "Staff Dashboard"
- Welcome message: "Welcome, Library Admin"
- "Staff Dashboard" link in navbar (instead of "My Bookings")

---

### Step 4: Staff Views Pending Booking

**On Staff Dashboard:**
1. See "Pending Approvals" section with badge showing count
2. Booking card displayed with:
   - Service name: "Main Library"
   - Student name: "John Doe"
   - Email: "student@ufl.edu"
   - Start/End times: Tomorrow 10:00-12:00
   - Reason: "Study session for exams"
   - Status badge: "Pending"

---

### Step 5: Staff Approves Booking

**Click "Approve" button:**
1. Modal dialog opens
2. Title: "Approve Booking"
3. Shows booking info: "Main Library - John Doe"
4. Optional "Additional Notes" textarea
5. Two buttons: "Cancel" and "Approve Booking"

**Add optional note:**
- Type: "Approved for exam prep session"
- Click "Approve Booking"

**Expected Result:**
- Modal closes
- Booking card disappears from pending list
- Notification: "Booking approved successfully"
- Dashboard refreshes to show updated count

---

### Step 6: Verify Student Can See Approval

**Student logs back in:**
1. Go to: `http://localhost:3000/login`
2. Login with student credentials
3. Go to: "My Bookings" link in navbar

**Expected Result:**
- Booking shows with status: "Approved"
- ApprovedBy shows staff name (optional, if implemented)
- Student can see approval notes (if shown)

---

## 🧪 Alternative Test: Staff Rejects Booking

**Instead of Approve (Step 5), click "Reject":**

1. Modal dialog opens
2. Title: "Reject Booking"
3. Add rejection notes: "Time slot not available"
4. Click "Reject Booking"

**Expected Result:**
- Booking status changed to "rejected"
- Rejection notes saved
- Student sees rejection when viewing bookings

---

## 🔍 Detailed API Testing

### Test 1: Get Pending Bookings

```bash
# Get staff ID from login
STAFF_ID="3fbe01b6-079e-4b2b-aa30-bc24026d799a"

# Get pending bookings
curl -X GET "http://localhost:8080/api/approval/staff/$STAFF_ID/pending" \
  -H "Authorization: Bearer token-here"

# Expected Response: Array of pending bookings for staff's service
[
  {
    "id": "booking-id",
    "userId": "student-id",
    "serviceId": "library-service-id",
    "status": "pending",
    "startTime": "2026-02-04T10:00:00Z",
    "endTime": "2026-02-04T12:00:00Z",
    "notes": "Study session for exams",
    "approvedBy": "",
    "approvalNotes": "",
    "user": {
      "id": "student-id",
      "firstName": "John",
      "lastName": "Doe",
      "email": "student@ufl.edu",
      "phone": "352-555-1234"
    },
    "service": {
      "id": "library-service-id",
      "name": "Main Library",
      "category": "library"
    }
  }
]
```

### Test 2: Approve Booking

```bash
BOOKING_ID="9d1f5c22-e4f0-4b5b-996f-e212c20c9631"
STAFF_ID="3fbe01b6-079e-4b2b-aa30-bc24026d799a"

curl -X PUT "http://localhost:8080/api/approval/bookings/$BOOKING_ID/approve" \
  -H "Authorization: Bearer token-here" \
  -H "Content-Type: application/json" \
  -d '{
    "notes": "Approved for exam prep session"
  }'

# Expected Response: Updated booking with status: "approved"
{
  "id": "booking-id",
  "status": "approved",
  "approvedBy": "staff-id",
  "approvalNotes": "Approved for exam prep session",
  ...
}
```

### Test 3: Reject Booking

```bash
curl -X PUT "http://localhost:8080/api/approval/bookings/$BOOKING_ID/reject" \
  -H "Authorization: Bearer token-here" \
  -H "Content-Type: application/json" \
  -d '{
    "notes": "Time slot not available"
  }'

# Expected Response: Updated booking with status: "rejected"
```

### Test 4: Verify Service Isolation

**Create booking for wrong service:**

```bash
# Try to have library staff approve a dining booking
LIBRARY_STAFF_ID="3fbe01b6-079e-4b2b-aa30-bc24026d799a"
DINING_BOOKING_ID="some-dining-booking-id"

curl -X PUT "http://localhost:8080/api/approval/bookings/$DINING_BOOKING_ID/approve" \
  -H "Authorization: Bearer library-token" \
  -H "Content-Type: application/json" \
  -d '{"notes": "Approved"}'

# Expected Response: 403 Forbidden or error
# "This booking does not belong to your service"
```

---

## 📊 Expected Database State After Tests

### Users Table
```
| Email                  | Role  | ServiceID                            |
|-----------------------|-------|--------------------------------------|
| library-admin@ufl.edu  | staff | bc06e055-2076-4654-8020-cd443d68aef4|
| student@ufl.edu        |student| NULL                                 |
```

### Bookings Table
```
| ID                | ServiceID                            | Status   | ApprovedBy                           | Notes                      |
|------------------|--------------------------------------|----------|--------------------------------------|----------------------------|
| 9d1f5c22-...     | bc06e055-2076-...                   | approved | 3fbe01b6-079e-4b2b-aa30-bc24026d799a| Approved for exam prep     |
```

---

## ✅ Verification Checklist

Use this checklist to verify all functionality works:

### Staff Dashboard
- [ ] Staff sees "Staff Dashboard" link in navbar
- [ ] Dashboard loads without errors
- [ ] Pending bookings count displays correctly
- [ ] Each booking card shows all required information
- [ ] Approve button is clickable
- [ ] Reject button is clickable

### Approval Modal
- [ ] Modal opens when clicking Approve/Reject
- [ ] Modal title changes based on action
- [ ] Notes textarea works
- [ ] Cancel button closes modal
- [ ] Action button submits request

### Authorization
- [ ] Library staff can only see library bookings
- [ ] Staff cannot approve other service bookings
- [ ] Student cannot access staff dashboard
- [ ] Admin (if implemented) has proper access

### Status Updates
- [ ] Booking status changes to "approved" when approved
- [ ] Booking status changes to "rejected" when rejected
- [ ] Dashboard refreshes after action
- [ ] Student sees updated status in "My Bookings"

### API Responses
- [ ] Login returns role and serviceId
- [ ] Pending bookings endpoint returns correct data
- [ ] Approve endpoint updates booking correctly
- [ ] Reject endpoint updates booking correctly

---

## 🐛 Troubleshooting

### Issue: "Staff Dashboard" link not showing
- **Cause**: User not logged in as staff or role not set correctly
- **Solution**: Login with staff email/password and verify response includes `"role": "staff"`

### Issue: No pending bookings showing
- **Cause**: Booking's serviceId doesn't match staff's serviceId
- **Solution**: Verify booking created with correct service ID matching staff's service

### Issue: Cannot approve booking
- **Cause**: Authorization error (staff can't approve other service's bookings)
- **Solution**: Create booking for correct service or use correct staff account

### Issue: Modal not opening
- **Cause**: JavaScript error or component not loaded
- **Solution**: Check browser console for errors, refresh page

### Issue: Changes not persisting
- **Cause**: Backend not saving data
- **Solution**: Check backend logs, verify database connection

---

## 🎓 Learning Notes

This system demonstrates:
1. **Role-Based Access Control** - Different features for different user roles
2. **Service Isolation** - Staff can only see their service's data
3. **Workflow State Management** - Booking moves through status states
4. **RESTful API Design** - Clean endpoints, proper HTTP methods
5. **Modal UI Patterns** - Non-blocking user interactions
6. **Authorization Validation** - Backend checks permissions

---

**Happy Testing! 🚀**
