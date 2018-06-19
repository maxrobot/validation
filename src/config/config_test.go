// Copyright (c) 2018 Clearmatics Technologies Ltd

package config_test

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"strings"
	"testing"

	"github.com/validation/src/config"
)

func Test_Read_ValidSetupJson(t *testing.T) {
	path := findPath() + "../setup.json"
	setup := config.Read(path)

	assert.Equal(t, "8501", setup.Port_to)
	assert.Equal(t, "127.0.0.1", setup.Addr_to)
	assert.Equal(t, "8502", setup.Port_from)
	assert.Equal(t, "127.0.0.1", setup.Addr_from)
}

func findPath() string {
	_, path, _, _ := runtime.Caller(0)
	pathSlice := strings.Split(path, "/")
	return strings.Trim(path, pathSlice[len(pathSlice)-1])
}
