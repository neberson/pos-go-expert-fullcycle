package entity

import (
	"fmt"
	"regexp"
)

var (
	ErrInvalidCep = fmt.Errorf("invalid zipcode")
)

type Cep struct {
	Cep string `json:"cep"`
}

func NewCep(cep string) *Cep {
	return &Cep{Cep: cep}
}

func (c *Cep) Validate() error {
	if !regexp.MustCompile(`^\d{8}$`).MatchString(c.Cep) {
		return ErrInvalidCep
	}
	return nil
}
