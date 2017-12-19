package rtf

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

var destinations = []string{
	"aftncn", "aftnsep", "aftnsepc", "annotation", "atnauthor", "atndate", "atnicn", "atnid",
	"atnparent", "atnref", "atntime", "atrfend", "atrfstart", "author", "background",
	"bkmkend", "bkmkstart", "blipuid", "buptim", "category", "colorschememapping",
	"colortbl", "comment", "company", "creatim", "datafield", "datastore", "defchp", "defpap",
	"do", "doccomm", "docvar", "dptxbxtext", "ebcend", "ebcstart", "factoidname", "falt",
	"fchars", "ffdeftext", "ffentrymcr", "ffexitmcr", "ffformat", "ffhelptext", "ffl",
	"ffname", "ffstattext", "field", "file", "filetbl", "fldinst", "fldrslt", "fldtype",
	"fname", "fontemb", "fontfile", "fonttbl", "footer", "footerf", "footerl", "footerr",
	"footnote", "formfield", "ftncn", "ftnsep", "ftnsepc", "g", "generator", "gridtbl",
	"header", "headerf", "headerl", "headerr", "hl", "hlfr", "hlinkbase", "hlloc", "hlsrc",
	"hsv", "htmltag", "info", "keycode", "keywords", "latentstyles", "lchars", "levelnumbers",
	"leveltext", "lfolevel", "linkval", "list", "listlevel", "listname", "listoverride",
	"listoverridetable", "listpicture", "liststylename", "listtable", "listtext",
	"lsdlockedexcept", "macc", "maccPr", "mailmerge", "maln", "malnScr", "manager", "margPr",
	"mbar", "mbarPr", "mbaseJc", "mbegChr", "mborderBox", "mborderBoxPr", "mbox", "mboxPr",
	"mchr", "mcount", "mctrlPr", "md", "mdeg", "mdegHide", "mden", "mdiff", "mdPr", "me",
	"mendChr", "meqArr", "meqArrPr", "mf", "mfName", "mfPr", "mfunc", "mfuncPr", "mgroupChr",
	"mgroupChrPr", "mgrow", "mhideBot", "mhideLeft", "mhideRight", "mhideTop", "mhtmltag",
	"mlim", "mlimloc", "mlimlow", "mlimlowPr", "mlimupp", "mlimuppPr", "mm", "mmaddfieldname",
	"mmath", "mmathPict", "mmathPr", "mmaxdist", "mmc", "mmcJc", "mmconnectstr",
	"mmconnectstrdata", "mmcPr", "mmcs", "mmdatasource", "mmheadersource", "mmmailsubject",
	"mmodso", "mmodsofilter", "mmodsofldmpdata", "mmodsomappedname", "mmodsoname",
	"mmodsorecipdata", "mmodsosort", "mmodsosrc", "mmodsotable", "mmodsoudl",
	"mmodsoudldata", "mmodsouniquetag", "mmPr", "mmquery", "mmr", "mnary", "mnaryPr",
	"mnoBreak", "mnum", "mobjDist", "moMath", "moMathPara", "moMathParaPr", "mopEmu",
	"mphant", "mphantPr", "mplcHide", "mpos", "mr", "mrad", "mradPr", "mrPr", "msepChr",
	"mshow", "mshp", "msPre", "msPrePr", "msSub", "msSubPr", "msSubSup", "msSubSupPr", "msSup",
	"msSupPr", "mstrikeBLTR", "mstrikeH", "mstrikeTLBR", "mstrikeV", "msub", "msubHide",
	"msup", "msupHide", "mtransp", "mtype", "mvertJc", "mvfmf", "mvfml", "mvtof", "mvtol",
	"mzeroAsc", "mzeroDesc", "mzeroWid", "nesttableprops", "nextfile", "nonesttables",
	"objalias", "objclass", "objdata", "object", "objname", "objsect", "objtime", "oldcprops",
	"oldpprops", "oldsprops", "oldtprops", "oleclsid", "operator", "panose", "password",
	"passwordhash", "pgp", "pgptbl", "picprop", "pict", "pn", "pnseclvl", "pntext", "pntxta",
	"pntxtb", "printim", "private", "propname", "protend", "protstart", "protusertbl", "pxe",
	"result", "revtbl", "revtim", "rsidtbl", "rxe", "shp", "shpgrp", "shpinst",
	"shppict", "shprslt", "shptxt", "sn", "sp", "staticval", "stylesheet", "subject", "sv",
	"svb", "tc", "template", "themedata", "title", "txe", "ud", "upr", "userprops",
	"wgrffmtfilter", "windowcaption", "writereservation", "writereservhash", "xe", "xform",
	"xmlattrname", "xmlattrvalue", "xmlclose", "xmlname", "xmlnstbl",
	"xmlopen",
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

var rtfRegex, _ = regexp.Compile("(?i)" + `\\([a-z]{1,32})(-?\d{1,10})?[ ]?|\\'([0-9a-f]{2})|\\([^a-z])|([{}])|[\r\n]+|(.)`)

type stackEntry struct {
	NumberOfCharactersToSkip int
	Ignorable                bool
}

func newStackEntry(
	numberOfCharactersToSkip int,
	ignorable bool,
) stackEntry {
	return stackEntry{
		NumberOfCharactersToSkip: numberOfCharactersToSkip,
		Ignorable:                ignorable,
	}
}

func isDestination(word string) bool {
	for _, destination := range destinations {
		if destination == word {
			return true
		}
	}
	return false
}

// Removes rtf characters from string and returns the new string.
func StripRichTextFormat(
	inputRtf string,
) string {
	if inputRtf == "" {
		return ""
	}

	var stack stack.Stack
	var ignorable bool
	ucskip := 1
	curskip := 0

	matches := rtfRegex.FindAllStringSubmatch(inputRtf, -1)

	if len(matches) == 0 {
		return inputRtf
	}

	var returnString string

	for _, match := range matches {
		word := match[1]
		arg := match[2]
		hex := match[3]
		character := match[4]
		brace := match[5]
		tchar := match[6]

		if brace != "" {
			curskip = 0
			if brace == "{" {
				stack.Push(newStackEntry(ucskip, ignorable))
			} else if brace == "}" {
				entry := stack.Pop().(stackEntry)
				ucskip = entry.NumberOfCharactersToSkip
				ignorable = entry.Ignorable
			}
		} else if character != "" {
			curskip = 0
			if character == "~" {
				if !ignorable {
					returnString += "\xA0"
				}
			} else if strings.Contains("{}\\", character) {
				if !ignorable {
					returnString += character
				}
			} else if character == "*" {
				ignorable = true
			}
		} else if word != "" {
			curskip = 0
			if isDestination(word) {
				ignorable = true
			} else if ignorable {
			} else if specialCharacters[word] != "" {
				returnString += specialCharacters[word]
			} else if word == "uc" {
				i, _ := strconv.Atoi(arg)
				ucskip = i
			} else if word == "u" {
				c, _ := strconv.Atoi(arg)
				if c < 0 {
					c += 0x10000
				}
				returnString += strconv.Itoa(c)
				curskip = ucskip
			}
		} else if hex != "" {
			if curskip > 0 {
				curskip -= 1
			} else if !ignorable {
				c, _ := strconv.ParseInt(hex, 16, 0)
				returnString += strconv.Itoa(int(c))
			}
		} else if tchar != "" {
			if curskip > 0 {
				curskip -= 1
			} else if ignorable == false {
				returnString += tchar
			}
		}
	}

	return returnString
}
