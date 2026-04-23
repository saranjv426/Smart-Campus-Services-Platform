# Sprint 3 Frontend - VoiceOver Recording Guide

**Duration**: 3-5 minutes  
**Format**: Screen recording with narration (MP4)  
**Audience**: Instructors and peer reviewers  

---

## 🎬 Recording Schedule & Talking Points

### SEGMENT 1: Introduction (30 seconds)

**Time**: 0:00 - 0:30  
**Visual**: Show desktop → Open browser to admin dashboard

```
"Hello, I'm [Your Name], student ID [ID]. 

I completed the frontend work for Sprint 3 of the Smart Campus Services 
Platform. Our team tackled three major features for the admin dashboard.

In this video, I'll demonstrate:
- The comprehensive bookings overview dashboard
- Admin approval and rejection workflows
- Service management interface

These features give campus administrators powerful tools to manage 
bookings and services across the platform. Let me show you what I built."
```

**Action Items**:
- [ ] Clear home screen
- [ ] Open chrome/browser
- [ ] Navigate to http://localhost:3001
- [ ] Login as admin@ufl.edu / admin123
- [ ] Show you're on admin dashboard

---

### SEGMENT 2: Statistics & Overview (45 seconds)

**Time**: 0:30 - 1:15  
**Visual**: Show statistics cards, scroll through bookings table

```
"On the admin dashboard, you can see four statistics cards at the top:
- Total bookings across the entire platform
- Pending bookings awaiting admin action  
- Approved bookings
- Rejected bookings

These numbers update in real-time as I approve or reject bookings. Below 
the statistics is a comprehensive table showing all student bookings.

The table includes:
- Service name - which campus service was booked
- Student name and email - who made the booking
- Start and end times - when the booking is scheduled
- Status - whether it's pending, approved, or rejected
- Created date - when the booking was made
- Action buttons - approve or reject for pending bookings

This table is sortable and responsive, adapting to different screen sizes."
```

**Action Items**:
- [ ] Point to each statistics card
- [ ] Read the numbers out loud
- [ ] Scroll through the bookings table
- [ ] Show all 8 columns visible
- [ ] Pause on different bookings
- [ ] Highlight pending bookings (they have action buttons)

---

### SEGMENT 3: Filtering - Status (45 seconds)

**Time**: 1:15 - 2:00  
**Visual**: Show status filter buttons, click through each one

```
"One key feature is filtering. I added status filters - you can view 
bookings by their approval status.

Here I have four filter buttons: Pending, Approved, Rejected, and All.

Let me click 'Pending' - notice the table updates instantly to show only 
pending bookings. These are the bookings waiting for admin action.

Now I'll click 'Approved' - the table refreshes to show only approved 
bookings. These are bookings we've already approved.

Clicking 'Rejected' shows rejected bookings.

And 'All' shows every booking regardless of status.

This filtering happens instantly without a page reload. Users get immediate 
feedback, which creates a smooth, professional experience."
```

**Action Items**:
- [ ] Click "Pending" button - pause 1 second to show table update
- [ ] Say the count: "I see X pending bookings"
- [ ] Click "Approved" button - pause 1 second
- [ ] Say the count: "I see X approved bookings"
- [ ] Click "Rejected" button - pause 1 second
- [ ] Click "All" button - pause 1 second

---

### SEGMENT 4: Filtering - Service Category (45 seconds)

**Time**: 2:00 - 2:45  
**Visual**: Show service category dropdown, select different categories

```
"Beyond status filtering, I also added service category filtering. 

Below the status buttons is a dropdown menu showing service categories: 
All, Library, Dining, Transportation, and Healthcare.

Let me select 'Library' - the table updates to show only bookings for 
library services. I can see bookings for 'Campus Library' and 'Study Rooms'.

If I select 'Dining', the table refreshes to show only dining service 
bookings.

Let me try 'Transportation' - now it shows shuttle and parking service 
bookings.

'Healthcare' shows health center and counseling service bookings.

These filters work together too. I can filter by 'Pending' status AND 
'Library' category at the same time to see pending library bookings 
specifically. This powerful combination gives admins fine-grained control 
over managing the platform."
```

