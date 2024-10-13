# Chain Traverser

Chain Traverser is a service that indexes the Ethereum blockchain in memory and provides graph data through an API. It enables users to explore transaction relationships and find paths between addresses on the Ethereum network.

DEMO: <https://dictynna.com/>

## Features

- In-memory indexing of Ethereum blockchain data
- Graph traversal for address relationships
- Path finding between addresses
- Local Ethereum node integration support
- Redis-based caching for improved performance

## Setup

### Environment Variables

Mandatory environment variables are listed in the `.env_example` file. Additional configuration options can be found in `internal/config/config.go`.

### Running with a Local Ethereum Node

To use a local Ethereum node:

1. Ensure the node is available on host 0.0.0.0.
2. Host execution and consensus clients on your local machine (tested with geth and lighthouse).

If hosting locally is not possible, create an SSH tunnel:

```ssh
Host ALIAS
  HostName HOSTNAME
  IdentityFile KEY_PATH
  User USER
  LocalForward 0.0.0.0:8545 127.0.0.1:8545
```

Note: Keep the 0.0.0.0 host for container accessibility.

### Docker Compose Setup (Development Only)

Use `docker-compose-eth-local.yaml` to start the following services:

1. `redis`: Storage for indexed data
2. `indexer`: Ethereum blockchain data indexer
3. `price_indexer`: Cryptocurrency price data indexer
4. `api`: Request handling API service

To start the services:

```sh
docker compose -f docker-compose-eth-local.yaml up --build
```

### Running with an External Ethereum Node

Option 1: Custom Docker Compose

- Create a new file (e.g., `docker-compose-external-node.yaml`) with your ETH_NODE_URL variable.

Option 2: Manual Setup

- Set up and run services manually for more flexibility and control.

## API Endpoints

1. `GET /ping/`: Health check
2. `GET /orb/eth/{address}`: Fetch graph data for an Ethereum address
3. `GET /orb/eth/paths/{addressFrom}/to/{addressTo}`: Find paths between two Ethereum addresses (Experimental)

### Graph Data Endpoint Parameters

- `address`: Starting Ethereum address for graph traversal
- `depth` (query): Graph traversal depth (default: 1)
- `flow` (query): Transaction direction ("input", "output", "all"; default: "all")
- `fromBlock` (query): Starting block number (optional)
- `toBlock` (query): Ending block number (optional)
- `algo` (query): Traversal algorithm ("dfs", "bfs"; default: "dfs")
- `collapseTrxs` (query): Collapse multiple transactions between same addresses (default: true)

example

```sh
curl -XGET 'http://localhost:8080/orb/eth/0x5B282a9456ea00a63f9412B76B2d14775B9a9b48?depth=1&flow=all'
```

## Performance Considerations

For optimal performance:

1. Co-locate Ethereum node, Redis instance, and Golang processes, preferably on the same machine.
2. If using separate machines, ensure they are in the same data center or network to minimize latency.

## License

This project is licensed under the [GNU License](LICENSE).

## Support

For issues, feature requests, or questions, please [open an issue](https://github.com/guzun-corp/chain-traverser/issues) on our GitHub repository.
