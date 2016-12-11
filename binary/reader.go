// Package binary provides abstraction for working with binary files.
package binary /* import "go-otserv.org/encoding/binary" */

// Reader interface for binary files
type Reader interface {
	Double() (float64, error)
	Discard(int) (int, error)
	Float() (float32, error)
	Int8() (int8, error)
	Int16() (int16, error)
	Int32() (int32, error)
	Int64() (int64, error)
	String() (string, error)
	UInt8() (uint8, error)
	UInt16() (uint16, error)
	UInt32() (uint32, error)
	UInt64() (uint64, error)
}
