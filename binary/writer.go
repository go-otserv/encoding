package binary

// Writer interface for binary files
type Writer interface {
	Double(float64) error
	Float(float32) error
	Int8(int8) error
	Int16(int16) error
	Int32(int32) error
	Int64(int64) error
	String(string) error
	UInt8(uint8) error
	UInt16(uint16) error
	UInt32(uint32) error
	UInt64(uint64) error
}
