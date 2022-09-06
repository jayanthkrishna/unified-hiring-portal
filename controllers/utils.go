package controllers

import (
	"fmt"

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

func retrieve_email(c *fiber.Ctx) (string, error) {
	claims, err := token_verify(c)

	if err != nil {
		return "", c.JSON(fiber.Map{
			"Error": err,
		})
	}

	return claims["jti"].(string), nil
}
