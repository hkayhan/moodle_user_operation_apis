package main

import (
	"copy_users_for_moodle/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	// .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	//db.Init()

	// PostgreSQL veritabanına bağlan
	//database.ConnectDB()

	// Users tablosunu oluştur
	//models.CreateUsersTable()

	// Sunucuyu başlat
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Varsayılan port
	}

	r := routes.SetupRouter()

	log.Println("Starting server on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
