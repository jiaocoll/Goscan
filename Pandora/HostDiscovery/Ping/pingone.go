package Ping

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func PingOne(ip string) bool {
	sysType := runtime.GOOS
	if sysType == "windows"{
		cmd := exec.Command("ping", "-n", "2", ip)
		var output bytes.Buffer
		cmd.Stdout = &output
		cmd.Run()
		if strings.Contains(output.String(), "TTL=") && strings.Contains(output.String(), ip)  {
			return true
		}
	}else if sysType == "linux" {
		cmd := exec.Command("ping", "-c", "2", ip)
		var output bytes.Buffer
		cmd.Stdout = &output
		cmd.Run()
		if strings.Contains(output.String(), "ttl=") && strings.Contains(output.String(), ip) {
		return true
		}
		}
	return false
}

func CScan(target string) (result []string) {
	ip := strings.Replace(target, "/24", "", -1)
	ips := strings.Split(ip,".")
	ip = ips[0] + "." + ips[1] + "." + ips[2]
	var aliveip = []string{}
	var wg sync.WaitGroup
	for i:=1;i<255;i++{
		ip := ip + "." + strconv.Itoa(i)
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			if (PingOne(ip)){
				aliveip = append(aliveip, ip)
			}
		}(ip)
	}
	wg.Wait()
	result = aliveip
	return result
}

func PingScan(ip string){
	if strings.Contains(ip,"/24") {
		ips := CScan(ip)
		fmt.Println("总存活主机个数为:",len(ips))
		for _,ip := range ips{
			fmt.Println("[+]存活的主机:",ip)
		}
	}else {
		if (PingOne(ip)){
			fmt.Println("[+]存活的主机:",ip)
		}
	}
}


