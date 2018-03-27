all:
	c-for-go -ccincl -out .. slurm.yml

clean:
	rm -f doc.go types.go const.go \
	      cgo_helpers.go cgo_helpers.c cgo_helpers.h \
	      slurm.go

test:
	go build
