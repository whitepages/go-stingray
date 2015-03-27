package stingray

import (
	"encoding/json"
	"net/http"
)

// A GlobalSetting is a Stingray global_setting shaping class.
type GlobalSetting struct {
	jsonResource            `json:"-"`
	GlobalSettingProperties `json:"properties"`
}

type GlobalSettingProperties struct {
	Basic struct {
	} `json:"basic"`
}

func (r *GlobalSetting) endpoint() string {
	return "global_settings"
}

func (r *GlobalSetting) String() string {
	s, _ := jsonMarshal(r)
	return string(s)
}

func (r *GlobalSetting) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewGlobalSetting(name string) *GlobalSetting {
	r := new(GlobalSetting)
	r.setName(name)
	return r
}

func (c *Client) GetGlobalSetting(name string) (*GlobalSetting, *http.Response, error) {
	r := NewGlobalSetting(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListGlobalSettings() ([]string, *http.Response, error) {
	return c.List(&GlobalSetting{})
}
