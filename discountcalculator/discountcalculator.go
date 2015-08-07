package discountcalculator

type discountCalculator struct {
	strategy     func() (rate float64, strategyCode int)
	strategyCode int
}

func New() *discountCalculator {
	return &discountCalculator{}
}

func (calculator *discountCalculator) DiscountFor(customer *customer) *discount {
	switch c := customer.category; c {
	case STANDARD:
		return NewDiscount(STANDARD_DISCOUNT)
	case SILVER:
		return NewDiscount(SILVER_DISCOUNT)
	case GOLD:
		return NewDiscount(GOLD_DISCOUNT)
	case PREMIUM:
		return NewDiscount(PREMIUM_DISCOUNT)
	}

	return nil
}

func (calculator *discountCalculator) SpecialDiscountFor(customer *customer, couponType int) *discount {
	customerRate := calculator.DiscountFor(customer)

	var d *discount
	switch couponType {
	case BIRTHDAY_ANNIVERSARY:
		d = NewDiscount(BIRTHDAY_DISCOUNT)
	}

	d.rate = calculator.addRates(customerRate.rate, d.rate, 2)
	return d
}

func (calculator *discountCalculator) addRates(rate1, rate2 float64, decimal int) float64 {
	precision := 10 * float64(decimal)
	return ((rate1 * precision) + (rate2 * precision)) / precision
}
