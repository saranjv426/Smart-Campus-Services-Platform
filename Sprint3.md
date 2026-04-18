# Sprint 3 Report - Smart Campus Services Platform

A comprehensive Sprint 3 summary for the Smart Campus Services Platform covering completed work, frontend and backend unit tests, and the updated backend API delivered during this sprint.

**Status**: ✅ **Sprint 3 Complete** | **Last Updated**: April 13, 2026 | **Sprint Focus**: Admin Dashboard, Booking Management, Service Management, and Backend Authorization Improvements

---

## 🎯 Sprint 3 Overview

### Team Members
- Venkata Sai Saran Jonnalagadda - 11114995
- Srikar Panuganti - 38909216
- Keerthi Reddy Gudibandi - 13652831
- Vishnu Sai Padyala - 32712860

### Sprint Goal
- Build a complete administrator workflow for monitoring bookings, approving and rejecting requests, managing services, and strengthening backend protection and test coverage for the new admin flows.

### Sprint 3 Completion Summary

| Area | Outcome | Status |
|------|---------|--------|
| Admin dashboard bookings overview | Completed | ✅ |
| Admin approve/reject workflow | Completed | ✅ |
| Admin create/delete service UI | Completed | ✅ |
| Backend role protection on approval routes | Completed | ✅ |
| Backend admin get-all-users endpoint | Completed | ✅ |
| Booking approval/rejection notification verification | Completed | ✅ |
| Expanded frontend and backend automated tests | Completed | ✅ |
| Updated backend API documentation | Completed | ✅ |

---

## 📋 Work Completed In Sprint 3

## 1. Admin Dashboard - View All Bookings

### Description
Implemented a comprehensive admin dashboard that displays all bookings across the platform in one place. The dashboard supports monitoring, filtering, and acting on bookings in real time.

### Completed Work
- Added a bookings overview tab for administrators.
- Displayed all bookings with booking metadata and status indicators.
- Added summary statistics cards for:
  - total bookings
  - pending bookings
  - approved bookings
  - rejected bookings
- Added frontend filters for:
  - booking status
  - service category
- Integrated dashboard with backend approval endpoints for fetching all bookings.
- Supported responsive layouts for desktop, tablet, and mobile views.

### Key Frontend Files
- `frontend/src/pages/AdminDashboard.js`
- `frontend/src/styles/AdminDashboard.css`

### Backend Integration
- `GET /api/approval/admin/:userId/all`
- `GET /api/approval/admin/:userId/pending`

---

## 2. Admin Approve Booking

### Description
Implemented admin approval functionality for pending bookings so administrators can confirm valid booking requests directly from the dashboard.

### Completed Work
- Added approve action buttons for pending bookings.
- Added approval modal with booking details.
- Allowed optional approval notes for auditability.
- Updated frontend state after approval without requiring a full page reload.
- Persisted approval status in the backend.
- Generated student notifications when bookings are approved.

### Key Frontend Files
- `frontend/src/pages/AdminDashboard.js`
- `frontend/src/styles/AdminDashboard.css`

### Key Backend Files
- `backend/handlers/approval.go`
- `backend/router/routes.go`
- `backend/middleware/auth.go`

### Backend Integration
- `PUT /api/approval/admin/:userId/bookings/:id/approve`

---

## 3. Admin Reject Booking

### Description
Implemented admin rejection functionality for pending bookings so administrators can reject conflicting or invalid requests with recorded reasoning.

### Completed Work
- Added reject action buttons for pending bookings.
- Added rejection modal requiring or encouraging a reason.
- Updated booking status to `Rejected`.
- Removed rejected items from the pending view and preserved them in booking history.
- Created student-facing rejection notifications.
- Stored rejection reason in approval notes for auditability.

### Key Frontend Files
- `frontend/src/pages/AdminDashboard.js`
- `frontend/src/styles/AdminDashboard.css`

### Key Backend Files
- `backend/handlers/approval.go`
- `backend/router/routes.go`

