package cachedentries

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"github.com/enniomara/shortify-alfred/api"
)

type entry struct {
	Name string `json:"name"`
}

func GetEntries() ([]entry, error) {
	jsonFile, err := os.Open("entries.json")
	if err != nil {
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

func SaveEntries(apiEntries []api.Entry) error {
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
