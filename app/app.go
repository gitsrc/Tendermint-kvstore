package app

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/tendermint/tendermint/abci/example/code"
	"github.com/tendermint/tendermint/libs/kv"

	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store"
	"github.com/tendermint/tendermint/abci/types"
)

type PersisApplication struct {
	types.Application

	// store function
	cms        store.CommitKVStore
	kvStoreKey store.StoreKey
}

func NewPersisApplication(db dbm.DB) *PersisApplication {
	app := &PersisApplication{
		Application: types.NewBaseApplication(),
	}
	app.initDB(db)
	return app
}

func (app *PersisApplication) initDB(db dbm.DB) (err error) {
	app.cms, err = iavl.LoadStore(db, store.CommitID{}, store.PruneNothing, false)
	if err != nil {
		return err
	}
	return nil
}

func (app *PersisApplication) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {

	var key, value []byte
	parts := bytes.Split(req.Tx, []byte("="))
	if len(parts) == 2 {
		key, value = parts[0], parts[1]
	} else {
		key, value = req.Tx, req.Tx
	}
	//iavlStore := app.cms.GetCommitStore(app.kvStoreKey).(*iavl.Store)
	app.cms.Set(key, value)
	//iavlStore.Set(key, value)
	commit := app.cms.Commit()
	events := []types.Event{
		{
			Type: "app",
			Attributes: []kv.Pair{
				{Key: []byte("creator"), Value: []byte("Cosmoshi Netowoko")},
				{Key: []byte("key"), Value: key},
				{Key: []byte("hash"), Value: commit.Hash},
			},
		},
	}
	fmt.Println("---------- commit hash : ", commit.String())
	return types.ResponseDeliverTx{Code: code.CodeTypeOK, Events: events}
}

//0xC4B8100E0A64BDF915E3545CEAFA98C1DF3BFC5BAF85665B6BF294CCF8C98350
func (app *PersisApplication) Commit() types.ResponseCommit {
	appHash := app.cms.Commit()
	commit := app.cms.Commit()
	//commit := iavlStore.Commit()
	fmt.Printf("===========commit hash : %s\n", commit.String())
	return types.ResponseCommit{
		Data: appHash.Hash,
	}
}

func (app *PersisApplication) Query(req types.RequestQuery) types.ResponseQuery {
	iavlStore := app.cms.(*iavl.Store)
	fmt.Printf("custom query data : %s\n", req.Data)
	res := iavlStore.Query(types.RequestQuery{
		Path:  "/key", // required path to get key/value+proof
		Data:  req.Data,
		Prove: true,
	})
	fmt.Printf("proof : %s\n", res.Proof.String())
	return res
}
