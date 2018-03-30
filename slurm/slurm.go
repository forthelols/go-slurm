// WARNING: This file has automatically been generated on Fri, 30 Mar 2018 11:30:42 CEST.
// By https://git.io/c-for-go. DO NOT EDIT.

package slurm

/*
#cgo pkg-config: slurm
#include <slurm/slurm.h>
#include "slurm_helpers.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// FreeSubmitResponseResponseMsg function as declared in slurm/slurm.h:3364
func FreeSubmitResponseResponseMsg(msg *SubmitResponseMsg) {
	cmsg, _ := msg.PassRef()
	C.slurm_free_submit_response_response_msg(cmsg)
}

// SetTrigger function as declared in slurm/slurm.h:4993
func SetTrigger(triggerSet *TriggerInfo) int32 {
	ctriggerSet, _ := triggerSet.PassRef()
	__ret := C.slurm_set_trigger(ctriggerSet)
	__v := (int32)(__ret)
	return __v
}

// ClearTrigger function as declared in slurm/slurm.h:4999
func ClearTrigger(triggerClear *TriggerInfo) int32 {
	ctriggerClear, _ := triggerClear.PassRef()
	__ret := C.slurm_clear_trigger(ctriggerClear)
	__v := (int32)(__ret)
	return __v
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

// SubmitBatchJob function as declared in slurm/slurm_helpers.h:12
func SubmitBatchJob(jobDescMsg *JobDescMsg) *SubmitResponseMsg {
	cjobDescMsg, _ := jobDescMsg.PassRef()
	__ret := C.wrap_slurm_submit_batch_job(cjobDescMsg)
	__v := NewSubmitResponseMsgRef(unsafe.Pointer(__ret))
	return __v
}
