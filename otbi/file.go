package otbi /* import "go-otserv.org/encoding/otbi" */

import (
	"fmt"
	"os"

	"go-otserv.org/encoding/otb"
)

type File struct {
	Root         *otb.Node
	VersionMinor uint32
	VersionMajor uint32
	Build        uint32
	Items        map[uint16]*Item
	ItemsByCID   map[uint16]*Item
}

// Open opens given file for reading
func Open(path string) (*File, error) {
	root, err := otb.Open(path)
	if err != nil {
		return nil, err
	}
	otbifh := &File{
		root, 0, 0, 0, make(map[uint16]*Item, 30000), make(map[uint16]*Item, 30000),
	}

	header, err := root.UInt32()
	if err != nil {
		return otbifh, err
	}
	if header != 0 {
		return otbifh, fmt.Errorf("OTB file error, expected header: four 0 bytes (uint32), %x found", header)
	}
	root.UInt8()  //rootAttr 0x01
	root.UInt16() //dataLen
	if otbifh.VersionMajor, err = root.UInt32(); err != nil {
		return otbifh, err
	}
	if otbifh.VersionMinor, err = root.UInt32(); err != nil {
		return otbifh, err
	}
	if otbifh.Build, err = root.UInt32(); err != nil {
		return otbifh, err
	}

	if otbifh.VersionMajor == 0xFFFFFFFF {
		fmt.Fprintln(os.Stderr, "Warning - items.otb uses generic client version")
	} else if otbifh.VersionMajor != 3 {
		return otbifh, fmt.Errorf("Old version of items.otb (%d.x), newer version required", otbifh.VersionMajor)
	} else if otbifh.VersionMinor < 57 /* 10.98 */ {
		return otbifh, fmt.Errorf("Old version of items.otb (%d.%d), newer version required", otbifh.VersionMajor, otbifh.VersionMinor)
	}
	return otbifh, nil
}

func (otbifh *File) Deserialize() error {
	node := otbifh.Root.Child
	for node != nil {
		item, err := DeserializeItem(node)
		if err != nil {
			return err
		}
		otbifh.Items[item.ServerID] = item
		otbifh.ItemsByCID[item.ClientID] = item
		node = node.Next
	}
	return nil
}
