package main

import (
	aw "github.com/deanishe/awgo"
	"github.com/enniomara/shortify-alfred/cachedentries"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	entries := cachedentries.GetEntries()
	for _, entry := range entries {
		wf.NewItem(entry.Name)
	}
	wf.WarnEmpty("No matching items", "Try something different")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
