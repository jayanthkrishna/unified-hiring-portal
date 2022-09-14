package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func token_verify(c *fiber.Ctx) (jwt.MapClaims, error) {
	cookie := c.Cookies("jwt")
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cookie, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
	fmt.Println(token.Valid)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)

		return nil, c.JSON(fiber.Map{
			"message": "UnAuthorized",
		})
	}

	return claims, nil
}

func retrieve_id(c *fiber.Ctx) (uint, error) {
	claims, err := token_verify(c)

	if err != nil {
		return 0, c.JSON(fiber.Map{
			"Error": err,
		})
	}
	u, _ := strconv.ParseUint(claims["jti"].(string), 10, 64)
	return uint(u), nil
}
