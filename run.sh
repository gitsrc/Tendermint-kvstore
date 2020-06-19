#!/bin/sh

# building kvstore app
go build .

# starting app server
./kvstore &
APP_PID=`ps -ef | grep "kvstore" | grep -v grep | awk '{print $2}'`
echo "kvstore running in pid $APP_PID\n"

sleep 3s


# starting tendermint as client
tendermint unsafe_reset_all
tendermint node  >node.log 2>node.log &
NODE_PID=`ps -ef | grep "tendermint" | grep -v grep | awk '{print $2}'`
echo "tendermint node running in pid $NODE_PID\n"

sleep 3s

# write kv pair to chain
echo "writing (name, cosmos) to blockchain ..."
echo `curl -s 'localhost:26657/broadcast_tx_commit?tx="name=cosmos"'`
echo "writing (name, cosmos) to blockchain done\n"

sleep 3s

# write kv pair to chain
echo "writing (token, atom) to blockchain ..."
echo `curl -s 'localhost:26657/broadcast_tx_commit?tx="token=atom"'`
echo "writing (token, atom) to blockchain done\n"

sleep 3s

KEY=name
VALUE=cosmos

# query kv pair
# curl -s 'localhost:26657/abci_query?data="'"$KEY"'"'
# sleep 3s

# query apphash
# curl -s 'localhost:26657/commit'
# sleep 3s

APPHASH=`curl -s 'localhost:26657/commit' | jq '.result.signed_header.header.app_hash' | sed 's/"//g'`
echo "apphash is $APPHASH\n"
sleep 3s

PROOF=`curl -s 'localhost:26657/abci_query?data="'"$KEY"'"' | jq .result.response.proof | tr -d " \t\n\r"`
echo "proof for ($KEY, $VALUE) is:\n$PROOF\n"
sleep 3s

# building proof verification tool
cd verify
# verify proof related with kv pair
echo "verify ($KEY, $VALUE) with above apphash and proof"

./verify -key=$KEY  -value=$VALUE -root=$APPHASH -proof=$PROOF


echo "cleaning up"
cd ..
rm -rf data
rm node.log
kill -9 $APP_PID
kill -9 $NODE_PID
