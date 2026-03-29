package models

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(255);not null;index"`
	Price       float64 `gorm:"type:numeric(10,2);not null"`
	Description string  `gorm:"type:text"`
	Stock       int     `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// relations
	// CartItems  []CartItem
	// OrderItems []OrderItem
}

type Cart struct {
	ID        uint
	UserID    uint
	Items     []CartItem `gorm:"foreignKey:CartID"`
	CreatedAt time.Time
}

type CartItem struct {
	ID        uint
	CartID    uint
	ProductID uint
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductID"`
}

// type OrderItem struct {
// 	ID        uint    `gorm:"primaryKey"`
// 	OrderID   uint    `gorm:"not null"`
// 	ProductID uint    `gorm:"not null"`
// 	Price     float64 `gorm:"type:numeric(10,2);not null"` // snapshot
// 	Quantity  int     `gorm:"not null"`

// 	// relations
// 	Product Product
// }

// type User struct {
// 	ID        uint   `gorm:"primaryKey"`
// 	Email     string `gorm:"type:varchar(255);uniqueIndex;not null"`
// 	Password  string `gorm:"type:text;not null"`
// 	Name      string `gorm:"type:varchar(255)"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time

// 	// relations
// 	Carts  []CartItem
// 	Orders []Order
// }

// type Order struct {
// 	ID         uint    `gorm:"primaryKey"`
// 	UserID     uint    `gorm:"not null"`
// 	TotalPrice float64 `gorm:"type:numeric(10,2)"`
// 	Status     string  `gorm:"type:varchar(50);default:'pending'"`
// 	CreatedAt  time.Time

// 	// relations
// 	User       User
// 	OrderItems []OrderItem `gorm:"constraint:OnDelete:CASCADE"`
// }
