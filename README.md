# Smart Campus Services Platform

Smart Campus Services Platform is a comprehensive full-stack, role-based web application for discovering, booking, reviewing, and managing campus services. Students can browse services, create bookings, review services, and manage their profiles. Staff and administrators can review booking requests, manage services, and monitor platform activity through dedicated dashboards.

The platform centralizes service discovery, booking workflows, notifications, analytics, and role-based dashboards into one system, making campus operations more organized, efficient, and easier to manage.

**Status**: ✅ Active Project  
**Frontend**: React  
**Backend**: Go + Gin + GORM + SQLite

---

## Team Members

- Venkata Sai Saran Jonnalagadda - 11114995
- Srikar Panuganti - 38909216
- Keerthi Reddy Gudibandi - 13652831
- Vishnu Sai Padyala - 32712860

---

## Project Overview

### What This Project Does

- Students can browse and book campus services
- Students can search, filter, and review services
- Students can manage personal bookings and profile information
- Staff can review pending booking requests for assigned services
- Admins can monitor bookings across the platform and manage services
- Users receive notifications for booking status updates
- The UI supports search, filtering, dashboards, and role-based workflows

### Main Roles

#### Student
- Register and log in
- Browse, search, and filter services
- View service details
- Create bookings
- View booking history
- Filter bookings by status
- Cancel eligible bookings
- Edit profile information
- Submit and manage reviews
- View booking update notifications

#### Staff
- Access staff dashboard
- View pending booking requests for assigned services
- Approve or reject booking requests
- Add notes while processing requests

#### Admin
- Access admin dashboard
- View platform-wide booking statistics
- View all bookings
- Filter bookings by status and service
- Approve or reject bookings
- Create new services
- Update service state
- Delete outdated services

---

## Key Features

### For Students
- Service discovery by category
- Search and filtering for services
- Detailed service pages with reviews and booking form
- Booking creation and booking history
- Booking filtering and cancellation
- Profile editing and review management
- Notification support for booking updates

### For Staff
- Pending booking review dashboard
- Approval and rejection workflow
- Optional notes for approval and rejection actions

### For Admins
- Dashboard with platform-wide booking overview
- Booking statistics cards
- Advanced filtering by status and service
- Approval and rejection of bookings
- Service catalog management
- Service creation and deletion
- Service activation state support via backend API

---

## Requirements

Install these tools before running the application:

| Tool | Required version | Used for |
| --- | --- | --- |
| Git | Current stable version | Cloning and version control |
| Go | 1.21 or newer | Backend API |
| Node.js | 14 or newer | React frontend |
| npm | Included with Node.js | Frontend dependencies and scripts |
| SQLite | No separate server required | Local database storage |

The backend stores data in a local SQLite file at `backend/data/smart_campus.db` by default.

---

## Technology Stack

### Frontend
- React
- React Router
- Axios
- Jest
- React Testing Library
- Cypress
- CSS3

### Backend
- Go
- Gin
- GORM
- SQLite

### Infrastructure / Local Development
- Frontend default port: `3000`
- Backend default port: `8080`
- Database file: `backend/data/smart_campus.db`

---

## Project Structure

```text
Smart-Campus-Services-Platform/
├── backend/
│   ├── config/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── router/
│   ├── testutil/
│   ├── validation/
│   ├── data/
│   └── main.go
├── frontend/
│   ├── cypress/
│   ├── public/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── services/
│   │   └── styles/
│   └── package.json
├── TESTING_GUIDE.md
├── Sprint2.md
├── Sprint3.md
├── Sprint4.md
└── README.md
```

---

## Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/saranjv426/Smart-Campus-Services-Platform.git
cd Smart-Campus-Services-Platform
```

### 2. Install Backend Dependencies

```bash
cd backend
go mod download
cd ..
```

### 3. Install Frontend Dependencies

```bash
cd frontend
npm install
cd ..
```

### 4. Configure Frontend Environment

Create the frontend environment file if needed:

```bash
cp frontend/.env.example frontend/.env
```

Default value:

```env
REACT_APP_API_URL=http://localhost:8080/api
```

The backend does not require a `.env` file for standard local development, though optional backend environment variables are supported.

---

## Running the Application

### Run the Backend

Open a terminal and run:

```bash
cd backend
go run .
```

The backend runs at:

```text
http://localhost:8080
```

### Run the Frontend

Open a second terminal and run:

```bash
cd frontend
npm start
```

The frontend runs at:

```text
http://localhost:3000
```

Then open `http://localhost:3000` in a browser.

