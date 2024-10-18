package kclient

import (
	"os"
	"testing"
)

func Test_ClientInit(t *testing.T) {
	os.Setenv("ENV_ROLE", "test")
	MustNewSingleClientConf("./", "client", "yaml")
}

func Test_ClientsInit(t *testing.T) {
	os.Setenv("ENV_ROLE", "test")
	MustNewMultiClientConf("./", "clients", "yaml")
}
