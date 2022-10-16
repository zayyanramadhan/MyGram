package repositories

import (
	"final_MyGram/models"
	"log"

	"gorm.io/gorm"
)

type UserRepo interface {
	UpdateUserRepository(Id int, user *models.UserUpdate) (*models.User, error)
	DeleteUserRepository(Id int) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (r *userRepo) UpdateUserRepository(Id int, user *models.UserUpdate) (*models.User, error) {
	userCheckId := models.User{}
	err := r.db.Select("id", "email", "user_name", "age", "created_at", "updated_at").First(&userCheckId, Id)
	if err.Error != nil {
		log.Printf("Error get data user id with err: %s", err.Error)
		return nil, err.Error
	}

	err = r.db.Model(&userCheckId).Updates(models.User{Email: user.Email, User_Name: user.UserName})
	if err.Error != nil {
		log.Printf("Error update data user id with err: %s", err.Error)
		return nil, err.Error
	}

	return &userCheckId, nil
}

func (r *userRepo) DeleteUserRepository(Id int) error {
	userCheckId := models.User{}
	err := r.db.Select("id", "email", "user_name", "age", "created_at", "updated_at").First(&userCheckId, Id)
	if err.Error != nil {
		log.Printf("Error get data user id with err: %s", err.Error)
		return err.Error
	}
	r.db.Delete(userCheckId, Id)
	return nil
}
