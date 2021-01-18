package PortScan

import (
	"Pandora/HostDiscovery/Ping"
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
				fmt.Println("[*]"+ip+":")
				fmt.Println("PORT\tSTATUS")
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
		fmt.Println("PORT\tSTATUS")
		for _,port := range aliveport{
			fmt.Println(port+"\tOpen")
		}
		fmt.Println("开放端口个数为:",len(aliveport))
	}
}

func BannerCheck(ip string, ports string)(Ip string,Ports []string,Banners []string){
	var banner string
	portresult := PortScanOne(ip,ports)
	for _, port := range portresult{
		tcpAddr := ip + ":" + port
		conn, err := net.DialTimeout("tcp",tcpAddr,time.Second*2)
		if err == nil {
			fmt.Fprintf(conn,"\r\n\r\n")
			conn.SetReadDeadline(time.Now().Add(time.Second*2))
			buff := make([]byte,1024)
			num,_:= conn.Read(buff)
			banner = string(buff[:num])
			if banner == "" {
				fmt.Fprintf(conn,"GET / HTTP/1.1\r\n\r\n")
				conn.SetReadDeadline(time.Now().Add(time.Second*2))
				buff := make([]byte,1024)
				num,_:= conn.Read(buff)
				banner = string(buff[:num])
			}
			banner = strings.Replace(banner,"\r\n","",-1)
			if strings.Contains(banner, "HTTP") {
				Banners = append(Banners,"http")
			}else if strings.Contains(banner,"SSH") {
				Banners = append(Banners,"ssh")
			}else if strings.Contains(banner,"FTP") {
				Banners = append(Banners,"ftp")
			}else if strings.Contains(banner,"TELNET") {
				Banners = append(Banners,"telnet")
			}else {
				Banners = append(Banners,"")
			}
		}
	}
	return ip,portresult,Banners

}

func Bannerscan(ip string, ports string){
	r1,r2,r3 := BannerCheck(ip,ports)
	fmt.Println(r1,":")
	fmt.Println("PORT\tSTATUS\tSERVER")
	for i:=0;i<len(r2);i++{
		fmt.Println(r2[i]+"\tOpen\t"+r3[i])
	}
}
