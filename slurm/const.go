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

const (
	// SLURM_VERSION_NUMBER as defined in slurm/slurm.h:142
	SLURM_VERSION_NUMBER = 1116933
	// SLURM_PENDING_STEP as defined in slurm/slurm.h:167
	SLURM_PENDING_STEP = 4294967293
	// SLURM_BATCH_SCRIPT as defined in slurm/slurm.h:169
	SLURM_BATCH_SCRIPT = 4294967294
	// SLURM_EXTERN_CONT as defined in slurm/slurm.h:171
	SLURM_EXTERN_CONT = 4294967295
	// SLURM_DIST_STATE_BASE as defined in slurm/slurm.h:795
	SLURM_DIST_STATE_BASE = 65535
	// SLURM_DIST_STATE_FLAGS as defined in slurm/slurm.h:796
	SLURM_DIST_STATE_FLAGS = 16711680
	// SLURM_DIST_PACK_NODES as defined in slurm/slurm.h:797
	SLURM_DIST_PACK_NODES = 8388608
	// SLURM_DIST_NO_PACK_NODES as defined in slurm/slurm.h:798
	SLURM_DIST_NO_PACK_NODES = 4194304
	// SLURM_DIST_NODEMASK as defined in slurm/slurm.h:800
	SLURM_DIST_NODEMASK = 61455
	// SLURM_DIST_SOCKMASK as defined in slurm/slurm.h:801
	SLURM_DIST_SOCKMASK = 61680
	// SLURM_DIST_COREMASK as defined in slurm/slurm.h:802
	SLURM_DIST_COREMASK = 65280
	// SLURM_DIST_NODESOCKMASK as defined in slurm/slurm.h:803
	SLURM_DIST_NODESOCKMASK = 61695
	// SLURM_SSL_SIGNATURE_LENGTH as defined in slurm/slurm.h:934
	SLURM_SSL_SIGNATURE_LENGTH = 128
	// SLURM_POWER_FLAGS_LEVEL as defined in slurm/slurm.h:1378
	SLURM_POWER_FLAGS_LEVEL = 1
	// TRIGGER_FLAG_PERM as defined in slurm/slurm.h:3096
	TRIGGER_FLAG_PERM = 1
	// TRIGGER_RES_TYPE_JOB as defined in slurm/slurm.h:3098
	TRIGGER_RES_TYPE_JOB = 1
	// TRIGGER_RES_TYPE_NODE as defined in slurm/slurm.h:3099
	TRIGGER_RES_TYPE_NODE = 2
	// TRIGGER_RES_TYPE_SLURMCTLD as defined in slurm/slurm.h:3100
	TRIGGER_RES_TYPE_SLURMCTLD = 3
	// TRIGGER_RES_TYPE_SLURMDBD as defined in slurm/slurm.h:3101
	TRIGGER_RES_TYPE_SLURMDBD = 4
	// TRIGGER_RES_TYPE_DATABASE as defined in slurm/slurm.h:3102
	TRIGGER_RES_TYPE_DATABASE = 5
	// TRIGGER_RES_TYPE_FRONT_END as defined in slurm/slurm.h:3103
	TRIGGER_RES_TYPE_FRONT_END = 6
	// TRIGGER_RES_TYPE_OTHER as defined in slurm/slurm.h:3104
	TRIGGER_RES_TYPE_OTHER = 7
	// TRIGGER_TYPE_UP as defined in slurm/slurm.h:3106
	TRIGGER_TYPE_UP = 1
	// TRIGGER_TYPE_DOWN as defined in slurm/slurm.h:3107
	TRIGGER_TYPE_DOWN = 2
	// TRIGGER_TYPE_FAIL as defined in slurm/slurm.h:3108
	TRIGGER_TYPE_FAIL = 4
	// TRIGGER_TYPE_TIME as defined in slurm/slurm.h:3109
	TRIGGER_TYPE_TIME = 8
	// TRIGGER_TYPE_FINI as defined in slurm/slurm.h:3110
	TRIGGER_TYPE_FINI = 16
	// TRIGGER_TYPE_RECONFIG as defined in slurm/slurm.h:3111
	TRIGGER_TYPE_RECONFIG = 32
	// TRIGGER_TYPE_BLOCK_ERR as defined in slurm/slurm.h:3112
	TRIGGER_TYPE_BLOCK_ERR = 64
	// TRIGGER_TYPE_IDLE as defined in slurm/slurm.h:3113
	TRIGGER_TYPE_IDLE = 128
	// TRIGGER_TYPE_DRAINED as defined in slurm/slurm.h:3114
	TRIGGER_TYPE_DRAINED = 256
	// TRIGGER_TYPE_PRI_CTLD_FAIL as defined in slurm/slurm.h:3115
	TRIGGER_TYPE_PRI_CTLD_FAIL = 512
	// TRIGGER_TYPE_PRI_CTLD_RES_OP as defined in slurm/slurm.h:3116
	TRIGGER_TYPE_PRI_CTLD_RES_OP = 1024
	// TRIGGER_TYPE_PRI_CTLD_RES_CTRL as defined in slurm/slurm.h:3117
	TRIGGER_TYPE_PRI_CTLD_RES_CTRL = 2048
	// TRIGGER_TYPE_PRI_CTLD_ACCT_FULL as defined in slurm/slurm.h:3118
	TRIGGER_TYPE_PRI_CTLD_ACCT_FULL = 4096
	// TRIGGER_TYPE_BU_CTLD_FAIL as defined in slurm/slurm.h:3119
	TRIGGER_TYPE_BU_CTLD_FAIL = 8192
	// TRIGGER_TYPE_BU_CTLD_RES_OP as defined in slurm/slurm.h:3120
	TRIGGER_TYPE_BU_CTLD_RES_OP = 16384
	// TRIGGER_TYPE_BU_CTLD_AS_CTRL as defined in slurm/slurm.h:3121
	TRIGGER_TYPE_BU_CTLD_AS_CTRL = 32768
	// TRIGGER_TYPE_PRI_DBD_FAIL as defined in slurm/slurm.h:3122
	TRIGGER_TYPE_PRI_DBD_FAIL = 65536
	// TRIGGER_TYPE_PRI_DBD_RES_OP as defined in slurm/slurm.h:3123
	TRIGGER_TYPE_PRI_DBD_RES_OP = 131072
	// TRIGGER_TYPE_PRI_DB_FAIL as defined in slurm/slurm.h:3124
	TRIGGER_TYPE_PRI_DB_FAIL = 262144
	// TRIGGER_TYPE_PRI_DB_RES_OP as defined in slurm/slurm.h:3125
	TRIGGER_TYPE_PRI_DB_RES_OP = 524288
	// TRIGGER_TYPE_BURST_BUFFER as defined in slurm/slurm.h:3126
	TRIGGER_TYPE_BURST_BUFFER = 1048576
	// SLURM_SUCCESS as defined in slurm/slurm_errno.h:56
	SLURM_SUCCESS = 0
	// SLURM_ERROR as defined in slurm/slurm_errno.h:57
	SLURM_ERROR = -1
	// SLURM_FAILURE as defined in slurm/slurm_errno.h:58
	SLURM_FAILURE = -1
	// SLURM_SOCKET_ERROR as defined in slurm/slurm_errno.h:61
	SLURM_SOCKET_ERROR = -1
	// SLURM_PROTOCOL_SUCCESS as defined in slurm/slurm_errno.h:62
	SLURM_PROTOCOL_SUCCESS = 0
	// SLURM_PROTOCOL_ERROR as defined in slurm/slurm_errno.h:63
	SLURM_PROTOCOL_ERROR = -1
)

