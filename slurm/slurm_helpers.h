#pragma once
#include <slurm/slurm.h>
#include <stdint.h>

// Slurm API docs don't say nothing about the state of output pointer arguments.
// Since all slurm_free_* calls never fail, we use a conservative approach in
// case of error:
//   1. explicitly free the response struct regardless of the value of the
//   output pointer, and
//   2. return NULL


// wrap_slurm_submit_batch_job - issue RPC to submit a job for later execution
// IN  job_desc_msg - description of batch job request
// RET a pointer to a valid submit_response_msg_t allocated by the underlying
//     call to slurm_submit_batch_job, otherwise return NULL and set errno to
//     indicate the error
// NOTE: free the response using slurm_free_submit_response_response_msg
//       (freeing a returned NULL is not needed but allowed)
extern submit_response_msg_t* wrap_slurm_submit_batch_job(job_desc_msg_t *job_desc_msg);

// wrap_slurm_load_job - issue RPC to get job information for one job ID
// IN  job_id -  ID of job we want information about
// IN  show_flags - job filtering options
// RET a pointer to a valid job_info_msg_t allocated by the underlying
//     call to slurm_load_job, otherwise return NULL and set errno to
//     indicate the error
// NOTE: free the response using slurm_free_job_info_msg
//       (freeing a returned NULL is not needed but allowed)
extern job_info_msg_t* wrap_slurm_load_job(uint32_t job_id, uint16_t show_flags);
