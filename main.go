package main

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
)

const (
	File  = "file.txt"
	Api   = "https://discordapp.com/api/invite/"
	Proxy = ""
)

func main() {
	file, err := os.Open(File)
	if err != nil {
		return

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup
	for scanner.Scan() {
		invite := scanner.Text()
		if len(invite) == 0 {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			proxyUrl, err := url.Parse(Proxy)
			if err != nil {
				return

			}
			http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
			resp, err := http.Get(Api + invite)
			if err != nil {
				return

			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return

			}
			if strings.Contains(string(body), "Unknown Invite") {
				color.Green("[+] %s", invite)
				return
			}
			color.Red("[-] %s", invite)
		}()
	}
	wg.Wait()
	if err := scanner.Err(); err != nil {
		return

	}
}
