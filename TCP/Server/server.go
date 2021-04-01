package main

import (
	"fmt"
	"log"
	"net"
	"svcMESH_test/utils"
)

const IP = "localhost"
const PORT = 8080

func main() {
	ip := net.ParseIP(IP)
	laddr := &net.TCPAddr{
		IP: ip,
		Port: PORT,
	}
	ln, err := net.ListenTCP("tcp", laddr)
	if err != nil{
		log.Fatalf("listen tcp err! %s", err)
		return
	}

	for {
		tcpCon, err := ln.AcceptTCP()
		if err != nil {
			log.Println("tcp connect err!")
			continue
		}
		log.Println("accept a new connect")
		go handlerConn(tcpCon)
	}
}

func handlerConn(con *net.TCPConn){
	defer con.Close()

	var acceptStr = ""
	var ip = utils.GetIPInfo()
	var name = utils.GetHostName()
	ms := fmt.Sprintf("I am a %s ip%+v", name, ip)

	for {
		var buf = make([]byte, 1024)
		log.Println("start to read from conn")
		n, err := con.Read(buf)
		if err != nil{
			log.Println("conn read error", err)
			return
		}
		acceptStr = string(buf[:n])
		log.Printf("read %d bytes, content is %s\n", n, acceptStr)
		con.Write([]byte(ms))
	}
}