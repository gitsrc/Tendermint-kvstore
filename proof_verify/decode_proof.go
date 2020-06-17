package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tendermint/tendermint/crypto/merkle"
)

func DecodeKeyValAndProof() (key string, value string, proof *merkle.Proof) {
	data := readProof()
	if len(data) != 8 {
		panic(fmt.Sprintf("invalid proof data, %v ", data))
	}

	key = strings.TrimSpace(strings.Split(data[3], "key:")[1])
	value = strings.TrimSpace(strings.Split(data[5], "value:")[1])
	proof = decodeProof(data[len(data)-1])
	return
}

func readProof() []string {
	proofStr := readFile(*proofFile)
	data := strings.Split(string(proofStr), "->")
	return data
}

func decodeProof(proofData string) *merkle.Proof {
	proof := &merkle.Proof{
		Ops: make([]merkle.ProofOp, 1),
	}
	result := make([]byte, 0, 40)
	preFix := "Data:[]uint8{"
	begin := strings.Index(proofData, preFix)
	if begin == -1 {
		panic(fmt.Sprintf("not find proof data"))
	}
	proofData = proofData[begin+len(preFix):]
	end := strings.Index(proofData, "},")
	if end == -1 {
		panic(fmt.Sprintf("not find proof data end"))
	}
	elems := strings.Split(proofData[:end], ",")
	for _, v := range elems {
		num, err := strconv.ParseInt(strings.Split(v, "0x")[1], 16, 16)
		if err != nil {
			panic(fmt.Sprintf("string(%s) convert to number failed: %v", v, err))
		}
		result = append(result, byte(num))
	}
	proof.Ops[0] = merkle.ProofOp{
		Data: result,
	}
	return proof
}
