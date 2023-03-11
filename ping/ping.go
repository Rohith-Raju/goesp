package ping

import (
	"fmt"
	"net"
	"time"

	"github.com/tatsushid/go-fastping"
)

func Ping(ipAddress string) (time.Duration, error) {
	var returnTime time.Duration
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ipAddress)
	if err != nil {
		return 0, err
	}
	fmt.Printf("Pinging : %s\n", ipAddress)
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		returnTime = rtt
	}
	p.OnIdle = func() {
		fmt.Println("finished pinging")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return returnTime, nil
}
