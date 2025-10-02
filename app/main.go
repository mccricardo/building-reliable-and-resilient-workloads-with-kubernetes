package main

import (
	"fmt"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/heavy", func(c *fiber.Ctx) error {
		// Simulate a CPU-intensive task
		for i := 0; i < 200000; i++ {
			isPrime(i)
		}
		return c.SendString("Done with heavy lifting!\n")
	})

	app.Get("/ready", func(c *fiber.Ctx) error {
		return c.SendString("Ready")
	})

	app.Get("/live", func(c *fiber.Ctx) error {
		return c.SendString("Live")
	})

	app.Get("/start", func(c *fiber.Ctx) error {
		fmt.Println("Starting up...")
		time.Sleep(15 * time.Second)
		fmt.Println("Startup complete!")
		return c.SendString("Started")
	})

	app.Get("/ready_fail", func(c *fiber.Ctx) error {
		return c.Status(500).SendString("Ready probe failed")
	})

	app.Get("/live_fail", func(c *fiber.Ctx) error {
		return c.Status(500).SendString("Live probe failed")
	})

	app.Get("/start_fail", func(c *fiber.Ctx) error {
		return c.Status(500).SendString("Start probe failed")
	})

	app.Listen(":3000")
}