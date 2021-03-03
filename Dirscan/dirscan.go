package Dirscan

import (
	"bufio"
	"github.com/fatih/color"
	"log"
	"net/http"
	"os"
)


var (
 	dirs []string
)



func dircheck(url string) bool{
	resp,err := http.Get(url)
	if nil != resp {
		defer resp.Body.Close()
	}
	if err != nil {
		return false
	}
	if resp.StatusCode == 200 || resp.StatusCode == 403 {
		return true
	}
	return false
}

func Dirscan(url string,dirfile string){

	infile, err := os.OpenFile(dirfile,os.O_RDONLY,1)
	if err != nil {
		log.Println(color.RedString("[ERROR]")+":",err)
	}
	tmps := bufio.NewScanner(infile)
	for tmps.Scan() {
		tmp := tmps.Text()
		dirs = append(dirs, tmp)

	}
	for _,dir := range dirs{
		target := url + dir
		if dircheck(target) {
			log.Println(color.CyanString(target))
		}
	}

}
