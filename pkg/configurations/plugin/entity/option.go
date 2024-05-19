package entity

// Option сущность настройки
type Option struct {
	Id     int            `gorm:"type:uuid;primaryKey" json:"id"`
	Conf   *Configuration `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ConfId int            `gorm:"foreignKey:Id;uniqueIndex:idx_conf_id_code" json:"confId"`
	Code   string         `gorm:"uniqueIndex:idx_conf_id_code" json:"code"`
	Name   string         `json:"name"`
	Type   string         `json:"type"`
	Value  string         `json:"value"`
}

type OptionShort struct {
	Code  string `json:"code" validate:"required,min=3,max=50,sluggable"`
	Value string `json:"value" validate:"required,min=3,max=50"`
}

type OptionsList struct {
	Total int      `json:"total"`
	Items []Option `json:"items"`
}

func (Option) TableName() string {
	return "options"
}
