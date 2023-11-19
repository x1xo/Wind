package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/x1xo/wind/client"
	"github.com/x1xo/wind/utils"
)


func GetPresence(c *fiber.Ctx, config *utils.Config) error {
	id := c.Query("id", "")
	if id == "" {
		return c.JSON(fiber.Map{
			"error": "Invalid Id query parameter.",
		})
	}

	presence, err := client.GetClient().State.Presence(config.Discord.Guild_ID, id)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Something went wrong while fetching member. Maybe the user is not in the guild.",
		})
	}

	return c.JSON(presence)
}
