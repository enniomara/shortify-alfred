package cachedentries

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type entry struct {
	Name string
}

func GetEntries() []entry {
	jsonFile, err := os.Open("entries.json")
	if err != nil {
		log.Fatalf("Error when opening jsonfile: %s", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var entries []entry

	err = json.Unmarshal(byteValue, &entries)
	return entries
}
