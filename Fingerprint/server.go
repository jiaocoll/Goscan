package Fingerprint

import (
	"net/http"
	"strings"
)

func ServerCheck(header http.Header,banner string)string{
	if header != nil{
		return "HTTP"
	}
	if strings.Contains(banner,"ssh") {
		return "SSH"
	}
	if strings.Contains(banner,"ftp") {
		return "FTP"
	}
	if strings.Contains(banner,"smb") {
		return "SMB"
	}else {
		return "UNKOWN"
	}
}
