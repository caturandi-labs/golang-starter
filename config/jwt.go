package config

import "caturandi-labs/golang-starter/utils"

type JwtConfig struct {
	Secret []byte

}

func NewJwtConfig() *JwtConfig {
	return &JwtConfig{
		Secret: []byte(utils.GetIni("jwt_secret", "JWT_SECRET", "thisissecretkey")),
	}
}