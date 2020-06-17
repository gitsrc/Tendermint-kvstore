package main

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/os"
)

func readFile(filePath string) []byte {
	if !os.FileExists(filePath) {
		panic(fmt.Sprintf("file not exist, filepath : %s", filePath))
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("read file failed: %v", err))
	}
	return data
}
