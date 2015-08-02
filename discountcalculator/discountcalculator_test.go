package discountcalculator

import "testing"

func TestWhenGivenAStandardCustomer_ThenPerformsStandardDiscountCalculation(t *testing.T) {
	discountCalculator := New()
	customer := NewCustomer(STANDARD)
	discountCalculator.CalcFor(customer)
	if discountCalculator.strategy() != STANDARD_DISCOUNT {
		t.Errorf("Expected calculation strategy to be %s, but get %s\n", "STANDARD", discountCalculator.strategy)
	}
}

func TestWhenGivenASilverCustomer_ThenPerformsSilverDiscountCalculation(t *testing.T) {
	discountCalculator := New()
	customer := NewCustomer(SILVER)
	discountCalculator.CalcFor(customer)
	if discountCalculator.strategy() != SILVER_DISCOUNT {
		t.Errorf("Expected calculation strategy to be %s, but get %s\n", "SILVER", discountCalculator.strategy)
	}
}

func TestWhenGivenAGoldCustomer_ThenPerformsGoldDiscountCalculation(t *testing.T) {
	discountCalculator := New()
	customer := NewCustomer(GOLD)
	discountCalculator.CalcFor(customer)
	if discountCalculator.strategy() != GOLD_DISCOUNT {
		t.Errorf("Expected calculation strategy to be %s, but get %s\n", "GOLD", discountCalculator.strategy)
	}
}

func TestWhenGivenAPremiumCustomer_ThenPerformsPremiumDiscountCalculation(t *testing.T) {
	discountCalculator := New()
	customer := NewCustomer(PREMIUM)
	discountCalculator.CalcFor(customer)
	if discountCalculator.strategy() != PREMIUM_DISCOUNT {
		t.Errorf("Expected calculation strategy to be %s, but get %s\n", "PREMIUM", discountCalculator.strategy)
	}
}
