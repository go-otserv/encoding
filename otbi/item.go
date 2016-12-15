package otbi /* import "go-otserv.org/encoding/otbi" */

import (
	"encoding/xml"
	"fmt"
	"os"

	"go-otserv.org/encoding/otb"
)

type Item struct {
	XMLName    xml.Name    `xml:"item"`
	ServerID   uint16      `xml:"id,attr"`
	ClientID   uint16      `xml:"cid,attr"`
	Type       string      `xml:"type,attr"`
	Name       string      `xml:"name,attr"`
	Attributes []Attribute `xml:"Attribute,omitempty"`
	Flags      []*Flag     `xml:"Flag,omitempty"`
}

func DeserializeItem(node *otb.Node) (*Item, error) {
	item := &Item{}
	if err := item.deserializeFlags(node); err != nil {
		return nil, err
	}
	if err := item.deserializeAttributes(node); err != nil {
		return nil, err
	}
	return item, nil
}

func (i *Item) deserializeFlags(node *otb.Node) error {
	flags, err := node.UInt32()
	if err != nil {
		return err
	}
	i.Flags = make([]*Flag, 5)
	for _, flagOp := range flagsList {
		if (FlagOp(flags) & flagOp) == flagOp {
			flagName := flagToNameMap[flagOp]
			if flagName != "" {
				i.Flags = append(i.Flags, &Flag{Name: flagName})
			}
		}
	}
	return nil
}

func (i *Item) deserializeAttributes(node *otb.Node) error {
	i.Attributes = make([]Attribute, 7)
	for node.Len() > 0 {
		attr, err := node.UInt8()
		if err != nil {
			return err
		}
		dataLen, err := node.UInt16()
		if err != nil {
			return err
		}
		switch AttrOp(attr) {
		case OpServerID:
			serverID, _ := node.UInt16()
			if err != nil {
				return err
			}
			if serverID > 30000 && serverID < 30100 {
				serverID -= 30000
			}
			i.ServerID = serverID
		case OpClientID:
			clientID, err := node.UInt16()
			if err != nil {
				return err
			}
			i.ClientID = clientID
		case OpSpeed:
			speed, err := node.UInt16()
			if err != nil {
				return err
			}
			i.Attributes = append(i.Attributes, NewSpeed(speed))
		case OpLight2:
			lightLevel, err := node.UInt16()
			if err != nil {
				return err
			}
			lightColor, err := node.UInt16()
			if err != nil {
				return err
			}
			i.Attributes = append(i.Attributes, NewLight2(lightLevel, lightColor))
		case OpTopOrder:
			topOrder, err := node.UInt8()
			if err != nil {
				return err
			}
			i.Attributes = append(i.Attributes, NewTopOrder(topOrder))
		case OpWareID:
			wareID, err := node.UInt16()
			if err != nil {
				return err
			}
			i.Attributes = append(i.Attributes, NewWareID(wareID))
		default:
			fmt.Fprintf(os.Stderr, "UNKNOWN ATTR: 0x%x\n", attr)
			var i uint16
			for i = 0; i < dataLen; i++ {
				_, _ = node.UInt8()
			}
		}
	}
	return nil
}
