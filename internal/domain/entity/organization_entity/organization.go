package organization_entity

type Organization struct {
	ID    int    `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Phone string `gorm:"column:phone"`
	Inn   string `gorm:"column:inn"`
}

func (m *Organization) TableName() string {
	return "organizations"
}
