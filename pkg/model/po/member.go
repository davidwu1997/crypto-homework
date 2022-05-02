package po

type Member struct {
	ID
	Login    string `gorm:"<-:create;column:login"`
	PassWord string `gorm:"column:pass_word"`
	Name     string `gorm:"column:name"`

	CreatedAt
	UpdatedAt
}

func (t *Member) TableName() string {
	return "member"
}
