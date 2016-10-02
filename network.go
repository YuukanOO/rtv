package main

import (
	"net"
	"strings"
)

func getNetworkInformations() (*NetworkInfo, error) {
	interfaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	for _, v := range interfaces {

		if v.Flags&net.FlagLoopback == 0 && !strings.Contains(v.Name, "vir") {
			addresses, err := v.Addrs()

			if err != nil {
				return nil, err
			}

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
				}, nil
			}
		}
	}

	return nil, nil
}