### Backend Integration
- `PUT /api/approval/admin/:userId/bookings/:id/reject`

---

## 4. Create New Service

### Description
Built an administrator workflow for creating new services directly from the admin dashboard through a dedicated modal form.

### Completed Work
- Added a service management tab in the admin dashboard.
- Added a create-service form modal with fields for:
  - service name
  - category
  - description
  - location
  - phone
  - email
  - hours
  - image URL
- Added validation for required fields.
- Set newly created services to active by default.
- Displayed newly created services immediately in the service list.
- Added success and error handling for submission.

### Key Frontend Files
- `frontend/src/pages/AdminDashboard.js`
- `frontend/src/styles/AdminDashboard.css`

### Backend Integration
- `POST /api/services`
- `GET /api/services`

---

## 5. Delete Service

### Description
Implemented an administrator workflow for deleting outdated or unavailable services from the system using a confirmation modal.

### Completed Work
- Added delete buttons on service cards.
- Added a confirmation modal before deletion.
- Displayed the selected service name inside the confirmation dialog.
- Removed deleted services from the UI immediately after success.
- Added success and failure messaging.
- Preserved a clear admin workflow for canceling the delete action.

### Key Frontend Files
- `frontend/src/pages/AdminDashboard.js`
- `frontend/src/styles/AdminDashboard.css`

### Backend Integration
- `DELETE /api/services/:id`

---

## 6. Backend Improvement - Role-Based Protection On Approval Routes

### Description
Strengthened backend authorization by applying `RequireRoles(...)` middleware to approval routes. Before this update, approval endpoints were authenticated but not consistently restricted by user role.

### Completed Work
- Restricted staff approval endpoints to `staff` users only.
- Restricted admin approval endpoints to `admin` users only.
- Updated lightweight token parsing so auth middleware places role information into request context.
- Added router tests confirming:
  - missing token returns `401`
  - wrong role returns `403`
  - correct role is allowed

### Key Backend Files
- `backend/middleware/auth.go`
- `backend/router/routes.go`
- `backend/main.go`
- `backend/router/routes_test.go`

### Result
- Staff users can only access staff approval routes.
- Admin users can only access admin approval routes.

---

## 7. Backend Improvement - Admin Get All Users Endpoint

### Description
Added a protected backend endpoint for retrieving all users, intended for administrator management features.

### Completed Work
- Added `GET /api/users` to return all users.
- Protected the endpoint with authentication and `RequireRoles("admin")`.
- Added tests for:
  - admin success
  - non-admin forbidden
  - auth-required behavior

### Key Backend Files
- `backend/handlers/user.go`
- `backend/router/routes.go`
- `backend/main.go`
- `backend/handlers/user_test.go`
- `backend/router/routes_test.go`

### Result
- Only admins can fetch the full user list.

---

## 8. Backend Improvement - Notification On Booking Approval And Rejection

### Description
Verified and strengthened notification behavior for approval and rejection workflows so students consistently receive notifications for booking status changes.

### Completed Work
- Verified notification creation for staff approval flow.
- Added explicit backend coverage for:
  - staff approve notification
  - staff reject notification
  - admin approve notification
  - admin reject notification
- Consolidated notification creation logic in approval handlers for consistency.

### Key Backend Files
- `backend/handlers/approval.go`
- `backend/handlers/approval_test.go`

### Result
- Booking approval and rejection flows now have stronger regression protection.

---

## 9. Backend Improvement - Service Activation Toggle

### Description
Added support for toggling a service’s active state with a focused endpoint rather than requiring a full service update payload.

### Completed Work
- Added `PATCH /api/services/:id/active`.
- Supported updating only the `isActive` field.
- Preserved all other service fields.
- Added tests covering:
  - activating a service
  - deactivating a service
  - invalid payload
  - missing service

### Key Backend Files
- `backend/handlers/service.go`
- `backend/handlers/service_test.go`
- `backend/router/routes.go`

---

## 10. Backend Improvement - Booking Status Filter

