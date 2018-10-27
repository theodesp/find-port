package findport

import (
	"net"
)

func DetectOpenPort() (int, error) {
	openPort, err := tryListen()
	if err != nil {
		return -1, err
	} else {
		return openPort, nil
	}
}

func tryListen() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return -1, err
	}
	return listener.Addr().(*net.TCPAddr).Port, nil
}
