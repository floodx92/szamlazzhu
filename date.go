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
	// Decode the XML element content as a string
	var s string
	err := decoder.DecodeElement(&s, &start)
	if err != nil {
		return err
	}

	// Unmarshal the string as a Date
	return d.UnmarshalText([]byte(s))
}

// MarshalXML overrides the default time.Time.MarshalXML
func (d *Date) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	// Format the date as a string in the "2006-01-02" format
	if d.Time.IsZero() {
		// If the date is zero, do not output anything.
		return nil // Optionally, handle empty date representation as per your requirements
	}
	formattedDate := d.Time.Format("2006-01-02")

	// Create a simple struct that represents the XML element content
	type dateElement struct {
		XMLName xml.Name
		Date    string `xml:",chardata"`
	}

	// Use the dateElement struct to marshal the date as XML element content
	return encoder.EncodeElement(dateElement{XMLName: start.Name, Date: formattedDate}, start)
}
