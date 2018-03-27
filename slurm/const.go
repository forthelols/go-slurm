// WARNING: This file has automatically been generated on Wed, 28 Mar 2018 00:50:28 CEST.
// By https://git.io/c-for-go. DO NOT EDIT.

package slurm

/*
#cgo pkg-config: slurm
#include <slurm/slurm.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// task_dist_states_t as declared in slurm/slurm.h:793
type task_dist_states_t int32

// task_dist_states_t enumeration from slurm/slurm.h:793
const ()

// cpu_bind_type_t as declared in slurm/slurm.h:845
type cpu_bind_type_t int32

// cpu_bind_type_t enumeration from slurm/slurm.h:845
const ()

// mem_bind_type_t as declared in slurm/slurm.h:880
type mem_bind_type_t int32

// mem_bind_type_t enumeration from slurm/slurm.h:880
const ()

// accel_bind_type_t as declared in slurm/slurm.h:887
type accel_bind_type_t int32

// accel_bind_type_t enumeration from slurm/slurm.h:887
const ()

// job_states as declared in slurm/slurm.h:204
type job_states int32

// job_states enumeration from slurm/slurm.h:204
const ()

// job_state_reason as declared in slurm/slurm.h:272
type job_state_reason int32

// job_state_reason enumeration from slurm/slurm.h:272
const ()

// job_acct_types as declared in slurm/slurm.h:536
type job_acct_types int32

// job_acct_types enumeration from slurm/slurm.h:536
const ()

// connection_type as declared in slurm/slurm.h:559
type connection_type int32

// connection_type enumeration from slurm/slurm.h:559
const ()

// node_use_type as declared in slurm/slurm.h:570
type node_use_type int32

// node_use_type enumeration from slurm/slurm.h:570
const ()

// select_plugin_type as declared in slurm/slurm.h:577
type select_plugin_type int32

// select_plugin_type enumeration from slurm/slurm.h:577
const ()

// switch_plugin_type as declared in slurm/slurm.h:590
type switch_plugin_type int32

// switch_plugin_type enumeration from slurm/slurm.h:590
const ()

// select_jobdata_type as declared in slurm/slurm.h:597
type select_jobdata_type int32

// select_jobdata_type enumeration from slurm/slurm.h:597
const ()

// select_nodedata_type as declared in slurm/slurm.h:626
type select_nodedata_type int32

// select_nodedata_type enumeration from slurm/slurm.h:626
const ()

// select_print_mode as declared in slurm/slurm.h:645
type select_print_mode int32

// select_print_mode enumeration from slurm/slurm.h:645
const ()

// select_node_cnt as declared in slurm/slurm.h:665
type select_node_cnt int32

// select_node_cnt enumeration from slurm/slurm.h:665
const ()

// acct_gather_profile_info as declared in slurm/slurm.h:677
type acct_gather_profile_info int32

// acct_gather_profile_info enumeration from slurm/slurm.h:677
const ()

// jobacct_data_type as declared in slurm/slurm.h:696
type jobacct_data_type int32

// jobacct_data_type enumeration from slurm/slurm.h:696
const ()

// acct_energy_type as declared in slurm/slurm.h:723
type acct_energy_type int32

// acct_energy_type enumeration from slurm/slurm.h:723
const ()

// node_states as declared in slurm/slurm.h:894
type node_states int32

// node_states enumeration from slurm/slurm.h:894
const ()

// ctx_keys as declared in slurm/slurm.h:950
type ctx_keys int32

// ctx_keys enumeration from slurm/slurm.h:950
const ()

// suspend_opts as declared in slurm/slurm.h:1900
type suspend_opts int32

// suspend_opts enumeration from slurm/slurm.h:1900
const ()

const ()

// __codecvt_result as declared in include/libio.h:181
type __codecvt_result int32

// __codecvt_result enumeration from include/libio.h:181
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

// __socket_type as declared in bits/socket_type.h:24
type __socket_type int32

// __socket_type enumeration from bits/socket_type.h:24
const ()
