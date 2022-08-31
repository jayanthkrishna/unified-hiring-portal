package controllers

import (
	"context"
	"jwt-auth-go/database"
	"jwt-auth-go/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "secret"

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!!")
}

func Register(c *fiber.Ctx) error {
	var data models.User

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	data.Password = password
	data.ID = primitive.NewObjectID()

	DB := database.DB
	_, err := DB.InsertOne(context.TODO(), &data)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		// var e mongo.WriteException
		e := err.(mongo.WriteException)
		return c.JSON(fiber.Map{
			"code":    e.WriteErrors[0].Code,
			"message": e.WriteErrors[0].Message,
		})
	}
	return c.JSON(data)
}

func Login(c *fiber.Ctx) error {
	var data models.User

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	filter := bson.M{"email": data.Email}
	err := database.DB.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		c.Status(fiber.StatusNotFound)

		return c.JSON(fiber.Map{
			"message": "User Doesnt Exist",
		})
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{
			"message": "Password Incorrect",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Name,
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
