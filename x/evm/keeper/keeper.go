package keeper

import (
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/VolumeFi/whoops"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	xchain "github.com/palomachain/paloma/internal/x-chain"
	keeperutil "github.com/palomachain/paloma/util/keeper"
	"github.com/palomachain/paloma/util/liblog"
	"github.com/palomachain/paloma/x/consensus/keeper/consensus"
	consensustypes "github.com/palomachain/paloma/x/consensus/types"
	"github.com/palomachain/paloma/x/evm/types"
	gravitymoduletypes "github.com/palomachain/paloma/x/gravity/types"
	ptypes "github.com/palomachain/paloma/x/paloma/types"
	schedulertypes "github.com/palomachain/paloma/x/scheduler/types"
	valsettypes "github.com/palomachain/paloma/x/valset/types"
)

const (
	maxPower                     = 1 << 32
	thresholdForConsensus uint64 = 2_863_311_530
)

const (
	ConsensusTurnstoneMessage     = "evm-turnstone-message"
	ConsensusGetValidatorBalances = "validators-balances"
	ConsensusCollectFundEvents    = "collect-fund-events"
	SignaturePrefix               = "\x19Ethereum Signed Message:\n32"
)

var _ ptypes.ExternalChainSupporterKeeper = Keeper{}

type supportedChainInfo struct {
	subqueue              string
	batch                 bool
	msgType               any
	processAttesationFunc func(Keeper) func(ctx sdk.Context, q consensus.Queuer, msg consensustypes.QueuedSignedMessageI) error
}

var SupportedConsensusQueues = []supportedChainInfo{
	{
		subqueue: ConsensusTurnstoneMessage,
		batch:    false,
		msgType:  &types.Message{},
		processAttesationFunc: func(k Keeper) func(ctx sdk.Context, q consensus.Queuer, msg consensustypes.QueuedSignedMessageI) error {
			return k.attestRouter
		},
	},
	{
		subqueue: ConsensusGetValidatorBalances,
		batch:    false,
		msgType:  &types.ValidatorBalancesAttestation{},
		processAttesationFunc: func(k Keeper) func(ctx sdk.Context, q consensus.Queuer, msg consensustypes.QueuedSignedMessageI) error {
			return k.attestValidatorBalances
		},
	},
	{
		batch:    false,
		subqueue: ConsensusCollectFundEvents,
		msgType:  &types.CollectFunds{},
		processAttesationFunc: func(k Keeper) func(ctx sdk.Context, q consensus.Queuer, msg consensustypes.QueuedSignedMessageI) error {
			return k.attestCollectedFunds
		},
	},
}

func init() {
	// just a check to ensure that there are no duplicates in the supported chain infos
	visited := make(map[string]struct{})
	for _, c := range SupportedConsensusQueues {
		if _, ok := visited[c.subqueue]; ok {
			panic(fmt.Sprintf("cannot have two queues with the same subqueue: %s", c.subqueue))
		}
		visited[c.subqueue] = struct{}{}
	}
}

var _ valsettypes.OnSnapshotBuiltListener = &Keeper{}

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	memKey     storetypes.StoreKey
	paramstore paramtypes.Subspace

	ConsensusKeeper types.ConsensusKeeper
	SchedulerKeeper types.SchedulerKeeper
	Valset          types.ValsetKeeper
	Gravity         types.GravityKeeper
	ider            keeperutil.IDGenerator
	msgSender       types.MsgSender
	msgAssigner     types.MsgAssigner
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	consensusKeeper types.ConsensusKeeper,
	valsetKeeper types.ValsetKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	k := &Keeper{
		cdc:             cdc,
		storeKey:        storeKey,
		memKey:          memKey,
		paramstore:      ps,
		ConsensusKeeper: consensusKeeper,
		Valset:          valsetKeeper,
		msgSender: msgSender{
			ConsensusKeeper: consensusKeeper,
			cdc:             cdc,
		},
		msgAssigner: MsgAssigner{
			valsetKeeper,
		},
	}

	k.ider = keeperutil.NewIDGenerator(keeperutil.StoreGetterFn(k.provideSmartContractStore), []byte("id-key"))

	return k
}

