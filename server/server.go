package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/net/context"
)

/*ServerThing is the server type*/
type ServerThing struct {
}

/*Server object */
func Server() *ServerThing {
	return &ServerThing{}
}

/*Run starts the server */
func (s *ServerThing) Run(ctx context.Context) error {
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
