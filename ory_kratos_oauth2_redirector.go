package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Configuration struct {
	port                 string
	loginEndpoint        string
	registrationEndpoint string
}

func (c *Configuration) load() {
	c.port = os.Getenv("PORT")
	c.loginEndpoint = os.Getenv("LOGIN_ENDPOINT")
	c.registrationEndpoint = os.Getenv("REGISTRATION_ENDPOINT")
}

func (c *Configuration) display() {
	fmt.Println(
		"Configuration:\n##############",
		"\nPORT: ", c.port,
		"\nLOGIN_ENDPOINT: ", c.loginEndpoint,
		"\nREGISTRATION_ENDPOINT: ", c.registrationEndpoint,
	)
}

func main() {
	c := &Configuration{}
	c.load()
	c.display()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		AppName: "Ory Kratos OAuth2 Redirector",
		Views:   engine,
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		endpoint := c.loginEndpoint

		var traits map[string]interface{}

		if ctx.Query("traits") != "" {
			endpoint = c.registrationEndpoint
			traitsErr := json.Unmarshal([]byte(ctx.Query("traits")), &traits)

			if traitsErr != nil {
				return ctx.SendStatus(400)
			}
		}

		return ctx.Render("index", fiber.Map{
			"endpoint":   endpoint,
			"provider":   ctx.Query("provider"),
			"csrf_token": ctx.Query("csrf_token"),
			"flow":       ctx.Query("flow"),
			"traits":     traits,
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", c.port)))
}