func (k Keeper) PickValidatorForMessage(ctx sdk.Context, chainReferenceID string, requirements *xchain.JobRequirements) (string, error) {
	weights, err := k.GetRelayWeights(ctx, chainReferenceID)
	if err != nil {
		return "", err
	}
	return k.msgAssigner.PickValidatorForMessage(ctx, weights, chainReferenceID, requirements)
}

func (k Keeper) Logger(ctx sdk.Context) liblog.Logr {
	return liblog.FromSDKLogger(ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName)))
}

func (k Keeper) ChangeMinOnChainBalance(ctx sdk.Context, chainReferenceID string, balance *big.Int) error {
	ci, err := k.GetChainInfo(ctx, chainReferenceID)
	if err != nil {
		return err
	}
	ci.MinOnChainBalance = balance.Text(10)
	return k.updateChainInfo(ctx, ci)
}

func (k Keeper) SupportedQueues(ctx sdk.Context) ([]consensus.SupportsConsensusQueueAction, error) {
	chains, err := k.GetAllChainInfos(ctx)
	if err != nil {
		return nil, err
	}

	res := []consensus.SupportsConsensusQueueAction{}

	for _, chainInfo := range chains {
		// if !chainInfo.IsActive() {
		// 	continue
		// }
		for _, queueInfo := range SupportedConsensusQueues {
			queue := consensustypes.Queue(queueInfo.subqueue, xchainType, xchain.ReferenceID(chainInfo.ChainReferenceID))
			opts := *consensus.ApplyOpts(nil,
				consensus.WithChainInfo(xchainType, chainInfo.ChainReferenceID),
				consensus.WithQueueTypeName(queue),
				consensus.WithStaticTypeCheck(queueInfo.msgType),
				consensus.WithBytesToSignCalc(
					consensustypes.BytesToSignFunc(func(msg consensustypes.ConsensusMsg, salt consensustypes.Salt) []byte {
						k := msg.(interface {
							Keccak256(uint64) []byte
						})
						return k.Keccak256(salt.Nonce)
					}),
				),
				consensus.WithVerifySignature(func(bz []byte, sig []byte, address []byte) bool {
					receivedAddr := common.BytesToAddress(address)

					bytesToVerify := crypto.Keccak256(append(
						[]byte(SignaturePrefix),
						bz...,
					))
					recoveredPk, err := crypto.Ecrecover(bytesToVerify, sig)
					if err != nil {
						return false
					}
					pk, err := crypto.UnmarshalPubkey(recoveredPk)
					if err != nil {
						return false
					}
					recoveredAddr := crypto.PubkeyToAddress(*pk)
					return receivedAddr.Hex() == recoveredAddr.Hex()
				}),
			)

			res = append(res, consensus.SupportsConsensusQueueAction{
				QueueOptions:                 opts,
				ProcessMessageForAttestation: queueInfo.processAttesationFunc(k),
			})
			k.Logger(ctx).Debug("supported-queues", "chain-id", chainInfo.ChainReferenceID, "queue", queue)
		}
	}

	return res, nil
}

func (k Keeper) GetAllChainInfos(ctx sdk.Context) ([]*types.ChainInfo, error) {
	_, all, err := keeperutil.IterAll[*types.ChainInfo](k.chainInfoStore(ctx), k.cdc)
	return all, err
}

func (k Keeper) GetChainInfo(ctx sdk.Context, targetChainReferenceID string) (*types.ChainInfo, error) {
	res, err := keeperutil.Load[*types.ChainInfo](k.chainInfoStore(ctx), k.cdc, []byte(targetChainReferenceID))
	if errors.Is(err, keeperutil.ErrNotFound) {
		return nil, ErrChainNotFound.Format(targetChainReferenceID)
	}
	return res, nil
}

