package repositories

import (
	"final_MyGram/models"
	"log"

	"gorm.io/gorm"
)

type AuthRepo interface {
	RegisterRepository(user *models.User) error
	LoginRepository(user *models.UserLogin) (*models.UserLogin, error)
	GetIdByEmail(email string) (error, int)
}

type authRepo struct {
	db *gorm.DB
}

func NewauthRepository(db *gorm.DB) AuthRepo {
	return &authRepo{db}
}

func (r *authRepo) RegisterRepository(user *models.User) error {
	err := r.db.Create(&user)
	if err.Error != nil {
		return err.Error
	}
	r.db.Last(&user)
	return nil
}

func (r *authRepo) LoginRepository(user *models.UserLogin) (*models.UserLogin, error) {
	var userLogin models.UserLogin
	if err := r.db.Table("users").Where("users.email = ?", user.Email).First(&userLogin).Error; err != nil {
		log.Printf("Error get data user with err: %s", err)
		return nil, err
	}
	return &userLogin, nil
}

func (r *authRepo) GetIdByEmail(email string) (error, int) {
	user := models.User{}

	err := r.db.Select("id").Where("email = ?", email).First(&user)
	if err.Error != nil {
		log.Printf("Error get data user email with err: %s", err.Error)
		return err.Error, -1
	}
	return nil, user.Id
}
