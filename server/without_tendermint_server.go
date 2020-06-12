package server

import (
	"fmt"

	"github.com/ludete/kvstore/app"

	"github.com/tendermint/tendermint/abci/server"
	tmos "github.com/tendermint/tendermint/libs/os"
)

func AppWithoutTenderMint() {
	app := app.CreateKVStore(nil)

	server, err := server.NewServer("127.0.0.1:26658", "socket", app)
	if err != nil {
		fmt.Println("create server failed: ", err)
		return
	}
	if err := server.Start(); err != nil {
		fmt.Println("start server failed: ", err)
		return
	}

	// Stop upon receiving SIGTERM or CTRL-C.
	tmos.TrapSignal(nil, func() {
		// Cleanup
		server.Stop()
	})

	// Run forever.
	select {}
}
