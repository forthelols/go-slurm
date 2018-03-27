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
import "unsafe"

// slurm_addr_t as declared in slurm/slurm.h:78
type slurm_addr_t struct {
	sin_family    sa_family_t
	sin_port      in_port_t
	sin_zero      [8]byte
	ref4e93d89    *C.slurm_addr_t
	allocs4e93d89 interface{}
}

// slurmdb_cluster_rec_t as declared in slurm/slurm.h:84
type slurmdb_cluster_rec_t C.slurmdb_cluster_rec_t

// slurm_cred_t as declared in slurm/slurm.h:89
type slurm_cred_t C.slurm_cred_t

// switch_jobinfo_t as declared in slurm/slurm.h:95
type switch_jobinfo_t C.switch_jobinfo_t

// job_resources_t as declared in slurm/slurm.h:102
type job_resources_t C.job_resources_t

// select_jobinfo_t as declared in slurm/slurm.h:109
type select_jobinfo_t C.select_jobinfo_t

// select_nodeinfo_t as declared in slurm/slurm.h:110
type select_nodeinfo_t C.select_nodeinfo_t

// jobacctinfo_t as declared in slurm/slurm.h:116
type jobacctinfo_t C.jobacctinfo_t

// allocation_msg_thread_t as declared in slurm/slurm.h:123
type allocation_msg_thread_t C.allocation_msg_thread_t

// sbcast_cred_t as declared in slurm/slurm.h:128
type sbcast_cred_t C.sbcast_cred_t

// hostlist_t as declared in slurm/slurm.h:1091
type hostlist_t C.hostlist_t

// List as declared in slurm/slurm.h:1226
type List C.List

// ListIterator as declared in slurm/slurm.h:1231
type ListIterator C.ListIterator

// ListDelF type as declared in slurm/slurm.h:1236
type ListDelF func(x unsafe.Pointer)

// ListCmpF type as declared in slurm/slurm.h:1243
type ListCmpF func(x unsafe.Pointer, y unsafe.Pointer) int32

// ListFindF type as declared in slurm/slurm.h:1250
type ListFindF func(x unsafe.Pointer, key unsafe.Pointer) int32

// ListForF type as declared in slurm/slurm.h:1256
type ListForF func(x unsafe.Pointer, arg unsafe.Pointer) int32

// bitstr_t type as declared in slurm/slurm.h:1363
type bitstr_t int

// bitoff_t type as declared in slurm/slurm.h:1366
type bitoff_t int

// dynamic_plugin_data_t as declared in slurm/slurm.h:1386
type dynamic_plugin_data_t struct {
	data           unsafe.Pointer
	plugin_id      uint32_t
	ref6c951b34    *C.dynamic_plugin_data_t
	allocs6c951b34 interface{}
}

// acct_gather_energy_t as declared in slurm/slurm.h:1395
type acct_gather_energy_t struct {
	base_consumed_energy     uint64_t
	base_watts               uint32_t
	consumed_energy          uint64_t
	current_watts            uint32_t
	previous_consumed_energy uint64_t
	poll_time                time_t
	ref3c370a53              *C.acct_gather_energy_t
	allocs3c370a53           interface{}
}

// ext_sensors_data_t as declared in slurm/slurm.h:1402
type ext_sensors_data_t struct {
	consumed_energy    uint64_t
	temperature        uint32_t
	energy_update_time time_t
	current_watts      uint32_t
	refed768141        *C.ext_sensors_data_t
	allocsed768141     interface{}
}

// power_mgmt_data_t as declared in slurm/slurm.h:1416
type power_mgmt_data_t struct {
	cap_watts      uint32_t
	current_watts  uint32_t
	joule_counter  uint64_t
	new_cap_watts  uint32_t
	max_watts      uint32_t
	min_watts      uint32_t
	new_job_time   time_t
	state          uint16_t
	time_usec      uint64_t
	ref13703811    *C.power_mgmt_data_t
	allocs13703811 interface{}
}

// job_desc_msg_t as declared in slurm/slurm.h:1611
type job_desc_msg_t struct {
	account             []byte
	acctg_freq          []byte
	admin_comment       []byte
	alloc_node          []byte
	alloc_resp_port     uint16_t
	alloc_sid           uint32_t
	argc                uint32_t
	argv                [][]byte
	array_inx           []byte
	array_bitmap        unsafe.Pointer
	begin_time          time_t
	bitflags            uint32_t
	burst_buffer        []byte
	ckpt_interval       uint16_t
	ckpt_dir            []byte
	clusters            []byte
	cluster_features    []byte
	comment             []byte
	contiguous          uint16_t
	core_spec           uint16_t
	cpu_bind            []byte
	cpu_bind_type       uint16_t
	cpu_freq_min        uint32_t
	cpu_freq_max        uint32_t
	cpu_freq_gov        uint32_t
	deadline            time_t
	delay_boot          uint32_t
	dependency          []byte
	end_time            time_t
	environment         [][]byte
	env_size            uint32_t
	extra               []byte
	exc_nodes           []byte
	features            []byte
	fed_siblings_active uint64_t
	fed_siblings_viable uint64_t
	gres                []byte
	group_id            uint32_t
	immediate           uint16_t
	job_id              uint32_t
	job_id_str          []byte
	kill_on_node_fail   uint16_t
	licenses            []byte
	mail_type           uint16_t
	mail_user           []byte
	mcs_label           []byte
	mem_bind            []byte
	mem_bind_type       uint16_t
	name                []byte
	network             []byte
	nice                uint32_t
	num_tasks           uint32_t
	open_mode           uint8_t
	origin_cluster      []byte
	other_port          uint16_t
	overcommit          uint8_t
	pack_job_offset     uint32_t
	partition           []byte
	plane_size          uint16_t
	power_flags         uint8_t
	priority            uint32_t
	profile             uint32_t
	qos                 []byte
	reboot              uint16_t
	resp_host           []byte
	restart_cnt         uint16_t
	req_nodes           []byte
	requeue             uint16_t
	reservation         []byte
	script              []byte
	shared              uint16_t
	spank_job_env       [][]byte
	spank_job_env_size  uint32_t
	task_dist           uint32_t
	time_limit          uint32_t
	time_min            uint32_t
	user_id             uint32_t
	wait_all_nodes      uint16_t
	warn_flags          uint16_t
	warn_signal         uint16_t
	warn_time           uint16_t
	work_dir            []byte
	cpus_per_task       uint16_t
	min_cpus            uint32_t
	max_cpus            uint32_t
	min_nodes           uint32_t
	max_nodes           uint32_t
	boards_per_node     uint16_t
	sockets_per_board   uint16_t
	sockets_per_node    uint16_t
	cores_per_socket    uint16_t
	threads_per_core    uint16_t
	ntasks_per_node     uint16_t
	ntasks_per_socket   uint16_t
	ntasks_per_core     uint16_t
	ntasks_per_board    uint16_t
	pn_min_cpus         uint16_t
	pn_min_memory       uint64_t
	pn_min_tmp_disk     uint32_t
	geometry            [5]uint16_t
	conn_type           [5]uint16_t
	rotate              uint16_t
	blrtsimage          []byte
	linuximage          []byte
	mloaderimage        []byte
	ramdiskimage        []byte
	req_switch          uint32_t
	select_jobinfo      []dynamic_plugin_data_t
	std_err             []byte
	std_in              []byte
	std_out             []byte
	tres_req_cnt        []uint64_t
	wait4switch         uint32_t
	wckey               []byte
	x11                 uint16_t
	x11_magic_cookie    []byte
	x11_target_port     uint16_t
	ref1da77736         *C.job_desc_msg_t
	allocs1da77736      interface{}
}

