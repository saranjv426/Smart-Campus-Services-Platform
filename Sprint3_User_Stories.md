# Sprint 3 User Stories - Smart Campus Services Platform

This document collects the Sprint 3 issues in a consistent format using:

- `Description`
- `User Story`
- `Acceptance Criteria`

It includes both backend-focused and admin dashboard/frontend-focused Sprint 3 work items.

---

## 1. Strengthen User and Notification Handler Test Coverage

### Description
Expand and standardize backend tests for user and notification handlers. Basic happy-path tests already exist in `user_test.go` and `notification_test.go`, but this issue covers missing negative cases, validation behavior, and edge conditions so those handlers are more fully protected against regressions.

### User Story
As a developer, I want stronger automated coverage for user and notification handlers so that profile and notification changes can be made confidently without breaking key API behavior.

### Acceptance Criteria
- User handler tests cover:
  - get user success
  - get user not found
  - update user success
  - update user with invalid payload
  - get profile success with preloaded bookings and reviews
  - get profile not found
- Notification handler tests cover:
  - get notifications success
  - create notification success with `isRead=false` by default
  - create notification invalid payload
  - mark notification as read success
  - mark notification as read for missing notification
- Tests assert both HTTP status codes and important response fields.
- Tests use the project’s existing test DB and test helper patterns.
- The full backend test suite passes after the new tests are added.

---

## 2. Toggle Service Active State Endpoint

### Description
Add a dedicated endpoint to toggle a service’s `isActive` state. The `Service` model already includes `isActive`, and `ListServices` already supports `activeOnly=true`, but there is no focused endpoint for admins to activate or deactivate a service without sending a full update payload.

### User Story
As an admin, I want to quickly activate or deactivate a campus service so that unavailable services can be hidden from active use without deleting the service record.

### Acceptance Criteria
- A protected endpoint is added for toggling service active state, for example `PATCH /api/services/:id/active`.
- The endpoint accepts a boolean `isActive` field in the request body.
- If the service exists, its `isActive` value is updated and the updated service is returned.
- If the service does not exist, the API returns `404 Not Found`.
- Invalid request bodies return `400 Bad Request`.
- Existing service fields are not changed when only active status is updated.
- `GET /api/services?activeOnly=true` reflects the updated active state correctly.
- Automated tests cover:
  - activating a service
  - deactivating a service
  - invalid payload
  - missing service
  - protected route requiring auth

---

## 3. Filter User Bookings by Status

### Description
Add support for filtering user bookings by a `status` query parameter on `GET /api/bookings/user/:userId`. Previously `booking.go` returned all bookings for a user, while the frontend filtered them client-side. This issue moves that filtering into the backend so clients can request only the statuses they need.

### User Story
As a student, I want to fetch only bookings with a specific status so that I can quickly view pending, approved, rejected, or cancelled requests without downloading and filtering everything on the client.

### Acceptance Criteria
- `GET /api/bookings/user/:userId` accepts an optional `status` query parameter.
- When `status` is omitted, the endpoint returns all bookings for that user.
- When `status` is provided, the endpoint returns only bookings for that user matching that status.
- Supported statuses align with the booking model, such as:
  - `pending`
  - `approved`
  - `rejected`
  - `completed`
  - `cancelled`
- Invalid status values return `400 Bad Request` with a clear error message.
- Returned bookings still preload service details as they do today.
- Automated tests cover:
  - no status filter
  - valid status filter
  - invalid status filter
  - no matching records returning an empty array

---

## 4. Create New Service

### Description
Add an admin-facing create service workflow in the dashboard so administrators can add new campus services through a structured form instead of relying on manual backend-only operations.

### User Story
As an Administrator, I want to create a new campus service through an intuitive form interface so that I can add new services to the platform.

### Acceptance Criteria
- Admin can access a `Create Service` form from the admin dashboard.
- The form includes fields for:
  - Service Name
  - Category
  - Description
  - Location
  - Phone
  - Email
  - Hours
  - Service Image URL
- The form validates all required fields before submission.
- Admin receives a success confirmation message upon service creation.
- New service appears immediately in the services list.
- Service is set to `Active` by default when created.
- Cancel button allows admin to discard changes and return to dashboard.
- Form displays validation errors for invalid inputs such as:
  - email format
  - empty required fields
- Service image URL is validated before submission.

---

## 5. Delete Service

### Description
Add an admin-facing delete service workflow so administrators can remove services that are no longer offered and keep the platform current and accurate.

