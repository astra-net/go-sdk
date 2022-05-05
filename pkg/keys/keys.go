package keys

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/astra-net/astra-network/accounts/keystore"
	ethCommon "github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/crypto"

	homedir "github.com/mitchellh/go-homedir"
)

func checkAndMakeKeyDirIfNeeded() string {
	userDir, _ := homedir.Dir()
	astraCLIDir := path.Join(userDir, ".astra_cli", "keystore")
	if _, err := os.Stat(astraCLIDir); os.IsNotExist(err) {
		// Double check with Leo what is right file persmission
		os.Mkdir(astraCLIDir, 0700)
	}

	return astraCLIDir
}

func ListKeys(keystoreDir string) {
	astraCLIDir := checkAndMakeKeyDirIfNeeded()
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(astraCLIDir, scryptN, scryptP)
	// keystore.KeyStore
	allAccounts := ks.Accounts()
	fmt.Printf("Astra Address:%s File URL:\n", strings.Repeat(" ", ethCommon.AddressLength*2))
	for _, account := range allAccounts {
		fmt.Printf("%s\t\t %s\n", account.Address, account.URL)
	}
}

func AddNewKey(password string) {
	astraCLIDir := checkAndMakeKeyDirIfNeeded()
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	ks := keystore.NewKeyStore(astraCLIDir, scryptN, scryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		fmt.Printf("new account error: %v\n", err)
	}
	fmt.Printf("account: %s\n", account.Address)
	fmt.Printf("URL: %s\n", account.URL)
}
