package keeper_test

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"math/big"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/VolumeFi/whoops"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// 	"github.com/palomachain/paloma/app"
// 	"github.com/palomachain/paloma/testutil"
// 	"github.com/palomachain/paloma/testutil/rand"
// 	"github.com/palomachain/paloma/testutil/sample"
// 	consensustypes "github.com/palomachain/paloma/x/consensus/types"
// 	"github.com/palomachain/paloma/x/evm/keeper"
// 	"github.com/palomachain/paloma/x/evm/types"
// 	valsettypes "github.com/palomachain/paloma/x/valset/types"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// var (
// 	contractAbi         = string(whoops.Must(ioutil.ReadFile("testdata/sample-abi.json")))
// 	contractBytecodeStr = string(whoops.Must(ioutil.ReadFile("testdata/sample-bytecode.out")))
// )

// func genValidators(numValidators, totalConsPower int) []stakingtypes.Validator {
// 	return testutil.GenValidators(numValidators, totalConsPower)
// }

// func TestEndToEndForEvmArbitraryCall(t *testing.T) {
// 	chainType, chainReferenceID := consensustypes.ChainTypeEVM, "eth-main"
// 	a := app.NewTestApp(t, false)
// 	ctx := a.NewContext(false)

// 	newChain := &types.AddChainProposal{
// 		ChainReferenceID:  "eth-main",
// 		Title:             "bla",
// 		Description:       "bla",
// 		BlockHeight:       uint64(123),
// 		BlockHashAtHeight: "0x1234",
// 	}

// 	err := a.EvmKeeper.AddSupportForNewChain(
// 		ctx,
// 		newChain.GetChainReferenceID(),
// 		newChain.GetChainID(),
// 		newChain.GetBlockHeight(),
// 		newChain.GetBlockHashAtHeight(),
// 		big.NewInt(55),
// 	)
// 	require.NoError(t, err)

// 	err = a.EvmKeeper.ActivateChainReferenceID(ctx, newChain.ChainReferenceID, &types.SmartContract{Id: 123}, "addr", []byte("abc"))
// 	require.NoError(t, err)

// 	validators := genValidators(25, 25000)
// 	for _, val := range validators {
// 		a.StakingKeeper.SetValidator(ctx, val)
// 	}

// 	for _, validator := range validators {
// 		valAddr, err := validator.GetConsAddr()
// 		require.NoError(t, err)
// 		pubKey, err := validator.ConsPubKey()
// 		require.NoError(t, err)
// 		err = a.ValsetKeeper.AddExternalChainInfo(ctx, validator.GetOperator(), []*valsettypes.ExternalChainInfo{
// 			{
// 				ChainType:        "evm",
// 				ChainReferenceID: newChain.GetChainReferenceID(),
// 				Address:          valAddr.String(),
// 				Pubkey:           pubKey.Bytes(),
// 			},
// 		})
// 		require.NoError(t, err)
// 	}

// 	_, err = a.ValsetKeeper.TriggerSnapshotBuild(ctx)
// 	require.NoError(t, err)

// 	smartContractAddr := common.BytesToAddress(rand.Bytes(5))
// 	_, err = a.EvmKeeper.AddSmartContractExecutionToConsensus(
// 		ctx,
// 		chainReferenceID,
// 		"",
// 		&types.SubmitLogicCall{
// 			Payload: func() []byte {
// 				evm := whoops.Must(abi.JSON(strings.NewReader(sample.SimpleABI)))
// 				return whoops.Must(evm.Pack("store", big.NewInt(1337)))
// 			}(),
// 			HexContractAddress: smartContractAddr.Hex(),
// 			Abi:                []byte(sample.SimpleABI),
// 			Deadline:           1337,
// 		},
// 	)

// 	require.NoError(t, err)

// 	private, err := crypto.GenerateKey()
// 	require.NoError(t, err)

// 	accAddr := crypto.PubkeyToAddress(private.PublicKey)
// 	err = a.ValsetKeeper.AddExternalChainInfo(ctx, validators[0].GetOperator(), []*valsettypes.ExternalChainInfo{
// 		{
// 			ChainType:        chainType,
// 			ChainReferenceID: chainReferenceID,
// 			Address:          accAddr.Hex(),
// 			Pubkey:           accAddr[:],
// 		},
// 	})

// 	require.NoError(t, err)
// 	queue := consensustypes.Queue(keeper.ConsensusTurnstoneMessage, chainType, chainReferenceID)
// 	msgs, err := a.ConsensusKeeper.GetMessagesForSigning(ctx, queue, validators[0].GetOperator())

// 	for _, msg := range msgs {
// 		sigbz, err := crypto.Sign(
// 			crypto.Keccak256(
// 				[]byte(keeper.SignaturePrefix),
// 				msg.GetBytesToSign(),
// 			),
// 			private,
// 		)
// 		require.NoError(t, err)
// 		err = a.ConsensusKeeper.AddMessageSignature(
// 			ctx,
// 			validators[0].GetOperator(),
// 			[]*consensustypes.ConsensusMessageSignature{
// 				{
// 					Id:              msg.GetId(),
// 					QueueTypeName:   queue,
// 					Signature:       sigbz,
// 					SignedByAddress: accAddr.Hex(),
// 				},
// 			},
// 		)
// 		require.NoError(t, err)
// 	}
// }

// func TestFirstSnapshot_OnSnapshotBuilt(t *testing.T) {
// 	a := app.NewTestApp(t, false)
// 	ctx := a.NewContext(false)

// 	newChain := &types.AddChainProposal{
// 		ChainReferenceID:  "bob",
// 		Title:             "bla",
// 		Description:       "bla",
// 		BlockHeight:       uint64(123),
// 		BlockHashAtHeight: "0x1234",
// 	}
// 	err := a.EvmKeeper.AddSupportForNewChain(
// 		ctx,
// 		newChain.GetChainReferenceID(),
// 		newChain.GetChainID(),
// 		newChain.GetBlockHeight(),
// 		newChain.GetBlockHashAtHeight(),
// 		big.NewInt(55),
// 	)
// 	require.NoError(t, err)
// 	err = a.EvmKeeper.ActivateChainReferenceID(
// 		ctx,
// 		newChain.ChainReferenceID,
// 		&types.SmartContract{
// 			Id: 123,
// 		},
// 		"addr",
// 		[]byte("abc"),
// 	)
// 	require.NoError(t, err)

