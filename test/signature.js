// Copyright (c) 2016-2018 Clearmatics Technologies Ltd
// SPDX-License-Identifier: LGPL-3.0+

const Web3 = require('web3');
const Web3Utils = require('web3-utils');
const Web3Abi = require('web3-eth-abi');
const Web3Accounts = require('web3-eth-accounts');
const rlp = require('rlp');

const Recover = artifacts.require("Recover");


const web3 = new Web3();

web3.setProvider(new web3.providers.HttpProvider('http://localhost:8501'));

function hexToBytes(hex) {
    for (var bytes = [], c = 0; c < hex.length; c += 2)
    bytes.push(parseInt(hex.substr(c, 2), 16));
    return bytes;
}

function bytesToHex(bytes) {
    for (var hex = [], i = 0; i < bytes.length; i++) {
        hex.push((bytes[i] >>> 4).toString(16));
        hex.push((bytes[i] & 0xF).toString(16));
    }
    return hex.join("");
}

contract('signatures.js', (accounts) => {
  const joinHex = arr => '0x' + arr.map(el => el.slice(2)).join('')

  it('Test: Recover', async () => {
    const recover = await Recover.new();
    const getOwnerReceipt = await recover.GetOwner();

    const coinbase = accounts[0];

    const hashData = Web3Utils.sha3("Test Data");

    const sig = web3.eth.sign(coinbase, hashData);

    const ecrecoveryReceipt = await recover.VerifyData(hashData, sig);
    const ecrecoveryExpected = ecrecoveryReceipt.logs[0].args['owner'];
    assert.equal(coinbase, ecrecoveryExpected);
  })

  it.only('test web3', async () => {
    const recover = await Recover.new();
    const accounts = web3.eth.accounts;

    // Get a single block
    const block = web3.eth.getBlock(12);

    // Decompose the values in the block to hash
    const parentHash = block.parentHash;
    const sha3Uncles = block.sha3Uncles;
    const coinbase = accounts[0]
    const root = block.stateRoot;
    const txHash = block.transactionsRoot;
    const receiptHash = block.receiptsRoot;
    const logsBloom = block.logsBloom;
    const difficulty = block.difficulty;
    const number = block.number;
    const gasLimit = block.gasLimit;
    const gasUsed = block.gasUsed;
    const timestamp = block.timestamp;
    const extraData = block.extraData;
    const mixHash = block.mixHash;
    const nonce = block.nonce;

    // Remove last 65 Bytes of extraData
    const extraBytes = hexToBytes(extraData);
    const extraBytesShort = extraBytes.splice(1, extraBytes.length-66);
    const extraDataSignature = '0x' + bytesToHex(extraBytes.splice(extraBytes.length-65));
    const extraDataShort = '0x' + bytesToHex(extraBytesShort);

    const header = [
      parentHash
    ];

    // const mergedHeader = joinHex(header);
    const encodedHeader = rlp.encode(header);

    const headerHash = Web3Utils.sha3(encodedHeader);
    console.log(headerHash)

    const sig = web3.eth.sign(coinbase, headerHash)

    const ecrecoveryReceipt = await recover.VerifyData(headerHash, extraDataSignature);
    const ecrecoveryExpected = ecrecoveryReceipt.logs[0].args['owner'];
    assert.equal(coinbase, ecrecoveryExpected);
  })
});
