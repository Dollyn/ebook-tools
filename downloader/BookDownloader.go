package downloader

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
)


type UrlMode int

const Pattern UrlMode = 0
const FullText UrlMode = 1

type Chapter struct {
	Title string
	Body  string
	next  string
}

type BookDownloader struct {
	UrlPattern string
	Start      string

	TitleSelector string
	ContentSelector string
	NextSelector string

	GBK bool

	ChapterPerFile  int
	FileNamePattern string

	fileIndex int
	fo        *os.File
	nextUrl   string

	decoder *encoding.Decoder
}

func (d *BookDownloader) DownLoad() {
	d.fileIndex = 0
	d.nextUrl = d.Start

	if d.GBK {
		d.decoder = simplifiedchinese.GBK.NewDecoder()
	} else {
		d.decoder = unicode.UTF8.NewDecoder()
	}

	i := 1
	switchToNextFile(d)
	fmt.Printf("%d: %s", d.fileIndex, d.fo.Name())

	for {
		if i == d.ChapterPerFile {
			switchToNextFile(d)
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
		d.fo.WriteString("\n")

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

	file, err := os.Create(fmt.Sprintf(d.FileNamePattern, d.fileIndex))
	fmt.Println(file.Name())
	d.fo = file

	d.fileIndex++
	if err != nil {
		panic(err)
	}
}

func (d *BookDownloader) readChapter(i string) *Chapter {
	//url := fmt.Sprintf("http://mianzhuan.wddsnxn.org/%d.html", i)
	url := fmt.Sprintf(d.UrlPattern, i)
	fmt.Println("reading ", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	result := &Chapter{}

	title := doc.Find(d.TitleSelector)
	t, _ := d.decoder.String(title.Text())
	t = strings.TrimPrefix(t, "章节目录")
	result.Title = strings.TrimSpace(t)
	fmt.Println(result.Title)

	text := doc.Find(d.ContentSelector).Get(0)
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
			text, _ = d.decoder.String(text)
			result.Body += text
			result.Body += "\n"
		}
		node = node.NextSibling
	}

	foot := ""

	foot, _ = doc.Find(d.NextSelector).First().Attr("href")

	result.next, _ = d.decoder.String(foot)

	fmt.Println(result.Body)
	fmt.Println("next: ", result.next)
	return result
}
