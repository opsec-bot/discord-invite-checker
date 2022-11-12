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
	InFile  = "in.txt"
	OutFile = "out.txt"
	Api     = "https://discordapp.com/api/invite/"
	Proxy   = ""
)

func main() {
	if _, err := os.Stat(OutFile); os.IsNotExist(err) {
		_, err := os.Create(OutFile)
		if err != nil {
			return
		}
	}
	err := ioutil.WriteFile(OutFile, []byte(""), 0644)
	if err != nil {
		return
	}

	file, err := os.Open(InFile)
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
				color.Green("[+] %s > Unclaimed!", invite)

				f, err := os.OpenFile(OutFile, os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					panic(err)
				}

				defer f.Close()

				if _, err = f.WriteString(invite + "\n"); err != nil {
					panic(err)
				}
			} else {
				color.Red("[-] %s > Claimed!", invite)
				return
			}
		}()
	}
	wg.Wait()
	if err := scanner.Err(); err != nil {
		return

	}
}
