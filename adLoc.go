package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

var env = runtime.GOOS

func CheckSite(site string, url string) {
	finalUrl := url + "/" + site
	req, err := http.Get(finalUrl)
	if err != nil {
		log.Fatal(err)
	}
	if req.StatusCode == 200 {
		color.Green("/" + site + " --> " + "Admin login panel found!\n")
		color.Green("Opening in browser ...\n")
		//open in browser
		switch env {
		case "linux":
			_ = exec.Command("/usr/bin/xdg-open", finalUrl).Start()
		case "windows":
			_ = exec.Command("C:\\Windows\\System32\\rundll32", "url.dll,FileProtocolHandler", finalUrl).Start()
		}

	} else {
		color.Red("/" + site + " --> " + "Admin login panel not present! ")
	}

}
func main() {

	banner := `    
	_____________________              
	___    |_____  /__  / _____________
	__  /| |  __  /__  /  _  __ \  ___/
	_  ___ / /_/ / _  /___/ /_/ / /__  
	/_/  |_\__,_/  /_____/\____/\___/  

	Coded by 6en6ar :)
																	 
	`
	url := flag.String("u", "", "Enter an url you wish to scan")
	flag.Parse()
	if *url == "" {
		color.Red("Please enter an url with -u")
		os.Exit(1)
	}
	color.Green(banner)
	file, err := os.Open("admin.txt")
	if err != nil {
		color.Red("Error opening file!")
	}
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	color.Green("Checking admin panels on --> " + *url)
	for scan.Scan() {
		CheckSite(scan.Text(), *url)
	}

}
