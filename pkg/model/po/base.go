package po

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

var (
	node *snowflake.Node
)

func init() {
	var err error
	if node, err = snowflake.NewNode(1); err != nil {
		panic(err)
	}
}

type ID struct {
	ID uint64 `gorm:"<-:create;column:id;primary_key;NOT NULL;comment:表格不重複主鍵"`
}

func (p *ID) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == 0 {
		p.ID = uint64(uint64(reverseInt64(node.Generate().Int64())))
	}
	return nil
}

func reverseInt64(n int64) int64 {
	newInt := int64(0)
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}

type CreatedAt struct {
	CreatedAt *time.Time `gorm:"<-:create;column:created_at;type:DATETIME(6);index:idx_order,sort:desc;NOT NULL;default:CURRENT_TIMESTAMP(6);comment:資料產生時間點"`
}

type UpdatedAt struct {
	UpdatedAt *time.Time `gorm:"column:updated_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:資料最後修改時點"`
}

type DeletedAt struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:DATETIME(6);comment:資料軟刪除時間點"`
}
