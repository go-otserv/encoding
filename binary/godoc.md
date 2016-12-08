# Package `binary`
## Overview
Package binary provides abstraction for working with binary files.

## Index

* Types
  * [BufferedFile](#BufferedFile)
	 * [func OpenBufferedFile(path string) (*BufferedFile, error)](#OpenBufferedFile)
	 * [func (fh *BufferedFile) Discard(n int) (int, error)](#BufferedFile-Discard)
	 * [func (fh *BufferedFile) Double() (float64, error)](#BufferedFile-Double)
	 * [func (fh *BufferedFile) Float() (float32, error)](#BufferedFile-Float)
	 * [func (fh *BufferedFile) Int16() (int16, error)](#BufferedFile-Int16)
	 * [func (fh *BufferedFile) Int32() (int32, error)](#BufferedFile-Int32)
	 * [func (fh *BufferedFile) Int64() (int64, error)](#BufferedFile-Int64)
	 * [func (fh *BufferedFile) Int8() (int8, error)](#BufferedFile-Int8)
	 * [func (fh *BufferedFile) String() (string, error)](#BufferedFile-String)
	 * [func (fh *BufferedFile) UInt16() (uint16, error)](#BufferedFile-UInt16)
	 * [func (fh *BufferedFile) UInt32() (uint32, error)](#BufferedFile-UInt32)
	 * [func (fh *BufferedFile) UInt64() (uint64, error)](#BufferedFile-UInt64)
	 * [func (fh *BufferedFile) UInt8() (uint8, error)](#BufferedFile-UInt8)

  * [File](#File)
	 * [func OpenFile(path string) (*File, error)](#OpenFile)
	 * [func (fh *File) Discard(n int) (int, error)](#File-Discard)
	 * [func (fh *File) Double() (float64, error)](#File-Double)
	 * [func (fh *File) Float() (float32, error)](#File-Float)
	 * [func (fh *File) Int16() (int16, error)](#File-Int16)
	 * [func (fh *File) Int32() (int32, error)](#File-Int32)
	 * [func (fh *File) Int64() (int64, error)](#File-Int64)
	 * [func (fh *File) Int8() (int8, error)](#File-Int8)
	 * [func (fh *File) String() (string, error)](#File-String)
	 * [func (fh *File) UInt16() (uint16, error)](#File-UInt16)
	 * [func (fh *File) UInt32() (uint32, error)](#File-UInt32)
	 * [func (fh *File) UInt64() (uint64, error)](#File-UInt64)
	 * [func (fh *File) UInt8() (uint8, error)](#File-UInt8)

  * [Reader](#Reader)

  * [Writer](#Writer)

## Types

### type <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L10" name="BufferedFile">BufferedFile</a> [¶](#BufferedFile)
```go
type BufferedFile struct {
	bufio.Reader
}
```
BufferedFile extends bufio.Reader implementing Reader interface  

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L16" name="OpenBufferedFile">OpenBufferedFile</a> [¶](#OpenBufferedFile)
```go
func OpenBufferedFile(path string) (*BufferedFile, error)
```
OpenBufferedFile opens file at given path for reading. If there is an
error, it will be of type *PathError

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L36" name="BufferedFile-Discard">Discard</a> [¶](#BufferedFile-Discard)
```go
func (fh *BufferedFile) Discard(n int) (int, error)
```
Discard n bytes

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L27" name="BufferedFile-Double">Double</a> [¶](#BufferedFile-Double)
```go
func (fh *BufferedFile) Double() (float64, error)
```
Double reads float64 from BufferedFile

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L42" name="BufferedFile-Float">Float</a> [¶](#BufferedFile-Float)
```go
func (fh *BufferedFile) Float() (float32, error)
```
Float reads float32 from BufferedFile

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L60" name="BufferedFile-Int16">Int16</a> [¶](#BufferedFile-Int16)
```go
func (fh *BufferedFile) Int16() (int16, error)
```
Int16 reads int16 from BufferedFile

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L69" name="BufferedFile-Int32">Int32</a> [¶](#BufferedFile-Int32)
```go
func (fh *BufferedFile) Int32() (int32, error)
```
Int32 reads int32 from BufferedFile

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L78" name="BufferedFile-Int64">Int64</a> [¶](#BufferedFile-Int64)
```go
func (fh *BufferedFile) Int64() (int64, error)
```
Int64 reads int64 from BufferedFile

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L51" name="BufferedFile-Int8">Int8</a> [¶](#BufferedFile-Int8)
```go
func (fh *BufferedFile) Int8() (int8, error)
```
Int8 reads int8 from BufferedFile

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L86" name="BufferedFile-String">String</a> [¶](#BufferedFile-String)
```go
func (fh *BufferedFile) String() (string, error)
```

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L108" name="BufferedFile-UInt16">UInt16</a> [¶](#BufferedFile-UInt16)
```go
func (fh *BufferedFile) UInt16() (uint16, error)
```
UInt16 reads uint16 from BufferedFile

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L117" name="BufferedFile-UInt32">UInt32</a> [¶](#BufferedFile-UInt32)
```go
func (fh *BufferedFile) UInt32() (uint32, error)
```
UInt32 reads uint32 from BufferedFile

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L126" name="BufferedFile-UInt64">UInt64</a> [¶](#BufferedFile-UInt64)
```go
func (fh *BufferedFile) UInt64() (uint64, error)
```
UInt64 reads uint64 from BufferedFile

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/buffered_file.go#L99" name="BufferedFile-UInt8">UInt8</a> [¶](#BufferedFile-UInt8)
```go
func (fh *BufferedFile) UInt8() (uint8, error)
```
UInt8 reads uint8 from BufferedFile

### type <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L9" name="File">File</a> [¶](#File)
```go
type File struct {
	os.File
}
```
File is os.File extension implementing Reader interface  

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L15" name="OpenFile">OpenFile</a> [¶](#OpenFile)
```go
func OpenFile(path string) (*File, error)
```
OpenFile opens file at given path for reading. If there is an error, it
will be of type *PathError

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L33" name="File-Discard">Discard</a> [¶](#File-Discard)
```go
func (fh *File) Discard(n int) (int, error)
```
Discard n bytes

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L24" name="File-Double">Double</a> [¶](#File-Double)
```go
func (fh *File) Double() (float64, error)
```
Double reads float64 from File

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L39" name="File-Float">Float</a> [¶](#File-Float)
```go
func (fh *File) Float() (float32, error)
```
Float reads float32 from File

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L57" name="File-Int16">Int16</a> [¶](#File-Int16)
```go
func (fh *File) Int16() (int16, error)
```
Int16 reads int16 from File

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L66" name="File-Int32">Int32</a> [¶](#File-Int32)
```go
func (fh *File) Int32() (int32, error)
```
Int32 reads int32 from File

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L75" name="File-Int64">Int64</a> [¶](#File-Int64)
```go
func (fh *File) Int64() (int64, error)
```
Int64 reads int64 from File

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L48" name="File-Int8">Int8</a> [¶](#File-Int8)
```go
func (fh *File) Int8() (int8, error)
```
Int8 reads int8 from File

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L86" name="File-String">String</a> [¶](#File-String)
```go
func (fh *File) String() (string, error)
```
String reads string from File
Function first reads uint16 which denotes string length N, then N bytes are
read and returned as string

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L108" name="File-UInt16">UInt16</a> [¶](#File-UInt16)
```go
func (fh *File) UInt16() (uint16, error)
```
UInt16 reads uint16 from File

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L117" name="File-UInt32">UInt32</a> [¶](#File-UInt32)
```go
func (fh *File) UInt32() (uint32, error)
```
UInt32 reads uint32 from File

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L126" name="File-UInt64">UInt64</a> [¶](#File-UInt64)
```go
func (fh *File) UInt64() (uint64, error)
```
UInt64 reads uint64 from File

#### func <a href="https://github.com/go-otserv/encoding/blob/master/binary/file.go#L99" name="File-UInt8">UInt8</a> [¶](#File-UInt8)
```go
func (fh *File) UInt8() (uint8, error)
```
UInt8 reads uint8 from File

### type <a href="https://github.com/go-otserv/encoding/blob/master/binary/reader.go#L5" name="Reader">Reader</a> [¶](#Reader)
```go
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
```
Reader interface for binary files  

### type <a href="https://github.com/go-otserv/encoding/blob/master/binary/writer.go#L4" name="Writer">Writer</a> [¶](#Writer)
```go
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
```
Writer interface for binary files  

***
_Last updated 8 Dec 2016_
