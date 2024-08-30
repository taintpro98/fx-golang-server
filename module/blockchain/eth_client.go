package blockchain

import (
	"context"
	"fx-golang-server/config"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IEthClient interface {
	GetBalance(ctx context.Context, address string) (*big.Int, error)
}

type ethClient struct {
	cnf    config.BlockchainConfig
	client *ethclient.Client
}

func NewEthClient(cfg *config.Config) (IEthClient, error) {
	client, err := ethclient.Dial(cfg.Blockchain.Url)
	if err != nil {
		return nil, err
	}
	return &ethClient{
		cnf:    cfg.Blockchain,
		client: client,
	}, nil
}

func (v *ethClient) GetBalance(ctx context.Context, address string) (*big.Int, error) {
	addr := common.HexToAddress(address)
	balance, err := v.client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