// 	validators := genValidators(25, 25000)
// 	for _, val := range validators {
// 		a.StakingKeeper.SetValidator(ctx, val)
// 		err = a.ValsetKeeper.AddExternalChainInfo(ctx, val.GetOperator(), []*valsettypes.ExternalChainInfo{
// 			{
// 				ChainType:        "evm",
// 				ChainReferenceID: "bob",
// 				Address:          rand.ETHAddress().Hex(),
// 				Pubkey:           []byte("pk" + rand.ETHAddress().Hex()),
// 			},
// 		})
// 		require.NoError(t, err)
// 	}

// 	queue := fmt.Sprintf("evm/%s/%s", newChain.GetChainReferenceID(), keeper.ConsensusTurnstoneMessage)

// 	msgs, err := a.ConsensusKeeper.GetMessagesFromQueue(ctx, queue, 100)
// 	require.NoError(t, err)
// 	require.Empty(t, msgs)

// 	_, err = a.ValsetKeeper.TriggerSnapshotBuild(ctx)
// 	require.NoError(t, err)

// 	msgs, err = a.ConsensusKeeper.GetMessagesFromQueue(ctx, queue, 100)
// 	require.NoError(t, err)
// 	require.Len(t, msgs, 1)
// }

// func TestRecentPublishedSnapshot_OnSnapshotBuilt(t *testing.T) {
// 	a := app.NewTestApp(t, false)
// 	ctx := a.NewContext(false)

// 	newChain := &types.AddChainProposal{
// 		ChainReferenceID:  "bob",
// 		Title:             "bla",
// 		Description:       "bla",
// 		BlockHeight:       uint64(123),
// 		BlockHashAtHeight: "0x1234",
// 	}
// 	err := a.EvmKeeper.AddSupportForNewChain(
// 		ctx,
// 		newChain.GetChainReferenceID(),
// 		newChain.GetChainID(),
// 		newChain.GetBlockHeight(),
// 		newChain.GetBlockHashAtHeight(),
// 		big.NewInt(55),
// 	)
// 	require.NoError(t, err)
// 	err = a.EvmKeeper.ActivateChainReferenceID(
// 		ctx,
// 		newChain.ChainReferenceID,
// 		&types.SmartContract{
// 			Id: 123,
// 		},
// 		"addr",
// 		[]byte("abc"),
// 	)
// 	require.NoError(t, err)

// 	validators := genValidators(25, 25000)
// 	for _, val := range validators {
// 		a.StakingKeeper.SetValidator(ctx, val)
// 		err = a.ValsetKeeper.AddExternalChainInfo(ctx, val.GetOperator(), []*valsettypes.ExternalChainInfo{
// 			{
// 				ChainType:        "evm",
// 				ChainReferenceID: "bob",
// 				Address:          rand.ETHAddress().Hex(),
// 				Pubkey:           []byte("pk" + rand.ETHAddress().Hex()),
// 			},
// 		})
// 		require.NoError(t, err)
// 	}

// 	queue := fmt.Sprintf("evm/%s/%s", newChain.GetChainReferenceID(), keeper.ConsensusTurnstoneMessage)

// 	msgs, err := a.ConsensusKeeper.GetMessagesFromQueue(ctx, queue, 1)
// 	require.NoError(t, err)
// 	require.Empty(t, msgs)

// 	// Remove the listeners to set current state
// 	snapshotListeners := a.ValsetKeeper.SnapshotListeners
// 	a.ValsetKeeper.SnapshotListeners = []valsettypes.OnSnapshotBuiltListener{}

// 	_, err = a.ValsetKeeper.TriggerSnapshotBuild(ctx)
// 	require.NoError(t, err)

// 	msgs, err = a.ConsensusKeeper.GetMessagesFromQueue(ctx, queue, 100)
// 	require.NoError(t, err)
// 	require.Len(t, msgs, 0)

// 	latestSnapshot, err := a.ValsetKeeper.GetCurrentSnapshot(ctx)
// 	require.NoError(t, err)

// 	latestSnapshot.Chains = []string{"bob"}
// 	err = a.ValsetKeeper.SaveModifiedSnapshot(ctx, latestSnapshot)
// 	require.NoError(t, err)

// 	// Add two validators to make this new snapshot worthy
// 	validators = genValidators(2, 25000)
// 	for _, val := range validators {
// 		a.StakingKeeper.SetValidator(ctx, val)
// 		err = a.ValsetKeeper.AddExternalChainInfo(ctx, val.GetOperator(), []*valsettypes.ExternalChainInfo{
// 			{
// 				ChainType:        "evm",
// 				ChainReferenceID: "bob",
// 				Address:          rand.ETHAddress().Hex(),
// 				Pubkey:           []byte("pk" + rand.ETHAddress().Hex()),
// 			},
// 		})
// 		require.NoError(t, err)
// 	}

// 	// Add the listeners back on for the test
// 	a.ValsetKeeper.SnapshotListeners = snapshotListeners

// 	_, err = a.ValsetKeeper.TriggerSnapshotBuild(ctx)
// 	require.NoError(t, err)

// 	msgs, err = a.ConsensusKeeper.GetMessagesFromQueue(ctx, queue, 100)
// 	require.NoError(t, err)
// 	require.Len(t, msgs, 0) // We don't expect a message because there is already a recent snapshot for the chain
// }

// func TestOldPublishedSnapshot_OnSnapshotBuilt(t *testing.T) {
// 	a := app.NewTestApp(t, false)
// 	ctx := a.NewContext(false)

// 	newChain := &types.AddChainProposal{
// 		ChainReferenceID:  "bob",
// 		Title:             "bla",
// 		Description:       "bla",
// 		BlockHeight:       uint64(123),
// 		BlockHashAtHeight: "0x1234",
// 	}
// 	err := a.EvmKeeper.AddSupportForNewChain(
// 		ctx,
// 		newChain.GetChainReferenceID(),
// 		newChain.GetChainID(),
// 		newChain.GetBlockHeight(),
// 		newChain.GetBlockHashAtHeight(),
// 		big.NewInt(55),
// 	)
// 	require.NoError(t, err)
// 	err = a.EvmKeeper.ActivateChainReferenceID(
// 		ctx,
// 		newChain.ChainReferenceID,
// 		&types.SmartContract{
// 			Id: 123,
// 		},
// 		"addr",
// 		[]byte("abc"),
// 	)
// 	require.NoError(t, err)

