package discountcalculator

const (
	STANDARD_CHECKOUT = iota + 1
	EXPRESS_CHECKOUT
)

func standardCheckout(invoiceTotal float64, d *discount) (balance float64, checkoutCode int) {
	return invoiceTotal * (1.00 - d.rate), STANDARD_CHECKOUT
}

func expressCheckout(invoiceTotal float64, d *discount) (balance float64, checkoutCode int) {
	return invoiceTotal * (1.00 - d.rate), EXPRESS_CHECKOUT
}
