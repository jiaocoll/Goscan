package Fingerprint

import (
	"net/http"
	"strings"
)

func OSCheck(banner http.Header) string{
	server := banner.Get("server")
	x_powered_by := banner.Get("x_powered_by")
	if strings.Contains(server,"win") || strings.Contains(x_powered_by,"win") {
		return "Windows Server"
	}
	if strings.Contains(server,"CentOS") || strings.Contains(x_powered_by,"CentOS") {
		return "CentOS"
	}
	if strings.Contains(server,"Ubuntu") || strings.Contains(x_powered_by,"Ubuntu") {
		return "Ubuntu"
	}
	if strings.Contains(server,"Darwin") || strings.Contains(x_powered_by,"Darwin") {
		return "Darwin"
	}
	if strings.Contains(server,"Debian") || strings.Contains(x_powered_by,"Debian") || strings.Contains(x_powered_by,"dotdeb") || strings.Contains(x_powered_by,"sarge\\;version:\\1") || strings.Contains(x_powered_by,"etch\\;version:\\1") || strings.Contains(x_powered_by,"lenny\\;version:\\1") || strings.Contains(x_powered_by,"squeeze\\;version:\\1") || strings.Contains(x_powered_by,"wheezy\\;version:\\1") || strings.Contains(x_powered_by,"jessie\\;version:\\1"){
		return "Debian"
	}
	if strings.Contains(server,"Fedora") {
		return "Fedora"
	}
	if strings.Contains(server,"FreeBSD") {
		return "FreeBSD"
	}
	if strings.Contains(x_powered_by,"gentoo") {
		return "Gentoo"
	}
	if strings.Contains(server,"Unix") {
		return "UNIX"
	}
	if strings.Contains(server,"Red Hat") || strings.Contains(x_powered_by,"Red Hat") {
		return "Red Hat"
	}
	if strings.Contains(server,"SUSE") || strings.Contains(x_powered_by,"SUSE") {
		return "SUSE"
	}
	if strings.Contains(server,"Scientific Linux") || strings.Contains(x_powered_by,"Scientific Linux") {
		return "Scientific Linux"
	}
	if strings.Contains(server,"SunOS") || strings.Contains(x_powered_by,"SunOS") {
		return "SunOS"
	}
	if strings.Contains(server,"BWS") {
		return "BWS"
	}
	return "UNKNOW"
}

