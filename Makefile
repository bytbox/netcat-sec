include ${GOROOT}/src/Make.inc

TARG = netcat-tls
GOFILES = nc.go

include ${GOROOT}/src/Make.cmd

fmt:
	gofmt -w ${GOFILES}

