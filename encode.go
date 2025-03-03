package base58

/*
#cgo CFLAGS: -Wall -fPIC
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include "libbase58.h"

*/
import "C"
import (
	"fmt"
	"unsafe"
)

// EncodeString - encodes given string
func EncodeString(src string) (string, error) {
	if len(src) == 0 {
		return "", nil
	}

	outSize := C.size_t(encodeLen(len(src)) + 1)
	dstData := make([]byte, outSize)

	if !C.encode_base58(
		unsafe.Pointer(unsafe.StringData(src)),
		C.size_t(len(src)),
		(*C.char)(unsafe.Pointer(&dstData[0])),
		&outSize,
	) {
		errno := C.get_errno()

		return "", fmt.Errorf("%s (errno:%d)", C.GoString(C.strerror(errno)), errno)
	}

	return unsafe.String(&dstData[0], outSize), nil
}

// EncodeBytes - encode given bytes to dst slice. Return a number of encoded bytes or error
func EncodeBytes(src, dst []byte) (int, error) {
	if len(src) == 0 {
		return 0, nil
	}

	var outSize C.size_t
	if encLen := encodeLen(len(src)); len(dst) < encLen {
		return 0, fmt.Errorf("given dst length %d is too small to encode", len(dst))
	} else {
		outSize = C.size_t(encLen + 1) // add +1 for preventing allocation err
	}

	if !C.encode_base58(
		unsafe.Pointer(&src[0]),
		C.size_t(len(src)),
		(*C.char)(unsafe.Pointer(&dst[0])),
		&outSize,
	) {
		errno := C.get_errno()

		return 0, fmt.Errorf("%s (errno:%d)", C.GoString(C.strerror(errno)), errno)
	}

	return int(outSize), nil
}

// encodeLen - returns a encoded bytes count by given count
func encodeLen(l int) int {
	return l * 3 / 2
}
