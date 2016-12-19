package xmli

import (
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"
)

type Items struct {
	XMLName xml.Name `xml:"items"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	XMLName    xml.Name    `xml:"item"`
	ID         uint16      `xml:"id,attr,omitempty"`
	FromID     uint16      `xml:"fromid,attr,omitempty"`
	ToID       uint16      `xml:"toid,attr,omitempty"`
	Article    string      `xml:"article,attr,omitempty"`
	Name       string      `xml:"name,attr,omitempty"`
	Attributes []Attribute `xml:"attribute,omitempty"`
}

type Attribute struct {
	XMLName    xml.Name    `xml:"attribute"`
	Key        string      `xml:"key,attr"`
	Value      string      `xml:"value,attr,omitempty"`
	Attributes []Attribute `xml:"attribute,omitempty"`
}

func Open(path string) (*Items, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	items := new(Items)
	decoder := xml.NewDecoder(fh)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
