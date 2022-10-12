package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

/*
Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Input site")
		return
	}
	site := os.Args[len(os.Args)-1]

	Download(site)
}

func Download(site string) error {
	if len(site) < 8 {
		fmt.Println("Write http prot")
		return errors.New("Need http protocol")
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{Transport: tr,
		Timeout: timeout}
	req, err := http.NewRequest("GET", site, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	req.Header.Add("Accept", "text/html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	nm := make([]rune, 0, 5)

	if site[7:8] == "/" {
		for _, v := range site[8:] {
			if string(v) == "/" {
				break
			}
			nm = append(nm, v)

		}

		os.Mkdir(string(nm), 0777)
	} else {
		for _, v := range site[7:] {
			if string(v) == "/" {
				break
			}
			nm = append(nm, v)

		}
		os.Mkdir(string(nm), 0777)
	}
	mainFile, err := os.Create(string(nm) + "/index.html")
	data, err := ioutil.ReadAll(resp.Body)
	mainFile.Write(data)
	reg := regexp.MustCompile(`(http|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	result := reg.FindAll(data, -1)
	urls := make([]string, 0, len(result))
	for i := 0; i < len(result); i++ {
		urls = append(urls, string(result[i]))
	}

	for i, url := range urls {
		num := strconv.Itoa(i)
		downUrls(url, string(nm), num)

	}
	return nil
}

func downUrls(url, path string, i string) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{Transport: tr,
		Timeout: timeout}
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "text/html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	File, _ := os.Create(path + "/index" + i + ".html")
	data, _ := ioutil.ReadAll(resp.Body)
	File.Write(data)

}