// MissingChains returns the chains in this keeper that aren't in the input slice
func (k Keeper) MissingChains(ctx sdk.Context, inputChainReferenceIDs []string) ([]string, error) {
	allChains, err := k.GetAllChainInfos(ctx)
	if err != nil {
		k.Logger(ctx).Error("Unable to get chains infos from keeper")
		return nil, err
	}

	// Build a map to use for efficient comparison
	supportedChainMap := make(map[string]bool, len(inputChainReferenceIDs))
	for _, chainReferenceID := range inputChainReferenceIDs {
		supportedChainMap[chainReferenceID] = true
	}

	// Walk through all chains and aggregate the ones not supported
	var unsuportedChainReferenceIDs []string
	for _, chain := range allChains {
		chainReferenceID := chain.GetChainReferenceID()
		if !chain.IsActive() {
			continue
		}
		if _, found := supportedChainMap[chainReferenceID]; !found {
			unsuportedChainReferenceIDs = append(unsuportedChainReferenceIDs, chainReferenceID)
		}
	}
	return unsuportedChainReferenceIDs, nil
}

func (k Keeper) updateChainInfo(ctx sdk.Context, chainInfo *types.ChainInfo) error {
	return keeperutil.Save(k.chainInfoStore(ctx), k.cdc, []byte(chainInfo.GetChainReferenceID()), chainInfo)
}

func (k Keeper) AddSupportForNewChain(
	ctx sdk.Context,
	chainReferenceID string,
	chainID uint64,
	blockHeight uint64,
	blockHashAtHeight string,
	minimumOnChainBalance *big.Int,
) error {
	_, err := k.GetChainInfo(ctx, chainReferenceID)
	switch {
	case err == nil:
		return ErrCannotAddSupportForChainThatExists.Format(chainReferenceID)
	case errors.Is(err, ErrChainNotFound):
		// we want chain not to exist when adding a new one!
	default:
		return whoops.Wrap(ErrUnexpectedError, err)
	}
	all, err := k.GetAllChainInfos(ctx)
	if err != nil {
		return err
	}
	for _, existing := range all {
		if existing.GetChainID() == chainID {
			return ErrCannotAddSupportForChainThatExists.Format(chainReferenceID).
				WrapS("chain with chainID %d already exists", chainID)
		}
	}

	chainInfo := &types.ChainInfo{
		ChainID:              chainID,
		ChainReferenceID:     chainReferenceID,
		ReferenceBlockHeight: blockHeight,
		ReferenceBlockHash:   blockHashAtHeight,
		MinOnChainBalance:    minimumOnChainBalance.Text(10),
		RelayWeights: &types.RelayWeights{
			Fee:           "1.0",
			Uptime:        "1.0",
			SuccessRate:   "1.0",
			ExecutionTime: "1.0",
		},
	}

	err = k.updateChainInfo(ctx, chainInfo)
	if err != nil {
		return err
	}

	k.TryDeployingLastCompassContractToAllChains(ctx)
	return nil
}

func (k Keeper) ActivateChainReferenceID(
	ctx sdk.Context,
	chainReferenceID string,
	smartContract *types.SmartContract,
	smartContractAddr string,
	smartContractUniqueID []byte,
) (retErr error) {
	defer func() {
		args := []any{
			"chain-reference-id", chainReferenceID,
			"smart-contract-id", smartContract.GetId(),
			"smart-contract-addr", smartContractAddr,
			"smart-contract-unique-id", smartContractUniqueID,
		}
		if retErr != nil {
			args = append(args, "err", retErr)
		}

		if retErr != nil {
			k.Logger(ctx).Error("error while activating chain with a new smart contract", args...)
		} else {
			k.Logger(ctx).Info("activated chain with a new smart contract", args...)
		}
	}()
	chainInfo, err := k.GetChainInfo(ctx, chainReferenceID)
	if err != nil {
		return err
	}
	// if this is called with version lower than the current one, then do nothing
	if chainInfo.GetActiveSmartContractID() >= smartContract.GetId() {
		return nil
	}
	chainInfo.Status = types.ChainInfo_ACTIVE
	chainInfo.Abi = smartContract.GetAbiJSON()
	chainInfo.Bytecode = smartContract.GetBytecode()
	chainInfo.ActiveSmartContractID = smartContract.GetId()

	chainInfo.SmartContractAddr = smartContractAddr
	chainInfo.SmartContractUniqueID = smartContractUniqueID

	k.DeleteSmartContractDeploymentByContractID(ctx, smartContract.GetId(), chainInfo.GetChainReferenceID())

	return k.updateChainInfo(ctx, chainInfo)
}

