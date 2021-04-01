package utils

import (
	"fmt"
	"log"
	"net"
	"os"
)

func GetIPInfo() []string {
	addrs, err := net.InterfaceAddrs()
	strs := make([]string, 0)
	if err != nil {
		fmt.Println(err)
		return strs
	}

	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				strs = append(strs, ipnet.IP.String())
			}
		}
	}
	return strs
}

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("Get HostName err: ", err)
		return "No-Hostname"
	}
	return hostname
}