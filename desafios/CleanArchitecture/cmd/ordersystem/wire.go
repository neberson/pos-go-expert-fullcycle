//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/event"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/infra/database"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/pkg/events"
)

var SetOrderRepository = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(order.RepositoryInterface), new(*database.OrderRepository)),
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
