package gobindings

import (
)

var MyString string

var MyInteger8		int8
var MyInteger16 	int16
var MyInteger32 	int32
var MyInteger64 	int64

// These types generate error on "gomobile bind ..."
//var MyUInteger8	uint8 //unsupported
//var MyUInteger16 	uint16 //unsupported
//var MyUInteger32 	uint32 //unsupported
//var MyUInteger64 	uint64 //unsupported

var MyFloat32	float32
var MyFloat64	float64

var MyBoolean	bool

var MyError	error

var MyInt	int
var MyRune	rune
var MyByte	byte

// These types are skipped but doesn't generate any error
var MyPointer	*int // skipped variable MyPointer with unsupported type: *types.Pointer
var MySlice	[]int // skipped variable MySlice with unsupported type: *types.Slice
var MyArray	[10]int // skipped variable MyArray with unsupported type: *types.Slice
var MyMap	map[string]int // skipped variable MyMap with unsupported type: *types.Map

// These type is OK
var MyByteSlice		[]byte // skipped variable MySlice with unsupported type: *types.Slice

type MyStruct struct {
	MyInt		int
	privateInt	int
}

func (ms *MyStruct) Inc() {
	ms.MyInt++
}

type MyInterface interface {
	MethodOne()
}

func NormalFunc() {

}

func NormalFuncReturn() string {
	return "some string"
}

func NormalFuncReturnError() (string, error) {
	return "some other string", nil
}




