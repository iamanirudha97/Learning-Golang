package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, str := range strings {
		floatPrice, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Converting text to float failed")
			return nil, errors.New("failed to convert string to float")
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}
