package discountcalculator

const (
	STANDARD_DISCOUNT = iota
	SILVER_DISCOUNT
	GOLD_DISCOUNT
	PREMIUM_DISCOUNT
)

func standardDiscount() int {
	return STANDARD_DISCOUNT
}

func silverDiscount() int {
	return SILVER_DISCOUNT
}

func goldDiscount() int {
	return GOLD_DISCOUNT
}

func premiumDiscount() int {
	return PREMIUM_DISCOUNT
}
