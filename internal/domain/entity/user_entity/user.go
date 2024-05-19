package user_entity

type User struct {
	ID             int    `gorm:"column:id"`
	Name           string `gorm:"column:name"`
	Role           string `gorm:"column:role"`
	Phone          string `gorm:"column:phone"`
	OrganizationID int    `gorm:"column:organization_id"`
}

func (m *User) TableName() string {
	return "users"
}
