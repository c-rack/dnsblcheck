/*

dnsblcheck

Reads one or more DNSBL hosts from stdin and checks if the IP address is listed

-------------------------------------------------------------------------------

Copyright (c) 2013 Constantin Rack

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	reverse_ip string
)

func check(dnsbl string, quit chan int) {
	hostname := reverse_ip + dnsbl
	ips, err := net.LookupIP(hostname)

	if err == nil {
		if ips[0].IsLoopback() {
			text, _ := net.LookupTXT(hostname)
			fmt.Println(text)
			os.Exit(1)
		}
	}
	quit <- 1
}

func getIp() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		printUsage()
		os.Exit(1)
	}

	ip := flag.Arg(0)

	if net.ParseIP(ip) == nil {
		fmt.Println("Not an IP address:", ip)
		os.Exit(1)
	}

	tokens := strings.Split(ip, ".")
	reverse_ip = tokens[3] + "." + tokens[2] + "." + tokens[1] + "." + tokens[0] + "."
}

func printUsage() {
	fmt.Println("dnsblcheck reads one or more DNSBL hosts from stdin and checks if the IP address is listed")
	fmt.Println("Usage:   dnsblcheck <ip>")
	fmt.Println("Example: cat dnsbl.txt | dnsblcheck 127.0.0.1")
}

func main() {
	getIp()
	coroutines := 0
	quit := make(chan int)
	in := bufio.NewReader(os.Stdin)
	for {
		dnsbl, err := in.ReadString('\n')
		if err != nil {
			break
		}
		coroutines++
		go check(strings.TrimSpace(dnsbl), quit)
	}
	for j := 0; j < coroutines; _, j = <-quit, j+1 {
	}
	fmt.Println("Not listed")
}
