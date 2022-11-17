package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type configuration struct {
	port     string
	endpoint string
}

func (c *configuration) load() {
	c.port = os.Getenv("PORT")
}

func (c *configuration) display() {
	fmt.Println(
		"Configuration:\n##############",
		"\nPORT: ", c.port,
		"\nENDPOINT: ", c.endpoint,
	)
}

func main() {
	c := &configuration{}
	c.load()
	c.display()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		AppName: "Ory Kratos OAuth2 Redirector",
		Views:   engine,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"endpoint":   c.endpoint,
			"provider":   ctx.Query("provider"),
			"csrf_token": ctx.Query("csrf_token"),
			"flow":       ctx.Query("flow"),
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", c.port)))
}
