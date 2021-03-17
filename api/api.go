package api

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/raysandeep/Agora-Token-Server-Golang/utils"

	"github.com/gofiber/fiber/v2"
)

func createRTCToken(c *fiber.Ctx) error {
	channel := c.Params("channel")
	uid := int(rand.Uint32())
	rtcToken, err := utils.GetRtcToken(channel, uid)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"msg": http.StatusInternalServerError,
			"err": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code":      http.StatusOK,
		"rtc_token": rtcToken,
	})
}

func createRTMToken(c *fiber.Ctx) error {
	uid := c.Params("uid")
	rtmToken, err := utils.GetRtmToken(fmt.Sprint(uid))
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"msg": http.StatusInternalServerError,
			"err": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"code":      http.StatusOK,
		"rtm_token": rtmToken,
	})
}

func createTokens(c *fiber.Ctx) error {
	channel := c.Params("channel")
	uid := int(rand.Uint32())
	rtcToken, err := utils.GetRtcToken(channel, uid)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"msg": http.StatusInternalServerError,
			"err": err.Error(),
		})
	}
	rtmToken, err := utils.GetRtmToken(fmt.Sprint(uid))
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"msg": http.StatusInternalServerError,
			"err": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"code":      http.StatusOK,
		"rtc_token": rtcToken,
		"rtm_token": rtmToken,
	})
}

// MountRoutes mounts all routes declared here
func MountRoutes(app *fiber.App) {
	app.Get("/api/get/rtc/:channel", createRTCToken)
	app.Get("/api/get/rtm/:uid", createRTMToken)
	app.Get("/api/tokens/:channel", createTokens)
}
