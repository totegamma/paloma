package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/palomachain/paloma/testutil/sample"
	"github.com/palomachain/paloma/x/valset/types"
	"github.com/stretchr/testify/require"
)

func TestMsgSetPublicAccessData_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetPublicAccessData
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetPublicAccessData{
				Creator: "invalid_address",
				Metadata: types.MsgMetadata{
					Creator: "invalid_address",
					Signers: []string{"invalid_address"},
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetPublicAccessData{
				Creator: sample.AccAddress(),
				Metadata: types.MsgMetadata{
					Creator: sample.AccAddress(),
					Signers: []string{sample.AccAddress()},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
