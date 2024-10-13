package traverser

import "github.com/shopspring/decimal"

type Tx struct {
	From           string
	To             string
	TxHash         string
	FlowByCurrency map[string]decimal.Decimal
	TotalUsdFlow   decimal.Decimal
}

type Addr struct {
	Hash         string
	Cnt          int64
	NeedTraverse bool
}

type Graph struct {
	Addrs *map[string]Addr
	Txs   *map[string]Tx
}

const GRAPH_LIMIT = 500_000
