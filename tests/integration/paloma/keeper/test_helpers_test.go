package keeper_test

import (
	"context"
	"fmt"
	"strings"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/upgrade"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	db "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/version"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil/integration"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/onsi/ginkgo/v2"
	params2 "github.com/palomachain/paloma/app/params"
	xchain "github.com/palomachain/paloma/internal/x-chain"
	"github.com/palomachain/paloma/x/consensus"
	consensusmodulekeeper "github.com/palomachain/paloma/x/consensus/keeper"
	consensusmoduletypes "github.com/palomachain/paloma/x/consensus/types"
	"github.com/palomachain/paloma/x/evm"
	evmmodulekeeper "github.com/palomachain/paloma/x/evm/keeper"
	evmmoduletypes "github.com/palomachain/paloma/x/evm/types"
	gravitymoduletypes "github.com/palomachain/paloma/x/gravity/types"
	"github.com/palomachain/paloma/x/paloma"
	palomakeeper "github.com/palomachain/paloma/x/paloma/keeper"
	palomamoduletypes "github.com/palomachain/paloma/x/paloma/types"
	"github.com/palomachain/paloma/x/scheduler"
	"github.com/palomachain/paloma/x/scheduler/keeper"
	schedulertypes "github.com/palomachain/paloma/x/scheduler/types"
	treasurymoduletypes "github.com/palomachain/paloma/x/treasury/types"
	"github.com/palomachain/paloma/x/valset"
	valsetmodulekeeper "github.com/palomachain/paloma/x/valset/keeper"
	valsetmoduletypes "github.com/palomachain/paloma/x/valset/types"
	"github.com/spf13/cast"
)

const (
	minimumPigeonVersion = "v1.10.0"
)

var (
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		gravitymoduletypes.ModuleName:  {authtypes.Minter, authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		ibcfeetypes.ModuleName:         nil,
		icatypes.ModuleName:            nil,
		wasmtypes.ModuleName:           {authtypes.Burner},
		treasurymoduletypes.ModuleName: {authtypes.Burner, authtypes.Minter},
	}
)

type fixture struct {
	ctx               sdk.Context
	codec             codec.Codec
	queryClient       v1.QueryClient
	legacyQueryClient v1beta1.QueryClient
	consensusKeeper   consensusmodulekeeper.Keeper
	evmKeeper         evmmodulekeeper.Keeper
	valsetKeeper      valsetmodulekeeper.Keeper
	schedulerKeeper   schedulertypes.Keeper
	stakingKeeper     stakingkeeper.Keeper
	paramsKeeper      paramskeeper.Keeper
	palomaKeeper      palomakeeper.Keeper
	upgradeKeeper     upgradekeeper.Keeper
}