**Action Items**:
- [ ] Click dropdown - show all categories
- [ ] Select "Library" - pause for table update
- [ ] Say what you see: "2 library bookings"
- [ ] Select "Dining" - pause for update
- [ ] Select "Transportation" - pause for update
- [ ] Select "Healthcare" - pause for update
- [ ] Select "All" to reset
- [ ] Show combined filtering: keep "Pending" + select "Library"
- [ ] Explain: "Now I see only pending library bookings"

---

### SEGMENT 5: Approval Workflow (60 seconds)

**Time**: 2:45 - 3:45  
**Visual**: Click approve button, show modal, add notes, confirm

```
"Now let me show the approval workflow. As an admin, I can approve or 
reject any booking.

I'll click the green 'Approve' button on this pending booking for the 
Campus Library.

A modal dialog appears with the booking details and a notes field. The 
admin can add optional notes explaining the approval decision. This creates 
an audit trail.

I'll type some notes: 'Approved - Library space available.'

Now I click confirm, and - notice the status in the table immediately 
changes from 'Pending' to 'Approved' in green. The modal closes and the 
table refreshes without a page reload.

The backend also sends a notification to the student letting them know 
their booking was approved.

Let me show rejection too - I'll click the red 'Reject' button on another 
pending booking.

A similar modal appears. I can add a rejection reason: 'Service fully 
booked for requested time.'

Click confirm - the booking status changes to 'Rejected' in red.

This workflow is intuitive and gives admins control while maintaining 
data integrity through confirmations."
```

**Action Items**:
- [ ] Find a pending booking (with Pending status)
- [ ] Click green "Approve" button
- [ ] Wait for modal to appear
- [ ] Type notes in the textarea
- [ ] Click "Confirm" or "Submit" button
- [ ] Watch table update - status should turn green "Approved"
- [ ] Take a screenshot or pause moment
- [ ] Find another pending booking
- [ ] Click red "Reject" button
- [ ] Type rejection reason
- [ ] Click confirm
- [ ] Watch status turn red "Rejected"

---

### SEGMENT 6: Service Management Tab (60 seconds)

**Time**: 3:45 - 4:45  
**Visual**: Click services tab, show grid, create service, delete service

```
"Let me show the third feature - Service Management.

At the top of the dashboard, I have two tabs: 'Bookings Overview' and 
'Manage Services'. Let me click the 'Manage Services' tab.

Now I see a grid displaying all campus services. Each service card shows 
the service name, category, and action buttons.

There's also an 'Add New Service' button. Let me click it to create a 
new service.

A modal form appears with fields for:
- Service name
- Description
- Category dropdown
- Location
- Phone
- Email
- Operating hours
- Image URL

Let me fill in the details for a new service... I'll name it 'Night Study 
Session', category is 'Library', I'll add a description and location.

Click 'Create Service' - and the service saves successfully. The modal 
closes and a new card appears in the grid.

To delete a service, I click the delete button on any service card. A 
confirmation modal appears asking 'Are you sure?' - this prevents 
accidental deletions.

I confirm the deletion, and the service card disappears from the grid.

This complete service management interface gives admins control to expand 
or modify the platform's service offerings in real-time."
```

**Action Items**:
- [ ] Click "Manage Services" tab (or 🏢 Services tab)
- [ ] Wait for grid to load
- [ ] Point out 9 service cards in the grid
- [ ] Click "Add New Service" or "➕ Create Service" button
- [ ] Modal appears - fill in form:
  - Name: "Night Study Session"
  - Description: "Late night study support"
  - Category: Select "library"
  - Location: "Main Library"
  - Phone: "555-9999"
  - Email: "study@campus.edu"
  - Hours: "Mon-Fri 6:00 PM - 1:00 AM"
