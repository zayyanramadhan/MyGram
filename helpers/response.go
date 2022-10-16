package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HandlerValidationResponse struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type HandlerDBResponse struct {
	Status  int         `json:"status"`
	Message error       `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHandlerResponse(message string, data interface{}) *HandlerResponse {
	return &HandlerResponse{
		Message: message,
		Data:    data,
	}
}

func NewHandlerValidationResponse(message, data interface{}) *HandlerValidationResponse {
	return &HandlerValidationResponse{
		Message: message,
		Data:    data,
	}
}

func NewHandlerDBResponse(message error, data interface{}) *HandlerDBResponse {
	return &HandlerDBResponse{
		Message: message,
		Data:    data,
	}
}

func (response *HandlerResponse) Success(c *gin.Context) {
	response.Status = http.StatusOK
	c.JSON(http.StatusOK, response)
}

func (response *HandlerResponse) SuccessCreate(c *gin.Context) {
	response.Status = http.StatusCreated
	c.JSON(http.StatusCreated, response)
}

func (response *HandlerResponse) Failed(c *gin.Context) {
	response.Status = http.StatusInternalServerError
	c.JSON(http.StatusInternalServerError, response)
}

func (response *HandlerResponse) BadRequest(c *gin.Context) {
	response.Status = http.StatusBadRequest
	c.JSON(http.StatusBadRequest, response)
}

func (response *HandlerResponse) Unauthorized(c *gin.Context) {
	response.Status = http.StatusUnauthorized
	c.JSON(http.StatusUnauthorized, response)
}

func (response *HandlerValidationResponse) Failed(c *gin.Context) {
	response.Status = http.StatusInternalServerError
	c.JSON(http.StatusInternalServerError, response)
}

func (response *HandlerValidationResponse) BadRequest(c *gin.Context) {
	response.Status = http.StatusBadRequest
	c.JSON(http.StatusBadRequest, response)
}

func (response *HandlerDBResponse) Failed(c *gin.Context) {
	response.Status = http.StatusInternalServerError
	c.JSON(http.StatusInternalServerError, response)
}

type JWTResponse struct {
	Token string `json:"token"`
}
