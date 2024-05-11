package service

import (
	"context"
	"fmt"
	"myapp/model"
	"myapp/tools"
	"time"

	"gorm.io/gorm"
)

func (s *Service) RegisterUser(ctx context.Context, input model.NewUser) (*model.UserToken, error) {
	password, err := tools.HashAndSalt(input.Password)
	if err != nil {
		panic(err)
	}
	user := model.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}
	if err := s.DB.Model(&user).Omit("updated_at").Create(&user).Error; err != nil {
		panic(err)
	}
	token := tools.TokenCreate(user.ID)

	return &model.UserToken{
		Message: "Success",
		Token:   token,
	}, nil
}

func (s *Service) LoginUser(ctx context.Context, input model.UserLogin) (*model.UserToken, error) {
	var user *model.User

	if err := s.DB.Model(&user).Scopes(tools.IsDeletedAtNull).Where("email LIKE ?", input.Email).First(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	if err := tools.ComparePasswords(user.Password, input.Password); err != nil {
		return nil, err
	}
	token := tools.TokenCreate(user.ID)

	return &model.UserToken{
		Message: "Success",
		Token:   token,
	}, nil
}

func (s *Service) UpdateUser(ctx context.Context, input model.UpdateUser) (string, error) {
	var user *model.User
	if err := s.DB.Model(&user).Scopes(tools.IsDeletedAtNull).Where("email LIKE ?", input.Email).First(&user).Update("updated_at", time.Now().UTC()).Error; err == gorm.ErrRecordNotFound {
		return "Email Not Found", nil
	} else if err != nil {
		return "Failed", nil
	}
	hashedPassword, err := tools.HashAndSalt(input.NewPassword)
	if err != nil {
		return "Failed", err
	}
	if err := s.DB.Model(&user).Scopes(tools.IsDeletedAtNull).Update("password", hashedPassword).Error; err != nil {
		return "Failed", nil
	}

	return "Success, Password Updated", nil
}

func (s *Service) DeleteUser(ctx context.Context, id int) (string, error) {
	var user *model.User
	if err := s.DB.Model(&user).Scopes(tools.IsDeletedAtNull).Where("id = ?", id).Omit("updated_at").Update("deleted_at", time.Now().UTC()).Error; err != nil {
		panic(err)
	}
	return "Delete User Success", nil
}

func (s *Service) UserGetAll(ctx context.Context) ([]*model.User, error) {
	var (
		user []*model.User
	)
	if err := s.DB.Model(&user).Scopes(tools.IsDeletedAtNull).Find(&user).Error; err != nil {
		panic(err)
	}
	return user, nil
}

func (s *Service) UserGetByID(ctx context.Context, id int) (*model.User, error) {
	var (
		user model.User
	)
	if err := s.DB.Model(&user).Scopes(tools.IsDeletedAtNull).Where("id = ?", id).First(&user).Error; err == gorm.ErrRecordNotFound {
		panic(fmt.Errorf("user not found"))
	} else if err != nil {
		panic(err)
	}
	return &user, nil
}
