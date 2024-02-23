package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Transaction struct {
	From   string
	To     string
	TxHash string
}

const LAST_BLOCK_FILE = "last_handled_block.txt"

func getNextBlockNumber() (*big.Int, error) {
	f, err := os.Open(LAST_BLOCK_FILE)
	blockNumber := 0
	if err != nil {
		if os.IsNotExist(err) {
			blockNumber = 19045008
		} else {
			// Other error occurred while opening the file
			return nil, err
		}
	}
	defer f.Close()

	if blockNumber == 0 {
		_, err = fmt.Fscanf(f, "%d", &blockNumber)
		if err != nil {
			return nil, err
		}
	}

	return big.NewInt(int64(blockNumber + 1)), nil
}

func setBlockNumber(blockNumber *big.Int) error {
	f, err := os.Create(LAST_BLOCK_FILE)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "%d", blockNumber)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	dbUser := ""
	dbPassword := ""
	dbUri := "bolt://localhost:7687" // scheme://host(:port) (default port is 7687)
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(dbUser, dbPassword, ""))
	ctx := context.Background()
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Viola! Connected to Memgraph!")
	}

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Ethereum node")

	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// fmt.Println("head: %+v\n", header.Number.String())
	prev_time := time.Now()
	i := 0
	for {
		select {
		case <-stop:
			// Handle interrupt, then return to terminate the goroutine
			fmt.Println("Interrupt received, terminating...")
			return
		default:

			blockNumber, err := getNextBlockNumber()

			// fmt.Printf("start block: %s \n", blockNumber)

			block, err := client.BlockByNumber(ctx, blockNumber)
			if err != nil {
				fmt.Printf("fetch block by number error: %s\n", err)
				time.Sleep(5 * time.Second)
				continue
			}
			// fmt.Println("got block: %s\n", block.Time())
			if i%100 == 0 {

				bloksPerSec := (time.Since(prev_time).Seconds()) / 100

				fmt.Println("%d | Blocks %d | p/s: %f", time.Now(), block.Number(), bloksPerSec)

				prev_time = time.Now()
			}

			blockAddresses := []string{}
			transactions := []Transaction{}

			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

			for _, tx := range block.Transactions() {
				if tx.To() == nil {
					fmt.Printf("to is nil, skip; tx: %s\n", tx.Hash().Hex())
					continue
				}
				// fmt.Println("To: %s\n", tx.To().Hex())
				// fmt.Println("Tx: %s\n", tx.Hash().Hex())

				if from, err := types.Sender(types.NewLondonSigner(big.NewInt(1)), tx); err == nil {
					// dbgHash := "0xd2b54c3babae07614c2263a10c32d820c3bb1de4594e559c9f73f46320cf11a2"

					// if tx.Hash().Hex() == dbgHash {
					// 	fmt.Printf("Tx: %s\n", tx.Hash().Hex())
					// 	fmt.Printf("Tx2: %s\n", tx.Hash())
					// 	fmt.Printf("From: %s\n", from.Hex())
					// 	fmt.Printf("to: %s\n", tx.To().Hex())
					// 	fmt.Printf("to: %s\n", common.HexToAddress(tx.To().Hex()))
					// 	blockAddresses = append(blockAddresses, tx.To().Hex())
					// 	blockAddresses = append(blockAddresses, from.Hex())
					// 	fmt.Println("addresses: %s\n", blockAddresses)
					// 	transactions = append(transactions, Transaction{
					// 		From:   from.Hex(),
					// 		To:     tx.To().Hex(),
					// 		TxHash: tx.Hash().Hex(),
					// 	})

					// }
					blockAddresses = append(blockAddresses, tx.To().Hex())
					blockAddresses = append(blockAddresses, from.Hex())

					transactions = append(transactions, Transaction{
						From:   from.Hex(),
						To:     tx.To().Hex(),
						TxHash: tx.Hash().Hex(),
					})

				}
			}
			// fmt.Printf("Addresses before deduplication: %d\n", len(blockAddresses))

			deduplicate(blockAddresses)

			start := time.Now()

			_, err = neo4j.ExecuteQuery(ctx, driver, `
					WITH $hashes AS batch
					UNWIND batch AS value
					MERGE (n:Address {hash: value})
				`, map[string]interface{}{
				"hashes": blockAddresses,
			}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(""))

			elapsed := time.Since(start)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Addresses inserted; execution time: %s\n", elapsed)

			// start = time.Now()

			data := "from_hash,to_hash,txhash,timestamp\n"
			for _, tx := range transactions {
				data += fmt.Sprintf("%s,%s,%s,%d\n", tx.From, tx.To, tx.TxHash, block.Time())
			}

			// Write the query to a file
			queryFile, err := os.Create("/var/log/memgraph/query.csv")
			if err != nil {
				log.Fatal(err)
			}
			defer queryFile.Close()

			_, err = queryFile.WriteString(data)
			if err != nil {
				log.Fatal(err)
			}

			query1 := `LOAD CSV FROM "/var/log/memgraph/query.csv" WITH HEADER AS row
			MATCH (a:Address {hash: row.from_hash}), (b:Address {hash: row.to_hash})
			CREATE (a)-[:TX{txhash:row.txhash, timestamp: row.timestamp}]->(b);`

			_, err = neo4j.ExecuteQuery(ctx, driver, query1, map[string]interface{}{}, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase(""))
			if err != nil {
				panic(err)
			}
			setBlockNumber(block.Number())
			i += 1
			// elapsed = time.Since(start)

			// time.Sleep(3 * time.Second)

			// fmt.Printf("Transactions inserted; execution time: %s\n", elapsed)
			// fmt.Printf("finish block: %s \n", blockNumber)
		}
	}
}

