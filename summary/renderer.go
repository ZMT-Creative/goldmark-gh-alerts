package summary

import (
	"strings"

	"github.com/ZMT-Creative/goldmark-gh-alerts/kinds"
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
	titleCaser cases.Caser
}

func NewAlertsHeaderHTMLRendererWithIcons(icons Icons, opts ...html.Option) renderer.NodeRenderer {
    r := &AlertsHeaderHTMLRenderer{
        Config:    html.NewConfig(),
        Icons:     icons,
        titleCaser: cases.Title(language.English, cases.Compact),
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
	reg.Register(kinds.KindAlertsHeader, r.renderAlertsHeader)
}

// isNoIconKind returns true if the kind string indicates that no icon should be rendered.
func isNoIconKind(kind string) bool {
    switch kind {
    case "noicon", "no-icon", "nil", "null":
        return true
    default:
        return false
    }
}

func (r *AlertsHeaderHTMLRenderer) renderAlertsHeader(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
    if entering {
        w.WriteString(`<div class="gh-alert-title"><p>`)
        var kind string = ""

        if t, ok := node.AttributeString("kind"); ok {
            kind = strings.ToLower(t.(string))
            icon, ok := r.Icons[kind]
            if ok {
                w.WriteString(icon)
            // Check if the kind indicates no icon should be rendered.
            // if it's not a "no icon" kind, we can try to find a default icon.
            } else if !isNoIconKind(kind) {
                for _, v := range []string{"note", "info", "default"} {
                    icon, ok = r.Icons[v]
                    if ok {
                        w.WriteString(icon)
                        break
                    }
                }
            }
            if _, ok := node.AttributeString("title"); ok {
                // do nothing
            } else {
                w.WriteString(r.titleCaser.String(kind))
            }
        }
    } else {
        w.WriteString(`</p></div>`)
    }
    return gast.WalkContinue, nil
}
