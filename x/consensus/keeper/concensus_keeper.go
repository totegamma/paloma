package keeper

import (
	"github.com/VolumeFi/whoops"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/palomachain/paloma/util/liblog"
	"github.com/palomachain/paloma/util/slice"
	"github.com/palomachain/paloma/x/consensus/keeper/consensus"
	"github.com/palomachain/paloma/x/consensus/types"
	evmtypes "github.com/palomachain/paloma/x/evm/types"
	valsettypes "github.com/palomachain/paloma/x/valset/types"
)

var defaultResponseMessageCount = 1000

// getConsensusQueue gets the consensus queue for the given type.
func (k Keeper) getConsensusQueue(ctx sdk.Context, queueTypeName string) (consensus.Queuer, error) {
	for _, q := range k.registry.slice {
		supportedQueues, err := q.SupportedQueues(ctx)
		if err != nil {
			return nil, err
		}
		opts := func() *consensus.SupportsConsensusQueueAction {
			for _, q := range supportedQueues {
				if q.QueueTypeName == queueTypeName {
					return &q
				}
			}
			return nil
		}()
		if opts == nil {
			continue
		}

		if opts.Ider.Zero() {
			opts.Ider = k.ider
		}

		if opts.Sg == nil {
			opts.Sg = k
		}

		if opts.Cdc == nil {
			opts.Cdc = k.cdc
		}

		if opts.Batched {
			return consensus.NewBatchQueue(opts.QueueOptions), nil
		}
		return consensus.NewQueue(opts.QueueOptions), nil
	}
	return nil, ErrConsensusQueueNotImplemented.Format(queueTypeName)
}

func (k Keeper) RemoveConsensusQueue(ctx sdk.Context, queueTypeName string) error {
	cq, err := k.getConsensusQueue(ctx, queueTypeName)
	if err != nil {
		return err
	}
	consensus.RemoveQueueCompletely(ctx, cq)
	return nil
}

func (k Keeper) PutMessageInQueue(ctx sdk.Context, queueTypeName string, msg consensus.ConsensusMsg, opts *consensus.PutOptions) (uint64, error) {
	cq, err := k.getConsensusQueue(ctx, queueTypeName)
	if err != nil {
		k.Logger(ctx).Error("error while getting consensus queue", "error", err)
		return 0, err
	}
	msgID, err := cq.Put(ctx, msg, opts)
	if err != nil {
		k.Logger(ctx).Error("error while putting message into queue", "error", err)
		return 0, err
	}
	k.Logger(ctx).Info(
		"put message into consensus queue",
		"queue-type-name", queueTypeName,
		"message-id", msgID,
	)
	return msgID, nil
}

// GetMessagesForSigning returns messages for a single validator that needs to be signed.
func (k Keeper) GetMessagesForSigning(ctx sdk.Context, queueTypeName string, valAddress sdk.ValAddress) (msgs []types.QueuedSignedMessageI, err error) {
	msgs, err = k.GetMessagesFromQueue(ctx, queueTypeName, 0)
	if err != nil {
		return nil, err
	}

	// Filter out already signed messages
	msgs = slice.Filter(msgs, func(msg types.QueuedSignedMessageI) bool {
		for _, signData := range msg.GetSignData() {
			if signData.ValAddress.Equals(valAddress) {
				return false
			}
		}
		return true
	})

	if len(msgs) > defaultResponseMessageCount {
		msgs = msgs[:defaultResponseMessageCount]
	}

	return msgs, nil
}

