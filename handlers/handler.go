package handlers

import (
	"caturandi-labs/golang-starter/config"
	"caturandi-labs/golang-starter/ent"
)

type Handlers struct {
	Client *ent.Client
	Config *config.Config
}

func NewHandler(client *ent.Client, config *config.Config) *Handlers {
	return &Handlers{
		Client: client,
		Config: config,
	}
}