---

## Seeding Local Data

After the backend is running, seed the local SQLite database with development data:

```bash
curl -X POST http://localhost:8080/api/seed
```

The seed endpoint is intended for local development and creates sample users, services, bookings, and related data used by the application flows.

---

## Using the Application

A common local workflow is:

1. Start the backend with `go run .`
2. Start the frontend with `npm start`
3. Seed the database with `POST /api/seed`
4. Register or log in from the frontend
5. Browse services from the home or services page
6. Open a service detail page to create bookings or reviews
7. Use staff or admin dashboards to review and manage bookings and services

---

## Configuration

### Backend Environment Variables

| Variable | Default | Description |
| --- | --- | --- |
| `PORT` | `8080` | Backend API port |
| `DB_PATH` | `data/smart_campus.db` | SQLite database file path |
| `GIN_MODE` | `debug` | Gin runtime mode, such as `debug` or `release` |

Example backend `.env` file:

```env
PORT=8080
DB_PATH=data/smart_campus.db
GIN_MODE=debug
```

### Frontend Environment Variables

| Variable | Default | Description |
| --- | --- | --- |
| `REACT_APP_API_URL` | `http://localhost:8080/api` | Backend API base URL used by the React app |

---

## Frontend Pages and Components

### Main Pages
- Home Page
- Services Page
- Service Detail Page
- Login Page
- Register Page
- Bookings Page
- Profile Page
- Admin Dashboard
- Staff Dashboard

### Shared Areas
- Navbar
- Footer
- API service helpers

---

## Frontend Testing

Run all frontend unit tests:

```bash
cd frontend
npm test -- --watchAll=false
```

Run a specific frontend test file:

```bash
cd frontend
npm test -- --watchAll=false src/pages/Profile.test.js
```

Run multiple specific files:

```bash
cd frontend
npm test -- --watchAll=false src/pages/Profile.test.js src/pages/StaffDashboard.test.js src/pages/AdminDashboard.test.js src/pages/Bookings.test.js
```

### Main Frontend Test Files

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

### Frontend Coverage Areas

- App and route behavior
- Login and registration flows
- Profile rendering and profile editing
- Services listing and filtering
- Service detail rendering and booking form behavior
- Booking list rendering, filtering, refresh, and cancellation
- Staff dashboard approve/reject workflow
- Admin dashboard statistics, filtering, booking actions, and service management

---

## Cypress End-to-End Testing

Open Cypress:

```bash
cd frontend
npm run cypress:open
```

Run Cypress in headless mode:

```bash
cd frontend
npm run cypress:run
```

### Current Cypress Flow

The current Cypress test covers a basic end-to-end navigation flow:

- loads the application home page
- clicks the **Get Started** button
- verifies navigation to the Register page
- confirms Register page content is displayed

### Cypress Test File

- `frontend/cypress/e2e/spec.cy.js`

More detailed testing guidance is available in `TESTING_GUIDE.md`.

---

## Backend Testing

Run all backend tests:

```bash
cd backend
go test ./...
```

If your environment blocks Go from writing to the default build cache, use a local cache:

```bash
cd backend
GOCACHE=$(pwd)/.gocache go test ./...
```

Or:

```bash
cd backend
GOCACHE=/tmp/go-build-cache go test ./...
```

Run only backend handler tests:

```bash
cd backend
go test ./handlers
```

### Main Backend Test Files

- `backend/handlers/approval_test.go`
- `backend/handlers/auth_test.go`
- `backend/handlers/booking_test.go`
- `backend/handlers/notification_test.go`
- `backend/handlers/review_test.go`
- `backend/handlers/service_test.go`
- `backend/handlers/test_helpers_test.go`
- `backend/handlers/user_test.go`
- `backend/main_test.go`

