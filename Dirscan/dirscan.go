package Dirscan

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"os"
	"time"
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
		fmt.Fprintln(color.Output,time.Now().Format("2006-01-02 15:04:05"),color.RedString("[ERROR]")+":",err)
	}
	tmps := bufio.NewScanner(infile)
	for tmps.Scan() {
		tmp := tmps.Text()
		dirs = append(dirs, tmp)

	}
	for _,dir := range dirs{
		target := url + dir
		if dircheck(target) {
			fmt.Fprintln(color.Output,time.Now().Format("2006/01/02 15:04:05"),color.HiCyanString(target))
		}
	}

}
