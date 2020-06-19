package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	"github.com/tendermint/tendermint/crypto/merkle"
)

func main() {
	keyPtr := flag.String("key", "", "key")
	valuePtr := flag.String("value", "", "value")
	rootPtr := flag.String("root", "", "root hash in base64 encoded form")
	proofPtr := flag.String("proof", "", "proof in json form")

	flag.Parse()

	fmt.Printf("proof: \n%s\n", *proofPtr)
	fmt.Printf(" root: %s\n", *rootPtr)
	fmt.Printf("  key: %s\n", *keyPtr)
	fmt.Printf("value: %s\n", *valuePtr)

	err := verify(*proofPtr, *rootPtr, *keyPtr, *valuePtr)
	if err != nil {
		fmt.Printf("kv onchain verify (%s,%s) failed\n", *keyPtr, *valuePtr)
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	fmt.Printf("kv onchain verify (%s,%s) succeeded\n", *keyPtr, *valuePtr)
}

func verify(proofJson, rootBase64, key, value string) error {

	merkleProof := &merkle.Proof{Ops: make([]merkle.ProofOp, 1)}
	err := merkleProof.UnmarshalJSON([]byte(proofJson))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return err
	}

	rootBytes, err := base64.StdEncoding.DecodeString(rootBase64)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return err
	}

	fmt.Printf("proof: %s\n root: %s\n key: %s\n value: %s\n",
		proofJson,
		rootBase64,
		key,
		value)

	prf := rootmulti.DefaultProofRuntime()
	return prf.VerifyValue(merkleProof, rootBytes, fmt.Sprintf("/%s", key), []byte(value))
}
