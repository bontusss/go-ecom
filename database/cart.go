package database

import "errors"

var (
	ErrCantFindProduct        = errors.New("Cannot find product")
	ErrCantDecodeProduct      = errors.New("Cannot find product")
	ErrUserIdInvalid          = errors.New("User ID invalid")
	ErrCantUpdateUser         = errors.New("Cannot update user")
	ErrCantRemoveItemFromcart = errors.New("Cannot remove item from cart")
	ErrCantGetItem            = errors.New("Cannot get item")
	ErrCantBuyCartItem        = errors.New("Cannot update the purchase")
)

func AddProductTocart() {}

func RemoveProductFromcart() {

}

func BuyFromCart() {

}

func InstantBuy() {

}
