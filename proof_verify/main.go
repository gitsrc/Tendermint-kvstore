package main

import (
	"encoding/hex"
	"fmt"

	"github.com/tendermint/tendermint/crypto/merkle"

	"github.com/cosmos/cosmos-sdk/store/rootmulti"
)

func main() {
	commit, err := hex.DecodeString("64")
	if err != nil {
		panic(err)
	}
	fmt.Println(commit)
	data := readFile("")
	verifyProof(data)
}

func readFile(filePath string) []byte {
	return nil
}

func verifyProof(data []byte) bool {
	prf := rootmulti.DefaultProofRuntime()
	commit, err := hex.DecodeString("647BCF46010345822E0AB477949CF4911456D4E39F3ECD08C131B3F7CDE04049")
	if err != nil {
		panic(err)
	}
	prof := &merkle.Proof{
		Ops: []merkle.ProofOp{
			{
				Type: "iavl:v",
				Key:  []uint8{0x61, 0x62, 0x63, 0x64},
				Data: []uint8{0x2f, 0xa, 0x2d, 0x1a, 0x2b, 0xa, 0x4, 0x61, 0x62, 0x63, 0x64, 0x12, 0x20, 0x88, 0xd4,
					0x26, 0x6f, 0xd4, 0xe6, 0x33, 0x8d, 0x13, 0xb8, 0x45, 0xfc, 0xf2, 0x89, 0x57, 0x9d, 0x20, 0x9c,
					0x89, 0x78, 0x23, 0xb9, 0x21, 0x7d, 0xa3, 0xe1, 0x61, 0x93, 0x6f, 0x3, 0x15, 0x89, 0x18, 0xbd, 0x2},
			},
		},
	}
	//proof := merkle.Proof{}
	//proof.Unmarshal(data)
	err = prf.VerifyValue(prof, commit, "/abcd", []byte("abcd"))
	if err != nil {
		panic(err)
	}
	return true
}
