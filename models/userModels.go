package models

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	User_Name string    `json:"username" valid:"required" gorm:"unique" validate:"required" example:"zayyan"`
	Email     string    `json:"email" valid:"required,email" gorm:"unique" validate:"required" example:"zayyan@mail.com"`
	Password  string    `json:"password,omitempty" valid:"required" validate:"required" example:"zayyan"`
	Age       int       `json:"age" valid:"required" validate:"required" example:"1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegisterRequest struct {
	Age      int    `json:"age" valid:"required" validate:"required,numeric,gt=8" example:"9"`
	Email    string `json:"email" valid:"required,email" validate:"required,email" example:"zayyan@mail.com"`
	Password string `json:"password,omitempty" valid:"required" validate:"required,min=6" example:"zayyan"`
	UserName string `json:"username" valid:"required" validate:"required" example:"zayyan"`
}

func RegisterToModel(req *UserRegisterRequest) (*User, error) {
	user := &User{
		Age:       req.Age,
		Email:     req.Email,
		Password:  req.Password,
		User_Name: req.UserName,
	}

	return user, nil
}

type UserRegisterResponse struct {
	Age      int    `json:"age" valid:"required" validate:"required" example:"1"`
	Email    string `json:"email" valid:"required,email" validate:"required" example:"zayyan@mail.com"`
	Id       int    `json:"id" example:"1"`
	UserName string `json:"username" valid:"required" validate:"required" example:"zayyan"`
}

func RegisterToResponse(req *User) (*UserRegisterResponse, error) {
	res := &UserRegisterResponse{
		Id:       req.Id,
		Age:      req.Age,
		Email:    req.Email,
		UserName: req.User_Name,
	}

	return res, nil
}

type UserLogin struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" valid:"required,email" gorm:"unique" validate:"required"`
	Password string `json:"password,omitempty" valid:"required" validate:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" valid:"required,email" gorm:"unique" validate:"required" example:"zayyan@mail.com"`
	Password string `json:"password,omitempty" valid:"required" validate:"required" example:"zayyan"`
}

func LoginToModel(req *UserLoginRequest) (*UserLogin, error) {
	user := &UserLogin{
		Email:    req.Email,
		Password: req.Password,
	}

	return user, nil
}

type UserUpdate struct {
	Email    string `json:"email" gorm:"unique" validate:"required" example:"zayyan@mail.com"`
	UserName string `json:"username" validate:"required" example:"zayyan"`
}

type UserUpdateResponse struct {
	Id        int       `json:"id" example:"1"`
	Email     string    `json:"email" valid:"required,email" validate:"required" example:"zayyan@mail.com"`
	Age       int       `json:"age" valid:"required" validate:"required" example:"1"`
	UserName  string    `json:"username" valid:"required" validate:"required" example:"zayyan"`
	CreatedAt time.Time `json:"created_at"`
}

func UpdateToResponse(data *User) (*UserUpdateResponse, error) {
	user := &UserUpdateResponse{
		Age:       data.Age,
		Email:     data.Email,
		Id:        data.Id,
		UserName:  data.User_Name,
		CreatedAt: data.CreatedAt,
	}

	return user, nil
}