### User Story
As an Administrator, I want to delete services that are no longer available through the admin interface so that I can keep the platform up-to-date with current offerings.

### Acceptance Criteria
- Admin can select a service from the services list to delete.
- A confirmation modal appears before deletion to prevent accidental removal.
- Confirmation modal displays the service name and warns about associated bookings.
- Upon confirmation, the service is deleted from the database.
- Success message confirms service deletion.
- Service is immediately removed from the UI.
- Deletion is prevented or warned if service has active or pending bookings.
- Admin can cancel the deletion action in the confirmation modal.
- Deleted services are permanently removed and not just marked inactive.
- Error handling displays if deletion fails, such as a database constraint error.

---

## 6. Approve Booking

### Description
Add admin approval functionality so administrators can review and approve pending service bookings from the dashboard with an auditable workflow and student notification support.

### User Story
As an Administrator, I want to approve pending service bookings from students and staff so that I can manage and confirm service reservations.

### Acceptance Criteria
- Admin can view pending bookings in a dedicated section of the dashboard.
- Each booking displays:
  - Student Name
  - Service Name
  - Booking Date and Time
  - Status
- Admin can click an `Approve` button on any pending booking.
- Upon approval, booking status changes to `Approved`.
- Automatic notification is sent to the student confirming approval.
- Approved booking is moved out of the pending section.
- Success confirmation message displays for the admin.
- Booking details modal or view shows complete information before approval.
- Admin can add notes or comments when approving, optionally.
- Approval is timestamped in the database.

---

## 7. Reject Booking

### Description
Add admin rejection functionality so administrators can reject invalid, inappropriate, or conflicting bookings and capture the reason for the rejection.

### User Story
As an Administrator, I want to reject inappropriate or duplicate bookings so that I can maintain service quality and prevent scheduling conflicts.

### Acceptance Criteria
- Admin can click a `Reject` button on any pending booking.
- A modal appears for admin to enter rejection reason.
- Rejection reason field is required before submission.
- Upon rejection, booking status changes to `Rejected`.
- Automatic notification is sent to the student with rejection reason.
- Rejected booking is removed from the pending section and appears in rejection history.
- Success confirmation message displays for the admin.
- Admin can view rejection reason in booking history or booking details.
- Rejection is timestamped and attributed to the admin user.
- Rejected bookings cannot be reactivated without explicit re-creation.

---

## 8. View All Bookings Dashboard

### Description
Build a comprehensive admin dashboard that shows all system bookings in one place with filtering, sorting, search, and summary statistics to support administrative monitoring and decision-making.

### User Story
As an Administrator, I want to view a comprehensive dashboard showing all system bookings with filtering and sorting so that I can monitor platform activity and manage bookings effectively.

### Acceptance Criteria
- Dashboard displays a table or list of all bookings with columns:
  - Student Name
  - Service Name
  - Booking Date and Time
  - Status
  - Actions
- Bookings are displayed with status indicators for:
  - Pending
  - Approved
  - Rejected
  - Completed
  - Cancelled
- Dashboard shows total counts for:
  - Total Bookings
  - Pending
  - Approved
  - Rejected
  - Completed
- Admin can filter bookings by status:
  - All
  - Pending
  - Approved
  - Rejected
  - Completed
- Admin can filter bookings by service category.
- Admin can sort bookings by:
  - date
  - student name
  - status
- Dashboard loads all bookings efficiently, with pagination if there are 50 or more bookings.
- Each booking row has action buttons:
  - View Details
  - Approve
  - Reject
  - Cancel
- Search functionality allows admin to find bookings by student name or service name.
- Date range filter allows filtering bookings by creation date or booking date.
- Dashboard uses color coding or icons to visually distinguish booking statuses.
- Last refresh timestamp is displayed.
- Admin can manually refresh the dashboard to update data.

---

## 9. View Booking Details Modal

### Description
Add a booking details modal to the admin dashboard so administrators can inspect complete booking information before taking action such as approval or rejection.

### User Story
As an Administrator, I want to click on any booking to view complete details so that I can make informed decisions about approvals and rejections.

### Acceptance Criteria
- Clicking `View Details` opens a modal with full booking information.
- Modal displays:
  - Student details: name, email, phone
  - Service details: name, description, location
  - Booking details: date, time, duration
  - Status history
