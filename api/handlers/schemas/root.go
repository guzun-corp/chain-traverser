package schemas

import (
	"strconv"

	"github.com/shopspring/decimal"
)

type Node struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Cnt    int64  `json:"cnt"`
	Picked bool   `json:"picked"`
	Type   string `json:"type"`
}

type Edge struct {
	Id             string                     `json:"id"`
	From           string                     `json:"start"`
	To             string                     `json:"end"`
	FlowByCurrency map[string]decimal.Decimal `json:"flow_by_currency"`
	TotalUsdFlow   decimal.Decimal            `json:"total_usd_flow"`
}

type CollapsedEdge struct {
	// contains all the transactions between two addresses
	Id             string                     `json:"id"`
	From           string                     `json:"start"`
	To             string                     `json:"end"`
	Count          int                        `json:"value"`
	FlowByCurrency map[string]decimal.Decimal `json:"flow_by_currency"`
	TotalUsdFlow   decimal.Decimal            `json:"total_usd_flow"`
}

type GraphCollapsed struct {
	Nodes []Node          `json:"nodes"`
	Edges []CollapsedEdge `json:"edges"`
}

type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

func CollapseTxs(txs *[]Edge) *[]CollapsedEdge {
	cnt := 0
	txsMap := make(map[string]CollapsedEdge)

	for _, tx := range *txs {
		key := tx.From + "-" + tx.To
		edge, exists := txsMap[key]
		if exists {
			edge.Count++
			edge.TotalUsdFlow = edge.TotalUsdFlow.Add(tx.TotalUsdFlow)
			for currency, amount := range tx.FlowByCurrency {
				edge.FlowByCurrency[currency] = edge.FlowByCurrency[currency].Add(amount)
			}
			txsMap[key] = edge
		} else {
			txsMap[key] = CollapsedEdge{
				From:           tx.From,
				To:             tx.To,
				Count:          1,
				FlowByCurrency: tx.FlowByCurrency,
				TotalUsdFlow:   tx.TotalUsdFlow,
				Id:             strconv.Itoa(cnt),
			}
			cnt++
		}
	}

	collapsedTxs := make([]CollapsedEdge, 0)

	for _, edge := range txsMap {
		collapsedTxs = append(collapsedTxs, edge)
	}

	return &collapsedTxs
}
