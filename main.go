package main

import (
	"net"
	"os"
)

func main() {
	nInfo, err := getNetworkInformations()

	if err != nil {
		panic(err)
	}

	tv := &TVInfo{
		Model: os.Args[1],
		IP:    net.ParseIP(os.Args[2]),
	}

	controller := NewSamsungController()
	err = controller.Connect(nInfo, tv)

	if err != nil {
		panic(err)
	}

	defer controller.Close()

	err = controller.SendKey(nInfo, tv, os.Args[3])

	if err != nil {
		panic(err)
	}
}
