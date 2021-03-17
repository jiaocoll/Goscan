# Goscan
Go主机发现和端口扫描
```
Goscan.exe [选项] [ip]
Options:
  -dict string
        字典文件
  -ds string
        输入目标url进行目录扫描,例如:-ds https://www.example.com
  -h, --help
        help, 帮助命令
  -ip string
        masscan-nmap扫描 ip地址范围,有三种有效格式,1:单独的IPv4地址 2:类似10.0.0.1-10.0.0.233的范围地址 3:CIDR地址,类似 于0.0.0.0/0,多个目标可以用逗号隔开
  -port string
        masscan-nmap扫描 port端口范围 例如-port 1-100或80,22,24
  -pv string
        输入目标ip进行端口扫描,判断是否开启并获得banner,速度较慢,例如:-pv 22,80或22-8080
  -sP string
        PingScan, 输入目标ip进行Ping扫描,多个ip使用逗号分隔,支持C段扫描,例如:-sP 127.0.0.1,192.168.10.1或192.168.10.1/24
  -sT string
        输入目标ip进行TCP连接扫描,例如:-sT 192.168.10.1,192,168,10,45
  -tP string
        PortScan, 输入目标ip进行端口扫描,仅判断是否开启,速度较快,例如:-tP 22,80或22-8080
```