// slurm_job_info_t as declared in slurm/slurm.h:1747
type slurm_job_info_t struct {
	account                 []byte
	admin_comment           []byte
	alloc_node              []byte
	alloc_sid               uint32_t
	array_bitmap            unsafe.Pointer
	array_job_id            uint32_t
	array_task_id           uint32_t
	array_max_tasks         uint32_t
	array_task_str          []byte
	assoc_id                uint32_t
	batch_flag              uint16_t
	batch_host              []byte
	bitflags                uint32_t
	boards_per_node         uint16_t
	burst_buffer            []byte
	burst_buffer_state      []byte
	cluster                 []byte
	cluster_features        []byte
	command                 []byte
	comment                 []byte
	contiguous              uint16_t
	core_spec               uint16_t
	cores_per_socket        uint16_t
	billable_tres           float64
	cpus_per_task           uint16_t
	cpu_freq_min            uint32_t
	cpu_freq_max            uint32_t
	cpu_freq_gov            uint32_t
	deadline                time_t
	delay_boot              uint32_t
	dependency              []byte
	derived_ec              uint32_t
	eligible_time           time_t
	end_time                time_t
	exc_nodes               []byte
	exc_node_inx            []int32_t
	exit_code               uint32_t
	features                []byte
	fed_origin_str          []byte
	fed_siblings_active     uint64_t
	fed_siblings_active_str []byte
	fed_siblings_viable     uint64_t
	fed_siblings_viable_str []byte
	gres                    []byte
	gres_detail_cnt         uint32_t
	gres_detail_str         [][]byte
	group_id                uint32_t
	job_id                  uint32_t
	job_resrcs              []job_resources_t
	job_state               uint32_t
	last_sched_eval         time_t
	licenses                []byte
	max_cpus                uint32_t
	max_nodes               uint32_t
	mcs_label               []byte
	name                    []byte
	network                 []byte
	nodes                   []byte
	nice                    uint32_t
	node_inx                []int32_t
	ntasks_per_core         uint16_t
	ntasks_per_node         uint16_t
	ntasks_per_socket       uint16_t
	ntasks_per_board        uint16_t
	num_cpus                uint32_t
	num_nodes               uint32_t
	num_tasks               uint32_t
	pack_job_id             uint32_t
	pack_job_id_set         []byte
	pack_job_offset         uint32_t
	partition               []byte
	pn_min_memory           uint64_t
	pn_min_cpus             uint16_t
	pn_min_tmp_disk         uint32_t
	power_flags             uint8_t
	preempt_time            time_t
	pre_sus_time            time_t
	priority                uint32_t
	profile                 uint32_t
	qos                     []byte
	reboot                  uint8_t
	req_nodes               []byte
	req_node_inx            []int32_t
	req_switch              uint32_t
	requeue                 uint16_t
	resize_time             time_t
	restart_cnt             uint16_t
	resv_name               []byte
	sched_nodes             []byte
	select_jobinfo          []dynamic_plugin_data_t
	shared                  uint16_t
	show_flags              uint16_t
	sockets_per_board       uint16_t
	sockets_per_node        uint16_t
	start_time              time_t
	start_protocol_ver      uint16_t
	state_desc              []byte
	state_reason            uint16_t
	std_err                 []byte
	std_in                  []byte
	std_out                 []byte
	submit_time             time_t
	suspend_time            time_t
	time_limit              uint32_t
	time_min                uint32_t
	threads_per_core        uint16_t
	tres_req_str            []byte
	tres_alloc_str          []byte
	user_id                 uint32_t
	user_name               []byte
	wait4switch             uint32_t
	wckey                   []byte
	work_dir                []byte
	refc83953cc             *C.slurm_job_info_t
	allocsc83953cc          interface{}
}

// priority_factors_object_t as declared in slurm/slurm.h:1767
type priority_factors_object_t struct {
	cluster_name   []byte
	job_id         uint32_t
	partition      []byte
	user_id        uint32_t
	priority_age   float64
	priority_fs    float64
	priority_js    float64
	priority_part  float64
	priority_qos   float64
	priority_tres  []float64
	tres_cnt       uint32_t
	tres_names     [][]byte
	tres_weights   []float64
	nice           uint32_t
	ref5b40eea1    *C.priority_factors_object_t
	allocs5b40eea1 interface{}
}

// priority_factors_response_msg_t as declared in slurm/slurm.h:1771
type priority_factors_response_msg_t struct {
	priority_factors_list List
	ref63319104           *C.priority_factors_response_msg_t
	allocs63319104        interface{}
}

// job_info_t as declared in slurm/slurm.h:1782
type job_info_t struct {
	account                 []byte
	admin_comment           []byte
	alloc_node              []byte
	alloc_sid               uint32_t
	array_bitmap            unsafe.Pointer
	array_job_id            uint32_t
	array_task_id           uint32_t
	array_max_tasks         uint32_t
	array_task_str          []byte
	assoc_id                uint32_t
	batch_flag              uint16_t
	batch_host              []byte
	bitflags                uint32_t
	boards_per_node         uint16_t
	burst_buffer            []byte
	burst_buffer_state      []byte
	cluster                 []byte
	cluster_features        []byte
	command                 []byte
	comment                 []byte
	contiguous              uint16_t
	core_spec               uint16_t
	cores_per_socket        uint16_t
	billable_tres           float64
	cpus_per_task           uint16_t
	cpu_freq_min            uint32_t
	cpu_freq_max            uint32_t
	cpu_freq_gov            uint32_t
	deadline                time_t
	delay_boot              uint32_t
	dependency              []byte
	derived_ec              uint32_t
	eligible_time           time_t
	end_time                time_t
	exc_nodes               []byte
	exc_node_inx            []int32_t
	exit_code               uint32_t
	features                []byte
	fed_origin_str          []byte
	fed_siblings_active     uint64_t
	fed_siblings_active_str []byte
	fed_siblings_viable     uint64_t
	fed_siblings_viable_str []byte
	gres                    []byte
	gres_detail_cnt         uint32_t
	gres_detail_str         [][]byte
	group_id                uint32_t
	job_id                  uint32_t
	job_resrcs              []job_resources_t
	job_state               uint32_t
	last_sched_eval         time_t
	licenses                []byte
	max_cpus                uint32_t
	max_nodes               uint32_t
	mcs_label               []byte
	name                    []byte
	network                 []byte
	nodes                   []byte
	nice                    uint32_t
	node_inx                []int32_t
	ntasks_per_core         uint16_t
	ntasks_per_node         uint16_t
	ntasks_per_socket       uint16_t
	ntasks_per_board        uint16_t
	num_cpus                uint32_t
	num_nodes               uint32_t
	num_tasks               uint32_t
	pack_job_id             uint32_t
	pack_job_id_set         []byte
	pack_job_offset         uint32_t
	partition               []byte
	pn_min_memory           uint64_t
	pn_min_cpus             uint16_t
	pn_min_tmp_disk         uint32_t
	power_flags             uint8_t
	preempt_time            time_t
	pre_sus_time            time_t
	priority                uint32_t
	profile                 uint32_t
	qos                     []byte
	reboot                  uint8_t
	req_nodes               []byte
	req_node_inx            []int32_t
	req_switch              uint32_t
	requeue                 uint16_t
	resize_time             time_t
	restart_cnt             uint16_t
	resv_name               []byte
	sched_nodes             []byte
	select_jobinfo          []dynamic_plugin_data_t
	shared                  uint16_t
	show_flags              uint16_t
	sockets_per_board       uint16_t
	sockets_per_node        uint16_t
	start_time              time_t
	start_protocol_ver      uint16_t
	state_desc              []byte
	state_reason            uint16_t
	std_err                 []byte
	std_in                  []byte
	std_out                 []byte
	submit_time             time_t
	suspend_time            time_t
	time_limit              uint32_t
	time_min                uint32_t
	threads_per_core        uint16_t
	tres_req_str            []byte
	tres_alloc_str          []byte
	user_id                 uint32_t
	user_name               []byte
	wait4switch             uint32_t
	wckey                   []byte
	work_dir                []byte
	refc83953cc             *C.slurm_job_info_t
	allocsc83953cc          interface{}
}

// job_info_msg_t as declared in slurm/slurm.h:1789
type job_info_msg_t struct {
	last_update    time_t
	record_count   uint32_t
	job_array      []slurm_job_info_t
	ref99f73760    *C.job_info_msg_t
	allocs99f73760 interface{}
}

// step_update_request_msg_t as declared in slurm/slurm.h:1800
type step_update_request_msg_t struct {
	end_time       time_t
	exit_code      uint32_t
	job_id         uint32_t
	jobacct        []jobacctinfo_t
	name           []byte
	start_time     time_t
	step_id        uint32_t
	time_limit     uint32_t
	refe10a901b    *C.step_update_request_msg_t
	allocse10a901b interface{}
}

// slurm_step_layout_req_t as declared in slurm/slurm.h:1812
type slurm_step_layout_req_t struct {
	node_list      []byte
	cpus_per_node  []uint16_t
	cpu_count_reps []uint32_t
	num_hosts      uint32_t
	num_tasks      uint32_t
	cpus_per_task  []uint16_t
	cpus_task_reps []uint32_t
	task_dist      uint32_t
	plane_size     uint16_t
	ref36330321    *C.slurm_step_layout_req_t
	allocs36330321 interface{}
}

// slurm_step_layout_t as declared in slurm/slurm.h:1834
type slurm_step_layout_t struct {
	front_end          []byte
	node_cnt           uint32_t
	node_list          []byte
	plane_size         uint16_t
	start_protocol_ver uint16_t
	tasks              []uint16_t
	task_cnt           uint32_t
	task_dist          uint32_t
	tids               [][]uint32_t
	ref5f1a1283        *C.slurm_step_layout_t
	allocs5f1a1283     interface{}
}

// slurm_step_io_fds_t as declared in slurm/slurm.h:1842
type slurm_step_io_fds_t struct {
	ref76cac59f    *C.slurm_step_io_fds_t
	allocs76cac59f interface{}
}

// launch_tasks_response_msg_t as declared in slurm/slurm.h:1857
type launch_tasks_response_msg_t struct {
	job_id        uint32_t
	step_id       uint32_t
	return_code   uint32_t
	node_name     []byte
	srun_node_id  uint32_t
	count_of_pids uint32_t
	local_pids    []uint32_t
	task_ids      []uint32_t
	ref9da0a9b    *C.launch_tasks_response_msg_t
	allocs9da0a9b interface{}
}

// task_exit_msg_t as declared in slurm/slurm.h:1865
type task_exit_msg_t struct {
	num_tasks      uint32_t
	task_id_list   []uint32_t
	return_code    uint32_t
	job_id         uint32_t
	step_id        uint32_t
	ref968ced3f    *C.task_exit_msg_t
	allocs968ced3f interface{}
}

