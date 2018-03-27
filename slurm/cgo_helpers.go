// WARNING: This file has automatically been generated on Wed, 28 Mar 2018 00:34:26 CEST.
// By https://git.io/c-for-go. DO NOT EDIT.

package slurm

/*
#cgo pkg-config: slurm
#include "slurm/slurm.h"
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

// allocSlurm_addr_tMemory allocates memory for type C.slurm_addr_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_addr_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_addr_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_addr_tValue = unsafe.Sizeof([1]C.slurm_addr_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_addr_t) Ref() *C.slurm_addr_t {
	if x == nil {
		return nil
	}
	return x.ref4e93d89
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_addr_t) Free() {
	if x != nil && x.allocs4e93d89 != nil {
		x.allocs4e93d89.(*cgoAllocMap).Free()
		x.ref4e93d89 = nil
	}
}

// Newslurm_addr_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_addr_tRef(ref unsafe.Pointer) *slurm_addr_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_addr_t)
	obj.ref4e93d89 = (*C.slurm_addr_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_addr_t) PassRef() (*C.slurm_addr_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref4e93d89 != nil {
		return x.ref4e93d89, nil
	}
	mem4e93d89 := allocSlurm_addr_tMemory(1)
	ref4e93d89 := (*C.slurm_addr_t)(mem4e93d89)
	allocs4e93d89 := new(cgoAllocMap)
	allocs4e93d89.Add(mem4e93d89)

	var csin_family_allocs *cgoAllocMap
	ref4e93d89.sin_family, csin_family_allocs = (C.sa_family_t)(x.sin_family), cgoAllocsUnknown
	allocs4e93d89.Borrow(csin_family_allocs)

	var csin_port_allocs *cgoAllocMap
	ref4e93d89.sin_port, csin_port_allocs = (C.in_port_t)(x.sin_port), cgoAllocsUnknown
	allocs4e93d89.Borrow(csin_port_allocs)

	var csin_zero_allocs *cgoAllocMap
	ref4e93d89.sin_zero, csin_zero_allocs = *(*[8]C.uchar)(unsafe.Pointer(&x.sin_zero)), cgoAllocsUnknown
	allocs4e93d89.Borrow(csin_zero_allocs)

	x.ref4e93d89 = ref4e93d89
	x.allocs4e93d89 = allocs4e93d89
	return ref4e93d89, allocs4e93d89

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_addr_t) PassValue() (C.slurm_addr_t, *cgoAllocMap) {
	if x.ref4e93d89 != nil {
		return *x.ref4e93d89, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_addr_t) Deref() {
	if x.ref4e93d89 == nil {
		return
	}
	x.sin_family = (sa_family_t)(x.ref4e93d89.sin_family)
	x.sin_port = (in_port_t)(x.ref4e93d89.sin_port)
	x.sin_zero = *(*[8]byte)(unsafe.Pointer(&x.ref4e93d89.sin_zero))
}

// Ref returns a reference to C object as it is.
func (x *slurmdb_cluster_rec_t) Ref() *C.slurmdb_cluster_rec_t {
	if x == nil {
		return nil
	}
	return (*C.slurmdb_cluster_rec_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *slurmdb_cluster_rec_t) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// Newslurmdb_cluster_rec_tRef converts the C object reference into a raw struct reference without wrapping.
func Newslurmdb_cluster_rec_tRef(ref unsafe.Pointer) *slurmdb_cluster_rec_t {
	return (*slurmdb_cluster_rec_t)(ref)
}

// Newslurmdb_cluster_rec_t allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func Newslurmdb_cluster_rec_t() *slurmdb_cluster_rec_t {
	return (*slurmdb_cluster_rec_t)(allocSlurmdb_cluster_rec_tMemory(1))
}

// allocSlurmdb_cluster_rec_tMemory allocates memory for type C.slurmdb_cluster_rec_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurmdb_cluster_rec_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurmdb_cluster_rec_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurmdb_cluster_rec_tValue = unsafe.Sizeof([1]C.slurmdb_cluster_rec_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *slurmdb_cluster_rec_t) PassRef() *C.slurmdb_cluster_rec_t {
	if x == nil {
		x = (*slurmdb_cluster_rec_t)(allocSlurmdb_cluster_rec_tMemory(1))
	}
	return (*C.slurmdb_cluster_rec_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *slurm_cred_t) Ref() *C.slurm_cred_t {
	if x == nil {
		return nil
	}
	return (*C.slurm_cred_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *slurm_cred_t) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// Newslurm_cred_tRef converts the C object reference into a raw struct reference without wrapping.
func Newslurm_cred_tRef(ref unsafe.Pointer) *slurm_cred_t {
	return (*slurm_cred_t)(ref)
}

// Newslurm_cred_t allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func Newslurm_cred_t() *slurm_cred_t {
	return (*slurm_cred_t)(allocSlurm_cred_tMemory(1))
}

// allocSlurm_cred_tMemory allocates memory for type C.slurm_cred_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_cred_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_cred_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_cred_tValue = unsafe.Sizeof([1]C.slurm_cred_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *slurm_cred_t) PassRef() *C.slurm_cred_t {
	if x == nil {
		x = (*slurm_cred_t)(allocSlurm_cred_tMemory(1))
	}
	return (*C.slurm_cred_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *switch_jobinfo_t) Ref() *C.switch_jobinfo_t {
	if x == nil {
		return nil
	}
	return (*C.switch_jobinfo_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *switch_jobinfo_t) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// Newswitch_jobinfo_tRef converts the C object reference into a raw struct reference without wrapping.
func Newswitch_jobinfo_tRef(ref unsafe.Pointer) *switch_jobinfo_t {
	return (*switch_jobinfo_t)(ref)
}

// Newswitch_jobinfo_t allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func Newswitch_jobinfo_t() *switch_jobinfo_t {
	return (*switch_jobinfo_t)(allocSwitch_jobinfo_tMemory(1))
}

// allocSwitch_jobinfo_tMemory allocates memory for type C.switch_jobinfo_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSwitch_jobinfo_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSwitch_jobinfo_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSwitch_jobinfo_tValue = unsafe.Sizeof([1]C.switch_jobinfo_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *switch_jobinfo_t) PassRef() *C.switch_jobinfo_t {
	if x == nil {
		x = (*switch_jobinfo_t)(allocSwitch_jobinfo_tMemory(1))
	}
	return (*C.switch_jobinfo_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *job_resources_t) Ref() *C.job_resources_t {
	if x == nil {
		return nil
	}
	return (*C.job_resources_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *job_resources_t) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// Newjob_resources_tRef converts the C object reference into a raw struct reference without wrapping.
func Newjob_resources_tRef(ref unsafe.Pointer) *job_resources_t {
	return (*job_resources_t)(ref)
}

// Newjob_resources_t allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func Newjob_resources_t() *job_resources_t {
	return (*job_resources_t)(allocJob_resources_tMemory(1))
}

// allocJob_resources_tMemory allocates memory for type C.job_resources_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_resources_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_resources_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_resources_tValue = unsafe.Sizeof([1]C.job_resources_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *job_resources_t) PassRef() *C.job_resources_t {
	if x == nil {
		x = (*job_resources_t)(allocJob_resources_tMemory(1))
	}
	return (*C.job_resources_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *select_jobinfo_t) Ref() *C.select_jobinfo_t {
	if x == nil {
		return nil
	}
	return (*C.select_jobinfo_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *select_jobinfo_t) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// Newselect_jobinfo_tRef converts the C object reference into a raw struct reference without wrapping.
func Newselect_jobinfo_tRef(ref unsafe.Pointer) *select_jobinfo_t {
	return (*select_jobinfo_t)(ref)
}

// Newselect_jobinfo_t allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func Newselect_jobinfo_t() *select_jobinfo_t {
	return (*select_jobinfo_t)(allocSelect_jobinfo_tMemory(1))
}

// allocSelect_jobinfo_tMemory allocates memory for type C.select_jobinfo_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSelect_jobinfo_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSelect_jobinfo_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSelect_jobinfo_tValue = unsafe.Sizeof([1]C.select_jobinfo_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *select_jobinfo_t) PassRef() *C.select_jobinfo_t {
	if x == nil {
		x = (*select_jobinfo_t)(allocSelect_jobinfo_tMemory(1))
	}
	return (*C.select_jobinfo_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *select_nodeinfo_t) Ref() *C.select_nodeinfo_t {
	if x == nil {
		return nil
	}
	return (*C.select_nodeinfo_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *select_nodeinfo_t) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// Newselect_nodeinfo_tRef converts the C object reference into a raw struct reference without wrapping.
func Newselect_nodeinfo_tRef(ref unsafe.Pointer) *select_nodeinfo_t {
	return (*select_nodeinfo_t)(ref)
}

// Newselect_nodeinfo_t allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func Newselect_nodeinfo_t() *select_nodeinfo_t {
	return (*select_nodeinfo_t)(allocSelect_nodeinfo_tMemory(1))
}

// allocSelect_nodeinfo_tMemory allocates memory for type C.select_nodeinfo_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSelect_nodeinfo_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSelect_nodeinfo_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSelect_nodeinfo_tValue = unsafe.Sizeof([1]C.select_nodeinfo_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *select_nodeinfo_t) PassRef() *C.select_nodeinfo_t {
	if x == nil {
		x = (*select_nodeinfo_t)(allocSelect_nodeinfo_tMemory(1))
	}
	return (*C.select_nodeinfo_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *jobacctinfo_t) Ref() *C.jobacctinfo_t {
	if x == nil {
		return nil
	}
	return (*C.jobacctinfo_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *jobacctinfo_t) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// Newjobacctinfo_tRef converts the C object reference into a raw struct reference without wrapping.
func Newjobacctinfo_tRef(ref unsafe.Pointer) *jobacctinfo_t {
	return (*jobacctinfo_t)(ref)
}

// Newjobacctinfo_t allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func Newjobacctinfo_t() *jobacctinfo_t {
	return (*jobacctinfo_t)(allocJobacctinfo_tMemory(1))
}

// allocJobacctinfo_tMemory allocates memory for type C.jobacctinfo_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJobacctinfo_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJobacctinfo_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJobacctinfo_tValue = unsafe.Sizeof([1]C.jobacctinfo_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *jobacctinfo_t) PassRef() *C.jobacctinfo_t {
	if x == nil {
		x = (*jobacctinfo_t)(allocJobacctinfo_tMemory(1))
	}
	return (*C.jobacctinfo_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *allocation_msg_thread_t) Ref() *C.allocation_msg_thread_t {
	if x == nil {
		return nil
	}
	return (*C.allocation_msg_thread_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *allocation_msg_thread_t) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// Newallocation_msg_thread_tRef converts the C object reference into a raw struct reference without wrapping.
func Newallocation_msg_thread_tRef(ref unsafe.Pointer) *allocation_msg_thread_t {
	return (*allocation_msg_thread_t)(ref)
}

// Newallocation_msg_thread_t allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func Newallocation_msg_thread_t() *allocation_msg_thread_t {
	return (*allocation_msg_thread_t)(allocAllocation_msg_thread_tMemory(1))
}

// allocAllocation_msg_thread_tMemory allocates memory for type C.allocation_msg_thread_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAllocation_msg_thread_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAllocation_msg_thread_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAllocation_msg_thread_tValue = unsafe.Sizeof([1]C.allocation_msg_thread_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *allocation_msg_thread_t) PassRef() *C.allocation_msg_thread_t {
	if x == nil {
		x = (*allocation_msg_thread_t)(allocAllocation_msg_thread_tMemory(1))
	}
	return (*C.allocation_msg_thread_t)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *sbcast_cred_t) Ref() *C.sbcast_cred_t {
	if x == nil {
		return nil
	}
	return (*C.sbcast_cred_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *sbcast_cred_t) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// Newsbcast_cred_tRef converts the C object reference into a raw struct reference without wrapping.
func Newsbcast_cred_tRef(ref unsafe.Pointer) *sbcast_cred_t {
	return (*sbcast_cred_t)(ref)
}

// Newsbcast_cred_t allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func Newsbcast_cred_t() *sbcast_cred_t {
	return (*sbcast_cred_t)(allocSbcast_cred_tMemory(1))
}

// allocSbcast_cred_tMemory allocates memory for type C.sbcast_cred_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSbcast_cred_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSbcast_cred_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSbcast_cred_tValue = unsafe.Sizeof([1]C.sbcast_cred_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *sbcast_cred_t) PassRef() *C.sbcast_cred_t {
	if x == nil {
		x = (*sbcast_cred_t)(allocSbcast_cred_tMemory(1))
	}
	return (*C.sbcast_cred_t)(unsafe.Pointer(x))
}

func (x ListDelF) PassRef() (ref *C.ListDelF, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if listDelF88473E32Func == nil {
		listDelF88473E32Func = x
	}
	return (*C.ListDelF)(C.ListDelF_88473e32), nil
}

func (x ListDelF) PassValue() (ref C.ListDelF, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if listDelF88473E32Func == nil {
		listDelF88473E32Func = x
	}
	return (C.ListDelF)(C.ListDelF_88473e32), nil
}

func NewListDelFRef(ref unsafe.Pointer) *ListDelF {
	return (*ListDelF)(ref)
}

//export listDelF88473E32
func listDelF88473E32(cx unsafe.Pointer) {
	if listDelF88473E32Func != nil {
		x88473e32 := (unsafe.Pointer)(unsafe.Pointer(cx))
		listDelF88473E32Func(x88473e32)
		return
	}
	panic("callback func has not been set (race?)")
}

var listDelF88473E32Func ListDelF

func (x ListCmpF) PassRef() (ref *C.ListCmpF, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if listCmpFFDF40A6EFunc == nil {
		listCmpFFDF40A6EFunc = x
	}
	return (*C.ListCmpF)(C.ListCmpF_fdf40a6e), nil
}

func (x ListCmpF) PassValue() (ref C.ListCmpF, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if listCmpFFDF40A6EFunc == nil {
		listCmpFFDF40A6EFunc = x
	}
	return (C.ListCmpF)(C.ListCmpF_fdf40a6e), nil
}

func NewListCmpFRef(ref unsafe.Pointer) *ListCmpF {
	return (*ListCmpF)(ref)
}

//export listCmpFFDF40A6E
func listCmpFFDF40A6E(cx unsafe.Pointer, cy unsafe.Pointer) C.int {
	if listCmpFFDF40A6EFunc != nil {
		xfdf40a6e := (unsafe.Pointer)(unsafe.Pointer(cx))
		yfdf40a6e := (unsafe.Pointer)(unsafe.Pointer(cy))
		retfdf40a6e := listCmpFFDF40A6EFunc(xfdf40a6e, yfdf40a6e)
		ret, _ := (C.int)(retfdf40a6e), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var listCmpFFDF40A6EFunc ListCmpF

func (x ListFindF) PassRef() (ref *C.ListFindF, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if listFindF11C26300Func == nil {
		listFindF11C26300Func = x
	}
	return (*C.ListFindF)(C.ListFindF_11c26300), nil
}

func (x ListFindF) PassValue() (ref C.ListFindF, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if listFindF11C26300Func == nil {
		listFindF11C26300Func = x
	}
	return (C.ListFindF)(C.ListFindF_11c26300), nil
}

func NewListFindFRef(ref unsafe.Pointer) *ListFindF {
	return (*ListFindF)(ref)
}

//export listFindF11C26300
func listFindF11C26300(cx unsafe.Pointer, ckey unsafe.Pointer) C.int {
	if listFindF11C26300Func != nil {
		x11c26300 := (unsafe.Pointer)(unsafe.Pointer(cx))
		key11c26300 := (unsafe.Pointer)(unsafe.Pointer(ckey))
		ret11c26300 := listFindF11C26300Func(x11c26300, key11c26300)
		ret, _ := (C.int)(ret11c26300), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var listFindF11C26300Func ListFindF

func (x ListForF) PassRef() (ref *C.ListForF, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if listForFFB984CB0Func == nil {
		listForFFB984CB0Func = x
	}
	return (*C.ListForF)(C.ListForF_fb984cb0), nil
}

func (x ListForF) PassValue() (ref C.ListForF, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if listForFFB984CB0Func == nil {
		listForFFB984CB0Func = x
	}
	return (C.ListForF)(C.ListForF_fb984cb0), nil
}

func NewListForFRef(ref unsafe.Pointer) *ListForF {
	return (*ListForF)(ref)
}

//export listForFFB984CB0
func listForFFB984CB0(cx unsafe.Pointer, carg unsafe.Pointer) C.int {
	if listForFFB984CB0Func != nil {
		xfb984cb0 := (unsafe.Pointer)(unsafe.Pointer(cx))
		argfb984cb0 := (unsafe.Pointer)(unsafe.Pointer(carg))
		retfb984cb0 := listForFFB984CB0Func(xfb984cb0, argfb984cb0)
		ret, _ := (C.int)(retfb984cb0), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var listForFFB984CB0Func ListForF

// allocDynamic_plugin_data_tMemory allocates memory for type C.dynamic_plugin_data_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocDynamic_plugin_data_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfDynamic_plugin_data_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfDynamic_plugin_data_tValue = unsafe.Sizeof([1]C.dynamic_plugin_data_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *dynamic_plugin_data_t) Ref() *C.dynamic_plugin_data_t {
	if x == nil {
		return nil
	}
	return x.ref6c951b34
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *dynamic_plugin_data_t) Free() {
	if x != nil && x.allocs6c951b34 != nil {
		x.allocs6c951b34.(*cgoAllocMap).Free()
		x.ref6c951b34 = nil
	}
}

// Newdynamic_plugin_data_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newdynamic_plugin_data_tRef(ref unsafe.Pointer) *dynamic_plugin_data_t {
	if ref == nil {
		return nil
	}
	obj := new(dynamic_plugin_data_t)
	obj.ref6c951b34 = (*C.dynamic_plugin_data_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *dynamic_plugin_data_t) PassRef() (*C.dynamic_plugin_data_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6c951b34 != nil {
		return x.ref6c951b34, nil
	}
	mem6c951b34 := allocDynamic_plugin_data_tMemory(1)
	ref6c951b34 := (*C.dynamic_plugin_data_t)(mem6c951b34)
	allocs6c951b34 := new(cgoAllocMap)
	allocs6c951b34.Add(mem6c951b34)

	var cdata_allocs *cgoAllocMap
	ref6c951b34.data, cdata_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.data)), cgoAllocsUnknown
	allocs6c951b34.Borrow(cdata_allocs)

	var cplugin_id_allocs *cgoAllocMap
	ref6c951b34.plugin_id, cplugin_id_allocs = (C.uint32_t)(x.plugin_id), cgoAllocsUnknown
	allocs6c951b34.Borrow(cplugin_id_allocs)

	x.ref6c951b34 = ref6c951b34
	x.allocs6c951b34 = allocs6c951b34
	return ref6c951b34, allocs6c951b34

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x dynamic_plugin_data_t) PassValue() (C.dynamic_plugin_data_t, *cgoAllocMap) {
	if x.ref6c951b34 != nil {
		return *x.ref6c951b34, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *dynamic_plugin_data_t) Deref() {
	if x.ref6c951b34 == nil {
		return
	}
	x.data = (unsafe.Pointer)(unsafe.Pointer(x.ref6c951b34.data))
	x.plugin_id = (uint32_t)(x.ref6c951b34.plugin_id)
}

// allocAcct_gather_energy_tMemory allocates memory for type C.acct_gather_energy_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAcct_gather_energy_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAcct_gather_energy_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAcct_gather_energy_tValue = unsafe.Sizeof([1]C.acct_gather_energy_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *acct_gather_energy_t) Ref() *C.acct_gather_energy_t {
	if x == nil {
		return nil
	}
	return x.ref3c370a53
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *acct_gather_energy_t) Free() {
	if x != nil && x.allocs3c370a53 != nil {
		x.allocs3c370a53.(*cgoAllocMap).Free()
		x.ref3c370a53 = nil
	}
}

// Newacct_gather_energy_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newacct_gather_energy_tRef(ref unsafe.Pointer) *acct_gather_energy_t {
	if ref == nil {
		return nil
	}
	obj := new(acct_gather_energy_t)
	obj.ref3c370a53 = (*C.acct_gather_energy_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *acct_gather_energy_t) PassRef() (*C.acct_gather_energy_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3c370a53 != nil {
		return x.ref3c370a53, nil
	}
	mem3c370a53 := allocAcct_gather_energy_tMemory(1)
	ref3c370a53 := (*C.acct_gather_energy_t)(mem3c370a53)
	allocs3c370a53 := new(cgoAllocMap)
	allocs3c370a53.Add(mem3c370a53)

	var cbase_consumed_energy_allocs *cgoAllocMap
	ref3c370a53.base_consumed_energy, cbase_consumed_energy_allocs = (C.uint64_t)(x.base_consumed_energy), cgoAllocsUnknown
	allocs3c370a53.Borrow(cbase_consumed_energy_allocs)

	var cbase_watts_allocs *cgoAllocMap
	ref3c370a53.base_watts, cbase_watts_allocs = (C.uint32_t)(x.base_watts), cgoAllocsUnknown
	allocs3c370a53.Borrow(cbase_watts_allocs)

	var cconsumed_energy_allocs *cgoAllocMap
	ref3c370a53.consumed_energy, cconsumed_energy_allocs = (C.uint64_t)(x.consumed_energy), cgoAllocsUnknown
	allocs3c370a53.Borrow(cconsumed_energy_allocs)

	var ccurrent_watts_allocs *cgoAllocMap
	ref3c370a53.current_watts, ccurrent_watts_allocs = (C.uint32_t)(x.current_watts), cgoAllocsUnknown
	allocs3c370a53.Borrow(ccurrent_watts_allocs)

	var cprevious_consumed_energy_allocs *cgoAllocMap
	ref3c370a53.previous_consumed_energy, cprevious_consumed_energy_allocs = (C.uint64_t)(x.previous_consumed_energy), cgoAllocsUnknown
	allocs3c370a53.Borrow(cprevious_consumed_energy_allocs)

	var cpoll_time_allocs *cgoAllocMap
	ref3c370a53.poll_time, cpoll_time_allocs = (C.time_t)(x.poll_time), cgoAllocsUnknown
	allocs3c370a53.Borrow(cpoll_time_allocs)

	x.ref3c370a53 = ref3c370a53
	x.allocs3c370a53 = allocs3c370a53
	return ref3c370a53, allocs3c370a53

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x acct_gather_energy_t) PassValue() (C.acct_gather_energy_t, *cgoAllocMap) {
	if x.ref3c370a53 != nil {
		return *x.ref3c370a53, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *acct_gather_energy_t) Deref() {
	if x.ref3c370a53 == nil {
		return
	}
	x.base_consumed_energy = (uint64_t)(x.ref3c370a53.base_consumed_energy)
	x.base_watts = (uint32_t)(x.ref3c370a53.base_watts)
	x.consumed_energy = (uint64_t)(x.ref3c370a53.consumed_energy)
	x.current_watts = (uint32_t)(x.ref3c370a53.current_watts)
	x.previous_consumed_energy = (uint64_t)(x.ref3c370a53.previous_consumed_energy)
	x.poll_time = (time_t)(x.ref3c370a53.poll_time)
}

// allocExt_sensors_data_tMemory allocates memory for type C.ext_sensors_data_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocExt_sensors_data_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfExt_sensors_data_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfExt_sensors_data_tValue = unsafe.Sizeof([1]C.ext_sensors_data_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *ext_sensors_data_t) Ref() *C.ext_sensors_data_t {
	if x == nil {
		return nil
	}
	return x.refed768141
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *ext_sensors_data_t) Free() {
	if x != nil && x.allocsed768141 != nil {
		x.allocsed768141.(*cgoAllocMap).Free()
		x.refed768141 = nil
	}
}

// Newext_sensors_data_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newext_sensors_data_tRef(ref unsafe.Pointer) *ext_sensors_data_t {
	if ref == nil {
		return nil
	}
	obj := new(ext_sensors_data_t)
	obj.refed768141 = (*C.ext_sensors_data_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *ext_sensors_data_t) PassRef() (*C.ext_sensors_data_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refed768141 != nil {
		return x.refed768141, nil
	}
	memed768141 := allocExt_sensors_data_tMemory(1)
	refed768141 := (*C.ext_sensors_data_t)(memed768141)
	allocsed768141 := new(cgoAllocMap)
	allocsed768141.Add(memed768141)

	var cconsumed_energy_allocs *cgoAllocMap
	refed768141.consumed_energy, cconsumed_energy_allocs = (C.uint64_t)(x.consumed_energy), cgoAllocsUnknown
	allocsed768141.Borrow(cconsumed_energy_allocs)

	var ctemperature_allocs *cgoAllocMap
	refed768141.temperature, ctemperature_allocs = (C.uint32_t)(x.temperature), cgoAllocsUnknown
	allocsed768141.Borrow(ctemperature_allocs)

	var cenergy_update_time_allocs *cgoAllocMap
	refed768141.energy_update_time, cenergy_update_time_allocs = (C.time_t)(x.energy_update_time), cgoAllocsUnknown
	allocsed768141.Borrow(cenergy_update_time_allocs)

	var ccurrent_watts_allocs *cgoAllocMap
	refed768141.current_watts, ccurrent_watts_allocs = (C.uint32_t)(x.current_watts), cgoAllocsUnknown
	allocsed768141.Borrow(ccurrent_watts_allocs)

	x.refed768141 = refed768141
	x.allocsed768141 = allocsed768141
	return refed768141, allocsed768141

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x ext_sensors_data_t) PassValue() (C.ext_sensors_data_t, *cgoAllocMap) {
	if x.refed768141 != nil {
		return *x.refed768141, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *ext_sensors_data_t) Deref() {
	if x.refed768141 == nil {
		return
	}
	x.consumed_energy = (uint64_t)(x.refed768141.consumed_energy)
	x.temperature = (uint32_t)(x.refed768141.temperature)
	x.energy_update_time = (time_t)(x.refed768141.energy_update_time)
	x.current_watts = (uint32_t)(x.refed768141.current_watts)
}

// allocPower_mgmt_data_tMemory allocates memory for type C.power_mgmt_data_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPower_mgmt_data_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPower_mgmt_data_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPower_mgmt_data_tValue = unsafe.Sizeof([1]C.power_mgmt_data_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *power_mgmt_data_t) Ref() *C.power_mgmt_data_t {
	if x == nil {
		return nil
	}
	return x.ref13703811
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *power_mgmt_data_t) Free() {
	if x != nil && x.allocs13703811 != nil {
		x.allocs13703811.(*cgoAllocMap).Free()
		x.ref13703811 = nil
	}
}

// Newpower_mgmt_data_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newpower_mgmt_data_tRef(ref unsafe.Pointer) *power_mgmt_data_t {
	if ref == nil {
		return nil
	}
	obj := new(power_mgmt_data_t)
	obj.ref13703811 = (*C.power_mgmt_data_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *power_mgmt_data_t) PassRef() (*C.power_mgmt_data_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref13703811 != nil {
		return x.ref13703811, nil
	}
	mem13703811 := allocPower_mgmt_data_tMemory(1)
	ref13703811 := (*C.power_mgmt_data_t)(mem13703811)
	allocs13703811 := new(cgoAllocMap)
	allocs13703811.Add(mem13703811)

	var ccap_watts_allocs *cgoAllocMap
	ref13703811.cap_watts, ccap_watts_allocs = (C.uint32_t)(x.cap_watts), cgoAllocsUnknown
	allocs13703811.Borrow(ccap_watts_allocs)

	var ccurrent_watts_allocs *cgoAllocMap
	ref13703811.current_watts, ccurrent_watts_allocs = (C.uint32_t)(x.current_watts), cgoAllocsUnknown
	allocs13703811.Borrow(ccurrent_watts_allocs)

	var cjoule_counter_allocs *cgoAllocMap
	ref13703811.joule_counter, cjoule_counter_allocs = (C.uint64_t)(x.joule_counter), cgoAllocsUnknown
	allocs13703811.Borrow(cjoule_counter_allocs)

	var cnew_cap_watts_allocs *cgoAllocMap
	ref13703811.new_cap_watts, cnew_cap_watts_allocs = (C.uint32_t)(x.new_cap_watts), cgoAllocsUnknown
	allocs13703811.Borrow(cnew_cap_watts_allocs)

	var cmax_watts_allocs *cgoAllocMap
	ref13703811.max_watts, cmax_watts_allocs = (C.uint32_t)(x.max_watts), cgoAllocsUnknown
	allocs13703811.Borrow(cmax_watts_allocs)

	var cmin_watts_allocs *cgoAllocMap
	ref13703811.min_watts, cmin_watts_allocs = (C.uint32_t)(x.min_watts), cgoAllocsUnknown
	allocs13703811.Borrow(cmin_watts_allocs)

	var cnew_job_time_allocs *cgoAllocMap
	ref13703811.new_job_time, cnew_job_time_allocs = (C.time_t)(x.new_job_time), cgoAllocsUnknown
	allocs13703811.Borrow(cnew_job_time_allocs)

	var cstate_allocs *cgoAllocMap
	ref13703811.state, cstate_allocs = (C.uint16_t)(x.state), cgoAllocsUnknown
	allocs13703811.Borrow(cstate_allocs)

	var ctime_usec_allocs *cgoAllocMap
	ref13703811.time_usec, ctime_usec_allocs = (C.uint64_t)(x.time_usec), cgoAllocsUnknown
	allocs13703811.Borrow(ctime_usec_allocs)

	x.ref13703811 = ref13703811
	x.allocs13703811 = allocs13703811
	return ref13703811, allocs13703811

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x power_mgmt_data_t) PassValue() (C.power_mgmt_data_t, *cgoAllocMap) {
	if x.ref13703811 != nil {
		return *x.ref13703811, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *power_mgmt_data_t) Deref() {
	if x.ref13703811 == nil {
		return
	}
	x.cap_watts = (uint32_t)(x.ref13703811.cap_watts)
	x.current_watts = (uint32_t)(x.ref13703811.current_watts)
	x.joule_counter = (uint64_t)(x.ref13703811.joule_counter)
	x.new_cap_watts = (uint32_t)(x.ref13703811.new_cap_watts)
	x.max_watts = (uint32_t)(x.ref13703811.max_watts)
	x.min_watts = (uint32_t)(x.ref13703811.min_watts)
	x.new_job_time = (time_t)(x.ref13703811.new_job_time)
	x.state = (uint16_t)(x.ref13703811.state)
	x.time_usec = (uint64_t)(x.ref13703811.time_usec)
}

// allocJob_desc_msg_tMemory allocates memory for type C.job_desc_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_desc_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_desc_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_desc_msg_tValue = unsafe.Sizeof([1]C.job_desc_msg_t{})

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

// unpackSDynamic_plugin_data_t transforms a sliced Go data structure into plain C format.
func unpackSDynamic_plugin_data_t(x []dynamic_plugin_data_t) (unpacked *C.dynamic_plugin_data_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.dynamic_plugin_data_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocDynamic_plugin_data_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.dynamic_plugin_data_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.dynamic_plugin_data_t)(unsafe.Pointer(h.Data))
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

// packSDynamic_plugin_data_t reads sliced Go data structure out from plain C format.
func packSDynamic_plugin_data_t(v []dynamic_plugin_data_t, ptr0 *C.dynamic_plugin_data_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfDynamic_plugin_data_tValue]C.dynamic_plugin_data_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newdynamic_plugin_data_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_desc_msg_t) Ref() *C.job_desc_msg_t {
	if x == nil {
		return nil
	}
	return x.ref1da77736
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_desc_msg_t) Free() {
	if x != nil && x.allocs1da77736 != nil {
		x.allocs1da77736.(*cgoAllocMap).Free()
		x.ref1da77736 = nil
	}
}

// Newjob_desc_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_desc_msg_tRef(ref unsafe.Pointer) *job_desc_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(job_desc_msg_t)
	obj.ref1da77736 = (*C.job_desc_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_desc_msg_t) PassRef() (*C.job_desc_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref1da77736 != nil {
		return x.ref1da77736, nil
	}
	mem1da77736 := allocJob_desc_msg_tMemory(1)
	ref1da77736 := (*C.job_desc_msg_t)(mem1da77736)
	allocs1da77736 := new(cgoAllocMap)
	allocs1da77736.Add(mem1da77736)

	var caccount_allocs *cgoAllocMap
	ref1da77736.account, caccount_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.account)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(caccount_allocs)

	var cacctg_freq_allocs *cgoAllocMap
	ref1da77736.acctg_freq, cacctg_freq_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.acctg_freq)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cacctg_freq_allocs)

	var cadmin_comment_allocs *cgoAllocMap
	ref1da77736.admin_comment, cadmin_comment_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.admin_comment)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cadmin_comment_allocs)

	var calloc_node_allocs *cgoAllocMap
	ref1da77736.alloc_node, calloc_node_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.alloc_node)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(calloc_node_allocs)

	var calloc_resp_port_allocs *cgoAllocMap
	ref1da77736.alloc_resp_port, calloc_resp_port_allocs = (C.uint16_t)(x.alloc_resp_port), cgoAllocsUnknown
	allocs1da77736.Borrow(calloc_resp_port_allocs)

	var calloc_sid_allocs *cgoAllocMap
	ref1da77736.alloc_sid, calloc_sid_allocs = (C.uint32_t)(x.alloc_sid), cgoAllocsUnknown
	allocs1da77736.Borrow(calloc_sid_allocs)

	var cargc_allocs *cgoAllocMap
	ref1da77736.argc, cargc_allocs = (C.uint32_t)(x.argc), cgoAllocsUnknown
	allocs1da77736.Borrow(cargc_allocs)

	var cargv_allocs *cgoAllocMap
	ref1da77736.argv, cargv_allocs = unpackSSByte(x.argv)
	allocs1da77736.Borrow(cargv_allocs)

	var carray_inx_allocs *cgoAllocMap
	ref1da77736.array_inx, carray_inx_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.array_inx)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(carray_inx_allocs)

	var carray_bitmap_allocs *cgoAllocMap
	ref1da77736.array_bitmap, carray_bitmap_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.array_bitmap)), cgoAllocsUnknown
	allocs1da77736.Borrow(carray_bitmap_allocs)

	var cbegin_time_allocs *cgoAllocMap
	ref1da77736.begin_time, cbegin_time_allocs = (C.time_t)(x.begin_time), cgoAllocsUnknown
	allocs1da77736.Borrow(cbegin_time_allocs)

	var cbitflags_allocs *cgoAllocMap
	ref1da77736.bitflags, cbitflags_allocs = (C.uint32_t)(x.bitflags), cgoAllocsUnknown
	allocs1da77736.Borrow(cbitflags_allocs)

	var cburst_buffer_allocs *cgoAllocMap
	ref1da77736.burst_buffer, cburst_buffer_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.burst_buffer)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cburst_buffer_allocs)

	var cckpt_interval_allocs *cgoAllocMap
	ref1da77736.ckpt_interval, cckpt_interval_allocs = (C.uint16_t)(x.ckpt_interval), cgoAllocsUnknown
	allocs1da77736.Borrow(cckpt_interval_allocs)

	var cckpt_dir_allocs *cgoAllocMap
	ref1da77736.ckpt_dir, cckpt_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ckpt_dir)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cckpt_dir_allocs)

	var cclusters_allocs *cgoAllocMap
	ref1da77736.clusters, cclusters_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.clusters)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cclusters_allocs)

	var ccluster_features_allocs *cgoAllocMap
	ref1da77736.cluster_features, ccluster_features_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster_features)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(ccluster_features_allocs)

	var ccomment_allocs *cgoAllocMap
	ref1da77736.comment, ccomment_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.comment)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(ccomment_allocs)

	var ccontiguous_allocs *cgoAllocMap
	ref1da77736.contiguous, ccontiguous_allocs = (C.uint16_t)(x.contiguous), cgoAllocsUnknown
	allocs1da77736.Borrow(ccontiguous_allocs)

	var ccore_spec_allocs *cgoAllocMap
	ref1da77736.core_spec, ccore_spec_allocs = (C.uint16_t)(x.core_spec), cgoAllocsUnknown
	allocs1da77736.Borrow(ccore_spec_allocs)

	var ccpu_bind_allocs *cgoAllocMap
	ref1da77736.cpu_bind, ccpu_bind_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cpu_bind)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpu_bind_allocs)

	var ccpu_bind_type_allocs *cgoAllocMap
	ref1da77736.cpu_bind_type, ccpu_bind_type_allocs = (C.uint16_t)(x.cpu_bind_type), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpu_bind_type_allocs)

	var ccpu_freq_min_allocs *cgoAllocMap
	ref1da77736.cpu_freq_min, ccpu_freq_min_allocs = (C.uint32_t)(x.cpu_freq_min), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpu_freq_min_allocs)

	var ccpu_freq_max_allocs *cgoAllocMap
	ref1da77736.cpu_freq_max, ccpu_freq_max_allocs = (C.uint32_t)(x.cpu_freq_max), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpu_freq_max_allocs)

	var ccpu_freq_gov_allocs *cgoAllocMap
	ref1da77736.cpu_freq_gov, ccpu_freq_gov_allocs = (C.uint32_t)(x.cpu_freq_gov), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpu_freq_gov_allocs)

	var cdeadline_allocs *cgoAllocMap
	ref1da77736.deadline, cdeadline_allocs = (C.time_t)(x.deadline), cgoAllocsUnknown
	allocs1da77736.Borrow(cdeadline_allocs)

	var cdelay_boot_allocs *cgoAllocMap
	ref1da77736.delay_boot, cdelay_boot_allocs = (C.uint32_t)(x.delay_boot), cgoAllocsUnknown
	allocs1da77736.Borrow(cdelay_boot_allocs)

	var cdependency_allocs *cgoAllocMap
	ref1da77736.dependency, cdependency_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.dependency)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cdependency_allocs)

	var cend_time_allocs *cgoAllocMap
	ref1da77736.end_time, cend_time_allocs = (C.time_t)(x.end_time), cgoAllocsUnknown
	allocs1da77736.Borrow(cend_time_allocs)

	var cenvironment_allocs *cgoAllocMap
	ref1da77736.environment, cenvironment_allocs = unpackSSByte(x.environment)
	allocs1da77736.Borrow(cenvironment_allocs)

	var cenv_size_allocs *cgoAllocMap
	ref1da77736.env_size, cenv_size_allocs = (C.uint32_t)(x.env_size), cgoAllocsUnknown
	allocs1da77736.Borrow(cenv_size_allocs)

	var cextra_allocs *cgoAllocMap
	ref1da77736.extra, cextra_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.extra)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cextra_allocs)

	var cexc_nodes_allocs *cgoAllocMap
	ref1da77736.exc_nodes, cexc_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.exc_nodes)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cexc_nodes_allocs)

	var cfeatures_allocs *cgoAllocMap
	ref1da77736.features, cfeatures_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.features)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cfeatures_allocs)

	var cfed_siblings_active_allocs *cgoAllocMap
	ref1da77736.fed_siblings_active, cfed_siblings_active_allocs = (C.uint64_t)(x.fed_siblings_active), cgoAllocsUnknown
	allocs1da77736.Borrow(cfed_siblings_active_allocs)

	var cfed_siblings_viable_allocs *cgoAllocMap
	ref1da77736.fed_siblings_viable, cfed_siblings_viable_allocs = (C.uint64_t)(x.fed_siblings_viable), cgoAllocsUnknown
	allocs1da77736.Borrow(cfed_siblings_viable_allocs)

	var cgres_allocs *cgoAllocMap
	ref1da77736.gres, cgres_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.gres)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cgres_allocs)

	var cgroup_id_allocs *cgoAllocMap
	ref1da77736.group_id, cgroup_id_allocs = (C.uint32_t)(x.group_id), cgoAllocsUnknown
	allocs1da77736.Borrow(cgroup_id_allocs)

	var cimmediate_allocs *cgoAllocMap
	ref1da77736.immediate, cimmediate_allocs = (C.uint16_t)(x.immediate), cgoAllocsUnknown
	allocs1da77736.Borrow(cimmediate_allocs)

	var cjob_id_allocs *cgoAllocMap
	ref1da77736.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs1da77736.Borrow(cjob_id_allocs)

	var cjob_id_str_allocs *cgoAllocMap
	ref1da77736.job_id_str, cjob_id_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_id_str)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cjob_id_str_allocs)

	var ckill_on_node_fail_allocs *cgoAllocMap
	ref1da77736.kill_on_node_fail, ckill_on_node_fail_allocs = (C.uint16_t)(x.kill_on_node_fail), cgoAllocsUnknown
	allocs1da77736.Borrow(ckill_on_node_fail_allocs)

	var clicenses_allocs *cgoAllocMap
	ref1da77736.licenses, clicenses_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.licenses)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(clicenses_allocs)

	var cmail_type_allocs *cgoAllocMap
	ref1da77736.mail_type, cmail_type_allocs = (C.uint16_t)(x.mail_type), cgoAllocsUnknown
	allocs1da77736.Borrow(cmail_type_allocs)

	var cmail_user_allocs *cgoAllocMap
	ref1da77736.mail_user, cmail_user_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mail_user)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cmail_user_allocs)

	var cmcs_label_allocs *cgoAllocMap
	ref1da77736.mcs_label, cmcs_label_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mcs_label)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cmcs_label_allocs)

	var cmem_bind_allocs *cgoAllocMap
	ref1da77736.mem_bind, cmem_bind_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mem_bind)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cmem_bind_allocs)

	var cmem_bind_type_allocs *cgoAllocMap
	ref1da77736.mem_bind_type, cmem_bind_type_allocs = (C.uint16_t)(x.mem_bind_type), cgoAllocsUnknown
	allocs1da77736.Borrow(cmem_bind_type_allocs)

	var cname_allocs *cgoAllocMap
	ref1da77736.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cname_allocs)

	var cnetwork_allocs *cgoAllocMap
	ref1da77736.network, cnetwork_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.network)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cnetwork_allocs)

	var cnice_allocs *cgoAllocMap
	ref1da77736.nice, cnice_allocs = (C.uint32_t)(x.nice), cgoAllocsUnknown
	allocs1da77736.Borrow(cnice_allocs)

	var cnum_tasks_allocs *cgoAllocMap
	ref1da77736.num_tasks, cnum_tasks_allocs = (C.uint32_t)(x.num_tasks), cgoAllocsUnknown
	allocs1da77736.Borrow(cnum_tasks_allocs)

	var copen_mode_allocs *cgoAllocMap
	ref1da77736.open_mode, copen_mode_allocs = (C.uint8_t)(x.open_mode), cgoAllocsUnknown
	allocs1da77736.Borrow(copen_mode_allocs)

	var corigin_cluster_allocs *cgoAllocMap
	ref1da77736.origin_cluster, corigin_cluster_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.origin_cluster)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(corigin_cluster_allocs)

	var cother_port_allocs *cgoAllocMap
	ref1da77736.other_port, cother_port_allocs = (C.uint16_t)(x.other_port), cgoAllocsUnknown
	allocs1da77736.Borrow(cother_port_allocs)

	var covercommit_allocs *cgoAllocMap
	ref1da77736.overcommit, covercommit_allocs = (C.uint8_t)(x.overcommit), cgoAllocsUnknown
	allocs1da77736.Borrow(covercommit_allocs)

	var cpack_job_offset_allocs *cgoAllocMap
	ref1da77736.pack_job_offset, cpack_job_offset_allocs = (C.uint32_t)(x.pack_job_offset), cgoAllocsUnknown
	allocs1da77736.Borrow(cpack_job_offset_allocs)

	var cpartition_allocs *cgoAllocMap
	ref1da77736.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partition)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cpartition_allocs)

	var cplane_size_allocs *cgoAllocMap
	ref1da77736.plane_size, cplane_size_allocs = (C.uint16_t)(x.plane_size), cgoAllocsUnknown
	allocs1da77736.Borrow(cplane_size_allocs)

	var cpower_flags_allocs *cgoAllocMap
	ref1da77736.power_flags, cpower_flags_allocs = (C.uint8_t)(x.power_flags), cgoAllocsUnknown
	allocs1da77736.Borrow(cpower_flags_allocs)

	var cpriority_allocs *cgoAllocMap
	ref1da77736.priority, cpriority_allocs = (C.uint32_t)(x.priority), cgoAllocsUnknown
	allocs1da77736.Borrow(cpriority_allocs)

	var cprofile_allocs *cgoAllocMap
	ref1da77736.profile, cprofile_allocs = (C.uint32_t)(x.profile), cgoAllocsUnknown
	allocs1da77736.Borrow(cprofile_allocs)

	var cqos_allocs *cgoAllocMap
	ref1da77736.qos, cqos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.qos)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cqos_allocs)

	var creboot_allocs *cgoAllocMap
	ref1da77736.reboot, creboot_allocs = (C.uint16_t)(x.reboot), cgoAllocsUnknown
	allocs1da77736.Borrow(creboot_allocs)

	var cresp_host_allocs *cgoAllocMap
	ref1da77736.resp_host, cresp_host_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.resp_host)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cresp_host_allocs)

	var crestart_cnt_allocs *cgoAllocMap
	ref1da77736.restart_cnt, crestart_cnt_allocs = (C.uint16_t)(x.restart_cnt), cgoAllocsUnknown
	allocs1da77736.Borrow(crestart_cnt_allocs)

	var creq_nodes_allocs *cgoAllocMap
	ref1da77736.req_nodes, creq_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.req_nodes)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(creq_nodes_allocs)

	var crequeue_allocs *cgoAllocMap
	ref1da77736.requeue, crequeue_allocs = (C.uint16_t)(x.requeue), cgoAllocsUnknown
	allocs1da77736.Borrow(crequeue_allocs)

	var creservation_allocs *cgoAllocMap
	ref1da77736.reservation, creservation_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.reservation)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(creservation_allocs)

	var cscript_allocs *cgoAllocMap
	ref1da77736.script, cscript_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.script)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cscript_allocs)

	var cshared_allocs *cgoAllocMap
	ref1da77736.shared, cshared_allocs = (C.uint16_t)(x.shared), cgoAllocsUnknown
	allocs1da77736.Borrow(cshared_allocs)

	var cspank_job_env_allocs *cgoAllocMap
	ref1da77736.spank_job_env, cspank_job_env_allocs = unpackSSByte(x.spank_job_env)
	allocs1da77736.Borrow(cspank_job_env_allocs)

	var cspank_job_env_size_allocs *cgoAllocMap
	ref1da77736.spank_job_env_size, cspank_job_env_size_allocs = (C.uint32_t)(x.spank_job_env_size), cgoAllocsUnknown
	allocs1da77736.Borrow(cspank_job_env_size_allocs)

	var ctask_dist_allocs *cgoAllocMap
	ref1da77736.task_dist, ctask_dist_allocs = (C.uint32_t)(x.task_dist), cgoAllocsUnknown
	allocs1da77736.Borrow(ctask_dist_allocs)

	var ctime_limit_allocs *cgoAllocMap
	ref1da77736.time_limit, ctime_limit_allocs = (C.uint32_t)(x.time_limit), cgoAllocsUnknown
	allocs1da77736.Borrow(ctime_limit_allocs)

	var ctime_min_allocs *cgoAllocMap
	ref1da77736.time_min, ctime_min_allocs = (C.uint32_t)(x.time_min), cgoAllocsUnknown
	allocs1da77736.Borrow(ctime_min_allocs)

	var cuser_id_allocs *cgoAllocMap
	ref1da77736.user_id, cuser_id_allocs = (C.uint32_t)(x.user_id), cgoAllocsUnknown
	allocs1da77736.Borrow(cuser_id_allocs)

	var cwait_all_nodes_allocs *cgoAllocMap
	ref1da77736.wait_all_nodes, cwait_all_nodes_allocs = (C.uint16_t)(x.wait_all_nodes), cgoAllocsUnknown
	allocs1da77736.Borrow(cwait_all_nodes_allocs)

	var cwarn_flags_allocs *cgoAllocMap
	ref1da77736.warn_flags, cwarn_flags_allocs = (C.uint16_t)(x.warn_flags), cgoAllocsUnknown
	allocs1da77736.Borrow(cwarn_flags_allocs)

	var cwarn_signal_allocs *cgoAllocMap
	ref1da77736.warn_signal, cwarn_signal_allocs = (C.uint16_t)(x.warn_signal), cgoAllocsUnknown
	allocs1da77736.Borrow(cwarn_signal_allocs)

	var cwarn_time_allocs *cgoAllocMap
	ref1da77736.warn_time, cwarn_time_allocs = (C.uint16_t)(x.warn_time), cgoAllocsUnknown
	allocs1da77736.Borrow(cwarn_time_allocs)

	var cwork_dir_allocs *cgoAllocMap
	ref1da77736.work_dir, cwork_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.work_dir)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cwork_dir_allocs)

	var ccpus_per_task_allocs *cgoAllocMap
	ref1da77736.cpus_per_task, ccpus_per_task_allocs = (C.uint16_t)(x.cpus_per_task), cgoAllocsUnknown
	allocs1da77736.Borrow(ccpus_per_task_allocs)

	var cmin_cpus_allocs *cgoAllocMap
	ref1da77736.min_cpus, cmin_cpus_allocs = (C.uint32_t)(x.min_cpus), cgoAllocsUnknown
	allocs1da77736.Borrow(cmin_cpus_allocs)

	var cmax_cpus_allocs *cgoAllocMap
	ref1da77736.max_cpus, cmax_cpus_allocs = (C.uint32_t)(x.max_cpus), cgoAllocsUnknown
	allocs1da77736.Borrow(cmax_cpus_allocs)

	var cmin_nodes_allocs *cgoAllocMap
	ref1da77736.min_nodes, cmin_nodes_allocs = (C.uint32_t)(x.min_nodes), cgoAllocsUnknown
	allocs1da77736.Borrow(cmin_nodes_allocs)

	var cmax_nodes_allocs *cgoAllocMap
	ref1da77736.max_nodes, cmax_nodes_allocs = (C.uint32_t)(x.max_nodes), cgoAllocsUnknown
	allocs1da77736.Borrow(cmax_nodes_allocs)

	var cboards_per_node_allocs *cgoAllocMap
	ref1da77736.boards_per_node, cboards_per_node_allocs = (C.uint16_t)(x.boards_per_node), cgoAllocsUnknown
	allocs1da77736.Borrow(cboards_per_node_allocs)

	var csockets_per_board_allocs *cgoAllocMap
	ref1da77736.sockets_per_board, csockets_per_board_allocs = (C.uint16_t)(x.sockets_per_board), cgoAllocsUnknown
	allocs1da77736.Borrow(csockets_per_board_allocs)

	var csockets_per_node_allocs *cgoAllocMap
	ref1da77736.sockets_per_node, csockets_per_node_allocs = (C.uint16_t)(x.sockets_per_node), cgoAllocsUnknown
	allocs1da77736.Borrow(csockets_per_node_allocs)

	var ccores_per_socket_allocs *cgoAllocMap
	ref1da77736.cores_per_socket, ccores_per_socket_allocs = (C.uint16_t)(x.cores_per_socket), cgoAllocsUnknown
	allocs1da77736.Borrow(ccores_per_socket_allocs)

	var cthreads_per_core_allocs *cgoAllocMap
	ref1da77736.threads_per_core, cthreads_per_core_allocs = (C.uint16_t)(x.threads_per_core), cgoAllocsUnknown
	allocs1da77736.Borrow(cthreads_per_core_allocs)

	var cntasks_per_node_allocs *cgoAllocMap
	ref1da77736.ntasks_per_node, cntasks_per_node_allocs = (C.uint16_t)(x.ntasks_per_node), cgoAllocsUnknown
	allocs1da77736.Borrow(cntasks_per_node_allocs)

	var cntasks_per_socket_allocs *cgoAllocMap
	ref1da77736.ntasks_per_socket, cntasks_per_socket_allocs = (C.uint16_t)(x.ntasks_per_socket), cgoAllocsUnknown
	allocs1da77736.Borrow(cntasks_per_socket_allocs)

	var cntasks_per_core_allocs *cgoAllocMap
	ref1da77736.ntasks_per_core, cntasks_per_core_allocs = (C.uint16_t)(x.ntasks_per_core), cgoAllocsUnknown
	allocs1da77736.Borrow(cntasks_per_core_allocs)

	var cntasks_per_board_allocs *cgoAllocMap
	ref1da77736.ntasks_per_board, cntasks_per_board_allocs = (C.uint16_t)(x.ntasks_per_board), cgoAllocsUnknown
	allocs1da77736.Borrow(cntasks_per_board_allocs)

	var cpn_min_cpus_allocs *cgoAllocMap
	ref1da77736.pn_min_cpus, cpn_min_cpus_allocs = (C.uint16_t)(x.pn_min_cpus), cgoAllocsUnknown
	allocs1da77736.Borrow(cpn_min_cpus_allocs)

	var cpn_min_memory_allocs *cgoAllocMap
	ref1da77736.pn_min_memory, cpn_min_memory_allocs = (C.uint64_t)(x.pn_min_memory), cgoAllocsUnknown
	allocs1da77736.Borrow(cpn_min_memory_allocs)

	var cpn_min_tmp_disk_allocs *cgoAllocMap
	ref1da77736.pn_min_tmp_disk, cpn_min_tmp_disk_allocs = (C.uint32_t)(x.pn_min_tmp_disk), cgoAllocsUnknown
	allocs1da77736.Borrow(cpn_min_tmp_disk_allocs)

	var cgeometry_allocs *cgoAllocMap
	ref1da77736.geometry, cgeometry_allocs = *(*[5]C.uint16_t)(unsafe.Pointer(&x.geometry)), cgoAllocsUnknown
	allocs1da77736.Borrow(cgeometry_allocs)

	var cconn_type_allocs *cgoAllocMap
	ref1da77736.conn_type, cconn_type_allocs = *(*[5]C.uint16_t)(unsafe.Pointer(&x.conn_type)), cgoAllocsUnknown
	allocs1da77736.Borrow(cconn_type_allocs)

	var crotate_allocs *cgoAllocMap
	ref1da77736.rotate, crotate_allocs = (C.uint16_t)(x.rotate), cgoAllocsUnknown
	allocs1da77736.Borrow(crotate_allocs)

	var cblrtsimage_allocs *cgoAllocMap
	ref1da77736.blrtsimage, cblrtsimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.blrtsimage)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cblrtsimage_allocs)

	var clinuximage_allocs *cgoAllocMap
	ref1da77736.linuximage, clinuximage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.linuximage)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(clinuximage_allocs)

	var cmloaderimage_allocs *cgoAllocMap
	ref1da77736.mloaderimage, cmloaderimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mloaderimage)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cmloaderimage_allocs)

	var cramdiskimage_allocs *cgoAllocMap
	ref1da77736.ramdiskimage, cramdiskimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ramdiskimage)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cramdiskimage_allocs)

	var creq_switch_allocs *cgoAllocMap
	ref1da77736.req_switch, creq_switch_allocs = (C.uint32_t)(x.req_switch), cgoAllocsUnknown
	allocs1da77736.Borrow(creq_switch_allocs)

	var cselect_jobinfo_allocs *cgoAllocMap
	ref1da77736.select_jobinfo, cselect_jobinfo_allocs = unpackSDynamic_plugin_data_t(x.select_jobinfo)
	allocs1da77736.Borrow(cselect_jobinfo_allocs)

	var cstd_err_allocs *cgoAllocMap
	ref1da77736.std_err, cstd_err_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.std_err)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cstd_err_allocs)

	var cstd_in_allocs *cgoAllocMap
	ref1da77736.std_in, cstd_in_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.std_in)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cstd_in_allocs)

	var cstd_out_allocs *cgoAllocMap
	ref1da77736.std_out, cstd_out_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.std_out)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cstd_out_allocs)

	var ctres_req_cnt_allocs *cgoAllocMap
	ref1da77736.tres_req_cnt, ctres_req_cnt_allocs = (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_req_cnt)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(ctres_req_cnt_allocs)

	var cwait4switch_allocs *cgoAllocMap
	ref1da77736.wait4switch, cwait4switch_allocs = (C.uint32_t)(x.wait4switch), cgoAllocsUnknown
	allocs1da77736.Borrow(cwait4switch_allocs)

	var cwckey_allocs *cgoAllocMap
	ref1da77736.wckey, cwckey_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.wckey)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cwckey_allocs)

	var cx11_allocs *cgoAllocMap
	ref1da77736.x11, cx11_allocs = (C.uint16_t)(x.x11), cgoAllocsUnknown
	allocs1da77736.Borrow(cx11_allocs)

	var cx11_magic_cookie_allocs *cgoAllocMap
	ref1da77736.x11_magic_cookie, cx11_magic_cookie_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.x11_magic_cookie)).Data)), cgoAllocsUnknown
	allocs1da77736.Borrow(cx11_magic_cookie_allocs)

	var cx11_target_port_allocs *cgoAllocMap
	ref1da77736.x11_target_port, cx11_target_port_allocs = (C.uint16_t)(x.x11_target_port), cgoAllocsUnknown
	allocs1da77736.Borrow(cx11_target_port_allocs)

	x.ref1da77736 = ref1da77736
	x.allocs1da77736 = allocs1da77736
	return ref1da77736, allocs1da77736

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_desc_msg_t) PassValue() (C.job_desc_msg_t, *cgoAllocMap) {
	if x.ref1da77736 != nil {
		return *x.ref1da77736, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_desc_msg_t) Deref() {
	if x.ref1da77736 == nil {
		return
	}
	hxfc4425b := (*sliceHeader)(unsafe.Pointer(&x.account))
	hxfc4425b.Data = uintptr(unsafe.Pointer(x.ref1da77736.account))
	hxfc4425b.Cap = 0x7fffffff
	// hxfc4425b.Len = ?

	hxf95e7c8 := (*sliceHeader)(unsafe.Pointer(&x.acctg_freq))
	hxf95e7c8.Data = uintptr(unsafe.Pointer(x.ref1da77736.acctg_freq))
	hxf95e7c8.Cap = 0x7fffffff
	// hxf95e7c8.Len = ?

	hxff2234b := (*sliceHeader)(unsafe.Pointer(&x.admin_comment))
	hxff2234b.Data = uintptr(unsafe.Pointer(x.ref1da77736.admin_comment))
	hxff2234b.Cap = 0x7fffffff
	// hxff2234b.Len = ?

	hxff73280 := (*sliceHeader)(unsafe.Pointer(&x.alloc_node))
	hxff73280.Data = uintptr(unsafe.Pointer(x.ref1da77736.alloc_node))
	hxff73280.Cap = 0x7fffffff
	// hxff73280.Len = ?

	x.alloc_resp_port = (uint16_t)(x.ref1da77736.alloc_resp_port)
	x.alloc_sid = (uint32_t)(x.ref1da77736.alloc_sid)
	x.argc = (uint32_t)(x.ref1da77736.argc)
	packSSByte(x.argv, x.ref1da77736.argv)
	hxfa3f05c := (*sliceHeader)(unsafe.Pointer(&x.array_inx))
	hxfa3f05c.Data = uintptr(unsafe.Pointer(x.ref1da77736.array_inx))
	hxfa3f05c.Cap = 0x7fffffff
	// hxfa3f05c.Len = ?

	x.array_bitmap = (unsafe.Pointer)(unsafe.Pointer(x.ref1da77736.array_bitmap))
	x.begin_time = (time_t)(x.ref1da77736.begin_time)
	x.bitflags = (uint32_t)(x.ref1da77736.bitflags)
	hxf0d18b7 := (*sliceHeader)(unsafe.Pointer(&x.burst_buffer))
	hxf0d18b7.Data = uintptr(unsafe.Pointer(x.ref1da77736.burst_buffer))
	hxf0d18b7.Cap = 0x7fffffff
	// hxf0d18b7.Len = ?

	x.ckpt_interval = (uint16_t)(x.ref1da77736.ckpt_interval)
	hxf2fab0d := (*sliceHeader)(unsafe.Pointer(&x.ckpt_dir))
	hxf2fab0d.Data = uintptr(unsafe.Pointer(x.ref1da77736.ckpt_dir))
	hxf2fab0d.Cap = 0x7fffffff
	// hxf2fab0d.Len = ?

	hxf69fe70 := (*sliceHeader)(unsafe.Pointer(&x.clusters))
	hxf69fe70.Data = uintptr(unsafe.Pointer(x.ref1da77736.clusters))
	hxf69fe70.Cap = 0x7fffffff
	// hxf69fe70.Len = ?

	hxf65bf54 := (*sliceHeader)(unsafe.Pointer(&x.cluster_features))
	hxf65bf54.Data = uintptr(unsafe.Pointer(x.ref1da77736.cluster_features))
	hxf65bf54.Cap = 0x7fffffff
	// hxf65bf54.Len = ?

	hxf3b8dbd := (*sliceHeader)(unsafe.Pointer(&x.comment))
	hxf3b8dbd.Data = uintptr(unsafe.Pointer(x.ref1da77736.comment))
	hxf3b8dbd.Cap = 0x7fffffff
	// hxf3b8dbd.Len = ?

	x.contiguous = (uint16_t)(x.ref1da77736.contiguous)
	x.core_spec = (uint16_t)(x.ref1da77736.core_spec)
	hxf7a6dff := (*sliceHeader)(unsafe.Pointer(&x.cpu_bind))
	hxf7a6dff.Data = uintptr(unsafe.Pointer(x.ref1da77736.cpu_bind))
	hxf7a6dff.Cap = 0x7fffffff
	// hxf7a6dff.Len = ?

	x.cpu_bind_type = (uint16_t)(x.ref1da77736.cpu_bind_type)
	x.cpu_freq_min = (uint32_t)(x.ref1da77736.cpu_freq_min)
	x.cpu_freq_max = (uint32_t)(x.ref1da77736.cpu_freq_max)
	x.cpu_freq_gov = (uint32_t)(x.ref1da77736.cpu_freq_gov)
	x.deadline = (time_t)(x.ref1da77736.deadline)
	x.delay_boot = (uint32_t)(x.ref1da77736.delay_boot)
	hxfe48d67 := (*sliceHeader)(unsafe.Pointer(&x.dependency))
	hxfe48d67.Data = uintptr(unsafe.Pointer(x.ref1da77736.dependency))
	hxfe48d67.Cap = 0x7fffffff
	// hxfe48d67.Len = ?

	x.end_time = (time_t)(x.ref1da77736.end_time)
	packSSByte(x.environment, x.ref1da77736.environment)
	x.env_size = (uint32_t)(x.ref1da77736.env_size)
	hxf058b18 := (*sliceHeader)(unsafe.Pointer(&x.extra))
	hxf058b18.Data = uintptr(unsafe.Pointer(x.ref1da77736.extra))
	hxf058b18.Cap = 0x7fffffff
	// hxf058b18.Len = ?

	hxff6bc57 := (*sliceHeader)(unsafe.Pointer(&x.exc_nodes))
	hxff6bc57.Data = uintptr(unsafe.Pointer(x.ref1da77736.exc_nodes))
	hxff6bc57.Cap = 0x7fffffff
	// hxff6bc57.Len = ?

	hxf5fa529 := (*sliceHeader)(unsafe.Pointer(&x.features))
	hxf5fa529.Data = uintptr(unsafe.Pointer(x.ref1da77736.features))
	hxf5fa529.Cap = 0x7fffffff
	// hxf5fa529.Len = ?

	x.fed_siblings_active = (uint64_t)(x.ref1da77736.fed_siblings_active)
	x.fed_siblings_viable = (uint64_t)(x.ref1da77736.fed_siblings_viable)
	hxf21690b := (*sliceHeader)(unsafe.Pointer(&x.gres))
	hxf21690b.Data = uintptr(unsafe.Pointer(x.ref1da77736.gres))
	hxf21690b.Cap = 0x7fffffff
	// hxf21690b.Len = ?

	x.group_id = (uint32_t)(x.ref1da77736.group_id)
	x.immediate = (uint16_t)(x.ref1da77736.immediate)
	x.job_id = (uint32_t)(x.ref1da77736.job_id)
	hxf1231c9 := (*sliceHeader)(unsafe.Pointer(&x.job_id_str))
	hxf1231c9.Data = uintptr(unsafe.Pointer(x.ref1da77736.job_id_str))
	hxf1231c9.Cap = 0x7fffffff
	// hxf1231c9.Len = ?

	x.kill_on_node_fail = (uint16_t)(x.ref1da77736.kill_on_node_fail)
	hxf04b15b := (*sliceHeader)(unsafe.Pointer(&x.licenses))
	hxf04b15b.Data = uintptr(unsafe.Pointer(x.ref1da77736.licenses))
	hxf04b15b.Cap = 0x7fffffff
	// hxf04b15b.Len = ?

	x.mail_type = (uint16_t)(x.ref1da77736.mail_type)
	hxf2f888b := (*sliceHeader)(unsafe.Pointer(&x.mail_user))
	hxf2f888b.Data = uintptr(unsafe.Pointer(x.ref1da77736.mail_user))
	hxf2f888b.Cap = 0x7fffffff
	// hxf2f888b.Len = ?

	hxf5d1de2 := (*sliceHeader)(unsafe.Pointer(&x.mcs_label))
	hxf5d1de2.Data = uintptr(unsafe.Pointer(x.ref1da77736.mcs_label))
	hxf5d1de2.Cap = 0x7fffffff
	// hxf5d1de2.Len = ?

	hxfe53d34 := (*sliceHeader)(unsafe.Pointer(&x.mem_bind))
	hxfe53d34.Data = uintptr(unsafe.Pointer(x.ref1da77736.mem_bind))
	hxfe53d34.Cap = 0x7fffffff
	// hxfe53d34.Len = ?

	x.mem_bind_type = (uint16_t)(x.ref1da77736.mem_bind_type)
	hxf547023 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf547023.Data = uintptr(unsafe.Pointer(x.ref1da77736.name))
	hxf547023.Cap = 0x7fffffff
	// hxf547023.Len = ?

	hxf5ebb88 := (*sliceHeader)(unsafe.Pointer(&x.network))
	hxf5ebb88.Data = uintptr(unsafe.Pointer(x.ref1da77736.network))
	hxf5ebb88.Cap = 0x7fffffff
	// hxf5ebb88.Len = ?

	x.nice = (uint32_t)(x.ref1da77736.nice)
	x.num_tasks = (uint32_t)(x.ref1da77736.num_tasks)
	x.open_mode = (uint8_t)(x.ref1da77736.open_mode)
	hxff20e84 := (*sliceHeader)(unsafe.Pointer(&x.origin_cluster))
	hxff20e84.Data = uintptr(unsafe.Pointer(x.ref1da77736.origin_cluster))
	hxff20e84.Cap = 0x7fffffff
	// hxff20e84.Len = ?

	x.other_port = (uint16_t)(x.ref1da77736.other_port)
	x.overcommit = (uint8_t)(x.ref1da77736.overcommit)
	x.pack_job_offset = (uint32_t)(x.ref1da77736.pack_job_offset)
	hxfa26a4d := (*sliceHeader)(unsafe.Pointer(&x.partition))
	hxfa26a4d.Data = uintptr(unsafe.Pointer(x.ref1da77736.partition))
	hxfa26a4d.Cap = 0x7fffffff
	// hxfa26a4d.Len = ?

	x.plane_size = (uint16_t)(x.ref1da77736.plane_size)
	x.power_flags = (uint8_t)(x.ref1da77736.power_flags)
	x.priority = (uint32_t)(x.ref1da77736.priority)
	x.profile = (uint32_t)(x.ref1da77736.profile)
	hxfe48098 := (*sliceHeader)(unsafe.Pointer(&x.qos))
	hxfe48098.Data = uintptr(unsafe.Pointer(x.ref1da77736.qos))
	hxfe48098.Cap = 0x7fffffff
	// hxfe48098.Len = ?

	x.reboot = (uint16_t)(x.ref1da77736.reboot)
	hxffe3496 := (*sliceHeader)(unsafe.Pointer(&x.resp_host))
	hxffe3496.Data = uintptr(unsafe.Pointer(x.ref1da77736.resp_host))
	hxffe3496.Cap = 0x7fffffff
	// hxffe3496.Len = ?

	x.restart_cnt = (uint16_t)(x.ref1da77736.restart_cnt)
	hxf5d48a6 := (*sliceHeader)(unsafe.Pointer(&x.req_nodes))
	hxf5d48a6.Data = uintptr(unsafe.Pointer(x.ref1da77736.req_nodes))
	hxf5d48a6.Cap = 0x7fffffff
	// hxf5d48a6.Len = ?

	x.requeue = (uint16_t)(x.ref1da77736.requeue)
	hxf685469 := (*sliceHeader)(unsafe.Pointer(&x.reservation))
	hxf685469.Data = uintptr(unsafe.Pointer(x.ref1da77736.reservation))
	hxf685469.Cap = 0x7fffffff
	// hxf685469.Len = ?

	hxf03a9a7 := (*sliceHeader)(unsafe.Pointer(&x.script))
	hxf03a9a7.Data = uintptr(unsafe.Pointer(x.ref1da77736.script))
	hxf03a9a7.Cap = 0x7fffffff
	// hxf03a9a7.Len = ?

	x.shared = (uint16_t)(x.ref1da77736.shared)
	packSSByte(x.spank_job_env, x.ref1da77736.spank_job_env)
	x.spank_job_env_size = (uint32_t)(x.ref1da77736.spank_job_env_size)
	x.task_dist = (uint32_t)(x.ref1da77736.task_dist)
	x.time_limit = (uint32_t)(x.ref1da77736.time_limit)
	x.time_min = (uint32_t)(x.ref1da77736.time_min)
	x.user_id = (uint32_t)(x.ref1da77736.user_id)
	x.wait_all_nodes = (uint16_t)(x.ref1da77736.wait_all_nodes)
	x.warn_flags = (uint16_t)(x.ref1da77736.warn_flags)
	x.warn_signal = (uint16_t)(x.ref1da77736.warn_signal)
	x.warn_time = (uint16_t)(x.ref1da77736.warn_time)
	hxfe93325 := (*sliceHeader)(unsafe.Pointer(&x.work_dir))
	hxfe93325.Data = uintptr(unsafe.Pointer(x.ref1da77736.work_dir))
	hxfe93325.Cap = 0x7fffffff
	// hxfe93325.Len = ?

	x.cpus_per_task = (uint16_t)(x.ref1da77736.cpus_per_task)
	x.min_cpus = (uint32_t)(x.ref1da77736.min_cpus)
	x.max_cpus = (uint32_t)(x.ref1da77736.max_cpus)
	x.min_nodes = (uint32_t)(x.ref1da77736.min_nodes)
	x.max_nodes = (uint32_t)(x.ref1da77736.max_nodes)
	x.boards_per_node = (uint16_t)(x.ref1da77736.boards_per_node)
	x.sockets_per_board = (uint16_t)(x.ref1da77736.sockets_per_board)
	x.sockets_per_node = (uint16_t)(x.ref1da77736.sockets_per_node)
	x.cores_per_socket = (uint16_t)(x.ref1da77736.cores_per_socket)
	x.threads_per_core = (uint16_t)(x.ref1da77736.threads_per_core)
	x.ntasks_per_node = (uint16_t)(x.ref1da77736.ntasks_per_node)
	x.ntasks_per_socket = (uint16_t)(x.ref1da77736.ntasks_per_socket)
	x.ntasks_per_core = (uint16_t)(x.ref1da77736.ntasks_per_core)
	x.ntasks_per_board = (uint16_t)(x.ref1da77736.ntasks_per_board)
	x.pn_min_cpus = (uint16_t)(x.ref1da77736.pn_min_cpus)
	x.pn_min_memory = (uint64_t)(x.ref1da77736.pn_min_memory)
	x.pn_min_tmp_disk = (uint32_t)(x.ref1da77736.pn_min_tmp_disk)
	x.geometry = *(*[5]uint16_t)(unsafe.Pointer(&x.ref1da77736.geometry))
	x.conn_type = *(*[5]uint16_t)(unsafe.Pointer(&x.ref1da77736.conn_type))
	x.rotate = (uint16_t)(x.ref1da77736.rotate)
	hxf09ea94 := (*sliceHeader)(unsafe.Pointer(&x.blrtsimage))
	hxf09ea94.Data = uintptr(unsafe.Pointer(x.ref1da77736.blrtsimage))
	hxf09ea94.Cap = 0x7fffffff
	// hxf09ea94.Len = ?

	hxfd687ee := (*sliceHeader)(unsafe.Pointer(&x.linuximage))
	hxfd687ee.Data = uintptr(unsafe.Pointer(x.ref1da77736.linuximage))
	hxfd687ee.Cap = 0x7fffffff
	// hxfd687ee.Len = ?

	hxf15a567 := (*sliceHeader)(unsafe.Pointer(&x.mloaderimage))
	hxf15a567.Data = uintptr(unsafe.Pointer(x.ref1da77736.mloaderimage))
	hxf15a567.Cap = 0x7fffffff
	// hxf15a567.Len = ?

	hxf8aebb5 := (*sliceHeader)(unsafe.Pointer(&x.ramdiskimage))
	hxf8aebb5.Data = uintptr(unsafe.Pointer(x.ref1da77736.ramdiskimage))
	hxf8aebb5.Cap = 0x7fffffff
	// hxf8aebb5.Len = ?

	x.req_switch = (uint32_t)(x.ref1da77736.req_switch)
	packSDynamic_plugin_data_t(x.select_jobinfo, x.ref1da77736.select_jobinfo)
	hxf5d30cf := (*sliceHeader)(unsafe.Pointer(&x.std_err))
	hxf5d30cf.Data = uintptr(unsafe.Pointer(x.ref1da77736.std_err))
	hxf5d30cf.Cap = 0x7fffffff
	// hxf5d30cf.Len = ?

	hxf882e98 := (*sliceHeader)(unsafe.Pointer(&x.std_in))
	hxf882e98.Data = uintptr(unsafe.Pointer(x.ref1da77736.std_in))
	hxf882e98.Cap = 0x7fffffff
	// hxf882e98.Len = ?

	hxf992404 := (*sliceHeader)(unsafe.Pointer(&x.std_out))
	hxf992404.Data = uintptr(unsafe.Pointer(x.ref1da77736.std_out))
	hxf992404.Cap = 0x7fffffff
	// hxf992404.Len = ?

	hxf8e0dd2 := (*sliceHeader)(unsafe.Pointer(&x.tres_req_cnt))
	hxf8e0dd2.Data = uintptr(unsafe.Pointer(x.ref1da77736.tres_req_cnt))
	hxf8e0dd2.Cap = 0x7fffffff
	// hxf8e0dd2.Len = ?

	x.wait4switch = (uint32_t)(x.ref1da77736.wait4switch)
	hxf44d909 := (*sliceHeader)(unsafe.Pointer(&x.wckey))
	hxf44d909.Data = uintptr(unsafe.Pointer(x.ref1da77736.wckey))
	hxf44d909.Cap = 0x7fffffff
	// hxf44d909.Len = ?

	x.x11 = (uint16_t)(x.ref1da77736.x11)
	hxfa835e7 := (*sliceHeader)(unsafe.Pointer(&x.x11_magic_cookie))
	hxfa835e7.Data = uintptr(unsafe.Pointer(x.ref1da77736.x11_magic_cookie))
	hxfa835e7.Cap = 0x7fffffff
	// hxfa835e7.Len = ?

	x.x11_target_port = (uint16_t)(x.ref1da77736.x11_target_port)
}

// allocSlurm_job_info_tMemory allocates memory for type C.slurm_job_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_job_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_job_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_job_info_tValue = unsafe.Sizeof([1]C.slurm_job_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_job_info_t) Ref() *C.slurm_job_info_t {
	if x == nil {
		return nil
	}
	return x.refc83953cc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_job_info_t) Free() {
	if x != nil && x.allocsc83953cc != nil {
		x.allocsc83953cc.(*cgoAllocMap).Free()
		x.refc83953cc = nil
	}
}

// Newslurm_job_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_job_info_tRef(ref unsafe.Pointer) *slurm_job_info_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_job_info_t)
	obj.refc83953cc = (*C.slurm_job_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_job_info_t) PassRef() (*C.slurm_job_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc83953cc != nil {
		return x.refc83953cc, nil
	}
	memc83953cc := allocSlurm_job_info_tMemory(1)
	refc83953cc := (*C.slurm_job_info_t)(memc83953cc)
	allocsc83953cc := new(cgoAllocMap)
	allocsc83953cc.Add(memc83953cc)

	var caccount_allocs *cgoAllocMap
	refc83953cc.account, caccount_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.account)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(caccount_allocs)

	var cadmin_comment_allocs *cgoAllocMap
	refc83953cc.admin_comment, cadmin_comment_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.admin_comment)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cadmin_comment_allocs)

	var calloc_node_allocs *cgoAllocMap
	refc83953cc.alloc_node, calloc_node_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.alloc_node)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(calloc_node_allocs)

	var calloc_sid_allocs *cgoAllocMap
	refc83953cc.alloc_sid, calloc_sid_allocs = (C.uint32_t)(x.alloc_sid), cgoAllocsUnknown
	allocsc83953cc.Borrow(calloc_sid_allocs)

	var carray_bitmap_allocs *cgoAllocMap
	refc83953cc.array_bitmap, carray_bitmap_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.array_bitmap)), cgoAllocsUnknown
	allocsc83953cc.Borrow(carray_bitmap_allocs)

	var carray_job_id_allocs *cgoAllocMap
	refc83953cc.array_job_id, carray_job_id_allocs = (C.uint32_t)(x.array_job_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(carray_job_id_allocs)

	var carray_task_id_allocs *cgoAllocMap
	refc83953cc.array_task_id, carray_task_id_allocs = (C.uint32_t)(x.array_task_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(carray_task_id_allocs)

	var carray_max_tasks_allocs *cgoAllocMap
	refc83953cc.array_max_tasks, carray_max_tasks_allocs = (C.uint32_t)(x.array_max_tasks), cgoAllocsUnknown
	allocsc83953cc.Borrow(carray_max_tasks_allocs)

	var carray_task_str_allocs *cgoAllocMap
	refc83953cc.array_task_str, carray_task_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.array_task_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(carray_task_str_allocs)

	var cassoc_id_allocs *cgoAllocMap
	refc83953cc.assoc_id, cassoc_id_allocs = (C.uint32_t)(x.assoc_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(cassoc_id_allocs)

	var cbatch_flag_allocs *cgoAllocMap
	refc83953cc.batch_flag, cbatch_flag_allocs = (C.uint16_t)(x.batch_flag), cgoAllocsUnknown
	allocsc83953cc.Borrow(cbatch_flag_allocs)

	var cbatch_host_allocs *cgoAllocMap
	refc83953cc.batch_host, cbatch_host_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.batch_host)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cbatch_host_allocs)

	var cbitflags_allocs *cgoAllocMap
	refc83953cc.bitflags, cbitflags_allocs = (C.uint32_t)(x.bitflags), cgoAllocsUnknown
	allocsc83953cc.Borrow(cbitflags_allocs)

	var cboards_per_node_allocs *cgoAllocMap
	refc83953cc.boards_per_node, cboards_per_node_allocs = (C.uint16_t)(x.boards_per_node), cgoAllocsUnknown
	allocsc83953cc.Borrow(cboards_per_node_allocs)

	var cburst_buffer_allocs *cgoAllocMap
	refc83953cc.burst_buffer, cburst_buffer_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.burst_buffer)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cburst_buffer_allocs)

	var cburst_buffer_state_allocs *cgoAllocMap
	refc83953cc.burst_buffer_state, cburst_buffer_state_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.burst_buffer_state)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cburst_buffer_state_allocs)

	var ccluster_allocs *cgoAllocMap
	refc83953cc.cluster, ccluster_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccluster_allocs)

	var ccluster_features_allocs *cgoAllocMap
	refc83953cc.cluster_features, ccluster_features_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster_features)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccluster_features_allocs)

	var ccommand_allocs *cgoAllocMap
	refc83953cc.command, ccommand_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.command)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccommand_allocs)

	var ccomment_allocs *cgoAllocMap
	refc83953cc.comment, ccomment_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.comment)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccomment_allocs)

	var ccontiguous_allocs *cgoAllocMap
	refc83953cc.contiguous, ccontiguous_allocs = (C.uint16_t)(x.contiguous), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccontiguous_allocs)

	var ccore_spec_allocs *cgoAllocMap
	refc83953cc.core_spec, ccore_spec_allocs = (C.uint16_t)(x.core_spec), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccore_spec_allocs)

	var ccores_per_socket_allocs *cgoAllocMap
	refc83953cc.cores_per_socket, ccores_per_socket_allocs = (C.uint16_t)(x.cores_per_socket), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccores_per_socket_allocs)

	var cbillable_tres_allocs *cgoAllocMap
	refc83953cc.billable_tres, cbillable_tres_allocs = (C.double)(x.billable_tres), cgoAllocsUnknown
	allocsc83953cc.Borrow(cbillable_tres_allocs)

	var ccpus_per_task_allocs *cgoAllocMap
	refc83953cc.cpus_per_task, ccpus_per_task_allocs = (C.uint16_t)(x.cpus_per_task), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccpus_per_task_allocs)

	var ccpu_freq_min_allocs *cgoAllocMap
	refc83953cc.cpu_freq_min, ccpu_freq_min_allocs = (C.uint32_t)(x.cpu_freq_min), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccpu_freq_min_allocs)

	var ccpu_freq_max_allocs *cgoAllocMap
	refc83953cc.cpu_freq_max, ccpu_freq_max_allocs = (C.uint32_t)(x.cpu_freq_max), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccpu_freq_max_allocs)

	var ccpu_freq_gov_allocs *cgoAllocMap
	refc83953cc.cpu_freq_gov, ccpu_freq_gov_allocs = (C.uint32_t)(x.cpu_freq_gov), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccpu_freq_gov_allocs)

	var cdeadline_allocs *cgoAllocMap
	refc83953cc.deadline, cdeadline_allocs = (C.time_t)(x.deadline), cgoAllocsUnknown
	allocsc83953cc.Borrow(cdeadline_allocs)

	var cdelay_boot_allocs *cgoAllocMap
	refc83953cc.delay_boot, cdelay_boot_allocs = (C.uint32_t)(x.delay_boot), cgoAllocsUnknown
	allocsc83953cc.Borrow(cdelay_boot_allocs)

	var cdependency_allocs *cgoAllocMap
	refc83953cc.dependency, cdependency_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.dependency)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cdependency_allocs)

	var cderived_ec_allocs *cgoAllocMap
	refc83953cc.derived_ec, cderived_ec_allocs = (C.uint32_t)(x.derived_ec), cgoAllocsUnknown
	allocsc83953cc.Borrow(cderived_ec_allocs)

	var celigible_time_allocs *cgoAllocMap
	refc83953cc.eligible_time, celigible_time_allocs = (C.time_t)(x.eligible_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(celigible_time_allocs)

	var cend_time_allocs *cgoAllocMap
	refc83953cc.end_time, cend_time_allocs = (C.time_t)(x.end_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(cend_time_allocs)

	var cexc_nodes_allocs *cgoAllocMap
	refc83953cc.exc_nodes, cexc_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.exc_nodes)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cexc_nodes_allocs)

	var cexc_node_inx_allocs *cgoAllocMap
	refc83953cc.exc_node_inx, cexc_node_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.exc_node_inx)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cexc_node_inx_allocs)

	var cexit_code_allocs *cgoAllocMap
	refc83953cc.exit_code, cexit_code_allocs = (C.uint32_t)(x.exit_code), cgoAllocsUnknown
	allocsc83953cc.Borrow(cexit_code_allocs)

	var cfeatures_allocs *cgoAllocMap
	refc83953cc.features, cfeatures_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.features)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfeatures_allocs)

	var cfed_origin_str_allocs *cgoAllocMap
	refc83953cc.fed_origin_str, cfed_origin_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.fed_origin_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfed_origin_str_allocs)

	var cfed_siblings_active_allocs *cgoAllocMap
	refc83953cc.fed_siblings_active, cfed_siblings_active_allocs = (C.uint64_t)(x.fed_siblings_active), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfed_siblings_active_allocs)

	var cfed_siblings_active_str_allocs *cgoAllocMap
	refc83953cc.fed_siblings_active_str, cfed_siblings_active_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.fed_siblings_active_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfed_siblings_active_str_allocs)

	var cfed_siblings_viable_allocs *cgoAllocMap
	refc83953cc.fed_siblings_viable, cfed_siblings_viable_allocs = (C.uint64_t)(x.fed_siblings_viable), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfed_siblings_viable_allocs)

	var cfed_siblings_viable_str_allocs *cgoAllocMap
	refc83953cc.fed_siblings_viable_str, cfed_siblings_viable_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.fed_siblings_viable_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfed_siblings_viable_str_allocs)

	var cgres_allocs *cgoAllocMap
	refc83953cc.gres, cgres_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.gres)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cgres_allocs)

	var cgres_detail_cnt_allocs *cgoAllocMap
	refc83953cc.gres_detail_cnt, cgres_detail_cnt_allocs = (C.uint32_t)(x.gres_detail_cnt), cgoAllocsUnknown
	allocsc83953cc.Borrow(cgres_detail_cnt_allocs)

	var cgres_detail_str_allocs *cgoAllocMap
	refc83953cc.gres_detail_str, cgres_detail_str_allocs = unpackSSByte(x.gres_detail_str)
	allocsc83953cc.Borrow(cgres_detail_str_allocs)

	var cgroup_id_allocs *cgoAllocMap
	refc83953cc.group_id, cgroup_id_allocs = (C.uint32_t)(x.group_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(cgroup_id_allocs)

	var cjob_id_allocs *cgoAllocMap
	refc83953cc.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(cjob_id_allocs)

	var cjob_resrcs_allocs *cgoAllocMap
	refc83953cc.job_resrcs, cjob_resrcs_allocs = (*C.job_resources_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_resrcs)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cjob_resrcs_allocs)

	var cjob_state_allocs *cgoAllocMap
	refc83953cc.job_state, cjob_state_allocs = (C.uint32_t)(x.job_state), cgoAllocsUnknown
	allocsc83953cc.Borrow(cjob_state_allocs)

	var clast_sched_eval_allocs *cgoAllocMap
	refc83953cc.last_sched_eval, clast_sched_eval_allocs = (C.time_t)(x.last_sched_eval), cgoAllocsUnknown
	allocsc83953cc.Borrow(clast_sched_eval_allocs)

	var clicenses_allocs *cgoAllocMap
	refc83953cc.licenses, clicenses_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.licenses)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(clicenses_allocs)

	var cmax_cpus_allocs *cgoAllocMap
	refc83953cc.max_cpus, cmax_cpus_allocs = (C.uint32_t)(x.max_cpus), cgoAllocsUnknown
	allocsc83953cc.Borrow(cmax_cpus_allocs)

	var cmax_nodes_allocs *cgoAllocMap
	refc83953cc.max_nodes, cmax_nodes_allocs = (C.uint32_t)(x.max_nodes), cgoAllocsUnknown
	allocsc83953cc.Borrow(cmax_nodes_allocs)

	var cmcs_label_allocs *cgoAllocMap
	refc83953cc.mcs_label, cmcs_label_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mcs_label)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cmcs_label_allocs)

	var cname_allocs *cgoAllocMap
	refc83953cc.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cname_allocs)

	var cnetwork_allocs *cgoAllocMap
	refc83953cc.network, cnetwork_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.network)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnetwork_allocs)

	var cnodes_allocs *cgoAllocMap
	refc83953cc.nodes, cnodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.nodes)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnodes_allocs)

	var cnice_allocs *cgoAllocMap
	refc83953cc.nice, cnice_allocs = (C.uint32_t)(x.nice), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnice_allocs)

	var cnode_inx_allocs *cgoAllocMap
	refc83953cc.node_inx, cnode_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_inx)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnode_inx_allocs)

	var cntasks_per_core_allocs *cgoAllocMap
	refc83953cc.ntasks_per_core, cntasks_per_core_allocs = (C.uint16_t)(x.ntasks_per_core), cgoAllocsUnknown
	allocsc83953cc.Borrow(cntasks_per_core_allocs)

	var cntasks_per_node_allocs *cgoAllocMap
	refc83953cc.ntasks_per_node, cntasks_per_node_allocs = (C.uint16_t)(x.ntasks_per_node), cgoAllocsUnknown
	allocsc83953cc.Borrow(cntasks_per_node_allocs)

	var cntasks_per_socket_allocs *cgoAllocMap
	refc83953cc.ntasks_per_socket, cntasks_per_socket_allocs = (C.uint16_t)(x.ntasks_per_socket), cgoAllocsUnknown
	allocsc83953cc.Borrow(cntasks_per_socket_allocs)

	var cntasks_per_board_allocs *cgoAllocMap
	refc83953cc.ntasks_per_board, cntasks_per_board_allocs = (C.uint16_t)(x.ntasks_per_board), cgoAllocsUnknown
	allocsc83953cc.Borrow(cntasks_per_board_allocs)

	var cnum_cpus_allocs *cgoAllocMap
	refc83953cc.num_cpus, cnum_cpus_allocs = (C.uint32_t)(x.num_cpus), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnum_cpus_allocs)

	var cnum_nodes_allocs *cgoAllocMap
	refc83953cc.num_nodes, cnum_nodes_allocs = (C.uint32_t)(x.num_nodes), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnum_nodes_allocs)

	var cnum_tasks_allocs *cgoAllocMap
	refc83953cc.num_tasks, cnum_tasks_allocs = (C.uint32_t)(x.num_tasks), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnum_tasks_allocs)

	var cpack_job_id_allocs *cgoAllocMap
	refc83953cc.pack_job_id, cpack_job_id_allocs = (C.uint32_t)(x.pack_job_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpack_job_id_allocs)

	var cpack_job_id_set_allocs *cgoAllocMap
	refc83953cc.pack_job_id_set, cpack_job_id_set_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.pack_job_id_set)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpack_job_id_set_allocs)

	var cpack_job_offset_allocs *cgoAllocMap
	refc83953cc.pack_job_offset, cpack_job_offset_allocs = (C.uint32_t)(x.pack_job_offset), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpack_job_offset_allocs)

	var cpartition_allocs *cgoAllocMap
	refc83953cc.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partition)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpartition_allocs)

	var cpn_min_memory_allocs *cgoAllocMap
	refc83953cc.pn_min_memory, cpn_min_memory_allocs = (C.uint64_t)(x.pn_min_memory), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpn_min_memory_allocs)

	var cpn_min_cpus_allocs *cgoAllocMap
	refc83953cc.pn_min_cpus, cpn_min_cpus_allocs = (C.uint16_t)(x.pn_min_cpus), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpn_min_cpus_allocs)

	var cpn_min_tmp_disk_allocs *cgoAllocMap
	refc83953cc.pn_min_tmp_disk, cpn_min_tmp_disk_allocs = (C.uint32_t)(x.pn_min_tmp_disk), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpn_min_tmp_disk_allocs)

	var cpower_flags_allocs *cgoAllocMap
	refc83953cc.power_flags, cpower_flags_allocs = (C.uint8_t)(x.power_flags), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpower_flags_allocs)

	var cpreempt_time_allocs *cgoAllocMap
	refc83953cc.preempt_time, cpreempt_time_allocs = (C.time_t)(x.preempt_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpreempt_time_allocs)

	var cpre_sus_time_allocs *cgoAllocMap
	refc83953cc.pre_sus_time, cpre_sus_time_allocs = (C.time_t)(x.pre_sus_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpre_sus_time_allocs)

	var cpriority_allocs *cgoAllocMap
	refc83953cc.priority, cpriority_allocs = (C.uint32_t)(x.priority), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpriority_allocs)

	var cprofile_allocs *cgoAllocMap
	refc83953cc.profile, cprofile_allocs = (C.uint32_t)(x.profile), cgoAllocsUnknown
	allocsc83953cc.Borrow(cprofile_allocs)

	var cqos_allocs *cgoAllocMap
	refc83953cc.qos, cqos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.qos)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cqos_allocs)

	var creboot_allocs *cgoAllocMap
	refc83953cc.reboot, creboot_allocs = (C.uint8_t)(x.reboot), cgoAllocsUnknown
	allocsc83953cc.Borrow(creboot_allocs)

	var creq_nodes_allocs *cgoAllocMap
	refc83953cc.req_nodes, creq_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.req_nodes)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(creq_nodes_allocs)

	var creq_node_inx_allocs *cgoAllocMap
	refc83953cc.req_node_inx, creq_node_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.req_node_inx)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(creq_node_inx_allocs)

	var creq_switch_allocs *cgoAllocMap
	refc83953cc.req_switch, creq_switch_allocs = (C.uint32_t)(x.req_switch), cgoAllocsUnknown
	allocsc83953cc.Borrow(creq_switch_allocs)

	var crequeue_allocs *cgoAllocMap
	refc83953cc.requeue, crequeue_allocs = (C.uint16_t)(x.requeue), cgoAllocsUnknown
	allocsc83953cc.Borrow(crequeue_allocs)

	var cresize_time_allocs *cgoAllocMap
	refc83953cc.resize_time, cresize_time_allocs = (C.time_t)(x.resize_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(cresize_time_allocs)

	var crestart_cnt_allocs *cgoAllocMap
	refc83953cc.restart_cnt, crestart_cnt_allocs = (C.uint16_t)(x.restart_cnt), cgoAllocsUnknown
	allocsc83953cc.Borrow(crestart_cnt_allocs)

	var cresv_name_allocs *cgoAllocMap
	refc83953cc.resv_name, cresv_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.resv_name)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cresv_name_allocs)

	var csched_nodes_allocs *cgoAllocMap
	refc83953cc.sched_nodes, csched_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.sched_nodes)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(csched_nodes_allocs)

	var cselect_jobinfo_allocs *cgoAllocMap
	refc83953cc.select_jobinfo, cselect_jobinfo_allocs = unpackSDynamic_plugin_data_t(x.select_jobinfo)
	allocsc83953cc.Borrow(cselect_jobinfo_allocs)

	var cshared_allocs *cgoAllocMap
	refc83953cc.shared, cshared_allocs = (C.uint16_t)(x.shared), cgoAllocsUnknown
	allocsc83953cc.Borrow(cshared_allocs)

	var cshow_flags_allocs *cgoAllocMap
	refc83953cc.show_flags, cshow_flags_allocs = (C.uint16_t)(x.show_flags), cgoAllocsUnknown
	allocsc83953cc.Borrow(cshow_flags_allocs)

	var csockets_per_board_allocs *cgoAllocMap
	refc83953cc.sockets_per_board, csockets_per_board_allocs = (C.uint16_t)(x.sockets_per_board), cgoAllocsUnknown
	allocsc83953cc.Borrow(csockets_per_board_allocs)

	var csockets_per_node_allocs *cgoAllocMap
	refc83953cc.sockets_per_node, csockets_per_node_allocs = (C.uint16_t)(x.sockets_per_node), cgoAllocsUnknown
	allocsc83953cc.Borrow(csockets_per_node_allocs)

	var cstart_time_allocs *cgoAllocMap
	refc83953cc.start_time, cstart_time_allocs = (C.time_t)(x.start_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstart_time_allocs)

	var cstart_protocol_ver_allocs *cgoAllocMap
	refc83953cc.start_protocol_ver, cstart_protocol_ver_allocs = (C.uint16_t)(x.start_protocol_ver), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstart_protocol_ver_allocs)

	var cstate_desc_allocs *cgoAllocMap
	refc83953cc.state_desc, cstate_desc_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.state_desc)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstate_desc_allocs)

	var cstate_reason_allocs *cgoAllocMap
	refc83953cc.state_reason, cstate_reason_allocs = (C.uint16_t)(x.state_reason), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstate_reason_allocs)

	var cstd_err_allocs *cgoAllocMap
	refc83953cc.std_err, cstd_err_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.std_err)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstd_err_allocs)

	var cstd_in_allocs *cgoAllocMap
	refc83953cc.std_in, cstd_in_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.std_in)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstd_in_allocs)

	var cstd_out_allocs *cgoAllocMap
	refc83953cc.std_out, cstd_out_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.std_out)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstd_out_allocs)

	var csubmit_time_allocs *cgoAllocMap
	refc83953cc.submit_time, csubmit_time_allocs = (C.time_t)(x.submit_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(csubmit_time_allocs)

	var csuspend_time_allocs *cgoAllocMap
	refc83953cc.suspend_time, csuspend_time_allocs = (C.time_t)(x.suspend_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(csuspend_time_allocs)

	var ctime_limit_allocs *cgoAllocMap
	refc83953cc.time_limit, ctime_limit_allocs = (C.uint32_t)(x.time_limit), cgoAllocsUnknown
	allocsc83953cc.Borrow(ctime_limit_allocs)

	var ctime_min_allocs *cgoAllocMap
	refc83953cc.time_min, ctime_min_allocs = (C.uint32_t)(x.time_min), cgoAllocsUnknown
	allocsc83953cc.Borrow(ctime_min_allocs)

	var cthreads_per_core_allocs *cgoAllocMap
	refc83953cc.threads_per_core, cthreads_per_core_allocs = (C.uint16_t)(x.threads_per_core), cgoAllocsUnknown
	allocsc83953cc.Borrow(cthreads_per_core_allocs)

	var ctres_req_str_allocs *cgoAllocMap
	refc83953cc.tres_req_str, ctres_req_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_req_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ctres_req_str_allocs)

	var ctres_alloc_str_allocs *cgoAllocMap
	refc83953cc.tres_alloc_str, ctres_alloc_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_alloc_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ctres_alloc_str_allocs)

	var cuser_id_allocs *cgoAllocMap
	refc83953cc.user_id, cuser_id_allocs = (C.uint32_t)(x.user_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(cuser_id_allocs)

	var cuser_name_allocs *cgoAllocMap
	refc83953cc.user_name, cuser_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.user_name)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cuser_name_allocs)

	var cwait4switch_allocs *cgoAllocMap
	refc83953cc.wait4switch, cwait4switch_allocs = (C.uint32_t)(x.wait4switch), cgoAllocsUnknown
	allocsc83953cc.Borrow(cwait4switch_allocs)

	var cwckey_allocs *cgoAllocMap
	refc83953cc.wckey, cwckey_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.wckey)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cwckey_allocs)

	var cwork_dir_allocs *cgoAllocMap
	refc83953cc.work_dir, cwork_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.work_dir)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cwork_dir_allocs)

	x.refc83953cc = refc83953cc
	x.allocsc83953cc = allocsc83953cc
	return refc83953cc, allocsc83953cc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_job_info_t) PassValue() (C.slurm_job_info_t, *cgoAllocMap) {
	if x.refc83953cc != nil {
		return *x.refc83953cc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_job_info_t) Deref() {
	if x.refc83953cc == nil {
		return
	}
	hxf8eae10 := (*sliceHeader)(unsafe.Pointer(&x.account))
	hxf8eae10.Data = uintptr(unsafe.Pointer(x.refc83953cc.account))
	hxf8eae10.Cap = 0x7fffffff
	// hxf8eae10.Len = ?

	hxfeb55cf := (*sliceHeader)(unsafe.Pointer(&x.admin_comment))
	hxfeb55cf.Data = uintptr(unsafe.Pointer(x.refc83953cc.admin_comment))
	hxfeb55cf.Cap = 0x7fffffff
	// hxfeb55cf.Len = ?

	hxf458096 := (*sliceHeader)(unsafe.Pointer(&x.alloc_node))
	hxf458096.Data = uintptr(unsafe.Pointer(x.refc83953cc.alloc_node))
	hxf458096.Cap = 0x7fffffff
	// hxf458096.Len = ?

	x.alloc_sid = (uint32_t)(x.refc83953cc.alloc_sid)
	x.array_bitmap = (unsafe.Pointer)(unsafe.Pointer(x.refc83953cc.array_bitmap))
	x.array_job_id = (uint32_t)(x.refc83953cc.array_job_id)
	x.array_task_id = (uint32_t)(x.refc83953cc.array_task_id)
	x.array_max_tasks = (uint32_t)(x.refc83953cc.array_max_tasks)
	hxf9aab83 := (*sliceHeader)(unsafe.Pointer(&x.array_task_str))
	hxf9aab83.Data = uintptr(unsafe.Pointer(x.refc83953cc.array_task_str))
	hxf9aab83.Cap = 0x7fffffff
	// hxf9aab83.Len = ?

	x.assoc_id = (uint32_t)(x.refc83953cc.assoc_id)
	x.batch_flag = (uint16_t)(x.refc83953cc.batch_flag)
	hxf8b35a8 := (*sliceHeader)(unsafe.Pointer(&x.batch_host))
	hxf8b35a8.Data = uintptr(unsafe.Pointer(x.refc83953cc.batch_host))
	hxf8b35a8.Cap = 0x7fffffff
	// hxf8b35a8.Len = ?

	x.bitflags = (uint32_t)(x.refc83953cc.bitflags)
	x.boards_per_node = (uint16_t)(x.refc83953cc.boards_per_node)
	hxf8959c2 := (*sliceHeader)(unsafe.Pointer(&x.burst_buffer))
	hxf8959c2.Data = uintptr(unsafe.Pointer(x.refc83953cc.burst_buffer))
	hxf8959c2.Cap = 0x7fffffff
	// hxf8959c2.Len = ?

	hxfb029a7 := (*sliceHeader)(unsafe.Pointer(&x.burst_buffer_state))
	hxfb029a7.Data = uintptr(unsafe.Pointer(x.refc83953cc.burst_buffer_state))
	hxfb029a7.Cap = 0x7fffffff
	// hxfb029a7.Len = ?

	hxf7d15a2 := (*sliceHeader)(unsafe.Pointer(&x.cluster))
	hxf7d15a2.Data = uintptr(unsafe.Pointer(x.refc83953cc.cluster))
	hxf7d15a2.Cap = 0x7fffffff
	// hxf7d15a2.Len = ?

	hxf8dbbe5 := (*sliceHeader)(unsafe.Pointer(&x.cluster_features))
	hxf8dbbe5.Data = uintptr(unsafe.Pointer(x.refc83953cc.cluster_features))
	hxf8dbbe5.Cap = 0x7fffffff
	// hxf8dbbe5.Len = ?

	hxf766ff8 := (*sliceHeader)(unsafe.Pointer(&x.command))
	hxf766ff8.Data = uintptr(unsafe.Pointer(x.refc83953cc.command))
	hxf766ff8.Cap = 0x7fffffff
	// hxf766ff8.Len = ?

	hxf9b1633 := (*sliceHeader)(unsafe.Pointer(&x.comment))
	hxf9b1633.Data = uintptr(unsafe.Pointer(x.refc83953cc.comment))
	hxf9b1633.Cap = 0x7fffffff
	// hxf9b1633.Len = ?

	x.contiguous = (uint16_t)(x.refc83953cc.contiguous)
	x.core_spec = (uint16_t)(x.refc83953cc.core_spec)
	x.cores_per_socket = (uint16_t)(x.refc83953cc.cores_per_socket)
	x.billable_tres = (float64)(x.refc83953cc.billable_tres)
	x.cpus_per_task = (uint16_t)(x.refc83953cc.cpus_per_task)
	x.cpu_freq_min = (uint32_t)(x.refc83953cc.cpu_freq_min)
	x.cpu_freq_max = (uint32_t)(x.refc83953cc.cpu_freq_max)
	x.cpu_freq_gov = (uint32_t)(x.refc83953cc.cpu_freq_gov)
	x.deadline = (time_t)(x.refc83953cc.deadline)
	x.delay_boot = (uint32_t)(x.refc83953cc.delay_boot)
	hxf502c9a := (*sliceHeader)(unsafe.Pointer(&x.dependency))
	hxf502c9a.Data = uintptr(unsafe.Pointer(x.refc83953cc.dependency))
	hxf502c9a.Cap = 0x7fffffff
	// hxf502c9a.Len = ?

	x.derived_ec = (uint32_t)(x.refc83953cc.derived_ec)
	x.eligible_time = (time_t)(x.refc83953cc.eligible_time)
	x.end_time = (time_t)(x.refc83953cc.end_time)
	hxf4a9453 := (*sliceHeader)(unsafe.Pointer(&x.exc_nodes))
	hxf4a9453.Data = uintptr(unsafe.Pointer(x.refc83953cc.exc_nodes))
	hxf4a9453.Cap = 0x7fffffff
	// hxf4a9453.Len = ?

	hxf1a1416 := (*sliceHeader)(unsafe.Pointer(&x.exc_node_inx))
	hxf1a1416.Data = uintptr(unsafe.Pointer(x.refc83953cc.exc_node_inx))
	hxf1a1416.Cap = 0x7fffffff
	// hxf1a1416.Len = ?

	x.exit_code = (uint32_t)(x.refc83953cc.exit_code)
	hxf92be66 := (*sliceHeader)(unsafe.Pointer(&x.features))
	hxf92be66.Data = uintptr(unsafe.Pointer(x.refc83953cc.features))
	hxf92be66.Cap = 0x7fffffff
	// hxf92be66.Len = ?

	hxf4b5187 := (*sliceHeader)(unsafe.Pointer(&x.fed_origin_str))
	hxf4b5187.Data = uintptr(unsafe.Pointer(x.refc83953cc.fed_origin_str))
	hxf4b5187.Cap = 0x7fffffff
	// hxf4b5187.Len = ?

	x.fed_siblings_active = (uint64_t)(x.refc83953cc.fed_siblings_active)
	hxf177f79 := (*sliceHeader)(unsafe.Pointer(&x.fed_siblings_active_str))
	hxf177f79.Data = uintptr(unsafe.Pointer(x.refc83953cc.fed_siblings_active_str))
	hxf177f79.Cap = 0x7fffffff
	// hxf177f79.Len = ?

	x.fed_siblings_viable = (uint64_t)(x.refc83953cc.fed_siblings_viable)
	hxfaa359c := (*sliceHeader)(unsafe.Pointer(&x.fed_siblings_viable_str))
	hxfaa359c.Data = uintptr(unsafe.Pointer(x.refc83953cc.fed_siblings_viable_str))
	hxfaa359c.Cap = 0x7fffffff
	// hxfaa359c.Len = ?

	hxfa897de := (*sliceHeader)(unsafe.Pointer(&x.gres))
	hxfa897de.Data = uintptr(unsafe.Pointer(x.refc83953cc.gres))
	hxfa897de.Cap = 0x7fffffff
	// hxfa897de.Len = ?

	x.gres_detail_cnt = (uint32_t)(x.refc83953cc.gres_detail_cnt)
	packSSByte(x.gres_detail_str, x.refc83953cc.gres_detail_str)
	x.group_id = (uint32_t)(x.refc83953cc.group_id)
	x.job_id = (uint32_t)(x.refc83953cc.job_id)
	hxf08bba9 := (*sliceHeader)(unsafe.Pointer(&x.job_resrcs))
	hxf08bba9.Data = uintptr(unsafe.Pointer(x.refc83953cc.job_resrcs))
	hxf08bba9.Cap = 0x7fffffff
	// hxf08bba9.Len = ?

	x.job_state = (uint32_t)(x.refc83953cc.job_state)
	x.last_sched_eval = (time_t)(x.refc83953cc.last_sched_eval)
	hxfd3aa9c := (*sliceHeader)(unsafe.Pointer(&x.licenses))
	hxfd3aa9c.Data = uintptr(unsafe.Pointer(x.refc83953cc.licenses))
	hxfd3aa9c.Cap = 0x7fffffff
	// hxfd3aa9c.Len = ?

	x.max_cpus = (uint32_t)(x.refc83953cc.max_cpus)
	x.max_nodes = (uint32_t)(x.refc83953cc.max_nodes)
	hxfb2f596 := (*sliceHeader)(unsafe.Pointer(&x.mcs_label))
	hxfb2f596.Data = uintptr(unsafe.Pointer(x.refc83953cc.mcs_label))
	hxfb2f596.Cap = 0x7fffffff
	// hxfb2f596.Len = ?

	hxf11683e := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf11683e.Data = uintptr(unsafe.Pointer(x.refc83953cc.name))
	hxf11683e.Cap = 0x7fffffff
	// hxf11683e.Len = ?

	hxfd9261b := (*sliceHeader)(unsafe.Pointer(&x.network))
	hxfd9261b.Data = uintptr(unsafe.Pointer(x.refc83953cc.network))
	hxfd9261b.Cap = 0x7fffffff
	// hxfd9261b.Len = ?

	hxf77d2ac := (*sliceHeader)(unsafe.Pointer(&x.nodes))
	hxf77d2ac.Data = uintptr(unsafe.Pointer(x.refc83953cc.nodes))
	hxf77d2ac.Cap = 0x7fffffff
	// hxf77d2ac.Len = ?

	x.nice = (uint32_t)(x.refc83953cc.nice)
	hxff6a91e := (*sliceHeader)(unsafe.Pointer(&x.node_inx))
	hxff6a91e.Data = uintptr(unsafe.Pointer(x.refc83953cc.node_inx))
	hxff6a91e.Cap = 0x7fffffff
	// hxff6a91e.Len = ?

	x.ntasks_per_core = (uint16_t)(x.refc83953cc.ntasks_per_core)
	x.ntasks_per_node = (uint16_t)(x.refc83953cc.ntasks_per_node)
	x.ntasks_per_socket = (uint16_t)(x.refc83953cc.ntasks_per_socket)
	x.ntasks_per_board = (uint16_t)(x.refc83953cc.ntasks_per_board)
	x.num_cpus = (uint32_t)(x.refc83953cc.num_cpus)
	x.num_nodes = (uint32_t)(x.refc83953cc.num_nodes)
	x.num_tasks = (uint32_t)(x.refc83953cc.num_tasks)
	x.pack_job_id = (uint32_t)(x.refc83953cc.pack_job_id)
	hxf971c70 := (*sliceHeader)(unsafe.Pointer(&x.pack_job_id_set))
	hxf971c70.Data = uintptr(unsafe.Pointer(x.refc83953cc.pack_job_id_set))
	hxf971c70.Cap = 0x7fffffff
	// hxf971c70.Len = ?

	x.pack_job_offset = (uint32_t)(x.refc83953cc.pack_job_offset)
	hxf047235 := (*sliceHeader)(unsafe.Pointer(&x.partition))
	hxf047235.Data = uintptr(unsafe.Pointer(x.refc83953cc.partition))
	hxf047235.Cap = 0x7fffffff
	// hxf047235.Len = ?

	x.pn_min_memory = (uint64_t)(x.refc83953cc.pn_min_memory)
	x.pn_min_cpus = (uint16_t)(x.refc83953cc.pn_min_cpus)
	x.pn_min_tmp_disk = (uint32_t)(x.refc83953cc.pn_min_tmp_disk)
	x.power_flags = (uint8_t)(x.refc83953cc.power_flags)
	x.preempt_time = (time_t)(x.refc83953cc.preempt_time)
	x.pre_sus_time = (time_t)(x.refc83953cc.pre_sus_time)
	x.priority = (uint32_t)(x.refc83953cc.priority)
	x.profile = (uint32_t)(x.refc83953cc.profile)
	hxf612a5d := (*sliceHeader)(unsafe.Pointer(&x.qos))
	hxf612a5d.Data = uintptr(unsafe.Pointer(x.refc83953cc.qos))
	hxf612a5d.Cap = 0x7fffffff
	// hxf612a5d.Len = ?

	x.reboot = (uint8_t)(x.refc83953cc.reboot)
	hxff58be3 := (*sliceHeader)(unsafe.Pointer(&x.req_nodes))
	hxff58be3.Data = uintptr(unsafe.Pointer(x.refc83953cc.req_nodes))
	hxff58be3.Cap = 0x7fffffff
	// hxff58be3.Len = ?

	hxf79f3d5 := (*sliceHeader)(unsafe.Pointer(&x.req_node_inx))
	hxf79f3d5.Data = uintptr(unsafe.Pointer(x.refc83953cc.req_node_inx))
	hxf79f3d5.Cap = 0x7fffffff
	// hxf79f3d5.Len = ?

	x.req_switch = (uint32_t)(x.refc83953cc.req_switch)
	x.requeue = (uint16_t)(x.refc83953cc.requeue)
	x.resize_time = (time_t)(x.refc83953cc.resize_time)
	x.restart_cnt = (uint16_t)(x.refc83953cc.restart_cnt)
	hxfbb2d22 := (*sliceHeader)(unsafe.Pointer(&x.resv_name))
	hxfbb2d22.Data = uintptr(unsafe.Pointer(x.refc83953cc.resv_name))
	hxfbb2d22.Cap = 0x7fffffff
	// hxfbb2d22.Len = ?

	hxff3831d := (*sliceHeader)(unsafe.Pointer(&x.sched_nodes))
	hxff3831d.Data = uintptr(unsafe.Pointer(x.refc83953cc.sched_nodes))
	hxff3831d.Cap = 0x7fffffff
	// hxff3831d.Len = ?

	packSDynamic_plugin_data_t(x.select_jobinfo, x.refc83953cc.select_jobinfo)
	x.shared = (uint16_t)(x.refc83953cc.shared)
	x.show_flags = (uint16_t)(x.refc83953cc.show_flags)
	x.sockets_per_board = (uint16_t)(x.refc83953cc.sockets_per_board)
	x.sockets_per_node = (uint16_t)(x.refc83953cc.sockets_per_node)
	x.start_time = (time_t)(x.refc83953cc.start_time)
	x.start_protocol_ver = (uint16_t)(x.refc83953cc.start_protocol_ver)
	hxf3b2498 := (*sliceHeader)(unsafe.Pointer(&x.state_desc))
	hxf3b2498.Data = uintptr(unsafe.Pointer(x.refc83953cc.state_desc))
	hxf3b2498.Cap = 0x7fffffff
	// hxf3b2498.Len = ?

	x.state_reason = (uint16_t)(x.refc83953cc.state_reason)
	hxfe80fc5 := (*sliceHeader)(unsafe.Pointer(&x.std_err))
	hxfe80fc5.Data = uintptr(unsafe.Pointer(x.refc83953cc.std_err))
	hxfe80fc5.Cap = 0x7fffffff
	// hxfe80fc5.Len = ?

	hxfc03c44 := (*sliceHeader)(unsafe.Pointer(&x.std_in))
	hxfc03c44.Data = uintptr(unsafe.Pointer(x.refc83953cc.std_in))
	hxfc03c44.Cap = 0x7fffffff
	// hxfc03c44.Len = ?

	hxf32b611 := (*sliceHeader)(unsafe.Pointer(&x.std_out))
	hxf32b611.Data = uintptr(unsafe.Pointer(x.refc83953cc.std_out))
	hxf32b611.Cap = 0x7fffffff
	// hxf32b611.Len = ?

	x.submit_time = (time_t)(x.refc83953cc.submit_time)
	x.suspend_time = (time_t)(x.refc83953cc.suspend_time)
	x.time_limit = (uint32_t)(x.refc83953cc.time_limit)
	x.time_min = (uint32_t)(x.refc83953cc.time_min)
	x.threads_per_core = (uint16_t)(x.refc83953cc.threads_per_core)
	hxf187e71 := (*sliceHeader)(unsafe.Pointer(&x.tres_req_str))
	hxf187e71.Data = uintptr(unsafe.Pointer(x.refc83953cc.tres_req_str))
	hxf187e71.Cap = 0x7fffffff
	// hxf187e71.Len = ?

	hxf13e50a := (*sliceHeader)(unsafe.Pointer(&x.tres_alloc_str))
	hxf13e50a.Data = uintptr(unsafe.Pointer(x.refc83953cc.tres_alloc_str))
	hxf13e50a.Cap = 0x7fffffff
	// hxf13e50a.Len = ?

	x.user_id = (uint32_t)(x.refc83953cc.user_id)
	hxf9227c2 := (*sliceHeader)(unsafe.Pointer(&x.user_name))
	hxf9227c2.Data = uintptr(unsafe.Pointer(x.refc83953cc.user_name))
	hxf9227c2.Cap = 0x7fffffff
	// hxf9227c2.Len = ?

	x.wait4switch = (uint32_t)(x.refc83953cc.wait4switch)
	hxfa7b2ab := (*sliceHeader)(unsafe.Pointer(&x.wckey))
	hxfa7b2ab.Data = uintptr(unsafe.Pointer(x.refc83953cc.wckey))
	hxfa7b2ab.Cap = 0x7fffffff
	// hxfa7b2ab.Len = ?

	hxfa34053 := (*sliceHeader)(unsafe.Pointer(&x.work_dir))
	hxfa34053.Data = uintptr(unsafe.Pointer(x.refc83953cc.work_dir))
	hxfa34053.Cap = 0x7fffffff
	// hxfa34053.Len = ?

}

// allocPriority_factors_object_tMemory allocates memory for type C.priority_factors_object_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPriority_factors_object_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPriority_factors_object_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPriority_factors_object_tValue = unsafe.Sizeof([1]C.priority_factors_object_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *priority_factors_object_t) Ref() *C.priority_factors_object_t {
	if x == nil {
		return nil
	}
	return x.ref5b40eea1
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *priority_factors_object_t) Free() {
	if x != nil && x.allocs5b40eea1 != nil {
		x.allocs5b40eea1.(*cgoAllocMap).Free()
		x.ref5b40eea1 = nil
	}
}

// Newpriority_factors_object_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newpriority_factors_object_tRef(ref unsafe.Pointer) *priority_factors_object_t {
	if ref == nil {
		return nil
	}
	obj := new(priority_factors_object_t)
	obj.ref5b40eea1 = (*C.priority_factors_object_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *priority_factors_object_t) PassRef() (*C.priority_factors_object_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5b40eea1 != nil {
		return x.ref5b40eea1, nil
	}
	mem5b40eea1 := allocPriority_factors_object_tMemory(1)
	ref5b40eea1 := (*C.priority_factors_object_t)(mem5b40eea1)
	allocs5b40eea1 := new(cgoAllocMap)
	allocs5b40eea1.Add(mem5b40eea1)

	var ccluster_name_allocs *cgoAllocMap
	ref5b40eea1.cluster_name, ccluster_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster_name)).Data)), cgoAllocsUnknown
	allocs5b40eea1.Borrow(ccluster_name_allocs)

	var cjob_id_allocs *cgoAllocMap
	ref5b40eea1.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs5b40eea1.Borrow(cjob_id_allocs)

	var cpartition_allocs *cgoAllocMap
	ref5b40eea1.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partition)).Data)), cgoAllocsUnknown
	allocs5b40eea1.Borrow(cpartition_allocs)

	var cuser_id_allocs *cgoAllocMap
	ref5b40eea1.user_id, cuser_id_allocs = (C.uint32_t)(x.user_id), cgoAllocsUnknown
	allocs5b40eea1.Borrow(cuser_id_allocs)

	var cpriority_age_allocs *cgoAllocMap
	ref5b40eea1.priority_age, cpriority_age_allocs = (C.double)(x.priority_age), cgoAllocsUnknown
	allocs5b40eea1.Borrow(cpriority_age_allocs)

	var cpriority_fs_allocs *cgoAllocMap
	ref5b40eea1.priority_fs, cpriority_fs_allocs = (C.double)(x.priority_fs), cgoAllocsUnknown
	allocs5b40eea1.Borrow(cpriority_fs_allocs)

	var cpriority_js_allocs *cgoAllocMap
	ref5b40eea1.priority_js, cpriority_js_allocs = (C.double)(x.priority_js), cgoAllocsUnknown
	allocs5b40eea1.Borrow(cpriority_js_allocs)

	var cpriority_part_allocs *cgoAllocMap
	ref5b40eea1.priority_part, cpriority_part_allocs = (C.double)(x.priority_part), cgoAllocsUnknown
	allocs5b40eea1.Borrow(cpriority_part_allocs)

	var cpriority_qos_allocs *cgoAllocMap
	ref5b40eea1.priority_qos, cpriority_qos_allocs = (C.double)(x.priority_qos), cgoAllocsUnknown
	allocs5b40eea1.Borrow(cpriority_qos_allocs)

	var cpriority_tres_allocs *cgoAllocMap
	ref5b40eea1.priority_tres, cpriority_tres_allocs = (*C.double)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.priority_tres)).Data)), cgoAllocsUnknown
	allocs5b40eea1.Borrow(cpriority_tres_allocs)

	var ctres_cnt_allocs *cgoAllocMap
	ref5b40eea1.tres_cnt, ctres_cnt_allocs = (C.uint32_t)(x.tres_cnt), cgoAllocsUnknown
	allocs5b40eea1.Borrow(ctres_cnt_allocs)

	var ctres_names_allocs *cgoAllocMap
	ref5b40eea1.tres_names, ctres_names_allocs = unpackSSByte(x.tres_names)
	allocs5b40eea1.Borrow(ctres_names_allocs)

	var ctres_weights_allocs *cgoAllocMap
	ref5b40eea1.tres_weights, ctres_weights_allocs = (*C.double)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_weights)).Data)), cgoAllocsUnknown
	allocs5b40eea1.Borrow(ctres_weights_allocs)

	var cnice_allocs *cgoAllocMap
	ref5b40eea1.nice, cnice_allocs = (C.uint32_t)(x.nice), cgoAllocsUnknown
	allocs5b40eea1.Borrow(cnice_allocs)

	x.ref5b40eea1 = ref5b40eea1
	x.allocs5b40eea1 = allocs5b40eea1
	return ref5b40eea1, allocs5b40eea1

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x priority_factors_object_t) PassValue() (C.priority_factors_object_t, *cgoAllocMap) {
	if x.ref5b40eea1 != nil {
		return *x.ref5b40eea1, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *priority_factors_object_t) Deref() {
	if x.ref5b40eea1 == nil {
		return
	}
	hxfe3c193 := (*sliceHeader)(unsafe.Pointer(&x.cluster_name))
	hxfe3c193.Data = uintptr(unsafe.Pointer(x.ref5b40eea1.cluster_name))
	hxfe3c193.Cap = 0x7fffffff
	// hxfe3c193.Len = ?

	x.job_id = (uint32_t)(x.ref5b40eea1.job_id)
	hxfe8267c := (*sliceHeader)(unsafe.Pointer(&x.partition))
	hxfe8267c.Data = uintptr(unsafe.Pointer(x.ref5b40eea1.partition))
	hxfe8267c.Cap = 0x7fffffff
	// hxfe8267c.Len = ?

	x.user_id = (uint32_t)(x.ref5b40eea1.user_id)
	x.priority_age = (float64)(x.ref5b40eea1.priority_age)
	x.priority_fs = (float64)(x.ref5b40eea1.priority_fs)
	x.priority_js = (float64)(x.ref5b40eea1.priority_js)
	x.priority_part = (float64)(x.ref5b40eea1.priority_part)
	x.priority_qos = (float64)(x.ref5b40eea1.priority_qos)
	hxf9143ee := (*sliceHeader)(unsafe.Pointer(&x.priority_tres))
	hxf9143ee.Data = uintptr(unsafe.Pointer(x.ref5b40eea1.priority_tres))
	hxf9143ee.Cap = 0x7fffffff
	// hxf9143ee.Len = ?

	x.tres_cnt = (uint32_t)(x.ref5b40eea1.tres_cnt)
	packSSByte(x.tres_names, x.ref5b40eea1.tres_names)
	hxf349858 := (*sliceHeader)(unsafe.Pointer(&x.tres_weights))
	hxf349858.Data = uintptr(unsafe.Pointer(x.ref5b40eea1.tres_weights))
	hxf349858.Cap = 0x7fffffff
	// hxf349858.Len = ?

	x.nice = (uint32_t)(x.ref5b40eea1.nice)
}

// allocPriority_factors_response_msg_tMemory allocates memory for type C.priority_factors_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPriority_factors_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPriority_factors_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPriority_factors_response_msg_tValue = unsafe.Sizeof([1]C.priority_factors_response_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *priority_factors_response_msg_t) Ref() *C.priority_factors_response_msg_t {
	if x == nil {
		return nil
	}
	return x.ref63319104
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *priority_factors_response_msg_t) Free() {
	if x != nil && x.allocs63319104 != nil {
		x.allocs63319104.(*cgoAllocMap).Free()
		x.ref63319104 = nil
	}
}

// Newpriority_factors_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newpriority_factors_response_msg_tRef(ref unsafe.Pointer) *priority_factors_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(priority_factors_response_msg_t)
	obj.ref63319104 = (*C.priority_factors_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *priority_factors_response_msg_t) PassRef() (*C.priority_factors_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref63319104 != nil {
		return x.ref63319104, nil
	}
	mem63319104 := allocPriority_factors_response_msg_tMemory(1)
	ref63319104 := (*C.priority_factors_response_msg_t)(mem63319104)
	allocs63319104 := new(cgoAllocMap)
	allocs63319104.Add(mem63319104)

	var cpriority_factors_list_allocs *cgoAllocMap
	ref63319104.priority_factors_list, cpriority_factors_list_allocs = *(*C.List)(unsafe.Pointer(&x.priority_factors_list)), cgoAllocsUnknown
	allocs63319104.Borrow(cpriority_factors_list_allocs)

	x.ref63319104 = ref63319104
	x.allocs63319104 = allocs63319104
	return ref63319104, allocs63319104

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x priority_factors_response_msg_t) PassValue() (C.priority_factors_response_msg_t, *cgoAllocMap) {
	if x.ref63319104 != nil {
		return *x.ref63319104, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *priority_factors_response_msg_t) Deref() {
	if x.ref63319104 == nil {
		return
	}
	x.priority_factors_list = *(*List)(unsafe.Pointer(&x.ref63319104.priority_factors_list))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_info_t) Ref() *C.slurm_job_info_t {
	if x == nil {
		return nil
	}
	return x.refc83953cc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_info_t) Free() {
	if x != nil && x.allocsc83953cc != nil {
		x.allocsc83953cc.(*cgoAllocMap).Free()
		x.refc83953cc = nil
	}
}

// Newjob_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_info_tRef(ref unsafe.Pointer) *job_info_t {
	if ref == nil {
		return nil
	}
	obj := new(job_info_t)
	obj.refc83953cc = (*C.slurm_job_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_info_t) PassRef() (*C.slurm_job_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc83953cc != nil {
		return x.refc83953cc, nil
	}
	memc83953cc := allocSlurm_job_info_tMemory(1)
	refc83953cc := (*C.slurm_job_info_t)(memc83953cc)
	allocsc83953cc := new(cgoAllocMap)
	allocsc83953cc.Add(memc83953cc)

	var caccount_allocs *cgoAllocMap
	refc83953cc.account, caccount_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.account)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(caccount_allocs)

	var cadmin_comment_allocs *cgoAllocMap
	refc83953cc.admin_comment, cadmin_comment_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.admin_comment)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cadmin_comment_allocs)

	var calloc_node_allocs *cgoAllocMap
	refc83953cc.alloc_node, calloc_node_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.alloc_node)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(calloc_node_allocs)

	var calloc_sid_allocs *cgoAllocMap
	refc83953cc.alloc_sid, calloc_sid_allocs = (C.uint32_t)(x.alloc_sid), cgoAllocsUnknown
	allocsc83953cc.Borrow(calloc_sid_allocs)

	var carray_bitmap_allocs *cgoAllocMap
	refc83953cc.array_bitmap, carray_bitmap_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.array_bitmap)), cgoAllocsUnknown
	allocsc83953cc.Borrow(carray_bitmap_allocs)

	var carray_job_id_allocs *cgoAllocMap
	refc83953cc.array_job_id, carray_job_id_allocs = (C.uint32_t)(x.array_job_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(carray_job_id_allocs)

	var carray_task_id_allocs *cgoAllocMap
	refc83953cc.array_task_id, carray_task_id_allocs = (C.uint32_t)(x.array_task_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(carray_task_id_allocs)

	var carray_max_tasks_allocs *cgoAllocMap
	refc83953cc.array_max_tasks, carray_max_tasks_allocs = (C.uint32_t)(x.array_max_tasks), cgoAllocsUnknown
	allocsc83953cc.Borrow(carray_max_tasks_allocs)

	var carray_task_str_allocs *cgoAllocMap
	refc83953cc.array_task_str, carray_task_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.array_task_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(carray_task_str_allocs)

	var cassoc_id_allocs *cgoAllocMap
	refc83953cc.assoc_id, cassoc_id_allocs = (C.uint32_t)(x.assoc_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(cassoc_id_allocs)

	var cbatch_flag_allocs *cgoAllocMap
	refc83953cc.batch_flag, cbatch_flag_allocs = (C.uint16_t)(x.batch_flag), cgoAllocsUnknown
	allocsc83953cc.Borrow(cbatch_flag_allocs)

	var cbatch_host_allocs *cgoAllocMap
	refc83953cc.batch_host, cbatch_host_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.batch_host)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cbatch_host_allocs)

	var cbitflags_allocs *cgoAllocMap
	refc83953cc.bitflags, cbitflags_allocs = (C.uint32_t)(x.bitflags), cgoAllocsUnknown
	allocsc83953cc.Borrow(cbitflags_allocs)

	var cboards_per_node_allocs *cgoAllocMap
	refc83953cc.boards_per_node, cboards_per_node_allocs = (C.uint16_t)(x.boards_per_node), cgoAllocsUnknown
	allocsc83953cc.Borrow(cboards_per_node_allocs)

	var cburst_buffer_allocs *cgoAllocMap
	refc83953cc.burst_buffer, cburst_buffer_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.burst_buffer)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cburst_buffer_allocs)

	var cburst_buffer_state_allocs *cgoAllocMap
	refc83953cc.burst_buffer_state, cburst_buffer_state_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.burst_buffer_state)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cburst_buffer_state_allocs)

	var ccluster_allocs *cgoAllocMap
	refc83953cc.cluster, ccluster_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccluster_allocs)

	var ccluster_features_allocs *cgoAllocMap
	refc83953cc.cluster_features, ccluster_features_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster_features)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccluster_features_allocs)

	var ccommand_allocs *cgoAllocMap
	refc83953cc.command, ccommand_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.command)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccommand_allocs)

	var ccomment_allocs *cgoAllocMap
	refc83953cc.comment, ccomment_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.comment)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccomment_allocs)

	var ccontiguous_allocs *cgoAllocMap
	refc83953cc.contiguous, ccontiguous_allocs = (C.uint16_t)(x.contiguous), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccontiguous_allocs)

	var ccore_spec_allocs *cgoAllocMap
	refc83953cc.core_spec, ccore_spec_allocs = (C.uint16_t)(x.core_spec), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccore_spec_allocs)

	var ccores_per_socket_allocs *cgoAllocMap
	refc83953cc.cores_per_socket, ccores_per_socket_allocs = (C.uint16_t)(x.cores_per_socket), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccores_per_socket_allocs)

	var cbillable_tres_allocs *cgoAllocMap
	refc83953cc.billable_tres, cbillable_tres_allocs = (C.double)(x.billable_tres), cgoAllocsUnknown
	allocsc83953cc.Borrow(cbillable_tres_allocs)

	var ccpus_per_task_allocs *cgoAllocMap
	refc83953cc.cpus_per_task, ccpus_per_task_allocs = (C.uint16_t)(x.cpus_per_task), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccpus_per_task_allocs)

	var ccpu_freq_min_allocs *cgoAllocMap
	refc83953cc.cpu_freq_min, ccpu_freq_min_allocs = (C.uint32_t)(x.cpu_freq_min), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccpu_freq_min_allocs)

	var ccpu_freq_max_allocs *cgoAllocMap
	refc83953cc.cpu_freq_max, ccpu_freq_max_allocs = (C.uint32_t)(x.cpu_freq_max), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccpu_freq_max_allocs)

	var ccpu_freq_gov_allocs *cgoAllocMap
	refc83953cc.cpu_freq_gov, ccpu_freq_gov_allocs = (C.uint32_t)(x.cpu_freq_gov), cgoAllocsUnknown
	allocsc83953cc.Borrow(ccpu_freq_gov_allocs)

	var cdeadline_allocs *cgoAllocMap
	refc83953cc.deadline, cdeadline_allocs = (C.time_t)(x.deadline), cgoAllocsUnknown
	allocsc83953cc.Borrow(cdeadline_allocs)

	var cdelay_boot_allocs *cgoAllocMap
	refc83953cc.delay_boot, cdelay_boot_allocs = (C.uint32_t)(x.delay_boot), cgoAllocsUnknown
	allocsc83953cc.Borrow(cdelay_boot_allocs)

	var cdependency_allocs *cgoAllocMap
	refc83953cc.dependency, cdependency_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.dependency)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cdependency_allocs)

	var cderived_ec_allocs *cgoAllocMap
	refc83953cc.derived_ec, cderived_ec_allocs = (C.uint32_t)(x.derived_ec), cgoAllocsUnknown
	allocsc83953cc.Borrow(cderived_ec_allocs)

	var celigible_time_allocs *cgoAllocMap
	refc83953cc.eligible_time, celigible_time_allocs = (C.time_t)(x.eligible_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(celigible_time_allocs)

	var cend_time_allocs *cgoAllocMap
	refc83953cc.end_time, cend_time_allocs = (C.time_t)(x.end_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(cend_time_allocs)

	var cexc_nodes_allocs *cgoAllocMap
	refc83953cc.exc_nodes, cexc_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.exc_nodes)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cexc_nodes_allocs)

	var cexc_node_inx_allocs *cgoAllocMap
	refc83953cc.exc_node_inx, cexc_node_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.exc_node_inx)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cexc_node_inx_allocs)

	var cexit_code_allocs *cgoAllocMap
	refc83953cc.exit_code, cexit_code_allocs = (C.uint32_t)(x.exit_code), cgoAllocsUnknown
	allocsc83953cc.Borrow(cexit_code_allocs)

	var cfeatures_allocs *cgoAllocMap
	refc83953cc.features, cfeatures_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.features)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfeatures_allocs)

	var cfed_origin_str_allocs *cgoAllocMap
	refc83953cc.fed_origin_str, cfed_origin_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.fed_origin_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfed_origin_str_allocs)

	var cfed_siblings_active_allocs *cgoAllocMap
	refc83953cc.fed_siblings_active, cfed_siblings_active_allocs = (C.uint64_t)(x.fed_siblings_active), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfed_siblings_active_allocs)

	var cfed_siblings_active_str_allocs *cgoAllocMap
	refc83953cc.fed_siblings_active_str, cfed_siblings_active_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.fed_siblings_active_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfed_siblings_active_str_allocs)

	var cfed_siblings_viable_allocs *cgoAllocMap
	refc83953cc.fed_siblings_viable, cfed_siblings_viable_allocs = (C.uint64_t)(x.fed_siblings_viable), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfed_siblings_viable_allocs)

	var cfed_siblings_viable_str_allocs *cgoAllocMap
	refc83953cc.fed_siblings_viable_str, cfed_siblings_viable_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.fed_siblings_viable_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cfed_siblings_viable_str_allocs)

	var cgres_allocs *cgoAllocMap
	refc83953cc.gres, cgres_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.gres)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cgres_allocs)

	var cgres_detail_cnt_allocs *cgoAllocMap
	refc83953cc.gres_detail_cnt, cgres_detail_cnt_allocs = (C.uint32_t)(x.gres_detail_cnt), cgoAllocsUnknown
	allocsc83953cc.Borrow(cgres_detail_cnt_allocs)

	var cgres_detail_str_allocs *cgoAllocMap
	refc83953cc.gres_detail_str, cgres_detail_str_allocs = unpackSSByte(x.gres_detail_str)
	allocsc83953cc.Borrow(cgres_detail_str_allocs)

	var cgroup_id_allocs *cgoAllocMap
	refc83953cc.group_id, cgroup_id_allocs = (C.uint32_t)(x.group_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(cgroup_id_allocs)

	var cjob_id_allocs *cgoAllocMap
	refc83953cc.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(cjob_id_allocs)

	var cjob_resrcs_allocs *cgoAllocMap
	refc83953cc.job_resrcs, cjob_resrcs_allocs = (*C.job_resources_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_resrcs)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cjob_resrcs_allocs)

	var cjob_state_allocs *cgoAllocMap
	refc83953cc.job_state, cjob_state_allocs = (C.uint32_t)(x.job_state), cgoAllocsUnknown
	allocsc83953cc.Borrow(cjob_state_allocs)

	var clast_sched_eval_allocs *cgoAllocMap
	refc83953cc.last_sched_eval, clast_sched_eval_allocs = (C.time_t)(x.last_sched_eval), cgoAllocsUnknown
	allocsc83953cc.Borrow(clast_sched_eval_allocs)

	var clicenses_allocs *cgoAllocMap
	refc83953cc.licenses, clicenses_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.licenses)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(clicenses_allocs)

	var cmax_cpus_allocs *cgoAllocMap
	refc83953cc.max_cpus, cmax_cpus_allocs = (C.uint32_t)(x.max_cpus), cgoAllocsUnknown
	allocsc83953cc.Borrow(cmax_cpus_allocs)

	var cmax_nodes_allocs *cgoAllocMap
	refc83953cc.max_nodes, cmax_nodes_allocs = (C.uint32_t)(x.max_nodes), cgoAllocsUnknown
	allocsc83953cc.Borrow(cmax_nodes_allocs)

	var cmcs_label_allocs *cgoAllocMap
	refc83953cc.mcs_label, cmcs_label_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mcs_label)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cmcs_label_allocs)

	var cname_allocs *cgoAllocMap
	refc83953cc.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cname_allocs)

	var cnetwork_allocs *cgoAllocMap
	refc83953cc.network, cnetwork_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.network)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnetwork_allocs)

	var cnodes_allocs *cgoAllocMap
	refc83953cc.nodes, cnodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.nodes)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnodes_allocs)

	var cnice_allocs *cgoAllocMap
	refc83953cc.nice, cnice_allocs = (C.uint32_t)(x.nice), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnice_allocs)

	var cnode_inx_allocs *cgoAllocMap
	refc83953cc.node_inx, cnode_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_inx)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnode_inx_allocs)

	var cntasks_per_core_allocs *cgoAllocMap
	refc83953cc.ntasks_per_core, cntasks_per_core_allocs = (C.uint16_t)(x.ntasks_per_core), cgoAllocsUnknown
	allocsc83953cc.Borrow(cntasks_per_core_allocs)

	var cntasks_per_node_allocs *cgoAllocMap
	refc83953cc.ntasks_per_node, cntasks_per_node_allocs = (C.uint16_t)(x.ntasks_per_node), cgoAllocsUnknown
	allocsc83953cc.Borrow(cntasks_per_node_allocs)

	var cntasks_per_socket_allocs *cgoAllocMap
	refc83953cc.ntasks_per_socket, cntasks_per_socket_allocs = (C.uint16_t)(x.ntasks_per_socket), cgoAllocsUnknown
	allocsc83953cc.Borrow(cntasks_per_socket_allocs)

	var cntasks_per_board_allocs *cgoAllocMap
	refc83953cc.ntasks_per_board, cntasks_per_board_allocs = (C.uint16_t)(x.ntasks_per_board), cgoAllocsUnknown
	allocsc83953cc.Borrow(cntasks_per_board_allocs)

	var cnum_cpus_allocs *cgoAllocMap
	refc83953cc.num_cpus, cnum_cpus_allocs = (C.uint32_t)(x.num_cpus), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnum_cpus_allocs)

	var cnum_nodes_allocs *cgoAllocMap
	refc83953cc.num_nodes, cnum_nodes_allocs = (C.uint32_t)(x.num_nodes), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnum_nodes_allocs)

	var cnum_tasks_allocs *cgoAllocMap
	refc83953cc.num_tasks, cnum_tasks_allocs = (C.uint32_t)(x.num_tasks), cgoAllocsUnknown
	allocsc83953cc.Borrow(cnum_tasks_allocs)

	var cpack_job_id_allocs *cgoAllocMap
	refc83953cc.pack_job_id, cpack_job_id_allocs = (C.uint32_t)(x.pack_job_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpack_job_id_allocs)

	var cpack_job_id_set_allocs *cgoAllocMap
	refc83953cc.pack_job_id_set, cpack_job_id_set_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.pack_job_id_set)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpack_job_id_set_allocs)

	var cpack_job_offset_allocs *cgoAllocMap
	refc83953cc.pack_job_offset, cpack_job_offset_allocs = (C.uint32_t)(x.pack_job_offset), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpack_job_offset_allocs)

	var cpartition_allocs *cgoAllocMap
	refc83953cc.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partition)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpartition_allocs)

	var cpn_min_memory_allocs *cgoAllocMap
	refc83953cc.pn_min_memory, cpn_min_memory_allocs = (C.uint64_t)(x.pn_min_memory), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpn_min_memory_allocs)

	var cpn_min_cpus_allocs *cgoAllocMap
	refc83953cc.pn_min_cpus, cpn_min_cpus_allocs = (C.uint16_t)(x.pn_min_cpus), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpn_min_cpus_allocs)

	var cpn_min_tmp_disk_allocs *cgoAllocMap
	refc83953cc.pn_min_tmp_disk, cpn_min_tmp_disk_allocs = (C.uint32_t)(x.pn_min_tmp_disk), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpn_min_tmp_disk_allocs)

	var cpower_flags_allocs *cgoAllocMap
	refc83953cc.power_flags, cpower_flags_allocs = (C.uint8_t)(x.power_flags), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpower_flags_allocs)

	var cpreempt_time_allocs *cgoAllocMap
	refc83953cc.preempt_time, cpreempt_time_allocs = (C.time_t)(x.preempt_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpreempt_time_allocs)

	var cpre_sus_time_allocs *cgoAllocMap
	refc83953cc.pre_sus_time, cpre_sus_time_allocs = (C.time_t)(x.pre_sus_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpre_sus_time_allocs)

	var cpriority_allocs *cgoAllocMap
	refc83953cc.priority, cpriority_allocs = (C.uint32_t)(x.priority), cgoAllocsUnknown
	allocsc83953cc.Borrow(cpriority_allocs)

	var cprofile_allocs *cgoAllocMap
	refc83953cc.profile, cprofile_allocs = (C.uint32_t)(x.profile), cgoAllocsUnknown
	allocsc83953cc.Borrow(cprofile_allocs)

	var cqos_allocs *cgoAllocMap
	refc83953cc.qos, cqos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.qos)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cqos_allocs)

	var creboot_allocs *cgoAllocMap
	refc83953cc.reboot, creboot_allocs = (C.uint8_t)(x.reboot), cgoAllocsUnknown
	allocsc83953cc.Borrow(creboot_allocs)

	var creq_nodes_allocs *cgoAllocMap
	refc83953cc.req_nodes, creq_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.req_nodes)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(creq_nodes_allocs)

	var creq_node_inx_allocs *cgoAllocMap
	refc83953cc.req_node_inx, creq_node_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.req_node_inx)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(creq_node_inx_allocs)

	var creq_switch_allocs *cgoAllocMap
	refc83953cc.req_switch, creq_switch_allocs = (C.uint32_t)(x.req_switch), cgoAllocsUnknown
	allocsc83953cc.Borrow(creq_switch_allocs)

	var crequeue_allocs *cgoAllocMap
	refc83953cc.requeue, crequeue_allocs = (C.uint16_t)(x.requeue), cgoAllocsUnknown
	allocsc83953cc.Borrow(crequeue_allocs)

	var cresize_time_allocs *cgoAllocMap
	refc83953cc.resize_time, cresize_time_allocs = (C.time_t)(x.resize_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(cresize_time_allocs)

	var crestart_cnt_allocs *cgoAllocMap
	refc83953cc.restart_cnt, crestart_cnt_allocs = (C.uint16_t)(x.restart_cnt), cgoAllocsUnknown
	allocsc83953cc.Borrow(crestart_cnt_allocs)

	var cresv_name_allocs *cgoAllocMap
	refc83953cc.resv_name, cresv_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.resv_name)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cresv_name_allocs)

	var csched_nodes_allocs *cgoAllocMap
	refc83953cc.sched_nodes, csched_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.sched_nodes)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(csched_nodes_allocs)

	var cselect_jobinfo_allocs *cgoAllocMap
	refc83953cc.select_jobinfo, cselect_jobinfo_allocs = unpackSDynamic_plugin_data_t(x.select_jobinfo)
	allocsc83953cc.Borrow(cselect_jobinfo_allocs)

	var cshared_allocs *cgoAllocMap
	refc83953cc.shared, cshared_allocs = (C.uint16_t)(x.shared), cgoAllocsUnknown
	allocsc83953cc.Borrow(cshared_allocs)

	var cshow_flags_allocs *cgoAllocMap
	refc83953cc.show_flags, cshow_flags_allocs = (C.uint16_t)(x.show_flags), cgoAllocsUnknown
	allocsc83953cc.Borrow(cshow_flags_allocs)

	var csockets_per_board_allocs *cgoAllocMap
	refc83953cc.sockets_per_board, csockets_per_board_allocs = (C.uint16_t)(x.sockets_per_board), cgoAllocsUnknown
	allocsc83953cc.Borrow(csockets_per_board_allocs)

	var csockets_per_node_allocs *cgoAllocMap
	refc83953cc.sockets_per_node, csockets_per_node_allocs = (C.uint16_t)(x.sockets_per_node), cgoAllocsUnknown
	allocsc83953cc.Borrow(csockets_per_node_allocs)

	var cstart_time_allocs *cgoAllocMap
	refc83953cc.start_time, cstart_time_allocs = (C.time_t)(x.start_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstart_time_allocs)

	var cstart_protocol_ver_allocs *cgoAllocMap
	refc83953cc.start_protocol_ver, cstart_protocol_ver_allocs = (C.uint16_t)(x.start_protocol_ver), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstart_protocol_ver_allocs)

	var cstate_desc_allocs *cgoAllocMap
	refc83953cc.state_desc, cstate_desc_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.state_desc)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstate_desc_allocs)

	var cstate_reason_allocs *cgoAllocMap
	refc83953cc.state_reason, cstate_reason_allocs = (C.uint16_t)(x.state_reason), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstate_reason_allocs)

	var cstd_err_allocs *cgoAllocMap
	refc83953cc.std_err, cstd_err_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.std_err)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstd_err_allocs)

	var cstd_in_allocs *cgoAllocMap
	refc83953cc.std_in, cstd_in_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.std_in)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstd_in_allocs)

	var cstd_out_allocs *cgoAllocMap
	refc83953cc.std_out, cstd_out_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.std_out)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cstd_out_allocs)

	var csubmit_time_allocs *cgoAllocMap
	refc83953cc.submit_time, csubmit_time_allocs = (C.time_t)(x.submit_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(csubmit_time_allocs)

	var csuspend_time_allocs *cgoAllocMap
	refc83953cc.suspend_time, csuspend_time_allocs = (C.time_t)(x.suspend_time), cgoAllocsUnknown
	allocsc83953cc.Borrow(csuspend_time_allocs)

	var ctime_limit_allocs *cgoAllocMap
	refc83953cc.time_limit, ctime_limit_allocs = (C.uint32_t)(x.time_limit), cgoAllocsUnknown
	allocsc83953cc.Borrow(ctime_limit_allocs)

	var ctime_min_allocs *cgoAllocMap
	refc83953cc.time_min, ctime_min_allocs = (C.uint32_t)(x.time_min), cgoAllocsUnknown
	allocsc83953cc.Borrow(ctime_min_allocs)

	var cthreads_per_core_allocs *cgoAllocMap
	refc83953cc.threads_per_core, cthreads_per_core_allocs = (C.uint16_t)(x.threads_per_core), cgoAllocsUnknown
	allocsc83953cc.Borrow(cthreads_per_core_allocs)

	var ctres_req_str_allocs *cgoAllocMap
	refc83953cc.tres_req_str, ctres_req_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_req_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ctres_req_str_allocs)

	var ctres_alloc_str_allocs *cgoAllocMap
	refc83953cc.tres_alloc_str, ctres_alloc_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_alloc_str)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(ctres_alloc_str_allocs)

	var cuser_id_allocs *cgoAllocMap
	refc83953cc.user_id, cuser_id_allocs = (C.uint32_t)(x.user_id), cgoAllocsUnknown
	allocsc83953cc.Borrow(cuser_id_allocs)

	var cuser_name_allocs *cgoAllocMap
	refc83953cc.user_name, cuser_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.user_name)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cuser_name_allocs)

	var cwait4switch_allocs *cgoAllocMap
	refc83953cc.wait4switch, cwait4switch_allocs = (C.uint32_t)(x.wait4switch), cgoAllocsUnknown
	allocsc83953cc.Borrow(cwait4switch_allocs)

	var cwckey_allocs *cgoAllocMap
	refc83953cc.wckey, cwckey_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.wckey)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cwckey_allocs)

	var cwork_dir_allocs *cgoAllocMap
	refc83953cc.work_dir, cwork_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.work_dir)).Data)), cgoAllocsUnknown
	allocsc83953cc.Borrow(cwork_dir_allocs)

	x.refc83953cc = refc83953cc
	x.allocsc83953cc = allocsc83953cc
	return refc83953cc, allocsc83953cc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_info_t) PassValue() (C.slurm_job_info_t, *cgoAllocMap) {
	if x.refc83953cc != nil {
		return *x.refc83953cc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_info_t) Deref() {
	if x.refc83953cc == nil {
		return
	}
	hxf7c2645 := (*sliceHeader)(unsafe.Pointer(&x.account))
	hxf7c2645.Data = uintptr(unsafe.Pointer(x.refc83953cc.account))
	hxf7c2645.Cap = 0x7fffffff
	// hxf7c2645.Len = ?

	hxfc83749 := (*sliceHeader)(unsafe.Pointer(&x.admin_comment))
	hxfc83749.Data = uintptr(unsafe.Pointer(x.refc83953cc.admin_comment))
	hxfc83749.Cap = 0x7fffffff
	// hxfc83749.Len = ?

	hxfa45f91 := (*sliceHeader)(unsafe.Pointer(&x.alloc_node))
	hxfa45f91.Data = uintptr(unsafe.Pointer(x.refc83953cc.alloc_node))
	hxfa45f91.Cap = 0x7fffffff
	// hxfa45f91.Len = ?

	x.alloc_sid = (uint32_t)(x.refc83953cc.alloc_sid)
	x.array_bitmap = (unsafe.Pointer)(unsafe.Pointer(x.refc83953cc.array_bitmap))
	x.array_job_id = (uint32_t)(x.refc83953cc.array_job_id)
	x.array_task_id = (uint32_t)(x.refc83953cc.array_task_id)
	x.array_max_tasks = (uint32_t)(x.refc83953cc.array_max_tasks)
	hxf59053c := (*sliceHeader)(unsafe.Pointer(&x.array_task_str))
	hxf59053c.Data = uintptr(unsafe.Pointer(x.refc83953cc.array_task_str))
	hxf59053c.Cap = 0x7fffffff
	// hxf59053c.Len = ?

	x.assoc_id = (uint32_t)(x.refc83953cc.assoc_id)
	x.batch_flag = (uint16_t)(x.refc83953cc.batch_flag)
	hxf6301b1 := (*sliceHeader)(unsafe.Pointer(&x.batch_host))
	hxf6301b1.Data = uintptr(unsafe.Pointer(x.refc83953cc.batch_host))
	hxf6301b1.Cap = 0x7fffffff
	// hxf6301b1.Len = ?

	x.bitflags = (uint32_t)(x.refc83953cc.bitflags)
	x.boards_per_node = (uint16_t)(x.refc83953cc.boards_per_node)
	hxfb2b86e := (*sliceHeader)(unsafe.Pointer(&x.burst_buffer))
	hxfb2b86e.Data = uintptr(unsafe.Pointer(x.refc83953cc.burst_buffer))
	hxfb2b86e.Cap = 0x7fffffff
	// hxfb2b86e.Len = ?

	hxfe8254a := (*sliceHeader)(unsafe.Pointer(&x.burst_buffer_state))
	hxfe8254a.Data = uintptr(unsafe.Pointer(x.refc83953cc.burst_buffer_state))
	hxfe8254a.Cap = 0x7fffffff
	// hxfe8254a.Len = ?

	hxf1c615c := (*sliceHeader)(unsafe.Pointer(&x.cluster))
	hxf1c615c.Data = uintptr(unsafe.Pointer(x.refc83953cc.cluster))
	hxf1c615c.Cap = 0x7fffffff
	// hxf1c615c.Len = ?

	hxf1ab1b2 := (*sliceHeader)(unsafe.Pointer(&x.cluster_features))
	hxf1ab1b2.Data = uintptr(unsafe.Pointer(x.refc83953cc.cluster_features))
	hxf1ab1b2.Cap = 0x7fffffff
	// hxf1ab1b2.Len = ?

	hxfbbed48 := (*sliceHeader)(unsafe.Pointer(&x.command))
	hxfbbed48.Data = uintptr(unsafe.Pointer(x.refc83953cc.command))
	hxfbbed48.Cap = 0x7fffffff
	// hxfbbed48.Len = ?

	hxf9cdd70 := (*sliceHeader)(unsafe.Pointer(&x.comment))
	hxf9cdd70.Data = uintptr(unsafe.Pointer(x.refc83953cc.comment))
	hxf9cdd70.Cap = 0x7fffffff
	// hxf9cdd70.Len = ?

	x.contiguous = (uint16_t)(x.refc83953cc.contiguous)
	x.core_spec = (uint16_t)(x.refc83953cc.core_spec)
	x.cores_per_socket = (uint16_t)(x.refc83953cc.cores_per_socket)
	x.billable_tres = (float64)(x.refc83953cc.billable_tres)
	x.cpus_per_task = (uint16_t)(x.refc83953cc.cpus_per_task)
	x.cpu_freq_min = (uint32_t)(x.refc83953cc.cpu_freq_min)
	x.cpu_freq_max = (uint32_t)(x.refc83953cc.cpu_freq_max)
	x.cpu_freq_gov = (uint32_t)(x.refc83953cc.cpu_freq_gov)
	x.deadline = (time_t)(x.refc83953cc.deadline)
	x.delay_boot = (uint32_t)(x.refc83953cc.delay_boot)
	hxf84a54f := (*sliceHeader)(unsafe.Pointer(&x.dependency))
	hxf84a54f.Data = uintptr(unsafe.Pointer(x.refc83953cc.dependency))
	hxf84a54f.Cap = 0x7fffffff
	// hxf84a54f.Len = ?

	x.derived_ec = (uint32_t)(x.refc83953cc.derived_ec)
	x.eligible_time = (time_t)(x.refc83953cc.eligible_time)
	x.end_time = (time_t)(x.refc83953cc.end_time)
	hxf6a06fa := (*sliceHeader)(unsafe.Pointer(&x.exc_nodes))
	hxf6a06fa.Data = uintptr(unsafe.Pointer(x.refc83953cc.exc_nodes))
	hxf6a06fa.Cap = 0x7fffffff
	// hxf6a06fa.Len = ?

	hxf5ba44e := (*sliceHeader)(unsafe.Pointer(&x.exc_node_inx))
	hxf5ba44e.Data = uintptr(unsafe.Pointer(x.refc83953cc.exc_node_inx))
	hxf5ba44e.Cap = 0x7fffffff
	// hxf5ba44e.Len = ?

	x.exit_code = (uint32_t)(x.refc83953cc.exit_code)
	hxf1ae993 := (*sliceHeader)(unsafe.Pointer(&x.features))
	hxf1ae993.Data = uintptr(unsafe.Pointer(x.refc83953cc.features))
	hxf1ae993.Cap = 0x7fffffff
	// hxf1ae993.Len = ?

	hxf250a52 := (*sliceHeader)(unsafe.Pointer(&x.fed_origin_str))
	hxf250a52.Data = uintptr(unsafe.Pointer(x.refc83953cc.fed_origin_str))
	hxf250a52.Cap = 0x7fffffff
	// hxf250a52.Len = ?

	x.fed_siblings_active = (uint64_t)(x.refc83953cc.fed_siblings_active)
	hxfe21083 := (*sliceHeader)(unsafe.Pointer(&x.fed_siblings_active_str))
	hxfe21083.Data = uintptr(unsafe.Pointer(x.refc83953cc.fed_siblings_active_str))
	hxfe21083.Cap = 0x7fffffff
	// hxfe21083.Len = ?

	x.fed_siblings_viable = (uint64_t)(x.refc83953cc.fed_siblings_viable)
	hxf9c9c15 := (*sliceHeader)(unsafe.Pointer(&x.fed_siblings_viable_str))
	hxf9c9c15.Data = uintptr(unsafe.Pointer(x.refc83953cc.fed_siblings_viable_str))
	hxf9c9c15.Cap = 0x7fffffff
	// hxf9c9c15.Len = ?

	hxfa3a85e := (*sliceHeader)(unsafe.Pointer(&x.gres))
	hxfa3a85e.Data = uintptr(unsafe.Pointer(x.refc83953cc.gres))
	hxfa3a85e.Cap = 0x7fffffff
	// hxfa3a85e.Len = ?

	x.gres_detail_cnt = (uint32_t)(x.refc83953cc.gres_detail_cnt)
	packSSByte(x.gres_detail_str, x.refc83953cc.gres_detail_str)
	x.group_id = (uint32_t)(x.refc83953cc.group_id)
	x.job_id = (uint32_t)(x.refc83953cc.job_id)
	hxf000188 := (*sliceHeader)(unsafe.Pointer(&x.job_resrcs))
	hxf000188.Data = uintptr(unsafe.Pointer(x.refc83953cc.job_resrcs))
	hxf000188.Cap = 0x7fffffff
	// hxf000188.Len = ?

	x.job_state = (uint32_t)(x.refc83953cc.job_state)
	x.last_sched_eval = (time_t)(x.refc83953cc.last_sched_eval)
	hxf23c7f0 := (*sliceHeader)(unsafe.Pointer(&x.licenses))
	hxf23c7f0.Data = uintptr(unsafe.Pointer(x.refc83953cc.licenses))
	hxf23c7f0.Cap = 0x7fffffff
	// hxf23c7f0.Len = ?

	x.max_cpus = (uint32_t)(x.refc83953cc.max_cpus)
	x.max_nodes = (uint32_t)(x.refc83953cc.max_nodes)
	hxf750b42 := (*sliceHeader)(unsafe.Pointer(&x.mcs_label))
	hxf750b42.Data = uintptr(unsafe.Pointer(x.refc83953cc.mcs_label))
	hxf750b42.Cap = 0x7fffffff
	// hxf750b42.Len = ?

	hxfe6ec59 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxfe6ec59.Data = uintptr(unsafe.Pointer(x.refc83953cc.name))
	hxfe6ec59.Cap = 0x7fffffff
	// hxfe6ec59.Len = ?

	hxfbc67d2 := (*sliceHeader)(unsafe.Pointer(&x.network))
	hxfbc67d2.Data = uintptr(unsafe.Pointer(x.refc83953cc.network))
	hxfbc67d2.Cap = 0x7fffffff
	// hxfbc67d2.Len = ?

	hxfe29d0c := (*sliceHeader)(unsafe.Pointer(&x.nodes))
	hxfe29d0c.Data = uintptr(unsafe.Pointer(x.refc83953cc.nodes))
	hxfe29d0c.Cap = 0x7fffffff
	// hxfe29d0c.Len = ?

	x.nice = (uint32_t)(x.refc83953cc.nice)
	hxfdeace4 := (*sliceHeader)(unsafe.Pointer(&x.node_inx))
	hxfdeace4.Data = uintptr(unsafe.Pointer(x.refc83953cc.node_inx))
	hxfdeace4.Cap = 0x7fffffff
	// hxfdeace4.Len = ?

	x.ntasks_per_core = (uint16_t)(x.refc83953cc.ntasks_per_core)
	x.ntasks_per_node = (uint16_t)(x.refc83953cc.ntasks_per_node)
	x.ntasks_per_socket = (uint16_t)(x.refc83953cc.ntasks_per_socket)
	x.ntasks_per_board = (uint16_t)(x.refc83953cc.ntasks_per_board)
	x.num_cpus = (uint32_t)(x.refc83953cc.num_cpus)
	x.num_nodes = (uint32_t)(x.refc83953cc.num_nodes)
	x.num_tasks = (uint32_t)(x.refc83953cc.num_tasks)
	x.pack_job_id = (uint32_t)(x.refc83953cc.pack_job_id)
	hxf1a6041 := (*sliceHeader)(unsafe.Pointer(&x.pack_job_id_set))
	hxf1a6041.Data = uintptr(unsafe.Pointer(x.refc83953cc.pack_job_id_set))
	hxf1a6041.Cap = 0x7fffffff
	// hxf1a6041.Len = ?

	x.pack_job_offset = (uint32_t)(x.refc83953cc.pack_job_offset)
	hxfe999fb := (*sliceHeader)(unsafe.Pointer(&x.partition))
	hxfe999fb.Data = uintptr(unsafe.Pointer(x.refc83953cc.partition))
	hxfe999fb.Cap = 0x7fffffff
	// hxfe999fb.Len = ?

	x.pn_min_memory = (uint64_t)(x.refc83953cc.pn_min_memory)
	x.pn_min_cpus = (uint16_t)(x.refc83953cc.pn_min_cpus)
	x.pn_min_tmp_disk = (uint32_t)(x.refc83953cc.pn_min_tmp_disk)
	x.power_flags = (uint8_t)(x.refc83953cc.power_flags)
	x.preempt_time = (time_t)(x.refc83953cc.preempt_time)
	x.pre_sus_time = (time_t)(x.refc83953cc.pre_sus_time)
	x.priority = (uint32_t)(x.refc83953cc.priority)
	x.profile = (uint32_t)(x.refc83953cc.profile)
	hxf61194e := (*sliceHeader)(unsafe.Pointer(&x.qos))
	hxf61194e.Data = uintptr(unsafe.Pointer(x.refc83953cc.qos))
	hxf61194e.Cap = 0x7fffffff
	// hxf61194e.Len = ?

	x.reboot = (uint8_t)(x.refc83953cc.reboot)
	hxf31fd96 := (*sliceHeader)(unsafe.Pointer(&x.req_nodes))
	hxf31fd96.Data = uintptr(unsafe.Pointer(x.refc83953cc.req_nodes))
	hxf31fd96.Cap = 0x7fffffff
	// hxf31fd96.Len = ?

	hxf8e12bb := (*sliceHeader)(unsafe.Pointer(&x.req_node_inx))
	hxf8e12bb.Data = uintptr(unsafe.Pointer(x.refc83953cc.req_node_inx))
	hxf8e12bb.Cap = 0x7fffffff
	// hxf8e12bb.Len = ?

	x.req_switch = (uint32_t)(x.refc83953cc.req_switch)
	x.requeue = (uint16_t)(x.refc83953cc.requeue)
	x.resize_time = (time_t)(x.refc83953cc.resize_time)
	x.restart_cnt = (uint16_t)(x.refc83953cc.restart_cnt)
	hxff7f3fa := (*sliceHeader)(unsafe.Pointer(&x.resv_name))
	hxff7f3fa.Data = uintptr(unsafe.Pointer(x.refc83953cc.resv_name))
	hxff7f3fa.Cap = 0x7fffffff
	// hxff7f3fa.Len = ?

	hxf96661a := (*sliceHeader)(unsafe.Pointer(&x.sched_nodes))
	hxf96661a.Data = uintptr(unsafe.Pointer(x.refc83953cc.sched_nodes))
	hxf96661a.Cap = 0x7fffffff
	// hxf96661a.Len = ?

	packSDynamic_plugin_data_t(x.select_jobinfo, x.refc83953cc.select_jobinfo)
	x.shared = (uint16_t)(x.refc83953cc.shared)
	x.show_flags = (uint16_t)(x.refc83953cc.show_flags)
	x.sockets_per_board = (uint16_t)(x.refc83953cc.sockets_per_board)
	x.sockets_per_node = (uint16_t)(x.refc83953cc.sockets_per_node)
	x.start_time = (time_t)(x.refc83953cc.start_time)
	x.start_protocol_ver = (uint16_t)(x.refc83953cc.start_protocol_ver)
	hxf216f94 := (*sliceHeader)(unsafe.Pointer(&x.state_desc))
	hxf216f94.Data = uintptr(unsafe.Pointer(x.refc83953cc.state_desc))
	hxf216f94.Cap = 0x7fffffff
	// hxf216f94.Len = ?

	x.state_reason = (uint16_t)(x.refc83953cc.state_reason)
	hxf124faa := (*sliceHeader)(unsafe.Pointer(&x.std_err))
	hxf124faa.Data = uintptr(unsafe.Pointer(x.refc83953cc.std_err))
	hxf124faa.Cap = 0x7fffffff
	// hxf124faa.Len = ?

	hxf401b49 := (*sliceHeader)(unsafe.Pointer(&x.std_in))
	hxf401b49.Data = uintptr(unsafe.Pointer(x.refc83953cc.std_in))
	hxf401b49.Cap = 0x7fffffff
	// hxf401b49.Len = ?

	hxf987077 := (*sliceHeader)(unsafe.Pointer(&x.std_out))
	hxf987077.Data = uintptr(unsafe.Pointer(x.refc83953cc.std_out))
	hxf987077.Cap = 0x7fffffff
	// hxf987077.Len = ?

	x.submit_time = (time_t)(x.refc83953cc.submit_time)
	x.suspend_time = (time_t)(x.refc83953cc.suspend_time)
	x.time_limit = (uint32_t)(x.refc83953cc.time_limit)
	x.time_min = (uint32_t)(x.refc83953cc.time_min)
	x.threads_per_core = (uint16_t)(x.refc83953cc.threads_per_core)
	hxf5eacde := (*sliceHeader)(unsafe.Pointer(&x.tres_req_str))
	hxf5eacde.Data = uintptr(unsafe.Pointer(x.refc83953cc.tres_req_str))
	hxf5eacde.Cap = 0x7fffffff
	// hxf5eacde.Len = ?

	hxfdcd67f := (*sliceHeader)(unsafe.Pointer(&x.tres_alloc_str))
	hxfdcd67f.Data = uintptr(unsafe.Pointer(x.refc83953cc.tres_alloc_str))
	hxfdcd67f.Cap = 0x7fffffff
	// hxfdcd67f.Len = ?

	x.user_id = (uint32_t)(x.refc83953cc.user_id)
	hxffc0254 := (*sliceHeader)(unsafe.Pointer(&x.user_name))
	hxffc0254.Data = uintptr(unsafe.Pointer(x.refc83953cc.user_name))
	hxffc0254.Cap = 0x7fffffff
	// hxffc0254.Len = ?

	x.wait4switch = (uint32_t)(x.refc83953cc.wait4switch)
	hxfadc325 := (*sliceHeader)(unsafe.Pointer(&x.wckey))
	hxfadc325.Data = uintptr(unsafe.Pointer(x.refc83953cc.wckey))
	hxfadc325.Cap = 0x7fffffff
	// hxfadc325.Len = ?

	hxf4df87a := (*sliceHeader)(unsafe.Pointer(&x.work_dir))
	hxf4df87a.Data = uintptr(unsafe.Pointer(x.refc83953cc.work_dir))
	hxf4df87a.Cap = 0x7fffffff
	// hxf4df87a.Len = ?

}

// allocJob_info_msg_tMemory allocates memory for type C.job_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_info_msg_tValue = unsafe.Sizeof([1]C.job_info_msg_t{})

// unpackSSlurm_job_info_t transforms a sliced Go data structure into plain C format.
func unpackSSlurm_job_info_t(x []slurm_job_info_t) (unpacked *C.slurm_job_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurm_job_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurm_job_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurm_job_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurm_job_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSSlurm_job_info_t reads sliced Go data structure out from plain C format.
func packSSlurm_job_info_t(v []slurm_job_info_t, ptr0 *C.slurm_job_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSlurm_job_info_tValue]C.slurm_job_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newslurm_job_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_info_msg_t) Ref() *C.job_info_msg_t {
	if x == nil {
		return nil
	}
	return x.ref99f73760
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_info_msg_t) Free() {
	if x != nil && x.allocs99f73760 != nil {
		x.allocs99f73760.(*cgoAllocMap).Free()
		x.ref99f73760 = nil
	}
}

// Newjob_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_info_msg_tRef(ref unsafe.Pointer) *job_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(job_info_msg_t)
	obj.ref99f73760 = (*C.job_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_info_msg_t) PassRef() (*C.job_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref99f73760 != nil {
		return x.ref99f73760, nil
	}
	mem99f73760 := allocJob_info_msg_tMemory(1)
	ref99f73760 := (*C.job_info_msg_t)(mem99f73760)
	allocs99f73760 := new(cgoAllocMap)
	allocs99f73760.Add(mem99f73760)

	var clast_update_allocs *cgoAllocMap
	ref99f73760.last_update, clast_update_allocs = (C.time_t)(x.last_update), cgoAllocsUnknown
	allocs99f73760.Borrow(clast_update_allocs)

	var crecord_count_allocs *cgoAllocMap
	ref99f73760.record_count, crecord_count_allocs = (C.uint32_t)(x.record_count), cgoAllocsUnknown
	allocs99f73760.Borrow(crecord_count_allocs)

	var cjob_array_allocs *cgoAllocMap
	ref99f73760.job_array, cjob_array_allocs = unpackSSlurm_job_info_t(x.job_array)
	allocs99f73760.Borrow(cjob_array_allocs)

	x.ref99f73760 = ref99f73760
	x.allocs99f73760 = allocs99f73760
	return ref99f73760, allocs99f73760

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_info_msg_t) PassValue() (C.job_info_msg_t, *cgoAllocMap) {
	if x.ref99f73760 != nil {
		return *x.ref99f73760, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_info_msg_t) Deref() {
	if x.ref99f73760 == nil {
		return
	}
	x.last_update = (time_t)(x.ref99f73760.last_update)
	x.record_count = (uint32_t)(x.ref99f73760.record_count)
	packSSlurm_job_info_t(x.job_array, x.ref99f73760.job_array)
}

// allocStep_update_request_msg_tMemory allocates memory for type C.step_update_request_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStep_update_request_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStep_update_request_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStep_update_request_msg_tValue = unsafe.Sizeof([1]C.step_update_request_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *step_update_request_msg_t) Ref() *C.step_update_request_msg_t {
	if x == nil {
		return nil
	}
	return x.refe10a901b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *step_update_request_msg_t) Free() {
	if x != nil && x.allocse10a901b != nil {
		x.allocse10a901b.(*cgoAllocMap).Free()
		x.refe10a901b = nil
	}
}

// Newstep_update_request_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newstep_update_request_msg_tRef(ref unsafe.Pointer) *step_update_request_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(step_update_request_msg_t)
	obj.refe10a901b = (*C.step_update_request_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *step_update_request_msg_t) PassRef() (*C.step_update_request_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refe10a901b != nil {
		return x.refe10a901b, nil
	}
	meme10a901b := allocStep_update_request_msg_tMemory(1)
	refe10a901b := (*C.step_update_request_msg_t)(meme10a901b)
	allocse10a901b := new(cgoAllocMap)
	allocse10a901b.Add(meme10a901b)

	var cend_time_allocs *cgoAllocMap
	refe10a901b.end_time, cend_time_allocs = (C.time_t)(x.end_time), cgoAllocsUnknown
	allocse10a901b.Borrow(cend_time_allocs)

	var cexit_code_allocs *cgoAllocMap
	refe10a901b.exit_code, cexit_code_allocs = (C.uint32_t)(x.exit_code), cgoAllocsUnknown
	allocse10a901b.Borrow(cexit_code_allocs)

	var cjob_id_allocs *cgoAllocMap
	refe10a901b.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocse10a901b.Borrow(cjob_id_allocs)

	var cjobacct_allocs *cgoAllocMap
	refe10a901b.jobacct, cjobacct_allocs = (*C.jobacctinfo_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.jobacct)).Data)), cgoAllocsUnknown
	allocse10a901b.Borrow(cjobacct_allocs)

	var cname_allocs *cgoAllocMap
	refe10a901b.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocse10a901b.Borrow(cname_allocs)

	var cstart_time_allocs *cgoAllocMap
	refe10a901b.start_time, cstart_time_allocs = (C.time_t)(x.start_time), cgoAllocsUnknown
	allocse10a901b.Borrow(cstart_time_allocs)

	var cstep_id_allocs *cgoAllocMap
	refe10a901b.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocse10a901b.Borrow(cstep_id_allocs)

	var ctime_limit_allocs *cgoAllocMap
	refe10a901b.time_limit, ctime_limit_allocs = (C.uint32_t)(x.time_limit), cgoAllocsUnknown
	allocse10a901b.Borrow(ctime_limit_allocs)

	x.refe10a901b = refe10a901b
	x.allocse10a901b = allocse10a901b
	return refe10a901b, allocse10a901b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x step_update_request_msg_t) PassValue() (C.step_update_request_msg_t, *cgoAllocMap) {
	if x.refe10a901b != nil {
		return *x.refe10a901b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *step_update_request_msg_t) Deref() {
	if x.refe10a901b == nil {
		return
	}
	x.end_time = (time_t)(x.refe10a901b.end_time)
	x.exit_code = (uint32_t)(x.refe10a901b.exit_code)
	x.job_id = (uint32_t)(x.refe10a901b.job_id)
	hxf3dfc2a := (*sliceHeader)(unsafe.Pointer(&x.jobacct))
	hxf3dfc2a.Data = uintptr(unsafe.Pointer(x.refe10a901b.jobacct))
	hxf3dfc2a.Cap = 0x7fffffff
	// hxf3dfc2a.Len = ?

	hxf08e5b7 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf08e5b7.Data = uintptr(unsafe.Pointer(x.refe10a901b.name))
	hxf08e5b7.Cap = 0x7fffffff
	// hxf08e5b7.Len = ?

	x.start_time = (time_t)(x.refe10a901b.start_time)
	x.step_id = (uint32_t)(x.refe10a901b.step_id)
	x.time_limit = (uint32_t)(x.refe10a901b.time_limit)
}

// allocSlurm_step_layout_req_tMemory allocates memory for type C.slurm_step_layout_req_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_step_layout_req_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_step_layout_req_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_step_layout_req_tValue = unsafe.Sizeof([1]C.slurm_step_layout_req_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_step_layout_req_t) Ref() *C.slurm_step_layout_req_t {
	if x == nil {
		return nil
	}
	return x.ref36330321
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_step_layout_req_t) Free() {
	if x != nil && x.allocs36330321 != nil {
		x.allocs36330321.(*cgoAllocMap).Free()
		x.ref36330321 = nil
	}
}

// Newslurm_step_layout_req_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_step_layout_req_tRef(ref unsafe.Pointer) *slurm_step_layout_req_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_step_layout_req_t)
	obj.ref36330321 = (*C.slurm_step_layout_req_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_step_layout_req_t) PassRef() (*C.slurm_step_layout_req_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref36330321 != nil {
		return x.ref36330321, nil
	}
	mem36330321 := allocSlurm_step_layout_req_tMemory(1)
	ref36330321 := (*C.slurm_step_layout_req_t)(mem36330321)
	allocs36330321 := new(cgoAllocMap)
	allocs36330321.Add(mem36330321)

	var cnode_list_allocs *cgoAllocMap
	ref36330321.node_list, cnode_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_list)).Data)), cgoAllocsUnknown
	allocs36330321.Borrow(cnode_list_allocs)

	var ccpus_per_node_allocs *cgoAllocMap
	ref36330321.cpus_per_node, ccpus_per_node_allocs = (*C.uint16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cpus_per_node)).Data)), cgoAllocsUnknown
	allocs36330321.Borrow(ccpus_per_node_allocs)

	var ccpu_count_reps_allocs *cgoAllocMap
	ref36330321.cpu_count_reps, ccpu_count_reps_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cpu_count_reps)).Data)), cgoAllocsUnknown
	allocs36330321.Borrow(ccpu_count_reps_allocs)

	var cnum_hosts_allocs *cgoAllocMap
	ref36330321.num_hosts, cnum_hosts_allocs = (C.uint32_t)(x.num_hosts), cgoAllocsUnknown
	allocs36330321.Borrow(cnum_hosts_allocs)

	var cnum_tasks_allocs *cgoAllocMap
	ref36330321.num_tasks, cnum_tasks_allocs = (C.uint32_t)(x.num_tasks), cgoAllocsUnknown
	allocs36330321.Borrow(cnum_tasks_allocs)

	var ccpus_per_task_allocs *cgoAllocMap
	ref36330321.cpus_per_task, ccpus_per_task_allocs = (*C.uint16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cpus_per_task)).Data)), cgoAllocsUnknown
	allocs36330321.Borrow(ccpus_per_task_allocs)

	var ccpus_task_reps_allocs *cgoAllocMap
	ref36330321.cpus_task_reps, ccpus_task_reps_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cpus_task_reps)).Data)), cgoAllocsUnknown
	allocs36330321.Borrow(ccpus_task_reps_allocs)

	var ctask_dist_allocs *cgoAllocMap
	ref36330321.task_dist, ctask_dist_allocs = (C.uint32_t)(x.task_dist), cgoAllocsUnknown
	allocs36330321.Borrow(ctask_dist_allocs)

	var cplane_size_allocs *cgoAllocMap
	ref36330321.plane_size, cplane_size_allocs = (C.uint16_t)(x.plane_size), cgoAllocsUnknown
	allocs36330321.Borrow(cplane_size_allocs)

	x.ref36330321 = ref36330321
	x.allocs36330321 = allocs36330321
	return ref36330321, allocs36330321

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_step_layout_req_t) PassValue() (C.slurm_step_layout_req_t, *cgoAllocMap) {
	if x.ref36330321 != nil {
		return *x.ref36330321, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_step_layout_req_t) Deref() {
	if x.ref36330321 == nil {
		return
	}
	hxf6deab7 := (*sliceHeader)(unsafe.Pointer(&x.node_list))
	hxf6deab7.Data = uintptr(unsafe.Pointer(x.ref36330321.node_list))
	hxf6deab7.Cap = 0x7fffffff
	// hxf6deab7.Len = ?

	hxf4d737a := (*sliceHeader)(unsafe.Pointer(&x.cpus_per_node))
	hxf4d737a.Data = uintptr(unsafe.Pointer(x.ref36330321.cpus_per_node))
	hxf4d737a.Cap = 0x7fffffff
	// hxf4d737a.Len = ?

	hxf5a38e8 := (*sliceHeader)(unsafe.Pointer(&x.cpu_count_reps))
	hxf5a38e8.Data = uintptr(unsafe.Pointer(x.ref36330321.cpu_count_reps))
	hxf5a38e8.Cap = 0x7fffffff
	// hxf5a38e8.Len = ?

	x.num_hosts = (uint32_t)(x.ref36330321.num_hosts)
	x.num_tasks = (uint32_t)(x.ref36330321.num_tasks)
	hxf51492a := (*sliceHeader)(unsafe.Pointer(&x.cpus_per_task))
	hxf51492a.Data = uintptr(unsafe.Pointer(x.ref36330321.cpus_per_task))
	hxf51492a.Cap = 0x7fffffff
	// hxf51492a.Len = ?

	hxf108fc8 := (*sliceHeader)(unsafe.Pointer(&x.cpus_task_reps))
	hxf108fc8.Data = uintptr(unsafe.Pointer(x.ref36330321.cpus_task_reps))
	hxf108fc8.Cap = 0x7fffffff
	// hxf108fc8.Len = ?

	x.task_dist = (uint32_t)(x.ref36330321.task_dist)
	x.plane_size = (uint16_t)(x.ref36330321.plane_size)
}

// allocSlurm_step_layout_tMemory allocates memory for type C.slurm_step_layout_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_step_layout_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_step_layout_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_step_layout_tValue = unsafe.Sizeof([1]C.slurm_step_layout_t{})

// allocPUint32_tMemory allocates memory for type *C.uint32_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPUint32_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPUint32_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPUint32_tValue = unsafe.Sizeof([1]*C.uint32_t{})

// unpackSSUUint32_t transforms a sliced Go data structure into plain C format.
func unpackSSUUint32_t(x [][]uint32_t) (unpacked **C.uint32_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.uint32_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPUint32_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.uint32_t)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.uint32_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.uint32_t)(unsafe.Pointer(h.Data))
	return
}

// packSSUUint32_t reads sliced Go data structure out from plain C format.
func packSSUUint32_t(v [][]uint32_t, ptr0 **C.uint32_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.uint32_t)(unsafe.Pointer(ptr0)))[i0]
		hxf166788 := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxf166788.Data = uintptr(unsafe.Pointer(ptr1))
		hxf166788.Cap = 0x7fffffff
		// hxf166788.Len = ?
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_step_layout_t) Ref() *C.slurm_step_layout_t {
	if x == nil {
		return nil
	}
	return x.ref5f1a1283
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_step_layout_t) Free() {
	if x != nil && x.allocs5f1a1283 != nil {
		x.allocs5f1a1283.(*cgoAllocMap).Free()
		x.ref5f1a1283 = nil
	}
}

// Newslurm_step_layout_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_step_layout_tRef(ref unsafe.Pointer) *slurm_step_layout_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_step_layout_t)
	obj.ref5f1a1283 = (*C.slurm_step_layout_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_step_layout_t) PassRef() (*C.slurm_step_layout_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5f1a1283 != nil {
		return x.ref5f1a1283, nil
	}
	mem5f1a1283 := allocSlurm_step_layout_tMemory(1)
	ref5f1a1283 := (*C.slurm_step_layout_t)(mem5f1a1283)
	allocs5f1a1283 := new(cgoAllocMap)
	allocs5f1a1283.Add(mem5f1a1283)

	var cfront_end_allocs *cgoAllocMap
	ref5f1a1283.front_end, cfront_end_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.front_end)).Data)), cgoAllocsUnknown
	allocs5f1a1283.Borrow(cfront_end_allocs)

	var cnode_cnt_allocs *cgoAllocMap
	ref5f1a1283.node_cnt, cnode_cnt_allocs = (C.uint32_t)(x.node_cnt), cgoAllocsUnknown
	allocs5f1a1283.Borrow(cnode_cnt_allocs)

	var cnode_list_allocs *cgoAllocMap
	ref5f1a1283.node_list, cnode_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_list)).Data)), cgoAllocsUnknown
	allocs5f1a1283.Borrow(cnode_list_allocs)

	var cplane_size_allocs *cgoAllocMap
	ref5f1a1283.plane_size, cplane_size_allocs = (C.uint16_t)(x.plane_size), cgoAllocsUnknown
	allocs5f1a1283.Borrow(cplane_size_allocs)

	var cstart_protocol_ver_allocs *cgoAllocMap
	ref5f1a1283.start_protocol_ver, cstart_protocol_ver_allocs = (C.uint16_t)(x.start_protocol_ver), cgoAllocsUnknown
	allocs5f1a1283.Borrow(cstart_protocol_ver_allocs)

	var ctasks_allocs *cgoAllocMap
	ref5f1a1283.tasks, ctasks_allocs = (*C.uint16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tasks)).Data)), cgoAllocsUnknown
	allocs5f1a1283.Borrow(ctasks_allocs)

	var ctask_cnt_allocs *cgoAllocMap
	ref5f1a1283.task_cnt, ctask_cnt_allocs = (C.uint32_t)(x.task_cnt), cgoAllocsUnknown
	allocs5f1a1283.Borrow(ctask_cnt_allocs)

	var ctask_dist_allocs *cgoAllocMap
	ref5f1a1283.task_dist, ctask_dist_allocs = (C.uint32_t)(x.task_dist), cgoAllocsUnknown
	allocs5f1a1283.Borrow(ctask_dist_allocs)

	var ctids_allocs *cgoAllocMap
	ref5f1a1283.tids, ctids_allocs = unpackSSUUint32_t(x.tids)
	allocs5f1a1283.Borrow(ctids_allocs)

	x.ref5f1a1283 = ref5f1a1283
	x.allocs5f1a1283 = allocs5f1a1283
	return ref5f1a1283, allocs5f1a1283

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_step_layout_t) PassValue() (C.slurm_step_layout_t, *cgoAllocMap) {
	if x.ref5f1a1283 != nil {
		return *x.ref5f1a1283, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_step_layout_t) Deref() {
	if x.ref5f1a1283 == nil {
		return
	}
	hxff268a1 := (*sliceHeader)(unsafe.Pointer(&x.front_end))
	hxff268a1.Data = uintptr(unsafe.Pointer(x.ref5f1a1283.front_end))
	hxff268a1.Cap = 0x7fffffff
	// hxff268a1.Len = ?

	x.node_cnt = (uint32_t)(x.ref5f1a1283.node_cnt)
	hxff5babc := (*sliceHeader)(unsafe.Pointer(&x.node_list))
	hxff5babc.Data = uintptr(unsafe.Pointer(x.ref5f1a1283.node_list))
	hxff5babc.Cap = 0x7fffffff
	// hxff5babc.Len = ?

	x.plane_size = (uint16_t)(x.ref5f1a1283.plane_size)
	x.start_protocol_ver = (uint16_t)(x.ref5f1a1283.start_protocol_ver)
	hxf91b269 := (*sliceHeader)(unsafe.Pointer(&x.tasks))
	hxf91b269.Data = uintptr(unsafe.Pointer(x.ref5f1a1283.tasks))
	hxf91b269.Cap = 0x7fffffff
	// hxf91b269.Len = ?

	x.task_cnt = (uint32_t)(x.ref5f1a1283.task_cnt)
	x.task_dist = (uint32_t)(x.ref5f1a1283.task_dist)
	packSSUUint32_t(x.tids, x.ref5f1a1283.tids)
}

// allocSlurm_step_io_fds_tMemory allocates memory for type C.slurm_step_io_fds_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_step_io_fds_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_step_io_fds_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_step_io_fds_tValue = unsafe.Sizeof([1]C.slurm_step_io_fds_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_step_io_fds_t) Ref() *C.slurm_step_io_fds_t {
	if x == nil {
		return nil
	}
	return x.ref76cac59f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_step_io_fds_t) Free() {
	if x != nil && x.allocs76cac59f != nil {
		x.allocs76cac59f.(*cgoAllocMap).Free()
		x.ref76cac59f = nil
	}
}

// Newslurm_step_io_fds_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_step_io_fds_tRef(ref unsafe.Pointer) *slurm_step_io_fds_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_step_io_fds_t)
	obj.ref76cac59f = (*C.slurm_step_io_fds_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_step_io_fds_t) PassRef() (*C.slurm_step_io_fds_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref76cac59f != nil {
		return x.ref76cac59f, nil
	}
	mem76cac59f := allocSlurm_step_io_fds_tMemory(1)
	ref76cac59f := (*C.slurm_step_io_fds_t)(mem76cac59f)
	allocs76cac59f := new(cgoAllocMap)
	allocs76cac59f.Add(mem76cac59f)

	x.ref76cac59f = ref76cac59f
	x.allocs76cac59f = allocs76cac59f
	return ref76cac59f, allocs76cac59f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_step_io_fds_t) PassValue() (C.slurm_step_io_fds_t, *cgoAllocMap) {
	if x.ref76cac59f != nil {
		return *x.ref76cac59f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_step_io_fds_t) Deref() {
	if x.ref76cac59f == nil {
		return
	}
}

// allocLaunch_tasks_response_msg_tMemory allocates memory for type C.launch_tasks_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocLaunch_tasks_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfLaunch_tasks_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfLaunch_tasks_response_msg_tValue = unsafe.Sizeof([1]C.launch_tasks_response_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *launch_tasks_response_msg_t) Ref() *C.launch_tasks_response_msg_t {
	if x == nil {
		return nil
	}
	return x.ref9da0a9b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *launch_tasks_response_msg_t) Free() {
	if x != nil && x.allocs9da0a9b != nil {
		x.allocs9da0a9b.(*cgoAllocMap).Free()
		x.ref9da0a9b = nil
	}
}

// Newlaunch_tasks_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newlaunch_tasks_response_msg_tRef(ref unsafe.Pointer) *launch_tasks_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(launch_tasks_response_msg_t)
	obj.ref9da0a9b = (*C.launch_tasks_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *launch_tasks_response_msg_t) PassRef() (*C.launch_tasks_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9da0a9b != nil {
		return x.ref9da0a9b, nil
	}
	mem9da0a9b := allocLaunch_tasks_response_msg_tMemory(1)
	ref9da0a9b := (*C.launch_tasks_response_msg_t)(mem9da0a9b)
	allocs9da0a9b := new(cgoAllocMap)
	allocs9da0a9b.Add(mem9da0a9b)

	var cjob_id_allocs *cgoAllocMap
	ref9da0a9b.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs9da0a9b.Borrow(cjob_id_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref9da0a9b.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocs9da0a9b.Borrow(cstep_id_allocs)

	var creturn_code_allocs *cgoAllocMap
	ref9da0a9b.return_code, creturn_code_allocs = (C.uint32_t)(x.return_code), cgoAllocsUnknown
	allocs9da0a9b.Borrow(creturn_code_allocs)

	var cnode_name_allocs *cgoAllocMap
	ref9da0a9b.node_name, cnode_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_name)).Data)), cgoAllocsUnknown
	allocs9da0a9b.Borrow(cnode_name_allocs)

	var csrun_node_id_allocs *cgoAllocMap
	ref9da0a9b.srun_node_id, csrun_node_id_allocs = (C.uint32_t)(x.srun_node_id), cgoAllocsUnknown
	allocs9da0a9b.Borrow(csrun_node_id_allocs)

	var ccount_of_pids_allocs *cgoAllocMap
	ref9da0a9b.count_of_pids, ccount_of_pids_allocs = (C.uint32_t)(x.count_of_pids), cgoAllocsUnknown
	allocs9da0a9b.Borrow(ccount_of_pids_allocs)

	var clocal_pids_allocs *cgoAllocMap
	ref9da0a9b.local_pids, clocal_pids_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.local_pids)).Data)), cgoAllocsUnknown
	allocs9da0a9b.Borrow(clocal_pids_allocs)

	var ctask_ids_allocs *cgoAllocMap
	ref9da0a9b.task_ids, ctask_ids_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.task_ids)).Data)), cgoAllocsUnknown
	allocs9da0a9b.Borrow(ctask_ids_allocs)

	x.ref9da0a9b = ref9da0a9b
	x.allocs9da0a9b = allocs9da0a9b
	return ref9da0a9b, allocs9da0a9b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x launch_tasks_response_msg_t) PassValue() (C.launch_tasks_response_msg_t, *cgoAllocMap) {
	if x.ref9da0a9b != nil {
		return *x.ref9da0a9b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *launch_tasks_response_msg_t) Deref() {
	if x.ref9da0a9b == nil {
		return
	}
	x.job_id = (uint32_t)(x.ref9da0a9b.job_id)
	x.step_id = (uint32_t)(x.ref9da0a9b.step_id)
	x.return_code = (uint32_t)(x.ref9da0a9b.return_code)
	hxf463f61 := (*sliceHeader)(unsafe.Pointer(&x.node_name))
	hxf463f61.Data = uintptr(unsafe.Pointer(x.ref9da0a9b.node_name))
	hxf463f61.Cap = 0x7fffffff
	// hxf463f61.Len = ?

	x.srun_node_id = (uint32_t)(x.ref9da0a9b.srun_node_id)
	x.count_of_pids = (uint32_t)(x.ref9da0a9b.count_of_pids)
	hxf4a3cd3 := (*sliceHeader)(unsafe.Pointer(&x.local_pids))
	hxf4a3cd3.Data = uintptr(unsafe.Pointer(x.ref9da0a9b.local_pids))
	hxf4a3cd3.Cap = 0x7fffffff
	// hxf4a3cd3.Len = ?

	hxf5ebcba := (*sliceHeader)(unsafe.Pointer(&x.task_ids))
	hxf5ebcba.Data = uintptr(unsafe.Pointer(x.ref9da0a9b.task_ids))
	hxf5ebcba.Cap = 0x7fffffff
	// hxf5ebcba.Len = ?

}

// allocTask_exit_msg_tMemory allocates memory for type C.task_exit_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocTask_exit_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfTask_exit_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfTask_exit_msg_tValue = unsafe.Sizeof([1]C.task_exit_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *task_exit_msg_t) Ref() *C.task_exit_msg_t {
	if x == nil {
		return nil
	}
	return x.ref968ced3f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *task_exit_msg_t) Free() {
	if x != nil && x.allocs968ced3f != nil {
		x.allocs968ced3f.(*cgoAllocMap).Free()
		x.ref968ced3f = nil
	}
}

// Newtask_exit_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newtask_exit_msg_tRef(ref unsafe.Pointer) *task_exit_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(task_exit_msg_t)
	obj.ref968ced3f = (*C.task_exit_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *task_exit_msg_t) PassRef() (*C.task_exit_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref968ced3f != nil {
		return x.ref968ced3f, nil
	}
	mem968ced3f := allocTask_exit_msg_tMemory(1)
	ref968ced3f := (*C.task_exit_msg_t)(mem968ced3f)
	allocs968ced3f := new(cgoAllocMap)
	allocs968ced3f.Add(mem968ced3f)

	var cnum_tasks_allocs *cgoAllocMap
	ref968ced3f.num_tasks, cnum_tasks_allocs = (C.uint32_t)(x.num_tasks), cgoAllocsUnknown
	allocs968ced3f.Borrow(cnum_tasks_allocs)

	var ctask_id_list_allocs *cgoAllocMap
	ref968ced3f.task_id_list, ctask_id_list_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.task_id_list)).Data)), cgoAllocsUnknown
	allocs968ced3f.Borrow(ctask_id_list_allocs)

	var creturn_code_allocs *cgoAllocMap
	ref968ced3f.return_code, creturn_code_allocs = (C.uint32_t)(x.return_code), cgoAllocsUnknown
	allocs968ced3f.Borrow(creturn_code_allocs)

	var cjob_id_allocs *cgoAllocMap
	ref968ced3f.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs968ced3f.Borrow(cjob_id_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref968ced3f.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocs968ced3f.Borrow(cstep_id_allocs)

	x.ref968ced3f = ref968ced3f
	x.allocs968ced3f = allocs968ced3f
	return ref968ced3f, allocs968ced3f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x task_exit_msg_t) PassValue() (C.task_exit_msg_t, *cgoAllocMap) {
	if x.ref968ced3f != nil {
		return *x.ref968ced3f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *task_exit_msg_t) Deref() {
	if x.ref968ced3f == nil {
		return
	}
	x.num_tasks = (uint32_t)(x.ref968ced3f.num_tasks)
	hxfba3d6e := (*sliceHeader)(unsafe.Pointer(&x.task_id_list))
	hxfba3d6e.Data = uintptr(unsafe.Pointer(x.ref968ced3f.task_id_list))
	hxfba3d6e.Cap = 0x7fffffff
	// hxfba3d6e.Len = ?

	x.return_code = (uint32_t)(x.ref968ced3f.return_code)
	x.job_id = (uint32_t)(x.ref968ced3f.job_id)
	x.step_id = (uint32_t)(x.ref968ced3f.step_id)
}

// allocSrun_ping_msg_tMemory allocates memory for type C.srun_ping_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSrun_ping_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSrun_ping_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSrun_ping_msg_tValue = unsafe.Sizeof([1]C.srun_ping_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *srun_ping_msg_t) Ref() *C.srun_ping_msg_t {
	if x == nil {
		return nil
	}
	return x.ref9d7c08b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *srun_ping_msg_t) Free() {
	if x != nil && x.allocs9d7c08b != nil {
		x.allocs9d7c08b.(*cgoAllocMap).Free()
		x.ref9d7c08b = nil
	}
}

// Newsrun_ping_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newsrun_ping_msg_tRef(ref unsafe.Pointer) *srun_ping_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(srun_ping_msg_t)
	obj.ref9d7c08b = (*C.srun_ping_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *srun_ping_msg_t) PassRef() (*C.srun_ping_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9d7c08b != nil {
		return x.ref9d7c08b, nil
	}
	mem9d7c08b := allocSrun_ping_msg_tMemory(1)
	ref9d7c08b := (*C.srun_ping_msg_t)(mem9d7c08b)
	allocs9d7c08b := new(cgoAllocMap)
	allocs9d7c08b.Add(mem9d7c08b)

	var cjob_id_allocs *cgoAllocMap
	ref9d7c08b.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs9d7c08b.Borrow(cjob_id_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref9d7c08b.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocs9d7c08b.Borrow(cstep_id_allocs)

	x.ref9d7c08b = ref9d7c08b
	x.allocs9d7c08b = allocs9d7c08b
	return ref9d7c08b, allocs9d7c08b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x srun_ping_msg_t) PassValue() (C.srun_ping_msg_t, *cgoAllocMap) {
	if x.ref9d7c08b != nil {
		return *x.ref9d7c08b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *srun_ping_msg_t) Deref() {
	if x.ref9d7c08b == nil {
		return
	}
	x.job_id = (uint32_t)(x.ref9d7c08b.job_id)
	x.step_id = (uint32_t)(x.ref9d7c08b.step_id)
}

// allocSrun_job_complete_msg_tMemory allocates memory for type C.srun_job_complete_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSrun_job_complete_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSrun_job_complete_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSrun_job_complete_msg_tValue = unsafe.Sizeof([1]C.srun_job_complete_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *srun_job_complete_msg_t) Ref() *C.srun_job_complete_msg_t {
	if x == nil {
		return nil
	}
	return x.ref12083467
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *srun_job_complete_msg_t) Free() {
	if x != nil && x.allocs12083467 != nil {
		x.allocs12083467.(*cgoAllocMap).Free()
		x.ref12083467 = nil
	}
}

// Newsrun_job_complete_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newsrun_job_complete_msg_tRef(ref unsafe.Pointer) *srun_job_complete_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(srun_job_complete_msg_t)
	obj.ref12083467 = (*C.srun_job_complete_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *srun_job_complete_msg_t) PassRef() (*C.srun_job_complete_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref12083467 != nil {
		return x.ref12083467, nil
	}
	mem12083467 := allocSrun_job_complete_msg_tMemory(1)
	ref12083467 := (*C.srun_job_complete_msg_t)(mem12083467)
	allocs12083467 := new(cgoAllocMap)
	allocs12083467.Add(mem12083467)

	var cjob_id_allocs *cgoAllocMap
	ref12083467.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs12083467.Borrow(cjob_id_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref12083467.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocs12083467.Borrow(cstep_id_allocs)

	x.ref12083467 = ref12083467
	x.allocs12083467 = allocs12083467
	return ref12083467, allocs12083467

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x srun_job_complete_msg_t) PassValue() (C.srun_job_complete_msg_t, *cgoAllocMap) {
	if x.ref12083467 != nil {
		return *x.ref12083467, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *srun_job_complete_msg_t) Deref() {
	if x.ref12083467 == nil {
		return
	}
	x.job_id = (uint32_t)(x.ref12083467.job_id)
	x.step_id = (uint32_t)(x.ref12083467.step_id)
}

// allocSrun_timeout_msg_tMemory allocates memory for type C.srun_timeout_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSrun_timeout_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSrun_timeout_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSrun_timeout_msg_tValue = unsafe.Sizeof([1]C.srun_timeout_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *srun_timeout_msg_t) Ref() *C.srun_timeout_msg_t {
	if x == nil {
		return nil
	}
	return x.ref51c6caaf
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *srun_timeout_msg_t) Free() {
	if x != nil && x.allocs51c6caaf != nil {
		x.allocs51c6caaf.(*cgoAllocMap).Free()
		x.ref51c6caaf = nil
	}
}

// Newsrun_timeout_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newsrun_timeout_msg_tRef(ref unsafe.Pointer) *srun_timeout_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(srun_timeout_msg_t)
	obj.ref51c6caaf = (*C.srun_timeout_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *srun_timeout_msg_t) PassRef() (*C.srun_timeout_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref51c6caaf != nil {
		return x.ref51c6caaf, nil
	}
	mem51c6caaf := allocSrun_timeout_msg_tMemory(1)
	ref51c6caaf := (*C.srun_timeout_msg_t)(mem51c6caaf)
	allocs51c6caaf := new(cgoAllocMap)
	allocs51c6caaf.Add(mem51c6caaf)

	var cjob_id_allocs *cgoAllocMap
	ref51c6caaf.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs51c6caaf.Borrow(cjob_id_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref51c6caaf.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocs51c6caaf.Borrow(cstep_id_allocs)

	var ctimeout_allocs *cgoAllocMap
	ref51c6caaf.timeout, ctimeout_allocs = (C.time_t)(x.timeout), cgoAllocsUnknown
	allocs51c6caaf.Borrow(ctimeout_allocs)

	x.ref51c6caaf = ref51c6caaf
	x.allocs51c6caaf = allocs51c6caaf
	return ref51c6caaf, allocs51c6caaf

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x srun_timeout_msg_t) PassValue() (C.srun_timeout_msg_t, *cgoAllocMap) {
	if x.ref51c6caaf != nil {
		return *x.ref51c6caaf, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *srun_timeout_msg_t) Deref() {
	if x.ref51c6caaf == nil {
		return
	}
	x.job_id = (uint32_t)(x.ref51c6caaf.job_id)
	x.step_id = (uint32_t)(x.ref51c6caaf.step_id)
	x.timeout = (time_t)(x.ref51c6caaf.timeout)
}

// allocSrun_user_msg_tMemory allocates memory for type C.srun_user_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSrun_user_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSrun_user_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSrun_user_msg_tValue = unsafe.Sizeof([1]C.srun_user_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *srun_user_msg_t) Ref() *C.srun_user_msg_t {
	if x == nil {
		return nil
	}
	return x.refbe9c6c1f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *srun_user_msg_t) Free() {
	if x != nil && x.allocsbe9c6c1f != nil {
		x.allocsbe9c6c1f.(*cgoAllocMap).Free()
		x.refbe9c6c1f = nil
	}
}

// Newsrun_user_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newsrun_user_msg_tRef(ref unsafe.Pointer) *srun_user_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(srun_user_msg_t)
	obj.refbe9c6c1f = (*C.srun_user_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *srun_user_msg_t) PassRef() (*C.srun_user_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe9c6c1f != nil {
		return x.refbe9c6c1f, nil
	}
	membe9c6c1f := allocSrun_user_msg_tMemory(1)
	refbe9c6c1f := (*C.srun_user_msg_t)(membe9c6c1f)
	allocsbe9c6c1f := new(cgoAllocMap)
	allocsbe9c6c1f.Add(membe9c6c1f)

	var cjob_id_allocs *cgoAllocMap
	refbe9c6c1f.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocsbe9c6c1f.Borrow(cjob_id_allocs)

	var cmsg_allocs *cgoAllocMap
	refbe9c6c1f.msg, cmsg_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.msg)).Data)), cgoAllocsUnknown
	allocsbe9c6c1f.Borrow(cmsg_allocs)

	x.refbe9c6c1f = refbe9c6c1f
	x.allocsbe9c6c1f = allocsbe9c6c1f
	return refbe9c6c1f, allocsbe9c6c1f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x srun_user_msg_t) PassValue() (C.srun_user_msg_t, *cgoAllocMap) {
	if x.refbe9c6c1f != nil {
		return *x.refbe9c6c1f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *srun_user_msg_t) Deref() {
	if x.refbe9c6c1f == nil {
		return
	}
	x.job_id = (uint32_t)(x.refbe9c6c1f.job_id)
	hxf88c1ce := (*sliceHeader)(unsafe.Pointer(&x.msg))
	hxf88c1ce.Data = uintptr(unsafe.Pointer(x.refbe9c6c1f.msg))
	hxf88c1ce.Cap = 0x7fffffff
	// hxf88c1ce.Len = ?

}

// allocSrun_node_fail_msg_tMemory allocates memory for type C.srun_node_fail_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSrun_node_fail_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSrun_node_fail_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSrun_node_fail_msg_tValue = unsafe.Sizeof([1]C.srun_node_fail_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *srun_node_fail_msg_t) Ref() *C.srun_node_fail_msg_t {
	if x == nil {
		return nil
	}
	return x.refcad68cc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *srun_node_fail_msg_t) Free() {
	if x != nil && x.allocscad68cc != nil {
		x.allocscad68cc.(*cgoAllocMap).Free()
		x.refcad68cc = nil
	}
}

// Newsrun_node_fail_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newsrun_node_fail_msg_tRef(ref unsafe.Pointer) *srun_node_fail_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(srun_node_fail_msg_t)
	obj.refcad68cc = (*C.srun_node_fail_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *srun_node_fail_msg_t) PassRef() (*C.srun_node_fail_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refcad68cc != nil {
		return x.refcad68cc, nil
	}
	memcad68cc := allocSrun_node_fail_msg_tMemory(1)
	refcad68cc := (*C.srun_node_fail_msg_t)(memcad68cc)
	allocscad68cc := new(cgoAllocMap)
	allocscad68cc.Add(memcad68cc)

	var cjob_id_allocs *cgoAllocMap
	refcad68cc.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocscad68cc.Borrow(cjob_id_allocs)

	var cnodelist_allocs *cgoAllocMap
	refcad68cc.nodelist, cnodelist_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.nodelist)).Data)), cgoAllocsUnknown
	allocscad68cc.Borrow(cnodelist_allocs)

	var cstep_id_allocs *cgoAllocMap
	refcad68cc.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocscad68cc.Borrow(cstep_id_allocs)

	x.refcad68cc = refcad68cc
	x.allocscad68cc = allocscad68cc
	return refcad68cc, allocscad68cc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x srun_node_fail_msg_t) PassValue() (C.srun_node_fail_msg_t, *cgoAllocMap) {
	if x.refcad68cc != nil {
		return *x.refcad68cc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *srun_node_fail_msg_t) Deref() {
	if x.refcad68cc == nil {
		return
	}
	x.job_id = (uint32_t)(x.refcad68cc.job_id)
	hxf9b6716 := (*sliceHeader)(unsafe.Pointer(&x.nodelist))
	hxf9b6716.Data = uintptr(unsafe.Pointer(x.refcad68cc.nodelist))
	hxf9b6716.Cap = 0x7fffffff
	// hxf9b6716.Len = ?

	x.step_id = (uint32_t)(x.refcad68cc.step_id)
}

// allocSrun_step_missing_msg_tMemory allocates memory for type C.srun_step_missing_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSrun_step_missing_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSrun_step_missing_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSrun_step_missing_msg_tValue = unsafe.Sizeof([1]C.srun_step_missing_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *srun_step_missing_msg_t) Ref() *C.srun_step_missing_msg_t {
	if x == nil {
		return nil
	}
	return x.ref7629d094
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *srun_step_missing_msg_t) Free() {
	if x != nil && x.allocs7629d094 != nil {
		x.allocs7629d094.(*cgoAllocMap).Free()
		x.ref7629d094 = nil
	}
}

// Newsrun_step_missing_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newsrun_step_missing_msg_tRef(ref unsafe.Pointer) *srun_step_missing_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(srun_step_missing_msg_t)
	obj.ref7629d094 = (*C.srun_step_missing_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *srun_step_missing_msg_t) PassRef() (*C.srun_step_missing_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref7629d094 != nil {
		return x.ref7629d094, nil
	}
	mem7629d094 := allocSrun_step_missing_msg_tMemory(1)
	ref7629d094 := (*C.srun_step_missing_msg_t)(mem7629d094)
	allocs7629d094 := new(cgoAllocMap)
	allocs7629d094.Add(mem7629d094)

	var cjob_id_allocs *cgoAllocMap
	ref7629d094.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs7629d094.Borrow(cjob_id_allocs)

	var cnodelist_allocs *cgoAllocMap
	ref7629d094.nodelist, cnodelist_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.nodelist)).Data)), cgoAllocsUnknown
	allocs7629d094.Borrow(cnodelist_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref7629d094.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocs7629d094.Borrow(cstep_id_allocs)

	x.ref7629d094 = ref7629d094
	x.allocs7629d094 = allocs7629d094
	return ref7629d094, allocs7629d094

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x srun_step_missing_msg_t) PassValue() (C.srun_step_missing_msg_t, *cgoAllocMap) {
	if x.ref7629d094 != nil {
		return *x.ref7629d094, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *srun_step_missing_msg_t) Deref() {
	if x.ref7629d094 == nil {
		return
	}
	x.job_id = (uint32_t)(x.ref7629d094.job_id)
	hxf814ae1 := (*sliceHeader)(unsafe.Pointer(&x.nodelist))
	hxf814ae1.Data = uintptr(unsafe.Pointer(x.ref7629d094.nodelist))
	hxf814ae1.Cap = 0x7fffffff
	// hxf814ae1.Len = ?

	x.step_id = (uint32_t)(x.ref7629d094.step_id)
}

// allocSuspend_msg_tMemory allocates memory for type C.suspend_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSuspend_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSuspend_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSuspend_msg_tValue = unsafe.Sizeof([1]C.suspend_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *suspend_msg_t) Ref() *C.suspend_msg_t {
	if x == nil {
		return nil
	}
	return x.ref4b652065
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *suspend_msg_t) Free() {
	if x != nil && x.allocs4b652065 != nil {
		x.allocs4b652065.(*cgoAllocMap).Free()
		x.ref4b652065 = nil
	}
}

// Newsuspend_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newsuspend_msg_tRef(ref unsafe.Pointer) *suspend_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(suspend_msg_t)
	obj.ref4b652065 = (*C.suspend_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *suspend_msg_t) PassRef() (*C.suspend_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref4b652065 != nil {
		return x.ref4b652065, nil
	}
	mem4b652065 := allocSuspend_msg_tMemory(1)
	ref4b652065 := (*C.suspend_msg_t)(mem4b652065)
	allocs4b652065 := new(cgoAllocMap)
	allocs4b652065.Add(mem4b652065)

	var cop_allocs *cgoAllocMap
	ref4b652065.op, cop_allocs = (C.uint16_t)(x.op), cgoAllocsUnknown
	allocs4b652065.Borrow(cop_allocs)

	var cjob_id_allocs *cgoAllocMap
	ref4b652065.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs4b652065.Borrow(cjob_id_allocs)

	var cjob_id_str_allocs *cgoAllocMap
	ref4b652065.job_id_str, cjob_id_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_id_str)).Data)), cgoAllocsUnknown
	allocs4b652065.Borrow(cjob_id_str_allocs)

	x.ref4b652065 = ref4b652065
	x.allocs4b652065 = allocs4b652065
	return ref4b652065, allocs4b652065

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x suspend_msg_t) PassValue() (C.suspend_msg_t, *cgoAllocMap) {
	if x.ref4b652065 != nil {
		return *x.ref4b652065, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *suspend_msg_t) Deref() {
	if x.ref4b652065 == nil {
		return
	}
	x.op = (uint16_t)(x.ref4b652065.op)
	x.job_id = (uint32_t)(x.ref4b652065.job_id)
	hxf3b5d02 := (*sliceHeader)(unsafe.Pointer(&x.job_id_str))
	hxf3b5d02.Data = uintptr(unsafe.Pointer(x.ref4b652065.job_id_str))
	hxf3b5d02.Cap = 0x7fffffff
	// hxf3b5d02.Len = ?

}

// allocTop_job_msg_tMemory allocates memory for type C.top_job_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocTop_job_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfTop_job_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfTop_job_msg_tValue = unsafe.Sizeof([1]C.top_job_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *top_job_msg_t) Ref() *C.top_job_msg_t {
	if x == nil {
		return nil
	}
	return x.ref874ee789
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *top_job_msg_t) Free() {
	if x != nil && x.allocs874ee789 != nil {
		x.allocs874ee789.(*cgoAllocMap).Free()
		x.ref874ee789 = nil
	}
}

// Newtop_job_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newtop_job_msg_tRef(ref unsafe.Pointer) *top_job_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(top_job_msg_t)
	obj.ref874ee789 = (*C.top_job_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *top_job_msg_t) PassRef() (*C.top_job_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref874ee789 != nil {
		return x.ref874ee789, nil
	}
	mem874ee789 := allocTop_job_msg_tMemory(1)
	ref874ee789 := (*C.top_job_msg_t)(mem874ee789)
	allocs874ee789 := new(cgoAllocMap)
	allocs874ee789.Add(mem874ee789)

	var cop_allocs *cgoAllocMap
	ref874ee789.op, cop_allocs = (C.uint16_t)(x.op), cgoAllocsUnknown
	allocs874ee789.Borrow(cop_allocs)

	var cjob_id_allocs *cgoAllocMap
	ref874ee789.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs874ee789.Borrow(cjob_id_allocs)

	var cjob_id_str_allocs *cgoAllocMap
	ref874ee789.job_id_str, cjob_id_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_id_str)).Data)), cgoAllocsUnknown
	allocs874ee789.Borrow(cjob_id_str_allocs)

	x.ref874ee789 = ref874ee789
	x.allocs874ee789 = allocs874ee789
	return ref874ee789, allocs874ee789

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x top_job_msg_t) PassValue() (C.top_job_msg_t, *cgoAllocMap) {
	if x.ref874ee789 != nil {
		return *x.ref874ee789, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *top_job_msg_t) Deref() {
	if x.ref874ee789 == nil {
		return
	}
	x.op = (uint16_t)(x.ref874ee789.op)
	x.job_id = (uint32_t)(x.ref874ee789.job_id)
	hxf560981 := (*sliceHeader)(unsafe.Pointer(&x.job_id_str))
	hxf560981.Data = uintptr(unsafe.Pointer(x.ref874ee789.job_id_str))
	hxf560981.Cap = 0x7fffffff
	// hxf560981.Len = ?

}

// allocSlurm_step_ctx_params_tMemory allocates memory for type C.slurm_step_ctx_params_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_step_ctx_params_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_step_ctx_params_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_step_ctx_params_tValue = unsafe.Sizeof([1]C.slurm_step_ctx_params_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_step_ctx_params_t) Ref() *C.slurm_step_ctx_params_t {
	if x == nil {
		return nil
	}
	return x.ref7db1e503
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_step_ctx_params_t) Free() {
	if x != nil && x.allocs7db1e503 != nil {
		x.allocs7db1e503.(*cgoAllocMap).Free()
		x.ref7db1e503 = nil
	}
}

// Newslurm_step_ctx_params_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_step_ctx_params_tRef(ref unsafe.Pointer) *slurm_step_ctx_params_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_step_ctx_params_t)
	obj.ref7db1e503 = (*C.slurm_step_ctx_params_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_step_ctx_params_t) PassRef() (*C.slurm_step_ctx_params_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref7db1e503 != nil {
		return x.ref7db1e503, nil
	}
	mem7db1e503 := allocSlurm_step_ctx_params_tMemory(1)
	ref7db1e503 := (*C.slurm_step_ctx_params_t)(mem7db1e503)
	allocs7db1e503 := new(cgoAllocMap)
	allocs7db1e503.Add(mem7db1e503)

	var cckpt_interval_allocs *cgoAllocMap
	ref7db1e503.ckpt_interval, cckpt_interval_allocs = (C.uint16_t)(x.ckpt_interval), cgoAllocsUnknown
	allocs7db1e503.Borrow(cckpt_interval_allocs)

	var ccpu_count_allocs *cgoAllocMap
	ref7db1e503.cpu_count, ccpu_count_allocs = (C.uint32_t)(x.cpu_count), cgoAllocsUnknown
	allocs7db1e503.Borrow(ccpu_count_allocs)

	var ccpu_freq_min_allocs *cgoAllocMap
	ref7db1e503.cpu_freq_min, ccpu_freq_min_allocs = (C.uint32_t)(x.cpu_freq_min), cgoAllocsUnknown
	allocs7db1e503.Borrow(ccpu_freq_min_allocs)

	var ccpu_freq_max_allocs *cgoAllocMap
	ref7db1e503.cpu_freq_max, ccpu_freq_max_allocs = (C.uint32_t)(x.cpu_freq_max), cgoAllocsUnknown
	allocs7db1e503.Borrow(ccpu_freq_max_allocs)

	var ccpu_freq_gov_allocs *cgoAllocMap
	ref7db1e503.cpu_freq_gov, ccpu_freq_gov_allocs = (C.uint32_t)(x.cpu_freq_gov), cgoAllocsUnknown
	allocs7db1e503.Borrow(ccpu_freq_gov_allocs)

	var cexclusive_allocs *cgoAllocMap
	ref7db1e503.exclusive, cexclusive_allocs = (C.uint16_t)(x.exclusive), cgoAllocsUnknown
	allocs7db1e503.Borrow(cexclusive_allocs)

	var cfeatures_allocs *cgoAllocMap
	ref7db1e503.features, cfeatures_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.features)).Data)), cgoAllocsUnknown
	allocs7db1e503.Borrow(cfeatures_allocs)

	var cimmediate_allocs *cgoAllocMap
	ref7db1e503.immediate, cimmediate_allocs = (C.uint16_t)(x.immediate), cgoAllocsUnknown
	allocs7db1e503.Borrow(cimmediate_allocs)

	var cjob_id_allocs *cgoAllocMap
	ref7db1e503.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs7db1e503.Borrow(cjob_id_allocs)

	var cpn_min_memory_allocs *cgoAllocMap
	ref7db1e503.pn_min_memory, cpn_min_memory_allocs = (C.uint64_t)(x.pn_min_memory), cgoAllocsUnknown
	allocs7db1e503.Borrow(cpn_min_memory_allocs)

	var cckpt_dir_allocs *cgoAllocMap
	ref7db1e503.ckpt_dir, cckpt_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ckpt_dir)).Data)), cgoAllocsUnknown
	allocs7db1e503.Borrow(cckpt_dir_allocs)

	var cgres_allocs *cgoAllocMap
	ref7db1e503.gres, cgres_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.gres)).Data)), cgoAllocsUnknown
	allocs7db1e503.Borrow(cgres_allocs)

	var cname_allocs *cgoAllocMap
	ref7db1e503.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs7db1e503.Borrow(cname_allocs)

	var cnetwork_allocs *cgoAllocMap
	ref7db1e503.network, cnetwork_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.network)).Data)), cgoAllocsUnknown
	allocs7db1e503.Borrow(cnetwork_allocs)

	var cprofile_allocs *cgoAllocMap
	ref7db1e503.profile, cprofile_allocs = (C.uint32_t)(x.profile), cgoAllocsUnknown
	allocs7db1e503.Borrow(cprofile_allocs)

	var cno_kill_allocs *cgoAllocMap
	ref7db1e503.no_kill, cno_kill_allocs = (C.uint8_t)(x.no_kill), cgoAllocsUnknown
	allocs7db1e503.Borrow(cno_kill_allocs)

	var cmin_nodes_allocs *cgoAllocMap
	ref7db1e503.min_nodes, cmin_nodes_allocs = (C.uint32_t)(x.min_nodes), cgoAllocsUnknown
	allocs7db1e503.Borrow(cmin_nodes_allocs)

	var cmax_nodes_allocs *cgoAllocMap
	ref7db1e503.max_nodes, cmax_nodes_allocs = (C.uint32_t)(x.max_nodes), cgoAllocsUnknown
	allocs7db1e503.Borrow(cmax_nodes_allocs)

	var cnode_list_allocs *cgoAllocMap
	ref7db1e503.node_list, cnode_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_list)).Data)), cgoAllocsUnknown
	allocs7db1e503.Borrow(cnode_list_allocs)

	var covercommit_allocs *cgoAllocMap
	ref7db1e503.overcommit, covercommit_allocs = (C._Bool)(x.overcommit), cgoAllocsUnknown
	allocs7db1e503.Borrow(covercommit_allocs)

	var cplane_size_allocs *cgoAllocMap
	ref7db1e503.plane_size, cplane_size_allocs = (C.uint16_t)(x.plane_size), cgoAllocsUnknown
	allocs7db1e503.Borrow(cplane_size_allocs)

	var crelative_allocs *cgoAllocMap
	ref7db1e503.relative, crelative_allocs = (C.uint16_t)(x.relative), cgoAllocsUnknown
	allocs7db1e503.Borrow(crelative_allocs)

	var cresv_port_cnt_allocs *cgoAllocMap
	ref7db1e503.resv_port_cnt, cresv_port_cnt_allocs = (C.uint16_t)(x.resv_port_cnt), cgoAllocsUnknown
	allocs7db1e503.Borrow(cresv_port_cnt_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref7db1e503.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocs7db1e503.Borrow(cstep_id_allocs)

	var ctask_count_allocs *cgoAllocMap
	ref7db1e503.task_count, ctask_count_allocs = (C.uint32_t)(x.task_count), cgoAllocsUnknown
	allocs7db1e503.Borrow(ctask_count_allocs)

	var ctask_dist_allocs *cgoAllocMap
	ref7db1e503.task_dist, ctask_dist_allocs = (C.uint32_t)(x.task_dist), cgoAllocsUnknown
	allocs7db1e503.Borrow(ctask_dist_allocs)

	var ctime_limit_allocs *cgoAllocMap
	ref7db1e503.time_limit, ctime_limit_allocs = (C.uint32_t)(x.time_limit), cgoAllocsUnknown
	allocs7db1e503.Borrow(ctime_limit_allocs)

	var cuid_allocs *cgoAllocMap
	ref7db1e503.uid, cuid_allocs = (C.uid_t)(x.uid), cgoAllocsUnknown
	allocs7db1e503.Borrow(cuid_allocs)

	var cverbose_level_allocs *cgoAllocMap
	ref7db1e503.verbose_level, cverbose_level_allocs = (C.uint16_t)(x.verbose_level), cgoAllocsUnknown
	allocs7db1e503.Borrow(cverbose_level_allocs)

	x.ref7db1e503 = ref7db1e503
	x.allocs7db1e503 = allocs7db1e503
	return ref7db1e503, allocs7db1e503

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_step_ctx_params_t) PassValue() (C.slurm_step_ctx_params_t, *cgoAllocMap) {
	if x.ref7db1e503 != nil {
		return *x.ref7db1e503, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_step_ctx_params_t) Deref() {
	if x.ref7db1e503 == nil {
		return
	}
	x.ckpt_interval = (uint16_t)(x.ref7db1e503.ckpt_interval)
	x.cpu_count = (uint32_t)(x.ref7db1e503.cpu_count)
	x.cpu_freq_min = (uint32_t)(x.ref7db1e503.cpu_freq_min)
	x.cpu_freq_max = (uint32_t)(x.ref7db1e503.cpu_freq_max)
	x.cpu_freq_gov = (uint32_t)(x.ref7db1e503.cpu_freq_gov)
	x.exclusive = (uint16_t)(x.ref7db1e503.exclusive)
	hxf31c255 := (*sliceHeader)(unsafe.Pointer(&x.features))
	hxf31c255.Data = uintptr(unsafe.Pointer(x.ref7db1e503.features))
	hxf31c255.Cap = 0x7fffffff
	// hxf31c255.Len = ?

	x.immediate = (uint16_t)(x.ref7db1e503.immediate)
	x.job_id = (uint32_t)(x.ref7db1e503.job_id)
	x.pn_min_memory = (uint64_t)(x.ref7db1e503.pn_min_memory)
	hxf8fb793 := (*sliceHeader)(unsafe.Pointer(&x.ckpt_dir))
	hxf8fb793.Data = uintptr(unsafe.Pointer(x.ref7db1e503.ckpt_dir))
	hxf8fb793.Cap = 0x7fffffff
	// hxf8fb793.Len = ?

	hxf2d756f := (*sliceHeader)(unsafe.Pointer(&x.gres))
	hxf2d756f.Data = uintptr(unsafe.Pointer(x.ref7db1e503.gres))
	hxf2d756f.Cap = 0x7fffffff
	// hxf2d756f.Len = ?

	hxf782cbf := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf782cbf.Data = uintptr(unsafe.Pointer(x.ref7db1e503.name))
	hxf782cbf.Cap = 0x7fffffff
	// hxf782cbf.Len = ?

	hxf5edf98 := (*sliceHeader)(unsafe.Pointer(&x.network))
	hxf5edf98.Data = uintptr(unsafe.Pointer(x.ref7db1e503.network))
	hxf5edf98.Cap = 0x7fffffff
	// hxf5edf98.Len = ?

	x.profile = (uint32_t)(x.ref7db1e503.profile)
	x.no_kill = (uint8_t)(x.ref7db1e503.no_kill)
	x.min_nodes = (uint32_t)(x.ref7db1e503.min_nodes)
	x.max_nodes = (uint32_t)(x.ref7db1e503.max_nodes)
	hxf6a0436 := (*sliceHeader)(unsafe.Pointer(&x.node_list))
	hxf6a0436.Data = uintptr(unsafe.Pointer(x.ref7db1e503.node_list))
	hxf6a0436.Cap = 0x7fffffff
	// hxf6a0436.Len = ?

	x.overcommit = (bool)(x.ref7db1e503.overcommit)
	x.plane_size = (uint16_t)(x.ref7db1e503.plane_size)
	x.relative = (uint16_t)(x.ref7db1e503.relative)
	x.resv_port_cnt = (uint16_t)(x.ref7db1e503.resv_port_cnt)
	x.step_id = (uint32_t)(x.ref7db1e503.step_id)
	x.task_count = (uint32_t)(x.ref7db1e503.task_count)
	x.task_dist = (uint32_t)(x.ref7db1e503.task_dist)
	x.time_limit = (uint32_t)(x.ref7db1e503.time_limit)
	x.uid = (uid_t)(x.ref7db1e503.uid)
	x.verbose_level = (uint16_t)(x.ref7db1e503.verbose_level)
}

// allocSlurm_step_launch_params_tMemory allocates memory for type C.slurm_step_launch_params_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_step_launch_params_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_step_launch_params_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_step_launch_params_tValue = unsafe.Sizeof([1]C.slurm_step_launch_params_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_step_launch_params_t) Ref() *C.slurm_step_launch_params_t {
	if x == nil {
		return nil
	}
	return x.ref17ff30d9
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_step_launch_params_t) Free() {
	if x != nil && x.allocs17ff30d9 != nil {
		x.allocs17ff30d9.(*cgoAllocMap).Free()
		x.ref17ff30d9 = nil
	}
}

// Newslurm_step_launch_params_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_step_launch_params_tRef(ref unsafe.Pointer) *slurm_step_launch_params_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_step_launch_params_t)
	obj.ref17ff30d9 = (*C.slurm_step_launch_params_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_step_launch_params_t) PassRef() (*C.slurm_step_launch_params_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref17ff30d9 != nil {
		return x.ref17ff30d9, nil
	}
	mem17ff30d9 := allocSlurm_step_launch_params_tMemory(1)
	ref17ff30d9 := (*C.slurm_step_launch_params_t)(mem17ff30d9)
	allocs17ff30d9 := new(cgoAllocMap)
	allocs17ff30d9.Add(mem17ff30d9)

	var calias_list_allocs *cgoAllocMap
	ref17ff30d9.alias_list, calias_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.alias_list)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(calias_list_allocs)

	var cargc_allocs *cgoAllocMap
	ref17ff30d9.argc, cargc_allocs = (C.uint32_t)(x.argc), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cargc_allocs)

	var cargv_allocs *cgoAllocMap
	ref17ff30d9.argv, cargv_allocs = unpackSSByte(x.argv)
	allocs17ff30d9.Borrow(cargv_allocs)

	var cenvc_allocs *cgoAllocMap
	ref17ff30d9.envc, cenvc_allocs = (C.uint32_t)(x.envc), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cenvc_allocs)

	var cenv_allocs *cgoAllocMap
	ref17ff30d9.env, cenv_allocs = unpackSSByte(x.env)
	allocs17ff30d9.Borrow(cenv_allocs)

	var ccwd_allocs *cgoAllocMap
	ref17ff30d9.cwd, ccwd_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cwd)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(ccwd_allocs)

	var cuser_managed_io_allocs *cgoAllocMap
	ref17ff30d9.user_managed_io, cuser_managed_io_allocs = (C._Bool)(x.user_managed_io), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cuser_managed_io_allocs)

	var cmsg_timeout_allocs *cgoAllocMap
	ref17ff30d9.msg_timeout, cmsg_timeout_allocs = (C.uint32_t)(x.msg_timeout), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cmsg_timeout_allocs)

	var cntasks_per_board_allocs *cgoAllocMap
	ref17ff30d9.ntasks_per_board, cntasks_per_board_allocs = (C.uint16_t)(x.ntasks_per_board), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cntasks_per_board_allocs)

	var cntasks_per_core_allocs *cgoAllocMap
	ref17ff30d9.ntasks_per_core, cntasks_per_core_allocs = (C.uint16_t)(x.ntasks_per_core), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cntasks_per_core_allocs)

	var cntasks_per_socket_allocs *cgoAllocMap
	ref17ff30d9.ntasks_per_socket, cntasks_per_socket_allocs = (C.uint16_t)(x.ntasks_per_socket), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cntasks_per_socket_allocs)

	var cbuffered_stdio_allocs *cgoAllocMap
	ref17ff30d9.buffered_stdio, cbuffered_stdio_allocs = (C._Bool)(x.buffered_stdio), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cbuffered_stdio_allocs)

	var clabelio_allocs *cgoAllocMap
	ref17ff30d9.labelio, clabelio_allocs = (C._Bool)(x.labelio), cgoAllocsUnknown
	allocs17ff30d9.Borrow(clabelio_allocs)

	var cremote_output_filename_allocs *cgoAllocMap
	ref17ff30d9.remote_output_filename, cremote_output_filename_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.remote_output_filename)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cremote_output_filename_allocs)

	var cremote_error_filename_allocs *cgoAllocMap
	ref17ff30d9.remote_error_filename, cremote_error_filename_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.remote_error_filename)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cremote_error_filename_allocs)

	var cremote_input_filename_allocs *cgoAllocMap
	ref17ff30d9.remote_input_filename, cremote_input_filename_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.remote_input_filename)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cremote_input_filename_allocs)

	var clocal_fds_allocs *cgoAllocMap
	ref17ff30d9.local_fds, clocal_fds_allocs = x.local_fds.PassValue()
	allocs17ff30d9.Borrow(clocal_fds_allocs)

	var cgid_allocs *cgoAllocMap
	ref17ff30d9.gid, cgid_allocs = (C.uint32_t)(x.gid), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cgid_allocs)

	var cmulti_prog_allocs *cgoAllocMap
	ref17ff30d9.multi_prog, cmulti_prog_allocs = (C._Bool)(x.multi_prog), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cmulti_prog_allocs)

	var cno_alloc_allocs *cgoAllocMap
	ref17ff30d9.no_alloc, cno_alloc_allocs = (C._Bool)(x.no_alloc), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cno_alloc_allocs)

	var cslurmd_debug_allocs *cgoAllocMap
	ref17ff30d9.slurmd_debug, cslurmd_debug_allocs = (C.uint32_t)(x.slurmd_debug), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cslurmd_debug_allocs)

	var cnode_offset_allocs *cgoAllocMap
	ref17ff30d9.node_offset, cnode_offset_allocs = (C.uint32_t)(x.node_offset), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cnode_offset_allocs)

	var cpack_jobid_allocs *cgoAllocMap
	ref17ff30d9.pack_jobid, cpack_jobid_allocs = (C.uint32_t)(x.pack_jobid), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cpack_jobid_allocs)

	var cpack_nnodes_allocs *cgoAllocMap
	ref17ff30d9.pack_nnodes, cpack_nnodes_allocs = (C.uint32_t)(x.pack_nnodes), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cpack_nnodes_allocs)

	var cpack_ntasks_allocs *cgoAllocMap
	ref17ff30d9.pack_ntasks, cpack_ntasks_allocs = (C.uint32_t)(x.pack_ntasks), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cpack_ntasks_allocs)

	var cpack_task_cnts_allocs *cgoAllocMap
	ref17ff30d9.pack_task_cnts, cpack_task_cnts_allocs = (*C.uint16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.pack_task_cnts)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cpack_task_cnts_allocs)

	var cpack_tids_allocs *cgoAllocMap
	ref17ff30d9.pack_tids, cpack_tids_allocs = unpackSSUUint32_t(x.pack_tids)
	allocs17ff30d9.Borrow(cpack_tids_allocs)

	var cpack_offset_allocs *cgoAllocMap
	ref17ff30d9.pack_offset, cpack_offset_allocs = (C.uint32_t)(x.pack_offset), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cpack_offset_allocs)

	var cpack_task_offset_allocs *cgoAllocMap
	ref17ff30d9.pack_task_offset, cpack_task_offset_allocs = (C.uint32_t)(x.pack_task_offset), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cpack_task_offset_allocs)

	var cpack_node_list_allocs *cgoAllocMap
	ref17ff30d9.pack_node_list, cpack_node_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.pack_node_list)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cpack_node_list_allocs)

	var cparallel_debug_allocs *cgoAllocMap
	ref17ff30d9.parallel_debug, cparallel_debug_allocs = (C._Bool)(x.parallel_debug), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cparallel_debug_allocs)

	var cprofile_allocs *cgoAllocMap
	ref17ff30d9.profile, cprofile_allocs = (C.uint32_t)(x.profile), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cprofile_allocs)

	var ctask_prolog_allocs *cgoAllocMap
	ref17ff30d9.task_prolog, ctask_prolog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.task_prolog)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(ctask_prolog_allocs)

	var ctask_epilog_allocs *cgoAllocMap
	ref17ff30d9.task_epilog, ctask_epilog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.task_epilog)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(ctask_epilog_allocs)

	var ccpu_bind_type_allocs *cgoAllocMap
	ref17ff30d9.cpu_bind_type, ccpu_bind_type_allocs = (C.uint16_t)(x.cpu_bind_type), cgoAllocsUnknown
	allocs17ff30d9.Borrow(ccpu_bind_type_allocs)

	var ccpu_bind_allocs *cgoAllocMap
	ref17ff30d9.cpu_bind, ccpu_bind_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cpu_bind)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(ccpu_bind_allocs)

	var ccpu_freq_min_allocs *cgoAllocMap
	ref17ff30d9.cpu_freq_min, ccpu_freq_min_allocs = (C.uint32_t)(x.cpu_freq_min), cgoAllocsUnknown
	allocs17ff30d9.Borrow(ccpu_freq_min_allocs)

	var ccpu_freq_max_allocs *cgoAllocMap
	ref17ff30d9.cpu_freq_max, ccpu_freq_max_allocs = (C.uint32_t)(x.cpu_freq_max), cgoAllocsUnknown
	allocs17ff30d9.Borrow(ccpu_freq_max_allocs)

	var ccpu_freq_gov_allocs *cgoAllocMap
	ref17ff30d9.cpu_freq_gov, ccpu_freq_gov_allocs = (C.uint32_t)(x.cpu_freq_gov), cgoAllocsUnknown
	allocs17ff30d9.Borrow(ccpu_freq_gov_allocs)

	var cmem_bind_type_allocs *cgoAllocMap
	ref17ff30d9.mem_bind_type, cmem_bind_type_allocs = (C.uint16_t)(x.mem_bind_type), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cmem_bind_type_allocs)

	var cmem_bind_allocs *cgoAllocMap
	ref17ff30d9.mem_bind, cmem_bind_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mem_bind)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cmem_bind_allocs)

	var caccel_bind_type_allocs *cgoAllocMap
	ref17ff30d9.accel_bind_type, caccel_bind_type_allocs = (C.uint16_t)(x.accel_bind_type), cgoAllocsUnknown
	allocs17ff30d9.Borrow(caccel_bind_type_allocs)

	var cmax_sockets_allocs *cgoAllocMap
	ref17ff30d9.max_sockets, cmax_sockets_allocs = (C.uint16_t)(x.max_sockets), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cmax_sockets_allocs)

	var cmax_cores_allocs *cgoAllocMap
	ref17ff30d9.max_cores, cmax_cores_allocs = (C.uint16_t)(x.max_cores), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cmax_cores_allocs)

	var cmax_threads_allocs *cgoAllocMap
	ref17ff30d9.max_threads, cmax_threads_allocs = (C.uint16_t)(x.max_threads), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cmax_threads_allocs)

	var ccpus_per_task_allocs *cgoAllocMap
	ref17ff30d9.cpus_per_task, ccpus_per_task_allocs = (C.uint16_t)(x.cpus_per_task), cgoAllocsUnknown
	allocs17ff30d9.Borrow(ccpus_per_task_allocs)

	var ctask_dist_allocs *cgoAllocMap
	ref17ff30d9.task_dist, ctask_dist_allocs = (C.uint32_t)(x.task_dist), cgoAllocsUnknown
	allocs17ff30d9.Borrow(ctask_dist_allocs)

	var cpartition_allocs *cgoAllocMap
	ref17ff30d9.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partition)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cpartition_allocs)

	var cpreserve_env_allocs *cgoAllocMap
	ref17ff30d9.preserve_env, cpreserve_env_allocs = (C._Bool)(x.preserve_env), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cpreserve_env_allocs)

	var cmpi_plugin_name_allocs *cgoAllocMap
	ref17ff30d9.mpi_plugin_name, cmpi_plugin_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mpi_plugin_name)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cmpi_plugin_name_allocs)

	var copen_mode_allocs *cgoAllocMap
	ref17ff30d9.open_mode, copen_mode_allocs = (C.uint8_t)(x.open_mode), cgoAllocsUnknown
	allocs17ff30d9.Borrow(copen_mode_allocs)

	var cacctg_freq_allocs *cgoAllocMap
	ref17ff30d9.acctg_freq, cacctg_freq_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.acctg_freq)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cacctg_freq_allocs)

	var cpty_allocs *cgoAllocMap
	ref17ff30d9.pty, cpty_allocs = (C._Bool)(x.pty), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cpty_allocs)

	var cckpt_dir_allocs *cgoAllocMap
	ref17ff30d9.ckpt_dir, cckpt_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ckpt_dir)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cckpt_dir_allocs)

	var crestart_dir_allocs *cgoAllocMap
	ref17ff30d9.restart_dir, crestart_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.restart_dir)).Data)), cgoAllocsUnknown
	allocs17ff30d9.Borrow(crestart_dir_allocs)

	var cspank_job_env_allocs *cgoAllocMap
	ref17ff30d9.spank_job_env, cspank_job_env_allocs = unpackSSByte(x.spank_job_env)
	allocs17ff30d9.Borrow(cspank_job_env_allocs)

	var cspank_job_env_size_allocs *cgoAllocMap
	ref17ff30d9.spank_job_env_size, cspank_job_env_size_allocs = (C.uint32_t)(x.spank_job_env_size), cgoAllocsUnknown
	allocs17ff30d9.Borrow(cspank_job_env_size_allocs)

	x.ref17ff30d9 = ref17ff30d9
	x.allocs17ff30d9 = allocs17ff30d9
	return ref17ff30d9, allocs17ff30d9

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_step_launch_params_t) PassValue() (C.slurm_step_launch_params_t, *cgoAllocMap) {
	if x.ref17ff30d9 != nil {
		return *x.ref17ff30d9, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_step_launch_params_t) Deref() {
	if x.ref17ff30d9 == nil {
		return
	}
	hxffd8c28 := (*sliceHeader)(unsafe.Pointer(&x.alias_list))
	hxffd8c28.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.alias_list))
	hxffd8c28.Cap = 0x7fffffff
	// hxffd8c28.Len = ?

	x.argc = (uint32_t)(x.ref17ff30d9.argc)
	packSSByte(x.argv, x.ref17ff30d9.argv)
	x.envc = (uint32_t)(x.ref17ff30d9.envc)
	packSSByte(x.env, x.ref17ff30d9.env)
	hxfe252d8 := (*sliceHeader)(unsafe.Pointer(&x.cwd))
	hxfe252d8.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.cwd))
	hxfe252d8.Cap = 0x7fffffff
	// hxfe252d8.Len = ?

	x.user_managed_io = (bool)(x.ref17ff30d9.user_managed_io)
	x.msg_timeout = (uint32_t)(x.ref17ff30d9.msg_timeout)
	x.ntasks_per_board = (uint16_t)(x.ref17ff30d9.ntasks_per_board)
	x.ntasks_per_core = (uint16_t)(x.ref17ff30d9.ntasks_per_core)
	x.ntasks_per_socket = (uint16_t)(x.ref17ff30d9.ntasks_per_socket)
	x.buffered_stdio = (bool)(x.ref17ff30d9.buffered_stdio)
	x.labelio = (bool)(x.ref17ff30d9.labelio)
	hxfbced30 := (*sliceHeader)(unsafe.Pointer(&x.remote_output_filename))
	hxfbced30.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.remote_output_filename))
	hxfbced30.Cap = 0x7fffffff
	// hxfbced30.Len = ?

	hxf98de04 := (*sliceHeader)(unsafe.Pointer(&x.remote_error_filename))
	hxf98de04.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.remote_error_filename))
	hxf98de04.Cap = 0x7fffffff
	// hxf98de04.Len = ?

	hxf57ac7e := (*sliceHeader)(unsafe.Pointer(&x.remote_input_filename))
	hxf57ac7e.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.remote_input_filename))
	hxf57ac7e.Cap = 0x7fffffff
	// hxf57ac7e.Len = ?

	x.local_fds = *Newslurm_step_io_fds_tRef(unsafe.Pointer(&x.ref17ff30d9.local_fds))
	x.gid = (uint32_t)(x.ref17ff30d9.gid)
	x.multi_prog = (bool)(x.ref17ff30d9.multi_prog)
	x.no_alloc = (bool)(x.ref17ff30d9.no_alloc)
	x.slurmd_debug = (uint32_t)(x.ref17ff30d9.slurmd_debug)
	x.node_offset = (uint32_t)(x.ref17ff30d9.node_offset)
	x.pack_jobid = (uint32_t)(x.ref17ff30d9.pack_jobid)
	x.pack_nnodes = (uint32_t)(x.ref17ff30d9.pack_nnodes)
	x.pack_ntasks = (uint32_t)(x.ref17ff30d9.pack_ntasks)
	hxfcb77fa := (*sliceHeader)(unsafe.Pointer(&x.pack_task_cnts))
	hxfcb77fa.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.pack_task_cnts))
	hxfcb77fa.Cap = 0x7fffffff
	// hxfcb77fa.Len = ?

	packSSUUint32_t(x.pack_tids, x.ref17ff30d9.pack_tids)
	x.pack_offset = (uint32_t)(x.ref17ff30d9.pack_offset)
	x.pack_task_offset = (uint32_t)(x.ref17ff30d9.pack_task_offset)
	hxfc08a21 := (*sliceHeader)(unsafe.Pointer(&x.pack_node_list))
	hxfc08a21.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.pack_node_list))
	hxfc08a21.Cap = 0x7fffffff
	// hxfc08a21.Len = ?

	x.parallel_debug = (bool)(x.ref17ff30d9.parallel_debug)
	x.profile = (uint32_t)(x.ref17ff30d9.profile)
	hxf90b062 := (*sliceHeader)(unsafe.Pointer(&x.task_prolog))
	hxf90b062.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.task_prolog))
	hxf90b062.Cap = 0x7fffffff
	// hxf90b062.Len = ?

	hxf57c7f8 := (*sliceHeader)(unsafe.Pointer(&x.task_epilog))
	hxf57c7f8.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.task_epilog))
	hxf57c7f8.Cap = 0x7fffffff
	// hxf57c7f8.Len = ?

	x.cpu_bind_type = (uint16_t)(x.ref17ff30d9.cpu_bind_type)
	hxffbeb31 := (*sliceHeader)(unsafe.Pointer(&x.cpu_bind))
	hxffbeb31.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.cpu_bind))
	hxffbeb31.Cap = 0x7fffffff
	// hxffbeb31.Len = ?

	x.cpu_freq_min = (uint32_t)(x.ref17ff30d9.cpu_freq_min)
	x.cpu_freq_max = (uint32_t)(x.ref17ff30d9.cpu_freq_max)
	x.cpu_freq_gov = (uint32_t)(x.ref17ff30d9.cpu_freq_gov)
	x.mem_bind_type = (uint16_t)(x.ref17ff30d9.mem_bind_type)
	hxf6a3e01 := (*sliceHeader)(unsafe.Pointer(&x.mem_bind))
	hxf6a3e01.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.mem_bind))
	hxf6a3e01.Cap = 0x7fffffff
	// hxf6a3e01.Len = ?

	x.accel_bind_type = (uint16_t)(x.ref17ff30d9.accel_bind_type)
	x.max_sockets = (uint16_t)(x.ref17ff30d9.max_sockets)
	x.max_cores = (uint16_t)(x.ref17ff30d9.max_cores)
	x.max_threads = (uint16_t)(x.ref17ff30d9.max_threads)
	x.cpus_per_task = (uint16_t)(x.ref17ff30d9.cpus_per_task)
	x.task_dist = (uint32_t)(x.ref17ff30d9.task_dist)
	hxf95ff39 := (*sliceHeader)(unsafe.Pointer(&x.partition))
	hxf95ff39.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.partition))
	hxf95ff39.Cap = 0x7fffffff
	// hxf95ff39.Len = ?

	x.preserve_env = (bool)(x.ref17ff30d9.preserve_env)
	hxf31b5a3 := (*sliceHeader)(unsafe.Pointer(&x.mpi_plugin_name))
	hxf31b5a3.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.mpi_plugin_name))
	hxf31b5a3.Cap = 0x7fffffff
	// hxf31b5a3.Len = ?

	x.open_mode = (uint8_t)(x.ref17ff30d9.open_mode)
	hxfb391ad := (*sliceHeader)(unsafe.Pointer(&x.acctg_freq))
	hxfb391ad.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.acctg_freq))
	hxfb391ad.Cap = 0x7fffffff
	// hxfb391ad.Len = ?

	x.pty = (bool)(x.ref17ff30d9.pty)
	hxf0fe205 := (*sliceHeader)(unsafe.Pointer(&x.ckpt_dir))
	hxf0fe205.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.ckpt_dir))
	hxf0fe205.Cap = 0x7fffffff
	// hxf0fe205.Len = ?

	hxfa71b27 := (*sliceHeader)(unsafe.Pointer(&x.restart_dir))
	hxfa71b27.Data = uintptr(unsafe.Pointer(x.ref17ff30d9.restart_dir))
	hxfa71b27.Cap = 0x7fffffff
	// hxfa71b27.Len = ?

	packSSByte(x.spank_job_env, x.ref17ff30d9.spank_job_env)
	x.spank_job_env_size = (uint32_t)(x.ref17ff30d9.spank_job_env_size)
}

// allocSlurm_step_launch_callbacks_tMemory allocates memory for type C.slurm_step_launch_callbacks_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_step_launch_callbacks_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_step_launch_callbacks_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_step_launch_callbacks_tValue = unsafe.Sizeof([1]C.slurm_step_launch_callbacks_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_step_launch_callbacks_t) Ref() *C.slurm_step_launch_callbacks_t {
	if x == nil {
		return nil
	}
	return x.refaf4fc3bc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_step_launch_callbacks_t) Free() {
	if x != nil && x.allocsaf4fc3bc != nil {
		x.allocsaf4fc3bc.(*cgoAllocMap).Free()
		x.refaf4fc3bc = nil
	}
}

// Newslurm_step_launch_callbacks_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_step_launch_callbacks_tRef(ref unsafe.Pointer) *slurm_step_launch_callbacks_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_step_launch_callbacks_t)
	obj.refaf4fc3bc = (*C.slurm_step_launch_callbacks_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_step_launch_callbacks_t) PassRef() (*C.slurm_step_launch_callbacks_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refaf4fc3bc != nil {
		return x.refaf4fc3bc, nil
	}
	memaf4fc3bc := allocSlurm_step_launch_callbacks_tMemory(1)
	refaf4fc3bc := (*C.slurm_step_launch_callbacks_t)(memaf4fc3bc)
	allocsaf4fc3bc := new(cgoAllocMap)
	allocsaf4fc3bc.Add(memaf4fc3bc)

	var cstep_complete_allocs *cgoAllocMap
	refaf4fc3bc.step_complete, cstep_complete_allocs = x.step_complete.PassRef()
	allocsaf4fc3bc.Borrow(cstep_complete_allocs)

	var cstep_signal_allocs *cgoAllocMap
	refaf4fc3bc.step_signal, cstep_signal_allocs = x.step_signal.PassRef()
	allocsaf4fc3bc.Borrow(cstep_signal_allocs)

	var cstep_timeout_allocs *cgoAllocMap
	refaf4fc3bc.step_timeout, cstep_timeout_allocs = x.step_timeout.PassRef()
	allocsaf4fc3bc.Borrow(cstep_timeout_allocs)

	var ctask_start_allocs *cgoAllocMap
	refaf4fc3bc.task_start, ctask_start_allocs = x.task_start.PassRef()
	allocsaf4fc3bc.Borrow(ctask_start_allocs)

	var ctask_finish_allocs *cgoAllocMap
	refaf4fc3bc.task_finish, ctask_finish_allocs = x.task_finish.PassRef()
	allocsaf4fc3bc.Borrow(ctask_finish_allocs)

	x.refaf4fc3bc = refaf4fc3bc
	x.allocsaf4fc3bc = allocsaf4fc3bc
	return refaf4fc3bc, allocsaf4fc3bc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_step_launch_callbacks_t) PassValue() (C.slurm_step_launch_callbacks_t, *cgoAllocMap) {
	if x.refaf4fc3bc != nil {
		return *x.refaf4fc3bc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_step_launch_callbacks_t) Deref() {
	if x.refaf4fc3bc == nil {
		return
	}
	x.step_complete = NewRef(unsafe.Pointer(x.refaf4fc3bc.step_complete))
	x.step_signal = NewRef(unsafe.Pointer(x.refaf4fc3bc.step_signal))
	x.step_timeout = NewRef(unsafe.Pointer(x.refaf4fc3bc.step_timeout))
	x.task_start = NewRef(unsafe.Pointer(x.refaf4fc3bc.task_start))
	x.task_finish = NewRef(unsafe.Pointer(x.refaf4fc3bc.task_finish))
}

// allocSlurm_allocation_callbacks_tMemory allocates memory for type C.slurm_allocation_callbacks_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_allocation_callbacks_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_allocation_callbacks_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_allocation_callbacks_tValue = unsafe.Sizeof([1]C.slurm_allocation_callbacks_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_allocation_callbacks_t) Ref() *C.slurm_allocation_callbacks_t {
	if x == nil {
		return nil
	}
	return x.ref681611eb
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_allocation_callbacks_t) Free() {
	if x != nil && x.allocs681611eb != nil {
		x.allocs681611eb.(*cgoAllocMap).Free()
		x.ref681611eb = nil
	}
}

// Newslurm_allocation_callbacks_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_allocation_callbacks_tRef(ref unsafe.Pointer) *slurm_allocation_callbacks_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_allocation_callbacks_t)
	obj.ref681611eb = (*C.slurm_allocation_callbacks_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_allocation_callbacks_t) PassRef() (*C.slurm_allocation_callbacks_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref681611eb != nil {
		return x.ref681611eb, nil
	}
	mem681611eb := allocSlurm_allocation_callbacks_tMemory(1)
	ref681611eb := (*C.slurm_allocation_callbacks_t)(mem681611eb)
	allocs681611eb := new(cgoAllocMap)
	allocs681611eb.Add(mem681611eb)

	var cping_allocs *cgoAllocMap
	ref681611eb.ping, cping_allocs = x.ping.PassRef()
	allocs681611eb.Borrow(cping_allocs)

	var cjob_complete_allocs *cgoAllocMap
	ref681611eb.job_complete, cjob_complete_allocs = x.job_complete.PassRef()
	allocs681611eb.Borrow(cjob_complete_allocs)

	var ctimeout_allocs *cgoAllocMap
	ref681611eb.timeout, ctimeout_allocs = x.timeout.PassRef()
	allocs681611eb.Borrow(ctimeout_allocs)

	var cuser_msg_allocs *cgoAllocMap
	ref681611eb.user_msg, cuser_msg_allocs = x.user_msg.PassRef()
	allocs681611eb.Borrow(cuser_msg_allocs)

	var cnode_fail_allocs *cgoAllocMap
	ref681611eb.node_fail, cnode_fail_allocs = x.node_fail.PassRef()
	allocs681611eb.Borrow(cnode_fail_allocs)

	var cjob_suspend_allocs *cgoAllocMap
	ref681611eb.job_suspend, cjob_suspend_allocs = x.job_suspend.PassRef()
	allocs681611eb.Borrow(cjob_suspend_allocs)

	x.ref681611eb = ref681611eb
	x.allocs681611eb = allocs681611eb
	return ref681611eb, allocs681611eb

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_allocation_callbacks_t) PassValue() (C.slurm_allocation_callbacks_t, *cgoAllocMap) {
	if x.ref681611eb != nil {
		return *x.ref681611eb, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_allocation_callbacks_t) Deref() {
	if x.ref681611eb == nil {
		return
	}
	x.ping = NewRef(unsafe.Pointer(x.ref681611eb.ping))
	x.job_complete = NewRef(unsafe.Pointer(x.ref681611eb.job_complete))
	x.timeout = NewRef(unsafe.Pointer(x.ref681611eb.timeout))
	x.user_msg = NewRef(unsafe.Pointer(x.ref681611eb.user_msg))
	x.node_fail = NewRef(unsafe.Pointer(x.ref681611eb.node_fail))
	x.job_suspend = NewRef(unsafe.Pointer(x.ref681611eb.job_suspend))
}

// allocSlurm_trigger_callbacks_tMemory allocates memory for type C.slurm_trigger_callbacks_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_trigger_callbacks_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_trigger_callbacks_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_trigger_callbacks_tValue = unsafe.Sizeof([1]C.slurm_trigger_callbacks_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_trigger_callbacks_t) Ref() *C.slurm_trigger_callbacks_t {
	if x == nil {
		return nil
	}
	return x.ref3554fae
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_trigger_callbacks_t) Free() {
	if x != nil && x.allocs3554fae != nil {
		x.allocs3554fae.(*cgoAllocMap).Free()
		x.ref3554fae = nil
	}
}

// Newslurm_trigger_callbacks_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_trigger_callbacks_tRef(ref unsafe.Pointer) *slurm_trigger_callbacks_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_trigger_callbacks_t)
	obj.ref3554fae = (*C.slurm_trigger_callbacks_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_trigger_callbacks_t) PassRef() (*C.slurm_trigger_callbacks_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3554fae != nil {
		return x.ref3554fae, nil
	}
	mem3554fae := allocSlurm_trigger_callbacks_tMemory(1)
	ref3554fae := (*C.slurm_trigger_callbacks_t)(mem3554fae)
	allocs3554fae := new(cgoAllocMap)
	allocs3554fae.Add(mem3554fae)

	var cacct_full_allocs *cgoAllocMap
	ref3554fae.acct_full, cacct_full_allocs = x.acct_full.PassRef()
	allocs3554fae.Borrow(cacct_full_allocs)

	var cdbd_fail_allocs *cgoAllocMap
	ref3554fae.dbd_fail, cdbd_fail_allocs = x.dbd_fail.PassRef()
	allocs3554fae.Borrow(cdbd_fail_allocs)

	var cdbd_resumed_allocs *cgoAllocMap
	ref3554fae.dbd_resumed, cdbd_resumed_allocs = x.dbd_resumed.PassRef()
	allocs3554fae.Borrow(cdbd_resumed_allocs)

	var cdb_fail_allocs *cgoAllocMap
	ref3554fae.db_fail, cdb_fail_allocs = x.db_fail.PassRef()
	allocs3554fae.Borrow(cdb_fail_allocs)

	var cdb_resumed_allocs *cgoAllocMap
	ref3554fae.db_resumed, cdb_resumed_allocs = x.db_resumed.PassRef()
	allocs3554fae.Borrow(cdb_resumed_allocs)

	x.ref3554fae = ref3554fae
	x.allocs3554fae = allocs3554fae
	return ref3554fae, allocs3554fae

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_trigger_callbacks_t) PassValue() (C.slurm_trigger_callbacks_t, *cgoAllocMap) {
	if x.ref3554fae != nil {
		return *x.ref3554fae, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_trigger_callbacks_t) Deref() {
	if x.ref3554fae == nil {
		return
	}
	x.acct_full = NewRef(unsafe.Pointer(x.ref3554fae.acct_full))
	x.dbd_fail = NewRef(unsafe.Pointer(x.ref3554fae.dbd_fail))
	x.dbd_resumed = NewRef(unsafe.Pointer(x.ref3554fae.dbd_resumed))
	x.db_fail = NewRef(unsafe.Pointer(x.ref3554fae.db_fail))
	x.db_resumed = NewRef(unsafe.Pointer(x.ref3554fae.db_resumed))
}

// allocJob_step_info_tMemory allocates memory for type C.job_step_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_step_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_step_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_step_info_tValue = unsafe.Sizeof([1]C.job_step_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_step_info_t) Ref() *C.job_step_info_t {
	if x == nil {
		return nil
	}
	return x.refc1c43052
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_step_info_t) Free() {
	if x != nil && x.allocsc1c43052 != nil {
		x.allocsc1c43052.(*cgoAllocMap).Free()
		x.refc1c43052 = nil
	}
}

// Newjob_step_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_step_info_tRef(ref unsafe.Pointer) *job_step_info_t {
	if ref == nil {
		return nil
	}
	obj := new(job_step_info_t)
	obj.refc1c43052 = (*C.job_step_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_step_info_t) PassRef() (*C.job_step_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc1c43052 != nil {
		return x.refc1c43052, nil
	}
	memc1c43052 := allocJob_step_info_tMemory(1)
	refc1c43052 := (*C.job_step_info_t)(memc1c43052)
	allocsc1c43052 := new(cgoAllocMap)
	allocsc1c43052.Add(memc1c43052)

	var carray_job_id_allocs *cgoAllocMap
	refc1c43052.array_job_id, carray_job_id_allocs = (C.uint32_t)(x.array_job_id), cgoAllocsUnknown
	allocsc1c43052.Borrow(carray_job_id_allocs)

	var carray_task_id_allocs *cgoAllocMap
	refc1c43052.array_task_id, carray_task_id_allocs = (C.uint32_t)(x.array_task_id), cgoAllocsUnknown
	allocsc1c43052.Borrow(carray_task_id_allocs)

	var cckpt_dir_allocs *cgoAllocMap
	refc1c43052.ckpt_dir, cckpt_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ckpt_dir)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(cckpt_dir_allocs)

	var cckpt_interval_allocs *cgoAllocMap
	refc1c43052.ckpt_interval, cckpt_interval_allocs = (C.uint16_t)(x.ckpt_interval), cgoAllocsUnknown
	allocsc1c43052.Borrow(cckpt_interval_allocs)

	var ccluster_allocs *cgoAllocMap
	refc1c43052.cluster, ccluster_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(ccluster_allocs)

	var cgres_allocs *cgoAllocMap
	refc1c43052.gres, cgres_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.gres)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(cgres_allocs)

	var cjob_id_allocs *cgoAllocMap
	refc1c43052.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocsc1c43052.Borrow(cjob_id_allocs)

	var cname_allocs *cgoAllocMap
	refc1c43052.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(cname_allocs)

	var cnetwork_allocs *cgoAllocMap
	refc1c43052.network, cnetwork_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.network)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(cnetwork_allocs)

	var cnodes_allocs *cgoAllocMap
	refc1c43052.nodes, cnodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.nodes)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(cnodes_allocs)

	var cnode_inx_allocs *cgoAllocMap
	refc1c43052.node_inx, cnode_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_inx)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(cnode_inx_allocs)

	var cnum_cpus_allocs *cgoAllocMap
	refc1c43052.num_cpus, cnum_cpus_allocs = (C.uint32_t)(x.num_cpus), cgoAllocsUnknown
	allocsc1c43052.Borrow(cnum_cpus_allocs)

	var ccpu_freq_min_allocs *cgoAllocMap
	refc1c43052.cpu_freq_min, ccpu_freq_min_allocs = (C.uint32_t)(x.cpu_freq_min), cgoAllocsUnknown
	allocsc1c43052.Borrow(ccpu_freq_min_allocs)

	var ccpu_freq_max_allocs *cgoAllocMap
	refc1c43052.cpu_freq_max, ccpu_freq_max_allocs = (C.uint32_t)(x.cpu_freq_max), cgoAllocsUnknown
	allocsc1c43052.Borrow(ccpu_freq_max_allocs)

	var ccpu_freq_gov_allocs *cgoAllocMap
	refc1c43052.cpu_freq_gov, ccpu_freq_gov_allocs = (C.uint32_t)(x.cpu_freq_gov), cgoAllocsUnknown
	allocsc1c43052.Borrow(ccpu_freq_gov_allocs)

	var cnum_tasks_allocs *cgoAllocMap
	refc1c43052.num_tasks, cnum_tasks_allocs = (C.uint32_t)(x.num_tasks), cgoAllocsUnknown
	allocsc1c43052.Borrow(cnum_tasks_allocs)

	var cpartition_allocs *cgoAllocMap
	refc1c43052.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partition)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(cpartition_allocs)

	var cresv_ports_allocs *cgoAllocMap
	refc1c43052.resv_ports, cresv_ports_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.resv_ports)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(cresv_ports_allocs)

	var crun_time_allocs *cgoAllocMap
	refc1c43052.run_time, crun_time_allocs = (C.time_t)(x.run_time), cgoAllocsUnknown
	allocsc1c43052.Borrow(crun_time_allocs)

	var cselect_jobinfo_allocs *cgoAllocMap
	refc1c43052.select_jobinfo, cselect_jobinfo_allocs = unpackSDynamic_plugin_data_t(x.select_jobinfo)
	allocsc1c43052.Borrow(cselect_jobinfo_allocs)

	var csrun_host_allocs *cgoAllocMap
	refc1c43052.srun_host, csrun_host_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.srun_host)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(csrun_host_allocs)

	var csrun_pid_allocs *cgoAllocMap
	refc1c43052.srun_pid, csrun_pid_allocs = (C.uint32_t)(x.srun_pid), cgoAllocsUnknown
	allocsc1c43052.Borrow(csrun_pid_allocs)

	var cstart_time_allocs *cgoAllocMap
	refc1c43052.start_time, cstart_time_allocs = (C.time_t)(x.start_time), cgoAllocsUnknown
	allocsc1c43052.Borrow(cstart_time_allocs)

	var cstart_protocol_ver_allocs *cgoAllocMap
	refc1c43052.start_protocol_ver, cstart_protocol_ver_allocs = (C.uint16_t)(x.start_protocol_ver), cgoAllocsUnknown
	allocsc1c43052.Borrow(cstart_protocol_ver_allocs)

	var cstate_allocs *cgoAllocMap
	refc1c43052.state, cstate_allocs = (C.uint32_t)(x.state), cgoAllocsUnknown
	allocsc1c43052.Borrow(cstate_allocs)

	var cstep_id_allocs *cgoAllocMap
	refc1c43052.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocsc1c43052.Borrow(cstep_id_allocs)

	var ctask_dist_allocs *cgoAllocMap
	refc1c43052.task_dist, ctask_dist_allocs = (C.uint32_t)(x.task_dist), cgoAllocsUnknown
	allocsc1c43052.Borrow(ctask_dist_allocs)

	var ctime_limit_allocs *cgoAllocMap
	refc1c43052.time_limit, ctime_limit_allocs = (C.uint32_t)(x.time_limit), cgoAllocsUnknown
	allocsc1c43052.Borrow(ctime_limit_allocs)

	var ctres_alloc_str_allocs *cgoAllocMap
	refc1c43052.tres_alloc_str, ctres_alloc_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_alloc_str)).Data)), cgoAllocsUnknown
	allocsc1c43052.Borrow(ctres_alloc_str_allocs)

	var cuser_id_allocs *cgoAllocMap
	refc1c43052.user_id, cuser_id_allocs = (C.uint32_t)(x.user_id), cgoAllocsUnknown
	allocsc1c43052.Borrow(cuser_id_allocs)

	x.refc1c43052 = refc1c43052
	x.allocsc1c43052 = allocsc1c43052
	return refc1c43052, allocsc1c43052

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_step_info_t) PassValue() (C.job_step_info_t, *cgoAllocMap) {
	if x.refc1c43052 != nil {
		return *x.refc1c43052, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_step_info_t) Deref() {
	if x.refc1c43052 == nil {
		return
	}
	x.array_job_id = (uint32_t)(x.refc1c43052.array_job_id)
	x.array_task_id = (uint32_t)(x.refc1c43052.array_task_id)
	hxf9e7b82 := (*sliceHeader)(unsafe.Pointer(&x.ckpt_dir))
	hxf9e7b82.Data = uintptr(unsafe.Pointer(x.refc1c43052.ckpt_dir))
	hxf9e7b82.Cap = 0x7fffffff
	// hxf9e7b82.Len = ?

	x.ckpt_interval = (uint16_t)(x.refc1c43052.ckpt_interval)
	hxf1018d1 := (*sliceHeader)(unsafe.Pointer(&x.cluster))
	hxf1018d1.Data = uintptr(unsafe.Pointer(x.refc1c43052.cluster))
	hxf1018d1.Cap = 0x7fffffff
	// hxf1018d1.Len = ?

	hxff52d5c := (*sliceHeader)(unsafe.Pointer(&x.gres))
	hxff52d5c.Data = uintptr(unsafe.Pointer(x.refc1c43052.gres))
	hxff52d5c.Cap = 0x7fffffff
	// hxff52d5c.Len = ?

	x.job_id = (uint32_t)(x.refc1c43052.job_id)
	hxf20b4dc := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf20b4dc.Data = uintptr(unsafe.Pointer(x.refc1c43052.name))
	hxf20b4dc.Cap = 0x7fffffff
	// hxf20b4dc.Len = ?

	hxf18f2b7 := (*sliceHeader)(unsafe.Pointer(&x.network))
	hxf18f2b7.Data = uintptr(unsafe.Pointer(x.refc1c43052.network))
	hxf18f2b7.Cap = 0x7fffffff
	// hxf18f2b7.Len = ?

	hxfd26ef9 := (*sliceHeader)(unsafe.Pointer(&x.nodes))
	hxfd26ef9.Data = uintptr(unsafe.Pointer(x.refc1c43052.nodes))
	hxfd26ef9.Cap = 0x7fffffff
	// hxfd26ef9.Len = ?

	hxf5c35cd := (*sliceHeader)(unsafe.Pointer(&x.node_inx))
	hxf5c35cd.Data = uintptr(unsafe.Pointer(x.refc1c43052.node_inx))
	hxf5c35cd.Cap = 0x7fffffff
	// hxf5c35cd.Len = ?

	x.num_cpus = (uint32_t)(x.refc1c43052.num_cpus)
	x.cpu_freq_min = (uint32_t)(x.refc1c43052.cpu_freq_min)
	x.cpu_freq_max = (uint32_t)(x.refc1c43052.cpu_freq_max)
	x.cpu_freq_gov = (uint32_t)(x.refc1c43052.cpu_freq_gov)
	x.num_tasks = (uint32_t)(x.refc1c43052.num_tasks)
	hxfc81a0f := (*sliceHeader)(unsafe.Pointer(&x.partition))
	hxfc81a0f.Data = uintptr(unsafe.Pointer(x.refc1c43052.partition))
	hxfc81a0f.Cap = 0x7fffffff
	// hxfc81a0f.Len = ?

	hxf45c1ee := (*sliceHeader)(unsafe.Pointer(&x.resv_ports))
	hxf45c1ee.Data = uintptr(unsafe.Pointer(x.refc1c43052.resv_ports))
	hxf45c1ee.Cap = 0x7fffffff
	// hxf45c1ee.Len = ?

	x.run_time = (time_t)(x.refc1c43052.run_time)
	packSDynamic_plugin_data_t(x.select_jobinfo, x.refc1c43052.select_jobinfo)
	hxfb97b1e := (*sliceHeader)(unsafe.Pointer(&x.srun_host))
	hxfb97b1e.Data = uintptr(unsafe.Pointer(x.refc1c43052.srun_host))
	hxfb97b1e.Cap = 0x7fffffff
	// hxfb97b1e.Len = ?

	x.srun_pid = (uint32_t)(x.refc1c43052.srun_pid)
	x.start_time = (time_t)(x.refc1c43052.start_time)
	x.start_protocol_ver = (uint16_t)(x.refc1c43052.start_protocol_ver)
	x.state = (uint32_t)(x.refc1c43052.state)
	x.step_id = (uint32_t)(x.refc1c43052.step_id)
	x.task_dist = (uint32_t)(x.refc1c43052.task_dist)
	x.time_limit = (uint32_t)(x.refc1c43052.time_limit)
	hxfc041c4 := (*sliceHeader)(unsafe.Pointer(&x.tres_alloc_str))
	hxfc041c4.Data = uintptr(unsafe.Pointer(x.refc1c43052.tres_alloc_str))
	hxfc041c4.Cap = 0x7fffffff
	// hxfc041c4.Len = ?

	x.user_id = (uint32_t)(x.refc1c43052.user_id)
}

// allocJob_step_info_response_msg_tMemory allocates memory for type C.job_step_info_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_step_info_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_step_info_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_step_info_response_msg_tValue = unsafe.Sizeof([1]C.job_step_info_response_msg_t{})

// unpackSJob_step_info_t transforms a sliced Go data structure into plain C format.
func unpackSJob_step_info_t(x []job_step_info_t) (unpacked *C.job_step_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_step_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_step_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_step_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_step_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSJob_step_info_t reads sliced Go data structure out from plain C format.
func packSJob_step_info_t(v []job_step_info_t, ptr0 *C.job_step_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfJob_step_info_tValue]C.job_step_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newjob_step_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_step_info_response_msg_t) Ref() *C.job_step_info_response_msg_t {
	if x == nil {
		return nil
	}
	return x.ref3f513490
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_step_info_response_msg_t) Free() {
	if x != nil && x.allocs3f513490 != nil {
		x.allocs3f513490.(*cgoAllocMap).Free()
		x.ref3f513490 = nil
	}
}

// Newjob_step_info_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_step_info_response_msg_tRef(ref unsafe.Pointer) *job_step_info_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(job_step_info_response_msg_t)
	obj.ref3f513490 = (*C.job_step_info_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_step_info_response_msg_t) PassRef() (*C.job_step_info_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3f513490 != nil {
		return x.ref3f513490, nil
	}
	mem3f513490 := allocJob_step_info_response_msg_tMemory(1)
	ref3f513490 := (*C.job_step_info_response_msg_t)(mem3f513490)
	allocs3f513490 := new(cgoAllocMap)
	allocs3f513490.Add(mem3f513490)

	var clast_update_allocs *cgoAllocMap
	ref3f513490.last_update, clast_update_allocs = (C.time_t)(x.last_update), cgoAllocsUnknown
	allocs3f513490.Borrow(clast_update_allocs)

	var cjob_step_count_allocs *cgoAllocMap
	ref3f513490.job_step_count, cjob_step_count_allocs = (C.uint32_t)(x.job_step_count), cgoAllocsUnknown
	allocs3f513490.Borrow(cjob_step_count_allocs)

	var cjob_steps_allocs *cgoAllocMap
	ref3f513490.job_steps, cjob_steps_allocs = unpackSJob_step_info_t(x.job_steps)
	allocs3f513490.Borrow(cjob_steps_allocs)

	x.ref3f513490 = ref3f513490
	x.allocs3f513490 = allocs3f513490
	return ref3f513490, allocs3f513490

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_step_info_response_msg_t) PassValue() (C.job_step_info_response_msg_t, *cgoAllocMap) {
	if x.ref3f513490 != nil {
		return *x.ref3f513490, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_step_info_response_msg_t) Deref() {
	if x.ref3f513490 == nil {
		return
	}
	x.last_update = (time_t)(x.ref3f513490.last_update)
	x.job_step_count = (uint32_t)(x.ref3f513490.job_step_count)
	packSJob_step_info_t(x.job_steps, x.ref3f513490.job_steps)
}

// allocJob_step_pids_tMemory allocates memory for type C.job_step_pids_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_step_pids_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_step_pids_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_step_pids_tValue = unsafe.Sizeof([1]C.job_step_pids_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_step_pids_t) Ref() *C.job_step_pids_t {
	if x == nil {
		return nil
	}
	return x.refe8082d8e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_step_pids_t) Free() {
	if x != nil && x.allocse8082d8e != nil {
		x.allocse8082d8e.(*cgoAllocMap).Free()
		x.refe8082d8e = nil
	}
}

// Newjob_step_pids_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_step_pids_tRef(ref unsafe.Pointer) *job_step_pids_t {
	if ref == nil {
		return nil
	}
	obj := new(job_step_pids_t)
	obj.refe8082d8e = (*C.job_step_pids_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_step_pids_t) PassRef() (*C.job_step_pids_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refe8082d8e != nil {
		return x.refe8082d8e, nil
	}
	meme8082d8e := allocJob_step_pids_tMemory(1)
	refe8082d8e := (*C.job_step_pids_t)(meme8082d8e)
	allocse8082d8e := new(cgoAllocMap)
	allocse8082d8e.Add(meme8082d8e)

	var cnode_name_allocs *cgoAllocMap
	refe8082d8e.node_name, cnode_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_name)).Data)), cgoAllocsUnknown
	allocse8082d8e.Borrow(cnode_name_allocs)

	var cpid_allocs *cgoAllocMap
	refe8082d8e.pid, cpid_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.pid)).Data)), cgoAllocsUnknown
	allocse8082d8e.Borrow(cpid_allocs)

	var cpid_cnt_allocs *cgoAllocMap
	refe8082d8e.pid_cnt, cpid_cnt_allocs = (C.uint32_t)(x.pid_cnt), cgoAllocsUnknown
	allocse8082d8e.Borrow(cpid_cnt_allocs)

	x.refe8082d8e = refe8082d8e
	x.allocse8082d8e = allocse8082d8e
	return refe8082d8e, allocse8082d8e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_step_pids_t) PassValue() (C.job_step_pids_t, *cgoAllocMap) {
	if x.refe8082d8e != nil {
		return *x.refe8082d8e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_step_pids_t) Deref() {
	if x.refe8082d8e == nil {
		return
	}
	hxfb5aa01 := (*sliceHeader)(unsafe.Pointer(&x.node_name))
	hxfb5aa01.Data = uintptr(unsafe.Pointer(x.refe8082d8e.node_name))
	hxfb5aa01.Cap = 0x7fffffff
	// hxfb5aa01.Len = ?

	hxfac6ee5 := (*sliceHeader)(unsafe.Pointer(&x.pid))
	hxfac6ee5.Data = uintptr(unsafe.Pointer(x.refe8082d8e.pid))
	hxfac6ee5.Cap = 0x7fffffff
	// hxfac6ee5.Len = ?

	x.pid_cnt = (uint32_t)(x.refe8082d8e.pid_cnt)
}

// allocJob_step_pids_response_msg_tMemory allocates memory for type C.job_step_pids_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_step_pids_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_step_pids_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_step_pids_response_msg_tValue = unsafe.Sizeof([1]C.job_step_pids_response_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_step_pids_response_msg_t) Ref() *C.job_step_pids_response_msg_t {
	if x == nil {
		return nil
	}
	return x.ref8647e5b2
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_step_pids_response_msg_t) Free() {
	if x != nil && x.allocs8647e5b2 != nil {
		x.allocs8647e5b2.(*cgoAllocMap).Free()
		x.ref8647e5b2 = nil
	}
}

// Newjob_step_pids_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_step_pids_response_msg_tRef(ref unsafe.Pointer) *job_step_pids_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(job_step_pids_response_msg_t)
	obj.ref8647e5b2 = (*C.job_step_pids_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_step_pids_response_msg_t) PassRef() (*C.job_step_pids_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref8647e5b2 != nil {
		return x.ref8647e5b2, nil
	}
	mem8647e5b2 := allocJob_step_pids_response_msg_tMemory(1)
	ref8647e5b2 := (*C.job_step_pids_response_msg_t)(mem8647e5b2)
	allocs8647e5b2 := new(cgoAllocMap)
	allocs8647e5b2.Add(mem8647e5b2)

	var cjob_id_allocs *cgoAllocMap
	ref8647e5b2.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs8647e5b2.Borrow(cjob_id_allocs)

	var cpid_list_allocs *cgoAllocMap
	ref8647e5b2.pid_list, cpid_list_allocs = *(*C.List)(unsafe.Pointer(&x.pid_list)), cgoAllocsUnknown
	allocs8647e5b2.Borrow(cpid_list_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref8647e5b2.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocs8647e5b2.Borrow(cstep_id_allocs)

	x.ref8647e5b2 = ref8647e5b2
	x.allocs8647e5b2 = allocs8647e5b2
	return ref8647e5b2, allocs8647e5b2

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_step_pids_response_msg_t) PassValue() (C.job_step_pids_response_msg_t, *cgoAllocMap) {
	if x.ref8647e5b2 != nil {
		return *x.ref8647e5b2, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_step_pids_response_msg_t) Deref() {
	if x.ref8647e5b2 == nil {
		return
	}
	x.job_id = (uint32_t)(x.ref8647e5b2.job_id)
	x.pid_list = *(*List)(unsafe.Pointer(&x.ref8647e5b2.pid_list))
	x.step_id = (uint32_t)(x.ref8647e5b2.step_id)
}

// allocJob_step_stat_tMemory allocates memory for type C.job_step_stat_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_step_stat_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_step_stat_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_step_stat_tValue = unsafe.Sizeof([1]C.job_step_stat_t{})

// unpackSJob_step_pids_t transforms a sliced Go data structure into plain C format.
func unpackSJob_step_pids_t(x []job_step_pids_t) (unpacked *C.job_step_pids_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_step_pids_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_step_pids_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_step_pids_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_step_pids_t)(unsafe.Pointer(h.Data))
	return
}

// packSJob_step_pids_t reads sliced Go data structure out from plain C format.
func packSJob_step_pids_t(v []job_step_pids_t, ptr0 *C.job_step_pids_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfJob_step_pids_tValue]C.job_step_pids_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newjob_step_pids_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_step_stat_t) Ref() *C.job_step_stat_t {
	if x == nil {
		return nil
	}
	return x.refc47deaa4
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_step_stat_t) Free() {
	if x != nil && x.allocsc47deaa4 != nil {
		x.allocsc47deaa4.(*cgoAllocMap).Free()
		x.refc47deaa4 = nil
	}
}

// Newjob_step_stat_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_step_stat_tRef(ref unsafe.Pointer) *job_step_stat_t {
	if ref == nil {
		return nil
	}
	obj := new(job_step_stat_t)
	obj.refc47deaa4 = (*C.job_step_stat_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_step_stat_t) PassRef() (*C.job_step_stat_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc47deaa4 != nil {
		return x.refc47deaa4, nil
	}
	memc47deaa4 := allocJob_step_stat_tMemory(1)
	refc47deaa4 := (*C.job_step_stat_t)(memc47deaa4)
	allocsc47deaa4 := new(cgoAllocMap)
	allocsc47deaa4.Add(memc47deaa4)

	var cjobacct_allocs *cgoAllocMap
	refc47deaa4.jobacct, cjobacct_allocs = (*C.jobacctinfo_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.jobacct)).Data)), cgoAllocsUnknown
	allocsc47deaa4.Borrow(cjobacct_allocs)

	var cnum_tasks_allocs *cgoAllocMap
	refc47deaa4.num_tasks, cnum_tasks_allocs = (C.uint32_t)(x.num_tasks), cgoAllocsUnknown
	allocsc47deaa4.Borrow(cnum_tasks_allocs)

	var creturn_code_allocs *cgoAllocMap
	refc47deaa4.return_code, creturn_code_allocs = (C.uint32_t)(x.return_code), cgoAllocsUnknown
	allocsc47deaa4.Borrow(creturn_code_allocs)

	var cstep_pids_allocs *cgoAllocMap
	refc47deaa4.step_pids, cstep_pids_allocs = unpackSJob_step_pids_t(x.step_pids)
	allocsc47deaa4.Borrow(cstep_pids_allocs)

	x.refc47deaa4 = refc47deaa4
	x.allocsc47deaa4 = allocsc47deaa4
	return refc47deaa4, allocsc47deaa4

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_step_stat_t) PassValue() (C.job_step_stat_t, *cgoAllocMap) {
	if x.refc47deaa4 != nil {
		return *x.refc47deaa4, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_step_stat_t) Deref() {
	if x.refc47deaa4 == nil {
		return
	}
	hxf8d479b := (*sliceHeader)(unsafe.Pointer(&x.jobacct))
	hxf8d479b.Data = uintptr(unsafe.Pointer(x.refc47deaa4.jobacct))
	hxf8d479b.Cap = 0x7fffffff
	// hxf8d479b.Len = ?

	x.num_tasks = (uint32_t)(x.refc47deaa4.num_tasks)
	x.return_code = (uint32_t)(x.refc47deaa4.return_code)
	packSJob_step_pids_t(x.step_pids, x.refc47deaa4.step_pids)
}

// allocJob_step_stat_response_msg_tMemory allocates memory for type C.job_step_stat_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_step_stat_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_step_stat_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_step_stat_response_msg_tValue = unsafe.Sizeof([1]C.job_step_stat_response_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_step_stat_response_msg_t) Ref() *C.job_step_stat_response_msg_t {
	if x == nil {
		return nil
	}
	return x.reff14ee411
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_step_stat_response_msg_t) Free() {
	if x != nil && x.allocsf14ee411 != nil {
		x.allocsf14ee411.(*cgoAllocMap).Free()
		x.reff14ee411 = nil
	}
}

// Newjob_step_stat_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_step_stat_response_msg_tRef(ref unsafe.Pointer) *job_step_stat_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(job_step_stat_response_msg_t)
	obj.reff14ee411 = (*C.job_step_stat_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_step_stat_response_msg_t) PassRef() (*C.job_step_stat_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff14ee411 != nil {
		return x.reff14ee411, nil
	}
	memf14ee411 := allocJob_step_stat_response_msg_tMemory(1)
	reff14ee411 := (*C.job_step_stat_response_msg_t)(memf14ee411)
	allocsf14ee411 := new(cgoAllocMap)
	allocsf14ee411.Add(memf14ee411)

	var cjob_id_allocs *cgoAllocMap
	reff14ee411.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocsf14ee411.Borrow(cjob_id_allocs)

	var cstats_list_allocs *cgoAllocMap
	reff14ee411.stats_list, cstats_list_allocs = *(*C.List)(unsafe.Pointer(&x.stats_list)), cgoAllocsUnknown
	allocsf14ee411.Borrow(cstats_list_allocs)

	var cstep_id_allocs *cgoAllocMap
	reff14ee411.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocsf14ee411.Borrow(cstep_id_allocs)

	x.reff14ee411 = reff14ee411
	x.allocsf14ee411 = allocsf14ee411
	return reff14ee411, allocsf14ee411

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_step_stat_response_msg_t) PassValue() (C.job_step_stat_response_msg_t, *cgoAllocMap) {
	if x.reff14ee411 != nil {
		return *x.reff14ee411, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_step_stat_response_msg_t) Deref() {
	if x.reff14ee411 == nil {
		return
	}
	x.job_id = (uint32_t)(x.reff14ee411.job_id)
	x.stats_list = *(*List)(unsafe.Pointer(&x.reff14ee411.stats_list))
	x.step_id = (uint32_t)(x.reff14ee411.step_id)
}

// allocNode_info_tMemory allocates memory for type C.node_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNode_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNode_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNode_info_tValue = unsafe.Sizeof([1]C.node_info_t{})

// unpackSAcct_gather_energy_t transforms a sliced Go data structure into plain C format.
func unpackSAcct_gather_energy_t(x []acct_gather_energy_t) (unpacked *C.acct_gather_energy_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.acct_gather_energy_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocAcct_gather_energy_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.acct_gather_energy_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.acct_gather_energy_t)(unsafe.Pointer(h.Data))
	return
}

// unpackSExt_sensors_data_t transforms a sliced Go data structure into plain C format.
func unpackSExt_sensors_data_t(x []ext_sensors_data_t) (unpacked *C.ext_sensors_data_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.ext_sensors_data_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocExt_sensors_data_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.ext_sensors_data_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.ext_sensors_data_t)(unsafe.Pointer(h.Data))
	return
}

// unpackSPower_mgmt_data_t transforms a sliced Go data structure into plain C format.
func unpackSPower_mgmt_data_t(x []power_mgmt_data_t) (unpacked *C.power_mgmt_data_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.power_mgmt_data_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPower_mgmt_data_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.power_mgmt_data_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.power_mgmt_data_t)(unsafe.Pointer(h.Data))
	return
}

// packSAcct_gather_energy_t reads sliced Go data structure out from plain C format.
func packSAcct_gather_energy_t(v []acct_gather_energy_t, ptr0 *C.acct_gather_energy_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfAcct_gather_energy_tValue]C.acct_gather_energy_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newacct_gather_energy_tRef(unsafe.Pointer(&ptr1))
	}
}

// packSExt_sensors_data_t reads sliced Go data structure out from plain C format.
func packSExt_sensors_data_t(v []ext_sensors_data_t, ptr0 *C.ext_sensors_data_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfExt_sensors_data_tValue]C.ext_sensors_data_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newext_sensors_data_tRef(unsafe.Pointer(&ptr1))
	}
}

// packSPower_mgmt_data_t reads sliced Go data structure out from plain C format.
func packSPower_mgmt_data_t(v []power_mgmt_data_t, ptr0 *C.power_mgmt_data_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPower_mgmt_data_tValue]C.power_mgmt_data_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newpower_mgmt_data_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *node_info_t) Ref() *C.node_info_t {
	if x == nil {
		return nil
	}
	return x.ref3ba4f3c4
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *node_info_t) Free() {
	if x != nil && x.allocs3ba4f3c4 != nil {
		x.allocs3ba4f3c4.(*cgoAllocMap).Free()
		x.ref3ba4f3c4 = nil
	}
}

// Newnode_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newnode_info_tRef(ref unsafe.Pointer) *node_info_t {
	if ref == nil {
		return nil
	}
	obj := new(node_info_t)
	obj.ref3ba4f3c4 = (*C.node_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *node_info_t) PassRef() (*C.node_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3ba4f3c4 != nil {
		return x.ref3ba4f3c4, nil
	}
	mem3ba4f3c4 := allocNode_info_tMemory(1)
	ref3ba4f3c4 := (*C.node_info_t)(mem3ba4f3c4)
	allocs3ba4f3c4 := new(cgoAllocMap)
	allocs3ba4f3c4.Add(mem3ba4f3c4)

	var carch_allocs *cgoAllocMap
	ref3ba4f3c4.arch, carch_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.arch)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(carch_allocs)

	var cboards_allocs *cgoAllocMap
	ref3ba4f3c4.boards, cboards_allocs = (C.uint16_t)(x.boards), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cboards_allocs)

	var cboot_time_allocs *cgoAllocMap
	ref3ba4f3c4.boot_time, cboot_time_allocs = (C.time_t)(x.boot_time), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cboot_time_allocs)

	var ccluster_name_allocs *cgoAllocMap
	ref3ba4f3c4.cluster_name, ccluster_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster_name)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(ccluster_name_allocs)

	var ccores_allocs *cgoAllocMap
	ref3ba4f3c4.cores, ccores_allocs = (C.uint16_t)(x.cores), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(ccores_allocs)

	var ccore_spec_cnt_allocs *cgoAllocMap
	ref3ba4f3c4.core_spec_cnt, ccore_spec_cnt_allocs = (C.uint16_t)(x.core_spec_cnt), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(ccore_spec_cnt_allocs)

	var ccpu_load_allocs *cgoAllocMap
	ref3ba4f3c4.cpu_load, ccpu_load_allocs = (C.uint32_t)(x.cpu_load), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(ccpu_load_allocs)

	var cfree_mem_allocs *cgoAllocMap
	ref3ba4f3c4.free_mem, cfree_mem_allocs = (C.uint64_t)(x.free_mem), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cfree_mem_allocs)

	var ccpus_allocs *cgoAllocMap
	ref3ba4f3c4.cpus, ccpus_allocs = (C.uint16_t)(x.cpus), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(ccpus_allocs)

	var ccpu_spec_list_allocs *cgoAllocMap
	ref3ba4f3c4.cpu_spec_list, ccpu_spec_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cpu_spec_list)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(ccpu_spec_list_allocs)

	var cenergy_allocs *cgoAllocMap
	ref3ba4f3c4.energy, cenergy_allocs = unpackSAcct_gather_energy_t(x.energy)
	allocs3ba4f3c4.Borrow(cenergy_allocs)

	var cext_sensors_allocs *cgoAllocMap
	ref3ba4f3c4.ext_sensors, cext_sensors_allocs = unpackSExt_sensors_data_t(x.ext_sensors)
	allocs3ba4f3c4.Borrow(cext_sensors_allocs)

	var cpower_allocs *cgoAllocMap
	ref3ba4f3c4.power, cpower_allocs = unpackSPower_mgmt_data_t(x.power)
	allocs3ba4f3c4.Borrow(cpower_allocs)

	var cfeatures_allocs *cgoAllocMap
	ref3ba4f3c4.features, cfeatures_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.features)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cfeatures_allocs)

	var cfeatures_act_allocs *cgoAllocMap
	ref3ba4f3c4.features_act, cfeatures_act_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.features_act)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cfeatures_act_allocs)

	var cgres_allocs *cgoAllocMap
	ref3ba4f3c4.gres, cgres_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.gres)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cgres_allocs)

	var cgres_drain_allocs *cgoAllocMap
	ref3ba4f3c4.gres_drain, cgres_drain_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.gres_drain)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cgres_drain_allocs)

	var cgres_used_allocs *cgoAllocMap
	ref3ba4f3c4.gres_used, cgres_used_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.gres_used)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cgres_used_allocs)

	var cmcs_label_allocs *cgoAllocMap
	ref3ba4f3c4.mcs_label, cmcs_label_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mcs_label)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cmcs_label_allocs)

	var cmem_spec_limit_allocs *cgoAllocMap
	ref3ba4f3c4.mem_spec_limit, cmem_spec_limit_allocs = (C.uint64_t)(x.mem_spec_limit), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cmem_spec_limit_allocs)

	var cname_allocs *cgoAllocMap
	ref3ba4f3c4.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cname_allocs)

	var cnode_addr_allocs *cgoAllocMap
	ref3ba4f3c4.node_addr, cnode_addr_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_addr)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cnode_addr_allocs)

	var cnode_hostname_allocs *cgoAllocMap
	ref3ba4f3c4.node_hostname, cnode_hostname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_hostname)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cnode_hostname_allocs)

	var cnode_state_allocs *cgoAllocMap
	ref3ba4f3c4.node_state, cnode_state_allocs = (C.uint32_t)(x.node_state), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cnode_state_allocs)

	var cos_allocs *cgoAllocMap
	ref3ba4f3c4.os, cos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.os)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cos_allocs)

	var cowner_allocs *cgoAllocMap
	ref3ba4f3c4.owner, cowner_allocs = (C.uint32_t)(x.owner), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cowner_allocs)

	var cpartitions_allocs *cgoAllocMap
	ref3ba4f3c4.partitions, cpartitions_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partitions)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cpartitions_allocs)

	var cport_allocs *cgoAllocMap
	ref3ba4f3c4.port, cport_allocs = (C.uint16_t)(x.port), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cport_allocs)

	var creal_memory_allocs *cgoAllocMap
	ref3ba4f3c4.real_memory, creal_memory_allocs = (C.uint64_t)(x.real_memory), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(creal_memory_allocs)

	var creason_allocs *cgoAllocMap
	ref3ba4f3c4.reason, creason_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.reason)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(creason_allocs)

	var creason_time_allocs *cgoAllocMap
	ref3ba4f3c4.reason_time, creason_time_allocs = (C.time_t)(x.reason_time), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(creason_time_allocs)

	var creason_uid_allocs *cgoAllocMap
	ref3ba4f3c4.reason_uid, creason_uid_allocs = (C.uint32_t)(x.reason_uid), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(creason_uid_allocs)

	var cselect_nodeinfo_allocs *cgoAllocMap
	ref3ba4f3c4.select_nodeinfo, cselect_nodeinfo_allocs = unpackSDynamic_plugin_data_t(x.select_nodeinfo)
	allocs3ba4f3c4.Borrow(cselect_nodeinfo_allocs)

	var cslurmd_start_time_allocs *cgoAllocMap
	ref3ba4f3c4.slurmd_start_time, cslurmd_start_time_allocs = (C.time_t)(x.slurmd_start_time), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cslurmd_start_time_allocs)

	var csockets_allocs *cgoAllocMap
	ref3ba4f3c4.sockets, csockets_allocs = (C.uint16_t)(x.sockets), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(csockets_allocs)

	var cthreads_allocs *cgoAllocMap
	ref3ba4f3c4.threads, cthreads_allocs = (C.uint16_t)(x.threads), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cthreads_allocs)

	var ctmp_disk_allocs *cgoAllocMap
	ref3ba4f3c4.tmp_disk, ctmp_disk_allocs = (C.uint32_t)(x.tmp_disk), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(ctmp_disk_allocs)

	var cweight_allocs *cgoAllocMap
	ref3ba4f3c4.weight, cweight_allocs = (C.uint32_t)(x.weight), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cweight_allocs)

	var ctres_fmt_str_allocs *cgoAllocMap
	ref3ba4f3c4.tres_fmt_str, ctres_fmt_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_fmt_str)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(ctres_fmt_str_allocs)

	var cversion_allocs *cgoAllocMap
	ref3ba4f3c4.version, cversion_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.version)).Data)), cgoAllocsUnknown
	allocs3ba4f3c4.Borrow(cversion_allocs)

	x.ref3ba4f3c4 = ref3ba4f3c4
	x.allocs3ba4f3c4 = allocs3ba4f3c4
	return ref3ba4f3c4, allocs3ba4f3c4

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x node_info_t) PassValue() (C.node_info_t, *cgoAllocMap) {
	if x.ref3ba4f3c4 != nil {
		return *x.ref3ba4f3c4, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *node_info_t) Deref() {
	if x.ref3ba4f3c4 == nil {
		return
	}
	hxfb6a89c := (*sliceHeader)(unsafe.Pointer(&x.arch))
	hxfb6a89c.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.arch))
	hxfb6a89c.Cap = 0x7fffffff
	// hxfb6a89c.Len = ?

	x.boards = (uint16_t)(x.ref3ba4f3c4.boards)
	x.boot_time = (time_t)(x.ref3ba4f3c4.boot_time)
	hxf9aa700 := (*sliceHeader)(unsafe.Pointer(&x.cluster_name))
	hxf9aa700.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.cluster_name))
	hxf9aa700.Cap = 0x7fffffff
	// hxf9aa700.Len = ?

	x.cores = (uint16_t)(x.ref3ba4f3c4.cores)
	x.core_spec_cnt = (uint16_t)(x.ref3ba4f3c4.core_spec_cnt)
	x.cpu_load = (uint32_t)(x.ref3ba4f3c4.cpu_load)
	x.free_mem = (uint64_t)(x.ref3ba4f3c4.free_mem)
	x.cpus = (uint16_t)(x.ref3ba4f3c4.cpus)
	hxf8a0e97 := (*sliceHeader)(unsafe.Pointer(&x.cpu_spec_list))
	hxf8a0e97.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.cpu_spec_list))
	hxf8a0e97.Cap = 0x7fffffff
	// hxf8a0e97.Len = ?

	packSAcct_gather_energy_t(x.energy, x.ref3ba4f3c4.energy)
	packSExt_sensors_data_t(x.ext_sensors, x.ref3ba4f3c4.ext_sensors)
	packSPower_mgmt_data_t(x.power, x.ref3ba4f3c4.power)
	hxf7f9e99 := (*sliceHeader)(unsafe.Pointer(&x.features))
	hxf7f9e99.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.features))
	hxf7f9e99.Cap = 0x7fffffff
	// hxf7f9e99.Len = ?

	hxf65874b := (*sliceHeader)(unsafe.Pointer(&x.features_act))
	hxf65874b.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.features_act))
	hxf65874b.Cap = 0x7fffffff
	// hxf65874b.Len = ?

	hxfdd8619 := (*sliceHeader)(unsafe.Pointer(&x.gres))
	hxfdd8619.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.gres))
	hxfdd8619.Cap = 0x7fffffff
	// hxfdd8619.Len = ?

	hxfd898f3 := (*sliceHeader)(unsafe.Pointer(&x.gres_drain))
	hxfd898f3.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.gres_drain))
	hxfd898f3.Cap = 0x7fffffff
	// hxfd898f3.Len = ?

	hxf40d6d7 := (*sliceHeader)(unsafe.Pointer(&x.gres_used))
	hxf40d6d7.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.gres_used))
	hxf40d6d7.Cap = 0x7fffffff
	// hxf40d6d7.Len = ?

	hxfb69baf := (*sliceHeader)(unsafe.Pointer(&x.mcs_label))
	hxfb69baf.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.mcs_label))
	hxfb69baf.Cap = 0x7fffffff
	// hxfb69baf.Len = ?

	x.mem_spec_limit = (uint64_t)(x.ref3ba4f3c4.mem_spec_limit)
	hxf6d059c := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf6d059c.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.name))
	hxf6d059c.Cap = 0x7fffffff
	// hxf6d059c.Len = ?

	hxf34c802 := (*sliceHeader)(unsafe.Pointer(&x.node_addr))
	hxf34c802.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.node_addr))
	hxf34c802.Cap = 0x7fffffff
	// hxf34c802.Len = ?

	hxfed1236 := (*sliceHeader)(unsafe.Pointer(&x.node_hostname))
	hxfed1236.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.node_hostname))
	hxfed1236.Cap = 0x7fffffff
	// hxfed1236.Len = ?

	x.node_state = (uint32_t)(x.ref3ba4f3c4.node_state)
	hxf8b14c9 := (*sliceHeader)(unsafe.Pointer(&x.os))
	hxf8b14c9.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.os))
	hxf8b14c9.Cap = 0x7fffffff
	// hxf8b14c9.Len = ?

	x.owner = (uint32_t)(x.ref3ba4f3c4.owner)
	hxf0ba811 := (*sliceHeader)(unsafe.Pointer(&x.partitions))
	hxf0ba811.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.partitions))
	hxf0ba811.Cap = 0x7fffffff
	// hxf0ba811.Len = ?

	x.port = (uint16_t)(x.ref3ba4f3c4.port)
	x.real_memory = (uint64_t)(x.ref3ba4f3c4.real_memory)
	hxf5977e7 := (*sliceHeader)(unsafe.Pointer(&x.reason))
	hxf5977e7.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.reason))
	hxf5977e7.Cap = 0x7fffffff
	// hxf5977e7.Len = ?

	x.reason_time = (time_t)(x.ref3ba4f3c4.reason_time)
	x.reason_uid = (uint32_t)(x.ref3ba4f3c4.reason_uid)
	packSDynamic_plugin_data_t(x.select_nodeinfo, x.ref3ba4f3c4.select_nodeinfo)
	x.slurmd_start_time = (time_t)(x.ref3ba4f3c4.slurmd_start_time)
	x.sockets = (uint16_t)(x.ref3ba4f3c4.sockets)
	x.threads = (uint16_t)(x.ref3ba4f3c4.threads)
	x.tmp_disk = (uint32_t)(x.ref3ba4f3c4.tmp_disk)
	x.weight = (uint32_t)(x.ref3ba4f3c4.weight)
	hxf8e3b4d := (*sliceHeader)(unsafe.Pointer(&x.tres_fmt_str))
	hxf8e3b4d.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.tres_fmt_str))
	hxf8e3b4d.Cap = 0x7fffffff
	// hxf8e3b4d.Len = ?

	hxf9c745d := (*sliceHeader)(unsafe.Pointer(&x.version))
	hxf9c745d.Data = uintptr(unsafe.Pointer(x.ref3ba4f3c4.version))
	hxf9c745d.Cap = 0x7fffffff
	// hxf9c745d.Len = ?

}

// allocNode_info_msg_tMemory allocates memory for type C.node_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNode_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNode_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNode_info_msg_tValue = unsafe.Sizeof([1]C.node_info_msg_t{})

// unpackSNode_info_t transforms a sliced Go data structure into plain C format.
func unpackSNode_info_t(x []node_info_t) (unpacked *C.node_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.node_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocNode_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.node_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.node_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSNode_info_t reads sliced Go data structure out from plain C format.
func packSNode_info_t(v []node_info_t, ptr0 *C.node_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfNode_info_tValue]C.node_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newnode_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *node_info_msg_t) Ref() *C.node_info_msg_t {
	if x == nil {
		return nil
	}
	return x.ref27202897
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *node_info_msg_t) Free() {
	if x != nil && x.allocs27202897 != nil {
		x.allocs27202897.(*cgoAllocMap).Free()
		x.ref27202897 = nil
	}
}

// Newnode_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newnode_info_msg_tRef(ref unsafe.Pointer) *node_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(node_info_msg_t)
	obj.ref27202897 = (*C.node_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *node_info_msg_t) PassRef() (*C.node_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref27202897 != nil {
		return x.ref27202897, nil
	}
	mem27202897 := allocNode_info_msg_tMemory(1)
	ref27202897 := (*C.node_info_msg_t)(mem27202897)
	allocs27202897 := new(cgoAllocMap)
	allocs27202897.Add(mem27202897)

	var clast_update_allocs *cgoAllocMap
	ref27202897.last_update, clast_update_allocs = (C.time_t)(x.last_update), cgoAllocsUnknown
	allocs27202897.Borrow(clast_update_allocs)

	var cnode_scaling_allocs *cgoAllocMap
	ref27202897.node_scaling, cnode_scaling_allocs = (C.uint32_t)(x.node_scaling), cgoAllocsUnknown
	allocs27202897.Borrow(cnode_scaling_allocs)

	var crecord_count_allocs *cgoAllocMap
	ref27202897.record_count, crecord_count_allocs = (C.uint32_t)(x.record_count), cgoAllocsUnknown
	allocs27202897.Borrow(crecord_count_allocs)

	var cnode_array_allocs *cgoAllocMap
	ref27202897.node_array, cnode_array_allocs = unpackSNode_info_t(x.node_array)
	allocs27202897.Borrow(cnode_array_allocs)

	x.ref27202897 = ref27202897
	x.allocs27202897 = allocs27202897
	return ref27202897, allocs27202897

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x node_info_msg_t) PassValue() (C.node_info_msg_t, *cgoAllocMap) {
	if x.ref27202897 != nil {
		return *x.ref27202897, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *node_info_msg_t) Deref() {
	if x.ref27202897 == nil {
		return
	}
	x.last_update = (time_t)(x.ref27202897.last_update)
	x.node_scaling = (uint32_t)(x.ref27202897.node_scaling)
	x.record_count = (uint32_t)(x.ref27202897.record_count)
	packSNode_info_t(x.node_array, x.ref27202897.node_array)
}

// allocFront_end_info_tMemory allocates memory for type C.front_end_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFront_end_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFront_end_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfFront_end_info_tValue = unsafe.Sizeof([1]C.front_end_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *front_end_info_t) Ref() *C.front_end_info_t {
	if x == nil {
		return nil
	}
	return x.refd6c05f8c
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *front_end_info_t) Free() {
	if x != nil && x.allocsd6c05f8c != nil {
		x.allocsd6c05f8c.(*cgoAllocMap).Free()
		x.refd6c05f8c = nil
	}
}

// Newfront_end_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newfront_end_info_tRef(ref unsafe.Pointer) *front_end_info_t {
	if ref == nil {
		return nil
	}
	obj := new(front_end_info_t)
	obj.refd6c05f8c = (*C.front_end_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *front_end_info_t) PassRef() (*C.front_end_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd6c05f8c != nil {
		return x.refd6c05f8c, nil
	}
	memd6c05f8c := allocFront_end_info_tMemory(1)
	refd6c05f8c := (*C.front_end_info_t)(memd6c05f8c)
	allocsd6c05f8c := new(cgoAllocMap)
	allocsd6c05f8c.Add(memd6c05f8c)

	var callow_groups_allocs *cgoAllocMap
	refd6c05f8c.allow_groups, callow_groups_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_groups)).Data)), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(callow_groups_allocs)

	var callow_users_allocs *cgoAllocMap
	refd6c05f8c.allow_users, callow_users_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_users)).Data)), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(callow_users_allocs)

	var cboot_time_allocs *cgoAllocMap
	refd6c05f8c.boot_time, cboot_time_allocs = (C.time_t)(x.boot_time), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(cboot_time_allocs)

	var cdeny_groups_allocs *cgoAllocMap
	refd6c05f8c.deny_groups, cdeny_groups_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.deny_groups)).Data)), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(cdeny_groups_allocs)

	var cdeny_users_allocs *cgoAllocMap
	refd6c05f8c.deny_users, cdeny_users_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.deny_users)).Data)), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(cdeny_users_allocs)

	var cname_allocs *cgoAllocMap
	refd6c05f8c.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(cname_allocs)

	var cnode_state_allocs *cgoAllocMap
	refd6c05f8c.node_state, cnode_state_allocs = (C.uint32_t)(x.node_state), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(cnode_state_allocs)

	var creason_allocs *cgoAllocMap
	refd6c05f8c.reason, creason_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.reason)).Data)), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(creason_allocs)

	var creason_time_allocs *cgoAllocMap
	refd6c05f8c.reason_time, creason_time_allocs = (C.time_t)(x.reason_time), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(creason_time_allocs)

	var creason_uid_allocs *cgoAllocMap
	refd6c05f8c.reason_uid, creason_uid_allocs = (C.uint32_t)(x.reason_uid), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(creason_uid_allocs)

	var cslurmd_start_time_allocs *cgoAllocMap
	refd6c05f8c.slurmd_start_time, cslurmd_start_time_allocs = (C.time_t)(x.slurmd_start_time), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(cslurmd_start_time_allocs)

	var cversion_allocs *cgoAllocMap
	refd6c05f8c.version, cversion_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.version)).Data)), cgoAllocsUnknown
	allocsd6c05f8c.Borrow(cversion_allocs)

	x.refd6c05f8c = refd6c05f8c
	x.allocsd6c05f8c = allocsd6c05f8c
	return refd6c05f8c, allocsd6c05f8c

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x front_end_info_t) PassValue() (C.front_end_info_t, *cgoAllocMap) {
	if x.refd6c05f8c != nil {
		return *x.refd6c05f8c, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *front_end_info_t) Deref() {
	if x.refd6c05f8c == nil {
		return
	}
	hxf33c4e5 := (*sliceHeader)(unsafe.Pointer(&x.allow_groups))
	hxf33c4e5.Data = uintptr(unsafe.Pointer(x.refd6c05f8c.allow_groups))
	hxf33c4e5.Cap = 0x7fffffff
	// hxf33c4e5.Len = ?

	hxf52f7e7 := (*sliceHeader)(unsafe.Pointer(&x.allow_users))
	hxf52f7e7.Data = uintptr(unsafe.Pointer(x.refd6c05f8c.allow_users))
	hxf52f7e7.Cap = 0x7fffffff
	// hxf52f7e7.Len = ?

	x.boot_time = (time_t)(x.refd6c05f8c.boot_time)
	hxf9ef33c := (*sliceHeader)(unsafe.Pointer(&x.deny_groups))
	hxf9ef33c.Data = uintptr(unsafe.Pointer(x.refd6c05f8c.deny_groups))
	hxf9ef33c.Cap = 0x7fffffff
	// hxf9ef33c.Len = ?

	hxfb07884 := (*sliceHeader)(unsafe.Pointer(&x.deny_users))
	hxfb07884.Data = uintptr(unsafe.Pointer(x.refd6c05f8c.deny_users))
	hxfb07884.Cap = 0x7fffffff
	// hxfb07884.Len = ?

	hxf98f69d := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf98f69d.Data = uintptr(unsafe.Pointer(x.refd6c05f8c.name))
	hxf98f69d.Cap = 0x7fffffff
	// hxf98f69d.Len = ?

	x.node_state = (uint32_t)(x.refd6c05f8c.node_state)
	hxf4aecb4 := (*sliceHeader)(unsafe.Pointer(&x.reason))
	hxf4aecb4.Data = uintptr(unsafe.Pointer(x.refd6c05f8c.reason))
	hxf4aecb4.Cap = 0x7fffffff
	// hxf4aecb4.Len = ?

	x.reason_time = (time_t)(x.refd6c05f8c.reason_time)
	x.reason_uid = (uint32_t)(x.refd6c05f8c.reason_uid)
	x.slurmd_start_time = (time_t)(x.refd6c05f8c.slurmd_start_time)
	hxf3d886e := (*sliceHeader)(unsafe.Pointer(&x.version))
	hxf3d886e.Data = uintptr(unsafe.Pointer(x.refd6c05f8c.version))
	hxf3d886e.Cap = 0x7fffffff
	// hxf3d886e.Len = ?

}

// allocFront_end_info_msg_tMemory allocates memory for type C.front_end_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFront_end_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFront_end_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfFront_end_info_msg_tValue = unsafe.Sizeof([1]C.front_end_info_msg_t{})

// unpackSFront_end_info_t transforms a sliced Go data structure into plain C format.
func unpackSFront_end_info_t(x []front_end_info_t) (unpacked *C.front_end_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.front_end_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocFront_end_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.front_end_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.front_end_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSFront_end_info_t reads sliced Go data structure out from plain C format.
func packSFront_end_info_t(v []front_end_info_t, ptr0 *C.front_end_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfFront_end_info_tValue]C.front_end_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newfront_end_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *front_end_info_msg_t) Ref() *C.front_end_info_msg_t {
	if x == nil {
		return nil
	}
	return x.ref76c8cc27
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *front_end_info_msg_t) Free() {
	if x != nil && x.allocs76c8cc27 != nil {
		x.allocs76c8cc27.(*cgoAllocMap).Free()
		x.ref76c8cc27 = nil
	}
}

// Newfront_end_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newfront_end_info_msg_tRef(ref unsafe.Pointer) *front_end_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(front_end_info_msg_t)
	obj.ref76c8cc27 = (*C.front_end_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *front_end_info_msg_t) PassRef() (*C.front_end_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref76c8cc27 != nil {
		return x.ref76c8cc27, nil
	}
	mem76c8cc27 := allocFront_end_info_msg_tMemory(1)
	ref76c8cc27 := (*C.front_end_info_msg_t)(mem76c8cc27)
	allocs76c8cc27 := new(cgoAllocMap)
	allocs76c8cc27.Add(mem76c8cc27)

	var clast_update_allocs *cgoAllocMap
	ref76c8cc27.last_update, clast_update_allocs = (C.time_t)(x.last_update), cgoAllocsUnknown
	allocs76c8cc27.Borrow(clast_update_allocs)

	var crecord_count_allocs *cgoAllocMap
	ref76c8cc27.record_count, crecord_count_allocs = (C.uint32_t)(x.record_count), cgoAllocsUnknown
	allocs76c8cc27.Borrow(crecord_count_allocs)

	var cfront_end_array_allocs *cgoAllocMap
	ref76c8cc27.front_end_array, cfront_end_array_allocs = unpackSFront_end_info_t(x.front_end_array)
	allocs76c8cc27.Borrow(cfront_end_array_allocs)

	x.ref76c8cc27 = ref76c8cc27
	x.allocs76c8cc27 = allocs76c8cc27
	return ref76c8cc27, allocs76c8cc27

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x front_end_info_msg_t) PassValue() (C.front_end_info_msg_t, *cgoAllocMap) {
	if x.ref76c8cc27 != nil {
		return *x.ref76c8cc27, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *front_end_info_msg_t) Deref() {
	if x.ref76c8cc27 == nil {
		return
	}
	x.last_update = (time_t)(x.ref76c8cc27.last_update)
	x.record_count = (uint32_t)(x.ref76c8cc27.record_count)
	packSFront_end_info_t(x.front_end_array, x.ref76c8cc27.front_end_array)
}

// allocTopo_info_tMemory allocates memory for type C.topo_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocTopo_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfTopo_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfTopo_info_tValue = unsafe.Sizeof([1]C.topo_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *topo_info_t) Ref() *C.topo_info_t {
	if x == nil {
		return nil
	}
	return x.ref1711603e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *topo_info_t) Free() {
	if x != nil && x.allocs1711603e != nil {
		x.allocs1711603e.(*cgoAllocMap).Free()
		x.ref1711603e = nil
	}
}

// Newtopo_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newtopo_info_tRef(ref unsafe.Pointer) *topo_info_t {
	if ref == nil {
		return nil
	}
	obj := new(topo_info_t)
	obj.ref1711603e = (*C.topo_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *topo_info_t) PassRef() (*C.topo_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref1711603e != nil {
		return x.ref1711603e, nil
	}
	mem1711603e := allocTopo_info_tMemory(1)
	ref1711603e := (*C.topo_info_t)(mem1711603e)
	allocs1711603e := new(cgoAllocMap)
	allocs1711603e.Add(mem1711603e)

	var clevel_allocs *cgoAllocMap
	ref1711603e.level, clevel_allocs = (C.uint16_t)(x.level), cgoAllocsUnknown
	allocs1711603e.Borrow(clevel_allocs)

	var clink_speed_allocs *cgoAllocMap
	ref1711603e.link_speed, clink_speed_allocs = (C.uint32_t)(x.link_speed), cgoAllocsUnknown
	allocs1711603e.Borrow(clink_speed_allocs)

	var cname_allocs *cgoAllocMap
	ref1711603e.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs1711603e.Borrow(cname_allocs)

	var cnodes_allocs *cgoAllocMap
	ref1711603e.nodes, cnodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.nodes)).Data)), cgoAllocsUnknown
	allocs1711603e.Borrow(cnodes_allocs)

	var cswitches_allocs *cgoAllocMap
	ref1711603e.switches, cswitches_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.switches)).Data)), cgoAllocsUnknown
	allocs1711603e.Borrow(cswitches_allocs)

	x.ref1711603e = ref1711603e
	x.allocs1711603e = allocs1711603e
	return ref1711603e, allocs1711603e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x topo_info_t) PassValue() (C.topo_info_t, *cgoAllocMap) {
	if x.ref1711603e != nil {
		return *x.ref1711603e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *topo_info_t) Deref() {
	if x.ref1711603e == nil {
		return
	}
	x.level = (uint16_t)(x.ref1711603e.level)
	x.link_speed = (uint32_t)(x.ref1711603e.link_speed)
	hxf49b243 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf49b243.Data = uintptr(unsafe.Pointer(x.ref1711603e.name))
	hxf49b243.Cap = 0x7fffffff
	// hxf49b243.Len = ?

	hxf92e87e := (*sliceHeader)(unsafe.Pointer(&x.nodes))
	hxf92e87e.Data = uintptr(unsafe.Pointer(x.ref1711603e.nodes))
	hxf92e87e.Cap = 0x7fffffff
	// hxf92e87e.Len = ?

	hxfddeff8 := (*sliceHeader)(unsafe.Pointer(&x.switches))
	hxfddeff8.Data = uintptr(unsafe.Pointer(x.ref1711603e.switches))
	hxfddeff8.Cap = 0x7fffffff
	// hxfddeff8.Len = ?

}

// allocTopo_info_response_msg_tMemory allocates memory for type C.topo_info_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocTopo_info_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfTopo_info_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfTopo_info_response_msg_tValue = unsafe.Sizeof([1]C.topo_info_response_msg_t{})

// unpackSTopo_info_t transforms a sliced Go data structure into plain C format.
func unpackSTopo_info_t(x []topo_info_t) (unpacked *C.topo_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.topo_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocTopo_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.topo_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.topo_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSTopo_info_t reads sliced Go data structure out from plain C format.
func packSTopo_info_t(v []topo_info_t, ptr0 *C.topo_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfTopo_info_tValue]C.topo_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newtopo_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *topo_info_response_msg_t) Ref() *C.topo_info_response_msg_t {
	if x == nil {
		return nil
	}
	return x.refbb64b6ce
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *topo_info_response_msg_t) Free() {
	if x != nil && x.allocsbb64b6ce != nil {
		x.allocsbb64b6ce.(*cgoAllocMap).Free()
		x.refbb64b6ce = nil
	}
}

// Newtopo_info_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newtopo_info_response_msg_tRef(ref unsafe.Pointer) *topo_info_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(topo_info_response_msg_t)
	obj.refbb64b6ce = (*C.topo_info_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *topo_info_response_msg_t) PassRef() (*C.topo_info_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbb64b6ce != nil {
		return x.refbb64b6ce, nil
	}
	membb64b6ce := allocTopo_info_response_msg_tMemory(1)
	refbb64b6ce := (*C.topo_info_response_msg_t)(membb64b6ce)
	allocsbb64b6ce := new(cgoAllocMap)
	allocsbb64b6ce.Add(membb64b6ce)

	var crecord_count_allocs *cgoAllocMap
	refbb64b6ce.record_count, crecord_count_allocs = (C.uint32_t)(x.record_count), cgoAllocsUnknown
	allocsbb64b6ce.Borrow(crecord_count_allocs)

	var ctopo_array_allocs *cgoAllocMap
	refbb64b6ce.topo_array, ctopo_array_allocs = unpackSTopo_info_t(x.topo_array)
	allocsbb64b6ce.Borrow(ctopo_array_allocs)

	x.refbb64b6ce = refbb64b6ce
	x.allocsbb64b6ce = allocsbb64b6ce
	return refbb64b6ce, allocsbb64b6ce

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x topo_info_response_msg_t) PassValue() (C.topo_info_response_msg_t, *cgoAllocMap) {
	if x.refbb64b6ce != nil {
		return *x.refbb64b6ce, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *topo_info_response_msg_t) Deref() {
	if x.refbb64b6ce == nil {
		return
	}
	x.record_count = (uint32_t)(x.refbb64b6ce.record_count)
	packSTopo_info_t(x.topo_array, x.refbb64b6ce.topo_array)
}

// allocJob_alloc_info_msg_tMemory allocates memory for type C.job_alloc_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_alloc_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_alloc_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_alloc_info_msg_tValue = unsafe.Sizeof([1]C.job_alloc_info_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_alloc_info_msg_t) Ref() *C.job_alloc_info_msg_t {
	if x == nil {
		return nil
	}
	return x.ref8a9412bc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_alloc_info_msg_t) Free() {
	if x != nil && x.allocs8a9412bc != nil {
		x.allocs8a9412bc.(*cgoAllocMap).Free()
		x.ref8a9412bc = nil
	}
}

// Newjob_alloc_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_alloc_info_msg_tRef(ref unsafe.Pointer) *job_alloc_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(job_alloc_info_msg_t)
	obj.ref8a9412bc = (*C.job_alloc_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_alloc_info_msg_t) PassRef() (*C.job_alloc_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref8a9412bc != nil {
		return x.ref8a9412bc, nil
	}
	mem8a9412bc := allocJob_alloc_info_msg_tMemory(1)
	ref8a9412bc := (*C.job_alloc_info_msg_t)(mem8a9412bc)
	allocs8a9412bc := new(cgoAllocMap)
	allocs8a9412bc.Add(mem8a9412bc)

	var cjob_id_allocs *cgoAllocMap
	ref8a9412bc.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs8a9412bc.Borrow(cjob_id_allocs)

	var creq_cluster_allocs *cgoAllocMap
	ref8a9412bc.req_cluster, creq_cluster_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.req_cluster)).Data)), cgoAllocsUnknown
	allocs8a9412bc.Borrow(creq_cluster_allocs)

	x.ref8a9412bc = ref8a9412bc
	x.allocs8a9412bc = allocs8a9412bc
	return ref8a9412bc, allocs8a9412bc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_alloc_info_msg_t) PassValue() (C.job_alloc_info_msg_t, *cgoAllocMap) {
	if x.ref8a9412bc != nil {
		return *x.ref8a9412bc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_alloc_info_msg_t) Deref() {
	if x.ref8a9412bc == nil {
		return
	}
	x.job_id = (uint32_t)(x.ref8a9412bc.job_id)
	hxf25d7e1 := (*sliceHeader)(unsafe.Pointer(&x.req_cluster))
	hxf25d7e1.Data = uintptr(unsafe.Pointer(x.ref8a9412bc.req_cluster))
	hxf25d7e1.Cap = 0x7fffffff
	// hxf25d7e1.Len = ?

}

// allocLayout_info_msg_tMemory allocates memory for type C.layout_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocLayout_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfLayout_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfLayout_info_msg_tValue = unsafe.Sizeof([1]C.layout_info_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *layout_info_msg_t) Ref() *C.layout_info_msg_t {
	if x == nil {
		return nil
	}
	return x.ref6987b71d
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *layout_info_msg_t) Free() {
	if x != nil && x.allocs6987b71d != nil {
		x.allocs6987b71d.(*cgoAllocMap).Free()
		x.ref6987b71d = nil
	}
}

// Newlayout_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newlayout_info_msg_tRef(ref unsafe.Pointer) *layout_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(layout_info_msg_t)
	obj.ref6987b71d = (*C.layout_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *layout_info_msg_t) PassRef() (*C.layout_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6987b71d != nil {
		return x.ref6987b71d, nil
	}
	mem6987b71d := allocLayout_info_msg_tMemory(1)
	ref6987b71d := (*C.layout_info_msg_t)(mem6987b71d)
	allocs6987b71d := new(cgoAllocMap)
	allocs6987b71d.Add(mem6987b71d)

	var crecord_count_allocs *cgoAllocMap
	ref6987b71d.record_count, crecord_count_allocs = (C.uint32_t)(x.record_count), cgoAllocsUnknown
	allocs6987b71d.Borrow(crecord_count_allocs)

	var crecords_allocs *cgoAllocMap
	ref6987b71d.records, crecords_allocs = unpackSSByte(x.records)
	allocs6987b71d.Borrow(crecords_allocs)

	x.ref6987b71d = ref6987b71d
	x.allocs6987b71d = allocs6987b71d
	return ref6987b71d, allocs6987b71d

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x layout_info_msg_t) PassValue() (C.layout_info_msg_t, *cgoAllocMap) {
	if x.ref6987b71d != nil {
		return *x.ref6987b71d, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *layout_info_msg_t) Deref() {
	if x.ref6987b71d == nil {
		return
	}
	x.record_count = (uint32_t)(x.ref6987b71d.record_count)
	packSSByte(x.records, x.ref6987b71d.records)
}

// allocUpdate_layout_msg_tMemory allocates memory for type C.update_layout_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUpdate_layout_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUpdate_layout_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfUpdate_layout_msg_tValue = unsafe.Sizeof([1]C.update_layout_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *update_layout_msg_t) Ref() *C.update_layout_msg_t {
	if x == nil {
		return nil
	}
	return x.ref63746d9e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *update_layout_msg_t) Free() {
	if x != nil && x.allocs63746d9e != nil {
		x.allocs63746d9e.(*cgoAllocMap).Free()
		x.ref63746d9e = nil
	}
}

// Newupdate_layout_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newupdate_layout_msg_tRef(ref unsafe.Pointer) *update_layout_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(update_layout_msg_t)
	obj.ref63746d9e = (*C.update_layout_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *update_layout_msg_t) PassRef() (*C.update_layout_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref63746d9e != nil {
		return x.ref63746d9e, nil
	}
	mem63746d9e := allocUpdate_layout_msg_tMemory(1)
	ref63746d9e := (*C.update_layout_msg_t)(mem63746d9e)
	allocs63746d9e := new(cgoAllocMap)
	allocs63746d9e.Add(mem63746d9e)

	var clayout_allocs *cgoAllocMap
	ref63746d9e.layout, clayout_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.layout)).Data)), cgoAllocsUnknown
	allocs63746d9e.Borrow(clayout_allocs)

	var carg_allocs *cgoAllocMap
	ref63746d9e.arg, carg_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.arg)).Data)), cgoAllocsUnknown
	allocs63746d9e.Borrow(carg_allocs)

	x.ref63746d9e = ref63746d9e
	x.allocs63746d9e = allocs63746d9e
	return ref63746d9e, allocs63746d9e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x update_layout_msg_t) PassValue() (C.update_layout_msg_t, *cgoAllocMap) {
	if x.ref63746d9e != nil {
		return *x.ref63746d9e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *update_layout_msg_t) Deref() {
	if x.ref63746d9e == nil {
		return
	}
	hxfd12df3 := (*sliceHeader)(unsafe.Pointer(&x.layout))
	hxfd12df3.Data = uintptr(unsafe.Pointer(x.ref63746d9e.layout))
	hxfd12df3.Cap = 0x7fffffff
	// hxfd12df3.Len = ?

	hxfb30835 := (*sliceHeader)(unsafe.Pointer(&x.arg))
	hxfb30835.Data = uintptr(unsafe.Pointer(x.ref63746d9e.arg))
	hxfb30835.Cap = 0x7fffffff
	// hxfb30835.Len = ?

}

// allocStep_alloc_info_msg_tMemory allocates memory for type C.step_alloc_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStep_alloc_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStep_alloc_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStep_alloc_info_msg_tValue = unsafe.Sizeof([1]C.step_alloc_info_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *step_alloc_info_msg_t) Ref() *C.step_alloc_info_msg_t {
	if x == nil {
		return nil
	}
	return x.ref38ca7831
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *step_alloc_info_msg_t) Free() {
	if x != nil && x.allocs38ca7831 != nil {
		x.allocs38ca7831.(*cgoAllocMap).Free()
		x.ref38ca7831 = nil
	}
}

// Newstep_alloc_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newstep_alloc_info_msg_tRef(ref unsafe.Pointer) *step_alloc_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(step_alloc_info_msg_t)
	obj.ref38ca7831 = (*C.step_alloc_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *step_alloc_info_msg_t) PassRef() (*C.step_alloc_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref38ca7831 != nil {
		return x.ref38ca7831, nil
	}
	mem38ca7831 := allocStep_alloc_info_msg_tMemory(1)
	ref38ca7831 := (*C.step_alloc_info_msg_t)(mem38ca7831)
	allocs38ca7831 := new(cgoAllocMap)
	allocs38ca7831.Add(mem38ca7831)

	var cjob_id_allocs *cgoAllocMap
	ref38ca7831.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs38ca7831.Borrow(cjob_id_allocs)

	var cpack_job_offset_allocs *cgoAllocMap
	ref38ca7831.pack_job_offset, cpack_job_offset_allocs = (C.uint32_t)(x.pack_job_offset), cgoAllocsUnknown
	allocs38ca7831.Borrow(cpack_job_offset_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref38ca7831.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocs38ca7831.Borrow(cstep_id_allocs)

	x.ref38ca7831 = ref38ca7831
	x.allocs38ca7831 = allocs38ca7831
	return ref38ca7831, allocs38ca7831

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x step_alloc_info_msg_t) PassValue() (C.step_alloc_info_msg_t, *cgoAllocMap) {
	if x.ref38ca7831 != nil {
		return *x.ref38ca7831, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *step_alloc_info_msg_t) Deref() {
	if x.ref38ca7831 == nil {
		return
	}
	x.job_id = (uint32_t)(x.ref38ca7831.job_id)
	x.pack_job_offset = (uint32_t)(x.ref38ca7831.pack_job_offset)
	x.step_id = (uint32_t)(x.ref38ca7831.step_id)
}

// allocPowercap_info_msg_tMemory allocates memory for type C.powercap_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPowercap_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPowercap_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPowercap_info_msg_tValue = unsafe.Sizeof([1]C.powercap_info_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *powercap_info_msg_t) Ref() *C.powercap_info_msg_t {
	if x == nil {
		return nil
	}
	return x.refe64943c1
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *powercap_info_msg_t) Free() {
	if x != nil && x.allocse64943c1 != nil {
		x.allocse64943c1.(*cgoAllocMap).Free()
		x.refe64943c1 = nil
	}
}

// Newpowercap_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newpowercap_info_msg_tRef(ref unsafe.Pointer) *powercap_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(powercap_info_msg_t)
	obj.refe64943c1 = (*C.powercap_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *powercap_info_msg_t) PassRef() (*C.powercap_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refe64943c1 != nil {
		return x.refe64943c1, nil
	}
	meme64943c1 := allocPowercap_info_msg_tMemory(1)
	refe64943c1 := (*C.powercap_info_msg_t)(meme64943c1)
	allocse64943c1 := new(cgoAllocMap)
	allocse64943c1.Add(meme64943c1)

	var cpower_cap_allocs *cgoAllocMap
	refe64943c1.power_cap, cpower_cap_allocs = (C.uint32_t)(x.power_cap), cgoAllocsUnknown
	allocse64943c1.Borrow(cpower_cap_allocs)

	var cpower_floor_allocs *cgoAllocMap
	refe64943c1.power_floor, cpower_floor_allocs = (C.uint32_t)(x.power_floor), cgoAllocsUnknown
	allocse64943c1.Borrow(cpower_floor_allocs)

	var cpower_change_allocs *cgoAllocMap
	refe64943c1.power_change, cpower_change_allocs = (C.uint32_t)(x.power_change), cgoAllocsUnknown
	allocse64943c1.Borrow(cpower_change_allocs)

	var cmin_watts_allocs *cgoAllocMap
	refe64943c1.min_watts, cmin_watts_allocs = (C.uint32_t)(x.min_watts), cgoAllocsUnknown
	allocse64943c1.Borrow(cmin_watts_allocs)

	var ccur_max_watts_allocs *cgoAllocMap
	refe64943c1.cur_max_watts, ccur_max_watts_allocs = (C.uint32_t)(x.cur_max_watts), cgoAllocsUnknown
	allocse64943c1.Borrow(ccur_max_watts_allocs)

	var cadj_max_watts_allocs *cgoAllocMap
	refe64943c1.adj_max_watts, cadj_max_watts_allocs = (C.uint32_t)(x.adj_max_watts), cgoAllocsUnknown
	allocse64943c1.Borrow(cadj_max_watts_allocs)

	var cmax_watts_allocs *cgoAllocMap
	refe64943c1.max_watts, cmax_watts_allocs = (C.uint32_t)(x.max_watts), cgoAllocsUnknown
	allocse64943c1.Borrow(cmax_watts_allocs)

	x.refe64943c1 = refe64943c1
	x.allocse64943c1 = allocse64943c1
	return refe64943c1, allocse64943c1

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x powercap_info_msg_t) PassValue() (C.powercap_info_msg_t, *cgoAllocMap) {
	if x.refe64943c1 != nil {
		return *x.refe64943c1, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *powercap_info_msg_t) Deref() {
	if x.refe64943c1 == nil {
		return
	}
	x.power_cap = (uint32_t)(x.refe64943c1.power_cap)
	x.power_floor = (uint32_t)(x.refe64943c1.power_floor)
	x.power_change = (uint32_t)(x.refe64943c1.power_change)
	x.min_watts = (uint32_t)(x.refe64943c1.min_watts)
	x.cur_max_watts = (uint32_t)(x.refe64943c1.cur_max_watts)
	x.adj_max_watts = (uint32_t)(x.refe64943c1.adj_max_watts)
	x.max_watts = (uint32_t)(x.refe64943c1.max_watts)
}

// allocUpdate_powercap_msg_tMemory allocates memory for type C.update_powercap_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUpdate_powercap_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUpdate_powercap_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfUpdate_powercap_msg_tValue = unsafe.Sizeof([1]C.update_powercap_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *update_powercap_msg_t) Ref() *C.update_powercap_msg_t {
	if x == nil {
		return nil
	}
	return x.ref38054dec
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *update_powercap_msg_t) Free() {
	if x != nil && x.allocs38054dec != nil {
		x.allocs38054dec.(*cgoAllocMap).Free()
		x.ref38054dec = nil
	}
}

// Newupdate_powercap_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newupdate_powercap_msg_tRef(ref unsafe.Pointer) *update_powercap_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(update_powercap_msg_t)
	obj.ref38054dec = (*C.update_powercap_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *update_powercap_msg_t) PassRef() (*C.update_powercap_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref38054dec != nil {
		return x.ref38054dec, nil
	}
	mem38054dec := allocUpdate_powercap_msg_tMemory(1)
	ref38054dec := (*C.update_powercap_msg_t)(mem38054dec)
	allocs38054dec := new(cgoAllocMap)
	allocs38054dec.Add(mem38054dec)

	var cpower_cap_allocs *cgoAllocMap
	ref38054dec.power_cap, cpower_cap_allocs = (C.uint32_t)(x.power_cap), cgoAllocsUnknown
	allocs38054dec.Borrow(cpower_cap_allocs)

	var cpower_floor_allocs *cgoAllocMap
	ref38054dec.power_floor, cpower_floor_allocs = (C.uint32_t)(x.power_floor), cgoAllocsUnknown
	allocs38054dec.Borrow(cpower_floor_allocs)

	var cpower_change_allocs *cgoAllocMap
	ref38054dec.power_change, cpower_change_allocs = (C.uint32_t)(x.power_change), cgoAllocsUnknown
	allocs38054dec.Borrow(cpower_change_allocs)

	var cmin_watts_allocs *cgoAllocMap
	ref38054dec.min_watts, cmin_watts_allocs = (C.uint32_t)(x.min_watts), cgoAllocsUnknown
	allocs38054dec.Borrow(cmin_watts_allocs)

	var ccur_max_watts_allocs *cgoAllocMap
	ref38054dec.cur_max_watts, ccur_max_watts_allocs = (C.uint32_t)(x.cur_max_watts), cgoAllocsUnknown
	allocs38054dec.Borrow(ccur_max_watts_allocs)

	var cadj_max_watts_allocs *cgoAllocMap
	ref38054dec.adj_max_watts, cadj_max_watts_allocs = (C.uint32_t)(x.adj_max_watts), cgoAllocsUnknown
	allocs38054dec.Borrow(cadj_max_watts_allocs)

	var cmax_watts_allocs *cgoAllocMap
	ref38054dec.max_watts, cmax_watts_allocs = (C.uint32_t)(x.max_watts), cgoAllocsUnknown
	allocs38054dec.Borrow(cmax_watts_allocs)

	x.ref38054dec = ref38054dec
	x.allocs38054dec = allocs38054dec
	return ref38054dec, allocs38054dec

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x update_powercap_msg_t) PassValue() (C.update_powercap_msg_t, *cgoAllocMap) {
	if x.ref38054dec != nil {
		return *x.ref38054dec, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *update_powercap_msg_t) Deref() {
	if x.ref38054dec == nil {
		return
	}
	x.power_cap = (uint32_t)(x.ref38054dec.power_cap)
	x.power_floor = (uint32_t)(x.ref38054dec.power_floor)
	x.power_change = (uint32_t)(x.ref38054dec.power_change)
	x.min_watts = (uint32_t)(x.ref38054dec.min_watts)
	x.cur_max_watts = (uint32_t)(x.ref38054dec.cur_max_watts)
	x.adj_max_watts = (uint32_t)(x.ref38054dec.adj_max_watts)
	x.max_watts = (uint32_t)(x.ref38054dec.max_watts)
}

// allocAcct_gather_node_resp_msg_tMemory allocates memory for type C.acct_gather_node_resp_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAcct_gather_node_resp_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAcct_gather_node_resp_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAcct_gather_node_resp_msg_tValue = unsafe.Sizeof([1]C.acct_gather_node_resp_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *acct_gather_node_resp_msg_t) Ref() *C.acct_gather_node_resp_msg_t {
	if x == nil {
		return nil
	}
	return x.ref94035ba4
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *acct_gather_node_resp_msg_t) Free() {
	if x != nil && x.allocs94035ba4 != nil {
		x.allocs94035ba4.(*cgoAllocMap).Free()
		x.ref94035ba4 = nil
	}
}

// Newacct_gather_node_resp_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newacct_gather_node_resp_msg_tRef(ref unsafe.Pointer) *acct_gather_node_resp_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(acct_gather_node_resp_msg_t)
	obj.ref94035ba4 = (*C.acct_gather_node_resp_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *acct_gather_node_resp_msg_t) PassRef() (*C.acct_gather_node_resp_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref94035ba4 != nil {
		return x.ref94035ba4, nil
	}
	mem94035ba4 := allocAcct_gather_node_resp_msg_tMemory(1)
	ref94035ba4 := (*C.acct_gather_node_resp_msg_t)(mem94035ba4)
	allocs94035ba4 := new(cgoAllocMap)
	allocs94035ba4.Add(mem94035ba4)

	var cenergy_allocs *cgoAllocMap
	ref94035ba4.energy, cenergy_allocs = unpackSAcct_gather_energy_t(x.energy)
	allocs94035ba4.Borrow(cenergy_allocs)

	var cnode_name_allocs *cgoAllocMap
	ref94035ba4.node_name, cnode_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_name)).Data)), cgoAllocsUnknown
	allocs94035ba4.Borrow(cnode_name_allocs)

	var csensor_cnt_allocs *cgoAllocMap
	ref94035ba4.sensor_cnt, csensor_cnt_allocs = (C.uint16_t)(x.sensor_cnt), cgoAllocsUnknown
	allocs94035ba4.Borrow(csensor_cnt_allocs)

	x.ref94035ba4 = ref94035ba4
	x.allocs94035ba4 = allocs94035ba4
	return ref94035ba4, allocs94035ba4

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x acct_gather_node_resp_msg_t) PassValue() (C.acct_gather_node_resp_msg_t, *cgoAllocMap) {
	if x.ref94035ba4 != nil {
		return *x.ref94035ba4, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *acct_gather_node_resp_msg_t) Deref() {
	if x.ref94035ba4 == nil {
		return
	}
	packSAcct_gather_energy_t(x.energy, x.ref94035ba4.energy)
	hxf0dfcee := (*sliceHeader)(unsafe.Pointer(&x.node_name))
	hxf0dfcee.Data = uintptr(unsafe.Pointer(x.ref94035ba4.node_name))
	hxf0dfcee.Cap = 0x7fffffff
	// hxf0dfcee.Len = ?

	x.sensor_cnt = (uint16_t)(x.ref94035ba4.sensor_cnt)
}

// allocAcct_gather_energy_req_msg_tMemory allocates memory for type C.acct_gather_energy_req_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAcct_gather_energy_req_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAcct_gather_energy_req_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAcct_gather_energy_req_msg_tValue = unsafe.Sizeof([1]C.acct_gather_energy_req_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *acct_gather_energy_req_msg_t) Ref() *C.acct_gather_energy_req_msg_t {
	if x == nil {
		return nil
	}
	return x.ref47835ed
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *acct_gather_energy_req_msg_t) Free() {
	if x != nil && x.allocs47835ed != nil {
		x.allocs47835ed.(*cgoAllocMap).Free()
		x.ref47835ed = nil
	}
}

// Newacct_gather_energy_req_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newacct_gather_energy_req_msg_tRef(ref unsafe.Pointer) *acct_gather_energy_req_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(acct_gather_energy_req_msg_t)
	obj.ref47835ed = (*C.acct_gather_energy_req_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *acct_gather_energy_req_msg_t) PassRef() (*C.acct_gather_energy_req_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref47835ed != nil {
		return x.ref47835ed, nil
	}
	mem47835ed := allocAcct_gather_energy_req_msg_tMemory(1)
	ref47835ed := (*C.acct_gather_energy_req_msg_t)(mem47835ed)
	allocs47835ed := new(cgoAllocMap)
	allocs47835ed.Add(mem47835ed)

	var cdelta_allocs *cgoAllocMap
	ref47835ed.delta, cdelta_allocs = (C.uint16_t)(x.delta), cgoAllocsUnknown
	allocs47835ed.Borrow(cdelta_allocs)

	x.ref47835ed = ref47835ed
	x.allocs47835ed = allocs47835ed
	return ref47835ed, allocs47835ed

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x acct_gather_energy_req_msg_t) PassValue() (C.acct_gather_energy_req_msg_t, *cgoAllocMap) {
	if x.ref47835ed != nil {
		return *x.ref47835ed, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *acct_gather_energy_req_msg_t) Deref() {
	if x.ref47835ed == nil {
		return
	}
	x.delta = (uint16_t)(x.ref47835ed.delta)
}

// allocPartition_info_tMemory allocates memory for type C.partition_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPartition_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPartition_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPartition_info_tValue = unsafe.Sizeof([1]C.partition_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *partition_info_t) Ref() *C.partition_info_t {
	if x == nil {
		return nil
	}
	return x.refc2f47d8f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *partition_info_t) Free() {
	if x != nil && x.allocsc2f47d8f != nil {
		x.allocsc2f47d8f.(*cgoAllocMap).Free()
		x.refc2f47d8f = nil
	}
}

// Newpartition_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newpartition_info_tRef(ref unsafe.Pointer) *partition_info_t {
	if ref == nil {
		return nil
	}
	obj := new(partition_info_t)
	obj.refc2f47d8f = (*C.partition_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *partition_info_t) PassRef() (*C.partition_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc2f47d8f != nil {
		return x.refc2f47d8f, nil
	}
	memc2f47d8f := allocPartition_info_tMemory(1)
	refc2f47d8f := (*C.partition_info_t)(memc2f47d8f)
	allocsc2f47d8f := new(cgoAllocMap)
	allocsc2f47d8f.Add(memc2f47d8f)

	var callow_alloc_nodes_allocs *cgoAllocMap
	refc2f47d8f.allow_alloc_nodes, callow_alloc_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_alloc_nodes)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(callow_alloc_nodes_allocs)

	var callow_accounts_allocs *cgoAllocMap
	refc2f47d8f.allow_accounts, callow_accounts_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_accounts)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(callow_accounts_allocs)

	var callow_groups_allocs *cgoAllocMap
	refc2f47d8f.allow_groups, callow_groups_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_groups)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(callow_groups_allocs)

	var callow_qos_allocs *cgoAllocMap
	refc2f47d8f.allow_qos, callow_qos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_qos)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(callow_qos_allocs)

	var calternate_allocs *cgoAllocMap
	refc2f47d8f.alternate, calternate_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.alternate)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(calternate_allocs)

	var cbilling_weights_str_allocs *cgoAllocMap
	refc2f47d8f.billing_weights_str, cbilling_weights_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.billing_weights_str)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cbilling_weights_str_allocs)

	var ccluster_name_allocs *cgoAllocMap
	refc2f47d8f.cluster_name, ccluster_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster_name)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(ccluster_name_allocs)

	var ccr_type_allocs *cgoAllocMap
	refc2f47d8f.cr_type, ccr_type_allocs = (C.uint16_t)(x.cr_type), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(ccr_type_allocs)

	var cdef_mem_per_cpu_allocs *cgoAllocMap
	refc2f47d8f.def_mem_per_cpu, cdef_mem_per_cpu_allocs = (C.uint64_t)(x.def_mem_per_cpu), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cdef_mem_per_cpu_allocs)

	var cdefault_time_allocs *cgoAllocMap
	refc2f47d8f.default_time, cdefault_time_allocs = (C.uint32_t)(x.default_time), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cdefault_time_allocs)

	var cdeny_accounts_allocs *cgoAllocMap
	refc2f47d8f.deny_accounts, cdeny_accounts_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.deny_accounts)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cdeny_accounts_allocs)

	var cdeny_qos_allocs *cgoAllocMap
	refc2f47d8f.deny_qos, cdeny_qos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.deny_qos)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cdeny_qos_allocs)

	var cflags_allocs *cgoAllocMap
	refc2f47d8f.flags, cflags_allocs = (C.uint16_t)(x.flags), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cflags_allocs)

	var cgrace_time_allocs *cgoAllocMap
	refc2f47d8f.grace_time, cgrace_time_allocs = (C.uint32_t)(x.grace_time), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cgrace_time_allocs)

	var cmax_cpus_per_node_allocs *cgoAllocMap
	refc2f47d8f.max_cpus_per_node, cmax_cpus_per_node_allocs = (C.uint32_t)(x.max_cpus_per_node), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cmax_cpus_per_node_allocs)

	var cmax_mem_per_cpu_allocs *cgoAllocMap
	refc2f47d8f.max_mem_per_cpu, cmax_mem_per_cpu_allocs = (C.uint64_t)(x.max_mem_per_cpu), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cmax_mem_per_cpu_allocs)

	var cmax_nodes_allocs *cgoAllocMap
	refc2f47d8f.max_nodes, cmax_nodes_allocs = (C.uint32_t)(x.max_nodes), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cmax_nodes_allocs)

	var cmax_share_allocs *cgoAllocMap
	refc2f47d8f.max_share, cmax_share_allocs = (C.uint16_t)(x.max_share), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cmax_share_allocs)

	var cmax_time_allocs *cgoAllocMap
	refc2f47d8f.max_time, cmax_time_allocs = (C.uint32_t)(x.max_time), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cmax_time_allocs)

	var cmin_nodes_allocs *cgoAllocMap
	refc2f47d8f.min_nodes, cmin_nodes_allocs = (C.uint32_t)(x.min_nodes), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cmin_nodes_allocs)

	var cname_allocs *cgoAllocMap
	refc2f47d8f.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cname_allocs)

	var cnode_inx_allocs *cgoAllocMap
	refc2f47d8f.node_inx, cnode_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_inx)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cnode_inx_allocs)

	var cnodes_allocs *cgoAllocMap
	refc2f47d8f.nodes, cnodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.nodes)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cnodes_allocs)

	var cover_time_limit_allocs *cgoAllocMap
	refc2f47d8f.over_time_limit, cover_time_limit_allocs = (C.uint16_t)(x.over_time_limit), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cover_time_limit_allocs)

	var cpreempt_mode_allocs *cgoAllocMap
	refc2f47d8f.preempt_mode, cpreempt_mode_allocs = (C.uint16_t)(x.preempt_mode), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cpreempt_mode_allocs)

	var cpriority_job_factor_allocs *cgoAllocMap
	refc2f47d8f.priority_job_factor, cpriority_job_factor_allocs = (C.uint16_t)(x.priority_job_factor), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cpriority_job_factor_allocs)

	var cpriority_tier_allocs *cgoAllocMap
	refc2f47d8f.priority_tier, cpriority_tier_allocs = (C.uint16_t)(x.priority_tier), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cpriority_tier_allocs)

	var cqos_char_allocs *cgoAllocMap
	refc2f47d8f.qos_char, cqos_char_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.qos_char)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cqos_char_allocs)

	var cstate_up_allocs *cgoAllocMap
	refc2f47d8f.state_up, cstate_up_allocs = (C.uint16_t)(x.state_up), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(cstate_up_allocs)

	var ctotal_cpus_allocs *cgoAllocMap
	refc2f47d8f.total_cpus, ctotal_cpus_allocs = (C.uint32_t)(x.total_cpus), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(ctotal_cpus_allocs)

	var ctotal_nodes_allocs *cgoAllocMap
	refc2f47d8f.total_nodes, ctotal_nodes_allocs = (C.uint32_t)(x.total_nodes), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(ctotal_nodes_allocs)

	var ctres_fmt_str_allocs *cgoAllocMap
	refc2f47d8f.tres_fmt_str, ctres_fmt_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_fmt_str)).Data)), cgoAllocsUnknown
	allocsc2f47d8f.Borrow(ctres_fmt_str_allocs)

	x.refc2f47d8f = refc2f47d8f
	x.allocsc2f47d8f = allocsc2f47d8f
	return refc2f47d8f, allocsc2f47d8f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x partition_info_t) PassValue() (C.partition_info_t, *cgoAllocMap) {
	if x.refc2f47d8f != nil {
		return *x.refc2f47d8f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *partition_info_t) Deref() {
	if x.refc2f47d8f == nil {
		return
	}
	hxf60ed55 := (*sliceHeader)(unsafe.Pointer(&x.allow_alloc_nodes))
	hxf60ed55.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.allow_alloc_nodes))
	hxf60ed55.Cap = 0x7fffffff
	// hxf60ed55.Len = ?

	hxf7ffdbf := (*sliceHeader)(unsafe.Pointer(&x.allow_accounts))
	hxf7ffdbf.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.allow_accounts))
	hxf7ffdbf.Cap = 0x7fffffff
	// hxf7ffdbf.Len = ?

	hxfb7ec97 := (*sliceHeader)(unsafe.Pointer(&x.allow_groups))
	hxfb7ec97.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.allow_groups))
	hxfb7ec97.Cap = 0x7fffffff
	// hxfb7ec97.Len = ?

	hxf79e6eb := (*sliceHeader)(unsafe.Pointer(&x.allow_qos))
	hxf79e6eb.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.allow_qos))
	hxf79e6eb.Cap = 0x7fffffff
	// hxf79e6eb.Len = ?

	hxf75e1cc := (*sliceHeader)(unsafe.Pointer(&x.alternate))
	hxf75e1cc.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.alternate))
	hxf75e1cc.Cap = 0x7fffffff
	// hxf75e1cc.Len = ?

	hxff6b567 := (*sliceHeader)(unsafe.Pointer(&x.billing_weights_str))
	hxff6b567.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.billing_weights_str))
	hxff6b567.Cap = 0x7fffffff
	// hxff6b567.Len = ?

	hxfc280cb := (*sliceHeader)(unsafe.Pointer(&x.cluster_name))
	hxfc280cb.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.cluster_name))
	hxfc280cb.Cap = 0x7fffffff
	// hxfc280cb.Len = ?

	x.cr_type = (uint16_t)(x.refc2f47d8f.cr_type)
	x.def_mem_per_cpu = (uint64_t)(x.refc2f47d8f.def_mem_per_cpu)
	x.default_time = (uint32_t)(x.refc2f47d8f.default_time)
	hxf976962 := (*sliceHeader)(unsafe.Pointer(&x.deny_accounts))
	hxf976962.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.deny_accounts))
	hxf976962.Cap = 0x7fffffff
	// hxf976962.Len = ?

	hxf27c4d8 := (*sliceHeader)(unsafe.Pointer(&x.deny_qos))
	hxf27c4d8.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.deny_qos))
	hxf27c4d8.Cap = 0x7fffffff
	// hxf27c4d8.Len = ?

	x.flags = (uint16_t)(x.refc2f47d8f.flags)
	x.grace_time = (uint32_t)(x.refc2f47d8f.grace_time)
	x.max_cpus_per_node = (uint32_t)(x.refc2f47d8f.max_cpus_per_node)
	x.max_mem_per_cpu = (uint64_t)(x.refc2f47d8f.max_mem_per_cpu)
	x.max_nodes = (uint32_t)(x.refc2f47d8f.max_nodes)
	x.max_share = (uint16_t)(x.refc2f47d8f.max_share)
	x.max_time = (uint32_t)(x.refc2f47d8f.max_time)
	x.min_nodes = (uint32_t)(x.refc2f47d8f.min_nodes)
	hxfc0f2e9 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxfc0f2e9.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.name))
	hxfc0f2e9.Cap = 0x7fffffff
	// hxfc0f2e9.Len = ?

	hxfcf5b8d := (*sliceHeader)(unsafe.Pointer(&x.node_inx))
	hxfcf5b8d.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.node_inx))
	hxfcf5b8d.Cap = 0x7fffffff
	// hxfcf5b8d.Len = ?

	hxf571c37 := (*sliceHeader)(unsafe.Pointer(&x.nodes))
	hxf571c37.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.nodes))
	hxf571c37.Cap = 0x7fffffff
	// hxf571c37.Len = ?

	x.over_time_limit = (uint16_t)(x.refc2f47d8f.over_time_limit)
	x.preempt_mode = (uint16_t)(x.refc2f47d8f.preempt_mode)
	x.priority_job_factor = (uint16_t)(x.refc2f47d8f.priority_job_factor)
	x.priority_tier = (uint16_t)(x.refc2f47d8f.priority_tier)
	hxf4f53c1 := (*sliceHeader)(unsafe.Pointer(&x.qos_char))
	hxf4f53c1.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.qos_char))
	hxf4f53c1.Cap = 0x7fffffff
	// hxf4f53c1.Len = ?

	x.state_up = (uint16_t)(x.refc2f47d8f.state_up)
	x.total_cpus = (uint32_t)(x.refc2f47d8f.total_cpus)
	x.total_nodes = (uint32_t)(x.refc2f47d8f.total_nodes)
	hxf45018f := (*sliceHeader)(unsafe.Pointer(&x.tres_fmt_str))
	hxf45018f.Data = uintptr(unsafe.Pointer(x.refc2f47d8f.tres_fmt_str))
	hxf45018f.Cap = 0x7fffffff
	// hxf45018f.Len = ?

}

// allocDelete_part_msg_tMemory allocates memory for type C.delete_part_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocDelete_part_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfDelete_part_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfDelete_part_msg_tValue = unsafe.Sizeof([1]C.delete_part_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *delete_part_msg_t) Ref() *C.delete_part_msg_t {
	if x == nil {
		return nil
	}
	return x.ref2d75a55b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *delete_part_msg_t) Free() {
	if x != nil && x.allocs2d75a55b != nil {
		x.allocs2d75a55b.(*cgoAllocMap).Free()
		x.ref2d75a55b = nil
	}
}

// Newdelete_part_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newdelete_part_msg_tRef(ref unsafe.Pointer) *delete_part_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(delete_part_msg_t)
	obj.ref2d75a55b = (*C.delete_part_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *delete_part_msg_t) PassRef() (*C.delete_part_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref2d75a55b != nil {
		return x.ref2d75a55b, nil
	}
	mem2d75a55b := allocDelete_part_msg_tMemory(1)
	ref2d75a55b := (*C.delete_part_msg_t)(mem2d75a55b)
	allocs2d75a55b := new(cgoAllocMap)
	allocs2d75a55b.Add(mem2d75a55b)

	var cname_allocs *cgoAllocMap
	ref2d75a55b.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs2d75a55b.Borrow(cname_allocs)

	x.ref2d75a55b = ref2d75a55b
	x.allocs2d75a55b = allocs2d75a55b
	return ref2d75a55b, allocs2d75a55b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x delete_part_msg_t) PassValue() (C.delete_part_msg_t, *cgoAllocMap) {
	if x.ref2d75a55b != nil {
		return *x.ref2d75a55b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *delete_part_msg_t) Deref() {
	if x.ref2d75a55b == nil {
		return
	}
	hxfe6fc23 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxfe6fc23.Data = uintptr(unsafe.Pointer(x.ref2d75a55b.name))
	hxfe6fc23.Cap = 0x7fffffff
	// hxfe6fc23.Len = ?

}

// allocResource_allocation_response_msg_tMemory allocates memory for type C.resource_allocation_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocResource_allocation_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfResource_allocation_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfResource_allocation_response_msg_tValue = unsafe.Sizeof([1]C.resource_allocation_response_msg_t{})

// unpackSSlurm_addr_t transforms a sliced Go data structure into plain C format.
func unpackSSlurm_addr_t(x []slurm_addr_t) (unpacked *C.slurm_addr_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurm_addr_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurm_addr_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurm_addr_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurm_addr_t)(unsafe.Pointer(h.Data))
	return
}

// packSSlurm_addr_t reads sliced Go data structure out from plain C format.
func packSSlurm_addr_t(v []slurm_addr_t, ptr0 *C.slurm_addr_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSlurm_addr_tValue]C.slurm_addr_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newslurm_addr_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *resource_allocation_response_msg_t) Ref() *C.resource_allocation_response_msg_t {
	if x == nil {
		return nil
	}
	return x.refa10376ab
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *resource_allocation_response_msg_t) Free() {
	if x != nil && x.allocsa10376ab != nil {
		x.allocsa10376ab.(*cgoAllocMap).Free()
		x.refa10376ab = nil
	}
}

// Newresource_allocation_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newresource_allocation_response_msg_tRef(ref unsafe.Pointer) *resource_allocation_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(resource_allocation_response_msg_t)
	obj.refa10376ab = (*C.resource_allocation_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *resource_allocation_response_msg_t) PassRef() (*C.resource_allocation_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refa10376ab != nil {
		return x.refa10376ab, nil
	}
	mema10376ab := allocResource_allocation_response_msg_tMemory(1)
	refa10376ab := (*C.resource_allocation_response_msg_t)(mema10376ab)
	allocsa10376ab := new(cgoAllocMap)
	allocsa10376ab.Add(mema10376ab)

	var caccount_allocs *cgoAllocMap
	refa10376ab.account, caccount_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.account)).Data)), cgoAllocsUnknown
	allocsa10376ab.Borrow(caccount_allocs)

	var cjob_id_allocs *cgoAllocMap
	refa10376ab.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocsa10376ab.Borrow(cjob_id_allocs)

	var calias_list_allocs *cgoAllocMap
	refa10376ab.alias_list, calias_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.alias_list)).Data)), cgoAllocsUnknown
	allocsa10376ab.Borrow(calias_list_allocs)

	var ccpu_freq_min_allocs *cgoAllocMap
	refa10376ab.cpu_freq_min, ccpu_freq_min_allocs = (C.uint32_t)(x.cpu_freq_min), cgoAllocsUnknown
	allocsa10376ab.Borrow(ccpu_freq_min_allocs)

	var ccpu_freq_max_allocs *cgoAllocMap
	refa10376ab.cpu_freq_max, ccpu_freq_max_allocs = (C.uint32_t)(x.cpu_freq_max), cgoAllocsUnknown
	allocsa10376ab.Borrow(ccpu_freq_max_allocs)

	var ccpu_freq_gov_allocs *cgoAllocMap
	refa10376ab.cpu_freq_gov, ccpu_freq_gov_allocs = (C.uint32_t)(x.cpu_freq_gov), cgoAllocsUnknown
	allocsa10376ab.Borrow(ccpu_freq_gov_allocs)

	var ccpus_per_node_allocs *cgoAllocMap
	refa10376ab.cpus_per_node, ccpus_per_node_allocs = (*C.uint16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cpus_per_node)).Data)), cgoAllocsUnknown
	allocsa10376ab.Borrow(ccpus_per_node_allocs)

	var ccpu_count_reps_allocs *cgoAllocMap
	refa10376ab.cpu_count_reps, ccpu_count_reps_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cpu_count_reps)).Data)), cgoAllocsUnknown
	allocsa10376ab.Borrow(ccpu_count_reps_allocs)

	var cenv_size_allocs *cgoAllocMap
	refa10376ab.env_size, cenv_size_allocs = (C.uint32_t)(x.env_size), cgoAllocsUnknown
	allocsa10376ab.Borrow(cenv_size_allocs)

	var cenvironment_allocs *cgoAllocMap
	refa10376ab.environment, cenvironment_allocs = unpackSSByte(x.environment)
	allocsa10376ab.Borrow(cenvironment_allocs)

	var cerror_code_allocs *cgoAllocMap
	refa10376ab.error_code, cerror_code_allocs = (C.uint32_t)(x.error_code), cgoAllocsUnknown
	allocsa10376ab.Borrow(cerror_code_allocs)

	var cjob_submit_user_msg_allocs *cgoAllocMap
	refa10376ab.job_submit_user_msg, cjob_submit_user_msg_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_submit_user_msg)).Data)), cgoAllocsUnknown
	allocsa10376ab.Borrow(cjob_submit_user_msg_allocs)

	var cnode_addr_allocs *cgoAllocMap
	refa10376ab.node_addr, cnode_addr_allocs = unpackSSlurm_addr_t(x.node_addr)
	allocsa10376ab.Borrow(cnode_addr_allocs)

	var cnode_cnt_allocs *cgoAllocMap
	refa10376ab.node_cnt, cnode_cnt_allocs = (C.uint32_t)(x.node_cnt), cgoAllocsUnknown
	allocsa10376ab.Borrow(cnode_cnt_allocs)

	var cnode_list_allocs *cgoAllocMap
	refa10376ab.node_list, cnode_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_list)).Data)), cgoAllocsUnknown
	allocsa10376ab.Borrow(cnode_list_allocs)

	var cntasks_per_board_allocs *cgoAllocMap
	refa10376ab.ntasks_per_board, cntasks_per_board_allocs = (C.uint16_t)(x.ntasks_per_board), cgoAllocsUnknown
	allocsa10376ab.Borrow(cntasks_per_board_allocs)

	var cntasks_per_core_allocs *cgoAllocMap
	refa10376ab.ntasks_per_core, cntasks_per_core_allocs = (C.uint16_t)(x.ntasks_per_core), cgoAllocsUnknown
	allocsa10376ab.Borrow(cntasks_per_core_allocs)

	var cntasks_per_socket_allocs *cgoAllocMap
	refa10376ab.ntasks_per_socket, cntasks_per_socket_allocs = (C.uint16_t)(x.ntasks_per_socket), cgoAllocsUnknown
	allocsa10376ab.Borrow(cntasks_per_socket_allocs)

	var cnum_cpu_groups_allocs *cgoAllocMap
	refa10376ab.num_cpu_groups, cnum_cpu_groups_allocs = (C.uint32_t)(x.num_cpu_groups), cgoAllocsUnknown
	allocsa10376ab.Borrow(cnum_cpu_groups_allocs)

	var cpartition_allocs *cgoAllocMap
	refa10376ab.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partition)).Data)), cgoAllocsUnknown
	allocsa10376ab.Borrow(cpartition_allocs)

	var cpn_min_memory_allocs *cgoAllocMap
	refa10376ab.pn_min_memory, cpn_min_memory_allocs = (C.uint64_t)(x.pn_min_memory), cgoAllocsUnknown
	allocsa10376ab.Borrow(cpn_min_memory_allocs)

	var cqos_allocs *cgoAllocMap
	refa10376ab.qos, cqos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.qos)).Data)), cgoAllocsUnknown
	allocsa10376ab.Borrow(cqos_allocs)

	var cresv_name_allocs *cgoAllocMap
	refa10376ab.resv_name, cresv_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.resv_name)).Data)), cgoAllocsUnknown
	allocsa10376ab.Borrow(cresv_name_allocs)

	var cselect_jobinfo_allocs *cgoAllocMap
	refa10376ab.select_jobinfo, cselect_jobinfo_allocs = unpackSDynamic_plugin_data_t(x.select_jobinfo)
	allocsa10376ab.Borrow(cselect_jobinfo_allocs)

	var cworking_cluster_rec_allocs *cgoAllocMap
	refa10376ab.working_cluster_rec, cworking_cluster_rec_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.working_cluster_rec)), cgoAllocsUnknown
	allocsa10376ab.Borrow(cworking_cluster_rec_allocs)

	x.refa10376ab = refa10376ab
	x.allocsa10376ab = allocsa10376ab
	return refa10376ab, allocsa10376ab

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x resource_allocation_response_msg_t) PassValue() (C.resource_allocation_response_msg_t, *cgoAllocMap) {
	if x.refa10376ab != nil {
		return *x.refa10376ab, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *resource_allocation_response_msg_t) Deref() {
	if x.refa10376ab == nil {
		return
	}
	hxf5fa3ec := (*sliceHeader)(unsafe.Pointer(&x.account))
	hxf5fa3ec.Data = uintptr(unsafe.Pointer(x.refa10376ab.account))
	hxf5fa3ec.Cap = 0x7fffffff
	// hxf5fa3ec.Len = ?

	x.job_id = (uint32_t)(x.refa10376ab.job_id)
	hxf0b70ee := (*sliceHeader)(unsafe.Pointer(&x.alias_list))
	hxf0b70ee.Data = uintptr(unsafe.Pointer(x.refa10376ab.alias_list))
	hxf0b70ee.Cap = 0x7fffffff
	// hxf0b70ee.Len = ?

	x.cpu_freq_min = (uint32_t)(x.refa10376ab.cpu_freq_min)
	x.cpu_freq_max = (uint32_t)(x.refa10376ab.cpu_freq_max)
	x.cpu_freq_gov = (uint32_t)(x.refa10376ab.cpu_freq_gov)
	hxf988202 := (*sliceHeader)(unsafe.Pointer(&x.cpus_per_node))
	hxf988202.Data = uintptr(unsafe.Pointer(x.refa10376ab.cpus_per_node))
	hxf988202.Cap = 0x7fffffff
	// hxf988202.Len = ?

	hxfe0b64a := (*sliceHeader)(unsafe.Pointer(&x.cpu_count_reps))
	hxfe0b64a.Data = uintptr(unsafe.Pointer(x.refa10376ab.cpu_count_reps))
	hxfe0b64a.Cap = 0x7fffffff
	// hxfe0b64a.Len = ?

	x.env_size = (uint32_t)(x.refa10376ab.env_size)
	packSSByte(x.environment, x.refa10376ab.environment)
	x.error_code = (uint32_t)(x.refa10376ab.error_code)
	hxf4cafca := (*sliceHeader)(unsafe.Pointer(&x.job_submit_user_msg))
	hxf4cafca.Data = uintptr(unsafe.Pointer(x.refa10376ab.job_submit_user_msg))
	hxf4cafca.Cap = 0x7fffffff
	// hxf4cafca.Len = ?

	packSSlurm_addr_t(x.node_addr, x.refa10376ab.node_addr)
	x.node_cnt = (uint32_t)(x.refa10376ab.node_cnt)
	hxf0467fd := (*sliceHeader)(unsafe.Pointer(&x.node_list))
	hxf0467fd.Data = uintptr(unsafe.Pointer(x.refa10376ab.node_list))
	hxf0467fd.Cap = 0x7fffffff
	// hxf0467fd.Len = ?

	x.ntasks_per_board = (uint16_t)(x.refa10376ab.ntasks_per_board)
	x.ntasks_per_core = (uint16_t)(x.refa10376ab.ntasks_per_core)
	x.ntasks_per_socket = (uint16_t)(x.refa10376ab.ntasks_per_socket)
	x.num_cpu_groups = (uint32_t)(x.refa10376ab.num_cpu_groups)
	hxf5d2d3f := (*sliceHeader)(unsafe.Pointer(&x.partition))
	hxf5d2d3f.Data = uintptr(unsafe.Pointer(x.refa10376ab.partition))
	hxf5d2d3f.Cap = 0x7fffffff
	// hxf5d2d3f.Len = ?

	x.pn_min_memory = (uint64_t)(x.refa10376ab.pn_min_memory)
	hxf02b567 := (*sliceHeader)(unsafe.Pointer(&x.qos))
	hxf02b567.Data = uintptr(unsafe.Pointer(x.refa10376ab.qos))
	hxf02b567.Cap = 0x7fffffff
	// hxf02b567.Len = ?

	hxfdf9a6d := (*sliceHeader)(unsafe.Pointer(&x.resv_name))
	hxfdf9a6d.Data = uintptr(unsafe.Pointer(x.refa10376ab.resv_name))
	hxfdf9a6d.Cap = 0x7fffffff
	// hxfdf9a6d.Len = ?

	packSDynamic_plugin_data_t(x.select_jobinfo, x.refa10376ab.select_jobinfo)
	x.working_cluster_rec = (unsafe.Pointer)(unsafe.Pointer(x.refa10376ab.working_cluster_rec))
}

// allocPartition_info_msg_tMemory allocates memory for type C.partition_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPartition_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPartition_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPartition_info_msg_tValue = unsafe.Sizeof([1]C.partition_info_msg_t{})

// unpackSPartition_info_t transforms a sliced Go data structure into plain C format.
func unpackSPartition_info_t(x []partition_info_t) (unpacked *C.partition_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.partition_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPartition_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.partition_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.partition_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSPartition_info_t reads sliced Go data structure out from plain C format.
func packSPartition_info_t(v []partition_info_t, ptr0 *C.partition_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPartition_info_tValue]C.partition_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newpartition_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *partition_info_msg_t) Ref() *C.partition_info_msg_t {
	if x == nil {
		return nil
	}
	return x.reffe44d6cd
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *partition_info_msg_t) Free() {
	if x != nil && x.allocsfe44d6cd != nil {
		x.allocsfe44d6cd.(*cgoAllocMap).Free()
		x.reffe44d6cd = nil
	}
}

// Newpartition_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newpartition_info_msg_tRef(ref unsafe.Pointer) *partition_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(partition_info_msg_t)
	obj.reffe44d6cd = (*C.partition_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *partition_info_msg_t) PassRef() (*C.partition_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reffe44d6cd != nil {
		return x.reffe44d6cd, nil
	}
	memfe44d6cd := allocPartition_info_msg_tMemory(1)
	reffe44d6cd := (*C.partition_info_msg_t)(memfe44d6cd)
	allocsfe44d6cd := new(cgoAllocMap)
	allocsfe44d6cd.Add(memfe44d6cd)

	var clast_update_allocs *cgoAllocMap
	reffe44d6cd.last_update, clast_update_allocs = (C.time_t)(x.last_update), cgoAllocsUnknown
	allocsfe44d6cd.Borrow(clast_update_allocs)

	var crecord_count_allocs *cgoAllocMap
	reffe44d6cd.record_count, crecord_count_allocs = (C.uint32_t)(x.record_count), cgoAllocsUnknown
	allocsfe44d6cd.Borrow(crecord_count_allocs)

	var cpartition_array_allocs *cgoAllocMap
	reffe44d6cd.partition_array, cpartition_array_allocs = unpackSPartition_info_t(x.partition_array)
	allocsfe44d6cd.Borrow(cpartition_array_allocs)

	x.reffe44d6cd = reffe44d6cd
	x.allocsfe44d6cd = allocsfe44d6cd
	return reffe44d6cd, allocsfe44d6cd

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x partition_info_msg_t) PassValue() (C.partition_info_msg_t, *cgoAllocMap) {
	if x.reffe44d6cd != nil {
		return *x.reffe44d6cd, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *partition_info_msg_t) Deref() {
	if x.reffe44d6cd == nil {
		return
	}
	x.last_update = (time_t)(x.reffe44d6cd.last_update)
	x.record_count = (uint32_t)(x.reffe44d6cd.record_count)
	packSPartition_info_t(x.partition_array, x.reffe44d6cd.partition_array)
}

// allocWill_run_response_msg_tMemory allocates memory for type C.will_run_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocWill_run_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfWill_run_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfWill_run_response_msg_tValue = unsafe.Sizeof([1]C.will_run_response_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *will_run_response_msg_t) Ref() *C.will_run_response_msg_t {
	if x == nil {
		return nil
	}
	return x.reffc5b749c
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *will_run_response_msg_t) Free() {
	if x != nil && x.allocsfc5b749c != nil {
		x.allocsfc5b749c.(*cgoAllocMap).Free()
		x.reffc5b749c = nil
	}
}

// Newwill_run_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newwill_run_response_msg_tRef(ref unsafe.Pointer) *will_run_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(will_run_response_msg_t)
	obj.reffc5b749c = (*C.will_run_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *will_run_response_msg_t) PassRef() (*C.will_run_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reffc5b749c != nil {
		return x.reffc5b749c, nil
	}
	memfc5b749c := allocWill_run_response_msg_tMemory(1)
	reffc5b749c := (*C.will_run_response_msg_t)(memfc5b749c)
	allocsfc5b749c := new(cgoAllocMap)
	allocsfc5b749c.Add(memfc5b749c)

	var cjob_id_allocs *cgoAllocMap
	reffc5b749c.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocsfc5b749c.Borrow(cjob_id_allocs)

	var cjob_submit_user_msg_allocs *cgoAllocMap
	reffc5b749c.job_submit_user_msg, cjob_submit_user_msg_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_submit_user_msg)).Data)), cgoAllocsUnknown
	allocsfc5b749c.Borrow(cjob_submit_user_msg_allocs)

	var cnode_list_allocs *cgoAllocMap
	reffc5b749c.node_list, cnode_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_list)).Data)), cgoAllocsUnknown
	allocsfc5b749c.Borrow(cnode_list_allocs)

	var cpreemptee_job_id_allocs *cgoAllocMap
	reffc5b749c.preemptee_job_id, cpreemptee_job_id_allocs = *(*C.List)(unsafe.Pointer(&x.preemptee_job_id)), cgoAllocsUnknown
	allocsfc5b749c.Borrow(cpreemptee_job_id_allocs)

	var cproc_cnt_allocs *cgoAllocMap
	reffc5b749c.proc_cnt, cproc_cnt_allocs = (C.uint32_t)(x.proc_cnt), cgoAllocsUnknown
	allocsfc5b749c.Borrow(cproc_cnt_allocs)

	var cstart_time_allocs *cgoAllocMap
	reffc5b749c.start_time, cstart_time_allocs = (C.time_t)(x.start_time), cgoAllocsUnknown
	allocsfc5b749c.Borrow(cstart_time_allocs)

	var csys_usage_per_allocs *cgoAllocMap
	reffc5b749c.sys_usage_per, csys_usage_per_allocs = (C.double)(x.sys_usage_per), cgoAllocsUnknown
	allocsfc5b749c.Borrow(csys_usage_per_allocs)

	x.reffc5b749c = reffc5b749c
	x.allocsfc5b749c = allocsfc5b749c
	return reffc5b749c, allocsfc5b749c

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x will_run_response_msg_t) PassValue() (C.will_run_response_msg_t, *cgoAllocMap) {
	if x.reffc5b749c != nil {
		return *x.reffc5b749c, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *will_run_response_msg_t) Deref() {
	if x.reffc5b749c == nil {
		return
	}
	x.job_id = (uint32_t)(x.reffc5b749c.job_id)
	hxf687e32 := (*sliceHeader)(unsafe.Pointer(&x.job_submit_user_msg))
	hxf687e32.Data = uintptr(unsafe.Pointer(x.reffc5b749c.job_submit_user_msg))
	hxf687e32.Cap = 0x7fffffff
	// hxf687e32.Len = ?

	hxf99a62b := (*sliceHeader)(unsafe.Pointer(&x.node_list))
	hxf99a62b.Data = uintptr(unsafe.Pointer(x.reffc5b749c.node_list))
	hxf99a62b.Cap = 0x7fffffff
	// hxf99a62b.Len = ?

	x.preemptee_job_id = *(*List)(unsafe.Pointer(&x.reffc5b749c.preemptee_job_id))
	x.proc_cnt = (uint32_t)(x.reffc5b749c.proc_cnt)
	x.start_time = (time_t)(x.reffc5b749c.start_time)
	x.sys_usage_per = (float64)(x.reffc5b749c.sys_usage_per)
}

// allocBlock_job_info_tMemory allocates memory for type C.block_job_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBlock_job_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBlock_job_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfBlock_job_info_tValue = unsafe.Sizeof([1]C.block_job_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *block_job_info_t) Ref() *C.block_job_info_t {
	if x == nil {
		return nil
	}
	return x.refb25fc08b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *block_job_info_t) Free() {
	if x != nil && x.allocsb25fc08b != nil {
		x.allocsb25fc08b.(*cgoAllocMap).Free()
		x.refb25fc08b = nil
	}
}

// Newblock_job_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newblock_job_info_tRef(ref unsafe.Pointer) *block_job_info_t {
	if ref == nil {
		return nil
	}
	obj := new(block_job_info_t)
	obj.refb25fc08b = (*C.block_job_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *block_job_info_t) PassRef() (*C.block_job_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb25fc08b != nil {
		return x.refb25fc08b, nil
	}
	memb25fc08b := allocBlock_job_info_tMemory(1)
	refb25fc08b := (*C.block_job_info_t)(memb25fc08b)
	allocsb25fc08b := new(cgoAllocMap)
	allocsb25fc08b.Add(memb25fc08b)

	var ccnodes_allocs *cgoAllocMap
	refb25fc08b.cnodes, ccnodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cnodes)).Data)), cgoAllocsUnknown
	allocsb25fc08b.Borrow(ccnodes_allocs)

	var ccnode_inx_allocs *cgoAllocMap
	refb25fc08b.cnode_inx, ccnode_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cnode_inx)).Data)), cgoAllocsUnknown
	allocsb25fc08b.Borrow(ccnode_inx_allocs)

	var cjob_id_allocs *cgoAllocMap
	refb25fc08b.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocsb25fc08b.Borrow(cjob_id_allocs)

	var cjob_ptr_allocs *cgoAllocMap
	refb25fc08b.job_ptr, cjob_ptr_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.job_ptr)), cgoAllocsUnknown
	allocsb25fc08b.Borrow(cjob_ptr_allocs)

	var cuser_id_allocs *cgoAllocMap
	refb25fc08b.user_id, cuser_id_allocs = (C.uint32_t)(x.user_id), cgoAllocsUnknown
	allocsb25fc08b.Borrow(cuser_id_allocs)

	var cuser_name_allocs *cgoAllocMap
	refb25fc08b.user_name, cuser_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.user_name)).Data)), cgoAllocsUnknown
	allocsb25fc08b.Borrow(cuser_name_allocs)

	x.refb25fc08b = refb25fc08b
	x.allocsb25fc08b = allocsb25fc08b
	return refb25fc08b, allocsb25fc08b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x block_job_info_t) PassValue() (C.block_job_info_t, *cgoAllocMap) {
	if x.refb25fc08b != nil {
		return *x.refb25fc08b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *block_job_info_t) Deref() {
	if x.refb25fc08b == nil {
		return
	}
	hxf411be2 := (*sliceHeader)(unsafe.Pointer(&x.cnodes))
	hxf411be2.Data = uintptr(unsafe.Pointer(x.refb25fc08b.cnodes))
	hxf411be2.Cap = 0x7fffffff
	// hxf411be2.Len = ?

	hxf869803 := (*sliceHeader)(unsafe.Pointer(&x.cnode_inx))
	hxf869803.Data = uintptr(unsafe.Pointer(x.refb25fc08b.cnode_inx))
	hxf869803.Cap = 0x7fffffff
	// hxf869803.Len = ?

	x.job_id = (uint32_t)(x.refb25fc08b.job_id)
	x.job_ptr = (unsafe.Pointer)(unsafe.Pointer(x.refb25fc08b.job_ptr))
	x.user_id = (uint32_t)(x.refb25fc08b.user_id)
	hxf76ae99 := (*sliceHeader)(unsafe.Pointer(&x.user_name))
	hxf76ae99.Data = uintptr(unsafe.Pointer(x.refb25fc08b.user_name))
	hxf76ae99.Cap = 0x7fffffff
	// hxf76ae99.Len = ?

}

// allocBlock_info_tMemory allocates memory for type C.block_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBlock_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBlock_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfBlock_info_tValue = unsafe.Sizeof([1]C.block_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *block_info_t) Ref() *C.block_info_t {
	if x == nil {
		return nil
	}
	return x.refc618b0ea
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *block_info_t) Free() {
	if x != nil && x.allocsc618b0ea != nil {
		x.allocsc618b0ea.(*cgoAllocMap).Free()
		x.refc618b0ea = nil
	}
}

// Newblock_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newblock_info_tRef(ref unsafe.Pointer) *block_info_t {
	if ref == nil {
		return nil
	}
	obj := new(block_info_t)
	obj.refc618b0ea = (*C.block_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *block_info_t) PassRef() (*C.block_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc618b0ea != nil {
		return x.refc618b0ea, nil
	}
	memc618b0ea := allocBlock_info_tMemory(1)
	refc618b0ea := (*C.block_info_t)(memc618b0ea)
	allocsc618b0ea := new(cgoAllocMap)
	allocsc618b0ea.Add(memc618b0ea)

	var cbg_block_id_allocs *cgoAllocMap
	refc618b0ea.bg_block_id, cbg_block_id_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.bg_block_id)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cbg_block_id_allocs)

	var cblrtsimage_allocs *cgoAllocMap
	refc618b0ea.blrtsimage, cblrtsimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.blrtsimage)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cblrtsimage_allocs)

	var cconn_type_allocs *cgoAllocMap
	refc618b0ea.conn_type, cconn_type_allocs = *(*[5]C.uint16_t)(unsafe.Pointer(&x.conn_type)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cconn_type_allocs)

	var ccnode_cnt_allocs *cgoAllocMap
	refc618b0ea.cnode_cnt, ccnode_cnt_allocs = (C.uint32_t)(x.cnode_cnt), cgoAllocsUnknown
	allocsc618b0ea.Borrow(ccnode_cnt_allocs)

	var ccnode_err_cnt_allocs *cgoAllocMap
	refc618b0ea.cnode_err_cnt, ccnode_err_cnt_allocs = (C.uint32_t)(x.cnode_err_cnt), cgoAllocsUnknown
	allocsc618b0ea.Borrow(ccnode_err_cnt_allocs)

	var cionode_inx_allocs *cgoAllocMap
	refc618b0ea.ionode_inx, cionode_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ionode_inx)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cionode_inx_allocs)

	var cionode_str_allocs *cgoAllocMap
	refc618b0ea.ionode_str, cionode_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ionode_str)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cionode_str_allocs)

	var cjob_list_allocs *cgoAllocMap
	refc618b0ea.job_list, cjob_list_allocs = *(*C.List)(unsafe.Pointer(&x.job_list)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cjob_list_allocs)

	var clinuximage_allocs *cgoAllocMap
	refc618b0ea.linuximage, clinuximage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.linuximage)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(clinuximage_allocs)

	var cmloaderimage_allocs *cgoAllocMap
	refc618b0ea.mloaderimage, cmloaderimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mloaderimage)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cmloaderimage_allocs)

	var cmp_inx_allocs *cgoAllocMap
	refc618b0ea.mp_inx, cmp_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mp_inx)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cmp_inx_allocs)

	var cmp_str_allocs *cgoAllocMap
	refc618b0ea.mp_str, cmp_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mp_str)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cmp_str_allocs)

	var cnode_use_allocs *cgoAllocMap
	refc618b0ea.node_use, cnode_use_allocs = (C.uint16_t)(x.node_use), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cnode_use_allocs)

	var cramdiskimage_allocs *cgoAllocMap
	refc618b0ea.ramdiskimage, cramdiskimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ramdiskimage)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cramdiskimage_allocs)

	var creason_allocs *cgoAllocMap
	refc618b0ea.reason, creason_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.reason)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(creason_allocs)

	var cstate_allocs *cgoAllocMap
	refc618b0ea.state, cstate_allocs = (C.uint16_t)(x.state), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cstate_allocs)

	x.refc618b0ea = refc618b0ea
	x.allocsc618b0ea = allocsc618b0ea
	return refc618b0ea, allocsc618b0ea

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x block_info_t) PassValue() (C.block_info_t, *cgoAllocMap) {
	if x.refc618b0ea != nil {
		return *x.refc618b0ea, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *block_info_t) Deref() {
	if x.refc618b0ea == nil {
		return
	}
	hxf9cb50d := (*sliceHeader)(unsafe.Pointer(&x.bg_block_id))
	hxf9cb50d.Data = uintptr(unsafe.Pointer(x.refc618b0ea.bg_block_id))
	hxf9cb50d.Cap = 0x7fffffff
	// hxf9cb50d.Len = ?

	hxfadb11b := (*sliceHeader)(unsafe.Pointer(&x.blrtsimage))
	hxfadb11b.Data = uintptr(unsafe.Pointer(x.refc618b0ea.blrtsimage))
	hxfadb11b.Cap = 0x7fffffff
	// hxfadb11b.Len = ?

	x.conn_type = *(*[5]uint16_t)(unsafe.Pointer(&x.refc618b0ea.conn_type))
	x.cnode_cnt = (uint32_t)(x.refc618b0ea.cnode_cnt)
	x.cnode_err_cnt = (uint32_t)(x.refc618b0ea.cnode_err_cnt)
	hxf5c4bf0 := (*sliceHeader)(unsafe.Pointer(&x.ionode_inx))
	hxf5c4bf0.Data = uintptr(unsafe.Pointer(x.refc618b0ea.ionode_inx))
	hxf5c4bf0.Cap = 0x7fffffff
	// hxf5c4bf0.Len = ?

	hxfc6eccc := (*sliceHeader)(unsafe.Pointer(&x.ionode_str))
	hxfc6eccc.Data = uintptr(unsafe.Pointer(x.refc618b0ea.ionode_str))
	hxfc6eccc.Cap = 0x7fffffff
	// hxfc6eccc.Len = ?

	x.job_list = *(*List)(unsafe.Pointer(&x.refc618b0ea.job_list))
	hxf4b8985 := (*sliceHeader)(unsafe.Pointer(&x.linuximage))
	hxf4b8985.Data = uintptr(unsafe.Pointer(x.refc618b0ea.linuximage))
	hxf4b8985.Cap = 0x7fffffff
	// hxf4b8985.Len = ?

	hxf7af91f := (*sliceHeader)(unsafe.Pointer(&x.mloaderimage))
	hxf7af91f.Data = uintptr(unsafe.Pointer(x.refc618b0ea.mloaderimage))
	hxf7af91f.Cap = 0x7fffffff
	// hxf7af91f.Len = ?

	hxff09baa := (*sliceHeader)(unsafe.Pointer(&x.mp_inx))
	hxff09baa.Data = uintptr(unsafe.Pointer(x.refc618b0ea.mp_inx))
	hxff09baa.Cap = 0x7fffffff
	// hxff09baa.Len = ?

	hxfc612c4 := (*sliceHeader)(unsafe.Pointer(&x.mp_str))
	hxfc612c4.Data = uintptr(unsafe.Pointer(x.refc618b0ea.mp_str))
	hxfc612c4.Cap = 0x7fffffff
	// hxfc612c4.Len = ?

	x.node_use = (uint16_t)(x.refc618b0ea.node_use)
	hxfc072e2 := (*sliceHeader)(unsafe.Pointer(&x.ramdiskimage))
	hxfc072e2.Data = uintptr(unsafe.Pointer(x.refc618b0ea.ramdiskimage))
	hxfc072e2.Cap = 0x7fffffff
	// hxfc072e2.Len = ?

	hxfbebdf9 := (*sliceHeader)(unsafe.Pointer(&x.reason))
	hxfbebdf9.Data = uintptr(unsafe.Pointer(x.refc618b0ea.reason))
	hxfbebdf9.Cap = 0x7fffffff
	// hxfbebdf9.Len = ?

	x.state = (uint16_t)(x.refc618b0ea.state)
}

// allocBlock_info_msg_tMemory allocates memory for type C.block_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBlock_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBlock_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfBlock_info_msg_tValue = unsafe.Sizeof([1]C.block_info_msg_t{})

// unpackSBlock_info_t transforms a sliced Go data structure into plain C format.
func unpackSBlock_info_t(x []block_info_t) (unpacked *C.block_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.block_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocBlock_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.block_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.block_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSBlock_info_t reads sliced Go data structure out from plain C format.
func packSBlock_info_t(v []block_info_t, ptr0 *C.block_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfBlock_info_tValue]C.block_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newblock_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *block_info_msg_t) Ref() *C.block_info_msg_t {
	if x == nil {
		return nil
	}
	return x.ref9e6dac16
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *block_info_msg_t) Free() {
	if x != nil && x.allocs9e6dac16 != nil {
		x.allocs9e6dac16.(*cgoAllocMap).Free()
		x.ref9e6dac16 = nil
	}
}

// Newblock_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newblock_info_msg_tRef(ref unsafe.Pointer) *block_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(block_info_msg_t)
	obj.ref9e6dac16 = (*C.block_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *block_info_msg_t) PassRef() (*C.block_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9e6dac16 != nil {
		return x.ref9e6dac16, nil
	}
	mem9e6dac16 := allocBlock_info_msg_tMemory(1)
	ref9e6dac16 := (*C.block_info_msg_t)(mem9e6dac16)
	allocs9e6dac16 := new(cgoAllocMap)
	allocs9e6dac16.Add(mem9e6dac16)

	var cblock_array_allocs *cgoAllocMap
	ref9e6dac16.block_array, cblock_array_allocs = unpackSBlock_info_t(x.block_array)
	allocs9e6dac16.Borrow(cblock_array_allocs)

	var clast_update_allocs *cgoAllocMap
	ref9e6dac16.last_update, clast_update_allocs = (C.time_t)(x.last_update), cgoAllocsUnknown
	allocs9e6dac16.Borrow(clast_update_allocs)

	var crecord_count_allocs *cgoAllocMap
	ref9e6dac16.record_count, crecord_count_allocs = (C.uint32_t)(x.record_count), cgoAllocsUnknown
	allocs9e6dac16.Borrow(crecord_count_allocs)

	x.ref9e6dac16 = ref9e6dac16
	x.allocs9e6dac16 = allocs9e6dac16
	return ref9e6dac16, allocs9e6dac16

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x block_info_msg_t) PassValue() (C.block_info_msg_t, *cgoAllocMap) {
	if x.ref9e6dac16 != nil {
		return *x.ref9e6dac16, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *block_info_msg_t) Deref() {
	if x.ref9e6dac16 == nil {
		return
	}
	packSBlock_info_t(x.block_array, x.ref9e6dac16.block_array)
	x.last_update = (time_t)(x.ref9e6dac16.last_update)
	x.record_count = (uint32_t)(x.ref9e6dac16.record_count)
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *update_block_msg_t) Ref() *C.block_info_t {
	if x == nil {
		return nil
	}
	return x.refc618b0ea
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *update_block_msg_t) Free() {
	if x != nil && x.allocsc618b0ea != nil {
		x.allocsc618b0ea.(*cgoAllocMap).Free()
		x.refc618b0ea = nil
	}
}

// Newupdate_block_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newupdate_block_msg_tRef(ref unsafe.Pointer) *update_block_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(update_block_msg_t)
	obj.refc618b0ea = (*C.block_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *update_block_msg_t) PassRef() (*C.block_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc618b0ea != nil {
		return x.refc618b0ea, nil
	}
	memc618b0ea := allocBlock_info_tMemory(1)
	refc618b0ea := (*C.block_info_t)(memc618b0ea)
	allocsc618b0ea := new(cgoAllocMap)
	allocsc618b0ea.Add(memc618b0ea)

	var cbg_block_id_allocs *cgoAllocMap
	refc618b0ea.bg_block_id, cbg_block_id_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.bg_block_id)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cbg_block_id_allocs)

	var cblrtsimage_allocs *cgoAllocMap
	refc618b0ea.blrtsimage, cblrtsimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.blrtsimage)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cblrtsimage_allocs)

	var cconn_type_allocs *cgoAllocMap
	refc618b0ea.conn_type, cconn_type_allocs = *(*[5]C.uint16_t)(unsafe.Pointer(&x.conn_type)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cconn_type_allocs)

	var ccnode_cnt_allocs *cgoAllocMap
	refc618b0ea.cnode_cnt, ccnode_cnt_allocs = (C.uint32_t)(x.cnode_cnt), cgoAllocsUnknown
	allocsc618b0ea.Borrow(ccnode_cnt_allocs)

	var ccnode_err_cnt_allocs *cgoAllocMap
	refc618b0ea.cnode_err_cnt, ccnode_err_cnt_allocs = (C.uint32_t)(x.cnode_err_cnt), cgoAllocsUnknown
	allocsc618b0ea.Borrow(ccnode_err_cnt_allocs)

	var cionode_inx_allocs *cgoAllocMap
	refc618b0ea.ionode_inx, cionode_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ionode_inx)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cionode_inx_allocs)

	var cionode_str_allocs *cgoAllocMap
	refc618b0ea.ionode_str, cionode_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ionode_str)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cionode_str_allocs)

	var cjob_list_allocs *cgoAllocMap
	refc618b0ea.job_list, cjob_list_allocs = *(*C.List)(unsafe.Pointer(&x.job_list)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cjob_list_allocs)

	var clinuximage_allocs *cgoAllocMap
	refc618b0ea.linuximage, clinuximage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.linuximage)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(clinuximage_allocs)

	var cmloaderimage_allocs *cgoAllocMap
	refc618b0ea.mloaderimage, cmloaderimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mloaderimage)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cmloaderimage_allocs)

	var cmp_inx_allocs *cgoAllocMap
	refc618b0ea.mp_inx, cmp_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mp_inx)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cmp_inx_allocs)

	var cmp_str_allocs *cgoAllocMap
	refc618b0ea.mp_str, cmp_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mp_str)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cmp_str_allocs)

	var cnode_use_allocs *cgoAllocMap
	refc618b0ea.node_use, cnode_use_allocs = (C.uint16_t)(x.node_use), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cnode_use_allocs)

	var cramdiskimage_allocs *cgoAllocMap
	refc618b0ea.ramdiskimage, cramdiskimage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ramdiskimage)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cramdiskimage_allocs)

	var creason_allocs *cgoAllocMap
	refc618b0ea.reason, creason_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.reason)).Data)), cgoAllocsUnknown
	allocsc618b0ea.Borrow(creason_allocs)

	var cstate_allocs *cgoAllocMap
	refc618b0ea.state, cstate_allocs = (C.uint16_t)(x.state), cgoAllocsUnknown
	allocsc618b0ea.Borrow(cstate_allocs)

	x.refc618b0ea = refc618b0ea
	x.allocsc618b0ea = allocsc618b0ea
	return refc618b0ea, allocsc618b0ea

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x update_block_msg_t) PassValue() (C.block_info_t, *cgoAllocMap) {
	if x.refc618b0ea != nil {
		return *x.refc618b0ea, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *update_block_msg_t) Deref() {
	if x.refc618b0ea == nil {
		return
	}
	hxf872ae4 := (*sliceHeader)(unsafe.Pointer(&x.bg_block_id))
	hxf872ae4.Data = uintptr(unsafe.Pointer(x.refc618b0ea.bg_block_id))
	hxf872ae4.Cap = 0x7fffffff
	// hxf872ae4.Len = ?

	hxf775df1 := (*sliceHeader)(unsafe.Pointer(&x.blrtsimage))
	hxf775df1.Data = uintptr(unsafe.Pointer(x.refc618b0ea.blrtsimage))
	hxf775df1.Cap = 0x7fffffff
	// hxf775df1.Len = ?

	x.conn_type = *(*[5]uint16_t)(unsafe.Pointer(&x.refc618b0ea.conn_type))
	x.cnode_cnt = (uint32_t)(x.refc618b0ea.cnode_cnt)
	x.cnode_err_cnt = (uint32_t)(x.refc618b0ea.cnode_err_cnt)
	hxfdda13e := (*sliceHeader)(unsafe.Pointer(&x.ionode_inx))
	hxfdda13e.Data = uintptr(unsafe.Pointer(x.refc618b0ea.ionode_inx))
	hxfdda13e.Cap = 0x7fffffff
	// hxfdda13e.Len = ?

	hxfe56971 := (*sliceHeader)(unsafe.Pointer(&x.ionode_str))
	hxfe56971.Data = uintptr(unsafe.Pointer(x.refc618b0ea.ionode_str))
	hxfe56971.Cap = 0x7fffffff
	// hxfe56971.Len = ?

	x.job_list = *(*List)(unsafe.Pointer(&x.refc618b0ea.job_list))
	hxfbce0c9 := (*sliceHeader)(unsafe.Pointer(&x.linuximage))
	hxfbce0c9.Data = uintptr(unsafe.Pointer(x.refc618b0ea.linuximage))
	hxfbce0c9.Cap = 0x7fffffff
	// hxfbce0c9.Len = ?

	hxf5eb5ca := (*sliceHeader)(unsafe.Pointer(&x.mloaderimage))
	hxf5eb5ca.Data = uintptr(unsafe.Pointer(x.refc618b0ea.mloaderimage))
	hxf5eb5ca.Cap = 0x7fffffff
	// hxf5eb5ca.Len = ?

	hxfcf0e04 := (*sliceHeader)(unsafe.Pointer(&x.mp_inx))
	hxfcf0e04.Data = uintptr(unsafe.Pointer(x.refc618b0ea.mp_inx))
	hxfcf0e04.Cap = 0x7fffffff
	// hxfcf0e04.Len = ?

	hxfd6f73b := (*sliceHeader)(unsafe.Pointer(&x.mp_str))
	hxfd6f73b.Data = uintptr(unsafe.Pointer(x.refc618b0ea.mp_str))
	hxfd6f73b.Cap = 0x7fffffff
	// hxfd6f73b.Len = ?

	x.node_use = (uint16_t)(x.refc618b0ea.node_use)
	hxfbe1d87 := (*sliceHeader)(unsafe.Pointer(&x.ramdiskimage))
	hxfbe1d87.Data = uintptr(unsafe.Pointer(x.refc618b0ea.ramdiskimage))
	hxfbe1d87.Cap = 0x7fffffff
	// hxfbe1d87.Len = ?

	hxfe194bd := (*sliceHeader)(unsafe.Pointer(&x.reason))
	hxfe194bd.Data = uintptr(unsafe.Pointer(x.refc618b0ea.reason))
	hxfe194bd.Cap = 0x7fffffff
	// hxfe194bd.Len = ?

	x.state = (uint16_t)(x.refc618b0ea.state)
}

// allocResv_core_spec_tMemory allocates memory for type C.resv_core_spec_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocResv_core_spec_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfResv_core_spec_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfResv_core_spec_tValue = unsafe.Sizeof([1]C.resv_core_spec_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *resv_core_spec_t) Ref() *C.resv_core_spec_t {
	if x == nil {
		return nil
	}
	return x.ref256bbcbd
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *resv_core_spec_t) Free() {
	if x != nil && x.allocs256bbcbd != nil {
		x.allocs256bbcbd.(*cgoAllocMap).Free()
		x.ref256bbcbd = nil
	}
}

// Newresv_core_spec_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newresv_core_spec_tRef(ref unsafe.Pointer) *resv_core_spec_t {
	if ref == nil {
		return nil
	}
	obj := new(resv_core_spec_t)
	obj.ref256bbcbd = (*C.resv_core_spec_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *resv_core_spec_t) PassRef() (*C.resv_core_spec_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref256bbcbd != nil {
		return x.ref256bbcbd, nil
	}
	mem256bbcbd := allocResv_core_spec_tMemory(1)
	ref256bbcbd := (*C.resv_core_spec_t)(mem256bbcbd)
	allocs256bbcbd := new(cgoAllocMap)
	allocs256bbcbd.Add(mem256bbcbd)

	var cnode_name_allocs *cgoAllocMap
	ref256bbcbd.node_name, cnode_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_name)).Data)), cgoAllocsUnknown
	allocs256bbcbd.Borrow(cnode_name_allocs)

	var ccore_id_allocs *cgoAllocMap
	ref256bbcbd.core_id, ccore_id_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.core_id)).Data)), cgoAllocsUnknown
	allocs256bbcbd.Borrow(ccore_id_allocs)

	x.ref256bbcbd = ref256bbcbd
	x.allocs256bbcbd = allocs256bbcbd
	return ref256bbcbd, allocs256bbcbd

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x resv_core_spec_t) PassValue() (C.resv_core_spec_t, *cgoAllocMap) {
	if x.ref256bbcbd != nil {
		return *x.ref256bbcbd, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *resv_core_spec_t) Deref() {
	if x.ref256bbcbd == nil {
		return
	}
	hxf0df3c3 := (*sliceHeader)(unsafe.Pointer(&x.node_name))
	hxf0df3c3.Data = uintptr(unsafe.Pointer(x.ref256bbcbd.node_name))
	hxf0df3c3.Cap = 0x7fffffff
	// hxf0df3c3.Len = ?

	hxfa541ca := (*sliceHeader)(unsafe.Pointer(&x.core_id))
	hxfa541ca.Data = uintptr(unsafe.Pointer(x.ref256bbcbd.core_id))
	hxfa541ca.Cap = 0x7fffffff
	// hxfa541ca.Len = ?

}

// allocReserve_info_tMemory allocates memory for type C.reserve_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocReserve_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfReserve_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfReserve_info_tValue = unsafe.Sizeof([1]C.reserve_info_t{})

// unpackSResv_core_spec_t transforms a sliced Go data structure into plain C format.
func unpackSResv_core_spec_t(x []resv_core_spec_t) (unpacked *C.resv_core_spec_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.resv_core_spec_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocResv_core_spec_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.resv_core_spec_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.resv_core_spec_t)(unsafe.Pointer(h.Data))
	return
}

// packSResv_core_spec_t reads sliced Go data structure out from plain C format.
func packSResv_core_spec_t(v []resv_core_spec_t, ptr0 *C.resv_core_spec_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfResv_core_spec_tValue]C.resv_core_spec_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newresv_core_spec_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *reserve_info_t) Ref() *C.reserve_info_t {
	if x == nil {
		return nil
	}
	return x.refb9d691a5
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *reserve_info_t) Free() {
	if x != nil && x.allocsb9d691a5 != nil {
		x.allocsb9d691a5.(*cgoAllocMap).Free()
		x.refb9d691a5 = nil
	}
}

// Newreserve_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newreserve_info_tRef(ref unsafe.Pointer) *reserve_info_t {
	if ref == nil {
		return nil
	}
	obj := new(reserve_info_t)
	obj.refb9d691a5 = (*C.reserve_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *reserve_info_t) PassRef() (*C.reserve_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb9d691a5 != nil {
		return x.refb9d691a5, nil
	}
	memb9d691a5 := allocReserve_info_tMemory(1)
	refb9d691a5 := (*C.reserve_info_t)(memb9d691a5)
	allocsb9d691a5 := new(cgoAllocMap)
	allocsb9d691a5.Add(memb9d691a5)

	var caccounts_allocs *cgoAllocMap
	refb9d691a5.accounts, caccounts_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.accounts)).Data)), cgoAllocsUnknown
	allocsb9d691a5.Borrow(caccounts_allocs)

	var cburst_buffer_allocs *cgoAllocMap
	refb9d691a5.burst_buffer, cburst_buffer_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.burst_buffer)).Data)), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cburst_buffer_allocs)

	var ccore_cnt_allocs *cgoAllocMap
	refb9d691a5.core_cnt, ccore_cnt_allocs = (C.uint32_t)(x.core_cnt), cgoAllocsUnknown
	allocsb9d691a5.Borrow(ccore_cnt_allocs)

	var ccore_spec_cnt_allocs *cgoAllocMap
	refb9d691a5.core_spec_cnt, ccore_spec_cnt_allocs = (C.uint32_t)(x.core_spec_cnt), cgoAllocsUnknown
	allocsb9d691a5.Borrow(ccore_spec_cnt_allocs)

	var ccore_spec_allocs *cgoAllocMap
	refb9d691a5.core_spec, ccore_spec_allocs = unpackSResv_core_spec_t(x.core_spec)
	allocsb9d691a5.Borrow(ccore_spec_allocs)

	var cend_time_allocs *cgoAllocMap
	refb9d691a5.end_time, cend_time_allocs = (C.time_t)(x.end_time), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cend_time_allocs)

	var cfeatures_allocs *cgoAllocMap
	refb9d691a5.features, cfeatures_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.features)).Data)), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cfeatures_allocs)

	var cflags_allocs *cgoAllocMap
	refb9d691a5.flags, cflags_allocs = (C.uint32_t)(x.flags), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cflags_allocs)

	var clicenses_allocs *cgoAllocMap
	refb9d691a5.licenses, clicenses_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.licenses)).Data)), cgoAllocsUnknown
	allocsb9d691a5.Borrow(clicenses_allocs)

	var cname_allocs *cgoAllocMap
	refb9d691a5.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cname_allocs)

	var cnode_cnt_allocs *cgoAllocMap
	refb9d691a5.node_cnt, cnode_cnt_allocs = (C.uint32_t)(x.node_cnt), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cnode_cnt_allocs)

	var cnode_inx_allocs *cgoAllocMap
	refb9d691a5.node_inx, cnode_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_inx)).Data)), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cnode_inx_allocs)

	var cnode_list_allocs *cgoAllocMap
	refb9d691a5.node_list, cnode_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_list)).Data)), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cnode_list_allocs)

	var cpartition_allocs *cgoAllocMap
	refb9d691a5.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partition)).Data)), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cpartition_allocs)

	var cstart_time_allocs *cgoAllocMap
	refb9d691a5.start_time, cstart_time_allocs = (C.time_t)(x.start_time), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cstart_time_allocs)

	var cresv_watts_allocs *cgoAllocMap
	refb9d691a5.resv_watts, cresv_watts_allocs = (C.uint32_t)(x.resv_watts), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cresv_watts_allocs)

	var ctres_str_allocs *cgoAllocMap
	refb9d691a5.tres_str, ctres_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_str)).Data)), cgoAllocsUnknown
	allocsb9d691a5.Borrow(ctres_str_allocs)

	var cusers_allocs *cgoAllocMap
	refb9d691a5.users, cusers_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.users)).Data)), cgoAllocsUnknown
	allocsb9d691a5.Borrow(cusers_allocs)

	x.refb9d691a5 = refb9d691a5
	x.allocsb9d691a5 = allocsb9d691a5
	return refb9d691a5, allocsb9d691a5

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x reserve_info_t) PassValue() (C.reserve_info_t, *cgoAllocMap) {
	if x.refb9d691a5 != nil {
		return *x.refb9d691a5, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *reserve_info_t) Deref() {
	if x.refb9d691a5 == nil {
		return
	}
	hxf165f6f := (*sliceHeader)(unsafe.Pointer(&x.accounts))
	hxf165f6f.Data = uintptr(unsafe.Pointer(x.refb9d691a5.accounts))
	hxf165f6f.Cap = 0x7fffffff
	// hxf165f6f.Len = ?

	hxfadc2dd := (*sliceHeader)(unsafe.Pointer(&x.burst_buffer))
	hxfadc2dd.Data = uintptr(unsafe.Pointer(x.refb9d691a5.burst_buffer))
	hxfadc2dd.Cap = 0x7fffffff
	// hxfadc2dd.Len = ?

	x.core_cnt = (uint32_t)(x.refb9d691a5.core_cnt)
	x.core_spec_cnt = (uint32_t)(x.refb9d691a5.core_spec_cnt)
	packSResv_core_spec_t(x.core_spec, x.refb9d691a5.core_spec)
	x.end_time = (time_t)(x.refb9d691a5.end_time)
	hxf4dda01 := (*sliceHeader)(unsafe.Pointer(&x.features))
	hxf4dda01.Data = uintptr(unsafe.Pointer(x.refb9d691a5.features))
	hxf4dda01.Cap = 0x7fffffff
	// hxf4dda01.Len = ?

	x.flags = (uint32_t)(x.refb9d691a5.flags)
	hxfbec3ee := (*sliceHeader)(unsafe.Pointer(&x.licenses))
	hxfbec3ee.Data = uintptr(unsafe.Pointer(x.refb9d691a5.licenses))
	hxfbec3ee.Cap = 0x7fffffff
	// hxfbec3ee.Len = ?

	hxf2f6263 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf2f6263.Data = uintptr(unsafe.Pointer(x.refb9d691a5.name))
	hxf2f6263.Cap = 0x7fffffff
	// hxf2f6263.Len = ?

	x.node_cnt = (uint32_t)(x.refb9d691a5.node_cnt)
	hxf6f3a66 := (*sliceHeader)(unsafe.Pointer(&x.node_inx))
	hxf6f3a66.Data = uintptr(unsafe.Pointer(x.refb9d691a5.node_inx))
	hxf6f3a66.Cap = 0x7fffffff
	// hxf6f3a66.Len = ?

	hxfa2d4f3 := (*sliceHeader)(unsafe.Pointer(&x.node_list))
	hxfa2d4f3.Data = uintptr(unsafe.Pointer(x.refb9d691a5.node_list))
	hxfa2d4f3.Cap = 0x7fffffff
	// hxfa2d4f3.Len = ?

	hxfdd7e32 := (*sliceHeader)(unsafe.Pointer(&x.partition))
	hxfdd7e32.Data = uintptr(unsafe.Pointer(x.refb9d691a5.partition))
	hxfdd7e32.Cap = 0x7fffffff
	// hxfdd7e32.Len = ?

	x.start_time = (time_t)(x.refb9d691a5.start_time)
	x.resv_watts = (uint32_t)(x.refb9d691a5.resv_watts)
	hxf6475f2 := (*sliceHeader)(unsafe.Pointer(&x.tres_str))
	hxf6475f2.Data = uintptr(unsafe.Pointer(x.refb9d691a5.tres_str))
	hxf6475f2.Cap = 0x7fffffff
	// hxf6475f2.Len = ?

	hxf561254 := (*sliceHeader)(unsafe.Pointer(&x.users))
	hxf561254.Data = uintptr(unsafe.Pointer(x.refb9d691a5.users))
	hxf561254.Cap = 0x7fffffff
	// hxf561254.Len = ?

}

// allocReserve_info_msg_tMemory allocates memory for type C.reserve_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocReserve_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfReserve_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfReserve_info_msg_tValue = unsafe.Sizeof([1]C.reserve_info_msg_t{})

// unpackSReserve_info_t transforms a sliced Go data structure into plain C format.
func unpackSReserve_info_t(x []reserve_info_t) (unpacked *C.reserve_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.reserve_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocReserve_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.reserve_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.reserve_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSReserve_info_t reads sliced Go data structure out from plain C format.
func packSReserve_info_t(v []reserve_info_t, ptr0 *C.reserve_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfReserve_info_tValue]C.reserve_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newreserve_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *reserve_info_msg_t) Ref() *C.reserve_info_msg_t {
	if x == nil {
		return nil
	}
	return x.reff1565ec7
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *reserve_info_msg_t) Free() {
	if x != nil && x.allocsf1565ec7 != nil {
		x.allocsf1565ec7.(*cgoAllocMap).Free()
		x.reff1565ec7 = nil
	}
}

// Newreserve_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newreserve_info_msg_tRef(ref unsafe.Pointer) *reserve_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(reserve_info_msg_t)
	obj.reff1565ec7 = (*C.reserve_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *reserve_info_msg_t) PassRef() (*C.reserve_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff1565ec7 != nil {
		return x.reff1565ec7, nil
	}
	memf1565ec7 := allocReserve_info_msg_tMemory(1)
	reff1565ec7 := (*C.reserve_info_msg_t)(memf1565ec7)
	allocsf1565ec7 := new(cgoAllocMap)
	allocsf1565ec7.Add(memf1565ec7)

	var clast_update_allocs *cgoAllocMap
	reff1565ec7.last_update, clast_update_allocs = (C.time_t)(x.last_update), cgoAllocsUnknown
	allocsf1565ec7.Borrow(clast_update_allocs)

	var crecord_count_allocs *cgoAllocMap
	reff1565ec7.record_count, crecord_count_allocs = (C.uint32_t)(x.record_count), cgoAllocsUnknown
	allocsf1565ec7.Borrow(crecord_count_allocs)

	var creservation_array_allocs *cgoAllocMap
	reff1565ec7.reservation_array, creservation_array_allocs = unpackSReserve_info_t(x.reservation_array)
	allocsf1565ec7.Borrow(creservation_array_allocs)

	x.reff1565ec7 = reff1565ec7
	x.allocsf1565ec7 = allocsf1565ec7
	return reff1565ec7, allocsf1565ec7

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x reserve_info_msg_t) PassValue() (C.reserve_info_msg_t, *cgoAllocMap) {
	if x.reff1565ec7 != nil {
		return *x.reff1565ec7, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *reserve_info_msg_t) Deref() {
	if x.reff1565ec7 == nil {
		return
	}
	x.last_update = (time_t)(x.reff1565ec7.last_update)
	x.record_count = (uint32_t)(x.reff1565ec7.record_count)
	packSReserve_info_t(x.reservation_array, x.reff1565ec7.reservation_array)
}

// allocResv_desc_msg_tMemory allocates memory for type C.resv_desc_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocResv_desc_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfResv_desc_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfResv_desc_msg_tValue = unsafe.Sizeof([1]C.resv_desc_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *resv_desc_msg_t) Ref() *C.resv_desc_msg_t {
	if x == nil {
		return nil
	}
	return x.refc2221b7f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *resv_desc_msg_t) Free() {
	if x != nil && x.allocsc2221b7f != nil {
		x.allocsc2221b7f.(*cgoAllocMap).Free()
		x.refc2221b7f = nil
	}
}

// Newresv_desc_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newresv_desc_msg_tRef(ref unsafe.Pointer) *resv_desc_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(resv_desc_msg_t)
	obj.refc2221b7f = (*C.resv_desc_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *resv_desc_msg_t) PassRef() (*C.resv_desc_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc2221b7f != nil {
		return x.refc2221b7f, nil
	}
	memc2221b7f := allocResv_desc_msg_tMemory(1)
	refc2221b7f := (*C.resv_desc_msg_t)(memc2221b7f)
	allocsc2221b7f := new(cgoAllocMap)
	allocsc2221b7f.Add(memc2221b7f)

	var caccounts_allocs *cgoAllocMap
	refc2221b7f.accounts, caccounts_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.accounts)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(caccounts_allocs)

	var cburst_buffer_allocs *cgoAllocMap
	refc2221b7f.burst_buffer, cburst_buffer_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.burst_buffer)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cburst_buffer_allocs)

	var ccore_cnt_allocs *cgoAllocMap
	refc2221b7f.core_cnt, ccore_cnt_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.core_cnt)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(ccore_cnt_allocs)

	var cduration_allocs *cgoAllocMap
	refc2221b7f.duration, cduration_allocs = (C.uint32_t)(x.duration), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cduration_allocs)

	var cend_time_allocs *cgoAllocMap
	refc2221b7f.end_time, cend_time_allocs = (C.time_t)(x.end_time), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cend_time_allocs)

	var cfeatures_allocs *cgoAllocMap
	refc2221b7f.features, cfeatures_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.features)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cfeatures_allocs)

	var cflags_allocs *cgoAllocMap
	refc2221b7f.flags, cflags_allocs = (C.uint32_t)(x.flags), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cflags_allocs)

	var clicenses_allocs *cgoAllocMap
	refc2221b7f.licenses, clicenses_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.licenses)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(clicenses_allocs)

	var cname_allocs *cgoAllocMap
	refc2221b7f.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cname_allocs)

	var cnode_cnt_allocs *cgoAllocMap
	refc2221b7f.node_cnt, cnode_cnt_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_cnt)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cnode_cnt_allocs)

	var cnode_list_allocs *cgoAllocMap
	refc2221b7f.node_list, cnode_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_list)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cnode_list_allocs)

	var cpartition_allocs *cgoAllocMap
	refc2221b7f.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partition)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cpartition_allocs)

	var cstart_time_allocs *cgoAllocMap
	refc2221b7f.start_time, cstart_time_allocs = (C.time_t)(x.start_time), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cstart_time_allocs)

	var cresv_watts_allocs *cgoAllocMap
	refc2221b7f.resv_watts, cresv_watts_allocs = (C.uint32_t)(x.resv_watts), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cresv_watts_allocs)

	var ctres_str_allocs *cgoAllocMap
	refc2221b7f.tres_str, ctres_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_str)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(ctres_str_allocs)

	var cusers_allocs *cgoAllocMap
	refc2221b7f.users, cusers_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.users)).Data)), cgoAllocsUnknown
	allocsc2221b7f.Borrow(cusers_allocs)

	x.refc2221b7f = refc2221b7f
	x.allocsc2221b7f = allocsc2221b7f
	return refc2221b7f, allocsc2221b7f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x resv_desc_msg_t) PassValue() (C.resv_desc_msg_t, *cgoAllocMap) {
	if x.refc2221b7f != nil {
		return *x.refc2221b7f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *resv_desc_msg_t) Deref() {
	if x.refc2221b7f == nil {
		return
	}
	hxf9f4fdc := (*sliceHeader)(unsafe.Pointer(&x.accounts))
	hxf9f4fdc.Data = uintptr(unsafe.Pointer(x.refc2221b7f.accounts))
	hxf9f4fdc.Cap = 0x7fffffff
	// hxf9f4fdc.Len = ?

	hxff43d70 := (*sliceHeader)(unsafe.Pointer(&x.burst_buffer))
	hxff43d70.Data = uintptr(unsafe.Pointer(x.refc2221b7f.burst_buffer))
	hxff43d70.Cap = 0x7fffffff
	// hxff43d70.Len = ?

	hxfdefba3 := (*sliceHeader)(unsafe.Pointer(&x.core_cnt))
	hxfdefba3.Data = uintptr(unsafe.Pointer(x.refc2221b7f.core_cnt))
	hxfdefba3.Cap = 0x7fffffff
	// hxfdefba3.Len = ?

	x.duration = (uint32_t)(x.refc2221b7f.duration)
	x.end_time = (time_t)(x.refc2221b7f.end_time)
	hxfb5b17c := (*sliceHeader)(unsafe.Pointer(&x.features))
	hxfb5b17c.Data = uintptr(unsafe.Pointer(x.refc2221b7f.features))
	hxfb5b17c.Cap = 0x7fffffff
	// hxfb5b17c.Len = ?

	x.flags = (uint32_t)(x.refc2221b7f.flags)
	hxf34d99d := (*sliceHeader)(unsafe.Pointer(&x.licenses))
	hxf34d99d.Data = uintptr(unsafe.Pointer(x.refc2221b7f.licenses))
	hxf34d99d.Cap = 0x7fffffff
	// hxf34d99d.Len = ?

	hxf0b92c4 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf0b92c4.Data = uintptr(unsafe.Pointer(x.refc2221b7f.name))
	hxf0b92c4.Cap = 0x7fffffff
	// hxf0b92c4.Len = ?

	hxf9d4ad3 := (*sliceHeader)(unsafe.Pointer(&x.node_cnt))
	hxf9d4ad3.Data = uintptr(unsafe.Pointer(x.refc2221b7f.node_cnt))
	hxf9d4ad3.Cap = 0x7fffffff
	// hxf9d4ad3.Len = ?

	hxf1f4227 := (*sliceHeader)(unsafe.Pointer(&x.node_list))
	hxf1f4227.Data = uintptr(unsafe.Pointer(x.refc2221b7f.node_list))
	hxf1f4227.Cap = 0x7fffffff
	// hxf1f4227.Len = ?

	hxf4ab867 := (*sliceHeader)(unsafe.Pointer(&x.partition))
	hxf4ab867.Data = uintptr(unsafe.Pointer(x.refc2221b7f.partition))
	hxf4ab867.Cap = 0x7fffffff
	// hxf4ab867.Len = ?

	x.start_time = (time_t)(x.refc2221b7f.start_time)
	x.resv_watts = (uint32_t)(x.refc2221b7f.resv_watts)
	hxfc015e5 := (*sliceHeader)(unsafe.Pointer(&x.tres_str))
	hxfc015e5.Data = uintptr(unsafe.Pointer(x.refc2221b7f.tres_str))
	hxfc015e5.Cap = 0x7fffffff
	// hxfc015e5.Len = ?

	hxfa55a25 := (*sliceHeader)(unsafe.Pointer(&x.users))
	hxfa55a25.Data = uintptr(unsafe.Pointer(x.refc2221b7f.users))
	hxfa55a25.Cap = 0x7fffffff
	// hxfa55a25.Len = ?

}

// allocReserve_response_msg_tMemory allocates memory for type C.reserve_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocReserve_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfReserve_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfReserve_response_msg_tValue = unsafe.Sizeof([1]C.reserve_response_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *reserve_response_msg_t) Ref() *C.reserve_response_msg_t {
	if x == nil {
		return nil
	}
	return x.ref18195ccc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *reserve_response_msg_t) Free() {
	if x != nil && x.allocs18195ccc != nil {
		x.allocs18195ccc.(*cgoAllocMap).Free()
		x.ref18195ccc = nil
	}
}

// Newreserve_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newreserve_response_msg_tRef(ref unsafe.Pointer) *reserve_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(reserve_response_msg_t)
	obj.ref18195ccc = (*C.reserve_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *reserve_response_msg_t) PassRef() (*C.reserve_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref18195ccc != nil {
		return x.ref18195ccc, nil
	}
	mem18195ccc := allocReserve_response_msg_tMemory(1)
	ref18195ccc := (*C.reserve_response_msg_t)(mem18195ccc)
	allocs18195ccc := new(cgoAllocMap)
	allocs18195ccc.Add(mem18195ccc)

	var cname_allocs *cgoAllocMap
	ref18195ccc.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs18195ccc.Borrow(cname_allocs)

	x.ref18195ccc = ref18195ccc
	x.allocs18195ccc = allocs18195ccc
	return ref18195ccc, allocs18195ccc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x reserve_response_msg_t) PassValue() (C.reserve_response_msg_t, *cgoAllocMap) {
	if x.ref18195ccc != nil {
		return *x.ref18195ccc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *reserve_response_msg_t) Deref() {
	if x.ref18195ccc == nil {
		return
	}
	hxfccb7ac := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxfccb7ac.Data = uintptr(unsafe.Pointer(x.ref18195ccc.name))
	hxfccb7ac.Cap = 0x7fffffff
	// hxfccb7ac.Len = ?

}

// allocReservation_name_msg_tMemory allocates memory for type C.reservation_name_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocReservation_name_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfReservation_name_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfReservation_name_msg_tValue = unsafe.Sizeof([1]C.reservation_name_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *reservation_name_msg_t) Ref() *C.reservation_name_msg_t {
	if x == nil {
		return nil
	}
	return x.ref5755f979
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *reservation_name_msg_t) Free() {
	if x != nil && x.allocs5755f979 != nil {
		x.allocs5755f979.(*cgoAllocMap).Free()
		x.ref5755f979 = nil
	}
}

// Newreservation_name_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newreservation_name_msg_tRef(ref unsafe.Pointer) *reservation_name_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(reservation_name_msg_t)
	obj.ref5755f979 = (*C.reservation_name_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *reservation_name_msg_t) PassRef() (*C.reservation_name_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5755f979 != nil {
		return x.ref5755f979, nil
	}
	mem5755f979 := allocReservation_name_msg_tMemory(1)
	ref5755f979 := (*C.reservation_name_msg_t)(mem5755f979)
	allocs5755f979 := new(cgoAllocMap)
	allocs5755f979.Add(mem5755f979)

	var cname_allocs *cgoAllocMap
	ref5755f979.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs5755f979.Borrow(cname_allocs)

	x.ref5755f979 = ref5755f979
	x.allocs5755f979 = allocs5755f979
	return ref5755f979, allocs5755f979

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x reservation_name_msg_t) PassValue() (C.reservation_name_msg_t, *cgoAllocMap) {
	if x.ref5755f979 != nil {
		return *x.ref5755f979, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *reservation_name_msg_t) Deref() {
	if x.ref5755f979 == nil {
		return
	}
	hxfdd589d := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxfdd589d.Data = uintptr(unsafe.Pointer(x.ref5755f979.name))
	hxfdd589d.Cap = 0x7fffffff
	// hxfdd589d.Len = ?

}

// allocSlurm_ctl_conf_tMemory allocates memory for type C.slurm_ctl_conf_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_ctl_conf_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_ctl_conf_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_ctl_conf_tValue = unsafe.Sizeof([1]C.slurm_ctl_conf_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_ctl_conf_t) Ref() *C.slurm_ctl_conf_t {
	if x == nil {
		return nil
	}
	return x.ref59adaf2e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_ctl_conf_t) Free() {
	if x != nil && x.allocs59adaf2e != nil {
		x.allocs59adaf2e.(*cgoAllocMap).Free()
		x.ref59adaf2e = nil
	}
}

// Newslurm_ctl_conf_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_ctl_conf_tRef(ref unsafe.Pointer) *slurm_ctl_conf_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_ctl_conf_t)
	obj.ref59adaf2e = (*C.slurm_ctl_conf_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_ctl_conf_t) PassRef() (*C.slurm_ctl_conf_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref59adaf2e != nil {
		return x.ref59adaf2e, nil
	}
	mem59adaf2e := allocSlurm_ctl_conf_tMemory(1)
	ref59adaf2e := (*C.slurm_ctl_conf_t)(mem59adaf2e)
	allocs59adaf2e := new(cgoAllocMap)
	allocs59adaf2e.Add(mem59adaf2e)

	var clast_update_allocs *cgoAllocMap
	ref59adaf2e.last_update, clast_update_allocs = (C.time_t)(x.last_update), cgoAllocsUnknown
	allocs59adaf2e.Borrow(clast_update_allocs)

	var caccounting_storage_tres_allocs *cgoAllocMap
	ref59adaf2e.accounting_storage_tres, caccounting_storage_tres_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.accounting_storage_tres)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(caccounting_storage_tres_allocs)

	var caccounting_storage_enforce_allocs *cgoAllocMap
	ref59adaf2e.accounting_storage_enforce, caccounting_storage_enforce_allocs = (C.uint16_t)(x.accounting_storage_enforce), cgoAllocsUnknown
	allocs59adaf2e.Borrow(caccounting_storage_enforce_allocs)

	var caccounting_storage_backup_host_allocs *cgoAllocMap
	ref59adaf2e.accounting_storage_backup_host, caccounting_storage_backup_host_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.accounting_storage_backup_host)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(caccounting_storage_backup_host_allocs)

	var caccounting_storage_host_allocs *cgoAllocMap
	ref59adaf2e.accounting_storage_host, caccounting_storage_host_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.accounting_storage_host)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(caccounting_storage_host_allocs)

	var caccounting_storage_loc_allocs *cgoAllocMap
	ref59adaf2e.accounting_storage_loc, caccounting_storage_loc_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.accounting_storage_loc)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(caccounting_storage_loc_allocs)

	var caccounting_storage_pass_allocs *cgoAllocMap
	ref59adaf2e.accounting_storage_pass, caccounting_storage_pass_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.accounting_storage_pass)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(caccounting_storage_pass_allocs)

	var caccounting_storage_port_allocs *cgoAllocMap
	ref59adaf2e.accounting_storage_port, caccounting_storage_port_allocs = (C.uint32_t)(x.accounting_storage_port), cgoAllocsUnknown
	allocs59adaf2e.Borrow(caccounting_storage_port_allocs)

	var caccounting_storage_type_allocs *cgoAllocMap
	ref59adaf2e.accounting_storage_type, caccounting_storage_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.accounting_storage_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(caccounting_storage_type_allocs)

	var caccounting_storage_user_allocs *cgoAllocMap
	ref59adaf2e.accounting_storage_user, caccounting_storage_user_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.accounting_storage_user)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(caccounting_storage_user_allocs)

	var cacctng_store_job_comment_allocs *cgoAllocMap
	ref59adaf2e.acctng_store_job_comment, cacctng_store_job_comment_allocs = (C.uint16_t)(x.acctng_store_job_comment), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cacctng_store_job_comment_allocs)

	var cacct_gather_conf_allocs *cgoAllocMap
	ref59adaf2e.acct_gather_conf, cacct_gather_conf_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.acct_gather_conf)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cacct_gather_conf_allocs)

	var cacct_gather_energy_type_allocs *cgoAllocMap
	ref59adaf2e.acct_gather_energy_type, cacct_gather_energy_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.acct_gather_energy_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cacct_gather_energy_type_allocs)

	var cacct_gather_profile_type_allocs *cgoAllocMap
	ref59adaf2e.acct_gather_profile_type, cacct_gather_profile_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.acct_gather_profile_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cacct_gather_profile_type_allocs)

	var cacct_gather_interconnect_type_allocs *cgoAllocMap
	ref59adaf2e.acct_gather_interconnect_type, cacct_gather_interconnect_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.acct_gather_interconnect_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cacct_gather_interconnect_type_allocs)

	var cacct_gather_filesystem_type_allocs *cgoAllocMap
	ref59adaf2e.acct_gather_filesystem_type, cacct_gather_filesystem_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.acct_gather_filesystem_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cacct_gather_filesystem_type_allocs)

	var cacct_gather_node_freq_allocs *cgoAllocMap
	ref59adaf2e.acct_gather_node_freq, cacct_gather_node_freq_allocs = (C.uint16_t)(x.acct_gather_node_freq), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cacct_gather_node_freq_allocs)

	var cauthinfo_allocs *cgoAllocMap
	ref59adaf2e.authinfo, cauthinfo_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.authinfo)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cauthinfo_allocs)

	var cauthtype_allocs *cgoAllocMap
	ref59adaf2e.authtype, cauthtype_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.authtype)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cauthtype_allocs)

	var cbackup_addr_allocs *cgoAllocMap
	ref59adaf2e.backup_addr, cbackup_addr_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.backup_addr)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cbackup_addr_allocs)

	var cbackup_controller_allocs *cgoAllocMap
	ref59adaf2e.backup_controller, cbackup_controller_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.backup_controller)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cbackup_controller_allocs)

	var cbatch_start_timeout_allocs *cgoAllocMap
	ref59adaf2e.batch_start_timeout, cbatch_start_timeout_allocs = (C.uint16_t)(x.batch_start_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cbatch_start_timeout_allocs)

	var cbb_type_allocs *cgoAllocMap
	ref59adaf2e.bb_type, cbb_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.bb_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cbb_type_allocs)

	var cboot_time_allocs *cgoAllocMap
	ref59adaf2e.boot_time, cboot_time_allocs = (C.time_t)(x.boot_time), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cboot_time_allocs)

	var ccheckpoint_type_allocs *cgoAllocMap
	ref59adaf2e.checkpoint_type, ccheckpoint_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.checkpoint_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ccheckpoint_type_allocs)

	var cchos_loc_allocs *cgoAllocMap
	ref59adaf2e.chos_loc, cchos_loc_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.chos_loc)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cchos_loc_allocs)

	var ccore_spec_plugin_allocs *cgoAllocMap
	ref59adaf2e.core_spec_plugin, ccore_spec_plugin_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.core_spec_plugin)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ccore_spec_plugin_allocs)

	var ccluster_name_allocs *cgoAllocMap
	ref59adaf2e.cluster_name, ccluster_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster_name)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ccluster_name_allocs)

	var ccomplete_wait_allocs *cgoAllocMap
	ref59adaf2e.complete_wait, ccomplete_wait_allocs = (C.uint16_t)(x.complete_wait), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ccomplete_wait_allocs)

	var ccontrol_addr_allocs *cgoAllocMap
	ref59adaf2e.control_addr, ccontrol_addr_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.control_addr)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ccontrol_addr_allocs)

	var ccontrol_machine_allocs *cgoAllocMap
	ref59adaf2e.control_machine, ccontrol_machine_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.control_machine)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ccontrol_machine_allocs)

	var ccpu_freq_def_allocs *cgoAllocMap
	ref59adaf2e.cpu_freq_def, ccpu_freq_def_allocs = (C.uint32_t)(x.cpu_freq_def), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ccpu_freq_def_allocs)

	var ccpu_freq_govs_allocs *cgoAllocMap
	ref59adaf2e.cpu_freq_govs, ccpu_freq_govs_allocs = (C.uint32_t)(x.cpu_freq_govs), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ccpu_freq_govs_allocs)

	var ccrypto_type_allocs *cgoAllocMap
	ref59adaf2e.crypto_type, ccrypto_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.crypto_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ccrypto_type_allocs)

	var cdebug_flags_allocs *cgoAllocMap
	ref59adaf2e.debug_flags, cdebug_flags_allocs = (C.uint64_t)(x.debug_flags), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cdebug_flags_allocs)

	var cdef_mem_per_cpu_allocs *cgoAllocMap
	ref59adaf2e.def_mem_per_cpu, cdef_mem_per_cpu_allocs = (C.uint64_t)(x.def_mem_per_cpu), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cdef_mem_per_cpu_allocs)

	var cdisable_root_jobs_allocs *cgoAllocMap
	ref59adaf2e.disable_root_jobs, cdisable_root_jobs_allocs = (C.uint16_t)(x.disable_root_jobs), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cdisable_root_jobs_allocs)

	var ceio_timeout_allocs *cgoAllocMap
	ref59adaf2e.eio_timeout, ceio_timeout_allocs = (C.uint16_t)(x.eio_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ceio_timeout_allocs)

	var cenforce_part_limits_allocs *cgoAllocMap
	ref59adaf2e.enforce_part_limits, cenforce_part_limits_allocs = (C.uint16_t)(x.enforce_part_limits), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cenforce_part_limits_allocs)

	var cepilog_allocs *cgoAllocMap
	ref59adaf2e.epilog, cepilog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.epilog)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cepilog_allocs)

	var cepilog_msg_time_allocs *cgoAllocMap
	ref59adaf2e.epilog_msg_time, cepilog_msg_time_allocs = (C.uint32_t)(x.epilog_msg_time), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cepilog_msg_time_allocs)

	var cepilog_slurmctld_allocs *cgoAllocMap
	ref59adaf2e.epilog_slurmctld, cepilog_slurmctld_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.epilog_slurmctld)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cepilog_slurmctld_allocs)

	var cext_sensors_type_allocs *cgoAllocMap
	ref59adaf2e.ext_sensors_type, cext_sensors_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.ext_sensors_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cext_sensors_type_allocs)

	var cext_sensors_freq_allocs *cgoAllocMap
	ref59adaf2e.ext_sensors_freq, cext_sensors_freq_allocs = (C.uint16_t)(x.ext_sensors_freq), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cext_sensors_freq_allocs)

	var cext_sensors_conf_allocs *cgoAllocMap
	ref59adaf2e.ext_sensors_conf, cext_sensors_conf_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.ext_sensors_conf)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cext_sensors_conf_allocs)

	var cfast_schedule_allocs *cgoAllocMap
	ref59adaf2e.fast_schedule, cfast_schedule_allocs = (C.uint16_t)(x.fast_schedule), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cfast_schedule_allocs)

	var cfed_params_allocs *cgoAllocMap
	ref59adaf2e.fed_params, cfed_params_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.fed_params)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cfed_params_allocs)

	var cfirst_job_id_allocs *cgoAllocMap
	ref59adaf2e.first_job_id, cfirst_job_id_allocs = (C.uint32_t)(x.first_job_id), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cfirst_job_id_allocs)

	var cfs_dampening_factor_allocs *cgoAllocMap
	ref59adaf2e.fs_dampening_factor, cfs_dampening_factor_allocs = (C.uint16_t)(x.fs_dampening_factor), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cfs_dampening_factor_allocs)

	var cget_env_timeout_allocs *cgoAllocMap
	ref59adaf2e.get_env_timeout, cget_env_timeout_allocs = (C.uint16_t)(x.get_env_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cget_env_timeout_allocs)

	var cgres_plugins_allocs *cgoAllocMap
	ref59adaf2e.gres_plugins, cgres_plugins_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.gres_plugins)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cgres_plugins_allocs)

	var cgroup_time_allocs *cgoAllocMap
	ref59adaf2e.group_time, cgroup_time_allocs = (C.uint16_t)(x.group_time), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cgroup_time_allocs)

	var cgroup_force_allocs *cgoAllocMap
	ref59adaf2e.group_force, cgroup_force_allocs = (C.uint16_t)(x.group_force), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cgroup_force_allocs)

	var chash_val_allocs *cgoAllocMap
	ref59adaf2e.hash_val, chash_val_allocs = (C.uint32_t)(x.hash_val), cgoAllocsUnknown
	allocs59adaf2e.Borrow(chash_val_allocs)

	var chealth_check_interval_allocs *cgoAllocMap
	ref59adaf2e.health_check_interval, chealth_check_interval_allocs = (C.uint16_t)(x.health_check_interval), cgoAllocsUnknown
	allocs59adaf2e.Borrow(chealth_check_interval_allocs)

	var chealth_check_node_state_allocs *cgoAllocMap
	ref59adaf2e.health_check_node_state, chealth_check_node_state_allocs = (C.uint16_t)(x.health_check_node_state), cgoAllocsUnknown
	allocs59adaf2e.Borrow(chealth_check_node_state_allocs)

	var chealth_check_program_allocs *cgoAllocMap
	ref59adaf2e.health_check_program, chealth_check_program_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.health_check_program)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(chealth_check_program_allocs)

	var cinactive_limit_allocs *cgoAllocMap
	ref59adaf2e.inactive_limit, cinactive_limit_allocs = (C.uint16_t)(x.inactive_limit), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cinactive_limit_allocs)

	var cjob_acct_gather_freq_allocs *cgoAllocMap
	ref59adaf2e.job_acct_gather_freq, cjob_acct_gather_freq_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_acct_gather_freq)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_acct_gather_freq_allocs)

	var cjob_acct_gather_type_allocs *cgoAllocMap
	ref59adaf2e.job_acct_gather_type, cjob_acct_gather_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_acct_gather_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_acct_gather_type_allocs)

	var cjob_acct_gather_params_allocs *cgoAllocMap
	ref59adaf2e.job_acct_gather_params, cjob_acct_gather_params_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_acct_gather_params)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_acct_gather_params_allocs)

	var cjob_ckpt_dir_allocs *cgoAllocMap
	ref59adaf2e.job_ckpt_dir, cjob_ckpt_dir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_ckpt_dir)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_ckpt_dir_allocs)

	var cjob_comp_host_allocs *cgoAllocMap
	ref59adaf2e.job_comp_host, cjob_comp_host_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_comp_host)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_comp_host_allocs)

	var cjob_comp_loc_allocs *cgoAllocMap
	ref59adaf2e.job_comp_loc, cjob_comp_loc_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_comp_loc)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_comp_loc_allocs)

	var cjob_comp_pass_allocs *cgoAllocMap
	ref59adaf2e.job_comp_pass, cjob_comp_pass_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_comp_pass)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_comp_pass_allocs)

	var cjob_comp_port_allocs *cgoAllocMap
	ref59adaf2e.job_comp_port, cjob_comp_port_allocs = (C.uint32_t)(x.job_comp_port), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_comp_port_allocs)

	var cjob_comp_type_allocs *cgoAllocMap
	ref59adaf2e.job_comp_type, cjob_comp_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_comp_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_comp_type_allocs)

	var cjob_comp_user_allocs *cgoAllocMap
	ref59adaf2e.job_comp_user, cjob_comp_user_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_comp_user)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_comp_user_allocs)

	var cjob_container_plugin_allocs *cgoAllocMap
	ref59adaf2e.job_container_plugin, cjob_container_plugin_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_container_plugin)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_container_plugin_allocs)

	var cjob_credential_private_key_allocs *cgoAllocMap
	ref59adaf2e.job_credential_private_key, cjob_credential_private_key_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_credential_private_key)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_credential_private_key_allocs)

	var cjob_credential_public_certificate_allocs *cgoAllocMap
	ref59adaf2e.job_credential_public_certificate, cjob_credential_public_certificate_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_credential_public_certificate)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_credential_public_certificate_allocs)

	var cjob_file_append_allocs *cgoAllocMap
	ref59adaf2e.job_file_append, cjob_file_append_allocs = (C.uint16_t)(x.job_file_append), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_file_append_allocs)

	var cjob_requeue_allocs *cgoAllocMap
	ref59adaf2e.job_requeue, cjob_requeue_allocs = (C.uint16_t)(x.job_requeue), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_requeue_allocs)

	var cjob_submit_plugins_allocs *cgoAllocMap
	ref59adaf2e.job_submit_plugins, cjob_submit_plugins_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_submit_plugins)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cjob_submit_plugins_allocs)

	var ckeep_alive_time_allocs *cgoAllocMap
	ref59adaf2e.keep_alive_time, ckeep_alive_time_allocs = (C.uint16_t)(x.keep_alive_time), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ckeep_alive_time_allocs)

	var ckill_on_bad_exit_allocs *cgoAllocMap
	ref59adaf2e.kill_on_bad_exit, ckill_on_bad_exit_allocs = (C.uint16_t)(x.kill_on_bad_exit), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ckill_on_bad_exit_allocs)

	var ckill_wait_allocs *cgoAllocMap
	ref59adaf2e.kill_wait, ckill_wait_allocs = (C.uint16_t)(x.kill_wait), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ckill_wait_allocs)

	var claunch_params_allocs *cgoAllocMap
	ref59adaf2e.launch_params, claunch_params_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.launch_params)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(claunch_params_allocs)

	var claunch_type_allocs *cgoAllocMap
	ref59adaf2e.launch_type, claunch_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.launch_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(claunch_type_allocs)

	var clayouts_allocs *cgoAllocMap
	ref59adaf2e.layouts, clayouts_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.layouts)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(clayouts_allocs)

	var clicenses_allocs *cgoAllocMap
	ref59adaf2e.licenses, clicenses_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.licenses)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(clicenses_allocs)

	var clicenses_used_allocs *cgoAllocMap
	ref59adaf2e.licenses_used, clicenses_used_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.licenses_used)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(clicenses_used_allocs)

	var clog_fmt_allocs *cgoAllocMap
	ref59adaf2e.log_fmt, clog_fmt_allocs = (C.uint16_t)(x.log_fmt), cgoAllocsUnknown
	allocs59adaf2e.Borrow(clog_fmt_allocs)

	var cmail_domain_allocs *cgoAllocMap
	ref59adaf2e.mail_domain, cmail_domain_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mail_domain)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmail_domain_allocs)

	var cmail_prog_allocs *cgoAllocMap
	ref59adaf2e.mail_prog, cmail_prog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mail_prog)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmail_prog_allocs)

	var cmax_array_sz_allocs *cgoAllocMap
	ref59adaf2e.max_array_sz, cmax_array_sz_allocs = (C.uint32_t)(x.max_array_sz), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmax_array_sz_allocs)

	var cmax_job_cnt_allocs *cgoAllocMap
	ref59adaf2e.max_job_cnt, cmax_job_cnt_allocs = (C.uint32_t)(x.max_job_cnt), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmax_job_cnt_allocs)

	var cmax_job_id_allocs *cgoAllocMap
	ref59adaf2e.max_job_id, cmax_job_id_allocs = (C.uint32_t)(x.max_job_id), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmax_job_id_allocs)

	var cmax_mem_per_cpu_allocs *cgoAllocMap
	ref59adaf2e.max_mem_per_cpu, cmax_mem_per_cpu_allocs = (C.uint64_t)(x.max_mem_per_cpu), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmax_mem_per_cpu_allocs)

	var cmax_step_cnt_allocs *cgoAllocMap
	ref59adaf2e.max_step_cnt, cmax_step_cnt_allocs = (C.uint32_t)(x.max_step_cnt), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmax_step_cnt_allocs)

	var cmax_tasks_per_node_allocs *cgoAllocMap
	ref59adaf2e.max_tasks_per_node, cmax_tasks_per_node_allocs = (C.uint16_t)(x.max_tasks_per_node), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmax_tasks_per_node_allocs)

	var cmcs_plugin_allocs *cgoAllocMap
	ref59adaf2e.mcs_plugin, cmcs_plugin_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mcs_plugin)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmcs_plugin_allocs)

	var cmcs_plugin_params_allocs *cgoAllocMap
	ref59adaf2e.mcs_plugin_params, cmcs_plugin_params_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mcs_plugin_params)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmcs_plugin_params_allocs)

	var cmem_limit_enforce_allocs *cgoAllocMap
	ref59adaf2e.mem_limit_enforce, cmem_limit_enforce_allocs = (C.uint16_t)(x.mem_limit_enforce), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmem_limit_enforce_allocs)

	var cmin_job_age_allocs *cgoAllocMap
	ref59adaf2e.min_job_age, cmin_job_age_allocs = (C.uint32_t)(x.min_job_age), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmin_job_age_allocs)

	var cmpi_default_allocs *cgoAllocMap
	ref59adaf2e.mpi_default, cmpi_default_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mpi_default)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmpi_default_allocs)

	var cmpi_params_allocs *cgoAllocMap
	ref59adaf2e.mpi_params, cmpi_params_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.mpi_params)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmpi_params_allocs)

	var cmsg_aggr_params_allocs *cgoAllocMap
	ref59adaf2e.msg_aggr_params, cmsg_aggr_params_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.msg_aggr_params)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmsg_aggr_params_allocs)

	var cmsg_timeout_allocs *cgoAllocMap
	ref59adaf2e.msg_timeout, cmsg_timeout_allocs = (C.uint16_t)(x.msg_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cmsg_timeout_allocs)

	var ctcp_timeout_allocs *cgoAllocMap
	ref59adaf2e.tcp_timeout, ctcp_timeout_allocs = (C.uint16_t)(x.tcp_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ctcp_timeout_allocs)

	var cnext_job_id_allocs *cgoAllocMap
	ref59adaf2e.next_job_id, cnext_job_id_allocs = (C.uint32_t)(x.next_job_id), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cnext_job_id_allocs)

	var cnode_features_plugins_allocs *cgoAllocMap
	ref59adaf2e.node_features_plugins, cnode_features_plugins_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_features_plugins)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cnode_features_plugins_allocs)

	var cnode_prefix_allocs *cgoAllocMap
	ref59adaf2e.node_prefix, cnode_prefix_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_prefix)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cnode_prefix_allocs)

	var cover_time_limit_allocs *cgoAllocMap
	ref59adaf2e.over_time_limit, cover_time_limit_allocs = (C.uint16_t)(x.over_time_limit), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cover_time_limit_allocs)

	var cplugindir_allocs *cgoAllocMap
	ref59adaf2e.plugindir, cplugindir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.plugindir)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cplugindir_allocs)

	var cplugstack_allocs *cgoAllocMap
	ref59adaf2e.plugstack, cplugstack_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.plugstack)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cplugstack_allocs)

	var cpower_parameters_allocs *cgoAllocMap
	ref59adaf2e.power_parameters, cpower_parameters_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.power_parameters)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpower_parameters_allocs)

	var cpower_plugin_allocs *cgoAllocMap
	ref59adaf2e.power_plugin, cpower_plugin_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.power_plugin)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpower_plugin_allocs)

	var cpreempt_mode_allocs *cgoAllocMap
	ref59adaf2e.preempt_mode, cpreempt_mode_allocs = (C.uint16_t)(x.preempt_mode), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpreempt_mode_allocs)

	var cpreempt_type_allocs *cgoAllocMap
	ref59adaf2e.preempt_type, cpreempt_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.preempt_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpreempt_type_allocs)

	var cpriority_decay_hl_allocs *cgoAllocMap
	ref59adaf2e.priority_decay_hl, cpriority_decay_hl_allocs = (C.uint32_t)(x.priority_decay_hl), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_decay_hl_allocs)

	var cpriority_calc_period_allocs *cgoAllocMap
	ref59adaf2e.priority_calc_period, cpriority_calc_period_allocs = (C.uint32_t)(x.priority_calc_period), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_calc_period_allocs)

	var cpriority_favor_small_allocs *cgoAllocMap
	ref59adaf2e.priority_favor_small, cpriority_favor_small_allocs = (C.uint16_t)(x.priority_favor_small), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_favor_small_allocs)

	var cpriority_flags_allocs *cgoAllocMap
	ref59adaf2e.priority_flags, cpriority_flags_allocs = (C.uint16_t)(x.priority_flags), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_flags_allocs)

	var cpriority_max_age_allocs *cgoAllocMap
	ref59adaf2e.priority_max_age, cpriority_max_age_allocs = (C.uint32_t)(x.priority_max_age), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_max_age_allocs)

	var cpriority_params_allocs *cgoAllocMap
	ref59adaf2e.priority_params, cpriority_params_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.priority_params)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_params_allocs)

	var cpriority_reset_period_allocs *cgoAllocMap
	ref59adaf2e.priority_reset_period, cpriority_reset_period_allocs = (C.uint16_t)(x.priority_reset_period), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_reset_period_allocs)

	var cpriority_type_allocs *cgoAllocMap
	ref59adaf2e.priority_type, cpriority_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.priority_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_type_allocs)

	var cpriority_weight_age_allocs *cgoAllocMap
	ref59adaf2e.priority_weight_age, cpriority_weight_age_allocs = (C.uint32_t)(x.priority_weight_age), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_weight_age_allocs)

	var cpriority_weight_fs_allocs *cgoAllocMap
	ref59adaf2e.priority_weight_fs, cpriority_weight_fs_allocs = (C.uint32_t)(x.priority_weight_fs), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_weight_fs_allocs)

	var cpriority_weight_js_allocs *cgoAllocMap
	ref59adaf2e.priority_weight_js, cpriority_weight_js_allocs = (C.uint32_t)(x.priority_weight_js), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_weight_js_allocs)

	var cpriority_weight_part_allocs *cgoAllocMap
	ref59adaf2e.priority_weight_part, cpriority_weight_part_allocs = (C.uint32_t)(x.priority_weight_part), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_weight_part_allocs)

	var cpriority_weight_qos_allocs *cgoAllocMap
	ref59adaf2e.priority_weight_qos, cpriority_weight_qos_allocs = (C.uint32_t)(x.priority_weight_qos), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_weight_qos_allocs)

	var cpriority_weight_tres_allocs *cgoAllocMap
	ref59adaf2e.priority_weight_tres, cpriority_weight_tres_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.priority_weight_tres)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpriority_weight_tres_allocs)

	var cprivate_data_allocs *cgoAllocMap
	ref59adaf2e.private_data, cprivate_data_allocs = (C.uint16_t)(x.private_data), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cprivate_data_allocs)

	var cproctrack_type_allocs *cgoAllocMap
	ref59adaf2e.proctrack_type, cproctrack_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.proctrack_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cproctrack_type_allocs)

	var cprolog_allocs *cgoAllocMap
	ref59adaf2e.prolog, cprolog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.prolog)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cprolog_allocs)

	var cprolog_epilog_timeout_allocs *cgoAllocMap
	ref59adaf2e.prolog_epilog_timeout, cprolog_epilog_timeout_allocs = (C.uint16_t)(x.prolog_epilog_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cprolog_epilog_timeout_allocs)

	var cprolog_slurmctld_allocs *cgoAllocMap
	ref59adaf2e.prolog_slurmctld, cprolog_slurmctld_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.prolog_slurmctld)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cprolog_slurmctld_allocs)

	var cpropagate_prio_process_allocs *cgoAllocMap
	ref59adaf2e.propagate_prio_process, cpropagate_prio_process_allocs = (C.uint16_t)(x.propagate_prio_process), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpropagate_prio_process_allocs)

	var cprolog_flags_allocs *cgoAllocMap
	ref59adaf2e.prolog_flags, cprolog_flags_allocs = (C.uint16_t)(x.prolog_flags), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cprolog_flags_allocs)

	var cpropagate_rlimits_allocs *cgoAllocMap
	ref59adaf2e.propagate_rlimits, cpropagate_rlimits_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.propagate_rlimits)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpropagate_rlimits_allocs)

	var cpropagate_rlimits_except_allocs *cgoAllocMap
	ref59adaf2e.propagate_rlimits_except, cpropagate_rlimits_except_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.propagate_rlimits_except)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cpropagate_rlimits_except_allocs)

	var creboot_program_allocs *cgoAllocMap
	ref59adaf2e.reboot_program, creboot_program_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.reboot_program)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(creboot_program_allocs)

	var creconfig_flags_allocs *cgoAllocMap
	ref59adaf2e.reconfig_flags, creconfig_flags_allocs = (C.uint16_t)(x.reconfig_flags), cgoAllocsUnknown
	allocs59adaf2e.Borrow(creconfig_flags_allocs)

	var crequeue_exit_allocs *cgoAllocMap
	ref59adaf2e.requeue_exit, crequeue_exit_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.requeue_exit)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(crequeue_exit_allocs)

	var crequeue_exit_hold_allocs *cgoAllocMap
	ref59adaf2e.requeue_exit_hold, crequeue_exit_hold_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.requeue_exit_hold)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(crequeue_exit_hold_allocs)

	var cresume_program_allocs *cgoAllocMap
	ref59adaf2e.resume_program, cresume_program_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.resume_program)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cresume_program_allocs)

	var cresume_rate_allocs *cgoAllocMap
	ref59adaf2e.resume_rate, cresume_rate_allocs = (C.uint16_t)(x.resume_rate), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cresume_rate_allocs)

	var cresume_timeout_allocs *cgoAllocMap
	ref59adaf2e.resume_timeout, cresume_timeout_allocs = (C.uint16_t)(x.resume_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cresume_timeout_allocs)

	var cresv_epilog_allocs *cgoAllocMap
	ref59adaf2e.resv_epilog, cresv_epilog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.resv_epilog)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cresv_epilog_allocs)

	var cresv_over_run_allocs *cgoAllocMap
	ref59adaf2e.resv_over_run, cresv_over_run_allocs = (C.uint16_t)(x.resv_over_run), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cresv_over_run_allocs)

	var cresv_prolog_allocs *cgoAllocMap
	ref59adaf2e.resv_prolog, cresv_prolog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.resv_prolog)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cresv_prolog_allocs)

	var cret2service_allocs *cgoAllocMap
	ref59adaf2e.ret2service, cret2service_allocs = (C.uint16_t)(x.ret2service), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cret2service_allocs)

	var croute_plugin_allocs *cgoAllocMap
	ref59adaf2e.route_plugin, croute_plugin_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.route_plugin)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(croute_plugin_allocs)

	var csalloc_default_command_allocs *cgoAllocMap
	ref59adaf2e.salloc_default_command, csalloc_default_command_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.salloc_default_command)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csalloc_default_command_allocs)

	var csbcast_parameters_allocs *cgoAllocMap
	ref59adaf2e.sbcast_parameters, csbcast_parameters_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.sbcast_parameters)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csbcast_parameters_allocs)

	var csched_logfile_allocs *cgoAllocMap
	ref59adaf2e.sched_logfile, csched_logfile_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.sched_logfile)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csched_logfile_allocs)

	var csched_log_level_allocs *cgoAllocMap
	ref59adaf2e.sched_log_level, csched_log_level_allocs = (C.uint16_t)(x.sched_log_level), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csched_log_level_allocs)

	var csched_params_allocs *cgoAllocMap
	ref59adaf2e.sched_params, csched_params_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.sched_params)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csched_params_allocs)

	var csched_time_slice_allocs *cgoAllocMap
	ref59adaf2e.sched_time_slice, csched_time_slice_allocs = (C.uint16_t)(x.sched_time_slice), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csched_time_slice_allocs)

	var cschedtype_allocs *cgoAllocMap
	ref59adaf2e.schedtype, cschedtype_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.schedtype)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cschedtype_allocs)

	var cselect_type_allocs *cgoAllocMap
	ref59adaf2e.select_type, cselect_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.select_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cselect_type_allocs)

	var cselect_conf_key_pairs_allocs *cgoAllocMap
	ref59adaf2e.select_conf_key_pairs, cselect_conf_key_pairs_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.select_conf_key_pairs)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cselect_conf_key_pairs_allocs)

	var cselect_type_param_allocs *cgoAllocMap
	ref59adaf2e.select_type_param, cselect_type_param_allocs = (C.uint16_t)(x.select_type_param), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cselect_type_param_allocs)

	var cslurm_conf_allocs *cgoAllocMap
	ref59adaf2e.slurm_conf, cslurm_conf_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.slurm_conf)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurm_conf_allocs)

	var cslurm_user_id_allocs *cgoAllocMap
	ref59adaf2e.slurm_user_id, cslurm_user_id_allocs = (C.uint32_t)(x.slurm_user_id), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurm_user_id_allocs)

	var cslurm_user_name_allocs *cgoAllocMap
	ref59adaf2e.slurm_user_name, cslurm_user_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.slurm_user_name)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurm_user_name_allocs)

	var cslurmd_user_id_allocs *cgoAllocMap
	ref59adaf2e.slurmd_user_id, cslurmd_user_id_allocs = (C.uint32_t)(x.slurmd_user_id), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmd_user_id_allocs)

	var cslurmd_user_name_allocs *cgoAllocMap
	ref59adaf2e.slurmd_user_name, cslurmd_user_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.slurmd_user_name)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmd_user_name_allocs)

	var cslurmctld_debug_allocs *cgoAllocMap
	ref59adaf2e.slurmctld_debug, cslurmctld_debug_allocs = (C.uint16_t)(x.slurmctld_debug), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmctld_debug_allocs)

	var cslurmctld_logfile_allocs *cgoAllocMap
	ref59adaf2e.slurmctld_logfile, cslurmctld_logfile_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.slurmctld_logfile)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmctld_logfile_allocs)

	var cslurmctld_pidfile_allocs *cgoAllocMap
	ref59adaf2e.slurmctld_pidfile, cslurmctld_pidfile_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.slurmctld_pidfile)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmctld_pidfile_allocs)

	var cslurmctld_plugstack_allocs *cgoAllocMap
	ref59adaf2e.slurmctld_plugstack, cslurmctld_plugstack_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.slurmctld_plugstack)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmctld_plugstack_allocs)

	var cslurmctld_port_allocs *cgoAllocMap
	ref59adaf2e.slurmctld_port, cslurmctld_port_allocs = (C.uint32_t)(x.slurmctld_port), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmctld_port_allocs)

	var cslurmctld_port_count_allocs *cgoAllocMap
	ref59adaf2e.slurmctld_port_count, cslurmctld_port_count_allocs = (C.uint16_t)(x.slurmctld_port_count), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmctld_port_count_allocs)

	var cslurmctld_syslog_debug_allocs *cgoAllocMap
	ref59adaf2e.slurmctld_syslog_debug, cslurmctld_syslog_debug_allocs = (C.uint16_t)(x.slurmctld_syslog_debug), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmctld_syslog_debug_allocs)

	var cslurmctld_timeout_allocs *cgoAllocMap
	ref59adaf2e.slurmctld_timeout, cslurmctld_timeout_allocs = (C.uint16_t)(x.slurmctld_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmctld_timeout_allocs)

	var cslurmd_debug_allocs *cgoAllocMap
	ref59adaf2e.slurmd_debug, cslurmd_debug_allocs = (C.uint16_t)(x.slurmd_debug), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmd_debug_allocs)

	var cslurmd_logfile_allocs *cgoAllocMap
	ref59adaf2e.slurmd_logfile, cslurmd_logfile_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.slurmd_logfile)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmd_logfile_allocs)

	var cslurmd_pidfile_allocs *cgoAllocMap
	ref59adaf2e.slurmd_pidfile, cslurmd_pidfile_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.slurmd_pidfile)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmd_pidfile_allocs)

	var cslurmd_port_allocs *cgoAllocMap
	ref59adaf2e.slurmd_port, cslurmd_port_allocs = (C.uint32_t)(x.slurmd_port), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmd_port_allocs)

	var cslurmd_spooldir_allocs *cgoAllocMap
	ref59adaf2e.slurmd_spooldir, cslurmd_spooldir_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.slurmd_spooldir)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmd_spooldir_allocs)

	var cslurmd_syslog_debug_allocs *cgoAllocMap
	ref59adaf2e.slurmd_syslog_debug, cslurmd_syslog_debug_allocs = (C.uint16_t)(x.slurmd_syslog_debug), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmd_syslog_debug_allocs)

	var cslurmd_timeout_allocs *cgoAllocMap
	ref59adaf2e.slurmd_timeout, cslurmd_timeout_allocs = (C.uint16_t)(x.slurmd_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cslurmd_timeout_allocs)

	var csrun_epilog_allocs *cgoAllocMap
	ref59adaf2e.srun_epilog, csrun_epilog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.srun_epilog)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csrun_epilog_allocs)

	var csrun_port_range_allocs *cgoAllocMap
	ref59adaf2e.srun_port_range, csrun_port_range_allocs = (*C.uint16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.srun_port_range)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csrun_port_range_allocs)

	var csrun_prolog_allocs *cgoAllocMap
	ref59adaf2e.srun_prolog, csrun_prolog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.srun_prolog)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csrun_prolog_allocs)

	var cstate_save_location_allocs *cgoAllocMap
	ref59adaf2e.state_save_location, cstate_save_location_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.state_save_location)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cstate_save_location_allocs)

	var csuspend_exc_nodes_allocs *cgoAllocMap
	ref59adaf2e.suspend_exc_nodes, csuspend_exc_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.suspend_exc_nodes)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csuspend_exc_nodes_allocs)

	var csuspend_exc_parts_allocs *cgoAllocMap
	ref59adaf2e.suspend_exc_parts, csuspend_exc_parts_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.suspend_exc_parts)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csuspend_exc_parts_allocs)

	var csuspend_program_allocs *cgoAllocMap
	ref59adaf2e.suspend_program, csuspend_program_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.suspend_program)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csuspend_program_allocs)

	var csuspend_rate_allocs *cgoAllocMap
	ref59adaf2e.suspend_rate, csuspend_rate_allocs = (C.uint16_t)(x.suspend_rate), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csuspend_rate_allocs)

	var csuspend_time_allocs *cgoAllocMap
	ref59adaf2e.suspend_time, csuspend_time_allocs = (C.uint32_t)(x.suspend_time), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csuspend_time_allocs)

	var csuspend_timeout_allocs *cgoAllocMap
	ref59adaf2e.suspend_timeout, csuspend_timeout_allocs = (C.uint16_t)(x.suspend_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(csuspend_timeout_allocs)

	var cswitch_type_allocs *cgoAllocMap
	ref59adaf2e.switch_type, cswitch_type_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.switch_type)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cswitch_type_allocs)

	var ctask_epilog_allocs *cgoAllocMap
	ref59adaf2e.task_epilog, ctask_epilog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.task_epilog)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ctask_epilog_allocs)

	var ctask_plugin_allocs *cgoAllocMap
	ref59adaf2e.task_plugin, ctask_plugin_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.task_plugin)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ctask_plugin_allocs)

	var ctask_plugin_param_allocs *cgoAllocMap
	ref59adaf2e.task_plugin_param, ctask_plugin_param_allocs = (C.uint32_t)(x.task_plugin_param), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ctask_plugin_param_allocs)

	var ctask_prolog_allocs *cgoAllocMap
	ref59adaf2e.task_prolog, ctask_prolog_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.task_prolog)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ctask_prolog_allocs)

	var ctmp_fs_allocs *cgoAllocMap
	ref59adaf2e.tmp_fs, ctmp_fs_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tmp_fs)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ctmp_fs_allocs)

	var ctopology_param_allocs *cgoAllocMap
	ref59adaf2e.topology_param, ctopology_param_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.topology_param)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ctopology_param_allocs)

	var ctopology_plugin_allocs *cgoAllocMap
	ref59adaf2e.topology_plugin, ctopology_plugin_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.topology_plugin)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ctopology_plugin_allocs)

	var ctrack_wckey_allocs *cgoAllocMap
	ref59adaf2e.track_wckey, ctrack_wckey_allocs = (C.uint16_t)(x.track_wckey), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ctrack_wckey_allocs)

	var ctree_width_allocs *cgoAllocMap
	ref59adaf2e.tree_width, ctree_width_allocs = (C.uint16_t)(x.tree_width), cgoAllocsUnknown
	allocs59adaf2e.Borrow(ctree_width_allocs)

	var cunkillable_program_allocs *cgoAllocMap
	ref59adaf2e.unkillable_program, cunkillable_program_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.unkillable_program)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cunkillable_program_allocs)

	var cunkillable_timeout_allocs *cgoAllocMap
	ref59adaf2e.unkillable_timeout, cunkillable_timeout_allocs = (C.uint16_t)(x.unkillable_timeout), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cunkillable_timeout_allocs)

	var cuse_pam_allocs *cgoAllocMap
	ref59adaf2e.use_pam, cuse_pam_allocs = (C.uint16_t)(x.use_pam), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cuse_pam_allocs)

	var cuse_spec_resources_allocs *cgoAllocMap
	ref59adaf2e.use_spec_resources, cuse_spec_resources_allocs = (C.uint16_t)(x.use_spec_resources), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cuse_spec_resources_allocs)

	var cversion_allocs *cgoAllocMap
	ref59adaf2e.version, cversion_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.version)).Data)), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cversion_allocs)

	var cvsize_factor_allocs *cgoAllocMap
	ref59adaf2e.vsize_factor, cvsize_factor_allocs = (C.uint16_t)(x.vsize_factor), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cvsize_factor_allocs)

	var cwait_time_allocs *cgoAllocMap
	ref59adaf2e.wait_time, cwait_time_allocs = (C.uint16_t)(x.wait_time), cgoAllocsUnknown
	allocs59adaf2e.Borrow(cwait_time_allocs)

	x.ref59adaf2e = ref59adaf2e
	x.allocs59adaf2e = allocs59adaf2e
	return ref59adaf2e, allocs59adaf2e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_ctl_conf_t) PassValue() (C.slurm_ctl_conf_t, *cgoAllocMap) {
	if x.ref59adaf2e != nil {
		return *x.ref59adaf2e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_ctl_conf_t) Deref() {
	if x.ref59adaf2e == nil {
		return
	}
	x.last_update = (time_t)(x.ref59adaf2e.last_update)
	hxfb26947 := (*sliceHeader)(unsafe.Pointer(&x.accounting_storage_tres))
	hxfb26947.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.accounting_storage_tres))
	hxfb26947.Cap = 0x7fffffff
	// hxfb26947.Len = ?

	x.accounting_storage_enforce = (uint16_t)(x.ref59adaf2e.accounting_storage_enforce)
	hxf2c419e := (*sliceHeader)(unsafe.Pointer(&x.accounting_storage_backup_host))
	hxf2c419e.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.accounting_storage_backup_host))
	hxf2c419e.Cap = 0x7fffffff
	// hxf2c419e.Len = ?

	hxf2b936c := (*sliceHeader)(unsafe.Pointer(&x.accounting_storage_host))
	hxf2b936c.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.accounting_storage_host))
	hxf2b936c.Cap = 0x7fffffff
	// hxf2b936c.Len = ?

	hxfdeadf1 := (*sliceHeader)(unsafe.Pointer(&x.accounting_storage_loc))
	hxfdeadf1.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.accounting_storage_loc))
	hxfdeadf1.Cap = 0x7fffffff
	// hxfdeadf1.Len = ?

	hxf4c0f12 := (*sliceHeader)(unsafe.Pointer(&x.accounting_storage_pass))
	hxf4c0f12.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.accounting_storage_pass))
	hxf4c0f12.Cap = 0x7fffffff
	// hxf4c0f12.Len = ?

	x.accounting_storage_port = (uint32_t)(x.ref59adaf2e.accounting_storage_port)
	hxf96e840 := (*sliceHeader)(unsafe.Pointer(&x.accounting_storage_type))
	hxf96e840.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.accounting_storage_type))
	hxf96e840.Cap = 0x7fffffff
	// hxf96e840.Len = ?

	hxfca0752 := (*sliceHeader)(unsafe.Pointer(&x.accounting_storage_user))
	hxfca0752.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.accounting_storage_user))
	hxfca0752.Cap = 0x7fffffff
	// hxfca0752.Len = ?

	x.acctng_store_job_comment = (uint16_t)(x.ref59adaf2e.acctng_store_job_comment)
	x.acct_gather_conf = (unsafe.Pointer)(unsafe.Pointer(x.ref59adaf2e.acct_gather_conf))
	hxf81f8fc := (*sliceHeader)(unsafe.Pointer(&x.acct_gather_energy_type))
	hxf81f8fc.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.acct_gather_energy_type))
	hxf81f8fc.Cap = 0x7fffffff
	// hxf81f8fc.Len = ?

	hxf87995f := (*sliceHeader)(unsafe.Pointer(&x.acct_gather_profile_type))
	hxf87995f.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.acct_gather_profile_type))
	hxf87995f.Cap = 0x7fffffff
	// hxf87995f.Len = ?

	hxf1c56f2 := (*sliceHeader)(unsafe.Pointer(&x.acct_gather_interconnect_type))
	hxf1c56f2.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.acct_gather_interconnect_type))
	hxf1c56f2.Cap = 0x7fffffff
	// hxf1c56f2.Len = ?

	hxfd282b8 := (*sliceHeader)(unsafe.Pointer(&x.acct_gather_filesystem_type))
	hxfd282b8.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.acct_gather_filesystem_type))
	hxfd282b8.Cap = 0x7fffffff
	// hxfd282b8.Len = ?

	x.acct_gather_node_freq = (uint16_t)(x.ref59adaf2e.acct_gather_node_freq)
	hxf5aa0aa := (*sliceHeader)(unsafe.Pointer(&x.authinfo))
	hxf5aa0aa.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.authinfo))
	hxf5aa0aa.Cap = 0x7fffffff
	// hxf5aa0aa.Len = ?

	hxfe70a7f := (*sliceHeader)(unsafe.Pointer(&x.authtype))
	hxfe70a7f.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.authtype))
	hxfe70a7f.Cap = 0x7fffffff
	// hxfe70a7f.Len = ?

	hxf4796ba := (*sliceHeader)(unsafe.Pointer(&x.backup_addr))
	hxf4796ba.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.backup_addr))
	hxf4796ba.Cap = 0x7fffffff
	// hxf4796ba.Len = ?

	hxfbafbb2 := (*sliceHeader)(unsafe.Pointer(&x.backup_controller))
	hxfbafbb2.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.backup_controller))
	hxfbafbb2.Cap = 0x7fffffff
	// hxfbafbb2.Len = ?

	x.batch_start_timeout = (uint16_t)(x.ref59adaf2e.batch_start_timeout)
	hxfb6c15f := (*sliceHeader)(unsafe.Pointer(&x.bb_type))
	hxfb6c15f.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.bb_type))
	hxfb6c15f.Cap = 0x7fffffff
	// hxfb6c15f.Len = ?

	x.boot_time = (time_t)(x.ref59adaf2e.boot_time)
	hxfc04a06 := (*sliceHeader)(unsafe.Pointer(&x.checkpoint_type))
	hxfc04a06.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.checkpoint_type))
	hxfc04a06.Cap = 0x7fffffff
	// hxfc04a06.Len = ?

	hxfbb751b := (*sliceHeader)(unsafe.Pointer(&x.chos_loc))
	hxfbb751b.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.chos_loc))
	hxfbb751b.Cap = 0x7fffffff
	// hxfbb751b.Len = ?

	hxf023cba := (*sliceHeader)(unsafe.Pointer(&x.core_spec_plugin))
	hxf023cba.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.core_spec_plugin))
	hxf023cba.Cap = 0x7fffffff
	// hxf023cba.Len = ?

	hxfe4fa36 := (*sliceHeader)(unsafe.Pointer(&x.cluster_name))
	hxfe4fa36.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.cluster_name))
	hxfe4fa36.Cap = 0x7fffffff
	// hxfe4fa36.Len = ?

	x.complete_wait = (uint16_t)(x.ref59adaf2e.complete_wait)
	hxfeadef6 := (*sliceHeader)(unsafe.Pointer(&x.control_addr))
	hxfeadef6.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.control_addr))
	hxfeadef6.Cap = 0x7fffffff
	// hxfeadef6.Len = ?

	hxf83c399 := (*sliceHeader)(unsafe.Pointer(&x.control_machine))
	hxf83c399.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.control_machine))
	hxf83c399.Cap = 0x7fffffff
	// hxf83c399.Len = ?

	x.cpu_freq_def = (uint32_t)(x.ref59adaf2e.cpu_freq_def)
	x.cpu_freq_govs = (uint32_t)(x.ref59adaf2e.cpu_freq_govs)
	hxff3958a := (*sliceHeader)(unsafe.Pointer(&x.crypto_type))
	hxff3958a.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.crypto_type))
	hxff3958a.Cap = 0x7fffffff
	// hxff3958a.Len = ?

	x.debug_flags = (uint64_t)(x.ref59adaf2e.debug_flags)
	x.def_mem_per_cpu = (uint64_t)(x.ref59adaf2e.def_mem_per_cpu)
	x.disable_root_jobs = (uint16_t)(x.ref59adaf2e.disable_root_jobs)
	x.eio_timeout = (uint16_t)(x.ref59adaf2e.eio_timeout)
	x.enforce_part_limits = (uint16_t)(x.ref59adaf2e.enforce_part_limits)
	hxff327fc := (*sliceHeader)(unsafe.Pointer(&x.epilog))
	hxff327fc.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.epilog))
	hxff327fc.Cap = 0x7fffffff
	// hxff327fc.Len = ?

	x.epilog_msg_time = (uint32_t)(x.ref59adaf2e.epilog_msg_time)
	hxf3ec017 := (*sliceHeader)(unsafe.Pointer(&x.epilog_slurmctld))
	hxf3ec017.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.epilog_slurmctld))
	hxf3ec017.Cap = 0x7fffffff
	// hxf3ec017.Len = ?

	hxfc8eeef := (*sliceHeader)(unsafe.Pointer(&x.ext_sensors_type))
	hxfc8eeef.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.ext_sensors_type))
	hxfc8eeef.Cap = 0x7fffffff
	// hxfc8eeef.Len = ?

	x.ext_sensors_freq = (uint16_t)(x.ref59adaf2e.ext_sensors_freq)
	x.ext_sensors_conf = (unsafe.Pointer)(unsafe.Pointer(x.ref59adaf2e.ext_sensors_conf))
	x.fast_schedule = (uint16_t)(x.ref59adaf2e.fast_schedule)
	hxf96fbf0 := (*sliceHeader)(unsafe.Pointer(&x.fed_params))
	hxf96fbf0.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.fed_params))
	hxf96fbf0.Cap = 0x7fffffff
	// hxf96fbf0.Len = ?

	x.first_job_id = (uint32_t)(x.ref59adaf2e.first_job_id)
	x.fs_dampening_factor = (uint16_t)(x.ref59adaf2e.fs_dampening_factor)
	x.get_env_timeout = (uint16_t)(x.ref59adaf2e.get_env_timeout)
	hxff7d0c6 := (*sliceHeader)(unsafe.Pointer(&x.gres_plugins))
	hxff7d0c6.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.gres_plugins))
	hxff7d0c6.Cap = 0x7fffffff
	// hxff7d0c6.Len = ?

	x.group_time = (uint16_t)(x.ref59adaf2e.group_time)
	x.group_force = (uint16_t)(x.ref59adaf2e.group_force)
	x.hash_val = (uint32_t)(x.ref59adaf2e.hash_val)
	x.health_check_interval = (uint16_t)(x.ref59adaf2e.health_check_interval)
	x.health_check_node_state = (uint16_t)(x.ref59adaf2e.health_check_node_state)
	hxf140acc := (*sliceHeader)(unsafe.Pointer(&x.health_check_program))
	hxf140acc.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.health_check_program))
	hxf140acc.Cap = 0x7fffffff
	// hxf140acc.Len = ?

	x.inactive_limit = (uint16_t)(x.ref59adaf2e.inactive_limit)
	hxf8a0e97 := (*sliceHeader)(unsafe.Pointer(&x.job_acct_gather_freq))
	hxf8a0e97.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_acct_gather_freq))
	hxf8a0e97.Cap = 0x7fffffff
	// hxf8a0e97.Len = ?

	hxf619c1f := (*sliceHeader)(unsafe.Pointer(&x.job_acct_gather_type))
	hxf619c1f.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_acct_gather_type))
	hxf619c1f.Cap = 0x7fffffff
	// hxf619c1f.Len = ?

	hxf502907 := (*sliceHeader)(unsafe.Pointer(&x.job_acct_gather_params))
	hxf502907.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_acct_gather_params))
	hxf502907.Cap = 0x7fffffff
	// hxf502907.Len = ?

	hxfec8dad := (*sliceHeader)(unsafe.Pointer(&x.job_ckpt_dir))
	hxfec8dad.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_ckpt_dir))
	hxfec8dad.Cap = 0x7fffffff
	// hxfec8dad.Len = ?

	hxfc7c07a := (*sliceHeader)(unsafe.Pointer(&x.job_comp_host))
	hxfc7c07a.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_comp_host))
	hxfc7c07a.Cap = 0x7fffffff
	// hxfc7c07a.Len = ?

	hxf1c0c28 := (*sliceHeader)(unsafe.Pointer(&x.job_comp_loc))
	hxf1c0c28.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_comp_loc))
	hxf1c0c28.Cap = 0x7fffffff
	// hxf1c0c28.Len = ?

	hxf8f3f85 := (*sliceHeader)(unsafe.Pointer(&x.job_comp_pass))
	hxf8f3f85.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_comp_pass))
	hxf8f3f85.Cap = 0x7fffffff
	// hxf8f3f85.Len = ?

	x.job_comp_port = (uint32_t)(x.ref59adaf2e.job_comp_port)
	hxf6762ad := (*sliceHeader)(unsafe.Pointer(&x.job_comp_type))
	hxf6762ad.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_comp_type))
	hxf6762ad.Cap = 0x7fffffff
	// hxf6762ad.Len = ?

	hxf30b007 := (*sliceHeader)(unsafe.Pointer(&x.job_comp_user))
	hxf30b007.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_comp_user))
	hxf30b007.Cap = 0x7fffffff
	// hxf30b007.Len = ?

	hxfa18b5d := (*sliceHeader)(unsafe.Pointer(&x.job_container_plugin))
	hxfa18b5d.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_container_plugin))
	hxfa18b5d.Cap = 0x7fffffff
	// hxfa18b5d.Len = ?

	hxf672095 := (*sliceHeader)(unsafe.Pointer(&x.job_credential_private_key))
	hxf672095.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_credential_private_key))
	hxf672095.Cap = 0x7fffffff
	// hxf672095.Len = ?

	hxf07d19e := (*sliceHeader)(unsafe.Pointer(&x.job_credential_public_certificate))
	hxf07d19e.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_credential_public_certificate))
	hxf07d19e.Cap = 0x7fffffff
	// hxf07d19e.Len = ?

	x.job_file_append = (uint16_t)(x.ref59adaf2e.job_file_append)
	x.job_requeue = (uint16_t)(x.ref59adaf2e.job_requeue)
	hxf836955 := (*sliceHeader)(unsafe.Pointer(&x.job_submit_plugins))
	hxf836955.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.job_submit_plugins))
	hxf836955.Cap = 0x7fffffff
	// hxf836955.Len = ?

	x.keep_alive_time = (uint16_t)(x.ref59adaf2e.keep_alive_time)
	x.kill_on_bad_exit = (uint16_t)(x.ref59adaf2e.kill_on_bad_exit)
	x.kill_wait = (uint16_t)(x.ref59adaf2e.kill_wait)
	hxfdb0667 := (*sliceHeader)(unsafe.Pointer(&x.launch_params))
	hxfdb0667.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.launch_params))
	hxfdb0667.Cap = 0x7fffffff
	// hxfdb0667.Len = ?

	hxff3bbef := (*sliceHeader)(unsafe.Pointer(&x.launch_type))
	hxff3bbef.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.launch_type))
	hxff3bbef.Cap = 0x7fffffff
	// hxff3bbef.Len = ?

	hxf4bcfb3 := (*sliceHeader)(unsafe.Pointer(&x.layouts))
	hxf4bcfb3.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.layouts))
	hxf4bcfb3.Cap = 0x7fffffff
	// hxf4bcfb3.Len = ?

	hxf987d33 := (*sliceHeader)(unsafe.Pointer(&x.licenses))
	hxf987d33.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.licenses))
	hxf987d33.Cap = 0x7fffffff
	// hxf987d33.Len = ?

	hxf8eb1cd := (*sliceHeader)(unsafe.Pointer(&x.licenses_used))
	hxf8eb1cd.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.licenses_used))
	hxf8eb1cd.Cap = 0x7fffffff
	// hxf8eb1cd.Len = ?

	x.log_fmt = (uint16_t)(x.ref59adaf2e.log_fmt)
	hxf36a66e := (*sliceHeader)(unsafe.Pointer(&x.mail_domain))
	hxf36a66e.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.mail_domain))
	hxf36a66e.Cap = 0x7fffffff
	// hxf36a66e.Len = ?

	hxf867581 := (*sliceHeader)(unsafe.Pointer(&x.mail_prog))
	hxf867581.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.mail_prog))
	hxf867581.Cap = 0x7fffffff
	// hxf867581.Len = ?

	x.max_array_sz = (uint32_t)(x.ref59adaf2e.max_array_sz)
	x.max_job_cnt = (uint32_t)(x.ref59adaf2e.max_job_cnt)
	x.max_job_id = (uint32_t)(x.ref59adaf2e.max_job_id)
	x.max_mem_per_cpu = (uint64_t)(x.ref59adaf2e.max_mem_per_cpu)
	x.max_step_cnt = (uint32_t)(x.ref59adaf2e.max_step_cnt)
	x.max_tasks_per_node = (uint16_t)(x.ref59adaf2e.max_tasks_per_node)
	hxfd23ed4 := (*sliceHeader)(unsafe.Pointer(&x.mcs_plugin))
	hxfd23ed4.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.mcs_plugin))
	hxfd23ed4.Cap = 0x7fffffff
	// hxfd23ed4.Len = ?

	hxf346633 := (*sliceHeader)(unsafe.Pointer(&x.mcs_plugin_params))
	hxf346633.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.mcs_plugin_params))
	hxf346633.Cap = 0x7fffffff
	// hxf346633.Len = ?

	x.mem_limit_enforce = (uint16_t)(x.ref59adaf2e.mem_limit_enforce)
	x.min_job_age = (uint32_t)(x.ref59adaf2e.min_job_age)
	hxfff3f9d := (*sliceHeader)(unsafe.Pointer(&x.mpi_default))
	hxfff3f9d.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.mpi_default))
	hxfff3f9d.Cap = 0x7fffffff
	// hxfff3f9d.Len = ?

	hxfbc17f8 := (*sliceHeader)(unsafe.Pointer(&x.mpi_params))
	hxfbc17f8.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.mpi_params))
	hxfbc17f8.Cap = 0x7fffffff
	// hxfbc17f8.Len = ?

	hxf5716c8 := (*sliceHeader)(unsafe.Pointer(&x.msg_aggr_params))
	hxf5716c8.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.msg_aggr_params))
	hxf5716c8.Cap = 0x7fffffff
	// hxf5716c8.Len = ?

	x.msg_timeout = (uint16_t)(x.ref59adaf2e.msg_timeout)
	x.tcp_timeout = (uint16_t)(x.ref59adaf2e.tcp_timeout)
	x.next_job_id = (uint32_t)(x.ref59adaf2e.next_job_id)
	hxf8c458d := (*sliceHeader)(unsafe.Pointer(&x.node_features_plugins))
	hxf8c458d.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.node_features_plugins))
	hxf8c458d.Cap = 0x7fffffff
	// hxf8c458d.Len = ?

	hxf5389fd := (*sliceHeader)(unsafe.Pointer(&x.node_prefix))
	hxf5389fd.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.node_prefix))
	hxf5389fd.Cap = 0x7fffffff
	// hxf5389fd.Len = ?

	x.over_time_limit = (uint16_t)(x.ref59adaf2e.over_time_limit)
	hxfd0226a := (*sliceHeader)(unsafe.Pointer(&x.plugindir))
	hxfd0226a.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.plugindir))
	hxfd0226a.Cap = 0x7fffffff
	// hxfd0226a.Len = ?

	hxfc95d51 := (*sliceHeader)(unsafe.Pointer(&x.plugstack))
	hxfc95d51.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.plugstack))
	hxfc95d51.Cap = 0x7fffffff
	// hxfc95d51.Len = ?

	hxfc56416 := (*sliceHeader)(unsafe.Pointer(&x.power_parameters))
	hxfc56416.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.power_parameters))
	hxfc56416.Cap = 0x7fffffff
	// hxfc56416.Len = ?

	hxf76be2b := (*sliceHeader)(unsafe.Pointer(&x.power_plugin))
	hxf76be2b.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.power_plugin))
	hxf76be2b.Cap = 0x7fffffff
	// hxf76be2b.Len = ?

	x.preempt_mode = (uint16_t)(x.ref59adaf2e.preempt_mode)
	hxfd582b8 := (*sliceHeader)(unsafe.Pointer(&x.preempt_type))
	hxfd582b8.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.preempt_type))
	hxfd582b8.Cap = 0x7fffffff
	// hxfd582b8.Len = ?

	x.priority_decay_hl = (uint32_t)(x.ref59adaf2e.priority_decay_hl)
	x.priority_calc_period = (uint32_t)(x.ref59adaf2e.priority_calc_period)
	x.priority_favor_small = (uint16_t)(x.ref59adaf2e.priority_favor_small)
	x.priority_flags = (uint16_t)(x.ref59adaf2e.priority_flags)
	x.priority_max_age = (uint32_t)(x.ref59adaf2e.priority_max_age)
	hxf845227 := (*sliceHeader)(unsafe.Pointer(&x.priority_params))
	hxf845227.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.priority_params))
	hxf845227.Cap = 0x7fffffff
	// hxf845227.Len = ?

	x.priority_reset_period = (uint16_t)(x.ref59adaf2e.priority_reset_period)
	hxfb61bec := (*sliceHeader)(unsafe.Pointer(&x.priority_type))
	hxfb61bec.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.priority_type))
	hxfb61bec.Cap = 0x7fffffff
	// hxfb61bec.Len = ?

	x.priority_weight_age = (uint32_t)(x.ref59adaf2e.priority_weight_age)
	x.priority_weight_fs = (uint32_t)(x.ref59adaf2e.priority_weight_fs)
	x.priority_weight_js = (uint32_t)(x.ref59adaf2e.priority_weight_js)
	x.priority_weight_part = (uint32_t)(x.ref59adaf2e.priority_weight_part)
	x.priority_weight_qos = (uint32_t)(x.ref59adaf2e.priority_weight_qos)
	hxf9e710a := (*sliceHeader)(unsafe.Pointer(&x.priority_weight_tres))
	hxf9e710a.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.priority_weight_tres))
	hxf9e710a.Cap = 0x7fffffff
	// hxf9e710a.Len = ?

	x.private_data = (uint16_t)(x.ref59adaf2e.private_data)
	hxf77eea6 := (*sliceHeader)(unsafe.Pointer(&x.proctrack_type))
	hxf77eea6.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.proctrack_type))
	hxf77eea6.Cap = 0x7fffffff
	// hxf77eea6.Len = ?

	hxf420be6 := (*sliceHeader)(unsafe.Pointer(&x.prolog))
	hxf420be6.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.prolog))
	hxf420be6.Cap = 0x7fffffff
	// hxf420be6.Len = ?

	x.prolog_epilog_timeout = (uint16_t)(x.ref59adaf2e.prolog_epilog_timeout)
	hxf622631 := (*sliceHeader)(unsafe.Pointer(&x.prolog_slurmctld))
	hxf622631.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.prolog_slurmctld))
	hxf622631.Cap = 0x7fffffff
	// hxf622631.Len = ?

	x.propagate_prio_process = (uint16_t)(x.ref59adaf2e.propagate_prio_process)
	x.prolog_flags = (uint16_t)(x.ref59adaf2e.prolog_flags)
	hxfe19e70 := (*sliceHeader)(unsafe.Pointer(&x.propagate_rlimits))
	hxfe19e70.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.propagate_rlimits))
	hxfe19e70.Cap = 0x7fffffff
	// hxfe19e70.Len = ?

	hxf3f5e97 := (*sliceHeader)(unsafe.Pointer(&x.propagate_rlimits_except))
	hxf3f5e97.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.propagate_rlimits_except))
	hxf3f5e97.Cap = 0x7fffffff
	// hxf3f5e97.Len = ?

	hxfa02c27 := (*sliceHeader)(unsafe.Pointer(&x.reboot_program))
	hxfa02c27.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.reboot_program))
	hxfa02c27.Cap = 0x7fffffff
	// hxfa02c27.Len = ?

	x.reconfig_flags = (uint16_t)(x.ref59adaf2e.reconfig_flags)
	hxf353ef1 := (*sliceHeader)(unsafe.Pointer(&x.requeue_exit))
	hxf353ef1.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.requeue_exit))
	hxf353ef1.Cap = 0x7fffffff
	// hxf353ef1.Len = ?

	hxf05e166 := (*sliceHeader)(unsafe.Pointer(&x.requeue_exit_hold))
	hxf05e166.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.requeue_exit_hold))
	hxf05e166.Cap = 0x7fffffff
	// hxf05e166.Len = ?

	hxf6a81f7 := (*sliceHeader)(unsafe.Pointer(&x.resume_program))
	hxf6a81f7.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.resume_program))
	hxf6a81f7.Cap = 0x7fffffff
	// hxf6a81f7.Len = ?

	x.resume_rate = (uint16_t)(x.ref59adaf2e.resume_rate)
	x.resume_timeout = (uint16_t)(x.ref59adaf2e.resume_timeout)
	hxf93d498 := (*sliceHeader)(unsafe.Pointer(&x.resv_epilog))
	hxf93d498.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.resv_epilog))
	hxf93d498.Cap = 0x7fffffff
	// hxf93d498.Len = ?

	x.resv_over_run = (uint16_t)(x.ref59adaf2e.resv_over_run)
	hxf9e0055 := (*sliceHeader)(unsafe.Pointer(&x.resv_prolog))
	hxf9e0055.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.resv_prolog))
	hxf9e0055.Cap = 0x7fffffff
	// hxf9e0055.Len = ?

	x.ret2service = (uint16_t)(x.ref59adaf2e.ret2service)
	hxf694a16 := (*sliceHeader)(unsafe.Pointer(&x.route_plugin))
	hxf694a16.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.route_plugin))
	hxf694a16.Cap = 0x7fffffff
	// hxf694a16.Len = ?

	hxf322790 := (*sliceHeader)(unsafe.Pointer(&x.salloc_default_command))
	hxf322790.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.salloc_default_command))
	hxf322790.Cap = 0x7fffffff
	// hxf322790.Len = ?

	hxf691471 := (*sliceHeader)(unsafe.Pointer(&x.sbcast_parameters))
	hxf691471.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.sbcast_parameters))
	hxf691471.Cap = 0x7fffffff
	// hxf691471.Len = ?

	hxf9ef69f := (*sliceHeader)(unsafe.Pointer(&x.sched_logfile))
	hxf9ef69f.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.sched_logfile))
	hxf9ef69f.Cap = 0x7fffffff
	// hxf9ef69f.Len = ?

	x.sched_log_level = (uint16_t)(x.ref59adaf2e.sched_log_level)
	hxf27c55d := (*sliceHeader)(unsafe.Pointer(&x.sched_params))
	hxf27c55d.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.sched_params))
	hxf27c55d.Cap = 0x7fffffff
	// hxf27c55d.Len = ?

	x.sched_time_slice = (uint16_t)(x.ref59adaf2e.sched_time_slice)
	hxf5094fc := (*sliceHeader)(unsafe.Pointer(&x.schedtype))
	hxf5094fc.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.schedtype))
	hxf5094fc.Cap = 0x7fffffff
	// hxf5094fc.Len = ?

	hxff2b7d1 := (*sliceHeader)(unsafe.Pointer(&x.select_type))
	hxff2b7d1.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.select_type))
	hxff2b7d1.Cap = 0x7fffffff
	// hxff2b7d1.Len = ?

	x.select_conf_key_pairs = (unsafe.Pointer)(unsafe.Pointer(x.ref59adaf2e.select_conf_key_pairs))
	x.select_type_param = (uint16_t)(x.ref59adaf2e.select_type_param)
	hxf9ed7b5 := (*sliceHeader)(unsafe.Pointer(&x.slurm_conf))
	hxf9ed7b5.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.slurm_conf))
	hxf9ed7b5.Cap = 0x7fffffff
	// hxf9ed7b5.Len = ?

	x.slurm_user_id = (uint32_t)(x.ref59adaf2e.slurm_user_id)
	hxfa09f29 := (*sliceHeader)(unsafe.Pointer(&x.slurm_user_name))
	hxfa09f29.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.slurm_user_name))
	hxfa09f29.Cap = 0x7fffffff
	// hxfa09f29.Len = ?

	x.slurmd_user_id = (uint32_t)(x.ref59adaf2e.slurmd_user_id)
	hxf4ab0e6 := (*sliceHeader)(unsafe.Pointer(&x.slurmd_user_name))
	hxf4ab0e6.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.slurmd_user_name))
	hxf4ab0e6.Cap = 0x7fffffff
	// hxf4ab0e6.Len = ?

	x.slurmctld_debug = (uint16_t)(x.ref59adaf2e.slurmctld_debug)
	hxf1bab63 := (*sliceHeader)(unsafe.Pointer(&x.slurmctld_logfile))
	hxf1bab63.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.slurmctld_logfile))
	hxf1bab63.Cap = 0x7fffffff
	// hxf1bab63.Len = ?

	hxfcc7b44 := (*sliceHeader)(unsafe.Pointer(&x.slurmctld_pidfile))
	hxfcc7b44.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.slurmctld_pidfile))
	hxfcc7b44.Cap = 0x7fffffff
	// hxfcc7b44.Len = ?

	hxf93ab8c := (*sliceHeader)(unsafe.Pointer(&x.slurmctld_plugstack))
	hxf93ab8c.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.slurmctld_plugstack))
	hxf93ab8c.Cap = 0x7fffffff
	// hxf93ab8c.Len = ?

	x.slurmctld_port = (uint32_t)(x.ref59adaf2e.slurmctld_port)
	x.slurmctld_port_count = (uint16_t)(x.ref59adaf2e.slurmctld_port_count)
	x.slurmctld_syslog_debug = (uint16_t)(x.ref59adaf2e.slurmctld_syslog_debug)
	x.slurmctld_timeout = (uint16_t)(x.ref59adaf2e.slurmctld_timeout)
	x.slurmd_debug = (uint16_t)(x.ref59adaf2e.slurmd_debug)
	hxf503fac := (*sliceHeader)(unsafe.Pointer(&x.slurmd_logfile))
	hxf503fac.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.slurmd_logfile))
	hxf503fac.Cap = 0x7fffffff
	// hxf503fac.Len = ?

	hxf635566 := (*sliceHeader)(unsafe.Pointer(&x.slurmd_pidfile))
	hxf635566.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.slurmd_pidfile))
	hxf635566.Cap = 0x7fffffff
	// hxf635566.Len = ?

	x.slurmd_port = (uint32_t)(x.ref59adaf2e.slurmd_port)
	hxf5b79ae := (*sliceHeader)(unsafe.Pointer(&x.slurmd_spooldir))
	hxf5b79ae.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.slurmd_spooldir))
	hxf5b79ae.Cap = 0x7fffffff
	// hxf5b79ae.Len = ?

	x.slurmd_syslog_debug = (uint16_t)(x.ref59adaf2e.slurmd_syslog_debug)
	x.slurmd_timeout = (uint16_t)(x.ref59adaf2e.slurmd_timeout)
	hxf8bfe02 := (*sliceHeader)(unsafe.Pointer(&x.srun_epilog))
	hxf8bfe02.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.srun_epilog))
	hxf8bfe02.Cap = 0x7fffffff
	// hxf8bfe02.Len = ?

	hxf00db10 := (*sliceHeader)(unsafe.Pointer(&x.srun_port_range))
	hxf00db10.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.srun_port_range))
	hxf00db10.Cap = 0x7fffffff
	// hxf00db10.Len = ?

	hxf0e1b46 := (*sliceHeader)(unsafe.Pointer(&x.srun_prolog))
	hxf0e1b46.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.srun_prolog))
	hxf0e1b46.Cap = 0x7fffffff
	// hxf0e1b46.Len = ?

	hxf0680d5 := (*sliceHeader)(unsafe.Pointer(&x.state_save_location))
	hxf0680d5.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.state_save_location))
	hxf0680d5.Cap = 0x7fffffff
	// hxf0680d5.Len = ?

	hxf63665a := (*sliceHeader)(unsafe.Pointer(&x.suspend_exc_nodes))
	hxf63665a.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.suspend_exc_nodes))
	hxf63665a.Cap = 0x7fffffff
	// hxf63665a.Len = ?

	hxf6aa19c := (*sliceHeader)(unsafe.Pointer(&x.suspend_exc_parts))
	hxf6aa19c.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.suspend_exc_parts))
	hxf6aa19c.Cap = 0x7fffffff
	// hxf6aa19c.Len = ?

	hxf805226 := (*sliceHeader)(unsafe.Pointer(&x.suspend_program))
	hxf805226.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.suspend_program))
	hxf805226.Cap = 0x7fffffff
	// hxf805226.Len = ?

	x.suspend_rate = (uint16_t)(x.ref59adaf2e.suspend_rate)
	x.suspend_time = (uint32_t)(x.ref59adaf2e.suspend_time)
	x.suspend_timeout = (uint16_t)(x.ref59adaf2e.suspend_timeout)
	hxf6d0687 := (*sliceHeader)(unsafe.Pointer(&x.switch_type))
	hxf6d0687.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.switch_type))
	hxf6d0687.Cap = 0x7fffffff
	// hxf6d0687.Len = ?

	hxf92df9b := (*sliceHeader)(unsafe.Pointer(&x.task_epilog))
	hxf92df9b.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.task_epilog))
	hxf92df9b.Cap = 0x7fffffff
	// hxf92df9b.Len = ?

	hxf4972cc := (*sliceHeader)(unsafe.Pointer(&x.task_plugin))
	hxf4972cc.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.task_plugin))
	hxf4972cc.Cap = 0x7fffffff
	// hxf4972cc.Len = ?

	x.task_plugin_param = (uint32_t)(x.ref59adaf2e.task_plugin_param)
	hxf25fbc6 := (*sliceHeader)(unsafe.Pointer(&x.task_prolog))
	hxf25fbc6.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.task_prolog))
	hxf25fbc6.Cap = 0x7fffffff
	// hxf25fbc6.Len = ?

	hxf83e67b := (*sliceHeader)(unsafe.Pointer(&x.tmp_fs))
	hxf83e67b.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.tmp_fs))
	hxf83e67b.Cap = 0x7fffffff
	// hxf83e67b.Len = ?

	hxffda601 := (*sliceHeader)(unsafe.Pointer(&x.topology_param))
	hxffda601.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.topology_param))
	hxffda601.Cap = 0x7fffffff
	// hxffda601.Len = ?

	hxf3c3d06 := (*sliceHeader)(unsafe.Pointer(&x.topology_plugin))
	hxf3c3d06.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.topology_plugin))
	hxf3c3d06.Cap = 0x7fffffff
	// hxf3c3d06.Len = ?

	x.track_wckey = (uint16_t)(x.ref59adaf2e.track_wckey)
	x.tree_width = (uint16_t)(x.ref59adaf2e.tree_width)
	hxf824b9a := (*sliceHeader)(unsafe.Pointer(&x.unkillable_program))
	hxf824b9a.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.unkillable_program))
	hxf824b9a.Cap = 0x7fffffff
	// hxf824b9a.Len = ?

	x.unkillable_timeout = (uint16_t)(x.ref59adaf2e.unkillable_timeout)
	x.use_pam = (uint16_t)(x.ref59adaf2e.use_pam)
	x.use_spec_resources = (uint16_t)(x.ref59adaf2e.use_spec_resources)
	hxfbfce3a := (*sliceHeader)(unsafe.Pointer(&x.version))
	hxfbfce3a.Data = uintptr(unsafe.Pointer(x.ref59adaf2e.version))
	hxfbfce3a.Cap = 0x7fffffff
	// hxfbfce3a.Len = ?

	x.vsize_factor = (uint16_t)(x.ref59adaf2e.vsize_factor)
	x.wait_time = (uint16_t)(x.ref59adaf2e.wait_time)
}

// allocSlurmd_status_tMemory allocates memory for type C.slurmd_status_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurmd_status_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurmd_status_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurmd_status_tValue = unsafe.Sizeof([1]C.slurmd_status_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurmd_status_t) Ref() *C.slurmd_status_t {
	if x == nil {
		return nil
	}
	return x.ref72c806fa
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurmd_status_t) Free() {
	if x != nil && x.allocs72c806fa != nil {
		x.allocs72c806fa.(*cgoAllocMap).Free()
		x.ref72c806fa = nil
	}
}

// Newslurmd_status_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurmd_status_tRef(ref unsafe.Pointer) *slurmd_status_t {
	if ref == nil {
		return nil
	}
	obj := new(slurmd_status_t)
	obj.ref72c806fa = (*C.slurmd_status_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurmd_status_t) PassRef() (*C.slurmd_status_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref72c806fa != nil {
		return x.ref72c806fa, nil
	}
	mem72c806fa := allocSlurmd_status_tMemory(1)
	ref72c806fa := (*C.slurmd_status_t)(mem72c806fa)
	allocs72c806fa := new(cgoAllocMap)
	allocs72c806fa.Add(mem72c806fa)

	var cbooted_allocs *cgoAllocMap
	ref72c806fa.booted, cbooted_allocs = (C.time_t)(x.booted), cgoAllocsUnknown
	allocs72c806fa.Borrow(cbooted_allocs)

	var clast_slurmctld_msg_allocs *cgoAllocMap
	ref72c806fa.last_slurmctld_msg, clast_slurmctld_msg_allocs = (C.time_t)(x.last_slurmctld_msg), cgoAllocsUnknown
	allocs72c806fa.Borrow(clast_slurmctld_msg_allocs)

	var cslurmd_debug_allocs *cgoAllocMap
	ref72c806fa.slurmd_debug, cslurmd_debug_allocs = (C.uint16_t)(x.slurmd_debug), cgoAllocsUnknown
	allocs72c806fa.Borrow(cslurmd_debug_allocs)

	var cactual_cpus_allocs *cgoAllocMap
	ref72c806fa.actual_cpus, cactual_cpus_allocs = (C.uint16_t)(x.actual_cpus), cgoAllocsUnknown
	allocs72c806fa.Borrow(cactual_cpus_allocs)

	var cactual_boards_allocs *cgoAllocMap
	ref72c806fa.actual_boards, cactual_boards_allocs = (C.uint16_t)(x.actual_boards), cgoAllocsUnknown
	allocs72c806fa.Borrow(cactual_boards_allocs)

	var cactual_sockets_allocs *cgoAllocMap
	ref72c806fa.actual_sockets, cactual_sockets_allocs = (C.uint16_t)(x.actual_sockets), cgoAllocsUnknown
	allocs72c806fa.Borrow(cactual_sockets_allocs)

	var cactual_cores_allocs *cgoAllocMap
	ref72c806fa.actual_cores, cactual_cores_allocs = (C.uint16_t)(x.actual_cores), cgoAllocsUnknown
	allocs72c806fa.Borrow(cactual_cores_allocs)

	var cactual_threads_allocs *cgoAllocMap
	ref72c806fa.actual_threads, cactual_threads_allocs = (C.uint16_t)(x.actual_threads), cgoAllocsUnknown
	allocs72c806fa.Borrow(cactual_threads_allocs)

	var cactual_real_mem_allocs *cgoAllocMap
	ref72c806fa.actual_real_mem, cactual_real_mem_allocs = (C.uint64_t)(x.actual_real_mem), cgoAllocsUnknown
	allocs72c806fa.Borrow(cactual_real_mem_allocs)

	var cactual_tmp_disk_allocs *cgoAllocMap
	ref72c806fa.actual_tmp_disk, cactual_tmp_disk_allocs = (C.uint32_t)(x.actual_tmp_disk), cgoAllocsUnknown
	allocs72c806fa.Borrow(cactual_tmp_disk_allocs)

	var cpid_allocs *cgoAllocMap
	ref72c806fa.pid, cpid_allocs = (C.uint32_t)(x.pid), cgoAllocsUnknown
	allocs72c806fa.Borrow(cpid_allocs)

	var chostname_allocs *cgoAllocMap
	ref72c806fa.hostname, chostname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.hostname)).Data)), cgoAllocsUnknown
	allocs72c806fa.Borrow(chostname_allocs)

	var cslurmd_logfile_allocs *cgoAllocMap
	ref72c806fa.slurmd_logfile, cslurmd_logfile_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.slurmd_logfile)).Data)), cgoAllocsUnknown
	allocs72c806fa.Borrow(cslurmd_logfile_allocs)

	var cstep_list_allocs *cgoAllocMap
	ref72c806fa.step_list, cstep_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.step_list)).Data)), cgoAllocsUnknown
	allocs72c806fa.Borrow(cstep_list_allocs)

	var cversion_allocs *cgoAllocMap
	ref72c806fa.version, cversion_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.version)).Data)), cgoAllocsUnknown
	allocs72c806fa.Borrow(cversion_allocs)

	x.ref72c806fa = ref72c806fa
	x.allocs72c806fa = allocs72c806fa
	return ref72c806fa, allocs72c806fa

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurmd_status_t) PassValue() (C.slurmd_status_t, *cgoAllocMap) {
	if x.ref72c806fa != nil {
		return *x.ref72c806fa, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurmd_status_t) Deref() {
	if x.ref72c806fa == nil {
		return
	}
	x.booted = (time_t)(x.ref72c806fa.booted)
	x.last_slurmctld_msg = (time_t)(x.ref72c806fa.last_slurmctld_msg)
	x.slurmd_debug = (uint16_t)(x.ref72c806fa.slurmd_debug)
	x.actual_cpus = (uint16_t)(x.ref72c806fa.actual_cpus)
	x.actual_boards = (uint16_t)(x.ref72c806fa.actual_boards)
	x.actual_sockets = (uint16_t)(x.ref72c806fa.actual_sockets)
	x.actual_cores = (uint16_t)(x.ref72c806fa.actual_cores)
	x.actual_threads = (uint16_t)(x.ref72c806fa.actual_threads)
	x.actual_real_mem = (uint64_t)(x.ref72c806fa.actual_real_mem)
	x.actual_tmp_disk = (uint32_t)(x.ref72c806fa.actual_tmp_disk)
	x.pid = (uint32_t)(x.ref72c806fa.pid)
	hxff5c0c3 := (*sliceHeader)(unsafe.Pointer(&x.hostname))
	hxff5c0c3.Data = uintptr(unsafe.Pointer(x.ref72c806fa.hostname))
	hxff5c0c3.Cap = 0x7fffffff
	// hxff5c0c3.Len = ?

	hxfe48a96 := (*sliceHeader)(unsafe.Pointer(&x.slurmd_logfile))
	hxfe48a96.Data = uintptr(unsafe.Pointer(x.ref72c806fa.slurmd_logfile))
	hxfe48a96.Cap = 0x7fffffff
	// hxfe48a96.Len = ?

	hxfde274b := (*sliceHeader)(unsafe.Pointer(&x.step_list))
	hxfde274b.Data = uintptr(unsafe.Pointer(x.ref72c806fa.step_list))
	hxfde274b.Cap = 0x7fffffff
	// hxfde274b.Len = ?

	hxfb1ae22 := (*sliceHeader)(unsafe.Pointer(&x.version))
	hxfb1ae22.Data = uintptr(unsafe.Pointer(x.ref72c806fa.version))
	hxfb1ae22.Cap = 0x7fffffff
	// hxfb1ae22.Len = ?

}

// allocSubmit_response_msg_tMemory allocates memory for type C.submit_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSubmit_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSubmit_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSubmit_response_msg_tValue = unsafe.Sizeof([1]C.submit_response_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *submit_response_msg_t) Ref() *C.submit_response_msg_t {
	if x == nil {
		return nil
	}
	return x.ref6c6e601
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *submit_response_msg_t) Free() {
	if x != nil && x.allocs6c6e601 != nil {
		x.allocs6c6e601.(*cgoAllocMap).Free()
		x.ref6c6e601 = nil
	}
}

// Newsubmit_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newsubmit_response_msg_tRef(ref unsafe.Pointer) *submit_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(submit_response_msg_t)
	obj.ref6c6e601 = (*C.submit_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *submit_response_msg_t) PassRef() (*C.submit_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6c6e601 != nil {
		return x.ref6c6e601, nil
	}
	mem6c6e601 := allocSubmit_response_msg_tMemory(1)
	ref6c6e601 := (*C.submit_response_msg_t)(mem6c6e601)
	allocs6c6e601 := new(cgoAllocMap)
	allocs6c6e601.Add(mem6c6e601)

	var cjob_id_allocs *cgoAllocMap
	ref6c6e601.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs6c6e601.Borrow(cjob_id_allocs)

	var cstep_id_allocs *cgoAllocMap
	ref6c6e601.step_id, cstep_id_allocs = (C.uint32_t)(x.step_id), cgoAllocsUnknown
	allocs6c6e601.Borrow(cstep_id_allocs)

	var cerror_code_allocs *cgoAllocMap
	ref6c6e601.error_code, cerror_code_allocs = (C.uint32_t)(x.error_code), cgoAllocsUnknown
	allocs6c6e601.Borrow(cerror_code_allocs)

	var cjob_submit_user_msg_allocs *cgoAllocMap
	ref6c6e601.job_submit_user_msg, cjob_submit_user_msg_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.job_submit_user_msg)).Data)), cgoAllocsUnknown
	allocs6c6e601.Borrow(cjob_submit_user_msg_allocs)

	x.ref6c6e601 = ref6c6e601
	x.allocs6c6e601 = allocs6c6e601
	return ref6c6e601, allocs6c6e601

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x submit_response_msg_t) PassValue() (C.submit_response_msg_t, *cgoAllocMap) {
	if x.ref6c6e601 != nil {
		return *x.ref6c6e601, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *submit_response_msg_t) Deref() {
	if x.ref6c6e601 == nil {
		return
	}
	x.job_id = (uint32_t)(x.ref6c6e601.job_id)
	x.step_id = (uint32_t)(x.ref6c6e601.step_id)
	x.error_code = (uint32_t)(x.ref6c6e601.error_code)
	hxfcc0acc := (*sliceHeader)(unsafe.Pointer(&x.job_submit_user_msg))
	hxfcc0acc.Data = uintptr(unsafe.Pointer(x.ref6c6e601.job_submit_user_msg))
	hxfcc0acc.Cap = 0x7fffffff
	// hxfcc0acc.Len = ?

}

// allocUpdate_node_msg_tMemory allocates memory for type C.update_node_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUpdate_node_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUpdate_node_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfUpdate_node_msg_tValue = unsafe.Sizeof([1]C.update_node_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *update_node_msg_t) Ref() *C.update_node_msg_t {
	if x == nil {
		return nil
	}
	return x.ref7d01926e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *update_node_msg_t) Free() {
	if x != nil && x.allocs7d01926e != nil {
		x.allocs7d01926e.(*cgoAllocMap).Free()
		x.ref7d01926e = nil
	}
}

// Newupdate_node_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newupdate_node_msg_tRef(ref unsafe.Pointer) *update_node_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(update_node_msg_t)
	obj.ref7d01926e = (*C.update_node_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *update_node_msg_t) PassRef() (*C.update_node_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref7d01926e != nil {
		return x.ref7d01926e, nil
	}
	mem7d01926e := allocUpdate_node_msg_tMemory(1)
	ref7d01926e := (*C.update_node_msg_t)(mem7d01926e)
	allocs7d01926e := new(cgoAllocMap)
	allocs7d01926e.Add(mem7d01926e)

	var cfeatures_allocs *cgoAllocMap
	ref7d01926e.features, cfeatures_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.features)).Data)), cgoAllocsUnknown
	allocs7d01926e.Borrow(cfeatures_allocs)

	var cfeatures_act_allocs *cgoAllocMap
	ref7d01926e.features_act, cfeatures_act_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.features_act)).Data)), cgoAllocsUnknown
	allocs7d01926e.Borrow(cfeatures_act_allocs)

	var cgres_allocs *cgoAllocMap
	ref7d01926e.gres, cgres_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.gres)).Data)), cgoAllocsUnknown
	allocs7d01926e.Borrow(cgres_allocs)

	var cnode_addr_allocs *cgoAllocMap
	ref7d01926e.node_addr, cnode_addr_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_addr)).Data)), cgoAllocsUnknown
	allocs7d01926e.Borrow(cnode_addr_allocs)

	var cnode_hostname_allocs *cgoAllocMap
	ref7d01926e.node_hostname, cnode_hostname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_hostname)).Data)), cgoAllocsUnknown
	allocs7d01926e.Borrow(cnode_hostname_allocs)

	var cnode_names_allocs *cgoAllocMap
	ref7d01926e.node_names, cnode_names_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_names)).Data)), cgoAllocsUnknown
	allocs7d01926e.Borrow(cnode_names_allocs)

	var cnode_state_allocs *cgoAllocMap
	ref7d01926e.node_state, cnode_state_allocs = (C.uint32_t)(x.node_state), cgoAllocsUnknown
	allocs7d01926e.Borrow(cnode_state_allocs)

	var creason_allocs *cgoAllocMap
	ref7d01926e.reason, creason_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.reason)).Data)), cgoAllocsUnknown
	allocs7d01926e.Borrow(creason_allocs)

	var creason_uid_allocs *cgoAllocMap
	ref7d01926e.reason_uid, creason_uid_allocs = (C.uint32_t)(x.reason_uid), cgoAllocsUnknown
	allocs7d01926e.Borrow(creason_uid_allocs)

	var cweight_allocs *cgoAllocMap
	ref7d01926e.weight, cweight_allocs = (C.uint32_t)(x.weight), cgoAllocsUnknown
	allocs7d01926e.Borrow(cweight_allocs)

	x.ref7d01926e = ref7d01926e
	x.allocs7d01926e = allocs7d01926e
	return ref7d01926e, allocs7d01926e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x update_node_msg_t) PassValue() (C.update_node_msg_t, *cgoAllocMap) {
	if x.ref7d01926e != nil {
		return *x.ref7d01926e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *update_node_msg_t) Deref() {
	if x.ref7d01926e == nil {
		return
	}
	hxfd6d733 := (*sliceHeader)(unsafe.Pointer(&x.features))
	hxfd6d733.Data = uintptr(unsafe.Pointer(x.ref7d01926e.features))
	hxfd6d733.Cap = 0x7fffffff
	// hxfd6d733.Len = ?

	hxf07fabc := (*sliceHeader)(unsafe.Pointer(&x.features_act))
	hxf07fabc.Data = uintptr(unsafe.Pointer(x.ref7d01926e.features_act))
	hxf07fabc.Cap = 0x7fffffff
	// hxf07fabc.Len = ?

	hxfb12572 := (*sliceHeader)(unsafe.Pointer(&x.gres))
	hxfb12572.Data = uintptr(unsafe.Pointer(x.ref7d01926e.gres))
	hxfb12572.Cap = 0x7fffffff
	// hxfb12572.Len = ?

	hxf14b2ef := (*sliceHeader)(unsafe.Pointer(&x.node_addr))
	hxf14b2ef.Data = uintptr(unsafe.Pointer(x.ref7d01926e.node_addr))
	hxf14b2ef.Cap = 0x7fffffff
	// hxf14b2ef.Len = ?

	hxfb36399 := (*sliceHeader)(unsafe.Pointer(&x.node_hostname))
	hxfb36399.Data = uintptr(unsafe.Pointer(x.ref7d01926e.node_hostname))
	hxfb36399.Cap = 0x7fffffff
	// hxfb36399.Len = ?

	hxf8e0acd := (*sliceHeader)(unsafe.Pointer(&x.node_names))
	hxf8e0acd.Data = uintptr(unsafe.Pointer(x.ref7d01926e.node_names))
	hxf8e0acd.Cap = 0x7fffffff
	// hxf8e0acd.Len = ?

	x.node_state = (uint32_t)(x.ref7d01926e.node_state)
	hxfc730dc := (*sliceHeader)(unsafe.Pointer(&x.reason))
	hxfc730dc.Data = uintptr(unsafe.Pointer(x.ref7d01926e.reason))
	hxfc730dc.Cap = 0x7fffffff
	// hxfc730dc.Len = ?

	x.reason_uid = (uint32_t)(x.ref7d01926e.reason_uid)
	x.weight = (uint32_t)(x.ref7d01926e.weight)
}

// allocUpdate_front_end_msg_tMemory allocates memory for type C.update_front_end_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUpdate_front_end_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUpdate_front_end_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfUpdate_front_end_msg_tValue = unsafe.Sizeof([1]C.update_front_end_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *update_front_end_msg_t) Ref() *C.update_front_end_msg_t {
	if x == nil {
		return nil
	}
	return x.ref157a0a01
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *update_front_end_msg_t) Free() {
	if x != nil && x.allocs157a0a01 != nil {
		x.allocs157a0a01.(*cgoAllocMap).Free()
		x.ref157a0a01 = nil
	}
}

// Newupdate_front_end_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newupdate_front_end_msg_tRef(ref unsafe.Pointer) *update_front_end_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(update_front_end_msg_t)
	obj.ref157a0a01 = (*C.update_front_end_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *update_front_end_msg_t) PassRef() (*C.update_front_end_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref157a0a01 != nil {
		return x.ref157a0a01, nil
	}
	mem157a0a01 := allocUpdate_front_end_msg_tMemory(1)
	ref157a0a01 := (*C.update_front_end_msg_t)(mem157a0a01)
	allocs157a0a01 := new(cgoAllocMap)
	allocs157a0a01.Add(mem157a0a01)

	var cname_allocs *cgoAllocMap
	ref157a0a01.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs157a0a01.Borrow(cname_allocs)

	var cnode_state_allocs *cgoAllocMap
	ref157a0a01.node_state, cnode_state_allocs = (C.uint32_t)(x.node_state), cgoAllocsUnknown
	allocs157a0a01.Borrow(cnode_state_allocs)

	var creason_allocs *cgoAllocMap
	ref157a0a01.reason, creason_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.reason)).Data)), cgoAllocsUnknown
	allocs157a0a01.Borrow(creason_allocs)

	var creason_uid_allocs *cgoAllocMap
	ref157a0a01.reason_uid, creason_uid_allocs = (C.uint32_t)(x.reason_uid), cgoAllocsUnknown
	allocs157a0a01.Borrow(creason_uid_allocs)

	x.ref157a0a01 = ref157a0a01
	x.allocs157a0a01 = allocs157a0a01
	return ref157a0a01, allocs157a0a01

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x update_front_end_msg_t) PassValue() (C.update_front_end_msg_t, *cgoAllocMap) {
	if x.ref157a0a01 != nil {
		return *x.ref157a0a01, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *update_front_end_msg_t) Deref() {
	if x.ref157a0a01 == nil {
		return
	}
	hxf545083 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf545083.Data = uintptr(unsafe.Pointer(x.ref157a0a01.name))
	hxf545083.Cap = 0x7fffffff
	// hxf545083.Len = ?

	x.node_state = (uint32_t)(x.ref157a0a01.node_state)
	hxfbcfcc4 := (*sliceHeader)(unsafe.Pointer(&x.reason))
	hxfbcfcc4.Data = uintptr(unsafe.Pointer(x.ref157a0a01.reason))
	hxfbcfcc4.Cap = 0x7fffffff
	// hxfbcfcc4.Len = ?

	x.reason_uid = (uint32_t)(x.ref157a0a01.reason_uid)
}

// allocUpdate_part_msg_tMemory allocates memory for type C.update_part_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUpdate_part_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUpdate_part_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfUpdate_part_msg_tValue = unsafe.Sizeof([1]C.update_part_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *update_part_msg_t) Ref() *C.update_part_msg_t {
	if x == nil {
		return nil
	}
	return x.reff7d95771
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *update_part_msg_t) Free() {
	if x != nil && x.allocsf7d95771 != nil {
		x.allocsf7d95771.(*cgoAllocMap).Free()
		x.reff7d95771 = nil
	}
}

// Newupdate_part_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newupdate_part_msg_tRef(ref unsafe.Pointer) *update_part_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(update_part_msg_t)
	obj.reff7d95771 = (*C.update_part_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *update_part_msg_t) PassRef() (*C.update_part_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff7d95771 != nil {
		return x.reff7d95771, nil
	}
	memf7d95771 := allocUpdate_part_msg_tMemory(1)
	reff7d95771 := (*C.update_part_msg_t)(memf7d95771)
	allocsf7d95771 := new(cgoAllocMap)
	allocsf7d95771.Add(memf7d95771)

	var callow_alloc_nodes_allocs *cgoAllocMap
	reff7d95771.allow_alloc_nodes, callow_alloc_nodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_alloc_nodes)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(callow_alloc_nodes_allocs)

	var callow_accounts_allocs *cgoAllocMap
	reff7d95771.allow_accounts, callow_accounts_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_accounts)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(callow_accounts_allocs)

	var callow_groups_allocs *cgoAllocMap
	reff7d95771.allow_groups, callow_groups_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_groups)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(callow_groups_allocs)

	var callow_qos_allocs *cgoAllocMap
	reff7d95771.allow_qos, callow_qos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_qos)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(callow_qos_allocs)

	var calternate_allocs *cgoAllocMap
	reff7d95771.alternate, calternate_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.alternate)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(calternate_allocs)

	var cbilling_weights_str_allocs *cgoAllocMap
	reff7d95771.billing_weights_str, cbilling_weights_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.billing_weights_str)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(cbilling_weights_str_allocs)

	var ccluster_name_allocs *cgoAllocMap
	reff7d95771.cluster_name, ccluster_name_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.cluster_name)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(ccluster_name_allocs)

	var ccr_type_allocs *cgoAllocMap
	reff7d95771.cr_type, ccr_type_allocs = (C.uint16_t)(x.cr_type), cgoAllocsUnknown
	allocsf7d95771.Borrow(ccr_type_allocs)

	var cdef_mem_per_cpu_allocs *cgoAllocMap
	reff7d95771.def_mem_per_cpu, cdef_mem_per_cpu_allocs = (C.uint64_t)(x.def_mem_per_cpu), cgoAllocsUnknown
	allocsf7d95771.Borrow(cdef_mem_per_cpu_allocs)

	var cdefault_time_allocs *cgoAllocMap
	reff7d95771.default_time, cdefault_time_allocs = (C.uint32_t)(x.default_time), cgoAllocsUnknown
	allocsf7d95771.Borrow(cdefault_time_allocs)

	var cdeny_accounts_allocs *cgoAllocMap
	reff7d95771.deny_accounts, cdeny_accounts_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.deny_accounts)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(cdeny_accounts_allocs)

	var cdeny_qos_allocs *cgoAllocMap
	reff7d95771.deny_qos, cdeny_qos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.deny_qos)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(cdeny_qos_allocs)

	var cflags_allocs *cgoAllocMap
	reff7d95771.flags, cflags_allocs = (C.uint16_t)(x.flags), cgoAllocsUnknown
	allocsf7d95771.Borrow(cflags_allocs)

	var cgrace_time_allocs *cgoAllocMap
	reff7d95771.grace_time, cgrace_time_allocs = (C.uint32_t)(x.grace_time), cgoAllocsUnknown
	allocsf7d95771.Borrow(cgrace_time_allocs)

	var cmax_cpus_per_node_allocs *cgoAllocMap
	reff7d95771.max_cpus_per_node, cmax_cpus_per_node_allocs = (C.uint32_t)(x.max_cpus_per_node), cgoAllocsUnknown
	allocsf7d95771.Borrow(cmax_cpus_per_node_allocs)

	var cmax_mem_per_cpu_allocs *cgoAllocMap
	reff7d95771.max_mem_per_cpu, cmax_mem_per_cpu_allocs = (C.uint64_t)(x.max_mem_per_cpu), cgoAllocsUnknown
	allocsf7d95771.Borrow(cmax_mem_per_cpu_allocs)

	var cmax_nodes_allocs *cgoAllocMap
	reff7d95771.max_nodes, cmax_nodes_allocs = (C.uint32_t)(x.max_nodes), cgoAllocsUnknown
	allocsf7d95771.Borrow(cmax_nodes_allocs)

	var cmax_share_allocs *cgoAllocMap
	reff7d95771.max_share, cmax_share_allocs = (C.uint16_t)(x.max_share), cgoAllocsUnknown
	allocsf7d95771.Borrow(cmax_share_allocs)

	var cmax_time_allocs *cgoAllocMap
	reff7d95771.max_time, cmax_time_allocs = (C.uint32_t)(x.max_time), cgoAllocsUnknown
	allocsf7d95771.Borrow(cmax_time_allocs)

	var cmin_nodes_allocs *cgoAllocMap
	reff7d95771.min_nodes, cmin_nodes_allocs = (C.uint32_t)(x.min_nodes), cgoAllocsUnknown
	allocsf7d95771.Borrow(cmin_nodes_allocs)

	var cname_allocs *cgoAllocMap
	reff7d95771.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(cname_allocs)

	var cnode_inx_allocs *cgoAllocMap
	reff7d95771.node_inx, cnode_inx_allocs = (*C.int32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_inx)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(cnode_inx_allocs)

	var cnodes_allocs *cgoAllocMap
	reff7d95771.nodes, cnodes_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.nodes)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(cnodes_allocs)

	var cover_time_limit_allocs *cgoAllocMap
	reff7d95771.over_time_limit, cover_time_limit_allocs = (C.uint16_t)(x.over_time_limit), cgoAllocsUnknown
	allocsf7d95771.Borrow(cover_time_limit_allocs)

	var cpreempt_mode_allocs *cgoAllocMap
	reff7d95771.preempt_mode, cpreempt_mode_allocs = (C.uint16_t)(x.preempt_mode), cgoAllocsUnknown
	allocsf7d95771.Borrow(cpreempt_mode_allocs)

	var cpriority_job_factor_allocs *cgoAllocMap
	reff7d95771.priority_job_factor, cpriority_job_factor_allocs = (C.uint16_t)(x.priority_job_factor), cgoAllocsUnknown
	allocsf7d95771.Borrow(cpriority_job_factor_allocs)

	var cpriority_tier_allocs *cgoAllocMap
	reff7d95771.priority_tier, cpriority_tier_allocs = (C.uint16_t)(x.priority_tier), cgoAllocsUnknown
	allocsf7d95771.Borrow(cpriority_tier_allocs)

	var cqos_char_allocs *cgoAllocMap
	reff7d95771.qos_char, cqos_char_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.qos_char)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(cqos_char_allocs)

	var cstate_up_allocs *cgoAllocMap
	reff7d95771.state_up, cstate_up_allocs = (C.uint16_t)(x.state_up), cgoAllocsUnknown
	allocsf7d95771.Borrow(cstate_up_allocs)

	var ctotal_cpus_allocs *cgoAllocMap
	reff7d95771.total_cpus, ctotal_cpus_allocs = (C.uint32_t)(x.total_cpus), cgoAllocsUnknown
	allocsf7d95771.Borrow(ctotal_cpus_allocs)

	var ctotal_nodes_allocs *cgoAllocMap
	reff7d95771.total_nodes, ctotal_nodes_allocs = (C.uint32_t)(x.total_nodes), cgoAllocsUnknown
	allocsf7d95771.Borrow(ctotal_nodes_allocs)

	var ctres_fmt_str_allocs *cgoAllocMap
	reff7d95771.tres_fmt_str, ctres_fmt_str_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.tres_fmt_str)).Data)), cgoAllocsUnknown
	allocsf7d95771.Borrow(ctres_fmt_str_allocs)

	x.reff7d95771 = reff7d95771
	x.allocsf7d95771 = allocsf7d95771
	return reff7d95771, allocsf7d95771

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x update_part_msg_t) PassValue() (C.update_part_msg_t, *cgoAllocMap) {
	if x.reff7d95771 != nil {
		return *x.reff7d95771, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *update_part_msg_t) Deref() {
	if x.reff7d95771 == nil {
		return
	}
	hxf44eac1 := (*sliceHeader)(unsafe.Pointer(&x.allow_alloc_nodes))
	hxf44eac1.Data = uintptr(unsafe.Pointer(x.reff7d95771.allow_alloc_nodes))
	hxf44eac1.Cap = 0x7fffffff
	// hxf44eac1.Len = ?

	hxf92e286 := (*sliceHeader)(unsafe.Pointer(&x.allow_accounts))
	hxf92e286.Data = uintptr(unsafe.Pointer(x.reff7d95771.allow_accounts))
	hxf92e286.Cap = 0x7fffffff
	// hxf92e286.Len = ?

	hxf9e348a := (*sliceHeader)(unsafe.Pointer(&x.allow_groups))
	hxf9e348a.Data = uintptr(unsafe.Pointer(x.reff7d95771.allow_groups))
	hxf9e348a.Cap = 0x7fffffff
	// hxf9e348a.Len = ?

	hxf38e026 := (*sliceHeader)(unsafe.Pointer(&x.allow_qos))
	hxf38e026.Data = uintptr(unsafe.Pointer(x.reff7d95771.allow_qos))
	hxf38e026.Cap = 0x7fffffff
	// hxf38e026.Len = ?

	hxf2fd639 := (*sliceHeader)(unsafe.Pointer(&x.alternate))
	hxf2fd639.Data = uintptr(unsafe.Pointer(x.reff7d95771.alternate))
	hxf2fd639.Cap = 0x7fffffff
	// hxf2fd639.Len = ?

	hxf9ec8b8 := (*sliceHeader)(unsafe.Pointer(&x.billing_weights_str))
	hxf9ec8b8.Data = uintptr(unsafe.Pointer(x.reff7d95771.billing_weights_str))
	hxf9ec8b8.Cap = 0x7fffffff
	// hxf9ec8b8.Len = ?

	hxf38af53 := (*sliceHeader)(unsafe.Pointer(&x.cluster_name))
	hxf38af53.Data = uintptr(unsafe.Pointer(x.reff7d95771.cluster_name))
	hxf38af53.Cap = 0x7fffffff
	// hxf38af53.Len = ?

	x.cr_type = (uint16_t)(x.reff7d95771.cr_type)
	x.def_mem_per_cpu = (uint64_t)(x.reff7d95771.def_mem_per_cpu)
	x.default_time = (uint32_t)(x.reff7d95771.default_time)
	hxfa7a3b1 := (*sliceHeader)(unsafe.Pointer(&x.deny_accounts))
	hxfa7a3b1.Data = uintptr(unsafe.Pointer(x.reff7d95771.deny_accounts))
	hxfa7a3b1.Cap = 0x7fffffff
	// hxfa7a3b1.Len = ?

	hxf1bdf9c := (*sliceHeader)(unsafe.Pointer(&x.deny_qos))
	hxf1bdf9c.Data = uintptr(unsafe.Pointer(x.reff7d95771.deny_qos))
	hxf1bdf9c.Cap = 0x7fffffff
	// hxf1bdf9c.Len = ?

	x.flags = (uint16_t)(x.reff7d95771.flags)
	x.grace_time = (uint32_t)(x.reff7d95771.grace_time)
	x.max_cpus_per_node = (uint32_t)(x.reff7d95771.max_cpus_per_node)
	x.max_mem_per_cpu = (uint64_t)(x.reff7d95771.max_mem_per_cpu)
	x.max_nodes = (uint32_t)(x.reff7d95771.max_nodes)
	x.max_share = (uint16_t)(x.reff7d95771.max_share)
	x.max_time = (uint32_t)(x.reff7d95771.max_time)
	x.min_nodes = (uint32_t)(x.reff7d95771.min_nodes)
	hxf85affd := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf85affd.Data = uintptr(unsafe.Pointer(x.reff7d95771.name))
	hxf85affd.Cap = 0x7fffffff
	// hxf85affd.Len = ?

	hxf1e2dd4 := (*sliceHeader)(unsafe.Pointer(&x.node_inx))
	hxf1e2dd4.Data = uintptr(unsafe.Pointer(x.reff7d95771.node_inx))
	hxf1e2dd4.Cap = 0x7fffffff
	// hxf1e2dd4.Len = ?

	hxf8cc25f := (*sliceHeader)(unsafe.Pointer(&x.nodes))
	hxf8cc25f.Data = uintptr(unsafe.Pointer(x.reff7d95771.nodes))
	hxf8cc25f.Cap = 0x7fffffff
	// hxf8cc25f.Len = ?

	x.over_time_limit = (uint16_t)(x.reff7d95771.over_time_limit)
	x.preempt_mode = (uint16_t)(x.reff7d95771.preempt_mode)
	x.priority_job_factor = (uint16_t)(x.reff7d95771.priority_job_factor)
	x.priority_tier = (uint16_t)(x.reff7d95771.priority_tier)
	hxfc229de := (*sliceHeader)(unsafe.Pointer(&x.qos_char))
	hxfc229de.Data = uintptr(unsafe.Pointer(x.reff7d95771.qos_char))
	hxfc229de.Cap = 0x7fffffff
	// hxfc229de.Len = ?

	x.state_up = (uint16_t)(x.reff7d95771.state_up)
	x.total_cpus = (uint32_t)(x.reff7d95771.total_cpus)
	x.total_nodes = (uint32_t)(x.reff7d95771.total_nodes)
	hxfced0ea := (*sliceHeader)(unsafe.Pointer(&x.tres_fmt_str))
	hxfced0ea.Data = uintptr(unsafe.Pointer(x.reff7d95771.tres_fmt_str))
	hxfced0ea.Cap = 0x7fffffff
	// hxfced0ea.Len = ?

}

// allocJob_sbcast_cred_msg_tMemory allocates memory for type C.job_sbcast_cred_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_sbcast_cred_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_sbcast_cred_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_sbcast_cred_msg_tValue = unsafe.Sizeof([1]C.job_sbcast_cred_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_sbcast_cred_msg_t) Ref() *C.job_sbcast_cred_msg_t {
	if x == nil {
		return nil
	}
	return x.refcd77303
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_sbcast_cred_msg_t) Free() {
	if x != nil && x.allocscd77303 != nil {
		x.allocscd77303.(*cgoAllocMap).Free()
		x.refcd77303 = nil
	}
}

// Newjob_sbcast_cred_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_sbcast_cred_msg_tRef(ref unsafe.Pointer) *job_sbcast_cred_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(job_sbcast_cred_msg_t)
	obj.refcd77303 = (*C.job_sbcast_cred_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_sbcast_cred_msg_t) PassRef() (*C.job_sbcast_cred_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refcd77303 != nil {
		return x.refcd77303, nil
	}
	memcd77303 := allocJob_sbcast_cred_msg_tMemory(1)
	refcd77303 := (*C.job_sbcast_cred_msg_t)(memcd77303)
	allocscd77303 := new(cgoAllocMap)
	allocscd77303.Add(memcd77303)

	var cjob_id_allocs *cgoAllocMap
	refcd77303.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocscd77303.Borrow(cjob_id_allocs)

	var cnode_addr_allocs *cgoAllocMap
	refcd77303.node_addr, cnode_addr_allocs = unpackSSlurm_addr_t(x.node_addr)
	allocscd77303.Borrow(cnode_addr_allocs)

	var cnode_cnt_allocs *cgoAllocMap
	refcd77303.node_cnt, cnode_cnt_allocs = (C.uint32_t)(x.node_cnt), cgoAllocsUnknown
	allocscd77303.Borrow(cnode_cnt_allocs)

	var cnode_list_allocs *cgoAllocMap
	refcd77303.node_list, cnode_list_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.node_list)).Data)), cgoAllocsUnknown
	allocscd77303.Borrow(cnode_list_allocs)

	var csbcast_cred_allocs *cgoAllocMap
	refcd77303.sbcast_cred, csbcast_cred_allocs = (*C.sbcast_cred_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.sbcast_cred)).Data)), cgoAllocsUnknown
	allocscd77303.Borrow(csbcast_cred_allocs)

	x.refcd77303 = refcd77303
	x.allocscd77303 = allocscd77303
	return refcd77303, allocscd77303

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_sbcast_cred_msg_t) PassValue() (C.job_sbcast_cred_msg_t, *cgoAllocMap) {
	if x.refcd77303 != nil {
		return *x.refcd77303, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_sbcast_cred_msg_t) Deref() {
	if x.refcd77303 == nil {
		return
	}
	x.job_id = (uint32_t)(x.refcd77303.job_id)
	packSSlurm_addr_t(x.node_addr, x.refcd77303.node_addr)
	x.node_cnt = (uint32_t)(x.refcd77303.node_cnt)
	hxf870205 := (*sliceHeader)(unsafe.Pointer(&x.node_list))
	hxf870205.Data = uintptr(unsafe.Pointer(x.refcd77303.node_list))
	hxf870205.Cap = 0x7fffffff
	// hxf870205.Len = ?

	hxfa8a19a := (*sliceHeader)(unsafe.Pointer(&x.sbcast_cred))
	hxfa8a19a.Data = uintptr(unsafe.Pointer(x.refcd77303.sbcast_cred))
	hxfa8a19a.Cap = 0x7fffffff
	// hxfa8a19a.Len = ?

}

// Ref returns a reference to C object as it is.
func (x *slurm_step_ctx_t) Ref() *C.slurm_step_ctx_t {
	if x == nil {
		return nil
	}
	return (*C.slurm_step_ctx_t)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *slurm_step_ctx_t) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// Newslurm_step_ctx_tRef converts the C object reference into a raw struct reference without wrapping.
func Newslurm_step_ctx_tRef(ref unsafe.Pointer) *slurm_step_ctx_t {
	return (*slurm_step_ctx_t)(ref)
}

// Newslurm_step_ctx_t allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func Newslurm_step_ctx_t() *slurm_step_ctx_t {
	return (*slurm_step_ctx_t)(allocSlurm_step_ctx_tMemory(1))
}

// allocSlurm_step_ctx_tMemory allocates memory for type C.slurm_step_ctx_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_step_ctx_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_step_ctx_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_step_ctx_tValue = unsafe.Sizeof([1]C.slurm_step_ctx_t{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *slurm_step_ctx_t) PassRef() *C.slurm_step_ctx_t {
	if x == nil {
		x = (*slurm_step_ctx_t)(allocSlurm_step_ctx_tMemory(1))
	}
	return (*C.slurm_step_ctx_t)(unsafe.Pointer(x))
}

// allocStats_info_request_msg_tMemory allocates memory for type C.stats_info_request_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStats_info_request_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStats_info_request_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStats_info_request_msg_tValue = unsafe.Sizeof([1]C.stats_info_request_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *stats_info_request_msg_t) Ref() *C.stats_info_request_msg_t {
	if x == nil {
		return nil
	}
	return x.refc3862c55
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *stats_info_request_msg_t) Free() {
	if x != nil && x.allocsc3862c55 != nil {
		x.allocsc3862c55.(*cgoAllocMap).Free()
		x.refc3862c55 = nil
	}
}

// Newstats_info_request_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newstats_info_request_msg_tRef(ref unsafe.Pointer) *stats_info_request_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(stats_info_request_msg_t)
	obj.refc3862c55 = (*C.stats_info_request_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *stats_info_request_msg_t) PassRef() (*C.stats_info_request_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc3862c55 != nil {
		return x.refc3862c55, nil
	}
	memc3862c55 := allocStats_info_request_msg_tMemory(1)
	refc3862c55 := (*C.stats_info_request_msg_t)(memc3862c55)
	allocsc3862c55 := new(cgoAllocMap)
	allocsc3862c55.Add(memc3862c55)

	var ccommand_id_allocs *cgoAllocMap
	refc3862c55.command_id, ccommand_id_allocs = (C.uint16_t)(x.command_id), cgoAllocsUnknown
	allocsc3862c55.Borrow(ccommand_id_allocs)

	x.refc3862c55 = refc3862c55
	x.allocsc3862c55 = allocsc3862c55
	return refc3862c55, allocsc3862c55

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x stats_info_request_msg_t) PassValue() (C.stats_info_request_msg_t, *cgoAllocMap) {
	if x.refc3862c55 != nil {
		return *x.refc3862c55, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *stats_info_request_msg_t) Deref() {
	if x.refc3862c55 == nil {
		return
	}
	x.command_id = (uint16_t)(x.refc3862c55.command_id)
}

// allocStats_info_response_msg_tMemory allocates memory for type C.stats_info_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStats_info_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStats_info_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStats_info_response_msg_tValue = unsafe.Sizeof([1]C.stats_info_response_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *stats_info_response_msg_t) Ref() *C.stats_info_response_msg_t {
	if x == nil {
		return nil
	}
	return x.ref5aae8b04
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *stats_info_response_msg_t) Free() {
	if x != nil && x.allocs5aae8b04 != nil {
		x.allocs5aae8b04.(*cgoAllocMap).Free()
		x.ref5aae8b04 = nil
	}
}

// Newstats_info_response_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newstats_info_response_msg_tRef(ref unsafe.Pointer) *stats_info_response_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(stats_info_response_msg_t)
	obj.ref5aae8b04 = (*C.stats_info_response_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *stats_info_response_msg_t) PassRef() (*C.stats_info_response_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5aae8b04 != nil {
		return x.ref5aae8b04, nil
	}
	mem5aae8b04 := allocStats_info_response_msg_tMemory(1)
	ref5aae8b04 := (*C.stats_info_response_msg_t)(mem5aae8b04)
	allocs5aae8b04 := new(cgoAllocMap)
	allocs5aae8b04.Add(mem5aae8b04)

	var cparts_packed_allocs *cgoAllocMap
	ref5aae8b04.parts_packed, cparts_packed_allocs = (C.uint32_t)(x.parts_packed), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cparts_packed_allocs)

	var creq_time_allocs *cgoAllocMap
	ref5aae8b04.req_time, creq_time_allocs = (C.time_t)(x.req_time), cgoAllocsUnknown
	allocs5aae8b04.Borrow(creq_time_allocs)

	var creq_time_start_allocs *cgoAllocMap
	ref5aae8b04.req_time_start, creq_time_start_allocs = (C.time_t)(x.req_time_start), cgoAllocsUnknown
	allocs5aae8b04.Borrow(creq_time_start_allocs)

	var cserver_thread_count_allocs *cgoAllocMap
	ref5aae8b04.server_thread_count, cserver_thread_count_allocs = (C.uint32_t)(x.server_thread_count), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cserver_thread_count_allocs)

	var cagent_queue_size_allocs *cgoAllocMap
	ref5aae8b04.agent_queue_size, cagent_queue_size_allocs = (C.uint32_t)(x.agent_queue_size), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cagent_queue_size_allocs)

	var cdbd_agent_queue_size_allocs *cgoAllocMap
	ref5aae8b04.dbd_agent_queue_size, cdbd_agent_queue_size_allocs = (C.uint32_t)(x.dbd_agent_queue_size), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cdbd_agent_queue_size_allocs)

	var cschedule_cycle_max_allocs *cgoAllocMap
	ref5aae8b04.schedule_cycle_max, cschedule_cycle_max_allocs = (C.uint32_t)(x.schedule_cycle_max), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cschedule_cycle_max_allocs)

	var cschedule_cycle_last_allocs *cgoAllocMap
	ref5aae8b04.schedule_cycle_last, cschedule_cycle_last_allocs = (C.uint32_t)(x.schedule_cycle_last), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cschedule_cycle_last_allocs)

	var cschedule_cycle_sum_allocs *cgoAllocMap
	ref5aae8b04.schedule_cycle_sum, cschedule_cycle_sum_allocs = (C.uint32_t)(x.schedule_cycle_sum), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cschedule_cycle_sum_allocs)

	var cschedule_cycle_counter_allocs *cgoAllocMap
	ref5aae8b04.schedule_cycle_counter, cschedule_cycle_counter_allocs = (C.uint32_t)(x.schedule_cycle_counter), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cschedule_cycle_counter_allocs)

	var cschedule_cycle_depth_allocs *cgoAllocMap
	ref5aae8b04.schedule_cycle_depth, cschedule_cycle_depth_allocs = (C.uint32_t)(x.schedule_cycle_depth), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cschedule_cycle_depth_allocs)

	var cschedule_queue_len_allocs *cgoAllocMap
	ref5aae8b04.schedule_queue_len, cschedule_queue_len_allocs = (C.uint32_t)(x.schedule_queue_len), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cschedule_queue_len_allocs)

	var cjobs_submitted_allocs *cgoAllocMap
	ref5aae8b04.jobs_submitted, cjobs_submitted_allocs = (C.uint32_t)(x.jobs_submitted), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cjobs_submitted_allocs)

	var cjobs_started_allocs *cgoAllocMap
	ref5aae8b04.jobs_started, cjobs_started_allocs = (C.uint32_t)(x.jobs_started), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cjobs_started_allocs)

	var cjobs_completed_allocs *cgoAllocMap
	ref5aae8b04.jobs_completed, cjobs_completed_allocs = (C.uint32_t)(x.jobs_completed), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cjobs_completed_allocs)

	var cjobs_canceled_allocs *cgoAllocMap
	ref5aae8b04.jobs_canceled, cjobs_canceled_allocs = (C.uint32_t)(x.jobs_canceled), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cjobs_canceled_allocs)

	var cjobs_failed_allocs *cgoAllocMap
	ref5aae8b04.jobs_failed, cjobs_failed_allocs = (C.uint32_t)(x.jobs_failed), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cjobs_failed_allocs)

	var cjobs_running_allocs *cgoAllocMap
	ref5aae8b04.jobs_running, cjobs_running_allocs = (C.uint32_t)(x.jobs_running), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cjobs_running_allocs)

	var cjobs_running_ts_allocs *cgoAllocMap
	ref5aae8b04.jobs_running_ts, cjobs_running_ts_allocs = (C.time_t)(x.jobs_running_ts), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cjobs_running_ts_allocs)

	var cbf_backfilled_jobs_allocs *cgoAllocMap
	ref5aae8b04.bf_backfilled_jobs, cbf_backfilled_jobs_allocs = (C.uint32_t)(x.bf_backfilled_jobs), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_backfilled_jobs_allocs)

	var cbf_last_backfilled_jobs_allocs *cgoAllocMap
	ref5aae8b04.bf_last_backfilled_jobs, cbf_last_backfilled_jobs_allocs = (C.uint32_t)(x.bf_last_backfilled_jobs), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_last_backfilled_jobs_allocs)

	var cbf_backfilled_pack_jobs_allocs *cgoAllocMap
	ref5aae8b04.bf_backfilled_pack_jobs, cbf_backfilled_pack_jobs_allocs = (C.uint32_t)(x.bf_backfilled_pack_jobs), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_backfilled_pack_jobs_allocs)

	var cbf_cycle_counter_allocs *cgoAllocMap
	ref5aae8b04.bf_cycle_counter, cbf_cycle_counter_allocs = (C.uint32_t)(x.bf_cycle_counter), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_cycle_counter_allocs)

	var cbf_cycle_sum_allocs *cgoAllocMap
	ref5aae8b04.bf_cycle_sum, cbf_cycle_sum_allocs = (C.uint64_t)(x.bf_cycle_sum), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_cycle_sum_allocs)

	var cbf_cycle_last_allocs *cgoAllocMap
	ref5aae8b04.bf_cycle_last, cbf_cycle_last_allocs = (C.uint32_t)(x.bf_cycle_last), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_cycle_last_allocs)

	var cbf_cycle_max_allocs *cgoAllocMap
	ref5aae8b04.bf_cycle_max, cbf_cycle_max_allocs = (C.uint32_t)(x.bf_cycle_max), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_cycle_max_allocs)

	var cbf_last_depth_allocs *cgoAllocMap
	ref5aae8b04.bf_last_depth, cbf_last_depth_allocs = (C.uint32_t)(x.bf_last_depth), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_last_depth_allocs)

	var cbf_last_depth_try_allocs *cgoAllocMap
	ref5aae8b04.bf_last_depth_try, cbf_last_depth_try_allocs = (C.uint32_t)(x.bf_last_depth_try), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_last_depth_try_allocs)

	var cbf_depth_sum_allocs *cgoAllocMap
	ref5aae8b04.bf_depth_sum, cbf_depth_sum_allocs = (C.uint32_t)(x.bf_depth_sum), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_depth_sum_allocs)

	var cbf_depth_try_sum_allocs *cgoAllocMap
	ref5aae8b04.bf_depth_try_sum, cbf_depth_try_sum_allocs = (C.uint32_t)(x.bf_depth_try_sum), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_depth_try_sum_allocs)

	var cbf_queue_len_allocs *cgoAllocMap
	ref5aae8b04.bf_queue_len, cbf_queue_len_allocs = (C.uint32_t)(x.bf_queue_len), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_queue_len_allocs)

	var cbf_queue_len_sum_allocs *cgoAllocMap
	ref5aae8b04.bf_queue_len_sum, cbf_queue_len_sum_allocs = (C.uint32_t)(x.bf_queue_len_sum), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_queue_len_sum_allocs)

	var cbf_when_last_cycle_allocs *cgoAllocMap
	ref5aae8b04.bf_when_last_cycle, cbf_when_last_cycle_allocs = (C.time_t)(x.bf_when_last_cycle), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_when_last_cycle_allocs)

	var cbf_active_allocs *cgoAllocMap
	ref5aae8b04.bf_active, cbf_active_allocs = (C.uint32_t)(x.bf_active), cgoAllocsUnknown
	allocs5aae8b04.Borrow(cbf_active_allocs)

	var crpc_type_size_allocs *cgoAllocMap
	ref5aae8b04.rpc_type_size, crpc_type_size_allocs = (C.uint32_t)(x.rpc_type_size), cgoAllocsUnknown
	allocs5aae8b04.Borrow(crpc_type_size_allocs)

	var crpc_type_id_allocs *cgoAllocMap
	ref5aae8b04.rpc_type_id, crpc_type_id_allocs = (*C.uint16_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.rpc_type_id)).Data)), cgoAllocsUnknown
	allocs5aae8b04.Borrow(crpc_type_id_allocs)

	var crpc_type_cnt_allocs *cgoAllocMap
	ref5aae8b04.rpc_type_cnt, crpc_type_cnt_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.rpc_type_cnt)).Data)), cgoAllocsUnknown
	allocs5aae8b04.Borrow(crpc_type_cnt_allocs)

	var crpc_type_time_allocs *cgoAllocMap
	ref5aae8b04.rpc_type_time, crpc_type_time_allocs = (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.rpc_type_time)).Data)), cgoAllocsUnknown
	allocs5aae8b04.Borrow(crpc_type_time_allocs)

	var crpc_user_size_allocs *cgoAllocMap
	ref5aae8b04.rpc_user_size, crpc_user_size_allocs = (C.uint32_t)(x.rpc_user_size), cgoAllocsUnknown
	allocs5aae8b04.Borrow(crpc_user_size_allocs)

	var crpc_user_id_allocs *cgoAllocMap
	ref5aae8b04.rpc_user_id, crpc_user_id_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.rpc_user_id)).Data)), cgoAllocsUnknown
	allocs5aae8b04.Borrow(crpc_user_id_allocs)

	var crpc_user_cnt_allocs *cgoAllocMap
	ref5aae8b04.rpc_user_cnt, crpc_user_cnt_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.rpc_user_cnt)).Data)), cgoAllocsUnknown
	allocs5aae8b04.Borrow(crpc_user_cnt_allocs)

	var crpc_user_time_allocs *cgoAllocMap
	ref5aae8b04.rpc_user_time, crpc_user_time_allocs = (*C.uint64_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.rpc_user_time)).Data)), cgoAllocsUnknown
	allocs5aae8b04.Borrow(crpc_user_time_allocs)

	x.ref5aae8b04 = ref5aae8b04
	x.allocs5aae8b04 = allocs5aae8b04
	return ref5aae8b04, allocs5aae8b04

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x stats_info_response_msg_t) PassValue() (C.stats_info_response_msg_t, *cgoAllocMap) {
	if x.ref5aae8b04 != nil {
		return *x.ref5aae8b04, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *stats_info_response_msg_t) Deref() {
	if x.ref5aae8b04 == nil {
		return
	}
	x.parts_packed = (uint32_t)(x.ref5aae8b04.parts_packed)
	x.req_time = (time_t)(x.ref5aae8b04.req_time)
	x.req_time_start = (time_t)(x.ref5aae8b04.req_time_start)
	x.server_thread_count = (uint32_t)(x.ref5aae8b04.server_thread_count)
	x.agent_queue_size = (uint32_t)(x.ref5aae8b04.agent_queue_size)
	x.dbd_agent_queue_size = (uint32_t)(x.ref5aae8b04.dbd_agent_queue_size)
	x.schedule_cycle_max = (uint32_t)(x.ref5aae8b04.schedule_cycle_max)
	x.schedule_cycle_last = (uint32_t)(x.ref5aae8b04.schedule_cycle_last)
	x.schedule_cycle_sum = (uint32_t)(x.ref5aae8b04.schedule_cycle_sum)
	x.schedule_cycle_counter = (uint32_t)(x.ref5aae8b04.schedule_cycle_counter)
	x.schedule_cycle_depth = (uint32_t)(x.ref5aae8b04.schedule_cycle_depth)
	x.schedule_queue_len = (uint32_t)(x.ref5aae8b04.schedule_queue_len)
	x.jobs_submitted = (uint32_t)(x.ref5aae8b04.jobs_submitted)
	x.jobs_started = (uint32_t)(x.ref5aae8b04.jobs_started)
	x.jobs_completed = (uint32_t)(x.ref5aae8b04.jobs_completed)
	x.jobs_canceled = (uint32_t)(x.ref5aae8b04.jobs_canceled)
	x.jobs_failed = (uint32_t)(x.ref5aae8b04.jobs_failed)
	x.jobs_running = (uint32_t)(x.ref5aae8b04.jobs_running)
	x.jobs_running_ts = (time_t)(x.ref5aae8b04.jobs_running_ts)
	x.bf_backfilled_jobs = (uint32_t)(x.ref5aae8b04.bf_backfilled_jobs)
	x.bf_last_backfilled_jobs = (uint32_t)(x.ref5aae8b04.bf_last_backfilled_jobs)
	x.bf_backfilled_pack_jobs = (uint32_t)(x.ref5aae8b04.bf_backfilled_pack_jobs)
	x.bf_cycle_counter = (uint32_t)(x.ref5aae8b04.bf_cycle_counter)
	x.bf_cycle_sum = (uint64_t)(x.ref5aae8b04.bf_cycle_sum)
	x.bf_cycle_last = (uint32_t)(x.ref5aae8b04.bf_cycle_last)
	x.bf_cycle_max = (uint32_t)(x.ref5aae8b04.bf_cycle_max)
	x.bf_last_depth = (uint32_t)(x.ref5aae8b04.bf_last_depth)
	x.bf_last_depth_try = (uint32_t)(x.ref5aae8b04.bf_last_depth_try)
	x.bf_depth_sum = (uint32_t)(x.ref5aae8b04.bf_depth_sum)
	x.bf_depth_try_sum = (uint32_t)(x.ref5aae8b04.bf_depth_try_sum)
	x.bf_queue_len = (uint32_t)(x.ref5aae8b04.bf_queue_len)
	x.bf_queue_len_sum = (uint32_t)(x.ref5aae8b04.bf_queue_len_sum)
	x.bf_when_last_cycle = (time_t)(x.ref5aae8b04.bf_when_last_cycle)
	x.bf_active = (uint32_t)(x.ref5aae8b04.bf_active)
	x.rpc_type_size = (uint32_t)(x.ref5aae8b04.rpc_type_size)
	hxf7628ab := (*sliceHeader)(unsafe.Pointer(&x.rpc_type_id))
	hxf7628ab.Data = uintptr(unsafe.Pointer(x.ref5aae8b04.rpc_type_id))
	hxf7628ab.Cap = 0x7fffffff
	// hxf7628ab.Len = ?

	hxfbfd8e8 := (*sliceHeader)(unsafe.Pointer(&x.rpc_type_cnt))
	hxfbfd8e8.Data = uintptr(unsafe.Pointer(x.ref5aae8b04.rpc_type_cnt))
	hxfbfd8e8.Cap = 0x7fffffff
	// hxfbfd8e8.Len = ?

	hxf0bc7e8 := (*sliceHeader)(unsafe.Pointer(&x.rpc_type_time))
	hxf0bc7e8.Data = uintptr(unsafe.Pointer(x.ref5aae8b04.rpc_type_time))
	hxf0bc7e8.Cap = 0x7fffffff
	// hxf0bc7e8.Len = ?

	x.rpc_user_size = (uint32_t)(x.ref5aae8b04.rpc_user_size)
	hxf6bce8d := (*sliceHeader)(unsafe.Pointer(&x.rpc_user_id))
	hxf6bce8d.Data = uintptr(unsafe.Pointer(x.ref5aae8b04.rpc_user_id))
	hxf6bce8d.Cap = 0x7fffffff
	// hxf6bce8d.Len = ?

	hxf83dbdd := (*sliceHeader)(unsafe.Pointer(&x.rpc_user_cnt))
	hxf83dbdd.Data = uintptr(unsafe.Pointer(x.ref5aae8b04.rpc_user_cnt))
	hxf83dbdd.Cap = 0x7fffffff
	// hxf83dbdd.Len = ?

	hxffee363 := (*sliceHeader)(unsafe.Pointer(&x.rpc_user_time))
	hxffee363.Data = uintptr(unsafe.Pointer(x.ref5aae8b04.rpc_user_time))
	hxffee363.Cap = 0x7fffffff
	// hxffee363.Len = ?

}

// allocTrigger_info_tMemory allocates memory for type C.trigger_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocTrigger_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfTrigger_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfTrigger_info_tValue = unsafe.Sizeof([1]C.trigger_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *trigger_info_t) Ref() *C.trigger_info_t {
	if x == nil {
		return nil
	}
	return x.ref1497cae7
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *trigger_info_t) Free() {
	if x != nil && x.allocs1497cae7 != nil {
		x.allocs1497cae7.(*cgoAllocMap).Free()
		x.ref1497cae7 = nil
	}
}

// Newtrigger_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newtrigger_info_tRef(ref unsafe.Pointer) *trigger_info_t {
	if ref == nil {
		return nil
	}
	obj := new(trigger_info_t)
	obj.ref1497cae7 = (*C.trigger_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *trigger_info_t) PassRef() (*C.trigger_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref1497cae7 != nil {
		return x.ref1497cae7, nil
	}
	mem1497cae7 := allocTrigger_info_tMemory(1)
	ref1497cae7 := (*C.trigger_info_t)(mem1497cae7)
	allocs1497cae7 := new(cgoAllocMap)
	allocs1497cae7.Add(mem1497cae7)

	var cflags_allocs *cgoAllocMap
	ref1497cae7.flags, cflags_allocs = (C.uint16_t)(x.flags), cgoAllocsUnknown
	allocs1497cae7.Borrow(cflags_allocs)

	var ctrig_id_allocs *cgoAllocMap
	ref1497cae7.trig_id, ctrig_id_allocs = (C.uint32_t)(x.trig_id), cgoAllocsUnknown
	allocs1497cae7.Borrow(ctrig_id_allocs)

	var cres_type_allocs *cgoAllocMap
	ref1497cae7.res_type, cres_type_allocs = (C.uint16_t)(x.res_type), cgoAllocsUnknown
	allocs1497cae7.Borrow(cres_type_allocs)

	var cres_id_allocs *cgoAllocMap
	ref1497cae7.res_id, cres_id_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.res_id)).Data)), cgoAllocsUnknown
	allocs1497cae7.Borrow(cres_id_allocs)

	var ctrig_type_allocs *cgoAllocMap
	ref1497cae7.trig_type, ctrig_type_allocs = (C.uint32_t)(x.trig_type), cgoAllocsUnknown
	allocs1497cae7.Borrow(ctrig_type_allocs)

	var coffset_allocs *cgoAllocMap
	ref1497cae7.offset, coffset_allocs = (C.uint16_t)(x.offset), cgoAllocsUnknown
	allocs1497cae7.Borrow(coffset_allocs)

	var cuser_id_allocs *cgoAllocMap
	ref1497cae7.user_id, cuser_id_allocs = (C.uint32_t)(x.user_id), cgoAllocsUnknown
	allocs1497cae7.Borrow(cuser_id_allocs)

	var cprogram_allocs *cgoAllocMap
	ref1497cae7.program, cprogram_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.program)).Data)), cgoAllocsUnknown
	allocs1497cae7.Borrow(cprogram_allocs)

	x.ref1497cae7 = ref1497cae7
	x.allocs1497cae7 = allocs1497cae7
	return ref1497cae7, allocs1497cae7

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x trigger_info_t) PassValue() (C.trigger_info_t, *cgoAllocMap) {
	if x.ref1497cae7 != nil {
		return *x.ref1497cae7, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *trigger_info_t) Deref() {
	if x.ref1497cae7 == nil {
		return
	}
	x.flags = (uint16_t)(x.ref1497cae7.flags)
	x.trig_id = (uint32_t)(x.ref1497cae7.trig_id)
	x.res_type = (uint16_t)(x.ref1497cae7.res_type)
	hxf077a77 := (*sliceHeader)(unsafe.Pointer(&x.res_id))
	hxf077a77.Data = uintptr(unsafe.Pointer(x.ref1497cae7.res_id))
	hxf077a77.Cap = 0x7fffffff
	// hxf077a77.Len = ?

	x.trig_type = (uint32_t)(x.ref1497cae7.trig_type)
	x.offset = (uint16_t)(x.ref1497cae7.offset)
	x.user_id = (uint32_t)(x.ref1497cae7.user_id)
	hxf5db4ce := (*sliceHeader)(unsafe.Pointer(&x.program))
	hxf5db4ce.Data = uintptr(unsafe.Pointer(x.ref1497cae7.program))
	hxf5db4ce.Cap = 0x7fffffff
	// hxf5db4ce.Len = ?

}

// allocTrigger_info_msg_tMemory allocates memory for type C.trigger_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocTrigger_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfTrigger_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfTrigger_info_msg_tValue = unsafe.Sizeof([1]C.trigger_info_msg_t{})

// unpackSTrigger_info_t transforms a sliced Go data structure into plain C format.
func unpackSTrigger_info_t(x []trigger_info_t) (unpacked *C.trigger_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.trigger_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocTrigger_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.trigger_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.trigger_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSTrigger_info_t reads sliced Go data structure out from plain C format.
func packSTrigger_info_t(v []trigger_info_t, ptr0 *C.trigger_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfTrigger_info_tValue]C.trigger_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newtrigger_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *trigger_info_msg_t) Ref() *C.trigger_info_msg_t {
	if x == nil {
		return nil
	}
	return x.refe1a7ae31
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *trigger_info_msg_t) Free() {
	if x != nil && x.allocse1a7ae31 != nil {
		x.allocse1a7ae31.(*cgoAllocMap).Free()
		x.refe1a7ae31 = nil
	}
}

// Newtrigger_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newtrigger_info_msg_tRef(ref unsafe.Pointer) *trigger_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(trigger_info_msg_t)
	obj.refe1a7ae31 = (*C.trigger_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *trigger_info_msg_t) PassRef() (*C.trigger_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refe1a7ae31 != nil {
		return x.refe1a7ae31, nil
	}
	meme1a7ae31 := allocTrigger_info_msg_tMemory(1)
	refe1a7ae31 := (*C.trigger_info_msg_t)(meme1a7ae31)
	allocse1a7ae31 := new(cgoAllocMap)
	allocse1a7ae31.Add(meme1a7ae31)

	var crecord_count_allocs *cgoAllocMap
	refe1a7ae31.record_count, crecord_count_allocs = (C.uint32_t)(x.record_count), cgoAllocsUnknown
	allocse1a7ae31.Borrow(crecord_count_allocs)

	var ctrigger_array_allocs *cgoAllocMap
	refe1a7ae31.trigger_array, ctrigger_array_allocs = unpackSTrigger_info_t(x.trigger_array)
	allocse1a7ae31.Borrow(ctrigger_array_allocs)

	x.refe1a7ae31 = refe1a7ae31
	x.allocse1a7ae31 = allocse1a7ae31
	return refe1a7ae31, allocse1a7ae31

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x trigger_info_msg_t) PassValue() (C.trigger_info_msg_t, *cgoAllocMap) {
	if x.refe1a7ae31 != nil {
		return *x.refe1a7ae31, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *trigger_info_msg_t) Deref() {
	if x.refe1a7ae31 == nil {
		return
	}
	x.record_count = (uint32_t)(x.refe1a7ae31.record_count)
	packSTrigger_info_t(x.trigger_array, x.refe1a7ae31.trigger_array)
}

// allocSlurm_license_info_tMemory allocates memory for type C.slurm_license_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSlurm_license_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSlurm_license_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSlurm_license_info_tValue = unsafe.Sizeof([1]C.slurm_license_info_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *slurm_license_info_t) Ref() *C.slurm_license_info_t {
	if x == nil {
		return nil
	}
	return x.refb718b5ee
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *slurm_license_info_t) Free() {
	if x != nil && x.allocsb718b5ee != nil {
		x.allocsb718b5ee.(*cgoAllocMap).Free()
		x.refb718b5ee = nil
	}
}

// Newslurm_license_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newslurm_license_info_tRef(ref unsafe.Pointer) *slurm_license_info_t {
	if ref == nil {
		return nil
	}
	obj := new(slurm_license_info_t)
	obj.refb718b5ee = (*C.slurm_license_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *slurm_license_info_t) PassRef() (*C.slurm_license_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb718b5ee != nil {
		return x.refb718b5ee, nil
	}
	memb718b5ee := allocSlurm_license_info_tMemory(1)
	refb718b5ee := (*C.slurm_license_info_t)(memb718b5ee)
	allocsb718b5ee := new(cgoAllocMap)
	allocsb718b5ee.Add(memb718b5ee)

	var cname_allocs *cgoAllocMap
	refb718b5ee.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocsb718b5ee.Borrow(cname_allocs)

	var ctotal_allocs *cgoAllocMap
	refb718b5ee.total, ctotal_allocs = (C.uint32_t)(x.total), cgoAllocsUnknown
	allocsb718b5ee.Borrow(ctotal_allocs)

	var cin_use_allocs *cgoAllocMap
	refb718b5ee.in_use, cin_use_allocs = (C.uint32_t)(x.in_use), cgoAllocsUnknown
	allocsb718b5ee.Borrow(cin_use_allocs)

	var cavailable_allocs *cgoAllocMap
	refb718b5ee.available, cavailable_allocs = (C.uint32_t)(x.available), cgoAllocsUnknown
	allocsb718b5ee.Borrow(cavailable_allocs)

	var cremote_allocs *cgoAllocMap
	refb718b5ee.remote, cremote_allocs = (C.uint8_t)(x.remote), cgoAllocsUnknown
	allocsb718b5ee.Borrow(cremote_allocs)

	x.refb718b5ee = refb718b5ee
	x.allocsb718b5ee = allocsb718b5ee
	return refb718b5ee, allocsb718b5ee

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x slurm_license_info_t) PassValue() (C.slurm_license_info_t, *cgoAllocMap) {
	if x.refb718b5ee != nil {
		return *x.refb718b5ee, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *slurm_license_info_t) Deref() {
	if x.refb718b5ee == nil {
		return
	}
	hxf9094b2 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf9094b2.Data = uintptr(unsafe.Pointer(x.refb718b5ee.name))
	hxf9094b2.Cap = 0x7fffffff
	// hxf9094b2.Len = ?

	x.total = (uint32_t)(x.refb718b5ee.total)
	x.in_use = (uint32_t)(x.refb718b5ee.in_use)
	x.available = (uint32_t)(x.refb718b5ee.available)
	x.remote = (uint8_t)(x.refb718b5ee.remote)
}

// allocLicense_info_msg_tMemory allocates memory for type C.license_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocLicense_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfLicense_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfLicense_info_msg_tValue = unsafe.Sizeof([1]C.license_info_msg_t{})

// unpackSSlurm_license_info_t transforms a sliced Go data structure into plain C format.
func unpackSSlurm_license_info_t(x []slurm_license_info_t) (unpacked *C.slurm_license_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurm_license_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurm_license_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurm_license_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurm_license_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSSlurm_license_info_t reads sliced Go data structure out from plain C format.
func packSSlurm_license_info_t(v []slurm_license_info_t, ptr0 *C.slurm_license_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSlurm_license_info_tValue]C.slurm_license_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newslurm_license_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *license_info_msg_t) Ref() *C.license_info_msg_t {
	if x == nil {
		return nil
	}
	return x.refe48f58cb
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *license_info_msg_t) Free() {
	if x != nil && x.allocse48f58cb != nil {
		x.allocse48f58cb.(*cgoAllocMap).Free()
		x.refe48f58cb = nil
	}
}

// Newlicense_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newlicense_info_msg_tRef(ref unsafe.Pointer) *license_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(license_info_msg_t)
	obj.refe48f58cb = (*C.license_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *license_info_msg_t) PassRef() (*C.license_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refe48f58cb != nil {
		return x.refe48f58cb, nil
	}
	meme48f58cb := allocLicense_info_msg_tMemory(1)
	refe48f58cb := (*C.license_info_msg_t)(meme48f58cb)
	allocse48f58cb := new(cgoAllocMap)
	allocse48f58cb.Add(meme48f58cb)

	var clast_update_allocs *cgoAllocMap
	refe48f58cb.last_update, clast_update_allocs = (C.time_t)(x.last_update), cgoAllocsUnknown
	allocse48f58cb.Borrow(clast_update_allocs)

	var cnum_lic_allocs *cgoAllocMap
	refe48f58cb.num_lic, cnum_lic_allocs = (C.uint32_t)(x.num_lic), cgoAllocsUnknown
	allocse48f58cb.Borrow(cnum_lic_allocs)

	var clic_array_allocs *cgoAllocMap
	refe48f58cb.lic_array, clic_array_allocs = unpackSSlurm_license_info_t(x.lic_array)
	allocse48f58cb.Borrow(clic_array_allocs)

	x.refe48f58cb = refe48f58cb
	x.allocse48f58cb = allocse48f58cb
	return refe48f58cb, allocse48f58cb

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x license_info_msg_t) PassValue() (C.license_info_msg_t, *cgoAllocMap) {
	if x.refe48f58cb != nil {
		return *x.refe48f58cb, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *license_info_msg_t) Deref() {
	if x.refe48f58cb == nil {
		return
	}
	x.last_update = (time_t)(x.refe48f58cb.last_update)
	x.num_lic = (uint32_t)(x.refe48f58cb.num_lic)
	packSSlurm_license_info_t(x.lic_array, x.refe48f58cb.lic_array)
}

// allocJob_array_resp_msg_tMemory allocates memory for type C.job_array_resp_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_array_resp_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_array_resp_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_array_resp_msg_tValue = unsafe.Sizeof([1]C.job_array_resp_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_array_resp_msg_t) Ref() *C.job_array_resp_msg_t {
	if x == nil {
		return nil
	}
	return x.ref6ab3c2ed
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_array_resp_msg_t) Free() {
	if x != nil && x.allocs6ab3c2ed != nil {
		x.allocs6ab3c2ed.(*cgoAllocMap).Free()
		x.ref6ab3c2ed = nil
	}
}

// Newjob_array_resp_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_array_resp_msg_tRef(ref unsafe.Pointer) *job_array_resp_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(job_array_resp_msg_t)
	obj.ref6ab3c2ed = (*C.job_array_resp_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_array_resp_msg_t) PassRef() (*C.job_array_resp_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6ab3c2ed != nil {
		return x.ref6ab3c2ed, nil
	}
	mem6ab3c2ed := allocJob_array_resp_msg_tMemory(1)
	ref6ab3c2ed := (*C.job_array_resp_msg_t)(mem6ab3c2ed)
	allocs6ab3c2ed := new(cgoAllocMap)
	allocs6ab3c2ed.Add(mem6ab3c2ed)

	var cjob_array_count_allocs *cgoAllocMap
	ref6ab3c2ed.job_array_count, cjob_array_count_allocs = (C.uint32_t)(x.job_array_count), cgoAllocsUnknown
	allocs6ab3c2ed.Borrow(cjob_array_count_allocs)

	var cjob_array_id_allocs *cgoAllocMap
	ref6ab3c2ed.job_array_id, cjob_array_id_allocs = unpackSSByte(x.job_array_id)
	allocs6ab3c2ed.Borrow(cjob_array_id_allocs)

	var cerror_code_allocs *cgoAllocMap
	ref6ab3c2ed.error_code, cerror_code_allocs = (*C.uint32_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.error_code)).Data)), cgoAllocsUnknown
	allocs6ab3c2ed.Borrow(cerror_code_allocs)

	x.ref6ab3c2ed = ref6ab3c2ed
	x.allocs6ab3c2ed = allocs6ab3c2ed
	return ref6ab3c2ed, allocs6ab3c2ed

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_array_resp_msg_t) PassValue() (C.job_array_resp_msg_t, *cgoAllocMap) {
	if x.ref6ab3c2ed != nil {
		return *x.ref6ab3c2ed, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_array_resp_msg_t) Deref() {
	if x.ref6ab3c2ed == nil {
		return
	}
	x.job_array_count = (uint32_t)(x.ref6ab3c2ed.job_array_count)
	packSSByte(x.job_array_id, x.ref6ab3c2ed.job_array_id)
	hxfa6d8e6 := (*sliceHeader)(unsafe.Pointer(&x.error_code))
	hxfa6d8e6.Data = uintptr(unsafe.Pointer(x.ref6ab3c2ed.error_code))
	hxfa6d8e6.Cap = 0x7fffffff
	// hxfa6d8e6.Len = ?

}

// allocAssoc_mgr_info_msg_tMemory allocates memory for type C.assoc_mgr_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAssoc_mgr_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAssoc_mgr_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAssoc_mgr_info_msg_tValue = unsafe.Sizeof([1]C.assoc_mgr_info_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *assoc_mgr_info_msg_t) Ref() *C.assoc_mgr_info_msg_t {
	if x == nil {
		return nil
	}
	return x.refb9a59e94
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *assoc_mgr_info_msg_t) Free() {
	if x != nil && x.allocsb9a59e94 != nil {
		x.allocsb9a59e94.(*cgoAllocMap).Free()
		x.refb9a59e94 = nil
	}
}

// Newassoc_mgr_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newassoc_mgr_info_msg_tRef(ref unsafe.Pointer) *assoc_mgr_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(assoc_mgr_info_msg_t)
	obj.refb9a59e94 = (*C.assoc_mgr_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *assoc_mgr_info_msg_t) PassRef() (*C.assoc_mgr_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb9a59e94 != nil {
		return x.refb9a59e94, nil
	}
	memb9a59e94 := allocAssoc_mgr_info_msg_tMemory(1)
	refb9a59e94 := (*C.assoc_mgr_info_msg_t)(memb9a59e94)
	allocsb9a59e94 := new(cgoAllocMap)
	allocsb9a59e94.Add(memb9a59e94)

	var cassoc_list_allocs *cgoAllocMap
	refb9a59e94.assoc_list, cassoc_list_allocs = *(*C.List)(unsafe.Pointer(&x.assoc_list)), cgoAllocsUnknown
	allocsb9a59e94.Borrow(cassoc_list_allocs)

	var cqos_list_allocs *cgoAllocMap
	refb9a59e94.qos_list, cqos_list_allocs = *(*C.List)(unsafe.Pointer(&x.qos_list)), cgoAllocsUnknown
	allocsb9a59e94.Borrow(cqos_list_allocs)

	var ctres_cnt_allocs *cgoAllocMap
	refb9a59e94.tres_cnt, ctres_cnt_allocs = (C.uint32_t)(x.tres_cnt), cgoAllocsUnknown
	allocsb9a59e94.Borrow(ctres_cnt_allocs)

	var ctres_names_allocs *cgoAllocMap
	refb9a59e94.tres_names, ctres_names_allocs = unpackSSByte(x.tres_names)
	allocsb9a59e94.Borrow(ctres_names_allocs)

	var cuser_list_allocs *cgoAllocMap
	refb9a59e94.user_list, cuser_list_allocs = *(*C.List)(unsafe.Pointer(&x.user_list)), cgoAllocsUnknown
	allocsb9a59e94.Borrow(cuser_list_allocs)

	x.refb9a59e94 = refb9a59e94
	x.allocsb9a59e94 = allocsb9a59e94
	return refb9a59e94, allocsb9a59e94

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x assoc_mgr_info_msg_t) PassValue() (C.assoc_mgr_info_msg_t, *cgoAllocMap) {
	if x.refb9a59e94 != nil {
		return *x.refb9a59e94, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *assoc_mgr_info_msg_t) Deref() {
	if x.refb9a59e94 == nil {
		return
	}
	x.assoc_list = *(*List)(unsafe.Pointer(&x.refb9a59e94.assoc_list))
	x.qos_list = *(*List)(unsafe.Pointer(&x.refb9a59e94.qos_list))
	x.tres_cnt = (uint32_t)(x.refb9a59e94.tres_cnt)
	packSSByte(x.tres_names, x.refb9a59e94.tres_names)
	x.user_list = *(*List)(unsafe.Pointer(&x.refb9a59e94.user_list))
}

// allocAssoc_mgr_info_request_msg_tMemory allocates memory for type C.assoc_mgr_info_request_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAssoc_mgr_info_request_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAssoc_mgr_info_request_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAssoc_mgr_info_request_msg_tValue = unsafe.Sizeof([1]C.assoc_mgr_info_request_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *assoc_mgr_info_request_msg_t) Ref() *C.assoc_mgr_info_request_msg_t {
	if x == nil {
		return nil
	}
	return x.ref8173f97
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *assoc_mgr_info_request_msg_t) Free() {
	if x != nil && x.allocs8173f97 != nil {
		x.allocs8173f97.(*cgoAllocMap).Free()
		x.ref8173f97 = nil
	}
}

// Newassoc_mgr_info_request_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newassoc_mgr_info_request_msg_tRef(ref unsafe.Pointer) *assoc_mgr_info_request_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(assoc_mgr_info_request_msg_t)
	obj.ref8173f97 = (*C.assoc_mgr_info_request_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *assoc_mgr_info_request_msg_t) PassRef() (*C.assoc_mgr_info_request_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref8173f97 != nil {
		return x.ref8173f97, nil
	}
	mem8173f97 := allocAssoc_mgr_info_request_msg_tMemory(1)
	ref8173f97 := (*C.assoc_mgr_info_request_msg_t)(mem8173f97)
	allocs8173f97 := new(cgoAllocMap)
	allocs8173f97.Add(mem8173f97)

	var cacct_list_allocs *cgoAllocMap
	ref8173f97.acct_list, cacct_list_allocs = *(*C.List)(unsafe.Pointer(&x.acct_list)), cgoAllocsUnknown
	allocs8173f97.Borrow(cacct_list_allocs)

	var cflags_allocs *cgoAllocMap
	ref8173f97.flags, cflags_allocs = (C.uint32_t)(x.flags), cgoAllocsUnknown
	allocs8173f97.Borrow(cflags_allocs)

	var cqos_list_allocs *cgoAllocMap
	ref8173f97.qos_list, cqos_list_allocs = *(*C.List)(unsafe.Pointer(&x.qos_list)), cgoAllocsUnknown
	allocs8173f97.Borrow(cqos_list_allocs)

	var cuser_list_allocs *cgoAllocMap
	ref8173f97.user_list, cuser_list_allocs = *(*C.List)(unsafe.Pointer(&x.user_list)), cgoAllocsUnknown
	allocs8173f97.Borrow(cuser_list_allocs)

	x.ref8173f97 = ref8173f97
	x.allocs8173f97 = allocs8173f97
	return ref8173f97, allocs8173f97

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x assoc_mgr_info_request_msg_t) PassValue() (C.assoc_mgr_info_request_msg_t, *cgoAllocMap) {
	if x.ref8173f97 != nil {
		return *x.ref8173f97, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *assoc_mgr_info_request_msg_t) Deref() {
	if x.ref8173f97 == nil {
		return
	}
	x.acct_list = *(*List)(unsafe.Pointer(&x.ref8173f97.acct_list))
	x.flags = (uint32_t)(x.ref8173f97.flags)
	x.qos_list = *(*List)(unsafe.Pointer(&x.ref8173f97.qos_list))
	x.user_list = *(*List)(unsafe.Pointer(&x.ref8173f97.user_list))
}

// allocNetwork_callerid_msg_tMemory allocates memory for type C.network_callerid_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNetwork_callerid_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNetwork_callerid_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNetwork_callerid_msg_tValue = unsafe.Sizeof([1]C.network_callerid_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *network_callerid_msg_t) Ref() *C.network_callerid_msg_t {
	if x == nil {
		return nil
	}
	return x.refab82ee84
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *network_callerid_msg_t) Free() {
	if x != nil && x.allocsab82ee84 != nil {
		x.allocsab82ee84.(*cgoAllocMap).Free()
		x.refab82ee84 = nil
	}
}

// Newnetwork_callerid_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newnetwork_callerid_msg_tRef(ref unsafe.Pointer) *network_callerid_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(network_callerid_msg_t)
	obj.refab82ee84 = (*C.network_callerid_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *network_callerid_msg_t) PassRef() (*C.network_callerid_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refab82ee84 != nil {
		return x.refab82ee84, nil
	}
	memab82ee84 := allocNetwork_callerid_msg_tMemory(1)
	refab82ee84 := (*C.network_callerid_msg_t)(memab82ee84)
	allocsab82ee84 := new(cgoAllocMap)
	allocsab82ee84.Add(memab82ee84)

	var cip_src_allocs *cgoAllocMap
	refab82ee84.ip_src, cip_src_allocs = *(*[16]C.uchar)(unsafe.Pointer(&x.ip_src)), cgoAllocsUnknown
	allocsab82ee84.Borrow(cip_src_allocs)

	var cip_dst_allocs *cgoAllocMap
	refab82ee84.ip_dst, cip_dst_allocs = *(*[16]C.uchar)(unsafe.Pointer(&x.ip_dst)), cgoAllocsUnknown
	allocsab82ee84.Borrow(cip_dst_allocs)

	var cport_src_allocs *cgoAllocMap
	refab82ee84.port_src, cport_src_allocs = (C.uint32_t)(x.port_src), cgoAllocsUnknown
	allocsab82ee84.Borrow(cport_src_allocs)

	var cport_dst_allocs *cgoAllocMap
	refab82ee84.port_dst, cport_dst_allocs = (C.uint32_t)(x.port_dst), cgoAllocsUnknown
	allocsab82ee84.Borrow(cport_dst_allocs)

	var caf_allocs *cgoAllocMap
	refab82ee84.af, caf_allocs = (C.int32_t)(x.af), cgoAllocsUnknown
	allocsab82ee84.Borrow(caf_allocs)

	x.refab82ee84 = refab82ee84
	x.allocsab82ee84 = allocsab82ee84
	return refab82ee84, allocsab82ee84

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x network_callerid_msg_t) PassValue() (C.network_callerid_msg_t, *cgoAllocMap) {
	if x.refab82ee84 != nil {
		return *x.refab82ee84, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *network_callerid_msg_t) Deref() {
	if x.refab82ee84 == nil {
		return
	}
	x.ip_src = *(*[16]byte)(unsafe.Pointer(&x.refab82ee84.ip_src))
	x.ip_dst = *(*[16]byte)(unsafe.Pointer(&x.refab82ee84.ip_dst))
	x.port_src = (uint32_t)(x.refab82ee84.port_src)
	x.port_dst = (uint32_t)(x.refab82ee84.port_dst)
	x.af = (int32_t)(x.refab82ee84.af)
}

// allocJob_step_kill_msg_tMemory allocates memory for type C.job_step_kill_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_step_kill_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_step_kill_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_step_kill_msg_tValue = unsafe.Sizeof([1]C.job_step_kill_msg_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *job_step_kill_msg_t) Ref() *C.job_step_kill_msg_t {
	if x == nil {
		return nil
	}
	return x.refd4cb75f2
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *job_step_kill_msg_t) Free() {
	if x != nil && x.allocsd4cb75f2 != nil {
		x.allocsd4cb75f2.(*cgoAllocMap).Free()
		x.refd4cb75f2 = nil
	}
}

// Newjob_step_kill_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newjob_step_kill_msg_tRef(ref unsafe.Pointer) *job_step_kill_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(job_step_kill_msg_t)
	obj.refd4cb75f2 = (*C.job_step_kill_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *job_step_kill_msg_t) PassRef() (*C.job_step_kill_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd4cb75f2 != nil {
		return x.refd4cb75f2, nil
	}
	memd4cb75f2 := allocJob_step_kill_msg_tMemory(1)
	refd4cb75f2 := (*C.job_step_kill_msg_t)(memd4cb75f2)
	allocsd4cb75f2 := new(cgoAllocMap)
	allocsd4cb75f2.Add(memd4cb75f2)

	var cjob_id_allocs *cgoAllocMap
	refd4cb75f2.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocsd4cb75f2.Borrow(cjob_id_allocs)

	var csjob_id_allocs *cgoAllocMap
	refd4cb75f2.sjob_id, csjob_id_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.sjob_id)).Data)), cgoAllocsUnknown
	allocsd4cb75f2.Borrow(csjob_id_allocs)

	var cjob_step_id_allocs *cgoAllocMap
	refd4cb75f2.job_step_id, cjob_step_id_allocs = (C.uint32_t)(x.job_step_id), cgoAllocsUnknown
	allocsd4cb75f2.Borrow(cjob_step_id_allocs)

	var csignal_allocs *cgoAllocMap
	refd4cb75f2.signal, csignal_allocs = (C.uint16_t)(x.signal), cgoAllocsUnknown
	allocsd4cb75f2.Borrow(csignal_allocs)

	var cflags_allocs *cgoAllocMap
	refd4cb75f2.flags, cflags_allocs = (C.uint16_t)(x.flags), cgoAllocsUnknown
	allocsd4cb75f2.Borrow(cflags_allocs)

	var csibling_allocs *cgoAllocMap
	refd4cb75f2.sibling, csibling_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.sibling)).Data)), cgoAllocsUnknown
	allocsd4cb75f2.Borrow(csibling_allocs)

	x.refd4cb75f2 = refd4cb75f2
	x.allocsd4cb75f2 = allocsd4cb75f2
	return refd4cb75f2, allocsd4cb75f2

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x job_step_kill_msg_t) PassValue() (C.job_step_kill_msg_t, *cgoAllocMap) {
	if x.refd4cb75f2 != nil {
		return *x.refd4cb75f2, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *job_step_kill_msg_t) Deref() {
	if x.refd4cb75f2 == nil {
		return
	}
	x.job_id = (uint32_t)(x.refd4cb75f2.job_id)
	hxf241faf := (*sliceHeader)(unsafe.Pointer(&x.sjob_id))
	hxf241faf.Data = uintptr(unsafe.Pointer(x.refd4cb75f2.sjob_id))
	hxf241faf.Cap = 0x7fffffff
	// hxf241faf.Len = ?

	x.job_step_id = (uint32_t)(x.refd4cb75f2.job_step_id)
	x.signal = (uint16_t)(x.refd4cb75f2.signal)
	x.flags = (uint16_t)(x.refd4cb75f2.flags)
	hxfb50466 := (*sliceHeader)(unsafe.Pointer(&x.sibling))
	hxfb50466.Data = uintptr(unsafe.Pointer(x.refd4cb75f2.sibling))
	hxfb50466.Cap = 0x7fffffff
	// hxfb50466.Len = ?

}

// allocBurst_buffer_pool_tMemory allocates memory for type C.burst_buffer_pool_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBurst_buffer_pool_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBurst_buffer_pool_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfBurst_buffer_pool_tValue = unsafe.Sizeof([1]C.burst_buffer_pool_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *burst_buffer_pool_t) Ref() *C.burst_buffer_pool_t {
	if x == nil {
		return nil
	}
	return x.ref907a8fe1
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *burst_buffer_pool_t) Free() {
	if x != nil && x.allocs907a8fe1 != nil {
		x.allocs907a8fe1.(*cgoAllocMap).Free()
		x.ref907a8fe1 = nil
	}
}

// Newburst_buffer_pool_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newburst_buffer_pool_tRef(ref unsafe.Pointer) *burst_buffer_pool_t {
	if ref == nil {
		return nil
	}
	obj := new(burst_buffer_pool_t)
	obj.ref907a8fe1 = (*C.burst_buffer_pool_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *burst_buffer_pool_t) PassRef() (*C.burst_buffer_pool_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref907a8fe1 != nil {
		return x.ref907a8fe1, nil
	}
	mem907a8fe1 := allocBurst_buffer_pool_tMemory(1)
	ref907a8fe1 := (*C.burst_buffer_pool_t)(mem907a8fe1)
	allocs907a8fe1 := new(cgoAllocMap)
	allocs907a8fe1.Add(mem907a8fe1)

	var cgranularity_allocs *cgoAllocMap
	ref907a8fe1.granularity, cgranularity_allocs = (C.uint64_t)(x.granularity), cgoAllocsUnknown
	allocs907a8fe1.Borrow(cgranularity_allocs)

	var cname_allocs *cgoAllocMap
	ref907a8fe1.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs907a8fe1.Borrow(cname_allocs)

	var ctotal_space_allocs *cgoAllocMap
	ref907a8fe1.total_space, ctotal_space_allocs = (C.uint64_t)(x.total_space), cgoAllocsUnknown
	allocs907a8fe1.Borrow(ctotal_space_allocs)

	var cused_space_allocs *cgoAllocMap
	ref907a8fe1.used_space, cused_space_allocs = (C.uint64_t)(x.used_space), cgoAllocsUnknown
	allocs907a8fe1.Borrow(cused_space_allocs)

	var cunfree_space_allocs *cgoAllocMap
	ref907a8fe1.unfree_space, cunfree_space_allocs = (C.uint64_t)(x.unfree_space), cgoAllocsUnknown
	allocs907a8fe1.Borrow(cunfree_space_allocs)

	x.ref907a8fe1 = ref907a8fe1
	x.allocs907a8fe1 = allocs907a8fe1
	return ref907a8fe1, allocs907a8fe1

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x burst_buffer_pool_t) PassValue() (C.burst_buffer_pool_t, *cgoAllocMap) {
	if x.ref907a8fe1 != nil {
		return *x.ref907a8fe1, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *burst_buffer_pool_t) Deref() {
	if x.ref907a8fe1 == nil {
		return
	}
	x.granularity = (uint64_t)(x.ref907a8fe1.granularity)
	hxf2bca6d := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf2bca6d.Data = uintptr(unsafe.Pointer(x.ref907a8fe1.name))
	hxf2bca6d.Cap = 0x7fffffff
	// hxf2bca6d.Len = ?

	x.total_space = (uint64_t)(x.ref907a8fe1.total_space)
	x.used_space = (uint64_t)(x.ref907a8fe1.used_space)
	x.unfree_space = (uint64_t)(x.ref907a8fe1.unfree_space)
}

// allocBurst_buffer_resv_tMemory allocates memory for type C.burst_buffer_resv_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBurst_buffer_resv_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBurst_buffer_resv_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfBurst_buffer_resv_tValue = unsafe.Sizeof([1]C.burst_buffer_resv_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *burst_buffer_resv_t) Ref() *C.burst_buffer_resv_t {
	if x == nil {
		return nil
	}
	return x.ref9c7cf8ca
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *burst_buffer_resv_t) Free() {
	if x != nil && x.allocs9c7cf8ca != nil {
		x.allocs9c7cf8ca.(*cgoAllocMap).Free()
		x.ref9c7cf8ca = nil
	}
}

// Newburst_buffer_resv_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newburst_buffer_resv_tRef(ref unsafe.Pointer) *burst_buffer_resv_t {
	if ref == nil {
		return nil
	}
	obj := new(burst_buffer_resv_t)
	obj.ref9c7cf8ca = (*C.burst_buffer_resv_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *burst_buffer_resv_t) PassRef() (*C.burst_buffer_resv_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9c7cf8ca != nil {
		return x.ref9c7cf8ca, nil
	}
	mem9c7cf8ca := allocBurst_buffer_resv_tMemory(1)
	ref9c7cf8ca := (*C.burst_buffer_resv_t)(mem9c7cf8ca)
	allocs9c7cf8ca := new(cgoAllocMap)
	allocs9c7cf8ca.Add(mem9c7cf8ca)

	var caccount_allocs *cgoAllocMap
	ref9c7cf8ca.account, caccount_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.account)).Data)), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(caccount_allocs)

	var carray_job_id_allocs *cgoAllocMap
	ref9c7cf8ca.array_job_id, carray_job_id_allocs = (C.uint32_t)(x.array_job_id), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(carray_job_id_allocs)

	var carray_task_id_allocs *cgoAllocMap
	ref9c7cf8ca.array_task_id, carray_task_id_allocs = (C.uint32_t)(x.array_task_id), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(carray_task_id_allocs)

	var ccreate_time_allocs *cgoAllocMap
	ref9c7cf8ca.create_time, ccreate_time_allocs = (C.time_t)(x.create_time), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(ccreate_time_allocs)

	var cjob_id_allocs *cgoAllocMap
	ref9c7cf8ca.job_id, cjob_id_allocs = (C.uint32_t)(x.job_id), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(cjob_id_allocs)

	var cname_allocs *cgoAllocMap
	ref9c7cf8ca.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(cname_allocs)

	var cpartition_allocs *cgoAllocMap
	ref9c7cf8ca.partition, cpartition_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.partition)).Data)), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(cpartition_allocs)

	var cpool_allocs *cgoAllocMap
	ref9c7cf8ca.pool, cpool_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.pool)).Data)), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(cpool_allocs)

	var cqos_allocs *cgoAllocMap
	ref9c7cf8ca.qos, cqos_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.qos)).Data)), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(cqos_allocs)

	var csize_allocs *cgoAllocMap
	ref9c7cf8ca.size, csize_allocs = (C.uint64_t)(x.size), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(csize_allocs)

	var cstate_allocs *cgoAllocMap
	ref9c7cf8ca.state, cstate_allocs = (C.uint16_t)(x.state), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(cstate_allocs)

	var cuser_id_allocs *cgoAllocMap
	ref9c7cf8ca.user_id, cuser_id_allocs = (C.uint32_t)(x.user_id), cgoAllocsUnknown
	allocs9c7cf8ca.Borrow(cuser_id_allocs)

	x.ref9c7cf8ca = ref9c7cf8ca
	x.allocs9c7cf8ca = allocs9c7cf8ca
	return ref9c7cf8ca, allocs9c7cf8ca

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x burst_buffer_resv_t) PassValue() (C.burst_buffer_resv_t, *cgoAllocMap) {
	if x.ref9c7cf8ca != nil {
		return *x.ref9c7cf8ca, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *burst_buffer_resv_t) Deref() {
	if x.ref9c7cf8ca == nil {
		return
	}
	hxf48f9f4 := (*sliceHeader)(unsafe.Pointer(&x.account))
	hxf48f9f4.Data = uintptr(unsafe.Pointer(x.ref9c7cf8ca.account))
	hxf48f9f4.Cap = 0x7fffffff
	// hxf48f9f4.Len = ?

	x.array_job_id = (uint32_t)(x.ref9c7cf8ca.array_job_id)
	x.array_task_id = (uint32_t)(x.ref9c7cf8ca.array_task_id)
	x.create_time = (time_t)(x.ref9c7cf8ca.create_time)
	x.job_id = (uint32_t)(x.ref9c7cf8ca.job_id)
	hxf64ca81 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf64ca81.Data = uintptr(unsafe.Pointer(x.ref9c7cf8ca.name))
	hxf64ca81.Cap = 0x7fffffff
	// hxf64ca81.Len = ?

	hxf114a2f := (*sliceHeader)(unsafe.Pointer(&x.partition))
	hxf114a2f.Data = uintptr(unsafe.Pointer(x.ref9c7cf8ca.partition))
	hxf114a2f.Cap = 0x7fffffff
	// hxf114a2f.Len = ?

	hxf7dbefa := (*sliceHeader)(unsafe.Pointer(&x.pool))
	hxf7dbefa.Data = uintptr(unsafe.Pointer(x.ref9c7cf8ca.pool))
	hxf7dbefa.Cap = 0x7fffffff
	// hxf7dbefa.Len = ?

	hxf66c126 := (*sliceHeader)(unsafe.Pointer(&x.qos))
	hxf66c126.Data = uintptr(unsafe.Pointer(x.ref9c7cf8ca.qos))
	hxf66c126.Cap = 0x7fffffff
	// hxf66c126.Len = ?

	x.size = (uint64_t)(x.ref9c7cf8ca.size)
	x.state = (uint16_t)(x.ref9c7cf8ca.state)
	x.user_id = (uint32_t)(x.ref9c7cf8ca.user_id)
}

// allocBurst_buffer_use_tMemory allocates memory for type C.burst_buffer_use_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBurst_buffer_use_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBurst_buffer_use_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfBurst_buffer_use_tValue = unsafe.Sizeof([1]C.burst_buffer_use_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *burst_buffer_use_t) Ref() *C.burst_buffer_use_t {
	if x == nil {
		return nil
	}
	return x.ref4d977eac
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *burst_buffer_use_t) Free() {
	if x != nil && x.allocs4d977eac != nil {
		x.allocs4d977eac.(*cgoAllocMap).Free()
		x.ref4d977eac = nil
	}
}

// Newburst_buffer_use_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newburst_buffer_use_tRef(ref unsafe.Pointer) *burst_buffer_use_t {
	if ref == nil {
		return nil
	}
	obj := new(burst_buffer_use_t)
	obj.ref4d977eac = (*C.burst_buffer_use_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *burst_buffer_use_t) PassRef() (*C.burst_buffer_use_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref4d977eac != nil {
		return x.ref4d977eac, nil
	}
	mem4d977eac := allocBurst_buffer_use_tMemory(1)
	ref4d977eac := (*C.burst_buffer_use_t)(mem4d977eac)
	allocs4d977eac := new(cgoAllocMap)
	allocs4d977eac.Add(mem4d977eac)

	var cuser_id_allocs *cgoAllocMap
	ref4d977eac.user_id, cuser_id_allocs = (C.uint32_t)(x.user_id), cgoAllocsUnknown
	allocs4d977eac.Borrow(cuser_id_allocs)

	var cused_allocs *cgoAllocMap
	ref4d977eac.used, cused_allocs = (C.uint64_t)(x.used), cgoAllocsUnknown
	allocs4d977eac.Borrow(cused_allocs)

	x.ref4d977eac = ref4d977eac
	x.allocs4d977eac = allocs4d977eac
	return ref4d977eac, allocs4d977eac

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x burst_buffer_use_t) PassValue() (C.burst_buffer_use_t, *cgoAllocMap) {
	if x.ref4d977eac != nil {
		return *x.ref4d977eac, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *burst_buffer_use_t) Deref() {
	if x.ref4d977eac == nil {
		return
	}
	x.user_id = (uint32_t)(x.ref4d977eac.user_id)
	x.used = (uint64_t)(x.ref4d977eac.used)
}

// allocBurst_buffer_info_tMemory allocates memory for type C.burst_buffer_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBurst_buffer_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBurst_buffer_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfBurst_buffer_info_tValue = unsafe.Sizeof([1]C.burst_buffer_info_t{})

// unpackSBurst_buffer_pool_t transforms a sliced Go data structure into plain C format.
func unpackSBurst_buffer_pool_t(x []burst_buffer_pool_t) (unpacked *C.burst_buffer_pool_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.burst_buffer_pool_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocBurst_buffer_pool_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.burst_buffer_pool_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.burst_buffer_pool_t)(unsafe.Pointer(h.Data))
	return
}

// unpackSBurst_buffer_resv_t transforms a sliced Go data structure into plain C format.
func unpackSBurst_buffer_resv_t(x []burst_buffer_resv_t) (unpacked *C.burst_buffer_resv_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.burst_buffer_resv_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocBurst_buffer_resv_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.burst_buffer_resv_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.burst_buffer_resv_t)(unsafe.Pointer(h.Data))
	return
}

// unpackSBurst_buffer_use_t transforms a sliced Go data structure into plain C format.
func unpackSBurst_buffer_use_t(x []burst_buffer_use_t) (unpacked *C.burst_buffer_use_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.burst_buffer_use_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocBurst_buffer_use_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.burst_buffer_use_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.burst_buffer_use_t)(unsafe.Pointer(h.Data))
	return
}

// packSBurst_buffer_pool_t reads sliced Go data structure out from plain C format.
func packSBurst_buffer_pool_t(v []burst_buffer_pool_t, ptr0 *C.burst_buffer_pool_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfBurst_buffer_pool_tValue]C.burst_buffer_pool_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newburst_buffer_pool_tRef(unsafe.Pointer(&ptr1))
	}
}

// packSBurst_buffer_resv_t reads sliced Go data structure out from plain C format.
func packSBurst_buffer_resv_t(v []burst_buffer_resv_t, ptr0 *C.burst_buffer_resv_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfBurst_buffer_resv_tValue]C.burst_buffer_resv_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newburst_buffer_resv_tRef(unsafe.Pointer(&ptr1))
	}
}

// packSBurst_buffer_use_t reads sliced Go data structure out from plain C format.
func packSBurst_buffer_use_t(v []burst_buffer_use_t, ptr0 *C.burst_buffer_use_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfBurst_buffer_use_tValue]C.burst_buffer_use_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newburst_buffer_use_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *burst_buffer_info_t) Ref() *C.burst_buffer_info_t {
	if x == nil {
		return nil
	}
	return x.reff68d04d1
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *burst_buffer_info_t) Free() {
	if x != nil && x.allocsf68d04d1 != nil {
		x.allocsf68d04d1.(*cgoAllocMap).Free()
		x.reff68d04d1 = nil
	}
}

// Newburst_buffer_info_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newburst_buffer_info_tRef(ref unsafe.Pointer) *burst_buffer_info_t {
	if ref == nil {
		return nil
	}
	obj := new(burst_buffer_info_t)
	obj.reff68d04d1 = (*C.burst_buffer_info_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *burst_buffer_info_t) PassRef() (*C.burst_buffer_info_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff68d04d1 != nil {
		return x.reff68d04d1, nil
	}
	memf68d04d1 := allocBurst_buffer_info_tMemory(1)
	reff68d04d1 := (*C.burst_buffer_info_t)(memf68d04d1)
	allocsf68d04d1 := new(cgoAllocMap)
	allocsf68d04d1.Add(memf68d04d1)

	var callow_users_allocs *cgoAllocMap
	reff68d04d1.allow_users, callow_users_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.allow_users)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(callow_users_allocs)

	var cdefault_pool_allocs *cgoAllocMap
	reff68d04d1.default_pool, cdefault_pool_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.default_pool)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cdefault_pool_allocs)

	var ccreate_buffer_allocs *cgoAllocMap
	reff68d04d1.create_buffer, ccreate_buffer_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.create_buffer)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(ccreate_buffer_allocs)

	var cdeny_users_allocs *cgoAllocMap
	reff68d04d1.deny_users, cdeny_users_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.deny_users)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cdeny_users_allocs)

	var cdestroy_buffer_allocs *cgoAllocMap
	reff68d04d1.destroy_buffer, cdestroy_buffer_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.destroy_buffer)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cdestroy_buffer_allocs)

	var cflags_allocs *cgoAllocMap
	reff68d04d1.flags, cflags_allocs = (C.uint32_t)(x.flags), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cflags_allocs)

	var cget_sys_state_allocs *cgoAllocMap
	reff68d04d1.get_sys_state, cget_sys_state_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.get_sys_state)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cget_sys_state_allocs)

	var cgranularity_allocs *cgoAllocMap
	reff68d04d1.granularity, cgranularity_allocs = (C.uint64_t)(x.granularity), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cgranularity_allocs)

	var cpool_cnt_allocs *cgoAllocMap
	reff68d04d1.pool_cnt, cpool_cnt_allocs = (C.uint32_t)(x.pool_cnt), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cpool_cnt_allocs)

	var cpool_ptr_allocs *cgoAllocMap
	reff68d04d1.pool_ptr, cpool_ptr_allocs = unpackSBurst_buffer_pool_t(x.pool_ptr)
	allocsf68d04d1.Borrow(cpool_ptr_allocs)

	var cname_allocs *cgoAllocMap
	reff68d04d1.name, cname_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.name)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cname_allocs)

	var cother_timeout_allocs *cgoAllocMap
	reff68d04d1.other_timeout, cother_timeout_allocs = (C.uint32_t)(x.other_timeout), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cother_timeout_allocs)

	var cstage_in_timeout_allocs *cgoAllocMap
	reff68d04d1.stage_in_timeout, cstage_in_timeout_allocs = (C.uint32_t)(x.stage_in_timeout), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cstage_in_timeout_allocs)

	var cstage_out_timeout_allocs *cgoAllocMap
	reff68d04d1.stage_out_timeout, cstage_out_timeout_allocs = (C.uint32_t)(x.stage_out_timeout), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cstage_out_timeout_allocs)

	var cstart_stage_in_allocs *cgoAllocMap
	reff68d04d1.start_stage_in, cstart_stage_in_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.start_stage_in)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cstart_stage_in_allocs)

	var cstart_stage_out_allocs *cgoAllocMap
	reff68d04d1.start_stage_out, cstart_stage_out_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.start_stage_out)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cstart_stage_out_allocs)

	var cstop_stage_in_allocs *cgoAllocMap
	reff68d04d1.stop_stage_in, cstop_stage_in_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.stop_stage_in)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cstop_stage_in_allocs)

	var cstop_stage_out_allocs *cgoAllocMap
	reff68d04d1.stop_stage_out, cstop_stage_out_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.stop_stage_out)).Data)), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cstop_stage_out_allocs)

	var ctotal_space_allocs *cgoAllocMap
	reff68d04d1.total_space, ctotal_space_allocs = (C.uint64_t)(x.total_space), cgoAllocsUnknown
	allocsf68d04d1.Borrow(ctotal_space_allocs)

	var cunfree_space_allocs *cgoAllocMap
	reff68d04d1.unfree_space, cunfree_space_allocs = (C.uint64_t)(x.unfree_space), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cunfree_space_allocs)

	var cused_space_allocs *cgoAllocMap
	reff68d04d1.used_space, cused_space_allocs = (C.uint64_t)(x.used_space), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cused_space_allocs)

	var cvalidate_timeout_allocs *cgoAllocMap
	reff68d04d1.validate_timeout, cvalidate_timeout_allocs = (C.uint32_t)(x.validate_timeout), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cvalidate_timeout_allocs)

	var cbuffer_count_allocs *cgoAllocMap
	reff68d04d1.buffer_count, cbuffer_count_allocs = (C.uint32_t)(x.buffer_count), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cbuffer_count_allocs)

	var cburst_buffer_resv_ptr_allocs *cgoAllocMap
	reff68d04d1.burst_buffer_resv_ptr, cburst_buffer_resv_ptr_allocs = unpackSBurst_buffer_resv_t(x.burst_buffer_resv_ptr)
	allocsf68d04d1.Borrow(cburst_buffer_resv_ptr_allocs)

	var cuse_count_allocs *cgoAllocMap
	reff68d04d1.use_count, cuse_count_allocs = (C.uint32_t)(x.use_count), cgoAllocsUnknown
	allocsf68d04d1.Borrow(cuse_count_allocs)

	var cburst_buffer_use_ptr_allocs *cgoAllocMap
	reff68d04d1.burst_buffer_use_ptr, cburst_buffer_use_ptr_allocs = unpackSBurst_buffer_use_t(x.burst_buffer_use_ptr)
	allocsf68d04d1.Borrow(cburst_buffer_use_ptr_allocs)

	x.reff68d04d1 = reff68d04d1
	x.allocsf68d04d1 = allocsf68d04d1
	return reff68d04d1, allocsf68d04d1

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x burst_buffer_info_t) PassValue() (C.burst_buffer_info_t, *cgoAllocMap) {
	if x.reff68d04d1 != nil {
		return *x.reff68d04d1, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *burst_buffer_info_t) Deref() {
	if x.reff68d04d1 == nil {
		return
	}
	hxf70737f := (*sliceHeader)(unsafe.Pointer(&x.allow_users))
	hxf70737f.Data = uintptr(unsafe.Pointer(x.reff68d04d1.allow_users))
	hxf70737f.Cap = 0x7fffffff
	// hxf70737f.Len = ?

	hxfa8cf1d := (*sliceHeader)(unsafe.Pointer(&x.default_pool))
	hxfa8cf1d.Data = uintptr(unsafe.Pointer(x.reff68d04d1.default_pool))
	hxfa8cf1d.Cap = 0x7fffffff
	// hxfa8cf1d.Len = ?

	hxf86e299 := (*sliceHeader)(unsafe.Pointer(&x.create_buffer))
	hxf86e299.Data = uintptr(unsafe.Pointer(x.reff68d04d1.create_buffer))
	hxf86e299.Cap = 0x7fffffff
	// hxf86e299.Len = ?

	hxf51a478 := (*sliceHeader)(unsafe.Pointer(&x.deny_users))
	hxf51a478.Data = uintptr(unsafe.Pointer(x.reff68d04d1.deny_users))
	hxf51a478.Cap = 0x7fffffff
	// hxf51a478.Len = ?

	hxf2768c1 := (*sliceHeader)(unsafe.Pointer(&x.destroy_buffer))
	hxf2768c1.Data = uintptr(unsafe.Pointer(x.reff68d04d1.destroy_buffer))
	hxf2768c1.Cap = 0x7fffffff
	// hxf2768c1.Len = ?

	x.flags = (uint32_t)(x.reff68d04d1.flags)
	hxf2383cd := (*sliceHeader)(unsafe.Pointer(&x.get_sys_state))
	hxf2383cd.Data = uintptr(unsafe.Pointer(x.reff68d04d1.get_sys_state))
	hxf2383cd.Cap = 0x7fffffff
	// hxf2383cd.Len = ?

	x.granularity = (uint64_t)(x.reff68d04d1.granularity)
	x.pool_cnt = (uint32_t)(x.reff68d04d1.pool_cnt)
	packSBurst_buffer_pool_t(x.pool_ptr, x.reff68d04d1.pool_ptr)
	hxf266c09 := (*sliceHeader)(unsafe.Pointer(&x.name))
	hxf266c09.Data = uintptr(unsafe.Pointer(x.reff68d04d1.name))
	hxf266c09.Cap = 0x7fffffff
	// hxf266c09.Len = ?

	x.other_timeout = (uint32_t)(x.reff68d04d1.other_timeout)
	x.stage_in_timeout = (uint32_t)(x.reff68d04d1.stage_in_timeout)
	x.stage_out_timeout = (uint32_t)(x.reff68d04d1.stage_out_timeout)
	hxf01775c := (*sliceHeader)(unsafe.Pointer(&x.start_stage_in))
	hxf01775c.Data = uintptr(unsafe.Pointer(x.reff68d04d1.start_stage_in))
	hxf01775c.Cap = 0x7fffffff
	// hxf01775c.Len = ?

	hxf606340 := (*sliceHeader)(unsafe.Pointer(&x.start_stage_out))
	hxf606340.Data = uintptr(unsafe.Pointer(x.reff68d04d1.start_stage_out))
	hxf606340.Cap = 0x7fffffff
	// hxf606340.Len = ?

	hxfbc2dfa := (*sliceHeader)(unsafe.Pointer(&x.stop_stage_in))
	hxfbc2dfa.Data = uintptr(unsafe.Pointer(x.reff68d04d1.stop_stage_in))
	hxfbc2dfa.Cap = 0x7fffffff
	// hxfbc2dfa.Len = ?

	hxf8a16c3 := (*sliceHeader)(unsafe.Pointer(&x.stop_stage_out))
	hxf8a16c3.Data = uintptr(unsafe.Pointer(x.reff68d04d1.stop_stage_out))
	hxf8a16c3.Cap = 0x7fffffff
	// hxf8a16c3.Len = ?

	x.total_space = (uint64_t)(x.reff68d04d1.total_space)
	x.unfree_space = (uint64_t)(x.reff68d04d1.unfree_space)
	x.used_space = (uint64_t)(x.reff68d04d1.used_space)
	x.validate_timeout = (uint32_t)(x.reff68d04d1.validate_timeout)
	x.buffer_count = (uint32_t)(x.reff68d04d1.buffer_count)
	packSBurst_buffer_resv_t(x.burst_buffer_resv_ptr, x.reff68d04d1.burst_buffer_resv_ptr)
	x.use_count = (uint32_t)(x.reff68d04d1.use_count)
	packSBurst_buffer_use_t(x.burst_buffer_use_ptr, x.reff68d04d1.burst_buffer_use_ptr)
}

// allocBurst_buffer_info_msg_tMemory allocates memory for type C.burst_buffer_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBurst_buffer_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBurst_buffer_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfBurst_buffer_info_msg_tValue = unsafe.Sizeof([1]C.burst_buffer_info_msg_t{})

// unpackSBurst_buffer_info_t transforms a sliced Go data structure into plain C format.
func unpackSBurst_buffer_info_t(x []burst_buffer_info_t) (unpacked *C.burst_buffer_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.burst_buffer_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocBurst_buffer_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.burst_buffer_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.burst_buffer_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSBurst_buffer_info_t reads sliced Go data structure out from plain C format.
func packSBurst_buffer_info_t(v []burst_buffer_info_t, ptr0 *C.burst_buffer_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfBurst_buffer_info_tValue]C.burst_buffer_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newburst_buffer_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *burst_buffer_info_msg_t) Ref() *C.burst_buffer_info_msg_t {
	if x == nil {
		return nil
	}
	return x.ref51ffcb89
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *burst_buffer_info_msg_t) Free() {
	if x != nil && x.allocs51ffcb89 != nil {
		x.allocs51ffcb89.(*cgoAllocMap).Free()
		x.ref51ffcb89 = nil
	}
}

// Newburst_buffer_info_msg_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newburst_buffer_info_msg_tRef(ref unsafe.Pointer) *burst_buffer_info_msg_t {
	if ref == nil {
		return nil
	}
	obj := new(burst_buffer_info_msg_t)
	obj.ref51ffcb89 = (*C.burst_buffer_info_msg_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *burst_buffer_info_msg_t) PassRef() (*C.burst_buffer_info_msg_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref51ffcb89 != nil {
		return x.ref51ffcb89, nil
	}
	mem51ffcb89 := allocBurst_buffer_info_msg_tMemory(1)
	ref51ffcb89 := (*C.burst_buffer_info_msg_t)(mem51ffcb89)
	allocs51ffcb89 := new(cgoAllocMap)
	allocs51ffcb89.Add(mem51ffcb89)

	var cburst_buffer_array_allocs *cgoAllocMap
	ref51ffcb89.burst_buffer_array, cburst_buffer_array_allocs = unpackSBurst_buffer_info_t(x.burst_buffer_array)
	allocs51ffcb89.Borrow(cburst_buffer_array_allocs)

	var crecord_count_allocs *cgoAllocMap
	ref51ffcb89.record_count, crecord_count_allocs = (C.uint32_t)(x.record_count), cgoAllocsUnknown
	allocs51ffcb89.Borrow(crecord_count_allocs)

	x.ref51ffcb89 = ref51ffcb89
	x.allocs51ffcb89 = allocs51ffcb89
	return ref51ffcb89, allocs51ffcb89

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x burst_buffer_info_msg_t) PassValue() (C.burst_buffer_info_msg_t, *cgoAllocMap) {
	if x.ref51ffcb89 != nil {
		return *x.ref51ffcb89, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *burst_buffer_info_msg_t) Deref() {
	if x.ref51ffcb89 == nil {
		return
	}
	packSBurst_buffer_info_t(x.burst_buffer_array, x.ref51ffcb89.burst_buffer_array)
	x.record_count = (uint32_t)(x.ref51ffcb89.record_count)
}

// allocImaxdiv_tMemory allocates memory for type C.imaxdiv_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocImaxdiv_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfImaxdiv_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfImaxdiv_tValue = unsafe.Sizeof([1]C.imaxdiv_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *imaxdiv_t) Ref() *C.imaxdiv_t {
	if x == nil {
		return nil
	}
	return x.ref873cd79e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *imaxdiv_t) Free() {
	if x != nil && x.allocs873cd79e != nil {
		x.allocs873cd79e.(*cgoAllocMap).Free()
		x.ref873cd79e = nil
	}
}

// Newimaxdiv_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newimaxdiv_tRef(ref unsafe.Pointer) *imaxdiv_t {
	if ref == nil {
		return nil
	}
	obj := new(imaxdiv_t)
	obj.ref873cd79e = (*C.imaxdiv_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *imaxdiv_t) PassRef() (*C.imaxdiv_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref873cd79e != nil {
		return x.ref873cd79e, nil
	}
	mem873cd79e := allocImaxdiv_tMemory(1)
	ref873cd79e := (*C.imaxdiv_t)(mem873cd79e)
	allocs873cd79e := new(cgoAllocMap)
	allocs873cd79e.Add(mem873cd79e)

	var cquot_allocs *cgoAllocMap
	ref873cd79e.quot, cquot_allocs = (C.long)(x.quot), cgoAllocsUnknown
	allocs873cd79e.Borrow(cquot_allocs)

	var crem_allocs *cgoAllocMap
	ref873cd79e.rem, crem_allocs = (C.long)(x.rem), cgoAllocsUnknown
	allocs873cd79e.Borrow(crem_allocs)

	x.ref873cd79e = ref873cd79e
	x.allocs873cd79e = allocs873cd79e
	return ref873cd79e, allocs873cd79e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x imaxdiv_t) PassValue() (C.imaxdiv_t, *cgoAllocMap) {
	if x.ref873cd79e != nil {
		return *x.ref873cd79e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *imaxdiv_t) Deref() {
	if x.ref873cd79e == nil {
		return
	}
	x.quot = (int)(x.ref873cd79e.quot)
	x.rem = (int)(x.ref873cd79e.rem)
}

// allocFILEMemory allocates memory for type C.FILE in C.
// The caller is responsible for freeing the this memory via C.free.
func allocFILEMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfFILEValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfFILEValue = unsafe.Sizeof([1]C.FILE{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *FILE) Ref() *C.FILE {
	if x == nil {
		return nil
	}
	return x.refba0adba4
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *FILE) Free() {
	if x != nil && x.allocsba0adba4 != nil {
		x.allocsba0adba4.(*cgoAllocMap).Free()
		x.refba0adba4 = nil
	}
}

// NewFILERef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewFILERef(ref unsafe.Pointer) *FILE {
	if ref == nil {
		return nil
	}
	obj := new(FILE)
	obj.refba0adba4 = (*C.FILE)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *FILE) PassRef() (*C.FILE, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refba0adba4 != nil {
		return x.refba0adba4, nil
	}
	memba0adba4 := allocFILEMemory(1)
	refba0adba4 := (*C.FILE)(memba0adba4)
	allocsba0adba4 := new(cgoAllocMap)
	allocsba0adba4.Add(memba0adba4)

	var c_flags_allocs *cgoAllocMap
	refba0adba4._flags, c_flags_allocs = (C.int)(x._flags), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_flags_allocs)

	var c_IO_read_ptr_allocs *cgoAllocMap
	refba0adba4._IO_read_ptr, c_IO_read_ptr_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_read_ptr)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_read_ptr_allocs)

	var c_IO_read_end_allocs *cgoAllocMap
	refba0adba4._IO_read_end, c_IO_read_end_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_read_end)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_read_end_allocs)

	var c_IO_read_base_allocs *cgoAllocMap
	refba0adba4._IO_read_base, c_IO_read_base_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_read_base)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_read_base_allocs)

	var c_IO_write_base_allocs *cgoAllocMap
	refba0adba4._IO_write_base, c_IO_write_base_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_write_base)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_write_base_allocs)

	var c_IO_write_ptr_allocs *cgoAllocMap
	refba0adba4._IO_write_ptr, c_IO_write_ptr_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_write_ptr)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_write_ptr_allocs)

	var c_IO_write_end_allocs *cgoAllocMap
	refba0adba4._IO_write_end, c_IO_write_end_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_write_end)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_write_end_allocs)

	var c_IO_buf_base_allocs *cgoAllocMap
	refba0adba4._IO_buf_base, c_IO_buf_base_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_buf_base)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_buf_base_allocs)

	var c_IO_buf_end_allocs *cgoAllocMap
	refba0adba4._IO_buf_end, c_IO_buf_end_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_buf_end)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_buf_end_allocs)

	var c_IO_save_base_allocs *cgoAllocMap
	refba0adba4._IO_save_base, c_IO_save_base_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_save_base)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_save_base_allocs)

	var c_IO_backup_base_allocs *cgoAllocMap
	refba0adba4._IO_backup_base, c_IO_backup_base_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_backup_base)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_backup_base_allocs)

	var c_IO_save_end_allocs *cgoAllocMap
	refba0adba4._IO_save_end, c_IO_save_end_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x._IO_save_end)).Data)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_IO_save_end_allocs)

	var c_fileno_allocs *cgoAllocMap
	refba0adba4._fileno, c_fileno_allocs = (C.int)(x._fileno), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_fileno_allocs)

	var c_flags2_allocs *cgoAllocMap
	refba0adba4._flags2, c_flags2_allocs = (C.int)(x._flags2), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_flags2_allocs)

	var c_old_offset_allocs *cgoAllocMap
	refba0adba4._old_offset, c_old_offset_allocs = (C.__off_t)(x._old_offset), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_old_offset_allocs)

	var c_cur_column_allocs *cgoAllocMap
	refba0adba4._cur_column, c_cur_column_allocs = (C.ushort)(x._cur_column), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_cur_column_allocs)

	var c_vtable_offset_allocs *cgoAllocMap
	refba0adba4._vtable_offset, c_vtable_offset_allocs = (C.char)(x._vtable_offset), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_vtable_offset_allocs)

	var c_shortbuf_allocs *cgoAllocMap
	refba0adba4._shortbuf, c_shortbuf_allocs = *(*[1]C.char)(unsafe.Pointer(&x._shortbuf)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_shortbuf_allocs)

	var c_lock_allocs *cgoAllocMap
	refba0adba4._lock, c_lock_allocs = *(**C._IO_lock_t)(unsafe.Pointer(&x._lock)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_lock_allocs)

	var c_offset_allocs *cgoAllocMap
	refba0adba4._offset, c_offset_allocs = (C.__off64_t)(x._offset), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_offset_allocs)

	var c__pad1_allocs *cgoAllocMap
	refba0adba4.__pad1, c__pad1_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.__pad1)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c__pad1_allocs)

	var c__pad2_allocs *cgoAllocMap
	refba0adba4.__pad2, c__pad2_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.__pad2)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c__pad2_allocs)

	var c__pad3_allocs *cgoAllocMap
	refba0adba4.__pad3, c__pad3_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.__pad3)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c__pad3_allocs)

	var c__pad4_allocs *cgoAllocMap
	refba0adba4.__pad4, c__pad4_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.__pad4)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c__pad4_allocs)

	var c__pad5_allocs *cgoAllocMap
	refba0adba4.__pad5, c__pad5_allocs = (C.size_t)(x.__pad5), cgoAllocsUnknown
	allocsba0adba4.Borrow(c__pad5_allocs)

	var c_mode_allocs *cgoAllocMap
	refba0adba4._mode, c_mode_allocs = (C.int)(x._mode), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_mode_allocs)

	var c_unused2_allocs *cgoAllocMap
	refba0adba4._unused2, c_unused2_allocs = *(*[20]C.char)(unsafe.Pointer(&x._unused2)), cgoAllocsUnknown
	allocsba0adba4.Borrow(c_unused2_allocs)

	x.refba0adba4 = refba0adba4
	x.allocsba0adba4 = allocsba0adba4
	return refba0adba4, allocsba0adba4

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x FILE) PassValue() (C.FILE, *cgoAllocMap) {
	if x.refba0adba4 != nil {
		return *x.refba0adba4, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *FILE) Deref() {
	if x.refba0adba4 == nil {
		return
	}
	x._flags = (int32)(x.refba0adba4._flags)
	hxf8ce826 := (*sliceHeader)(unsafe.Pointer(&x._IO_read_ptr))
	hxf8ce826.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_read_ptr))
	hxf8ce826.Cap = 0x7fffffff
	// hxf8ce826.Len = ?

	hxf2f51bb := (*sliceHeader)(unsafe.Pointer(&x._IO_read_end))
	hxf2f51bb.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_read_end))
	hxf2f51bb.Cap = 0x7fffffff
	// hxf2f51bb.Len = ?

	hxfd8d280 := (*sliceHeader)(unsafe.Pointer(&x._IO_read_base))
	hxfd8d280.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_read_base))
	hxfd8d280.Cap = 0x7fffffff
	// hxfd8d280.Len = ?

	hxf02f99f := (*sliceHeader)(unsafe.Pointer(&x._IO_write_base))
	hxf02f99f.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_write_base))
	hxf02f99f.Cap = 0x7fffffff
	// hxf02f99f.Len = ?

	hxf48d2ac := (*sliceHeader)(unsafe.Pointer(&x._IO_write_ptr))
	hxf48d2ac.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_write_ptr))
	hxf48d2ac.Cap = 0x7fffffff
	// hxf48d2ac.Len = ?

	hxfdacaf0 := (*sliceHeader)(unsafe.Pointer(&x._IO_write_end))
	hxfdacaf0.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_write_end))
	hxfdacaf0.Cap = 0x7fffffff
	// hxfdacaf0.Len = ?

	hxfdc6d11 := (*sliceHeader)(unsafe.Pointer(&x._IO_buf_base))
	hxfdc6d11.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_buf_base))
	hxfdc6d11.Cap = 0x7fffffff
	// hxfdc6d11.Len = ?

	hxf659275 := (*sliceHeader)(unsafe.Pointer(&x._IO_buf_end))
	hxf659275.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_buf_end))
	hxf659275.Cap = 0x7fffffff
	// hxf659275.Len = ?

	hxfef6b75 := (*sliceHeader)(unsafe.Pointer(&x._IO_save_base))
	hxfef6b75.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_save_base))
	hxfef6b75.Cap = 0x7fffffff
	// hxfef6b75.Len = ?

	hxf2d2f6c := (*sliceHeader)(unsafe.Pointer(&x._IO_backup_base))
	hxf2d2f6c.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_backup_base))
	hxf2d2f6c.Cap = 0x7fffffff
	// hxf2d2f6c.Len = ?

	hxf3d1cf5 := (*sliceHeader)(unsafe.Pointer(&x._IO_save_end))
	hxf3d1cf5.Data = uintptr(unsafe.Pointer(x.refba0adba4._IO_save_end))
	hxf3d1cf5.Cap = 0x7fffffff
	// hxf3d1cf5.Len = ?

	x._fileno = (int32)(x.refba0adba4._fileno)
	x._flags2 = (int32)(x.refba0adba4._flags2)
	x._old_offset = (__off_t)(x.refba0adba4._old_offset)
	x._cur_column = (uint16)(x.refba0adba4._cur_column)
	x._vtable_offset = (byte)(x.refba0adba4._vtable_offset)
	x._shortbuf = *(*[1]byte)(unsafe.Pointer(&x.refba0adba4._shortbuf))
	x._lock = (*_IO_lock_t)(unsafe.Pointer(x.refba0adba4._lock))
	x._offset = (__off64_t)(x.refba0adba4._offset)
	x.__pad1 = (unsafe.Pointer)(unsafe.Pointer(x.refba0adba4.__pad1))
	x.__pad2 = (unsafe.Pointer)(unsafe.Pointer(x.refba0adba4.__pad2))
	x.__pad3 = (unsafe.Pointer)(unsafe.Pointer(x.refba0adba4.__pad3))
	x.__pad4 = (unsafe.Pointer)(unsafe.Pointer(x.refba0adba4.__pad4))
	x.__pad5 = (size_t)(x.refba0adba4.__pad5)
	x._mode = (int32)(x.refba0adba4._mode)
	x._unused2 = *(*[20]byte)(unsafe.Pointer(&x.refba0adba4._unused2))
}

// alloc_G_fpos_tMemory allocates memory for type C._G_fpos_t in C.
// The caller is responsible for freeing the this memory via C.free.
func alloc_G_fpos_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOf_G_fpos_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOf_G_fpos_tValue = unsafe.Sizeof([1]C._G_fpos_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *fpos_t) Ref() *C._G_fpos_t {
	if x == nil {
		return nil
	}
	return x.refc5e101c9
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *fpos_t) Free() {
	if x != nil && x.allocsc5e101c9 != nil {
		x.allocsc5e101c9.(*cgoAllocMap).Free()
		x.refc5e101c9 = nil
	}
}

// Newfpos_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newfpos_tRef(ref unsafe.Pointer) *fpos_t {
	if ref == nil {
		return nil
	}
	obj := new(fpos_t)
	obj.refc5e101c9 = (*C._G_fpos_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *fpos_t) PassRef() (*C._G_fpos_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc5e101c9 != nil {
		return x.refc5e101c9, nil
	}
	memc5e101c9 := alloc_G_fpos_tMemory(1)
	refc5e101c9 := (*C._G_fpos_t)(memc5e101c9)
	allocsc5e101c9 := new(cgoAllocMap)
	allocsc5e101c9.Add(memc5e101c9)

	var c__pos_allocs *cgoAllocMap
	refc5e101c9.__pos, c__pos_allocs = (C.__off_t)(x.__pos), cgoAllocsUnknown
	allocsc5e101c9.Borrow(c__pos_allocs)

	var c__state_allocs *cgoAllocMap
	refc5e101c9.__state, c__state_allocs = x.__state.PassValue()
	allocsc5e101c9.Borrow(c__state_allocs)

	x.refc5e101c9 = refc5e101c9
	x.allocsc5e101c9 = allocsc5e101c9
	return refc5e101c9, allocsc5e101c9

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x fpos_t) PassValue() (C._G_fpos_t, *cgoAllocMap) {
	if x.refc5e101c9 != nil {
		return *x.refc5e101c9, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *fpos_t) Deref() {
	if x.refc5e101c9 == nil {
		return
	}
	x.__pos = (__off_t)(x.refc5e101c9.__pos)
	x.__state = *New__mbstate_tRef(unsafe.Pointer(&x.refc5e101c9.__state))
}

// alloc__fsid_tMemory allocates memory for type C.__fsid_t in C.
// The caller is responsible for freeing the this memory via C.free.
func alloc__fsid_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOf__fsid_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOf__fsid_tValue = unsafe.Sizeof([1]C.__fsid_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *__fsid_t) Ref() *C.__fsid_t {
	if x == nil {
		return nil
	}
	return x.ref3fddfece
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *__fsid_t) Free() {
	if x != nil && x.allocs3fddfece != nil {
		x.allocs3fddfece.(*cgoAllocMap).Free()
		x.ref3fddfece = nil
	}
}

// New__fsid_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func New__fsid_tRef(ref unsafe.Pointer) *__fsid_t {
	if ref == nil {
		return nil
	}
	obj := new(__fsid_t)
	obj.ref3fddfece = (*C.__fsid_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *__fsid_t) PassRef() (*C.__fsid_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3fddfece != nil {
		return x.ref3fddfece, nil
	}
	mem3fddfece := alloc__fsid_tMemory(1)
	ref3fddfece := (*C.__fsid_t)(mem3fddfece)
	allocs3fddfece := new(cgoAllocMap)
	allocs3fddfece.Add(mem3fddfece)

	var c__val_allocs *cgoAllocMap
	ref3fddfece.__val, c__val_allocs = *(*[2]C.int)(unsafe.Pointer(&x.__val)), cgoAllocsUnknown
	allocs3fddfece.Borrow(c__val_allocs)

	x.ref3fddfece = ref3fddfece
	x.allocs3fddfece = allocs3fddfece
	return ref3fddfece, allocs3fddfece

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x __fsid_t) PassValue() (C.__fsid_t, *cgoAllocMap) {
	if x.ref3fddfece != nil {
		return *x.ref3fddfece, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *__fsid_t) Deref() {
	if x.ref3fddfece == nil {
		return
	}
	x.__val = *(*[2]int32)(unsafe.Pointer(&x.ref3fddfece.__val))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *_G_fpos_t) Ref() *C._G_fpos_t {
	if x == nil {
		return nil
	}
	return x.refc5e101c9
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *_G_fpos_t) Free() {
	if x != nil && x.allocsc5e101c9 != nil {
		x.allocsc5e101c9.(*cgoAllocMap).Free()
		x.refc5e101c9 = nil
	}
}

// New_G_fpos_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func New_G_fpos_tRef(ref unsafe.Pointer) *_G_fpos_t {
	if ref == nil {
		return nil
	}
	obj := new(_G_fpos_t)
	obj.refc5e101c9 = (*C._G_fpos_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *_G_fpos_t) PassRef() (*C._G_fpos_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc5e101c9 != nil {
		return x.refc5e101c9, nil
	}
	memc5e101c9 := alloc_G_fpos_tMemory(1)
	refc5e101c9 := (*C._G_fpos_t)(memc5e101c9)
	allocsc5e101c9 := new(cgoAllocMap)
	allocsc5e101c9.Add(memc5e101c9)

	var c__pos_allocs *cgoAllocMap
	refc5e101c9.__pos, c__pos_allocs = (C.__off_t)(x.__pos), cgoAllocsUnknown
	allocsc5e101c9.Borrow(c__pos_allocs)

	var c__state_allocs *cgoAllocMap
	refc5e101c9.__state, c__state_allocs = x.__state.PassValue()
	allocsc5e101c9.Borrow(c__state_allocs)

	x.refc5e101c9 = refc5e101c9
	x.allocsc5e101c9 = allocsc5e101c9
	return refc5e101c9, allocsc5e101c9

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x _G_fpos_t) PassValue() (C._G_fpos_t, *cgoAllocMap) {
	if x.refc5e101c9 != nil {
		return *x.refc5e101c9, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *_G_fpos_t) Deref() {
	if x.refc5e101c9 == nil {
		return
	}
	x.__pos = (__off_t)(x.refc5e101c9.__pos)
	x.__state = *New__mbstate_tRef(unsafe.Pointer(&x.refc5e101c9.__state))
}

// alloc_G_fpos64_tMemory allocates memory for type C._G_fpos64_t in C.
// The caller is responsible for freeing the this memory via C.free.
func alloc_G_fpos64_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOf_G_fpos64_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOf_G_fpos64_tValue = unsafe.Sizeof([1]C._G_fpos64_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *_G_fpos64_t) Ref() *C._G_fpos64_t {
	if x == nil {
		return nil
	}
	return x.ref70194a78
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *_G_fpos64_t) Free() {
	if x != nil && x.allocs70194a78 != nil {
		x.allocs70194a78.(*cgoAllocMap).Free()
		x.ref70194a78 = nil
	}
}

// New_G_fpos64_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func New_G_fpos64_tRef(ref unsafe.Pointer) *_G_fpos64_t {
	if ref == nil {
		return nil
	}
	obj := new(_G_fpos64_t)
	obj.ref70194a78 = (*C._G_fpos64_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *_G_fpos64_t) PassRef() (*C._G_fpos64_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref70194a78 != nil {
		return x.ref70194a78, nil
	}
	mem70194a78 := alloc_G_fpos64_tMemory(1)
	ref70194a78 := (*C._G_fpos64_t)(mem70194a78)
	allocs70194a78 := new(cgoAllocMap)
	allocs70194a78.Add(mem70194a78)

	var c__pos_allocs *cgoAllocMap
	ref70194a78.__pos, c__pos_allocs = (C.__off64_t)(x.__pos), cgoAllocsUnknown
	allocs70194a78.Borrow(c__pos_allocs)

	var c__state_allocs *cgoAllocMap
	ref70194a78.__state, c__state_allocs = x.__state.PassValue()
	allocs70194a78.Borrow(c__state_allocs)

	x.ref70194a78 = ref70194a78
	x.allocs70194a78 = allocs70194a78
	return ref70194a78, allocs70194a78

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x _G_fpos64_t) PassValue() (C._G_fpos64_t, *cgoAllocMap) {
	if x.ref70194a78 != nil {
		return *x.ref70194a78, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *_G_fpos64_t) Deref() {
	if x.ref70194a78 == nil {
		return
	}
	x.__pos = (__off64_t)(x.ref70194a78.__pos)
	x.__state = *New__mbstate_tRef(unsafe.Pointer(&x.ref70194a78.__state))
}

// alloc__mbstate_tMemory allocates memory for type C.__mbstate_t in C.
// The caller is responsible for freeing the this memory via C.free.
func alloc__mbstate_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOf__mbstate_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOf__mbstate_tValue = unsafe.Sizeof([1]C.__mbstate_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *__mbstate_t) Ref() *C.__mbstate_t {
	if x == nil {
		return nil
	}
	return x.refde5795dc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *__mbstate_t) Free() {
	if x != nil && x.allocsde5795dc != nil {
		x.allocsde5795dc.(*cgoAllocMap).Free()
		x.refde5795dc = nil
	}
}

// New__mbstate_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func New__mbstate_tRef(ref unsafe.Pointer) *__mbstate_t {
	if ref == nil {
		return nil
	}
	obj := new(__mbstate_t)
	obj.refde5795dc = (*C.__mbstate_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *__mbstate_t) PassRef() (*C.__mbstate_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refde5795dc != nil {
		return x.refde5795dc, nil
	}
	memde5795dc := alloc__mbstate_tMemory(1)
	refde5795dc := (*C.__mbstate_t)(memde5795dc)
	allocsde5795dc := new(cgoAllocMap)
	allocsde5795dc.Add(memde5795dc)

	var c__count_allocs *cgoAllocMap
	refde5795dc.__count, c__count_allocs = (C.int)(x.__count), cgoAllocsUnknown
	allocsde5795dc.Borrow(c__count_allocs)

	x.refde5795dc = refde5795dc
	x.allocsde5795dc = allocsde5795dc
	return refde5795dc, allocsde5795dc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x __mbstate_t) PassValue() (C.__mbstate_t, *cgoAllocMap) {
	if x.refde5795dc != nil {
		return *x.refde5795dc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *__mbstate_t) Deref() {
	if x.refde5795dc == nil {
		return
	}
	x.__count = (int32)(x.refde5795dc.__count)
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *fsid_t) Ref() *C.__fsid_t {
	if x == nil {
		return nil
	}
	return x.ref3fddfece
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *fsid_t) Free() {
	if x != nil && x.allocs3fddfece != nil {
		x.allocs3fddfece.(*cgoAllocMap).Free()
		x.ref3fddfece = nil
	}
}

// Newfsid_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newfsid_tRef(ref unsafe.Pointer) *fsid_t {
	if ref == nil {
		return nil
	}
	obj := new(fsid_t)
	obj.ref3fddfece = (*C.__fsid_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *fsid_t) PassRef() (*C.__fsid_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3fddfece != nil {
		return x.ref3fddfece, nil
	}
	mem3fddfece := alloc__fsid_tMemory(1)
	ref3fddfece := (*C.__fsid_t)(mem3fddfece)
	allocs3fddfece := new(cgoAllocMap)
	allocs3fddfece.Add(mem3fddfece)

	var c__val_allocs *cgoAllocMap
	ref3fddfece.__val, c__val_allocs = *(*[2]C.int)(unsafe.Pointer(&x.__val)), cgoAllocsUnknown
	allocs3fddfece.Borrow(c__val_allocs)

	x.ref3fddfece = ref3fddfece
	x.allocs3fddfece = allocs3fddfece
	return ref3fddfece, allocs3fddfece

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x fsid_t) PassValue() (C.__fsid_t, *cgoAllocMap) {
	if x.ref3fddfece != nil {
		return *x.ref3fddfece, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *fsid_t) Deref() {
	if x.ref3fddfece == nil {
		return
	}
	x.__val = *(*[2]int32)(unsafe.Pointer(&x.ref3fddfece.__val))
}

// alloc__sigset_tMemory allocates memory for type C.__sigset_t in C.
// The caller is responsible for freeing the this memory via C.free.
func alloc__sigset_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOf__sigset_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOf__sigset_tValue = unsafe.Sizeof([1]C.__sigset_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *sigset_t) Ref() *C.__sigset_t {
	if x == nil {
		return nil
	}
	return x.refa44732af
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *sigset_t) Free() {
	if x != nil && x.allocsa44732af != nil {
		x.allocsa44732af.(*cgoAllocMap).Free()
		x.refa44732af = nil
	}
}

// Newsigset_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newsigset_tRef(ref unsafe.Pointer) *sigset_t {
	if ref == nil {
		return nil
	}
	obj := new(sigset_t)
	obj.refa44732af = (*C.__sigset_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *sigset_t) PassRef() (*C.__sigset_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refa44732af != nil {
		return x.refa44732af, nil
	}
	mema44732af := alloc__sigset_tMemory(1)
	refa44732af := (*C.__sigset_t)(mema44732af)
	allocsa44732af := new(cgoAllocMap)
	allocsa44732af.Add(mema44732af)

	var c__val_allocs *cgoAllocMap
	refa44732af.__val, c__val_allocs = *(*[16]C.ulong)(unsafe.Pointer(&x.__val)), cgoAllocsUnknown
	allocsa44732af.Borrow(c__val_allocs)

	x.refa44732af = refa44732af
	x.allocsa44732af = allocsa44732af
	return refa44732af, allocsa44732af

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x sigset_t) PassValue() (C.__sigset_t, *cgoAllocMap) {
	if x.refa44732af != nil {
		return *x.refa44732af, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *sigset_t) Deref() {
	if x.refa44732af == nil {
		return
	}
	x.__val = *(*[16]uint)(unsafe.Pointer(&x.refa44732af.__val))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *__sigset_t) Ref() *C.__sigset_t {
	if x == nil {
		return nil
	}
	return x.refa44732af
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *__sigset_t) Free() {
	if x != nil && x.allocsa44732af != nil {
		x.allocsa44732af.(*cgoAllocMap).Free()
		x.refa44732af = nil
	}
}

// New__sigset_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func New__sigset_tRef(ref unsafe.Pointer) *__sigset_t {
	if ref == nil {
		return nil
	}
	obj := new(__sigset_t)
	obj.refa44732af = (*C.__sigset_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *__sigset_t) PassRef() (*C.__sigset_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refa44732af != nil {
		return x.refa44732af, nil
	}
	mema44732af := alloc__sigset_tMemory(1)
	refa44732af := (*C.__sigset_t)(mema44732af)
	allocsa44732af := new(cgoAllocMap)
	allocsa44732af.Add(mema44732af)

	var c__val_allocs *cgoAllocMap
	refa44732af.__val, c__val_allocs = *(*[16]C.ulong)(unsafe.Pointer(&x.__val)), cgoAllocsUnknown
	allocsa44732af.Borrow(c__val_allocs)

	x.refa44732af = refa44732af
	x.allocsa44732af = allocsa44732af
	return refa44732af, allocsa44732af

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x __sigset_t) PassValue() (C.__sigset_t, *cgoAllocMap) {
	if x.refa44732af != nil {
		return *x.refa44732af, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *__sigset_t) Deref() {
	if x.refa44732af == nil {
		return
	}
	x.__val = *(*[16]uint)(unsafe.Pointer(&x.refa44732af.__val))
}

// alloc__pthread_list_tMemory allocates memory for type C.__pthread_list_t in C.
// The caller is responsible for freeing the this memory via C.free.
func alloc__pthread_list_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOf__pthread_list_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOf__pthread_list_tValue = unsafe.Sizeof([1]C.__pthread_list_t{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *__pthread_list_t) Ref() *C.__pthread_list_t {
	if x == nil {
		return nil
	}
	return x.ref3eec1ba6
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *__pthread_list_t) Free() {
	if x != nil && x.allocs3eec1ba6 != nil {
		x.allocs3eec1ba6.(*cgoAllocMap).Free()
		x.ref3eec1ba6 = nil
	}
}

// New__pthread_list_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func New__pthread_list_tRef(ref unsafe.Pointer) *__pthread_list_t {
	if ref == nil {
		return nil
	}
	obj := new(__pthread_list_t)
	obj.ref3eec1ba6 = (*C.__pthread_list_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *__pthread_list_t) PassRef() (*C.__pthread_list_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3eec1ba6 != nil {
		return x.ref3eec1ba6, nil
	}
	mem3eec1ba6 := alloc__pthread_list_tMemory(1)
	ref3eec1ba6 := (*C.__pthread_list_t)(mem3eec1ba6)
	allocs3eec1ba6 := new(cgoAllocMap)
	allocs3eec1ba6.Add(mem3eec1ba6)

	x.ref3eec1ba6 = ref3eec1ba6
	x.allocs3eec1ba6 = allocs3eec1ba6
	return ref3eec1ba6, allocs3eec1ba6

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x __pthread_list_t) PassValue() (C.__pthread_list_t, *cgoAllocMap) {
	if x.ref3eec1ba6 != nil {
		return *x.ref3eec1ba6, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *__pthread_list_t) Deref() {
	if x.ref3eec1ba6 == nil {
		return
	}
}

// alloc__locale_tMemory allocates memory for type C.__locale_t in C.
// The caller is responsible for freeing the this memory via C.free.
func alloc__locale_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOf__locale_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOf__locale_tValue = unsafe.Sizeof([1]C.__locale_t{})

// allocA13PCharMemory allocates memory for type [13]*C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocA13PCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfA13PCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfA13PCharValue = unsafe.Sizeof([1][13]*C.char{})

// unpackA13String transforms a sliced Go data structure into plain C format.
func unpackA13String(x [13]string) (unpacked [13]*C.char, allocs *cgoAllocMap) {
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(*[13]*C.char) {
		go allocs.Free()
	})

	mem0 := allocA13PCharMemory(1)
	allocs.Add(mem0)
	v0 := (*[13]*C.char)(mem0)
	for i0 := range x {
		v0[i0], _ = unpackPCharString(x[i0])
	}
	unpacked = *(*[13]*C.char)(mem0)
	return
}

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// packA13String reads sliced Go data structure out from plain C format.
func packA13String(v *[13]string, ptr0 *[13]*C.char) {
	for i0 := range v {
		ptr1 := ptr0[i0]
		v[i0] = packPCharString(ptr1)
	}
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *__locale_t) Ref() *C.__locale_t {
	if x == nil {
		return nil
	}
	return x.refa5868e3f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *__locale_t) Free() {
	if x != nil && x.allocsa5868e3f != nil {
		x.allocsa5868e3f.(*cgoAllocMap).Free()
		x.refa5868e3f = nil
	}
}

// New__locale_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func New__locale_tRef(ref unsafe.Pointer) *__locale_t {
	if ref == nil {
		return nil
	}
	obj := new(__locale_t)
	obj.refa5868e3f = (*C.__locale_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *__locale_t) PassRef() (*C.__locale_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refa5868e3f != nil {
		return x.refa5868e3f, nil
	}
	mema5868e3f := alloc__locale_tMemory(1)
	refa5868e3f := (*C.__locale_t)(mema5868e3f)
	allocsa5868e3f := new(cgoAllocMap)
	allocsa5868e3f.Add(mema5868e3f)

	var c__ctype_b_allocs *cgoAllocMap
	refa5868e3f.__ctype_b, c__ctype_b_allocs = (*C.ushort)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.__ctype_b)).Data)), cgoAllocsUnknown
	allocsa5868e3f.Borrow(c__ctype_b_allocs)

	var c__ctype_tolower_allocs *cgoAllocMap
	refa5868e3f.__ctype_tolower, c__ctype_tolower_allocs = (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.__ctype_tolower)).Data)), cgoAllocsUnknown
	allocsa5868e3f.Borrow(c__ctype_tolower_allocs)

	var c__ctype_toupper_allocs *cgoAllocMap
	refa5868e3f.__ctype_toupper, c__ctype_toupper_allocs = (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.__ctype_toupper)).Data)), cgoAllocsUnknown
	allocsa5868e3f.Borrow(c__ctype_toupper_allocs)

	var c__names_allocs *cgoAllocMap
	refa5868e3f.__names, c__names_allocs = unpackA13String(x.__names)
	allocsa5868e3f.Borrow(c__names_allocs)

	x.refa5868e3f = refa5868e3f
	x.allocsa5868e3f = allocsa5868e3f
	return refa5868e3f, allocsa5868e3f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x __locale_t) PassValue() (C.__locale_t, *cgoAllocMap) {
	if x.refa5868e3f != nil {
		return *x.refa5868e3f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *__locale_t) Deref() {
	if x.refa5868e3f == nil {
		return
	}
	hxfa5d4a2 := (*sliceHeader)(unsafe.Pointer(&x.__ctype_b))
	hxfa5d4a2.Data = uintptr(unsafe.Pointer(x.refa5868e3f.__ctype_b))
	hxfa5d4a2.Cap = 0x7fffffff
	// hxfa5d4a2.Len = ?

	hxfa76696 := (*sliceHeader)(unsafe.Pointer(&x.__ctype_tolower))
	hxfa76696.Data = uintptr(unsafe.Pointer(x.refa5868e3f.__ctype_tolower))
	hxfa76696.Cap = 0x7fffffff
	// hxfa76696.Len = ?

	hxf5f87e7 := (*sliceHeader)(unsafe.Pointer(&x.__ctype_toupper))
	hxf5f87e7.Data = uintptr(unsafe.Pointer(x.refa5868e3f.__ctype_toupper))
	hxf5f87e7.Cap = 0x7fffffff
	// hxf5f87e7.Len = ?

	packA13String(&x.__names, (*[13]*C.char)(unsafe.Pointer(&x.refa5868e3f.__names)))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *locale_t) Ref() *C.__locale_t {
	if x == nil {
		return nil
	}
	return x.refa5868e3f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *locale_t) Free() {
	if x != nil && x.allocsa5868e3f != nil {
		x.allocsa5868e3f.(*cgoAllocMap).Free()
		x.refa5868e3f = nil
	}
}

// Newlocale_tRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func Newlocale_tRef(ref unsafe.Pointer) *locale_t {
	if ref == nil {
		return nil
	}
	obj := new(locale_t)
	obj.refa5868e3f = (*C.__locale_t)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *locale_t) PassRef() (*C.__locale_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refa5868e3f != nil {
		return x.refa5868e3f, nil
	}
	mema5868e3f := alloc__locale_tMemory(1)
	refa5868e3f := (*C.__locale_t)(mema5868e3f)
	allocsa5868e3f := new(cgoAllocMap)
	allocsa5868e3f.Add(mema5868e3f)

	var c__ctype_b_allocs *cgoAllocMap
	refa5868e3f.__ctype_b, c__ctype_b_allocs = (*C.ushort)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.__ctype_b)).Data)), cgoAllocsUnknown
	allocsa5868e3f.Borrow(c__ctype_b_allocs)

	var c__ctype_tolower_allocs *cgoAllocMap
	refa5868e3f.__ctype_tolower, c__ctype_tolower_allocs = (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.__ctype_tolower)).Data)), cgoAllocsUnknown
	allocsa5868e3f.Borrow(c__ctype_tolower_allocs)

	var c__ctype_toupper_allocs *cgoAllocMap
	refa5868e3f.__ctype_toupper, c__ctype_toupper_allocs = (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.__ctype_toupper)).Data)), cgoAllocsUnknown
	allocsa5868e3f.Borrow(c__ctype_toupper_allocs)

	var c__names_allocs *cgoAllocMap
	refa5868e3f.__names, c__names_allocs = unpackA13String(x.__names)
	allocsa5868e3f.Borrow(c__names_allocs)

	x.refa5868e3f = refa5868e3f
	x.allocsa5868e3f = allocsa5868e3f
	return refa5868e3f, allocsa5868e3f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x locale_t) PassValue() (C.__locale_t, *cgoAllocMap) {
	if x.refa5868e3f != nil {
		return *x.refa5868e3f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *locale_t) Deref() {
	if x.refa5868e3f == nil {
		return
	}
	hxfab9202 := (*sliceHeader)(unsafe.Pointer(&x.__ctype_b))
	hxfab9202.Data = uintptr(unsafe.Pointer(x.refa5868e3f.__ctype_b))
	hxfab9202.Cap = 0x7fffffff
	// hxfab9202.Len = ?

	hxf37812b := (*sliceHeader)(unsafe.Pointer(&x.__ctype_tolower))
	hxf37812b.Data = uintptr(unsafe.Pointer(x.refa5868e3f.__ctype_tolower))
	hxf37812b.Cap = 0x7fffffff
	// hxf37812b.Len = ?

	hxfe0b820 := (*sliceHeader)(unsafe.Pointer(&x.__ctype_toupper))
	hxfe0b820.Data = uintptr(unsafe.Pointer(x.refa5868e3f.__ctype_toupper))
	hxfe0b820.Cap = 0x7fffffff
	// hxfe0b820.Len = ?

	packA13String(&x.__names, (*[13]*C.char)(unsafe.Pointer(&x.refa5868e3f.__names)))
}

// unpackArgSFILE transforms a sliced Go data structure into plain C format.
func unpackArgSFILE(x []FILE) (unpacked *C.FILE, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.FILE) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocFILEMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.FILE)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.FILE)(unsafe.Pointer(h.Data))
	return
}

// packSFILE reads sliced Go data structure out from plain C format.
func packSFILE(v []FILE, ptr0 *C.FILE) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfFILEValue]C.FILE)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewFILERef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSBlock_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSBlock_info_msg_t(x []block_info_msg_t) (unpacked *C.block_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.block_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocBlock_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.block_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.block_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSBlock_info_msg_t reads sliced Go data structure out from plain C format.
func packSBlock_info_msg_t(v []block_info_msg_t, ptr0 *C.block_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfBlock_info_msg_tValue]C.block_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newblock_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSBlock_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSBlock_info_t(x []block_info_t) (unpacked *C.block_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.block_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocBlock_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.block_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.block_info_t)(unsafe.Pointer(h.Data))
	return
}

// allocPBlock_info_msg_tMemory allocates memory for type *C.block_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPBlock_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPBlock_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPBlock_info_msg_tValue = unsafe.Sizeof([1]*C.block_info_msg_t{})

// unpackArgSSBlock_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSBlock_info_msg_t(x [][]block_info_msg_t) (unpacked **C.block_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.block_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPBlock_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.block_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocBlock_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.block_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.block_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.block_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSBlock_info_msg_t reads sliced Go data structure out from plain C format.
func packSSBlock_info_msg_t(v [][]block_info_msg_t, ptr0 **C.block_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.block_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfBlock_info_msg_tValue]C.block_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newblock_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// allocUpdate_block_msg_tMemory allocates memory for type C.update_block_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocUpdate_block_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfUpdate_block_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfUpdate_block_msg_tValue = unsafe.Sizeof([1]C.update_block_msg_t{})

// unpackArgSUpdate_block_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSUpdate_block_msg_t(x []update_block_msg_t) (unpacked *C.update_block_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.update_block_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocUpdate_block_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.update_block_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.update_block_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSUpdate_block_msg_t reads sliced Go data structure out from plain C format.
func packSUpdate_block_msg_t(v []update_block_msg_t, ptr0 *C.update_block_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfUpdate_block_msg_tValue]C.update_block_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newupdate_block_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSJob_desc_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSJob_desc_msg_t(x []job_desc_msg_t) (unpacked *C.job_desc_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_desc_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_desc_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_desc_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_desc_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSJob_desc_msg_t reads sliced Go data structure out from plain C format.
func packSJob_desc_msg_t(v []job_desc_msg_t, ptr0 *C.job_desc_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfJob_desc_msg_tValue]C.job_desc_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newjob_desc_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPResource_allocation_response_msg_tMemory allocates memory for type *C.resource_allocation_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPResource_allocation_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPResource_allocation_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPResource_allocation_response_msg_tValue = unsafe.Sizeof([1]*C.resource_allocation_response_msg_t{})

// unpackArgSSResource_allocation_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSResource_allocation_response_msg_t(x [][]resource_allocation_response_msg_t) (unpacked **C.resource_allocation_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.resource_allocation_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPResource_allocation_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.resource_allocation_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocResource_allocation_response_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.resource_allocation_response_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.resource_allocation_response_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.resource_allocation_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSResource_allocation_response_msg_t reads sliced Go data structure out from plain C format.
func packSSResource_allocation_response_msg_t(v [][]resource_allocation_response_msg_t, ptr0 **C.resource_allocation_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.resource_allocation_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfResource_allocation_response_msg_tValue]C.resource_allocation_response_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newresource_allocation_response_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSResource_allocation_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSResource_allocation_response_msg_t(x []resource_allocation_response_msg_t) (unpacked *C.resource_allocation_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.resource_allocation_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocResource_allocation_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.resource_allocation_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.resource_allocation_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSResource_allocation_response_msg_t reads sliced Go data structure out from plain C format.
func packSResource_allocation_response_msg_t(v []resource_allocation_response_msg_t, ptr0 *C.resource_allocation_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfResource_allocation_response_msg_tValue]C.resource_allocation_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newresource_allocation_response_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSSlurm_allocation_callbacks_t transforms a sliced Go data structure into plain C format.
func unpackArgSSlurm_allocation_callbacks_t(x []slurm_allocation_callbacks_t) (unpacked *C.slurm_allocation_callbacks_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurm_allocation_callbacks_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurm_allocation_callbacks_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurm_allocation_callbacks_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurm_allocation_callbacks_t)(unsafe.Pointer(h.Data))
	return
}

// packSSlurm_allocation_callbacks_t reads sliced Go data structure out from plain C format.
func packSSlurm_allocation_callbacks_t(v []slurm_allocation_callbacks_t, ptr0 *C.slurm_allocation_callbacks_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSlurm_allocation_callbacks_tValue]C.slurm_allocation_callbacks_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newslurm_allocation_callbacks_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPSubmit_response_msg_tMemory allocates memory for type *C.submit_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPSubmit_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPSubmit_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPSubmit_response_msg_tValue = unsafe.Sizeof([1]*C.submit_response_msg_t{})

// unpackArgSSSubmit_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSSubmit_response_msg_t(x [][]submit_response_msg_t) (unpacked **C.submit_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.submit_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPSubmit_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.submit_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocSubmit_response_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.submit_response_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.submit_response_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.submit_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSSubmit_response_msg_t reads sliced Go data structure out from plain C format.
func packSSSubmit_response_msg_t(v [][]submit_response_msg_t, ptr0 **C.submit_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.submit_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfSubmit_response_msg_tValue]C.submit_response_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newsubmit_response_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSSubmit_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSubmit_response_msg_t(x []submit_response_msg_t) (unpacked *C.submit_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.submit_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSubmit_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.submit_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.submit_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSubmit_response_msg_t reads sliced Go data structure out from plain C format.
func packSSubmit_response_msg_t(v []submit_response_msg_t, ptr0 *C.submit_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSubmit_response_msg_tValue]C.submit_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newsubmit_response_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPWill_run_response_msg_tMemory allocates memory for type *C.will_run_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPWill_run_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPWill_run_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPWill_run_response_msg_tValue = unsafe.Sizeof([1]*C.will_run_response_msg_t{})

// unpackArgSSWill_run_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSWill_run_response_msg_t(x [][]will_run_response_msg_t) (unpacked **C.will_run_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.will_run_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPWill_run_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.will_run_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocWill_run_response_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.will_run_response_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.will_run_response_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.will_run_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSWill_run_response_msg_t reads sliced Go data structure out from plain C format.
func packSSWill_run_response_msg_t(v [][]will_run_response_msg_t, ptr0 **C.will_run_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.will_run_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfWill_run_response_msg_tValue]C.will_run_response_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newwill_run_response_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// allocPJob_sbcast_cred_msg_tMemory allocates memory for type *C.job_sbcast_cred_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPJob_sbcast_cred_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPJob_sbcast_cred_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPJob_sbcast_cred_msg_tValue = unsafe.Sizeof([1]*C.job_sbcast_cred_msg_t{})

// unpackArgSSJob_sbcast_cred_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSJob_sbcast_cred_msg_t(x [][]job_sbcast_cred_msg_t) (unpacked **C.job_sbcast_cred_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.job_sbcast_cred_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPJob_sbcast_cred_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.job_sbcast_cred_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocJob_sbcast_cred_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.job_sbcast_cred_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.job_sbcast_cred_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.job_sbcast_cred_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSJob_sbcast_cred_msg_t reads sliced Go data structure out from plain C format.
func packSSJob_sbcast_cred_msg_t(v [][]job_sbcast_cred_msg_t, ptr0 **C.job_sbcast_cred_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.job_sbcast_cred_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfJob_sbcast_cred_msg_tValue]C.job_sbcast_cred_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newjob_sbcast_cred_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSJob_sbcast_cred_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSJob_sbcast_cred_msg_t(x []job_sbcast_cred_msg_t) (unpacked *C.job_sbcast_cred_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_sbcast_cred_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_sbcast_cred_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_sbcast_cred_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_sbcast_cred_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSJob_sbcast_cred_msg_t reads sliced Go data structure out from plain C format.
func packSJob_sbcast_cred_msg_t(v []job_sbcast_cred_msg_t, ptr0 *C.job_sbcast_cred_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfJob_sbcast_cred_msg_tValue]C.job_sbcast_cred_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newjob_sbcast_cred_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPLicense_info_msg_tMemory allocates memory for type *C.license_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPLicense_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPLicense_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPLicense_info_msg_tValue = unsafe.Sizeof([1]*C.license_info_msg_t{})

// unpackArgSSLicense_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSLicense_info_msg_t(x [][]license_info_msg_t) (unpacked **C.license_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.license_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPLicense_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.license_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocLicense_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.license_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.license_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.license_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSLicense_info_msg_t reads sliced Go data structure out from plain C format.
func packSSLicense_info_msg_t(v [][]license_info_msg_t, ptr0 **C.license_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.license_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfLicense_info_msg_tValue]C.license_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newlicense_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSLicense_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSLicense_info_msg_t(x []license_info_msg_t) (unpacked *C.license_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.license_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocLicense_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.license_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.license_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSLicense_info_msg_t reads sliced Go data structure out from plain C format.
func packSLicense_info_msg_t(v []license_info_msg_t, ptr0 *C.license_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfLicense_info_msg_tValue]C.license_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newlicense_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSAssoc_mgr_info_request_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSAssoc_mgr_info_request_msg_t(x []assoc_mgr_info_request_msg_t) (unpacked *C.assoc_mgr_info_request_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.assoc_mgr_info_request_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocAssoc_mgr_info_request_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.assoc_mgr_info_request_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.assoc_mgr_info_request_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSAssoc_mgr_info_request_msg_t reads sliced Go data structure out from plain C format.
func packSAssoc_mgr_info_request_msg_t(v []assoc_mgr_info_request_msg_t, ptr0 *C.assoc_mgr_info_request_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfAssoc_mgr_info_request_msg_tValue]C.assoc_mgr_info_request_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newassoc_mgr_info_request_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPAssoc_mgr_info_msg_tMemory allocates memory for type *C.assoc_mgr_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPAssoc_mgr_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPAssoc_mgr_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPAssoc_mgr_info_msg_tValue = unsafe.Sizeof([1]*C.assoc_mgr_info_msg_t{})

// unpackArgSSAssoc_mgr_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSAssoc_mgr_info_msg_t(x [][]assoc_mgr_info_msg_t) (unpacked **C.assoc_mgr_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.assoc_mgr_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPAssoc_mgr_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.assoc_mgr_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocAssoc_mgr_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.assoc_mgr_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.assoc_mgr_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.assoc_mgr_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSAssoc_mgr_info_msg_t reads sliced Go data structure out from plain C format.
func packSSAssoc_mgr_info_msg_t(v [][]assoc_mgr_info_msg_t, ptr0 **C.assoc_mgr_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.assoc_mgr_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfAssoc_mgr_info_msg_tValue]C.assoc_mgr_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newassoc_mgr_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSAssoc_mgr_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSAssoc_mgr_info_msg_t(x []assoc_mgr_info_msg_t) (unpacked *C.assoc_mgr_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.assoc_mgr_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocAssoc_mgr_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.assoc_mgr_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.assoc_mgr_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSAssoc_mgr_info_msg_t reads sliced Go data structure out from plain C format.
func packSAssoc_mgr_info_msg_t(v []assoc_mgr_info_msg_t, ptr0 *C.assoc_mgr_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfAssoc_mgr_info_msg_tValue]C.assoc_mgr_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newassoc_mgr_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSJob_step_kill_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSJob_step_kill_msg_t(x []job_step_kill_msg_t) (unpacked *C.job_step_kill_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_step_kill_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_step_kill_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_step_kill_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_step_kill_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSJob_step_kill_msg_t reads sliced Go data structure out from plain C format.
func packSJob_step_kill_msg_t(v []job_step_kill_msg_t, ptr0 *C.job_step_kill_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfJob_step_kill_msg_tValue]C.job_step_kill_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newjob_step_kill_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSSlurm_step_ctx_params_t transforms a sliced Go data structure into plain C format.
func unpackArgSSlurm_step_ctx_params_t(x []slurm_step_ctx_params_t) (unpacked *C.slurm_step_ctx_params_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurm_step_ctx_params_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurm_step_ctx_params_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurm_step_ctx_params_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurm_step_ctx_params_t)(unsafe.Pointer(h.Data))
	return
}

// packSSlurm_step_ctx_params_t reads sliced Go data structure out from plain C format.
func packSSlurm_step_ctx_params_t(v []slurm_step_ctx_params_t, ptr0 *C.slurm_step_ctx_params_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSlurm_step_ctx_params_tValue]C.slurm_step_ctx_params_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newslurm_step_ctx_params_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSDynamic_plugin_data_t transforms a sliced Go data structure into plain C format.
func unpackArgSDynamic_plugin_data_t(x []dynamic_plugin_data_t) (unpacked *C.dynamic_plugin_data_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.dynamic_plugin_data_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocDynamic_plugin_data_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.dynamic_plugin_data_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.dynamic_plugin_data_t)(unsafe.Pointer(h.Data))
	return
}

// unpackArgSSlurm_step_launch_params_t transforms a sliced Go data structure into plain C format.
func unpackArgSSlurm_step_launch_params_t(x []slurm_step_launch_params_t) (unpacked *C.slurm_step_launch_params_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurm_step_launch_params_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurm_step_launch_params_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurm_step_launch_params_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurm_step_launch_params_t)(unsafe.Pointer(h.Data))
	return
}

// packSSlurm_step_launch_params_t reads sliced Go data structure out from plain C format.
func packSSlurm_step_launch_params_t(v []slurm_step_launch_params_t, ptr0 *C.slurm_step_launch_params_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSlurm_step_launch_params_tValue]C.slurm_step_launch_params_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newslurm_step_launch_params_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSSlurm_step_launch_callbacks_t transforms a sliced Go data structure into plain C format.
func unpackArgSSlurm_step_launch_callbacks_t(x []slurm_step_launch_callbacks_t) (unpacked *C.slurm_step_launch_callbacks_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurm_step_launch_callbacks_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurm_step_launch_callbacks_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurm_step_launch_callbacks_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurm_step_launch_callbacks_t)(unsafe.Pointer(h.Data))
	return
}

// packSSlurm_step_launch_callbacks_t reads sliced Go data structure out from plain C format.
func packSSlurm_step_launch_callbacks_t(v []slurm_step_launch_callbacks_t, ptr0 *C.slurm_step_launch_callbacks_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSlurm_step_launch_callbacks_tValue]C.slurm_step_launch_callbacks_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newslurm_step_launch_callbacks_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPSlurm_ctl_conf_tMemory allocates memory for type *C.slurm_ctl_conf_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPSlurm_ctl_conf_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPSlurm_ctl_conf_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPSlurm_ctl_conf_tValue = unsafe.Sizeof([1]*C.slurm_ctl_conf_t{})

// unpackArgSSSlurm_ctl_conf_t transforms a sliced Go data structure into plain C format.
func unpackArgSSSlurm_ctl_conf_t(x [][]slurm_ctl_conf_t) (unpacked **C.slurm_ctl_conf_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.slurm_ctl_conf_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPSlurm_ctl_conf_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.slurm_ctl_conf_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocSlurm_ctl_conf_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.slurm_ctl_conf_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.slurm_ctl_conf_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.slurm_ctl_conf_t)(unsafe.Pointer(h.Data))
	return
}

// packSSSlurm_ctl_conf_t reads sliced Go data structure out from plain C format.
func packSSSlurm_ctl_conf_t(v [][]slurm_ctl_conf_t, ptr0 **C.slurm_ctl_conf_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.slurm_ctl_conf_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfSlurm_ctl_conf_tValue]C.slurm_ctl_conf_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newslurm_ctl_conf_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSSlurm_ctl_conf_t transforms a sliced Go data structure into plain C format.
func unpackArgSSlurm_ctl_conf_t(x []slurm_ctl_conf_t) (unpacked *C.slurm_ctl_conf_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurm_ctl_conf_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurm_ctl_conf_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurm_ctl_conf_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurm_ctl_conf_t)(unsafe.Pointer(h.Data))
	return
}

// packSSlurm_ctl_conf_t reads sliced Go data structure out from plain C format.
func packSSlurm_ctl_conf_t(v []slurm_ctl_conf_t, ptr0 *C.slurm_ctl_conf_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSlurm_ctl_conf_tValue]C.slurm_ctl_conf_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newslurm_ctl_conf_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSNode_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSNode_info_msg_t(x []node_info_msg_t) (unpacked *C.node_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.node_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocNode_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.node_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.node_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSNode_info_msg_t reads sliced Go data structure out from plain C format.
func packSNode_info_msg_t(v []node_info_msg_t, ptr0 *C.node_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfNode_info_msg_tValue]C.node_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newnode_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPartition_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSPartition_info_msg_t(x []partition_info_msg_t) (unpacked *C.partition_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.partition_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPartition_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.partition_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.partition_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSPartition_info_msg_t reads sliced Go data structure out from plain C format.
func packSPartition_info_msg_t(v []partition_info_msg_t, ptr0 *C.partition_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPartition_info_msg_tValue]C.partition_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newpartition_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPSlurmd_status_tMemory allocates memory for type *C.slurmd_status_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPSlurmd_status_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPSlurmd_status_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPSlurmd_status_tValue = unsafe.Sizeof([1]*C.slurmd_status_t{})

// unpackArgSSSlurmd_status_t transforms a sliced Go data structure into plain C format.
func unpackArgSSSlurmd_status_t(x [][]slurmd_status_t) (unpacked **C.slurmd_status_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.slurmd_status_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPSlurmd_status_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.slurmd_status_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocSlurmd_status_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.slurmd_status_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.slurmd_status_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.slurmd_status_t)(unsafe.Pointer(h.Data))
	return
}

// packSSSlurmd_status_t reads sliced Go data structure out from plain C format.
func packSSSlurmd_status_t(v [][]slurmd_status_t, ptr0 **C.slurmd_status_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.slurmd_status_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfSlurmd_status_tValue]C.slurmd_status_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newslurmd_status_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSSlurmd_status_t transforms a sliced Go data structure into plain C format.
func unpackArgSSlurmd_status_t(x []slurmd_status_t) (unpacked *C.slurmd_status_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurmd_status_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurmd_status_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurmd_status_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurmd_status_t)(unsafe.Pointer(h.Data))
	return
}

// packSSlurmd_status_t reads sliced Go data structure out from plain C format.
func packSSlurmd_status_t(v []slurmd_status_t, ptr0 *C.slurmd_status_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSlurmd_status_tValue]C.slurmd_status_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newslurmd_status_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSStep_update_request_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSStep_update_request_msg_t(x []step_update_request_msg_t) (unpacked *C.step_update_request_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.step_update_request_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStep_update_request_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.step_update_request_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.step_update_request_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSStep_update_request_msg_t reads sliced Go data structure out from plain C format.
func packSStep_update_request_msg_t(v []step_update_request_msg_t, ptr0 *C.step_update_request_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStep_update_request_msg_tValue]C.step_update_request_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newstep_update_request_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPStats_info_response_msg_tMemory allocates memory for type *C.stats_info_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPStats_info_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPStats_info_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPStats_info_response_msg_tValue = unsafe.Sizeof([1]*C.stats_info_response_msg_t{})

// unpackArgSSStats_info_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSStats_info_response_msg_t(x [][]stats_info_response_msg_t) (unpacked **C.stats_info_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.stats_info_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPStats_info_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.stats_info_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocStats_info_response_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.stats_info_response_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.stats_info_response_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.stats_info_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSStats_info_response_msg_t reads sliced Go data structure out from plain C format.
func packSSStats_info_response_msg_t(v [][]stats_info_response_msg_t, ptr0 **C.stats_info_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.stats_info_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfStats_info_response_msg_tValue]C.stats_info_response_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newstats_info_response_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSStats_info_request_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSStats_info_request_msg_t(x []stats_info_request_msg_t) (unpacked *C.stats_info_request_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.stats_info_request_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStats_info_request_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.stats_info_request_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.stats_info_request_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSStats_info_request_msg_t reads sliced Go data structure out from plain C format.
func packSStats_info_request_msg_t(v []stats_info_request_msg_t, ptr0 *C.stats_info_request_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStats_info_request_msg_tValue]C.stats_info_request_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newstats_info_request_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSJob_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSJob_info_msg_t(x []job_info_msg_t) (unpacked *C.job_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSJob_info_msg_t reads sliced Go data structure out from plain C format.
func packSJob_info_msg_t(v []job_info_msg_t, ptr0 *C.job_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfJob_info_msg_tValue]C.job_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newjob_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPriority_factors_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSPriority_factors_response_msg_t(x []priority_factors_response_msg_t) (unpacked *C.priority_factors_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.priority_factors_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPriority_factors_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.priority_factors_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.priority_factors_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSPriority_factors_response_msg_t reads sliced Go data structure out from plain C format.
func packSPriority_factors_response_msg_t(v []priority_factors_response_msg_t, ptr0 *C.priority_factors_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPriority_factors_response_msg_tValue]C.priority_factors_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newpriority_factors_response_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocJob_info_tMemory allocates memory for type C.job_info_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocJob_info_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfJob_info_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfJob_info_tValue = unsafe.Sizeof([1]C.job_info_t{})

// unpackArgSJob_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSJob_info_t(x []job_info_t) (unpacked *C.job_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_info_t)(unsafe.Pointer(h.Data))
	return
}

// packSJob_info_t reads sliced Go data structure out from plain C format.
func packSJob_info_t(v []job_info_t, ptr0 *C.job_info_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfJob_info_tValue]C.job_info_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newjob_info_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPJob_info_msg_tMemory allocates memory for type *C.job_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPJob_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPJob_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPJob_info_msg_tValue = unsafe.Sizeof([1]*C.job_info_msg_t{})

// unpackArgSSJob_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSJob_info_msg_t(x [][]job_info_msg_t) (unpacked **C.job_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.job_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPJob_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.job_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocJob_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.job_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.job_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.job_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSJob_info_msg_t reads sliced Go data structure out from plain C format.
func packSSJob_info_msg_t(v [][]job_info_msg_t, ptr0 **C.job_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.job_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfJob_info_msg_tValue]C.job_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newjob_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// allocPPriority_factors_response_msg_tMemory allocates memory for type *C.priority_factors_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPPriority_factors_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPPriority_factors_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPPriority_factors_response_msg_tValue = unsafe.Sizeof([1]*C.priority_factors_response_msg_t{})

// unpackArgSSPriority_factors_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSPriority_factors_response_msg_t(x [][]priority_factors_response_msg_t) (unpacked **C.priority_factors_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.priority_factors_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPPriority_factors_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.priority_factors_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocPriority_factors_response_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.priority_factors_response_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.priority_factors_response_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.priority_factors_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSPriority_factors_response_msg_t reads sliced Go data structure out from plain C format.
func packSSPriority_factors_response_msg_t(v [][]priority_factors_response_msg_t, ptr0 **C.priority_factors_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.priority_factors_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfPriority_factors_response_msg_tValue]C.priority_factors_response_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newpriority_factors_response_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSSlurm_job_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSSlurm_job_info_t(x []slurm_job_info_t) (unpacked *C.slurm_job_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurm_job_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurm_job_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurm_job_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurm_job_info_t)(unsafe.Pointer(h.Data))
	return
}

// allocPJob_array_resp_msg_tMemory allocates memory for type *C.job_array_resp_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPJob_array_resp_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPJob_array_resp_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPJob_array_resp_msg_tValue = unsafe.Sizeof([1]*C.job_array_resp_msg_t{})

// unpackArgSSJob_array_resp_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSJob_array_resp_msg_t(x [][]job_array_resp_msg_t) (unpacked **C.job_array_resp_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.job_array_resp_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPJob_array_resp_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.job_array_resp_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocJob_array_resp_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.job_array_resp_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.job_array_resp_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.job_array_resp_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSJob_array_resp_msg_t reads sliced Go data structure out from plain C format.
func packSSJob_array_resp_msg_t(v [][]job_array_resp_msg_t, ptr0 **C.job_array_resp_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.job_array_resp_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfJob_array_resp_msg_tValue]C.job_array_resp_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newjob_array_resp_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// allocPJob_step_info_response_msg_tMemory allocates memory for type *C.job_step_info_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPJob_step_info_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPJob_step_info_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPJob_step_info_response_msg_tValue = unsafe.Sizeof([1]*C.job_step_info_response_msg_t{})

// unpackArgSSJob_step_info_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSJob_step_info_response_msg_t(x [][]job_step_info_response_msg_t) (unpacked **C.job_step_info_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.job_step_info_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPJob_step_info_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.job_step_info_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocJob_step_info_response_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.job_step_info_response_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.job_step_info_response_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.job_step_info_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSJob_step_info_response_msg_t reads sliced Go data structure out from plain C format.
func packSSJob_step_info_response_msg_t(v [][]job_step_info_response_msg_t, ptr0 **C.job_step_info_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.job_step_info_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfJob_step_info_response_msg_tValue]C.job_step_info_response_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newjob_step_info_response_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSJob_step_info_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSJob_step_info_response_msg_t(x []job_step_info_response_msg_t) (unpacked *C.job_step_info_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_step_info_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_step_info_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_step_info_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_step_info_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSJob_step_info_response_msg_t reads sliced Go data structure out from plain C format.
func packSJob_step_info_response_msg_t(v []job_step_info_response_msg_t, ptr0 *C.job_step_info_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfJob_step_info_response_msg_tValue]C.job_step_info_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newjob_step_info_response_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSJob_step_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSJob_step_info_t(x []job_step_info_t) (unpacked *C.job_step_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_step_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_step_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_step_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_step_info_t)(unsafe.Pointer(h.Data))
	return
}

// allocPJob_step_stat_response_msg_tMemory allocates memory for type *C.job_step_stat_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPJob_step_stat_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPJob_step_stat_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPJob_step_stat_response_msg_tValue = unsafe.Sizeof([1]*C.job_step_stat_response_msg_t{})

// unpackArgSSJob_step_stat_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSJob_step_stat_response_msg_t(x [][]job_step_stat_response_msg_t) (unpacked **C.job_step_stat_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.job_step_stat_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPJob_step_stat_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.job_step_stat_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocJob_step_stat_response_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.job_step_stat_response_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.job_step_stat_response_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.job_step_stat_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSJob_step_stat_response_msg_t reads sliced Go data structure out from plain C format.
func packSSJob_step_stat_response_msg_t(v [][]job_step_stat_response_msg_t, ptr0 **C.job_step_stat_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.job_step_stat_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfJob_step_stat_response_msg_tValue]C.job_step_stat_response_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newjob_step_stat_response_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// allocPJob_step_pids_response_msg_tMemory allocates memory for type *C.job_step_pids_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPJob_step_pids_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPJob_step_pids_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPJob_step_pids_response_msg_tValue = unsafe.Sizeof([1]*C.job_step_pids_response_msg_t{})

// unpackArgSSJob_step_pids_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSJob_step_pids_response_msg_t(x [][]job_step_pids_response_msg_t) (unpacked **C.job_step_pids_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.job_step_pids_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPJob_step_pids_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.job_step_pids_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocJob_step_pids_response_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.job_step_pids_response_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.job_step_pids_response_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.job_step_pids_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSJob_step_pids_response_msg_t reads sliced Go data structure out from plain C format.
func packSSJob_step_pids_response_msg_t(v [][]job_step_pids_response_msg_t, ptr0 **C.job_step_pids_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.job_step_pids_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfJob_step_pids_response_msg_tValue]C.job_step_pids_response_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newjob_step_pids_response_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSSlurm_step_layout_t transforms a sliced Go data structure into plain C format.
func unpackArgSSlurm_step_layout_t(x []slurm_step_layout_t) (unpacked *C.slurm_step_layout_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.slurm_step_layout_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSlurm_step_layout_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.slurm_step_layout_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.slurm_step_layout_t)(unsafe.Pointer(h.Data))
	return
}

// packSSlurm_step_layout_t reads sliced Go data structure out from plain C format.
func packSSlurm_step_layout_t(v []slurm_step_layout_t, ptr0 *C.slurm_step_layout_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSlurm_step_layout_tValue]C.slurm_step_layout_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newslurm_step_layout_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSJob_step_pids_t transforms a sliced Go data structure into plain C format.
func unpackArgSJob_step_pids_t(x []job_step_pids_t) (unpacked *C.job_step_pids_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_step_pids_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_step_pids_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_step_pids_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_step_pids_t)(unsafe.Pointer(h.Data))
	return
}

// unpackArgSJob_step_stat_t transforms a sliced Go data structure into plain C format.
func unpackArgSJob_step_stat_t(x []job_step_stat_t) (unpacked *C.job_step_stat_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_step_stat_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_step_stat_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_step_stat_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_step_stat_t)(unsafe.Pointer(h.Data))
	return
}

// packSJob_step_stat_t reads sliced Go data structure out from plain C format.
func packSJob_step_stat_t(v []job_step_stat_t, ptr0 *C.job_step_stat_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfJob_step_stat_tValue]C.job_step_stat_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newjob_step_stat_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPNode_info_msg_tMemory allocates memory for type *C.node_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPNode_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPNode_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPNode_info_msg_tValue = unsafe.Sizeof([1]*C.node_info_msg_t{})

// unpackArgSSNode_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSNode_info_msg_t(x [][]node_info_msg_t) (unpacked **C.node_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.node_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPNode_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.node_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocNode_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.node_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.node_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.node_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSNode_info_msg_t reads sliced Go data structure out from plain C format.
func packSSNode_info_msg_t(v [][]node_info_msg_t, ptr0 **C.node_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.node_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfNode_info_msg_tValue]C.node_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newnode_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// allocPAcct_gather_energy_tMemory allocates memory for type *C.acct_gather_energy_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPAcct_gather_energy_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPAcct_gather_energy_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPAcct_gather_energy_tValue = unsafe.Sizeof([1]*C.acct_gather_energy_t{})

// unpackArgSSAcct_gather_energy_t transforms a sliced Go data structure into plain C format.
func unpackArgSSAcct_gather_energy_t(x [][]acct_gather_energy_t) (unpacked **C.acct_gather_energy_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.acct_gather_energy_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPAcct_gather_energy_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.acct_gather_energy_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocAcct_gather_energy_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.acct_gather_energy_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.acct_gather_energy_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.acct_gather_energy_t)(unsafe.Pointer(h.Data))
	return
}

// packSSAcct_gather_energy_t reads sliced Go data structure out from plain C format.
func packSSAcct_gather_energy_t(v [][]acct_gather_energy_t, ptr0 **C.acct_gather_energy_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.acct_gather_energy_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfAcct_gather_energy_tValue]C.acct_gather_energy_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newacct_gather_energy_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSNode_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSNode_info_t(x []node_info_t) (unpacked *C.node_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.node_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocNode_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.node_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.node_info_t)(unsafe.Pointer(h.Data))
	return
}

// unpackArgSUpdate_node_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSUpdate_node_msg_t(x []update_node_msg_t) (unpacked *C.update_node_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.update_node_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocUpdate_node_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.update_node_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.update_node_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSUpdate_node_msg_t reads sliced Go data structure out from plain C format.
func packSUpdate_node_msg_t(v []update_node_msg_t, ptr0 *C.update_node_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfUpdate_node_msg_tValue]C.update_node_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newupdate_node_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPFront_end_info_msg_tMemory allocates memory for type *C.front_end_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPFront_end_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPFront_end_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPFront_end_info_msg_tValue = unsafe.Sizeof([1]*C.front_end_info_msg_t{})

// unpackArgSSFront_end_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSFront_end_info_msg_t(x [][]front_end_info_msg_t) (unpacked **C.front_end_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.front_end_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPFront_end_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.front_end_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocFront_end_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.front_end_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.front_end_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.front_end_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSFront_end_info_msg_t reads sliced Go data structure out from plain C format.
func packSSFront_end_info_msg_t(v [][]front_end_info_msg_t, ptr0 **C.front_end_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.front_end_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfFront_end_info_msg_tValue]C.front_end_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newfront_end_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSFront_end_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSFront_end_info_msg_t(x []front_end_info_msg_t) (unpacked *C.front_end_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.front_end_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocFront_end_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.front_end_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.front_end_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSFront_end_info_msg_t reads sliced Go data structure out from plain C format.
func packSFront_end_info_msg_t(v []front_end_info_msg_t, ptr0 *C.front_end_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfFront_end_info_msg_tValue]C.front_end_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newfront_end_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSFront_end_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSFront_end_info_t(x []front_end_info_t) (unpacked *C.front_end_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.front_end_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocFront_end_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.front_end_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.front_end_info_t)(unsafe.Pointer(h.Data))
	return
}

// unpackArgSUpdate_front_end_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSUpdate_front_end_msg_t(x []update_front_end_msg_t) (unpacked *C.update_front_end_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.update_front_end_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocUpdate_front_end_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.update_front_end_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.update_front_end_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSUpdate_front_end_msg_t reads sliced Go data structure out from plain C format.
func packSUpdate_front_end_msg_t(v []update_front_end_msg_t, ptr0 *C.update_front_end_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfUpdate_front_end_msg_tValue]C.update_front_end_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newupdate_front_end_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPTopo_info_response_msg_tMemory allocates memory for type *C.topo_info_response_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPTopo_info_response_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPTopo_info_response_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPTopo_info_response_msg_tValue = unsafe.Sizeof([1]*C.topo_info_response_msg_t{})

// unpackArgSSTopo_info_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSTopo_info_response_msg_t(x [][]topo_info_response_msg_t) (unpacked **C.topo_info_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.topo_info_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPTopo_info_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.topo_info_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocTopo_info_response_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.topo_info_response_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.topo_info_response_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.topo_info_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSTopo_info_response_msg_t reads sliced Go data structure out from plain C format.
func packSSTopo_info_response_msg_t(v [][]topo_info_response_msg_t, ptr0 **C.topo_info_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.topo_info_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfTopo_info_response_msg_tValue]C.topo_info_response_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newtopo_info_response_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSTopo_info_response_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSTopo_info_response_msg_t(x []topo_info_response_msg_t) (unpacked *C.topo_info_response_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.topo_info_response_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocTopo_info_response_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.topo_info_response_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.topo_info_response_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSTopo_info_response_msg_t reads sliced Go data structure out from plain C format.
func packSTopo_info_response_msg_t(v []topo_info_response_msg_t, ptr0 *C.topo_info_response_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfTopo_info_response_msg_tValue]C.topo_info_response_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newtopo_info_response_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSTopo_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSTopo_info_t(x []topo_info_t) (unpacked *C.topo_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.topo_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocTopo_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.topo_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.topo_info_t)(unsafe.Pointer(h.Data))
	return
}

// allocPPowercap_info_msg_tMemory allocates memory for type *C.powercap_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPPowercap_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPPowercap_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPPowercap_info_msg_tValue = unsafe.Sizeof([1]*C.powercap_info_msg_t{})

// unpackArgSSPowercap_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSPowercap_info_msg_t(x [][]powercap_info_msg_t) (unpacked **C.powercap_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.powercap_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPPowercap_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.powercap_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocPowercap_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.powercap_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.powercap_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.powercap_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSPowercap_info_msg_t reads sliced Go data structure out from plain C format.
func packSSPowercap_info_msg_t(v [][]powercap_info_msg_t, ptr0 **C.powercap_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.powercap_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfPowercap_info_msg_tValue]C.powercap_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newpowercap_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSPowercap_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSPowercap_info_msg_t(x []powercap_info_msg_t) (unpacked *C.powercap_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.powercap_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPowercap_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.powercap_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.powercap_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSPowercap_info_msg_t reads sliced Go data structure out from plain C format.
func packSPowercap_info_msg_t(v []powercap_info_msg_t, ptr0 *C.powercap_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPowercap_info_msg_tValue]C.powercap_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newpowercap_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSUpdate_powercap_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSUpdate_powercap_msg_t(x []update_powercap_msg_t) (unpacked *C.update_powercap_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.update_powercap_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocUpdate_powercap_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.update_powercap_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.update_powercap_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSUpdate_powercap_msg_t reads sliced Go data structure out from plain C format.
func packSUpdate_powercap_msg_t(v []update_powercap_msg_t, ptr0 *C.update_powercap_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfUpdate_powercap_msg_tValue]C.update_powercap_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newupdate_powercap_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSUpdate_part_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSUpdate_part_msg_t(x []update_part_msg_t) (unpacked *C.update_part_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.update_part_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocUpdate_part_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.update_part_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.update_part_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSUpdate_part_msg_t reads sliced Go data structure out from plain C format.
func packSUpdate_part_msg_t(v []update_part_msg_t, ptr0 *C.update_part_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfUpdate_part_msg_tValue]C.update_part_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newupdate_part_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPPartition_info_msg_tMemory allocates memory for type *C.partition_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPPartition_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPPartition_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPPartition_info_msg_tValue = unsafe.Sizeof([1]*C.partition_info_msg_t{})

// unpackArgSSPartition_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSPartition_info_msg_t(x [][]partition_info_msg_t) (unpacked **C.partition_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.partition_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPPartition_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.partition_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocPartition_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.partition_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.partition_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.partition_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSPartition_info_msg_t reads sliced Go data structure out from plain C format.
func packSSPartition_info_msg_t(v [][]partition_info_msg_t, ptr0 **C.partition_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.partition_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfPartition_info_msg_tValue]C.partition_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newpartition_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSPartition_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSPartition_info_t(x []partition_info_t) (unpacked *C.partition_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.partition_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPartition_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.partition_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.partition_info_t)(unsafe.Pointer(h.Data))
	return
}

// unpackArgSDelete_part_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSDelete_part_msg_t(x []delete_part_msg_t) (unpacked *C.delete_part_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.delete_part_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocDelete_part_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.delete_part_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.delete_part_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSDelete_part_msg_t reads sliced Go data structure out from plain C format.
func packSDelete_part_msg_t(v []delete_part_msg_t, ptr0 *C.delete_part_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfDelete_part_msg_tValue]C.delete_part_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newdelete_part_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSLayout_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSLayout_info_msg_t(x []layout_info_msg_t) (unpacked *C.layout_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.layout_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocLayout_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.layout_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.layout_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSLayout_info_msg_t reads sliced Go data structure out from plain C format.
func packSLayout_info_msg_t(v []layout_info_msg_t, ptr0 *C.layout_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfLayout_info_msg_tValue]C.layout_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newlayout_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPLayout_info_msg_tMemory allocates memory for type *C.layout_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPLayout_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPLayout_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPLayout_info_msg_tValue = unsafe.Sizeof([1]*C.layout_info_msg_t{})

// unpackArgSSLayout_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSLayout_info_msg_t(x [][]layout_info_msg_t) (unpacked **C.layout_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.layout_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPLayout_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.layout_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocLayout_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.layout_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.layout_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.layout_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSLayout_info_msg_t reads sliced Go data structure out from plain C format.
func packSSLayout_info_msg_t(v [][]layout_info_msg_t, ptr0 **C.layout_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.layout_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfLayout_info_msg_tValue]C.layout_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newlayout_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSUpdate_layout_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSUpdate_layout_msg_t(x []update_layout_msg_t) (unpacked *C.update_layout_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.update_layout_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocUpdate_layout_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.update_layout_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.update_layout_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSUpdate_layout_msg_t reads sliced Go data structure out from plain C format.
func packSUpdate_layout_msg_t(v []update_layout_msg_t, ptr0 *C.update_layout_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfUpdate_layout_msg_tValue]C.update_layout_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newupdate_layout_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSResv_desc_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSResv_desc_msg_t(x []resv_desc_msg_t) (unpacked *C.resv_desc_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.resv_desc_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocResv_desc_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.resv_desc_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.resv_desc_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSResv_desc_msg_t reads sliced Go data structure out from plain C format.
func packSResv_desc_msg_t(v []resv_desc_msg_t, ptr0 *C.resv_desc_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfResv_desc_msg_tValue]C.resv_desc_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newresv_desc_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSReservation_name_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSReservation_name_msg_t(x []reservation_name_msg_t) (unpacked *C.reservation_name_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.reservation_name_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocReservation_name_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.reservation_name_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.reservation_name_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSReservation_name_msg_t reads sliced Go data structure out from plain C format.
func packSReservation_name_msg_t(v []reservation_name_msg_t, ptr0 *C.reservation_name_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfReservation_name_msg_tValue]C.reservation_name_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newreservation_name_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPReserve_info_msg_tMemory allocates memory for type *C.reserve_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPReserve_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPReserve_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPReserve_info_msg_tValue = unsafe.Sizeof([1]*C.reserve_info_msg_t{})

// unpackArgSSReserve_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSReserve_info_msg_t(x [][]reserve_info_msg_t) (unpacked **C.reserve_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.reserve_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPReserve_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.reserve_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocReserve_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.reserve_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.reserve_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.reserve_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSReserve_info_msg_t reads sliced Go data structure out from plain C format.
func packSSReserve_info_msg_t(v [][]reserve_info_msg_t, ptr0 **C.reserve_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.reserve_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfReserve_info_msg_tValue]C.reserve_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newreserve_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSReserve_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSReserve_info_msg_t(x []reserve_info_msg_t) (unpacked *C.reserve_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.reserve_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocReserve_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.reserve_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.reserve_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSReserve_info_msg_t reads sliced Go data structure out from plain C format.
func packSReserve_info_msg_t(v []reserve_info_msg_t, ptr0 *C.reserve_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfReserve_info_msg_tValue]C.reserve_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newreserve_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSReserve_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSReserve_info_t(x []reserve_info_t) (unpacked *C.reserve_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.reserve_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocReserve_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.reserve_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.reserve_info_t)(unsafe.Pointer(h.Data))
	return
}

// unpackArgSJob_array_resp_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSJob_array_resp_msg_t(x []job_array_resp_msg_t) (unpacked *C.job_array_resp_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.job_array_resp_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocJob_array_resp_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.job_array_resp_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.job_array_resp_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSJob_array_resp_msg_t reads sliced Go data structure out from plain C format.
func packSJob_array_resp_msg_t(v []job_array_resp_msg_t, ptr0 *C.job_array_resp_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfJob_array_resp_msg_tValue]C.job_array_resp_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newjob_array_resp_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSSByte transforms a sliced Go data structure into plain C format.
func unpackArgSSByte(x [][]byte) (unpacked **C.char, allocs *cgoAllocMap) {
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

// unpackArgSTrigger_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSTrigger_info_t(x []trigger_info_t) (unpacked *C.trigger_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.trigger_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocTrigger_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.trigger_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.trigger_info_t)(unsafe.Pointer(h.Data))
	return
}

// allocPTrigger_info_msg_tMemory allocates memory for type *C.trigger_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPTrigger_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPTrigger_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPTrigger_info_msg_tValue = unsafe.Sizeof([1]*C.trigger_info_msg_t{})

// unpackArgSSTrigger_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSTrigger_info_msg_t(x [][]trigger_info_msg_t) (unpacked **C.trigger_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.trigger_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPTrigger_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.trigger_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocTrigger_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.trigger_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.trigger_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.trigger_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSTrigger_info_msg_t reads sliced Go data structure out from plain C format.
func packSSTrigger_info_msg_t(v [][]trigger_info_msg_t, ptr0 **C.trigger_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.trigger_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfTrigger_info_msg_tValue]C.trigger_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newtrigger_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSTrigger_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSTrigger_info_msg_t(x []trigger_info_msg_t) (unpacked *C.trigger_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.trigger_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocTrigger_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.trigger_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.trigger_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSTrigger_info_msg_t reads sliced Go data structure out from plain C format.
func packSTrigger_info_msg_t(v []trigger_info_msg_t, ptr0 *C.trigger_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfTrigger_info_msg_tValue]C.trigger_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newtrigger_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// allocPBurst_buffer_info_msg_tMemory allocates memory for type *C.burst_buffer_info_msg_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPBurst_buffer_info_msg_tMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPBurst_buffer_info_msg_tValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPBurst_buffer_info_msg_tValue = unsafe.Sizeof([1]*C.burst_buffer_info_msg_t{})

// unpackArgSSBurst_buffer_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSSBurst_buffer_info_msg_t(x [][]burst_buffer_info_msg_t) (unpacked **C.burst_buffer_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.burst_buffer_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPBurst_buffer_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.burst_buffer_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocBurst_buffer_info_msg_tMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.burst_buffer_info_msg_t)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.burst_buffer_info_msg_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.burst_buffer_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSSBurst_buffer_info_msg_t reads sliced Go data structure out from plain C format.
func packSSBurst_buffer_info_msg_t(v [][]burst_buffer_info_msg_t, ptr0 **C.burst_buffer_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.burst_buffer_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfBurst_buffer_info_msg_tValue]C.burst_buffer_info_msg_t)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *Newburst_buffer_info_msg_tRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSBurst_buffer_info_msg_t transforms a sliced Go data structure into plain C format.
func unpackArgSBurst_buffer_info_msg_t(x []burst_buffer_info_msg_t) (unpacked *C.burst_buffer_info_msg_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.burst_buffer_info_msg_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocBurst_buffer_info_msg_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.burst_buffer_info_msg_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.burst_buffer_info_msg_t)(unsafe.Pointer(h.Data))
	return
}

// packSBurst_buffer_info_msg_t reads sliced Go data structure out from plain C format.
func packSBurst_buffer_info_msg_t(v []burst_buffer_info_msg_t, ptr0 *C.burst_buffer_info_msg_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfBurst_buffer_info_msg_tValue]C.burst_buffer_info_msg_t)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *Newburst_buffer_info_msg_tRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSBurst_buffer_info_t transforms a sliced Go data structure into plain C format.
func unpackArgSBurst_buffer_info_t(x []burst_buffer_info_t) (unpacked *C.burst_buffer_info_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.burst_buffer_info_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocBurst_buffer_info_tMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.burst_buffer_info_t)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.burst_buffer_info_t)(unsafe.Pointer(h.Data))
	return
}
