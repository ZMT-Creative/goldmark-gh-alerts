package summary

import (
	"strings"

	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Icons map[string]string

type AlertsHeaderHTMLRenderer struct {
	html.Config
	Icons
}

func NewAlertsHeaderHTMLRendererWithIcons(icons Icons, opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsHeaderHTMLRenderer{
		Config: html.NewConfig(),
		Icons:  icons,
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func NewAlertsHeaderHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsHeaderHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *AlertsHeaderHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindAlertsHeader, r.renderAlertsHeader)
}

func (r *AlertsHeaderHTMLRenderer) renderAlertsHeader(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		caser := cases.Title(language.English, cases.Compact)
		w.WriteString(`<p class="markdown-alert-title">`)
		var kind string = ""

		if t, ok := node.AttributeString("kind"); ok {
			kind = strings.ToLower(t.(string))
			icon, ok := r.Icons[kind]
			if ok {
				w.WriteString(icon)
			}
			if _, ok := node.AttributeString("title"); ok {
				// do nothing
			} else {
				w.WriteString(caser.String(kind))
			}
		}
	} else {
		w.WriteString(`</p>`)
	}
	return gast.WalkContinue, nil
}
