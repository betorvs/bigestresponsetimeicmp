package usecase

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/betorvs/biggestresponsetimeicmp/gateway"

	"github.com/betorvs/biggestresponsetimeicmp/domain"
)

//var biggest is used to keep all ipAddr:ResponseTime
var biggest map[string]int64

//printJump
func printJump(host domain.JumpedHost) {
	addr := fmt.Sprintf("%v.%v.%v.%v", host.Address[0], host.Address[1], host.Address[2], host.Address[3])
	hostOrAddr := addr
	if host.Hostname != "" {
		hostOrAddr = host.Hostname
	}
	if host.Success {
		fmt.Printf("%-3d %v (%v)  %v\n", host.TTL, hostOrAddr, addr, host.ElapsedTime)
	} else {
		fmt.Printf("%-3d *\n", host.TTL)
	}
}

//FindBiggestResponseTime
func FindBiggestResponseTime(biggest map[[4]byte]int64) ([4]byte, int64) {
	big := int64(0)
	var key [4]byte
	for k, v := range biggest {
		if v > big {
			big = v
			key = k
		}
	}
	return key, big
}

//CalculateBiggestResponseTime func main usecase
func CalculateBiggestResponseTime(host string) {
	ipAddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	biggest := make(map[[4]byte]int64)
	fmt.Printf("Traceroute to %v (%v), %v hops max, %v byte packets\n", host, ipAddr, gateway.DefaultMaxJumps, gateway.DefaultPacketSize)
	c := make(chan domain.JumpedHost, 0)
	go func() {

		for {
			hop, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
			biggest[hop.Address] = int64(hop.ElapsedTime)
			printJump(hop)
		}
	}()

	_, err = gateway.Jumping(host, c)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	key, big := FindBiggestResponseTime(biggest)
	addr := fmt.Sprintf("%v.%v.%v.%v", key[0], key[1], key[2], key[3])
	fmt.Println("The biggest response time is: ", string(addr), "with", time.Duration(big))

}
