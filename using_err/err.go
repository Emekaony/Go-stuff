package using_err

import (
	"errors"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// there's an error, return the zero value for the int
		// and us errors.new(<message>) to convey that there was an error
		return float64(0.0), errors.New("There was an error during the division!")
	} else {
		return a / b, nil
	}
}
