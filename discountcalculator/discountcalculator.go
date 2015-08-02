package discountcalculator

type discountCalculator struct {
	strategy func() int
}

func New() *discountCalculator {
	return &discountCalculator{}
}

func (calculator *discountCalculator) CalcFor(customer *customer) {
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
}
