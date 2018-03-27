// WARNING: This file has automatically been generated on Wed, 28 Mar 2018 01:12:32 CEST.
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

// slurm_hostlist_create function as declared in slurm/slurm.h:1126
func slurm_hostlist_create(hostlist string) hostlist_t {
	chostlist, _ := unpackPCharString(hostlist)
	__ret := C.slurm_hostlist_create(chostlist)
	__v := *(*hostlist_t)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_hostlist_count function as declared in slurm/slurm.h:1132
func slurm_hostlist_count(hl hostlist_t) int32 {
	chl, _ := *(*C.hostlist_t)(unsafe.Pointer(&hl)), cgoAllocsUnknown
	__ret := C.slurm_hostlist_count(chl)
	__v := (int32)(__ret)
	return __v
}

// slurm_hostlist_destroy function as declared in slurm/slurm.h:1139
func slurm_hostlist_destroy(hl hostlist_t) {
	chl, _ := *(*C.hostlist_t)(unsafe.Pointer(&hl)), cgoAllocsUnknown
	C.slurm_hostlist_destroy(chl)
}

// slurm_hostlist_find function as declared in slurm/slurm.h:1148
func slurm_hostlist_find(hl hostlist_t, hostname string) int32 {
	chl, _ := *(*C.hostlist_t)(unsafe.Pointer(&hl)), cgoAllocsUnknown
	chostname, _ := unpackPCharString(hostname)
	__ret := C.slurm_hostlist_find(chl, chostname)
	__v := (int32)(__ret)
	return __v
}

// slurm_hostlist_push function as declared in slurm/slurm.h:1159
func slurm_hostlist_push(hl hostlist_t, hosts string) int32 {
	chl, _ := *(*C.hostlist_t)(unsafe.Pointer(&hl)), cgoAllocsUnknown
	chosts, _ := unpackPCharString(hosts)
	__ret := C.slurm_hostlist_push(chl, chosts)
	__v := (int32)(__ret)
	return __v
}

// slurm_hostlist_push_host function as declared in slurm/slurm.h:1169
func slurm_hostlist_push_host(hl hostlist_t, host string) int32 {
	chl, _ := *(*C.hostlist_t)(unsafe.Pointer(&hl)), cgoAllocsUnknown
	chost, _ := unpackPCharString(host)
	__ret := C.slurm_hostlist_push_host(chl, chost)
	__v := (int32)(__ret)
	return __v
}

// slurm_hostlist_ranged_string function as declared in slurm/slurm.h:1182
func slurm_hostlist_ranged_string(hl hostlist_t, n size_t, buf []byte) ssize_t {
	chl, _ := *(*C.hostlist_t)(unsafe.Pointer(&hl)), cgoAllocsUnknown
	cn, _ := (C.size_t)(n), cgoAllocsUnknown
	cbuf, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buf)).Data)), cgoAllocsUnknown
	__ret := C.slurm_hostlist_ranged_string(chl, cn, cbuf)
	__v := (ssize_t)(__ret)
	return __v
}

