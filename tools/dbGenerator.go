package tools

import "gorm.io/gorm"

func IsDeletedAtNull(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at IS NULL")
}