### Backend Coverage Areas

- Authentication handlers
- User profile handlers
- Service CRUD handlers
- Booking create/read/update/cancel handlers
- Review handlers
- Notification handlers
- Approval and rejection handlers
- Main package / setup tests

---

## Backend API Overview

The backend API is available under `/api`.

### Authentication

- `POST /api/auth/register`
- `POST /api/auth/login`
- `POST /api/auth/logout`
- `POST /api/auth/refresh`

### Users

- `GET /api/users/:id`
- `GET /api/users/:id/profile`
- `PUT /api/users/:id`
- `GET /api/users`

### Services

- `GET /api/services`
- `GET /api/services/:id`
- `GET /api/services/category/:category`
- `POST /api/services`
- `PUT /api/services/:id`
- `PATCH /api/services/:id/active`
- `DELETE /api/services/:id`

`GET /api/services` supports `q`, `category`, `activeOnly`, `limit`, `offset`, `sortBy`, and `sortOrder`.

### Bookings

- `POST /api/bookings`
- `GET /api/bookings/:id`
- `GET /api/bookings/user/:userId`
- `PUT /api/bookings/:id`
- `PATCH /api/bookings/:id/status`

### Reviews

- `POST /api/reviews`
- `GET /api/reviews/service/:serviceId`
- `GET /api/reviews/user/:userId`
- `GET /api/reviews/:id`
- `DELETE /api/reviews/:id`

### Notifications

- `GET /api/notifications/:userId`
- `POST /api/notifications`
- `PUT /api/notifications/:id/read`

### Approval Routes

#### Staff Approval Routes
- `GET /api/approval/staff/:staffId/pending`
- `GET /api/approval/staff/:staffId/all`
- `PUT /api/approval/bookings/:id/approve`
- `PUT /api/approval/bookings/:id/reject`

#### Admin Approval Routes
- `GET /api/approval/admin/:userId/pending`
- `GET /api/approval/admin/:userId/all`
- `PUT /api/approval/admin/:userId/bookings/:id/approve`
- `PUT /api/approval/admin/:userId/bookings/:id/reject`

### Health and Seed

- `GET /health`
- `POST /api/seed`

---

## Common Local Workflow Examples

### Start Backend
```bash
cd backend
go run .
```

### Start Frontend
```bash
cd frontend
npm start
```

### Seed Database
```bash
curl -X POST http://localhost:8080/api/seed
```

### Run Frontend Tests
```bash
cd frontend
npm test -- --watchAll=false
```

### Run Backend Tests
```bash
cd backend
go test ./...
```

---

## Troubleshooting

### Backend port already in use

```bash
cd backend
PORT=8081 go run .
```

### Frontend cannot reach the backend

- Confirm the backend is running at `http://localhost:8080`
- Confirm `frontend/.env` contains `REACT_APP_API_URL=http://localhost:8080/api`
- Restart the frontend after changing `.env`

### Database issues

- Confirm you are running backend commands from the `backend/` directory
- The default database file is `backend/data/smart_campus.db`
- For a clean local database, stop the backend, remove the SQLite file, restart the backend, and run the seed endpoint again

### Frontend tests failing

```bash
rm -rf node_modules package-lock.json
npm install
npm test -- --clearCache --watchAll=false
```

### Go test cache permission issues

```bash
cd backend
GOCACHE=/tmp/go-build-cache go test ./...
```

---

## Documentation Files

- `README.md` - setup, usage, testing, and API overview
- `TESTING_GUIDE.md` - testing instructions
- `Sprint2.md` - Sprint 2 report
- `Sprint3.md` - Sprint 3 report
- `Sprint4.md` - Sprint 4 report

---

## Project Summary

Smart Campus Services Platform provides a complete role-based workflow for students, staff, and administrators:

- Students can discover services, create bookings, manage their profiles, and track updates
- Staff can process pending booking requests for their assigned services
- Admins can oversee bookings platform-wide and manage available services

This makes the platform a centralized, structured, and user-friendly solution for campus service management.