// slurm_hostlist_ranged_string_malloc function as declared in slurm/slurm.h:1190
func slurm_hostlist_ranged_string_malloc(hl hostlist_t) *byte {
	chl, _ := *(*C.hostlist_t)(unsafe.Pointer(&hl)), cgoAllocsUnknown
	__ret := C.slurm_hostlist_ranged_string_malloc(chl)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_hostlist_ranged_string_xmalloc function as declared in slurm/slurm.h:1200
func slurm_hostlist_ranged_string_xmalloc(hl hostlist_t) *byte {
	chl, _ := *(*C.hostlist_t)(unsafe.Pointer(&hl)), cgoAllocsUnknown
	__ret := C.slurm_hostlist_ranged_string_xmalloc(chl)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_hostlist_shift function as declared in slurm/slurm.h:1211
func slurm_hostlist_shift(hl hostlist_t) *byte {
	chl, _ := *(*C.hostlist_t)(unsafe.Pointer(&hl)), cgoAllocsUnknown
	__ret := C.slurm_hostlist_shift(chl)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_hostlist_uniq function as declared in slurm/slurm.h:1218
func slurm_hostlist_uniq(hl hostlist_t) {
	chl, _ := *(*C.hostlist_t)(unsafe.Pointer(&hl)), cgoAllocsUnknown
	C.slurm_hostlist_uniq(chl)
}

// slurm_list_append function as declared in slurm/slurm.h:1268
func slurm_list_append(l List, x unsafe.Pointer) unsafe.Pointer {
	cl, _ := *(*C.List)(unsafe.Pointer(&l)), cgoAllocsUnknown
	cx, _ := x, cgoAllocsUnknown
	__ret := C.slurm_list_append(cl, cx)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_list_count function as declared in slurm/slurm.h:1274
func slurm_list_count(l List) int32 {
	cl, _ := *(*C.List)(unsafe.Pointer(&l)), cgoAllocsUnknown
	__ret := C.slurm_list_count(cl)
	__v := (int32)(__ret)
	return __v
}

// slurm_list_create function as declared in slurm/slurm.h:1285
func slurm_list_create(f ListDelF) List {
	cf, _ := f.PassValue()
	__ret := C.slurm_list_create(cf)
	__v := *(*List)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_list_destroy function as declared in slurm/slurm.h:1293
func slurm_list_destroy(l List) {
	cl, _ := *(*C.List)(unsafe.Pointer(&l)), cgoAllocsUnknown
	C.slurm_list_destroy(cl)
}

// slurm_list_find function as declared in slurm/slurm.h:1304
func slurm_list_find(i ListIterator, f ListFindF, key unsafe.Pointer) unsafe.Pointer {
	ci, _ := *(*C.ListIterator)(unsafe.Pointer(&i)), cgoAllocsUnknown
	cf, _ := f.PassValue()
	ckey, _ := key, cgoAllocsUnknown
	__ret := C.slurm_list_find(ci, cf, ckey)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_list_is_empty function as declared in slurm/slurm.h:1310
func slurm_list_is_empty(l List) int32 {
	cl, _ := *(*C.List)(unsafe.Pointer(&l)), cgoAllocsUnknown
	__ret := C.slurm_list_is_empty(cl)
	__v := (int32)(__ret)
	return __v
}

// slurm_list_iterator_create function as declared in slurm/slurm.h:1316
func slurm_list_iterator_create(l List) ListIterator {
	cl, _ := *(*C.List)(unsafe.Pointer(&l)), cgoAllocsUnknown
	__ret := C.slurm_list_iterator_create(cl)
	__v := *(*ListIterator)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_list_iterator_reset function as declared in slurm/slurm.h:1323
func slurm_list_iterator_reset(i ListIterator) {
	ci, _ := *(*C.ListIterator)(unsafe.Pointer(&i)), cgoAllocsUnknown
	C.slurm_list_iterator_reset(ci)
}

// slurm_list_iterator_destroy function as declared in slurm/slurm.h:1330
func slurm_list_iterator_destroy(i ListIterator) {
	ci, _ := *(*C.ListIterator)(unsafe.Pointer(&i)), cgoAllocsUnknown
	C.slurm_list_iterator_destroy(ci)
}

// slurm_list_next function as declared in slurm/slurm.h:1339
func slurm_list_next(i ListIterator) unsafe.Pointer {
	ci, _ := *(*C.ListIterator)(unsafe.Pointer(&i)), cgoAllocsUnknown
	__ret := C.slurm_list_next(ci)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_list_sort function as declared in slurm/slurm.h:1347
func slurm_list_sort(l List, f ListCmpF) {
	cl, _ := *(*C.List)(unsafe.Pointer(&l)), cgoAllocsUnknown
	cf, _ := f.PassValue()
	C.slurm_list_sort(cl, cf)
}

// slurm_list_pop function as declared in slurm/slurm.h:1354
func slurm_list_pop(l List) unsafe.Pointer {
	cl, _ := *(*C.List)(unsafe.Pointer(&l)), cgoAllocsUnknown
	__ret := C.slurm_list_pop(cl)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_print_block_info_msg function as declared in slurm/slurm.h:2450
func slurm_print_block_info_msg(out []FILE, info_ptr []block_info_msg_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cinfo_ptr, _ := unpackArgSBlock_info_msg_t(info_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_block_info_msg(cout, cinfo_ptr, cone_liner)
	packSBlock_info_msg_t(info_ptr, cinfo_ptr)
	packSFILE(out, cout)
}

// slurm_print_block_info function as declared in slurm/slurm.h:2461
func slurm_print_block_info(out []FILE, bg_info_ptr []block_info_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cbg_info_ptr, _ := unpackArgSBlock_info_t(bg_info_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_block_info(cout, cbg_info_ptr, cone_liner)
	packSBlock_info_t(bg_info_ptr, cbg_info_ptr)
	packSFILE(out, cout)
}

// slurm_sprint_block_info function as declared in slurm/slurm.h:2473
func slurm_sprint_block_info(bg_info_ptr []block_info_t, one_liner int32) *byte {
	cbg_info_ptr, _ := unpackArgSBlock_info_t(bg_info_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	__ret := C.slurm_sprint_block_info(cbg_info_ptr, cone_liner)
	packSBlock_info_t(bg_info_ptr, cbg_info_ptr)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_load_block_info function as declared in slurm/slurm.h:2485
func slurm_load_block_info(update_time time_t, block_info_msg_pptr [][]block_info_msg_t, show_flags uint16_t) int32 {
	cupdate_time, _ := (C.time_t)(update_time), cgoAllocsUnknown
	cblock_info_msg_pptr, _ := unpackArgSSBlock_info_msg_t(block_info_msg_pptr)
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	__ret := C.slurm_load_block_info(cupdate_time, cblock_info_msg_pptr, cshow_flags)
	packSSBlock_info_msg_t(block_info_msg_pptr, cblock_info_msg_pptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_block_info_msg function as declared in slurm/slurm.h:2495
func slurm_free_block_info_msg(block_info_msg []block_info_msg_t) {
	cblock_info_msg, _ := unpackArgSBlock_info_msg_t(block_info_msg)
	C.slurm_free_block_info_msg(cblock_info_msg)
	packSBlock_info_msg_t(block_info_msg, cblock_info_msg)
}

// slurm_update_block function as declared in slurm/slurm.h:2498
func slurm_update_block(block_msg []update_block_msg_t) int32 {
	cblock_msg, _ := unpackArgSUpdate_block_msg_t(block_msg)
	__ret := C.slurm_update_block(cblock_msg)
	packSUpdate_block_msg_t(block_msg, cblock_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_init_update_block_msg function as declared in slurm/slurm.h:2500
func slurm_init_update_block_msg(update_block_msg []update_block_msg_t) {
	cupdate_block_msg, _ := unpackArgSUpdate_block_msg_t(update_block_msg)
	C.slurm_init_update_block_msg(cupdate_block_msg)
	packSUpdate_block_msg_t(update_block_msg, cupdate_block_msg)
}

// slurm_init_job_desc_msg function as declared in slurm/slurm.h:3208
func slurm_init_job_desc_msg(job_desc_msg []job_desc_msg_t) {
	cjob_desc_msg, _ := unpackArgSJob_desc_msg_t(job_desc_msg)
	C.slurm_init_job_desc_msg(cjob_desc_msg)
	packSJob_desc_msg_t(job_desc_msg, cjob_desc_msg)
}

// slurm_allocate_resources function as declared in slurm/slurm.h:3221
func slurm_allocate_resources(job_desc_msg []job_desc_msg_t, job_alloc_resp_msg [][]resource_allocation_response_msg_t) int32 {
	cjob_desc_msg, _ := unpackArgSJob_desc_msg_t(job_desc_msg)
	cjob_alloc_resp_msg, _ := unpackArgSSResource_allocation_response_msg_t(job_alloc_resp_msg)
	__ret := C.slurm_allocate_resources(cjob_desc_msg, cjob_alloc_resp_msg)
	packSSResource_allocation_response_msg_t(job_alloc_resp_msg, cjob_alloc_resp_msg)
	packSJob_desc_msg_t(job_desc_msg, cjob_desc_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_allocate_resources_blocking function as declared in slurm/slurm.h:3242
func slurm_allocate_resources_blocking(user_req []job_desc_msg_t, timeout time_t, pending_callback *func(job_id uint32_t)) *resource_allocation_response_msg_t {
	cuser_req, _ := unpackArgSJob_desc_msg_t(user_req)
	ctimeout, _ := (C.time_t)(timeout), cgoAllocsUnknown
	cpending_callback, _ := pending_callback.PassRef()
	__ret := C.slurm_allocate_resources_blocking(cuser_req, ctimeout, cpending_callback)
	packSJob_desc_msg_t(user_req, cuser_req)
	__v := Newresource_allocation_response_msg_tRef(unsafe.Pointer(__ret))
	return __v
}

// slurm_free_resource_allocation_response_msg function as declared in slurm/slurm.h:3253
func slurm_free_resource_allocation_response_msg(msg []resource_allocation_response_msg_t) {
	cmsg, _ := unpackArgSResource_allocation_response_msg_t(msg)
	C.slurm_free_resource_allocation_response_msg(cmsg)
	packSResource_allocation_response_msg_t(msg, cmsg)
}

// slurm_allocate_pack_job_blocking function as declared in slurm/slurm.h:3274
func slurm_allocate_pack_job_blocking(job_req_list List, timeout time_t, pending_callback *func(job_id uint32_t)) List {
	cjob_req_list, _ := *(*C.List)(unsafe.Pointer(&job_req_list)), cgoAllocsUnknown
	ctimeout, _ := (C.time_t)(timeout), cgoAllocsUnknown
	cpending_callback, _ := pending_callback.PassRef()
	__ret := C.slurm_allocate_pack_job_blocking(cjob_req_list, ctimeout, cpending_callback)
	__v := *(*List)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_allocation_lookup function as declared in slurm/slurm.h:3285
func slurm_allocation_lookup(job_id uint32_t, resp [][]resource_allocation_response_msg_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cresp, _ := unpackArgSSResource_allocation_response_msg_t(resp)
	__ret := C.slurm_allocation_lookup(cjob_id, cresp)
	packSSResource_allocation_response_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_pack_job_lookup function as declared in slurm/slurm.h:3298
func slurm_pack_job_lookup(jobid uint32_t, resp []List) int32 {
	cjobid, _ := (C.uint32_t)(jobid), cgoAllocsUnknown
	cresp, _ := (*C.List)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&resp)).Data)), cgoAllocsUnknown
	__ret := C.slurm_pack_job_lookup(cjobid, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_read_hostfile function as declared in slurm/slurm.h:3315
func slurm_read_hostfile(filename []byte, n int32) *byte {
	cfilename, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&filename)).Data)), cgoAllocsUnknown
	cn, _ := (C.int)(n), cgoAllocsUnknown
	__ret := C.slurm_read_hostfile(cfilename, cn)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_allocation_msg_thr_create function as declared in slurm/slurm.h:3325
func slurm_allocation_msg_thr_create(port []uint16_t, callbacks []slurm_allocation_callbacks_t) *allocation_msg_thread_t {
	cport, _ := (*C.uint16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&port)).Data)), cgoAllocsUnknown
	ccallbacks, _ := (*C.slurm_allocation_callbacks_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&callbacks)).Data)), cgoAllocsUnknown
	__ret := C.slurm_allocation_msg_thr_create(cport, ccallbacks)
	__v := *(**allocation_msg_thread_t)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_allocation_msg_thr_destroy function as declared in slurm/slurm.h:3335
func slurm_allocation_msg_thr_destroy(msg_thr []allocation_msg_thread_t) {
	cmsg_thr, _ := (*C.allocation_msg_thread_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&msg_thr)).Data)), cgoAllocsUnknown
	C.slurm_allocation_msg_thr_destroy(cmsg_thr)
}

// slurm_submit_batch_job function as declared in slurm/slurm.h:3344
func slurm_submit_batch_job(job_desc_msg []job_desc_msg_t, slurm_alloc_msg [][]submit_response_msg_t) int32 {
	cjob_desc_msg, _ := unpackArgSJob_desc_msg_t(job_desc_msg)
	cslurm_alloc_msg, _ := unpackArgSSSubmit_response_msg_t(slurm_alloc_msg)
	__ret := C.slurm_submit_batch_job(cjob_desc_msg, cslurm_alloc_msg)
	packSSSubmit_response_msg_t(slurm_alloc_msg, cslurm_alloc_msg)
	packSJob_desc_msg_t(job_desc_msg, cjob_desc_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_submit_batch_pack_job function as declared in slurm/slurm.h:3355
func slurm_submit_batch_pack_job(job_req_list List, slurm_alloc_msg [][]submit_response_msg_t) int32 {
	cjob_req_list, _ := *(*C.List)(unsafe.Pointer(&job_req_list)), cgoAllocsUnknown
	cslurm_alloc_msg, _ := unpackArgSSSubmit_response_msg_t(slurm_alloc_msg)
	__ret := C.slurm_submit_batch_pack_job(cjob_req_list, cslurm_alloc_msg)
	packSSSubmit_response_msg_t(slurm_alloc_msg, cslurm_alloc_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_submit_response_response_msg function as declared in slurm/slurm.h:3364
func slurm_free_submit_response_response_msg(msg []submit_response_msg_t) {
	cmsg, _ := unpackArgSSubmit_response_msg_t(msg)
	C.slurm_free_submit_response_response_msg(cmsg)
	packSSubmit_response_msg_t(msg, cmsg)
}

// slurm_job_batch_script function as declared in slurm/slurm.h:3370
func slurm_job_batch_script(out []FILE, jobid uint32_t) int32 {
	cout, _ := unpackArgSFILE(out)
	cjobid, _ := (C.uint32_t)(jobid), cgoAllocsUnknown
	__ret := C.slurm_job_batch_script(cout, cjobid)
	packSFILE(out, cout)
	__v := (int32)(__ret)
	return __v
}

// slurm_job_will_run function as declared in slurm/slurm.h:3378
func slurm_job_will_run(job_desc_msg []job_desc_msg_t) int32 {
	cjob_desc_msg, _ := unpackArgSJob_desc_msg_t(job_desc_msg)
	__ret := C.slurm_job_will_run(cjob_desc_msg)
	packSJob_desc_msg_t(job_desc_msg, cjob_desc_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_pack_job_will_run function as declared in slurm/slurm.h:3387
func slurm_pack_job_will_run(job_req_list List) int32 {
	cjob_req_list, _ := *(*C.List)(unsafe.Pointer(&job_req_list)), cgoAllocsUnknown
	__ret := C.slurm_pack_job_will_run(cjob_req_list)
	__v := (int32)(__ret)
	return __v
}

// slurm_job_will_run2 function as declared in slurm/slurm.h:3398
func slurm_job_will_run2(req []job_desc_msg_t, will_run_resp [][]will_run_response_msg_t) int32 {
	creq, _ := unpackArgSJob_desc_msg_t(req)
	cwill_run_resp, _ := unpackArgSSWill_run_response_msg_t(will_run_resp)
	__ret := C.slurm_job_will_run2(creq, cwill_run_resp)
	packSSWill_run_response_msg_t(will_run_resp, cwill_run_resp)
	packSJob_desc_msg_t(req, creq)
	__v := (int32)(__ret)
	return __v
}

// slurm_sbcast_lookup function as declared in slurm/slurm.h:3411
func slurm_sbcast_lookup(job_id uint32_t, pack_job_offset uint32_t, step_id uint32_t, info [][]job_sbcast_cred_msg_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cpack_job_offset, _ := (C.uint32_t)(pack_job_offset), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	cinfo, _ := unpackArgSSJob_sbcast_cred_msg_t(info)
	__ret := C.slurm_sbcast_lookup(cjob_id, cpack_job_offset, cstep_id, cinfo)
	packSSJob_sbcast_cred_msg_t(info, cinfo)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_sbcast_cred_msg function as declared in slurm/slurm.h:3414
func slurm_free_sbcast_cred_msg(msg []job_sbcast_cred_msg_t) {
	cmsg, _ := unpackArgSJob_sbcast_cred_msg_t(msg)
	C.slurm_free_sbcast_cred_msg(cmsg)
	packSJob_sbcast_cred_msg_t(msg, cmsg)
}

// slurm_load_licenses function as declared in slurm/slurm.h:3423
func slurm_load_licenses(arg0 time_t, arg1 [][]license_info_msg_t, arg2 uint16_t) int32 {
	carg0, _ := (C.time_t)(arg0), cgoAllocsUnknown
	carg1, _ := unpackArgSSLicense_info_msg_t(arg1)
	carg2, _ := (C.uint16_t)(arg2), cgoAllocsUnknown
	__ret := C.slurm_load_licenses(carg0, carg1, carg2)
	packSSLicense_info_msg_t(arg1, carg1)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_license_info_msg function as declared in slurm/slurm.h:3424
func slurm_free_license_info_msg(arg0 []license_info_msg_t) {
	carg0, _ := unpackArgSLicense_info_msg_t(arg0)
	C.slurm_free_license_info_msg(carg0)
	packSLicense_info_msg_t(arg0, carg0)
}

// slurm_load_assoc_mgr_info function as declared in slurm/slurm.h:3432
func slurm_load_assoc_mgr_info(arg0 []assoc_mgr_info_request_msg_t, arg1 [][]assoc_mgr_info_msg_t) int32 {
	carg0, _ := unpackArgSAssoc_mgr_info_request_msg_t(arg0)
	carg1, _ := unpackArgSSAssoc_mgr_info_msg_t(arg1)
	__ret := C.slurm_load_assoc_mgr_info(carg0, carg1)
	packSSAssoc_mgr_info_msg_t(arg1, carg1)
	packSAssoc_mgr_info_request_msg_t(arg0, carg0)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_assoc_mgr_info_msg function as declared in slurm/slurm.h:3434
func slurm_free_assoc_mgr_info_msg(arg0 []assoc_mgr_info_msg_t) {
	carg0, _ := unpackArgSAssoc_mgr_info_msg_t(arg0)
	C.slurm_free_assoc_mgr_info_msg(carg0)
	packSAssoc_mgr_info_msg_t(arg0, carg0)
}

// slurm_free_assoc_mgr_info_request_members function as declared in slurm/slurm.h:3435
func slurm_free_assoc_mgr_info_request_members(arg0 []assoc_mgr_info_request_msg_t) {
	carg0, _ := unpackArgSAssoc_mgr_info_request_msg_t(arg0)
	C.slurm_free_assoc_mgr_info_request_members(carg0)
	packSAssoc_mgr_info_request_msg_t(arg0, carg0)
}

// slurm_free_assoc_mgr_info_request_msg function as declared in slurm/slurm.h:3436
func slurm_free_assoc_mgr_info_request_msg(arg0 []assoc_mgr_info_request_msg_t) {
	carg0, _ := unpackArgSAssoc_mgr_info_request_msg_t(arg0)
	C.slurm_free_assoc_mgr_info_request_msg(carg0)
	packSAssoc_mgr_info_request_msg_t(arg0, carg0)
}

// slurm_kill_job function as declared in slurm/slurm.h:3476
func slurm_kill_job(job_id uint32_t, signal uint16_t, flags uint16_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	csignal, _ := (C.uint16_t)(signal), cgoAllocsUnknown
	cflags, _ := (C.uint16_t)(flags), cgoAllocsUnknown
	__ret := C.slurm_kill_job(cjob_id, csignal, cflags)
	__v := (int32)(__ret)
	return __v
}

// slurm_kill_job_step function as declared in slurm/slurm.h:3485
func slurm_kill_job_step(job_id uint32_t, step_id uint32_t, signal uint16_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	csignal, _ := (C.uint16_t)(signal), cgoAllocsUnknown
	__ret := C.slurm_kill_job_step(cjob_id, cstep_id, csignal)
	__v := (int32)(__ret)
	return __v
}

// slurm_kill_job2 function as declared in slurm/slurm.h:3491
func slurm_kill_job2(job_id string, signal uint16_t, flags uint16_t) int32 {
	cjob_id, _ := unpackPCharString(job_id)
	csignal, _ := (C.uint16_t)(signal), cgoAllocsUnknown
	cflags, _ := (C.uint16_t)(flags), cgoAllocsUnknown
	__ret := C.slurm_kill_job2(cjob_id, csignal, cflags)
	__v := (int32)(__ret)
	return __v
}

// slurm_kill_job_msg function as declared in slurm/slurm.h:3500
func slurm_kill_job_msg(msg_type uint16_t, kill_msg []job_step_kill_msg_t) int32 {
	cmsg_type, _ := (C.uint16_t)(msg_type), cgoAllocsUnknown
	ckill_msg, _ := unpackArgSJob_step_kill_msg_t(kill_msg)
	__ret := C.slurm_kill_job_msg(cmsg_type, ckill_msg)
	packSJob_step_kill_msg_t(kill_msg, ckill_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_signal_job function as declared in slurm/slurm.h:3508
func slurm_signal_job(job_id uint32_t, signal uint16_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	csignal, _ := (C.uint16_t)(signal), cgoAllocsUnknown
	__ret := C.slurm_signal_job(cjob_id, csignal)
	__v := (int32)(__ret)
	return __v
}

// slurm_signal_job_step function as declared in slurm/slurm.h:3518
func slurm_signal_job_step(job_id uint32_t, step_id uint32_t, signal uint32_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	csignal, _ := (C.uint32_t)(signal), cgoAllocsUnknown
	__ret := C.slurm_signal_job_step(cjob_id, cstep_id, csignal)
	__v := (int32)(__ret)
	return __v
}

// slurm_complete_job function as declared in slurm/slurm.h:3533
func slurm_complete_job(job_id uint32_t, job_return_code uint32_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cjob_return_code, _ := (C.uint32_t)(job_return_code), cgoAllocsUnknown
	__ret := C.slurm_complete_job(cjob_id, cjob_return_code)
	__v := (int32)(__ret)
	return __v
}

// slurm_terminate_job_step function as declared in slurm/slurm.h:3546
func slurm_terminate_job_step(job_id uint32_t, step_id uint32_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	__ret := C.slurm_terminate_job_step(cjob_id, cstep_id)
	__v := (int32)(__ret)
	return __v
}

// slurm_step_ctx_params_t_init function as declared in slurm/slurm.h:3559
func slurm_step_ctx_params_t_init(ptr []slurm_step_ctx_params_t) {
	cptr, _ := unpackArgSSlurm_step_ctx_params_t(ptr)
	C.slurm_step_ctx_params_t_init(cptr)
	packSSlurm_step_ctx_params_t(ptr, cptr)
}

// slurm_step_ctx_create function as declared in slurm/slurm.h:3567
func slurm_step_ctx_create(step_params []slurm_step_ctx_params_t) *slurm_step_ctx_t {
	cstep_params, _ := unpackArgSSlurm_step_ctx_params_t(step_params)
	__ret := C.slurm_step_ctx_create(cstep_params)
	packSSlurm_step_ctx_params_t(step_params, cstep_params)
	__v := *(**slurm_step_ctx_t)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_step_ctx_create_timeout function as declared in slurm/slurm.h:3576
func slurm_step_ctx_create_timeout(step_params []slurm_step_ctx_params_t, timeout int32) *slurm_step_ctx_t {
	cstep_params, _ := unpackArgSSlurm_step_ctx_params_t(step_params)
	ctimeout, _ := (C.int)(timeout), cgoAllocsUnknown
	__ret := C.slurm_step_ctx_create_timeout(cstep_params, ctimeout)
	packSSlurm_step_ctx_params_t(step_params, cstep_params)
	__v := *(**slurm_step_ctx_t)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_step_retry_errno function as declared in slurm/slurm.h:3583
func slurm_step_retry_errno(rc int32) bool {
	crc, _ := (C.int)(rc), cgoAllocsUnknown
	__ret := C.slurm_step_retry_errno(crc)
	__v := (bool)(__ret)
	return __v
}

// slurm_step_ctx_create_no_alloc function as declared in slurm/slurm.h:3593
func slurm_step_ctx_create_no_alloc(step_params []slurm_step_ctx_params_t, step_id uint32_t) *slurm_step_ctx_t {
	cstep_params, _ := unpackArgSSlurm_step_ctx_params_t(step_params)
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	__ret := C.slurm_step_ctx_create_no_alloc(cstep_params, cstep_id)
	packSSlurm_step_ctx_params_t(step_params, cstep_params)
	__v := *(**slurm_step_ctx_t)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_jobinfo_ctx_get function as declared in slurm/slurm.h:3610
func slurm_jobinfo_ctx_get(jobinfo []dynamic_plugin_data_t, data_type int32, data unsafe.Pointer) int32 {
	cjobinfo, _ := unpackArgSDynamic_plugin_data_t(jobinfo)
	cdata_type, _ := (C.int)(data_type), cgoAllocsUnknown
	cdata, _ := data, cgoAllocsUnknown
	__ret := C.slurm_jobinfo_ctx_get(cjobinfo, cdata_type, cdata)
	packSDynamic_plugin_data_t(jobinfo, cjobinfo)
	__v := (int32)(__ret)
	return __v
}

// slurm_step_ctx_daemon_per_node_hack function as declared in slurm/slurm.h:3629
func slurm_step_ctx_daemon_per_node_hack(ctx []slurm_step_ctx_t, node_list []byte, node_cnt uint32_t, curr_task_num []uint32_t) int32 {
	cctx, _ := (*C.slurm_step_ctx_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ctx)).Data)), cgoAllocsUnknown
	cnode_list, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&node_list)).Data)), cgoAllocsUnknown
	cnode_cnt, _ := (C.uint32_t)(node_cnt), cgoAllocsUnknown
	ccurr_task_num, _ := (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&curr_task_num)).Data)), cgoAllocsUnknown
	__ret := C.slurm_step_ctx_daemon_per_node_hack(cctx, cnode_list, cnode_cnt, ccurr_task_num)
	__v := (int32)(__ret)
	return __v
}

// slurm_step_ctx_destroy function as declared in slurm/slurm.h:3639
func slurm_step_ctx_destroy(ctx []slurm_step_ctx_t) int32 {
	cctx, _ := (*C.slurm_step_ctx_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ctx)).Data)), cgoAllocsUnknown
	__ret := C.slurm_step_ctx_destroy(cctx)
	__v := (int32)(__ret)
	return __v
}

// slurm_step_launch_params_t_init function as declared in slurm/slurm.h:3648
func slurm_step_launch_params_t_init(ptr []slurm_step_launch_params_t) {
	cptr, _ := unpackArgSSlurm_step_launch_params_t(ptr)
	C.slurm_step_launch_params_t_init(cptr)
	packSSlurm_step_launch_params_t(ptr, cptr)
}

// slurm_step_launch function as declared in slurm/slurm.h:3658
func slurm_step_launch(ctx []slurm_step_ctx_t, params []slurm_step_launch_params_t, callbacks []slurm_step_launch_callbacks_t, pack_job_cnt int32) int32 {
	cctx, _ := (*C.slurm_step_ctx_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ctx)).Data)), cgoAllocsUnknown
	cparams, _ := unpackArgSSlurm_step_launch_params_t(params)
	ccallbacks, _ := (*C.slurm_step_launch_callbacks_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&callbacks)).Data)), cgoAllocsUnknown
	cpack_job_cnt, _ := (C.int)(pack_job_cnt), cgoAllocsUnknown
	__ret := C.slurm_step_launch(cctx, cparams, ccallbacks, cpack_job_cnt)
	packSSlurm_step_launch_params_t(params, cparams)
	__v := (int32)(__ret)
	return __v
}

// slurm_step_launch_add function as declared in slurm/slurm.h:3674
func slurm_step_launch_add(ctx []slurm_step_ctx_t, first_ctx []slurm_step_ctx_t, params []slurm_step_launch_params_t, node_list []byte, start_nodeid int32) int32 {
	cctx, _ := (*C.slurm_step_ctx_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ctx)).Data)), cgoAllocsUnknown
	cfirst_ctx, _ := (*C.slurm_step_ctx_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&first_ctx)).Data)), cgoAllocsUnknown
	cparams, _ := unpackArgSSlurm_step_launch_params_t(params)
	cnode_list, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&node_list)).Data)), cgoAllocsUnknown
	cstart_nodeid, _ := (C.int)(start_nodeid), cgoAllocsUnknown
	__ret := C.slurm_step_launch_add(cctx, cfirst_ctx, cparams, cnode_list, cstart_nodeid)
	packSSlurm_step_launch_params_t(params, cparams)
	__v := (int32)(__ret)
	return __v
}

