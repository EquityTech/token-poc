package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	model "github.com/ssb4/token-poc/model"
	"github.com/ssb4/token-poc/service"
)

// Controller accepts input from the Router and performs corresponding endpoint actions
type Controller struct {
	TokenService service.TokenService
}

// CreateToken - POST /tokens
func (c *Controller) CreateToken(w http.ResponseWriter, r *http.Request) {
	var token model.Token
	body, err := readBody(r.Body, "CreateToken")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &token); err != nil {
		if err2 := json.NewEncoder(w).Encode(err); err2 != nil {
			log.Println("[CreateToken] Error unmarshalling: ", err2)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Println("[CreateToken] Error unmarshalling data: ", err)
		return
	}
	if validationErrors := token.Validate(); len(validationErrors) > 0 {
		w.Header().Set("Content-type", "applciation/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErrors)
		return
	}

	err = c.TokenService.CreateToken(token)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

// readBody tries to read the request body and returns it as a byte slice
func readBody(requestBody io.ReadCloser, funcName string) ([]byte, error) {
	body, err := ioutil.ReadAll(io.LimitReader(requestBody, 1048576)) // ~ 1mb

	if err != nil {
		log.Println(funcName+" Error reading request body: ", err)
		return nil, err
	}
	if err := requestBody.Close(); err != nil {
		// Don't fail in this case but log it.
		log.Println(funcName+" Error closing body: ", err)
	}

	return body, nil
}
