// Copyright (c) 2016-2018 Clearmatics Technologies Ltd
// SPDX-License-Identifier: LGPL-3.0+
pragma solidity ^0.4.23;

import "./ECVerify.sol";

contract Recover {
	address Owner;

	event broadcastSig(address owner);
	event extraData(bytes header, bytes parentHash, bytes rootHash);

	constructor () public {
		Owner = msg.sender;
	}

	/*
	 * @param data  			data that has been signed
	 * @param sig    			signature of data
	 */
	function VerifyHash(bytes32 data, bytes sig) public {
		address sig_addr = ECVerify.ecrecovery(data, sig);

		emit broadcastSig(sig_addr);
	}

	/*
	* @param header  			header rlp encoded, with extraData signatures removed
	* @param sig    			extraData signatures
	*/
	function VerifyBlock(bytes header, bytes sig) public {
		bytes32 hashData = keccak256(header);
		address sig_addr = ECVerify.ecrecovery(hashData, sig);

		bytes memory parentHash = new bytes(32);
		bytes memory rootHash = new bytes(32);

		// get parentHash and rootHash
		extractData(parentHash, header, 4, 32);
		extractData(rootHash, header, 91, 32);

		emit extraData(header, parentHash, rootHash);
		emit broadcastSig(sig_addr);
	}

	/*
	* @param data	  			memory allocation for the data you need to extract
	* @param sig    			array from which the data should be extracted
	* @param start   			index which the data starts within the byte array
	* @param length  			total length of the data to be extracted
	*/
	function extractData(bytes data, bytes input, uint start, uint length) private pure {
		for (uint i=0; i<length; i++) {
			data[i] = input[start+i];
		}
	}

}