// slurm_step_launch_wait_start function as declared in slurm/slurm.h:3682
func slurm_step_launch_wait_start(ctx []slurm_step_ctx_t) int32 {
	cctx, _ := (*C.slurm_step_ctx_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ctx)).Data)), cgoAllocsUnknown
	__ret := C.slurm_step_launch_wait_start(cctx)
	__v := (int32)(__ret)
	return __v
}

// slurm_step_launch_wait_finish function as declared in slurm/slurm.h:3687
func slurm_step_launch_wait_finish(ctx []slurm_step_ctx_t) {
	cctx, _ := (*C.slurm_step_ctx_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ctx)).Data)), cgoAllocsUnknown
	C.slurm_step_launch_wait_finish(cctx)
}

// slurm_step_launch_abort function as declared in slurm/slurm.h:3694
func slurm_step_launch_abort(ctx []slurm_step_ctx_t) {
	cctx, _ := (*C.slurm_step_ctx_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ctx)).Data)), cgoAllocsUnknown
	C.slurm_step_launch_abort(cctx)
}

// slurm_step_launch_fwd_signal function as declared in slurm/slurm.h:3699
func slurm_step_launch_fwd_signal(ctx []slurm_step_ctx_t, signo int32) {
	cctx, _ := (*C.slurm_step_ctx_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&ctx)).Data)), cgoAllocsUnknown
	csigno, _ := (C.int)(signo), cgoAllocsUnknown
	C.slurm_step_launch_fwd_signal(cctx, csigno)
}

// slurm_mpi_plugin_init function as declared in slurm/slurm.h:3712
func slurm_mpi_plugin_init(plugin_name []byte) int32 {
	cplugin_name, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&plugin_name)).Data)), cgoAllocsUnknown
	__ret := C.slurm_mpi_plugin_init(cplugin_name)
	__v := (int32)(__ret)
	return __v
}

// slurm_api_version function as declared in slurm/slurm.h:3724
func slurm_api_version() int {
	__ret := C.slurm_api_version()
	__v := (int)(__ret)
	return __v
}

// slurm_load_ctl_conf function as declared in slurm/slurm.h:3735
func slurm_load_ctl_conf(update_time time_t, slurm_ctl_conf_ptr [][]slurm_ctl_conf_t) int32 {
	cupdate_time, _ := (C.time_t)(update_time), cgoAllocsUnknown
	cslurm_ctl_conf_ptr, _ := unpackArgSSSlurm_ctl_conf_t(slurm_ctl_conf_ptr)
	__ret := C.slurm_load_ctl_conf(cupdate_time, cslurm_ctl_conf_ptr)
	packSSSlurm_ctl_conf_t(slurm_ctl_conf_ptr, cslurm_ctl_conf_ptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_ctl_conf function as declared in slurm/slurm.h:3743
func slurm_free_ctl_conf(slurm_ctl_conf_ptr []slurm_ctl_conf_t) {
	cslurm_ctl_conf_ptr, _ := unpackArgSSlurm_ctl_conf_t(slurm_ctl_conf_ptr)
	C.slurm_free_ctl_conf(cslurm_ctl_conf_ptr)
	packSSlurm_ctl_conf_t(slurm_ctl_conf_ptr, cslurm_ctl_conf_ptr)
}

// slurm_print_ctl_conf function as declared in slurm/slurm.h:3751
func slurm_print_ctl_conf(out []FILE, slurm_ctl_conf_ptr []slurm_ctl_conf_t) {
	cout, _ := unpackArgSFILE(out)
	cslurm_ctl_conf_ptr, _ := unpackArgSSlurm_ctl_conf_t(slurm_ctl_conf_ptr)
	C.slurm_print_ctl_conf(cout, cslurm_ctl_conf_ptr)
	packSSlurm_ctl_conf_t(slurm_ctl_conf_ptr, cslurm_ctl_conf_ptr)
	packSFILE(out, cout)
}

// slurm_write_ctl_conf function as declared in slurm/slurm.h:3762
func slurm_write_ctl_conf(slurm_ctl_conf_ptr []slurm_ctl_conf_t, node_info_ptr []node_info_msg_t, part_info_ptr []partition_info_msg_t) {
	cslurm_ctl_conf_ptr, _ := unpackArgSSlurm_ctl_conf_t(slurm_ctl_conf_ptr)
	cnode_info_ptr, _ := unpackArgSNode_info_msg_t(node_info_ptr)
	cpart_info_ptr, _ := unpackArgSPartition_info_msg_t(part_info_ptr)
	C.slurm_write_ctl_conf(cslurm_ctl_conf_ptr, cnode_info_ptr, cpart_info_ptr)
	packSPartition_info_msg_t(part_info_ptr, cpart_info_ptr)
	packSNode_info_msg_t(node_info_ptr, cnode_info_ptr)
	packSSlurm_ctl_conf_t(slurm_ctl_conf_ptr, cslurm_ctl_conf_ptr)
}

// slurm_ctl_conf_2_key_pairs function as declared in slurm/slurm.h:3772
func slurm_ctl_conf_2_key_pairs(slurm_ctl_conf_ptr []slurm_ctl_conf_t) unsafe.Pointer {
	cslurm_ctl_conf_ptr, _ := unpackArgSSlurm_ctl_conf_t(slurm_ctl_conf_ptr)
	__ret := C.slurm_ctl_conf_2_key_pairs(cslurm_ctl_conf_ptr)
	packSSlurm_ctl_conf_t(slurm_ctl_conf_ptr, cslurm_ctl_conf_ptr)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_print_key_pairs function as declared in slurm/slurm.h:3781
func slurm_print_key_pairs(out []FILE, key_pairs unsafe.Pointer, title []byte) {
	cout, _ := unpackArgSFILE(out)
	ckey_pairs, _ := key_pairs, cgoAllocsUnknown
	ctitle, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&title)).Data)), cgoAllocsUnknown
	C.slurm_print_key_pairs(cout, ckey_pairs, ctitle)
	packSFILE(out, cout)
}

