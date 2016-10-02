package main

import (
	"encoding/base64"
	"fmt"
	"net"
)

// SamsungController represents a controller for samsung smart tvs.
type SamsungController struct {
	appString  string
	remoteName string
	handle     *net.TCPConn
}

// NewSamsungController instantiates a new controller for samsung smart TVs.
func NewSamsungController() *SamsungController {
	return &SamsungController{
		appString:  "iphone..iapp.samsung",
		remoteName: "RTV Remote Control",
	}
}

// Connect initialize the connection.
func (controller *SamsungController) Connect(emitter *NetworkInfo, receiver *TVInfo) error {
	conn, err := net.DialTCP("tcp", &net.TCPAddr{
		IP: emitter.IP,
	}, &net.TCPAddr{
		IP:   receiver.IP,
		Port: 55000,
	})

	if err != nil {
		return err
	}

	controller.handle = conn

	encoding := base64.StdEncoding

	encodedIP := encoding.EncodeToString([]byte(emitter.IP.String()))
	encodedMAC := encoding.EncodeToString([]byte(emitter.MAC))
	encodedRemoteName := encoding.EncodeToString([]byte(controller.remoteName))

	msgPart1 := fmt.Sprintf("%c%c%c%c%s%c%c%s%c%c%s", 0x64, 0x00, len(encodedIP), 0x00, encodedIP, len(encodedMAC), 0x00, encodedMAC, len(encodedRemoteName), 0x00, encodedRemoteName)
	part1 := fmt.Sprintf("%c%c%c%s%c%c%s", 0x00, len(controller.appString), 0x00, controller.appString, len(msgPart1), 0x00, msgPart1)

	_, err = controller.handle.Write([]byte(part1))

	if err != nil {
		return err
	}

	msgPart2 := fmt.Sprintf("%c%c", 0xc8, 0x00)
	part2 := fmt.Sprintf("%c%c%c%s%c%c%s", 0x00, len(controller.appString), 0x00, controller.appString, len(msgPart2), 0x00, msgPart2)

	_, err = controller.handle.Write([]byte(part2))

	return err
}

// SendKey sends a key to the TV.
func (controller *SamsungController) SendKey(emitter *NetworkInfo, receiver *TVInfo, key string) error {
	encoding := base64.StdEncoding
	encodedKey := encoding.EncodeToString([]byte(key))

	msgPart3 := fmt.Sprintf("%c%c%c%c%c%s", 0x00, 0x00, 0x00, len(encodedKey), 0x00, encodedKey)
	part3 := fmt.Sprintf("%c%c%c%s%c%c%s", 0x00, len(controller.appString), 0x00, controller.appString, len(msgPart3), 0x00, msgPart3)

	_, err := controller.handle.Write([]byte(part3))

	return err
}

// Close the connection.
func (controller *SamsungController) Close() error {
	return controller.handle.Close()
}
