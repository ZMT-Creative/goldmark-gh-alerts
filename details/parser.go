package details

import (
	"github.com/ZMT-Creative/goldmark-gh-alerts/body"
	"github.com/ZMT-Creative/goldmark-gh-alerts/kinds"
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
    // This is slightly modified code from https://github.com/yuin/goldmark.git
    // Originally written by Yusuke Inuzuka, licensed under MIT License

    line, _ := reader.PeekLine()
    w, pos := util.IndentWidth(line, reader.LineOffset())
    if w > 3 || pos >= len(line) || line[pos] != '>' {
        return false, 0
    }

    advanceBy := 1

    if pos+advanceBy >= len(line) || line[pos+advanceBy] == '\n' {
        return true, advanceBy
    }
    if line[pos+advanceBy] == ' ' || line[pos+advanceBy] == '\t' {
        advanceBy++
    }

    if line[pos+advanceBy-1] == '\t' {
        reader.SetPadding(2)
    }

    return true, advanceBy
}

func (b *alertParser) Open(parent gast.Node, reader text.Reader, pc parser.Context) (gast.Node, parser.State) {

  // check if we are inside of a block quote
	ok, advanceBy := b.process(reader)
	if !ok {
		return nil, parser.NoChildren
	}

	line, _ := reader.PeekLine()

	// empty blockquote
	if len(line) <= advanceBy {
		return nil, parser.NoChildren
	}

  // right after `>` and up to one space
    subline := line[advanceBy:]
    match := regex.FindSubmatch(subline)
    if match == nil {
        return nil, parser.NoChildren
    }

	kind := match[1]
	closed := match[2]
    title := match[3]

	alert := NewAlerts()

	alert.SetAttributeString("kind", kind)
	alert.SetAttributeString("closed", len(closed) != 0)
	alert.SetAttributeString("title", title)

	i := strings.Index(string(line), "]")
	reader.Advance(i)

	return alert, parser.HasChildren
}

func (b *alertParser) Continue(node gast.Node, reader text.Reader, pc parser.Context) parser.State {
	ok, advanceBy := b.process(reader)
	if !ok {
		return parser.Close
	}

	reader.Advance(advanceBy)

	return parser.Continue | parser.HasChildren
}

func (b *alertParser) Close(node gast.Node, reader text.Reader, pc parser.Context) {
    // Restructure nodes to have a proper AlertsHeader and AlertsBody hierarchy
    var header gast.Node
    var bodyChildren []gast.Node

    for c := node.FirstChild(); c != nil; {
        next := c.NextSibling()
        if c.Kind() == kinds.KindAlertsHeader {
            header = c
        } else {
            bodyChildren = append(bodyChildren, c)
        }
        c = next
    }

    // Re-parent children
    node.RemoveChildren(node)
    if header != nil {
        node.AppendChild(node, header)
    }

    if len(bodyChildren) > 0 {
        bodyNode := body.NewAlertsBody()
        for _, child := range bodyChildren {
            bodyNode.AppendChild(bodyNode, child)
        }
        node.AppendChild(node, bodyNode)
    }
}

func (b *alertParser) CanInterruptParagraph() bool {
	return true
}

func (b *alertParser) CanAcceptIndentedLine() bool {
	return false
}
