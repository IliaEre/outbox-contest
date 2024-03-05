package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"log"
	"net/http"
	"order-service/config"
	"strconv"
	"time"

	"order-service/handler"
)

func BindServer(cfg config.Config, handler handler.OrderHandler) *http.Server {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(otelmux.Middleware("order-server"))

	r.Handle("/api/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})
	}))

	r.HandleFunc("/api/v1/create-order", handler.SaveOrder).Methods("POST")

	server := &http.Server{
		Addr:         cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port),
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server will be started with address: %s", server.Addr)

	return server
}

func Run(server *http.Server) {
	go func() {
		fmt.Println("Server is running on port 8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
}
