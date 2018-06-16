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

  // This won't work because the ECVerify contract is not appending the ethereum message
  // it('Test: Recover', async () => {
  //   const recover = await Recover.new();
  //   // const getOwnerReceipt = await recover.GetOwner();
  //
  //   const signer = accounts[0];
  //
  //   const hashData = Web3Utils.sha3("Test Data");
  //
  //   const sig = web3.eth.sign(signer, hashData);
  //
  //   const ecrecoveryReceipt = await recover.VerifyHash(hashData, sig);
  //   const ecrecoveryExpected = ecrecoveryReceipt.logs[0].args['owner'];
  //   assert.equal(ecrecoveryExpected, signer);
  // })

  it('Test: VerifyBlockHash()', async () => {
    const recover = await Recover.new();
    const accounts = web3.eth.accounts;
    const signer = accounts[0];

    // Get a single block
    const block = web3.eth.getBlock(10);

    // Decompose the values in the block to hash
    const parentHash = block.parentHash;
    const sha3Uncles = block.sha3Uncles;
    const coinbase = block.miner;
    const root = block.stateRoot;
    const txHash = block.transactionsRoot;
    const receiptHash = block.receiptsRoot;
    const logsBloom = block.logsBloom;
    const difficulty = Web3Utils.toBN(block.difficulty);
    const number = Web3Utils.toBN(block.number);
    const gasLimit = block.gasLimit;
    const gasUsed = block.gasUsed;
    const timestamp = Web3Utils.toBN(block.timestamp);
    const extraData = block.extraData;
    const mixHash = block.mixHash;
    const nonce = block.nonce;

    const header = [
      parentHash,
      sha3Uncles,
      coinbase,
      root,
      txHash,
      receiptHash,
      logsBloom,
      difficulty,
      number,
      gasLimit,
      gasUsed,
      timestamp,
      extraData,
      mixHash,
      nonce
    ];

    const encodedHeader = rlp.encode(header);

    const headerHash = Web3Utils.sha3(encodedHeader);
    assert.equal(block.hash, headerHash);
  })

  it('Test: VerifySignedHash()', async () => {
    const recover = await Recover.new();
    const accounts = web3.eth.accounts;
    const signer = accounts[0];

    // Get a single block
    const block = web3.eth.getBlock(10);
    console.log(block);

    // Decompose the values in the block to hash
    const parentHash = block.parentHash;
    const sha3Uncles = block.sha3Uncles;difficulty
    const coinbase = block.miner;
    const root = block.stateRoot;
    const txHash = block.transactionsRoot;
    const receiptHash = block.receiptsRoot;
    const logsBloom = block.logsBloom;
    const difficulty = Web3Utils.toBN(block.difficulty);
    const number = Web3Utils.toBN(block.number);
    const gasLimit = block.gasLimit;
    const gasUsed = block.gasUsed;
    const timestamp = Web3Utils.toBN(block.timestamp);
    const extraData = block.extraData;
    const mixHash = block.mixHash;
    const nonce = block.nonce;

    // Remove last 65 Bytes of extraData
    const extraBytes = hexToBytes(extraData);
    const extraBytesShort = extraBytes.splice(1, extraBytes.length-66);
    const extraDataSignature = '0x' + bytesToHex(extraBytes.splice(extraBytes.length-65));
    const extraDataShort = '0x' + bytesToHex(extraBytesShort);

    const header = [
      parentHash,
      sha3Uncles,
      coinbase,
      root,
      txHash,
      receiptHash,
      logsBloom,
      difficulty,
      number,
      gasLimit,
      gasUsed,
      timestamp,
      extraDataShort,
      mixHash,
      nonce
    ];

    const encodedHeader = rlp.encode(header);

    const headerHash = Web3Utils.sha3(encodedHeader);

    const ecrecoveryReceipt = await recover.VerifyHash(headerHash, extraDataSignature);
    const ecrecoveryExpected = ecrecoveryReceipt.logs[0].args['owner'];
    assert.equal(ecrecoveryExpected, signer);
  })

  it('Test: VerifyBlock()', async () => {
    const recover = await Recover.new();
    const accounts = web3.eth.accounts;
    const signer = accounts[0];

    // Get a single block
    const block = web3.eth.getBlock(10);

    // Decompose the values in the block to hash
    const parentHash = block.parentHash;
    const sha3Uncles = block.sha3Uncles;
    const coinbase = block.miner;
    const root = block.stateRoot;
    const txHash = block.transactionsRoot;
    const receiptHash = block.receiptsRoot;
    const logsBloom = block.logsBloom;
    const difficulty = Web3Utils.toBN(block.difficulty);
    const number = Web3Utils.toBN(block.number);
    const gasLimit = block.gasLimit;
    const gasUsed = block.gasUsed;
    const timestamp = Web3Utils.toBN(block.timestamp);
    const extraData = block.extraData;
    const mixHash = block.mixHash;
    const nonce = block.nonce;

    // Remove last 65 Bytes of extraData
    const extraBytes = hexToBytes(extraData);
    const extraBytesShort = extraBytes.splice(1, extraBytes.length-66);
    const extraDataSignature = '0x' + bytesToHex(extraBytes.splice(extraBytes.length-65));
    const extraDataShort = '0x' + bytesToHex(extraBytesShort);

    const header = [
      parentHash,
      sha3Uncles,
      coinbase,
      root,
      txHash,
      receiptHash,
      logsBloom,
      difficulty,
      number,
      gasLimit,
      gasUsed,
      timestamp,
      extraDataShort,
      mixHash,
      nonceblockH
    ];

    const encodedHeader = '0x' + rlp.encode(header).toString('hex');
    const headerHash = Web3Utils.sha3(encodedHeader);

    const ecrecoveryReceipt = await recover.VerifyBlock(encodedHeader, extraDataSignature);
    const recoveredParentHash = ecrecoveryReceipt.logs[0].args['parentHash']
    const recoveredRootHash = ecrecoveryReceipt.logs[0].args['rootHash']
    const ecrecoveryExpected = ecrecoveryReceipt.logs[1].args['owner'];
    assert.equal(recoveredParentHash, parentHash);
    assert.equal(recoveredRootHash, root);
    assert.equal(ecrecoveryExpected, signer);
  })

  it.only('Test: ExtractSignature()', async () => {
    const recover = await Recover.new();
    const accounts = web3.eth.accounts;
    const signer = accounts[0];

    // Get a single block
    const block = web3.eth.getBlock(10);

    // Decompose the values in the block to hash
    const parentHash = block.parentHash;
    const sha3Uncles = block.sha3Uncles;
    const coinbase = block.miner;
    const root = block.stateRoot;
    const txHash = block.transactionsRoot;
    const receiptHash = block.receiptsRoot;
    const logsBloom = block.logsBloom;
    const difficulty = Web3Utils.toBN(block.difficulty);
    const number = Web3Utils.toBN(block.number);
    const gasLimit = block.gasLimit;
    const gasUsed = block.gasUsed;
    const timestamp = Web3Utils.toBN(block.timestamp);
    const extraData = block.extraData;
    const mixHash = block.mixHash;
    const nonce = block.nonce;

    // Remove last 65 Bytes of extraData
    const extraBytes = hexToBytes(extraData);
    const extraBytesShort = extraBytes.splice(1, extraBytes.length-66);
    const extraDataSignature = '0x' + bytesToHex(extraBytes.splice(extraBytes.length-65));
    const extraDataShort = '0x' + bytesToHex(extraBytesShort);

    const blockHeader = [
      parentHash,
      sha3Uncles,
      coinbase,
      root,
      txHash,
      receiptHash,
      logsBloom,
      difficulty,
      number,
      gasLimit,
      gasUsed,
      timestamp,
      extraData,
      mixHash,
      nonce
    ];

    const header = [
      parentHash,
      sha3Uncles,
      coinbase,
      root,
      txHash,
      receiptHash,
      logsBloom,
      difficulty,
      number,
      gasLimit,
      gasUsed,
      timestamp,
      extraDataShort,
      mixHash,
      nonce
    ];

    const encodedBlockHeader = '0x' + rlp.encode(blockHeader).toString('hex');
    const blockHeaderHash = Web3Utils.sha3(encodedBlockHeader);
    assert.equal(block.hash, blockHeaderHash);

    const encodedHeader = '0x' + rlp.encode(header).toString('hex');
    const headerHash = Web3Utils.sha3(encodedHeader);

    // The new prefixes should be calculated off chain
    const prefixHeader = '0x0214';
    const prefixExtraData = '0xa0';

    const ecrecoveryReceipt = await recover.ExtractHash(encodedBlockHeader, encodedHeader, prefixHeader, prefixExtraData);
    const recoveredBlockHash = ecrecoveryReceipt.logs[0].args['blockHash'];
    assert.equal(block.hash, recoveredBlockHash)
    console.log(encodedHeader);
    console.log(ecrecoveryReceipt.logs[1].args['header']);
    console.log(ecrecoveryReceipt.logs[2].args['header']);
    // console.log(ecrecoveryReceipt.logs[2].args['header']);
    // console.log(ecrecoveryReceipt.logs[3].args['header']);
    // console.log(headerHash);
    // console.log(ecrecoveryReceipt.logs[1].args['blockHash']);
    // console.log(signer);
    // console.log(ecrecoveryReceipt.logs[2].args['owner']);
    // assert.equal(recoveredParentHash, parentHash);
    // assert.equal(recoveredRootHash, root);
    // assert.equal(ecrecoveryExpected, signer);
  })
});
