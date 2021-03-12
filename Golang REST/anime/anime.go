package anime

import (
	"github.com/London-p/go-fiber-test/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Anime struct {
	gorm.Model
	Title string "json:'title'"
	Rank  int    "json:'rank'"
}

func GetAnimes(c *fiber.Ctx) {
	db := database.DBConn
	var animes []Anime
	db.Find(&animes)
	c.JSON(animes)
}

func GetAnime(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var anime Anime
	db.Find(&anime, id)
	c.JSON(anime)
}

func NewAnime(c *fiber.Ctx) {
	db := database.DBConn
	anime := new(Anime)
	if err := c.BodyParser(anime); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&anime)
	c.JSON(anime)
}

func DeleteAnime(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var anime Anime
	db.First(&anime, id)
	if anime.Title == "" {
		c.Status(500).Send("No anime found with given id")
		return
	}
	db.Delete(&anime)

	c.Send("Anime was deleted")
}
