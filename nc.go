package main

import (
	"bufio"
	//"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
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

	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	go func() {
		r := bufio.NewReader(os.Stdin)
		for {
			l, _, err := r.ReadLine()
			if err != nil {
				os.Exit(0)
			}
			_, err = conn.Write([]byte(string(l)+"\n"))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}()

	go func() {
		r := bufio.NewReader(conn)
		for {
			l, _, err := r.ReadLine()
			if err != nil {
				if err == io.EOF {
					os.Exit(0)
				}
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Print(string(l))
		}
	}()

	<-make(chan interface{})
}
