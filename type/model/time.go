package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Time represents a nullable epoch time in a millisecond precision
type Time struct {
	Time  time.Time
	Valid bool
}

// NewTime creates a new time
func NewTime(t time.Time, valid bool) Time {
	return Time{
		Time:  t,
		Valid: valid,
	}
}

// TimeFrom creates a new time that will always be valid
func TimeFrom(t time.Time) Time {
	return NewTime(t, true)
}

// TimeFromPtr creates a new time that will be null if t is nil
func TimeFromPtr(t *time.Time) Time {
	if t == nil {
		return NewTime(time.Time{}, false)
	}
	return NewTime(*t, true)
}

// Scan implements the Scanner interface.
// See https://golang.org/pkg/database/sql/#Scanner
func (t *Time) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case time.Time:
		t.Time = x
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("null: cannot scan type %T into null.Time: %v", value, value)
	}
	t.Valid = err == nil
	return err
}

// Value implements the Valuer interface.
// See https://golang.org/pkg/database/sql/driver/#Valuer
func (t Time) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

func (t Time) toEpochMsec() int64 {
	return t.Time.UnixNano() / 1e6
}

func (Time) timeFromEpochMsec(msec int64) time.Time {
	return time.Unix(msec/1e3, (msec%1e3)*1e3)
}

// MarshalJSON implements the Marshaler interface.
// See https://golang.org/pkg/encoding/json/#Marshaler
func (t Time) MarshalJSON() ([]byte, error) {
	return t.MarshalText()
}

// UnmarshalJSON implements the Unmarshaler interface.
// See https://golang.org/pkg/encoding/json/#UnmarshalJSON
func (t *Time) UnmarshalJSON(data []byte) error {
	return t.UnmarshalText(data)
}

// MarshalText implements the MarshalText interface.
// See https://golang.org/pkg/encoding/#MarshalText
func (t Time) MarshalText() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(t.toEpochMsec(), 10)), nil
}

// UnmarshalText implements the UnmarshalText interface.
// See https://golang.org/pkg/encoding/#UnmarshalText
func (t *Time) UnmarshalText(data []byte) error {
	str := string(data)
	if str == "" || str == "null" {
		t.Valid = false
		return nil
	}

	var epochMsec int64
	if err := json.Unmarshal(data, &epochMsec); err != nil {
		t.Valid = false
		return nil
	}

	t.Time = t.timeFromEpochMsec(epochMsec)

	return nil
}

// SetValid changes the time's value and sets it to be non-null
func (t *Time) SetValid(v time.Time) {
	t.Time = v
	t.Valid = true
}

// Ptr returns a pointer to the time's value, or a nil pointer if the time is null
func (t Time) Ptr() *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}
