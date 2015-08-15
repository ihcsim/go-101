package discountcalculator

type discountCalculator struct {
	strategy func(invoiceTotal float64, d *discount) (balance float64, code checkoutCode)
}

func New() *discountCalculator {
	return &discountCalculator{}
}

func (calculator *discountCalculator) Discount(customer *customer) *discount {
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

func (calculator *discountCalculator) SpecialDiscount(customer *customer, couponType int) *discount {
	customerDiscount := calculator.Discount(customer)

	var d *discount
	switch couponType {
	case BIRTHDAY_ANNIVERSARY:
		d = NewDiscount(BIRTHDAY_DISCOUNT)
	}

	decimal := 2
	d.addRates(customerDiscount, decimal)
	return d
}

func (calculator *discountCalculator) Checkout(customer *customer, invoiceTotal float64) (balance float64, code checkoutCode) {
	discount := calculator.Discount(customer)

	switch c := customer.category; c {
	case STANDARD:
		calculator.strategy = standardCheckout
	case SILVER:
		calculator.strategy = standardCheckout
	case GOLD:
		calculator.strategy = expressCheckout
	case PREMIUM:
		calculator.strategy = expressCheckout
	}

	return calculator.strategy(invoiceTotal, discount)
}

func (calculator *discountCalculator) CheckoutWithSpecialDiscount(customer *customer, couponType int, invoiceTotal float64) (balance float64, code checkoutCode) {
	discount := calculator.SpecialDiscount(customer, couponType)

	switch c := customer.category; c {
	case STANDARD:
		calculator.strategy = standardCheckout
	case SILVER:
		calculator.strategy = standardCheckout
	case GOLD:
		calculator.strategy = expressCheckout
	case PREMIUM:
		calculator.strategy = expressCheckout
	}

	return calculator.strategy(invoiceTotal, discount)
}
