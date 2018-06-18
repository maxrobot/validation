// Copyright (c) 2018 Clearmatics Technologies Ltd

package config_test

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"strings"
	"testing"

	"gitlab.clearmatics.net/dev/boe-poc/src/config"
)

func Test_ParseParameters_ValidUserFile(t *testing.T) {
	path := findPath() + "../user.json"
	commandLine := []string{path}

	setup := config.ParseParameters(commandLine)

	assert.Equal(t, "http://rt-poc2-2.azurewebsites.net", setup.APIURL)
	assert.Equal(t, "password", setup.Grant_type)
	assert.Equal(t, "svcp93lXc&", setup.Password)
	assert.Equal(t, "user1@scheme1.co.uk", setup.User)
}

func findPath() string {
	_, path, _, _ := runtime.Caller(0)
	pathSlice := strings.Split(path, "/")
	return strings.Trim(path, pathSlice[len(pathSlice)-1])
}
