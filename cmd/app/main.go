package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	env, err := newEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	app, err := newApp(context.TODO(), env)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelFunc := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)
	defer cancelFunc()

	<-app.run(ctx)
	if err := app.shutdown(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
