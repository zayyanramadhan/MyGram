package repositories

import (
	"final_MyGram/models"
	"log"

	"gorm.io/gorm"
)

type SocialMediaRepo interface {
	CreateSocialMediaRepository(authId int, SocialMedia *models.SocialMedia) (*models.SocialMedia, error)
	GetAllSocialMediaRepository() (*[]models.SocialMediaResponse, error)
	UpdateSocialMediaRepository(authId int, SocialMedia *models.SocialMedia) (*models.SocialMedia, error)
	DeleteSocialMediaRepository(SocialMediaId, authId int) error
}

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepo {
	return &socialMediaRepo{db}
}

func (r *socialMediaRepo) CreateSocialMediaRepository(authId int, socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	socialMedia.UserId = authId
	err := r.db.Create(&socialMedia)
	if err.Error != nil {
		return nil, err.Error
	}
	return socialMedia, nil
}

func (r *socialMediaRepo) GetAllSocialMediaRepository() (*[]models.SocialMediaResponse, error) {
	socialMedia := []models.SocialMediaResponse{}
	if err := r.db.Joins("User").Model(&models.SocialMedia{}).Find(&socialMedia).Error; err != nil {
		log.Printf("Error get data social media with err: %s", err)
		return nil, err
	}
	return &socialMedia, nil
}

func (r *socialMediaRepo) UpdateSocialMediaRepository(authId int, socialMedia *models.SocialMedia) (*models.SocialMedia, error) {
	socialMediaCheckAvail := models.SocialMedia{}
	err := r.db.Where("user_id = ?", authId).
		Where("id = ?", socialMedia.Id).
		First(&socialMediaCheckAvail)
	if err.Error != nil {
		log.Printf("Error get data social media detail with err: %s", err.Error)
		return nil, err.Error
	}
	err = r.db.Model(&socialMediaCheckAvail).Updates(socialMedia)
	if err.Error != nil {
		log.Printf("Error update data social media id with err: %s", err.Error)
		return nil, err.Error
	}
	return &socialMediaCheckAvail, nil
}

func (r *socialMediaRepo) DeleteSocialMediaRepository(socialMediaId, authId int) error {
	socialMediaCheckAvail := models.SocialMedia{}
	err := r.db.Where("user_id = ?", authId).
		Where("id = ?", socialMediaId).
		First(&socialMediaCheckAvail)
	if err.Error != nil {
		log.Printf("Error get data social media detail with err: %s", err.Error)
		return err.Error
	}
	r.db.Delete(socialMediaCheckAvail, socialMediaId)
	return nil
}
