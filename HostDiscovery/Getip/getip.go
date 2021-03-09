package main

import (
	"fmt"
	"github.com/fatih/color"
	"net"
	"sync"
	"time"
)

func Getipbyhost(hostname string) string {
	ip, err := net.ResolveIPAddr("ip",hostname)
	if err != nil {
		fmt.Fprintln(color.Output,time.Now().Format("2006/01/02 15:04:05"),color.YellowString("[WARNING]")+":",err)
	}
	return ip.String()
}

func Getip(hostnames []string) []string {
	var ips []string
	var wg sync.WaitGroup
	for _,hostname := range hostnames{
		wg.Add(1)
		go func(hostname string) {
			defer wg.Done()
			ip := Getipbyhost(hostname)
			ips = append(ips, ip)
		}(hostname)
	}
	wg.Wait()
	return ips
}
