package eth

import (
	"chain-traverser/internal/storage/redis"
	"fmt"

	"github.com/shopspring/decimal"
)

const (
	USD                  = "USD"
	ETH                  = "ETH"
	BTC                  = "BTC"
	MAX_PRICE_DEPTH_DAYS = 3
)

var DECIMAL_ONE = decimal.NewFromInt(1)
var CURRENCIES_REMAP = map[string]string{
	"wBTC":   BTC,
	"stETH":  ETH,
	"eETH":   ETH,
	"weETH":  ETH,
	"rETH":   ETH,
	"mETH":   ETH,
	"rsETH":  ETH,
	"cbETH":  ETH,
	"frxETH": ETH,
	"USDT":   USD,
	"USDC":   USD,
	"DAI":    USD,
	"sUSDe":  USD,
	"USDe":   USD,
	"TUSD":   USD,
	"FDUSD":  USD,
	"FRAX":   USD,
	"wMANA":  "MANA",
	"wCELO":  "CELO",
}
var CURRENCIES = []string{
	BTC,
	ETH,
	"BNB",
	"TONCOIN",
	"SHIB",
	"WBTC",
	"TRX",
	"LINK",
	"NEAR",
	"LEO",
	"UNI",
	"MNT",
	"OKB",
	"CRO",
	"RNDR",
	"ARB",
	"VEN",
	"IMX",
	"GRT",
	"INJ",
	"PEPE",
	"FET",
	"THETA",
	"FTM",
	"LDO",
	"BGB",
	"QNT",
	"BEAM",
	"ENA",
	"FLOKI",
	"WBT",
	"stkAAVE",
	"BTT",
	"ONDO",
	"RBN",
	"AGIX",
	"AXS",
	"SAND",
	"SNX",
	"CHZ",
	"WLD",
	"GNO",
	"KCS",
	"MANA",
	"PRIME",
	"APE",
	"AIOZ",
	"NEXO",
	"DEXE",
	"AXL",
	"Cake",
	"DYDX",
	"FRAX",
	"OM",
	"ILV",
	"BLUR",
	"XAUt",
	"PENDLE",
	"MX",
	"WOO",
	"OCEAN",
	"CRV",
	"IOTX",
	"ALT",
	"SKL",
	"ENJ",
	"NFT",
	"1INCH",
	"GMT",
	"PAXG",
	"ZIL",
	"ENS",
	"RPL",
	"WQTUM",
	"ZRX",
	"MEME",
	"CELO",
	"ELF",
	"HOT",
	"AMP",
}

// if db does not have price for the block, try to get it for the previous 10 days
func GetTokenPrice(blockTime uint64, redis *redis.RedisClient, currency string) (*decimal.Decimal, error) {
	c := remapCurrency(currency)
	if c == USD {
		return &DECIMAL_ONE, nil
	}

	var err error

	for i := 1; i < MAX_PRICE_DEPTH_DAYS; i++ {
		blockTime := normalizedBlockTime(blockTime, i)
		price, pErr := redis.GetPrice(blockTime, c)
		if price != nil {
			return price, nil
		}
		err = pErr
	}
	return nil, fmt.Errorf("price not found for %s: %w", c, err)
}

func remapCurrency(currency string) string {
	c, remapped := CURRENCIES_REMAP[currency]
	if !remapped {
		c = currency
	}
	return c
}

func IsCurrency(currency string) bool {
	if currency == USD {
		return true
	}
	_, ok := CURRENCIES_REMAP[currency]
	return ok
}