func initFixture(t ginkgo.FullGinkgoTInterface) *fixture {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("paloma", "pub")
	config.SetBech32PrefixForValidator("palomavaloper", "valoperpub")
	// db := dbm.NewMemDB()
	keys := storetypes.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey,
		distrtypes.StoreKey, stakingtypes.StoreKey,
		schedulertypes.StoreKey, evmmoduletypes.StoreKey,
		valsetmoduletypes.StoreKey, consensusmoduletypes.StoreKey,
		upgradetypes.StoreKey,
	)
	encCfg := moduletestutil.MakeTestEncodingConfig(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		gov.AppModuleBasic{},
		evm.AppModuleBasic{},
		scheduler.AppModuleBasic{},
		valset.AppModuleBasic{},
		consensus.AppModuleBasic{},
		paloma.AppModuleBasic{},
	)

	cdc := encCfg.Codec

	logger := log.NewTestLogger(t)
	cms := integration.CreateMultiStore(keys, logger)

	newCtx := sdk.NewContext(cms, cmtproto.Header{}, true, logger)

	//authority := authtypes.NewModuleAddress(types.ModuleName)

	maccPerms := map[string][]string{
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		schedulertypes.ModuleName:      {authtypes.Burner},
	}

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		address.Bech32Codec{
			Bech32Prefix: params2.AccountAddressPrefix,
		},
		"paloma",
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
	// interfaceRegistry, _ := codecTypes.NewInterfaceRegistryWithOptions(codecTypes.InterfaceRegistryOptions{
	// 	ProtoFiles: proto.HybridResolver,
	// 	SigningOptions: signing.Options{
	// 		CustomGetSigners: make(map[protoreflect.FullName]signing.GetSignersFunc),
	// 		AddressCodec: address.Bech32Codec{
	// 			Bech32Prefix: params2.AccountAddressPrefix,
	// 		},
	// 		ValidatorAddressCodec: address.Bech32Codec{
	// 			Bech32Prefix: params2.ValidatorAddressPrefix,
	// 		},
	// 	},
	// })
	appCodec := codec.NewProtoCodec(cdc.InterfaceRegistry())
	legacyAmino := codec.NewLegacyAmino()
	tkeys := storetypes.NewTransientStoreKeys(paramstypes.TStoreKey)
	paramsKeeper := initParamsKeeper(appCodec, legacyAmino, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[banktypes.StoreKey]),
		accountKeeper,
		BlockedAddresses(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		log.NewNopLogger(),
	)

	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[stakingtypes.StoreKey]),
		accountKeeper,
		bankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		authcodec.NewBech32Codec(params2.ValidatorAddressPrefix),
		authcodec.NewBech32Codec(params2.ConsNodeAddressPrefix),
	)

	valsetKeeper := *valsetmodulekeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[valsetmoduletypes.StoreKey]),
		getSubspace(valsetmoduletypes.ModuleName, paramsKeeper),
		stakingKeeper,
		minimumPigeonVersion,
		sdk.DefaultPowerReduction,
		authcodec.NewBech32Codec(params2.ValidatorAddressPrefix),
	)
	consensusRegistry := consensusmodulekeeper.NewRegistry()
	consensusKeeper := *consensusmodulekeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[consensusmoduletypes.StoreKey]),
		getSubspace(consensusmoduletypes.ModuleName, paramsKeeper),
		valsetKeeper,
		consensusRegistry,
	)
	evmKeeper := *evmmodulekeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[evmmoduletypes.StoreKey]),
		consensusKeeper,
		valsetKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		authcodec.NewBech32Codec(params2.ValidatorAddressPrefix),
	)
	consensusRegistry.Add(
		evmKeeper,
	)
	valsetKeeper.EvmKeeper = evmKeeper
	schedulerKeeper := *keeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[schedulertypes.StoreKey]),
		accountKeeper,
		evmKeeper,
		[]xchain.Bridge{
			evmKeeper,
		},
	)

	appOptions := make(simtestutil.AppOptionsMap, 0)

	appOpts := appOptions
	skipUpgradeHeights := map[int64]bool{}
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}
	homePath := cast.ToString(appOpts.Get(flags.FlagHome))

	bApp := *baseapp.NewBaseApp(
		"integration-app",
		logger,
		db.NewMemDB(),
		encCfg.TxConfig.TxDecoder(),
	)

	version.Version = "v5.1.6"
	bApp.SetVersion(version.Version)
	oldVersion := version.Version
	t.Cleanup(func() {
		version.Version = oldVersion
	})

	upgradeKeeper := *upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		runtime.NewKVStoreService(keys[upgradetypes.StoreKey]),
		appCodec,
		homePath,
		&bApp,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	semverVersion := bApp.Version()

	if !strings.HasPrefix(semverVersion, "v") {
		semverVersion = fmt.Sprintf("v%s", semverVersion)
	}
	palomaKeeper := palomakeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[palomamoduletypes.StoreKey]),
		getSubspace(palomamoduletypes.ModuleName, paramsKeeper),
		semverVersion,
		valsetKeeper,
		upgradeKeeper, //TODO
		authcodec.NewBech32Codec(params2.ValidatorAddressPrefix),
	)

	palomaKeeper.ExternalChains = []palomamoduletypes.ExternalChainSupporterKeeper{
		evmKeeper,
	}

	authModule := auth.NewAppModule(cdc, accountKeeper, authsims.RandomGenesisAccounts, nil)
	bankModule := bank.NewAppModule(cdc, bankKeeper, accountKeeper, nil)
	schedulerModule := scheduler.NewAppModule(cdc, schedulerKeeper, accountKeeper, bankKeeper)
	evmModule := evm.NewAppModule(cdc, evmKeeper, accountKeeper, bankKeeper)
	upgradeModule := upgrade.NewAppModule(&upgradeKeeper, address.NewBech32Codec("cosmos"))
	palomaModule := paloma.NewAppModule(cdc, *palomaKeeper, accountKeeper, bankKeeper)

	integrationApp := integration.NewIntegrationApp(newCtx, logger, keys, cdc, map[string]appmodule.AppModule{
		authtypes.ModuleName:         authModule,
		banktypes.ModuleName:         bankModule,
		schedulertypes.ModuleName:    schedulerModule,
		evmmoduletypes.ModuleName:    evmModule,
		palomamoduletypes.ModuleName: palomaModule,
		upgradetypes.ModuleName:      upgradeModule,
	})

	upgradeKeeper.SetUpgradeHandler(semverVersion, func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (vm module.VersionMap, err error) {
		return
	})

	sdkCtx := sdk.UnwrapSDKContext(integrationApp.Context())

	queryClient := v1.NewQueryClient(integrationApp.QueryHelper())
	legacyQueryClient := v1beta1.NewQueryClient(integrationApp.QueryHelper())
	return &fixture{
		ctx:               sdkCtx,
		codec:             appCodec,
		queryClient:       queryClient,
		legacyQueryClient: legacyQueryClient,
		consensusKeeper:   consensusKeeper,
		valsetKeeper:      valsetKeeper,
		schedulerKeeper:   schedulerKeeper,
		paramsKeeper:      paramsKeeper,
		evmKeeper:         evmKeeper,
		stakingKeeper:     *stakingKeeper,
		palomaKeeper:      *palomaKeeper,
		upgradeKeeper:     upgradeKeeper,
	}

}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range GetMaccPerms() {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	// allow the following addresses to receive funds
	delete(modAccAddrs, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	return modAccAddrs
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	keyTable := ibcclienttypes.ParamKeyTable()
	keyTable.RegisterParamSet(&ibcconnectiontypes.Params{})

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(schedulertypes.ModuleName)
	paramsKeeper.Subspace(consensusmoduletypes.ModuleName)
	paramsKeeper.Subspace(valsetmoduletypes.ModuleName)
	paramsKeeper.Subspace(wasmtypes.ModuleName)
	paramsKeeper.Subspace(evmmoduletypes.ModuleName)
	paramsKeeper.Subspace(gravitymoduletypes.ModuleName)

	paramsKeeper.Subspace(ibcexported.ModuleName).WithKeyTable(keyTable)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName).WithKeyTable(ibctransfertypes.ParamKeyTable())
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName).WithKeyTable(icacontrollertypes.ParamKeyTable())
	paramsKeeper.Subspace(icahosttypes.SubModuleName).WithKeyTable(icahosttypes.ParamKeyTable())

	return paramsKeeper
}

// NOTE: This is solely to be used for testing purposes.
func getSubspace(moduleName string, paramsKeeper paramskeeper.Keeper) paramstypes.Subspace {
	subspace, _ := paramsKeeper.GetSubspace(moduleName)
	return subspace
}
