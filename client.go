package vault

import (
	"fmt"
	"log"

	hashivault "github.com/hashicorp/vault/api"
)

func DefaultConfig() *hashivault.Config {
	fmt.Println("Testing ")
	return nil
	//return hashivault.DefaultConfig()
}

type Client struct {
	client *hashivault.Client
	Auth   *Auth
	KV2    *KV2
}

func NewClient(config *hashivault.Config) (*Client, error) {
	client, err := hashivault.NewClient(config)
	log.Println("TEST .........")
	log.Println("rESFSFS")
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
		Auth:   NewAuth(client),
		KV2:    &KV2{client: client}}, nil
}
