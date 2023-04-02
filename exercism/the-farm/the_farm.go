package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	message string
	details string
} 

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("%s %s", e.message, e.details)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
		
	amt, err := weightFodder.FodderAmount()
	//println("19")
	//println(amt, err)
	//println("amount", amt)

	//println("25")
	// if err != nil && err.Error() == ErrScaleMalfunction.Error() && amt > 0 {
	// 	return (amt * 2.0) / float64(cows), nil
	// }

	if err != nil && err.Error() == ErrScaleMalfunction.Error() && amt > 0 {
		return (amt * 2.0) / float64(cows), nil
	}



	if amt < 0 && (err == nil || err.Error() == ErrScaleMalfunction.Error()) {
		return 0.0, errors.New("negative fodder")
	}
	if err != nil {
		println(err.Error(), "amt ", amt)
		return 0.0, err
	}
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}
	if cows < 0 {
		//println("15")
		return 0.0, &SillyNephewError{message: "silly nephew, there cannot be", details: fmt.Sprintf("%d cows", cows)}
	}
	//println("29")

	//println("33")
	// if err != nil && err.Error() == ErrScaleMalfunction.Error() && amt > 0 {
	// 	return (amt * 2.0) / float64(cows), nil
	// }

	// if err != nil {
	// 	println(err.Error())
	// 	return 0.0, err
	// }
	// if amt < 0 {
	// 	return 0.0, errors.New("negative fodder")
	// }
	//println("37")
	return amt / float64(cows), nil
}
