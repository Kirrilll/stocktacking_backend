package item_entity

type Item struct {
	ID          int    `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	PhotoLink   string `gorm:"column:photo_link"`
	StorageID   int    `gorm:"column:storage_id"`
	UserID      int    `gorm:"column:user_id"`
	Comment     string `gorm:"column:comment"`
	Status      string `gorm:"column:status"`
	CategoryID  int    `gorm:"column:category_id"`
	BranchID    int    `gorm:"column:branch_id"`
}

func (m *Item) TableName() string {
	return "items"
}
