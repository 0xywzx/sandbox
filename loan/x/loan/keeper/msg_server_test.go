package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/0xywzx/loan/testutil/keeper"
	"github.com/0xywzx/loan/x/loan/keeper"
	"github.com/0xywzx/loan/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LoanKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
