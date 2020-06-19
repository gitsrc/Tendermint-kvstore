package app

import (
	"fmt"

	"github.com/tendermint/tendermint/abci/server"
	tmos "github.com/tendermint/tendermint/libs/os"
)

func RemoteAppViaTSP() {
	// create app
	db := createDB(initConfig())
	app := NewKVStoreApplication(db)

	// create server
	server, err := server.NewServer("127.0.0.1:26658", "socket", app)
	if err != nil {
		fmt.Println("create server failed: ", err)
		return
	}

	// start server
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
