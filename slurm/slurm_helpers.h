#pragma once
#include <slurm/slurm.h>

extern submit_response_msg_t* wrap_slurm_submit_batch_job(job_desc_msg_t *job_desc_msg);

