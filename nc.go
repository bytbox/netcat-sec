package main

import (
	"bufio"
	"crypto/tls"
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
	useTLS = flag.Bool("tls", true, "Use TLS")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("netcat-sec %s\n", VERSION)
	}

	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Need host and port")
		return
	}
	host, port := args[0], args[1]

	var (
		conn net.Conn
		err  error
	)

	if *useTLS {
		conn, err = tls.Dial("tcp", host+":"+port, nil)
	} else {
		conn, err = net.Dial("tcp", host+":"+port)
	}
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
			fmt.Println(string(l))
		}
	}()

	<-make(chan interface{})
}
