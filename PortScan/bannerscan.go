package PortScan

import (
	"Goscan/Fingerprint"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type webbanner struct {
	header http.Header
	body string
}

func newWebbanner(header http.Header, body string) *webbanner {
	return &webbanner{header: header, body: body}
}


func GetBanner(ip string, port string) (*webbanner,bool) {
	var buff []byte
	var tmp = true
	tcpAddr := ip + ":" + port
	conn, err := net.DialTimeout("tcp",tcpAddr,time.Second*3)
	if err != nil{
		tmp = false
		return newWebbanner(nil,""),tmp
	}
	defer conn.Close()
	fmt.Fprintf(conn,"\r\n\r\n")
	conn.SetReadDeadline(time.Now().Add(time.Second*2))
	num,_:= conn.Read(buff)
	content := buff[:num]
	banner := newWebbanner(nil,string(content))
	return banner,tmp
}

func GetWebBanner(ip string, port string) (*webbanner,bool) {
	url := "http://"+ ip + ":" + port
	client := http.Client{Timeout: 3 * time.Second}
	resp,err := client.Get(url)
	if err != nil {
		return nil,false
	}
	body, _ := ioutil.ReadAll(resp.Body)
	header := resp.Header
	webbanner := newWebbanner(header,string(body))
	return webbanner,true
}

func BannerCheck(ip string, port string) (*webbanner,bool) {
	banner,flag := GetBanner(ip,port)
	if banner.header == nil && flag==true{
		banner,flag = GetWebBanner(ip,port)
	}
	return banner,flag
}

func Bannerscan(ip string, port string){
	if strings.Contains(port,","){
		ports := strings.Split(port,",")
		var wg sync.WaitGroup
		for _,p :=range ports{
			wg.Add(1)
			go func(ip string,p string) {
				banner,flag := BannerCheck(ip,p)
				if flag {
					OS := Fingerprint.OSCheck(banner.header)
					Server := Fingerprint.ServerCheck(banner.header,banner.body)
					fmt.Printf("PORT\tOS\tSERVER\n")
					fmt.Printf("%s\t%s\t%s\n",p,OS,Server)
				}
				defer wg.Done()
			}(ip,p)
		}
		wg.Wait()
	}
	if strings.Contains(port, "-") {
		tmps := strings.Split(port,"-")
		MIN,_ := strconv.Atoi(tmps[0])
		MAX,_ := strconv.Atoi(tmps[1])
		MAX = MAX + 1
		var wg sync.WaitGroup
		for p:=MIN;p<MAX;p++{
			wg.Add(1)
			go func(ip string,p string) {
				banner,flag := BannerCheck(ip,p)
				if flag {
					OS := Fingerprint.OSCheck(banner.header)
					Server := Fingerprint.ServerCheck(banner.header,banner.body)
					fmt.Printf("PORT\tOS\tSERVER\n")
					fmt.Printf("%s\t%s\t%s\n",p,OS,Server)
				}
				defer wg.Done()
			}(ip,strconv.Itoa(p))
		}
		wg.Wait()
	}else {
		banner,flag := BannerCheck(ip,port)
		if flag {
			OS := Fingerprint.OSCheck(banner.header)
			Server := Fingerprint.ServerCheck(banner.header,banner.body)
			fmt.Printf("PORT\tOS\tSERVER\n")
			fmt.Printf("%s\t%s\t%s\n",port,OS,Server)
		}
	}
}