package model

import "time"

type User struct {
	ID           int            `json:"id" gorm:"type:int;autoIncrement"`
	Name         string         `json:"name" gorm:"type:varchar(100);not null"`
	Age          int            `json:"age" gorm:"type:int;not null"`
	Transactions []*Transaction `json:"transactions" gorm:"-"`
}

type Transaction struct {
	ID        int                 `json:"id" gorm:"type:int;autoIncrement"`
	CreatedAt time.Time           `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UserID    int                 `json:"user_id" gorm:"type:int;not null"`
	Summary   *TransactionSummary `json:"summary" gorm:"-"`

	// Attribute For Foreign Key
	User               User                `json:"user" gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDELETE:CASCADE"`
	TransactionDetails []TransactionDetail `json:"transaction_details"`
}

type TransactionDetail struct {
	ID            int      `json:"id" gorm:"type:int;autoIncrement"`
	Name          string   `json:"name" gorm:"type:varchar(100);not null"`
	Description   *string  `json:"description" gorm:"type:varchar(100);null;default:NULL"`
	Price         float64  `json:"price" gorm:"type:decimal(15,2);not null"`
	Discount      *float64 `json:"discount" gorm:"type:decimal(15,2);null:default:NULL"`
	TransactionID int      `json:"transaction_id" gorm:"type:int;not null"`

	// Attribute For Foreign Key
	Transaction Transaction `json:"transaction" gorm:"foreignKey:transaction_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
