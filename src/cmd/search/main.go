package main

import (
	"flag"
	"log"
	"net/url"

	aw "github.com/deanishe/awgo"

	"github.com/enniomara/shortify-alfred/internal/api"
	"github.com/enniomara/shortify-alfred/internal/cachedentries"
	"github.com/enniomara/shortify-alfred/internal/config"
)

var wf *aw.Workflow
var endpoint string
var (
	setKey           string
	configItemToEdit string
	configMode       bool
	configHandler    config.ConfigHandler
)

func init() {
	flag.StringVar(&setKey, "set", "", "")
	flag.StringVar(&configItemToEdit, "edit", "", "")
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
		configHandler.Handle(setKey, configItemToEdit, query)
		return
	}

	endpointUrl, err := configHandler.GetURL()
	if err != nil {
		wf.Fatal("Configure API endpoint by running `sh-settings`.")
	}

	cachedEntries := cachedentries.NewAlfredCacheStorage(*wf.Cache)
	entries, err := cachedEntries.GetEntries()
	if err != nil {
		log.Printf("Failed to get entries: %s", err)
		wf.Fatal("Failed to get entries from cache")
	}

	if len(entries) == 0 {
		err := updateEntries(endpointUrl, cachedEntries)
		if err != nil {
			log.Printf("Error while updating entries: %s", err)
			wf.Fatal("Failed to update entries")
			wf.SendFeedback()
		}

		// the entries have been updated, we need to update them
		entries, err = cachedEntries.GetEntries()
		if err != nil {
			log.Printf("Failed to get entries: %s", err)
			wf.Fatal("Failed to get entries from cache")
		}
	}

	for _, entry := range entries {
		ref, err := url.Parse(entry.Name)
		if err != nil {
			log.Printf("Failed to parse entry as URL: %s", err)
			wf.Fatal("Unknown error while parsing entry")
		}
		wf.NewItem(entry.Name).Valid(true).Arg(endpointUrl.ResolveReference(ref).String())
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
func updateEntries(endpoint *url.URL, cachedEntries cachedentries.CachedEntries) error {
	apiEntries, err := api.GetEntries(endpoint)
	if err != nil {
		return err
	}

	err = cachedEntries.SaveEntries(apiEntries)
	if err != nil {
		return nil
	}
	return nil
}
