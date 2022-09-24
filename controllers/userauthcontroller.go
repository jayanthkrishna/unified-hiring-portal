package controllers

import (
	"strconv"
	"time"
	"unified-hiring-portal/database"
	"unified-hiring-portal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"net/http"
)

const secretKey = "secret"

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!!")
}

func Register(c *fiber.Ctx) error {
	var data models.User

	err := c.BodyParser(&data)

	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request Failed"})
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	data.Password = password

	DB := database.DB

	err = DB.Create(&data).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create a User"})
		return err
	}

	return c.JSON(data)
}

func Login(c *fiber.Ctx) error {
	var data models.User

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	err := database.DB.Where("Email = ?", data.Email).First(&user).Error

	if err != nil {
		c.Status(fiber.StatusNotFound)

		return c.JSON(fiber.Map{
			"message": "User Doesnt Exist",
		})
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{
			"message": "User Exist but Password Incorrect",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Name,
		Id:        strconv.FormatUint(uint64(user.ID), 10),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})

}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "UnAuthorized",
		})
	}

	claims := token.Claims

	return c.JSON(claims)
}
