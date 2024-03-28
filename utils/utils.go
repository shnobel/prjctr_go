package utils

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func GetDataFromJson[T any](fileName string) (*T, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var t T
	json.Unmarshal(byteValue, &t)
	return &t, nil
}
