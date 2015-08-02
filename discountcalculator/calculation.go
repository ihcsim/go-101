package discountcalculator

const (
	STANDARD_DISCOUNT = iota
	SILVER_DISCOUNT
	GOLD_DISCOUNT
	PREMIUM_DISCOUNT
)

const (
	STANDARD_DISCOUNT_RATE = 0.1
	SILVER_DISCOUNT_RATE   = 0.15
	GOLD_DISCOUNT_RATE     = 0.20
	PREMIUM_DISCOUNT_RATE  = 0.25
)

func standardDiscount() (rate float64, strategyCode int) {
	rate = STANDARD_DISCOUNT_RATE
	strategyCode = STANDARD_DISCOUNT
	return
}

func silverDiscount() (rate float64, strategyCode int) {
	rate = SILVER_DISCOUNT_RATE
	strategyCode = SILVER_DISCOUNT
	return
}

func goldDiscount() (rate float64, strategyCode int) {
	rate = GOLD_DISCOUNT_RATE
	strategyCode = GOLD_DISCOUNT
	return
}

func premiumDiscount() (rate float64, strategyCode int) {
	rate = PREMIUM_DISCOUNT_RATE
	strategyCode = PREMIUM_DISCOUNT
	return
}
