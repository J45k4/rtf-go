package rtf

import (
	"bytes"
	"golang.org/x/text/encoding/charmap"
	"regexp"
	"strconv"
	"strings"
)

var destinations = map[string]bool{
	"aftncn":             true,
	"aftnsep":            true,
	"aftnsepc":           true,
	"annotation":         true,
	"atnauthor":          true,
	"atndate":            true,
	"atnicn":             true,
	"atnid":              true,
	"atnparent":          true,
	"atnref":             true,
	"atntime":            true,
	"atrfend":            true,
	"atrfstart":          true,
	"author":             true,
	"background":         true,
	"bkmkend":            true,
	"bkmkstart":          true,
	"blipuid":            true,
	"buptim":             true,
	"category":           true,
	"colorschememapping": true,
	"colortbl":           true,
	"comment":            true,
	"company":            true,
	"creatim":            true,
	"datafield":          true,
	"datastore":          true,
	"defchp":             true,
	"defpap":             true,
	"do":                 true,
	"doccomm":            true,
	"docvar":             true,
	"dptxbxtext":         true,
	"ebcend":             true,
	"ebcstart":           true,
	"factoidname":        true,
	"falt":               true,
	"fchars":             true,
	"ffdeftext":          true,
	"ffentrymcr":         true,
	"ffexitmcr":          true,
	"ffformat":           true,
	"ffhelptext":         true,
	"ffl":                true,
	"ffname":             true,
	"ffstattext":         true,
	"field":              true,
	"file":               true,
	"filetbl":            true,
	"fldinst":            true,
	"fldrslt":            true,
	"fldtype":            true,
	"fname":              true,
	"fontemb":            true,
	"fontfile":           true,
	"fonttbl":            true,
	"footer":             true,
	"footerf":            true,
	"footerl":            true,
	"footerr":            true,
	"footnote":           true,
	"formfield":          true,
	"ftncn":              true,
	"ftnsep":             true,
	"ftnsepc":            true,
	"g":                  true,
	"generator":          true,
	"gridtbl":            true,
	"header":             true,
	"headerf":            true,
	"headerl":            true,
	"headerr":            true,
	"hl":                 true,
	"hlfr":               true,
	"hlinkbase":          true,
	"hlloc":              true,
	"hlsrc":              true,
	"hsv":                true,
	"htmltag":            true,
	"info":               true,
	"keycode":            true,
	"keywords":           true,
	"latentstyles":       true,
	"lchars":             true,
	"levelnumbers":       true,
	"leveltext":          true,
	"lfolevel":           true,
	"linkval":            true,
	"list":               true,
	"listlevel":          true,
	"listname":           true,
	"listoverride":       true,
	"listoverridetable":  true,
	"listpicture":        true,
	"liststylename":      true,
	"listtable":          true,
	"listtext":           true,
	"lsdlockedexcept":    true,
	"macc":               true,
	"maccPr":             true,
	"mailmerge":          true,
	"maln":               true,
	"malnScr":            true,
	"manager":            true,
	"margPr":             true,
	"mbar":               true,
	"mbarPr":             true,
	"mbaseJc":            true,
	"mbegChr":            true,
	"mborderBox":         true,
	"mborderBoxPr":       true,
	"mbox":               true,
	"mboxPr":             true,
	"mchr":               true,
	"mcount":             true,
	"mctrlPr":            true,
	"md":                 true,
	"mdeg":               true,
	"mdegHide":           true,
	"mden":               true,
	"mdiff":              true,
	"mdPr":               true,
	"me":                 true,
	"mendChr":            true,
	"meqArr":             true,
	"meqArrPr":           true,
	"mf":                 true,
	"mfName":             true,
	"mfPr":               true,
	"mfunc":              true,
	"mfuncPr":            true,
	"mgroupChr":          true,
	"mgroupChrPr":        true,
	"mgrow":              true,
	"mhideBot":           true,
	"mhideLeft":          true,
	"mhideRight":         true,
	"mhideTop":           true,
	"mhtmltag":           true,
	"mlim":               true,
	"mlimloc":            true,
	"mlimlow":            true,
	"mlimlowPr":          true,
	"mlimupp":            true,
	"mlimuppPr":          true,
	"mm":                 true,
	"mmaddfieldname":     true,
	"mmath":              true,
	"mmathPict":          true,
	"mmathPr":            true,
	"mmaxdist":           true,
	"mmc":                true,
	"mmcJc":              true,
	"mmconnectstr":       true,
	"mmconnectstrdata":   true,
	"mmcPr":              true,
	"mmcs":               true,
	"mmdatasource":       true,
	"mmheadersource":     true,
	"mmmailsubject":      true,
	"mmodso":             true,
	"mmodsofilter":       true,
	"mmodsofldmpdata":    true,
	"mmodsomappedname":   true,
	"mmodsoname":         true,
	"mmodsorecipdata":    true,
	"mmodsosort":         true,
	"mmodsosrc":          true,
	"mmodsotable":        true,
	"mmodsoudl":          true,
	"mmodsoudldata":      true,
	"mmodsouniquetag":    true,
	"mmPr":               true,
	"mmquery":            true,
	"mmr":                true,
	"mnary":              true,
	"mnaryPr":            true,
	"mnoBreak":           true,
	"mnum":               true,
	"mobjDist":           true,
	"moMath":             true,
	"moMathPara":         true,
	"moMathParaPr":       true,
	"mopEmu":             true,
	"mphant":             true,
	"mphantPr":           true,
	"mplcHide":           true,
	"mpos":               true,
	"mr":                 true,
	"mrad":               true,
	"mradPr":             true,
	"mrPr":               true,
	"msepChr":            true,
	"mshow":              true,
	"mshp":               true,
	"msPre":              true,
	"msPrePr":            true,
	"msSub":              true,
	"msSubPr":            true,
	"msSubSup":           true,
	"msSubSupPr":         true,
	"msSup":              true,
	"msSupPr":            true,
	"mstrikeBLTR":        true,
	"mstrikeH":           true,
	"mstrikeTLBR":        true,
	"mstrikeV":           true,
	"msub":               true,
	"msubHide":           true,
	"msup":               true,
	"msupHide":           true,
	"mtransp":            true,
	"mtype":              true,
	"mvertJc":            true,
	"mvfmf":              true,
	"mvfml":              true,
	"mvtof":              true,
	"mvtol":              true,
	"mzeroAsc":           true,
	"mzeroDesc":          true,
	"mzeroWid":           true,
	"nesttableprops":     true,
	"nextfile":           true,
	"nonesttables":       true,
	"objalias":           true,
	"objclass":           true,
	"objdata":            true,
	"object":             true,
	"objname":            true,
	"objsect":            true,
	"objtime":            true,
	"oldcprops":          true,
	"oldpprops":          true,
	"oldsprops":          true,
	"oldtprops":          true,
	"oleclsid":           true,
	"operator":           true,
	"panose":             true,
	"password":           true,
	"passwordhash":       true,
	"pgp":                true,
	"pgptbl":             true,
	"picprop":            true,
	"pict":               true,
	"pn":                 true,
	"pnseclvl":           true,
	"pntext":             true,
	"pntxta":             true,
	"pntxtb":             true,
	"printim":            true,
	"private":            true,
	"propname":           true,
	"protend":            true,
	"protstart":          true,
	"protusertbl":        true,
	"pxe":                true,
	"result":             true,
	"revtbl":             true,
	"revtim":             true,
	"rsidtbl":            true,
	"rxe":                true,
	"shp":                true,
	"shpgrp":             true,
	"shpinst":            true,
	"shppict":            true,
	"shprslt":            true,
	"shptxt":             true,
	"sn":                 true,
	"sp":                 true,
	"staticval":          true,
	"stylesheet":         true,
	"subject":            true,
	"sv":                 true,
	"svb":                true,
	"tc":                 true,
	"template":           true,
	"themedata":          true,
	"title":              true,
	"txe":                true,
	"ud":                 true,
	"upr":                true,
	"userprops":          true,
	"wgrffmtfilter":      true,
	"windowcaption":      true,
	"writereservation":   true,
	"writereservhash":    true,
	"xe":                 true,
	"xform":              true,
	"xmlattrname":        true,
	"xmlattrvalue":       true,
	"xmlclose":           true,
	"xmlname":            true,
	"xmlnstbl":           true,
	"xmlopen":            true,
}