- [ ] Click "Create" or "Submit"
- [ ] Show new card added to grid
- [ ] Count services: "Now we have 10 services"
- [ ] Click delete button on the new service
- [ ] Confirm deletion in modal
- [ ] Service disappears from grid

---

### SEGMENT 7: Responsive Design & Testing (45 seconds)

**Time**: 4:45 - 5:30  
**Visual**: Resize browser, show responsive layout, show tests

```
"One important aspect is responsive design. The dashboard works on all 
device sizes.

When I resize the browser to a mobile width - notice how the table adapts. 
On very narrow screens, the grid columns reorganize for readability.

Back to desktop size - the full layout displays.

Behind the scenes, I wrote comprehensive unit tests to ensure everything 
works correctly. Let me show you the test results.

Running the test suite... I have 42 unit tests covering:
- Tab navigation and switching between tabs
- Service creation with form validation
- Service deletion with confirmation
- Booking filtering by status and category
- Approval and rejection workflows
- Responsive design verification

All 42 tests pass, showing the admin dashboard is reliable and 
production-ready.

The test coverage for the AdminDashboard component is 85%, meaning 
most of the code is tested, significantly reducing the risk of bugs 
in production."
```

**Action Items**:
- [ ] Open browser devtools (F12)
- [ ] Click responsive design mode
- [ ] Select mobile device (iPhone 12)
- [ ] Show how table/grid adapts
- [ ] Scroll through mobile layout
- [ ] Return to desktop size
- [ ] Close devtools
- [ ] Open terminal in VS Code
- [ ] Run: `npm test -- --testPathPattern=AdminDashboard`
- [ ] Show test output with "PASS" and test names
- [ ] Show "42+ tests passed"
- [ ] Show coverage percentage

---

### SEGMENT 8: Closing Remarks (30 seconds)

**Time**: 5:30 - 6:00  
**Visual**: Back to dashboard, show overall view

```
"To summarize, in Sprint 3 I implemented:

✓ Issue #62: Admin Dashboard with real-time statistics and advanced filtering
✓ Issue #63: Approval and rejection workflows with audit notes
✓ Issue #64: Service management with create and delete operations

Every feature is:
- Fully functional and tested
- Responsive across all devices
- Integrated with the backend API
- Protected with proper validations
- Covered by comprehensive unit tests

The admin dashboard is now a professional, production-ready interface 
for managing the Smart Campus Services platform. Thank you for watching!"
```

**Action Items**:
- [ ] Show the full admin dashboard
- [ ] Do a final scroll through
- [ ] Point to key features visible
- [ ] Thank the viewer

---

## 🎙️ Voiceover Tips & Techniques

### Preparation
- [ ] Write out the exact wording you want to say
- [ ] Practice the narration 2-3 times before recording
- [ ] Have water nearby
- [ ] Minimize background noise
- [ ] Test microphone volume levels

### Recording Technique
- [ ] Speak clearly and at a moderate pace
- [ ] Pause for 1-2 seconds after each action for visibility
- [ ] Use natural inflection and enthusiasm
- [ ] Emphasize key features: "Notice how the table updates instantly..."
- [ ] Use specific numbers when possible: "9 services are now available"
- [ ] Explain the 'why' not just the 'what'

### Pacing Guide
```
Segment 1: 0:30 - 0:45 (Intro) - slow, clear
Segment 2: 0:45 - 1:30 (Stats) - moderate pace
Segment 3: 1:30 - 2:15 (Status filter) - show, then explain
Segment 4: 2:15 - 3:00 (Category filter) - show, then explain
Segment 5: 3:00 - 4:00 (Approval) - slower, demonstrate carefully
Segment 6: 4:00 - 5:00 (Services) - slower, demonstrate each step
Segment 7: 5:00 - 5:45 (Testing) - moderate pace
Segment 8: 5:45 - 6:00 (Closing) - clear and confident
```

