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

	// proof type
	typePrefix := "Type:\""
	typeIndex := strings.Index(proofData, typePrefix)
	if typeIndex == -1 {
		panic(fmt.Sprintf("not find type content in proof"))
	}
	proofData = proofData[typeIndex+len(typePrefix):]
	endIndex := strings.Index(proofData, "\",")
	typeData := proofData[:endIndex]

	// proof key
	key := make([]byte, 0, 40)
	keyPrefix := "{"
	keyIndex := strings.Index(proofData, keyPrefix)
	if keyIndex == -1 {
		panic(fmt.Sprintf("not find key content in proof, data : %s", proofData))
	}
	proofData = proofData[keyIndex+len(keyPrefix):]
	endIndex = strings.Index(proofData, "}")
	key = parseHex(proofData[:endIndex])

	// proof data
	data := make([]byte, 0, 40)
	preFix := "Data:[]uint8{"
	begin := strings.Index(proofData, preFix)
	if begin == -1 {
		panic(fmt.Sprintf("not find proof data in proof, data : %s", proofData))
	}
	proofData = proofData[begin+len(preFix):]
	end := strings.Index(proofData, "},")
	if end == -1 {
		panic(fmt.Sprintf("not find proof data end"))
	}
	data = parseHex(proofData[:end])
	proof.Ops[0] = merkle.ProofOp{
		Type: typeData,
		Key:  key,
		Data: data,
	}
	return proof
}

func parseHex(hexData string) (result []byte) {
	elms := strings.Split(hexData, ",")
	for _, v := range elms {
		num, err := strconv.ParseInt(strings.Split(v, "0x")[1], 16, 16)
		if err != nil {
			panic(fmt.Sprintf("string(%s) convert to number failed: %v", v, err))
		}
		result = append(result, byte(num))
	}
	return result
}
