package main

import (
	"flag"
	"fmt"
	"net"
)

const MAXPORTS int = 65535

func worker_bees(target string, port int, live chan int) {
	connection, err := net.Dial("tcp", fmt.Sprintf("%v:%v", target, port))
	if err != nil {
		live <- 0
		return
	}
	connection.Close()
	live <- port
}

func check_ports(target string) []int {
	var open_ports []int
	var idx int = 1
	live := make(chan int)

	for port := 1; port <= MAXPORTS; port++ {
		go worker_bees(target, port, live)
	}

	for range MAXPORTS {
		port := <-live
		if port != 0 {
			open_ports = append(open_ports, port)
		}

		progress := (float64(idx) / float64(MAXPORTS)) * 100
		fmt.Printf("\rProgress: %.2f%%", progress)
		idx++
	}

	fmt.Printf("\n")
	return open_ports
}

func main() {
	target := flag.String("t", "", "Target (example.com)")
	flag.Parse()
	if *target == "" {
		fmt.Printf("Usage: ./main -t <example.com>\n")
		return
	}

	open_ports := check_ports(*target)
	if len(open_ports) == 0 {
		fmt.Printf("[!] No Port Was Opened!\n")
		return
	}

	for _, port := range open_ports {
		fmt.Printf("[+] Port %v is Opened!\n", port)
	}
}
