package model

import (
	"time"

	"github.com/cao7113/hellogolang/config"

	"github.com/jinzhu/gorm"
)

type Guser struct {
	gorm.Model
	Name     string
	Email    string
	Birthday *time.Time
}

func init() {
	Migrate(config.Conn)
}

// Migrate create table and add indexes
func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&Guser{})
	db.Model(&Guser{}).AddUniqueIndex("idx_email", "email")
}

// TableName tells the table name of withdrawal_logs
func (Guser) TableName() string {
	return "gusers"
}
