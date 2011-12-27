package main

import (
	//"crypto/tls"
	"flag"
	"fmt"
	//"net"
)

const VERSION = `0.1`

// Flags
var (
	version = flag.Bool("V", false, "Display version information and exit")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("netcat-tls %s\n", VERSION)
	}
}