// 	validators := genValidators(25, 25000)
// 	for _, val := range validators {
// 		a.StakingKeeper.SetValidator(ctx, val)
// 		err = a.ValsetKeeper.AddExternalChainInfo(ctx, val.GetOperator(), []*valsettypes.ExternalChainInfo{
// 			{
// 				ChainType:        "evm",
// 				ChainReferenceID: "bob",
// 				Address:          rand.ETHAddress().Hex(),
// 				Pubkey:           []byte("pk" + rand.ETHAddress().Hex()),
// 			},
// 		})
// 		require.NoError(t, err)
// 	}

// 	queue := fmt.Sprintf("evm/%s/%s", newChain.GetChainReferenceID(), keeper.ConsensusTurnstoneMessage)

// 	msgs, err := a.ConsensusKeeper.GetMessagesFromQueue(ctx, queue, 1)
// 	require.NoError(t, err)
// 	require.Empty(t, msgs)

// 	// Remove the listeners to set current state
// 	snapshotListeners := a.ValsetKeeper.SnapshotListeners
// 	a.ValsetKeeper.SnapshotListeners = []valsettypes.OnSnapshotBuiltListener{}

// 	_, err = a.ValsetKeeper.TriggerSnapshotBuild(ctx)
// 	require.NoError(t, err)

// 	msgs, err = a.ConsensusKeeper.GetMessagesFromQueue(ctx, queue, 100)
// 	require.NoError(t, err)
// 	require.Len(t, msgs, 0)

// 	latestSnapshot, err := a.ValsetKeeper.GetCurrentSnapshot(ctx)
// 	require.NoError(t, err)

// 	// Age the latest snapshot by 30 days, 1 minute, set as active on chain
// 	latestSnapshot.Chains = []string{"bob"}
// 	latestSnapshot.CreatedAt = ctx.BlockTime().Add(-((30 * 24 * time.Hour) + time.Minute))
// 	err = a.ValsetKeeper.SaveModifiedSnapshot(ctx, latestSnapshot)
// 	require.NoError(t, err)

// 	// Add two validators to make this new snapshot worthy
// 	validators = genValidators(2, 25000)
// 	for _, val := range validators {
// 		a.StakingKeeper.SetValidator(ctx, val)
// 		err = a.ValsetKeeper.AddExternalChainInfo(ctx, val.GetOperator(), []*valsettypes.ExternalChainInfo{
// 			{
// 				ChainType:        "evm",
// 				ChainReferenceID: "bob",
// 				Address:          rand.ETHAddress().Hex(),
// 				Pubkey:           []byte("pk" + rand.ETHAddress().Hex()),
// 			},
// 		})
// 		require.NoError(t, err)
// 	}

// 	// Add the listeners back on for the test
// 	a.ValsetKeeper.SnapshotListeners = snapshotListeners

// 	_, err = a.ValsetKeeper.TriggerSnapshotBuild(ctx)
// 	require.NoError(t, err)

// 	msgs, err = a.ConsensusKeeper.GetMessagesFromQueue(ctx, queue, 100)
// 	require.NoError(t, err)
// 	require.Len(t, msgs, 1) // We expect a new message because the previous one is a week old
// }

// func TestInactiveChain_OnSnapshotBuilt(t *testing.T) {
// 	a := app.NewTestApp(t, false)
// 	ctx := a.NewContext(false)

// 	validators := genValidators(25, 25000)
// 	for _, val := range validators {
// 		a.StakingKeeper.SetValidator(ctx, val)
// 	}

// 	queue := fmt.Sprintf("evm/%s/%s", "bob", keeper.ConsensusTurnstoneMessage)

// 	_, err := a.ValsetKeeper.TriggerSnapshotBuild(ctx)
// 	require.NoError(t, err)

// 	_, err = a.ConsensusKeeper.GetMessagesFromQueue(ctx, queue, 100)
// 	require.Error(t, err) // We expect an error from this
// }

// func TestAddingSupportForNewChain(t *testing.T) {
// 	a := app.NewTestApp(t, false)
// 	ctx := a.NewContext(false)

// 	t.Run("with happy path there are no errors", func(t *testing.T) {
// 		newChain := &types.AddChainProposal{
// 			ChainReferenceID:  "bob",
// 			Title:             "bla",
// 			Description:       "bla",
// 			BlockHeight:       uint64(123),
// 			BlockHashAtHeight: "0x1234",
// 		}
// 		err := a.EvmKeeper.AddSupportForNewChain(
// 			ctx,
// 			newChain.GetChainReferenceID(),
// 			newChain.GetChainID(),
// 			newChain.GetBlockHeight(),
// 			newChain.GetBlockHashAtHeight(),
// 			big.NewInt(55),
// 		)
// 		require.NoError(t, err)

// 		gotChainInfo, err := a.EvmKeeper.GetChainInfo(ctx, newChain.GetChainReferenceID())
// 		require.NoError(t, err)

// 		require.Equal(t, newChain.GetChainReferenceID(), gotChainInfo.GetChainReferenceID())
// 		require.Equal(t, newChain.GetBlockHashAtHeight(), gotChainInfo.GetReferenceBlockHash())
// 		require.Equal(t, newChain.GetBlockHeight(), gotChainInfo.GetReferenceBlockHeight())
// 		t.Run("it returns an error if we try to add a chian whose chainID already exists", func(t *testing.T) {
// 			newChain.ChainReferenceID = "something_new"
// 			err := a.EvmKeeper.AddSupportForNewChain(
// 				ctx,
// 				newChain.GetChainReferenceID(),
// 				newChain.GetChainID(),
// 				newChain.GetBlockHeight(),
// 				newChain.GetBlockHashAtHeight(),
// 				big.NewInt(55),
// 			)
// 			require.ErrorIs(t, err, keeper.ErrCannotAddSupportForChainThatExists)
// 		})
// 	})