func (k Keeper) RemoveSupportForChain(ctx sdk.Context, proposal *types.RemoveChainProposal) error {
	_, err := k.GetChainInfo(ctx, proposal.GetChainReferenceID())
	if err != nil {
		return err
	}

	k.chainInfoStore(ctx).Delete([]byte(proposal.GetChainReferenceID()))

	for _, q := range SupportedConsensusQueues {
		queue := consensustypes.Queue(q.subqueue, xchainType, xchain.ReferenceID(proposal.GetChainReferenceID()))
		if e := k.ConsensusKeeper.RemoveConsensusQueue(ctx, queue); e != nil {
			k.Logger(ctx).Error("error removing consensus queue", "err", err, "referenceID", proposal.GetChainReferenceID())
		}
	}

	return nil
}

func (k Keeper) chainInfoStore(ctx sdk.Context) sdk.KVStore {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte("chain-info"))
}

func (k Keeper) PreJobExecution(ctx sdk.Context, job *schedulertypes.Job) error {
	router := job.GetRouting()
	chainReferenceID := router.GetChainReferenceID()
	chain, err := k.GetChainInfo(ctx, chainReferenceID)
	if err != nil {
		k.Logger(ctx).Error("couldn't get chain info",
			"chain-reference-id", chainReferenceID,
			"err", err,
		)
		return err
	}
	// Publish this valset if it differs from the current published valset for this chain
	return k.justInTimeValsetUpdate(ctx, chain)
}

func (k Keeper) justInTimeValsetUpdate(ctx sdk.Context, chain *types.ChainInfo) error {
	latestSnapshot, err := k.Valset.GetCurrentSnapshot(ctx)
	if err != nil {
		k.Logger(ctx).Error("couldn't get latest snapshot", "err", err)
		return err
	}
	if latestSnapshot == nil {
		// For some reason, GetCurrentShapshot is hiding the notFound errors and just returning nil, nil, so we need this
		err := errors.New("nil, nil returned from Valset.GetCurrentSnapshot")
		k.Logger(ctx).Error("unable to find current snapshot", "err", err)
		return err
	}

	chainReferenceID := chain.GetChainReferenceID()

	latestPublishedSnapshot, err := k.Valset.GetLatestSnapshotOnChain(ctx, chainReferenceID)
	if err != nil {
		k.Logger(ctx).Info("couldn't get latest published snapshot for chain.",
			"chain-reference-id", chain.GetChainReferenceID(),
			"err", err,
		)
		return err
	}

	latestValset := transformSnapshotToCompass(latestSnapshot, chainReferenceID, k.Logger(ctx))

	if latestPublishedSnapshot.GetId() == latestSnapshot.GetId() {
		k.Logger(ctx).Info("ignoring valset for chain because it is already most recent",
			"chain-reference-id", chain.GetChainReferenceID(),
			"valset-id", latestValset.GetValsetID(),
		)
		return nil
	}

	if !chain.IsActive() {
		k.Logger(ctx).Info("ignoring valset for chain as the chain is not yet active",
			"chain-reference-id", chain.GetChainReferenceID(),
			"valset-id", latestValset.GetValsetID(),
		)
		return nil
	}

	if !isEnoughToReachConsensus(latestValset) {
		k.Logger(ctx).Info("ignoring valset for chain as there aren't enough validators to form a consensus for this chain",
			"chain-reference-id", chain.GetChainReferenceID(),
			"valset-id", latestValset.GetValsetID(),
		)
		return nil
	}

	assignee, err := k.PickValidatorForMessage(ctx, chain.GetChainReferenceID(), nil)
	if err != nil {
		return err
	}

	err = k.msgSender.SendValsetMsgForChain(ctx, chain, latestValset, assignee)
	if err != nil {
		k.Logger(ctx).Error("unable to send valset message for chain",
			"chain", chain.GetChainReferenceID(),
			"err", err,
		)
	}

	return err
}

