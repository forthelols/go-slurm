#pragma once
#include <slurm/slurm.h>

/*
 * wrap_slurm_submit_batch_job - issue RPC to submit a job for later execution
 * NOTE: free the response using slurm_free_submit_response_response_msg
 * IN job_desc_msg - description of batch job request
 * RET a pointer to a valid submit_response_msg_t allocated by the underlying
 * call to slurm_submit_batch_job, otherwise return NULL and set errno to
 * indicate the error
 */
extern submit_response_msg_t* wrap_slurm_submit_batch_job(job_desc_msg_t *job_desc_msg);

