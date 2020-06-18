package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func DecodeCommitHash() []byte {
	rootHash := getRootHash()
	hash, err := hex.DecodeString(rootHash)
	if err != nil {
		panic(fmt.Sprintf("Convert hex hash failed : %v,\n actual rootHash: %s", err, rootHash))
	}
	return hash
}

func getRootHash() string {
	content, err := readEventRet()
	if err != nil || len(content) == 0 {
		panic(fmt.Sprintf("err : %v, content size : %d\n", err, len(content)))
	}
	data := strings.Split(content, "->")
	if len(data) != 3 {
		panic(fmt.Sprintf("Number of entries: %d, expected number : 3", len(data)))
	}
	okIndex := strings.Index(data[1], ":")
	isOk := strings.TrimSpace(data[1][okIndex+1:])
	rootIndex := strings.Index(data[2], ":")
	rootHash := strings.TrimSpace(data[2][rootIndex:])[4:]
	if isOk != "OK" {
		panic(fmt.Sprintf("RootHash result actual: [%s], expected: %s", isOk, "OK"))
	}
	return rootHash
}

func readEventRet() (string, error) {
	data := readFile(*appHashFile)
	return string(data), nil
}