func (k Keeper) PublishValsetToChain(ctx sdk.Context, valset types.Valset, chain *types.ChainInfo) error {
	if !chain.IsActive() {
		k.Logger(ctx).Info("ignoring valset for chain as the chain is not yet active",
			"chain-reference-id", chain.GetChainReferenceID(),
			"valset-id", valset.GetValsetID(),
		)
		return nil
	}

	if !isEnoughToReachConsensus(valset) {
		k.Logger(ctx).Info("ignoring valset for chain as there aren't enough validators to form a consensus for this chain",
			"chain-reference-id", chain.GetChainReferenceID(),
			"valset-id", valset.GetValsetID(),
		)
		return nil
	}

	assignee, err := k.PickValidatorForMessage(ctx, chain.GetChainReferenceID(), nil)
	if err != nil {
		k.Logger(ctx).Error("error picking a validator to run the message",
			"chain-reference-id", chain.GetChainReferenceID(),
			"valset-id", valset.GetValsetID(),
			"error", err,
		)
		return err
	}

	err = k.msgSender.SendValsetMsgForChain(ctx, chain, valset, assignee)
	if err != nil {
		k.Logger(ctx).Error("unable to send valset message for chain",
			"chain", chain.GetChainReferenceID(),
			"err", err,
		)
		return err
	}
	return nil
}

func (k Keeper) PublishSnapshotToAllChains(ctx sdk.Context, snapshot *valsettypes.Snapshot, forcePublish bool) error {
	chainInfos, err := k.GetAllChainInfos(ctx)
	if err != nil {
		return err
	}
	logger := k.Logger(ctx)
	for _, chain := range chainInfos {
		valset := transformSnapshotToCompass(snapshot, chain.GetChainReferenceID(), logger)

		latestActiveValset, _ := k.Valset.GetLatestSnapshotOnChain(ctx, chain.GetChainReferenceID())
		if latestActiveValset != nil && !forcePublish {
			latestActiveValsetAge := ctx.BlockTime().Sub(latestActiveValset.CreatedAt)

			// If it's been less than 1 month since publishing a valset, don't publish
			keepWarmDays := 30
			if latestActiveValsetAge < (time.Duration(keepWarmDays) * 24 * time.Hour) {
				k.Logger(ctx).Info(fmt.Sprintf("ignoring valset for chain because chain has had a valset update in the past %d days", keepWarmDays),
					"chain-reference-id", chain.GetChainReferenceID(),
					"current-block-height", ctx.BlockHeight(),
					"current-published-valset-id", latestActiveValset.GetId(),
					"current-published-valset-created-time", latestActiveValset.CreatedAt,
					"valset-id", valset.GetValsetID(),
				)
				continue
			}
		}

		err := k.PublishValsetToChain(ctx, valset, chain)
		if err != nil {
			k.Logger(ctx).Error(err.Error())
		}
	}
	return nil
}

func (k *Keeper) OnSnapshotBuilt(ctx sdk.Context, snapshot *valsettypes.Snapshot) {
	err := k.PublishSnapshotToAllChains(ctx, snapshot, false)
	if err != nil {
		panic(err)
	}

	k.TryDeployingLastCompassContractToAllChains(ctx)
}

type msgSender struct {
	ConsensusKeeper types.ConsensusKeeper
	cdc             codec.BinaryCodec
}

