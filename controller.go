package main

import "net"

// TVInfo represents a remote TV.
type TVInfo struct {
	Model string
	IP    net.IP
}

// NetworkInfo represents a device on the network.
type NetworkInfo struct {
	IP  net.IP
	MAC string
}

// Controller is the base interface implemented by vendor specific TVs.
type Controller interface {
	Connect(emitter *NetworkInfo, receiver *TVInfo) error
	SendKey(emitter *NetworkInfo, receiver *TVInfo, key string) error
	Close() error
}
