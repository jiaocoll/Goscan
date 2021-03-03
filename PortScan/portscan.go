package PortScan

import (
	"github.com/fatih/color"
	"Goscan/HostDiscovery/Ping"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)


func PortCheck(ip string, port string)(bool){
	ip = ip + ":" + port
	conn, err := net.DialTimeout("tcp",ip,time.Second*3)
	if err != nil{
		return false
	}else {
		conn.Close()
		return true
	}
}

func PortScanOne(ip string, ports string)(result []string){
	var aliveport []string
	if strings.Contains(ports,","){
		Ports := strings.Split(ports,",")
		var wg sync.WaitGroup
		for _, port := range Ports{
			wg.Add(1)
			go func(ip string, port string) {
				defer wg.Done()
				if (PortCheck(ip,port)){
					aliveport = append(aliveport,port)
				}
			}(ip,port)
		}
		wg.Wait()
		result = aliveport
		return result
	}else if strings.Contains(ports,"-") {
		Ports := strings.Split(ports,"-")
		var wg sync.WaitGroup
		MIN,_ := strconv.Atoi(Ports[0])
		MAX,_ := strconv.Atoi(Ports[1])
		MAX = MAX + 1
		for port:=MIN;port<MAX;port++{
			wg.Add(1)
			go func(ip string, port string) {
				defer wg.Done()
				if (PortCheck(ip,port)){
					aliveport = append(aliveport,port)
				}
			}(ip, strconv.Itoa(port))
		}
		wg.Wait()
		result = aliveport
		return result
	}
	result = aliveport
	return result
}

func PortScanMore(ips []string, ports string)(result []string){
	var wg sync.WaitGroup
	var aliveport []string
	for _, ip := range ips{
		wg.Add(1)
		go func(ip string, ports string) {
			defer wg.Done()
			aliveport = PortScanOne(ip,ports)
			if aliveport != nil{
				fmt.Println(color.HiMagentaString("[+]")+ip+":")
				fmt.Println(color.HiCyanString("PORT\tSTATUS"))
				for _,port := range aliveport{
					fmt.Println(port+"\tOpen")
				}
			}
		}(ip,ports)
	}
	wg.Wait()
	result = aliveport
	return result
}

func PortScan(ip string, ports string){
	if strings.Contains(ip,"/24") {
		ips := Ping.CScan(ip)
		PortScanMore(ips,ports)
	}else {
		aliveport := PortScanOne(ip,ports)
		fmt.Println(ip+":")
		fmt.Println(color.HiCyanString("PORT\tSTATUS"))
		for _,port := range aliveport{
			fmt.Println(port+"\tOpen")
		}
		fmt.Println(color.CyanString("开放端口个数为:"),len(aliveport))
	}
}


