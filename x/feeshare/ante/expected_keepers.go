package ante

// Used for the Juno ante handler so we can properly send 50% of fees to dAPP developers via fee share module

import (
	revtypes "github.com/CosmosContracts/juno/v12/x/feeshare/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BankKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type FeeShareKeeper interface {
	GetParams(ctx sdk.Context) revtypes.Params
	GetFeeShare(ctx sdk.Context, contract sdk.Address) (revtypes.FeeShare, bool)
}