// srun_ping_msg_t as declared in slurm/slurm.h:1870
type srun_ping_msg_t struct {
	job_id        uint32_t
	step_id       uint32_t
	ref9d7c08b    *C.srun_ping_msg_t
	allocs9d7c08b interface{}
}

// srun_job_complete_msg_t as declared in slurm/slurm.h:1875
type srun_job_complete_msg_t struct {
	job_id         uint32_t
	step_id        uint32_t
	ref12083467    *C.srun_job_complete_msg_t
	allocs12083467 interface{}
}

// srun_timeout_msg_t as declared in slurm/slurm.h:1881
type srun_timeout_msg_t struct {
	job_id         uint32_t
	step_id        uint32_t
	timeout        time_t
	ref51c6caaf    *C.srun_timeout_msg_t
	allocs51c6caaf interface{}
}

// srun_user_msg_t as declared in slurm/slurm.h:1886
type srun_user_msg_t struct {
	job_id         uint32_t
	msg            []byte
	refbe9c6c1f    *C.srun_user_msg_t
	allocsbe9c6c1f interface{}
}

// srun_node_fail_msg_t as declared in slurm/slurm.h:1892
type srun_node_fail_msg_t struct {
	job_id        uint32_t
	nodelist      []byte
	step_id       uint32_t
	refcad68cc    *C.srun_node_fail_msg_t
	allocscad68cc interface{}
}

// srun_step_missing_msg_t as declared in slurm/slurm.h:1898
type srun_step_missing_msg_t struct {
	job_id         uint32_t
	nodelist       []byte
	step_id        uint32_t
	ref7629d094    *C.srun_step_missing_msg_t
	allocs7629d094 interface{}
}

// suspend_msg_t as declared in slurm/slurm.h:1910
type suspend_msg_t struct {
	op             uint16_t
	job_id         uint32_t
	job_id_str     []byte
	ref4b652065    *C.suspend_msg_t
	allocs4b652065 interface{}
}

// top_job_msg_t as declared in slurm/slurm.h:1917
type top_job_msg_t struct {
	op             uint16_t
	job_id         uint32_t
	job_id_str     []byte
	ref874ee789    *C.top_job_msg_t
	allocs874ee789 interface{}
}

// slurm_step_ctx_params_t as declared in slurm/slurm.h:1960
type slurm_step_ctx_params_t struct {
	ckpt_interval  uint16_t
	cpu_count      uint32_t
	cpu_freq_min   uint32_t
	cpu_freq_max   uint32_t
	cpu_freq_gov   uint32_t
	exclusive      uint16_t
	features       []byte
	immediate      uint16_t
	job_id         uint32_t
	pn_min_memory  uint64_t
	ckpt_dir       []byte
	gres           []byte
	name           []byte
	network        []byte
	profile        uint32_t
	no_kill        uint8_t
	min_nodes      uint32_t
	max_nodes      uint32_t
	node_list      []byte
	overcommit     bool
	plane_size     uint16_t
	relative       uint16_t
	resv_port_cnt  uint16_t
	step_id        uint32_t
	task_count     uint32_t
	task_dist      uint32_t
	time_limit     uint32_t
	uid            uid_t
	verbose_level  uint16_t
	ref7db1e503    *C.slurm_step_ctx_params_t
	allocs7db1e503 interface{}
}

// slurm_step_launch_params_t as declared in slurm/slurm.h:2028
type slurm_step_launch_params_t struct {
	alias_list             []byte
	argc                   uint32_t
	argv                   [][]byte
	envc                   uint32_t
	env                    [][]byte
	cwd                    []byte
	user_managed_io        bool
	msg_timeout            uint32_t
	ntasks_per_board       uint16_t
	ntasks_per_core        uint16_t
	ntasks_per_socket      uint16_t
	buffered_stdio         bool
	labelio                bool
	remote_output_filename []byte
	remote_error_filename  []byte
	remote_input_filename  []byte
	local_fds              slurm_step_io_fds_t
	gid                    uint32_t
	multi_prog             bool
	no_alloc               bool
	slurmd_debug           uint32_t
	node_offset            uint32_t
	pack_jobid             uint32_t
	pack_nnodes            uint32_t
	pack_ntasks            uint32_t
	pack_task_cnts         []uint16_t
	pack_tids              [][]uint32_t
	pack_offset            uint32_t
	pack_task_offset       uint32_t
	pack_node_list         []byte
	parallel_debug         bool
	profile                uint32_t
	task_prolog            []byte
	task_epilog            []byte
	cpu_bind_type          uint16_t
	cpu_bind               []byte
	cpu_freq_min           uint32_t
	cpu_freq_max           uint32_t
	cpu_freq_gov           uint32_t
	mem_bind_type          uint16_t
	mem_bind               []byte
	accel_bind_type        uint16_t
	max_sockets            uint16_t
	max_cores              uint16_t
	max_threads            uint16_t
	cpus_per_task          uint16_t
	task_dist              uint32_t
	partition              []byte
	preserve_env           bool
	mpi_plugin_name        []byte
	open_mode              uint8_t
	acctg_freq             []byte
	pty                    bool
	ckpt_dir               []byte
	restart_dir            []byte
	spank_job_env          [][]byte
	spank_job_env_size     uint32_t
	ref17ff30d9            *C.slurm_step_launch_params_t
	allocs17ff30d9         interface{}
}

// slurm_step_launch_callbacks_t as declared in slurm/slurm.h:2036
type slurm_step_launch_callbacks_t struct {
	step_complete  *func(arg0 []srun_job_complete_msg_t)
	step_signal    *func(arg0 int32)
	step_timeout   *func(arg0 []srun_timeout_msg_t)
	task_start     *func(arg0 []launch_tasks_response_msg_t)
	task_finish    *func(arg0 []task_exit_msg_t)
	refaf4fc3bc    *C.slurm_step_launch_callbacks_t
	allocsaf4fc3bc interface{}
}

// slurm_allocation_callbacks_t as declared in slurm/slurm.h:2045
type slurm_allocation_callbacks_t struct {
	ping           *func(arg0 []srun_ping_msg_t)
	job_complete   *func(arg0 []srun_job_complete_msg_t)
	timeout        *func(arg0 []srun_timeout_msg_t)
	user_msg       *func(arg0 []srun_user_msg_t)
	node_fail      *func(arg0 []srun_node_fail_msg_t)
	job_suspend    *func(arg0 []suspend_msg_t)
	ref681611eb    *C.slurm_allocation_callbacks_t
	allocs681611eb interface{}
}

// slurm_trigger_callbacks_t as declared in slurm/slurm.h:2053
type slurm_trigger_callbacks_t struct {
	acct_full     *func()
	dbd_fail      *func()
	dbd_resumed   *func()
	db_fail       *func()
	db_resumed    *func()
	ref3554fae    *C.slurm_trigger_callbacks_t
	allocs3554fae interface{}
}

// job_step_info_t as declared in slurm/slurm.h:2093
type job_step_info_t struct {
	array_job_id       uint32_t
	array_task_id      uint32_t
	ckpt_dir           []byte
	ckpt_interval      uint16_t
	cluster            []byte
	gres               []byte
	job_id             uint32_t
	name               []byte
	network            []byte
	nodes              []byte
	node_inx           []int32_t
	num_cpus           uint32_t
	cpu_freq_min       uint32_t
	cpu_freq_max       uint32_t
	cpu_freq_gov       uint32_t
	num_tasks          uint32_t
	partition          []byte
	resv_ports         []byte
	run_time           time_t
	select_jobinfo     []dynamic_plugin_data_t
	srun_host          []byte
	srun_pid           uint32_t
	start_time         time_t
	start_protocol_ver uint16_t
	state              uint32_t
	step_id            uint32_t
	task_dist          uint32_t
	time_limit         uint32_t
	tres_alloc_str     []byte
	user_id            uint32_t
	refc1c43052        *C.job_step_info_t
	allocsc1c43052     interface{}
}

// job_step_info_response_msg_t as declared in slurm/slurm.h:2099
type job_step_info_response_msg_t struct {
	last_update    time_t
	job_step_count uint32_t
	job_steps      []job_step_info_t
	ref3f513490    *C.job_step_info_response_msg_t
	allocs3f513490 interface{}
}

// job_step_pids_t as declared in slurm/slurm.h:2105
type job_step_pids_t struct {
	node_name      []byte
	pid            []uint32_t
	pid_cnt        uint32_t
	refe8082d8e    *C.job_step_pids_t
	allocse8082d8e interface{}
}

// job_step_pids_response_msg_t as declared in slurm/slurm.h:2111
type job_step_pids_response_msg_t struct {
	job_id         uint32_t
	pid_list       List
	step_id        uint32_t
	ref8647e5b2    *C.job_step_pids_response_msg_t
	allocs8647e5b2 interface{}
}

// job_step_stat_t as declared in slurm/slurm.h:2118
type job_step_stat_t struct {
	jobacct        []jobacctinfo_t
	num_tasks      uint32_t
	return_code    uint32_t
	step_pids      []job_step_pids_t
	refc47deaa4    *C.job_step_stat_t
	allocsc47deaa4 interface{}
}

// job_step_stat_response_msg_t as declared in slurm/slurm.h:2124
type job_step_stat_response_msg_t struct {
	job_id         uint32_t
	stats_list     List
	step_id        uint32_t
	reff14ee411    *C.job_step_stat_response_msg_t
	allocsf14ee411 interface{}
}

