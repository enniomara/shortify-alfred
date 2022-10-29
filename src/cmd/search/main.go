package main

import (
	"flag"
	"fmt"
	"log"

	aw "github.com/deanishe/awgo"

	"github.com/enniomara/shortify-alfred/internal/api"
	"github.com/enniomara/shortify-alfred/internal/cachedentries"
	"github.com/enniomara/shortify-alfred/internal/config"
)

var wf *aw.Workflow
var endpoint string
var (
	setKey        string
	getKey        string
	configMode    bool
	configHandler config.ConfigHandler
)

func init() {
	flag.StringVar(&setKey, "set", "", "")
	flag.StringVar(&getKey, "get", "", "")
	flag.BoolVar(&configMode, "config", false, "")
	flag.Parse()

	wf = aw.New()
	configHandler = config.NewConfigHandler(wf)
}

func run() {
	wf.Args()
	query := flag.Arg(0)

	if configMode {
		// handle case when configuration mode is entered
		configHandler.Handle(setKey, getKey, query)
		return
	}

	endpointUrl := configHandler.GetURL()
	if endpointUrl == "" {
		wf.Fatal("Endpoint is empty. Make sure to set it.")
	}

	entries, err := cachedentries.GetEntries()
	if err != nil {
		log.Printf("Failed to get entries: %s", err)
		wf.Fatal("Failed to get entries from cache")
	}

	if len(entries) == 0 {
		wf.Warn("No entries found, updating. Please try again", "")
		updateEntries(configHandler.GetURL())
	}

	for _, entry := range entries {
		wf.NewItem(entry.Name).Valid(true).Arg(fmt.Sprintf("%s/%s", endpointUrl, entry.Name))
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

// Downloads entries from shortify and saves them to a cache
func updateEntries(endpoint string) error {
	apiEntries, err := api.GetEntries(endpoint)
	if err != nil {
		return err
	}

	err = cachedentries.SaveEntries(apiEntries)
	if err != nil {
		return nil
	}
	return nil
}
