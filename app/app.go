package app

import (
	"github.com/tendermint/tendermint/abci/types"
)

func CreateKVStore() types.Application {
	return nil
}

type PersisApplication struct {
	types.BaseApplication
}
