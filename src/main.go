package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/troodinc/trood/cmd"
)

func main() {
	ctx, shutdown := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer shutdown()

	root := cmd.RootCmd()
	if err := root.ExecuteContext(ctx); err != nil {
		log.Fatal(err)
	}
}
