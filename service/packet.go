package service

import (
	"context"
	"fmt"
	"myapp/model"
	"myapp/tools"
	"time"

	"gorm.io/gorm"
)

func (s *Service) CreatePacket(ctx context.Context, input model.NewPacket) (*model.Packet, error) {
	packet := model.Packet{
		Name:       input.Name,
		DataLimit:  input.DataLimit,
		Validity:   input.Validity,
		Price:      input.Price,
		ProviderID: input.ProviderID,
	}
	if err := s.DB.Model(&packet).Omit("updated_at").Create(&packet).Error; err != nil {
		panic(err)
	}
	return &packet, nil
}

func (s *Service) UpdatePacket(ctx context.Context, input model.NewPacket, id int) (*model.Packet, error) {
	var packet model.Packet
	if err := s.DB.Model(&packet).Scopes(tools.IsDeletedAtNull).Where("id = ?", id).Updates(map[string]interface{}{
		"id":          id,
		"name":        input.Name,
		"data_limit":  input.DataLimit,
		"validity":    input.Validity,
		"price":       input.Price,
		"provider_id": input.ProviderID,
		"updated_at":  time.Now().UTC(),
	}).Error; err != nil {
		panic(err)
	}
	return &packet, nil
}

func (s *Service) DeletePacket(ctx context.Context, id int) (string, error) {
	var packet model.Packet

	if err := s.DB.Model(&packet).Where("id = ", id).Omit("updated_at").Update("deleted_at", time.Now().UTC()).Error; err != nil {
		panic(err)
	}
	return "Success", nil
}

func (s *Service) PacketGetAll(ctx context.Context) ([]*model.Packet, error) {
	var packet []*model.Packet
	if err := s.DB.Model(&packet).Scopes(tools.IsDeletedAtNull).Find(&packet).Error; err != nil {
		return nil, err
	}
	return packet, nil
}

func (s *Service) PacketGetById(ctx context.Context, id int) (*model.Packet, error) {
	var packet model.Packet
	if err := s.DB.Model(&packet).Scopes(tools.IsDeletedAtNull).Where("id = ?", id).First(&packet).Error; err == gorm.ErrRecordNotFound {
		panic(fmt.Errorf("packet not found?"))
	} else if err != nil {
		panic(err)
	}
	return &packet, nil
}

func (s *Service) PacketGetByProviderId(ctx context.Context, providerID int) ([]*model.Packet, error) {
	var packets []*model.Packet
	if err := s.DB.Where("provider_id = ?", providerID).Find(&packets).Error; err != nil {
		return nil, err
	}
	return packets, nil
}
func (s *Service) PacketGetByUserId(ctx context.Context, id int) ([]*model.Packet, error) {
	var (
		userPacket []*model.UserPacket
		packets    []*model.Packet
		packetId   []int
	)
	if err := s.DB.Model(&userPacket).Where("user_id = ?", id).Find(&userPacket).Error; err != nil {
		panic(err)
	}
	for _, val := range userPacket {
		packetId = append(packetId, val.PacketID)
	}
	if err := s.DB.Model(&packets).Scopes(tools.IsDeletedAtNull).Where("id IN ?", packetId).Find(&packets).Error; err != nil {
		panic(err)
	}
	return packets, nil
}
