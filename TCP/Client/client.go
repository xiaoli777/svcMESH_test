package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	address := flag.String("address", "localhost:8080", "tcp destination")
	flag.Parse()

	con, err := net.Dial("tcp", *address)
	if err != nil {
		log.Println("connect server err:", err)
		return
	}
	defer con.Close()

	done := make(chan struct{}, 0)
	go func() {
		defer close(done)
		for {
			var buf = make([]byte, 1024)
			_, err = con.Read(buf)
			if err != nil{
				log.Println("Read Data err!", err)
				return
			}
			fmt.Println("recv: ", fmt.Sprintf("%s %s", string(buf), time.Now().String()))
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
			case <- done:
			case t:= <- ticker.C:
				msg := fmt.Sprint("%s %s", "hello! I am a client!", t.String())
				con.Write([]byte(msg))
		}
	}
}