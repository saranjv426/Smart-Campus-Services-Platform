# Sprint 3 Frontend - Setup & Running Instructions

**Developer**: [Your Name]  
**Role**: Frontend Developer  
**Sprint**: Sprint 3 (Weeks 11-13)  
**Status**: ✅ Complete and Ready for Submission

---

## 🚀 Quick Start (30 seconds)

If you already have the project set up:

```bash
# Terminal 1: Start Backend
cd backend
go run .
# Backend runs on http://localhost:8080

# Terminal 2: Start Frontend  
cd frontend
npm start
# Frontend runs on http://localhost:3001

# Login with admin credentials
Email: admin@ufl.edu
Password: admin123
```

---

## 📋 Prerequisites

### System Requirements
- Windows 10/11, macOS, or Linux
- Git installed
- Go 1.21+ (for backend)
- Node.js 16+ and npm (for frontend)

### Verify Installation
```bash
# Check Go
go version

# Check Node & npm
node -v
npm -v

# Check Git
git --version
```

---

## 🔧 Initial Setup

### Step 1: Clone Repository
```bash
git clone https://github.com/your-org/Smart-Campus-Services-Platform.git
cd Smart-Campus-Services-Platform
git checkout sprint3v
```

### Step 2: Backend Setup
```bash
cd backend

# Download dependencies
go mod download

# Run backend
go run .
```

**Expected Output:**
```
2026/04/11 10:30:00 Server starting on port 8080
[GIN-debug] GET    /api/services
[GIN-debug] POST   /api/services
[GIN-debug] DELETE /api/services/{id}
...
[GIN] Listening and serving HTTP on :8080
```

### Step 3: Frontend Setup (New Terminal)
```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm start
```

**Expected Output:**
```
Compiled successfully!
You can now view smart-campus-services-frontend in the browser.
Local: http://localhost:3001
```

### Step 4: Database Seeding (If Services Don't Load)
```bash
# Call the seed endpoint
curl -X POST http://localhost:8080/api/seed

# Response:
# {"message":"Database seeded successfully"}
```

---

## ✅ Verification Checklist

### Backend Running?
```bash
curl -X GET http://localhost:8080/health

# Should return:
# {"status":"healthy","message":"Smart Campus Services Platform is running"}
```

### Services Available?
```bash
curl -X GET http://localhost:8080/api/services

# Should return array of 9 services with names like:
# Campus Library, Main Dining Hall, Campus Shuttle Service, etc.
```

### Frontend Running?
- Open browser to http://localhost:3001
- You should see login screen or dashboard
- No console errors in developer tools (F12)

### Database Exists?
```bash
# On Windows
ls -la backend/data/smart_campus.db

# Should show the database file exists
```

---

## 👤 Test User Credentials

### Admin Account
```
Email: admin@ufl.edu
Password: admin123
Role: admin
```

### Staff Account (Optional)
```
Email: staff@ufl.edu
Password: staff123
Role: staff
```

### Student Account (Optional)
```
Email: student@ufl.edu
Password: student123
Role: student
```

---

## 🧪 Running Tests

### Frontend Tests
```bash
cd frontend

# Run all tests
npm test

# Run specific test file
npm test -- --testPathPattern=AdminDashboard

# Run with coverage report
npm test -- --coverage

# Run in watch mode (recommended for development)
npm test -- --watch
```

### Backend Tests
```bash
cd backend

# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run with coverage
go test -cover ./...
```

---

## 📊 Featured Demo Flow (For Recording)

### 1. Show Services Loading (2 minutes)
```
1. Open browser to http://localhost:3001
2. Login with admin@ufl.edu / admin123
3. Navigate to "Services" page (or Admin Dashboard Services tab)
4. Show 9 services displaying:
   - Campus Library
   - Main Dining Hall
   - Campus Shuttle Service
   - Campus Health Center
   - Reference Desk
   - Campus Cafe
   - Parking Services
   - Counseling Services
   - Study Rooms
```

### 2. Admin Dashboard Overview (1 minute)
```
1. On Admin Dashboard (if not there, click admin nav)
2. Show statistics cards:
   - Total Bookings
   - Pending Bookings
   - Approved Bookings
   - Rejected Bookings
3. Show bookings table with all columns
4. Scroll through multiple bookings
```

