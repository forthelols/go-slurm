#include "slurm_helpers.h"
#include <stddef.h>

submit_response_msg_t* wrap_slurm_submit_batch_job(job_desc_msg_t *job_desc_msg) {
    submit_response_msg_t* response;
    if (slurm_submit_batch_job(job_desc_msg, &response) != 0) {
        // API docs don't say nothing about the state of the
        // out pointer in case of failure. Let's free it and
        // return NULL just in case.
        slurm_free_submit_response_response_msg(response);
        response = NULL;
    }
    return response;
}

job_info_msg_t* wrap_slurm_load_job(uint32_t job_id, uint16_t show_flags) {
    job_info_msg_t* response;
    if (slurm_load_job(&response, job_id, show_flags) != 0) {
        // API docs don't say nothing about the state of the
        // out pointer in case of failure. Let's free it and
        // return NULL just in case.
        slurm_free_job_info_msg(response);
        response = NULL;
    }
    return response;
}

