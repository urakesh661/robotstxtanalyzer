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

	var byte []byte
	buffer := bytes.NewBuffer(byte)
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
		key := strings.Split(path, ":")
		word := strings.Split(key[1], "/")
		for _, url := range word {
			fmt.Printf("%s\n", url)
		}

	}

}
