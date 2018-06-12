# PoA Network
This repo allows you to set up a "proof of authority" ethereum blockchain with two separate peers who both engage in the sealing of blocks. The instructions are based on the tutorial of [Salanfe](https://hackernoon.com/setup-your-own-private-proof-of-authority-ethereum-network-with-geth-9a0a3750cda8) but has the more complicated parts already initialised.

## Launch the Network
In order to run the network first install an instance of [geth](https://geth.ethereum.org/downloads/) the directory structure and accounts have been set up a priori. Hence all that is required to launch the network is to follow these instructions:

### Initialise Nodes
```
$ geth --datadir node1/ init genesis.json
$ geth --datadir node2/ init genesis.json
```

### Launch the Bootnode
The boot node tells the peers how to connect with each other. In another terminal instance run:
```
$ bootnode -nodekey boot.key -verbosity 9 -addr :30310
$ INFO [06-07|12:16:21] UDP listener up                          self=enode://dcb1dbf8d710eb7d10e0e2db1e6d3370c4b048efe47c7a85c4b537b60b5c11832ef25f026915b803e928c1d93f01b853131e412c6308c4c6141d1504c78823c8@[::]:30310
```
As the peers communicate this terminal should fill with logs.

### Start and Attach to the Nodes
Each node must be launch, either as a background operation or on separate terminal instances. For node 1:
```
$ geth --datadir node1/ --syncmode 'full' --port 30311 --rpc --rpcaddr 'localhost' --rpcport 8501 --bootnodes 'enode://dcb1dbf8d710eb7d10e0e2db1e6d3370c4b048efe47c7a85c4b537b60b5c11832ef25f026915b803e928c1d93f01b853131e412c6308c4c6141d1504c78823c8@127.0.0.1:30310' --networkid 1515 --gasprice '1' -unlock '0x2be5ab0e43b6dc2908d5321cf318f35b80d0c10d' --password node1/password.txt --mine
```
then attach,
```
$ geth attach node1/geth.ipc
```
node 2:
```
$ geth --datadir node2/ --syncmode 'full' --port 30312 --rpc --rpcaddr 'localhost' --rpcport 8502 --bootnodes 'enode://dcb1dbf8d710eb7d10e0e2db1e6d3370c4b048efe47c7a85c4b537b60b5c11832ef25f026915b803e928c1d93f01b853131e412c6308c4c6141d1504c78823c8@127.0.0.1:30310' --networkid 1515 --gasprice '0' -unlock '0x8671e5e08d74f338ee1c462340842346d797afd3' --password node2/password.txt --mine
```
again attach,
```
$ geth attach node2/geth.ipc
```
Notice that IPC has been used to attach to the nodes, this allows the clique module to be used.


## Test