func (m msgSender) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (m msgSender) SendValsetMsgForChain(ctx sdk.Context, chainInfo *types.ChainInfo, valset types.Valset, assignee string) error {
	m.Logger(ctx).Info("snapshot was built and a new update valset message is being sent over",
		"chainInfo-reference-id", chainInfo.GetChainReferenceID(),
		"valset-id", valset.GetValsetID(),
	)

	// clear all other instances of the update valset from the queue
	m.Logger(ctx).Info("clearing previous instances of the update valset from the queue")
	queueName := consensustypes.Queue(ConsensusTurnstoneMessage, xchainType, xchain.ReferenceID(chainInfo.GetChainReferenceID()))
	messages, err := m.ConsensusKeeper.GetMessagesFromQueue(ctx, queueName, 0)
	if err != nil {
		m.Logger(ctx).Error("unable to get messages from queue", "err", err)
		return err
	}

	for _, msg := range messages {
		cmsg, err := msg.ConsensusMsg(m.cdc)
		if err != nil {
			m.Logger(ctx).Error("unable to unpack message", "err", err)
			return err
		}

		mmsg := cmsg.(*types.Message)
		act := mmsg.GetAction()
		if mmsg.GetTurnstoneID() != string(chainInfo.GetSmartContractUniqueID()) {
			return nil
		}
		if _, ok := act.(*types.Message_UpdateValset); ok {
			err := m.ConsensusKeeper.DeleteJob(ctx, queueName, msg.GetId())
			if err != nil {
				m.Logger(ctx).Error("unable to delete message", "err", err)
				return err
			}
		}
	}

	// put update valset message into the queue
	msgID, err := m.ConsensusKeeper.PutMessageInQueue(
		ctx,
		consensustypes.Queue(ConsensusTurnstoneMessage, xchainType, xchain.ReferenceID(chainInfo.GetChainReferenceID())),
		&types.Message{
			TurnstoneID:      string(chainInfo.GetSmartContractUniqueID()),
			ChainReferenceID: chainInfo.GetChainReferenceID(),
			Action: &types.Message_UpdateValset{
				UpdateValset: &types.UpdateValset{
					Valset: &valset,
				},
			},
			Assignee: assignee,
		}, nil,
	)
	if err != nil {
		m.Logger(ctx).Error("unable to put message in the queue", "err", err)
		return err
	}

	m.Logger(ctx).With("new-message-id", msgID).Debug("Valset update message added to consensus queue.")
	return nil
}

func (k Keeper) CheckExternalBalancesForChain(ctx sdk.Context, chainReferenceID string) error {
	snapshot, err := k.Valset.GetCurrentSnapshot(ctx)
	if err != nil {
		return err
	}

	var msg types.ValidatorBalancesAttestation
	msg.FromBlockTime = ctx.BlockTime().UTC()

	for _, val := range snapshot.GetValidators() {
		for _, ext := range val.GetExternalChainInfos() {
			if ext.GetChainReferenceID() == chainReferenceID && ext.GetChainType() == "evm" {
				msg.ValAddresses = append(msg.ValAddresses, val.GetAddress())
				msg.HexAddresses = append(msg.HexAddresses, ext.GetAddress())
				k.Logger(ctx).Debug("check-external-balances-for-chain",
					"chain-reference-id", chainReferenceID,
					"msg-val-address", val.GetAddress(),
					"msg-hex-address", ext.GetAddress(),
					"val-share-count", val.ShareCount,
					"ext-chain-balance", ext.Balance,
				)
			}
		}
	}

	if len(msg.ValAddresses) == 0 {
		return nil
	}
	_, err = k.ConsensusKeeper.PutMessageInQueue(
		ctx,
		consensustypes.Queue(ConsensusGetValidatorBalances, xchainType, chainReferenceID),
		&msg,
		&consensus.PutOptions{
			RequireSignatures: false,
			PublicAccessData:  []byte{1}, // anything because pigeon cares if public access data exists to be able to provide evidence
		},
	)

	return err
}

func isEnoughToReachConsensus(val types.Valset) bool {
	var sum uint64
	for _, power := range val.Powers {
		sum += power
	}

	return sum >= thresholdForConsensus
}

