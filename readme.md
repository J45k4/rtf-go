# Rtf

## Strip rtf markup from string

StripRichTextFormat function removes rtf markup from string and returns new string.


Rtf text
```
{\rtf1\ansi\deff0\sdeasyworship2
{\fonttbl{\f0 Tahoma;}}
{\colortbl ;}
{\pard\sdlistlevel0\qc\qdef\sdewparatemplatestyle101{\*\sdasfactor 1}{\*\sdasbaseline 72.9}\sdastextstyle101\plain\sdewtemplatestyle101\fs146{\*\sdfsreal 72.9}{\*\sdfsdef 72.9}\sdfsauto hello\par}
{\pard\sdslidemarker\sdlistlevel0\qc\qdef\sdewparatemplatestyle101\plain\sdewtemplatestyle101\fs146{\*\sdfsreal 72.9}{\*\sdfsdef 72.9}\sdfsauto\par}
{\pard\sdlistlevel0\qc\qdef\sdewparatemplatestyle101\plain\sdewtemplatestyle101\fs146{\*\sdfsreal 72.9}{\*\sdfsdef 72.9}\sdfsauto hello\par}
}
```
becomes 
```
hello

hello
```
code
``` go
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
```