package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

/*
gerar chaves
ssh-keygen -f ~/<name> -t rsa -m PKCS8 -b 2048

export keys
ssh-keygen -m PKCS8 -e
*/

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func main() {

	// load ssh keys
	err := loadSSHKeys()
	if err != nil {
		fmt.Println("erro loadSSHKeys")
		log.Fatal(err)
	}

	app := fiber.New()

	// Login route
	app.Post("/login", login)

	// Unauthenticated route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("not accessible")
	})

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    publicKey,
	}))

	app.Get("/restricted", func(c *fiber.Ctx) error {
		return c.SendString("server ok")
	})

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}

func login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "admin" || pass != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "admin",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(privateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}

func loadSSHKeys() error {

	privateByte, err := ioutil.ReadFile("/home/rafael/.ssh/jwtRS256")
	if err != nil {
		return err
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateByte)
	if err != nil {
		return err
	}

	publicByte, err := ioutil.ReadFile("/home/rafael/.ssh/jwtRS256.pub")
	if err != nil {
		return err
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicByte)
	if err != nil {
		return err
	}

	return nil
}
