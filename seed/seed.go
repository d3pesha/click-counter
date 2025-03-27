package seed

import (
	"database/sql"
	"log"
)

func Banners(db *sql.DB) {
	banners := []string{
		"Promo Banner", "Sale Banner", "Black Friday", "New Year Sale",
		"Summer Discounts", "Winter Clearance", "Back to School", "Holiday Special",
	}

	for _, name := range banners {
		_, err := db.Exec("INSERT INTO banners (name) VALUES ($1) ON CONFLICT DO NOTHING;", name)
		if err != nil {
			log.Fatalf("error inserting banner %s: %v", name, err)
		}
	}

	log.Println("Banners seeded successfully")
}
