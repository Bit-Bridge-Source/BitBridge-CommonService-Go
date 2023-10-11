package service

import (
	"errors"

	vault "github.com/hashicorp/vault/api"
)

type IVaultService interface {
	ReadSecret(key string) ([]byte, error)
}

type VaultService struct {
	Vault *vault.Client
}

func (v *VaultService) ReadSecret(key string) ([]byte, error) {
	secret, err := v.Vault.Logical().Read(key)
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

func NewVaultService(vaultAddress, token string) (*VaultService, error) {
	config := &vault.Config{
		Address: vaultAddress,
	}

	vaultClient, err := vault.NewClient(config)
	if err != nil {
		panic(err)
	}

	vaultClient.SetToken(token)

	return &VaultService{
		Vault: vaultClient,
	}, nil
}

// Ensure that VaultService implements IVaultClient.
var _ IVaultService = (*VaultService)(nil)
