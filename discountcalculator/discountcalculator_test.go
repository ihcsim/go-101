package discountcalculator

import "testing"

func TestDiscountFor_GivenACustomer_ThenReturnsTheCorrectDiscountRateAndCode(t *testing.T) {
	var tests = []struct {
		customer             *customer
		expectedDiscountRate float64
		expectedDiscountCode int
	}{
		{customer: NewCustomer(STANDARD), expectedDiscountRate: 0.1, expectedDiscountCode: STANDARD_DISCOUNT},
		{customer: NewCustomer(SILVER), expectedDiscountRate: 0.15, expectedDiscountCode: SILVER_DISCOUNT},
		{customer: NewCustomer(GOLD), expectedDiscountRate: 0.2, expectedDiscountCode: GOLD_DISCOUNT},
		{customer: NewCustomer(PREMIUM), expectedDiscountRate: 0.25, expectedDiscountCode: PREMIUM_DISCOUNT},
	}

	discountCalculator := New()
	for _, test := range tests {
		discount := discountCalculator.DiscountFor(test.customer)
		if discount.rate != test.expectedDiscountRate {
			t.Errorf("Expected discount rate to be %.2f, but get %.2f\n", test.expectedDiscountRate, discount.rate)
		}
		if discount.code != test.expectedDiscountCode {
			t.Errorf("Expected discount code to be %d, but get %d\n", test.expectedDiscountCode, discount.code)
		}
	}
}

func TestSpecialDiscountFor_WhenCustomerIsGivenSpecialDiscount_ThenReturnsTheCorrectDiscountRateAndCode(t *testing.T) {
	var tests = []struct {
		customer             *customer
		couponType           int
		expectedDiscountRate float64
		expectedDiscountCode int
	}{
		{customer: NewCustomer(STANDARD), couponType: BIRTHDAY_ANNIVERSARY, expectedDiscountRate: 0.15, expectedDiscountCode: BIRTHDAY_DISCOUNT},
		{customer: NewCustomer(SILVER), couponType: BIRTHDAY_ANNIVERSARY, expectedDiscountRate: 0.20, expectedDiscountCode: BIRTHDAY_DISCOUNT},
		{customer: NewCustomer(GOLD), couponType: BIRTHDAY_ANNIVERSARY, expectedDiscountRate: 0.25, expectedDiscountCode: BIRTHDAY_DISCOUNT},
		{customer: NewCustomer(PREMIUM), couponType: BIRTHDAY_ANNIVERSARY, expectedDiscountRate: 0.30, expectedDiscountCode: BIRTHDAY_DISCOUNT},
	}

	discountCalculator := New()
	for _, test := range tests {
		discount := discountCalculator.SpecialDiscountFor(test.customer, test.couponType)
		if discount.rate != test.expectedDiscountRate {
			t.Errorf("Expected discount rate with promotion to be %.2f, but get %.2f", test.expectedDiscountRate, discount.rate)
		}

		if discount.code != test.expectedDiscountCode {
			t.Errorf("Expected discount code to be %d, but get %d\n", test.expectedDiscountCode, discount.code)
		}
	}
}

func TestCheckout_GivenACustomer_ThenReturnsTheCorrectCheckoutBalanceAndCode(t *testing.T) {
	var tests = []struct {
		customer                *customer
		invoiceTotal            float64
		expectedCheckoutBalance float64
		expectedCheckoutCode    int
	}{
		{customer: NewCustomer(STANDARD), invoiceTotal: 10.0, expectedCheckoutBalance: 9.0, expectedCheckoutCode: STANDARD_CHECKOUT},
		{customer: NewCustomer(SILVER), invoiceTotal: 10.0, expectedCheckoutBalance: 8.50, expectedCheckoutCode: STANDARD_CHECKOUT},
		{customer: NewCustomer(GOLD), invoiceTotal: 10.0, expectedCheckoutBalance: 8.00, expectedCheckoutCode: EXPRESS_CHECKOUT},
		{customer: NewCustomer(PREMIUM), invoiceTotal: 10.0, expectedCheckoutBalance: 7.50, expectedCheckoutCode: EXPRESS_CHECKOUT},
	}

	invoiceTotal := 10.0
	for _, test := range tests {
		discountCalculator := New()
		balance, checkoutCode := discountCalculator.Checkout(test.customer, invoiceTotal)
		if checkoutCode != test.expectedCheckoutCode {
			t.Errorf("Expected checkout code to be %d, but get %d", test.expectedCheckoutCode, checkoutCode)
		}

		if balance != test.expectedCheckoutBalance {
			t.Errorf("Expected checkout balance to be %.2f, but get %.2f", test.expectedCheckoutBalance, balance)
		}
	}
}
