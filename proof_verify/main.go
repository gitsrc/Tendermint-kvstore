package main

import (
	"flag"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/rootmulti"

	"github.com/tendermint/tendermint/crypto/merkle"
)

var proofFile = flag.String("proof-file", "", "proof data file")
var appHashFile = flag.String("hash-file", "", "app hash data file")

func main() {
	hash := DecodeCommitHash()
	fmt.Println("hash : ", hash)
	key, val, proof := DecodeKeyValAndProof()
	fmt.Println("key: ", key, ", val: ", val, ", proof: ", proof)
	verifyProof(hash, key, val, proof)
}

func verifyProof(hash []byte, key, value string, proof *merkle.Proof) bool {
	prf := rootmulti.DefaultProofRuntime()
	err := prf.VerifyValue(proof, hash, fmt.Sprintf("/%s", key), []byte(value))
	if err != nil {
		panic(err)
	}
	fmt.Println("verify proof ok ....")
	return true
}