// node_info_t as declared in slurm/slurm.h:2176
type node_info_t struct {
	arch              []byte
	boards            uint16_t
	boot_time         time_t
	cluster_name      []byte
	cores             uint16_t
	core_spec_cnt     uint16_t
	cpu_load          uint32_t
	free_mem          uint64_t
	cpus              uint16_t
	cpu_spec_list     []byte
	energy            []acct_gather_energy_t
	ext_sensors       []ext_sensors_data_t
	power             []power_mgmt_data_t
	features          []byte
	features_act      []byte
	gres              []byte
	gres_drain        []byte
	gres_used         []byte
	mcs_label         []byte
	mem_spec_limit    uint64_t
	name              []byte
	node_addr         []byte
	node_hostname     []byte
	node_state        uint32_t
	os                []byte
	owner             uint32_t
	partitions        []byte
	port              uint16_t
	real_memory       uint64_t
	reason            []byte
	reason_time       time_t
	reason_uid        uint32_t
	select_nodeinfo   []dynamic_plugin_data_t
	slurmd_start_time time_t
	sockets           uint16_t
	threads           uint16_t
	tmp_disk          uint32_t
	weight            uint32_t
	tres_fmt_str      []byte
	version           []byte
	ref3ba4f3c4       *C.node_info_t
	allocs3ba4f3c4    interface{}
}

// node_info_msg_t as declared in slurm/slurm.h:2188
type node_info_msg_t struct {
	last_update    time_t
	node_scaling   uint32_t
	record_count   uint32_t
	node_array     []node_info_t
	ref27202897    *C.node_info_msg_t
	allocs27202897 interface{}
}

// front_end_info_t as declared in slurm/slurm.h:2207
type front_end_info_t struct {
	allow_groups      []byte
	allow_users       []byte
	boot_time         time_t
	deny_groups       []byte
	deny_users        []byte
	name              []byte
	node_state        uint32_t
	reason            []byte
	reason_time       time_t
	reason_uid        uint32_t
	slurmd_start_time time_t
	version           []byte
	refd6c05f8c       *C.front_end_info_t
	allocsd6c05f8c    interface{}
}

// front_end_info_msg_t as declared in slurm/slurm.h:2213
type front_end_info_msg_t struct {
	last_update     time_t
	record_count    uint32_t
	front_end_array []front_end_info_t
	ref76c8cc27     *C.front_end_info_msg_t
	allocs76c8cc27  interface{}
}

// topo_info_t as declared in slurm/slurm.h:2221
type topo_info_t struct {
	level          uint16_t
	link_speed     uint32_t
	name           []byte
	nodes          []byte
	switches       []byte
	ref1711603e    *C.topo_info_t
	allocs1711603e interface{}
}

// topo_info_response_msg_t as declared in slurm/slurm.h:2226
type topo_info_response_msg_t struct {
	record_count   uint32_t
	topo_array     []topo_info_t
	refbb64b6ce    *C.topo_info_response_msg_t
	allocsbb64b6ce interface{}
}

// job_alloc_info_msg_t as declared in slurm/slurm.h:2231
type job_alloc_info_msg_t struct {
	job_id         uint32_t
	req_cluster    []byte
	ref8a9412bc    *C.job_alloc_info_msg_t
	allocs8a9412bc interface{}
}

// layout_info_msg_t as declared in slurm/slurm.h:2236
type layout_info_msg_t struct {
	record_count   uint32_t
	records        [][]byte
	ref6987b71d    *C.layout_info_msg_t
	allocs6987b71d interface{}
}

// update_layout_msg_t as declared in slurm/slurm.h:2241
type update_layout_msg_t struct {
	layout         []byte
	arg            []byte
	ref63746d9e    *C.update_layout_msg_t
	allocs63746d9e interface{}
}

// step_alloc_info_msg_t as declared in slurm/slurm.h:2247
type step_alloc_info_msg_t struct {
	job_id          uint32_t
	pack_job_offset uint32_t
	step_id         uint32_t
	ref38ca7831     *C.step_alloc_info_msg_t
	allocs38ca7831  interface{}
}

// powercap_info_msg_t as declared in slurm/slurm.h:2259
type powercap_info_msg_t struct {
	power_cap      uint32_t
	power_floor    uint32_t
	power_change   uint32_t
	min_watts      uint32_t
	cur_max_watts  uint32_t
	adj_max_watts  uint32_t
	max_watts      uint32_t
	refe64943c1    *C.powercap_info_msg_t
	allocse64943c1 interface{}
}

// update_powercap_msg_t as declared in slurm/slurm.h:2261
type update_powercap_msg_t struct {
	power_cap      uint32_t
	power_floor    uint32_t
	power_change   uint32_t
	min_watts      uint32_t
	cur_max_watts  uint32_t
	adj_max_watts  uint32_t
	max_watts      uint32_t
	ref38054dec    *C.update_powercap_msg_t
	allocs38054dec interface{}
}

// acct_gather_node_resp_msg_t as declared in slurm/slurm.h:2267
type acct_gather_node_resp_msg_t struct {
	energy         []acct_gather_energy_t
	node_name      []byte
	sensor_cnt     uint16_t
	ref94035ba4    *C.acct_gather_node_resp_msg_t
	allocs94035ba4 interface{}
}

// acct_gather_energy_req_msg_t as declared in slurm/slurm.h:2271
type acct_gather_energy_req_msg_t struct {
	delta         uint16_t
	ref47835ed    *C.acct_gather_energy_req_msg_t
	allocs47835ed interface{}
}

// partition_info_t as declared in slurm/slurm.h:2336
type partition_info_t struct {
	allow_alloc_nodes   []byte
	allow_accounts      []byte
	allow_groups        []byte
	allow_qos           []byte
	alternate           []byte
	billing_weights_str []byte
	cluster_name        []byte
	cr_type             uint16_t
	def_mem_per_cpu     uint64_t
	default_time        uint32_t
	deny_accounts       []byte
	deny_qos            []byte
	flags               uint16_t
	grace_time          uint32_t
	max_cpus_per_node   uint32_t
	max_mem_per_cpu     uint64_t
	max_nodes           uint32_t
	max_share           uint16_t
	max_time            uint32_t
	min_nodes           uint32_t
	name                []byte
	node_inx            []int32_t
	nodes               []byte
	over_time_limit     uint16_t
	preempt_mode        uint16_t
	priority_job_factor uint16_t
	priority_tier       uint16_t
	qos_char            []byte
	state_up            uint16_t
	total_cpus          uint32_t
	total_nodes         uint32_t
	tres_fmt_str        []byte
	refc2f47d8f         *C.partition_info_t
	allocsc2f47d8f      interface{}
}

// delete_part_msg_t as declared in slurm/slurm.h:2340
type delete_part_msg_t struct {
	name           []byte
	ref2d75a55b    *C.delete_part_msg_t
	allocs2d75a55b interface{}
}

// resource_allocation_response_msg_t as declared in slurm/slurm.h:2377
type resource_allocation_response_msg_t struct {
	account             []byte
	job_id              uint32_t
	alias_list          []byte
	cpu_freq_min        uint32_t
	cpu_freq_max        uint32_t
	cpu_freq_gov        uint32_t
	cpus_per_node       []uint16_t
	cpu_count_reps      []uint32_t
	env_size            uint32_t
	environment         [][]byte
	error_code          uint32_t
	job_submit_user_msg []byte
	node_addr           []slurm_addr_t
	node_cnt            uint32_t
	node_list           []byte
	ntasks_per_board    uint16_t
	ntasks_per_core     uint16_t
	ntasks_per_socket   uint16_t
	num_cpu_groups      uint32_t
	partition           []byte
	pn_min_memory       uint64_t
	qos                 []byte
	resv_name           []byte
	select_jobinfo      []dynamic_plugin_data_t
	working_cluster_rec unsafe.Pointer
	refa10376ab         *C.resource_allocation_response_msg_t
	allocsa10376ab      interface{}
}

// partition_info_msg_t as declared in slurm/slurm.h:2383
type partition_info_msg_t struct {
	last_update     time_t
	record_count    uint32_t
	partition_array []partition_info_t
	reffe44d6cd     *C.partition_info_msg_t
	allocsfe44d6cd  interface{}
}

// will_run_response_msg_t as declared in slurm/slurm.h:2393
type will_run_response_msg_t struct {
	job_id              uint32_t
	job_submit_user_msg []byte
	node_list           []byte
	preemptee_job_id    List
	proc_cnt            uint32_t
	start_time          time_t
	sys_usage_per       float64
	reffc5b749c         *C.will_run_response_msg_t
	allocsfc5b749c      interface{}
}

// block_job_info_t as declared in slurm/slurm.h:2409
type block_job_info_t struct {
	cnodes         []byte
	cnode_inx      []int32_t
	job_id         uint32_t
	job_ptr        unsafe.Pointer
	user_id        uint32_t
	user_name      []byte
	refb25fc08b    *C.block_job_info_t
	allocsb25fc08b interface{}
}

