package config

import (
	"fmt"
	"net/url"

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
func (config *ConfigHandler) Handle(setKey, itemToEdit, query string) {
	// we want to save configuration for a key (done after editing)
	if setKey != "" {
		config.Set(setKey, query)
		return
	}

	// we want to edit configuration for a key
	if itemToEdit != "" {
		if query != "" {
			config.wf.
				NewItem(fmt.Sprintf("Set %s to \"%s\"", itemToEdit, query)).
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
	if key == "url" {
		_, err := url.ParseRequestURI(value)
		if err != nil {
			config.wf.Fatal("URL was invalid. Please enter a valid url.")
			return
		}

		if err := config.wf.Config.Set(key, value, false).Do(); err != nil {
			config.wf.FatalError(err)
			return
		}
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

func (config ConfigHandler) GetURL() (*url.URL, error) {
	u, err := url.Parse(config.config.Url);
	if err != nil {
		return nil, err
	}

	return u, nil
}
