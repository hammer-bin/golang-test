package main

import (
	vault "github.com/hashicorp/vault/api"
	"log"
)

func main() {
	config := vault.DefaultConfig()
	//config.Address = "http://127.0.0.1:8200"
	config.Address = "http://3.38.254.89:8200"
	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	//client.SetToken("hvs.ZlWL4QUKJQHt4yNWj6ejODNb")
	client.SetToken("hvs.v7GN8AmsJIteSHRW9UJjymiY")

	secretData := map[string]interface{}{
		"data": map[string]interface{}{
			"access": "1234SZP7123477E41234",
			"secret": "1234Og3tprzhSlbMu+12346OAw7i0c5uLNj31234",
		},
	}

	_, err = client.Logical().Write("secret/data/aws-secret-key", secretData)
	if err != nil {
		log.Fatalf("Unable to write secret: %v", err)
	}

	log.Println("Secret written successfully.")

	// Read a secret
	secret, err := client.Logical().Read("secret/data/aws-secret-key")
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatalf("Data type assertion failed: %T %#v", secret.Data["data"], secret.Data["data"])
	}
	for key, value := range data {
		log.Println(key, value)
	}

	value, ok := data["access"].(string)
	if !ok {
		log.Fatalf("Value type assertion failed: %T %#v", data["access"], data["access"])
	}

	if value != "1234SZP7123477E41234" {
		log.Fatalf("Unexpected password value %q retrieved from vault", value)
	}

	log.Println("Access granted")
}
