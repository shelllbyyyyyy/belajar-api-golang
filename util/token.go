package util

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken(id string, email string, exp float64) (tokenString string, err error) {
	privateKeyPath := getKeyPath("private.pem")

	key, e := LoadPrivate(privateKeyPath)
	if e != nil {
		log.Fatal(e)
	}

	claims := jwt.MapClaims{
		"id":   id,
		"email": email,
		"exp": jwt.NewNumericDate(time.Now().Add(time.Duration(exp) * time.Minute)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)

	tokenString, err = token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (id string, email string, err error) {
	publicKeyPath := getKeyPath("public.pem")

	key, e := LoadPublic(publicKeyPath)
	if e != nil {
		log.Fatal(e)
	}

	tokens, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return key, nil
	})

	if err != nil {
		return
	}

	claims, ok := tokens.Claims.(jwt.MapClaims)
	if ok && tokens.Valid {
		id = fmt.Sprintf("%v", claims["id"])
		email = fmt.Sprintf("%v", claims["email"])
		return
	}

	err = fmt.Errorf("unable to extract claims")
	return
}

func getKeyPath(name string) string {
	execDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	return filepath.Join(execDir, "configs", "keys", name)
}

func LoadPrivate(filepath string) (any, error) {
	pemData, e := os.ReadFile(filepath)
	if e != nil {
		return nil, e
	}
  
	pemBlock, _ := pem.Decode(pemData)
	if pemBlock == nil || pemBlock.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}
	priv, e := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
  
	return priv, e
}

func LoadPublic(filepath string) (any, error) {
	pemData, e := os.ReadFile(filepath)
	if e != nil {
		return nil, e
	}

	pemBlock, _ := pem.Decode(pemData)
	if pemBlock == nil || pemBlock.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	priv, e := x509.ParsePKIXPublicKey(pemBlock.Bytes)
  
	return priv, e
}