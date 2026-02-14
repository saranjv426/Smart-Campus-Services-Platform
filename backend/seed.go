package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"smart-campus-services/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func seedDatabase(db *gorm.DB) error {
	// Clear existing services
	db.Where("1 = 1").Delete(&models.Service{})

	libraryServices := []models.Service{
		{
			Name:        "Main Library",
			Description: "University's primary library offering extensive collections of books, journals, databases, and research materials. Features study rooms, computer labs, and collaborative spaces.",
			Category:    "library",
			Location:    "Norman Hall, Gainesville Campus",
			Phone:       "(352) 392-0621",
			Email:       "libinfo@ufl.edu",
			Hours:       "Mon-Fri: 8:00 AM - 10:00 PM, Sat: 10:00 AM - 6:00 PM, Sun: 12:00 PM - 8:00 PM",
			ImageURL:    "https://images.pexels.com/photos/3712552/pexels-photo-3712552.jpeg?w=500&h=300",
			Rating:      4.8,
			IsActive:    true,
		},
		{
			Name:        "Smathers Library",
			Description: "Humanities and social sciences library with specialized collections in literature, history, and social studies. Features individual study carrels and group study areas.",
			Category:    "library",
			Location:    "Smathers Library Building, West Campus",
			Phone:       "(352) 392-0621",
			Email:       "smathers@ufl.edu",
			Hours:       "Mon-Fri: 7:30 AM - 10:00 PM, Sat: 9:00 AM - 6:00 PM, Sun: 11:00 AM - 8:00 PM",
			ImageURL:    "https://images.pexels.com/photos/373076/pexels-photo-373076.jpeg?w=500&h=300",
			Rating:      4.6,
			IsActive:    true,
		},
		{
			Name:        "Science Library",
			Description: "Specialized library serving engineering, physics, chemistry, and biology students. Contains laboratory manuals, technical journals, and scientific databases.",
			Category:    "library",
			Location:    "Williamson Hall, Science Campus",
			Phone:       "(352) 392-0621",
			Email:       "scilib@ufl.edu",
			Hours:       "Mon-Fri: 8:00 AM - 11:00 PM, Sat: 10:00 AM - 8:00 PM, Sun: 12:00 PM - 10:00 PM",
			ImageURL:    "https://images.pexels.com/photos/3729557/pexels-photo-3729557.jpeg?w=500&h=300",
			Rating:      4.7,
			IsActive:    true,
		},
		{
			Name:        "Law Library",
			Description: "Law library with comprehensive collections of legal research materials, case reports, and law journals. Reserved for law school and graduate students.",
			Category:    "library",
			Location:    "Spessard Holland Law Center, Downtown Campus",
			Phone:       "(352) 392-2080",
			Email:       "lawlib@ufl.edu",
			Hours:       "Mon-Fri: 7:30 AM - Midnight, Sat: 9:00 AM - 9:00 PM, Sun: 10:00 AM - Midnight",
			ImageURL:    "https://images.pexels.com/photos/3721941/pexels-photo-3721941.jpeg?w=500&h=300",
			Rating:      4.9,
			IsActive:    true,
		},
		{
			Name:        "Health Sciences Library",
			Description: "Medical and health sciences focused library with databases on medicine, nursing, dentistry, and pharmacy. Features case study materials and clinical references.",
			Category:    "library",
			Location:    "Mowry Building, Health Science Campus",
			Phone:       "(352) 392-2080",
			Email:       "hslib@ufl.edu",
			Hours:       "Mon-Fri: 8:00 AM - 10:00 PM, Sat: 10:00 AM - 6:00 PM, Sun: 12:00 PM - 8:00 PM",
			ImageURL:    "https://images.pexels.com/photos/7974/pexels-photo.jpg?w=500&h=300",
			Rating:      4.5,
			IsActive:    true,
		},
		{
			Name:        "Engineering Library",
			Description: "Engineering-focused research library with technical manuals, design software, and engineering databases. Features computer-aided design (CAD) stations.",
			Category:    "library",
			Location:    "Weil Hall, Engineering Campus",
			Phone:       "(352) 392-0621",
			Email:       "englib@ufl.edu",
			Hours:       "Mon-Fri: 7:30 AM - 11:00 PM, Sat: 10:00 AM - 7:00 PM, Sun: 12:00 PM - 10:00 PM",
			ImageURL:    "https://images.pexels.com/photos/1974652/pexels-photo-1974652.jpeg?w=500&h=300",
			Rating:      4.7,
			IsActive:    true,
		},
		{
			Name:        "Library Study Rooms",
			Description: "Bookable private study rooms and group study spaces available across all campus libraries. Perfect for group projects and focused studying.",
			Category:    "library",
			Location:    "Multiple Locations - All UF Libraries",
			Phone:       "(352) 392-0621",
			Email:       "roomreservations@ufl.edu",
			Hours:       "Available during all library operating hours",
			ImageURL:    "https://images.pexels.com/photos/3775517/pexels-photo-3775517.jpeg?w=500&h=300",
			Rating:      4.8,
			IsActive:    true,
		},
		{
			Name:        "Library Computer Lab",
			Description: "State-of-the-art computer lab with high-performance workstations, specialized software, and high-speed internet for research and academic work.",
			Category:    "library",
			Location:    "Norman Hall, 3rd Floor",
			Phone:       "(352) 392-0621",
			Email:       "complab@ufl.edu",
			Hours:       "Mon-Fri: 8:00 AM - 9:00 PM, Sat: 10:00 AM - 6:00 PM, Sun: 12:00 PM - 8:00 PM",
				ImageURL:    "https://images.pexels.com/photos/7621/pexels-photo.jpg?w=500&h=300",
			Rating:      4.6,
			IsActive:    true,
		},
	}

	// Dining Services
	diningServices := []models.Service{
		{
			Name:        "Chick-fil-A @ The Hub",
			Description: "Premium fast-casual chicken restaurant featuring hand-breaded chicken sandwiches, salads, and quality beverages. Fresh prepared daily.",
			Category:    "dining",
			Location:    "The Hub, Central Campus",
			Phone:       "(352) 392-8400",
			Email:       "dining@ufl.edu",
			Hours:       "Mon-Fri: 7:00 AM - 8:00 PM, Sat: 10:30 AM - 6:00 PM",
				ImageURL:    "https://images.pexels.com/photos/1624487/pexels-photo-1624487.jpeg?w=500&h=300",
			Rating:      4.7,
			IsActive:    true,
		},
		{
			Name:        "Panda Express @ Reitz Union",
			Description: "Asian-inspired fast-casual dining offering wok-prepared entrees, rice, noodles, and appetizers. Popular choice for lunch and dinner.",
			Category:    "dining",
			Location:    "Reitz Union, Norman Hall",
			Phone:       "(352) 392-8401",
			Email:       "dining@ufl.edu",
			Hours:       "Sun: 10:30 AM - 7:00 PM, Mon-Fri: 10:30 AM - 8:00 PM, Sat: 10:30 AM - 7:00 PM",
				ImageURL:    "https://images.pexels.com/photos/1410235/pexels-photo-1410235.jpeg?w=500&h=300",
			Rating:      4.5,
			IsActive:    true,
		},
		{
			Name:        "Subway @ Reitz Union",
			Description: "Customizable submarine sandwiches and salads made to order. Fresh ingredients and endless combinations for every taste preference.",
			Category:    "dining",
			Location:    "Reitz Union, Norman Hall",
			Phone:       "(352) 392-8402",
			Email:       "dining@ufl.edu",
			Hours:       "Daily: 11:00 AM - 10:00 PM",
				ImageURL:    "https://images.pexels.com/photos/312418/pexels-photo-312418.jpeg?w=500&h=300",
			Rating:      4.4,
			IsActive:    true,
		},
		{
			Name:        "Starbucks @ Reitz Union",
			Description: "Premium coffee shop offering handcrafted espresso drinks, pastries, breakfast items, and specialty beverages for your daily caffeine fix.",
			Category:    "dining",
			Location:    "Reitz Union, Norman Hall",
			Phone:       "(352) 392-8403",
			Email:       "dining@ufl.edu",
			Hours:       "Sun: 9:00 AM - 5:00 PM, Mon-Fri: 7:00 AM - 9:00 PM, Sat: 9:00 AM - 5:00 PM",
				ImageURL:    "https://images.pexels.com/photos/312418/pexels-photo-312418.jpeg?w=500&h=300",
			Rating:      4.6,
			IsActive:    true,
		},
		{
			Name:        "The Food Hall @ Gator Corner",
			Description: "Modern dining hall with multiple vendor stations offering diverse cuisine options. Everything from comfort food to international dishes.",
			Category:    "dining",
			Location:    "Gator Corner Dining Hall",
			Phone:       "(352) 392-8404",
			Email:       "dining@ufl.edu",
			Hours:       "Sun: 8:30 AM - 9:00 PM, Mon-Fri: 7:00 AM - 9:00 PM, Sat: 8:30 AM - 9:00 PM",
				ImageURL:    "https://images.pexels.com/photos/821365/pexels-photo-821365.jpeg?w=500&h=300",
			Rating:      4.8,
			IsActive:    true,
		},
		{
			Name:        "Baba's Pizza @ Reitz Union",
			Description: "Fresh hand-tossed pizzas prepared to order with premium toppings. Traditional Italian recipes with modern creative options available.",
			Category:    "dining",
			Location:    "Reitz Union, Norman Hall",
			Phone:       "(352) 392-8405",
			Email:       "dining@ufl.edu",
			Hours:       "Mon-Fri: 10:30 AM - 8:00 PM, Sat-Sun: Closed",
				ImageURL:    "https://images.pexels.com/photos/3407704/pexels-photo-3407704.jpeg?w=500&h=300",
			Rating:      4.6,
			IsActive:    true,
		},
		{
			Name:        "Shake Smart @ SW Rec Center",
			Description: "Health-focused smoothie and shake bar with protein options, açai bowls, and fresh fruit beverages perfect for post-workout recovery.",
			Category:    "dining",
			Location:    "Southwest Recreation Center",
			Phone:       "(352) 392-8406",
			Email:       "dining@ufl.edu",
			Hours:       "Sun: 10:00 AM - 8:00 PM, Mon-Fri: 7:00 AM - 10:00 PM, Sat: 10:00 AM - 8:00 PM",
				ImageURL:    "https://images.pexels.com/photos/1410235/pexels-photo-1410235.jpeg?w=500&h=300",
			Rating:      4.7,
			IsActive:    true,
		},
		{
			Name:        "Choolaah Indian BBQ @ United Table",
			Description: "Authentic Indian cuisine featuring tandoori preparations, curries, and Indian street food. Bold flavors and fresh spices in every dish.",
			Category:    "dining",
			Location:    "Racquet Club Dining Center",
			Phone:       "(352) 392-8407",
			Email:       "dining@ufl.edu",
			Hours:       "Mon-Fri: 11:00 AM - 8:00 PM, Sat-Sun: Closed",
				ImageURL:    "https://images.pexels.com/photos/699953/pexels-photo-699953.jpeg?w=500&h=300",
			Rating:      4.8,
			IsActive:    true,
		},
		{
			Name:        "Ace Sushi @ United Table",
			Description: "Premium sushi bar featuring fresh fish, creative rolls, and authentic Japanese preparations. Skilled sushi chefs craft each piece to perfection.",
			Category:    "dining",
			Location:    "Racquet Club Dining Center",
			Phone:       "(352) 392-8408",
			Email:       "dining@ufl.edu",
			Hours:       "Mon-Fri: 11:00 AM - 8:00 PM, Sat-Sun: Closed",
				ImageURL:    "https://images.pexels.com/photos/2097090/pexels-photo-2097090.jpeg?w=500&h=300",
			Rating:      4.7,
			IsActive:    true,
		},
		{
			Name:        "The Eatery @ Broward Hall",
			Description: "All-you-can-eat dining hall featuring diverse menu options, made-to-order stations, and fresh prepared meals throughout the day.",
			Category:    "dining",
			Location:    "Broward Hall Dining Center",
			Phone:       "(352) 392-8409",
			Email:       "dining@ufl.edu",
			Hours:       "Sun: 8:30 AM - 9:00 PM, Mon-Fri: 7:00 AM - 9:00 PM, Sat: 8:30 AM - 9:00 PM",
				ImageURL:    "https://images.pexels.com/photos/821365/pexels-photo-821365.jpeg?w=500&h=300",
			Rating:      4.6,
			IsActive:    true,
		},
		{
			Name:        "The Halal Shack @ Reitz Union",
			Description: "Authentic halal cuisine featuring Mediterranean and Middle Eastern specialties. Flavorful grilled meats and fresh vegetable preparations.",
			Category:    "dining",
			Location:    "Reitz Union, Norman Hall",
			Phone:       "(352) 392-8410",
			Email:       "dining@ufl.edu",
			Hours:       "Mon-Fri: 10:30 AM - 8:00 PM, Sat-Sun: Closed",
				ImageURL:    "https://images.pexels.com/photos/1410235/pexels-photo-1410235.jpeg?w=500&h=300",
			Rating:      4.5,
			IsActive:    true,
		},
		{
			Name:        "Starbucks @ Library West",
			Description: "Coffee shop located in the library for late-night study sessions. Quality espresso drinks and snacks available during extended hours.",
			Category:    "dining",
			Location:    "Library West, University Avenue",
			Phone:       "(352) 392-8411",
			Email:       "dining@ufl.edu",
			Hours:       "Sun: 10:00 AM - 12:00 AM, Mon-Fri: 8:00 AM - 12:00 AM, Sat: 10:00 AM - 6:00 PM",
				ImageURL:    "https://images.pexels.com/photos/3645550/pexels-photo-3645550.jpeg?w=500&h=300",
			Rating:      4.5,
			IsActive:    true,
		},
		{
			Name:        "Einstein Bros Bagels @ Shands",
			Description: "Fresh-baked bagels with cream cheese spreads and breakfast sandwiches. Quick, nutritious breakfast option for early risers.",
			Category:    "dining",
			Location:    "Shands Sun Terrace, Medical Campus",
			Phone:       "(352) 392-8412",
			Email:       "dining@ufl.edu",
			Hours:       "Mon-Sat: 7:00 AM - 3:00 PM, Sun: Closed",
				ImageURL:    "https://images.pexels.com/photos/1092730/pexels-photo-1092730.jpeg?w=500&h=300",
			Rating:      4.4,
			IsActive:    true,
		},
		{
			Name:        "Melt Lab @ Broward",
			Description: "Late-night dessert spot featuring gourmet ice cream and sweet treats. Perfect for study breaks and late-night cravings.",
			Category:    "dining",
			Location:    "Broward Hall",
			Phone:       "(352) 392-8413",
			Email:       "dining@ufl.edu",
			Hours:       "Daily: 9:00 PM - 12:00 AM",
				ImageURL:    "https://images.pexels.com/photos/3407704/pexels-photo-3407704.jpeg?w=500&h=300",
			Rating:      4.8,
			IsActive:    true,
		},
		{
			Name:        "The Market @ Rawlings",
			Description: "Convenience market with grab-and-go snacks, beverages, prepared meals, and late-night food options. Open until 1:00 AM for midnight cravings.",
			Category:    "dining",
			Location:    "Rawlings Hall",
			Phone:       "(352) 392-8414",
			Email:       "dining@ufl.edu",
			Hours:       "Sun: 10:00 AM - 1:00 AM, Mon-Fri: 9:00 AM - 1:00 AM, Sat: 10:00 AM - 1:00 AM",
				ImageURL:    "https://images.pexels.com/photos/1410235/pexels-photo-1410235.jpeg?w=500&h=300",
			Rating:      4.3,
			IsActive:    true,
		},
	}

	// Transportation Services
	transportationServices := []models.Service{
		{
			Name:        "RTS Route 5 - Main Campus",
			Description: "Regional Transit System bus route serving central campus area. Connects dormitories, academic buildings, and libraries. Real-time tracking available via transit app. Service every 10-15 minutes during peak hours.",
			Category:    "transportation",
			Location:    "Campus-wide (Primary: Norman Hall Transit Hub)",
			Phone:       "(352) 392-8200",
			Email:       "transit@ufl.edu",
			Hours:       "Mon-Fri: 6:30 AM - 12:00 AM, Sat-Sun: 7:00 AM - 11:00 PM",
			ImageURL:    "https://images.pexels.com/photos/2574261/pexels-photo-2574261.jpeg?w=500&h=300",
			Rating:      4.4,
			IsActive:    true,
		},
		{
			Name:        "RTS Route 9 - Medical Campus",
			Description: "Direct route connecting main campus to Shands Medical Center and health facilities. Ideal for students and staff accessing health services. Modern articulated buses with air conditioning and real-time arrival information.",
			Category:    "transportation",
			Location:    "Main Campus to Medical Campus",
			Phone:       "(352) 392-8201",
			Email:       "transit@ufl.edu",
			Hours:       "Mon-Fri: 7:00 AM - 11:00 PM, Sat-Sun: 8:00 AM - 10:00 PM",
			ImageURL:    "https://images.pexels.com/photos/1761279/pexels-photo-1761279.jpeg?w=500&h=300",
			Rating:      4.5,
			IsActive:    true,
		},
		{
			Name:        "RTS Route 12 - Academic Core",
			Description: "Circular route through engineering, science, and business districts. Connects all major academic buildings. Stops at key research facilities and computer labs. Frequency: Every 8-12 minutes during day.",
			Category:    "transportation",
			Location:    "Engineering Campus Loop",
			Phone:       "(352) 392-8202",
			Email:       "transit@ufl.edu",
			Hours:       "Mon-Fri: 7:30 AM - 5:00 PM, Sat-Sun: 10:00 AM - 4:00 PM",
			ImageURL:    "https://images.pexels.com/photos/2574261/pexels-photo-2574261.jpeg?w=500&h=300",
			Rating:      4.6,
			IsActive:    true,
		},
		{
			Name:        "Campus Connector Express",
			Description: "Dedicated internal campus shuttle system designed for quick movement between major campus zones. Serves central hub, west campus, and academic districts. Free for all UF students and staff. Estimated travel times: 5-15 minutes per segment.",
			Category:    "transportation",
			Location:    "Multiple Campus Zones - Central Hub: Reitz Union",
			Phone:       "(352) 392-8250",
			Email:       "connector@ufl.edu",
			Hours:       "Mon-Fri: 7:00 AM - 9:00 PM, Sat-Sun: 9:00 AM - 6:00 PM",
			ImageURL:    "https://images.pexels.com/photos/170350/pexels-photo-170350.jpeg?w=500&h=300",
			Rating:      4.7,
			IsActive:    true,
		},
		{
			Name:        "Campus Connector Route A - West Loop",
			Description: "West campus route connecting athletic facilities, residential halls, and western academic buildings. Pickup/drop-off points: SW Recreation Center, UF Housing, Engineering Lab Complex. Service interval: 15-20 minutes.",
			Category:    "transportation",
			Location:    "West Campus Loop",
			Phone:       "(352) 392-8251",
			Email:       "connector@ufl.edu",
			Hours:       "Daily: 7:00 AM - 9:00 PM",
			ImageURL:    "https://images.pexels.com/photos/1440918/pexels-photo-1440918.jpeg?w=500&h=300",
			Rating:      4.6,
			IsActive:    true,
		},
		{
			Name:        "Lyft E-Bikes - Primary Hubs",
			Description: "Micromobility network with e-bike docking stations across campus. Prime pickup locations include Reitz Union, Norman Hall, SW Recreation Center, Engineering Hub. Usage: $1 unlock + $0.25/minute. Download Lyft app for real-time availability.",
			Category:    "transportation",
			Location:    "15+ Docking Stations Campus-wide",
			Phone:       "(352) 392-8300",
			Email:       "ekugly@ufl.edu",
			Hours:       "24/7 Availability",
			ImageURL:    "https://images.pexels.com/photos/2382898/pexels-photo-2382898.jpeg?w=500&h=300",
			Rating:      4.5,
			IsActive:    true,
		},
		{
			Name:        "Lyft E-Bikes - East Campus Hub",
			Description: "High-demand e-bike docking hub on east campus near libraries and student centers. Frequently restocked during peak hours. Popular for short trips between academic buildings. Maintenance: Daily inspections ensure bike safety.",
			Category:    "transportation",
			Location:    "East Campus Hub - Library West Complex",
			Phone:       "(352) 392-8301",
			Email:       "ekugly@ufl.edu",
			Hours:       "24/7 Availability",
			ImageURL:    "https://images.pexels.com/photos/2382898/pexels-photo-2382898.jpeg?w=500&h=300",
			Rating:      4.4,
			IsActive:    true,
		},
	}

	// Health Services
	healthServices := []models.Service{
		{
			Name:        "UF Student Health Services (SHS)",
			Description: "Primary health care facility for UF students offering comprehensive medical services. Services include: General physician visits, preventive care, urgent care, vaccinations, allergy testing, STI testing, mental health counseling. Insurance: UF Health Plan accepted. Visits typically $5 copay for students. Open to all enrolled students.",
			Category:    "health",
			Location:    "Shands Hospital - South Campus, Room 103",
			Phone:       "(352) 392-1161",
			Email:       "shs@ufl.edu",
			Hours:       "Mon-Fri: 8:00 AM - 5:00 PM (Urgent: until 9 PM Fri)",
			ImageURL:    "https://images.pexels.com/photos/3807517/pexels-photo-3807517.jpeg?w=500&h=300",
			Rating:      4.3,
			IsActive:    true,
		},
		{
			Name:        "UF Faculty & Staff Health Program",
			Description: "Comprehensive health services for UF faculty and staff members. Coverage includes: Annual physicals, occupational health services, wellness programs, laboratory services, medical advice line 24/7. Insurance: Varies by plan. Includes Blue Cross Blue Shield, Cigna, and UF Select.\nEligibility: Benefits-eligible faculty/staff. Retirees with active coverage welcome.",
			Category:    "health",
			Location:    "Shands Medical Plaza - 200 Tower Drive, Suite 200",
			Phone:       "(352) 273-7780",
			Email:       "facultyhealth@ufl.edu",
			Hours:       "Mon-Fri: 7:30 AM - 4:30 PM, Sat: 8:00 AM - 12:00 PM",
			ImageURL:    "https://images.pexels.com/photos/3957987/pexels-photo-3957987.jpeg?w=500&h=300",
			Rating:      4.4,
			IsActive:    true,
		},
		{
			Name:        "UF Health Urgent Care Center",
			Description: "Walk-in urgent care facility for non-life-threatening injuries and illnesses. Services: Minor injuries, cold/flu symptoms, urinary tract infections, sprains, cuts, and burns. No appointment needed. Average wait: 20-45 minutes. Insurance: Most plans accepted including UF Health Plan, Blue Cross, Cigna, and others.",
			Category:    "health",
			Location:    "Shands Medical Campus - Building 100",
			Phone:       "(352) 265-0111",
			Email:       "urgentcare@ufl.edu",
			Hours:       "Mon-Fri: 11:00 AM - 8:00 PM, Sat-Sun: 10:00 AM - 6:00 PM",
			ImageURL:    "https://images.pexels.com/photos/7088543/pexels-photo-7088543.jpeg?w=500&h=300",
			Rating:      4.2,
			IsActive:    true,
		},
		{
			Name:        "UF Health Mental Health Services",
			Description: "Mental health and counseling services for students. Services: Individual therapy, group counseling, crisis intervention, psychiatric medication management, peer support, stress management workshops. Insurance: UF Health Plan covers 8 free sessions/year. External insurance also accepted. Telehealth available.",
			Category:    "health",
			Location:    "Shands Medical Center - Behavioral Health Wing",
			Phone:       "(352) 392-1575",
			Email:       "mentalhealth@ufl.edu",
			Hours:       "Mon-Fri: 9:00 AM - 5:00 PM, Urgent: 24/7 Crisis Line",
			ImageURL:    "https://images.pexels.com/photos/7088465/pexels-photo-7088465.jpeg?w=500&h=300",
			Rating:      4.6,
			IsActive:    true,
		},
		{
			Name:        "UF Health Dental Services",
			Description: "Comprehensive dental care for students and faculty. Services: Cleanings, fillings, extractions, orthodontics, emergency dental treatment. Note: Limited availability for students; may require referral. Insurance: UF dental plans accepted. Faculty/staff: Most plans covered. Cost varies by service.",
			Category:    "health",
			Location:    "College of Dentistry - South Campus",
			Phone:       "(352) 273-5800",
			Email:       "dentistry@ufl.edu",
			Hours:       "Mon-Thu: 8:00 AM - 5:00 PM, Fri: 8:00 AM - 1:00 PM",
			ImageURL:    "https://images.pexels.com/photos/3779035/pexels-photo-3779035.jpeg?w=500&h=300",
			Rating:      4.5,
			IsActive:    true,
		},
		{
			Name:        "UF Health Pharmacy",
			Description: "On-campus pharmacy serving UF community. Services: Prescription fill (UF student rate $5 generic/$10 name-brand), medication counseling, preferred medication lists consultation, insurance verification. Insurance: All UF plans accepted. External insurance processed. 24-hour emergency pharmacy available.",
			Category:    "health",
			Location:    "Shands Hospital - 1st Floor, Room 150",
			Phone:       "(352) 265-0111 ext 4",
			Email:       "pharmacy@ufl.edu",
			Hours:       "Mon-Fri: 8:00 AM - 6:00 PM, Sat: 9:00 AM - 1:00 PM",
			ImageURL:    "https://images.pexels.com/photos/4021085/pexels-photo-4021085.jpeg?w=500&h=300",
			Rating:      4.4,
			IsActive:    true,
		},
		{
			Name:        "UF Health Specialty Services - Orthopedics",
			Description: "Specialized orthopedic care for sports injuries, joint problems, and bone/muscle issues. Services: Orthopedic evaluation, X-ray available, physical therapy referral, surgical consultation. Insurance: Student Health Plan (referral required), Faculty/Staff plans (Blue Cross, Cigna). Estimated visit: $25-50 copay.",
			Category:    "health",
			Location:    "Shands Medical Plaza - Sports Medicine Wing",
			Phone:       "(352) 265-5911",
			Email:       "orthopedics@ufl.edu",
			Hours:       "Mon-Fri: 8:00 AM - 5:00 PM (Urgent: Emergency Room 24/7)",
			ImageURL:    "https://images.pexels.com/photos/7088465/pexels-photo-7088465.jpeg?w=500&h=300",
			Rating:      4.7,
			IsActive:    true,
		},
		{
			Name:        "UF Health - Emergency Room",
			Description: "24-hour emergency department for life-threatening conditions and severe injuries. Services: Trauma care, emergency surgery, intensive care. Insurance: All insurance plans accepted. Uninsured patients: UF financial assistance available. Triaged by severity - critical cases seen first.",
			Category:    "health",
			Location:    "Shands Hospital - Main Entrance",
			Phone:       "911 or (352) 265-0111",
			Email:       "er@ufl.edu",
			Hours:       "Open 24/7",
			ImageURL:    "https://images.pexels.com/photos-7088543/pexels-photo.jpg?w=500&h=300",
			Rating:      4.5,
			IsActive:    true,
		},
	}

	// Create services first and store IDs
	var libraryServiceID string
	var diningServiceID string
	var transportationServiceID string
	var healthServiceID string
	for _, service := range libraryServices {
		if err := db.Create(&service).Error; err != nil {
			log.Printf("Failed to create service %s: %v", service.Name, err)
		} else {
			fmt.Printf("✓ Created service: %s\n", service.Name)
			if service.Name == "Main Library" {
				libraryServiceID = service.ID
			}
		}
	}

	for _, service := range diningServices {
		if err := db.Create(&service).Error; err != nil {
			log.Printf("Failed to create service %s: %v", service.Name, err)
		} else {
			fmt.Printf("✓ Created service: %s\n", service.Name)
			if service.Name == "Chick-fil-A @ The Hub" {
				diningServiceID = service.ID
			}
		}
	}

	fmt.Println("\n🚍 Adding Transportation Services...")
	for _, service := range transportationServices {
		if err := db.Create(&service).Error; err != nil {
			log.Printf("Failed to create service %s: %v", service.Name, err)
		} else {
			fmt.Printf("✓ Created service: %s\n", service.Name)
			if service.Name == "RTS Route 5 - Main Campus" {
				transportationServiceID = service.ID
			}
		}
	}

	fmt.Println("\n🏥 Adding Health Services...")
	for _, service := range healthServices {
		if err := db.Create(&service).Error; err != nil {
			log.Printf("Failed to create service %s: %v", service.Name, err)
		} else {
			fmt.Printf("✓ Created service: %s\n", service.Name)
			if service.Name == "UF Student Health Services (SHS)" {
				healthServiceID = service.ID
			}
		}
	}

	// Create staff accounts for each service
	staffAccounts := []models.User{
		{
			Email:     "library-admin@ufl.edu",
			Password:  "library123",
			FirstName: "Library",
			LastName:  "Admin",
			Phone:     "(352) 392-0621",
			Role:      "staff",
			ServiceID: libraryServiceID,
		},
		{
			Email:     "library-assistant@ufl.edu",
			Password:  "library123",
			FirstName: "Sarah",
			LastName:  "Mitchell",
			Phone:     "(352) 392-0622",
			Role:      "staff",
			ServiceID: libraryServiceID,
		},
		{
			Email:     "dining-admin@ufl.edu",
			Password:  "dining123",
			FirstName: "Dining",
			LastName:  "Manager",
			Phone:     "(352) 392-8400",
			Role:      "staff",
			ServiceID: diningServiceID,
		},
		{
			Email:     "dining-staff@ufl.edu",
			Password:  "dining123",
			FirstName: "John",
			LastName:  "Garcia",
			Phone:     "(352) 392-8401",
			Role:      "staff",
			ServiceID: diningServiceID,
		},
		{
			Email:     "transit-admin@ufl.edu",
			Password:  "transit123",
			FirstName: "Transit",
			LastName:  "Coordinator",
			Phone:     "(352) 392-8200",
			Role:      "staff",
			ServiceID: transportationServiceID,
		},
		{
			Email:     "transit-staff@ufl.edu",
			Password:  "transit123",
			FirstName: "Maria",
			LastName:  "Lopez",
			Phone:     "(352) 392-8201",
			Role:      "staff",
			ServiceID: transportationServiceID,
		},
		{
			Email:     "health-admin@ufl.edu",
			Password:  "health123",
			FirstName: "Health",
			LastName:  "Administrator",
			Phone:     "(352) 392-1161",
			Role:      "staff",
			ServiceID: healthServiceID,
		},
		{
			Email:     "health-nurse@ufl.edu",
			Password:  "health123",
			FirstName: "Emily",
			LastName:  "Johnson",
			Phone:     "(352) 392-1162",
			Role:      "staff",
			ServiceID: healthServiceID,
		},
	}

	fmt.Println("\n🔐 Creating staff accounts...")
	for _, staff := range staffAccounts {
		// Check if staff already exists
		var existingStaff models.User
		if err := db.Where("email = ?", staff.Email).First(&existingStaff).Error; err == nil {
			fmt.Printf("✓ Staff already exists: %s\n", staff.Email)
			continue
		}

		if err := db.Create(&staff).Error; err != nil {
			log.Printf("Failed to create staff %s: %v", staff.Email, err)
		} else {
			fmt.Printf("✓ Created staff account: %s (%s %s) - ServiceID: %s\n",
				staff.Email, staff.FirstName, staff.LastName, staff.ServiceID)
		}
	}

	fmt.Println("\n✓ Database seeding completed!")
	return nil
}

func seedDatabaseOnly() {
	dbPath := databasePath()
	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		log.Fatal("Failed to create database directory:", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := seedDatabase(db); err != nil {
		log.Fatal("Failed to seed database:", err)
	}
}
