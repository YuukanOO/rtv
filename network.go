package main

import (
	"net"
	"strings"
)

func getNetworkInformations() *NetworkInfo {
	interfaces, _ := net.Interfaces()

	for _, v := range interfaces {

		if v.Flags&net.FlagLoopback == 0 && !strings.Contains(v.Name, "vir") {
			addresses, _ := v.Addrs()

			for _, i := range addresses {

				var ip net.IP
				switch v := i.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}

				return &NetworkInfo{
					IP:  ip,
					MAC: v.HardwareAddr.String(),
				}
			}
		}
	}

	return nil
}
