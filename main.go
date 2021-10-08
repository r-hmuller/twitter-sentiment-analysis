package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"twitterSearchSentimentalAnalysis/services"
)

func main() {
	app := fiber.New()

	app.Get("/search", func(c *fiber.Ctx) error {
		param := c.Query("term")
		go services.DoTwitterSearch(param)
		return c.SendStatus(202)
	})

	log.Fatal(app.Listen(":9999"))
}
