package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/miekg/dns"
)

var maxConnection = 100
var iterCount = 10000000
var PingURL = "http://nginx-svc.default.svc.cluster.local:8080/ping"
var QueryURL = "http://nginx-svc.default.svc.cluster.local"

func main() {
	maxConnection = *flag.Int("n", 100, "-n=100")
	wg := sync.WaitGroup{}
	ch := make(chan int, maxConnection)

	for i := 0; i< iterCount; i++ {
		ch <- 1
		wg.Add(1)
		go func(index int64, cc chan int) {
			defer wg.Done()
			start := time.Now()
			PingTest(index)
			defer flushChan(cc, start)
		}(int64(i), ch)
	}
	wg.Wait()
}

func flushChan(ch chan int, t time.Time) {
	<-ch
	elapsed := time.Since(t)
	kk := float64(1) / elapsed.Seconds()
	fmt.Println("======= maxConnection", maxConnection, "=======", time.Now(), " qps ==", kk)
}

func dnsAQuery(i int64){
	client := dns.Client{}
	msg := dns.Msg{}
	msg.SetQuestion(QueryURL, dns.TypeA)

	ns := "172.17.0.1" + ":53"
	res, _, err := client.Exchange(&msg, ns)
	if err!= nil {
		fmt.Println("nameserver %s error: %v", ns,err)
		return
	}

	if len(res.Answer) > 0 {

	} else {
		fmt.Println("===", i, "error")
	}
}

func PingTest(i int64){
	url := PingURL
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed i = ", i, " time = ", time.Now(), "Get error =>", err)
		fmt.Println("Failed i = ", i, " time = ", time.Now(), "Get error =>", err)
		return
	}
	if resp == nil {
		log.Println("Failed i = ", i, " time = ", time.Now(), "======= error =======")
		fmt.Println("Failed i = ", i, " time = ", time.Now(), "======= error =======")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed i = ", i, " time = ", time.Now(), "ReadAll error =>", err)
		fmt.Println("Failed i = ", i, " time = ", time.Now(), "ReadAll error =>", err)
		return
	}
	fmt.Println("Successfully! i = ", i, " time = ", time.Now(), " ", string(body), " ==>", i)
}