package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoringNumber = 3
const monitoringDelay = 5

func main() {

	introduction()

	for {
		menuOptions()

		command := commandInput()

		switch command {
		case 1:
			startMonitor()
		case 2:
			showLog()
		case 0:
			fmt.Println("Finished.")
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
			os.Exit(-1)
		}
	}
}

func introduction() {
	fmt.Println("")
	name := "Luan Sapelli"
	version := 1.1
	fmt.Println("Hello", name, "| Type of name:", reflect.TypeOf(name))
	fmt.Println("Program Version:", version, "| Type of version:", reflect.TypeOf(version))
	fmt.Println("")
}

func menuOptions() {
	fmt.Println("1 - Start Monitor")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
	fmt.Println("")
}

func commandInput() int {
	var commandInput int
	fmt.Scan(&commandInput)
	fmt.Println("")
	return commandInput
}

func startMonitor() {
	fmt.Println("Monitoring...")

	sites := siteArchive()

	for i := 0; i < monitoringNumber; i++ {
		for i, site := range sites {
			fmt.Println("Monitoring site", i, site)
			siteStatus(site)
		}
		time.Sleep(monitoringDelay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func siteStatus(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "is online. Status code: ", resp.StatusCode)
		writeLog(site, true)
	} else {
		fmt.Println("Site:", site, "is offline. Status Code: ", resp.StatusCode)
		writeLog(site, false)
	}
}

func siteArchive() []string {
	var sites []string
	archive, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	reader := bufio.NewReader(archive)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}
	archive.Close()

	return sites
}

func writeLog(site string, status bool) {
	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	archive.WriteString(time.Now().Format("Mon 02/01/2006 15:04:05") + " - " + site + " - online:" + strconv.FormatBool(status) + "\n")
	//https://golang.org/src/time/format.go

	archive.Close()
}

func showLog() {
	archive, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(archive))
}
