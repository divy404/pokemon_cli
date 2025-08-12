package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Pokemon struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Types []struct {
		Type struct{
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`

}
func FetchPokemon(nameOrID string) (*Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", strings.ToLower(nameOrID))
	resp , err := http.Get(url) 
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("pokemon not found or API error: %s (status %d)",nameOrID, resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var p Pokemon
	if err := json.Unmarshal(body,&p); err != nil {
		return nil,err
	}
	return &p, nil
}
