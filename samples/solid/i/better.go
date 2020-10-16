package main

// Do instead

type PastaShop interface {
	OrderPasta()
}

type PizzaShop interface {
	OrderPizza()
}

type Restaurant interface {
	GetMenu()
}

type BetterItalianRestaurant interface {
	Restaurant
	PastaShop
	PizzaShop
}

type PizzaPlaceFromTheCorner interface {
	Restaurant
	PizzaShop
}
