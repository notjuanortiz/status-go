package statusaccounts

import (
	"errors"
	"fmt"

	"github.com/pborman/uuid"
	staccount "github.com/status-im/status-go/account"
	"github.com/status-im/status-go/extkeys"
)

var (
	errAccountNotFoundByID          = errors.New("account not found")
	errAccountCannotDeriveChildKeys = errors.New("selected account cannot derive child keys")
	errAccountManagerNotSet         = errors.New("account manager not set")
)

type generator struct {
	am       *staccount.Manager
	accounts map[string]*account
}

func newGenerator() *generator {
	return &generator{
		accounts: make(map[string]*account),
	}
}

func (g *generator) setAccountManager(am *staccount.Manager) {
	g.am = am
}

func (g *generator) generate(mnemonicPhraseLength int, n int) ([]CreatedAccountInfo, error) {
	entropyStrength, err := mnemonicPhraseLengthToEntropyStrenght(mnemonicPhraseLength)
	if err != nil {
		return nil, err
	}

	infos := make([]CreatedAccountInfo, 0)

	for i := 0; i < n; i++ {
		mnemonic := extkeys.NewMnemonic()
		mnemonicPhrase, err := mnemonic.MnemonicPhrase(entropyStrength, extkeys.EnglishLanguage)
		if err != nil {
			return nil, fmt.Errorf("can not create mnemonic seed: %v", err)
		}

		info, err := g.importMnemonic(mnemonicPhrase)
		if err != nil {
			return nil, err
		}

		infos = append(infos, info)
	}

	return infos, err
}

func (g *generator) importMnemonic(mnemonicPhrase string) (CreatedAccountInfo, error) {
	mnemonic := extkeys.NewMnemonic()
	masterExtendedKey, err := extkeys.NewMaster(mnemonic.MnemonicSeed(mnemonicPhrase, ""))
	if err != nil {
		return CreatedAccountInfo{}, fmt.Errorf("can not create master extended key: %v", err)
	}

	acc := &account{
		privateKey:  masterExtendedKey.ToECDSA(),
		extendedKey: masterExtendedKey,
	}

	id := uuid.NewRandom().String()
	g.accounts[id] = acc

	return acc.toCreatedAccountInfo(id, mnemonicPhrase), nil
}

func (g *generator) findAccount(accountID string) (*account, error) {
	acc, ok := g.accounts[accountID]
	if !ok {
		return nil, errAccountNotFoundByID
	}

	return acc, nil
}

func (g *generator) deriveAddresses(accountID string, pathStrings []string) (map[string]AccountInfo, error) {
	acc, err := g.findAccount(accountID)
	if err != nil {
		return nil, err
	}

	pathAccounts, err := g.deriveChildAccounts(acc, pathStrings)
	if err != nil {
		return nil, err
	}

	pathAccountsInfo := make(map[string]AccountInfo)

	for pathString, childAccount := range pathAccounts {
		pathAccountsInfo[pathString] = childAccount.toAccountInfo()
	}

	return pathAccountsInfo, nil
}

func (g *generator) storeAccount(accountID string, password string) (AccountInfo, error) {
	if g.am == nil {
		return AccountInfo{}, errAccountManagerNotSet
	}

	acc, err := g.findAccount(accountID)
	if err != nil {
		return AccountInfo{}, err
	}

	return g.store(acc, password)
}

func (g *generator) storeDerivedAccounts(accountID string, password string, pathStrings []string) (map[string]AccountInfo, error) {
	if g.am == nil {
		return nil, errAccountManagerNotSet
	}

	acc, err := g.findAccount(accountID)
	if err != nil {
		return nil, err
	}

	pathAccounts, err := g.deriveChildAccounts(acc, pathStrings)
	if err != nil {
		return nil, err
	}

	pathAccountsInfo := make(map[string]AccountInfo)

	for pathString, childAccount := range pathAccounts {
		info, err := g.store(childAccount, password)
		if err != nil {
			return nil, err
		}

		pathAccountsInfo[pathString] = info
	}

	return pathAccountsInfo, nil
}

func (g *generator) loadAccount(address string, password string) (IdentifiedAccountInfo, error) {
	if g.am == nil {
		return IdentifiedAccountInfo{}, errAccountManagerNotSet
	}

	_, key, err := g.am.AddressToDecryptedAccount(address, password)
	if err != nil {
		return IdentifiedAccountInfo{}, err
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return IdentifiedAccountInfo{}, err
	}

	id := uuid.String()

	account := &account{
		privateKey:  key.PrivateKey,
		extendedKey: key.ExtendedKey,
	}

	g.accounts[id] = account

	return IdentifiedAccountInfo{
		AccountInfo: account.toAccountInfo(),
		ID:          id,
	}, nil
}

func (g *generator) deriveChildAccounts(acc *account, pathStrings []string) (map[string]*account, error) {
	pathAccounts := make(map[string]*account)

	for _, pathString := range pathStrings {
		childAccount, err := g.deriveChildAccount(acc, pathString)
		if err != nil {
			return pathAccounts, err
		}

		pathAccounts[pathString] = childAccount
	}

	return pathAccounts, nil
}

func (g *generator) deriveChildAccount(acc *account, pathString string) (*account, error) {
	_, path, err := decodePath(pathString)
	if err != nil {
		return nil, err
	}

	if acc.extendedKey == nil {
		return nil, errAccountCannotDeriveChildKeys
	}

	childExtendedKey, err := acc.extendedKey.Derive(path)
	if err != nil {
		return nil, err
	}

	return &account{
		privateKey:  childExtendedKey.ToECDSA(),
		extendedKey: childExtendedKey,
	}, nil
}

func (g *generator) store(acc *account, password string) (AccountInfo, error) {
	address, publicKey, err := g.am.ImportSingleExtendedKey(acc.extendedKey, password)
	if err != nil {
		return AccountInfo{}, err
	}

	g.reset()

	return AccountInfo{
		Address:   address,
		PublicKey: publicKey,
	}, nil
}

func (g *generator) reset() {
	g.accounts = make(map[string]*account)
}
