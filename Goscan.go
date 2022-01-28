package main

import (
	"Goscan/Dirscan"
	"Goscan/Fingerprint"
	"Goscan/HostDiscovery/Masscan"
	"Goscan/HostDiscovery/Tcpscan"
	"github.com/fatih/color"
	"Goscan/HostDiscovery/Ping"
	"Goscan/PortScan"
	"flag"
	"fmt"
	"time"
)

var (
	help bool
	ip string
	pingscan string
	portscan string
	portbanner string
	dirscan string
	TCPscan string
	dict string
	masnmapip string
	masnmapport string
)

func init(){
	flag.BoolVar(&help,"h, --help",false,"help, 帮助命令")
	flag.StringVar(&dict,"dict","","字典文件")
	flag.StringVar(&pingscan,"sP","","PingScan, 输入目标ip进行Ping扫描,多个ip使用逗号分隔,支持C段扫描,例如:-sP 127.0.0.1,192.168.10.1或192.168.10.1/24")
	flag.StringVar(&portscan,"tP","","PortScan, 输入目标ip进行端口扫描,仅判断是否开启,速度较快,例如:-tP 22,80或22-8080")
	flag.StringVar(&portbanner,"pv","","输入目标ip进行端口扫描,判断是否开启并获得banner,速度较慢,例如:-pv 22,80或22-8080")
	flag.StringVar(&dirscan,"ds","","输入目标url进行目录扫描,例如:-ds https://www.example.com")
	flag.StringVar(&TCPscan,"sT","","输入目标ip进行TCP连接扫描,例如:-sT 192.168.10.1,192,168,10,45")
	flag.StringVar(&masnmapip,"ip","","masscan-nmap扫描 ip地址范围,有三种有效格式,1:单独的IPv4地址 2:类似10.0.0.1-10.0.0.233的范围地址 3:CIDR地址,类似于0.0.0.0/0,多个目标可以用逗号隔开")
	flag.StringVar(&masnmapport,"port","","masscan-nmap扫描 port端口范围 例如-port 1-100或80,22,24")
	flag.Usage = usage
	flag.Parse()
	ip = flag.Arg(0)
}

func usage(){
	fmt.Fprintf(color.Output,color.CyanString(`Ameng编写的Go语言主机发现和端口扫描工具——Goscan
Options:
`))
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
		fmt.Fprint(color.Output,color.HiMagentaString("花费时间为:"),end)
	}
	if  portscan != ""{
		start := time.Now()
		PortScan.PortScan(ip,portscan)
		end := time.Since(start)
		fmt.Fprint(color.Output,color.HiMagentaString("花费时间为:"),end)
	}
	if portbanner != ""{
		start := time.Now()
		PortScan.Bannerscan(ip,portbanner)
		end := time.Since(start)
		fmt.Fprint(color.Output,color.HiMagentaString("花费时间为:"),end)
	}
	if dirscan != ""{
		start := time.Now()
		Dirscan.Dirscan(dirscan,dict)
		end := time.Since(start)
		fmt.Fprint(color.Output,color.HiMagentaString("花费时间为:"),end)
	}
	if TCPscan != ""{
		start := time.Now()
		Tcpscan.Tcpscan(TCPscan)
		end := time.Since(start)
		fmt.Fprint(color.Output,color.HiMagentaString("花费时间为:"),end)
	}

	if masnmapip != "" && masnmapport != ""{
		targets := Masscan.Mscan(masnmapip,masnmapport)
		for _,target := range targets{
			Fingerprint.Nmapscan(target)
		}
	}



}