// 	t.Run("when chainReferenceID already exists then it returns an error", func(t *testing.T) {
// 		newChain := &types.AddChainProposal{
// 			ChainReferenceID:  "bob",
// 			Title:             "bla",
// 			Description:       "bla",
// 			BlockHeight:       uint64(123),
// 			BlockHashAtHeight: "0x1234",
// 		}
// 		err := a.EvmKeeper.AddSupportForNewChain(
// 			ctx,
// 			newChain.GetChainReferenceID(),
// 			newChain.GetChainID(),
// 			newChain.GetBlockHeight(),
// 			newChain.GetBlockHashAtHeight(),

// 			big.NewInt(55),
// 		)
// 		require.Error(t, err)
// 	})

// 	t.Run("activating chain", func(t *testing.T) {
// 		t.Run("if the chain does not exist it returns the error", func(t *testing.T) {
// 			err := a.EvmKeeper.ActivateChainReferenceID(ctx, "i don't exist", &types.SmartContract{}, "", []byte{})
// 			require.Error(t, err)
// 		})
// 		t.Run("works when chain exists", func(t *testing.T) {
// 			err := a.EvmKeeper.ActivateChainReferenceID(ctx, "bob", &types.SmartContract{Id: 123}, "addr", []byte("unique id"))
// 			require.NoError(t, err)
// 			gotChainInfo, err := a.EvmKeeper.GetChainInfo(ctx, "bob")
// 			require.NoError(t, err)

// 			require.Equal(t, "addr", gotChainInfo.GetSmartContractAddr())
// 			require.Equal(t, []byte("unique id"), gotChainInfo.GetSmartContractUniqueID())
// 		})
// 	})

// 	t.Run("removing chain", func(t *testing.T) {
// 		t.Run("if the chain does not exist it returns the error", func(t *testing.T) {
// 			err := a.EvmKeeper.RemoveSupportForChain(ctx, &types.RemoveChainProposal{
// 				ChainReferenceID: "i don't exist",
// 			})
// 			require.Error(t, err)
// 		})
// 		t.Run("works when chain exists", func(t *testing.T) {
// 			err := a.EvmKeeper.RemoveSupportForChain(ctx, &types.RemoveChainProposal{
// 				ChainReferenceID: "bob",
// 			})
// 			require.NoError(t, err)
// 			_, err = a.EvmKeeper.GetChainInfo(ctx, "bob")
// 			require.Error(t, keeper.ErrChainNotFound)
// 		})
// 	})
// }

// func TestKeeper_ValidatorSupportsAllChains(t *testing.T) {
// 	testcases := []struct {
// 		name     string
// 		setup    func(sdk.Context, app.TestApp) sdk.ValAddress
// 		expected bool
// 	}{
// 		{
// 			name: "returns true when all chains supported",
// 			setup: func(ctx sdk.Context, a app.TestApp) sdk.ValAddress {
// 				for i, chainId := range []string{"chain-1", "chain-2"} {
// 					newChain := &types.AddChainProposal{
// 						ChainReferenceID:  chainId,
// 						ChainID:           uint64(i),
// 						Title:             "bla",
// 						Description:       "bla",
// 						BlockHeight:       uint64(123),
// 						BlockHashAtHeight: "0x1234",
// 					}

// 					err := a.EvmKeeper.AddSupportForNewChain(
// 						ctx,
// 						newChain.GetChainReferenceID(),
// 						newChain.GetChainID(),
// 						newChain.GetBlockHeight(),
// 						newChain.GetBlockHashAtHeight(),
// 						big.NewInt(55),
// 					)
// 					require.NoError(t, err)

// 					err = a.EvmKeeper.ActivateChainReferenceID(ctx, newChain.ChainReferenceID, &types.SmartContract{Id: 123}, fmt.Sprintf("addr%d", i), []byte("abc"))
// 					require.NoError(t, err)
// 				}

// 				validator := genValidators(1, 1000)[0]
// 				a.StakingKeeper.SetValidator(ctx, validator)

// 				private, err := crypto.GenerateKey()
// 				require.NoError(t, err)

// 				accAddr := crypto.PubkeyToAddress(private.PublicKey)

// 				// Add support for both chains created
// 				externalChains := make([]*valsettypes.ExternalChainInfo, 2)
// 				for i, chainId := range []string{"chain-1", "chain-2"} {
// 					externalChains[i] = &valsettypes.ExternalChainInfo{
// 						ChainType:        "evm",
// 						ChainReferenceID: chainId,
// 						Address:          accAddr.Hex(),
// 						Pubkey:           accAddr[:],
// 					}
// 				}
// 				err = a.ValsetKeeper.AddExternalChainInfo(ctx, validator.GetOperator(), externalChains)
// 				require.NoError(t, err)

// 				return validator.GetOperator()
// 			},
// 			expected: true,
// 		},
// 		{
// 			name: "returns false when a chain is not supported",
// 			setup: func(ctx sdk.Context, a app.TestApp) sdk.ValAddress {
// 				for i, chainId := range []string{"chain-1", "chain-2"} {
// 					newChain := &types.AddChainProposal{
// 						ChainReferenceID:  chainId,
// 						ChainID:           uint64(i),
// 						Title:             "bla",
// 						Description:       "bla",
// 						BlockHeight:       uint64(123),
// 						BlockHashAtHeight: "0x1234",
// 					}

// 					err := a.EvmKeeper.AddSupportForNewChain(
// 						ctx,
// 						newChain.GetChainReferenceID(),
// 						newChain.GetChainID(),
// 						newChain.GetBlockHeight(),
// 						newChain.GetBlockHashAtHeight(),
// 						big.NewInt(55),
// 					)
// 					require.NoError(t, err)

// 					err = a.EvmKeeper.ActivateChainReferenceID(ctx, newChain.ChainReferenceID, &types.SmartContract{Id: 123}, fmt.Sprintf("addr%d", i), []byte("abc"))
// 					require.NoError(t, err)
// 				}

// 				validator := genValidators(1, 1000)[0]
// 				a.StakingKeeper.SetValidator(ctx, validator)

// 				private, err := crypto.GenerateKey()
// 				require.NoError(t, err)

// 				accAddr := crypto.PubkeyToAddress(private.PublicKey)

