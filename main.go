package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var strConvertBody string

func main() {

	var robotEntries []string

	domains := bufio.NewScanner(os.Stdin)

	for domains.Scan() {
		robotEntries = append(robotEntries, domains.Text())
	}
	for _, url := range robotEntries {
		var url = url + "/robots.txt"
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		strConvertBody = string(body)
		modifiedRobotsTxt()
	}
}

func modifiedRobotsTxt() {

	const (
		allow1 = "Disallow"
		allow2 = "Allow"
	)

	var byt []byte
	buffer := bytes.NewBuffer(byt)
	scanner := bufio.NewScanner(strings.NewReader(strConvertBody))

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), allow1) || strings.Contains(scanner.Text(), allow2) {
			buffer.Write(scanner.Bytes())
			buffer.WriteString("\n")
		}
	}

	robotsEntry := buffer.String()
	scanner = bufio.NewScanner(strings.NewReader(robotsEntry))
	for scanner.Scan() {
		path := scanner.Text()
		if strings.Contains(path, "https") {
			keyHttps := strings.Split(path, "https://")
			for _, urlHttps := range keyHttps {
				keyHttps := urlHttps[strings.LastIndex(urlHttps, " ")+1:]
				wordHttps := strings.Split(keyHttps, "/")
				for _, urlHttps := range wordHttps {
					if len(urlHttps) > 0 {
						if urlHttps == allow1 || urlHttps == allow2 {
							continue
						} else {
							fmt.Printf("%s\n", urlHttps)
						}
					}
				}
			}
		} else {
			key := strings.Split(path, ":")
			word := strings.Split(key[1], "/")
			for _, url := range word {
				if len(url) > 0 {
					if url == allow1 || url == allow2 {
						continue
					}
					fmt.Printf("%s\n", url)
				}
			}
		}
	}
}
