package server

import (
	"fmt"

	tmos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	dbm "github.com/tendermint/tm-db"

	"github.com/ludete/kvstore/app"
)

func AppWithTenderMint() {
	// init config
	cfg := initConfig()
	db := dbm.NewDB("app", dbm.GoLevelDBBackend, ".")
	app := app.NewPersisApplication(db)
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

//func createCmd() *cobra.Command {
//	ctx := server.NewDefaultContext()
//	_ = ctx
//	//rootCmd := &cobra.Command{
//	//	Use:               "smartd",
//	//	Short:             "CoinEx Smart Chain Daemon (server)",
//	//	PersistentPreRunE: server.PersistentPreRunEFn(ctx),
//	//}
//	return nil
//}
