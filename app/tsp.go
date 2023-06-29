package app

import (
	"fmt"
	"log"

	"github.com/tendermint/tendermint/abci/server"
	tmos "github.com/tendermint/tendermint/libs/os"
)

func RemoteAppViaTSP() {
	// create app
	_, err := createDB(initConfig())
	if err != nil {
		log.Panic(err)
	}
	//app := NewKVStoreApplication(db)
	app := NewKVStoreApplication()
	// create server
	server, err := server.NewServer("127.0.0.1:26658", "socket", app)
	if err != nil {
		fmt.Println("create server failed: ", err)
		return
	}

	log.Println(server.String())

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
