package main

import (
	aw "github.com/deanishe/awgo"
	"github.com/enniomara/shortify-alfred/cachedentries"
)

var (
	wf           *aw.Workflow
	query        string
	shouldUpdate bool
)

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(func() {})

	entries := cachedentries.GetEntries()
	for _, entry := range entries {
		wf.NewItem(entry.Name)
	}
	wf.WarnEmpty("No matching items", "Try something different")
	wf.SendFeedback()
}
