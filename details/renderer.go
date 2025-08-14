package details

import (
	"regexp"
	"strings"

	"github.com/ZMT-Creative/goldmark-gh-alerts/kinds"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"

	"fmt"
)

type AlertsHTMLRenderer struct {
	html.Config
}

func NewAlertsHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &AlertsHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

func (r *AlertsHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(kinds.KindAlerts, r.renderAlerts)
}

func (r *AlertsHTMLRenderer) renderAlerts(w util.BufWriter, source []byte, node gast.Node, entering bool) (gast.WalkStatus, error) {
    alertType := ""
	rawTitle := ""
    if t, ok := node.AttributeString("kind"); ok {
        if typeBytes, isBytes := t.([]uint8); isBytes {
            alertType = strings.ToLower(string(typeBytes))
        } else if typeStr, isStr := t.(string); isStr {
            alertType = strings.ToLower(typeStr)
        }
		// Check if the alertType is "noicon" or one of its variants
		if kinds.IsNoIconKind(alertType) {
			// If the alertType is "noicon", we want to use the "title" if it exists
			// If not, we can just use existing alertType as a fallback
			if rt, ok := node.AttributeString("title"); ok {
				if typeBytes, isBytes := rt.([]uint8); isBytes {
					rawTitle = strings.ToLower(string(typeBytes))
				} else if typeStr, isStr := rt.(string); isStr {
					rawTitle = strings.ToLower(typeStr)
				}
				// Create regular expressions for cleaning up the title for use in the HTML output
				reHTML := regexp.MustCompile(`<[^>]+>`)
				reWS := regexp.MustCompile(`\s+`)
				reMDInline := regexp.MustCompile(`\*\*|\*|__|_|~~|\\`)
				reMDLinks := regexp.MustCompile(`!?\[(.*?)\]\((.*?)\)`)
				reMDCode := regexp.MustCompile("`{1,3}[^`]*`{1,3}")
				// Clean up the raw title using the regular expressions
				rawTitle = reHTML.ReplaceAllString(rawTitle, "")
				rawTitle = reMDInline.ReplaceAllString(rawTitle, "")
				rawTitle = reMDLinks.ReplaceAllString(rawTitle, "")
				rawTitle = reMDCode.ReplaceAllString(rawTitle, "")
				rawTitle = strings.TrimSpace(rawTitle)
				rawTitle = reWS.ReplaceAllString(rawTitle, "-")
				// Set the alert type to the cleaned-up title
				alertType = rawTitle
			}
		}
    }

    if entering {
        fmt.Fprintf(w, `<div class="gh-alert gh-alert-%s" data-callout="%s">`, alertType, alertType)
    } else {
        w.WriteString("</div>\n")
    }
    return gast.WalkContinue, nil
}