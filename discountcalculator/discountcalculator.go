package discountcalculator

type discountCalculator struct {
	strategy     func(invoiceTotal float64, d *discount) (balance float64)
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

func (calculator *discountCalculator) ComputeBalance(customer *customer, invoiceTotal float64) float64 {
	discount := calculator.DiscountFor(customer)

	switch c := customer.category; c {
	case STANDARD:
		calculator.strategy = standardCalculation
		calculator.strategyCode = STANDARD_CALCULATION
	case SILVER:
		calculator.strategy = silverCalculation
		calculator.strategyCode = SILVER_CALCULATION
	case GOLD:
		calculator.strategy = goldCalculation
		calculator.strategyCode = GOLD_CALCULATION
	case PREMIUM:
		calculator.strategy = premiumCalculation
		calculator.strategyCode = PREMIUM_CALCULATION
	}

	return calculator.strategy(invoiceTotal, discount)
}

const (
	STANDARD_CALCULATION = iota + 1
	SILVER_CALCULATION
	GOLD_CALCULATION
	PREMIUM_CALCULATION
)

func standardCalculation(invoiceTotal float64, d *discount) float64 {
	return 0.0
}

func silverCalculation(invoiceTotal float64, d *discount) float64 {
	return 0.0
}

func goldCalculation(invoiceTotal float64, d *discount) float64 {
	return 0.0
}

func premiumCalculation(invoiceTotal float64, d *discount) float64 {
	return 0.0
}
