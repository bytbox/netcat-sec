include ${GOROOT}/src/Make.inc

TARG = netcat-sec
GOFILES = nc.go

include ${GOROOT}/src/Make.cmd

fmt:
	gofmt -w ${GOFILES}

