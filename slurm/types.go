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

// JobDescMsg as declared in slurm/slurm.h:1611
type JobDescMsg struct {
	Account           []byte
	AcctgFreq         []byte
	AdminComment      []byte
	AllocNode         []byte
	AllocRespPort     uint16
	AllocSid          uint32
	Argc              uint32
	Argv              [][]byte
	ArrayInx          []byte
	ArrayBitmap       unsafe.Pointer
	BeginTime         int
	Bitflags          uint32
	BurstBuffer       []byte
	CkptInterval      uint16
	CkptDir           []byte
	Clusters          []byte
	ClusterFeatures   []byte
	Comment           []byte
	Contiguous        uint16
	CoreSpec          uint16
	CpuBind           []byte
	CpuBindType       uint16
	CpuFreqMin        uint32
	CpuFreqMax        uint32
	CpuFreqGov        uint32
	Deadline          int
	DelayBoot         uint32
	Dependency        []byte
	EndTime           int
	Environment       [][]byte
	EnvSize           uint32
	Extra             []byte
	ExcNodes          []byte
	Features          []byte
	FedSiblingsActive uint
	FedSiblingsViable uint
	Gres              []byte
	GroupId           uint32
	Immediate         uint16
	JobId             uint32
	JobIdStr          []byte
	KillOnNodeFail    uint16
	Licenses          []byte
	MailType          uint16
	MailUser          []byte
	McsLabel          []byte
	MemBind           []byte
	MemBindType       uint16
	Name              []byte
	Network           []byte
	Nice              uint32
	NumTasks          uint32
	OpenMode          byte
	OriginCluster     []byte
	OtherPort         uint16
	Overcommit        byte
	PackJobOffset     uint32
	Partition         []byte
	PlaneSize         uint16
	PowerFlags        byte
	Priority          uint32
	Profile           uint32
	Qos               []byte
	Reboot            uint16
	RespHost          []byte
	RestartCnt        uint16
	ReqNodes          []byte
	Requeue           uint16
	Reservation       []byte
	Script            []byte
	Shared            uint16
	SpankJobEnv       [][]byte
	SpankJobEnvSize   uint32
	TaskDist          uint32
	TimeLimit         uint32
	TimeMin           uint32
	UserId            uint32
	WaitAllNodes      uint16
	WarnFlags         uint16
	WarnSignal        uint16
	WarnTime          uint16
	WorkDir           []byte
	CpusPerTask       uint16
	MinCpus           uint32
	MaxCpus           uint32
	MinNodes          uint32
	MaxNodes          uint32
	BoardsPerNode     uint16
	SocketsPerBoard   uint16
	SocketsPerNode    uint16
	CoresPerSocket    uint16
	ThreadsPerCore    uint16
	NtasksPerNode     uint16
	NtasksPerSocket   uint16
	NtasksPerCore     uint16
	NtasksPerBoard    uint16
	PnMinCpus         uint16
	PnMinMemory       uint
	PnMinTmpDisk      uint32
	Geometry          [5]uint16
	ConnType          [5]uint16
	Rotate            uint16
	Blrtsimage        []byte
	Linuximage        []byte
	Mloaderimage      []byte
	Ramdiskimage      []byte
	ReqSwitch         uint32
	StdErr            []byte
	StdIn             []byte
	StdOut            []byte
	TresReqCnt        []uint
	Wait4switch       uint32
	Wckey             []byte
	X11               uint16
	X11MagicCookie    []byte
	X11TargetPort     uint16
	ref1da77736       *C.job_desc_msg_t
	allocs1da77736    interface{}
}

// SubmitResponseMsg as declared in slurm/slurm.h:3000
type SubmitResponseMsg struct {
	JobId            uint32
	StepId           uint32
	ErrorCode        uint32
	JobSubmitUserMsg []byte
	ref6c6e601       *C.submit_response_msg_t
	allocs6c6e601    interface{}
}

// TriggerInfo as declared in slurm/slurm.h:3138
type TriggerInfo struct {
	Flags          uint16
	TrigId         uint32
	ResType        uint16
	ResId          []byte
	TrigType       uint32
	Offset         uint16
	UserId         uint32
	Program        []byte
	ref1497cae7    *C.trigger_info_t
	allocs1497cae7 interface{}
}
