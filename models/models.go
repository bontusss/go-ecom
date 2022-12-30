package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// A model representing a User object.
type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name     string             `json:"first_name" validate:"required, min=2, max=30"`
	Last_Name      string             `json:"last_name" validate:"required, min=2, max=30"`
	Password       string             `json:"password" validate:"required, min=6"`
	Email          string             `json:"email" validate:"email, required"`
	Phone          string             `json:"phone_number" validate:"required"`
	Token          string             `json:"token"`
	Refresh_Token  string             `json:"refresh_token"`
	Created_at     time.Time          `json:"created_at"`
	Updated_at     time.Time          `json:"updated_at"`
	User_ID        string             `json:"user_id"`
	UserCart       []ProductUser      `json:"user_cart" bson:"usercart"` // Use bson here because we are referencing another struct
	Address_Detail []Address          `json:"address" bson:"address"`
	Order_Status   []Order            `json:"order_status" bson:"orders"`
}

// Model representing a single product in the store.
type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name string             `json:"product_name"`
	Price        uint64             `json:"price"`
	Rating       uint8              `json:"rating"`
	Image        string             `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name string             `json:"product_name"`
	Price        uint64             `json:"price"`
	Rating       uint8              `json:"rating"`
	Image        string             `json:"image"`
}

// A model representing the Address Object
type Address struct {
	Address_ID  primitive.ObjectID `bson:"_id"`
	House       string             `json:"house"`
	Street      string             `json:"street"`
	City        string             `json:"city"`
	Postal_Code string             `json:"postal_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_list" bson:"order_list"`
	Ordered_at     time.Time          `json:"ordered_at"`
	Price          uint64             `json:"price"`
	Discount       int                `json:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment"`
}

type Payment struct {
	// Digital payment method
	Digital bool `json:"digital"`

	// Cash on delivery
	COD bool `json:"cod"`
}
