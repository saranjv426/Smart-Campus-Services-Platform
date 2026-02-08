# Service Details Display in Browse Services Section

## Data Structure

When you click on a service card in the "Browse Services" section, the application fetches detailed information from the backend and displays it in a comprehensive service detail view.

## Fields Displayed

### Service Information
- **Service Name** - Large heading showing the service name
- **Category Badge** - Shows the service category (e.g., "library")
- **Description** - Full detailed description of the service
- **Service Image** - High-quality image (from imageUrl)

### Contact & Location Details
- **Location** - Physical address or campus building
- **Phone** - Contact phone number
- **Email** - Contact email address
- **Hours** - Operating hours (days and times)
- **Rating** - Average star rating calculated from all reviews

### Service Metadata
All services are stored in the database with these fields:
```json
{
  "id": "UUID",
  "name": "Service Name",
  "description": "Detailed description",
  "category": "library",
  "imageUrl": "https://image-url.com/image.jpg",
  "location": "Building Name, Campus",
  "phone": "(XXX) XXX-XXXX",
  "email": "email@university.edu",
  "hours": "Operating hours",
  "rating": 4.8,
  "isActive": true,
  "createdAt": "ISO-8601 timestamp",
  "updatedAt": "ISO-8601 timestamp"
}
```

## Interactive Features on Service Detail Page

### 1. Booking System
**Location:** Below service information
- Click **"Book This Service"** button to expand booking form
- Fields to fill:
  - **Start Time** - When you want to use the service (datetime picker)
  - **End Time** - When you'll finish (datetime picker)
  - **Notes** - Optional special requests or notes
- Submit to create a booking
- Booking data is saved to database with status "pending"

### 2. Review & Rating System
**Location:** Reviews section at bottom
- **View All Reviews** - Displays all customer reviews with:
  - Reviewer name
  - Star rating (1-5 stars)
  - Comment/review text
  - Review date
- **Average Rating** - Automatically calculated from all reviews
- **Submit Review** - Form to leave your own review:
  - Select star rating
  - Write a comment
  - Submit to add to review list

## How Service Data is Retrieved

### Initial Page Load
```javascript
// Fetches service details by ID from backend
GET /api/services/{serviceId}
→ Returns: Service object with all fields

// Fetches related reviews
GET /api/reviews/service/{serviceId}
→ Returns: Array of review objects
```

### Real-time Updates
- When you submit a booking, the booking is stored in the database
- When you submit a review, it appears immediately on the page
- Average rating is recalculated when new reviews are added

## Example: Library Service Display

When you click "Main Library" service card:

### Header Section Shows:
```
[Library Image]
├─ Main Library
├─ Category: Library
└─ "University's primary library offering extensive collections..."

Meta Information:
├─ Location: Norman Hall, Gainesville Campus
├─ Phone: (352) 392-0621
├─ Email: libinfo@ufl.edu
├─ Hours: Mon-Fri: 8:00 AM - 10:00 PM...
└─ Rating: ⭐ 4.8/5
```

### Action Section:
```
[Book This Service Button]
```

### Booking Form (when expanded):
```
Start Time: [datetime picker]
End Time: [datetime picker]
Notes: [textarea]
[Confirm Booking Button]
```

### Reviews Section:
```
Reviews (N reviews)

[Leave a Review Form]
Rating: [dropdown 1-5 stars]
Comment: [textarea]
[Submit Review Button]

[Review List]
├─ John Doe ⭐⭐⭐⭐⭐
│  "Great facilities and helpful staff!"
│  2026-02-02
├─ Jane Smith ⭐⭐⭐⭐
│  "Good collection, quiet study areas"
│  2026-02-01
└─ More reviews...
```

## Customization

### Changing Service Details
To update any service information, modify the database entry:

```bash
# Backend API endpoint
PUT /api/services/{serviceId}
{
  "name": "Updated Name",
  "description": "Updated description",
  "hours": "New hours",
  "phone": "New phone",
  ...
}
```

### Adding New Services
Use the seed endpoint or API:

```bash
POST /api/services
{
  "name": "New Service",
  "description": "Description",
  "category": "library",
  "imageUrl": "https://...",
  "location": "Location",
  "phone": "Phone",
  "email": "Email",
  "hours": "Hours"
}
```

## Frontend Components

The service details are displayed using these React components:

- **ServiceDetail.js** - Main component that fetches and displays all data
- **Services.css** - Styling for service cards and detail pages
- **ServiceDetail.css** - Detailed styling for the service detail view

The component handles:
- Fetching service data
- Displaying service information
- Managing booking form state
- Handling review submissions
- Real-time updates to UI

All user interactions are processed through the API and stored in the SQLite database!
