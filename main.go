package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/SarjitDelivala/aws-ec2-dashboard/aws"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		instances, err := aws.GetInstanceList()
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
		return c.JSON(instances)
	})

	err := app.Listen(":8001")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
