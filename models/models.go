package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// A model representing a User object.
type User struct {
	ID             primitive.ObjectID
	First_Name     string
	Last_Name      string
	Password       string
	Email          string
	Phone          string
	Token          string
	Refresh_Token  string
	Created_at     time.Time
	Updated_at     time.Time
	User_ID        string
	UserCart       []ProductUser
	Address_Detail []Address
	Order_Status   []Order
}

// Model representing a single product in the store.
type Product struct {
	Product_ID   primitive.ObjectID
	Product_Name string
	Price        float64
	Rating       uint8
	Image        string
}

type ProductUser struct {
	Product_ID   primitive.ObjectID
	Product_Name string
	Price        uint64
	Rating       uint8
	Image        string
}

// A model representing the Address Object
type Address struct {
	Address_ID  primitive.ObjectID
	House       string
	Street      string
	City        string
	Postal_Code string
}

type Order struct {
	Order_ID       primitive.ObjectID
	Order_Cart     []ProductUser
	Ordered_at     time.Time
	Price          uint64
	Discount       int
	Payment_Method Payment
}

type Payment struct {
	// Digital payment method
	Digital bool

	// Cash on delivery
	COD bool
}
