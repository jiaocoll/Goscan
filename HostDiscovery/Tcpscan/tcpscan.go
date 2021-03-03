package Tcpscan

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	topport = []int{1,3,4,6,7,9,13,17,19,20,21,22,23,24,25,26,30,32,33,37,42,43,49,53,70,
		79,80,81,82,83,84,85,88,89,90,99,100,106,109,110,111,113,119,125,135,
		139,143,144,146,161,163,179,199,211,212,222,254,255,256,259,264,280,
		301,306,311,340,366,389,406,407,416,417,425,427,443,444,445,458,464,
		465,481,497,500,512,513,514,515,524,541,543,544,545,548,554,555,563,
		587,593,616,617,625,631,636,646,648,666,667,668,683,687,691,700,705,
		711,714,720,722,726,749,765,777,783,787,800,801,808,843,873,880,888,
		898,900,901,902,903,911,912,981,987,990,992,993,995,999,1000,1001,
		1002,1007,1009,1010,1011,1021,1022,1023,1024,1025,1026,1027,1028,
		1029,1030,1031,1032,1033,1034,1035,1036,1037,1038,1039,1040,1041,
		1042,1043,1044,1045,1046,1047,1048,1049,1050,1051,1052,1053,1054,
		1055,1056,1057,1058,1059,1060,1061,1062,1063,1064,1065,1066,1067,
		1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,
		1081,1082,1083,1084,1085,1086,1087,1088,1089,1090,1091,1092,1093,
		1094,1095,1096,1097,1098,1099,1100,1102,1104,1105,1106,1107,1108,
		}
)

func Portcheck(ip string, port string) bool{
	ip = ip + ":" + port
	conn, err := net.DialTimeout("tcp",ip,time.Second*3)
	if err != nil{
		return false
	}else {
		conn.Close()
		return true
	}
}

func TcpscanOne(ip string) bool{
	var result bool
	var wg sync.WaitGroup
	for _,port := range topport {
		wg.Add(1)
		go func(port string) {
			if (Portcheck(ip,port)){
				result = true
			}
			wg.Done()
		}(strconv.Itoa(port))
	}
	wg.Wait()
	return result
}

func Tcpscan(ip string) {
	ips := strings.Split(ip,",")
	var aliveip []string
	var wg sync.WaitGroup
	for _,ip := range ips{
		wg.Add(1)
		go func(ip string) {
			if (TcpscanOne(ip)){
				aliveip = append(aliveip, ip)
			}
			wg.Done()
		}(ip)
	}
	wg.Wait()
	fmt.Println(aliveip)
}