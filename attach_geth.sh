#!/bin/bash

NODE_NUMBER=$1
REMAINING_ARGS=${@:2}

re_number=^[0-9]+$
if ! [[ $NODE_NUMBER =~ $re_number ]]; then
  echo "First argument needs to be a number!"
  exit 1
fi

PHOTON_CMD=/usr/bin/geth
DATA_DIR="node$NODE_NUMBER"
IPC_PATH="ipc:$DATA_DIR/geth.ipc"
RPC="rpc:http://127.0.0.1:$((30310+$NODE_NUMBER))"

$PHOTON_CMD attach $IPC_PATH
