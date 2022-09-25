package controllers

import (
	"time"
	"unified-hiring-portal/database"
	"unified-hiring-portal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateToken(c *fiber.Ctx) error {

	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return err
	}

	res := models.Client{}

	err := database.DB.Where("ID = ?", client.ID).First(&res).Error

	if err != nil {
		return c.JSON(fiber.Map{
			"error":   err,
			"Message": "Client Doesnt Exist",
		})
	}

	if client.Secret != res.Secret {
		return c.JSON(fiber.Map{
			"Message": "Secrets Dont Match",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        client.ID.String(),
		Subject:   client.Name,
		Audience:  client.ClientUser,
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"access_token": token,
	})

}

func client_verify(c *fiber.Ctx) (jwt.MapClaims, error) {
	type Token struct {
		Token string `json:"access_token"`
	}
	access_token := Token{}
	c.BodyParser(&access_token)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(access_token.Token, &claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)

		return nil, c.JSON(fiber.Map{
			"message": "UnAuthorized",
		})
	}

	return claims, nil

}

func retrieve_Client_id(c *fiber.Ctx) (uuid.UUID, error) {
	claims, err := client_verify(c)

	if err != nil {
		return uuid.Nil, c.JSON(fiber.Map{
			"Error": err,
		})
	}
	return uuid.Parse(claims["jti"].(string))
}
