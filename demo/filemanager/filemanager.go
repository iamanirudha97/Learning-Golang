package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error Opening the file: ")
		fmt.Println(err)
		return nil, errors.New("failed to open the file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file Failed: ")
		fmt.Println(err)
		file.Close()
		return nil, errors.New("failed to read the file")
	}
	file.Close()
	return lines, nil
}

func WriteToJson(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("error occured while creating json file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to json")
	}
	file.Close()
	return nil
}
