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

// DecodeString - decodes given base58 string
func DecodeString(src string) (string, error) {
	if len(src) == 0 {
		return "", nil
	}

	outSize := C.size_t(decodeLen(len(src)) + 1)
	dstData := make([]byte, outSize)

	if !C.decode_base58(
		(*C.char)(unsafe.Pointer(unsafe.StringData(src))),
		C.size_t(len(src)),
		unsafe.Pointer(&dstData[0]),
		&outSize,
	) {
		errno := C.get_errno()

		return "", fmt.Errorf("%s (errno:%d)", C.GoString(C.strerror(errno)), errno)
	}

	return unsafe.String(&dstData[len(dstData)-int(outSize)], outSize), nil
}

// DecodeBytes - decodes given base58 bytes to dst, returns decoded count or error
func DecodeBytes(src, dst []byte) (int, error) {
	if len(src) == 0 || len(dst) == 0 {
		return 0, nil
	}

	var outSize C.size_t
	if decLen := decodeLen(len(src)); len(dst) < decLen {
		return 0, fmt.Errorf("given dst length %d is too small to decode", len(dst))
	} else {
		outSize = C.size_t(decLen + 1) // add +1 for preventing allocation err
	}

	if !C.decode_base58(
		(*C.char)(unsafe.Pointer(&src[0])),
		C.size_t(len(src)),
		unsafe.Pointer(&dst[0]),
		&outSize,
	) {
		errno := C.get_errno()

		return 0, fmt.Errorf("%s (errno:%d)", C.GoString(C.strerror(errno)), errno)
	}

	// our bytes are placed at the end of slice, shift it
	var shift = 0
	for i := 0; i < len(dst); i++ {
		if dst[i] == 0 {
			continue
		}

		dst[shift], dst[i] = dst[i], 0
		shift++
		if shift == int(outSize) {
			break
		}
	}

	return int(outSize), nil
}

func decodeLen(l int) int {
	return l * 4 / 3
}
