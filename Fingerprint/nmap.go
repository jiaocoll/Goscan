package Fingerprint

import (
	"fmt"
	"github.com/Ullaakut/nmap"
	"log"
	"strings"
)

func Nmapscan(target string) {

	tmptarget := strings.Split(target,":")
	ip := tmptarget[0]
	port := tmptarget[1]


	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5 minute timeout.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts(port),
		nmap.WithServiceInfo(),
		nmap.WithBinaryPath(`D:\Tools\nmap\nmap.exe`),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name,port.Service.Version)
		}
	}

}
