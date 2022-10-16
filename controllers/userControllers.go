package controllers

import (
	"final_MyGram/helpers"
	"final_MyGram/models"
	"final_MyGram/repositories"
	"final_MyGram/validations"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	authrepo repositories.AuthRepo
	userrepo repositories.UserRepo
}

func NewUserController(repo *repositories.UserRepo, authrepo *repositories.AuthRepo) *UserController {
	return &UserController{userrepo: *repo, authrepo: *authrepo}
}

// Update godoc
// @Summary Update
// @Decription Update
// @Tags user
// @Accept json
// @Produce json
// @Param userId path int true "userId"
// @Param user body models.UserUpdate true "user info"
// @Router /users/{userId} [PUT]
func (repo *UserController) UpdateUserController(c *gin.Context) {
	var user models.UserUpdate
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data userId", nil).BadRequest(c)
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}
	if err := validations.DoValidation(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}
	data, err := repo.userrepo.UpdateUserRepository(id, &user)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	response, err := models.UpdateToResponse(data)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully update user", response).Success(c)
}

// Delete godoc
// @Summary Delete
// @Decription Delete
// @Tags user
// @Accept json
// @Produce json
// @Router /users [DELETE]
func (repo *UserController) DeleteUserController(c *gin.Context) {

	email := c.GetString("email")
	err, authId := repo.authrepo.GetIdByEmail(email)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	err = repo.userrepo.DeleteUserRepository(authId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	helpers.NewHandlerResponse("Your account has been successfully deleted", nil).Success(c)
}
