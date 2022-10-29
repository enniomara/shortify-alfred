package config

import (
	"fmt"
	"log"

	aw "github.com/deanishe/awgo"
)

type config struct {
	Url   string `env:"url"`
	Token string
}
type ConfigHandler struct {
	wf     *aw.Workflow
	config *config
}

func NewConfigHandler(wf *aw.Workflow) ConfigHandler {
	defaultConfig := &config{}
	if err := wf.Config.To(defaultConfig); err != nil {
		panic("Could not load current Alfred configuration.")
	}

	c := ConfigHandler{
		wf:     wf,
		config: defaultConfig,
	}
	return c
}

// Handle handles the configuration GUI, i.e. setting and getting of config
// variables
func (config *ConfigHandler) Handle(setKey, getKey, query string) {
	// we want to save configuration for a key (done after editing)
	if setKey != "" {
		config.Set(setKey, query)
		return
	}

	// we want to edit configuration for a key
	if getKey != "" {
		if query != "" {
			config.wf.
				NewItem(fmt.Sprintf("Set %s to \"%s\"", getKey, query)).
				Valid(true).
				Var("value", query).
				Arg(query).
				Subtitle("↩ to save")
		}
		config.wf.SendFeedback()
		return
	}

	config.ShowItems()
	config.wf.SendFeedback()
	return
}

// Set the given configuration key to a value
func (config *ConfigHandler) Set(key string, value string) {
	log.Printf("Modifying %s to %s", key, value)
	if err := config.wf.Config.Set(key, value, false).Do(); err != nil {
		config.wf.FatalError(err)
	}
}

// Shows the possible configuration items
func (config *ConfigHandler) ShowItems() {
	config.wf.NewItem(fmt.Sprintf("Url: %s", config.config.Url)).
		Var("name", "url").
		// when the action is clicked, the input field will contain the old
		// URL, so that editing is easy
		Arg(config.config.Url).
		Valid(true).
		Autocomplete("url").
		Subtitle("↩ to edit")
}

func (config ConfigHandler) GetURL() string {
	return config.config.Url
}
