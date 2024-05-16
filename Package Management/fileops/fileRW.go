package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)

	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to write float to .txt file")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueStr := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueStr), 0644)
}