// block_info_t as declared in slurm/slurm.h:2433
type block_info_t struct {
	bg_block_id    []byte
	blrtsimage     []byte
	conn_type      [5]uint16_t
	cnode_cnt      uint32_t
	cnode_err_cnt  uint32_t
	ionode_inx     []int32_t
	ionode_str     []byte
	job_list       List
	linuximage     []byte
	mloaderimage   []byte
	mp_inx         []int32_t
	mp_str         []byte
	node_use       uint16_t
	ramdiskimage   []byte
	reason         []byte
	state          uint16_t
	refc618b0ea    *C.block_info_t
	allocsc618b0ea interface{}
}

// block_info_msg_t as declared in slurm/slurm.h:2439
type block_info_msg_t struct {
	block_array    []block_info_t
	last_update    time_t
	record_count   uint32_t
	ref9e6dac16    *C.block_info_msg_t
	allocs9e6dac16 interface{}
}

// update_block_msg_t as declared in slurm/slurm.h:2441
type update_block_msg_t struct {
	bg_block_id    []byte
	blrtsimage     []byte
	conn_type      [5]uint16_t
	cnode_cnt      uint32_t
	cnode_err_cnt  uint32_t
	ionode_inx     []int32_t
	ionode_str     []byte
	job_list       List
	linuximage     []byte
	mloaderimage   []byte
	mp_inx         []int32_t
	mp_str         []byte
	node_use       uint16_t
	ramdiskimage   []byte
	reason         []byte
	state          uint16_t
	refc618b0ea    *C.block_info_t
	allocsc618b0ea interface{}
}

// resv_core_spec_t as declared in slurm/slurm.h:2565
type resv_core_spec_t struct {
	node_name      []byte
	core_id        []byte
	ref256bbcbd    *C.resv_core_spec_t
	allocs256bbcbd interface{}
}

// reserve_info_t as declared in slurm/slurm.h:2588
type reserve_info_t struct {
	accounts       []byte
	burst_buffer   []byte
	core_cnt       uint32_t
	core_spec_cnt  uint32_t
	core_spec      []resv_core_spec_t
	end_time       time_t
	features       []byte
	flags          uint32_t
	licenses       []byte
	name           []byte
	node_cnt       uint32_t
	node_inx       []int32_t
	node_list      []byte
	partition      []byte
	start_time     time_t
	resv_watts     uint32_t
	tres_str       []byte
	users          []byte
	refb9d691a5    *C.reserve_info_t
	allocsb9d691a5 interface{}
}

// reserve_info_msg_t as declared in slurm/slurm.h:2594
type reserve_info_msg_t struct {
	last_update       time_t
	record_count      uint32_t
	reservation_array []reserve_info_t
	reff1565ec7       *C.reserve_info_msg_t
	allocsf1565ec7    interface{}
}

// resv_desc_msg_t as declared in slurm/slurm.h:2617
type resv_desc_msg_t struct {
	accounts       []byte
	burst_buffer   []byte
	core_cnt       []uint32_t
	duration       uint32_t
	end_time       time_t
	features       []byte
	flags          uint32_t
	licenses       []byte
	name           []byte
	node_cnt       []uint32_t
	node_list      []byte
	partition      []byte
	start_time     time_t
	resv_watts     uint32_t
	tres_str       []byte
	users          []byte
	refc2221b7f    *C.resv_desc_msg_t
	allocsc2221b7f interface{}
}

// reserve_response_msg_t as declared in slurm/slurm.h:2621
type reserve_response_msg_t struct {
	name           []byte
	ref18195ccc    *C.reserve_response_msg_t
	allocs18195ccc interface{}
}

// reservation_name_msg_t as declared in slurm/slurm.h:2626
type reservation_name_msg_t struct {
	name           []byte
	ref5755f979    *C.reservation_name_msg_t
	allocs5755f979 interface{}
}

// slurm_ctl_conf_t as declared in slurm/slurm.h:2975
type slurm_ctl_conf_t struct {
	last_update                       time_t
	accounting_storage_tres           []byte
	accounting_storage_enforce        uint16_t
	accounting_storage_backup_host    []byte
	accounting_storage_host           []byte
	accounting_storage_loc            []byte
	accounting_storage_pass           []byte
	accounting_storage_port           uint32_t
	accounting_storage_type           []byte
	accounting_storage_user           []byte
	acctng_store_job_comment          uint16_t
	acct_gather_conf                  unsafe.Pointer
	acct_gather_energy_type           []byte
	acct_gather_profile_type          []byte
	acct_gather_interconnect_type     []byte
	acct_gather_filesystem_type       []byte
	acct_gather_node_freq             uint16_t
	authinfo                          []byte
	authtype                          []byte
	backup_addr                       []byte
	backup_controller                 []byte
	batch_start_timeout               uint16_t
	bb_type                           []byte
	boot_time                         time_t
	checkpoint_type                   []byte
	chos_loc                          []byte
	core_spec_plugin                  []byte
	cluster_name                      []byte
	complete_wait                     uint16_t
	control_addr                      []byte
	control_machine                   []byte
	cpu_freq_def                      uint32_t
	cpu_freq_govs                     uint32_t
	crypto_type                       []byte
	debug_flags                       uint64_t
	def_mem_per_cpu                   uint64_t
	disable_root_jobs                 uint16_t
	eio_timeout                       uint16_t
	enforce_part_limits               uint16_t
	epilog                            []byte
	epilog_msg_time                   uint32_t
	epilog_slurmctld                  []byte
	ext_sensors_type                  []byte
	ext_sensors_freq                  uint16_t
	ext_sensors_conf                  unsafe.Pointer
	fast_schedule                     uint16_t
	fed_params                        []byte
	first_job_id                      uint32_t
	fs_dampening_factor               uint16_t
	get_env_timeout                   uint16_t
	gres_plugins                      []byte
	group_time                        uint16_t
	group_force                       uint16_t
	hash_val                          uint32_t
	health_check_interval             uint16_t
	health_check_node_state           uint16_t
	health_check_program              []byte
	inactive_limit                    uint16_t
	job_acct_gather_freq              []byte
	job_acct_gather_type              []byte
	job_acct_gather_params            []byte
	job_ckpt_dir                      []byte
	job_comp_host                     []byte
	job_comp_loc                      []byte
	job_comp_pass                     []byte
	job_comp_port                     uint32_t
	job_comp_type                     []byte
	job_comp_user                     []byte
	job_container_plugin              []byte
	job_credential_private_key        []byte
	job_credential_public_certificate []byte
	job_file_append                   uint16_t
	job_requeue                       uint16_t
	job_submit_plugins                []byte
	keep_alive_time                   uint16_t
	kill_on_bad_exit                  uint16_t
	kill_wait                         uint16_t
	launch_params                     []byte
	launch_type                       []byte
	layouts                           []byte
	licenses                          []byte
	licenses_used                     []byte
	log_fmt                           uint16_t
	mail_domain                       []byte
	mail_prog                         []byte
	max_array_sz                      uint32_t
	max_job_cnt                       uint32_t
	max_job_id                        uint32_t
	max_mem_per_cpu                   uint64_t
	max_step_cnt                      uint32_t
	max_tasks_per_node                uint16_t
	mcs_plugin                        []byte
	mcs_plugin_params                 []byte
	mem_limit_enforce                 uint16_t
	min_job_age                       uint32_t
	mpi_default                       []byte
	mpi_params                        []byte
	msg_aggr_params                   []byte
	msg_timeout                       uint16_t
	tcp_timeout                       uint16_t
	next_job_id                       uint32_t
	node_features_plugins             []byte
	node_prefix                       []byte
	over_time_limit                   uint16_t
	plugindir                         []byte
	plugstack                         []byte
	power_parameters                  []byte
	power_plugin                      []byte
	preempt_mode                      uint16_t
	preempt_type                      []byte
	priority_decay_hl                 uint32_t
	priority_calc_period              uint32_t
	priority_favor_small              uint16_t
	priority_flags                    uint16_t
	priority_max_age                  uint32_t
	priority_params                   []byte
	priority_reset_period             uint16_t
	priority_type                     []byte
	priority_weight_age               uint32_t
	priority_weight_fs                uint32_t
	priority_weight_js                uint32_t
	priority_weight_part              uint32_t
	priority_weight_qos               uint32_t
	priority_weight_tres              []byte
	private_data                      uint16_t
	proctrack_type                    []byte
	prolog                            []byte
	prolog_epilog_timeout             uint16_t
	prolog_slurmctld                  []byte
	propagate_prio_process            uint16_t
	prolog_flags                      uint16_t
	propagate_rlimits                 []byte
	propagate_rlimits_except          []byte
	reboot_program                    []byte
	reconfig_flags                    uint16_t
	requeue_exit                      []byte
	requeue_exit_hold                 []byte
	resume_program                    []byte
	resume_rate                       uint16_t
	resume_timeout                    uint16_t
	resv_epilog                       []byte
	resv_over_run                     uint16_t
	resv_prolog                       []byte
	ret2service                       uint16_t
	route_plugin                      []byte
	salloc_default_command            []byte
	sbcast_parameters                 []byte
	sched_logfile                     []byte
	sched_log_level                   uint16_t
	sched_params                      []byte
	sched_time_slice                  uint16_t
	schedtype                         []byte
	select_type                       []byte
	select_conf_key_pairs             unsafe.Pointer
	select_type_param                 uint16_t
	slurm_conf                        []byte
	slurm_user_id                     uint32_t
	slurm_user_name                   []byte
	slurmd_user_id                    uint32_t
	slurmd_user_name                  []byte
	slurmctld_debug                   uint16_t
	slurmctld_logfile                 []byte
	slurmctld_pidfile                 []byte
	slurmctld_plugstack               []byte
	slurmctld_port                    uint32_t
	slurmctld_port_count              uint16_t
	slurmctld_syslog_debug            uint16_t
	slurmctld_timeout                 uint16_t
	slurmd_debug                      uint16_t
	slurmd_logfile                    []byte
	slurmd_pidfile                    []byte
	slurmd_port                       uint32_t
	slurmd_spooldir                   []byte
	slurmd_syslog_debug               uint16_t
	slurmd_timeout                    uint16_t
	srun_epilog                       []byte
	srun_port_range                   []uint16_t
	srun_prolog                       []byte
	state_save_location               []byte
	suspend_exc_nodes                 []byte
	suspend_exc_parts                 []byte
	suspend_program                   []byte
	suspend_rate                      uint16_t
	suspend_time                      uint32_t
	suspend_timeout                   uint16_t
	switch_type                       []byte
	task_epilog                       []byte
	task_plugin                       []byte
	task_plugin_param                 uint32_t
	task_prolog                       []byte
	tmp_fs                            []byte
	topology_param                    []byte
	topology_plugin                   []byte
	track_wckey                       uint16_t
	tree_width                        uint16_t
	unkillable_program                []byte
	unkillable_timeout                uint16_t
	use_pam                           uint16_t
	use_spec_resources                uint16_t
	version                           []byte
	vsize_factor                      uint16_t
	wait_time                         uint16_t
	ref59adaf2e                       *C.slurm_ctl_conf_t
	allocs59adaf2e                    interface{}
}

