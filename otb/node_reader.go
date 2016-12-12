package otb /* import "go-otserv.org/encoding/otb" */

import (
	"encoding/binary"
	"fmt"
)

// Double reads float64 from Node
func (no *Node) Double() (float64, error) {
	var out float64
	if err := binary.Read(no.buf, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Discard n bytes
func (no *Node) Discard(n int) (int, error) {
	// TODO: fixme, Discarding / bath reading was bugged
	return 0, fmt.Errorf("Not implemented")
}

// Float reads float32 from Node
func (no *Node) Float() (float32, error) {
	var out float32
	if err := binary.Read(no.buf, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int8 reads int8 from Node
func (no *Node) Int8() (int8, error) {
	var out int8
	if err := binary.Read(no.buf, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int16 reads int16 from Node
func (no *Node) Int16() (int16, error) {
	var out int16
	if err := binary.Read(no.buf, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int32 reads int32 from Node
func (no *Node) Int32() (int32, error) {
	var out int32
	if err := binary.Read(no.buf, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int64 reads int64 from Node
func (no *Node) Int64() (int64, error) {
	var out int64
	if err := binary.Read(no.buf, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

func (no *Node) String() (string, error) {
	var strLen uint16
	if err := binary.Read(no.buf, binary.LittleEndian, &strLen); err != nil {
		return "", err
	}
	out := make([]byte, strLen, strLen)
	if _, err := no.buf.Read(out); err != nil {
		return "", err
	}
	return string(out), nil
}

// UInt8 reads uint8 from Node
func (no *Node) UInt8() (uint8, error) {
	var out uint8
	if err := binary.Read(no.buf, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt16 reads uint16 from Node
func (no *Node) UInt16() (uint16, error) {
	var out uint16
	if err := binary.Read(no.buf, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt32 reads uint32 from Node
func (no *Node) UInt32() (uint32, error) {
	var out uint32
	if err := binary.Read(no.buf, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt64 reads uint64 from Node
func (no *Node) UInt64() (uint64, error) {
	var out uint64
	if err := binary.Read(no.buf, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}
