package gateway

/*

This gateway is based on aeden/traceroute library.

*/

import (
	"errors"
	"net"
	"syscall"
	"time"

	"github.com/betorvs/biggestresponsetimeicmp/domain"
)

//DefaultPort const
const DefaultPort = 33434

//DefaultMaxJumps const
const DefaultMaxJumps = 20

//DefaultFirstJump const
const DefaultFirstJump = 1

//DefaultTimeoutMili const
const DefaultTimeoutMili = 500

//DefaultRetries const
const DefaultRetries = 3

//DefaultPacketSize const
const DefaultPacketSize = 52

// localInterface : local network interface
func localInterface() (addr [4]byte, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if len(ipnet.IP.To4()) == net.IPv4len {
				copy(addr[:], ipnet.IP.To4())
				return
			}
		}
	}
	err = errors.New("Are you connected to the Internet?")
	return
}

// resolveHostname : convert a hostname to a 4 byte IP address.
func resolveHostname(destinationHost string) (destinationAddr [4]byte, err error) {
	addrs, err := net.LookupHost(destinationHost)
	if err != nil {
		return
	}
	addr := addrs[0]

	ipAddr, err := net.ResolveIPAddr("ip", addr)
	if err != nil {
		return
	}
	copy(destinationAddr[:], ipAddr.IP.To4())
	return
}

//Jumping func do an icmp/udp test to given destinationHost
func Jumping(destinationHost string, c ...chan domain.JumpedHost) (spots domain.FalloutResult, err error) {
	spots.Places = []domain.JumpedHost{}
	destiny, _ := resolveHostname(destinationHost)
	spots.DestinationAddress = destiny
	networkInterface, err := localInterface()
	if err != nil {
		return
	}
	timeoutMili := (int64)(DefaultTimeoutMili)
	tml := syscall.NsecToTimeval(1000 * 1000 * timeoutMili)
	ttl := DefaultFirstJump
	retry := 0
	for {
		start := time.Now()
		inputInterface, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
		if err != nil {
			return spots, err
		}
		outputInterface, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
		if err != nil {
			return spots, err
		}
		syscall.SetsockoptInt(outputInterface, 0x0, syscall.IP_TTL, ttl)
		syscall.SetsockoptTimeval(inputInterface, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tml)
		defer syscall.Close(inputInterface)
		defer syscall.Close(outputInterface)
		syscall.Bind(inputInterface, &syscall.SockaddrInet4{Port: DefaultPort, Addr: networkInterface})
		syscall.Sendto(outputInterface, []byte{0x0}, 0, &syscall.SockaddrInet4{Port: DefaultPort, Addr: destiny})
		var p = make([]byte, DefaultPacketSize)
		n, from, err := syscall.Recvfrom(inputInterface, p, 0)
		elapsed := time.Since(start)
		if err == nil {
			actualAddr := from.(*syscall.SockaddrInet4).Addr
			visited := domain.JumpedHost{Success: true, Address: actualAddr, N: n, ElapsedTime: elapsed, TTL: ttl}
			currentHost, err := net.LookupAddr(visited.AddrToString())
			if err == nil {
				visited.Hostname = currentHost[0]
			}
			domain.Notify(visited, c)

			spots.Places = append(spots.Places, visited)

			ttl++
			retry = 0

			if ttl > DefaultMaxJumps || actualAddr == destiny {
				domain.CloseNotify(c)
				return spots, nil
			}
		} else {
			retry++
			if retry > DefaultRetries {
				domain.Notify(domain.JumpedHost{Success: false, TTL: ttl}, c)
				ttl++
				retry = 0
			}

			if ttl > DefaultMaxJumps {
				domain.CloseNotify(c)
				return spots, nil
			}
		}
	}
}
