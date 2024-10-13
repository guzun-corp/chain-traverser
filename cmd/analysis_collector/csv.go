package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func writeCsv(data map[string]Wallet) {
	now := time.Now().Format("20060102150405") // Format the current time as a string
	fname := fmt.Sprintf("eth_wallets_%s.csv", now)
	file, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Optional: Write header to CSV file
	if err := writer.Write([]string{"Address", "TxTotal", "TxIn", "TxOut", "TxTotalVolumeUsdOnDay", "TxInVolumeUsdOnDay", "TxOutVolumeUsdOnDay", "BalanceUsd"}); err != nil {
		panic(err)
	}

	for _, r := range data {
		if r.Address == "" {
			continue
		}
		record := []string{
			r.Address,
			fmt.Sprintf("%d", r.TxTotal),
			fmt.Sprintf("%d", r.TxIn),
			fmt.Sprintf("%d", r.TxOut),
			fmt.Sprintf("%f", r.TxTotalVolumeUsdOnDay),
			fmt.Sprintf("%f", r.TxInVolumeUsdOnDay),
			fmt.Sprintf("%f", r.TxOutVolumeUsdOnDay),
			fmt.Sprintf("%f", r.BalanceUsd),
		}
		if err := writer.Write(record); err != nil {
			panic(err) // handle error
		}
	}
}