// slurm_load_slurmd_status function as declared in slurm/slurm.h:3790
func slurm_load_slurmd_status(slurmd_status_ptr [][]slurmd_status_t) int32 {
	cslurmd_status_ptr, _ := unpackArgSSSlurmd_status_t(slurmd_status_ptr)
	__ret := C.slurm_load_slurmd_status(cslurmd_status_ptr)
	packSSSlurmd_status_t(slurmd_status_ptr, cslurmd_status_ptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_slurmd_status function as declared in slurm/slurm.h:3797
func slurm_free_slurmd_status(slurmd_status_ptr []slurmd_status_t) {
	cslurmd_status_ptr, _ := unpackArgSSlurmd_status_t(slurmd_status_ptr)
	C.slurm_free_slurmd_status(cslurmd_status_ptr)
	packSSlurmd_status_t(slurmd_status_ptr, cslurmd_status_ptr)
}

// slurm_print_slurmd_status function as declared in slurm/slurm.h:3805
func slurm_print_slurmd_status(out []FILE, slurmd_status_ptr []slurmd_status_t) {
	cout, _ := unpackArgSFILE(out)
	cslurmd_status_ptr, _ := unpackArgSSlurmd_status_t(slurmd_status_ptr)
	C.slurm_print_slurmd_status(cout, cslurmd_status_ptr)
	packSSlurmd_status_t(slurmd_status_ptr, cslurmd_status_ptr)
	packSFILE(out, cout)
}

// slurm_init_update_step_msg function as declared in slurm/slurm.h:3812
func slurm_init_update_step_msg(step_msg []step_update_request_msg_t) {
	cstep_msg, _ := unpackArgSStep_update_request_msg_t(step_msg)
	C.slurm_init_update_step_msg(cstep_msg)
	packSStep_update_request_msg_t(step_msg, cstep_msg)
}

// slurm_get_statistics function as declared in slurm/slurm.h:3815
func slurm_get_statistics(buf [][]stats_info_response_msg_t, req []stats_info_request_msg_t) int32 {
	cbuf, _ := unpackArgSSStats_info_response_msg_t(buf)
	creq, _ := unpackArgSStats_info_request_msg_t(req)
	__ret := C.slurm_get_statistics(cbuf, creq)
	packSStats_info_request_msg_t(req, creq)
	packSSStats_info_response_msg_t(buf, cbuf)
	__v := (int32)(__ret)
	return __v
}

// slurm_reset_statistics function as declared in slurm/slurm.h:3819
func slurm_reset_statistics(req []stats_info_request_msg_t) int32 {
	creq, _ := unpackArgSStats_info_request_msg_t(req)
	__ret := C.slurm_reset_statistics(creq)
	packSStats_info_request_msg_t(req, creq)
	__v := (int32)(__ret)
	return __v
}

// slurm_job_cpus_allocated_on_node_id function as declared in slurm/slurm.h:3833
func slurm_job_cpus_allocated_on_node_id(job_resrcs_ptr []job_resources_t, node_id int32) int32 {
	cjob_resrcs_ptr, _ := (*C.job_resources_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_resrcs_ptr)).Data)), cgoAllocsUnknown
	cnode_id, _ := (C.int)(node_id), cgoAllocsUnknown
	__ret := C.slurm_job_cpus_allocated_on_node_id(cjob_resrcs_ptr, cnode_id)
	__v := (int32)(__ret)
	return __v
}

// slurm_job_cpus_allocated_on_node function as declared in slurm/slurm.h:3844
func slurm_job_cpus_allocated_on_node(job_resrcs_ptr []job_resources_t, node_name string) int32 {
	cjob_resrcs_ptr, _ := (*C.job_resources_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_resrcs_ptr)).Data)), cgoAllocsUnknown
	cnode_name, _ := unpackPCharString(node_name)
	__ret := C.slurm_job_cpus_allocated_on_node(cjob_resrcs_ptr, cnode_name)
	__v := (int32)(__ret)
	return __v
}

// slurm_job_cpus_allocated_str_on_node_id function as declared in slurm/slurm.h:3857
func slurm_job_cpus_allocated_str_on_node_id(cpus []byte, cpus_len size_t, job_resrcs_ptr []job_resources_t, node_id int32) int32 {
	ccpus, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&cpus)).Data)), cgoAllocsUnknown
	ccpus_len, _ := (C.size_t)(cpus_len), cgoAllocsUnknown
	cjob_resrcs_ptr, _ := (*C.job_resources_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_resrcs_ptr)).Data)), cgoAllocsUnknown
	cnode_id, _ := (C.int)(node_id), cgoAllocsUnknown
	__ret := C.slurm_job_cpus_allocated_str_on_node_id(ccpus, ccpus_len, cjob_resrcs_ptr, cnode_id)
	__v := (int32)(__ret)
	return __v
}

// slurm_job_cpus_allocated_str_on_node function as declared in slurm/slurm.h:3872
func slurm_job_cpus_allocated_str_on_node(cpus []byte, cpus_len size_t, job_resrcs_ptr []job_resources_t, node_name string) int32 {
	ccpus, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&cpus)).Data)), cgoAllocsUnknown
	ccpus_len, _ := (C.size_t)(cpus_len), cgoAllocsUnknown
	cjob_resrcs_ptr, _ := (*C.job_resources_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_resrcs_ptr)).Data)), cgoAllocsUnknown
	cnode_name, _ := unpackPCharString(node_name)
	__ret := C.slurm_job_cpus_allocated_str_on_node(ccpus, ccpus_len, cjob_resrcs_ptr, cnode_name)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_job_info_msg function as declared in slurm/slurm.h:3886
func slurm_free_job_info_msg(job_buffer_ptr []job_info_msg_t) {
	cjob_buffer_ptr, _ := unpackArgSJob_info_msg_t(job_buffer_ptr)
	C.slurm_free_job_info_msg(cjob_buffer_ptr)
	packSJob_info_msg_t(job_buffer_ptr, cjob_buffer_ptr)
}

// slurm_free_priority_factors_response_msg function as declared in slurm/slurm.h:3894
func slurm_free_priority_factors_response_msg(factors_resp []priority_factors_response_msg_t) {
	cfactors_resp, _ := unpackArgSPriority_factors_response_msg_t(factors_resp)
	C.slurm_free_priority_factors_response_msg(cfactors_resp)
	packSPriority_factors_response_msg_t(factors_resp, cfactors_resp)
}

// slurm_get_end_time function as declared in slurm/slurm.h:3903
func slurm_get_end_time(jobid uint32_t, end_time_ptr []time_t) int32 {
	cjobid, _ := (C.uint32_t)(jobid), cgoAllocsUnknown
	cend_time_ptr, _ := (*C.time_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&end_time_ptr)).Data)), cgoAllocsUnknown
	__ret := C.slurm_get_end_time(cjobid, cend_time_ptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_get_job_stderr function as declared in slurm/slurm.h:3906
func slurm_get_job_stderr(buf []byte, buf_size int32, job_ptr []job_info_t) {
	cbuf, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buf)).Data)), cgoAllocsUnknown
	cbuf_size, _ := (C.int)(buf_size), cgoAllocsUnknown
	cjob_ptr, _ := unpackArgSJob_info_t(job_ptr)
	C.slurm_get_job_stderr(cbuf, cbuf_size, cjob_ptr)
	packSJob_info_t(job_ptr, cjob_ptr)
}

// slurm_get_job_stdin function as declared in slurm/slurm.h:3909
func slurm_get_job_stdin(buf []byte, buf_size int32, job_ptr []job_info_t) {
	cbuf, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buf)).Data)), cgoAllocsUnknown
	cbuf_size, _ := (C.int)(buf_size), cgoAllocsUnknown
	cjob_ptr, _ := unpackArgSJob_info_t(job_ptr)
	C.slurm_get_job_stdin(cbuf, cbuf_size, cjob_ptr)
	packSJob_info_t(job_ptr, cjob_ptr)
}

// slurm_get_job_stdout function as declared in slurm/slurm.h:3912
func slurm_get_job_stdout(buf []byte, buf_size int32, job_ptr []job_info_t) {
	cbuf, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&buf)).Data)), cgoAllocsUnknown
	cbuf_size, _ := (C.int)(buf_size), cgoAllocsUnknown
	cjob_ptr, _ := unpackArgSJob_info_t(job_ptr)
	C.slurm_get_job_stdout(cbuf, cbuf_size, cjob_ptr)
	packSJob_info_t(job_ptr, cjob_ptr)
}

// slurm_get_rem_time function as declared in slurm/slurm.h:3919
func slurm_get_rem_time(jobid uint32_t) int {
	cjobid, _ := (C.uint32_t)(jobid), cgoAllocsUnknown
	__ret := C.slurm_get_rem_time(cjobid)
	__v := (int)(__ret)
	return __v
}

// slurm_job_node_ready function as declared in slurm/slurm.h:3926
func slurm_job_node_ready(job_id uint32_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	__ret := C.slurm_job_node_ready(cjob_id)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_job function as declared in slurm/slurm.h:3936
func slurm_load_job(resp [][]job_info_msg_t, job_id uint32_t, show_flags uint16_t) int32 {
	cresp, _ := unpackArgSSJob_info_msg_t(resp)
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	__ret := C.slurm_load_job(cresp, cjob_id, cshow_flags)
	packSSJob_info_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_job_prio function as declared in slurm/slurm.h:3951
func slurm_load_job_prio(factors_resp [][]priority_factors_response_msg_t, job_id_list List, partitions []byte, uid_list List, show_flags uint16_t) int32 {
	cfactors_resp, _ := unpackArgSSPriority_factors_response_msg_t(factors_resp)
	cjob_id_list, _ := *(*C.List)(unsafe.Pointer(&job_id_list)), cgoAllocsUnknown
	cpartitions, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&partitions)).Data)), cgoAllocsUnknown
	cuid_list, _ := *(*C.List)(unsafe.Pointer(&uid_list)), cgoAllocsUnknown
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	__ret := C.slurm_load_job_prio(cfactors_resp, cjob_id_list, cpartitions, cuid_list, cshow_flags)
	packSSPriority_factors_response_msg_t(factors_resp, cfactors_resp)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_job_user function as declared in slurm/slurm.h:3964
func slurm_load_job_user(job_info_msg_pptr [][]job_info_msg_t, user_id uint32_t, show_flags uint16_t) int32 {
	cjob_info_msg_pptr, _ := unpackArgSSJob_info_msg_t(job_info_msg_pptr)
	cuser_id, _ := (C.uint32_t)(user_id), cgoAllocsUnknown
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	__ret := C.slurm_load_job_user(cjob_info_msg_pptr, cuser_id, cshow_flags)
	packSSJob_info_msg_t(job_info_msg_pptr, cjob_info_msg_pptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_jobs function as declared in slurm/slurm.h:3977
func slurm_load_jobs(update_time time_t, job_info_msg_pptr [][]job_info_msg_t, show_flags uint16_t) int32 {
	cupdate_time, _ := (C.time_t)(update_time), cgoAllocsUnknown
	cjob_info_msg_pptr, _ := unpackArgSSJob_info_msg_t(job_info_msg_pptr)
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	__ret := C.slurm_load_jobs(cupdate_time, cjob_info_msg_pptr, cshow_flags)
	packSSJob_info_msg_t(job_info_msg_pptr, cjob_info_msg_pptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_notify_job function as declared in slurm/slurm.h:3988
func slurm_notify_job(job_id uint32_t, message []byte) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cmessage, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&message)).Data)), cgoAllocsUnknown
	__ret := C.slurm_notify_job(cjob_id, cmessage)
	__v := (int32)(__ret)
	return __v
}

