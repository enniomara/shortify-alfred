package main

import (
	aw "github.com/deanishe/awgo"
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

	wf.WarnEmpty("No matching items", "Try something different")
	wf.SendFeedback()
}
