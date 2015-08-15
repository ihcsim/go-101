package discountcalculator

const (
	STANDARD = iota
	SILVER
	GOLD
	PREMIUM
)

type customer struct {
	category int
}

func NewCustomer(category int) *customer {
	return &customer{
		category: category,
	}
}
