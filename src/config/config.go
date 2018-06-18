// Copyright (c) 2018 Clearmatics Technologies Ltd

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Settings
type Setup struct {
	Port_to  		string `json:"rpc-port-to"`
	Addr_to  		string `json:"rpc-addr-to"`
	Port_from  	string `json:"rpc-port-from"`
	Addr_from  	string `json:"rpc-addr-from"`
}

func Read(config string) (setup Setup) {
	raw, err := ioutil.ReadFile(config)
	if err != nil {
		fmt.Print(err, "\n")
	}

	err = json.Unmarshal(raw, &setup)

	return setup
}
