package handlers

import (
	"caturandi-labs/golang-starter/ent/user"
	"caturandi-labs/golang-starter/middleware"
	"caturandi-labs/golang-starter/utils"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r registerRequest) validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.FirstName, validation.Required, validation.Length(3, 20)),
		validation.Field(&r.LastName, validation.Required, validation.Length(3, 20)),
		validation.Field(&r.Email, is.Email, validation.Required),
		validation.Field(&r.Password, validation.Required, validation.Length(8,32)),
	)
}

func (h *Handlers) UserRegister(ctx *fiber.Ctx) error {
	var request registerRequest

	err := ctx.BodyParser(&request)

	if err != nil {
		err = ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Invalid JSON Format",
		})

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
	}

	if err = request.validate(); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": err,
		})
		return nil
	}

	exist, _ := h.Client.User.Query().
					Where(user.Email(request.Email)).
					First(ctx.Context())

	if (exist != nil) {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "The Email Already taken",
		})
	}

	hashedPassword, err := utils.HashPassword(request.Password)

	if err != nil {
		utils.Errorf("Failed to hash user password : %s", err)
		return nil
	}

	_, err = h.Client.User.Create().
		SetEmail(request.Email).
		SetFirstName(request.FirstName).
		SetLastName(request.LastName).
		SetAvatar(request.Avatar).
		SetPassword(hashedPassword).
		Save(ctx.Context())

	if err != nil {
		utils.Errorf("Fail to create new User: ", err)
		return nil
	}

	_ = ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"error": false,
		"message": "Register Successful",
	})

	return nil
}

func (h *Handlers) UserLogin(ctx *fiber.Ctx) error {
	var request loginRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		err = ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Invalid JSON Format",
		})

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
		return nil
	}

	u, err := h.Client.User.Query().Where(user.Email(request.Email)).First(ctx.Context())

	if err != nil {
		_ = ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Invalid User",
		})
	}

	if err = utils.ComparePassword(request.Password, u.Password); err != nil {
		_ = ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Invalid Credentials",
		})
	}

	userToken, err := middleware.ClaimToken(u.ID)

	if err != nil {
		utils.Errorf("Token Generation Error : %s", err)
		return nil
	}


	response := map[string]interface{} {
		"firstname": u.FirstName,
		"lastName": u.LastName,
		"email": u.Email,
		"avatar": u.Avatar,
	}

	_ = ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data": response,
		"token": userToken,
	})
	return nil
}

func (h *Handlers) MeQuery(ctx *fiber.Ctx)  error {
	userid, err := middleware.GetUserIdFromContext(ctx)

	if err != nil {
		utils.Errorf("Error when get user from context: %s", err)
		return nil
	}

	uid, _ := uuid.Parse(userid)

	userInfo, err := h.Client.User.Query().Where(user.ID(uid)).First(ctx.Context())

	if err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"message": "Sorry, Can't Find the User",
		})
		return nil
	}

	response := map[string]interface{} {
		"firstname": userInfo.FirstName,
		"lastname": userInfo.LastName,
		"email": userInfo.Email,
		"avatar": userInfo.Avatar,
	}

	_ = ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"data": response,
	})

	return nil
}