// GetMessagesForRelaying returns messages for a single validator to relay.
func (k Keeper) GetMessagesForRelaying(ctx sdk.Context, queueTypeName string, valAddress sdk.ValAddress) (msgs []types.QueuedSignedMessageI, err error) {
	msgs, err = k.GetMessagesFromQueue(ctx, queueTypeName, 0)
	if err != nil {
		return nil, err
	}

	// Check for existing valset update messages on any target chains
	valsetUpdatesOnChainLkUp := make(map[string]uint64)
	for _, v := range msgs {
		cm, err := v.ConsensusMsg(k.cdc)
		if err != nil {
			liblog.FromSDKLogger(k.Logger(ctx)).WithError(err).Error("Failed to get consensus msg")
			continue
		}

		m, ok := cm.(*evmtypes.Message)
		if !ok {
			continue
		}

		action := m.GetAction()
		_, ok = action.(*evmtypes.Message_UpdateValset)
		if ok {
			if _, found := valsetUpdatesOnChainLkUp[m.GetChainReferenceID()]; found {
				// Looks like we already have a pending valset update for this chain,
				// we want to keep the earlierst message ID for a valset update we found,
				// so we can skip here.
				continue
			}
			valsetUpdatesOnChainLkUp[m.GetChainReferenceID()] = v.GetId()
		}
	}

	// Filter down to just messages for target chains without pending valset updates on them
	msgs = slice.Filter(msgs, func(msg types.QueuedSignedMessageI) bool {
		cm, err := msg.ConsensusMsg(k.cdc)
		if err != nil {
			// NO cross chain message, just return true
			liblog.FromSDKLogger(k.Logger(ctx)).WithError(err).Error("Failed to get consensus msg")
			return true
		}

		m, ok := cm.(*evmtypes.Message)
		if !ok {
			// NO cross chain message, just return true
			return true
		}

		// Cross chain message for relaying, return only if no pending valset update on target chain
		vuMid, found := valsetUpdatesOnChainLkUp[m.GetChainReferenceID()]
		if !found {
			return true
		}

		// Looks like there is a valset update for the target chain,
		// only return true if this message is younger than the valset update
		return msg.GetId() <= vuMid
	})

	// Filter down to just messages assigned to this validator
	msgs = slice.Filter(msgs, func(msg types.QueuedSignedMessageI) bool {
		var unpackedMsg evmtypes.TurnstoneMsg
		if err := k.cdc.UnpackAny(msg.GetMsg(), &unpackedMsg); err != nil {
			k.Logger(ctx).With("err", err).Error("Failed to unpack message")
			return false
		}

		return unpackedMsg.GetAssignee() == valAddress.String()
	})

	// Filter down to just messages that have neither publicAccessData nor errorData
	msgs = slice.Filter(msgs, func(msg types.QueuedSignedMessageI) bool {
		return msg.GetPublicAccessData() == nil && msg.GetErrorData() == nil
	})

	if len(msgs) > defaultResponseMessageCount {
		msgs = msgs[:defaultResponseMessageCount]
	}

	return msgs, nil
}

// GetMessagesForAttesting returns messages for a single validator to attest.
func (k Keeper) GetMessagesForAttesting(ctx sdk.Context, queueTypeName string, valAddress sdk.ValAddress) (msgs []types.QueuedSignedMessageI, err error) {
	msgs, err = k.GetMessagesFromQueue(ctx, queueTypeName, 0)
	if err != nil {
		return nil, err
	}

	// Filter down to just messages that have either publicAccessData or errorData
	msgs = slice.Filter(msgs, func(msg types.QueuedSignedMessageI) bool {
		return msg.GetPublicAccessData() != nil || msg.GetErrorData() != nil
	})

	// Filter out messages this validator has already attested to
	msgs = slice.Filter(msgs, func(msg types.QueuedSignedMessageI) bool {
		for _, evidence := range msg.GetEvidence() {
			if evidence.ValAddress.Equals(valAddress) {
				return false
			}
		}

		return true
	})

	if len(msgs) > defaultResponseMessageCount {
		msgs = msgs[:defaultResponseMessageCount]
	}

	return msgs, nil
}

// GetMessagesFromQueue gets N messages from the queue.
func (k Keeper) GetMessagesFromQueue(ctx sdk.Context, queueTypeName string, n int) (msgs []types.QueuedSignedMessageI, err error) {
	cq, err := k.getConsensusQueue(ctx, queueTypeName)
	if err != nil {
		k.Logger(ctx).Error("error while getting consensus queue", "err", err)
		return nil, err
	}
	msgs, err = cq.GetAll(ctx)

	if err != nil {
		k.Logger(ctx).Error("error while getting all messages from queue", "err", err)
		return nil, err
	}

	if n > 0 && len(msgs) > n {
		msgs = msgs[:n]
	}

	return
}

func (k Keeper) DeleteJob(ctx sdk.Context, queueTypeName string, id uint64) (err error) {
	cq, err := k.getConsensusQueue(ctx, queueTypeName)
	if err != nil {
		k.Logger(ctx).Error("error while getting consensus queue", "err", err)
		return err
	}
	return cq.Remove(ctx, id)
}

