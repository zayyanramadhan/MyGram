package controllers

import (
	"final_MyGram/helpers"
	"final_MyGram/models"
	"final_MyGram/repositories"
	"final_MyGram/validations"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	authrepo        repositories.AuthRepo
	socialMediarepo repositories.SocialMediaRepo
}

func NewSocialMediaController(repo *repositories.SocialMediaRepo, authrepo *repositories.AuthRepo) *SocialMediaController {
	return &SocialMediaController{socialMediarepo: *repo, authrepo: *authrepo}
}

// Create SocialMedia godoc
// @Summary Create SocialMedia
// @Decription Create SocialMedia
// @Tags socialmedia
// @Accept json
// @Produce json
// @Param SocialMedia body models.SocialMediaRequest true "SocialMedia info"
// @Router /socialmedias [POST]
func (repo *SocialMediaController) CreateSocialMediaController(c *gin.Context) {
	email := c.GetString("email")
	err, authId := repo.authrepo.GetIdByEmail(email)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	var socialMedia models.SocialMediaRequest
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&socialMedia); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	modelSocialMedia, err := models.SocialMediaToModel(&socialMedia)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	data, err := repo.socialMediarepo.CreateSocialMediaRepository(authId, modelSocialMedia)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	response, err := models.SocialMediaToResponse(data)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully create social media", response).SuccessCreate(c)
}

// Get SocialMedia godoc
// @Summary Get SocialMedia
// @Decription Get SocialMedia
// @Tags socialmedia
// @Accept json
// @Produce json
// @Router /socialmedias [GET]
func (repo *SocialMediaController) GetAllSocialMediaController(c *gin.Context) {
	data, err := repo.socialMediarepo.GetAllSocialMediaRepository()
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully get all social media", data).Success(c)
}

// Update SocialMedia godoc
// @Summary Update SocialMedia
// @Decription Update SocialMedia
// @Tags socialmedia
// @Accept json
// @Produce json
// @Param socialMediaId path int true "socialMediaId"
// @Param SocialMedia body models.SocialMediaRequest true "SocialMedia info"
// @Router /socialmedias/{socialMediaId} [PUT]
func (repo *SocialMediaController) UpdateSocialMediaController(c *gin.Context) {
	email := c.GetString("email")
	err, authId := repo.authrepo.GetIdByEmail(email)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	var socialMedia models.SocialMediaRequest

	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&socialMedia); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	modelSocialMedia, err := models.SocialMediaToModel(&socialMedia)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	modelSocialMedia.Id = socialMediaId

	data, err := repo.socialMediarepo.UpdateSocialMediaRepository(authId, modelSocialMedia)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	response, err := models.SocialMediaToResponseUpdate(data)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully update social media", response).Success(c)
}

// Delete SocialMedia godoc
// @Summary Delete SocialMedia
// @Decription Delete SocialMedia
// @Tags socialmedia
// @Accept json
// @Produce json
// @Param socialMediaId path int true "socialMediaId"
// @Router /socialmedias/{socialMediaId} [DELETE]
func (repo *SocialMediaController) DeleteSocialMediaController(c *gin.Context) {
	email := c.GetString("email")
	err, authId := repo.authrepo.GetIdByEmail(email)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}

	err = repo.socialMediarepo.DeleteSocialMediaRepository(socialMediaId, authId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Your social media has been successfully delete", nil).Success(c)
}
