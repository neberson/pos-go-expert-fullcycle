package graph

import "github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
}
