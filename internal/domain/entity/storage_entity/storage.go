package storage_entity

type Storage struct {
	ID              int    `gorm:"column:id"`
	Name            string `gorm:"column:name"`
	ParentStorageID int    `gorm:"column:parent_storage_id"`
	BranchID        int    `gorm:"column:branch_id"`
}

func (m *Storage) TableName() string {
	return "storages"
}
