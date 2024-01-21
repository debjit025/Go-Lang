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
		return 10000, errors.New("Failed to find the file.")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 10000, errors.New("Failed to Parse stored value.")
	}
	return value, nil
}

func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
