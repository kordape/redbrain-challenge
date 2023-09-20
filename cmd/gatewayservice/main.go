package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/kordape/redbrain-challenge/internal/http"
	"github.com/kordape/redbrain-challenge/pkg/encryptorapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Initializing gateway service")

	conn, err := grpc.Dial("encryption-service:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(fmt.Sprintf("Error establishing connection to encryption service: %s", err))
	}

	client := encryptorapi.NewEncryptorClient(conn)

	// HTTP Server
	handler := gin.New()
	http.NewRouter(
		handler,
		client,
	)
	httpServer := http.New(handler, http.Port("8080"))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Println(fmt.Sprintf("app - Run - signal: " + s.String()))
	case err := <-httpServer.Notify():
		// Shutdown
		err = httpServer.Shutdown()
		if err != nil {
			fmt.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
		}
	}
}