### Description
Moved booking status filtering into the backend so clients can request only the booking statuses they need.

### Completed Work
- Added optional `status` query parameter support to `GET /api/bookings/user/:userId`.
- Allowed valid statuses:
  - pending
  - approved
  - rejected
  - completed
  - cancelled
- Returned `400 Bad Request` for invalid status values.
- Preserved service preloading on returned bookings.
- Added tests for:
  - no filter
  - valid filter
  - invalid filter
  - no matching records

### Key Backend Files
- `backend/handlers/booking.go`
- `backend/handlers/booking_test.go`

---

## 🧪 Frontend Unit Tests

The following frontend test files are present and cover major UI flows, services, and page-level interactions:

### Test Files
- `frontend/src/App.test.js`
- `frontend/src/components/Footer.test.js`
- `frontend/src/components/Navbar.test.js`
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
- `frontend/src/services/api.test.js`

### Frontend Test Coverage Highlights
- App rendering and route behavior
- Navbar role-based links and logout behavior
- Login and registration flows
- Profile display and profile editing
- Services listing and service detail pages
- Booking page behavior
- Staff dashboard access control
- Admin dashboard workflows
- API service helpers and interceptor logic

### Sprint 3 Frontend Test Areas Especially Relevant
- `AdminDashboard.test.js`
  - dashboard rendering
  - admin-only behavior
  - bookings overview interactions
  - service management interactions
- `Navbar.test.js`
  - role-based navigation visibility
- `api.test.js`
  - authorization header behavior

---

## 🧪 Backend Unit Tests

The backend includes handler, middleware, router, model, validation, and utility tests.

### Backend Test Files
- `backend/config/config_test.go`
- `backend/handlers/approval_test.go`
- `backend/handlers/auth_test.go`
- `backend/handlers/booking_test.go`
- `backend/handlers/notification_test.go`
- `backend/handlers/review_test.go`
- `backend/handlers/service_test.go`
- `backend/handlers/test_helpers_test.go`
- `backend/handlers/user_test.go`
- `backend/main_test.go`
- `backend/middleware/middleware_test.go`
- `backend/middleware/test_helpers_test.go`
- `backend/models/models_test.go`
- `backend/router/routes_test.go`
- `backend/testutil/http_test.go`
- `backend/testutil/test_db_test.go`
- `backend/validation/validator_test.go`

### Backend Test Coverage By Area

#### User Handler Tests
- `TestGetUserReturnsUser`
- `TestGetAllUsersReturnsUsers`
- `TestGetUserReturnsNotFoundForMissingUser`
- `TestUpdateUserPersistsChanges`
- `TestUpdateUserReturnsBadRequestForInvalidPayload`
- `TestGetProfileReturnsBookingsAndReviews`
- `TestGetProfileReturnsNotFoundForMissingUser`

#### Notification Handler Tests
- `TestGetNotificationsReturnsUserNotifications`
- `TestCreateNotificationDefaultsUnread`
- `TestCreateNotificationReturnsBadRequestForInvalidPayload`
- `TestMarkAsReadUpdatesNotification`
- `TestMarkAsReadReturnsNotFoundForMissingNotification`
- `TestCreateNotificationRejectsMalformedJSON`

#### Booking Handler Tests
- `TestCreateBookingSetsPendingStatus`
- `TestGetBookingReturnsBookingWithRelations`
- `TestGetUserBookingsWithoutStatusFilterReturnsAllUserRecords`
- `TestGetUserBookingsWithValidStatusFilterReturnsMatchingRecords`
- `TestGetUserBookingsWithInvalidStatusFilterReturnsBadRequest`
- `TestGetUserBookingsWithNoMatchingStatusReturnsEmptyArray`
- `TestUpdateBookingPersistsChanges`
- `TestCancelBookingMarksStatusCancelled`

