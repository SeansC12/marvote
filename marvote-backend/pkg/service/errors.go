package service

import "fmt"

type ErrorFailedToLoadData struct {
	Status  int
	Message string
}

func (r *ErrorFailedToLoadData) Error() string {
	return fmt.Sprintf("status %d: err %v", 500, "failed to load data")
}
