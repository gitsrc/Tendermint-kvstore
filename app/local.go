package app

import (
	"fmt"
	"log"

	tmos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	dbm "github.com/tendermint/tm-db"
)

func LocalApp() {
	// init config
	cfg := initConfig()
	_, err := dbm.NewDB("app", dbm.GoLevelDBBackend, ".")
	if err != nil {
		log.Panic(err)
	}
	//app := NewKVStoreApplication(db)
	app := NewKVStoreApplication()
	nodeKey, err := p2p.LoadOrGenNodeKey(cfg.NodeKeyFile())
	if err != nil {
		fmt.Println(err)
		return
	}

	// create node
	node, err := node.NewNode(cfg,
		privval.LoadOrGenFilePV(cfg.PrivValidatorKeyFile(), cfg.PrivValidatorStateFile()),
		nodeKey,
		proxy.NewLocalClientCreator(app),
		node.DefaultGenesisDocProviderFunc(cfg),
		node.DefaultDBProvider,
		node.DefaultMetricsProvider(cfg.Instrumentation),
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run and stop
	if err := node.Start(); err != nil {
		fmt.Println(err)
		return
	}
	tmos.TrapSignal(nil, func() {
		// Cleanup
		if err := node.Stop(); err != nil {
			fmt.Println(err)
			return
		}
	})

	// Run forever.
	select {}
}
