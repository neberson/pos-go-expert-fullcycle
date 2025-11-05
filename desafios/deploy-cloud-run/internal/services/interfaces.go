package services

import (
	"context"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/deploy-cloud-run/internal/entity"
)

type CepServiceInterface interface {
	GetCepViaCep(ctx context.Context, cep string) (*entity.PostalAddress, error)
}
