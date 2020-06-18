package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/store/rootmulti"

	"github.com/tendermint/tendermint/crypto/merkle"
)

var (
	proofFile   = flag.String("proof-file", "", "proof data file")
	appHashFile = flag.String("hash-file", "", "app hash data file")
	help        = flag.Bool("help", false, "help usage")
)

func init() {
	flag.Parse()
}

func main() {
	if *help {
		flag.Usage()
		os.Exit(1)
	}

	hash := DecodeCommitHash()
	fmt.Println("root hash : ", hash)
	key, val, proof := DecodeKeyValAndProof()
	fmt.Printf("proof: key: %s, val: %s, proof: %v\n", key, val, *proof)
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
