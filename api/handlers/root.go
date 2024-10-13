package handlers

import (
	"encoding/json"
	"slices"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/valyala/fasthttp"

	"chain-traverser/api/handlers/schemas"
	"chain-traverser/api/handlers/utils"
	"chain-traverser/internal/config"
	"chain-traverser/internal/storage/redis"
	"chain-traverser/internal/traverser"
)

func CollectGraphHandler(c *fasthttp.RequestCtx) {
	start := time.Now()

	targetHash := c.UserValue("address")
	if targetHash == nil {
		c.Error("Bad Request", fasthttp.StatusBadRequest)
		return
	}

	depthStr := string(c.QueryArgs().Peek("depth"))
	depth, err := strconv.Atoi(depthStr)
	if err != nil {
		c.Error("Invalid depth parameter", fasthttp.StatusBadRequest)
		return
	}

	collapseTrxsStr := string(c.QueryArgs().Peek("collapseTrxs"))
	collapseTrxs := true
	if collapseTrxsStr == "false" {
		collapseTrxs = false
	}

	fromBlockStr := string(c.QueryArgs().Peek("fromBlock"))
	var fromBlock int
	if fromBlockStr == "" {
		fromBlock = 0
	} else {
		fromBlock, err = strconv.Atoi(fromBlockStr)
		if err != nil {
			c.Error("Invalid fromBlock parameter", fasthttp.StatusBadRequest)
			return
		}
	}

	toBlockStr := string(c.QueryArgs().Peek("toBlock"))
	var toBlock int
	if toBlockStr == "" {
		toBlock = 99999999
	} else {
		toBlock, err = strconv.Atoi(toBlockStr)
		if err != nil {
			c.Error("Invalid toBlock parameter", fasthttp.StatusBadRequest)
			return
		}
	}

	algoStr := string(c.QueryArgs().Peek("algo"))
	log.Info().Msgf("algoStr: %s", algoStr)
	var algo string
	if algoStr == "bfs" {
		algo = "bfs"
	} else {
		algo = "dfs"
	}

	flowOptions := []string{"input", "output", "all"}
	flowStr := string(c.QueryArgs().Peek("flow"))
	log.Info().Msgf("flowStr: %s", algoStr)
	var flow string
	if slices.Contains(flowOptions, flowStr) {
		flow = flowStr
	} else {
		c.Error("Invalid flow parameter", fasthttp.StatusBadRequest)
		return
	}

	var graph *traverser.Graph

	cfg, err := config.NewConfig()

	if err != nil {
		log.Err(err).Msg("error reading config")
		return
	}
	redis := redis.NewClient(&cfg.Redis)

	if algo == "dfs" {
		dfsParams := traverser.ParamsDFS{
			Address:        targetHash.(string),
			Depth:          depth,
			FromBlock:      fromBlock,
			ToBlock:        toBlock,
			Flow:           flow,
			GraphSizeLimit: 5_000,
		}

		graph, err = traverser.CollectDFS(dfsParams, redis)
		if err != nil {
			c.Error("Error collecting graph dfs", fasthttp.StatusInternalServerError)
			return
		}
	} else {
		graph, err = traverser.CollectBFS(targetHash.(string), depth, fromBlock, toBlock, redis)
		if err != nil {
			c.Error("Error collecting graph", fasthttp.StatusInternalServerError)
			return
		}

	}

	nodes := []schemas.Node{}
	edges := []schemas.Edge{}
	nodesMap := make(map[string]bool)

	for _, tx := range *graph.Txs {
		_, exists := (nodesMap)[tx.From]
		if !exists {
			nodesMap[tx.From] = true
		}
		_, exists = (nodesMap)[tx.To]
		if !exists {
			nodesMap[tx.To] = true
		}
		edges = append(edges, schemas.Edge{From: tx.From, To: tx.To, Id: tx.TxHash, FlowByCurrency: tx.FlowByCurrency, TotalUsdFlow: tx.TotalUsdFlow})
	}

	for n_hash := range nodesMap {
		node := utils.FetchAddress(n_hash, redis)
		node.Picked = targetHash == n_hash
		nodes = append(nodes, node)
	}

	if collapseTrxs {
		collapsedTrxs := schemas.CollapseTxs(&edges)

		log.Info().Msgf("got %d collapsed transactions in %s", len(*collapsedTrxs), time.Since(start))
		data := schemas.GraphCollapsed{Nodes: nodes, Edges: *collapsedTrxs}
		jsonData, err := json.Marshal(data)
		if err != nil {
			c.Error("Error encoding JSON", fasthttp.StatusInternalServerError)
			return
		}
		c.Write(jsonData)

	} else {
		data := schemas.Graph{Nodes: nodes, Edges: edges}
		jsonData, err := json.Marshal(data)
		if err != nil {
			c.Error("Error encoding JSON", fasthttp.StatusInternalServerError)
			return
		}
		c.Write(jsonData)

	}

	c.SetContentType("application/json")
	c.Response.Header.Set("Access-Control-Allow-Origin", "*")
	c.Response.Header.Set("Access-Control-Allow-Methods", "GET")
	c.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type")
	c.Response.SetStatusCode(fasthttp.StatusOK)
}
