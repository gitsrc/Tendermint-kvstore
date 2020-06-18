kvstore is a demo to tendermint app;

## build kvstore 

1. enter project root dir: `cd kvstore`
2. `go build .`

## build tendermint-abci

1. clone tendermint project: `git clone https://github.com/tendermint/tendermint.git`
2. enter project root dir: `cd tendermint`
3. `make tools && make install_abci`

 
## run program

1. run kvstore: `./kvstore`
2. run tendermint node: `tendermint unsafe_reset_all && tendermint node`
3. store value in app: `curl -s 'localhost:26657/broadcast_tx_commit?tx="abc=abc"'`

    ```
    {
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
                  "value": "yAbQ+KQFzlkd/f3FQS5YcjkkxS8GkSwkysk7WxRH67w=",
                  "index": false
                }
              ]
            }
          ],
          "codespace": ""
        },
        "hash": "095439D2B2EB4BAEEE3E1804B48536AE4E6476D4FE10ECE2D621952975FBA5E3",
        "height": "6"
      }
    }
    ```
4. query value in app: `curl -s 'localhost:26657/abci_query?data="abc"'`

   ```
   {
     "jsonrpc": "2.0",
     "id": -1,
     "result": {
       "response": {
         "code": 0,
         "log": "",
         "info": "",
         "index": "0",
         "key": "YWJj",
         "value": "YWJj",
         "proof": {
           "ops": [
             {
               "type": "iavl:v",
               "key": "YWJj",
               "data": "LgosGioKA2FiYxIgungWv48Bz+pBQUDeXa4iI7ADYaOWF3qctBD/YfIAFa0YthQ="
             }
           ]
         },
         "height": "2795",
         "codespace": ""
       }
     }
   }
   ```
4. query key proof in app: ` abci-cli query 0x616263 --path=/key --prove > proof.txt`
5. query rootHash in app: `abci-cli commit > hash.txt`


## build proof_verify

1. enter proof_verify dir: `cd proof_verify`
2. `go build .`
 
## verify proof

`./proof_verify -proof-file=../proof.txt -hash-file=../hash.txt`