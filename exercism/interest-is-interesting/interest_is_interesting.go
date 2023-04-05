package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	}
	if balance < 1000 {
		return 0.5
	}
	if balance < 5000 {
		return 1.621
	}
	return 2.475
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	//panic("Please implement the AnnualBalanceUpdate function")
	// for i := 0; i < 12; i++ {
	// 	balance += Interest(balance)
	// }

	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	//panic("Please implement the YearsBeforeDesiredBalance function")
	years := 0

	for balance < targetBalance {
		years++
		balance += Interest(balance)
	}
	return years
}
