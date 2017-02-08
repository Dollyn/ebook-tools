package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, _ := client.Get("http://mianzhuan.wddsnxn.org/247.html")

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		//case html.ErrorToken:
		//return z.Err()
		case html.TextToken:
			fmt.Print(z.Text())
			//case html.StartTagToken, html.EndTagToken:

		}
	}
}
