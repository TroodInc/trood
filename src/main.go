package main

import (
	"context"
	"github.com/troodinc/trood/cmd"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, shutdown := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer shutdown()

	root := cmd.NewRootCmd()
	if err := root.ExecuteContext(ctx); err != nil {
		log.Fatal(err)
	}
}
