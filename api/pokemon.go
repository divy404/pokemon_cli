package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type PokemonListResponse struct {
	Count int `json:"count"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

type TypeResponse struct {
	Pokemon []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon"`
}

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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var p Pokemon
	if err := json.Unmarshal(body,&p); err != nil {
		return nil,err
	}
	return &p, nil
}

func FetchPoekmonList(limit, offset int) (*PokemonListResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?limit=%d&offset=%d",limit,offset)
	resp , err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var list PokemonListResponse
	if err := json.Unmarshal(body, &list); err != nil {
		return nil,err
	}
	return &list,nil
}

func FetchPokemonByType(typeName string) (*TypeResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/type/%s",strings.ToLower(typeName))
	resp,err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("type not found: %s", typeName)
	}
	body, err := io.ReadAll(resp.Body)
	if err!= nil {
		return nil,err
	}
	var tr TypeResponse
	if err := json.Unmarshal(body,&tr); err != nil {
		return nil,err
	}
	return &tr, nil
}