// 				// Only add support for one of two chains created
// 				err = a.ValsetKeeper.AddExternalChainInfo(
// 					ctx,
// 					validator.GetOperator(),
// 					[]*valsettypes.ExternalChainInfo{
// 						{
// 							ChainType:        "evm",
// 							ChainReferenceID: "chain-1",
// 							Address:          accAddr.Hex(),
// 							Pubkey:           accAddr[:],
// 						},
// 					},
// 				)
// 				require.NoError(t, err)

// 				return validator.GetOperator()
// 			},
// 			expected: false,
// 		},
// 	}

// 	asserter := assert.New(t)
// 	for _, tt := range testcases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			a := app.NewTestApp(t, false)
// 			ctx := a.NewContext(false)

// 			validatorAddress := tt.setup(ctx, a)

// 			actual := a.ValsetKeeper.ValidatorSupportsAllChains(ctx, validatorAddress)
// 			asserter.Equal(tt.expected, actual)
// 		})
// 	}
// }

// func TestWithGinkgo(t *testing.T) {
// 	RegisterFailHandler(Fail)

// 	RunSpecs(t, "EVM keeper")
// }

// var _ = Describe("evm", func() {
// 	// smartContractAddr := common.BytesToAddress(rand.Bytes(5))
// 	// chainType, chainReferenceID := consensustypes.ChainTypeEVM, "eth-main"
// 	var a app.TestApp
// 	var ctx sdk.Context
// 	var validators []stakingtypes.Validator
// 	newChain := &types.AddChainProposal{
// 		ChainReferenceID:  "eth-main",
// 		Title:             "bla",
// 		Description:       "bla",
// 		BlockHeight:       uint64(123),
// 		BlockHashAtHeight: "0x1234",
// 	}
// 	smartContract := &types.SmartContract{
// 		Id:       1,
// 		AbiJSON:  contractAbi,
// 		Bytecode: common.FromHex(contractBytecodeStr),
// 	}
// 	smartContract2 := &types.SmartContract{
// 		Id:       2,
// 		AbiJSON:  contractAbi,
// 		Bytecode: common.FromHex(contractBytecodeStr),
// 	}

// 	BeforeEach(func() {
// 		a = app.NewTestApp(GinkgoT(), false)
// 		ctx = a.NewContext(false)
// 	})

// 	Context("multiple chains and smart contracts", func() {
// 		Describe("trying to add support for the same chain twice", func() {
// 			It("returns an error", func() {
// 				err := a.EvmKeeper.AddSupportForNewChain(
// 					ctx,
// 					newChain.GetChainReferenceID(),
// 					newChain.GetChainID(),
// 					newChain.GetBlockHeight(),
// 					newChain.GetBlockHashAtHeight(),
// 					big.NewInt(55),
// 				)
// 				Expect(err).To(BeNil())

// 				err = a.EvmKeeper.AddSupportForNewChain(
// 					ctx,
// 					newChain.GetChainReferenceID(),
// 					newChain.GetChainID(),
// 					newChain.GetBlockHeight(),
// 					newChain.GetBlockHashAtHeight(),
// 					big.NewInt(55),
// 				)
// 				Expect(err).To(MatchError(keeper.ErrCannotAddSupportForChainThatExists))
// 			})
// 		})

// 		Describe("ensuring that there can be two chains at the same time", func() {
// 			chain1 := &types.AddChainProposal{
// 				ChainReferenceID:  "chain1",
// 				Title:             "bla",
// 				Description:       "bla",
// 				BlockHeight:       uint64(456),
// 				BlockHashAtHeight: "0x1234",
// 				ChainID:           1,
// 			}
// 			chain2 := &types.AddChainProposal{
// 				ChainReferenceID:  "chain2",
// 				Title:             "bla",
// 				Description:       "bla",
// 				BlockHeight:       uint64(123),
// 				BlockHashAtHeight: "0x5678",
// 				ChainID:           2,
// 			}
// 			BeforeEach(func() {
// 				validators = genValidators(25, 25000)
// 				for _, val := range validators {
// 					a.StakingKeeper.SetValidator(ctx, val)
// 				}
// 			})

// 			JustBeforeEach(func() {
// 				for _, val := range validators {
// 					private1, err := crypto.GenerateKey()
// 					private2, err := crypto.GenerateKey()
// 					Expect(err).To(BeNil())
// 					accAddr1 := crypto.PubkeyToAddress(private1.PublicKey)
// 					accAddr2 := crypto.PubkeyToAddress(private2.PublicKey)
// 					err = a.ValsetKeeper.AddExternalChainInfo(ctx, val.GetOperator(), []*valsettypes.ExternalChainInfo{
// 						{
// 							ChainType:        "evm",
// 							ChainReferenceID: chain1.ChainReferenceID,
// 							Address:          accAddr1.Hex(),
// 							Pubkey:           []byte("pub key 1" + accAddr1.Hex()),
// 						},
// 						{
// 							ChainType:        "evm",
// 							ChainReferenceID: chain2.ChainReferenceID,
// 							Address:          accAddr2.Hex(),
// 							Pubkey:           []byte("pub key 2" + accAddr2.Hex()),
// 						},
// 					})
// 					Expect(err).To(BeNil())
// 				}
// 				_, err := a.ValsetKeeper.TriggerSnapshotBuild(ctx)
// 				Expect(err).To(BeNil())
// 			})

// 			BeforeEach(func() {
// 				By("adding chain1 works")
// 				err := a.EvmKeeper.AddSupportForNewChain(
// 					ctx,
// 					chain1.GetChainReferenceID(),
// 					chain1.GetChainID(),
// 					chain1.GetBlockHeight(),
// 					chain1.GetBlockHashAtHeight(),
// 					big.NewInt(55),
// 				)
// 				Expect(err).To(BeNil())

// 				By("adding chain2 works")
// 				err = a.EvmKeeper.AddSupportForNewChain(
// 					ctx,
// 					chain2.GetChainReferenceID(),
// 					chain2.GetChainID(),
// 					chain2.GetBlockHeight(),
// 					chain2.GetBlockHashAtHeight(),
// 					big.NewInt(55),
// 				)
// 				Expect(err).To(BeNil())
// 			})

