# Smart Campus Services Platform

Smart Campus Services Platform is a full-stack web application for discovering, booking, reviewing, and managing campus services. Students can browse services, create bookings, review services, and manage their profiles. Staff and administrators can manage services, review booking requests, and track service activity.

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

## Project Structure

```text
Smart-Campus-Services-Platform/
|-- backend/                 # Go, Gin, GORM, SQLite API
|   |-- handlers/            # HTTP handlers and handler tests
|   |-- models/              # Database models
|   |-- router/              # Route setup tests
|   |-- testutil/            # Shared backend test helpers
|   |-- validation/          # Request validation helpers
|   `-- main.go              # Backend entry point
|-- frontend/                # React application
|   |-- src/pages/           # Main UI pages
|   |-- src/components/      # Shared UI components
|   `-- src/services/        # API client helpers
|-- TESTING_GUIDE.md         # Backend and manual testing instructions
`-- README.md                # Front-page setup and usage guide
```

## Quick Start

Clone the repository:

```bash
git clone https://github.com/saranjv426/Smart-Campus-Services-Platform.git
cd Smart-Campus-Services-Platform
```

Install backend dependencies:

```bash
cd backend
go mod download
cd ..
```

Install frontend dependencies:

```bash
cd frontend
npm install
cd ..
```

Create the frontend environment file:

```bash
cp frontend/.env.example frontend/.env
```

The default value is:

```env
REACT_APP_API_URL=http://localhost:8080/api
```

The backend does not require a `.env` file for local development. Optional backend environment variables are listed below.

## Running The Application

Start the backend API from one terminal:

```bash
cd backend
go run main.go
```

The backend runs at `http://localhost:8080`.

Start the frontend from a second terminal:

```bash
cd frontend
npm start
```

The frontend runs at `http://localhost:3000`.

Open `http://localhost:3000` in a browser to use the application.

## Seeding Local Data

After the backend is running, seed the local SQLite database with development data:

```bash
curl -X POST http://localhost:8080/api/seed
```

The seed endpoint is intended for local development. It creates sample users, services, bookings, and related data used by the app flows.

## Using The Application

Common local workflow:

1. Start the backend with `go run main.go`.
2. Start the frontend with `npm start`.
3. Seed the database with `POST /api/seed`.
4. Register or log in from the frontend.
5. Browse services from the home or services page.
6. Open a service detail page to create bookings or reviews.
7. Use staff or admin dashboards to review and manage bookings and services.

Main user roles:

- `student`: browse services, create bookings, manage personal bookings, write reviews.
- `staff`: manage booking approvals for the staff member's assigned service.
- `admin`: manage broader service and booking workflows.

## Configuration

Backend environment variables are optional for local development:

| Variable | Default | Description |
| --- | --- | --- |
| `PORT` | `8080` | Backend API port |
| `DB_PATH` | `data/smart_campus.db` | SQLite database file path, relative to `backend/` when running from that directory |
| `GIN_MODE` | `debug` | Gin runtime mode, such as `debug` or `release` |

Example backend `.env` file in `backend/`:

```env
PORT=8080
DB_PATH=data/smart_campus.db
GIN_MODE=debug
```

Frontend environment variables:

| Variable | Default | Description |
| --- | --- | --- |
| `REACT_APP_API_URL` | `http://localhost:8080/api` | Backend API base URL used by the React app |

## Running Tests

Run all backend tests:

```bash
cd backend
go test ./...
```

If your environment blocks Go from writing to the default build cache, use a local temporary cache:

```bash
cd backend
GOCACHE=/tmp/go-build-cache go test ./...
```

Run only backend handler tests:

```bash
cd backend
go test ./handlers
```

Run frontend tests:

```bash
cd frontend
npm test
```

More detailed testing guidance is available in [TESTING_GUIDE.md](TESTING_GUIDE.md).

## API Overview

The backend API is available under `/api`.

### Authentication

```text
POST   /api/auth/register
POST   /api/auth/login
POST   /api/auth/logout
POST   /api/auth/refresh
```

### Services

```text
GET    /api/services
GET    /api/services/:id
POST   /api/services
PUT    /api/services/:id
PATCH  /api/services/:id/active
DELETE /api/services/:id
GET    /api/services/category/:category
```

`GET /api/services` supports `q`, `category`, `activeOnly`, `limit`, `offset`, `sortBy`, and `sortOrder`.

### Bookings

```text
POST   /api/bookings
GET    /api/bookings/:id
GET    /api/bookings/user/:userId
PUT    /api/bookings/:id
DELETE /api/bookings/:id
```

### Reviews

```text
POST   /api/reviews
GET    /api/reviews/service/:serviceId
GET    /api/reviews/user/:userId
GET    /api/reviews/:id
DELETE /api/reviews/:id
```

### Users

```text
GET    /api/users/:id
PUT    /api/users/:id
GET    /api/users/:id/profile
```

### Notifications

```text
GET    /api/notifications/:userId
POST   /api/notifications
PUT    /api/notifications/:id/read
```

### Approvals

```text
GET    /api/approval/staff/:staffId/pending
GET    /api/approval/staff/:staffId/all
PUT    /api/approval/bookings/:id/approve
PUT    /api/approval/bookings/:id/reject
GET    /api/approval/admin/:userId/pending
GET    /api/approval/admin/:userId/all
PUT    /api/approval/admin/:userId/bookings/:id/approve
PUT    /api/approval/admin/:userId/bookings/:id/reject
```

## Troubleshooting

Backend port already in use:

```bash
cd backend
PORT=8081 go run main.go
```

Frontend cannot reach the backend:

- Confirm the backend is running at `http://localhost:8080`.
- Confirm `frontend/.env` contains `REACT_APP_API_URL=http://localhost:8080/api`.
- Restart the frontend after changing `.env`.

Database issues:

- Confirm you are running backend commands from the `backend/` directory.
- The default database file is `backend/data/smart_campus.db`.
- For a clean local database, stop the backend, remove the SQLite file, restart the backend, and run the seed endpoint again.

Go test cache permission issues:

```bash
cd backend
GOCACHE=/tmp/go-build-cache go test ./...
```

## Project Members

- Venkata Sai Saran Jonnalagadda
- Srikar Panuganti
- Keerthi Reddy Gudibandi
- Vishnu Sai Padyala
