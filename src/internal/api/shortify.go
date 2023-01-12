package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Entry struct {
	Name string `json:"name"`
}

func GetEntries(endpoint *url.URL) ([]Entry, error) {
	ref, err := url.Parse("_entries")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	entriesUrl := endpoint.ResolveReference(ref)
	resp, err := http.Get(entriesUrl.String())
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
