package cachedentries

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/enniomara/shortify-alfred/internal/api"
)

type entry struct {
	Name string `json:"name"`
}

type CachedEntries interface {
	GetEntries() ([]entry, error)
	SaveEntries([]api.Entry) error
}

func NewFsStorage() CachedEntries {
	return fsStorage{}
}

type fsStorage struct {}

func (s fsStorage) GetEntries() ([]entry, error) {
	jsonFile, err := os.Open("entries.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []entry{}, nil
		}

		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var entries []entry

	err = json.Unmarshal(byteValue, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (s fsStorage) SaveEntries(apiEntries []api.Entry) error {
	entries := fromApiEntry(apiEntries)
	marshaledOutput, err := json.Marshal(entries)
	if err != nil {
		return err
	}

	f, err := os.Create("entries.json")
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write(marshaledOutput)
	if err != nil {
		return err
	}

	return nil
}

func fromApiEntry(apiEntries []api.Entry) []entry {
	var entries []entry
	for _, e := range apiEntries {
		entries = append(entries, entry{
			Name: e.Name,
		})
	}

	return entries
}
