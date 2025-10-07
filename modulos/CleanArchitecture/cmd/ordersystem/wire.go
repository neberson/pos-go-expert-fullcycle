//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

var SetOrderRepository = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(order.RepositoryInterface), new(*database.OrderRepository)),
)
