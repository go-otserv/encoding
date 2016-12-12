package otb /* import "go-otserv.org/encoding/otb" */

import (
	"fmt"

	bin "go-otserv.org/encoding/binary"
)

// Open opens given file for reading
func Open(path string) (*Node, error) {
	fh, err := bin.OpenFile(path)
	if err != nil {
		return nil, err
	}
	overallVersion, err := fh.UInt32() //overallVersion
	if err != nil {
		return nil, err
	}
	if overallVersion != 0 {
		return nil, fmt.Errorf("OTB file error, expected header: four 0 bytes (uint32), %x found", overallVersion)

	}

	nodeOp, err := fh.UInt8()
	if err != nil {
		return nil, err
	}
	if nodeOp != START {
		return nil, fmt.Errorf("OTB file error, expected node.START byte, %x found", nodeOp)
	}

	root := NewNode()
	parseNode(fh, root, root)
	return root, nil
}
