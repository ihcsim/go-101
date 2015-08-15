package discountcalculator

type checkoutCode int

const (
	STANDARD_CHECKOUT checkoutCode = iota
	EXPRESS_CHECKOUT
)

func (c checkoutCode) String() string {
	switch c {
	case STANDARD_CHECKOUT:
		return "STANDARD_CHECKOUT"
	case EXPRESS_CHECKOUT:
		return "EXPRESS_CHECKOUT"
	}
	return ""
}

func standardCheckout(invoiceTotal float64, d *discount) (balance float64, code checkoutCode) {
	return invoiceTotal * (1.00 - d.rate), STANDARD_CHECKOUT
}

func expressCheckout(invoiceTotal float64, d *discount) (balance float64, code checkoutCode) {
	return invoiceTotal * (1.00 - d.rate), EXPRESS_CHECKOUT
}
