package middleware

import (
	"caturandi-labs/golang-starter/config"
	"caturandi-labs/golang-starter/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func IsAuthenticated(config *config.Config) func(ctx *fiber.Ctx) error{
	return jwtware.New(
		jwtware.Config{
			SigningKey: []byte(config.Jwt.Secret),
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": true,
					"message": "Unauthorized Access " + string(err.Error()) ,
				})
				return nil
			},
		})
}

func GetUserIdFromContext(ctx *fiber.Ctx) (string, error) {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	strId := claims["sub"].(string)

	return strId, nil
}

func ClaimToken(id uuid.UUID) (string, error){
	config := config.New()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() // = 1month

	s, err := token.SignedString(config.Jwt.Secret)

	if err != nil {
		utils.Errorf("Error Claim Token : %s", err)
		return "", err
	}

	return s, nil
}