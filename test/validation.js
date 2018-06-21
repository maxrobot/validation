// Copyright (c) 2016-2018 Clearmatics Technologies Ltd
// SPDX-License-Identifier: LGPL-3.0+

const Util = require('ethereumjs-util');
const Web3 = require('web3');
const Web3Utils = require('web3-utils');
const Web3Abi = require('web3-eth-abi');
const Web3Accounts = require('web3-eth-accounts');
const rlp = require('rlp');

const Validation = artifacts.require("Validation");
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

contract.only('validation.js', (accounts) => {
  const joinHex = arr => '0x' + arr.map(el => el.slice(2)).join('')

  const validators = ["0x2be5ab0e43b6dc2908d5321cf318f35b80d0c10d", "0x8671e5e08d74f338ee1c462340842346d797afd3"];

  it('Test: Recover', async () => {
    const recover = await Recover.new();
    // const getOwnerReceipt = await recover.GetOwner();

    const signer = accounts[0];

    const hashData = Web3Utils.sha3("Test Data");

    console.log(Web3Utils.sha3("Test Data"))
    console.log(bytesToHex(Util.sha3("Test Data")))
    // const sig = web3.eth.sign(signer, Web3Utils.sha3("\x19Ethereum Signed Message:\n" + hashData.length +  hashData));

    const sig = web3.eth.sign(signer, hashData);
    const newSig = Util.fromRpcSig(sig);
    const ecrecover = Util.ecrecover(hashData, newSig.v, newSig.r, newSig.s);
    console.log(ecrecover)
    // const prefixHashData = Web3Utils.sha3("\x19Ethereum Signed Message:\n" + hashData.length +  hashData)
    //
    // const ecrecoveryReceipt = await recover.VerifyHash(prefixHashData, sig);
    // const ecrecoveryExpected = ecrecoveryReceipt.logs[0].args['owner'];
    // assert.equal(ecrecoveryExpected, signer);
  })

  it('Test: Recover', async () => {
    const recover = await Recover.new();
    // const getOwnerReceipt = await recover.GetOwner();

    const signer = accounts[0];

    const hashData = Web3Utils.sha3("Test Data");

    // const sig = web3.eth.sign(signer, Web3Utils.sha3("\x19Ethereum Signed Message:\n" + hashData.length +  hashData));

    const sig = web3.eth.sign(signer, hashData);
    const prefixHashData = Web3Utils.sha3("\x19Ethereum Signed Message:\n" + hashData.length +  hashData)

    const ecrecoveryReceipt = await recover.VerifyHash(prefixHashData, sig);
    const ecrecoveryExpected = ecrecoveryReceipt.logs[0].args['owner'];
    assert.equal(ecrecoveryExpected, signer);
  })

  it('Test: GetValidators()', async () => {
    const validation = await Validation.new(validators);
    const accounts = web3.eth.accounts;
    const signer = accounts[0];

    const validatorsReceipt = await validation.GetValidators();
    assert.equal(validators[0], validatorsReceipt[0])
  })

  it.only('Test: Authentic Submission - ValidateBlock()', async () => {
    const validation = await Validation.new(validators);
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

    const ecrecoveryReceipt = await validation.ValidateBlock(encodedBlockHeader, prefixHeader, prefixExtraData);
    const recoveredBlockHash = ecrecoveryReceipt.logs[0].args['blockHash'];
    const recoveredHash = ecrecoveryReceipt.logs[1].args['test'];
    const recoveredSignature = ecrecoveryReceipt.logs[2].args['owner'];
    assert.equal(block.hash, recoveredBlockHash)
    assert.equal(recoveredSignature, signer);

    console.log("Signature Hash Solidity: \n", recoveredHash)
    console.log("Signature Solidity: \n", extraDataSignature)
    console.log("Signer Solidity: \n", recoveredSignature)
  })

  it.only('Test: Authentic Submission Off-Chain - ValidateBlock()', async () => {
    const validation = await Validation.new(validators);
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

    // Create new signed hash
    const extraBytes = hexToBytes(extraData);
    const extraBytesShort = extraBytes.splice(1, extraBytes.length-66);
    const extraDataSignature = '0x' + bytesToHex(extraBytes.splice(extraBytes.length-65));
    const extraDataShort = '0x' + bytesToHex(extraBytesShort);

    // Make some changes to the block
    const newTxHash = Web3Utils.sha3("Test Data");
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

    // Encode and sign the new header
    const encodedHeader = '0x' + rlp.encode(header).toString('hex');
    const headerHash = Util.sha3(encodedHeader);

    const privateKey = Buffer.from('e176c157b5ae6413726c23094bb82198eb283030409624965231606ec0fbe65b', 'hex')

    const sig = Util.ecsign(headerHash, privateKey)
    if (this._chainId > 0) {
      sig.v += this._chainId * 2 + 8
    }

    const pubKey  = Util.ecrecover(headerHash, sig.v, sig.r, sig.s);
    const addrBuf = Util.pubToAddress(pubKey);
    const addr    = Util.bufferToHex(addrBuf);

    console.log(web3.eth.accounts[0],  addr);
    // Append signature to the end of extraData
    // const sigBytes = hexToBytes(sig);
    // const newExtraDataBytes = extraBytesShort.concat(sigBytes);
    // const newExtraData = '0x' + bytesToHex(newExtraDataBytes);
    //
    // const newBlockHeader = [
    //   parentHash,
    //   sha3Uncles,
    //   coinbase,
    //   root,
    //   newTxHash,
    //   receiptHash,
    //   logsBloom,
    //   difficulty,
    //   number,
    //   gasLimit,
    //   gasUsed,
    //   timestamp,
    //   newExtraData,
    //   mixHash,
    //   nonce
    // ];
    //
    // const encodedBlockHeader = '0x' + rlp.encode(newBlockHeader).toString('hex');
    // const blockHeaderHash = Web3Utils.sha3(encodedBlockHeader);
    //
    // // The new prefixes should be calculated off chain
    // const prefixHeader = '0x0214';
    // const prefixExtraData = '0xa0';
    //
    // const ecrecoveryReceipt = await validation.ValidateBlock(encodedBlockHeader, prefixHeader, prefixExtraData);
    // const recoveredBlockHash = ecrecoveryReceipt.logs[0].args['blockHash'];
    // const recoveredSignature = ecrecoveryReceipt.logs[1].args['owner'];
    // assert.equal(block.hash, recoveredBlockHash)
    // assert.equal(recoveredSignature, signer);

  })
});
