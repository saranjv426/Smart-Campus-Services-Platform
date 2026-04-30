# Sprint 4 Report - Smart Campus Services Platform

A comprehensive Sprint 4 summary for the Smart Campus Services Platform covering completed work, frontend and backend automated tests, documentation updates, and the current backend API reflected in the project.

**Status**: ✅ **Sprint 4 Complete** | **Sprint Focus**: Frontend Test Expansion, Booking Workflow Validation, Dashboard Reliability, Backend Test Expansion, and Documentation Updates

---

## 🎯 Sprint 4 Overview

### Team Members
- Venkata Sai Saran Jonnalagadda - 11114995
- Srikar Panuganti - 38909216
- Keerthi Reddy Gudibandi - 13652831
- Vishnu Sai Padyala - 32712860

### Sprint Goal
- Strengthen the reliability of the Smart Campus Services Platform by expanding frontend unit test coverage, validating end-to-end navigation, improving backend handler test coverage, and updating project/testing documentation for easier setup, execution, and evaluation.

### Sprint 4 Completion Summary

| Area | Outcome | Status |
|------|---------|--------|
| Profile page test coverage | Completed | ✅ |
| Staff dashboard test coverage | Completed | ✅ |
| Admin dashboard test coverage | Completed | ✅ |
| Bookings page cancel and filter tests | Completed | ✅ |
| Login, Register, Services, and Service Detail frontend test coverage | Completed | ✅ |
| Cypress basic navigation coverage | Completed | ✅ |
| Review handler backend tests | Completed | ✅ |
| Expanded service handler tests for update and delete | Completed | ✅ |
| README and testing guide updates | Completed | ✅ |
| Updated backend API documentation | Completed | ✅ |

---

## 📋 Work Completed In Sprint 4

## 1. Profile Page Tests

### Description
Expanded and refined frontend unit tests for the Profile page to verify profile rendering, user details, booking history, review display, profile editing, validation, and delete-review behavior.

### Completed Work
- Added and refined tests for profile loading and route-based profile fetching.
- Verified rendering of user details such as name, role, email, and other profile-facing information.
- Added coverage for recent bookings summary rendering.
- Added coverage for user reviews and empty review state.
- Tested edit-profile flow, including entering edit mode and updating form fields.
- Tested save behavior for profile updates.
- Added validation coverage for required fields.
- Covered delete-review flow and related success behavior.
- Covered error handling for profile fetch failures.

### Key Frontend Files
- `frontend/src/pages/Profile.test.js`
- `frontend/src/pages/Profile.js`

---

## 2. StaffDashboard & AdminDashboard Page Tests

### Description
Expanded frontend unit tests for both dashboard pages to validate booking-management workflows for staff and administrators.

### Completed Work
- Added tests for Staff Dashboard rendering and pending-booking display.
- Verified approve and reject button rendering for pending staff bookings.
- Tested modal opening, note entry, submission flow, and modal closing behavior for staff actions.
- Added error-handling coverage for staff approval/rejection workflow.
- Added tests for Admin Dashboard rendering and admin-only access behavior.
- Verified fetching and rendering of bookings and services on the admin side.
- Added coverage for booking statistics cards and dashboard summary UI.
- Tested admin approve and reject flows.
- Tested tab navigation between bookings overview and manage services sections.
- Tested booking filtering by status and service selection.

### Key Frontend Files
- `frontend/src/pages/StaffDashboard.test.js`
- `frontend/src/pages/AdminDashboard.test.js`
- `frontend/src/pages/StaffDashboard.js`
- `frontend/src/pages/AdminDashboard.js`

---

## 3. Bookings Page Tests – Cancel & Filter

### Description
Expanded frontend unit tests for the My Bookings page with a focus on status filtering, refresh behavior, cancellation flow, and edge cases.

### Completed Work
- Added tests for authenticated booking fetch on page load.
- Verified rendering of booking cards with service details, locations, notes, and status badges.
- Added status-filter tests for pending and all bookings views.
- Verified filter reset behavior when returning to the All view.
- Added cancel-button visibility tests for eligible pending bookings.
- Tested cancel-booking confirmation flow and booking-list refresh after cancellation.
- Tested refresh button behavior.
- Added empty-state coverage when no bookings are returned.
- Added unauthenticated-user error-state coverage.

### Key Frontend Files
- `frontend/src/pages/Bookings.test.js`
- `frontend/src/pages/Bookings.js`

---

## 4. Comprehensive Frontend Test Coverage for Authentication and Service Discovery Pages

### Description
Expanded frontend test coverage for the application’s core user-entry and service-discovery flows.

### Completed Work
- Added/refined unit tests for Login page form rendering, submission flow, token storage, and failure cases.
- Added/refined unit tests for Register page input handling, role selection, form submission, and error handling.
- Added/refined Services page tests for service rendering, search, and category filtering.
- Added/refined Service Detail page tests for service loading, review rendering, booking-form behavior, and validation.
- Verified core frontend entry flows needed by new users before they reach the dashboard and booking workflows.

