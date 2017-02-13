package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"os"
	"golang.org/x/net/html"
	"time"
	"strings"
)

type Chapter struct{
	Title string
	Body string
	next string
}

func main() {

	fo, err := os.Create("mianzhuan2.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	i := "6599206"
	for {
		chapter := ReadChapter(i)
		if chapter == nil {
			break
		}

		fo.WriteString(chapter.Title)
		fo.WriteString("\n")
		fo.WriteString(chapter.Body)


		fmt.Println("sleep...")
		time.Sleep(2000 * time.Millisecond)
	}

}

func ReadChapter(i string) *Chapter {
	//url := fmt.Sprintf("http://mianzhuan.wddsnxn.org/%d.html", i)
	url := fmt.Sprintf("http://www.szzyue.com/dushu/10/10326/%s.html", i)
	fmt.Println("reading ", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	result := &Chapter{}

	title := doc.Find("#amain > dl > dd:nth-child(2) > h1")
	result.Title = title.Text()

	text := doc.Find("#contents").Get(0)
	node := text.FirstChild
	for {
		if node == nil {
			break
		}

		if node.Type == html.TextNode {
			text := node.Data
			text = strings.TrimPrefix(text, "\"")
			text = strings.TrimSuffix(text, "\"")
			result.Body += node.Data
			result.Body += "\n"
		}
		node = node.NextSibling
	}
	fmt.Println(result.Body)
	return result
}