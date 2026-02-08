# Campus Services Setup Guide

## Overview
The Smart Campus Services Platform is now populated with **8 Library Services** that users can browse, filter, book, and review.

## Available Library Services

### 1. **Main Library**
- **Location:** Norman Hall, Gainesville Campus
- **Hours:** Mon-Fri: 8:00 AM - 10:00 PM, Sat: 10:00 AM - 6:00 PM, Sun: 12:00 PM - 8:00 PM
- **Phone:** (352) 392-0621
- **Email:** libinfo@ufl.edu
- **Rating:** 4.8/5
- **Description:** University's primary library offering extensive collections of books, journals, databases, and research materials. Features study rooms, computer labs, and collaborative spaces.

### 2. **Smathers Library**
- **Location:** Smathers Library Building, West Campus
- **Hours:** Mon-Fri: 7:30 AM - 10:00 PM, Sat: 9:00 AM - 6:00 PM, Sun: 11:00 AM - 8:00 PM
- **Phone:** (352) 392-0621
- **Email:** smathers@ufl.edu
- **Rating:** 4.6/5
- **Description:** Humanities and social sciences library with specialized collections in literature, history, and social studies. Features individual study carrels and group study areas.

### 3. **Science Library**
- **Location:** Williamson Hall, Science Campus
- **Hours:** Mon-Fri: 8:00 AM - 11:00 PM, Sat: 10:00 AM - 8:00 PM, Sun: 12:00 PM - 10:00 PM
- **Phone:** (352) 392-0621
- **Email:** scilib@ufl.edu
- **Rating:** 4.7/5
- **Description:** Specialized library serving engineering, physics, chemistry, and biology students. Contains laboratory manuals, technical journals, and scientific databases.

### 4. **Law Library**
- **Location:** Spessard Holland Law Center, Downtown Campus
- **Hours:** Mon-Fri: 7:30 AM - Midnight, Sat: 9:00 AM - 9:00 PM, Sun: 10:00 AM - Midnight
- **Phone:** (352) 392-2080
- **Email:** lawlib@ufl.edu
- **Rating:** 4.9/5
- **Description:** Law library with comprehensive collections of legal research materials, case reports, and law journals. Reserved for law school and graduate students.

### 5. **Health Sciences Library**
- **Location:** Mowry Building, Health Science Campus
- **Hours:** Mon-Fri: 8:00 AM - 10:00 PM, Sat: 10:00 AM - 6:00 PM, Sun: 12:00 PM - 8:00 PM
- **Phone:** (352) 392-2080
- **Email:** hslib@ufl.edu
- **Rating:** 4.5/5
- **Description:** Medical and health sciences focused library with databases on medicine, nursing, dentistry, and pharmacy. Features case study materials and clinical references.

### 6. **Engineering Library**
- **Location:** Weil Hall, Engineering Campus
- **Hours:** Mon-Fri: 7:30 AM - 11:00 PM, Sat: 10:00 AM - 7:00 PM, Sun: 12:00 PM - 10:00 PM
- **Phone:** (352) 392-0621
- **Email:** englib@ufl.edu
- **Rating:** 4.7/5
- **Description:** Engineering-focused research library with technical manuals, design software, and engineering databases. Features computer-aided design (CAD) stations.

### 7. **Library Study Rooms**
- **Location:** Multiple Locations - All UF Libraries
- **Hours:** Available during all library operating hours
- **Phone:** (352) 392-0621
- **Email:** roomreservations@ufl.edu
- **Rating:** 4.8/5
- **Description:** Bookable private study rooms and group study spaces available across all campus libraries. Perfect for group projects and focused studying.

### 8. **Library Computer Lab**
- **Location:** Norman Hall, 3rd Floor
- **Hours:** Mon-Fri: 8:00 AM - 9:00 PM, Sat: 10:00 AM - 6:00 PM, Sun: 12:00 PM - 8:00 PM
- **Phone:** (352) 392-0621
- **Email:** complab@ufl.edu
- **Rating:** 4.6/5
- **Description:** State-of-the-art computer lab with high-performance workstations, specialized software, and high-speed internet for research and academic work.

## How to Use the Services Section

### 1. **Browse Services**
- Navigate to the **"Services"** tab in the main navigation
- View all 8 library services displayed as cards
- Each card shows the service name, description preview, image, and location

### 2. **Filter by Category**
- Click on the **"Library"** category button to filter only library services
- Use the **"All"** button to see all services (when more categories are added)
- Filters update the display in real-time

### 3. **Search Services**
- Use the search box at the top to find services by name or keyword
- Search works across service names and descriptions
- Examples: Search "Law" to find the Law Library, or "Computer" to find the Computer Lab

### 4. **View Service Details**
- Click on any service card to open the **Service Detail page**
- See complete information including:
  - Full description
  - Location and contact details
  - Operating hours
  - Current average rating
  - High-quality service image
  - Customer reviews and ratings

### 5. **Make a Booking**
- On the Service Detail page, click **"Book This Service"**
- Fill in:
  - Start Time (when you want to use the service)
  - End Time (when you'll be done)
  - Notes (optional special requests)
- Click **"Confirm Booking"**
- Track bookings in the **"My Bookings"** section

### 6. **Leave a Review**
- On the Service Detail page, scroll to the **"Reviews"** section
- Click **"Leave a Review"**
- Select a star rating (1-5 stars)
- Add a comment about your experience
- Click **"Submit Review"**
- Your review appears immediately in the reviews list
- Average rating is automatically calculated from all reviews

## Backend API Endpoints

### Get All Services
```
GET /api/services
```

### Get Services by Category
```
GET /api/services/category/library
```

### Get Service Details
```
GET /api/services/{id}
```

### Create a Booking
```
POST /api/bookings
Body: {
  "userId": "user-id",
  "serviceId": "service-id",
  "startTime": "2026-02-10T10:00:00",
  "endTime": "2026-02-10T12:00:00",
  "notes": "Need quiet study space"
}
```

### Leave a Review
```
POST /api/reviews
Body: {
  "userId": "user-id",
  "serviceId": "service-id",
  "rating": 5,
  "comment": "Excellent service and facilities!"
}
```

## Next Steps

To add more service categories:

1. **Update the seed.go file** with additional services (e.g., Dining, Transportation, Health Services, etc.)
2. **Call the seed endpoint** to add them to the database:
   ```bash
   curl -X POST http://localhost:8080/api/seed
   ```
3. **New services appear immediately** in the Services page

The system is fully scalable and ready to accommodate multiple service categories with booking and review functionality!
