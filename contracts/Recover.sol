// Copyright (c) 2016-2018 Clearmatics Technologies Ltd
// SPDX-License-Identifier: LGPL-3.0+
pragma solidity ^0.4.18;

import "./ECVerify.sol";

contract Recover {
	address Owner;

	event broadcastSig(address owner);

	constructor () public {
		Owner = msg.sender;
	}

	function VerifyData(bytes32 data, bytes sig) public {
		address sig_addr = ECVerify.ecrecovery(data, sig);

		emit broadcastSig(sig_addr);

	}

  function GetOwner() public {
    emit broadcastSig(Owner);
  }
}
