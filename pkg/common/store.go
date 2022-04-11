package common

import (
	"github.com/astra-net/astra-network/accounts/keystore"
)

func KeyStoreForPath(p string) *keystore.KeyStore {
	return keystore.NewKeyStore(p, ScryptN, ScryptP)
}
