package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/0xywzx/cosmos/testutil/keeper"
	"github.com/0xywzx/cosmos/x/cosmos/keeper"
	"github.com/0xywzx/cosmos/x/cosmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmosKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
