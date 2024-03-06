package main

import (
	"bookstore-webapi/application"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	bookStoreApp := application.NewBookStoreApp(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := bookStoreApp.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
