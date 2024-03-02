package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"order-service/config"
	"order-service/handler"
	"os"
	"os/signal"
	"syscall"
	"time"

	order_svc "order-service/app/usecase/order"
	order_db "order-service/pkg/db"
	order_server "order-service/pkg/server"
	order_repo "order-service/repository/order"
)

func main() {
	cfg := config.ReadConfig()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	db := order_db.NewConnection(cfg)
	defer db.Close()

	repo := order_repo.NewOrderRepository(db)
	srv := order_svc.NewOrderService(repo)
	orderHndl := handler.NewOrderHandler(srv)

	server := order_server.BindServer(cfg, *orderHndl)
	order_server.Run(server)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %v", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}

	fmt.Println("Exiting the application")
}