#### Service Handler Tests
- `TestListServicesReturnsAllServices`
- `TestGetServiceReturnsNotFoundForMissingService`
- `TestCreateServicePersistsRecord`
- `TestUpdateServicePersistsChanges`
- `TestUpdateServiceActiveActivatesServiceWithoutChangingOtherFields`
- `TestUpdateServiceActiveDeactivatesServiceAndAffectsActiveOnlyFilter`
- `TestUpdateServiceActiveReturnsBadRequestForInvalidPayload`
- `TestUpdateServiceActiveReturnsNotFoundForMissingService`
- `TestUpdateServiceActiveRejectsMalformedJSON`
- `TestDeleteServiceRemovesRecord`
- `TestGetServicesByCategoryFiltersResults`

#### Approval Handler Tests
- `TestGetPendingBookingsReturnsStaffServiceBookings`
- `TestGetAllBookingsForServiceReturnsAllStatuses`
- `TestApproveBookingUpdatesStatusAndCreatesNotification`
- `TestRejectBookingCreatesNotification`
- `TestRejectBookingRequiresMatchingStaffService`
- `TestGetAllPendingBookingsRequiresAdmin`
- `TestGetAllBookingsReturnsAdminView`
- `TestAdminApproveBookingUpdatesStatus`
- `TestAdminRejectBookingUpdatesStatus`

#### Auth Handler Tests
- `TestRegisterCreatesUser`
- `TestRegisterRejectsDuplicateEmail`
- `TestLoginAuthenticatesUser`
- `TestLoginRejectsInvalidPassword`
- `TestLogoutReturnsSuccess`
- `TestRefreshTokenReturnsStructuredToken`

#### Review Handler Tests
- `TestCreateReviewPersistsRecordAndUpdatesServiceRating`
- `TestGetServiceReviewsReturnsReviewsForService`
- `TestGetReviewReturnsSingleReview`
- `TestGetUserReviewsReturnsReviewsForUser`
- `TestDeleteReviewRespectsOwnershipAndUpdatesRating`

#### Middleware And Router Tests
- `TestAuthRequiredRejectsMissingAuthorizationHeader`
- `TestAuthRequiredAllowsBearerTokenAndStoresIt`
- `TestRequireRolesRejectsForbiddenRole`
- `TestRequireRolesAllowsMatchingRole`
- `TestCORSSetsHeadersAndShortCircuitsOptions`
- `TestErrorHandlerWritesInternalServerErrorForUnhandledErrors`
- `TestRequestLoggerAddsRequestIDHeader`
- `TestRequestLoggerPreservesIncomingRequestID`
- `TestProtectedRoutesRequireAuth`
- `TestPublicRoutesRemainAccessibleWithoutAuth`
- `TestProtectedRouteAllowsBearerToken`
- `TestApprovalRoutesRequireMatchingRole`
- `TestAdminUsersRouteRequiresAdminRole`
- `TestAdminUsersRouteAllowsAdminRole`
- `TestApprovalRoutesAllowMatchingRoleMiddleware`
- `TestAuthRequiredParsesStructuredTokenClaims`

### Backend Test Result
- The backend test suite passes with:

```bash
cd backend
go test ./...
```

---

## 🔗 Updated Backend API Documentation

This section reflects the current backend API after Sprint 3 enhancements.

## Authentication

### `POST /api/auth/register`
- Register a new user.

### `POST /api/auth/login`
- Login and receive lightweight auth token.

### `POST /api/auth/logout`
- Protected logout endpoint.

### `POST /api/auth/refresh`
- Protected token refresh endpoint.

---

## Users

### `GET /api/users/:id`
- Fetch a single user by ID.

### `GET /api/users/:id/profile`
- Fetch a user profile with preloaded bookings and reviews.

### `PUT /api/users/:id`
- Protected endpoint for updating user profile fields.

### `GET /api/users`
- Protected admin-only endpoint.
- Returns all users in the system.
- Added during Sprint 3 backend work.

---

## Services

### `GET /api/services`
- Returns services list.
- Supports optional filtering and pagination behavior already present in the handler.

