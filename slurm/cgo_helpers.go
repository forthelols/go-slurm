// WARNING: This file has automatically been generated on Fri, 30 Mar 2018 11:18:58 CEST.
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
import (
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocJobDescMsgMemory allocates memory for type C.job_desc_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJobDescMsgMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJobDescMsgValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJobDescMsgValue = unsafe.Sizeof([1]C.job_desc_msg_t{})

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// allocPCharMemory allocates memory for type *C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPCharValue = unsafe.Sizeof([1]*C.char{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackSSByte transforms a sliced Go data structure into plain C format.
func unpackSSByte(x [][]byte) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.char)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}

// packSSByte reads sliced Go data structure out from plain C format.
func packSSByte(v [][]byte, ptr0 **C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.char)(unsafe.Pointer(ptr0)))[i0]
		hxfa9955c := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxfa9955c.Data = uintptr(unsafe.Pointer(ptr1))
		hxfa9955c.Cap = 0x7fffffff
		// hxfa9955c.Len = ?
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *JobDescMsg) Ref() *C.job_desc_msg_t {
	if x == nil {
		return nil
	}
	return x.ref1da77736
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *JobDescMsg) Free() {
	if x != nil && x.allocs1da77736 != nil {
		x.allocs1da77736.(*cgoAllocMap).Free()
		x.ref1da77736 = nil
	}
}

// NewJobDescMsgRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewJobDescMsgRef(ref unsafe.Pointer) *JobDescMsg {
	if ref == nil {
		return nil
	}
	obj := new(JobDescMsg)
	obj.ref1da77736 = (*C.job_desc_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *JobDescMsg) PassRef() (*C.job_desc_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref1da77736 != nil {
		return x.ref1da77736, nil
	}
	mem1da77736 := allocJobDescMsgMemory(1)
	ref1da77736 := (*C.job_desc_msg_t)(mem1da77736)
	allocs1da77736 := new(cgoAllocMap)
	allocs1da77736.Add(mem1da77736)

	var caccount_allocs *cgoAllocMap
	ref1da77736.account, caccount_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Account)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(caccount_allocs)

	var cacctg_freq_allocs *cgoAllocMap
	ref1da77736.acctg_freq, cacctg_freq_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.AcctgFreq)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cacctg_freq_allocs)

	var cadmin_comment_allocs *cgoAllocMap
	ref1da77736.admin_comment, cadmin_comment_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.AdminComment)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cadmin_comment_allocs)

	var calloc_node_allocs *cgoAllocMap
	ref1da77736.alloc_node, calloc_node_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.AllocNode)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(calloc_node_allocs)

	var calloc_resp_port_allocs *cgoAllocMap
	ref1da77736.alloc_resp_port, calloc_resp_port_allocs = (C.uint16_t)(x.AllocRespPort), cgoAllocsUnknown
	allocs1da77736.Borrow(calloc_resp_port_allocs)

	var calloc_sid_allocs *cgoAllocMap
	ref1da77736.alloc_sid, calloc_sid_allocs = (C.uint32_t)(x.AllocSid), cgoAllocsUnknown
	allocs1da77736.Borrow(calloc_sid_allocs)

	var cargc_allocs *cgoAllocMap
	ref1da77736.argc, cargc_allocs = (C.uint32_t)(x.Argc), cgoAllocsUnknown
	allocs1da77736.Borrow(cargc_allocs)

	var cargv_allocs *cgoAllocMap
	ref1da77736.argv, cargv_allocs = unpackSSByte(x.Argv)
	allocs1da77736.Borrow(cargv_allocs)

	var carray_inx_allocs *cgoAllocMap
	ref1da77736.array_inx, carray_inx_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ArrayInx)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(carray_inx_allocs)

	var carray_bitmap_allocs *cgoAllocMap
	ref1da77736.array_bitmap, carray_bitmap_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.ArrayBitmap)), cgoAllocsUnknown
	allocs1da77736.Borrow(carray_bitmap_allocs)

	var cbegin_time_allocs *cgoAllocMap
	ref1da77736.begin_time, cbegin_time_allocs = (C.time_t)(x.BeginTime), cgoAllocsUnknown
	allocs1da77736.Borrow(cbegin_time_allocs)

	var cbitflags_allocs *cgoAllocMap
	ref1da77736.bitflags, cbitflags_allocs = (C.uint32_t)(x.Bitflags), cgoAllocsUnknown
	allocs1da77736.Borrow(cbitflags_allocs)

	var cburst_buffer_allocs *cgoAllocMap
	ref1da77736.burst_buffer, cburst_buffer_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.BurstBuffer)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cburst_buffer_allocs)

	var cckpt_interval_allocs *cgoAllocMap
	ref1da77736.ckpt_interval, cckpt_interval_allocs = (C.uint16_t)(x.CkptInterval), cgoAllocsUnknown
	allocs1da77736.Borrow(cckpt_interval_allocs)

	var cckpt_dir_allocs *cgoAllocMap
	ref1da77736.ckpt_dir, cckpt_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.CkptDir)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cckpt_dir_allocs)

	var cclusters_allocs *cgoAllocMap
	ref1da77736.clusters, cclusters_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Clusters)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cclusters_allocs)

	var ccluster_features_allocs *cgoAllocMap
	ref1da77736.cluster_features, ccluster_features_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ClusterFeatures)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(ccluster_features_allocs)

	var ccomment_allocs *cgoAllocMap
	ref1da77736.comment, ccomment_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Comment)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(ccomment_allocs)

	var ccontiguous_allocs *cgoAllocMap
	ref1da77736.contiguous, ccontiguous_allocs = (C.uint16_t)(x.Contiguous), cgoAllocsUnknown
	allocs1da77736.Borrow(ccontiguous_allocs)

	var ccore_spec_allocs *cgoAllocMap
	ref1da77736.core_spec, ccore_spec_allocs = (C.uint16_t)(x.CoreSpec), cgoAllocsUnknown
	allocs1da77736.Borrow(ccore_spec_allocs)

	var ccpu_bind_allocs *cgoAllocMap
	ref1da77736.cpu_bind, ccpu_bind_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.CpuBind)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpu_bind_allocs)

	var ccpu_bind_type_allocs *cgoAllocMap
	ref1da77736.cpu_bind_type, ccpu_bind_type_allocs = (C.uint16_t)(x.CpuBindType), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpu_bind_type_allocs)

	var ccpu_freq_min_allocs *cgoAllocMap
	ref1da77736.cpu_freq_min, ccpu_freq_min_allocs = (C.uint32_t)(x.CpuFreqMin), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpu_freq_min_allocs)

	var ccpu_freq_max_allocs *cgoAllocMap
	ref1da77736.cpu_freq_max, ccpu_freq_max_allocs = (C.uint32_t)(x.CpuFreqMax), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpu_freq_max_allocs)

	var ccpu_freq_gov_allocs *cgoAllocMap
	ref1da77736.cpu_freq_gov, ccpu_freq_gov_allocs = (C.uint32_t)(x.CpuFreqGov), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpu_freq_gov_allocs)

	var cdeadline_allocs *cgoAllocMap
	ref1da77736.deadline, cdeadline_allocs = (C.time_t)(x.Deadline), cgoAllocsUnknown
	allocs1da77736.Borrow(cdeadline_allocs)

	var cdelay_boot_allocs *cgoAllocMap
	ref1da77736.delay_boot, cdelay_boot_allocs = (C.uint32_t)(x.DelayBoot), cgoAllocsUnknown
	allocs1da77736.Borrow(cdelay_boot_allocs)

	var cdependency_allocs *cgoAllocMap
	ref1da77736.dependency, cdependency_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Dependency)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cdependency_allocs)

	var cend_time_allocs *cgoAllocMap
	ref1da77736.end_time, cend_time_allocs = (C.time_t)(x.EndTime), cgoAllocsUnknown
	allocs1da77736.Borrow(cend_time_allocs)

	var cenvironment_allocs *cgoAllocMap
	ref1da77736.environment, cenvironment_allocs = unpackSSByte(x.Environment)
	allocs1da77736.Borrow(cenvironment_allocs)

	var cenv_size_allocs *cgoAllocMap
	ref1da77736.env_size, cenv_size_allocs = (C.uint32_t)(x.EnvSize), cgoAllocsUnknown
	allocs1da77736.Borrow(cenv_size_allocs)

	var cextra_allocs *cgoAllocMap
	ref1da77736.extra, cextra_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Extra)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cextra_allocs)

	var cexc_nodes_allocs *cgoAllocMap
	ref1da77736.exc_nodes, cexc_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ExcNodes)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cexc_nodes_allocs)

	var cfeatures_allocs *cgoAllocMap
	ref1da77736.features, cfeatures_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Features)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cfeatures_allocs)

	var cfed_siblings_active_allocs *cgoAllocMap
	ref1da77736.fed_siblings_active, cfed_siblings_active_allocs = (C.uint64_t)(x.FedSiblingsActive), cgoAllocsUnknown
	allocs1da77736.Borrow(cfed_siblings_active_allocs)

	var cfed_siblings_viable_allocs *cgoAllocMap
	ref1da77736.fed_siblings_viable, cfed_siblings_viable_allocs = (C.uint64_t)(x.FedSiblingsViable), cgoAllocsUnknown
	allocs1da77736.Borrow(cfed_siblings_viable_allocs)

	var cgres_allocs *cgoAllocMap
	ref1da77736.gres, cgres_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Gres)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cgres_allocs)

	var cgroup_id_allocs *cgoAllocMap
	ref1da77736.group_id, cgroup_id_allocs = (C.uint32_t)(x.GroupId), cgoAllocsUnknown
	allocs1da77736.Borrow(cgroup_id_allocs)

	var cimmediate_allocs *cgoAllocMap
	ref1da77736.immediate, cimmediate_allocs = (C.uint16_t)(x.Immediate), cgoAllocsUnknown
	allocs1da77736.Borrow(cimmediate_allocs)

	var cjob_id_allocs *cgoAllocMap
	ref1da77736.job_id, cjob_id_allocs = (C.uint32_t)(x.JobId), cgoAllocsUnknown
	allocs1da77736.Borrow(cjob_id_allocs)

	var cjob_id_str_allocs *cgoAllocMap
	ref1da77736.job_id_str, cjob_id_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.JobIdStr)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cjob_id_str_allocs)

	var ckill_on_node_fail_allocs *cgoAllocMap
	ref1da77736.kill_on_node_fail, ckill_on_node_fail_allocs = (C.uint16_t)(x.KillOnNodeFail), cgoAllocsUnknown
	allocs1da77736.Borrow(ckill_on_node_fail_allocs)

	var clicenses_allocs *cgoAllocMap
	ref1da77736.licenses, clicenses_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Licenses)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(clicenses_allocs)

	var cmail_type_allocs *cgoAllocMap
	ref1da77736.mail_type, cmail_type_allocs = (C.uint16_t)(x.MailType), cgoAllocsUnknown
	allocs1da77736.Borrow(cmail_type_allocs)

	var cmail_user_allocs *cgoAllocMap
	ref1da77736.mail_user, cmail_user_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.MailUser)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cmail_user_allocs)

	var cmcs_label_allocs *cgoAllocMap
	ref1da77736.mcs_label, cmcs_label_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.McsLabel)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cmcs_label_allocs)

	var cmem_bind_allocs *cgoAllocMap
	ref1da77736.mem_bind, cmem_bind_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.MemBind)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cmem_bind_allocs)

	var cmem_bind_type_allocs *cgoAllocMap
	ref1da77736.mem_bind_type, cmem_bind_type_allocs = (C.uint16_t)(x.MemBindType), cgoAllocsUnknown
	allocs1da77736.Borrow(cmem_bind_type_allocs)

	var cname_allocs *cgoAllocMap
	ref1da77736.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Name)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cname_allocs)

	var cnetwork_allocs *cgoAllocMap
	ref1da77736.network, cnetwork_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Network)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cnetwork_allocs)

	var cnice_allocs *cgoAllocMap
	ref1da77736.nice, cnice_allocs = (C.uint32_t)(x.Nice), cgoAllocsUnknown
	allocs1da77736.Borrow(cnice_allocs)

	var cnum_tasks_allocs *cgoAllocMap
	ref1da77736.num_tasks, cnum_tasks_allocs = (C.uint32_t)(x.NumTasks), cgoAllocsUnknown
	allocs1da77736.Borrow(cnum_tasks_allocs)

	var copen_mode_allocs *cgoAllocMap
	ref1da77736.open_mode, copen_mode_allocs = (C.uint8_t)(x.OpenMode), cgoAllocsUnknown
	allocs1da77736.Borrow(copen_mode_allocs)

	var corigin_cluster_allocs *cgoAllocMap
	ref1da77736.origin_cluster, corigin_cluster_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.OriginCluster)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(corigin_cluster_allocs)

	var cother_port_allocs *cgoAllocMap
	ref1da77736.other_port, cother_port_allocs = (C.uint16_t)(x.OtherPort), cgoAllocsUnknown
	allocs1da77736.Borrow(cother_port_allocs)

	var covercommit_allocs *cgoAllocMap
	ref1da77736.overcommit, covercommit_allocs = (C.uint8_t)(x.Overcommit), cgoAllocsUnknown
	allocs1da77736.Borrow(covercommit_allocs)

	var cpack_job_offset_allocs *cgoAllocMap
	ref1da77736.pack_job_offset, cpack_job_offset_allocs = (C.uint32_t)(x.PackJobOffset), cgoAllocsUnknown
	allocs1da77736.Borrow(cpack_job_offset_allocs)

	var cpartition_allocs *cgoAllocMap
	ref1da77736.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Partition)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cpartition_allocs)

	var cplane_size_allocs *cgoAllocMap
	ref1da77736.plane_size, cplane_size_allocs = (C.uint16_t)(x.PlaneSize), cgoAllocsUnknown
	allocs1da77736.Borrow(cplane_size_allocs)

	var cpower_flags_allocs *cgoAllocMap
	ref1da77736.power_flags, cpower_flags_allocs = (C.uint8_t)(x.PowerFlags), cgoAllocsUnknown
	allocs1da77736.Borrow(cpower_flags_allocs)

	var cpriority_allocs *cgoAllocMap
	ref1da77736.priority, cpriority_allocs = (C.uint32_t)(x.Priority), cgoAllocsUnknown
	allocs1da77736.Borrow(cpriority_allocs)

	var cprofile_allocs *cgoAllocMap
	ref1da77736.profile, cprofile_allocs = (C.uint32_t)(x.Profile), cgoAllocsUnknown
	allocs1da77736.Borrow(cprofile_allocs)

	var cqos_allocs *cgoAllocMap
	ref1da77736.qos, cqos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Qos)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cqos_allocs)

	var creboot_allocs *cgoAllocMap
	ref1da77736.reboot, creboot_allocs = (C.uint16_t)(x.Reboot), cgoAllocsUnknown
	allocs1da77736.Borrow(creboot_allocs)

	var cresp_host_allocs *cgoAllocMap
	ref1da77736.resp_host, cresp_host_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.RespHost)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cresp_host_allocs)

	var crestart_cnt_allocs *cgoAllocMap
	ref1da77736.restart_cnt, crestart_cnt_allocs = (C.uint16_t)(x.RestartCnt), cgoAllocsUnknown
	allocs1da77736.Borrow(crestart_cnt_allocs)

	var creq_nodes_allocs *cgoAllocMap
	ref1da77736.req_nodes, creq_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ReqNodes)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(creq_nodes_allocs)

	var crequeue_allocs *cgoAllocMap
	ref1da77736.requeue, crequeue_allocs = (C.uint16_t)(x.Requeue), cgoAllocsUnknown
	allocs1da77736.Borrow(crequeue_allocs)

	var creservation_allocs *cgoAllocMap
	ref1da77736.reservation, creservation_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Reservation)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(creservation_allocs)

	var cscript_allocs *cgoAllocMap
	ref1da77736.script, cscript_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Script)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cscript_allocs)

	var cshared_allocs *cgoAllocMap
	ref1da77736.shared, cshared_allocs = (C.uint16_t)(x.Shared), cgoAllocsUnknown
	allocs1da77736.Borrow(cshared_allocs)

	var cspank_job_env_allocs *cgoAllocMap
	ref1da77736.spank_job_env, cspank_job_env_allocs = unpackSSByte(x.SpankJobEnv)
	allocs1da77736.Borrow(cspank_job_env_allocs)

	var cspank_job_env_size_allocs *cgoAllocMap
	ref1da77736.spank_job_env_size, cspank_job_env_size_allocs = (C.uint32_t)(x.SpankJobEnvSize), cgoAllocsUnknown
	allocs1da77736.Borrow(cspank_job_env_size_allocs)

	var ctask_dist_allocs *cgoAllocMap
	ref1da77736.task_dist, ctask_dist_allocs = (C.uint32_t)(x.TaskDist), cgoAllocsUnknown
	allocs1da77736.Borrow(ctask_dist_allocs)

	var ctime_limit_allocs *cgoAllocMap
	ref1da77736.time_limit, ctime_limit_allocs = (C.uint32_t)(x.TimeLimit), cgoAllocsUnknown
	allocs1da77736.Borrow(ctime_limit_allocs)

	var ctime_min_allocs *cgoAllocMap
	ref1da77736.time_min, ctime_min_allocs = (C.uint32_t)(x.TimeMin), cgoAllocsUnknown
	allocs1da77736.Borrow(ctime_min_allocs)

	var cuser_id_allocs *cgoAllocMap
	ref1da77736.user_id, cuser_id_allocs = (C.uint32_t)(x.UserId), cgoAllocsUnknown
	allocs1da77736.Borrow(cuser_id_allocs)

	var cwait_all_nodes_allocs *cgoAllocMap
	ref1da77736.wait_all_nodes, cwait_all_nodes_allocs = (C.uint16_t)(x.WaitAllNodes), cgoAllocsUnknown
	allocs1da77736.Borrow(cwait_all_nodes_allocs)

	var cwarn_flags_allocs *cgoAllocMap
	ref1da77736.warn_flags, cwarn_flags_allocs = (C.uint16_t)(x.WarnFlags), cgoAllocsUnknown
	allocs1da77736.Borrow(cwarn_flags_allocs)

	var cwarn_signal_allocs *cgoAllocMap
	ref1da77736.warn_signal, cwarn_signal_allocs = (C.uint16_t)(x.WarnSignal), cgoAllocsUnknown
	allocs1da77736.Borrow(cwarn_signal_allocs)

	var cwarn_time_allocs *cgoAllocMap
	ref1da77736.warn_time, cwarn_time_allocs = (C.uint16_t)(x.WarnTime), cgoAllocsUnknown
	allocs1da77736.Borrow(cwarn_time_allocs)

	var cwork_dir_allocs *cgoAllocMap
	ref1da77736.work_dir, cwork_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.WorkDir)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cwork_dir_allocs)

	var ccpus_per_task_allocs *cgoAllocMap
	ref1da77736.cpus_per_task, ccpus_per_task_allocs = (C.uint16_t)(x.CpusPerTask), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpus_per_task_allocs)

	var cmin_cpus_allocs *cgoAllocMap
	ref1da77736.min_cpus, cmin_cpus_allocs = (C.uint32_t)(x.MinCpus), cgoAllocsUnknown
	allocs1da77736.Borrow(cmin_cpus_allocs)

	var cmax_cpus_allocs *cgoAllocMap
	ref1da77736.max_cpus, cmax_cpus_allocs = (C.uint32_t)(x.MaxCpus), cgoAllocsUnknown
	allocs1da77736.Borrow(cmax_cpus_allocs)

	var cmin_nodes_allocs *cgoAllocMap
	ref1da77736.min_nodes, cmin_nodes_allocs = (C.uint32_t)(x.MinNodes), cgoAllocsUnknown
	allocs1da77736.Borrow(cmin_nodes_allocs)

	var cmax_nodes_allocs *cgoAllocMap
	ref1da77736.max_nodes, cmax_nodes_allocs = (C.uint32_t)(x.MaxNodes), cgoAllocsUnknown
	allocs1da77736.Borrow(cmax_nodes_allocs)

	var cboards_per_node_allocs *cgoAllocMap
	ref1da77736.boards_per_node, cboards_per_node_allocs = (C.uint16_t)(x.BoardsPerNode), cgoAllocsUnknown
	allocs1da77736.Borrow(cboards_per_node_allocs)

	var csockets_per_board_allocs *cgoAllocMap
	ref1da77736.sockets_per_board, csockets_per_board_allocs = (C.uint16_t)(x.SocketsPerBoard), cgoAllocsUnknown
	allocs1da77736.Borrow(csockets_per_board_allocs)

	var csockets_per_node_allocs *cgoAllocMap
	ref1da77736.sockets_per_node, csockets_per_node_allocs = (C.uint16_t)(x.SocketsPerNode), cgoAllocsUnknown
	allocs1da77736.Borrow(csockets_per_node_allocs)

	var ccores_per_socket_allocs *cgoAllocMap
	ref1da77736.cores_per_socket, ccores_per_socket_allocs = (C.uint16_t)(x.CoresPerSocket), cgoAllocsUnknown
	allocs1da77736.Borrow(ccores_per_socket_allocs)

	var cthreads_per_core_allocs *cgoAllocMap
	ref1da77736.threads_per_core, cthreads_per_core_allocs = (C.uint16_t)(x.ThreadsPerCore), cgoAllocsUnknown
	allocs1da77736.Borrow(cthreads_per_core_allocs)

	var cntasks_per_node_allocs *cgoAllocMap
	ref1da77736.ntasks_per_node, cntasks_per_node_allocs = (C.uint16_t)(x.NtasksPerNode), cgoAllocsUnknown
	allocs1da77736.Borrow(cntasks_per_node_allocs)

	var cntasks_per_socket_allocs *cgoAllocMap
	ref1da77736.ntasks_per_socket, cntasks_per_socket_allocs = (C.uint16_t)(x.NtasksPerSocket), cgoAllocsUnknown
	allocs1da77736.Borrow(cntasks_per_socket_allocs)

	var cntasks_per_core_allocs *cgoAllocMap
	ref1da77736.ntasks_per_core, cntasks_per_core_allocs = (C.uint16_t)(x.NtasksPerCore), cgoAllocsUnknown
	allocs1da77736.Borrow(cntasks_per_core_allocs)

	var cntasks_per_board_allocs *cgoAllocMap
	ref1da77736.ntasks_per_board, cntasks_per_board_allocs = (C.uint16_t)(x.NtasksPerBoard), cgoAllocsUnknown
	allocs1da77736.Borrow(cntasks_per_board_allocs)

	var cpn_min_cpus_allocs *cgoAllocMap
	ref1da77736.pn_min_cpus, cpn_min_cpus_allocs = (C.uint16_t)(x.PnMinCpus), cgoAllocsUnknown
	allocs1da77736.Borrow(cpn_min_cpus_allocs)

	var cpn_min_memory_allocs *cgoAllocMap
	ref1da77736.pn_min_memory, cpn_min_memory_allocs = (C.uint64_t)(x.PnMinMemory), cgoAllocsUnknown
	allocs1da77736.Borrow(cpn_min_memory_allocs)

	var cpn_min_tmp_disk_allocs *cgoAllocMap
	ref1da77736.pn_min_tmp_disk, cpn_min_tmp_disk_allocs = (C.uint32_t)(x.PnMinTmpDisk), cgoAllocsUnknown
	allocs1da77736.Borrow(cpn_min_tmp_disk_allocs)

	var cgeometry_allocs *cgoAllocMap
	ref1da77736.geometry, cgeometry_allocs = *(*[5]C.uint16_t)(unsafe.Pointer(&x.Geometry)), cgoAllocsUnknown
	allocs1da77736.Borrow(cgeometry_allocs)

	var cconn_type_allocs *cgoAllocMap
	ref1da77736.conn_type, cconn_type_allocs = *(*[5]C.uint16_t)(unsafe.Pointer(&x.ConnType)), cgoAllocsUnknown
	allocs1da77736.Borrow(cconn_type_allocs)

	var crotate_allocs *cgoAllocMap
	ref1da77736.rotate, crotate_allocs = (C.uint16_t)(x.Rotate), cgoAllocsUnknown
	allocs1da77736.Borrow(crotate_allocs)

	var cblrtsimage_allocs *cgoAllocMap
	ref1da77736.blrtsimage, cblrtsimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Blrtsimage)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cblrtsimage_allocs)

	var clinuximage_allocs *cgoAllocMap
	ref1da77736.linuximage, clinuximage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Linuximage)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(clinuximage_allocs)

	var cmloaderimage_allocs *cgoAllocMap
	ref1da77736.mloaderimage, cmloaderimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Mloaderimage)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cmloaderimage_allocs)

	var cramdiskimage_allocs *cgoAllocMap
	ref1da77736.ramdiskimage, cramdiskimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Ramdiskimage)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cramdiskimage_allocs)

	var creq_switch_allocs *cgoAllocMap
	ref1da77736.req_switch, creq_switch_allocs = (C.uint32_t)(x.ReqSwitch), cgoAllocsUnknown
	allocs1da77736.Borrow(creq_switch_allocs)

	var cstd_err_allocs *cgoAllocMap
	ref1da77736.std_err, cstd_err_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.StdErr)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cstd_err_allocs)

	var cstd_in_allocs *cgoAllocMap
	ref1da77736.std_in, cstd_in_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.StdIn)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cstd_in_allocs)

	var cstd_out_allocs *cgoAllocMap
	ref1da77736.std_out, cstd_out_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.StdOut)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cstd_out_allocs)

	var ctres_req_cnt_allocs *cgoAllocMap
	ref1da77736.tres_req_cnt, ctres_req_cnt_allocs = (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.TresReqCnt)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(ctres_req_cnt_allocs)

	var cwait4switch_allocs *cgoAllocMap
	ref1da77736.wait4switch, cwait4switch_allocs = (C.uint32_t)(x.Wait4switch), cgoAllocsUnknown
	allocs1da77736.Borrow(cwait4switch_allocs)

	var cwckey_allocs *cgoAllocMap
	ref1da77736.wckey, cwckey_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Wckey)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cwckey_allocs)

	var cx11_allocs *cgoAllocMap
	ref1da77736.x11, cx11_allocs = (C.uint16_t)(x.X11), cgoAllocsUnknown
	allocs1da77736.Borrow(cx11_allocs)

	var cx11_magic_cookie_allocs *cgoAllocMap
	ref1da77736.x11_magic_cookie, cx11_magic_cookie_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.X11MagicCookie)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cx11_magic_cookie_allocs)

	var cx11_target_port_allocs *cgoAllocMap
	ref1da77736.x11_target_port, cx11_target_port_allocs = (C.uint16_t)(x.X11TargetPort), cgoAllocsUnknown
	allocs1da77736.Borrow(cx11_target_port_allocs)

	x.ref1da77736 = ref1da77736
	x.allocs1da77736 = allocs1da77736
	return ref1da77736, allocs1da77736

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x JobDescMsg) PassValue() (C.job_desc_msg_t, *cgoAllocMap) {
	if x.ref1da77736 != nil {
		return *x.ref1da77736, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *JobDescMsg) Deref() {
	if x.ref1da77736 == nil {
		return
	}
	hxfc4425b := (*sliceHeader)(unsafe.Pointer(&x.Account))
	hxfc4425b.Data = uintptr(unsafe.Pointer(x.ref1da77736.account))
	hxfc4425b.Cap = 0x7fffffff
	// hxfc4425b.Len = ?

	hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&x.AcctgFreq))
	hxf95e7c8.Data = uintptr(unsafe.Pointer(x.ref1da77736.acctg_freq))
	hxf95e7c8.Cap = 0x7fffffff
	// hxf95e7c8.Len = ?

	hxff2234b := (*sliceHeader)(unsafe.Pointer(&x.AdminComment))
	hxff2234b.Data = uintptr(unsafe.Pointer(x.ref1da77736.admin_comment))
	hxff2234b.Cap = 0x7fffffff
	// hxff2234b.Len = ?

	hxff73280 := (*sliceHeader)(unsafe.Pointer(&x.AllocNode))
	hxff73280.Data = uintptr(unsafe.Pointer(x.ref1da77736.alloc_node))
	hxff73280.Cap = 0x7fffffff
	// hxff73280.Len = ?

	x.AllocRespPort = (uint16)(x.ref1da77736.alloc_resp_port)
	x.AllocSid = (uint32)(x.ref1da77736.alloc_sid)
	x.Argc = (uint32)(x.ref1da77736.argc)
	packSSByte(x.Argv, x.ref1da77736.argv)
	hxfa3f05c := (*sliceHeader)(unsafe.Pointer(&x.ArrayInx))
	hxfa3f05c.Data = uintptr(unsafe.Pointer(x.ref1da77736.array_inx))
	hxfa3f05c.Cap = 0x7fffffff
	// hxfa3f05c.Len = ?

	x.ArrayBitmap = (unsafe.Pointer)(unsafe.Pointer(x.ref1da77736.array_bitmap))
	x.BeginTime = (int)(x.ref1da77736.begin_time)
	x.Bitflags = (uint32)(x.ref1da77736.bitflags)
	hxf0d18b7 := (*sliceHeader)(unsafe.Pointer(&x.BurstBuffer))
	hxf0d18b7.Data = uintptr(unsafe.Pointer(x.ref1da77736.burst_buffer))
	hxf0d18b7.Cap = 0x7fffffff
	// hxf0d18b7.Len = ?

	x.CkptInterval = (uint16)(x.ref1da77736.ckpt_interval)
	hxf2fab0d := (*sliceHeader)(unsafe.Pointer(&x.CkptDir))
	hxf2fab0d.Data = uintptr(unsafe.Pointer(x.ref1da77736.ckpt_dir))
	hxf2fab0d.Cap = 0x7fffffff
	// hxf2fab0d.Len = ?

	hxf69fe70 := (*sliceHeader)(unsafe.Pointer(&x.Clusters))
	hxf69fe70.Data = uintptr(unsafe.Pointer(x.ref1da77736.clusters))
	hxf69fe70.Cap = 0x7fffffff
	// hxf69fe70.Len = ?

	hxf65bf54 := (*sliceHeader)(unsafe.Pointer(&x.ClusterFeatures))
	hxf65bf54.Data = uintptr(unsafe.Pointer(x.ref1da77736.cluster_features))
	hxf65bf54.Cap = 0x7fffffff
	// hxf65bf54.Len = ?

	hxf3b8dbd := (*sliceHeader)(unsafe.Pointer(&x.Comment))
	hxf3b8dbd.Data = uintptr(unsafe.Pointer(x.ref1da77736.comment))
	hxf3b8dbd.Cap = 0x7fffffff
	// hxf3b8dbd.Len = ?

	x.Contiguous = (uint16)(x.ref1da77736.contiguous)
	x.CoreSpec = (uint16)(x.ref1da77736.core_spec)
	hxf7a6dff := (*sliceHeader)(unsafe.Pointer(&x.CpuBind))
	hxf7a6dff.Data = uintptr(unsafe.Pointer(x.ref1da77736.cpu_bind))
	hxf7a6dff.Cap = 0x7fffffff
	// hxf7a6dff.Len = ?

	x.CpuBindType = (uint16)(x.ref1da77736.cpu_bind_type)
	x.CpuFreqMin = (uint32)(x.ref1da77736.cpu_freq_min)
	x.CpuFreqMax = (uint32)(x.ref1da77736.cpu_freq_max)
	x.CpuFreqGov = (uint32)(x.ref1da77736.cpu_freq_gov)
	x.Deadline = (int)(x.ref1da77736.deadline)
	x.DelayBoot = (uint32)(x.ref1da77736.delay_boot)
	hxfe48d67 := (*sliceHeader)(unsafe.Pointer(&x.Dependency))
	hxfe48d67.Data = uintptr(unsafe.Pointer(x.ref1da77736.dependency))
	hxfe48d67.Cap = 0x7fffffff
	// hxfe48d67.Len = ?

	x.EndTime = (int)(x.ref1da77736.end_time)
	packSSByte(x.Environment, x.ref1da77736.environment)
	x.EnvSize = (uint32)(x.ref1da77736.env_size)
	hxf058b18 := (*sliceHeader)(unsafe.Pointer(&x.Extra))
	hxf058b18.Data = uintptr(unsafe.Pointer(x.ref1da77736.extra))
	hxf058b18.Cap = 0x7fffffff
	// hxf058b18.Len = ?

	hxff6bc57 := (*sliceHeader)(unsafe.Pointer(&x.ExcNodes))
	hxff6bc57.Data = uintptr(unsafe.Pointer(x.ref1da77736.exc_nodes))
	hxff6bc57.Cap = 0x7fffffff
	// hxff6bc57.Len = ?

	hxf5fa529 := (*sliceHeader)(unsafe.Pointer(&x.Features))
	hxf5fa529.Data = uintptr(unsafe.Pointer(x.ref1da77736.features))
	hxf5fa529.Cap = 0x7fffffff
	// hxf5fa529.Len = ?

	x.FedSiblingsActive = (uint)(x.ref1da77736.fed_siblings_active)
	x.FedSiblingsViable = (uint)(x.ref1da77736.fed_siblings_viable)
	hxf21690b := (*sliceHeader)(unsafe.Pointer(&x.Gres))
	hxf21690b.Data = uintptr(unsafe.Pointer(x.ref1da77736.gres))
	hxf21690b.Cap = 0x7fffffff
	// hxf21690b.Len = ?

	x.GroupId = (uint32)(x.ref1da77736.group_id)
	x.Immediate = (uint16)(x.ref1da77736.immediate)
	x.JobId = (uint32)(x.ref1da77736.job_id)
	hxf1231c9 := (*sliceHeader)(unsafe.Pointer(&x.JobIdStr))
	hxf1231c9.Data = uintptr(unsafe.Pointer(x.ref1da77736.job_id_str))
	hxf1231c9.Cap = 0x7fffffff
	// hxf1231c9.Len = ?

	x.KillOnNodeFail = (uint16)(x.ref1da77736.kill_on_node_fail)
	hxf04b15b := (*sliceHeader)(unsafe.Pointer(&x.Licenses))
	hxf04b15b.Data = uintptr(unsafe.Pointer(x.ref1da77736.licenses))
	hxf04b15b.Cap = 0x7fffffff
	// hxf04b15b.Len = ?

	x.MailType = (uint16)(x.ref1da77736.mail_type)
	hxf2f888b := (*sliceHeader)(unsafe.Pointer(&x.MailUser))
	hxf2f888b.Data = uintptr(unsafe.Pointer(x.ref1da77736.mail_user))
	hxf2f888b.Cap = 0x7fffffff
	// hxf2f888b.Len = ?

	hxf5d1de2 := (*sliceHeader)(unsafe.Pointer(&x.McsLabel))
	hxf5d1de2.Data = uintptr(unsafe.Pointer(x.ref1da77736.mcs_label))
	hxf5d1de2.Cap = 0x7fffffff
	// hxf5d1de2.Len = ?

	hxfe53d34 := (*sliceHeader)(unsafe.Pointer(&x.MemBind))
	hxfe53d34.Data = uintptr(unsafe.Pointer(x.ref1da77736.mem_bind))
	hxfe53d34.Cap = 0x7fffffff
	// hxfe53d34.Len = ?

	x.MemBindType = (uint16)(x.ref1da77736.mem_bind_type)
	hxf547023 := (*sliceHeader)(unsafe.Pointer(&x.Name))
	hxf547023.Data = uintptr(unsafe.Pointer(x.ref1da77736.name))
	hxf547023.Cap = 0x7fffffff
	// hxf547023.Len = ?

	hxf5ebb88 := (*sliceHeader)(unsafe.Pointer(&x.Network))
	hxf5ebb88.Data = uintptr(unsafe.Pointer(x.ref1da77736.network))
	hxf5ebb88.Cap = 0x7fffffff
	// hxf5ebb88.Len = ?

	x.Nice = (uint32)(x.ref1da77736.nice)
	x.NumTasks = (uint32)(x.ref1da77736.num_tasks)
	x.OpenMode = (byte)(x.ref1da77736.open_mode)
	hxff20e84 := (*sliceHeader)(unsafe.Pointer(&x.OriginCluster))
	hxff20e84.Data = uintptr(unsafe.Pointer(x.ref1da77736.origin_cluster))
	hxff20e84.Cap = 0x7fffffff
	// hxff20e84.Len = ?

	x.OtherPort = (uint16)(x.ref1da77736.other_port)
	x.Overcommit = (byte)(x.ref1da77736.overcommit)
	x.PackJobOffset = (uint32)(x.ref1da77736.pack_job_offset)
	hxfa26a4d := (*sliceHeader)(unsafe.Pointer(&x.Partition))
	hxfa26a4d.Data = uintptr(unsafe.Pointer(x.ref1da77736.partition))
	hxfa26a4d.Cap = 0x7fffffff
	// hxfa26a4d.Len = ?

	x.PlaneSize = (uint16)(x.ref1da77736.plane_size)
	x.PowerFlags = (byte)(x.ref1da77736.power_flags)
	x.Priority = (uint32)(x.ref1da77736.priority)
	x.Profile = (uint32)(x.ref1da77736.profile)
	hxfe48098 := (*sliceHeader)(unsafe.Pointer(&x.Qos))
	hxfe48098.Data = uintptr(unsafe.Pointer(x.ref1da77736.qos))
	hxfe48098.Cap = 0x7fffffff
	// hxfe48098.Len = ?

	x.Reboot = (uint16)(x.ref1da77736.reboot)
	hxffe3496 := (*sliceHeader)(unsafe.Pointer(&x.RespHost))
	hxffe3496.Data = uintptr(unsafe.Pointer(x.ref1da77736.resp_host))
	hxffe3496.Cap = 0x7fffffff
	// hxffe3496.Len = ?

	x.RestartCnt = (uint16)(x.ref1da77736.restart_cnt)
	hxf5d48a6 := (*sliceHeader)(unsafe.Pointer(&x.ReqNodes))
	hxf5d48a6.Data = uintptr(unsafe.Pointer(x.ref1da77736.req_nodes))
	hxf5d48a6.Cap = 0x7fffffff
	// hxf5d48a6.Len = ?

	x.Requeue = (uint16)(x.ref1da77736.requeue)
	hxf685469 := (*sliceHeader)(unsafe.Pointer(&x.Reservation))
	hxf685469.Data = uintptr(unsafe.Pointer(x.ref1da77736.reservation))
	hxf685469.Cap = 0x7fffffff
	// hxf685469.Len = ?

	hxf03a9a7 := (*sliceHeader)(unsafe.Pointer(&x.Script))
	hxf03a9a7.Data = uintptr(unsafe.Pointer(x.ref1da77736.script))
	hxf03a9a7.Cap = 0x7fffffff
	// hxf03a9a7.Len = ?

	x.Shared = (uint16)(x.ref1da77736.shared)
	packSSByte(x.SpankJobEnv, x.ref1da77736.spank_job_env)
	x.SpankJobEnvSize = (uint32)(x.ref1da77736.spank_job_env_size)
	x.TaskDist = (uint32)(x.ref1da77736.task_dist)
	x.TimeLimit = (uint32)(x.ref1da77736.time_limit)
	x.TimeMin = (uint32)(x.ref1da77736.time_min)
	x.UserId = (uint32)(x.ref1da77736.user_id)
	x.WaitAllNodes = (uint16)(x.ref1da77736.wait_all_nodes)
	x.WarnFlags = (uint16)(x.ref1da77736.warn_flags)
	x.WarnSignal = (uint16)(x.ref1da77736.warn_signal)
	x.WarnTime = (uint16)(x.ref1da77736.warn_time)
	hxfe93325 := (*sliceHeader)(unsafe.Pointer(&x.WorkDir))
	hxfe93325.Data = uintptr(unsafe.Pointer(x.ref1da77736.work_dir))
	hxfe93325.Cap = 0x7fffffff
	// hxfe93325.Len = ?

	x.CpusPerTask = (uint16)(x.ref1da77736.cpus_per_task)
	x.MinCpus = (uint32)(x.ref1da77736.min_cpus)
	x.MaxCpus = (uint32)(x.ref1da77736.max_cpus)
	x.MinNodes = (uint32)(x.ref1da77736.min_nodes)
	x.MaxNodes = (uint32)(x.ref1da77736.max_nodes)
	x.BoardsPerNode = (uint16)(x.ref1da77736.boards_per_node)
	x.SocketsPerBoard = (uint16)(x.ref1da77736.sockets_per_board)
	x.SocketsPerNode = (uint16)(x.ref1da77736.sockets_per_node)
	x.CoresPerSocket = (uint16)(x.ref1da77736.cores_per_socket)
	x.ThreadsPerCore = (uint16)(x.ref1da77736.threads_per_core)
	x.NtasksPerNode = (uint16)(x.ref1da77736.ntasks_per_node)
	x.NtasksPerSocket = (uint16)(x.ref1da77736.ntasks_per_socket)
	x.NtasksPerCore = (uint16)(x.ref1da77736.ntasks_per_core)
	x.NtasksPerBoard = (uint16)(x.ref1da77736.ntasks_per_board)
	x.PnMinCpus = (uint16)(x.ref1da77736.pn_min_cpus)
	x.PnMinMemory = (uint)(x.ref1da77736.pn_min_memory)
	x.PnMinTmpDisk = (uint32)(x.ref1da77736.pn_min_tmp_disk)
	x.Geometry = *(*[5]uint16)(unsafe.Pointer(&x.ref1da77736.geometry))
	x.ConnType = *(*[5]uint16)(unsafe.Pointer(&x.ref1da77736.conn_type))
	x.Rotate = (uint16)(x.ref1da77736.rotate)
	hxf09ea94 := (*sliceHeader)(unsafe.Pointer(&x.Blrtsimage))
	hxf09ea94.Data = uintptr(unsafe.Pointer(x.ref1da77736.blrtsimage))
	hxf09ea94.Cap = 0x7fffffff
	// hxf09ea94.Len = ?

	hxfd687ee := (*sliceHeader)(unsafe.Pointer(&x.Linuximage))
	hxfd687ee.Data = uintptr(unsafe.Pointer(x.ref1da77736.linuximage))
	hxfd687ee.Cap = 0x7fffffff
	// hxfd687ee.Len = ?

	hxf15a567 := (*sliceHeader)(unsafe.Pointer(&x.Mloaderimage))
	hxf15a567.Data = uintptr(unsafe.Pointer(x.ref1da77736.mloaderimage))
	hxf15a567.Cap = 0x7fffffff
	// hxf15a567.Len = ?

	hxf8aebb5 := (*sliceHeader)(unsafe.Pointer(&x.Ramdiskimage))
	hxf8aebb5.Data = uintptr(unsafe.Pointer(x.ref1da77736.ramdiskimage))
	hxf8aebb5.Cap = 0x7fffffff
	// hxf8aebb5.Len = ?

	x.ReqSwitch = (uint32)(x.ref1da77736.req_switch)
	hxf5d30cf := (*sliceHeader)(unsafe.Pointer(&x.StdErr))
	hxf5d30cf.Data = uintptr(unsafe.Pointer(x.ref1da77736.std_err))
	hxf5d30cf.Cap = 0x7fffffff
	// hxf5d30cf.Len = ?

	hxf882e98 := (*sliceHeader)(unsafe.Pointer(&x.StdIn))
	hxf882e98.Data = uintptr(unsafe.Pointer(x.ref1da77736.std_in))
	hxf882e98.Cap = 0x7fffffff
	// hxf882e98.Len = ?

	hxf992404 := (*sliceHeader)(unsafe.Pointer(&x.StdOut))
	hxf992404.Data = uintptr(unsafe.Pointer(x.ref1da77736.std_out))
	hxf992404.Cap = 0x7fffffff
	// hxf992404.Len = ?

	hxf8e0dd2 := (*sliceHeader)(unsafe.Pointer(&x.TresReqCnt))
	hxf8e0dd2.Data = uintptr(unsafe.Pointer(x.ref1da77736.tres_req_cnt))
	hxf8e0dd2.Cap = 0x7fffffff
	// hxf8e0dd2.Len = ?

	x.Wait4switch = (uint32)(x.ref1da77736.wait4switch)
	hxf44d909 := (*sliceHeader)(unsafe.Pointer(&x.Wckey))
	hxf44d909.Data = uintptr(unsafe.Pointer(x.ref1da77736.wckey))
	hxf44d909.Cap = 0x7fffffff
	// hxf44d909.Len = ?

	x.X11 = (uint16)(x.ref1da77736.x11)
	hxfa835e7 := (*sliceHeader)(unsafe.Pointer(&x.X11MagicCookie))
	hxfa835e7.Data = uintptr(unsafe.Pointer(x.ref1da77736.x11_magic_cookie))
	hxfa835e7.Cap = 0x7fffffff
	// hxfa835e7.Len = ?

	x.X11TargetPort = (uint16)(x.ref1da77736.x11_target_port)
}

