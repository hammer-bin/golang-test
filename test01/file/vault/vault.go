package main

import (
	"github.com/hashicorp/vault/api"
	"log"
)

var (
	/*vaultAddr = os.Getenv("VAULT_ADDR")
	roleId    = os.Getenv("VAULT_ROLE_ID")
	secretId  = os.Getenv("VAULT_SECRET_ID")*/
	vaultAddr = "http://127.0.0.1:8200"
	roleId    = "f8bc2998-60a2-5545-c85a-4320c94431ef"
	secretId  = "382cc795-141c-ce97-f6a5-16c0b1c1bf01"
)

type Client struct {
	Token         string
	Client        *api.Client
	LeaseDuration int
}

type Token struct {
	Auth struct {
		ClientToken   string `json:"client_token"`
		LeaseDuration int    `json:"lease_duration"`
	} `json:"auth"`
}

type appRoleLogin struct {
	RoleID   string `json:"role_id"`
	SecretID string `json:"secret_id"`
}

func NewVaultClient() (*Client, error) {
	vaultClient := Client{}

	client, err := api.NewClient(&api.Config{
		Address: vaultAddr,
	})

	vaultClient.Client = client

	return &vaultClient, err
}

func (vaultClient *Client) AuthUser() (string, error) {
	// step: create the token request
	// URL: PUT http://127.0.0.1:8200/v1/auth/approle/login  :: app 로그인 API
	request := vaultClient.Client.NewRequest("POST", "/v1/auth/approle/login")
	login := appRoleLogin{SecretID: secretId, RoleID: roleId}

	if err := request.SetJSONBody(login); err != nil {
		return "", err
	}

	// step: make the request
	resp, err := vaultClient.Client.RawRequest(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// step: parse and return auth
	secret, err := api.ParseSecret(resp.Body)
	if err != nil {
		return "", err
	}

	return secret.Auth.ClientToken, nil
}

// Generic secret value reader
func ReadSecretValues(secretPath string, userDesignation string, passDesignation string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	client, err := NewVaultClient()

	if err != nil {
		return nil, err
	}

	token, err := client.AuthUser()

	if err != nil {
		return nil, err
	}

	client.Client.SetToken(token)
	secretValues, err := client.Client.Logical().Read(secretPath)
	data1, ok := secretValues.Data["data"].(map[string]interface{})
	if !ok {
		log.Printf("Data type assertion failed: %T\n", secretValues.Data["data"])
		log.Printf("Data type assertion failed: %#v\n", secretValues.Data["data"])
	}
	for key, value := range data1 {
		log.Println(key, value)
	}

	if err != nil || secretValues == nil {
		return nil, err
	}

	data["db_name"] = secretValues.Data[userDesignation]
	data["url"] = secretValues.Data[passDesignation]

	data2, _ := data["db_name"].(map[string]interface{})
	println(len(data2))

	return data, err
}
