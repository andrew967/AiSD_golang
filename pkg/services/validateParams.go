package services

import "errors"

func ValidateParams(arr []float32) error {
	if len(arr) == 0 {
		return errors.New("Array is empty")
	}
	return nil
}
