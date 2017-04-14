package system

import (
	"encoding/json"
	"net/url"

	"github.com/ghodss/yaml"
)

// URL represents a url object that can be decoded from JSON and YAML
type URL struct {
	*url.URL
}

// UnmarshalJSON decodes the object from JSON bytes
func (u *URL) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	_u, err := url.Parse(str)
	u.URL = _u
	return err
}

// UnmarshalYAML decodes the object from YAML bytes
func (u *URL) UnmarshalYAML(b []byte) error {
	var str string
	if err := yaml.Unmarshal(b, &str); err != nil {
		return err
	}

	_u, err := url.Parse(str)
	u.URL = _u
	return err
}
