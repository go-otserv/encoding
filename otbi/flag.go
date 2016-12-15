package otbi /* import "go-otserv.org/encoding/otbi" */

import "encoding/xml"

type FlagOp uint32

const (
	OpBlockSolid         FlagOp = 1 << 0
	OpBlockProjectile           = 1 << 1
	OpBlockPathfind             = 1 << 2
	OpHasHeight                 = 1 << 3
	OpUseable                   = 1 << 4
	OpPickupable                = 1 << 5
	OpMoveable                  = 1 << 6
	OpStackable                 = 1 << 7
	OpFloorChangeDown           = 1 << 8  // unused
	OpFloorCOphangeNorth        = 1 << 9  // unused
	OpFloorChangeEast           = 1 << 10 // unused
	OpFloorChangeSouth          = 1 << 11 // unused
	OpFloorChangeWest           = 1 << 12 // unused
	OpAlwaysOnTop               = 1 << 13
	OpReadable                  = 1 << 14
	OpRotatable                 = 1 << 15
	OpHangable                  = 1 << 16
	OpVertical                  = 1 << 17
	OpHorizontal                = 1 << 18
	OpCannotDecay               = 1 << 19 // unused
	OpAllowDistread             = 1 << 20
	OpUnused                    = 1 << 21 // unused
	OpClientCharges             = 1 << 22 // deprecated
	OpLookThrough               = 1 << 23
	OpAnimation                 = 1 << 24
	OpFullTile                  = 1 << 25 // unused
	OpForceUse                  = 1 << 26
)

var flagsList []FlagOp = []FlagOp{
	OpBlockSolid,
	OpBlockProjectile,
	OpBlockPathfind,
	OpHasHeight,
	OpUseable,
	OpPickupable,
	OpMoveable,
	OpStackable,
	OpAlwaysOnTop,
	OpReadable,
	OpRotatable,
	OpHangable,
	OpVertical,
	OpHorizontal,
	OpAllowDistread,
	OpLookThrough,
	OpAnimation,
	OpForceUse,
}

var flagToNameMap map[FlagOp]string = map[FlagOp]string{
	OpBlockSolid:      "blockSolid",
	OpBlockProjectile: "blockProjectile",
	OpBlockPathfind:   "blockPathfind",
	OpHasHeight:       "hasHeight",
	OpUseable:         "useable",
	OpPickupable:      "pickupable",
	OpMoveable:        "notMoveable",
	OpStackable:       "stackable",
	OpAlwaysOnTop:     "alwaysOnTop",
	OpReadable:        "readable",
	OpRotatable:       "rotatable",
	OpHangable:        "hangable",
	OpVertical:        "vertical",
	OpHorizontal:      "horizontal",
	OpAllowDistread:   "allowDistread",
	OpLookThrough:     "lookThrough",
	OpAnimation:       "animation",
	OpForceUse:        "forceUse",
}

type Flag struct {
	XMLName xml.Name `xml:"flag"`
	Name    string   `xml:"name,attr"`
}

func NewFlag(name string) *Flag {
	return &Flag{Name: name}
}