// slurm_pid2jobid function as declared in slurm/slurm.h:3997
func slurm_pid2jobid(job_pid pid_t, job_id_ptr []uint32_t) int32 {
	cjob_pid, _ := (C.pid_t)(job_pid), cgoAllocsUnknown
	cjob_id_ptr, _ := (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_id_ptr)).Data)), cgoAllocsUnknown
	__ret := C.slurm_pid2jobid(cjob_pid, cjob_id_ptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_print_job_info function as declared in slurm/slurm.h:4006
func slurm_print_job_info(out []FILE, job_ptr []slurm_job_info_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cjob_ptr, _ := unpackArgSSlurm_job_info_t(job_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_job_info(cout, cjob_ptr, cone_liner)
	packSSlurm_job_info_t(job_ptr, cjob_ptr)
	packSFILE(out, cout)
}

// slurm_print_job_info_msg function as declared in slurm/slurm.h:4017
func slurm_print_job_info_msg(out []FILE, job_info_msg_ptr []job_info_msg_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cjob_info_msg_ptr, _ := unpackArgSJob_info_msg_t(job_info_msg_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_job_info_msg(cout, cjob_info_msg_ptr, cone_liner)
	packSJob_info_msg_t(job_info_msg_ptr, cjob_info_msg_ptr)
	packSFILE(out, cout)
}

// slurm_sprint_job_info function as declared in slurm/slurm.h:4029
func slurm_sprint_job_info(job_ptr []slurm_job_info_t, one_liner int32) *byte {
	cjob_ptr, _ := unpackArgSSlurm_job_info_t(job_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	__ret := C.slurm_sprint_job_info(cjob_ptr, cone_liner)
	packSSlurm_job_info_t(job_ptr, cjob_ptr)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_update_job function as declared in slurm/slurm.h:4038
func slurm_update_job(job_msg []job_desc_msg_t) int32 {
	cjob_msg, _ := unpackArgSJob_desc_msg_t(job_msg)
	__ret := C.slurm_update_job(cjob_msg)
	packSJob_desc_msg_t(job_msg, cjob_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_update_job2 function as declared in slurm/slurm.h:4048
func slurm_update_job2(job_msg []job_desc_msg_t, resp [][]job_array_resp_msg_t) int32 {
	cjob_msg, _ := unpackArgSJob_desc_msg_t(job_msg)
	cresp, _ := unpackArgSSJob_array_resp_msg_t(resp)
	__ret := C.slurm_update_job2(cjob_msg, cresp)
	packSSJob_array_resp_msg_t(resp, cresp)
	packSJob_desc_msg_t(job_msg, cjob_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_xlate_job_id function as declared in slurm/slurm.h:4059
func slurm_xlate_job_id(job_id_str []byte) uint32_t {
	cjob_id_str, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_id_str)).Data)), cgoAllocsUnknown
	__ret := C.slurm_xlate_job_id(cjob_id_str)
	__v := (uint32_t)(__ret)
	return __v
}

// slurm_get_job_steps function as declared in slurm/slurm.h:4080
func slurm_get_job_steps(update_time time_t, job_id uint32_t, step_id uint32_t, step_response_pptr [][]job_step_info_response_msg_t, show_flags uint16_t) int32 {
	cupdate_time, _ := (C.time_t)(update_time), cgoAllocsUnknown
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	cstep_response_pptr, _ := unpackArgSSJob_step_info_response_msg_t(step_response_pptr)
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	__ret := C.slurm_get_job_steps(cupdate_time, cjob_id, cstep_id, cstep_response_pptr, cshow_flags)
	packSSJob_step_info_response_msg_t(step_response_pptr, cstep_response_pptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_job_step_info_response_msg function as declared in slurm/slurm.h:4092
func slurm_free_job_step_info_response_msg(msg []job_step_info_response_msg_t) {
	cmsg, _ := unpackArgSJob_step_info_response_msg_t(msg)
	C.slurm_free_job_step_info_response_msg(cmsg)
	packSJob_step_info_response_msg_t(msg, cmsg)
}

// slurm_print_job_step_info_msg function as declared in slurm/slurm.h:4101
func slurm_print_job_step_info_msg(out []FILE, job_step_info_msg_ptr []job_step_info_response_msg_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cjob_step_info_msg_ptr, _ := unpackArgSJob_step_info_response_msg_t(job_step_info_msg_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_job_step_info_msg(cout, cjob_step_info_msg_ptr, cone_liner)
	packSJob_step_info_response_msg_t(job_step_info_msg_ptr, cjob_step_info_msg_ptr)
	packSFILE(out, cout)
}

// slurm_print_job_step_info function as declared in slurm/slurm.h:4112
func slurm_print_job_step_info(out []FILE, step_ptr []job_step_info_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cstep_ptr, _ := unpackArgSJob_step_info_t(step_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_job_step_info(cout, cstep_ptr, cone_liner)
	packSJob_step_info_t(step_ptr, cstep_ptr)
	packSFILE(out, cout)
}

// slurm_job_step_layout_get function as declared in slurm/slurm.h:4125
func slurm_job_step_layout_get(job_id uint32_t, step_id uint32_t) *slurm_step_layout_t {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	__ret := C.slurm_job_step_layout_get(cjob_id, cstep_id)
	__v := Newslurm_step_layout_tRef(unsafe.Pointer(__ret))
	return __v
}

// slurm_sprint_job_step_info function as declared in slurm/slurm.h:4136
func slurm_sprint_job_step_info(step_ptr []job_step_info_t, one_liner int32) *byte {
	cstep_ptr, _ := unpackArgSJob_step_info_t(step_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	__ret := C.slurm_sprint_job_step_info(cstep_ptr, cone_liner)
	packSJob_step_info_t(step_ptr, cstep_ptr)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_job_step_stat function as declared in slurm/slurm.h:4148
func slurm_job_step_stat(job_id uint32_t, step_id uint32_t, node_list []byte, use_protocol_ver uint16_t, resp [][]job_step_stat_response_msg_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	cnode_list, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&node_list)).Data)), cgoAllocsUnknown
	cuse_protocol_ver, _ := (C.uint16_t)(use_protocol_ver), cgoAllocsUnknown
	cresp, _ := unpackArgSSJob_step_stat_response_msg_t(resp)
	__ret := C.slurm_job_step_stat(cjob_id, cstep_id, cnode_list, cuse_protocol_ver, cresp)
	packSSJob_step_stat_response_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_job_step_get_pids function as declared in slurm/slurm.h:4163
func slurm_job_step_get_pids(job_id uint32_t, step_id uint32_t, node_list []byte, resp [][]job_step_pids_response_msg_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	cnode_list, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&node_list)).Data)), cgoAllocsUnknown
	cresp, _ := unpackArgSSJob_step_pids_response_msg_t(resp)
	__ret := C.slurm_job_step_get_pids(cjob_id, cstep_id, cnode_list, cresp)
	packSSJob_step_pids_response_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_job_step_layout_free function as declared in slurm/slurm.h:4168
func slurm_job_step_layout_free(layout []slurm_step_layout_t) {
	clayout, _ := unpackArgSSlurm_step_layout_t(layout)
	C.slurm_job_step_layout_free(clayout)
	packSSlurm_step_layout_t(layout, clayout)
}

// slurm_job_step_pids_free function as declared in slurm/slurm.h:4169
func slurm_job_step_pids_free(object []job_step_pids_t) {
	cobject, _ := unpackArgSJob_step_pids_t(object)
	C.slurm_job_step_pids_free(cobject)
	packSJob_step_pids_t(object, cobject)
}

// slurm_job_step_pids_response_msg_free function as declared in slurm/slurm.h:4170
func slurm_job_step_pids_response_msg_free(object unsafe.Pointer) {
	cobject, _ := object, cgoAllocsUnknown
	C.slurm_job_step_pids_response_msg_free(cobject)
}

// slurm_job_step_stat_free function as declared in slurm/slurm.h:4171
func slurm_job_step_stat_free(object []job_step_stat_t) {
	cobject, _ := unpackArgSJob_step_stat_t(object)
	C.slurm_job_step_stat_free(cobject)
	packSJob_step_stat_t(object, cobject)
}

// slurm_job_step_stat_response_msg_free function as declared in slurm/slurm.h:4172
func slurm_job_step_stat_response_msg_free(object unsafe.Pointer) {
	cobject, _ := object, cgoAllocsUnknown
	C.slurm_job_step_stat_response_msg_free(cobject)
}

// slurm_update_step function as declared in slurm/slurm.h:4177
func slurm_update_step(step_msg []step_update_request_msg_t) int32 {
	cstep_msg, _ := unpackArgSStep_update_request_msg_t(step_msg)
	__ret := C.slurm_update_step(cstep_msg)
	packSStep_update_request_msg_t(step_msg, cstep_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_node function as declared in slurm/slurm.h:4192
func slurm_load_node(update_time time_t, resp [][]node_info_msg_t, show_flags uint16_t) int32 {
	cupdate_time, _ := (C.time_t)(update_time), cgoAllocsUnknown
	cresp, _ := unpackArgSSNode_info_msg_t(resp)
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	__ret := C.slurm_load_node(cupdate_time, cresp, cshow_flags)
	packSSNode_info_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_node2 function as declared in slurm/slurm.h:4199
func slurm_load_node2(update_time time_t, resp [][]node_info_msg_t, show_flags uint16_t, cluster []slurmdb_cluster_rec_t) int32 {
	cupdate_time, _ := (C.time_t)(update_time), cgoAllocsUnknown
	cresp, _ := unpackArgSSNode_info_msg_t(resp)
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	ccluster, _ := (*C.slurmdb_cluster_rec_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&cluster)).Data)), cgoAllocsUnknown
	__ret := C.slurm_load_node2(cupdate_time, cresp, cshow_flags, ccluster)
	packSSNode_info_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_node_single function as declared in slurm/slurm.h:4212
func slurm_load_node_single(resp [][]node_info_msg_t, node_name []byte, show_flags uint16_t) int32 {
	cresp, _ := unpackArgSSNode_info_msg_t(resp)
	cnode_name, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&node_name)).Data)), cgoAllocsUnknown
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	__ret := C.slurm_load_node_single(cresp, cnode_name, cshow_flags)
	packSSNode_info_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_node_single2 function as declared in slurm/slurm.h:4219
func slurm_load_node_single2(resp [][]node_info_msg_t, node_name []byte, show_flags uint16_t, cluster []slurmdb_cluster_rec_t) int32 {
	cresp, _ := unpackArgSSNode_info_msg_t(resp)
	cnode_name, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&node_name)).Data)), cgoAllocsUnknown
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	ccluster, _ := (*C.slurmdb_cluster_rec_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&cluster)).Data)), cgoAllocsUnknown
	__ret := C.slurm_load_node_single2(cresp, cnode_name, cshow_flags, ccluster)
	packSSNode_info_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_populate_node_partitions function as declared in slurm/slurm.h:4226
func slurm_populate_node_partitions(node_buffer_ptr []node_info_msg_t, part_buffer_ptr []partition_info_msg_t) {
	cnode_buffer_ptr, _ := unpackArgSNode_info_msg_t(node_buffer_ptr)
	cpart_buffer_ptr, _ := unpackArgSPartition_info_msg_t(part_buffer_ptr)
	C.slurm_populate_node_partitions(cnode_buffer_ptr, cpart_buffer_ptr)
	packSPartition_info_msg_t(part_buffer_ptr, cpart_buffer_ptr)
	packSNode_info_msg_t(node_buffer_ptr, cnode_buffer_ptr)
}

// slurm_get_node_energy function as declared in slurm/slurm.h:4240
func slurm_get_node_energy(host []byte, delta uint16_t, sensors_cnt []uint16_t, energy [][]acct_gather_energy_t) int32 {
	chost, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&host)).Data)), cgoAllocsUnknown
	cdelta, _ := (C.uint16_t)(delta), cgoAllocsUnknown
	csensors_cnt, _ := (*C.uint16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&sensors_cnt)).Data)), cgoAllocsUnknown
	cenergy, _ := unpackArgSSAcct_gather_energy_t(energy)
	__ret := C.slurm_get_node_energy(chost, cdelta, csensors_cnt, cenergy)
	packSSAcct_gather_energy_t(energy, cenergy)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_node_info_msg function as declared in slurm/slurm.h:4250
func slurm_free_node_info_msg(node_buffer_ptr []node_info_msg_t) {
	cnode_buffer_ptr, _ := unpackArgSNode_info_msg_t(node_buffer_ptr)
	C.slurm_free_node_info_msg(cnode_buffer_ptr)
	packSNode_info_msg_t(node_buffer_ptr, cnode_buffer_ptr)
}