// allocSubmitResponseMsgMemory allocates memory for type C.submit_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSubmitResponseMsgMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSubmitResponseMsgValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSubmitResponseMsgValue = unsafe.Sizeof([1]C.submit_response_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *SubmitResponseMsg) Ref() *C.submit_response_msg_t {
	if x == nil {
		return nil
	}
	return x.ref6c6e601
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *SubmitResponseMsg) Free() {
	if x != nil && x.allocs6c6e601 != nil {
		x.allocs6c6e601.(*cgoAllocMap).Free()
		x.ref6c6e601 = nil
	}
}

// NewSubmitResponseMsgRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewSubmitResponseMsgRef(ref unsafe.Pointer) *SubmitResponseMsg {
	if ref == nil {
		return nil
	}
	obj := new(SubmitResponseMsg)
	obj.ref6c6e601 = (*C.submit_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *SubmitResponseMsg) PassRef() (*C.submit_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6c6e601 != nil {
		return x.ref6c6e601, nil
	}
	mem6c6e601 := allocSubmitResponseMsgMemory(1)
	ref6c6e601 := (*C.submit_response_msg_t)(mem6c6e601)
	allocs6c6e601 := new(cgoAllocMap)
	allocs6c6e601.Add(mem6c6e601)

	var cjob_id_allocs *cgoAllocMap
	ref6c6e601.job_id, cjob_id_allocs = (C.uint32_t)(x.JobId), cgoAllocsUnknown
	allocs6c6e601.Borrow(cjob_id_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref6c6e601.step_id, cstep_id_allocs = (C.uint32_t)(x.StepId), cgoAllocsUnknown
	allocs6c6e601.Borrow(cstep_id_allocs)

	var cerror_code_allocs *cgoAllocMap
	ref6c6e601.error_code, cerror_code_allocs = (C.uint32_t)(x.ErrorCode), cgoAllocsUnknown
	allocs6c6e601.Borrow(cerror_code_allocs)

	var cjob_submit_user_msg_allocs *cgoAllocMap
	ref6c6e601.job_submit_user_msg, cjob_submit_user_msg_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.JobSubmitUserMsg)).Data)), cgoAllocsUnknown
	allocs6c6e601.Borrow(cjob_submit_user_msg_allocs)

	x.ref6c6e601 = ref6c6e601
	x.allocs6c6e601 = allocs6c6e601
	return ref6c6e601, allocs6c6e601

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x SubmitResponseMsg) PassValue() (C.submit_response_msg_t, *cgoAllocMap) {
	if x.ref6c6e601 != nil {
		return *x.ref6c6e601, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *SubmitResponseMsg) Deref() {
	if x.ref6c6e601 == nil {
		return
	}
	x.JobId = (uint32)(x.ref6c6e601.job_id)
	x.StepId = (uint32)(x.ref6c6e601.step_id)
	x.ErrorCode = (uint32)(x.ref6c6e601.error_code)
	hxf8eae10 := (*sliceHeader)(unsafe.Pointer(&x.JobSubmitUserMsg))
	hxf8eae10.Data = uintptr(unsafe.Pointer(x.ref6c6e601.job_submit_user_msg))
	hxf8eae10.Cap = 0x7fffffff
	// hxf8eae10.Len = ?

}
