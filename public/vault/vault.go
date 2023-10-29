package vault

import (
	"errors"

	vault "github.com/hashicorp/vault/api"
)

type IVault interface {
	ReadSecret(key string) ([]byte, error)
}

type Vault struct {
	Client *vault.Client
}

func NewVault(vaultAddress string, token string) (*Vault, error) {
	client, err := vault.NewClient(&vault.Config{
		Address: vaultAddress,
	})
	if err != nil {
		return nil, err
	}
	client.SetToken(token)
	return &Vault{
		Client: client,
	}, nil
}

func (v *Vault) ReadSecret(key string) ([]byte, error) {
	secret, err := v.Client.Logical().Read(key)
	if err != nil {
		return nil, err
	}
	if secret == nil || secret.Data == nil {
		return nil, errors.New("secret not found or empty")
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid secret data")
	}

	value, ok := data["value"].(string)
	if !ok {
		return nil, errors.New("invalid secret data")
	}

	if value == "" {
		return nil, errors.New("invalid secret data")
	}

	return []byte(value), nil
}

// Ensure that Vault implements IVault
var _ IVault = (*Vault)(nil)
