package discountcalculator

import "testing"

func TestWhenGivenAStandardCustomer_ThenReturns10PercentDiscount(t *testing.T) {
	discountCalculator := New()
	customer := NewCustomer(STANDARD)
	rate := discountCalculator.RateFor(customer)

	if rate != 0.1 {
		t.Errorf("Expected discount rate to be %.2f, but get %.2f\n", 0.1, rate)
	}
	if discountCalculator.strategyCode != STANDARD_DISCOUNT {
		t.Errorf("Expected calculation strategy to be %s, but get %s\n", "STANDARD", discountCalculator.strategyCode)
	}
}

func TestWhenGivenASilverCustomer_ThenReturns15PercentDiscount(t *testing.T) {
	discountCalculator := New()
	customer := NewCustomer(SILVER)
	rate := discountCalculator.RateFor(customer)

	if rate != 0.15 {
		t.Errorf("Expected discount rate to be %.2f, but get %.2f\n", 0.15, rate)
	}

	if discountCalculator.strategyCode != SILVER_DISCOUNT {
		t.Errorf("Expected calculation strategy to be %s, but get %s\n", "SILVER", discountCalculator.strategyCode)
	}
}

func TestWhenGivenAGoldCustomer_ThenReturns20PercentDiscount(t *testing.T) {
	discountCalculator := New()
	customer := NewCustomer(GOLD)
	rate := discountCalculator.RateFor(customer)

	if rate != 0.2 {
		t.Errorf("Expected discount rate to be %.2f, but get %.2f\n", 0.2, rate)
	}

	if discountCalculator.strategyCode != GOLD_DISCOUNT {
		t.Errorf("Expected calculation strategy to be %s, but get %s\n", "GOLD", discountCalculator.strategyCode)
	}
}

func TestWhenGivenAPremiumCustomer_ThenReturns25PercentDiscount(t *testing.T) {
	discountCalculator := New()
	customer := NewCustomer(PREMIUM)
	rate := discountCalculator.RateFor(customer)

	if rate != 0.25 {
		t.Errorf("Expected discount rate to be %.2f, but get %.2f\n", 0.25, rate)
	}

	if discountCalculator.strategyCode != PREMIUM_DISCOUNT {
		t.Errorf("Expected calculation strategy to be %s, but get %s\n", "PREMIUM", discountCalculator.strategyCode)
	}
}
