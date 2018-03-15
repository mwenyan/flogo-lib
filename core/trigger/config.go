package trigger

import (
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"strconv"
	"time"
)

// Config is the configuration for a Trigger
type Config struct {
	Name     string                 `json:"name"`
	Id       string                 `json:"id"`
	Ref      string                 `json:"ref"`
	Settings map[string]interface{} `json:"settings"`
	Output   map[string]interface{} `json:"output"`

	Handlers []*HandlerConfig `json:"handlers"`

	//for backwards compatibility
	Outputs map[string]interface{} `json:"outputs"`
}

func (c *Config) FixUp(metadata *Metadata) {

	//for backwards compatibility
	if len(c.Output) == 0 {
		c.Output = c.Outputs
	}


	// fix up top-level outputs
	for name, value := range c.Output {

		attr, ok := metadata.Output[name]

		if ok {
			newValue, err := data.CoerceToValue(value, attr.Type())

			if err != nil {
				//todo handle error
			} else {
				c.Output[name] = newValue
			}
		}
	}

	// fix up handler outputs
	for _, hc := range c.Handlers {

		hc.parent = c

		//for backwards compatibility
		if hc.ActionId == "" {
			hc.ActionId = strconv.Itoa(time.Now().Nanosecond())
		}

		//for backwards compatibility
		if len(hc.Output) == 0 {
			hc.Output = hc.Outputs
		}

		// fix up outputs
		for name, value := range hc.Output {

			attr, ok := metadata.Output[name]

			if ok {
				newValue, err := data.CoerceToValue(value, attr.Type())

				if err != nil {
					//todo handle error
				} else {
					hc.Output[name] = newValue
				}
			}
		}
	}
}

func (c *Config) GetSetting(setting string) string {

	val, exists := data.GetValueWithResolver(c.Settings, setting)

	if !exists {
		return ""
	}

	strVal, err := data.CoerceToString(val)

	if err != nil {
		return ""
	}

	return strVal
}

type HandlerConfig struct {
	parent   *Config
	Settings map[string]interface{} `json:"settings"`
	Output   map[string]interface{} `json:"output"`
	Action   *action.Config

	//for backwards compatibility
	ActionId             string                 `json:"actionId"`
	ActionMappings       *data.IOMappings       `json:"actionMappings,omitempty"`
	Outputs              map[string]interface{} `json:"outputs"`
	ActionOutputMappings []*data.MappingDef     `json:"actionOutputMappings,omitempty"`
	ActionInputMappings  []*data.MappingDef     `json:"actionInputMappings,omitempty"`
}

func (hc *HandlerConfig) GetSetting(setting string) string {

	val, exists := data.GetValueWithResolver(hc.Settings, setting)

	if !exists {
		return ""
	}

	strVal, err := data.CoerceToString(val)

	if err != nil {
		return ""
	}

	return strVal
}
