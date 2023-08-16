package main

import (
	"context"
	"github.com/MGMCN/opa-demo/web"
)

func main() {
	var err error
	ctx := context.Background()

	apiServer := web.NewApiServer()
	err = apiServer.InitServer(ctx)
	if err != nil {
		// handle error
		return
	}

	err = apiServer.Serve(":3000")
	if err != nil {
		// handle error
		return
	}
}