// JobStates as declared in slurm/slurm.h:204
type JobStates int32

// JobStates enumeration from slurm/slurm.h:204
const ()

// JobStateReason as declared in slurm/slurm.h:272
type JobStateReason int32

// JobStateReason enumeration from slurm/slurm.h:272
const ()

// JobAcctTypes as declared in slurm/slurm.h:536
type JobAcctTypes int32

// JobAcctTypes enumeration from slurm/slurm.h:536
const ()

// ConnectionType as declared in slurm/slurm.h:559
type ConnectionType int32

// ConnectionType enumeration from slurm/slurm.h:559
const ()

// NodeUseType as declared in slurm/slurm.h:570
type NodeUseType int32

// NodeUseType enumeration from slurm/slurm.h:570
const ()

// SelectPluginType as declared in slurm/slurm.h:577
type SelectPluginType int32

// SelectPluginType enumeration from slurm/slurm.h:577
const ()

// SwitchPluginType as declared in slurm/slurm.h:590
type SwitchPluginType int32

// SwitchPluginType enumeration from slurm/slurm.h:590
const ()

// SelectJobdataType as declared in slurm/slurm.h:597
type SelectJobdataType int32

// SelectJobdataType enumeration from slurm/slurm.h:597
const ()

// SelectNodedataType as declared in slurm/slurm.h:626
type SelectNodedataType int32

// SelectNodedataType enumeration from slurm/slurm.h:626
const ()

// SelectPrintMode as declared in slurm/slurm.h:645
type SelectPrintMode int32

// SelectPrintMode enumeration from slurm/slurm.h:645
const ()

// SelectNodeCnt as declared in slurm/slurm.h:665
type SelectNodeCnt int32

// SelectNodeCnt enumeration from slurm/slurm.h:665
const ()

// AcctGatherProfileInfo as declared in slurm/slurm.h:677
type AcctGatherProfileInfo int32

// AcctGatherProfileInfo enumeration from slurm/slurm.h:677
const ()

// JobacctDataType as declared in slurm/slurm.h:696
type JobacctDataType int32

// JobacctDataType enumeration from slurm/slurm.h:696
const ()

// AcctEnergyType as declared in slurm/slurm.h:723
type AcctEnergyType int32

// AcctEnergyType enumeration from slurm/slurm.h:723
const ()

// NodeStates as declared in slurm/slurm.h:894
type NodeStates int32

// NodeStates enumeration from slurm/slurm.h:894
const ()

// CtxKeys as declared in slurm/slurm.h:950
type CtxKeys int32

// CtxKeys enumeration from slurm/slurm.h:950
const (
	SLURM_STEP_CTX_STEPID               = iota
	SLURM_STEP_CTX_TASKS                = 1
	SLURM_STEP_CTX_TID                  = 2
	SLURM_STEP_CTX_RESP                 = 3
	SLURM_STEP_CTX_CRED                 = 4
	SLURM_STEP_CTX_SWITCH_JOB           = 5
	SLURM_STEP_CTX_NUM_HOSTS            = 6
	SLURM_STEP_CTX_HOST                 = 7
	SLURM_STEP_CTX_JOBID                = 8
	SLURM_STEP_CTX_USER_MANAGED_SOCKETS = 9
	SLURM_STEP_CTX_NODE_LIST            = 10
	SLURM_STEP_CTX_TIDS                 = 11
)

// SuspendOpts as declared in slurm/slurm.h:1900
type SuspendOpts int32

// SuspendOpts enumeration from slurm/slurm.h:1900
const ()

