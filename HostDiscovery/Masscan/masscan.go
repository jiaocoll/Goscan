package Masscan


import (
	"github.com/dean2021/go-masscan"
	"log"
)



func Mscan(ips string,ports string) []string{
	var target []string

	m := masscan.New()

	// masscan可执行文件路径,默认不需要设置
	m.SetSystemPath(`D:\Tools\masscan-1.3.2\masscan.exe`)

	// 扫描端口范围
	m.SetPorts(ports)

	// 扫描IP范围
	m.SetRanges(ips)

	// 扫描速率
	m.SetRate("2000")

	// 开始扫描
	err := m.Run()
	if err != nil {
		log.Println("scanner failed:", err)
	}

	// 解析扫描结果
	results, err := m.Parse()
	if err != nil {
		log.Println("Parse scanner result:", err)
	}

	for _, result := range results {
		for _,port := range result.Ports{
			tmp := result.Address.Addr + ":" + port.Portid
			target = append(target, tmp)
		}
	}
	//返回目标列表，单个目标由ip和端口组成
	return target
}