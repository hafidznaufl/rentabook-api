package schema

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	BookID    uint           `gorm:"index"`
	Book      Book           `gorm:"foreignKey:BookID"`
	UserID    uint           `gorm:"index"`
	User      User           `gorm:"foreignKey:UserID"`
	Date      time.Time      `json:"date"`
	Return    time.Time      `json:"return"`
	Status    string         `json:"status"`
}
