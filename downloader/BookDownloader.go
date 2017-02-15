package downloader

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var decoder = simplifiedchinese.GBK.NewDecoder()

type Chapter struct {
	Title string
	Body  string
	next  string
}

type BookDownloader struct {
	UrlPattern string
	ChapterPerFile int

	fileIndex int
	fo        *os.File
	nextUrl   string
}

func (d BookDownloader) DownLoad() {
	d.fileIndex = 0
	d.ChapterPerFile = 200
	d.nextUrl = "5583810.html"

	i := 1
	switchToNextFile(&d)
	fmt.Printf("%d: %s", d.fileIndex, d.fo.Name())
	
	for {
		if i == d.ChapterPerFile {
			switchToNextFile(&d)
			fmt.Printf("%d: %s", d.fileIndex, d.fo.Name())
			i = 1
		}

		chapter := d.readChapter(d.nextUrl)
		if chapter == nil {
			break
		}

		d.fo.WriteString(chapter.Title)
		d.fo.WriteString("\n")
		d.fo.WriteString(chapter.Body)

		fmt.Println("sleep...")
		time.Sleep(2000 * time.Millisecond)
		d.nextUrl = chapter.next

		i++
	}
}

func switchToNextFile(d *BookDownloader) {
	fmt.Println("switch to next file ...")
	if d.fo != nil {
		d.fo.Close()
	}

	file, err := os.Create(fmt.Sprintf("mianzhuan_%d.txt", d.fileIndex))
	fmt.Println(file.Name())
	d.fo = file

	d.fileIndex++
	if err != nil {
		panic(err)
	}
}

func (d BookDownloader) readChapter(i string) *Chapter {
	//url := fmt.Sprintf("http://mianzhuan.wddsnxn.org/%d.html", i)
	url := fmt.Sprintf("http://www.szzyue.com/dushu/10/10326/%s", i)
	fmt.Println("reading ", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	result := &Chapter{}

	title := doc.Find("#amain > dl > dd:nth-child(2) > h1")
	t, _ := decoder.String(title.Text())
	t = strings.TrimPrefix(t, "章节目录")
	result.Title = strings.TrimSpace(t)
	fmt.Println(result.Title)

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
			text = strings.TrimSpace(text)
			text, _ = decoder.String(text)
			result.Body += text
			result.Body += "\n"
		}
		node = node.NextSibling
	}

	foot := ""
	doc.Find("#footlink").First().Find("a").Each(func(i int, s *goquery.Selection) {
		t, _ := decoder.String(s.Text())
		if t == "下一页" {
			foot, _ = s.Attr("href")
		}
	})

	result.next, _ = decoder.String(foot)

	fmt.Println(result.Body)
	fmt.Println("next: ", result.next)
	return result
}
