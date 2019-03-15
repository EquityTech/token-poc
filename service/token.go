package service

import (
	"html/template"
	"os"

	api "github.com/ssb4/token-poc/model"
)

// TokenService handles creating Solidity-based tokens from templates
type TokenService struct{}

// CreateToken generates a solidity contract based on a token request model
func (ts *TokenService) CreateToken(token api.Token) error {
	filename := token.Name + ".sol"
	t, _ := template.ParseFiles("templates/ERC20.tmpl")

	f, _ := os.Create(filename)
	defer f.Close()

	t.Execute(f, token)

	return nil
}
