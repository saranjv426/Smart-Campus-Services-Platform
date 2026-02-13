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
			ImageURL:    "https://images.unsplash.com/photo-1507842675196-a957f0e16486?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1524995997946-a1c2e315a42f?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1532012197267-da84d127e765?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1521587765099-8244a2520746?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1576091160550-112173f31c77?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1454165804606-c3d57bc86b40?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1552664730-d307ca884978?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1593642632823-8f785ba67e45?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1585521537066-0f3cf13b911b?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1552521516-c4c45915a724?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1565299585323-38d6b0865b47?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1549855255-e0635a0b934e?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1568901346375-23c9450c58cd?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1585521537066-0f3cf13b911b?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1590080876757-c9a1feec4d9f?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1565557623262-b51c2513a641?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1579584425555-c3ce17fd4351?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1559827260-dc66d52bef19?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1546069901-ba9599a7e63c?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1459925985917-135a9331d55d?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1578308553693-ba4f8c70ae0b?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1563805042-7684c019e1cb?w=500&h=300&fit=crop",
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
			ImageURL:    "https://images.unsplash.com/photo-1555521760-cb7aca801448?w=500&h=300&fit=crop",
			Rating:      4.3,
			IsActive:    true,
		},
	}

	// Create services first and store IDs
	var libraryServiceID string
	var diningServiceID string
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
