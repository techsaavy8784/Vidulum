package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBitcoinOwnership{}, "verify/BitcoinOwnership", nil)
	cdc.RegisterConcrete(&MsgSolanaOwnership{}, "verify/SolanaOwnership", nil)
	cdc.RegisterConcrete(&MsgEthereumOwnership{}, "verify/EthereumOwnership", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBitcoinOwnership{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSolanaOwnership{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEthereumOwnership{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
