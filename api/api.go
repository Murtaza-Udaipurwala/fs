package api

import (
	"os"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func route(app *fiber.App, c *Controller) {
	app.Get("/:id", c.Retrieve)
	app.Post("/", c.Create)
}

func Serve(s *Service) {
	app := fiber.New(fiber.Config{
		BodyLimit:     maxUploadSize,
		CaseSensitive: true,
	})

	c := NewController(s)
	route(app, c)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	log.Printf("Listening on port :%s", port)
	log.Fatal(app.Listen(":" + port))
}
