package utils

import "fmt"

func TakeFive(in int) (int, error) {
	if in < 5 {
		return 0, fmt.Errorf("the number supplied is less than 5")
	}
	return in - 5, nil
}
