package alerts

// GhAlerts is a extension for the goldmark.

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"

	alertParser "github.com/ZMT-Creative/goldmark-gh-alerts/internal/parser"
	alertRenderer "github.com/ZMT-Creative/goldmark-gh-alerts/internal/renderer"
)

type GhAlerts struct {
	alertRenderer.Icons
}

// Meta is a extension for the goldmark.
var GhAlertsExtension = &GhAlerts{}

// Extend implements goldmark.Extender.
func (e *GhAlerts) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(alertParser.NewAlertsParser(), 799),
			util.Prioritized(alertParser.NewAlertsHeaderParser(), 799),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(alertRenderer.NewAlertsHTMLRenderer(), 0),
			util.Prioritized(alertRenderer.NewAlertsHeaderHTMLRendererWithIcons(e.Icons), 0),
			util.Prioritized(alertRenderer.NewAlertsBodyHTMLRenderer(), 0),
		),
	)
}
