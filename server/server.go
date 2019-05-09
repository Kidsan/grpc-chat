package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/net/context"
)

type server struct {
}

func Server() *server {

}

func (s *server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	fmt.Printf(time.Now().String(), "starting chat server")
}

func main() {
	ctx := context.Background()
	var err error

	err = Server().Run(ctx)

	if err != nil {
		fmt.Printf("<<Process>>", err.Error())
		os.Exit(1)
	}
}
