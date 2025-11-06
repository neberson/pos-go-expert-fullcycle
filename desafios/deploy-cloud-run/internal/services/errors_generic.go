package services

import "fmt"

var (
	ErrNotFound       = fmt.Errorf("can not find zipcode")
	ErrToManyRequests = fmt.Errorf("many requests for the zip code service")
	ErrGeneric        = fmt.Errorf("error when querying service")
)
