package web

import (
	"encoding/json"
	"net/http"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/entity"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/internal/usecase"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/CleanArchitecture/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	eventDispatcher events.EventDispatcherInterface,
	orderRepository entity.OrderRepositoryInterface,
	orderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   eventDispatcher,
		OrderRepository:   orderRepository,
		OrderCreatedEvent: orderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	listOrders := usecase.NewListOrdersUseCase(h.OrderRepository)
	output, err := listOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