var specialCharacters = map[string]string{
	"par":       "\n",
	"sect":      "\n\n",
	"page":      "\n\n",
	"line":      "\n",
	"tab":       "\t",
	"emdash":    "\u2014",
	"endash":    "\u2013",
	"emspace":   "\u2003",
	"enspace":   "\u2002",
	"qmspace":   "\u2005",
	"bullet":    "\u2022",
	"lquote":    "\u2018",
	"rquote":    "\u2019",
	"ldblquote": "\u201C",
	"rdblquote": "\u201D",
}

var charmaps = map[string]*charmap.Charmap{
	"437": charmap.CodePage437,
	//	"708":  nil,
	//	"709":  nil,
	//	"710":  nil,
	//	"711":  nil,
	//	"720":  nil,
	//	"819":  nil,
	"850": charmap.CodePage850,
	"852": charmap.CodePage852,
	"860": charmap.CodePage860,
	"862": charmap.CodePage862,
	"863": charmap.CodePage863,
	//	"864":  nil,
	"865": charmap.CodePage865,
	"866": charmap.CodePage866,
	//	"874":  nil,
	//	"932":  nil,
	//	"936":  nil,
	//	"949":  nil,
	//	"950":  nil,
	"1250": charmap.Windows1250,
	"1251": charmap.Windows1251,
	"1252": charmap.Windows1252,
	"1253": charmap.Windows1253,
	"1254": charmap.Windows1254,
	"1255": charmap.Windows1255,
	"1256": charmap.Windows1256,
	"1257": charmap.Windows1257,
	"1258": charmap.Windows1258,
	//	"1361": nil,
}

