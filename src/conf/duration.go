package conf

import (
	"strconv"
	"strings"
	"time"
)

// Duration parse time.Duration from string
type Duration time.Duration

func (d *Duration) UnmarshalJSON(data []byte) (err error) {
	s := string(data)
	if s == "null" {
		return
	}

	s, err = strconv.Unquote(s)
	if err != nil {
		return
	}

	t, err := time.ParseDuration(s)
	if err != nil {
		return
	}

	*d = Duration(t)
	return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	v := time.Duration(d)
	data := append([]byte(`"`), []byte(v.String())...)
	data = append(data, []byte(`"`)...)
	return data, nil
}

// MarshalText ...
func (d Duration) MarshalText() ([]byte, error) {
	s := d.String()
	return []byte(s), nil
}

// String for %v
func (d Duration) String() string {
	v := time.Duration(d)
	return v.String()
}

// GoString for %#v
func (d Duration) GoString() string {
	return strings.Join([]string{"\"", "\""}, d.String())
}
