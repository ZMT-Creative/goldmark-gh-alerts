package ast

import (
	"github.com/ZMT-Creative/goldmark-gh-alerts/internal/constants"
	gast "github.com/yuin/goldmark/ast"
)

// Alerts represents an alert block node
type Alerts struct {
	gast.BaseBlock
}

func (n *Alerts) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

func (n *Alerts) Kind() gast.NodeKind {
	return constants.KindAlerts
}

func NewAlerts() *Alerts {
	return &Alerts{}
}

// AlertsHeader represents an alert header node
type AlertsHeader struct {
	gast.BaseBlock
}

func (n *AlertsHeader) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

func (n *AlertsHeader) Kind() gast.NodeKind {
	return constants.KindAlertsHeader
}

func NewAlertsHeader() *AlertsHeader {
	return &AlertsHeader{}
}

// AlertsBody represents an alert body node
type AlertsBody struct {
	gast.BaseBlock
}

func (n *AlertsBody) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

func (n *AlertsBody) Kind() gast.NodeKind {
	return constants.KindAlertsBody
}

func NewAlertsBody() *AlertsBody {
	return &AlertsBody{}
}