var rtfRegex = regexp.MustCompile(
	"(?i)" +
		`\\([a-z]{1,32})(-?\d{1,10})?[ ]?` +
		`|\\'([0-9a-f]{2})` +
		`|\\([^a-z])` +
		`|([{}])` +
		`|[\r\n]+` +
		`|(.)`)

type stackEntry struct {
	NumberOfCharactersToSkip int
	Ignorable                bool
}

func newStackEntry(numberOfCharactersToSkip int, ignorable bool) stackEntry {
	return stackEntry{
		NumberOfCharactersToSkip: numberOfCharactersToSkip,
		Ignorable:                ignorable,
	}
}

// Removes rtf characters from string and returns the new string.
func StripRichTextFormat(inputRtf string) string {
	var charMap *charmap.Charmap = nil
	var stack []stackEntry
	var ignorable bool
	ucskip := 1
	curskip := 0

	matches := rtfRegex.FindAllStringSubmatch(inputRtf, -1)
	var returnBuffer bytes.Buffer

	for _, match := range matches {
		word := match[1]
		arg := match[2]
		hex := match[3]
		character := match[4]
		brace := match[5]
		tchar := match[6]

		switch {
		case tchar != "":
			if curskip > 0 {
				curskip -= 1
			} else if !ignorable {
				returnBuffer.WriteString(tchar)
			}
		case brace != "":
			curskip = 0
			if brace == "{" {
				stack = append(
					stack, newStackEntry(ucskip, ignorable))
			} else if brace == "}" {
				l := len(stack)
				entry := stack[l-1]
				stack = stack[:l-1]
				ucskip = entry.NumberOfCharactersToSkip
				ignorable = entry.Ignorable
			}
		case character != "":
			curskip = 0
			if character == "~" {
				if !ignorable {
					returnBuffer.WriteString("\xA0")
				}
			} else if strings.Contains("{}\\", character) {
				if !ignorable {
					returnBuffer.WriteString(character)
				}
			} else if character == "*" {
				ignorable = true
			}
		case word != "":
			curskip = 0
			if destinations[word] {
				ignorable = true
			} else if ignorable {
			} else if specialCharacters[word] != "" {
				returnBuffer.WriteString(
					specialCharacters[word])
			} else if word == "ansicpg" {
				charMap = charmaps[arg]
			} else if word == "uc" {
				i, _ := strconv.Atoi(arg)
				ucskip = i
			} else if word == "u" {
				c, _ := strconv.Atoi(arg)
				if c < 0 {
					c += 0x10000
				}
				returnBuffer.WriteRune(rune(c))
				curskip = ucskip
			}
		case hex != "":
			if curskip > 0 {
				curskip -= 1
			} else if !ignorable {
				c, _ := strconv.ParseInt(hex, 16, 0)
				if charMap == nil {
					returnBuffer.WriteRune(rune(c))
				} else {
					returnBuffer.WriteRune(
						charMap.DecodeByte(byte(c)))
				}
			}
		}
	}
	return returnBuffer.String()
}
