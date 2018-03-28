// WARNING: This file has automatically been generated on Wed, 28 Mar 2018 14:55:54 CEST.
// By https://git.io/c-for-go. DO NOT EDIT.

package slurm

/*
#cgo pkg-config: slurm
#include <slurm/slurm.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// SubmitBatchJob function as declared in slurm/slurm.h:3344
func SubmitBatchJob(jobDescMsg *JobDescMsg, slurmAllocMsg **SubmitResponseMsg) int32 {
	cjobDescMsg, _ := jobDescMsg.PassRef()
	cslurmAllocMsg, _ := (**C.submit_response_msg_t)(unsafe.Pointer(slurmAllocMsg)), cgoAllocsUnknown
	__ret := C.slurm_submit_batch_job(cjobDescMsg, cslurmAllocMsg)
	__v := (int32)(__ret)
	return __v
}

// FreeSubmitResponseResponseMsg function as declared in slurm/slurm.h:3364
func FreeSubmitResponseResponseMsg(msg *SubmitResponseMsg) {
	cmsg, _ := (*C.submit_response_msg_t)(unsafe.Pointer(msg)), cgoAllocsUnknown
	C.slurm_free_submit_response_response_msg(cmsg)
}

// Strerror function as declared in slurm/slurm_errno.h:290
func Strerror(errnum int32) *byte {
	cerrnum, _ := (C.int)(errnum), cgoAllocsUnknown
	__ret := C.slurm_strerror(cerrnum)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// GetErrno function as declared in slurm/slurm_errno.h:296
func GetErrno() int32 {
	__ret := C.slurm_get_errno()
	__v := (int32)(__ret)
	return __v
}
