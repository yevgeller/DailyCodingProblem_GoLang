package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {

	out := make([]Record, 0)

	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
		}
	}

	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var fil = ByDaysPeriod(p) //simplify
	records := Filter(in, fil)
	return calcTotal(records)
}

func calcTotal(in []Record) float64 {
	total := 0.0

	for _, r := range in {
		total += r.Amount
	}

	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	//panic("Please implement the CategoryExpenses function")
	var byCategory = Filter(in, ByCategory(c))
	if len(byCategory) == 0 {
		return 0.0, errors.New("unknown category ")
	}
	var byDays = Filter(byCategory, ByDaysPeriod(p))

	return calcTotal(byDays), nil
}
