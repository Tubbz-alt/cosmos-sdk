package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgGrantAuthorization{}, "cosmos-sdk/GrantAuthorization", nil)
	cdc.RegisterConcrete(MsgRevokeAuthorization{}, "cosmos-sdk/RevokeAuthorization", nil)
	cdc.RegisterConcrete(MsgExecDelegated{}, "cosmos-sdk/ExecDelegated", nil)
	cdc.RegisterConcrete(SendCapability{}, "cosmos-sdk/SendCapability", nil)
	cdc.RegisterConcrete(CapabilityGrant{}, "cosmos-sdk/CapabilityGrant", nil)

	cdc.RegisterInterface((*Capability)(nil), nil)
}