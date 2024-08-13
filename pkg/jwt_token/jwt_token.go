package jwttoken

import (
	"crypto/rsa"
	"log/slog"
	"time"
	"token/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type ConfigJWTToken struct {
	ServiceAccount *config.ServiceAccount
	Logger         *slog.Logger
	Url            string
}

func Generate(cfg *ConfigJWTToken) (string, error) {
	privateKey, err := parseRSAPrivateFromPEM([]byte(cfg.ServiceAccount.PrivateKey))
	if err != nil {
		cfg.Logger.Error("Error parsing private key", slog.Any("error", err))
		return "", err
	}

	claims := jwt.RegisteredClaims{
		Issuer:    cfg.ServiceAccount.ServiceAccountId,
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(1 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		Audience:  []string{cfg.Url},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)
	token.Header["kid"] = cfg.ServiceAccount.Id

	signed, err := token.SignedString(privateKey)
	if err != nil {
		cfg.Logger.Error("Error signing token", slog.Any("error", err))
		return "", err
	}
	return signed, nil
}

func parseRSAPrivateFromPEM(key []byte) (*rsa.PrivateKey, error) {
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return nil, err
	}
	return rsaPrivateKey, nil
}
