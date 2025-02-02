package common

import (
	"errors"
	"os"

	"github.com/astra-net/astra-network/accounts/keystore"
)

const (
	DefaultConfigDirName               = ".astra_cli"
	DefaultConfigAccountAliasesDirName = "account-keys"
	DefaultCommandAliasesDirName       = "command"
	DefaultPassphrase                  = ""
	JSONRPCVersion                     = "2.0"
	Secp256k1PrivateKeyBytesLength     = 32
)

var (
	ScryptN          = keystore.StandardScryptN
	ScryptP          = keystore.StandardScryptP
	DebugRPC         = false
	DebugTransaction = false
	ErrNotAbsPath    = errors.New("keypath is not absolute path")
	ErrBadKeyLength  = errors.New("Invalid private key (wrong length)")
	ErrFoundNoKey    = errors.New("found no bls key file")
	ErrFoundNoPass   = errors.New("found no passphrase file")
)

func init() {
	if _, enabled := os.LookupEnv("ASTRA_RPC_DEBUG"); enabled != false {
		DebugRPC = true
	}
	if _, enabled := os.LookupEnv("ASTRA_TX_DEBUG"); enabled != false {
		DebugTransaction = true
	}
	if _, enabled := os.LookupEnv("ASTRA_ALL_DEBUG"); enabled != false {
		EnableAllVerbose()
	}
}

// EnableAllVerbose sets debug vars to true
func EnableAllVerbose() {
	DebugRPC = true
	DebugTransaction = true
}
