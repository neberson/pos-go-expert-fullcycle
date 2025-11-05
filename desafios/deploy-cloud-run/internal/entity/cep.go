package entity

import (
	"errors"
	"regexp"
)

const (
	ErrInvalidCep = "invalid zipcode"
)

type Cep struct {
	Cep string `json:"cep"`
}

func NewCep(cep string) *Cep {
	return &Cep{Cep: cep}
}

func (c *Cep) Validate() error {
	if !regexp.MustCompile(`^\d{8}$`).MatchString(c.Cep) {
		return errors.New(ErrInvalidCep)
	}
	return nil
}
