package main

import (
	"errors"
	"fmt"
	"github.com/sgahlot/go-postgres-service-quickstart/pkg/common"
	"github.com/sgahlot/go-postgres-service-quickstart/pkg/db"
	"log"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	SERVER_PORT = ":8080"
)

func makeRoute(service common.Service) http.Handler {
	ctx := common.GetContext()

	endPoints := common.EndPoints{
		InsertFruit: common.InsertFruit(service),
		DeleteFruit: common.DeleteFruits(service),
		GetFruit:    common.GetFruit(service),
		GetFruits:   common.GetFruits(service),
	}

	router := common.CreateHandlers(ctx, endPoints)

	return router
}

func main() {
	router := makeRoute(&db.FruitService{})

	errChan := make(chan error)
	go func() {
		log.Printf("Starting FruitShop server at port %s\n", SERVER_PORT)
		handler := router
		errChan <- http.ListenAndServe(SERVER_PORT, handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- errors.New(fmt.Sprintf("Exit status %v", <-c))
	}()

	fmt.Println(<-errChan)
}