// GetMessagesThatHaveReachedConsensus returns messages from a given
// queueTypeName that have reached consensus based on the latest snapshot
// available.
func (k Keeper) GetMessagesThatHaveReachedConsensus(ctx sdk.Context, queueTypeName string) ([]types.QueuedSignedMessageI, error) {
	var consensusReached []types.QueuedSignedMessageI

	err := whoops.Try(func() {
		cq, err := k.getConsensusQueue(ctx, queueTypeName)
		whoops.Assert(err)

		msgs := whoops.Must(cq.GetAll(ctx))
		if len(msgs) == 0 {
			return
		}
		snapshot := whoops.Must(k.valset.GetCurrentSnapshot(ctx))

		if len(snapshot.Validators) == 0 || snapshot.TotalShares.Equal(sdk.ZeroInt()) {
			return
		}

		validatorMap := make(map[string]valsettypes.Validator)
		for _, validator := range snapshot.GetValidators() {
			validatorMap[validator.Address.String()] = validator
		}

		for _, msg := range msgs {
			msgTotal := sdk.ZeroInt()
			// add shares of validators that have signed the message
			for _, signData := range msg.GetSignData() {
				signedValidator, ok := validatorMap[signData.ValAddress.String()]
				if !ok {
					k.Logger(ctx).Info("validator not found", "validator", signData.ValAddress)
					continue
				}
				msgTotal = msgTotal.Add(signedValidator.ShareCount)
			}

			// Now we need to check if the consensus was reached. We do this by
			// checking if there were at least 2/3 of total signatures in
			// staking power.
			// The formula goes like this:
			// msgTotal >= 2/3 * snapshotTotal
			// the "issue" now becomes: 2/3 which could be problematic as we
			// could lose precision using floating point arithmetic.
			// If we multiply both sides with 3, we don't need to do division.
			// 3 * msgTotal >= 2 * snapshotTotal
			if msgTotal.Mul(sdk.NewInt(3)).GTE(snapshot.TotalShares.Mul(sdk.NewInt(2))) {
				// consensus has been reached
				consensusReached = append(consensusReached, msg)
			}
		}
	})
	if err != nil {
		return nil, err
	}
	return consensusReached, nil
}

// AddMessageSignature adds signatures to the messages.
func (k Keeper) AddMessageSignature(
	ctx sdk.Context,
	valAddr sdk.ValAddress,
	msgs []*types.ConsensusMessageSignature,
) error {
	err := whoops.Try(func() {
		for _, msg := range msgs {
			cq := whoops.Must(
				k.getConsensusQueue(ctx, msg.GetQueueTypeName()),
			)
			chainType, chainReferenceID := cq.ChainInfo()

			publicKey := whoops.Must(k.valset.GetSigningKey(
				ctx,
				valAddr,
				chainType,
				chainReferenceID,
				msg.GetSignedByAddress(),
			))

			whoops.Assert(
				cq.AddSignature(
					ctx,
					msg.Id,
					&types.SignData{
						ValAddress:             valAddr,
						Signature:              msg.GetSignature(),
						ExternalAccountAddress: msg.GetSignedByAddress(),
						PublicKey:              publicKey,
					},
				),
			)

			k.Logger(ctx).Info("added message signature",
				"message-id", msg.GetId(),
				"queue-type-name", msg.GetQueueTypeName(),
				"signed-by-address", msg.GetSignedByAddress(),
				"chain-type", chainType,
				"chain-reference-id", chainReferenceID,
			)
		}
	})
	if err != nil {
		k.Logger(ctx).Error("error while adding messages signatures",
			"err", err,
		)
	}

	return err
}

func (k Keeper) AddMessageEvidence(
	ctx sdk.Context,
	valAddr sdk.ValAddress,
	msg *types.MsgAddEvidence,
) error {
	err := whoops.Try(func() {
		cq := whoops.Must(
			k.getConsensusQueue(ctx, msg.GetQueueTypeName()),
		)

		whoops.Assert(
			cq.AddEvidence(
				ctx,
				msg.GetMessageID(),
				&types.Evidence{
					ValAddress: valAddr,
					Proof:      msg.GetProof(),
				},
			),
		)
		chainType, chainReferenceID := cq.ChainInfo()
		k.Logger(ctx).Info("added message evidence",
			"message-id", msg.GetMessageID(),
			"queue-type-name", msg.GetQueueTypeName(),
			"chain-type", chainType,
			"chain-reference-id", chainReferenceID,
		)
	})
	if err != nil {
		k.Logger(ctx).Error("error while adding message evidence",
			"err", err,
		)
	}

	return err
}

