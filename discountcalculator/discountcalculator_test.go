package discountcalculator

import "testing"

func TestRateFor_WhenGivenCustomerType_ThenReturnsTheCorrectDiscountRateAndType(t *testing.T) {
	var tests = []struct {
		customer             *customer
		expectedRate         float64
		expectedDiscountType int
	}{
		{customer: NewCustomer(STANDARD), expectedRate: 0.1, expectedDiscountType: STANDARD_DISCOUNT},
		{customer: NewCustomer(SILVER), expectedRate: 0.15, expectedDiscountType: SILVER_DISCOUNT},
		{customer: NewCustomer(GOLD), expectedRate: 0.2, expectedDiscountType: GOLD_DISCOUNT},
		{customer: NewCustomer(PREMIUM), expectedRate: 0.25, expectedDiscountType: PREMIUM_DISCOUNT},
	}

	discountCalculator := New()
	for _, test := range tests {
		if rate := discountCalculator.RateFor(test.customer); rate != test.expectedRate {
			t.Errorf("Expected discount rate to be %.2f, but get %.2f\n", test.expectedRate, rate)
		}
		if discountType := discountCalculator.strategyCode; discountType != test.expectedDiscountType {
			t.Errorf("Expected calculation strategy to be %d, but get %d\n", test.expectedDiscountType, discountType)
		}
	}
}

func TestRateWithCouponFor_WhenGivenCustomerAndCouponType_ThenReturnsTheCorrectDiscountRateAndType(t *testing.T) {
	var tests = []struct {
		customer             *customer
		couponType           int
		expectedRate         float64
		expectedDiscountType int
	}{
		{customer: NewCustomer(STANDARD), couponType: BIRTHDAY_ANNIVERSARY, expectedRate: 0.15, expectedDiscountType: BIRTHDAY_DISCOUNT},
		{customer: NewCustomer(SILVER), couponType: BIRTHDAY_ANNIVERSARY, expectedRate: 0.20, expectedDiscountType: BIRTHDAY_DISCOUNT},
		{customer: NewCustomer(GOLD), couponType: BIRTHDAY_ANNIVERSARY, expectedRate: 0.25, expectedDiscountType: BIRTHDAY_DISCOUNT},
		{customer: NewCustomer(PREMIUM), couponType: BIRTHDAY_ANNIVERSARY, expectedRate: 0.30, expectedDiscountType: BIRTHDAY_DISCOUNT},
	}

	discountCalculator := New()
	for _, test := range tests {
		if rate := discountCalculator.RateWithCouponFor(test.customer, test.couponType); rate != test.expectedRate {
			t.Errorf("Expected discount rate with promotion to be %.2f, but get %.2f", test.expectedRate, rate)
		}

		if discountType := discountCalculator.strategyCode; discountType != test.expectedDiscountType {
			t.Errorf("Expected discount type to be %d, but get %d\n", test.expectedDiscountType, discountType)
		}
	}
}