// slurm_print_node_info_msg function as declared in slurm/slurm.h:4259
func slurm_print_node_info_msg(out []FILE, node_info_msg_ptr []node_info_msg_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cnode_info_msg_ptr, _ := unpackArgSNode_info_msg_t(node_info_msg_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_node_info_msg(cout, cnode_info_msg_ptr, cone_liner)
	packSNode_info_msg_t(node_info_msg_ptr, cnode_info_msg_ptr)
	packSFILE(out, cout)
}

// slurm_print_node_table function as declared in slurm/slurm.h:4271
func slurm_print_node_table(out []FILE, node_ptr []node_info_t, node_scaling int32, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cnode_ptr, _ := unpackArgSNode_info_t(node_ptr)
	cnode_scaling, _ := (C.int)(node_scaling), cgoAllocsUnknown
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_node_table(cout, cnode_ptr, cnode_scaling, cone_liner)
	packSNode_info_t(node_ptr, cnode_ptr)
	packSFILE(out, cout)
}

// slurm_sprint_node_table function as declared in slurm/slurm.h:4285
func slurm_sprint_node_table(node_ptr []node_info_t, node_scaling int32, one_liner int32) *byte {
	cnode_ptr, _ := unpackArgSNode_info_t(node_ptr)
	cnode_scaling, _ := (C.int)(node_scaling), cgoAllocsUnknown
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	__ret := C.slurm_sprint_node_table(cnode_ptr, cnode_scaling, cone_liner)
	packSNode_info_t(node_ptr, cnode_ptr)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_init_update_node_msg function as declared in slurm/slurm.h:4293
func slurm_init_update_node_msg(update_node_msg []update_node_msg_t) {
	cupdate_node_msg, _ := unpackArgSUpdate_node_msg_t(update_node_msg)
	C.slurm_init_update_node_msg(cupdate_node_msg)
	packSUpdate_node_msg_t(update_node_msg, cupdate_node_msg)
}

// slurm_update_node function as declared in slurm/slurm.h:4301
func slurm_update_node(node_msg []update_node_msg_t) int32 {
	cnode_msg, _ := unpackArgSUpdate_node_msg_t(node_msg)
	__ret := C.slurm_update_node(cnode_msg)
	packSUpdate_node_msg_t(node_msg, cnode_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_front_end function as declared in slurm/slurm.h:4316
func slurm_load_front_end(update_time time_t, resp [][]front_end_info_msg_t) int32 {
	cupdate_time, _ := (C.time_t)(update_time), cgoAllocsUnknown
	cresp, _ := unpackArgSSFront_end_info_msg_t(resp)
	__ret := C.slurm_load_front_end(cupdate_time, cresp)
	packSSFront_end_info_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_front_end_info_msg function as declared in slurm/slurm.h:4325
func slurm_free_front_end_info_msg(front_end_buffer_ptr []front_end_info_msg_t) {
	cfront_end_buffer_ptr, _ := unpackArgSFront_end_info_msg_t(front_end_buffer_ptr)
	C.slurm_free_front_end_info_msg(cfront_end_buffer_ptr)
	packSFront_end_info_msg_t(front_end_buffer_ptr, cfront_end_buffer_ptr)
}

// slurm_print_front_end_info_msg function as declared in slurm/slurm.h:4334
func slurm_print_front_end_info_msg(out []FILE, front_end_info_msg_ptr []front_end_info_msg_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cfront_end_info_msg_ptr, _ := unpackArgSFront_end_info_msg_t(front_end_info_msg_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_front_end_info_msg(cout, cfront_end_info_msg_ptr, cone_liner)
	packSFront_end_info_msg_t(front_end_info_msg_ptr, cfront_end_info_msg_ptr)
	packSFILE(out, cout)
}

// slurm_print_front_end_table function as declared in slurm/slurm.h:4344
func slurm_print_front_end_table(out []FILE, front_end_ptr []front_end_info_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cfront_end_ptr, _ := unpackArgSFront_end_info_t(front_end_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_front_end_table(cout, cfront_end_ptr, cone_liner)
	packSFront_end_info_t(front_end_ptr, cfront_end_ptr)
	packSFILE(out, cout)
}

// slurm_sprint_front_end_table function as declared in slurm/slurm.h:4356
func slurm_sprint_front_end_table(front_end_ptr []front_end_info_t, one_liner int32) *byte {
	cfront_end_ptr, _ := unpackArgSFront_end_info_t(front_end_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	__ret := C.slurm_sprint_front_end_table(cfront_end_ptr, cone_liner)
	packSFront_end_info_t(front_end_ptr, cfront_end_ptr)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_init_update_front_end_msg function as declared in slurm/slurm.h:4363
func slurm_init_update_front_end_msg(update_front_end_msg []update_front_end_msg_t) {
	cupdate_front_end_msg, _ := unpackArgSUpdate_front_end_msg_t(update_front_end_msg)
	C.slurm_init_update_front_end_msg(cupdate_front_end_msg)
	packSUpdate_front_end_msg_t(update_front_end_msg, cupdate_front_end_msg)
}

// slurm_update_front_end function as declared in slurm/slurm.h:4371
func slurm_update_front_end(front_end_msg []update_front_end_msg_t) int32 {
	cfront_end_msg, _ := unpackArgSUpdate_front_end_msg_t(front_end_msg)
	__ret := C.slurm_update_front_end(cfront_end_msg)
	packSUpdate_front_end_msg_t(front_end_msg, cfront_end_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_topo function as declared in slurm/slurm.h:4385
func slurm_load_topo(topo_info_msg_pptr [][]topo_info_response_msg_t) int32 {
	ctopo_info_msg_pptr, _ := unpackArgSSTopo_info_response_msg_t(topo_info_msg_pptr)
	__ret := C.slurm_load_topo(ctopo_info_msg_pptr)
	packSSTopo_info_response_msg_t(topo_info_msg_pptr, ctopo_info_msg_pptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_topo_info_msg function as declared in slurm/slurm.h:4393
func slurm_free_topo_info_msg(msg []topo_info_response_msg_t) {
	cmsg, _ := unpackArgSTopo_info_response_msg_t(msg)
	C.slurm_free_topo_info_msg(cmsg)
	packSTopo_info_response_msg_t(msg, cmsg)
}

// slurm_print_topo_info_msg function as declared in slurm/slurm.h:4403
func slurm_print_topo_info_msg(out []FILE, topo_info_msg_ptr []topo_info_response_msg_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	ctopo_info_msg_ptr, _ := unpackArgSTopo_info_response_msg_t(topo_info_msg_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_topo_info_msg(cout, ctopo_info_msg_ptr, cone_liner)
	packSTopo_info_response_msg_t(topo_info_msg_ptr, ctopo_info_msg_ptr)
	packSFILE(out, cout)
}

// slurm_print_topo_record function as declared in slurm/slurm.h:4416
func slurm_print_topo_record(out []FILE, topo_ptr []topo_info_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	ctopo_ptr, _ := unpackArgSTopo_info_t(topo_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_topo_record(cout, ctopo_ptr, cone_liner)
	packSTopo_info_t(topo_ptr, ctopo_ptr)
	packSFILE(out, cout)
}

// slurm_load_powercap function as declared in slurm/slurm.h:4430
func slurm_load_powercap(powercap_info_msg_pptr [][]powercap_info_msg_t) int32 {
	cpowercap_info_msg_pptr, _ := unpackArgSSPowercap_info_msg_t(powercap_info_msg_pptr)
	__ret := C.slurm_load_powercap(cpowercap_info_msg_pptr)
	packSSPowercap_info_msg_t(powercap_info_msg_pptr, cpowercap_info_msg_pptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_powercap_info_msg function as declared in slurm/slurm.h:4438
func slurm_free_powercap_info_msg(msg []powercap_info_msg_t) {
	cmsg, _ := unpackArgSPowercap_info_msg_t(msg)
	C.slurm_free_powercap_info_msg(cmsg)
	packSPowercap_info_msg_t(msg, cmsg)
}

// slurm_print_powercap_info_msg function as declared in slurm/slurm.h:4447
func slurm_print_powercap_info_msg(out []FILE, powercap_info_msg_ptr []powercap_info_msg_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cpowercap_info_msg_ptr, _ := unpackArgSPowercap_info_msg_t(powercap_info_msg_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_powercap_info_msg(cout, cpowercap_info_msg_ptr, cone_liner)
	packSPowercap_info_msg_t(powercap_info_msg_ptr, cpowercap_info_msg_ptr)
	packSFILE(out, cout)
}

// slurm_update_powercap function as declared in slurm/slurm.h:4456
func slurm_update_powercap(powercap_msg []update_powercap_msg_t) int32 {
	cpowercap_msg, _ := unpackArgSUpdate_powercap_msg_t(powercap_msg)
	__ret := C.slurm_update_powercap(cpowercap_msg)
	packSUpdate_powercap_msg_t(powercap_msg, cpowercap_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_get_select_jobinfo function as declared in slurm/slurm.h:4469
func slurm_get_select_jobinfo(jobinfo []dynamic_plugin_data_t, data_type select_jobdata_type, data unsafe.Pointer) int32 {
	cjobinfo, _ := unpackArgSDynamic_plugin_data_t(jobinfo)
	cdata_type, _ := (C.enum_select_jobdata_type)(data_type), cgoAllocsUnknown
	cdata, _ := data, cgoAllocsUnknown
	__ret := C.slurm_get_select_jobinfo(cjobinfo, cdata_type, cdata)
	packSDynamic_plugin_data_t(jobinfo, cjobinfo)
	__v := (int32)(__ret)
	return __v
}

// slurm_get_select_nodeinfo function as declared in slurm/slurm.h:4481
func slurm_get_select_nodeinfo(nodeinfo []dynamic_plugin_data_t, data_type select_nodedata_type, state node_states, data unsafe.Pointer) int32 {
	cnodeinfo, _ := unpackArgSDynamic_plugin_data_t(nodeinfo)
	cdata_type, _ := (C.enum_select_nodedata_type)(data_type), cgoAllocsUnknown
	cstate, _ := (C.enum_node_states)(state), cgoAllocsUnknown
	cdata, _ := data, cgoAllocsUnknown
	__ret := C.slurm_get_select_nodeinfo(cnodeinfo, cdata_type, cstate, cdata)
	packSDynamic_plugin_data_t(nodeinfo, cnodeinfo)
	__v := (int32)(__ret)
	return __v
}

// slurm_init_part_desc_msg function as declared in slurm/slurm.h:4495
func slurm_init_part_desc_msg(update_part_msg []update_part_msg_t) {
	cupdate_part_msg, _ := unpackArgSUpdate_part_msg_t(update_part_msg)
	C.slurm_init_part_desc_msg(cupdate_part_msg)
	packSUpdate_part_msg_t(update_part_msg, cupdate_part_msg)
}

// slurm_load_partitions function as declared in slurm/slurm.h:4507
func slurm_load_partitions(update_time time_t, part_buffer_ptr [][]partition_info_msg_t, show_flags uint16_t) int32 {
	cupdate_time, _ := (C.time_t)(update_time), cgoAllocsUnknown
	cpart_buffer_ptr, _ := unpackArgSSPartition_info_msg_t(part_buffer_ptr)
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	__ret := C.slurm_load_partitions(cupdate_time, cpart_buffer_ptr, cshow_flags)
	packSSPartition_info_msg_t(part_buffer_ptr, cpart_buffer_ptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_partitions2 function as declared in slurm/slurm.h:4515
func slurm_load_partitions2(update_time time_t, resp [][]partition_info_msg_t, show_flags uint16_t, cluster []slurmdb_cluster_rec_t) int32 {
	cupdate_time, _ := (C.time_t)(update_time), cgoAllocsUnknown
	cresp, _ := unpackArgSSPartition_info_msg_t(resp)
	cshow_flags, _ := (C.uint16_t)(show_flags), cgoAllocsUnknown
	ccluster, _ := (*C.slurmdb_cluster_rec_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&cluster)).Data)), cgoAllocsUnknown
	__ret := C.slurm_load_partitions2(cupdate_time, cresp, cshow_flags, ccluster)
	packSSPartition_info_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_partition_info_msg function as declared in slurm/slurm.h:4526
func slurm_free_partition_info_msg(part_info_ptr []partition_info_msg_t) {
	cpart_info_ptr, _ := unpackArgSPartition_info_msg_t(part_info_ptr)
	C.slurm_free_partition_info_msg(cpart_info_ptr)
	packSPartition_info_msg_t(part_info_ptr, cpart_info_ptr)
}

// slurm_print_partition_info_msg function as declared in slurm/slurm.h:4535
func slurm_print_partition_info_msg(out []FILE, part_info_ptr []partition_info_msg_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cpart_info_ptr, _ := unpackArgSPartition_info_msg_t(part_info_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_partition_info_msg(cout, cpart_info_ptr, cone_liner)
	packSPartition_info_msg_t(part_info_ptr, cpart_info_ptr)
	packSFILE(out, cout)
}

// slurm_print_partition_info function as declared in slurm/slurm.h:4544
func slurm_print_partition_info(out []FILE, part_ptr []partition_info_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cpart_ptr, _ := unpackArgSPartition_info_t(part_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_partition_info(cout, cpart_ptr, cone_liner)
	packSPartition_info_t(part_ptr, cpart_ptr)
	packSFILE(out, cout)
}

// slurm_sprint_partition_info function as declared in slurm/slurm.h:4556
func slurm_sprint_partition_info(part_ptr []partition_info_t, one_liner int32) *byte {
	cpart_ptr, _ := unpackArgSPartition_info_t(part_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	__ret := C.slurm_sprint_partition_info(cpart_ptr, cone_liner)
	packSPartition_info_t(part_ptr, cpart_ptr)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_create_partition function as declared in slurm/slurm.h:4564
func slurm_create_partition(part_msg []update_part_msg_t) int32 {
	cpart_msg, _ := unpackArgSUpdate_part_msg_t(part_msg)
	__ret := C.slurm_create_partition(cpart_msg)
	packSUpdate_part_msg_t(part_msg, cpart_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_update_partition function as declared in slurm/slurm.h:4572
func slurm_update_partition(part_msg []update_part_msg_t) int32 {
	cpart_msg, _ := unpackArgSUpdate_part_msg_t(part_msg)
	__ret := C.slurm_update_partition(cpart_msg)
	packSUpdate_part_msg_t(part_msg, cpart_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_delete_partition function as declared in slurm/slurm.h:4580
func slurm_delete_partition(part_msg []delete_part_msg_t) int32 {
	cpart_msg, _ := unpackArgSDelete_part_msg_t(part_msg)
	__ret := C.slurm_delete_partition(cpart_msg)
	packSDelete_part_msg_t(part_msg, cpart_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_print_layout_info function as declared in slurm/slurm.h:4585
func slurm_print_layout_info(out []FILE, layout_info_ptr []layout_info_msg_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	clayout_info_ptr, _ := unpackArgSLayout_info_msg_t(layout_info_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_layout_info(cout, clayout_info_ptr, cone_liner)
	packSLayout_info_msg_t(layout_info_ptr, clayout_info_ptr)
	packSFILE(out, cout)
}

// slurm_load_layout function as declared in slurm/slurm.h:4589
func slurm_load_layout(layout_type []byte, entities []byte, _type []byte, no_relation uint32_t, resp [][]layout_info_msg_t) int32 {
	clayout_type, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&layout_type)).Data)), cgoAllocsUnknown
	centities, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&entities)).Data)), cgoAllocsUnknown
	c_type, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&_type)).Data)), cgoAllocsUnknown
	cno_relation, _ := (C.uint32_t)(no_relation), cgoAllocsUnknown
	cresp, _ := unpackArgSSLayout_info_msg_t(resp)
	__ret := C.slurm_load_layout(clayout_type, centities, c_type, cno_relation, cresp)
	packSSLayout_info_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_update_layout function as declared in slurm/slurm.h:4595
func slurm_update_layout(layout_info_msg []update_layout_msg_t) int32 {
	clayout_info_msg, _ := unpackArgSUpdate_layout_msg_t(layout_info_msg)
	__ret := C.slurm_update_layout(clayout_info_msg)
	packSUpdate_layout_msg_t(layout_info_msg, clayout_info_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_layout_info_msg function as declared in slurm/slurm.h:4597
func slurm_free_layout_info_msg(layout_info_msg []layout_info_msg_t) {
	clayout_info_msg, _ := unpackArgSLayout_info_msg_t(layout_info_msg)
	C.slurm_free_layout_info_msg(clayout_info_msg)
	packSLayout_info_msg_t(layout_info_msg, clayout_info_msg)
}

// slurm_init_resv_desc_msg function as declared in slurm/slurm.h:4608
func slurm_init_resv_desc_msg(update_resv_msg []resv_desc_msg_t) {
	cupdate_resv_msg, _ := unpackArgSResv_desc_msg_t(update_resv_msg)
	C.slurm_init_resv_desc_msg(cupdate_resv_msg)
	packSResv_desc_msg_t(update_resv_msg, cupdate_resv_msg)
}

// slurm_create_reservation function as declared in slurm/slurm.h:4615
func slurm_create_reservation(resv_msg []resv_desc_msg_t) *byte {
	cresv_msg, _ := unpackArgSResv_desc_msg_t(resv_msg)
	__ret := C.slurm_create_reservation(cresv_msg)
	packSResv_desc_msg_t(resv_msg, cresv_msg)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_update_reservation function as declared in slurm/slurm.h:4623
func slurm_update_reservation(resv_msg []resv_desc_msg_t) int32 {
	cresv_msg, _ := unpackArgSResv_desc_msg_t(resv_msg)
	__ret := C.slurm_update_reservation(cresv_msg)
	packSResv_desc_msg_t(resv_msg, cresv_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_delete_reservation function as declared in slurm/slurm.h:4631
func slurm_delete_reservation(resv_msg []reservation_name_msg_t) int32 {
	cresv_msg, _ := unpackArgSReservation_name_msg_t(resv_msg)
	__ret := C.slurm_delete_reservation(cresv_msg)
	packSReservation_name_msg_t(resv_msg, cresv_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_reservations function as declared in slurm/slurm.h:4642
func slurm_load_reservations(update_time time_t, resp [][]reserve_info_msg_t) int32 {
	cupdate_time, _ := (C.time_t)(update_time), cgoAllocsUnknown
	cresp, _ := unpackArgSSReserve_info_msg_t(resp)
	__ret := C.slurm_load_reservations(cupdate_time, cresp)
	packSSReserve_info_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_print_reservation_info_msg function as declared in slurm/slurm.h:4652
func slurm_print_reservation_info_msg(out []FILE, resv_info_ptr []reserve_info_msg_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cresv_info_ptr, _ := unpackArgSReserve_info_msg_t(resv_info_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_reservation_info_msg(cout, cresv_info_ptr, cone_liner)
	packSReserve_info_msg_t(resv_info_ptr, cresv_info_ptr)
	packSFILE(out, cout)
}

// slurm_print_reservation_info function as declared in slurm/slurm.h:4663
func slurm_print_reservation_info(out []FILE, resv_ptr []reserve_info_t, one_liner int32) {
	cout, _ := unpackArgSFILE(out)
	cresv_ptr, _ := unpackArgSReserve_info_t(resv_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	C.slurm_print_reservation_info(cout, cresv_ptr, cone_liner)
	packSReserve_info_t(resv_ptr, cresv_ptr)
	packSFILE(out, cout)
}

// slurm_sprint_reservation_info function as declared in slurm/slurm.h:4675
func slurm_sprint_reservation_info(resv_ptr []reserve_info_t, one_liner int32) *byte {
	cresv_ptr, _ := unpackArgSReserve_info_t(resv_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	__ret := C.slurm_sprint_reservation_info(cresv_ptr, cone_liner)
	packSReserve_info_t(resv_ptr, cresv_ptr)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_free_reservation_info_msg function as declared in slurm/slurm.h:4683
func slurm_free_reservation_info_msg(resv_info_ptr []reserve_info_msg_t) {
	cresv_info_ptr, _ := unpackArgSReserve_info_msg_t(resv_info_ptr)
	C.slurm_free_reservation_info_msg(cresv_info_ptr)
	packSReserve_info_msg_t(resv_info_ptr, cresv_info_ptr)
}

// slurm_ping function as declared in slurm/slurm.h:4694
func slurm_ping(primary int32) int32 {
	cprimary, _ := (C.int)(primary), cgoAllocsUnknown
	__ret := C.slurm_ping(cprimary)
	__v := (int32)(__ret)
	return __v
}

// slurm_reconfigure function as declared in slurm/slurm.h:4701
func slurm_reconfigure() int32 {
	__ret := C.slurm_reconfigure()
	__v := (int32)(__ret)
	return __v
}

// slurm_shutdown function as declared in slurm/slurm.h:4712
func slurm_shutdown(options uint16_t) int32 {
	coptions, _ := (C.uint16_t)(options), cgoAllocsUnknown
	__ret := C.slurm_shutdown(coptions)
	__v := (int32)(__ret)
	return __v
}

// slurm_takeover function as declared in slurm/slurm.h:4720
func slurm_takeover() int32 {
	__ret := C.slurm_takeover()
	__v := (int32)(__ret)
	return __v
}

// slurm_set_debugflags function as declared in slurm/slurm.h:4728
func slurm_set_debugflags(debug_flags_plus uint64_t, debug_flags_minus uint64_t) int32 {
	cdebug_flags_plus, _ := (C.uint64_t)(debug_flags_plus), cgoAllocsUnknown
	cdebug_flags_minus, _ := (C.uint64_t)(debug_flags_minus), cgoAllocsUnknown
	__ret := C.slurm_set_debugflags(cdebug_flags_plus, cdebug_flags_minus)
	__v := (int32)(__ret)
	return __v
}

// slurm_set_debug_level function as declared in slurm/slurm.h:4736
func slurm_set_debug_level(debug_level uint32_t) int32 {
	cdebug_level, _ := (C.uint32_t)(debug_level), cgoAllocsUnknown
	__ret := C.slurm_set_debug_level(cdebug_level)
	__v := (int32)(__ret)
	return __v
}

// slurm_set_schedlog_level function as declared in slurm/slurm.h:4743
func slurm_set_schedlog_level(schedlog_level uint32_t) int32 {
	cschedlog_level, _ := (C.uint32_t)(schedlog_level), cgoAllocsUnknown
	__ret := C.slurm_set_schedlog_level(cschedlog_level)
	__v := (int32)(__ret)
	return __v
}

// slurm_set_fs_dampeningfactor function as declared in slurm/slurm.h:4750
func slurm_set_fs_dampeningfactor(factor uint16_t) int32 {
	cfactor, _ := (C.uint16_t)(factor), cgoAllocsUnknown
	__ret := C.slurm_set_fs_dampeningfactor(cfactor)
	__v := (int32)(__ret)
	return __v
}

// slurm_suspend function as declared in slurm/slurm.h:4761
func slurm_suspend(job_id uint32_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	__ret := C.slurm_suspend(cjob_id)
	__v := (int32)(__ret)
	return __v
}

// slurm_suspend2 function as declared in slurm/slurm.h:4771
func slurm_suspend2(job_id []byte, resp [][]job_array_resp_msg_t) int32 {
	cjob_id, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_id)).Data)), cgoAllocsUnknown
	cresp, _ := unpackArgSSJob_array_resp_msg_t(resp)
	__ret := C.slurm_suspend2(cjob_id, cresp)
	packSSJob_array_resp_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_resume function as declared in slurm/slurm.h:4778
func slurm_resume(job_id uint32_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	__ret := C.slurm_resume(cjob_id)
	__v := (int32)(__ret)
	return __v
}

// slurm_resume2 function as declared in slurm/slurm.h:4788
func slurm_resume2(job_id []byte, resp [][]job_array_resp_msg_t) int32 {
	cjob_id, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_id)).Data)), cgoAllocsUnknown
	cresp, _ := unpackArgSSJob_array_resp_msg_t(resp)
	__ret := C.slurm_resume2(cjob_id, cresp)
	packSSJob_array_resp_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_job_array_resp function as declared in slurm/slurm.h:4791
func slurm_free_job_array_resp(resp []job_array_resp_msg_t) {
	cresp, _ := unpackArgSJob_array_resp_msg_t(resp)
	C.slurm_free_job_array_resp(cresp)
	packSJob_array_resp_msg_t(resp, cresp)
}

// slurm_requeue function as declared in slurm/slurm.h:4807
func slurm_requeue(job_id uint32_t, state uint32_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstate, _ := (C.uint32_t)(state), cgoAllocsUnknown
	__ret := C.slurm_requeue(cjob_id, cstate)
	__v := (int32)(__ret)
	return __v
}

// slurm_requeue2 function as declared in slurm/slurm.h:4826
func slurm_requeue2(job_id []byte, state uint32_t, resp [][]job_array_resp_msg_t) int32 {
	cjob_id, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_id)).Data)), cgoAllocsUnknown
	cstate, _ := (C.uint32_t)(state), cgoAllocsUnknown
	cresp, _ := unpackArgSSJob_array_resp_msg_t(resp)
	__ret := C.slurm_requeue2(cjob_id, cstate, cresp)
	packSSJob_array_resp_msg_t(resp, cresp)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_able function as declared in slurm/slurm.h:4842
func slurm_checkpoint_able(job_id uint32_t, step_id uint32_t, start_time []time_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	cstart_time, _ := (*C.time_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&start_time)).Data)), cgoAllocsUnknown
	__ret := C.slurm_checkpoint_able(cjob_id, cstep_id, cstart_time)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_disable function as declared in slurm/slurm.h:4852
func slurm_checkpoint_disable(job_id uint32_t, step_id uint32_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	__ret := C.slurm_checkpoint_disable(cjob_id, cstep_id)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_enable function as declared in slurm/slurm.h:4861
func slurm_checkpoint_enable(job_id uint32_t, step_id uint32_t) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	__ret := C.slurm_checkpoint_enable(cjob_id, cstep_id)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_create function as declared in slurm/slurm.h:4872
func slurm_checkpoint_create(job_id uint32_t, step_id uint32_t, max_wait uint16_t, image_dir []byte) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	cmax_wait, _ := (C.uint16_t)(max_wait), cgoAllocsUnknown
	cimage_dir, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&image_dir)).Data)), cgoAllocsUnknown
	__ret := C.slurm_checkpoint_create(cjob_id, cstep_id, cmax_wait, cimage_dir)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_requeue function as declared in slurm/slurm.h:4885
func slurm_checkpoint_requeue(job_id uint32_t, max_wait uint16_t, image_dir []byte) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cmax_wait, _ := (C.uint16_t)(max_wait), cgoAllocsUnknown
	cimage_dir, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&image_dir)).Data)), cgoAllocsUnknown
	__ret := C.slurm_checkpoint_requeue(cjob_id, cmax_wait, cimage_dir)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_vacate function as declared in slurm/slurm.h:4898
func slurm_checkpoint_vacate(job_id uint32_t, step_id uint32_t, max_wait uint16_t, image_dir []byte) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	cmax_wait, _ := (C.uint16_t)(max_wait), cgoAllocsUnknown
	cimage_dir, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&image_dir)).Data)), cgoAllocsUnknown
	__ret := C.slurm_checkpoint_vacate(cjob_id, cstep_id, cmax_wait, cimage_dir)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_restart function as declared in slurm/slurm.h:4911
func slurm_checkpoint_restart(job_id uint32_t, step_id uint32_t, stick uint16_t, image_dir []byte) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	cstick, _ := (C.uint16_t)(stick), cgoAllocsUnknown
	cimage_dir, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&image_dir)).Data)), cgoAllocsUnknown
	__ret := C.slurm_checkpoint_restart(cjob_id, cstep_id, cstick, cimage_dir)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_complete function as declared in slurm/slurm.h:4926
func slurm_checkpoint_complete(job_id uint32_t, step_id uint32_t, begin_time time_t, error_code uint32_t, error_msg []byte) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	cbegin_time, _ := (C.time_t)(begin_time), cgoAllocsUnknown
	cerror_code, _ := (C.uint32_t)(error_code), cgoAllocsUnknown
	cerror_msg, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&error_msg)).Data)), cgoAllocsUnknown
	__ret := C.slurm_checkpoint_complete(cjob_id, cstep_id, cbegin_time, cerror_code, cerror_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_task_complete function as declared in slurm/slurm.h:4943
func slurm_checkpoint_task_complete(job_id uint32_t, step_id uint32_t, task_id uint32_t, begin_time time_t, error_code uint32_t, error_msg []byte) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	ctask_id, _ := (C.uint32_t)(task_id), cgoAllocsUnknown
	cbegin_time, _ := (C.time_t)(begin_time), cgoAllocsUnknown
	cerror_code, _ := (C.uint32_t)(error_code), cgoAllocsUnknown
	cerror_msg, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&error_msg)).Data)), cgoAllocsUnknown
	__ret := C.slurm_checkpoint_task_complete(cjob_id, cstep_id, ctask_id, cbegin_time, cerror_code, cerror_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_error function as declared in slurm/slurm.h:4963
func slurm_checkpoint_error(job_id uint32_t, step_id uint32_t, error_code []uint32_t, error_msg [][]byte) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint32_t)(step_id), cgoAllocsUnknown
	cerror_code, _ := (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&error_code)).Data)), cgoAllocsUnknown
	cerror_msg, _ := unpackArgSSByte(error_msg)
	__ret := C.slurm_checkpoint_error(cjob_id, cstep_id, cerror_code, cerror_msg)
	packSSByte(error_msg, cerror_msg)
	__v := (int32)(__ret)
	return __v
}

