package main

import (
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/enniomara/shortify-alfred/cachedentries"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	query := wf.Args()[0]
	entries, err := cachedentries.GetEntries()
	if err != nil {
		log.Printf("Failed to get entries: %s", err)
		wf.Fatal("Failed to get entries from cache")
	}

	for _, entry := range entries {
		wf.NewItem(entry.Name)
	}

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try something different")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
