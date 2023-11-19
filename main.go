package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/x1xo/wind/client"
	"github.com/x1xo/wind/routes"
	"github.com/x1xo/wind/store"
	"github.com/x1xo/wind/utils"
)

func main() {
	config, err := utils.ParseConfig()
	if err != nil {
		panic(err)
	}

	err = store.Init(config)
	if err != nil {
		panic(err)
	}

	_, err = client.InitClient(config)
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ServerHeader:          "Wind",
	})

	app.Use(limiter.New(limiter.Config{
		Max:        config.Server.LimiterMaxRequests,
		Expiration: config.Server.LimiterDuration,
	}))

	app.Use(func(c *fiber.Ctx) error {
		auth := c.Get("Authorization", "")
		if auth == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "authorization header is not present",
			})
		}

		id, err := (*store.GetStore()).GetIDByKey(auth)
		if err != nil {
			fmt.Println(err)
			return c.Status(500).JSON(fiber.Map{
				"error": "couldn't find a user with that api key",
			})
		}

		queryId := c.Query("id", "")
		if id != queryId {
			return c.Status(403).JSON(fiber.Map{
				"error": "that api is not registered for that id",
			})
		}

		return c.Next()
	})

	app.Get("/presence", func(c *fiber.Ctx) error { return routes.GetPresence(c, config) })

	fmt.Printf("%s Listening on port %d\n", utils.Format(utils.GREEN, "[INFO]"), config.Server.Port)
	app.Listen(fmt.Sprintf("%s:%d", config.Server.IP, config.Server.Port))

}
