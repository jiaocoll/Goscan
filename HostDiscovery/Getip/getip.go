package main

import (
	"github.com/fatih/color"
	"log"
	"net"
	"sync"
)

func Getipbyhost(hostname string) string {
	ip, err := net.ResolveIPAddr("ip",hostname)
	if err != nil {
		log.Println(color.YellowString("[WARNING]")+":",err)
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
