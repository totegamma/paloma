package keeper

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/VolumeFi/whoops"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	xchain "github.com/palomachain/paloma/internal/x-chain"
	"github.com/palomachain/paloma/util/slice"
	"github.com/palomachain/paloma/x/evm/types"
)

var xchainType = xchain.Type("evm")

var _ xchain.Bridge = Keeper{}

func (k Keeper) XChainType() xchain.Type {
	return xchainType
}

func (k Keeper) XChainReferenceIDs(ctx sdk.Context) []xchain.ReferenceID {
	chainInfos, err := k.GetAllChainInfos(ctx)
	if err != nil {
		panic(err)
	}

	return slice.Map(chainInfos, func(ci *types.ChainInfo) xchain.ReferenceID {
		return xchain.ReferenceID(ci.GetChainReferenceID())
	})
}

func (k Keeper) unmarshalJob(definition, payload []byte, chainReferenceID xchain.ReferenceID) (resdef types.JobDefinition, respay types.JobPayload, reserr error) {
	var jobDefinition types.JobDefinition
	var jobPayload types.JobPayload

	err := json.Unmarshal(definition, &jobDefinition)
	if err != nil {
		reserr = err
		return
	}
	err = json.Unmarshal(payload, &jobPayload)
	if err != nil {
		reserr = err
		return
	}

	return jobDefinition, jobPayload, nil
}

func (k Keeper) VerifyJob(ctx sdk.Context, definition, payload []byte, chainReferenceID xchain.ReferenceID) error {
	_, _, err := k.unmarshalJob(definition, payload, chainReferenceID)
	return err
}

// ExecuteJob schedules the definition and payload for execution via consensus queue
func (k Keeper) ExecuteJob(ctx sdk.Context, jcfg *xchain.JobConfiguration) (uint64, error) {
	def, load, err := k.unmarshalJob(jcfg.Definition, jcfg.Payload, jcfg.RefID)
	if err != nil {
		return 0, err
	}
	ci, err := k.GetChainInfo(ctx, jcfg.RefID)
	if err != nil {
		return 0, err
	}

	var hexBytes []byte
	switch {
	case jcfg.SenderAddress != nil:
		hexBytes, err = hex.DecodeString(addressToHex(jcfg.SenderAddress))
		if err != nil {
			return 0, err
		}
	case jcfg.ContractAddress != nil:
		hexBytes, err = hex.DecodeString(addressToHex(jcfg.ContractAddress))
		if err != nil {
			return 0, err
		}
	}

	// zero pad our byte array to 32 bytes
	appendSenderBytes, err := zeroPadBytes(hexBytes, 32)
	if err != nil {
		return 0, err
	}

	modifiedPayload := append(common.FromHex(load.GetHexPayload()), appendSenderBytes...)

	return k.AddSmartContractExecutionToConsensus(
		ctx,
		jcfg.RefID,
		string(ci.GetSmartContractUniqueID()),
		&types.SubmitLogicCall{
			HexContractAddress: def.GetAddress(),
			Abi:                common.FromHex(def.GetABI()),
			Payload:            modifiedPayload,
			Deadline:           ctx.BlockTime().Add(10 * time.Minute).Unix(),
			SenderAddress:      jcfg.SenderAddress,
			ContractAddress:    jcfg.ContractAddress,
			ExecutionRequirements: types.SubmitLogicCall_ExecutionRequirements{
				EnforceMEVRelay: jcfg.Requirements.EnforceMEVRelay,
			},
		},
	)
}

func addressToHex(address sdk.AccAddress) string {
	return fmt.Sprintf("%x", address.Bytes())
}

func zeroPadBytes(input []byte, size int) ([]byte, error) {
	inputLen := len(input)
	if inputLen > size {
		return nil, whoops.String(fmt.Sprintf("Can not zero pad byte array of size %d to %d", inputLen, size))
	}
	ret := make([]byte, size)
	copy(ret[size-inputLen:], input)

	return ret, nil
}
