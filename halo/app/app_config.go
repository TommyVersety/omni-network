package app

import (
	attestmodule "github.com/omni-network/omni/halo/attest/module"
	attesttypes "github.com/omni-network/omni/halo/attest/types"
	engevmmodule "github.com/omni-network/omni/halo/evmengine/module"
	engevmtypes "github.com/omni-network/omni/halo/evmengine/types"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	distrmodulev1 "cosmossdk.io/api/cosmos/distribution/module/v1"
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	stakingmodulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	"cosmossdk.io/core/appconfig"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Bech32HRP is the human-readable-part of the Bech32 address format.
const Bech32HRP = "omni"

// init initializes the Cosmos SDK configuration.
//
//nolint:gochecknoinits // Cosmos-style
func init() {
	// Set prefixes
	accountPubKeyPrefix := Bech32HRP + "pub"
	validatorAddressPrefix := Bech32HRP + "valoper"
	validatorPubKeyPrefix := Bech32HRP + "valoperpub"
	consNodeAddressPrefix := Bech32HRP + "valcons"
	consNodePubKeyPrefix := Bech32HRP + "valconspub"

	// Set and seal config
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(Bech32HRP, accountPubKeyPrefix)
	cfg.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	cfg.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	cfg.Seal()
}

// DepConfig returns the default app depinject config.
func DepConfig() depinject.Config {
	return depinject.Configs(
		appConfig,
		depinject.Supply(),
	)
}

//nolint:gochecknoglobals // Cosmos-style
var (
	genesisModuleOrder = []string{
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		genutiltypes.ModuleName,
	}

	beginBlockers = []string{
		distrtypes.ModuleName,   // Note: slashing happens after distr.BeginBlocker
		stakingtypes.ModuleName, // Note: staking module is required if HistoricalEntries param > 0
	}

	endBlockers = []string{
		stakingtypes.ModuleName,
		attesttypes.ModuleName,
	}

	// blocked account addresses.
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
		stakingtypes.BondedPoolName,
		stakingtypes.NotBondedPoolName,
	}

	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: distrtypes.ModuleName},
		{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
	}

	// appConfig application configuration (used by depinject).
	appConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			{
				Name: runtime.ModuleName,
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName:       Name,
					BeginBlockers: beginBlockers,
					EndBlockers:   endBlockers,
					InitGenesis:   genesisModuleOrder,
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
					},
				}),
			},
			{
				Name: authtypes.ModuleName,
				Config: appconfig.WrapAny(&authmodulev1.Module{
					ModuleAccountPermissions: moduleAccPerms,
					Bech32Prefix:             Bech32HRP,
				}),
			},
			{
				Name: "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{
					SkipAnteHandler: true, // Disable ante handler (since we don't have proper txs).
					SkipPostHandler: true,
				}),
			},
			{
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},
			{
				Name:   distrtypes.ModuleName,
				Config: appconfig.WrapAny(&distrmodulev1.Module{}),
			},
			{
				Name:   genutiltypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
			},
			{
				Name:   stakingtypes.ModuleName,
				Config: appconfig.WrapAny(&stakingmodulev1.Module{}),
			},
			{
				Name:   engevmtypes.ModuleName,
				Config: appconfig.WrapAny(&engevmmodule.Module{}),
			},
			{
				Name:   attesttypes.ModuleName,
				Config: appconfig.WrapAny(&attestmodule.Module{}),
			},
		},
	})
)