package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"os"
)

type Chapter struct{
	Title string
	Body string
}

func main() {

	fo, err := os.Create("mianzhuan.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	i := 244
	for {
		chapter := ReadChapter(i)
		fo.WriteString(chapter.Title)
		fo.WriteString(chapter.Body)
		break
	}

}

func ReadChapter(i int) *Chapter {
	url := fmt.Sprintf("http://mianzhuan.wddsnxn.org/%d.html", i)
	fmt.Println("reading ", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	result := &Chapter{}

	title := doc.Find("body > div.site.clearfix > div > div.chaptertitle.clearfix > h1").Each(func(i int, s *goquery.Selection) {
		fmt.Println(i)
	})
	result.Title = title.Text()

	text := doc.Find("#BookText").Get(0)
	node := text.FirstChild
	for {
		result.Body += node.Data
		result.Body += "\n"
		node = node.NextSibling
		if node == nil {
			break
		}
	}
	return result
}