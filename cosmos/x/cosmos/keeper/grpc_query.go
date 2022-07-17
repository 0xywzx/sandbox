package keeper

import (
	"github.com/0xywzx/cosmos/x/cosmos/types"
)

var _ types.QueryServer = Keeper{}
