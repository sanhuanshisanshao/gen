// Code generated by github.com/sanhuanshisanshao/gen. DO NOT EDIT.
// Code generated by github.com/sanhuanshisanshao/gen. DO NOT EDIT.
// Code generated by github.com/sanhuanshisanshao/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCreditCard = "credit_cards"

// CreditCard mapped from table <credit_cards>
type CreditCard struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Number        string         `gorm:"column:number" json:"number"`
	CustomerRefer int64          `gorm:"column:customer_refer" json:"customer_refer"`
	BankID        int64          `gorm:"column:bank_id" json:"bank_id"`
}

// TableName CreditCard's table name
func (*CreditCard) TableName() string {
	return TableNameCreditCard
}
