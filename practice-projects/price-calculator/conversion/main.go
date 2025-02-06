package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64 = make([]float64, len(strings))
	for _, string := range strings {
		price, err := strconv.ParseFloat(string, 64)

		if err != nil {
			return nil, errors.New("Failed to convert string to float.")
		}

		floats = append(floats, price)
	}
	return floats, nil
}
