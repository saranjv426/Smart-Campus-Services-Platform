# Sprint 1 Detailed Report

## Sprint Goal
Establish a production-style backend foundation for Smart Campus Services, deliver core API modules, and set up a reusable testing framework so future features can be built with lower risk and faster iteration.

## Sprint Scope
This sprint focused on backend-only work:
- API server setup and route organization
- database and model infrastructure
- seed/test data initialization
- service CRUD APIs
- authentication APIs
- user and notification APIs
- test utilities and baseline unit tests

## User Stories

### Student-facing stories
- As a student, I want to browse available campus services so I can quickly find the right resource.
- As a student, I want to view service details so I can make informed decisions before booking.
- As a student, I want to register and log in so I can access personalized data.
- As a student, I want to manage my profile so my account information stays current.
- As a student, I want to receive notifications so I can track booking and service updates.

### Staff/Admin-facing stories
- As a staff/admin user, I want CRUD APIs for services so I can maintain accurate service listings.
- As a staff/admin user, I want backend structures that support approvals/workflows so operations scale reliably.

### Engineering stories
- As a development team, we want clear backend layering (handlers, router, middleware, validation, models) so the system is maintainable.
- As a development team, we want test utilities and unit-test baseline coverage so regressions are caught early.

## Issues Planned for Sprint 1
- #4 Setup Project, Models & Database Infrastructure
- #5 Seed All Services & Test Data
- #6 Create All Service CRUD Handlers & Routes
- #7 Authentication System & Handlers
- #8 User & Notification Handler System
- #9 API Server Setup & Global Testing Framework

## Delivery Status by Planned Issue

### #4 Setup Project, Models & Database Infrastructure
**Status:** Completed

**Delivered:**
- Go backend module and project structure finalized.
- Core models implemented in `backend/models/models.go`:
  - `User`, `Service`, `Booking`, `Notification`, `Review`
- UUID-based record IDs via model hooks.
- Database bootstrap and migration flow in `backend/main.go`.
- SQLite configured for local development and quick setup.

**Outcome:**
- A stable relational backend data model was established and became the base for all API modules.

### #5 Seed All Services & Test Data
**Status:** Completed

**Delivered:**
- Seed logic implemented in `backend/seed.go`.
- Seed endpoint exposed via `POST /api/seed`.
- Initial service/test dataset available for demos and API validation.

**Outcome:**
- Team could test feature flows quickly without manual data entry.

### #6 Create All Service CRUD Handlers & Routes
**Status:** Completed

**Delivered:**
- Service handlers in `backend/handlers/service.go`.
- Routes wired through router setup.
- Implemented endpoints:
  - `GET /api/services`
  - `GET /api/services/:id`
  - `POST /api/services`
  - `PUT /api/services/:id`
  - `DELETE /api/services/:id`
  - `GET /api/services/category/:category`

**Outcome:**
- Complete Service API lifecycle delivered and ready for frontend integration.

### #7 Authentication System & Handlers
**Status:** Completed

**Delivered:**
- Auth handlers in `backend/handlers/auth.go`:
  - Register, Login, Logout, RefreshToken
- Request validation with Gin binding tags.
- Custom role validation integrated.
- Auth response includes `role` and `serviceId` for role-aware client behavior.

**Outcome:**
- Working authentication API contract established for frontend auth flows.

### #8 User & Notification Handler System
**Status:** Completed

**Delivered:**
- User handlers in `backend/handlers/user.go`:
  - `GetUser`, `UpdateUser`, `GetProfile`
- Notification handlers in `backend/handlers/notification.go`:
  - `GetNotifications`, `CreateNotification`, `MarkAsRead`

**Outcome:**
- Core account/profile and notification workflows now supported in backend.

### #9 API Server Setup & Global Testing Framework
**Status:** Completed

**Delivered:**
- Middleware extraction and setup:
  - `backend/middleware/cors.go`
  - `backend/middleware/request_logger.go`
  - `backend/middleware/error_handler.go`
  - `backend/middleware/auth.go`
- Centralized route registration:
  - `backend/router/routes.go`
- Validation initialization:
  - `backend/validation/validator.go`
- Test utilities:
  - `backend/testutil/test_db.go`
  - `backend/testutil/http.go`
- Handler unit tests added/expanded:
  - `backend/handlers/auth_test.go`
  - `backend/handlers/service_test.go`
  - `backend/handlers/booking_test.go`
  - `backend/handlers/user_test.go`
  - `backend/handlers/notification_test.go`

**Outcome:**
- Backend gained a maintainable global structure and repeatable test foundation.

## Successfully Completed (Summary)
All planned Sprint 1 backend issues were completed:
- #4 Completed
- #5 Completed
- #6 Completed
- #7 Completed
- #8 Completed
- #9 Completed

## Not Completed and Why
No planned Sprint 1 issue remained incomplete.

That said, some items were intentionally left as future hardening tasks (outside Sprint 1 scope):
- Production-grade JWT/token lifecycle hardening
- Password hashing enforcement across all auth paths
- Expanded authorization guards on every protected route
- Broader integration/end-to-end test suite

These were not delivery failures; they were deferred by scope prioritization to keep Sprint 1 focused on backend foundation and core API readiness.

## Testing and Validation Evidence
- APIs were exercised via Postman and CLI (`curl`) for endpoint-level verification.
- Unit tests validate both success and failure paths (validation errors, conflict, unauthorized, not found).
- In-memory DB test setup ensures deterministic test execution without polluting local runtime data.

## Sprint Retrospective (Backend)
### What went well
- Clear modular backend organization improved implementation speed.
- Test utility abstraction reduced duplicate setup code.
- Seeded data accelerated debugging and demos.

### What can be improved
- Add stricter API response standardization for all handlers.
- Increase negative-path test coverage for every endpoint.
- Add CI automation to run test suite on every PR.

### Next sprint recommendations
- Complete security hardening (JWT + bcrypt + tighter route protection).
- Add pagination/filtering enhancements where needed.
- Add integration tests spanning auth -> user -> booking -> notification flow.