// slurm_checkpoint_tasks function as declared in slurm/slurm.h:4978
func slurm_checkpoint_tasks(job_id uint32_t, step_id uint16_t, begin_time time_t, image_dir []byte, max_wait uint16_t, nodelist []byte) int32 {
	cjob_id, _ := (C.uint32_t)(job_id), cgoAllocsUnknown
	cstep_id, _ := (C.uint16_t)(step_id), cgoAllocsUnknown
	cbegin_time, _ := (C.time_t)(begin_time), cgoAllocsUnknown
	cimage_dir, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&image_dir)).Data)), cgoAllocsUnknown
	cmax_wait, _ := (C.uint16_t)(max_wait), cgoAllocsUnknown
	cnodelist, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&nodelist)).Data)), cgoAllocsUnknown
	__ret := C.slurm_checkpoint_tasks(cjob_id, cstep_id, cbegin_time, cimage_dir, cmax_wait, cnodelist)
	__v := (int32)(__ret)
	return __v
}

// slurm_set_trigger function as declared in slurm/slurm.h:4993
func slurm_set_trigger(trigger_set []trigger_info_t) int32 {
	ctrigger_set, _ := unpackArgSTrigger_info_t(trigger_set)
	__ret := C.slurm_set_trigger(ctrigger_set)
	packSTrigger_info_t(trigger_set, ctrigger_set)
	__v := (int32)(__ret)
	return __v
}

