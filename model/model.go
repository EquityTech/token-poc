package api

import "net/url"

// Token represents an ERC-20 compliant set of fields
type Token struct {
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	Decimals      int    `json:"decimals"`
	InitialSupply int    `json:"initialSupply"`
}

// Validate checks for existence of required fieldss
func (t *Token) Validate() url.Values {
	errs := url.Values{}

	if t.Name == "" {
		errs.Add("name", "Field is required")
	}
	if t.Symbol == "" {
		errs.Add("symbol", "Field is required")
	}
	if t.Decimals == 0 {
		errs.Add("decimals", "Field is required")
	}
	if t.InitialSupply == 0 {
		errs.Add("initialSupply", "Field is required")
	}

	return errs
}
