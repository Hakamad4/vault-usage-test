package main

import (
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"log"
)

type VaultParameters struct {
	// connection parameters
	address             string
	approleRoleID       string
	approleSecretIDFile string
	env                 string

	// the locations / field names of our two secrets
	apiKeyPath              string
	apiKeyMountPath         string
	apiKeyField             string
	databaseCredentialsPath string
}

type Vault struct {
	client *vault.Client
	params *VaultParameters
}

func (v *Vault) NewVaultClient(params *VaultParameters) (*Vault, *vault.Secret, error) {
	log.Printf("connecting to vault @ %s", params.address)

	config := vault.DefaultConfig()
	config.Address = params.address

	client, err := vault.NewClient(config)
	if err != nil {
		return nil, nil, err
	}

	v.client = client
	v.params = params

	err = v.login()
	if err != nil {
		return nil, nil, err
	}

	return v, nil, nil
}

func (v *Vault) login() error {
	if v.params.env == "local" {
		v.client.SetToken("myroot")
	} else {
		log.Panicf("not implemented yet")
	}
	return nil
}

func ReadVaultConfigs() {
	params := VaultParameters{
		address: "http://localhost:8200",
	}

	Vault := new(Vault)
	vaultClient, _, _ := Vault.NewVaultClient(&params)

	// Leitura dos segredos
	secret, err := vaultClient.client.Logical().Read("secret/data/mongodb")
	if err != nil {
		log.Panicf("Erro ao ler segredos: %s", err)
		return
	}

	if secret == nil {
		log.Fatalf("Nenhum segredo encontrado")
	}
	if secret.Warnings != nil {
		log.Printf("Aviso: %s", secret.Warnings)
	}

	data := secret.Data["data"].(map[string]interface{})
	mongoURL := data["MONGO_URL"].(string)
	mongoDatabase := data["MONGO_DATABASE"].(string)
	mongoPass := data["MONGO_PASS"].(string)

	// Usar os segredos conforme necessário
	fmt.Printf("MONGO_URL: %s\n", mongoURL)
	fmt.Printf("MONGO_DATABASE: %s\n", mongoDatabase)
	fmt.Printf("MONGO_PASS: %s\n", mongoPass)
}
