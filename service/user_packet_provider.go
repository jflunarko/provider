package service

import (
	"context"
	"myapp/model"
)

func (s *Service) BuyPacket(ctx context.Context, packetId int, userId int) (string, error) {
	var count int64

	err := s.DB.Model(&model.UserPacket{}).Where("user_id = ? AND packet_id = ?", userId, packetId).Count(&count).Error
	if err != nil {
		return "Error", nil
	}
	if count > 0 {
		return "You have already purchased this packet, Please choose another packet", nil
	}
	userPacket := &model.UserPacket{
		UserID:   userId,
		PacketID: packetId,
	}
	if err := s.DB.Model(&model.UserPacket{}).Create(&userPacket).Error; err != nil {
		panic(err)
	}
	return "Your Purchase Has Been Confirmed", nil
}

func (s *Service) CancelPacket(ctx context.Context, packetId int, userId int) (string, error) {
	var userPacket model.UserPacket

	result := s.DB.Model(&userPacket).Where("packet_id = ? AND user_id = ?", packetId, userId).Delete(&userPacket)
	if result.Error != nil {
		panic(result)
	}
	if result.RowsAffected == 0 {
		return "Invalid User!", nil
	}
	return "Your purchase has been Cancelled", nil
}

func (s *Service) ChooseProvider(ctx context.Context, providerId int, userId int) (string, error) {
	userProvider := &model.UserProvider{
		UserID:     userId,
		ProviderID: providerId,
	}
	if err := s.DB.Model(&model.UserProvider{}).Create(&userProvider).Error; err != nil {
		panic(err)
	}
	return "Success", nil
}

func (s *Service) CancelProvider(ctx context.Context, providerId int, userId int) (string, error) {
	var userProvider model.UserProvider

	result := s.DB.Model(&userProvider).Where("provider_id = ? AND user_id = ?", providerId, userId).Delete(&userProvider)
	if result.Error != nil {
		panic(result)
	}
	if result.RowsAffected == 0 {
		return "invalid user!", nil
	}
	return "Your provider has been Cancelled", nil
}