### 3. Filtering Feature (2 minutes)
```
1. Click status filter "Pending" - table updates
2. Click "Approved" - table updates
3. Click "All" - shows all bookings
4. Select service category "Library" from dropdown
5. Show filtering results
6. Combine filters: Pending + Library
```

### 4. Approval Workflow (1.5 minutes)
```
1. Click "Approve" button on pending booking
2. Modal appears with booking details
3. Add notes in the notes field
4. Click "Confirm"
5. Status in table changes to green "Approved"
6. Try rejection workflow similarly
```

### 5. Service Management (1.5 minutes)
```
1. Click "Manage Services" tab
2. Show service grid with 9 cards
3. Click "Add New Service"
4. Fill form fields
5. Click "Create"
6. New service appears in grid
7. Click delete on a service
8. Confirm deletion
9. Service disappears
```

### 6. Run Tests (1 minute)
```
1. Open terminal in VS Code
2. Run: npm test -- --testPathPattern=AdminDashboard
3. Show: "PASS src/pages/AdminDashboard.test.js"
4. Show test count: "42 passed"
5. Close test runner (Ctrl+C)
```

---

## 🔍 API Endpoints Reference

### Services API
```bash
# List all services
GET http://localhost:8080/api/services

# Create service
curl -X POST http://localhost:8080/api/services \
  -H "Authorization: Bearer {token}" \
  -H "Content-Type: application/json" \
  -d '{"name":"New Service","description":"...","category":"library","location":"..."}'

# Get specific service
GET http://localhost:8080/api/services/{id}

# Delete service
DELETE http://localhost:8080/api/services/{id} \
  -H "Authorization: Bearer {token}"
```

### Booking Approval API
```bash
# Get all bookings (admin)
GET http://localhost:8080/api/approval/admin/{userId}/all \
  -H "Authorization: Bearer {token}"

# Approve booking
PUT http://localhost:8080/api/approval/admin/{userId}/bookings/{id}/approve \
  -H "Authorization: Bearer {token}" \
  -H "Content-Type: application/json" \
  -d '{"notes":"Approved - space available"}'

# Reject booking
PUT http://localhost:8080/api/approval/admin/{userId}/bookings/{id}/reject \
  -H "Authorization: Bearer {token}" \
  -H "Content-Type: application/json" \
  -d '{"reason":"Service fully booked"}'
```

---

## 🐛 Troubleshooting

### Issue: Services Not Loading
**Symptom**: Empty services list or "404 Not Found"

**Solution**:
```bash
# 1. Check backend is running
curl http://localhost:8080/health

# 2. Seed database
curl -X POST http://localhost:8080/api/seed

# 3. Verify services exist
curl http://localhost:8080/api/services

# 4. Hard refresh frontend (Ctrl+Shift+R)
```

### Issue: Port Already in Use
**Symptom**: "Something is already running on port 8080"

**Solution**:
```bash
# Find process using port
netstat -ano | findstr :8080

# Kill process (Windows - use PID from above)
taskkill /PID {PID} /F

# Or start on different port
go run . -port 8081
```

### Issue: Tests Failing
**Symptom**: "FAIL" in test output

**Solution**:
```bash
# Clear cache
npm test -- --clearCache

# Run with verbose output
npm test -- --verbose

# Update snapshots if needed
npm test -- --updateSnapshot
```

### Issue: CORS Errors
**Symptom**: "Access to XMLHttpRequest has been blocked by CORS policy"

**Solution**:
1. Verify backend is running on port 8080
2. Check frontend `package.json` has proxy:
   ```json
   "proxy": "http://localhost:8080"
   ```
3. Restart frontend: `npm start`

### Issue: Authentication Fails
**Symptom**: "Invalid credentials" or "Unauthorized"

**Solution**:
1. Verify you're using correct credentials:
   - Email: admin@ufl.edu
   - Password: admin123
2. Check browser localStorage (F12 → Application → Storage)
3. Clear browser cache: Ctrl+Shift+Delete
4. Try incognito/private window

---

## 📁 Project Structure

