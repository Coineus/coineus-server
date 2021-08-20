package app

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	uuid "github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		log.Println("uuid create fail. ", err)
	}
	return id.String()
}

func CreateToken(userId int, mail, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = userId
	claims["uname"] = username
	claims["email"] = mail
	//claims["admin"] = isAdmin
	exp, _ := strconv.Atoi(os.Getenv("JWT_EXP_TIME"))
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(exp)).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func CheckPasswordHash(pass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err == nil
}

func GetUserClaims(c *fiber.Ctx) jwt.MapClaims {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
