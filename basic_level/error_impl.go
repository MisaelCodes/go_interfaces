package basic_level

// this exercise could be good for custom error
// errors that deal more with business logic rather
// than dev skills.
type divisionError struct {
	message string
}

func (de divisionError) Error() string {
	return de.message
}

func Divide[T int | float64](a, b T) (T, error) {
	if b == 0 {
		e := &divisionError{"can't divide by 0"}
		return 0, e
	}
	return a / b, nil
}