func deduplicate(blockAddresses []string) {
	uniqueAddresses := make(map[string]bool)

	uniqueBlockAddresses := make([]string, 0)

	for _, address := range blockAddresses {

		if !uniqueAddresses[address] {

			uniqueAddresses[address] = true
			uniqueBlockAddresses = append(uniqueBlockAddresses, address)
		}
	}

	blockAddresses = uniqueBlockAddresses
}

// MATCH (n)
// DETACH DELETE n

// DROP INDEX ON :Address
// FREE MEMORY
// STORAGE MODE ON_DISK_TRANSACTIONAL
// MATCH (a:Address {hash: '0xc72a27e9eb15401b90d257d0180ca6d0646fd650'}),(b:Address {hash: '0x883fb7864c8998d2a7bd8ff4b417d73c22cf8d2e'}) CREATE (a)-[r:TX{txhash: '0xe2ea2692ada7401394845413f0f96d3d5ec62101d03e73791a8229928891d6eb'}]->(b);
// MATCH (a:Address), (b:Address)
// WHERE a.hash = '0x03e5bef0e28d46bd2f1c3427050e9052be1b8b18' AND b.hash = '0x352e504813b9e0b30f9ca70efc27a52d298f6697'
// CREATE (a)-[r:RELATIONSHIP_TYPE{txhash:'0x052dba88175297910b7d54cdff44e0dd4bb35989d41ee7f4a21e36cae895a499'}]->(b)
// RETURN a, b, r;

// MATCH (n:Address {hash:'0x0b0409A70A279dBa59f07ebe4345AaFC9C64154a'})
// RETURN n;

// "from": "0x0b0409a70a279dba59f07ebe4345aafc9c64154a",
// "hash": "0xd2b54c3babae07614c2263a10c32d820c3bb1de4594e559c9f73f46320cf11a2",
// "to": "0xd9e1ce17f2641f24ae83637ab66a2cca9c378b9f",

// MATCH (a:Address {hash: '0x0b0409A70A279dBa59f07ebe4345AaFC9C64154a'}),
// (b:Address {hash: '0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F'})
// CREATE (a)-[:TX{txhash: '0xd2b54c3babae07614c2263a10c32d820c3bb1de4594e559c9f73f46320cf11a2'}]->(b);

// MATCH (p1),(p2)
// WHERE p1.hash = '0x0b0409a70a279dba59f07ebe4345aafc9c64154a' AND p2.hash = '0xd9e1ce17f2641f24ae83637ab66a2cca9c378b9f'
// CREATE (p1)-[r:TX{txhash: '0xd2b54c3babae07614c2263a10c32d820c3bb1de4594e559c9f73f46320cf11a2'}]->(p2);

// "blockHash": "0x13efda093a343af7a22ffbd57128b5b1343cd63af1353f0f13bc9d0c605448f5",
// "blockNumber": "0xb628d0",
// "from": "0x0b0409a70a279dba59f07ebe4345aafc9c64154a",
// "gas": "0x253f0",
// "gasPrice": "0x1caf4ad54e",
// "hash": "0xd2b54c3babae07614c2263a10c32d820c3bb1de4594e559c9f73f46320cf11a2",
// "input": "0x7ff36ab500000000000000000000000000000000000000000000000013ea14d41802d02a00000000000000000000000000000000000000000000000000000000000000800000000000000000000000000b0409a70a279dba59f07ebe4345aafc9c64154a000000000000000000000000000000000000000000000000000000006039f90b0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000dbdb4d16eda451d0503b854cf79d55697f90c8df",
// "nonce": "0x17d4",
// "to": "0xd9e1ce17f2641f24ae83637ab66a2cca9c378b9f",
// "transactionIndex": "0x4c",
// "value": "0x58d15e176280000",
// "type": "0x0",
// "chainId": "0x1",
// "v": "0x26",
// "r": "0xdfaaabe4b5a4168b3428d5d3bf0ca7cc7d0d70f1a8a96c5ec8f67b2afcc0b43e",
// "s": "0x340bd9f454d87c647609d34a1b4e0affdffa1d4c3c6008d007e5c9ce1e4cfab5"
