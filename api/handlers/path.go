package handlers

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/valyala/fasthttp"

	"chain-traverser/api/handlers/schemas"
	"chain-traverser/api/handlers/utils"
	"chain-traverser/internal/config"
	"chain-traverser/internal/storage/redis"
	"chain-traverser/internal/traverser"

	dominik "github.com/dominikbraun/graph"
)

type Params struct {
	FromHash  string
	ToHash    string
	FromBlock int
	ToBlock   int
}

func extractParams(c *fasthttp.RequestCtx) (Params, error) {
	var params Params

	fromHashRaw := c.UserValue("addressFrom")
	if fromHashRaw == nil {
		return params, errors.New("addressFrom invalid")
	}
	fromHash, ok := fromHashRaw.(string)
	if !ok {
		return params, errors.New("addressFrom required")
	}
	params.FromHash = fromHash

	toHashRaw := c.UserValue("addressTo")
	if toHashRaw == nil {
		return params, errors.New("addressTo required")
	}
	toHash, ok := toHashRaw.(string)
	if !ok {
		return params, errors.New("addressTo invalid")
	}
	params.ToHash = toHash

	fromBlockStr := string(c.QueryArgs().Peek("fromBlock"))
	if fromBlockStr != "" {
		fromBlock, err := strconv.Atoi(fromBlockStr)
		if err != nil {
			return params, errors.New("fromBlock invalid")
		}
		params.FromBlock = fromBlock
	}

	toBlockStr := string(c.QueryArgs().Peek("toBlock invalid"))
	if toBlockStr != "" {
		ToBlock, err := strconv.Atoi(toBlockStr)
		if err != nil {
			return params, errors.New("toBlock invalid")
		}
		params.ToBlock = ToBlock
	}
	if params.ToBlock == 0 {
		params.ToBlock = 99999999
	}

	return params, nil
}

func newDominikGraph(nodesMap map[string]bool, collapsedTrxs *[]schemas.CollapsedEdge) dominik.Graph[string, string] {
	g := dominik.New(dominik.StringHash)
	for n_hash := range nodesMap {
		g.AddVertex(n_hash)
	}
	for _, tx := range *collapsedTrxs {
		g.AddEdge(tx.From, tx.To)
	}
	return g
}

func fetchPathAddresses(paths [][]string, params Params, redis *redis.RedisClient) []schemas.Node {
	// fetch all nodes in the path, enrich with address-related data
	pathNodes := []schemas.Node{}
	if len(paths) == 0 {
		// if there is no path between two addresses, we just return these two addresses
		fromNode := utils.FetchAddress(params.FromHash, redis)
		fromNode.Picked = true
		toNode := utils.FetchAddress(params.ToHash, redis)
		toNode.Picked = true
		pathNodes = append(pathNodes, fromNode, toNode)
	} else {
		// if there is a path between two addresses, we return all nodes in the path
		for i := range paths {
			for _, pHash := range paths[i] {
				node := utils.FetchAddress(pHash, redis)
				node.Picked = params.ToHash == pHash || params.FromHash == pHash
				pathNodes = append(pathNodes, node)
			}
		}
	}
	return pathNodes
}

const PATH_GRAPH_LIMIT = 500_000
const PATH_GRAPH_DFS_MAX_DEPTH = 100

func CollectPathHandler(c *fasthttp.RequestCtx) {
	start := time.Now()

	params, vErr := extractParams(c)
	if vErr != nil {
		c.Error(vErr.Error(), fasthttp.StatusBadRequest)
		return
	}

	cfg, cfgErr := config.NewConfig()

	if cfgErr != nil {
		c.Error(cfgErr.Error(), fasthttp.StatusInternalServerError)
		return
	}
	redis := redis.NewClient(&cfg.Redis)

	var graph *traverser.Graph
	// dfsParams := traverser.ParamsDFS{
	// 	Address:        params.ToHash,
	// 	Depth:          PATH_GRAPH_DFS_MAX_DEPTH,
	// 	FromBlock:      params.FromBlock,
	// 	ToBlock:        params.ToBlock,
	// 	Flow:           "input",
	// 	GraphSizeLimit: PATH_GRAPH_LIMIT,
	// }
	dfsParams := traverser.ParamsDFS{
		Address:        params.FromHash,
		Depth:          PATH_GRAPH_DFS_MAX_DEPTH,
		FromBlock:      params.FromBlock,
		ToBlock:        params.ToBlock,
		Flow:           "output",
		GraphSizeLimit: PATH_GRAPH_LIMIT,
	}

	graph, err := traverser.CollectDFS(dfsParams, redis)
	if err != nil {
		c.Error("Error collecting graph dfs", fasthttp.StatusInternalServerError)
		return
	}

	edges := []schemas.Edge{}
	nodesMap := make(map[string]bool)
	// restore adresses from transactions
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
	collapsedTrxs := schemas.CollapseTxs(&edges)
	log.Info().Msgf("dfs collected %d nodes and %d edges (%d collapsed)", len(nodesMap), len(edges), len(*collapsedTrxs))
	g := newDominikGraph(nodesMap, collapsedTrxs)

	// find all paths between two addresses
	gSize, _ := g.Size()
	log.Info().Msgf("start paths collecting graph size = %d", gSize)

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	paths, err := utils.AllPathsBetween(g, params.FromHash, params.ToHash)
	if err != nil {
		log.Err(err).Msg("Error collecting paths")
	}
	log.Info().Msgf("all paths %s", paths)
	pathNodes := fetchPathAddresses(paths, params, redis)

	// paths edges are subset of all edges
	// We just return all edges in the graph
	data := schemas.GraphCollapsed{Nodes: pathNodes, Edges: *collapsedTrxs}
	jsonData, err := json.Marshal(data)
	if err != nil {
		c.Error("Error encoding JSON", fasthttp.StatusInternalServerError)
		return
	}
	c.Write(jsonData)

	c.SetContentType("application/json")
	c.Response.Header.Set("Access-Control-Allow-Origin", "*")
	c.Response.Header.Set("Access-Control-Allow-Methods", "GET")
	c.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type")
	c.Response.SetStatusCode(fasthttp.StatusOK)

	log.Info().Msgf("CollectPathHandler in %s", time.Since(start))
}
