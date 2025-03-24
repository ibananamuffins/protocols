package protocols

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

type MobyConfig struct {
	LPTDecimals uint   `json:"lpt_decimals"`
}

// MobyLPPriceProvider defines the provider for Moby LP price and TVL.
type MobyLPPriceProvider struct {
	lptAddress     common.Address
	lptManagerAddress common.Address
	logger      zerolog.Logger
	configBytes []byte
	config      *MobyConfig
	lptContract *sc.MobyOLP
	lptManagerContract *sc.MobyOLPManager
}

// NewMobyLPPriceProvider creates a new instance of the MobyLPPriceProvider.
func NewMobyLPPriceProvider(lptAddress common.Address, lptManagerAddress common.Address, logger zerolog.Logger, config []byte) *MobyLPPriceProvider {
	m := &MobyLPPriceProvider{
		lptAddress:     lptAddress,
		lptManagerAddress: lptManagerAddress,
		logger:      logger,
		configBytes: config,
	}
	return m
}

// Initialize checks the configuration/data provided and instantiates the Moby smart contract.
func (m *MobyLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	m.config = &MobyConfig{}
	err = json.Unmarshal(m.configBytes, m.config)
	if err != nil {
		m.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	m.lptContract, err = sc.NewMobyOLP(m.lptAddress, client)
	if err != nil {
		m.logger.Error().Err(err).Msg("failed to instatiate Moby OLP contract")
		return err
	}

	m.lptManagerContract, err = sc.NewMobyOLPManager(m.lptManagerAddress, client)
	if err != nil {
		m.logger.Error().Err(err).Msg("failed to instatiate Moby OLP Manager contract")
		return err
	}

	return nil
}

func (m *MobyLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	// Get price from MobyOLPManager (30 decimals)
	price, err := m.lptManagerContract.GetPrice(opts, true) // true for _maximise parameter
	if err != nil {
		m.logger.Error().Err(err).Msg("failed to fetch price from MobyOLPManager")
		return "", err
	}

	// Convert price to decimal with 30 decimals
	priceDecimal := decimal.NewFromBigInt(price, -30) // 30 decimals

	m.logger.Info().
		Str("price", priceDecimal.String()).
		Msg("LP token price fetched successfully")

	return priceDecimal.StringFixed(roundingDecimals), nil
}

func (m *MobyLPPriceProvider) TVL(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	// Get AUM from MobyOLPManager (30 decimals)
	aum, err := m.lptManagerContract.GetAum(opts, true) // false for _maximise parameter
	if err != nil {
		m.logger.Error().Err(err).Msg("failed to fetch AUM from MobyOLPManager")
		return "", err
	}

	// Convert AUM to decimal with 30 decimals
	aumDecimal := decimal.NewFromBigInt(aum, -30) // 30 decimals

	m.logger.Info().
		Str("tvl", aumDecimal.String()).
		Msg("successfully fetched TVL")

	return aumDecimal.StringFixed(roundingDecimals), nil
}

func (m *MobyLPPriceProvider) GetConfig(ctx context.Context, lptAddress string, ethClient *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(lptAddress) {
		err = fmt.Errorf("invalid smart contract address, '%s'", lptAddress)
		return nil, err
	}

	contract, err := sc.NewMobyOLP(common.HexToAddress(lptAddress), ethClient)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Moby OLP smart contract, %v", err)
		return nil, err
	}

	mc := &MobyConfig{}
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch decimals, %v", err)
		return nil, err
	}
	mc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(mc)
	if err != nil {
		return nil, err
	}

	return body, nil
}