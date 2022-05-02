package po

type Currency struct {
	ID
	Code string `gorm:"<-:create;column:code"`
	Name string `gorm:"column:name"`

	CreatedAt
	UpdatedAt
}

func (t *Currency) TableName() string {
	return "currency"
}
