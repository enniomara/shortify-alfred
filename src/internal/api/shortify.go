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

func EndpointFromUrl(url *url.URL) endpoint {
	return endpoint{url}
}

type endpoint struct {
	*url.URL
}

func (e endpoint) GetEntries() ([]Entry, error) {
	ref, _ := url.Parse("_entries")

	entriesUrl := e.ResolveReference(ref)
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

func (e endpoint) UrlOf(entryName string) *url.URL {
	ref, _ := url.Parse(entryName)

	return e.ResolveReference(ref)
}
