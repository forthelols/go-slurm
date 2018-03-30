all:
	c-for-go -nostamp -ccincl slurm.yml

clean:
	rm -f slurm/doc.go \
	      slurm/types.go \
	      slurm/const.go \
	      slurm/slurm.go \
	      slurm/cgo_helpers.go \
	      slurm/cgo_helpers.c \
	      slurm/cgo_helpers.h

test: all
	cd slurm && go build
