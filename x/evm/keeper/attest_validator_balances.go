package keeper

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/palomachain/paloma/x/consensus/keeper/consensus"
	consensustypes "github.com/palomachain/paloma/x/consensus/types"
	"github.com/palomachain/paloma/x/evm/types"
)

func (k Keeper) attestValidatorBalances(ctx context.Context, q consensus.Queuer, msg consensustypes.QueuedSignedMessageI) (retErr error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.Logger(sdkCtx).Debug("attest-validator-balances", "msg-id", msg.GetId(), "msg-nonce", msg.Nonce())
	if len(msg.GetEvidence()) == 0 {
		return nil
	}

	ctx, writeCache := sdkCtx.CacheContext()
	defer func() {
		if retErr == nil {
			writeCache()
		}
	}()

	consensusMsg, err := msg.ConsensusMsg(k.cdc)
	if err != nil {
		return err
	}

	request := consensusMsg.(*types.ValidatorBalancesAttestation)

	evidence, err := k.findEvidenceThatWon(ctx, msg.GetEvidence())
	if err != nil {
		if errors.Is(err, ErrConsensusNotAchieved) {
			return nil
		}
		return err
	}

	defer func() {
		// given that there was enough evidence for a proof, regardless of the outcome,
		// we should remove this from the queue as there isn't much that we can do about it.
		if err := q.Remove(sdkCtx, msg.GetId()); err != nil {
			k.Logger(sdkCtx).Error("error removing message, attestValidatorBalances", "msg-id", msg.GetId(), "msg-nonce", msg.Nonce())
		}
	}()

	_, chainReferenceID := q.ChainInfo()
	ci, err := k.GetChainInfo(ctx, chainReferenceID)
	if err != nil {
		return err
	}

	minBalance, err := ci.GetMinOnChainBalanceBigInt()
	if err != nil {
		return err
	}

	return k.processValidatorBalanceProof(sdkCtx, request, evidence, chainReferenceID, minBalance)
}

func (k Keeper) processValidatorBalanceProof(
	ctx sdk.Context,
	request *types.ValidatorBalancesAttestation,
	evidence any,
	chainReferenceID string,
	minBalance *big.Int,
) error {
	switch winner := evidence.(type) {
	case *types.ValidatorBalancesAttestationRes:
		for i := range request.GetHexAddresses() {
			valAddr := request.ValAddresses[i]
			hexAddr, balanceStr := common.HexToAddress(request.HexAddresses[i]), winner.Balances[i]
			balance, ok := new(big.Int).SetString(balanceStr, 10)
			if !ok {
				k.Logger(ctx).Error(
					"invalid balance string when attesting to EVM balance",
					"balance", balanceStr,
					"val-addr", valAddr,
					"eth-addr", hexAddr,
				)
				// WHAT TO DO NOW?!?!?! jail the poor fellow that has invalid balance format??
				// blame the flock for reporting this??!?
				continue
			}

			if err := k.Valset.SetValidatorBalance(ctx, valAddr, "evm", chainReferenceID, hexAddr.String(), balance); err != nil {
				k.Logger(ctx).Error(
					"error setting validator balance",
					"err", err,
					"val-addr", valAddr,
					"eth-addr", hexAddr,
				)
			}
			if balance.Cmp(minBalance) == -1 || balance.Cmp(big.NewInt(0)) == 0 {
				isJailed, err := k.Valset.IsJailed(ctx, valAddr)
				if err != nil {
					k.Logger(ctx).Error(
						"error in checking jailed validator",
						"err", err,
						"val-addr", valAddr,
						"eth-addr", hexAddr,
					)
				}

				if !isJailed {
					if err := k.Valset.Jail(ctx, valAddr, fmt.Sprintf(types.JailReasonNotEnoughFunds, chainReferenceID, balanceStr, minBalance)); err != nil {
						k.Logger(ctx).Error(
							"error jailing validator",
							"err", err,
							"val-addr", valAddr,
							"eth-addr", hexAddr,
						)
					}
				}
			}
		}
	default:
		return ErrUnexpectedError.WrapS("unknown type %t when attesting", winner)
	}

	return nil
}
