package szamlazzhu

import (
	"encoding/xml"
	"time"
)

// Date implements date fields as go time.Time
type Date struct {
	time.Time
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

// UnmarshalXMLAttr UnmarshalXML overrides the default time.Time.UnmarshalXML
func (d *Date) UnmarshalXMLAttr(attr xml.Attr) error {
	// Unmarshal the string as a Date
	return d.UnmarshalText([]byte(attr.Value))
}

func (d *Date) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	dateString := d.Format("2006-01-02")
	attr := xml.Attr{
		Name:  name,
		Value: dateString,
	}

	return attr, nil
}
