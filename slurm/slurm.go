// WARNING: This file has automatically been generated on Wed, 28 Mar 2018 14:32:22 CEST.
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
