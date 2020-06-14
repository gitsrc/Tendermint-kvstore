package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"

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
	dec := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString("43Fd0vw0I4Wkiks9odF241vKuS5gg8GY/04mBQHnEWc="))
	res1, err := ioutil.ReadAll(dec)
	if err != nil {
		panic(err)
	}
	fmt.Println("decode base64 : ", res1)
	prof := &merkle.Proof{
		Ops: []merkle.ProofOp{
			{
				Type: "iavl:v",
				Key:  []uint8{0x61, 0x62, 0x63},
				Data: []uint8{0x2d, 0xa, 0x2b, 0x1a, 0x29, 0xa, 0x3, 0x61, 0x62, 0x63, 0x12, 0x20,
					0xba, 0x78, 0x16, 0xbf, 0x8f, 0x1, 0xcf, 0xea, 0x41, 0x41, 0x40, 0xde, 0x5d,
					0xae, 0x22, 0x23, 0xb0, 0x3, 0x61, 0xa3, 0x96, 0x17, 0x7a, 0x9c, 0xb4, 0x10,
					0xff, 0x61, 0xf2, 0x0, 0x15, 0xad, 0x18, 0x5},
			},
		},
	}
	err = prf.VerifyValue(prof, res1, fmt.Sprintf("/%s", prof.Ops[0].Key), []byte("abc"))
	if err != nil {
		panic(err)
	}
	fmt.Println("ok ....")
	return true
}
