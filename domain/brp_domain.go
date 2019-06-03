package domain

import (
	"fmt"
	"time"
)

//JumpedHost : object from icmp test
type JumpedHost struct {
	Success     bool
	Address     [4]byte
	Hostname    string
	N           int
	ElapsedTime time.Duration
	TTL         int
}

//FalloutResult : resturned objects from icmp jumps
type FalloutResult struct {
	DestinationAddress [4]byte
	Places             []JumpedHost
}

//AddrToString do
func (host *JumpedHost) AddrToString() string {
	return fmt.Sprintf("%v.%v.%v.%v", host.Address[0], host.Address[1], host.Address[2], host.Address[3])
}

//HostOrAddressString do
func (host *JumpedHost) HostOrAddressString() string {
	hostOrAddr := host.AddrToString()
	if host.Hostname != "" {
		hostOrAddr = host.Hostname
	}
	return hostOrAddr
}

//Notify : used to update JumpedHost struct
func Notify(host JumpedHost, channels []chan JumpedHost) {
	for _, c := range channels {
		c <- host
	}
}

//CloseNotify : used to close channels
func CloseNotify(channels []chan JumpedHost) {
	for _, c := range channels {
		close(c)
	}
}
