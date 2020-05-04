// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 04 May 2020 20:49:10 CEST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package leopard

/*
#cgo LDFLAGS: -L./build -llibleopard -lstdc++
#cgo linux LDFLAGS: -fopenmp -ldl
#cgo linux CFLAGS: -fopenmp -DENABLE_OPENMP
#include "leopard.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// leoInit function as declared in leopard/leopard.h:105
func leoInit(version int32) int32 {
	cversion, _ := (C.int)(version), cgoAllocsUnknown
	__ret := C.leo_init_(cversion)
	__v := (int32)(__ret)
	return __v
}

// leoResultString function as declared in leopard/leopard.h:127
func leoResultString(result leopardresult) string {
	cresult, _ := (C.LeopardResult)(result), cgoAllocsUnknown
	__ret := C.leo_result_string(cresult)
	__v := packPCharString(__ret)
	return __v
}

// leoEncodeWorkCount function as declared in leopard/leopard.h:143
func leoEncodeWorkCount(originalCount uint32, recoveryCount uint32) uint32 {
	coriginalCount, _ := (C.uint)(originalCount), cgoAllocsUnknown
	crecoveryCount, _ := (C.uint)(recoveryCount), cgoAllocsUnknown
	__ret := C.leo_encode_work_count(coriginalCount, crecoveryCount)
	__v := (uint32)(__ret)
	return __v
}

// leoEncode function as declared in leopard/leopard.h:180
func leoEncode(bufferBytes uint64, originalCount uint32, recoveryCount uint32, workCount uint32, originalData []unsafe.Pointer, workData []unsafe.Pointer) leopardresult {
	cbufferBytes, _ := (C.uint64_t)(bufferBytes), cgoAllocsUnknown
	coriginalCount, _ := (C.uint)(originalCount), cgoAllocsUnknown
	crecoveryCount, _ := (C.uint)(recoveryCount), cgoAllocsUnknown
	cworkCount, _ := (C.uint)(workCount), cgoAllocsUnknown
	coriginalData, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&originalData)).Data)), cgoAllocsUnknown
	cworkData, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&workData)).Data)), cgoAllocsUnknown
	__ret := C.leo_encode(cbufferBytes, coriginalCount, crecoveryCount, cworkCount, coriginalData, cworkData)
	__v := (leopardresult)(__ret)
	return __v
}

// leoDecodeWorkCount function as declared in leopard/leopard.h:202
func leoDecodeWorkCount(originalCount uint32, recoveryCount uint32) uint32 {
	coriginalCount, _ := (C.uint)(originalCount), cgoAllocsUnknown
	crecoveryCount, _ := (C.uint)(recoveryCount), cgoAllocsUnknown
	__ret := C.leo_decode_work_count(coriginalCount, crecoveryCount)
	__v := (uint32)(__ret)
	return __v
}

// leoDecode function as declared in leopard/leopard.h:227
func leoDecode(bufferBytes uint64, originalCount uint32, recoveryCount uint32, workCount uint32, originalData []unsafe.Pointer, recoveryData []unsafe.Pointer, workData []unsafe.Pointer) leopardresult {
	cbufferBytes, _ := (C.uint64_t)(bufferBytes), cgoAllocsUnknown
	coriginalCount, _ := (C.uint)(originalCount), cgoAllocsUnknown
	crecoveryCount, _ := (C.uint)(recoveryCount), cgoAllocsUnknown
	cworkCount, _ := (C.uint)(workCount), cgoAllocsUnknown
	coriginalData, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&originalData)).Data)), cgoAllocsUnknown
	crecoveryData, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&recoveryData)).Data)), cgoAllocsUnknown
	cworkData, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&workData)).Data)), cgoAllocsUnknown
	__ret := C.leo_decode(cbufferBytes, coriginalCount, crecoveryCount, cworkCount, coriginalData, crecoveryData, cworkData)
	__v := (leopardresult)(__ret)
	return __v
}
