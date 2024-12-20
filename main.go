package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "Mayola's Website"})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// app.Get("/:values?", func(c *fiber.Ctx) error {
	// return c.SendString(c.Params("values"))
	// })
	// app.Get("/*", func(c *fiber.Ctx) error {
	// 	return c.SendString("que es eso: " + c.Params("*"))
	// })
	app.Static("/static", "./static", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	app.Listen(":3000")
}
