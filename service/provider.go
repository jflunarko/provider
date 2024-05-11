package service

import (
	"context"
	"fmt"
	"myapp/model"
	"myapp/tools"
	"time"

	"gorm.io/gorm"
)

func (s *Service) CreateProvider(ctx context.Context, input model.NewProvider) (*model.Provider, error) {
	provider := model.Provider{
		Provider:  input.Provider,
		Location:  input.Location,
		CreatedAt: time.Now().UTC(),
	}
	if err := s.DB.Model(&provider).Omit("updated_at").Create(&provider).Error; err != nil {
		panic(err)
	}

	return &provider, nil
}

func (s *Service) UpdateProvider(ctx context.Context, input model.UpdateProvider, id int) (*model.Provider, error) {
	var provider model.Provider
	if err := s.DB.Model(&provider).Scopes(tools.IsDeletedAtNull).Where("id = ?", id).Updates(map[string]interface{}{
		"provider":   input.Provider,
		"location":   input.Location,
		"updated_at": time.Now().UTC(),
	}).Error; err != nil {
		panic(err)
	}
	return &provider, nil
}

func (s *Service) DeleteProvider(ctx context.Context, id int) (string, error) {
	var provider model.Provider

	if err := s.DB.Model(&provider).Scopes(tools.IsDeletedAtNull).Where("id = ?", id).Omit("updated_at").Update("deleted_at", time.Now().UTC()).Error; err != nil {
		panic(err)
	}
	return "success", nil
}

func (s *Service) ProvidersGetAll(c context.Context) ([]*model.Provider, error) {
	var provider []*model.Provider
	if err := s.DB.Model(&provider).Scopes(tools.IsDeletedAtNull).Find(&provider).Error; err != nil {
		return nil, err
	}
	return provider, nil
}

func (s *Service) ProviderGetById(ctx context.Context, id int) (*model.Provider, error) {
	var provider model.Provider

	if err := s.DB.Model(&provider).Scopes(tools.IsDeletedAtNull).Where("id = ?", id).First(&provider).Error; err == gorm.ErrRecordNotFound {
		panic(fmt.Errorf("provider not found?"))
	} else if err != nil {
		panic(err)
	}
	return &provider, nil
}

func (s *Service) ProviderGetByUserId(ctx context.Context, id int) ([]*model.Provider, error) {
	var (
		userProvider []*model.UserProvider
		providers    []*model.Provider
		providerId   []int
	)
	if err := s.DB.Model(&userProvider).Where("user_id = ?", id).Find(&userProvider).Error; err != nil {
		panic(err)
	}
	for _, val := range userProvider {
		providerId = append(providerId, val.ProviderID)
	}
	if err := s.DB.Model(&providers).Scopes(tools.IsDeletedAtNull).Where("id = ?", providerId).Find(&providers).Error; err != nil {
		panic(err)
	}
	return providers, nil
}
