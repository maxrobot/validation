#!/bin/bash

NODE_NUMBER=$1
REMAINING_ARGS=${@:2}

re_number=^[0-9]+$
if ! [[ $NODE_NUMBER =~ $re_number ]]; then
  echo "First argument needs to be a number!"
  exit 1
fi

GENESIS_FILE="genesis.json"
PHOTON_CMD=/usr/bin/geth
IDENTITY="node_$NODE_NUMBER"
RPC_PORT=$((8500+$NODE_NUMBER))
RPC_ADDR="127.0.0.1"

GAS_PRICE=0
BLOCK_GAS_LIMIT=0xFFFFFFFF
DATA_DIR="node$NODE_NUMBER"
KEYSTORE_DIR="keystore"
BOOTNODES="enode://dcb1dbf8d710eb7d10e0e2db1e6d3370c4b048efe47c7a85c4b537b60b5c11832ef25f026915b803e928c1d93f01b853131e412c6308c4c6141d1504c78823c8@127.0.0.1:30310"
IPC_API="admin,eth,debug,miner,net,shh,txpool,personal,web3"
IPC_PATH="$DATA_DIR/geth.ipc"
NODE_PORT=$((30310+$NODE_NUMBER))
NETWORK_ID=1515
SYNCMODE="full"

#unlock addr and pass
ACCOUNT_ADDR_1="0x2be5ab0e43b6dc2908d5321cf318f35b80d0c10d"
ACCOUNT_PASS_1="password1"
ACCOUNT_ADDR_2="0x8671e5e08d74f338ee1c462340842346d797afd3"
ACCOUNT_PASS_2="password2"
if [ $NODE_NUMBER -eq 1 ]; then
  ACCOUNT_ADDR=$ACCOUNT_ADDR_1
  ACCOUNT_PASS=$ACCOUNT_PASS_1
  ETHERBASE=$ACCOUNT_ADDR_1
elif [ $NODE_NUMBER -eq 2 ]; then
  ACCOUNT_ADDR=$ACCOUNT_ADDR_2
  ACCOUNT_PASS=$ACCOUNT_PASS_2
  ETHERBASE=$ACCOUNT_ADDR_2
else
  ACCOUNT_ADDR=0
  ACCOUNT_PASS="xxx"
  ETHERBASE=0
fi

if [ ! -d $DATA_DIR ]; then
  echo "Creating data directory: $DATA_DIR"
  $PHOTON_CMD --datadir $DATA_DIR/ init $GENESIS_FILE
fi

$PHOTON_CMD \
  --syncmode $SYNCMODE \
  --port $NODE_PORT \
  --datadir $DATA_DIR/ \
  --ipcpath $IPC_PATH \
  --keystore $KEYSTORE_DIR \
  --rpc \
  --rpcport $RPC_PORT \
  --rpcaddr $RPC_ADDR \
  --rpccorsdomain "4444" \
  --rpcapi $IPC_API \
  --identity $IDENTITY \
  --bootnodes $BOOTNODES \
  --networkid $NETWORK_ID \
  --nodiscover \
  --mine \
  --gasprice $GAS_PRICE \
  --targetgaslimit $BLOCK_GAS_LIMIT \
  --unlock $ACCOUNT_ADDR \
  --password <(echo $ACCOUNT_PASS) \
  --etherbase $ETHERBASE \
  $REMAINING_ARGS
