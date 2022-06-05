package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Entry struct {
	Name string `json:"name"`
}

func GetEntries(endpoint string) ([]Entry, error) {
	resp, err := http.Get(endpoint + "/_entries")
	if err != nil {
		log.Print(err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var entries []Entry
	err = json.Unmarshal(body, &entries)
	if err != nil {
		log.Print("Could not unmarshal")
		return nil, err
	}
	return entries, nil
}