// 			Context("adding smart contract", func() {
// 				It("adds a new smart contract deployment", func() {
// 					By("simple assertion that two smart contracts share different ids", func() {
// 						Expect(smartContract.GetId()).NotTo(Equal(smartContract2.GetId()))
// 					})
// 					By("saving a new smart contract", func() {
// 						Expect(
// 							a.EvmKeeper.HasAnySmartContractDeployment(ctx, chain1.GetChainReferenceID()),
// 						).To(BeFalse())
// 						Expect(
// 							a.EvmKeeper.HasAnySmartContractDeployment(ctx, chain2.GetChainReferenceID()),
// 						).To(BeFalse())

// 						sc, err := a.EvmKeeper.SaveNewSmartContract(ctx, smartContract.GetAbiJSON(), smartContract.GetBytecode())
// 						Expect(err).To(BeNil())

// 						err = a.EvmKeeper.SetAsCompassContract(ctx, sc)
// 						Expect(err).To(BeNil())

// 						Expect(
// 							a.EvmKeeper.HasAnySmartContractDeployment(ctx, chain1.GetChainReferenceID()),
// 						).To(BeTrue())
// 						Expect(
// 							a.EvmKeeper.HasAnySmartContractDeployment(ctx, chain2.GetChainReferenceID()),
// 						).To(BeTrue())
// 					})

// 					By("removing a smart deployment for chain1 - it means that it was successfully uploaded", func() {
// 						a.EvmKeeper.DeleteSmartContractDeploymentByContractID(ctx, smartContract.GetId(), chain1.GetChainReferenceID())
// 						Expect(
// 							a.EvmKeeper.HasAnySmartContractDeployment(ctx, chain1.GetChainReferenceID()),
// 						).To(BeFalse())
// 						Expect(
// 							a.EvmKeeper.HasAnySmartContractDeployment(ctx, chain2.GetChainReferenceID()),
// 						).To(BeTrue())
// 					})

// 					By("activating a new smart contract it removes a deployment for chain1 but it doesn't for chain2", func() {
// 						err := a.EvmKeeper.ActivateChainReferenceID(ctx, chain1.GetChainReferenceID(), smartContract, "addr1", []byte("id1"))
// 						Expect(err).To(BeNil())
// 						Expect(
// 							a.EvmKeeper.HasAnySmartContractDeployment(ctx, chain1.GetChainReferenceID()),
// 						).To(BeFalse())
// 						Expect(
// 							a.EvmKeeper.HasAnySmartContractDeployment(ctx, chain2.GetChainReferenceID()),
// 						).To(BeTrue())

// 						By("verify that the chain's smart contract id has been deployed", func() {
// 							ci, err := a.EvmKeeper.GetChainInfo(ctx, chain1.GetChainReferenceID())
// 							Expect(err).To(BeNil())
// 							Expect(ci.GetActiveSmartContractID()).To(Equal(smartContract.GetId()))
// 						})
// 					})

// 					By("adding a new smart contract deployment deploys it to chain1 only", func() {
// 						sc, err := a.EvmKeeper.SaveNewSmartContract(ctx, smartContract2.GetAbiJSON(), smartContract2.GetBytecode())
// 						Expect(err).To(BeNil())
// 						err = a.EvmKeeper.SetAsCompassContract(ctx, sc)
// 						Expect(err).To(BeNil())
// 						Expect(
// 							a.EvmKeeper.HasAnySmartContractDeployment(ctx, chain1.GetChainReferenceID()),
// 						).To(BeTrue())
// 					})

// 					By("activating a new-new smart contract it deploys it to chain 1", func() {
// 						err := a.EvmKeeper.ActivateChainReferenceID(ctx, chain1.GetChainReferenceID(), smartContract2, "addr2", []byte("id2"))
// 						Expect(err).To(BeNil())
// 						Expect(
// 							a.EvmKeeper.HasAnySmartContractDeployment(ctx, chain2.GetChainReferenceID()),
// 						).To(BeTrue())
// 						By("verify that the chain's smart contract id has been deployed", func() {
// 							ci, err := a.EvmKeeper.GetChainInfo(ctx, chain1.GetChainReferenceID())
// 							Expect(err).To(BeNil())
// 							Expect(ci.GetActiveSmartContractID()).To(Equal(smartContract2.GetId()))
// 						})
// 					})
// 				})
// 			})
// 		})
// 	})

// 	Describe("on snapshot build", func() {
// 		var snapshot *valsettypes.Snapshot
// 		When("validator set is valid", func() {
// 			BeforeEach(func() {
// 				validators = genValidators(25, 25000)
// 				for _, val := range validators {
// 					a.StakingKeeper.SetValidator(ctx, val)
// 				}
// 			})

// 			When("evm chain and smart contract both exist", func() {
// 				BeforeEach(func() {
// 					for _, val := range validators {
// 						private, err := crypto.GenerateKey()
// 						Expect(err).To(BeNil())
// 						accAddr := crypto.PubkeyToAddress(private.PublicKey)
// 						err = a.ValsetKeeper.AddExternalChainInfo(ctx, val.GetOperator(), []*valsettypes.ExternalChainInfo{
// 							{
// 								ChainType:        "evm",
// 								ChainReferenceID: newChain.ChainReferenceID,
// 								Address:          accAddr.Hex(),
// 								Pubkey:           []byte("pub key" + accAddr.Hex()),
// 							},
// 							{
// 								ChainType:        "evm",
// 								ChainReferenceID: "new-chain",
// 								Address:          accAddr.Hex(),
// 								Pubkey:           []byte("pub key" + accAddr.Hex()),
// 							},
// 						})
// 						Expect(err).To(BeNil())
// 					}
// 					var err error
// 					snapshot, err = a.ValsetKeeper.TriggerSnapshotBuild(ctx)
// 					Expect(err).To(BeNil())
// 				})

// 				BeforeEach(func() {
// 					err := a.EvmKeeper.AddSupportForNewChain(
// 						ctx,
// 						newChain.GetChainReferenceID(),
// 						newChain.GetChainID(),
// 						newChain.GetBlockHeight(),
// 						newChain.GetBlockHashAtHeight(),
// 						big.NewInt(55),
// 					)
// 					Expect(err).To(BeNil())

// 					sc, err := a.EvmKeeper.SaveNewSmartContract(ctx, smartContract.GetAbiJSON(), smartContract.GetBytecode())
// 					Expect(err).To(BeNil())
// 					err = a.EvmKeeper.SetAsCompassContract(ctx, sc)
// 					Expect(err).To(BeNil())

