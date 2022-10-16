package models

import "time"

type Comment struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Message   string    `json:"message" validate:"required" example:"zayyan"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_id" example:"1"`
	User      *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Photo     *Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photo,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentRequest struct {
	Message string `json:"message" validate:"required" example:"zayyan"`
	PhotoId int    `json:"photo_id" example:"1"`
}

func CommentToModel(req *CommentRequest) (*Comment, error) {
	comment := &Comment{
		Message: req.Message,
		PhotoId: req.PhotoId,
	}

	return comment, nil
}

type CommentRequestUpdate struct {
	Message string `json:"message" validate:"required" example:"zayyan"`
}

func CommentUpdateToModel(req *CommentRequestUpdate) (*Comment, error) {
	comment := &Comment{
		Message: req.Message,
	}

	return comment, nil
}

type CommentResponse struct {
	Id        int                     `json:"id" gorm:"primaryKey"`
	Message   string                  `json:"message" validate:"required"`
	UserId    int                     `json:"user_id"`
	PhotoId   int                     `json:"photo_id"`
	User      *UserCommentRelational  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Photo     *PhotoCommentRelational `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photo,omitempty"`
	CreatedAt time.Time               `json:"created_at"`
}

func CommentToResponse(req *Comment) (*CommentResponse, error) {
	res := &CommentResponse{
		Id:        req.Id,
		Message:   req.Message,
		PhotoId:   req.PhotoId,
		UserId:    req.UserId,
		CreatedAt: req.CreatedAt,
	}

	return res, nil
}

type CommentResponseUpdate struct {
	Id        int                     `json:"id" gorm:"primaryKey"`
	Message   string                  `json:"message" validate:"required"`
	UserId    int                     `json:"user_id"`
	PhotoId   int                     `json:"photo_id"`
	User      *UserCommentRelational  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Photo     *PhotoCommentRelational `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photo,omitempty"`
	UpdatedAt time.Time               `json:"updated_at"`
}

func CommentToResponseUpdate(req *Comment) (*CommentResponseUpdate, error) {
	res := &CommentResponseUpdate{
		Id:        req.Id,
		Message:   req.Message,
		PhotoId:   req.PhotoId,
		UserId:    req.UserId,
		UpdatedAt: req.UpdatedAt,
	}

	return res, nil
}

type UserCommentRelational struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" validate:"required"`
	UserName string `json:"username" validate:"required"`
}

type PhotoCommentRelational struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" validate:"required"`
	UserId   int    `json:"user_id"`
}
