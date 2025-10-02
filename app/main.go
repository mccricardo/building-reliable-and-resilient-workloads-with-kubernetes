package main

import (
	"math"

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

	app.Listen(":3000")
}
