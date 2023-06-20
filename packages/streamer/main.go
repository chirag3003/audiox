package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

const sampleRate = 44100
const seconds = 1

func main() {
	app := fiber.New()
	file, _ := os.Open("music/music1.mp3")

	app.Get("/", func(c *fiber.Ctx) error {
		for {
			buffer := make([]byte, 100)
			_, err := file.Read(buffer)
			if err != nil {
				break
			}
			c.Write(buffer)
		}
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
