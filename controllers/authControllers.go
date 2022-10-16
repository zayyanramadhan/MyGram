package controllers

import (
	"final_MyGram/helpers"
	"final_MyGram/models"
	"final_MyGram/repositories"
	"final_MyGram/validations"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authrepo repositories.AuthRepo
}

func NewAuthController(repo *repositories.AuthRepo) *AuthController {
	return &AuthController{authrepo: *repo}
}

// Register godoc
// @Summary Register
// @Decription Register
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UserRegisterRequest true "user info"
// @Router /users/register [POST]
func (repo *AuthController) RegisterController(c *gin.Context) {
	var user models.UserRegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password with err: %s", err)
		helpers.NewHandlerResponse("Error hash", nil).Failed(c)
		return
	}
	user.Password = string(password)
	modelUser, err := models.RegisterToModel(&user)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	err = repo.authrepo.RegisterRepository(modelUser)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	responseUser, err := models.RegisterToResponse(modelUser)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	helpers.NewHandlerResponse("Successfully register", responseUser).SuccessCreate(c)
}

// Login godoc
// @Summary Login
// @Decription Login
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UserLoginRequest true "user info"
// @Router /users/login [POST]
func (repo *AuthController) LoginController(c *gin.Context) {
	var user models.UserLoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}
	if err := validations.DoValidation(&user); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}
	modelUser, err := models.LoginToModel(&user)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	userData, err := repo.authrepo.LoginRepository(modelUser)
	if err != nil {
		helpers.NewHandlerResponse("Email not found", nil).BadRequest(c)
		return
	}
	errHash := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if errHash != nil {
		log.Printf("Email or Password Incorrect with err: %s\n", errHash)
		helpers.NewHandlerResponse("Email or password is incorrect", nil).BadRequest(c)
		return
	}
	tokenString, err := helpers.GenerateJWT(userData)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		c.Abort()
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24 * 60),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	resJWT := helpers.JWTResponse{
		Token: tokenString,
	}
	helpers.NewHandlerResponse("Successfully Login", resJWT).Success(c)
}

// Logout godoc
// @Summary Logout
// @Decription Logout
// @Tags user
// @Accept json
// @Produce json
// @Router /users/logout [POST]
func (repo *AuthController) LogoutController(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",
		// Local
		SameSite: 2,
		HttpOnly: true,
	})
	helpers.NewHandlerResponse("Successfully logout", nil).Success(c)
}
