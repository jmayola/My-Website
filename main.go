package main

import (
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/jmayola/fiber/middleware"
	"github.com/jmayola/fiber/ws"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "Mayola's Website"})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Use(middleware.MiddleCsrf())

	// Or extend your config for customization
	// app.Get("/:values?", func(c *fiber.Ctx) error {
	// return c.SendString(c.Params("values"))
	// })

	// app.Get("/*", func(c *fiber.Ctx) error {
	// 	return c.SendString("que es eso: " + c.Params("*"))
	// })

	// // Match requests starting with /api or /home (multiple-prefix support)
	// app.Use([]string{"/api", "/home"}, func(c *fiber.Ctx) error {
	// 	return c.Next()
	// })
	// // app.Route("/test", func(api fiber.Router) {
	// // 	api.Get("/foo", handler).Name("foo") // /test/foo (name: test.foo)
	// // 	api.Get("/bar", handler).Name("bar") // /test/bar (name: test.bar)
	// // }, "test.")
	app.Use("/ws", func(c *fiber.Ctx) error {
		if c.Get("host") == "localhost:3000" {
			c.Locals("Host", "Localhost:3000")
			return c.Next()
		}
		return c.Status(403).SendString("Request origin not allowed")
	})
	// Upgraded websocket request
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			println("accepted")
			c.Locals("allowed", true)
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})
	app.Get("/ws", websocket.New(ws.GetWs))

	app.Static("/static", "./static", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	log.Fatal(app.Listen(":3000"))
}
