package model

import (
	"time"
)

type Currency struct {
	ID          int64     `gorm:"<-:create;column:id;type:bigint unsigned AUTO_INCREMENT;primaryKey;NOT NULL;comment:primary key"`
	Code        string    `gorm:"type:varchar(45)"`
	Name        string    `gorm:"type:varchar(45)"`
	CreatedAt   time.Time `gorm:"column:created_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6);NOT NULL;comment:資料產生時間點"`
	UpdatedAtAt time.Time `gorm:"column:updated_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:資料最後修改時點"`
}

func (s *Currency) TableName() string {
	return "currency"
}

type Member struct {
	ID        int64     `gorm:"<-:create;column:id;type:bigint unsigned AUTO_INCREMENT;primaryKey;NOT NULL;comment:primary key"`
	Login     string    `gorm:"type:varchar(45)"`
	PassWord  string    `gorm:"type:varchar(45)"`
	Name      string    `gorm:"type:varchar(45)"`
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6);NOT NULL;comment:資料產生時間點"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:資料最後修改時點"`
}

func (s *Member) TableName() string {
	return "member"
}

type Wallet struct {
	ID           int64     `gorm:"<-:create;column:id;type:bigint unsigned AUTO_INCREMENT;primaryKey;NOT NULL;comment:primary key"`
	MemberLogin  string    `gorm:"type:varchar(45)"`
	CurrencyCode string    `gorm:"type:varchar(45)"`
	Balance      float64   `gorm:"type:DECIMAL(20,8) unsigned;NOT NULL"`
	CreatedAt    time.Time `gorm:"column:created_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6);NOT NULL;comment:資料產生時間點"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:資料最後修改時點"`
}

func (s *Wallet) TableName() string {
	return "wallet"
}

type Withdraw struct {
	ID             int64     `gorm:"<-:create;column:id;type:bigint unsigned AUTO_INCREMENT;primaryKey;NOT NULL;comment:primary key"`
	WalletID       int64     `gorm:"type:bigint unsigned;NOT NULL"`
	TransferAmount float64   `gorm:"type:DECIMAL(20,8) unsigned;NOT NULL"`
	CreatedAt      time.Time `gorm:"column:created_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6);NOT NULL;comment:資料產生時間點"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:資料最後修改時點"`
}

func (s *Withdraw) TableName() string {
	return "withdraw"
}

type Deposit struct {
	ID             int64     `gorm:"<-:create;column:id;type:bigint unsigned AUTO_INCREMENT;primaryKey;NOT NULL;comment:primary key"`
	WalletID       int64     `gorm:"type:bigint unsigned;NOT NULL"`
	TransferAmount float64   `gorm:"type:DECIMAL(20,8) unsigned;NOT NULL"`
	CreatedAt      time.Time `gorm:"column:created_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6);NOT NULL;comment:資料產生時間點"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:資料最後修改時點"`
}

func (s *Deposit) TableName() string {
	return "deposit"
}

type Transfer struct {
	ID             int64     `gorm:"<-:create;column:id;type:bigint unsigned AUTO_INCREMENT;primaryKey;NOT NULL;comment:primary key"`
	FromWalletID   int64     `gorm:"type:bigint unsigned;NOT NULL"`
	TOWalletID     int64     `gorm:"type:bigint unsigned;NOT NULL"`
	TransferAmount float64   `gorm:"type:DECIMAL(20,8) unsigned;NOT NULL"`
	CreatedAt      time.Time `gorm:"column:created_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6);NOT NULL;comment:資料產生時間點"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:資料最後修改時點"`
}

func (s *Transfer) TableName() string {
	return "transfer"
}

type Transaction struct {
	ID             int64     `gorm:"<-:create;column:id;type:bigint unsigned AUTO_INCREMENT;primaryKey;NOT NULL;comment:primary key"`
	WalletID       int64     `gorm:"type:varchar(45)"`
	Method         string    `gorm:"type:varchar(45)"`
	ActionID       int64     `gorm:"column:action_id;type:bigint unsigned;NOT NULL"`
	TransferBefore float64   `gorm:"type:DECIMAL(20,8) unsigned;NOT NULL"`
	TransferAmount float64   `gorm:"type:DECIMAL(20,8) unsigned;NOT NULL"`
	TransferAfter  float64   `gorm:"type:DECIMAL(20,8) unsigned;NOT NULL"`
	CreatedAt      time.Time `gorm:"column:created_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6);NOT NULL;comment:資料產生時間點"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:DATETIME(6);default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:資料最後修改時點"`
}

func (s *Transaction) TableName() string {
	return "transaction"
}
