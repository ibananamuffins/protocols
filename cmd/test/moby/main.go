package main

import (
	"context"
	"flag"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols"
	"github.com/rs/zerolog"
)

func main() {
	// Create a zerolog logger
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// Command-line arguments
	lptAddressArg := flag.String("lptaddress", "", "Smart contract address")
	lptManagerAddressArg := flag.String("lptmanageraddress", "", "Smart contract address")
	rpcURLArg := flag.String("rpcurl", "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID", "Ethereum RPC URL")
	flag.Parse()

	// Validate required arguments
	missingArgs := []string{}
	if *lptAddressArg == "" {
		missingArgs = append(missingArgs, "lptaddress")
	}
	if *lptManagerAddressArg == "" {
		missingArgs = append(missingArgs, "lptmanageraddress")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -lptaddress <lpt-address> -lptmanageraddress <lptmanager-address> -rpcurl <rpc-url>").
			Msg("Missing required arguments")
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	cp := protocols.MobyLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *lptAddressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Str("lptaddress", *lptAddressArg).Msg("Failed to get config")
	}

	// Parse the smart contract address
	lptAddress := common.HexToAddress(*lptAddressArg)
	lptManagerAddress := common.HexToAddress(*lptManagerAddressArg)
	// Create a new MobyLPPriceProvider
	provider := protocols.NewMobyLPPriceProvider(lptAddress, lptManagerAddress, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize MobyLPPriceProvider")
	}

	// Fetch LP token price
	lpPrice, err := provider.LPTokenPrice(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch LP token price")
	} else {
		logger.Info().
			Str("LPTokenPrice (USD)", lpPrice).
			Msg("successfully fetched LP token price")
	}

	// Fetch TVL
	tvl, err := provider.TVL(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch TVL")
	} else {
		logger.Info().
			Str("TVL (USD)", tvl).
			Msg("successfully fetched TVL")
	}

}