### Key Frontend Files
- `frontend/src/pages/Login.test.js`
- `frontend/src/pages/Register.test.js`
- `frontend/src/pages/Services.test.js`
- `frontend/src/pages/ServiceDetail.test.js`

---

## 5. Cypress End-to-End Navigation Test

### Description
Maintained end-to-end browser validation for a basic user navigation flow using Cypress.

### Completed Work
- Verified application load at the home page.
- Clicked the **Get Started** button from the landing page.
- Confirmed navigation to the Register page.
- Confirmed that Register page content rendered successfully after navigation.

### Key Cypress File
- `frontend/cypress/e2e/spec.cy.js`

---

## 6. Documentation Updates - README and TESTING_GUIDE

### Description
Updated project documentation to make setup, execution, and testing easier for developers, teammates, and graders.

### Completed Work
- Updated README content to better explain project purpose and execution steps.
- Documented frontend and backend setup flow.
- Added or refined test-running instructions for frontend unit tests.
- Added or refined backend test-running instructions.
- Added or refined Cypress execution guidance.
- Improved testing-related documentation so project evaluation is easier and more reproducible.

### Key Documentation Files
- `README.md`
- `TESTING_GUIDE.md`

---

## 7. Backend Test Expansion - Review Handler Tests

### Description
Expanded backend review handler tests to improve regression protection for review-related workflows.

### Completed Work
- Verified review creation flow.
- Verified service-rating recalculation behavior after review creation.
- Verified service review listing behavior.
- Verified single review retrieval behavior.
- Verified user review retrieval behavior.
- Verified review deletion behavior and post-delete rating recalculation.
- Strengthened backend coverage around review ownership and correctness of review-side effects.

### Key Backend Files
- `backend/handlers/review_test.go`
- `backend/handlers/review.go`

---

## 8. Backend Test Expansion - Service Handler Update & Delete Coverage

### Description
Expanded service handler tests to strengthen backend coverage for service update and delete workflows.

### Completed Work
- Added or refined tests covering service update behavior.
- Added or refined tests covering service delete behavior.
- Verified persistence of updated service data.
- Verified successful removal of deleted services.
- Improved reliability of service-management backend flows that support the admin dashboard.

### Key Backend Files
- `backend/handlers/service_test.go`
- `backend/handlers/service.go`

---

## 🧪 Frontend Unit Tests

The following frontend test files are currently present and cover major UI flows, dashboards, profile functionality, booking workflows, authentication flows, and service discovery interactions.

### Frontend Test Files
- `frontend/src/pages/AdminDashboard.test.js`
- `frontend/src/pages/Bookings.test.js`
- `frontend/src/pages/Home.test.js`
- `frontend/src/pages/Login.test.js`
- `frontend/src/pages/Profile.test.js`
- `frontend/src/pages/ProfileEdit.test.js`
- `frontend/src/pages/Register.test.js`
- `frontend/src/pages/ServiceDetail.test.js`
- `frontend/src/pages/Services.test.js`
- `frontend/src/pages/StaffDashboard.test.js`

### Sprint 4 Frontend Test Areas Especially Relevant
- `Profile.test.js`
  - profile loading
  - user information rendering
  - bookings summary
  - reviews
  - edit profile flow
  - validation and delete review behavior
- `StaffDashboard.test.js`
  - pending bookings display
  - approve/reject button rendering
  - modal interaction
  - notes entry
  - approval/rejection submission
- `AdminDashboard.test.js`
  - admin-only access behavior
  - booking and service fetch behavior
  - statistics cards
  - tab navigation
  - filtering and approval workflows
- `Bookings.test.js`
  - authenticated booking loading
  - booking detail rendering
  - filter behavior
  - cancel-booking flow
  - refresh behavior
  - empty state and unauthenticated state
- `Login.test.js`
  - authentication flow and error handling
- `Register.test.js`
  - registration form submission and validation
- `Services.test.js`
  - services rendering, search, and category filtering
- `ServiceDetail.test.js`
  - service loading, booking-form behavior, and review rendering

### Frontend Unit Test Result
From the `frontend` directory:

```bash
npm test -- --watchAll=false
```

---

## 🧪 Cypress Tests

### Cypress Test File
- `frontend/cypress/e2e/spec.cy.js`

### Cypress Coverage
- Visits the application home page.
- Clicks the **Get Started** button.
- Verifies navigation to the Register page.
- Confirms Register page content renders successfully.

### Running Cypress Tests
From the `frontend` directory:

```bash
npm run cypress:run
```

Or run the specific spec:

```bash
npm run cypress:run -- --spec "cypress/e2e/spec.cy.js"
```

---

## 🧪 Backend Unit Tests

The backend includes handler, middleware, router, main-package, and helper tests used to validate business logic and endpoint behavior.

