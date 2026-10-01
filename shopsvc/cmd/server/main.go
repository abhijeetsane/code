package main

import (
	"database/sql"
	"log"
	nethttp "net/http"

	"shopsvc/internal/http"
	"shopsvc/internal/notification"
	orderservice "shopsvc/internal/order_service"
	"shopsvc/internal/payment"
	db "shopsvc/internal/platform/postgres"
)

func main() {
	conn, err := sql.Open("postgres", "postgres://localhost/shop")
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(conn)

	mgr := orderservice.NewOrderManager(
		store,
		payment.NewGateway("key"),
		&notification.EmailSender{},
	)

	log.Fatal(nethttp.ListenAndServe(":8080", http.NewRouter(mgr)))
}
