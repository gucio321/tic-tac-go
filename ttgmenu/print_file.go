package ttgmenu

import (
	"fmt"
	"io/ioutil"

	"github.com/jaytaylor/html2text"
	"github.com/russross/blackfriday"

	"github.com/gucio321/tic-tac-go/ttgcommon"
)

func readMarkdown(path string) []string {
	var data []byte

	var err error

	// nolint:gosec // this is ok
	if data, err = ioutil.ReadFile(path); err != nil {
		return []string{fmt.Sprintf("%s is missing.", path), "Visit https://github.com/gucio321/tic-tac-go to see it."}
	}

	html := blackfriday.MarkdownBasic(data)

	text, err := html2text.FromString(string(html), html2text.Options{PrettyTables: true})
	if err != nil {
		return []string{
			fmt.Sprintf("Error loading %s:", path), err.Error(),
			"Please raport it on https://github.com/gucio321/tic-tac-go",
		}
	}

	lines := ttgcommon.SplitIntoLinesWithMaxWidth(text, 70)

	return lines
}
