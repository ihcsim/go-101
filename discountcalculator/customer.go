package discountcalculator

const (
	STANDARD = iota
	SILVER   = iota
	GOLD     = iota
	PREMIUM  = iota
)

type customer struct {
	category int
}

func NewCustomer(category int) *customer {
	return &customer{
		category: category,
	}
}
