package discountcalculator

type discountCode int

const (
	STANDARD_DISCOUNT discountCode = iota
	SILVER_DISCOUNT
	GOLD_DISCOUNT
	PREMIUM_DISCOUNT
	BIRTHDAY_DISCOUNT
)

func (d discountCode) String() string {
	switch d {
	case STANDARD_DISCOUNT:
		return "STANDARD_DISCOUNT"
	case SILVER_DISCOUNT:
		return "SILVER_DISCOUNT"
	case GOLD_DISCOUNT:
		return "GOLD_DISCOUNT"
	case PREMIUM_DISCOUNT:
		return "PREMIUM_DISCOUNT"
	case BIRTHDAY_DISCOUNT:
		return "BIRTHDAY_DISCOUNT"
	}

	return ""
}

const (
	STANDARD_DISCOUNT_RATE = 0.1
	SILVER_DISCOUNT_RATE   = 0.15
	GOLD_DISCOUNT_RATE     = 0.20
	PREMIUM_DISCOUNT_RATE  = 0.25
	BIRTHDAY_DISCOUNT_RATE = 0.05
)

type discount struct {
	rate float64
	code discountCode
}

func NewDiscount(code discountCode) *discount {
	d := discount{}
	switch code {
	case STANDARD_DISCOUNT:
		d.standard()
	case SILVER_DISCOUNT:
		d.silver()
	case GOLD_DISCOUNT:
		d.gold()
	case PREMIUM_DISCOUNT:
		d.premium()
	case BIRTHDAY_DISCOUNT:
		d.birthdaySpecial()
	}

	return &d
}

func (d *discount) standard() {
	d.rate = STANDARD_DISCOUNT_RATE
	d.code = STANDARD_DISCOUNT
}

func (d *discount) silver() {
	d.rate = SILVER_DISCOUNT_RATE
	d.code = SILVER_DISCOUNT
}

func (d *discount) gold() {
	d.rate = GOLD_DISCOUNT_RATE
	d.code = GOLD_DISCOUNT
}

func (d *discount) premium() {
	d.rate = PREMIUM_DISCOUNT_RATE
	d.code = PREMIUM_DISCOUNT
}

func (d *discount) birthdaySpecial() {
	d.rate = BIRTHDAY_DISCOUNT_RATE
	d.code = BIRTHDAY_DISCOUNT
}

func (d *discount) addRates(di *discount, decimal int) {
	precision := 10 * float64(decimal)
	d.rate = ((d.rate * precision) + (di.rate * precision)) / precision
}