// slurmd_status_t as declared in slurm/slurm.h:2993
type slurmd_status_t struct {
	booted             time_t
	last_slurmctld_msg time_t
	slurmd_debug       uint16_t
	actual_cpus        uint16_t
	actual_boards      uint16_t
	actual_sockets     uint16_t
	actual_cores       uint16_t
	actual_threads     uint16_t
	actual_real_mem    uint64_t
	actual_tmp_disk    uint32_t
	pid                uint32_t
	hostname           []byte
	slurmd_logfile     []byte
	step_list          []byte
	version            []byte
	ref72c806fa        *C.slurmd_status_t
	allocs72c806fa     interface{}
}

// submit_response_msg_t as declared in slurm/slurm.h:3000
type submit_response_msg_t struct {
	job_id              uint32_t
	step_id             uint32_t
	error_code          uint32_t
	job_submit_user_msg []byte
	ref6c6e601          *C.submit_response_msg_t
	allocs6c6e601       interface{}
}

// update_node_msg_t as declared in slurm/slurm.h:3016
type update_node_msg_t struct {
	features       []byte
	features_act   []byte
	gres           []byte
	node_addr      []byte
	node_hostname  []byte
	node_names     []byte
	node_state     uint32_t
	reason         []byte
	reason_uid     uint32_t
	weight         uint32_t
	ref7d01926e    *C.update_node_msg_t
	allocs7d01926e interface{}
}

// update_front_end_msg_t as declared in slurm/slurm.h:3024
type update_front_end_msg_t struct {
	name           []byte
	node_state     uint32_t
	reason         []byte
	reason_uid     uint32_t
	ref157a0a01    *C.update_front_end_msg_t
	allocs157a0a01 interface{}
}

// update_part_msg_t as declared in slurm/slurm.h:3026
type update_part_msg_t struct {
	allow_alloc_nodes   []byte
	allow_accounts      []byte
	allow_groups        []byte
	allow_qos           []byte
	alternate           []byte
	billing_weights_str []byte
	cluster_name        []byte
	cr_type             uint16_t
	def_mem_per_cpu     uint64_t
	default_time        uint32_t
	deny_accounts       []byte
	deny_qos            []byte
	flags               uint16_t
	grace_time          uint32_t
	max_cpus_per_node   uint32_t
	max_mem_per_cpu     uint64_t
	max_nodes           uint32_t
	max_share           uint16_t
	max_time            uint32_t
	min_nodes           uint32_t
	name                []byte
	node_inx            []int32_t
	nodes               []byte
	over_time_limit     uint16_t
	preempt_mode        uint16_t
	priority_job_factor uint16_t
	priority_tier       uint16_t
	qos_char            []byte
	state_up            uint16_t
	total_cpus          uint32_t
	total_nodes         uint32_t
	tres_fmt_str        []byte
	reff7d95771         *C.update_part_msg_t
	allocsf7d95771      interface{}
}

// job_sbcast_cred_msg_t as declared in slurm/slurm.h:3034
type job_sbcast_cred_msg_t struct {
	job_id        uint32_t
	node_addr     []slurm_addr_t
	node_cnt      uint32_t
	node_list     []byte
	sbcast_cred   []sbcast_cred_t
	refcd77303    *C.job_sbcast_cred_msg_t
	allocscd77303 interface{}
}

// slurm_step_ctx_t as declared in slurm/slurm.h:3037
type slurm_step_ctx_t C.slurm_step_ctx_t

// stats_info_request_msg_t as declared in slurm/slurm.h:3043
type stats_info_request_msg_t struct {
	command_id     uint16_t
	refc3862c55    *C.stats_info_request_msg_t
	allocsc3862c55 interface{}
}

// stats_info_response_msg_t as declared in slurm/slurm.h:3094
type stats_info_response_msg_t struct {
	parts_packed            uint32_t
	req_time                time_t
	req_time_start          time_t
	server_thread_count     uint32_t
	agent_queue_size        uint32_t
	dbd_agent_queue_size    uint32_t
	schedule_cycle_max      uint32_t
	schedule_cycle_last     uint32_t
	schedule_cycle_sum      uint32_t
	schedule_cycle_counter  uint32_t
	schedule_cycle_depth    uint32_t
	schedule_queue_len      uint32_t
	jobs_submitted          uint32_t
	jobs_started            uint32_t
	jobs_completed          uint32_t
	jobs_canceled           uint32_t
	jobs_failed             uint32_t
	jobs_running            uint32_t
	jobs_running_ts         time_t
	bf_backfilled_jobs      uint32_t
	bf_last_backfilled_jobs uint32_t
	bf_backfilled_pack_jobs uint32_t
	bf_cycle_counter        uint32_t
	bf_cycle_sum            uint64_t
	bf_cycle_last           uint32_t
	bf_cycle_max            uint32_t
	bf_last_depth           uint32_t
	bf_last_depth_try       uint32_t
	bf_depth_sum            uint32_t
	bf_depth_try_sum        uint32_t
	bf_queue_len            uint32_t
	bf_queue_len_sum        uint32_t
	bf_when_last_cycle      time_t
	bf_active               uint32_t
	rpc_type_size           uint32_t
	rpc_type_id             []uint16_t
	rpc_type_cnt            []uint32_t
	rpc_type_time           []uint64_t
	rpc_user_size           uint32_t
	rpc_user_id             []uint32_t
	rpc_user_cnt            []uint32_t
	rpc_user_time           []uint64_t
	ref5aae8b04             *C.stats_info_response_msg_t
	allocs5aae8b04          interface{}
}

// trigger_info_t as declared in slurm/slurm.h:3138
type trigger_info_t struct {
	flags          uint16_t
	trig_id        uint32_t
	res_type       uint16_t
	res_id         []byte
	trig_type      uint32_t
	offset         uint16_t
	user_id        uint32_t
	program        []byte
	ref1497cae7    *C.trigger_info_t
	allocs1497cae7 interface{}
}

// trigger_info_msg_t as declared in slurm/slurm.h:3143
type trigger_info_msg_t struct {
	record_count   uint32_t
	trigger_array  []trigger_info_t
	refe1a7ae31    *C.trigger_info_msg_t
	allocse1a7ae31 interface{}
}

// slurm_license_info_t as declared in slurm/slurm.h:3155
type slurm_license_info_t struct {
	name           []byte
	total          uint32_t
	in_use         uint32_t
	available      uint32_t
	remote         uint8_t
	refb718b5ee    *C.slurm_license_info_t
	allocsb718b5ee interface{}
}

// license_info_msg_t as declared in slurm/slurm.h:3163
type license_info_msg_t struct {
	last_update    time_t
	num_lic        uint32_t
	lic_array      []slurm_license_info_t
	refe48f58cb    *C.license_info_msg_t
	allocse48f58cb interface{}
}

// job_array_resp_msg_t as declared in slurm/slurm.h:3169
type job_array_resp_msg_t struct {
	job_array_count uint32_t
	job_array_id    [][]byte
	error_code      []uint32_t
	ref6ab3c2ed     *C.job_array_resp_msg_t
	allocs6ab3c2ed  interface{}
}

// assoc_mgr_info_msg_t as declared in slurm/slurm.h:3178
type assoc_mgr_info_msg_t struct {
	assoc_list     List
	qos_list       List
	tres_cnt       uint32_t
	tres_names     [][]byte
	user_list      List
	refb9a59e94    *C.assoc_mgr_info_msg_t
	allocsb9a59e94 interface{}
}

