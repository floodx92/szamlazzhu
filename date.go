package szamlazzhu

import (
	"encoding/xml"
	"time"
)

// Date implements date fields as go time.Time
type Date struct {
	Time time.Time
}

// UnmarshalText unmarshals a date from "2006-01-02" format
func (d *Date) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		d.Time = time.Time{}
		return nil
	}
	t, err := time.Parse("2006-01-02", string(text))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// MarshalText marshals a Date to "2006-01-02" format
func (d *Date) MarshalText() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte{}, nil // Return an empty byte slice to indicate no value.
	}
	s := d.Time.Format("2006-01-02")
	return []byte(s), nil
}

// UnmarshalXML overrides the default time.Time.UnmarshalXML
func (d *Date) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var s string
	err := decoder.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	return d.UnmarshalText([]byte(s))
}

// MarshalXML overrides the default time.Time.MarshalXML
func (d *Date) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	s, err := d.MarshalText()
	if err != nil {
		return err
	}
	if len(s) == 0 {
		return nil // Simulate `omitempty` by not encoding zero dates.
	}
	var v = string(s)
	return encoder.EncodeElement(&v, start)
}
