package main

import (
	// "fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
)

func main() {
	filename := "./README.md"
	file, _ := os.OpenFile(filename, os.O_RDONLY, 0644)
	defer file.Close()
	b, _ := ioutil.ReadAll(file)
	unsafe := blackfriday.MarkdownCommon(b)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	of, _ := os.OpenFile("readme.html", os.O_CREATE|os.O_RDWR, 0666)
	defer of.Close()
	of.Write(html)
}
