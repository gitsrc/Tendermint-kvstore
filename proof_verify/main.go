package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/tendermint/tendermint/crypto/merkle"

	"github.com/cosmos/cosmos-sdk/store/rootmulti"
)

func main() {
	//num, err := hex.DecodeString("2d")
	//fmt.Println(num[0])
	//if err != nil {
	//	panic(fmt.Sprintf("string convert to number failed: %v", err))
	//}
	//hash := DecodeCommitHash()
	//fmt.Println(hash)
	key, proofData := decodeKeyAndProof()
	fmt.Println("content :", key, ", ", proofData)
	return
	commit, err := hex.DecodeString("64")
	if err != nil {
		panic(err)
	}
	fmt.Println(commit)
	data := readFile("")
	verifyProof(data)
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

	ret := `{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "check_tx": {
      "code": 0,
      "data": null,
      "log": "",
      "info": "",
      "gasWanted": "0",
      "gasUsed": "0",
      "events": [],
      "codespace": ""
    },
    "deliver_tx": {
      "code": 0,
      "data": null,
      "log": "",
      "info": "",
      "gasWanted": "0",
      "gasUsed": "0",
      "events": [
        {
          "type": "app",
          "attributes": [
            {
              "key": "Y3JlYXRvcg==",
              "value": "Q29zbW9zaGkgTmV0b3dva28=",
              "index": false
            },
            {
              "key": "a2V5",
              "value": "YWJj",
              "index": false
            },
            {
              "key": "aGFzaA==",
              "value": "27FccTZfxmBjuxOXh84B3M2ATOoqmFYMr7XsuQ8f7es=",
              "index": false
            }
          ]
        }
      ],
      "codespace": ""
    },
    "hash": "095439D2B2EB4BAEEE3E1804B48536AE4E6476D4FE10ECE2D621952975FBA5E3",
    "height": "8"
  }
}`
	return ret, nil
}

func decodeKeyAndProof() (string, []byte) {
	data := readProof()
	if len(data) != 8 {
		panic(fmt.Sprintf("invalid proof data, %v ", data))
	}
	key := strings.TrimSpace(strings.Split(data[3], "key:")[1])
	proof := decodeProof(data[len(data)-1])
	return key, proof
}

func decodeProof(proofData string) []byte {
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
	return result
}

func readProof() []string {
	proofStr := readFile("")
	rets := strings.Split(string(proofStr), "->")
	return rets
}

func readFile(filePath string) []byte {
	proofStr := `-> code: OK
	-> height: 54
	-> key: abc
	-> key.hex: 616263
	-> value: abc
	-> value.hex: 616263
	-> proof: &merkle.Proof{Ops:[]merkle.ProofOp{merkle.ProofOp{Type:"iavl:v", Key:[]uint8{0x61, 0x62, 0x63}, Data:[]uint8{0x2d, 0xa, 0x2b, 0x1a, 0x29, 0xa, 0x3, 0x61, 0x62, 0x63, 0x12, 0x20, 0xba, 0x78, 0x16, 0xbf, 0x8f, 0x1, 0xcf, 0xea, 0x41, 0x41, 0x40, 0xde, 0x5d, 0xae, 0x22, 0x23, 0xb0, 0x3, 0x61, 0xa3, 0x96, 0x17, 0x7a, 0x9c, 0xb4, 0x10, 0xff, 0x61, 0xf2, 0x0, 0x15, 0xad, 0x18, 0xf}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}}, XXX_NoUnkeyedLiteral:struct {}{}, XXX_unrecognized:[]uint8(nil), XXX_sizecache:0}`
	return []byte(proofStr)
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
