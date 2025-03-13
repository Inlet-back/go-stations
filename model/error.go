package model

import "fmt"

type ErrNotFound struct{
		Message string
}

func (e *ErrNotFound) Error() string {
	return  fmt.Sprintf("ERROR: %s", e.Message)
}