// 					err = a.EvmKeeper.ActivateChainReferenceID(ctx, newChain.ChainReferenceID, smartContract, "addr", []byte("abc"))
// 					Expect(err).To(BeNil())

// 					By("it should have upload smart contract message", func() {
// 						msgs, err := a.ConsensusKeeper.GetMessagesFromQueue(ctx, "evm/eth-main/evm-turnstone-message", 5)

// 						Expect(err).To(BeNil())
// 						Expect(len(msgs)).To(Equal(1))

// 						con, err := msgs[0].ConsensusMsg(a.AppCodec())
// 						Expect(err).To(BeNil())

// 						evmMsg, ok := con.(*types.Message)
// 						Expect(ok).To(BeTrue())

// 						_, ok = evmMsg.GetAction().(*types.Message_UploadSmartContract)
// 						Expect(ok).To(BeTrue())

// 						a.ConsensusKeeper.DeleteJob(ctx, "evm/eth-main/evm-turnstone-message", msgs[0].GetId())
// 					})
// 				})

// 				It("expects update valset message to exist", func() {
// 					a.EvmKeeper.OnSnapshotBuilt(ctx, snapshot)
// 					msgs, err := a.ConsensusKeeper.GetMessagesFromQueue(ctx, "evm/eth-main/evm-turnstone-message", 5)

// 					Expect(err).To(BeNil())
// 					Expect(len(msgs)).To(Equal(1))

// 					con, err := msgs[0].ConsensusMsg(a.AppCodec())
// 					Expect(err).To(BeNil())

// 					evmMsg, ok := con.(*types.Message)
// 					Expect(ok).To(BeTrue())

// 					_, ok = evmMsg.GetAction().(*types.Message_UpdateValset)
// 					Expect(ok).To(BeTrue())
// 				})

// 				When("adding another chain which is not yet active", func() {
// 					BeforeEach(func() {
// 						err := a.EvmKeeper.AddSupportForNewChain(
// 							ctx,
// 							"new-chain",
// 							123,
// 							uint64(123),
// 							"0x1234",
// 							big.NewInt(55),
// 						)
// 						Expect(err).To(BeNil())
// 					})

// 					It("tries to deploy a smart contract to it", func() {
// 						a.EvmKeeper.OnSnapshotBuilt(ctx, snapshot)
// 						msgs, err := a.ConsensusKeeper.GetMessagesFromQueue(ctx, "evm/new-chain/evm-turnstone-message", 5)
// 						Expect(err).To(BeNil())
// 						Expect(len(msgs)).To(Equal(1))

// 						con, err := msgs[0].ConsensusMsg(a.AppCodec())
// 						Expect(err).To(BeNil())

// 						evmMsg, ok := con.(*types.Message)
// 						Expect(ok).To(BeTrue())

// 						_, ok = evmMsg.GetAction().(*types.Message_UploadSmartContract)
// 						Expect(ok).To(BeTrue())
// 					})
// 				})

// 				When("there is another upload valset already in", func() {
// 					BeforeEach(func() {
// 						err := a.EvmKeeper.AddSupportForNewChain(
// 							ctx,
// 							"new-chain",
// 							123,
// 							uint64(123),
// 							"0x1234",
// 							big.NewInt(55),
// 						)
// 						Expect(err).To(BeNil())
// 						err = a.EvmKeeper.ActivateChainReferenceID(ctx, "new-chain", &types.SmartContract{Id: 123}, "addr", []byte("abc"))
// 						Expect(err).To(BeNil())
// 						for _, val := range validators {
// 							private, err := crypto.GenerateKey()
// 							Expect(err).To(BeNil())
// 							accAddr := crypto.PubkeyToAddress(private.PublicKey)
// 							err = a.ValsetKeeper.AddExternalChainInfo(ctx, val.GetOperator(), []*valsettypes.ExternalChainInfo{
// 								{
// 									ChainType:        "evm",
// 									ChainReferenceID: "new-chain",
// 									Address:          accAddr.Hex(),
// 									Pubkey:           []byte("pub key" + accAddr.Hex()),
// 								},
// 							})
// 							Expect(err).To(BeNil())
// 						}
// 					})
// 					BeforeEach(func() {
// 						msgs, err := a.ConsensusKeeper.GetMessagesFromQueue(ctx, "evm/new-chain/evm-turnstone-message", 5)
// 						Expect(err).To(BeNil())
// 						for _, msg := range msgs {
// 							// we are now clearing the deploy smart contract from the queue as we don't need it
// 							a.ConsensusKeeper.DeleteJob(ctx, "evm/new-chain/evm-turnstone-message", msg.GetId())
// 						}
// 						a.ConsensusKeeper.PutMessageInQueue(ctx, "evm/new-chain/evm-turnstone-message", &types.Message{
// 							TurnstoneID:      "abc",
// 							ChainReferenceID: "new-chain",
// 							Action: &types.Message_UpdateValset{
// 								UpdateValset: &types.UpdateValset{
// 									Valset: &types.Valset{
// 										ValsetID: 777,
// 									},
// 								},
// 							},
// 						}, nil)
// 					})
// 					It("deletes the old smart deployment", func() {
// 						a.EvmKeeper.OnSnapshotBuilt(ctx, snapshot)
// 						msgs, err := a.ConsensusKeeper.GetMessagesFromQueue(ctx, "evm/new-chain/evm-turnstone-message", 5)
// 						Expect(err).To(BeNil())
// 						Expect(len(msgs)).To(Equal(1))

// 						con, err := msgs[0].ConsensusMsg(a.AppCodec())
// 						Expect(err).To(BeNil())

// 						evmMsg, ok := con.(*types.Message)
// 						Expect(ok).To(BeTrue())

// 						vset, ok := evmMsg.GetAction().(*types.Message_UpdateValset)
// 						Expect(ok).To(BeTrue())
// 						Expect(vset.UpdateValset.GetValset().GetValsetID()).NotTo(Equal(uint64(777)))
// 						Expect(len(vset.UpdateValset.GetValset().GetValidators())).NotTo(BeZero())
// 					})
// 				})
// 			})
// 		})

