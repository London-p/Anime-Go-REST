package main

import (
	"fmt"

	"github.com/London-p/go-fiber-test/anime"
	"github.com/London-p/go-fiber-test/database"
	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/anime", anime.GetAnimes)
	app.Get("/api/v1/anime/:id", anime.GetAnime)
	app.Post("/api/v1/anime", anime.NewAnime)
	app.Delete("/api/v1/anime/:id", anime.DeleteAnime)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "animes.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection succesfully opened")

	database.DBConn.AutoMigrate(&anime.Anime{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}