func (k Keeper) SetMessagePublicAccessData(
	ctx sdk.Context,
	valAddr sdk.ValAddress,
	msg *types.MsgSetPublicAccessData,
) error {
	cq, err := k.getConsensusQueue(ctx, msg.GetQueueTypeName())
	if err != nil {
		return err
	}

	payload := &types.PublicAccessData{
		ValAddress: valAddr,
		Data:       msg.GetData(),
	}
	err = cq.SetPublicAccessData(ctx, msg.GetMessageID(), payload)
	if err != nil {
		k.Logger(ctx).Error("error while adding message public access data", "err", err)
		return err
	}

	chainType, chainReferenceID := cq.ChainInfo()
	k.Logger(ctx).Info("added message public access data",
		"message-id", msg.GetMessageID(),
		"queue-type-name", msg.GetQueueTypeName(),
		"chain-type", chainType,
		"chain-reference-id", chainReferenceID,
		"public-access-data", hexutil.Encode(payload.Data),
	)

	return nil
}

func (k Keeper) SetMessageErrorData(
	ctx sdk.Context,
	valAddr sdk.ValAddress,
	msg *types.MsgSetErrorData,
) error {
	cq, err := k.getConsensusQueue(ctx, msg.GetQueueTypeName())
	if err != nil {
		return err
	}

	payload := &types.ErrorData{
		ValAddress: valAddr,
		Data:       msg.GetData(),
	}
	err = cq.SetErrorData(ctx, msg.GetMessageID(), payload)
	if err != nil {
		k.Logger(ctx).Error("error while adding error data", "err", err)
		return err
	}

	chainType, chainReferenceID := cq.ChainInfo()
	k.Logger(ctx).Info("added error data",
		"message-id", msg.GetMessageID(),
		"queue-type-name", msg.GetQueueTypeName(),
		"chain-type", chainType,
		"chain-reference-id", chainReferenceID,
		"error-data", hexutil.Encode(payload.Data),
	)

	return nil
}

func (k Keeper) reassignMessageValidator(
	ctx sdk.Context,
	valAddr string,
	msgID uint64,
	queueTypeName string,
) error {
	cq, err := k.getConsensusQueue(ctx, queueTypeName)
	if err != nil {
		return err
	}

	chainType, chainReferenceID := cq.ChainInfo()
	k.Logger(ctx).Info("reassigning orphaned message",
		"message-id", msgID,
		"queue-type-name", queueTypeName,
		"chain-type", chainType,
		"chain-reference-id", chainReferenceID,
		"new-assignee", valAddr,
	)

	return cq.ReassignValidator(ctx, msgID, valAddr)
}

func nonceFromID(id uint64) []byte {
	return sdk.Uint64ToBigEndian(id)
}

func (k Keeper) queuedMessageToMessageToSign(msg types.QueuedSignedMessageI) *types.MessageToSign {
	consensusMsg, err := msg.ConsensusMsg(k.cdc)
	if err != nil {
		panic(err)
	}
	anyMsg, err := codectypes.NewAnyWithValue(consensusMsg)
	if err != nil {
		panic(err)
	}
	return &types.MessageToSign{
		Nonce:       nonceFromID(msg.GetId()),
		Id:          msg.GetId(),
		BytesToSign: msg.GetBytesToSign(),
		Msg:         anyMsg,
	}
}

func (k Keeper) queuedMessageToMessageWithSignatures(msg types.QueuedSignedMessageI) (types.MessageWithSignatures, error) {
	consensusMsg, err := msg.ConsensusMsg(k.cdc)
	if err != nil {
		return types.MessageWithSignatures{}, err
	}
	anyMsg, err := codectypes.NewAnyWithValue(consensusMsg)
	if err != nil {
		return types.MessageWithSignatures{}, err
	}

	var publicAccessData []byte

	if msg.GetPublicAccessData() != nil {
		publicAccessData = msg.GetPublicAccessData().GetData()
	}

	var errorData []byte

	if msg.GetErrorData() != nil {
		errorData = msg.GetErrorData().GetData()
	}

	respMsg := types.MessageWithSignatures{
		Nonce:            nonceFromID(msg.GetId()),
		Id:               msg.GetId(),
		BytesToSign:      msg.GetBytesToSign(),
		Msg:              anyMsg,
		PublicAccessData: publicAccessData,
		ErrorData:        errorData,
	}

	for _, signData := range msg.GetSignData() {
		respMsg.SignData = append(respMsg.SignData, &types.ValidatorSignature{
			ValAddress:             signData.GetValAddress(),
			Signature:              signData.GetSignature(),
			ExternalAccountAddress: signData.GetExternalAccountAddress(),
			PublicKey:              signData.GetPublicKey(),
		})
	}

	return respMsg, nil
}
