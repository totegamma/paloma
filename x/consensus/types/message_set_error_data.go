package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/palomachain/paloma/util/libmeta"
)

const TypeMsgSetErrorData = "set_error_data"

var _ sdk.Msg = &MsgSetErrorData{}

func (msg *MsgSetErrorData) Route() string {
	return RouterKey
}

func (msg *MsgSetErrorData) Type() string {
	return TypeMsgSetErrorData
}

func (msg *MsgSetErrorData) GetSigners() []sdk.AccAddress {
	return libmeta.GetSigners(msg)
}

func (msg *MsgSetErrorData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetErrorData) ValidateBasic() error {
	return libmeta.ValidateBasic(msg)
}
