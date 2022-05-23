package keeper_test

import (
	"context"
	"testing"

	keepertest "githhub.com/reemardelarosa/checkers/testutil/keeper"
	"githhub.com/reemardelarosa/checkers/x/checkers/keeper"
	"githhub.com/reemardelarosa/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
