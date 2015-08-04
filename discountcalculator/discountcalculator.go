package discountcalculator

type discountCalculator struct {
	strategy     func() (rate float64, strategyCode int)
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

func (calculator *discountCalculator) RateWithCouponFor(customer *customer, couponType int) float64 {
	customerRate := calculator.RateFor(customer)

	switch couponType {
	case BIRTHDAY_ANNIVERSARY:
		calculator.strategy = birthdayDiscount
	}

	couponRate, code := calculator.strategy()
	calculator.strategyCode = code

	decimal := 2
	return calculator.addRates(customerRate, couponRate, decimal)
}

func (calculator *discountCalculator) addRates(rate1, rate2 float64, decimal int) float64 {
	precision := 10 * float64(decimal)
	return ((rate1 * precision) + (rate2 * precision)) / precision
}
