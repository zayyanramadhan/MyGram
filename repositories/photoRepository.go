package repositories

import (
	"final_MyGram/models"
	"log"

	"gorm.io/gorm"
)

type PhotoRepo interface {
	CreatePhotoRepository(authId int, Photo *models.Photo) (*models.Photo, error)
	GetAllPhotoRepository() (*[]models.PhotoResponse, error)
	UpdatePhotoRepository(authId int, Photo *models.Photo) (*models.Photo, error)
	DeletePhotoRepository(PhotoId, authId int) error
}

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepo {
	return &photoRepo{db}
}

func (r *photoRepo) CreatePhotoRepository(Id int, photo *models.Photo) (*models.Photo, error) {
	photo.UserId = Id
	err := r.db.Create(&photo)
	if err.Error != nil {
		return nil, err.Error
	}
	return photo, nil
}

func (r *photoRepo) GetAllPhotoRepository() (*[]models.PhotoResponse, error) {
	photos := []models.PhotoResponse{}
	if err := r.db.Joins("User").Model(&models.Photo{}).Find(&photos).Error; err != nil {
		log.Printf("Error get data with err: %s", err)
		return nil, err
	}
	return &photos, nil
}

func (r *photoRepo) UpdatePhotoRepository(authId int, photo *models.Photo) (*models.Photo, error) {
	photoCheckAvail := models.Photo{}
	err := r.db.Where("user_id = ?", authId).
		Where("id = ?", photo.Id).
		First(&photoCheckAvail)
	if err.Error != nil {
		log.Printf("Error get data photo detail with err: %s", err.Error)
		return nil, err.Error
	}
	err = r.db.Model(&photoCheckAvail).Updates(photo)
	if err.Error != nil {
		log.Printf("Error update data photo id with err: %s", err.Error)
		return nil, err.Error
	}
	return &photoCheckAvail, nil
}

func (r *photoRepo) DeletePhotoRepository(photoId, authId int) error {
	photoCheckAvail := models.Photo{}
	err := r.db.Where("user_id = ?", authId).
		Where("id = ?", photoId).
		First(&photoCheckAvail)
	if err.Error != nil {
		log.Printf("Error delete photo detail with err: %s", err.Error)
		return err.Error
	}
	r.db.Delete(photoCheckAvail, photoId)
	return nil
}
