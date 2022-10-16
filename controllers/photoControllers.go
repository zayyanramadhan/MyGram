package controllers

import (
	"final_MyGram/helpers"
	"final_MyGram/models"
	"final_MyGram/repositories"
	"final_MyGram/validations"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	authrepo  repositories.AuthRepo
	photorepo repositories.PhotoRepo
}

func NewPhotoController(repo *repositories.PhotoRepo, authrepo *repositories.AuthRepo) *PhotoController {
	return &PhotoController{photorepo: *repo, authrepo: *authrepo}
}

// Create Photo godoc
// @Summary Create Photo
// @Decription Create Photo
// @Tags photo
// @Accept json
// @Produce json
// @Param Photo body models.PhotoRequest true "Photo info"
// @Router /photos [POST]
func (repo *PhotoController) CreatePhotoController(c *gin.Context) {
	email := c.GetString("email")
	err, authId := repo.authrepo.GetIdByEmail(email)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	var photo models.PhotoRequest
	if err := c.ShouldBindJSON(&photo); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&photo); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	modelPhoto, err := models.PhotoToModel(&photo)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	data, err := repo.photorepo.CreatePhotoRepository(authId, modelPhoto)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	response, err := models.PhotoToResponse(data)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully create photo", response).SuccessCreate(c)
}

// Get Photo godoc
// @Summary Get Photo
// @Decription Get Photo
// @Tags photo
// @Accept json
// @Produce json
// @Router /photos [GET]
func (repo *PhotoController) GetAllPhotoController(c *gin.Context) {
	data, err := repo.photorepo.GetAllPhotoRepository()
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully get all photo", data).Success(c)
}

// Update Photo godoc
// @Summary Update Photo
// @Decription Update Photo
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "photoId"
// @Param Photo body models.PhotoRequest true "Photo info"
// @Router /photos/{photoId} [PUT]
func (repo *PhotoController) UpdatePhotoController(c *gin.Context) {
	email := c.GetString("email")
	err, authId := repo.authrepo.GetIdByEmail(email)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	var photo models.PhotoRequest

	if err := c.ShouldBindJSON(&photo); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}
	modelPhoto, err := models.PhotoToModel(&photo)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	modelPhoto.Id = photoId
	if err := validations.DoValidation(&photo); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}
	data, err := repo.photorepo.UpdatePhotoRepository(authId, modelPhoto)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	response, err := models.PhotoToResponseUpdate(data)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully update photo", response).Success(c)
}

// Delete Photo godoc
// @Summary Delete Photo
// @Decription Delete Photo
// @Tags photo
// @Accept json
// @Produce json
// @Param photoId path int true "photoId"
// @Router /photos/{photoId} [DELETE]
func (repo *PhotoController) DeletePhotoController(c *gin.Context) {
	email := c.GetString("email")
	err, authId := repo.authrepo.GetIdByEmail(email)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	err = repo.photorepo.DeletePhotoRepository(photoId, authId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Your photo has been successfully delete", nil).Success(c)
}
