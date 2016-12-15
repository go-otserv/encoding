package otbi /* import "go-otserv.org/encoding/otbi" */

import "encoding/xml"

type AttrOp uint16

const (
	OpFirst        AttrOp = 0x10
	OpServerID            = OpFirst
	OpClientID            = 0x11
	OpName                = 0x12
	OpDescr               = 0x13
	OpSpeed               = 0x14
	OpSlot                = 0x15
	OpMaxItems            = 0x16
	OpWeigth              = 0x17
	OpWeapon              = 0x18
	OpAmu                 = 0x19
	OpArmor               = 0x1A
	OpMagLevel            = 0x1B
	OpMagFieldType        = 0x1C
	OpWriteable           = 0x1D
	OpRotateTo            = 0x1E
	OpDecay               = 0x1F
	OpSpriteHash          = 0x20
	OpMinimapColor        = 0x21
	OpAttr07              = 0x22
	OpAttr08              = 0x23
	OpLight               = 0x24

	OpDecay2     = 0x25 //deprecated
	OpWeapon2    = 0x26 //deprecated
	OpAmu2       = 0x27 //deprecated
	OpArmor2     = 0x28 //deprecated
	OpWriteable2 = 0x29 //deprecated
	OpLight2     = 0x2A
	OpTopOrder   = 0x2B
	OpWriteable3 = 0x2C //deprecated
	OpWareID     = 0x2D

	OpLast = 0x2E
)

type Attribute interface {
	attrName() string
}

type AttributeBase struct {
	XMLName xml.Name `xml:"attr"`
	Name    string   `xml:"name,attr"`
}

type Description struct {
	AttributeBase
	Text string `xml:"text,attr"`
}

// attrName implements Attribute interface
func (attr Description) attrName() string { return attr.Name }

// NewDescription creates new Description attribute
func NewDescription(val string) *Description {
	return &Description{AttributeBase{Name: "description"}, val}
}

type Speed struct {
	AttributeBase
	TextLen uint16 `xml:"val,attr"`
}

// attrName implements Attribute interface
func (attr Speed) attrName() string { return attr.Name }

// NewSpeed creates new Speed attribute
func NewSpeed(val uint16) *Speed {
	return &Speed{AttributeBase{Name: "speed"}, val}
}

type Light2 struct {
	AttributeBase
	Level uint16 `xml:"level,attr"`
	Color uint16 `xml:"color,attr"`
}

// attrName implements Attribute interface
func (attr Light2) attrName() string { return attr.Name }

// NewLight2 creates new Light2 attribute
func NewLight2(level, color uint16) *Light2 {
	return &Light2{AttributeBase{Name: "light2"}, level, color}
}

type TopOrder struct {
	AttributeBase
	Val uint8 `xml:"val,attr"`
}

// attrName implements Attribute interface
func (attr TopOrder) attrName() string { return attr.Name }

// NewTopOrder creates new TopOrder attribute
func NewTopOrder(val uint8) *TopOrder {
	return &TopOrder{AttributeBase{Name: "topOrder"}, val}
}

type WareID struct {
	AttributeBase
	Val uint16 `xml:"val,attr"`
}

// attrName implements Attribute interface
func (attr WareID) attrName() string { return attr.Name }

// NewWareID creates new WareID attribute
func NewWareID(val uint16) *WareID {
	return &WareID{AttributeBase{Name: "wareID"}, val}
}