// assoc_mgr_info_request_msg_t as declared in slurm/slurm.h:3189
type assoc_mgr_info_request_msg_t struct {
	acct_list     List
	flags         uint32_t
	qos_list      List
	user_list     List
	ref8173f97    *C.assoc_mgr_info_request_msg_t
	allocs8173f97 interface{}
}

// network_callerid_msg_t as declared in slurm/slurm.h:3197
type network_callerid_msg_t struct {
	ip_src         [16]byte
	ip_dst         [16]byte
	port_src       uint32_t
	port_dst       uint32_t
	af             int32_t
	refab82ee84    *C.network_callerid_msg_t
	allocsab82ee84 interface{}
}

// job_step_kill_msg_t as declared in slurm/slurm.h:3450
type job_step_kill_msg_t struct {
	job_id         uint32_t
	sjob_id        []byte
	job_step_id    uint32_t
	signal         uint16_t
	flags          uint16_t
	sibling        []byte
	refd4cb75f2    *C.job_step_kill_msg_t
	allocsd4cb75f2 interface{}
}

// burst_buffer_pool_t as declared in slurm/slurm.h:5065
type burst_buffer_pool_t struct {
	granularity    uint64_t
	name           []byte
	total_space    uint64_t
	used_space     uint64_t
	unfree_space   uint64_t
	ref907a8fe1    *C.burst_buffer_pool_t
	allocs907a8fe1 interface{}
}

// burst_buffer_resv_t as declared in slurm/slurm.h:5080
type burst_buffer_resv_t struct {
	account        []byte
	array_job_id   uint32_t
	array_task_id  uint32_t
	create_time    time_t
	job_id         uint32_t
	name           []byte
	partition      []byte
	pool           []byte
	qos            []byte
	size           uint64_t
	state          uint16_t
	user_id        uint32_t
	ref9c7cf8ca    *C.burst_buffer_resv_t
	allocs9c7cf8ca interface{}
}

// burst_buffer_use_t as declared in slurm/slurm.h:5085
type burst_buffer_use_t struct {
	user_id        uint32_t
	used           uint64_t
	ref4d977eac    *C.burst_buffer_use_t
	allocs4d977eac interface{}
}

// burst_buffer_info_t as declared in slurm/slurm.h:5116
type burst_buffer_info_t struct {
	allow_users           []byte
	default_pool          []byte
	create_buffer         []byte
	deny_users            []byte
	destroy_buffer        []byte
	flags                 uint32_t
	get_sys_state         []byte
	granularity           uint64_t
	pool_cnt              uint32_t
	pool_ptr              []burst_buffer_pool_t
	name                  []byte
	other_timeout         uint32_t
	stage_in_timeout      uint32_t
	stage_out_timeout     uint32_t
	start_stage_in        []byte
	start_stage_out       []byte
	stop_stage_in         []byte
	stop_stage_out        []byte
	total_space           uint64_t
	unfree_space          uint64_t
	used_space            uint64_t
	validate_timeout      uint32_t
	buffer_count          uint32_t
	burst_buffer_resv_ptr []burst_buffer_resv_t
	use_count             uint32_t
	burst_buffer_use_ptr  []burst_buffer_use_t
	reff68d04d1           *C.burst_buffer_info_t
	allocsf68d04d1        interface{}
}

// burst_buffer_info_msg_t as declared in slurm/slurm.h:5121
type burst_buffer_info_msg_t struct {
	burst_buffer_array []burst_buffer_info_t
	record_count       uint32_t
	ref51ffcb89        *C.burst_buffer_info_msg_t
	allocs51ffcb89     interface{}
}

// __gwchar_t type as declared in include/inttypes.h:38
type __gwchar_t int32

// imaxdiv_t as declared in include/inttypes.h:282
type imaxdiv_t struct {
	quot           int
	rem            int
	ref873cd79e    *C.imaxdiv_t
	allocs873cd79e interface{}
}

// int8_t type as declared in include/stdint.h:36
type int8_t byte

// int16_t type as declared in include/stdint.h:37
type int16_t int16

// int32_t type as declared in include/stdint.h:38
type int32_t int32

// int64_t type as declared in include/stdint.h:40
type int64_t int

// uint8_t type as declared in include/stdint.h:48
type uint8_t byte

// uint16_t type as declared in include/stdint.h:49
type uint16_t uint16

// uint32_t type as declared in include/stdint.h:51
type uint32_t uint32

// uint64_t type as declared in include/stdint.h:55
type uint64_t uint

// int_least8_t type as declared in include/stdint.h:65
type int_least8_t byte

// int_least16_t type as declared in include/stdint.h:66
type int_least16_t int16

// int_least32_t type as declared in include/stdint.h:67
type int_least32_t int32

// int_least64_t type as declared in include/stdint.h:69
type int_least64_t int

// uint_least8_t type as declared in include/stdint.h:76
type uint_least8_t byte

// uint_least16_t type as declared in include/stdint.h:77
type uint_least16_t uint16

// uint_least32_t type as declared in include/stdint.h:78
type uint_least32_t uint32

// uint_least64_t type as declared in include/stdint.h:80
type uint_least64_t uint

// int_fast8_t type as declared in include/stdint.h:90
type int_fast8_t byte

// int_fast16_t type as declared in include/stdint.h:92
type int_fast16_t int

// int_fast32_t type as declared in include/stdint.h:93
type int_fast32_t int

// int_fast64_t type as declared in include/stdint.h:94
type int_fast64_t int

// uint_fast8_t type as declared in include/stdint.h:103
type uint_fast8_t byte

// uint_fast16_t type as declared in include/stdint.h:105
type uint_fast16_t uint

// uint_fast32_t type as declared in include/stdint.h:106
type uint_fast32_t uint

// uint_fast64_t type as declared in include/stdint.h:107
type uint_fast64_t uint

// intptr_t type as declared in include/stdint.h:119
type intptr_t int

// uintptr_t type as declared in include/stdint.h:122
type uintptr_t uint

// intmax_t type as declared in include/stdint.h:134
type intmax_t int

// uintmax_t type as declared in include/stdint.h:135
type uintmax_t uint

// size_t type as declared in include/stddef.h:212
type size_t uint

// wchar_t type as declared in include/stddef.h:324
type wchar_t int32

// wint_t type as declared in include/stddef.h:353
type wint_t uint32

// FILE as declared in include/stdio.h:48
type FILE struct {
	_flags          int32
	_IO_read_ptr    []byte
	_IO_read_end    []byte
	_IO_read_base   []byte
	_IO_write_base  []byte
	_IO_write_ptr   []byte
	_IO_write_end   []byte
	_IO_buf_base    []byte
	_IO_buf_end     []byte
	_IO_save_base   []byte
	_IO_backup_base []byte
	_IO_save_end    []byte
	_fileno         int32
	_flags2         int32
	_old_offset     __off_t
	_cur_column     uint16
	_vtable_offset  byte
	_shortbuf       [1]byte
	_lock           *_IO_lock_t
	_offset         __off64_t
	__pad1          unsafe.Pointer
	__pad2          unsafe.Pointer
	__pad3          unsafe.Pointer
	__pad4          unsafe.Pointer
	__pad5          size_t
	_mode           int32
	_unused2        [20]byte
	refba0adba4     *C.FILE
	allocsba0adba4  interface{}
}

// off_t type as declared in include/stdio.h:90
type off_t int

// ssize_t type as declared in include/stdio.h:102
type ssize_t int

// fpos_t as declared in include/stdio.h:110
type fpos_t struct {
	__pos          __off_t
	__state        __mbstate_t
	refc5e101c9    *C._G_fpos_t
	allocsc5e101c9 interface{}
}

// __int8_t type as declared in bits/types.h:36
type __int8_t byte

// __uint8_t type as declared in bits/types.h:37
type __uint8_t byte

// __int16_t type as declared in bits/types.h:38
type __int16_t int16

// __uint16_t type as declared in bits/types.h:39
type __uint16_t uint16

// __int32_t type as declared in bits/types.h:40
type __int32_t int32

// __uint32_t type as declared in bits/types.h:41
type __uint32_t uint32

// __int64_t type as declared in bits/types.h:43
type __int64_t int

// __uint64_t type as declared in bits/types.h:44
type __uint64_t uint

// __quad_t type as declared in bits/types.h:52
type __quad_t int

// __u_quad_t type as declared in bits/types.h:53
type __u_quad_t uint

// __dev_t type as declared in bits/types.h:133
type __dev_t uint

// __uid_t type as declared in bits/types.h:134
type __uid_t uint32

// __gid_t type as declared in bits/types.h:135
type __gid_t uint32

// __ino_t type as declared in bits/types.h:136
type __ino_t uint

// __ino64_t type as declared in bits/types.h:137
type __ino64_t uint

// __mode_t type as declared in bits/types.h:138
type __mode_t uint32

// __nlink_t type as declared in bits/types.h:139
type __nlink_t uint

// __off_t type as declared in bits/types.h:140
type __off_t int

// __off64_t type as declared in bits/types.h:141
type __off64_t int

// __pid_t type as declared in bits/types.h:142
type __pid_t int32

// __fsid_t as declared in bits/types.h:143
type __fsid_t struct {
	__val          [2]int32
	ref3fddfece    *C.__fsid_t
	allocs3fddfece interface{}
}

// __clock_t type as declared in bits/types.h:144
type __clock_t int

// __rlim_t type as declared in bits/types.h:145
type __rlim_t uint

