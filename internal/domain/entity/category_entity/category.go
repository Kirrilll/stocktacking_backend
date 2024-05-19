package category_entity

type Category struct {
	ID     int    `gorm:"column:id"`
	Name   string `gorm:"column:name"`
	ConfID int    `gorm:"column:conf_id"`
}

func (m *Category) TableName() string {
	return "categories"
}
