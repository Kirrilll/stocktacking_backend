package report

type Report struct {
	ID           int    `gorm:"column:id"`
	CreatedAt    string `gorm:"column:created_at"`
	ItemID       int    `gorm:"column:item_id"`
	ClosedAt     string `gorm:"column:closed_at"`
	SatusStart   string `gorm:"column:satus_start"`
	CommentStart string `gorm:"column:comment_start"`
	UserStartID  int    `gorm:"column:user_start_id"`
	StatusEnd    string `gorm:"column:status_end"`
	CommentEnd   string `gorm:"column:comment_end"`
	UserEndID    int    `gorm:"column:user_end_id"`
	StorageID    int    `gorm:"column:storage_id"`
}

func (m *Report) TableName() string {
	return "reports"
}
