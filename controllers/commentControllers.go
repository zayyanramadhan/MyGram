package controllers

import (
	"final_MyGram/helpers"
	"final_MyGram/models"
	"final_MyGram/repositories"
	"final_MyGram/validations"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	authrepo    repositories.AuthRepo
	commentrepo repositories.CommentRepo
}

func NewCommentController(repo *repositories.CommentRepo, authrepo *repositories.AuthRepo) *CommentController {
	return &CommentController{commentrepo: *repo, authrepo: *authrepo}
}

// Create Comment godoc
// @Summary Create Comment
// @Decription Create Comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body models.CommentRequest true "comment info"
// @Router /comments [POST]
func (repo *CommentController) CreateCommentController(c *gin.Context) {
	var comment models.CommentRequest
	email := c.GetString("email")
	err, authId := repo.authrepo.GetIdByEmail(email)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	if err := c.ShouldBindJSON(&comment); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&comment); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}
	modelComment, err := models.CommentToModel(&comment)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}

	data, err := repo.commentrepo.CreateCommentRepository(authId, modelComment)
	if err != nil {
		helpers.NewHandlerResponse("Photo ID Not found", nil).BadRequest(c)
		return
	}
	response, err := models.CommentToResponseUpdate(data)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully create comment", response).SuccessCreate(c)
}

// Get Comment godoc
// @Summary Get Comment
// @Decription Get Comment
// @Tags comment
// @Accept json
// @Produce json
// @Router /comments [GET]
func (repo *CommentController) GetAllCommentController(c *gin.Context) {
	data, err := repo.commentrepo.GetAllCommentRepository()
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully get all comment", data).Success(c)
}

// Update Comment godoc
// @Summary Update Comment
// @Decription Update Comment
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "commentId"
// @Param comment body models.CommentRequestUpdate true "comment info"
// @Router /comments/{commentId} [PUT]
func (repo *CommentController) UpdateCommentController(c *gin.Context) {
	email := c.GetString("email")
	err, authId := repo.authrepo.GetIdByEmail(email)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	var comment models.CommentRequestUpdate

	if err := c.ShouldBindJSON(&comment); err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).BadRequest(c)
		return
	}

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}

	if err := validations.DoValidation(&comment); err != nil {
		helpers.NewHandlerValidationResponse(err, nil).BadRequest(c)
		return
	}

	modelComment, err := models.CommentUpdateToModel(&comment)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	modelComment.Id = commentId

	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	data, err := repo.commentrepo.UpdateCommentRepository(authId, modelComment)
	if err != nil {
		helpers.NewHandlerResponse("Comment ID Not found", nil).BadRequest(c)
		return
	}
	response, err := models.CommentToResponse(data)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Successfully update comment", response).Success(c)
}

// Delete Comment godoc
// @Summary Delete Comment
// @Decription Delete Comment
// @Tags comment
// @Accept json
// @Produce json
// @Param commentId path int true "commentId"
// @Router /comments/{commentId} [DELETE]
func (repo *CommentController) DeleteCommentController(c *gin.Context) {
	email := c.GetString("email")
	err, authId := repo.authrepo.GetIdByEmail(email)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		helpers.NewHandlerResponse("Error convert data", nil).BadRequest(c)
		return
	}
	err = repo.commentrepo.DeleteCommentRepository(commentId, authId)
	if err != nil {
		helpers.NewHandlerResponse(err.Error(), nil).Failed(c)
		return
	}
	helpers.NewHandlerResponse("Your comment has been successfully delete", nil).Success(c)
}
