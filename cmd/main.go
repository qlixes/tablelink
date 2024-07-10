package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/auth/login", func(r *fiber.Ctx) error {

	})

	app.Get("/users/user", func(r *fiber.Ctx) error {

	})

	app.Get("/users/user", func(r *fiber.Ctx) error {

	})

	app.Get("/users/user", func(r *fiber.Ctx) error {

	})

	app.Get("/users/user", func(r *fiber.Ctx) error {

	})

	log.Fatal(app.Listen(":3000"))
}
