package main

import (
	"fmt"
	"gek_net"
)

func getActiveNetworkInterface() error {
	netInterfaces, err := gek_net.GetNetInterfaces()
	if err != nil {
		return err
	}

	for _, netInterface := range netInterfaces {
		if !netInterface.Flag.Up {
			continue
		}
		fmt.Printf("Active network interface: %s\n", netInterface.Name)
		for _, ipv4 := range netInterface.Ipv4 {
			fmt.Printf("IPv4: %s\n", ipv4)
		}
		for _, ipv6 := range netInterface.Ipv6 {
			fmt.Printf("IPv6: %s\n", ipv6)
		}
	}

	return err
}