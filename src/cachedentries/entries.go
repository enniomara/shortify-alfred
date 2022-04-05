package cachedentries

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type entry struct {
	Name string
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