### Backend Test Files
- `backend/handlers/approval_test.go`
- `backend/handlers/auth_test.go`
- `backend/handlers/booking_test.go`
- `backend/handlers/notification_test.go`
- `backend/handlers/review_test.go`
- `backend/handlers/service_test.go`
- `backend/handlers/test_helpers_test.go`
- `backend/handlers/user_test.go`
- `backend/main_test.go`

### Sprint 4 Backend Test Areas Especially Relevant
- `review_test.go`
  - create review
  - get service reviews
  - get single review
  - get user reviews
  - delete review and rating recalculation
- `service_test.go`
  - update service persistence
  - delete service behavior
  - additional service reliability checks
- `approval_test.go`
  - continued coverage for booking approval/rejection flows used by dashboards
- `booking_test.go`
  - booking retrieval and cancellation behavior supporting frontend flows
- `auth_test.go`
  - authentication behavior used by login and protected routes
- `user_test.go`
  - profile-related backend behavior supporting frontend profile tests

### Backend Test Result
From the `backend` directory:

```bash
go test ./...
```

---

## 🔗 Updated Backend API Documentation

Sprint 4 documentation reflects the current backend API used by the application and validated through frontend and backend testing. The documented backend areas include authentication, users, services, bookings, approval, notifications, and reviews.

## Authentication

### `POST /api/auth/register`
- Register a new user.

### `POST /api/auth/login`
- Authenticate an existing user.

### `POST /api/auth/logout`
- Logout endpoint.

### `POST /api/auth/refresh`
- Token refresh endpoint.

---

## Users

### `GET /api/users/:id`
- Fetch a single user by ID.

### `GET /api/users/:id/profile`
- Fetch a user profile with preloaded bookings and reviews.

### `PUT /api/users/:id`
- Update user profile-facing fields.

### `GET /api/users`
- Protected admin-only endpoint returning all users.

---

## Services

### `GET /api/services`
- Returns the services list.

### `GET /api/services/:id`
- Returns a single service with related review data.

### `GET /api/services/category/:category`
- Returns services for a specific category.

### `POST /api/services`
- Protected endpoint to create a new service.

### `PUT /api/services/:id`
- Protected endpoint to update service details.

### `PATCH /api/services/:id/active`
- Protected endpoint to update only the active state of a service.

### `DELETE /api/services/:id`
- Protected endpoint to delete a service.

---

## Bookings

### `POST /api/bookings`
- Protected endpoint to create a booking.

### `GET /api/bookings/:id`
- Returns booking details with related user and service information.

### `GET /api/bookings/user/:userId`
- Returns bookings for a user.
- Supports optional status filtering.

### `PUT /api/bookings/:id`
- Protected endpoint to update booking details.

### `PATCH /api/bookings/:id/status`
- Protected endpoint used for booking cancellation / status update behavior.

---

## Approval

### Staff Approval Routes

#### `GET /api/approval/staff/:staffId/pending`
- Staff-only route returning pending bookings for the staff member’s service.

#### `GET /api/approval/staff/:staffId/all`
- Staff-only route returning all bookings for the staff member’s service.

#### `PUT /api/approval/bookings/:id/approve`
- Staff-only route to approve a booking.

#### `PUT /api/approval/bookings/:id/reject`
- Staff-only route to reject a booking.

### Admin Approval Routes

#### `GET /api/approval/admin/:userId/pending`
- Admin-only route returning all pending bookings across the platform.

#### `GET /api/approval/admin/:userId/all`
- Admin-only route returning all bookings across the platform.

#### `PUT /api/approval/admin/:userId/bookings/:id/approve`
- Admin-only route to approve any booking.

#### `PUT /api/approval/admin/:userId/bookings/:id/reject`
- Admin-only route to reject any booking.

---

## Notifications

### `GET /api/notifications/:userId`
- Protected endpoint returning notifications for a user.

### `POST /api/notifications`
- Protected endpoint creating a notification.

### `PUT /api/notifications/:id/read`
- Protected endpoint marking a notification as read.

---

## Reviews

### `GET /api/reviews/service/:serviceId`
- Returns reviews for a service.

### `GET /api/reviews/user/:userId`
- Returns reviews created by a specific user.

### `GET /api/reviews/:id`
- Returns a single review.

### `POST /api/reviews`
- Protected endpoint creating a review.

### `DELETE /api/reviews/:id`
- Protected endpoint deleting a review.

---

## 📌 Sprint 4 Deliverables Summary

Sprint 4 successfully delivered:
- expanded frontend unit testing for profile, dashboard, bookings, authentication, and service pages
- Cypress validation for a core user navigation flow
- stronger backend review and service handler test coverage
- improved testing and setup documentation through README and TESTING_GUIDE updates
- updated backend API documentation aligned with the current tested application behavior

Sprint 4 improved the reliability, maintainability, and evaluability of the Smart Campus Services Platform by strengthening both automated test coverage and supporting project documentation.
