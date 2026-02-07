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
