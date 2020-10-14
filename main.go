package main

import (
	"log"

	"github.com/linkpoolio/bridges"
	"github.com/rocksideio/chainlink-metatx-adapter/app"
)

func main() {
	adapter, err := app.NewRelayerAdapterFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	bridges.NewServer(adapter).Start(9090)
}
