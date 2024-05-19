package branch_entity

type Branch struct {
	ID             int    `gorm:"column:id"`
	Name           string `gorm:"column:name"`
	OrganizationID int    `gorm:"column:organization_id"`
	Address        string `gorm:"column:address"`
	Lat            string `gorm:"column:lat"`
	Lon            string `gorm:"column:lon"`
}

func (m *Branch) TableName() string {
	return "branches"
}