### Audio Quality Checklist
- [ ] Microphone is 3-6 inches from mouth
- [ ] No clipping (audio not too loud)
- [ ] No background hum or noise
- [ ] Consistent volume throughout
- [ ] Audible but not overwhelming
- [ ] No background music or distracting sounds

---

## 📹 Recording Software Options

### Option 1: OBS Studio (Free, Professional)
```
1. Download: https://obsproject.com/
2. Add Source: Window Capture (Browser)
3. Add Source: Audio Input Capture (Microphone)
4. Set output: 1080p, 30fps, 3000 kbps
5. Start Recording
6. Click through UI as you narrate
7. Stop and save
```

### Option 2: ScreenFlow (Mac, $99)
```
1. Open ScreenFlow
2. Select screen to record (browser)
3. Enable microphone input
4. Start recording
5. Perform actions while speaking
6. Export as MP4
```

### Option 3: Camtasia (Windows/Mac, $99)
```
1. Open Camtasia
2. Click "Record screen"
3. Enable microphone
4. Record audio narration while acting
5. Export as MP4
```

### Option 4: Built-in (Windows 10/11, Free)
```
1. Open Xbox Game Bar (Win + G)
2. Click "Start Recording"
3. Your narration and actions are recorded
4. Find the video in Videos > Captures
```

---

## ✅ Final Checklist Before Uploading

- [ ] Audio is clear and easy to understand
- [ ] Video resolution is at least 1080p
- [ ] Total duration is 3-5 minutes
- [ ] All three issues are demonstrated
- [ ] Tests are shown passing
- [ ] No sensitive information visible
- [ ] No spelling errors in text shown
- [ ] Microphone captured clearly
- [ ] Background noise is minimal
- [ ] Video file is saved as MP4
- [ ] File size is reasonable (< 500MB)
- [ ] You've watched the entire video once

---

## 🚨 Common Video Issues & Solutions

| Issue | Solution |
|-------|----------|
| Audio too quiet | Check recording levels, re-record closer to mic |
| Audio too loud | Reduce microphone in volume settings |
| Video choppy | Close background applications, use lower resolution |
| Microphone not working | Check Windows sound settings, test before recording |
| Can't find recording | Check Videos/Captures folder or OBS output folder |
| Video won't upload | Convert to H.264 codec if needed |
| Background noise | Record in quiet room, use noise gate in OBS |
| Speaking too fast | Practice narration slower, focus on clarity |

---

## 📊 Recording Quality Checklist

**Technical Quality** (60% of evaluation)
- [ ] Video is 1920x1080p minimum
- [ ] Frame rate is consistent (30fps+)
- [ ] Audio is clear and understandable
- [ ] No dropped frames or stuttering
- [ ] Color is accurate and visible

**Content Quality** (40% of evaluation)
- [ ] All three issues demonstrated
- [ ] Tests are shown passing
- [ ] Features work correctly
- [ ] Narration is clear and informative
- [ ] Pacing is appropriate
- [ ] Professional presentation

---

## 🎬 Final Recording Workflow

1. **Prepare** (5 min before recording)
   - Start backend server
   - Start frontend server
   - Login to admin account
   - Close unnecessary apps
   - Test microphone
   - Open recording software

2. **Record** (5-6 minutes)
   - Start recording
   - Narrate introduction (0:30)
   - Narrate statistics (0:45)
   - Narrate status filtering (0:45)
   - Narrate category filtering (0:45)
   - Narrate approval workflow (1:00)
   - Narrate service management (1:00)
   - Narrate testing (0:45)
   - Narrate closing (0:30)

3. **Post-Process** (2 min)
   - Export video as MP4
   - Check file is playable
   - Check audio/video sync
   - Save with clear filename
   - Note the file size

4. **Upload** (1 min)
   - Upload to GitHub releases or discussions
   - Generate link
   - Add link to submission

---

**You're ready to record! 🎉 Good luck!**