// __rlim64_t type as declared in bits/types.h:146
type __rlim64_t uint

// __id_t type as declared in bits/types.h:147
type __id_t uint32

// __time_t type as declared in bits/types.h:148
type __time_t int

// __useconds_t type as declared in bits/types.h:149
type __useconds_t uint32

// __suseconds_t type as declared in bits/types.h:150
type __suseconds_t int

// __daddr_t type as declared in bits/types.h:152
type __daddr_t int32

// __key_t type as declared in bits/types.h:153
type __key_t int32

// __clockid_t type as declared in bits/types.h:156
type __clockid_t int32

// __timer_t type as declared in bits/types.h:159
type __timer_t unsafe.Pointer

// __blksize_t type as declared in bits/types.h:162
type __blksize_t int

// __blkcnt_t type as declared in bits/types.h:167
type __blkcnt_t int

// __blkcnt64_t type as declared in bits/types.h:168
type __blkcnt64_t int

// __fsblkcnt_t type as declared in bits/types.h:171
type __fsblkcnt_t uint

// __fsblkcnt64_t type as declared in bits/types.h:172
type __fsblkcnt64_t uint

// __fsfilcnt_t type as declared in bits/types.h:175
type __fsfilcnt_t uint

// __fsfilcnt64_t type as declared in bits/types.h:176
type __fsfilcnt64_t uint

// __fsword_t type as declared in bits/types.h:179
type __fsword_t int

// __ssize_t type as declared in bits/types.h:181
type __ssize_t int

// __syscall_slong_t type as declared in bits/types.h:184
type __syscall_slong_t int

// __syscall_ulong_t type as declared in bits/types.h:186
type __syscall_ulong_t uint

// __loff_t type as declared in bits/types.h:190
type __loff_t int

// __qaddr_t type as declared in bits/types.h:191
type __qaddr_t []int

// __caddr_t type as declared in bits/types.h:192
type __caddr_t []byte

// __intptr_t type as declared in bits/types.h:195
type __intptr_t int

// __socklen_t type as declared in bits/types.h:198
type __socklen_t uint32

// _IO_lock_t type as declared in include/libio.h:155
type _IO_lock_t [0]byte

// _G_fpos_t as declared in include/_G_config.h:25
type _G_fpos_t struct {
	__pos          __off_t
	__state        __mbstate_t
	refc5e101c9    *C._G_fpos_t
	allocsc5e101c9 interface{}
}

// _G_fpos64_t as declared in include/_G_config.h:30
type _G_fpos64_t struct {
	__pos          __off64_t
	__state        __mbstate_t
	ref70194a78    *C._G_fpos64_t
	allocs70194a78 interface{}
}

// __mbstate_t as declared in include/wchar.h:94
type __mbstate_t struct {
	__count        int32
	refde5795dc    *C.__mbstate_t
	allocsde5795dc interface{}
}

// quad_t type as declared in sys/types.h:37
type quad_t int

// u_quad_t type as declared in sys/types.h:38
type u_quad_t uint

// fsid_t as declared in sys/types.h:39
type fsid_t struct {
	__val          [2]int32
	ref3fddfece    *C.__fsid_t
	allocs3fddfece interface{}
}

// loff_t type as declared in sys/types.h:44
type loff_t int

// ino_t type as declared in sys/types.h:48
type ino_t uint

// dev_t type as declared in sys/types.h:60
type dev_t uint

// gid_t type as declared in sys/types.h:65
type gid_t uint32

// mode_t type as declared in sys/types.h:70
type mode_t uint32

// nlink_t type as declared in sys/types.h:75
type nlink_t uint

// uid_t type as declared in sys/types.h:80
type uid_t uint32

// pid_t type as declared in sys/types.h:98
type pid_t int32

// id_t type as declared in sys/types.h:104
type id_t uint32

// daddr_t type as declared in sys/types.h:115
type daddr_t int32

// caddr_t type as declared in sys/types.h:116
type caddr_t []byte

// key_t type as declared in sys/types.h:122
type key_t int32

// u_int8_t type as declared in sys/types.h:173
type u_int8_t byte

// u_int16_t type as declared in sys/types.h:174
type u_int16_t uint16

// u_int32_t type as declared in sys/types.h:175
type u_int32_t uint32

// u_int64_t type as declared in sys/types.h:177
type u_int64_t uint

// register_t type as declared in sys/types.h:182
type register_t int32

// blksize_t type as declared in sys/types.h:228
type blksize_t int

// blkcnt_t type as declared in sys/types.h:235
type blkcnt_t int

// fsblkcnt_t type as declared in sys/types.h:239
type fsblkcnt_t uint

// fsfilcnt_t type as declared in sys/types.h:243
type fsfilcnt_t uint

// clock_t type as declared in include/time.h:59
type clock_t int

// time_t type as declared in include/time.h:75
type time_t int

// clockid_t type as declared in include/time.h:91
type clockid_t int32

// timer_t type as declared in include/time.h:103
type timer_t unsafe.Pointer

// sigset_t as declared in sys/select.h:37
type sigset_t struct {
	__val          [16]uint
	refa44732af    *C.__sigset_t
	allocsa44732af interface{}
}

// suseconds_t type as declared in sys/select.h:48
type suseconds_t int

// __sig_atomic_t type as declared in bits/sigset.h:23
type __sig_atomic_t int32

// __sigset_t as declared in bits/sigset.h:31
type __sigset_t struct {
	__val          [16]uint
	refa44732af    *C.__sigset_t
	allocsa44732af interface{}
}

// pthread_t type as declared in bits/pthreadtypes.h:60
type pthread_t uint

// pthread_attr_t as declared in bits/pthreadtypes.h:69
const sizeofpthread_attr_t = unsafe.Sizeof(C.pthread_attr_t{})

type pthread_attr_t [sizeofpthread_attr_t]byte

// __pthread_list_t as declared in bits/pthreadtypes.h:79
type __pthread_list_t struct {
	ref3eec1ba6    *C.__pthread_list_t
	allocs3eec1ba6 interface{}
}

// pthread_mutex_t as declared in bits/pthreadtypes.h:118
const sizeofpthread_mutex_t = unsafe.Sizeof(C.pthread_mutex_t{})

type pthread_mutex_t [sizeofpthread_mutex_t]byte

// pthread_mutexattr_t as declared in bits/pthreadtypes.h:124
const sizeofpthread_mutexattr_t = unsafe.Sizeof(C.pthread_mutexattr_t{})

type pthread_mutexattr_t [sizeofpthread_mutexattr_t]byte

// pthread_cond_t as declared in bits/pthreadtypes.h:144
const sizeofpthread_cond_t = unsafe.Sizeof(C.pthread_cond_t{})

type pthread_cond_t [sizeofpthread_cond_t]byte

// pthread_condattr_t as declared in bits/pthreadtypes.h:150
const sizeofpthread_condattr_t = unsafe.Sizeof(C.pthread_condattr_t{})

type pthread_condattr_t [sizeofpthread_condattr_t]byte

// pthread_key_t type as declared in bits/pthreadtypes.h:154
type pthread_key_t uint32

// pthread_once_t type as declared in bits/pthreadtypes.h:158
type pthread_once_t int32

// pthread_rwlock_t as declared in bits/pthreadtypes.h:204
const sizeofpthread_rwlock_t = unsafe.Sizeof(C.pthread_rwlock_t{})

type pthread_rwlock_t [sizeofpthread_rwlock_t]byte

// pthread_rwlockattr_t as declared in bits/pthreadtypes.h:210
const sizeofpthread_rwlockattr_t = unsafe.Sizeof(C.pthread_rwlockattr_t{})

type pthread_rwlockattr_t [sizeofpthread_rwlockattr_t]byte

// pthread_spinlock_t type as declared in bits/pthreadtypes.h:216
type pthread_spinlock_t int32

// pthread_barrier_t as declared in bits/pthreadtypes.h:225
const sizeofpthread_barrier_t = unsafe.Sizeof(C.pthread_barrier_t{})

type pthread_barrier_t [sizeofpthread_barrier_t]byte

// pthread_barrierattr_t as declared in bits/pthreadtypes.h:231
const sizeofpthread_barrierattr_t = unsafe.Sizeof(C.pthread_barrierattr_t{})

type pthread_barrierattr_t [sizeofpthread_barrierattr_t]byte

// __locale_t as declared in include/xlocale.h:39
type __locale_t struct {
	__ctype_b       []uint16
	__ctype_tolower []int32
	__ctype_toupper []int32
	__names         [13]string
	refa5868e3f     *C.__locale_t
	allocsa5868e3f  interface{}
}

// locale_t as declared in include/xlocale.h:42
type locale_t struct {
	__ctype_b       []uint16
	__ctype_tolower []int32
	__ctype_toupper []int32
	__names         [13]string
	refa5868e3f     *C.__locale_t
	allocsa5868e3f  interface{}
}

// useconds_t type as declared in include/unistd.h:255
type useconds_t uint32

// socklen_t type as declared in include/unistd.h:274
type socklen_t uint32

// in_addr_t type as declared in netinet/in.h:31
type in_addr_t uint32

// in_port_t type as declared in netinet/in.h:118
type in_port_t uint16

// sa_family_t type as declared in bits/sockaddr.h:28
type sa_family_t uint16
