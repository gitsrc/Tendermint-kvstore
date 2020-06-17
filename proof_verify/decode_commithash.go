package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JsonResult struct {
	Result struct {
		DeliverTx struct {
			Events []struct {
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"events"`
		} `json:"deliver_tx"`
	} `json:"result"`
}

func DecodeCommitHash() []byte {
	hash := getCommitHash()
	dec := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(hash))
	res1, err := ioutil.ReadAll(dec)
	if err != nil {
		panic(err)
	}
	return res1
}

func getCommitHash() string {
	content, err := readEventRet()
	if err != nil || len(content) == 0 {
		panic(fmt.Sprintf("err : %v, content size : %d\n", err, len(content)))
	}
	var r JsonResult
	if err := json.Unmarshal([]byte(content), &r); err != nil {
		panic(err)
	}
	if len(r.Result.DeliverTx.Events) == 0 || len(r.Result.DeliverTx.Events[0].Attributes) != 3 {
		panic(fmt.Sprintf("unvalid json content, %s", content))
	}
	hash := r.Result.DeliverTx.Events[0].Attributes[2].Value
	return hash
}

func readEventRet() (string, error) {
	data := readFile(*appHashFile)
	return string(data), nil
}
