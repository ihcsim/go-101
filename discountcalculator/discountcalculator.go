package discountcalculator

type discountCalculator struct {
	strategy func() (rate float64, strategyCode int)
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
	customerDiscount := calculator.DiscountFor(customer)

	var d *discount
	switch couponType {
	case BIRTHDAY_ANNIVERSARY:
		d = NewDiscount(BIRTHDAY_DISCOUNT)
	}

	decimal := 2
	d.addRates(customerDiscount, decimal)
	return d
}
