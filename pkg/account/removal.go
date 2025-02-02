package account

import (
	"fmt"
	"os"
	"path"

	"github.com/astra-net/go-sdk/pkg/common"
	"github.com/astra-net/go-sdk/pkg/store"
	"github.com/mitchellh/go-homedir"
)

// RemoveAccount - removes an account from the keystore
func RemoveAccount(name string) error {
	accountExists := store.DoesNamedAccountExist(name)

	if !accountExists {
		return fmt.Errorf("account %s doesn't exist", name)
	}

	uDir, _ := homedir.Dir()
	astraCLIDir := path.Join(uDir, common.DefaultConfigDirName, common.DefaultConfigAccountAliasesDirName)
	accountDir := fmt.Sprintf("%s/%s", astraCLIDir, name)
	os.RemoveAll(accountDir)

	return nil
}
