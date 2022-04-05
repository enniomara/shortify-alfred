package main

import (
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/enniomara/shortify-alfred/actions"
	"github.com/enniomara/shortify-alfred/api"
	"github.com/enniomara/shortify-alfred/cachedentries"
)

var wf *aw.Workflow
var endpoint string

func init() {
	updateMagic := aw.AddMagic(actions.NewUpdateAction(func() error {
		apiEntries, err := api.GetEntries(endpoint)
		if err != nil {
			return err
		}

		err = cachedentries.SaveEntries(apiEntries)
		if err != nil {
			return nil
		}
		return nil
	}))

	wf = aw.New(updateMagic)
}

func run() {
	query := wf.Args()[0]

	config := aw.NewConfig()
	endpointUrl := config.Get("shortify_url")
	if endpointUrl == "" {
		wf.Fatal("Endpoint is empty. Make sure to set it.")
	}

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
