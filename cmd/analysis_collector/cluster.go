package main

import (
	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
	"github.com/muesli/kmeans/plotter"
)

func createClasters(data []NormalizedWallet) {
	var d clusters.Observations
	for _, w := range data {
		d = append(d, clusters.Coordinates{
			// w.BalanceUsd,
			w.TxTotal,
			// w.TxIn,
			// w.TxOut,
			w.BalanceUsd,
			// w.TxInVolumeUsdOnDay,
			// w.TxOutVolumeUsdOnDay,
		})
	}

	// Partition the data points into 16 clusters
	// km := kmeans.New()
	// km, _ := kmeans.NewWithOptions(0.05, nil)
	km, _ := kmeans.NewWithOptions(0.01, plotter.SimplePlotter{})

	km.Partition(d, 4)

	// if err != nil {
	// 	log.Err(err).Msg("error creating clusters")
	// 	return
	// }
	// for _, c := range clusters {
	// 	fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0], c.Center[1])
	// 	fmt.Printf("Matching data points: %+v\n\n", c.Observations)
	// }
}