### `GET /api/services/:id`
- Returns a single service with related review data.

### `GET /api/services/category/:category`
- Returns services for a specific category.

### `POST /api/services`
- Protected endpoint to create a new service.

### `PUT /api/services/:id`
- Protected endpoint to update service details.

### `PATCH /api/services/:id/active`
- Protected endpoint to activate or deactivate a service.
- Request body:

```json
{
  "isActive": true
}
```

### `DELETE /api/services/:id`
- Protected endpoint to delete a service.

---

## Bookings

### `POST /api/bookings`
- Protected endpoint to create a booking.

### `GET /api/bookings/:id`
- Returns booking details with related user and service information.

### `GET /api/bookings/user/:userId`
- Returns all bookings for a user.
- Supports optional `status` query parameter.

Example:

```http
GET /api/bookings/user/123?status=pending
```

Supported status values:
- `pending`
- `approved`
- `rejected`
- `completed`
- `cancelled`

### `PUT /api/bookings/:id`
- Protected endpoint to update booking details.

### `PATCH /api/bookings/:id/status`
- Protected endpoint to cancel a booking by setting status to `cancelled`.

---

## Approval

All approval routes are protected by authentication. Sprint 3 added role-based protection so users can only access the approval routes that match their role.

### Staff Approval Routes

#### `GET /api/approval/staff/:staffId/pending`
- Staff-only route.
- Returns pending bookings for the staff member’s service.

#### `GET /api/approval/staff/:staffId/all`
- Staff-only route.
- Returns all bookings for the staff member’s service.

#### `PUT /api/approval/bookings/:id/approve`
- Staff-only route.
- Approves a booking for the staff member’s service.

Request example:

```json
{
  "status": "approved",
  "approvalNotes": "Approved for requested slot",
  "staffId": "staff-user-id"
}
```

#### `PUT /api/approval/bookings/:id/reject`
- Staff-only route.
- Rejects a booking for the staff member’s service.

Request example:

```json
{
  "status": "rejected",
  "approvalNotes": "Conflicts with service availability",
  "staffId": "staff-user-id"
}
```

### Admin Approval Routes

#### `GET /api/approval/admin/:userId/pending`
- Admin-only route.
- Returns all pending bookings across the platform.

#### `GET /api/approval/admin/:userId/all`
- Admin-only route.
- Returns all bookings across the platform.

#### `PUT /api/approval/admin/:userId/bookings/:id/approve`
- Admin-only route.
- Approves any booking.

#### `PUT /api/approval/admin/:userId/bookings/:id/reject`
- Admin-only route.
- Rejects any booking.

### Approval Notification Behavior
- Student notifications are created when bookings are approved.
- Student notifications are created when bookings are rejected.
- This applies to both staff-driven and admin-driven approval flows.

---

## Notifications

### `GET /api/notifications/:userId`
- Protected endpoint returning notifications for a user.

### `POST /api/notifications`
- Protected endpoint creating a notification.

Request example:

```json
{
  "userId": "user-id",
  "title": "Booking Approved",
  "message": "Your booking has been approved.",
  "type": "booking_approval"
}
```

### `PUT /api/notifications/:id/read`
- Protected endpoint marking a notification as read.

---

## Reviews

### `GET /api/reviews/service/:serviceId`
- Returns reviews for a service.

### `GET /api/reviews/:id`
- Returns a single review.

### `POST /api/reviews`
- Protected endpoint creating a review.

### `DELETE /api/reviews/:id`
- Protected endpoint deleting a review.

---

## 📌 Sprint 3 Deliverables Summary

Sprint 3 successfully delivered:
- a full admin dashboard for viewing all bookings
- admin approval and rejection workflows
- create and delete service UI
- stronger backend role-based protection
- admin-only all-users endpoint
- booking notification verification for approval and rejection flows
- expanded unit test coverage across frontend and backend
- updated backend API documentation aligned with implemented routes

This sprint significantly improved both administrator capabilities and backend reliability for the Smart Campus Services Platform.
