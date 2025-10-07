//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/entity"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/event"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/infra/database"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/infra/web"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/usecase"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/pkg/events"
)

var SetOrderRepository = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	orderRepository := database.NewOrderRepository(db)
	return usecase.NewListOrdersUseCase(orderRepository)
}
