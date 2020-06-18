package main

import (
	"testing"
)

func TestDecodeCommitHash(t *testing.T) {
	//dec := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(
	//	"6ALSQoKTUxy0wMDvMcPHa9zdwNxn+0a9Mdl3BSJnQSI="))
	//res1, err := ioutil.ReadAll(dec)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(res1)

	//	content := `I[2020-06-18|13:46:51.109] Starting socketClient service                module=abci-client impl=socketClient
	//-> code: OK
	//-> data.hex: 0xE802D2428293531CB4C0C0EF31C3C76BDCDDC0DC67FB46BD31D9770522674122`
	//	data := strings.Split(content, "->")
	//	fmt.Println("data size : ", len(data), ", content : ", data)

	//hash, err := hex.DecodeString("E802D2428293531CB4C0C0EF31C3C76BDCDDC0DC67FB46BD31D9770522674122")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(hash)
	*help = false
	*proofFile = "proof.txt"
	DecodeKeyValAndProof()
}
