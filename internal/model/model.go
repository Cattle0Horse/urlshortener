package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Model) CreateTime() int64 {
	return m.CreatedAt.UnixMilli()
}

func (m *Model) UpdateTime() int64 {
	return m.UpdatedAt.UnixMilli()
}
