package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	count    = flag.Int("c", 4, "Number of Pings: <= 0 means forever")
	interval = flag.Duration("i", time.Second, "Interval between Pings")
	timeout  = flag.Duration("w", 5*time.Second, "time to wait for reply")
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] host:port\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Print("host:port is mandatory\n\n")
		flag.Usage()
		os.Exit(1)
	}

	target := flag.Arg(0)
	fmt.Println("ping", target)

	if *count <= 0 {
		fmt.Println("ctrl+c stops ping")
	}

	msg := 0

	for (*count <= 0) || (msg < *count) {
		msg++
		fmt.Print(msg, "  ")

		start := time.Now()
		c, err := net.DialTimeout("tcp", target, *timeout)
		dur := time.Since(start)

		if err != nil {
			fmt.Printf("fail in %s: %v\n", dur, err)
			if nErr, ok := err.(net.Error); !ok || !nErr.Temporary() {
				os.Exit(1)
			}
		} else {
			_ = c.Close()
			//fmt.Print(dur, " :")
			fmt.Printf("%.4v connect from: %s to: %s \n", dur, c.LocalAddr().String(), c.RemoteAddr().String())
		}
		time.Sleep(*interval)
	}
}
