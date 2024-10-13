package main

import (
	"fmt"
	"math"

	"github.com/rs/zerolog/log"
)

// calculateMean computes the mean of a slice of float64
func calculateMean(data []float64) float64 {
	var sum float64
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// calculateStdDev computes the standard deviation of a slice of float64
func calculateStdDev(data []float64, mean float64) float64 {
	var sum float64
	for _, value := range data {
		sum += math.Pow(value-mean, 2)
	}
	variance := sum / float64(len(data))
	return math.Sqrt(variance)
}

// normalizeData normalizes the data using the 3-sigma method and identifies outliers
func normalizeData(data []float64) ([]float64, []int) {
	mean := calculateMean(data)
	stdDev := calculateStdDev(data, mean)
	normalized := make([]float64, 0)
	removedIndexes := make([]int, 0)

	for i, value := range data {
		zScore := (value - mean) / stdDev
		if math.Abs(zScore) > 3 {
			// This value is an outlier, more than 3 standard deviations from the mean
			removedIndexes = append(removedIndexes, i)
		} else {
			normalized = append(normalized, zScore)
		}
	}
	return normalized, removedIndexes
}

func do() {
	data := []float64{10, 12, 23, 23, 16, 23, 21, 100, 999999}
	normalizedData, removedIndexes := normalizeData(data)
	fmt.Println("Normalized Data:", normalizedData)
	fmt.Println("Removed Indexes:", removedIndexes)
}

func normValueFloat(value, max float64) float64 {
	return value / max
}
func normValueInt(value, max int64) float64 {
	return float64(value) / float64(max)
}

func normalizeWalletsOld(data FetchingResult) []NormalizedWallet {
	do()
	nWallets := make([]NormalizedWallet, len(data.Wallets))

	for _, w := range data.Wallets {
		if w.Address == "" {
			continue
		}
		nw := NormalizedWallet{
			Address:               w.Address,
			TxTotal:               normValueInt(w.TxTotal, data.Ref.MaxTxTotal),
			TxIn:                  normValueInt(w.TxIn, data.Ref.MaxTxIn),
			TxOut:                 normValueInt(w.TxOut, data.Ref.MaxTxOut),
			TxTotalVolumeUsdOnDay: normValueFloat(w.TxTotalVolumeUsdOnDay, data.Ref.MaxTxTotalVolumeUsdOnDay),
			TxInVolumeUsdOnDay:    normValueFloat(w.TxInVolumeUsdOnDay, data.Ref.MaxTxInVolumeUsdOnDay),
			TxOutVolumeUsdOnDay:   normValueFloat(w.TxOutVolumeUsdOnDay, data.Ref.MaxTxOutVolumeUsdOnDay),
			BalanceUsd:            normValueFloat(w.TxOutVolumeUsdOnDay, data.Ref.MaxTxOutVolumeUsdOnDay),
		}
		// skip edge cases
		// if nw.TxTotal > 0.41 {
		// 	continue
		// }

		// if nw.TxTotalVolumeUsdOnDay > 0.07 {
		// 	continue
		// }

		if nw.TxTotal > 1 || nw.TxTotal < 0 || nw.TxTotalVolumeUsdOnDay > 1 || nw.TxTotalVolumeUsdOnDay < 0 {
			log.Info().Msgf("wallet: %+v", w)
			log.Info().Msgf("wallet: %+v", nw)

		} else {
			//log.Debug().Msgf("wallet: %+v", w)
			nWallets = append(nWallets, nw)
		}

	}
	return nWallets
}