```
Smart-Campus-Services-Platform/
├── backend/
│   ├── main.go                 # Entry point
│   ├── seed.go                 # Database seeding (9 services)
│   ├── handlers/
│   │   ├── approval.go
│   │   ├── booking.go
│   │   ├── service.go
│   │   ├── auth.go
│   │   └── *.test.go           # Unit tests
│   ├── models/
│   │   └── models.go           # Data models
│   ├── middleware/
│   ├── config/
│   ├── data/
│   │   └── smart_campus.db     # SQLite database
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── src/
│   │   ├── pages/
│   │   │   ├── AdminDashboard.js       # SPRINT 3 ✅
│   │   │   ├── AdminDashboard.test.js  # SPRINT 3 ✅
│   │   │   └── *.js                    # Other pages
│   │   ├── styles/
│   │   │   ├── AdminDashboard.css      # SPRINT 3 ✅
│   │   │   └── *.css
│   │   ├── services/
│   │   │   └── api.js                  # Axios API client
│   │   ├── components/
│   │   ├── App.js
│   │   └── index.js
│   ├── public/
│   ├── package.json
│   ├── jest.config.js
│   ├── cypress.config.js
│   └── README.md
│
├── Sprint3.md                          # SPRINT 3 REPORT ✅
├── FRONTEND_SUBMISSION_INSTRUCTIONS.md # FOR SUBMISSION ✅
├── FRONTEND_VOICEOVER_GUIDE.md         # FOR VIDEO RECORDING ✅
├── SUBMISSION_REQUIREMENTS.md          # REQUIREMENTS CHECKLIST ✅
├── README.md                           # Main project README
└── .gitignore
```

---

## 📝 Important Environment Variables

### Backend
```
PORT=8080              # Server port
GIN_MODE=debug         # For development (debug, release)
DB_PATH=data/smart_campus.db  # Database location
```

### Frontend
```
REACT_APP_API_URL=http://localhost:8080  # Backend URL
```

---

## 🎯 Files for Submission

### Required Files
- [x] `frontend/src/pages/AdminDashboard.js` - Main component
- [x] `frontend/src/pages/AdminDashboard.test.js` - Tests
- [x] `frontend/src/styles/AdminDashboard.css` - Styles
- [x] `Sprint3.md` - Documentation
- [x] `FRONTEND_SUBMISSION_INSTRUCTIONS.md` - Instructions
- [x] `FRONTEND_VOICEOVER_GUIDE.md` - Recording guide
- [x] Video recording (MP4, 3-5 minutes)

### Optional Files
- `SUBMISSION_REQUIREMENTS.md` - This checklist
- `backend/handlers/booking_test.go` - Backend fix
- `backend/seed.go` - Backend enhancement

---

## ✅ Pre-Submission Verification

Run this checklist before submitting:

```bash
# 1. Verify branch
git branch
# Should show: * sprint3v

# 2. Verify no uncommitted changes
git status
# Should show: nothing to commit, working tree clean

# 3. Run frontend tests
cd frontend
npm test -- --testPathPattern=AdminDashboard --watchAll=false
# Should show: "PASS" and "42 passed"

# 4. Verify backend
cd backend
go test ./...
# Should show: "ok" for all packages

# 5. Verify API endpoints
curl http://localhost:8080/api/services | grep -c "Campus"
# Should return non-zero count

# 6. Check documentation
ls -la *.md
# Should show: Sprint3.md, FRONTEND_*.md, SUBMISSION_*.md
```

---

## 🎬 Recording Instructions

See `FRONTEND_VOICEOVER_GUIDE.md` for detailed narration script and timing.

Quick summary:
1. Start backend on port 8080
2. Start frontend on port 3001
3. Open OBS Studio or screen recording tool
4. Record 3-5 minute narrated presentation
5. Show all 3 features (bookings, approval, services)
6. Show tests passing
7. Export as MP4

---

## 📞 Support

### Common Issues
- See "Troubleshooting" section above
- Check browser console for errors (F12)
- Check backend terminal for error logs
- Review `Sprint3.md` for more details

### Documentation References
- [Sprint3.md](Sprint3.md) - Complete feature documentation
- [FRONTEND_SUBMISSION_INSTRUCTIONS.md](FRONTEND_SUBMISSION_INSTRUCTIONS.md) - Submission guidelines
- [FRONTEND_VOICEOVER_GUIDE.md](FRONTEND_VOICEOVER_GUIDE.md) - Recording narration guide

---

**Ready to submit? Follow the SUBMISSION_REQUIREMENTS.md checklist!** ✅
