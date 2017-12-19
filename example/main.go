package main

import (
	"io/ioutil"
	"os"

	"github.com/J45k4/rtf"
)

func main() {
	b, _ := ioutil.ReadFile("rtftext.rtf")
	f, _ := os.Create("text.txt")
	f.WriteString(rtf.StripRichTextFormat(string(b)))
}
