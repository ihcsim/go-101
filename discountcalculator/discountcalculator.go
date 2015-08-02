package discountcalculator

type discountCalculator struct {
	strategy     func() (float64, int)
	strategyCode int
}

func New() *discountCalculator {
	return &discountCalculator{}
}

func (calculator *discountCalculator) RateFor(customer *customer) float64 {
	switch c := customer.category; c {
	case STANDARD:
		calculator.strategy = standardDiscount
	case SILVER:
		calculator.strategy = silverDiscount
	case GOLD:
		calculator.strategy = goldDiscount
	case PREMIUM:
		calculator.strategy = premiumDiscount
	}

	rate, code := calculator.strategy()
	calculator.strategyCode = code
	return rate
}
