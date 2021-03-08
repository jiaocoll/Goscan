# Goscan
简易的Go主机发现和端口扫描
目前是demo，有待完善
```
Pandora.exe [选项] [ip]
Options:
  -h, --help
        help, 帮助命令
  -p string
        PortScan, 输入目标ip进行端口扫描,仅判断是否开启,速度较快,例如:-p 22,80或22-8080
  -pv string
        输入目标ip进行端口扫描,判断是否开启并获得banner,速度较慢,例如:-pv 22,80或22-8080
  -sP string
        PingScan, 输入目标ip进行Ping扫描,多个ip使用逗号分隔,支持C段扫描,例如:-sP 127.0.0.1,192.168.10.1或192.168.10.1/24
  -sT string
        输入目标ip进行TCP连接扫描,例如:-sT 192.168.10.1,192,168,10,45
  -ds string
        输入目标url进行目录扫描,例如:-ds https://www.example.com
```
