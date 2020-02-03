package converter

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/gomarkdown/markdown"
)

const (
	openTag  = `[code]`
	closeTag = `[/code]`
)

func Convert(r io.Reader) (io.Reader, error) {
	md, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("could not read content: %w", err)
	}
	bbcode := openTag + string(markdown.ToHTML(md, nil, nil)) + closeTag
	return strings.NewReader(bbcode), nil
}
