package models

import (
	"time"
)

type Photo struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" validate:"required" example:"zayyan"`
	Caption   string    `json:"caption" example:"zayyan"`
	PhotoURL  string    `json:"photo_url" validate:"required" example:"zayyan"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

type PhotoRequest struct {
	Title    string `json:"title" validate:"required" example:"zayyan"`
	Caption  string `json:"caption" example:"zayyan"`
	PhotoURL string `json:"photo_url" validate:"required" example:"zayyan"`
}

func PhotoToModel(req *PhotoRequest) (*Photo, error) {
	photo := &Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
	}

	return photo, nil
}

type PhotoResponse struct {
	Id        int             `json:"id" gorm:"primaryKey"`
	Title     string          `json:"title" validate:"required"`
	Caption   string          `json:"caption"`
	PhotoURL  string          `json:"photo_url" validate:"required"`
	UserId    int             `json:"user_id"`
	User      *UserRelational `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
}

func PhotoToResponse(req *Photo) (*PhotoResponse, error) {
	res := &PhotoResponse{
		Id:        req.Id,
		Title:     req.Title,
		Caption:   req.Caption,
		PhotoURL:  req.PhotoURL,
		UserId:    req.UserId,
		CreatedAt: req.CreatedAt,
	}

	return res, nil
}

type PhotoResponseUpdate struct {
	Id        int             `json:"id" gorm:"primaryKey"`
	Title     string          `json:"title" validate:"required"`
	Caption   string          `json:"caption"`
	PhotoURL  string          `json:"photo_url" validate:"required"`
	UserId    int             `json:"user_id"`
	User      *UserRelational `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func PhotoToResponseUpdate(req *Photo) (*PhotoResponseUpdate, error) {
	res := &PhotoResponseUpdate{
		Id:        req.Id,
		Title:     req.Title,
		Caption:   req.Caption,
		PhotoURL:  req.PhotoURL,
		UserId:    req.UserId,
		UpdatedAt: req.UpdatedAt,
	}

	return res, nil
}

type UserRelational struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" validate:"required"`
	UserName string `json:"username" validate:"required"`
}
