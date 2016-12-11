package binary /* import "go-otserv.org/encoding/binary" */

import (
	"encoding/binary"
	"os"
)

// File is os.File extension implementing Reader interface
type File struct {
	os.File
}

// OpenFile opens file at given path for reading. If there is an error, it
// will be of type *PathError
func OpenFile(path string) (*File, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &File{*fh}, nil
}

// Double reads float64 from File
func (fh *File) Double() (float64, error) {
	var out float64
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Discard n bytes
func (fh *File) Discard(n int) (int, error) {
	m, err := fh.Seek(int64(n), 1)
	return int(m), err
}

// Float reads float32 from File
func (fh *File) Float() (float32, error) {
	var out float32
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int8 reads int8 from File
func (fh *File) Int8() (int8, error) {
	var out int8
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int16 reads int16 from File
func (fh *File) Int16() (int16, error) {
	var out int16
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int32 reads int32 from File
func (fh *File) Int32() (int32, error) {
	var out int32
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// Int64 reads int64 from File
func (fh *File) Int64() (int64, error) {
	var out int64
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// String reads string from File
// Function first reads uint16 which denotes string length N, then N bytes are
// read and returned as string
func (fh *File) String() (string, error) {
	var strLen uint16
	if err := binary.Read(fh, binary.LittleEndian, &strLen); err != nil {
		return "", err
	}
	out := make([]byte, strLen, strLen)
	if _, err := fh.Read(out); err != nil {
		return "", err
	}
	return string(out), nil
}

// UInt8 reads uint8 from File
func (fh *File) UInt8() (uint8, error) {
	var out uint8
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt16 reads uint16 from File
func (fh *File) UInt16() (uint16, error) {
	var out uint16
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt32 reads uint32 from File
func (fh *File) UInt32() (uint32, error) {
	var out uint32
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}

// UInt64 reads uint64 from File
func (fh *File) UInt64() (uint64, error) {
	var out uint64
	if err := binary.Read(fh, binary.LittleEndian, &out); err != nil {
		return 0, err
	}
	return out, nil
}
