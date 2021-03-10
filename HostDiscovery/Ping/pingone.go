package Ping

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/panjf2000/ants/v2"
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
	var aliveip []string
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(50000, func(i interface{}) {
		if(PingOne(i.(string))){
			aliveip = append(aliveip, i.(string))
		}
		wg.Done()
	})
	for i:=1;i<255;i++{
		ip := ip + "." + strconv.Itoa(i)
		wg.Add(1)
		_ = p.Invoke(ip)
	}
	wg.Wait()
	result = aliveip
	return result
}


func PingScan(target string){
	ips := strings.Split(target,",")
	for _,ip := range ips{
		if strings.Contains(ip,"/24") {
			ips := CScan(ip)
			fmt.Fprintln(color.Output,color.HiCyanString("总存活主机个数为:"),len(ips))
			for _,ip := range ips{
				fmt.Fprintln(color.Output,color.HiBlueString("[+]存活的主机:"),ip)
			}
		}else {
			if (PingOne(ip)){
				fmt.Fprintln(color.Output,color.HiBlueString("[+]存活的主机:"),ip)
			}
		}
	}

}


