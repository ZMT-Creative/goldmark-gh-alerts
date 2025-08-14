package constants

import gast "github.com/yuin/goldmark/ast"

// Node kinds for different alert components
var (
	// KindAlerts is the NodeKind for the alert block.
	KindAlerts = gast.NewNodeKind("Alerts")

	// KindAlertsHeader is the NodeKind for the alert header.
	KindAlertsHeader = gast.NewNodeKind("AlertsHeader")

	// KindAlertsBody is the NodeKind for the alert body.
	KindAlertsBody = gast.NewNodeKind("AlertsBody")
)

// IsNoIconKind returns true if the kind string indicates that no icon should be rendered.
func IsNoIconKind(kind string) bool {
	switch kind {
	case "noicon", "no-icon", "nil", "null":
		return true
	default:
		return false
	}
}
