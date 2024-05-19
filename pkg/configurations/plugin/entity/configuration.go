package entity

import (
	"github.com/google/uuid"
)

// Configuration сущность конфигурации
type Configuration struct {
	Id   uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Code string    `json:"code"`
	Name string    `json:"name"`
}

type ConfigurationsList struct {
	Total int             `json:"total"`
	Items []Configuration `json:"items"`
}

func (Configuration) TableName() string {
	return "configurations"
}
