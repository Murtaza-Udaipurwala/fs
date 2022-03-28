package api

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

	app.Use(limiter.New(limiter.Config{
		Expiration: time.Second * 30,
		Max:        3,
		Next: func(ctx *fiber.Ctx) bool {
			if ctx.Method() == "GET" {
				return true
			}

			return false
		},
	}))

	app.Use(cors.New())

	c := NewController(s)
	route(app, c)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	log.Printf("Listening on port :%s", port)
	log.Fatal(app.Listen(":" + port))
}
