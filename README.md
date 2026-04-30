# Smart Campus Services Platform 🎓

Welcome to the Smart Campus Services Platform! This guide will help you understand what this project is all about and how you can get it running on your own computer in just a few minutes.

## Team Members:
Venkata Sai Saran Jonnalagadda - 11114995 
Srikar Panuganti - 38909216 
Keerthi Reddy Gudibandi - 13652831 
Vishnu Sai Padyala - 32712860


> [!IMPORTANT]
> **🔐 Faculty Testing Credentials**
> For grading and testing purposes, please use the following seeded accounts to bypass registration and test the platform immediately:
> - **Student Account:** `student@campus.edu` / Password: `password123`
> - **Staff Account:** `staff@campus.edu` / Password: `password123`
> - **Admin Account:** `admin@ufl.edu` / Password: `admin123`
>
> *(Note: The platform is fully secured. Random or incorrect credentials will be actively blocked by the backend API).*

---
## 📖 Project Overview & Inspiration

**What inspired us?**
University life can be overwhelming. Students constantly need to book study rooms, find out when the dining hall closes, or schedule a visit to the health center. However, this usually means jumping between five different confusing websites. We were inspired to build a single, unified platform where all campus services are accessible in one beautifully designed, easy-to-use place. 

**What is it?**
The Smart Campus Services Platform is a one-stop-shop for university students and staff. It allows students to discover campus services, book appointments, and leave reviews. It also provides administrators with a powerful dashboard to approve requests, update service information, and monitor campus activity.

---

## 🏗️ What We Did

We built this platform completely from scratch. Here are the core things we accomplished:

1. **Service Discovery:** Created a visual catalog where students can search and filter services by categories like "Dining", "Health", and "Library".
2. **Booking System:** Developed a robust calendar and time-slot system allowing students to reserve services without overlap.
3. **Admin Dashboard:** Built a control center for university staff to approve or reject student bookings, update service details, and upload photos.
4. **Review System:** Added the ability for students to leave 1-to-5 star ratings and comments based on their experiences.
5. **Real-time Notifications:** Implemented alerts so students know instantly when their booking is approved.


## 🚀 How to Launch the Website

To run this platform on your computer, you need to start both the "Backend" (which stores our data) and the "Frontend" (the visual website you click on).

**Step 1: Start the Data Server (Backend)**
1. Open your computer's terminal (Command Prompt or Mac Terminal).
2. Go into the `backend` folder by typing `cd backend`.
3. Start the server and load some test data by running: 
   ```bash
   go run main.go seed.go
   ```
   *(This will create some sample services like the Library and Dining Hall so you have something to look at!)*

**Step 2: Start the Website (Frontend)**
1. Open a second terminal window.
2. Go into the `frontend` folder by typing `cd frontend`.
3. Run these two commands:
   ```bash
   npm install
   npm start
   ```
4. The website will automatically open in your web browser at `http://localhost:3000`. You're all set!

---




---

## 📍 Where We Are Now

We have successfully completed the final phase (Sprint 4) of our project! 
The platform is fully functional, visually polished, and rigorously tested. We spent our final sprint ensuring that everything works perfectly by writing extensive automated tests for both the website interactions and the data servers.

---

## 📊 How Well We Did (Project Statistics)

We are extremely proud of the final product. Here are some quick stats that show the scale and quality of what we built:

- **100% Core Features Completed:** From searching to booking to admin approvals, everything works.
- **Over 4,000 Lines of Code:** Carefully written and organized.
- **42+ Frontend Tests:** Ensuring buttons, forms, and pages always respond correctly.
- **85%+ Test Coverage:** The vast majority of our code is automatically checked for bugs every time we make a change.
- **9 Pre-loaded Services:** The app comes ready-to-use with libraries, cafeterias, and health centers populated with high-quality images.

Thank you for visiting our project! Enjoy exploring the Smart Campus.
