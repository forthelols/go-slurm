#include "slurm_helpers.h"
#include <stddef.h>

submit_response_msg_t* wrap_slurm_submit_batch_job(job_desc_msg_t *job_desc_msg) {
    submit_response_msg_t* response;
    if (slurm_submit_batch_job(job_desc_msg, &response) != 0) {
        // API doesn't say nothing about the state of
        // out pointer in case of failure. Let's free
        // it and return NULL just in case.
        slurm_free_submit_response_response_msg(response);
        response = NULL;
    }
    return response;
}
