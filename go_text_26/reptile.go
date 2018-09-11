package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	//  method   url    forminformation
	request, _ := http.NewRequest(http.MethodGet, "https://www.bilibili.com/", nil)
	// add headInformation
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")

	client := http.Client{
		//req= change url  via[]=wait change urls
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
