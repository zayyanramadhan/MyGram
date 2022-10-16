package models

import "time"

type SocialMedia struct {
	Id             int       `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" validate:"required"`
	SocialMediaURL string    `json:"social_media_url" validate:"required"`
	UserId         int       `json:"user_id"`
	User           *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaRequest struct {
	Name           string `json:"name" validate:"required" example:"zayyan"`
	SocialMediaURL string `json:"social_media_url" validate:"required" example:"zayyan"`
}

func SocialMediaToModel(req *SocialMediaRequest) (*SocialMedia, error) {
	socialMedia := &SocialMedia{
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
	}

	return socialMedia, nil
}

type SocialMediaResponse struct {
	Id             int                        `json:"id" gorm:"primaryKey"`
	Name           string                     `json:"name" validate:"required"`
	SocialMediaURL string                     `json:"social_media_url" validate:"required"`
	UserId         int                        `json:"user_id"`
	User           *UserSocialMediaRelational `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	CreatedAt      time.Time                  `json:"created_at"`
}

func SocialMediaToResponse(req *SocialMedia) (*SocialMediaResponseUpdate, error) {
	res := &SocialMediaResponseUpdate{
		Id:             req.Id,
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserId:         req.UserId,
		UpdatedAt:      req.UpdatedAt,
	}

	return res, nil
}

type SocialMediaResponseUpdate struct {
	Id             int                        `json:"id" gorm:"primaryKey"`
	Name           string                     `json:"name" validate:"required"`
	SocialMediaURL string                     `json:"social_media_url" validate:"required"`
	UserId         int                        `json:"user_id"`
	User           *UserSocialMediaRelational `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	UpdatedAt      time.Time                  `json:"updated_at"`
}

func SocialMediaToResponseUpdate(req *SocialMedia) (*SocialMediaResponseUpdate, error) {
	res := &SocialMediaResponseUpdate{
		Id:             req.Id,
		Name:           req.Name,
		SocialMediaURL: req.SocialMediaURL,
		UserId:         req.UserId,
		UpdatedAt:      req.UpdatedAt,
	}

	return res, nil
}

type UserSocialMediaRelational struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	UserName string `json:"username" validate:"required"`
}