func transformSnapshotToCompass(snapshot *valsettypes.Snapshot, chainReferenceID string, logger log.Logger) types.Valset {
	var totalShares sdk.Int
	if snapshot != nil {
		totalShares = snapshot.TotalShares
	}
	logger.Info("transformSnapshotToCompass",
		"snapshot-id", snapshot.GetId(),
		"snapshot-height", snapshot.GetHeight(),
		"snapshot-total-shares", totalShares,
		"snapshot-validators-length", len(snapshot.GetValidators()),
	)
	validators := make([]valsettypes.Validator, len(snapshot.GetValidators()))
	copy(validators, snapshot.GetValidators())

	sort.SliceStable(validators, func(i, j int) bool {
		// doing GTE because we want a reverse sort
		return validators[i].ShareCount.GTE(validators[j].ShareCount)
	})

	totalPowerInt := sdk.NewInt(0)
	for _, val := range validators {
		totalPowerInt = totalPowerInt.Add(val.ShareCount)
	}

	totalPower := totalPowerInt.Int64()

	valset := types.Valset{
		ValsetID: snapshot.GetId(),
	}

	logger.Info("transformSnapshotToCompass",
		"total-power", totalPower,
		"valset-id", valset.ValsetID,
	)

	for _, val := range validators {
		for _, ext := range val.GetExternalChainInfos() {
			if strings.ToLower(ext.GetChainType()) == xchainType && ext.GetChainReferenceID() == chainReferenceID {
				power := maxPower * (float64(val.ShareCount.Int64()) / float64(totalPower))

				valset.Validators = append(valset.Validators, ext.Address)
				valset.Powers = append(valset.Powers, uint64(power))
			}
		}
	}

	return valset
}

func (k Keeper) ModuleName() string { return types.ModuleName }

func generateSmartContractID(ctx sdk.Context) (res [32]byte) {
	b := []byte(fmt.Sprintf("%d", ctx.BlockHeight()))
	copy(res[:], b)
	return
}

func (k Keeper) SetRelayWeights(ctx sdk.Context, chainReferenceID string, weights *types.RelayWeights) error {
	chainInfo, err := k.GetChainInfo(ctx, chainReferenceID)
	if err != nil {
		return err
	}

	chainInfo.RelayWeights = weights

	return keeperutil.Save(k.chainInfoStore(ctx), k.cdc, []byte(chainReferenceID), chainInfo)
}

func (k Keeper) GetRelayWeights(ctx sdk.Context, chainReferenceID string) (*types.RelayWeights, error) {
	chainInfo, err := k.GetChainInfo(ctx, chainReferenceID)
	if err != nil {
		return &types.RelayWeights{}, err
	}

	return chainInfo.RelayWeights, nil
}

func (k Keeper) GetEthAddressByValidator(ctx sdk.Context, validator sdk.ValAddress, chainReferenceId string) (ethAddress *gravitymoduletypes.EthAddress, found bool, err error) {
	chainInfos, err := k.Valset.GetValidatorChainInfos(ctx, validator)
	if err != nil {
		return ethAddress, false, err
	}
	for _, chainInfo := range chainInfos {
		if chainInfo.GetChainReferenceID() == chainReferenceId {
			ethAddress = &gravitymoduletypes.EthAddress{}
			err = ethAddress.SetAddress(chainInfo.GetAddress())
			if err != nil {
				return ethAddress, false, err
			}
			return ethAddress, true, nil
		}
	}
	return ethAddress, false, nil
}

func (k Keeper) GetValidatorAddressByEthAddress(ctx sdk.Context, ethAddr gravitymoduletypes.EthAddress, chainReferenceId string) (valAddr sdk.ValAddress, found bool, err error) {
	validatorsExternalAccounts, err := k.Valset.GetAllChainInfos(ctx)
	if err != nil {
		return valAddr, false, err
	}
	for _, validatorExternalAccounts := range validatorsExternalAccounts {
		for _, chainInfo := range validatorExternalAccounts.ExternalChainInfo {
			if chainInfo.GetChainReferenceID() == chainReferenceId && ethAddr.GetAddress().String() == chainInfo.GetAddress() {
				return validatorExternalAccounts.Address, true, nil
			}
		}
	}
	return
}
