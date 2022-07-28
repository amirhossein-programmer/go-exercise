package main

import "github.com/gofiber/fiber/v2"

type Ninja struct {
	name   string
	weapon string
}

func getNinja(ctx *fiber.Ctx) error {
	var wallace = Ninja{
		name:   "wallace",
		weapon: "Ninja Star",
	}
	return ctx.Status(fiber.StatusOK).JSON(wallace)
}
func createNinja(ctx *fiber.Ctx) error {
	body := new(Ninja)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	ninja := Ninja{
		name:   body.name,
		weapon: body.weapon,
	}
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func main() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world!")
	})
	app.Get("/ninja", getNinja)
	app.Post("/ninja", createNinja)
	app.Listen(":3000")
}