// 		When("validator set is too tiny", func() {
// 			BeforeEach(func() {
// 				validators = genValidators(25, 25000)[:5]
// 				for _, val := range validators {
// 					a.StakingKeeper.SetValidator(ctx, val)
// 				}
// 				_, err := a.ValsetKeeper.TriggerSnapshotBuild(ctx)
// 				Expect(err).To(BeNil())
// 			})

// 			Context("evm chain and smart contract both exist", func() {
// 				BeforeEach(func() {
// 					err := a.EvmKeeper.AddSupportForNewChain(
// 						ctx,
// 						newChain.GetChainReferenceID(),
// 						newChain.GetChainID(),
// 						newChain.GetBlockHeight(),
// 						newChain.GetBlockHashAtHeight(),
// 						big.NewInt(55),
// 					)
// 					Expect(err).To(BeNil())
// 					sc, err := a.EvmKeeper.SaveNewSmartContract(ctx, smartContract.GetAbiJSON(), smartContract.GetBytecode())
// 					Expect(err).To(BeNil())
// 					err = a.EvmKeeper.SetAsCompassContract(ctx, sc)
// 					Expect(err).To(BeNil())
// 				})

// 				It("doesn't put any message into a queue", func() {
// 					msgs, err := a.ConsensusKeeper.GetMessagesFromQueue(ctx, "evm/eth-main/evm-turnstone-message", 5)
// 					Expect(err).To(BeNil())
// 					Expect(msgs).To(BeZero())
// 				})
// 			})
// 		})
// 	})
// })

// var _ = Describe("change min on chain balance", func() {
// 	var a app.TestApp
// 	var ctx sdk.Context
// 	newChain := &types.AddChainProposal{
// 		ChainReferenceID:  "eth-main",
// 		Title:             "bla",
// 		Description:       "bla",
// 		BlockHeight:       uint64(123),
// 		BlockHashAtHeight: "0x1234",
// 	}

// 	BeforeEach(func() {
// 		a = app.NewTestApp(GinkgoT(), false)
// 		ctx = a.NewContext(false)
// 	})

// 	When("chain info exists", func() {
// 		BeforeEach(func() {
// 			err := a.EvmKeeper.AddSupportForNewChain(ctx, newChain.GetChainReferenceID(), newChain.GetChainID(), 1, "a", big.NewInt(55))
// 			Expect(err).To(BeNil())
// 		})

// 		BeforeEach(func() {
// 			ci, err := a.EvmKeeper.GetChainInfo(ctx, newChain.GetChainReferenceID())
// 			Expect(err).To(BeNil())
// 			balance, err := ci.GetMinOnChainBalanceBigInt()
// 			Expect(err).To(BeNil())
// 			Expect(balance.Text(10)).To(Equal(big.NewInt(55).Text(10)))
// 		})

// 		It("changes the on chain balance", func() {
// 			err := a.EvmKeeper.ChangeMinOnChainBalance(ctx, newChain.GetChainReferenceID(), big.NewInt(888))
// 			Expect(err).To(BeNil())

// 			ci, err := a.EvmKeeper.GetChainInfo(ctx, newChain.GetChainReferenceID())
// 			Expect(err).To(BeNil())
// 			balance, err := ci.GetMinOnChainBalanceBigInt()
// 			Expect(err).To(BeNil())
// 			Expect(balance.Text(10)).To(Equal(big.NewInt(888).Text(10)))
// 		})
// 	})

// 	When("chain info does not exists", func() {
// 		It("returns an error", func() {
// 			err := a.EvmKeeper.ChangeMinOnChainBalance(ctx, newChain.GetChainReferenceID(), big.NewInt(888))
// 			Expect(err).To(MatchError(keeper.ErrChainNotFound))
// 		})
// 	})
// })

// var _ = Describe("change relay weights", func() {
// 	var a app.TestApp
// 	var ctx sdk.Context
// 	newChain := &types.AddChainProposal{
// 		ChainReferenceID:  "eth-main",
// 		Title:             "bla",
// 		Description:       "bla",
// 		BlockHeight:       uint64(123),
// 		BlockHashAtHeight: "0x1234",
// 	}

// 	BeforeEach(func() {
// 		a = app.NewTestApp(GinkgoT(), false)
// 		ctx = a.NewContext(false)
// 	})

// 	When("chain info exists", func() {
// 		BeforeEach(func() {
// 			err := a.EvmKeeper.AddSupportForNewChain(ctx, newChain.GetChainReferenceID(), newChain.GetChainID(), 1, "a", big.NewInt(55))
// 			Expect(err).To(BeNil())
// 		})

// 		BeforeEach(func() {
// 			ci, err := a.EvmKeeper.GetChainInfo(ctx, newChain.GetChainReferenceID())
// 			Expect(err).To(BeNil())
// 			weights := ci.GetRelayWeights()
// 			Expect(weights).To(Equal(&types.RelayWeights{
// 				Fee:           "1.0",
// 				Uptime:        "1.0",
// 				SuccessRate:   "1.0",
// 				ExecutionTime: "1.0",
// 			}))
// 		})

// 		It("changes the relay weights", func() {
// 			newWeights := &types.RelayWeights{
// 				Fee:           "0.12",
// 				Uptime:        "0.34",
// 				SuccessRate:   "0.56",
// 				ExecutionTime: "0.78",
// 			}
// 			err := a.EvmKeeper.SetRelayWeights(ctx, newChain.GetChainReferenceID(), newWeights)
// 			Expect(err).To(BeNil())

// 			ci, err := a.EvmKeeper.GetChainInfo(ctx, newChain.GetChainReferenceID())
// 			Expect(err).To(BeNil())
// 			weights := ci.GetRelayWeights()
// 			Expect(weights).To(Equal(newWeights))
// 		})
// 	})

// 	When("chain info does not exists", func() {
// 		It("returns an error", func() {
// 			err := a.EvmKeeper.SetRelayWeights(ctx, newChain.GetChainReferenceID(), &types.RelayWeights{
// 				Fee:           "0.12",
// 				Uptime:        "0.34",
// 				SuccessRate:   "0.56",
// 				ExecutionTime: "0.78",
// 			})
// 			Expect(err).To(MatchError(keeper.ErrChainNotFound))
// 		})
// 	})
// })
