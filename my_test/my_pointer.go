package main

import (
	"reflect"
	"unsafe"
)

func main() {

}

func string2Byte(s string) []byte {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&s))

	res := &reflect.SliceHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
		Cap:  sliceHeader.Cap,
	}

	return *(*[]byte)(unsafe.Pointer(&res))
}

func byte2String(b []byte) string {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&b))

	sh := &reflect.StringHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
	}

	// 将这个pointer转成string类型的
	return *(*string)(unsafe.Pointer(&sh))
}
