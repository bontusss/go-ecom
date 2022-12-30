package models

// A model representing a User object.
type User struct {
	ID
	First_Name
	Last_Name
	Password
	Email
	Phone
	Token
	Refresh_Token
	Created_at
	Updated_at
	User_ID
	UserCart
	Address_Detail
	Order_Status
}

// Model representing a single product in the store.
type Product struct {
	Product_ID
	Product_Name
	Price
	Rating
	Image
}

type ProductUser struct {
	Product_ID
	Product_Name
	Price
	Rating
	Image
}

// A model representing the Address Object
type Address struct {
	Address_ID
	House
	Street
	City
	Postal_Code
}

type Order struct{}

type Payment struct{}