const (
	// SLURM_UNEXPECTED_MSG_ERROR as declared in slurm/slurm_errno.h:67
	SLURM_UNEXPECTED_MSG_ERROR = 1000
	// SLURM_COMMUNICATIONS_CONNECTION_ERROR as declared in slurm/slurm_errno.h:68
	SLURM_COMMUNICATIONS_CONNECTION_ERROR = 1001
	// SLURM_COMMUNICATIONS_SEND_ERROR as declared in slurm/slurm_errno.h:69
	SLURM_COMMUNICATIONS_SEND_ERROR = 1002
	// SLURM_COMMUNICATIONS_RECEIVE_ERROR as declared in slurm/slurm_errno.h:70
	SLURM_COMMUNICATIONS_RECEIVE_ERROR = 1003
	// SLURM_COMMUNICATIONS_SHUTDOWN_ERROR as declared in slurm/slurm_errno.h:71
	SLURM_COMMUNICATIONS_SHUTDOWN_ERROR = 1004
	// SLURM_PROTOCOL_VERSION_ERROR as declared in slurm/slurm_errno.h:72
	SLURM_PROTOCOL_VERSION_ERROR = 1005
	// SLURM_PROTOCOL_IO_STREAM_VERSION_ERROR as declared in slurm/slurm_errno.h:73
	SLURM_PROTOCOL_IO_STREAM_VERSION_ERROR = 1006
	// SLURM_PROTOCOL_AUTHENTICATION_ERROR as declared in slurm/slurm_errno.h:74
	SLURM_PROTOCOL_AUTHENTICATION_ERROR = 1007
	// SLURM_PROTOCOL_INSANE_MSG_LENGTH as declared in slurm/slurm_errno.h:75
	SLURM_PROTOCOL_INSANE_MSG_LENGTH = 1008
	// SLURM_MPI_PLUGIN_NAME_INVALID as declared in slurm/slurm_errno.h:76
	SLURM_MPI_PLUGIN_NAME_INVALID = 1009
	// SLURM_MPI_PLUGIN_PRELAUNCH_SETUP_FAILED as declared in slurm/slurm_errno.h:77
	SLURM_MPI_PLUGIN_PRELAUNCH_SETUP_FAILED = 1010
	// SLURM_PLUGIN_NAME_INVALID as declared in slurm/slurm_errno.h:78
	SLURM_PLUGIN_NAME_INVALID = 1011
	// SLURM_UNKNOWN_FORWARD_ADDR as declared in slurm/slurm_errno.h:79
	SLURM_UNKNOWN_FORWARD_ADDR = 1012
	// SLURMCTLD_COMMUNICATIONS_CONNECTION_ERROR as declared in slurm/slurm_errno.h:82
	SLURMCTLD_COMMUNICATIONS_CONNECTION_ERROR = 1800
	// SLURMCTLD_COMMUNICATIONS_SEND_ERROR as declared in slurm/slurm_errno.h:83
	SLURMCTLD_COMMUNICATIONS_SEND_ERROR = 1801
	// SLURMCTLD_COMMUNICATIONS_RECEIVE_ERROR as declared in slurm/slurm_errno.h:84
	SLURMCTLD_COMMUNICATIONS_RECEIVE_ERROR = 1802
	// SLURMCTLD_COMMUNICATIONS_SHUTDOWN_ERROR as declared in slurm/slurm_errno.h:85
	SLURMCTLD_COMMUNICATIONS_SHUTDOWN_ERROR = 1803
	// SLURM_NO_CHANGE_IN_DATA as declared in slurm/slurm_errno.h:88
	SLURM_NO_CHANGE_IN_DATA = 1900
	// ESLURM_INVALID_PARTITION_NAME as declared in slurm/slurm_errno.h:91
	ESLURM_INVALID_PARTITION_NAME = 2000
	// ESLURM_DEFAULT_PARTITION_NOT_SET as declared in slurm/slurm_errno.h:92
	ESLURM_DEFAULT_PARTITION_NOT_SET = 2001
	// ESLURM_ACCESS_DENIED as declared in slurm/slurm_errno.h:93
	ESLURM_ACCESS_DENIED = 2002
	// ESLURM_JOB_MISSING_REQUIRED_PARTITION_GROUP as declared in slurm/slurm_errno.h:94
	ESLURM_JOB_MISSING_REQUIRED_PARTITION_GROUP = 2003
	// ESLURM_REQUESTED_NODES_NOT_IN_PARTITION as declared in slurm/slurm_errno.h:95
	ESLURM_REQUESTED_NODES_NOT_IN_PARTITION = 2004
	// ESLURM_TOO_MANY_REQUESTED_CPUS as declared in slurm/slurm_errno.h:96
	ESLURM_TOO_MANY_REQUESTED_CPUS = 2005
	// ESLURM_INVALID_NODE_COUNT as declared in slurm/slurm_errno.h:97
	ESLURM_INVALID_NODE_COUNT = 2006
	// ESLURM_ERROR_ON_DESC_TO_RECORD_COPY as declared in slurm/slurm_errno.h:98
	ESLURM_ERROR_ON_DESC_TO_RECORD_COPY = 2007
	// ESLURM_JOB_MISSING_SIZE_SPECIFICATION as declared in slurm/slurm_errno.h:99
	ESLURM_JOB_MISSING_SIZE_SPECIFICATION = 2008
	// ESLURM_JOB_SCRIPT_MISSING as declared in slurm/slurm_errno.h:100
	ESLURM_JOB_SCRIPT_MISSING = 2009
	// ESLURM_USER_ID_MISSING as declared in slurm/slurm_errno.h:101
	ESLURM_USER_ID_MISSING = 2010
	// ESLURM_DUPLICATE_JOB_ID as declared in slurm/slurm_errno.h:102
	ESLURM_DUPLICATE_JOB_ID = 2011
	// ESLURM_PATHNAME_TOO_LONG as declared in slurm/slurm_errno.h:103
	ESLURM_PATHNAME_TOO_LONG = 2012
	// ESLURM_NOT_TOP_PRIORITY as declared in slurm/slurm_errno.h:104
	ESLURM_NOT_TOP_PRIORITY = 2013
	// ESLURM_REQUESTED_NODE_CONFIG_UNAVAILABLE as declared in slurm/slurm_errno.h:105
	ESLURM_REQUESTED_NODE_CONFIG_UNAVAILABLE = 2014
	// ESLURM_REQUESTED_PART_CONFIG_UNAVAILABLE as declared in slurm/slurm_errno.h:106
	ESLURM_REQUESTED_PART_CONFIG_UNAVAILABLE = 2015
	// ESLURM_NODES_BUSY as declared in slurm/slurm_errno.h:107
	ESLURM_NODES_BUSY = 2016
	// ESLURM_INVALID_JOB_ID as declared in slurm/slurm_errno.h:108
	ESLURM_INVALID_JOB_ID = 2017
	// ESLURM_INVALID_NODE_NAME as declared in slurm/slurm_errno.h:109
	ESLURM_INVALID_NODE_NAME = 2018
	// ESLURM_WRITING_TO_FILE as declared in slurm/slurm_errno.h:110
	ESLURM_WRITING_TO_FILE = 2019
	// ESLURM_TRANSITION_STATE_NO_UPDATE as declared in slurm/slurm_errno.h:111
	ESLURM_TRANSITION_STATE_NO_UPDATE = 2020
	// ESLURM_ALREADY_DONE as declared in slurm/slurm_errno.h:112
	ESLURM_ALREADY_DONE = 2021
	// ESLURM_INTERCONNECT_FAILURE as declared in slurm/slurm_errno.h:113
	ESLURM_INTERCONNECT_FAILURE = 2022
	// ESLURM_BAD_DIST as declared in slurm/slurm_errno.h:114
	ESLURM_BAD_DIST = 2023
	// ESLURM_JOB_PENDING as declared in slurm/slurm_errno.h:115
	ESLURM_JOB_PENDING = 2024
	// ESLURM_BAD_TASK_COUNT as declared in slurm/slurm_errno.h:116
	ESLURM_BAD_TASK_COUNT = 2025
	// ESLURM_INVALID_JOB_CREDENTIAL as declared in slurm/slurm_errno.h:117
	ESLURM_INVALID_JOB_CREDENTIAL = 2026
	// ESLURM_IN_STANDBY_MODE as declared in slurm/slurm_errno.h:118
	ESLURM_IN_STANDBY_MODE = 2027
	// ESLURM_INVALID_NODE_STATE as declared in slurm/slurm_errno.h:119
	ESLURM_INVALID_NODE_STATE = 2028
	// ESLURM_INVALID_FEATURE as declared in slurm/slurm_errno.h:120
	ESLURM_INVALID_FEATURE = 2029
	// ESLURM_INVALID_AUTHTYPE_CHANGE as declared in slurm/slurm_errno.h:121
	ESLURM_INVALID_AUTHTYPE_CHANGE = 2030
	// ESLURM_INVALID_CHECKPOINT_TYPE_CHANGE as declared in slurm/slurm_errno.h:122
	ESLURM_INVALID_CHECKPOINT_TYPE_CHANGE = 2031
	// ESLURM_INVALID_SCHEDTYPE_CHANGE as declared in slurm/slurm_errno.h:123
	ESLURM_INVALID_SCHEDTYPE_CHANGE = 2032
	// ESLURM_INVALID_SELECTTYPE_CHANGE as declared in slurm/slurm_errno.h:124
	ESLURM_INVALID_SELECTTYPE_CHANGE = 2033
	// ESLURM_INVALID_SWITCHTYPE_CHANGE as declared in slurm/slurm_errno.h:125
	ESLURM_INVALID_SWITCHTYPE_CHANGE = 2034
	// ESLURM_FRAGMENTATION as declared in slurm/slurm_errno.h:126
	ESLURM_FRAGMENTATION = 2035
	// ESLURM_NOT_SUPPORTED as declared in slurm/slurm_errno.h:127
	ESLURM_NOT_SUPPORTED = 2036
	// ESLURM_DISABLED as declared in slurm/slurm_errno.h:128
	ESLURM_DISABLED = 2037
	// ESLURM_DEPENDENCY as declared in slurm/slurm_errno.h:129
	ESLURM_DEPENDENCY = 2038
	// ESLURM_BATCH_ONLY as declared in slurm/slurm_errno.h:130
	ESLURM_BATCH_ONLY = 2039
	// ESLURM_TASKDIST_ARBITRARY_UNSUPPORTED as declared in slurm/slurm_errno.h:131
	ESLURM_TASKDIST_ARBITRARY_UNSUPPORTED = 2040
	// ESLURM_TASKDIST_REQUIRES_OVERCOMMIT as declared in slurm/slurm_errno.h:132
	ESLURM_TASKDIST_REQUIRES_OVERCOMMIT = 2041
	// ESLURM_JOB_HELD as declared in slurm/slurm_errno.h:133
	ESLURM_JOB_HELD = 2042
	// ESLURM_INVALID_CRYPTO_TYPE_CHANGE as declared in slurm/slurm_errno.h:134
	ESLURM_INVALID_CRYPTO_TYPE_CHANGE = 2043
	// ESLURM_INVALID_TASK_MEMORY as declared in slurm/slurm_errno.h:135
	ESLURM_INVALID_TASK_MEMORY = 2044
	// ESLURM_INVALID_ACCOUNT as declared in slurm/slurm_errno.h:136
	ESLURM_INVALID_ACCOUNT = 2045
	// ESLURM_INVALID_PARENT_ACCOUNT as declared in slurm/slurm_errno.h:137
	ESLURM_INVALID_PARENT_ACCOUNT = 2046
	// ESLURM_SAME_PARENT_ACCOUNT as declared in slurm/slurm_errno.h:138
	ESLURM_SAME_PARENT_ACCOUNT = 2047
	// ESLURM_INVALID_LICENSES as declared in slurm/slurm_errno.h:139
	ESLURM_INVALID_LICENSES = 2048
	// ESLURM_NEED_RESTART as declared in slurm/slurm_errno.h:140
	ESLURM_NEED_RESTART = 2049
	// ESLURM_ACCOUNTING_POLICY as declared in slurm/slurm_errno.h:141
	ESLURM_ACCOUNTING_POLICY = 2050
	// ESLURM_INVALID_TIME_LIMIT as declared in slurm/slurm_errno.h:142
	ESLURM_INVALID_TIME_LIMIT = 2051
	// ESLURM_RESERVATION_ACCESS as declared in slurm/slurm_errno.h:143
	ESLURM_RESERVATION_ACCESS = 2052
	// ESLURM_RESERVATION_INVALID as declared in slurm/slurm_errno.h:144
	ESLURM_RESERVATION_INVALID = 2053
	// ESLURM_INVALID_TIME_VALUE as declared in slurm/slurm_errno.h:145
	ESLURM_INVALID_TIME_VALUE = 2054
	// ESLURM_RESERVATION_BUSY as declared in slurm/slurm_errno.h:146
	ESLURM_RESERVATION_BUSY = 2055
	// ESLURM_RESERVATION_NOT_USABLE as declared in slurm/slurm_errno.h:147
	ESLURM_RESERVATION_NOT_USABLE = 2056
	// ESLURM_INVALID_WCKEY as declared in slurm/slurm_errno.h:148
	ESLURM_INVALID_WCKEY = 2057
	// ESLURM_RESERVATION_OVERLAP as declared in slurm/slurm_errno.h:149
	ESLURM_RESERVATION_OVERLAP = 2058
	// ESLURM_PORTS_BUSY as declared in slurm/slurm_errno.h:150
	ESLURM_PORTS_BUSY = 2059
	// ESLURM_PORTS_INVALID as declared in slurm/slurm_errno.h:151
	ESLURM_PORTS_INVALID = 2060
	// ESLURM_PROLOG_RUNNING as declared in slurm/slurm_errno.h:152
	ESLURM_PROLOG_RUNNING = 2061
	// ESLURM_NO_STEPS as declared in slurm/slurm_errno.h:153
	ESLURM_NO_STEPS = 2062
	// ESLURM_INVALID_BLOCK_STATE as declared in slurm/slurm_errno.h:154
	ESLURM_INVALID_BLOCK_STATE = 2063
	// ESLURM_INVALID_BLOCK_LAYOUT as declared in slurm/slurm_errno.h:155
	ESLURM_INVALID_BLOCK_LAYOUT = 2064
	// ESLURM_INVALID_BLOCK_NAME as declared in slurm/slurm_errno.h:156
	ESLURM_INVALID_BLOCK_NAME = 2065
	// ESLURM_INVALID_QOS as declared in slurm/slurm_errno.h:157
	ESLURM_INVALID_QOS = 2066
	// ESLURM_QOS_PREEMPTION_LOOP as declared in slurm/slurm_errno.h:158
	ESLURM_QOS_PREEMPTION_LOOP = 2067
	// ESLURM_NODE_NOT_AVAIL as declared in slurm/slurm_errno.h:159
	ESLURM_NODE_NOT_AVAIL = 2068
	// ESLURM_INVALID_CPU_COUNT as declared in slurm/slurm_errno.h:160
	ESLURM_INVALID_CPU_COUNT = 2069
	// ESLURM_PARTITION_NOT_AVAIL as declared in slurm/slurm_errno.h:161
	ESLURM_PARTITION_NOT_AVAIL = 2070
	// ESLURM_CIRCULAR_DEPENDENCY as declared in slurm/slurm_errno.h:162
	ESLURM_CIRCULAR_DEPENDENCY = 2071
	// ESLURM_INVALID_GRES as declared in slurm/slurm_errno.h:163
	ESLURM_INVALID_GRES = 2072
	// ESLURM_JOB_NOT_PENDING as declared in slurm/slurm_errno.h:164
	ESLURM_JOB_NOT_PENDING = 2073
	// ESLURM_QOS_THRES as declared in slurm/slurm_errno.h:165
	ESLURM_QOS_THRES = 2074
	// ESLURM_PARTITION_IN_USE as declared in slurm/slurm_errno.h:166
	ESLURM_PARTITION_IN_USE = 2075
	// ESLURM_STEP_LIMIT as declared in slurm/slurm_errno.h:167
	ESLURM_STEP_LIMIT = 2076
	// ESLURM_JOB_SUSPENDED as declared in slurm/slurm_errno.h:168
	ESLURM_JOB_SUSPENDED = 2077
	// ESLURM_CAN_NOT_START_IMMEDIATELY as declared in slurm/slurm_errno.h:169
	ESLURM_CAN_NOT_START_IMMEDIATELY = 2078
	// ESLURM_INTERCONNECT_BUSY as declared in slurm/slurm_errno.h:170
	ESLURM_INTERCONNECT_BUSY = 2079
	// ESLURM_RESERVATION_EMPTY as declared in slurm/slurm_errno.h:171
	ESLURM_RESERVATION_EMPTY = 2080
	// ESLURM_INVALID_ARRAY as declared in slurm/slurm_errno.h:172
	ESLURM_INVALID_ARRAY = 2081
	// ESLURM_RESERVATION_NAME_DUP as declared in slurm/slurm_errno.h:173
	ESLURM_RESERVATION_NAME_DUP = 2082
	// ESLURM_JOB_STARTED as declared in slurm/slurm_errno.h:174
	ESLURM_JOB_STARTED = 2083
	// ESLURM_JOB_FINISHED as declared in slurm/slurm_errno.h:175
	ESLURM_JOB_FINISHED = 2084
	// ESLURM_JOB_NOT_RUNNING as declared in slurm/slurm_errno.h:176
	ESLURM_JOB_NOT_RUNNING = 2085
	// ESLURM_JOB_NOT_PENDING_NOR_RUNNING as declared in slurm/slurm_errno.h:177
	ESLURM_JOB_NOT_PENDING_NOR_RUNNING = 2086
	// ESLURM_JOB_NOT_SUSPENDED as declared in slurm/slurm_errno.h:178
	ESLURM_JOB_NOT_SUSPENDED = 2087
	// ESLURM_JOB_NOT_FINISHED as declared in slurm/slurm_errno.h:179
	ESLURM_JOB_NOT_FINISHED = 2088
	// ESLURM_TRIGGER_DUP as declared in slurm/slurm_errno.h:180
	ESLURM_TRIGGER_DUP = 2089
	// ESLURM_INTERNAL as declared in slurm/slurm_errno.h:181
	ESLURM_INTERNAL = 2090
	// ESLURM_INVALID_BURST_BUFFER_CHANGE as declared in slurm/slurm_errno.h:182
	ESLURM_INVALID_BURST_BUFFER_CHANGE = 2091
	// ESLURM_BURST_BUFFER_PERMISSION as declared in slurm/slurm_errno.h:183
	ESLURM_BURST_BUFFER_PERMISSION = 2092
	// ESLURM_BURST_BUFFER_LIMIT as declared in slurm/slurm_errno.h:184
	ESLURM_BURST_BUFFER_LIMIT = 2093
	// ESLURM_INVALID_BURST_BUFFER_REQUEST as declared in slurm/slurm_errno.h:185
	ESLURM_INVALID_BURST_BUFFER_REQUEST = 2094
	// ESLURM_PRIO_RESET_FAIL as declared in slurm/slurm_errno.h:186
	ESLURM_PRIO_RESET_FAIL = 2095
	// ESLURM_POWER_NOT_AVAIL as declared in slurm/slurm_errno.h:187
	ESLURM_POWER_NOT_AVAIL = 2096
	// ESLURM_POWER_RESERVED as declared in slurm/slurm_errno.h:188
	ESLURM_POWER_RESERVED = 2097
	// ESLURM_INVALID_POWERCAP as declared in slurm/slurm_errno.h:189
	ESLURM_INVALID_POWERCAP = 2098
	// ESLURM_INVALID_MCS_LABEL as declared in slurm/slurm_errno.h:190
	ESLURM_INVALID_MCS_LABEL = 2099
	// ESLURM_BURST_BUFFER_WAIT as declared in slurm/slurm_errno.h:191
	ESLURM_BURST_BUFFER_WAIT = 2100
	// ESLURM_PARTITION_DOWN as declared in slurm/slurm_errno.h:192
	ESLURM_PARTITION_DOWN = 2101
	// ESLURM_DUPLICATE_GRES as declared in slurm/slurm_errno.h:193
	ESLURM_DUPLICATE_GRES = 2102
	// ESLURM_JOB_SETTING_DB_INX as declared in slurm/slurm_errno.h:194
	ESLURM_JOB_SETTING_DB_INX = 2103
	// ESLURM_RSV_ALREADY_STARTED as declared in slurm/slurm_errno.h:195
	ESLURM_RSV_ALREADY_STARTED = 2104
	// ESLURM_SUBMISSIONS_DISABLED as declared in slurm/slurm_errno.h:196
	ESLURM_SUBMISSIONS_DISABLED = 2105
	// ESLURM_NOT_PACK_JOB as declared in slurm/slurm_errno.h:197
	ESLURM_NOT_PACK_JOB = 2106
	// ESLURM_NOT_PACK_JOB_LEADER as declared in slurm/slurm_errno.h:198
	ESLURM_NOT_PACK_JOB_LEADER = 2107
	// ESLURM_NOT_PACK_WHOLE as declared in slurm/slurm_errno.h:199
	ESLURM_NOT_PACK_WHOLE = 2108
	// ESLURM_CORE_RESERVATION_UPDATE as declared in slurm/slurm_errno.h:200
	ESLURM_CORE_RESERVATION_UPDATE = 2109
	// ESLURM_DUPLICATE_STEP_ID as declared in slurm/slurm_errno.h:201
	ESLURM_DUPLICATE_STEP_ID = 2110
	// ESLURM_INVALID_CORE_CNT as declared in slurm/slurm_errno.h:202
	ESLURM_INVALID_CORE_CNT = 2111
	// ESLURM_X11_NOT_AVAIL as declared in slurm/slurm_errno.h:203
	ESLURM_X11_NOT_AVAIL = 2112
	// ESLURM_SWITCH_MIN as declared in slurm/slurm_errno.h:206
	ESLURM_SWITCH_MIN = 3000
	// ESLURM_SWITCH_MAX as declared in slurm/slurm_errno.h:207
	ESLURM_SWITCH_MAX = 3099
	// ESLURM_JOBCOMP_MIN as declared in slurm/slurm_errno.h:208
	ESLURM_JOBCOMP_MIN = 3100
	// ESLURM_JOBCOMP_MAX as declared in slurm/slurm_errno.h:209
	ESLURM_JOBCOMP_MAX = 3199
	// ESLURM_SCHED_MIN as declared in slurm/slurm_errno.h:210
	ESLURM_SCHED_MIN = 3200
	// ESLURM_SCHED_MAX as declared in slurm/slurm_errno.h:211
	ESLURM_SCHED_MAX = 3299
	// ESLURMD_KILL_TASK_FAILED as declared in slurm/slurm_errno.h:216
	ESLURMD_KILL_TASK_FAILED = 4001
	// ESLURMD_KILL_JOB_ALREADY_COMPLETE as declared in slurm/slurm_errno.h:217
	ESLURMD_KILL_JOB_ALREADY_COMPLETE = 4002
	// ESLURMD_INVALID_ACCT_FREQ as declared in slurm/slurm_errno.h:218
	ESLURMD_INVALID_ACCT_FREQ = 4003
	// ESLURMD_INVALID_JOB_CREDENTIAL as declared in slurm/slurm_errno.h:219
	ESLURMD_INVALID_JOB_CREDENTIAL = 4004
	// ESLURMD_UID_NOT_FOUND as declared in slurm/slurm_errno.h:220
	ESLURMD_UID_NOT_FOUND = 4005
	// ESLURMD_GID_NOT_FOUND as declared in slurm/slurm_errno.h:221
	ESLURMD_GID_NOT_FOUND = 4006
	// ESLURMD_CREDENTIAL_EXPIRED as declared in slurm/slurm_errno.h:222
	ESLURMD_CREDENTIAL_EXPIRED = 4007
	// ESLURMD_CREDENTIAL_REVOKED as declared in slurm/slurm_errno.h:223
	ESLURMD_CREDENTIAL_REVOKED = 4008
	// ESLURMD_CREDENTIAL_REPLAYED as declared in slurm/slurm_errno.h:224
	ESLURMD_CREDENTIAL_REPLAYED = 4009
	// ESLURMD_CREATE_BATCH_DIR_ERROR as declared in slurm/slurm_errno.h:225
	ESLURMD_CREATE_BATCH_DIR_ERROR = 4010
	// ESLURMD_MODIFY_BATCH_DIR_ERROR as declared in slurm/slurm_errno.h:226
	ESLURMD_MODIFY_BATCH_DIR_ERROR = 4011
	// ESLURMD_CREATE_BATCH_SCRIPT_ERROR as declared in slurm/slurm_errno.h:227
	ESLURMD_CREATE_BATCH_SCRIPT_ERROR = 4012
	// ESLURMD_MODIFY_BATCH_SCRIPT_ERROR as declared in slurm/slurm_errno.h:228
	ESLURMD_MODIFY_BATCH_SCRIPT_ERROR = 4013
	// ESLURMD_SETUP_ENVIRONMENT_ERROR as declared in slurm/slurm_errno.h:229
	ESLURMD_SETUP_ENVIRONMENT_ERROR = 4014
	// ESLURMD_SHARED_MEMORY_ERROR as declared in slurm/slurm_errno.h:230
	ESLURMD_SHARED_MEMORY_ERROR = 4015
	// ESLURMD_SET_UID_OR_GID_ERROR as declared in slurm/slurm_errno.h:231
	ESLURMD_SET_UID_OR_GID_ERROR = 4016
	// ESLURMD_SET_SID_ERROR as declared in slurm/slurm_errno.h:232
	ESLURMD_SET_SID_ERROR = 4017
	// ESLURMD_CANNOT_SPAWN_IO_THREAD as declared in slurm/slurm_errno.h:233
	ESLURMD_CANNOT_SPAWN_IO_THREAD = 4018
	// ESLURMD_FORK_FAILED as declared in slurm/slurm_errno.h:234
	ESLURMD_FORK_FAILED = 4019
	// ESLURMD_EXECVE_FAILED as declared in slurm/slurm_errno.h:235
	ESLURMD_EXECVE_FAILED = 4020
	// ESLURMD_IO_ERROR as declared in slurm/slurm_errno.h:236
	ESLURMD_IO_ERROR = 4021
	// ESLURMD_PROLOG_FAILED as declared in slurm/slurm_errno.h:237
	ESLURMD_PROLOG_FAILED = 4022
	// ESLURMD_EPILOG_FAILED as declared in slurm/slurm_errno.h:238
	ESLURMD_EPILOG_FAILED = 4023
	// ESLURMD_SESSION_KILLED as declared in slurm/slurm_errno.h:239
	ESLURMD_SESSION_KILLED = 4024
	// ESLURMD_TOOMANYSTEPS as declared in slurm/slurm_errno.h:240
	ESLURMD_TOOMANYSTEPS = 4025
	// ESLURMD_STEP_EXISTS as declared in slurm/slurm_errno.h:241
	ESLURMD_STEP_EXISTS = 4026
	// ESLURMD_JOB_NOTRUNNING as declared in slurm/slurm_errno.h:242
	ESLURMD_JOB_NOTRUNNING = 4027
	// ESLURMD_STEP_SUSPENDED as declared in slurm/slurm_errno.h:243
	ESLURMD_STEP_SUSPENDED = 4028
	// ESLURMD_STEP_NOTSUSPENDED as declared in slurm/slurm_errno.h:244
	ESLURMD_STEP_NOTSUSPENDED = 4029
	// ESCRIPT_CHDIR_FAILED as declared in slurm/slurm_errno.h:247
	ESCRIPT_CHDIR_FAILED = 4100
	// ESCRIPT_OPEN_OUTPUT_FAILED as declared in slurm/slurm_errno.h:248
	ESCRIPT_OPEN_OUTPUT_FAILED = 4101
	// ESCRIPT_NON_ZERO_RETURN as declared in slurm/slurm_errno.h:249
	ESCRIPT_NON_ZERO_RETURN = 4102
	// SLURM_PROTOCOL_SOCKET_IMPL_ZERO_RECV_LENGTH as declared in slurm/slurm_errno.h:252
	SLURM_PROTOCOL_SOCKET_IMPL_ZERO_RECV_LENGTH = 5000
	// SLURM_PROTOCOL_SOCKET_IMPL_NEGATIVE_RECV_LENGTH as declared in slurm/slurm_errno.h:253
	SLURM_PROTOCOL_SOCKET_IMPL_NEGATIVE_RECV_LENGTH = 5001
	// SLURM_PROTOCOL_SOCKET_IMPL_NOT_ALL_DATA_SENT as declared in slurm/slurm_errno.h:254
	SLURM_PROTOCOL_SOCKET_IMPL_NOT_ALL_DATA_SENT = 5002
	// ESLURM_PROTOCOL_INCOMPLETE_PACKET as declared in slurm/slurm_errno.h:255
	ESLURM_PROTOCOL_INCOMPLETE_PACKET = 5003
	// SLURM_PROTOCOL_SOCKET_IMPL_TIMEOUT as declared in slurm/slurm_errno.h:256
	SLURM_PROTOCOL_SOCKET_IMPL_TIMEOUT = 5004
	// SLURM_PROTOCOL_SOCKET_ZERO_BYTES_SENT as declared in slurm/slurm_errno.h:257
	SLURM_PROTOCOL_SOCKET_ZERO_BYTES_SENT = 5005
	// ESLURM_AUTH_CRED_INVALID as declared in slurm/slurm_errno.h:260
	ESLURM_AUTH_CRED_INVALID = 6000
	// ESLURM_AUTH_FOPEN_ERROR as declared in slurm/slurm_errno.h:261
	ESLURM_AUTH_FOPEN_ERROR = 6001
	// ESLURM_AUTH_NET_ERROR as declared in slurm/slurm_errno.h:262
	ESLURM_AUTH_NET_ERROR = 6002
	// ESLURM_AUTH_UNABLE_TO_SIGN as declared in slurm/slurm_errno.h:263
	ESLURM_AUTH_UNABLE_TO_SIGN = 6003
	// ESLURM_DB_CONNECTION as declared in slurm/slurm_errno.h:266
	ESLURM_DB_CONNECTION = 7000
	// ESLURM_JOBS_RUNNING_ON_ASSOC as declared in slurm/slurm_errno.h:267
	ESLURM_JOBS_RUNNING_ON_ASSOC = 7001
	// ESLURM_CLUSTER_DELETED as declared in slurm/slurm_errno.h:268
	ESLURM_CLUSTER_DELETED = 7002
	// ESLURM_ONE_CHANGE as declared in slurm/slurm_errno.h:269
	ESLURM_ONE_CHANGE = 7003
	// ESLURM_BAD_NAME as declared in slurm/slurm_errno.h:270
	ESLURM_BAD_NAME = 7004
	// ESLURM_OVER_ALLOCATE as declared in slurm/slurm_errno.h:271
	ESLURM_OVER_ALLOCATE = 7005
	// ESLURM_RESULT_TOO_LARGE as declared in slurm/slurm_errno.h:272
	ESLURM_RESULT_TOO_LARGE = 7006
	// ESLURM_DB_QUERY_TOO_WIDE as declared in slurm/slurm_errno.h:273
	ESLURM_DB_QUERY_TOO_WIDE = 7007
	// ESLURM_FED_CLUSTER_MAX_CNT as declared in slurm/slurm_errno.h:276
	ESLURM_FED_CLUSTER_MAX_CNT = 7100
	// ESLURM_FED_CLUSTER_MULTIPLE_ASSIGNMENT as declared in slurm/slurm_errno.h:277
	ESLURM_FED_CLUSTER_MULTIPLE_ASSIGNMENT = 7101
	// ESLURM_INVALID_CLUSTER_FEATURE as declared in slurm/slurm_errno.h:278
	ESLURM_INVALID_CLUSTER_FEATURE = 7102
	// ESLURM_JOB_NOT_FEDERATED as declared in slurm/slurm_errno.h:279
	ESLURM_JOB_NOT_FEDERATED = 7103
	// ESLURM_INVALID_CLUSTER_NAME as declared in slurm/slurm_errno.h:280
	ESLURM_INVALID_CLUSTER_NAME = 7104
	// ESLURM_FED_JOB_LOCK as declared in slurm/slurm_errno.h:281
	ESLURM_FED_JOB_LOCK = 7105
	// ESLURM_FED_NO_VALID_CLUSTERS as declared in slurm/slurm_errno.h:282
	ESLURM_FED_NO_VALID_CLUSTERS = 7106
	// ESLURM_MISSING_TIME_LIMIT as declared in slurm/slurm_errno.h:285
	ESLURM_MISSING_TIME_LIMIT = 8000
	// ESLURM_INVALID_KNL as declared in slurm/slurm_errno.h:286
	ESLURM_INVALID_KNL = 8001
)

// _CodecvtResult as declared in include/libio.h:181
type _CodecvtResult int32

// _CodecvtResult enumeration from include/libio.h:181
const ()

const ()

const ()

const ()

const ()

const ()

const ()

const ()

const ()

const ()

// _SocketType as declared in bits/socket_type.h:24
type _SocketType int32

// _SocketType enumeration from bits/socket_type.h:24
const ()
