# Quick Start Reference Guide

## 🚀 Getting Started in 5 Minutes

### Step 1: Backend Setup (Terminal 1)
```bash
cd backend
# Create .env file
DB_PATH=data/smart_campus.db
PORT=8080
GIN_MODE=debug

# Download dependencies
go mod download

# Run server
go run main.go
```
✅ Backend ready at http://localhost:8080

### Step 2: Frontend Setup (Terminal 2)
```bash
cd frontend
# Install dependencies
npm install

# Create .env file
REACT_APP_API_URL=http://localhost:8080/api

# Start development server
npm start
```
✅ Frontend ready at http://localhost:3000

### Step 3: Database Setup
SQLite is embedded. The database file is created automatically on backend startup.
✅ Database ready

## 📋 File Overview

### Backend Files
| File | Purpose |
|------|---------|
| `main.go` | Server setup, routes, middleware |
| `handlers/*.go` | Request handlers for each domain |
| `models/models.go` | Database models & relationships |
| `config/config.go` | Configuration constants |
| `go.mod` | Dependencies |

### Frontend Files
| Directory | Purpose |
|-----------|---------|
| `pages/` | Page components (Home, Services, etc.) |
| `components/` | Reusable components (Navbar, Footer) |
| `services/api.js` | API client for backend |
| `styles/` | CSS files for each page |
| `App.js` | Main routing and app structure |

## 🔗 Key API Endpoints

```
Authentication
POST /api/auth/register
POST /api/auth/login
POST /api/auth/logout

Services
GET /api/services
GET /api/services/:id
GET /api/services/category/:category

Bookings
POST /api/bookings
GET /api/bookings/user/:userId
DELETE /api/bookings/:id

Reviews
POST /api/reviews
GET /api/reviews/service/:serviceId

Users
GET /api/users/:id
PUT /api/users/:id
```

## 🎯 Key Features

### User Features
- 👤 Register & Login
- 🔍 Search services
- 📅 Book services
- ⭐ Leave reviews
- 👥 Manage profile

### Admin Features
- ➕ Create services
- ✏️ Edit services
- 🗑️ Delete services
- 📊 View analytics

## 🛠️ Useful Commands

### Backend
```bash
# Run server
go run main.go

# Build executable
go build -o smart-campus-backend

# Run tests
go test ./...

# Format code
go fmt ./...
```

### Frontend
```bash
# Start development
npm start

# Build for production
npm run build

# Run tests
npm test

# Clean dependencies
npm cache clean --force
```

### Database
```bash
# Open the SQLite database (optional)
sqlite3 data/smart_campus.db

# Example queries
SELECT * FROM users;
SELECT * FROM services;
SELECT * FROM bookings;
```

## 📝 Testing the Application

### 1. Create Account
- Go to http://localhost:3000
- Click "Register"
- Fill in details (role: student)

### 2. Browse Services
- Click "Services"
- Use search or filter by category
- Click "View Details"

### 3. Book a Service
- On service detail page
- Click "Book This Service"
- Select date/time
- Confirm booking

### 4. Leave a Review
- Still on service detail page
- Scroll to Reviews section
- Rate and comment
- Submit

### 5. Check Bookings
- Click "My Bookings" in navbar
- View all your bookings
- Filter by status

## 🌐 Frontend Routes

```
/ → Home page
/services → Services listing
/services/:id → Service details
/bookings → User bookings
/login → Login page
/register → Registration page
/profile/:id → User profile
```

## 🎨 Color Scheme

| Element | Color | Hex |
|---------|-------|-----|
| Primary | UF Blue | #0033a0 |
| Secondary | UF Orange | #ff6600 |
| Success | Green | #4caf50 |
| Error | Red | #f44336 |
| Warning | Orange | #ff9800 |

## 🔐 Default Test Credentials

After registration, use these for testing:
```
Email: student@uf.edu
Password: password123
Role: student
```

## ⚠️ Common Issues & Solutions

### Backend won't start
```
Error: Database connection failed
Solution: Check the backend is running and `.env` is correct
```

### Frontend shows CORS error
```
Error: CORS policy error
Solution: Make sure backend is running on :8080
```

### Bookings won't save
```
Error: 400 Bad Request
Solution: Check startTime is before endTime
```

### Can't login
```
Error: Invalid credentials
Solution: Make sure account is registered, use exact email
```

## 📞 Support Files

- **README.md** - Project overview
- **SETUP_GUIDE.md** - Detailed setup
- **ARCHITECTURE.md** - System design
- **backend/README.md** - Backend docs
- **frontend/README.md** - Frontend docs

## 🚀 Deployment Checklist

- [ ] Set `GIN_MODE=release` in backend
- [ ] Update database credentials for production
- [ ] Add JWT authentication
- [ ] Hash passwords with bcrypt
- [ ] Enable HTTPS
- [ ] Add rate limiting
- [ ] Set up monitoring
- [ ] Configure backups
- [ ] Update CORS for production domain

## 📚 Technology Stack

| Layer | Technology |
|-------|-----------|
| Frontend | React 18 + React Router + Axios |
| Backend | Go + Gin Framework |
| Database | SQLite + GORM |
| Styling | CSS3 + CSS Variables |
| Auth | Token-based (localStorage) |

## 🎓 Learning Resources

- **React**: https://react.dev
- **Go**: https://go.dev
- **Gin**: https://github.com/gin-gonic/gin
- **SQLite**: https://www.sqlite.org
- **GORM**: https://gorm.io

---

**Happy Coding! 🎉**

Created: February 1, 2026
Version: 1.0.0
