package details

import (
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"

	"regexp"
	"strings"
)

type alertParser struct{}

var defaultAlertsParser = &alertParser{}

func NewAlertsParser() parser.BlockParser {
	return defaultAlertsParser
}

func (b *alertParser) Trigger() []byte {
	return []byte{'>'}
}

var regex = regexp.MustCompile(`^\[!(?P<kind>[\w]+)\](?P<closed>-{0,1})($|\s+(?P<title>.*))`)

func (b *alertParser) process(reader text.Reader) (bool, int) {
	// This is slighlty modified code from https://github.com/yuin/goldmark.git
	// Originally written by Yusuke Inuzuka, licensed under MIT License

	line, _ := reader.PeekLine()
	w, pos := util.IndentWidth(line, reader.LineOffset())
	if w > 3 || pos >= len(line) || line[pos] != '>' {
		return false, 0
	}

	advance_by := 1

	if pos + advance_by >= len(line) || line[pos+advance_by] == '\n' {
		return true, advance_by
	}
	if line[pos+advance_by] == ' ' || line[pos+advance_by] == '\t' {
		advance_by++
	}

	if line[pos+advance_by-1] == '\t' {
		reader.SetPadding(2)
	}

	return true, advance_by
}

func (b *alertParser) Open(parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {

  // check if we are inside of a block quote
	ok, advance_by := b.process(reader)
	if !ok {
		return nil, parser.NoChildren
	}

	line, _ := reader.PeekLine()

  // empty blockquote
  if len(line) <= advance_by {
    return nil, parser.NoChildren
  }

  // right after `>` and up to one space
  subline := line[advance_by:]
	if !regex.Match(subline) {
		return nil, parser.NoChildren
	}

	match := regex.FindSubmatch(subline)

	kind := match[1]
	closed := match[2]

	alert := NewAlerts()

	alert.SetAttributeString("kind", kind)
	alert.SetAttributeString("closed", len(closed) != 0)

	i := strings.Index(string(line), "]")
	reader.Advance(i)

	return alert, parser.HasChildren
}

func (b *alertParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	ok, advance_by := b.process(reader)
	if !ok {
		return parser.Close
	}

	reader.Advance(advance_by)

	return parser.Continue | parser.HasChildren
}

func (b *alertParser) Close(node gast.Node, reader text.Reader, pc parser.Context) {
	// nothing to do
}

func (b *alertParser) CanInterruptParagraph() bool {
	return true
}

func (b *alertParser) CanAcceptIndentedLine() bool {
	return false
}