- Modal shows any notes or comments associated with the booking.
- Admin can approve or reject the booking directly from the details modal.
- Modal has a close button to return to the dashboard.
- Booking details are read-only unless admin is taking an action.
- Modal is responsive and works well on different screen sizes.
- Modal displays review or rating for the service if available.
- Modal shows any special requests or requirements from the student.
- Admin can see timestamp of when booking was created and last modified.

---

## 10. Notification on Booking Approval and Rejection

### Description
Verify and strengthen notification behavior for booking approval and rejection flows. The approval handlers already created notification records in some paths, but this issue ensures that approval and rejection actions consistently create notifications for students across both staff-driven and admin-driven approval flows. The issue also expands automated tests so these behaviors are protected from regressions.

### User Story
As a student, I want to receive a notification whenever my booking is approved or rejected so that I am always informed about changes to my booking status.

### Acceptance Criteria
- Staff approval creates a notification for the booking’s student.
- Staff rejection creates a notification for the booking’s student.
- Admin approval creates a notification for the booking’s student.
- Admin rejection creates a notification for the booking’s student.
- Approval notifications include:
  - title indicating approval
  - type set to `booking_approval`
- Rejection notifications include:
  - title indicating rejection
  - type set to `booking_rejection`
- Notification records are created with `isRead=false` by default.
- Notification messages include meaningful context such as:
  - service name
  - approval notes or rejection reason where applicable
- Approval handler tests cover:
  - staff approve updates booking and creates notification
  - staff reject creates notification
  - admin approve updates booking and creates notification
  - admin reject updates booking and creates notification
- Tests assert both booking status changes and important notification fields.
- The full backend test suite passes after the new tests are added.

---

## 11. GET All Users Endpoint for Admin

### Description
Add a backend endpoint that returns all users for admin management purposes and protect it with `RequireRoles("admin")`. The user handler already supported fetching a single user and a user profile, but there was no admin-only endpoint to retrieve the full user list. This issue adds `GET /api/users`, wires it through routing, and ensures that only admin users can access it.

### User Story
As an admin, I want to fetch the full list of users so that I can support administrative user management features while keeping that data restricted from non-admin roles.

### Acceptance Criteria
- A new endpoint exists: `GET /api/users`.
- The endpoint returns all users from the database.
- The endpoint is protected by authentication.
- The endpoint is additionally protected by `RequireRoles("admin")`.
- An admin token receives `200 OK` and a JSON array of users.
- A staff token receives `403 Forbidden`.
- A student token receives `403 Forbidden`.
- The response includes important user fields such as:
  - `id`
  - `email`
  - `firstName`
  - `lastName`
  - `role`
- Automated tests cover:
  - get all users success
  - protected route requiring auth
  - non-admin access rejected
  - admin access allowed
- The full backend test suite passes after the endpoint and tests are added.

---

## 12. RequireRoles Middleware on Approval Routes

### Description
Add role-based authorization to approval-related backend routes by applying the existing `RequireRoles(...)` middleware in routing. The approval endpoints were already protected by authentication, but they were not consistently restricted by user role. This issue ensures that staff-only approval routes can only be accessed by staff users, and admin approval routes can only be accessed by admin users. It also aligns the lightweight auth token flow so role information is available in request context for middleware enforcement.

### User Story
As a developer, I want approval routes to enforce role-based access so that only authorized staff and admin users can perform approval-related actions and the backend is protected from incorrect cross-role access.

### Acceptance Criteria
- Approval routes are protected by authentication and role middleware.
- Staff-only approval endpoints can only be accessed by users with the `staff` role.
- Admin-only approval endpoints can only be accessed by users with the `admin` role.
- A staff token can access:
  - `GET /api/approval/staff/:staffId/pending`
  - `GET /api/approval/staff/:staffId/all`
  - `PUT /api/approval/bookings/:id/approve`
  - `PUT /api/approval/bookings/:id/reject`
- A staff token is rejected with `403 Forbidden` on admin approval endpoints.
- An admin token can access:
  - `GET /api/approval/admin/:userId/pending`
  - `GET /api/approval/admin/:userId/all`
  - `PUT /api/approval/admin/:userId/bookings/:id/approve`
  - `PUT /api/approval/admin/:userId/bookings/:id/reject`
- An admin token is rejected with `403 Forbidden` on staff approval endpoints.
- Auth middleware populates request context with role information from the lightweight token format.
- Route tests verify:
  - missing authorization returns `401 Unauthorized`
  - wrong role returns `403 Forbidden`
  - correct role is allowed
- The full backend test suite passes after the changes are added.
