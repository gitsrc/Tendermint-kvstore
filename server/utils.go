package server

import (
	"flag"
	"path/filepath"

	dbm "github.com/tendermint/tm-db"

	cfg "github.com/tendermint/tendermint/config"
)

var (
	home = flag.String("home", ".", "Dir in ")
)

func initConfig() *cfg.Config {
	flag.Parse()
	return &cfg.Config{
		BaseConfig: cfg.BaseConfig{
			RootDir: *home,
		},
	}
}

func createDB(cfg *cfg.Config) dbm.DB {
	return dbm.NewDB("app", dbm.GoLevelDBBackend, filepath.Join(cfg.RootDir, "data"))
}
