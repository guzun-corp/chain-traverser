package eth

import (
	"chain-traverser/internal/blockchain/eth/erc20"
	"chain-traverser/internal/config"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

type ERC20Transaction struct {
	To     string
	Value  decimal.Decimal
	Ticker string
}

type Erc20Token struct {
	Instance     *erc20.Erc20
	Ticker       string
	Contract     string
	Denomination int
}

type Erc20Cache map[string]Erc20Token

func initErc20Token(contractAddr string, client *ethclient.Client) (*Erc20Token, error) {
	tokenAddress := common.HexToAddress(contractAddr)
	instance, err := erc20.NewErc20(tokenAddress, client)
	if err != nil {
		log.Err(err).Msg("failed to init token instance")
		return nil, err
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Err(err).Msg("failed to fetch decimals")
		return nil, err
	}
	denomination := decimals.Uint64() // Convert *big.Int to uint8
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Err(err).Msg("failed to fetch symbol")
		return nil, err
	}
	log.Info().Msgf("instance %s uploaded successfully", symbol)
	return &Erc20Token{Instance: instance, Ticker: symbol, Contract: contractAddr, Denomination: int(denomination)}, nil

}

func newErc20Cache(client *ethclient.Client) Erc20Cache {
	cache := make(Erc20Cache)
	for _, contract := range CONTRACTS_TO_TRACK {
		token, err := initErc20Token(contract, client)
		if err != nil {
			log.Err(err).Msg("failed to init token")
			continue
		}
		validContract := common.HexToAddress(contract).Hex()
		cache[validContract] = *token
	}
	return cache
}

type ContractProps struct {
	Ticker       string
	Contract     string
	Denomination int
	RawABI       string
	ABI          *abi.ABI
}

type EthClient struct {
	Client       *ethclient.Client
	tokenCache   map[string]Erc20Token
	addrByTicker map[string]string
	// we expect all erc20 tokens to have the same abi for functions we care about
	// use it for decoding transaction input (to address and value)
	erc20abi *abi.ABI
}

func initErc20abi() (*abi.ABI, error) {
	abi, err := abi.JSON(strings.NewReader(USDT_ABI))
	if err != nil {
		log.Err(err).Msg("failed to parse abi")
		return nil, err
	}
	return &abi, nil
}

func NewEthClient(cfg *config.EthConfig) (*EthClient, error) {
	client, err := ethclient.Dial(cfg.NodeUrl)
	if err != nil {
		log.Err(err).Msg("error connecting to Ethereum node")
		return nil, err
	}

	tokenCache := newErc20Cache(client)
	var addrByTicker = make(map[string]string)
	for k, v := range tokenCache {
		addrByTicker[v.Ticker] = k
	}

	abi, err := initErc20abi()
	if err != nil {
		log.Err(err).Msg("error uploading erc20 abi")
		return nil, err
	}

	return &EthClient{Client: client, tokenCache: tokenCache, erc20abi: abi, addrByTicker: addrByTicker}, nil
}

func (c *EthClient) GetToken(contractAddr string) *Erc20Token {
	token, ok := c.tokenCache[contractAddr]
	if !ok {
		return nil
	}
	return &token
}

func (c *EthClient) GetTokenByTicker(ticker string) *Erc20Token {
	addr, ok := c.addrByTicker[ticker]
	if !ok {
		return nil
	}
	token := c.GetToken(addr)
	return token
}

func (c *EthClient) HandleERC20(tx types.Transaction) (*ERC20Transaction, error) {
	// abi decider https://bia.is/tools/abi-decoder/
	// transactions for test
	// 0xe4ffe0f5426c1143880b72e37ce585aed04ebdc6176aa5dfcf88a6839220bfb0 p2p 53.29 USDT
	// {
	// 	"name": "transfer",
	// 	"params": [
	// 	  {
	// 		"name": "_to",
	// 		"value": "0xd268231fc42e0c3dfdc5ece98b1039f9f2fafeee",
	// 		"type": "address"
	// 	  },
	// 	  {
	// 		"name": "_value",
	// 		"value": "53294745",
	// 		"type": "uint256"
	// 	  }
	// 	]
	//   }

	token := c.GetToken(tx.To().Hex())
	if token == nil {
		err := fmt.Errorf("no token %s", tx.To().Hex())
		return nil, err
	}
	method, err := c.erc20abi.MethodById(tx.Data())

	if err != nil {
		return nil, err
	}
	if method.Name != "transfer" && method.Name != "transferFrom" {
		if method.Name != "approve" {
			log.Warn().Msgf("unexpected method tx: %s, method: %s", tx.Hash().Hex(), method.Name)
		}
		return nil, errors.New("wrong method")
	}
	if len(tx.Data()) < 4 {
		log.Warn().Msgf("No data in tx %s", tx.Hash().Hex())
		return nil, errors.New("no data in tx")
	}
	results, err := method.Inputs.Unpack(tx.Data()[4:])
	if err != nil {
		txhash := tx.Hash()
		log.Err(err).Msgf("failed to unpack tx:%s", txhash)
		return nil, err
	}

	var to common.Address
	var amount *big.Int

	if len(results) == 3 {
		to = results[1].(common.Address)
		amount = results[2].(*big.Int)
	} else if len(results) == 2 {
		to = results[0].(common.Address)
		amount = results[1].(*big.Int)
	}

	floatAmount, _ := amount.Float64()

	adjustedAmount := floatAmount / math.Pow10(int(token.Denomination))
	amountString := decimal.NewFromFloat(adjustedAmount)

	return &ERC20Transaction{To: to.Hex(), Value: amountString, Ticker: token.Ticker}, nil
}
