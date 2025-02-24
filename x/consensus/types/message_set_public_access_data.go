package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/palomachain/paloma/util/libmeta"
)

const TypeMsgSetPublicAccessData = "set_public_access_data"

var _ sdk.Msg = &MsgSetPublicAccessData{}

func (msg *MsgSetPublicAccessData) Route() string {
	return RouterKey
}

func (msg *MsgSetPublicAccessData) Type() string {
	return TypeMsgSetPublicAccessData
}

func (msg *MsgSetPublicAccessData) GetSigners() []sdk.AccAddress {
	return libmeta.GetSigners(msg)
}

func (msg *MsgSetPublicAccessData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPublicAccessData) ValidateBasic() error {
	return libmeta.ValidateBasic(msg)
}
