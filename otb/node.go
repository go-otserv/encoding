package otb /* import "go-otserv.org/encoding/otb" */

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	bin "go-otserv.org/encoding/binary"
)

const (
	// END op marks end of Node block
	END = 255
	// START op marks start of Node block
	START = 254
	// ESCAPECHAR marks escape char
	ESCAPECHAR = 253
)

// Node is basic component of otb file
type Node struct {
	buf   *bytes.Buffer
	Type  uint8
	Next  *Node
	Child *Node
}

// NewNode returns new Node instance
func NewNode() *Node {
	return &Node{}
}

func (node *Node) BufLeft() int {
	return node.buf.Len()
}

func parseNode(fh *bin.File, root, node *Node) error {
	current := node
	buf := new(bytes.Buffer)
	propertiesSet := false

	current.Type, _ = fh.UInt8()
	for {
		byt, _ := fh.UInt8()
		switch byt {
		case START:
			// fmt.Println("START - TOP")
			child := NewNode()
			current.buf = buf
			propertiesSet = true
			current.Child = child
			parseNode(fh, root, child)
		case END:
			// fmt.Println("END - TOP")
			if !propertiesSet {
				current.buf = buf
			}
			if node == root {
				return nil
			}
			byt2, _ := fh.UInt8()
			switch byt2 {
			case START:
				// fmt.Println("START - BOTTOM")
				next := NewNode()
				current.Next = next
				current = next
				current.Type, _ = fh.UInt8()
				propertiesSet = false
				buf = new(bytes.Buffer)
				break
			case END:
				// fmt.Println("up and back")
				fh.Seek(int64(-1), os.SEEK_CUR)
				return nil
			default:
				fmt.Printf("(%d)\n", byt2)
				panic("invalid file format")
			}
		case ESCAPECHAR:
			// fmt.Println("ESCAPECHAR - TOP")
			ch, _ := fh.UInt8()
			binary.Write(buf, binary.LittleEndian, ch)
			break
		default:
			// fmt.Println("DEFAULT - TOP")
			binary.Write(buf, binary.LittleEndian, byt)
		}
	}
}
