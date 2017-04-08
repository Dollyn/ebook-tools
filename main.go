package main

import "github.com/dollyn/ebook-tools/downloader"

func main() {
	//bd := &downloader.BookDownloader{
	//	UrlPattern: "http://www.szzyue.com/dushu/10/10326/%s",
	//	Start:      "5583810.html",
	//
	//	TitleSelector:   "#amain > dl > dd:nth-child(2) > h1",
	//	ContentSelector: "#contents",
	//	NextSelector:    "#footlink > a:nth-child(3)",
	//
	//	ChapterPerFile:  300,
	//	FileNamePattern: "mianzhuan_%d.txt"}
	//bd.DownLoad()

	//bd := &downloader.BookDownloader{
	//	UrlPattern: "%s",
	//	Start:      "http://www.wddsnxn.org/mindiaojuyiwenlu/244.html",
	//
	//	TitleSelector:   "body > div.site.clearfix > div > div.chaptertitle.clearfix > h1",
	//	ContentSelector: "#BookText",
	//	NextSelector:    "body > div.site.clearfix > div > div:nth-child(6) > div > h2 > a:nth-child(3)",
	//
	//	ChapterPerFile:  300,
	//	FileNamePattern: "mindiaoju_%d.txt"}
	//bd.DownLoad()

	bd := &downloader.BookDownloader{
		UrlPattern: "%s",
		Start:      "http://mianzhuan.wddsnxn.org/1088.html",

		TitleSelector:   "body > div.site.clearfix > div > div.chaptertitle.clearfix > h1",
		ContentSelector: "#BookText",
		NextSelector:    "body > div.site.clearfix > div > div:nth-child(6) > div > h2 > a:nth-child(3)",

		ChapterPerFile:  300,
		FileIndex: 4,
		FileNamePattern: "mianzhuan_%d.txt"}
	bd.DownLoad()

}
