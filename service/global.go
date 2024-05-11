package service

import (
	"fmt"
	"myapp/config"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func GetService() *Service {
	s := Service{
		DB: config.GetDB(),
	}
	return &s
}
func GetTransaction() *Service { //digunakan jika insert data banyak, dan jika salah satu data error, tidak ada yang terjadi pada data yg lain
	fmt.Println("begin...")
	s := Service{
		DB: config.GetDB().Begin(),
	}

	return &s
}

func (s *Service) Rollback(err ...interface{}) error {
	s.DB.Rollback()
	fmt.Println("rollback...")

	if len(err) > 0 && err[0] != nil {
		return err[0].(error)
	}

	return nil
}
func (s *Service) Commit() error { //commit digunakan agar data masuk ke database
	if err := s.DB.Commit().Error; err != nil {
		return err
	}

	fmt.Println("commit...")

	return nil
}

func (s *Service) ErrorCheck(err ...interface{}) error {
	if len(err) > 0 && err[0] != nil {
		return err[0].(error)
	}

	return nil
}
