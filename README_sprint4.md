# Smart Campus Services Platform

Smart Campus Services Platform is a full-stack role-based campus service management application that helps students discover and book campus services, allows staff to review and act on pending booking requests, and enables administrators to monitor bookings and manage services across the platform.

The system centralizes service discovery, booking workflows, notifications, and role-based dashboards into one platform, making campus operations more organized, efficient, and easier to manage.

---

## Team Members

- Venkata Sai Saran Jonnalagadda - 11114995
- Srikar Panuganti - 38909216
- Keerthi Reddy Gudibandi - 13652831
- Vishnu Sai Padyala - 32712860

---

## Project Features

### Student Functionality
- Register and log in with role-based access
- Browse campus services
- Search and filter services
- View service details
- Create bookings for services
- View booking history
- Filter bookings by status
- Cancel eligible bookings
- View and edit profile information
- Submit and manage reviews
- View notifications for booking updates

### Staff Functionality
- Access staff dashboard
- View pending booking requests for assigned services
- Approve or reject booking requests
- Add notes while processing requests

### Admin Functionality
- Access admin dashboard
- View platform-wide booking statistics
- View all bookings across the platform
- Filter bookings by status and service
- Approve or reject bookings
- Manage available services
- Create new services
- Delete outdated services

---

## Tech Stack

### Frontend
- React
- React Router
- Axios
- Jest
- React Testing Library
- Cypress

### Backend
- Go
- Gin
- GORM
- SQLite

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
├── Sprint2.md
├── Sprint3.md
├── Sprint4.md
└── README.md
```

---

## How to Run the Project

## 1. Clone the Repository

```bash
git clone https://github.com/saranjv426/Smart-Campus-Services-Platform.git
cd Smart-Campus-Services-Platform
```

---

## 2. Run the Backend

Open a terminal and run:

```bash
cd backend
go run .
```

The backend will start on:

```text
http://localhost:8080
```

---

## 3. Run the Frontend

Open a second terminal and run:

```bash
cd frontend
npm install
npm start
```

The frontend will start on:

```text
http://localhost:3000
```

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

Examples:

```bash
cd frontend
npm test -- --watchAll=false src/pages/Profile.test.js
npm test -- --watchAll=false src/pages/StaffDashboard.test.js
npm test -- --watchAll=false src/pages/AdminDashboard.test.js
npm test -- --watchAll=false src/pages/Bookings.test.js
```

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

Current Cypress flow includes:
- loading the application home page
- clicking the **Get Started** button
- verifying navigation to the Register page
- confirming Register page content is displayed

---

## Backend Testing

Run all backend tests:

```bash
cd backend
go test ./...
```

If needed, use a local Go cache:

```bash
cd backend
GOCACHE=$(pwd)/.gocache go test ./...
```

---

## Main Frontend Test Files

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

---

## Main Backend Test Files

- `backend/handlers/approval_test.go`
- `backend/handlers/auth_test.go`
- `backend/handlers/booking_test.go`
- `backend/handlers/notification_test.go`
- `backend/handlers/review_test.go`
- `backend/handlers/service_test.go`
- `backend/handlers/test_helpers_test.go`
- `backend/handlers/user_test.go`
- `backend/main_test.go`

---

## Backend API Overview

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

### Bookings
- `POST /api/bookings`
- `GET /api/bookings/:id`
- `GET /api/bookings/user/:userId`
- `PUT /api/bookings/:id`
- `PATCH /api/bookings/:id/status`

### Approval
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

### Notifications
- `GET /api/notifications/:userId`
- `POST /api/notifications`
- `PUT /api/notifications/:id/read`

### Reviews
- `POST /api/reviews`
- `GET /api/reviews/service/:serviceId`
- `GET /api/reviews/user/:userId`
- `GET /api/reviews/:id`
- `DELETE /api/reviews/:id`

---

## Project Summary

Smart Campus Services Platform provides a complete role-based workflow for students, staff, and administrators:

- Students can discover services, create bookings, manage their profiles, and track updates
- Staff can process pending booking requests for their assigned services
- Admins can oversee bookings platform-wide and manage available services

This makes the platform a centralized, structured, and user-friendly solution for campus service management.
