package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	err := DownloadFile("//users/will/source/a.zip", "http://ipv4.download.thinkbroadband.com/20MB.zip")

	fmt.Println(err)

}

func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
