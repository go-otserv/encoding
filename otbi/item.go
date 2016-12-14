package otbi /* import "go-otserv.org/encoding/otbi" */

type ItemAttr uint16

const (
	First        ItemAttr = 0x10
	ServerID              = First
	ClientID              = 0x11
	Name                  = 0x12
	Descr                 = 0x13
	Speed                 = 0x14
	Slot                  = 0x15
	MaxItems              = 0x16
	Weigth                = 0x17
	Weapon                = 0x18
	Amu                   = 0x19
	Armor                 = 0x1A
	MagLevel              = 0x1B
	MagFieldType          = 0x1C
	Writeable             = 0x1D
	RotateTo              = 0x1E
	Decay                 = 0x1F
	SpriteHash            = 0x20
	MinimapColor          = 0x21
	Attr07                = 0x22
	Attr08                = 0x23
	Light                 = 0x24

	Decay2     = 0x25 //deprecated
	Weapon2    = 0x26 //deprecated
	Amu2       = 0x27 //deprecated
	Armor2     = 0x28 //deprecated
	Writeable2 = 0x29 //deprecated
	Light2     = 0x2A
	TopOrder   = 0x2B
	Writeable3 = 0x2C //deprecated
	WareID     = 0x2D

	Last = 0x2E
)
