package entity

// Configuration сущность конфигурации
type Configuration struct {
	Id   int    `gorm:"type:uuid;primaryKey;" json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ConfigurationsList struct {
	Total int             `json:"total"`
	Items []Configuration `json:"items"`
}

func (Configuration) TableName() string {
	return "configurations"
}
