package http

import (
	"encoding/json"
	nethttp "net/http"

	orderservice "shopsvc/internal/order_service"
)

type Router struct {
	orders *orderservice.OrderManager
}

func NewRouter(o *orderservice.OrderManager) nethttp.Handler {
	r := &Router{orders: o}
	mux := nethttp.NewServeMux()
	mux.HandleFunc("POST /orders", r.createOrder)
	return mux
}

func (rt *Router) createOrder(w nethttp.ResponseWriter, req *nethttp.Request) {
	_ = json.NewDecoder(req.Body)
	// ...
}
