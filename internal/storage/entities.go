package storage

type Labels struct {
	Prime    string    `json:"prime"`     // Coinbase 1, Binance 2, etc
	Type     string    `json:"type"`      // Exchange, DEX, etc
	Seconary *[]string `json:"secondary"` // other labels
}

type TxNumberStat struct {
	Total  int64
	Input  int64
	Output int64
}

type TxAmountStat struct {
	Total  int64
	Input  int64
	Output int64
}
