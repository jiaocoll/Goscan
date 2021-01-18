package main

import (
	"Pandora/HostDiscovery/Ping"
	"Pandora/PortScan"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	help bool
	ip string
	pingscan string
	portscan string
	portbanner string

)

func init(){
	flag.BoolVar(&help,"h, --help",false,"help, 帮助命令")
	flag.StringVar(&pingscan,"sP","","PingScan, 输入目标ip进行Ping扫描,多个ip使用逗号分隔,支持C段扫描,例如:-sP 127.0.0.1,192.168.10.1或192.168.10.1/24")
	flag.StringVar(&portscan,"p","","PortScan, 输入目标ip进行端口扫描,仅判断是否开启,速度较快,例如:-p 22,80或22-8080")
	flag.StringVar(&portbanner,"pv","","输入目标ip进行端口扫描,判断是否开启并获得banner,速度较慢,例如:-pv 22,80或22-8080")
	flag.Usage = usage
	flag.Parse()
	ip = flag.Arg(0)
}

func usage(){
	fmt.Fprintf(os.Stderr,`Ameng编写的Go语言综合渗透工具——Pandora
Options:
`)
	flag.PrintDefaults()
}

func main(){


	if help {
		flag.Usage()
	}
	if pingscan != "" {
		start := time.Now()
		Ping.PingScan(pingscan)
		end := time.Since(start)
		fmt.Print("花费时间为:",end)
	}
	if  portscan != ""{
		start := time.Now()
		PortScan.PortScan(ip,portscan)
		end := time.Since(start)
		fmt.Print("花费时间为:",end)
	}
	if portbanner != ""{
		start := time.Now()
		PortScan.Bannerscan(ip,portbanner)
		end := time.Since(start)
		fmt.Print("花费时间为:",end)
	}




}