// slurm_clear_trigger function as declared in slurm/slurm.h:4999
func slurm_clear_trigger(trigger_clear []trigger_info_t) int32 {
	ctrigger_clear, _ := unpackArgSTrigger_info_t(trigger_clear)
	__ret := C.slurm_clear_trigger(ctrigger_clear)
	packSTrigger_info_t(trigger_clear, ctrigger_clear)
	__v := (int32)(__ret)
	return __v
}

// slurm_get_triggers function as declared in slurm/slurm.h:5006
func slurm_get_triggers(trigger_get [][]trigger_info_msg_t) int32 {
	ctrigger_get, _ := unpackArgSSTrigger_info_msg_t(trigger_get)
	__ret := C.slurm_get_triggers(ctrigger_get)
	packSSTrigger_info_msg_t(trigger_get, ctrigger_get)
	__v := (int32)(__ret)
	return __v
}

// slurm_pull_trigger function as declared in slurm/slurm.h:5012
func slurm_pull_trigger(trigger_pull []trigger_info_t) int32 {
	ctrigger_pull, _ := unpackArgSTrigger_info_t(trigger_pull)
	__ret := C.slurm_pull_trigger(ctrigger_pull)
	packSTrigger_info_t(trigger_pull, ctrigger_pull)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_trigger_msg function as declared in slurm/slurm.h:5018
func slurm_free_trigger_msg(trigger_free []trigger_info_msg_t) {
	ctrigger_free, _ := unpackArgSTrigger_info_msg_t(trigger_free)
	C.slurm_free_trigger_msg(ctrigger_free)
	packSTrigger_info_msg_t(trigger_free, ctrigger_free)
}

// slurm_init_trigger_msg function as declared in slurm/slurm.h:5024
func slurm_init_trigger_msg(trigger_info_msg []trigger_info_t) {
	ctrigger_info_msg, _ := unpackArgSTrigger_info_t(trigger_info_msg)
	C.slurm_init_trigger_msg(ctrigger_info_msg)
	packSTrigger_info_t(trigger_info_msg, ctrigger_info_msg)
}

// slurm_load_burst_buffer_info function as declared in slurm/slurm.h:5137
func slurm_load_burst_buffer_info(burst_buffer_info_msg_pptr [][]burst_buffer_info_msg_t) int32 {
	cburst_buffer_info_msg_pptr, _ := unpackArgSSBurst_buffer_info_msg_t(burst_buffer_info_msg_pptr)
	__ret := C.slurm_load_burst_buffer_info(cburst_buffer_info_msg_pptr)
	packSSBurst_buffer_info_msg_t(burst_buffer_info_msg_pptr, cburst_buffer_info_msg_pptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_free_burst_buffer_info_msg function as declared in slurm/slurm.h:5145
func slurm_free_burst_buffer_info_msg(burst_buffer_info_msg []burst_buffer_info_msg_t) {
	cburst_buffer_info_msg, _ := unpackArgSBurst_buffer_info_msg_t(burst_buffer_info_msg)
	C.slurm_free_burst_buffer_info_msg(cburst_buffer_info_msg)
	packSBurst_buffer_info_msg_t(burst_buffer_info_msg, cburst_buffer_info_msg)
}

// slurm_print_burst_buffer_info_msg function as declared in slurm/slurm.h:5155
func slurm_print_burst_buffer_info_msg(out []FILE, info_ptr []burst_buffer_info_msg_t, one_liner int32, verbosity int32) {
	cout, _ := unpackArgSFILE(out)
	cinfo_ptr, _ := unpackArgSBurst_buffer_info_msg_t(info_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	cverbosity, _ := (C.int)(verbosity), cgoAllocsUnknown
	C.slurm_print_burst_buffer_info_msg(cout, cinfo_ptr, cone_liner, cverbosity)
	packSBurst_buffer_info_msg_t(info_ptr, cinfo_ptr)
	packSFILE(out, cout)
}

// slurm_print_burst_buffer_record function as declared in slurm/slurm.h:5171
func slurm_print_burst_buffer_record(out []FILE, burst_buffer_ptr []burst_buffer_info_t, one_liner int32, verbose int32) {
	cout, _ := unpackArgSFILE(out)
	cburst_buffer_ptr, _ := unpackArgSBurst_buffer_info_t(burst_buffer_ptr)
	cone_liner, _ := (C.int)(one_liner), cgoAllocsUnknown
	cverbose, _ := (C.int)(verbose), cgoAllocsUnknown
	C.slurm_print_burst_buffer_record(cout, cburst_buffer_ptr, cone_liner, cverbose)
	packSBurst_buffer_info_t(burst_buffer_ptr, cburst_buffer_ptr)
	packSFILE(out, cout)
}

// slurm_network_callerid function as declared in slurm/slurm.h:5186
func slurm_network_callerid(req network_callerid_msg_t, job_id []uint32_t, node_name []byte, node_name_size int32) int32 {
	creq, _ := req.PassValue()
	cjob_id, _ := (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_id)).Data)), cgoAllocsUnknown
	cnode_name, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&node_name)).Data)), cgoAllocsUnknown
	cnode_name_size, _ := (C.int)(node_name_size), cgoAllocsUnknown
	__ret := C.slurm_network_callerid(creq, cjob_id, cnode_name, cnode_name_size)
	__v := (int32)(__ret)
	return __v
}

// slurm_top_job function as declared in slurm/slurm.h:5196
func slurm_top_job(job_id_str []byte) int32 {
	cjob_id_str, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&job_id_str)).Data)), cgoAllocsUnknown
	__ret := C.slurm_top_job(cjob_id_str)
	__v := (int32)(__ret)
	return __v
}

// slurm_load_federation function as declared in slurm/slurm.h:5210
func slurm_load_federation(fed_pptr []unsafe.Pointer) int32 {
	cfed_pptr, _ := (*unsafe.Pointer)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&fed_pptr)).Data)), cgoAllocsUnknown
	__ret := C.slurm_load_federation(cfed_pptr)
	__v := (int32)(__ret)
	return __v
}

// slurm_print_federation function as declared in slurm/slurm.h:5216
func slurm_print_federation(fed unsafe.Pointer) {
	cfed, _ := fed, cgoAllocsUnknown
	C.slurm_print_federation(cfed)
}

// slurm_destroy_federation_rec function as declared in slurm/slurm.h:5222
func slurm_destroy_federation_rec(fed unsafe.Pointer) {
	cfed, _ := fed, cgoAllocsUnknown
	C.slurm_destroy_federation_rec(cfed)
}

// slurm_strerror function as declared in slurm/slurm_errno.h:290
func slurm_strerror(errnum int32) *byte {
	cerrnum, _ := (C.int)(errnum), cgoAllocsUnknown
	__ret := C.slurm_strerror(cerrnum)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// slurm_seterrno function as declared in slurm/slurm_errno.h:293
func slurm_seterrno(errnum int32) {
	cerrnum, _ := (C.int)(errnum), cgoAllocsUnknown
	C.slurm_seterrno(cerrnum)
}

// slurm_get_errno function as declared in slurm/slurm_errno.h:296
func slurm_get_errno() int32 {
	__ret := C.slurm_get_errno()
	__v := (int32)(__ret)
	return __v
}

// slurm_perror function as declared in slurm/slurm_errno.h:299
func slurm_perror(msg []byte) {
	cmsg, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&msg)).Data)), cgoAllocsUnknown
	C.slurm_perror(cmsg)
}
