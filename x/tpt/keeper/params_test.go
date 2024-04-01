package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "tpt/testutil/keeper"
	"tpt/x/tpt/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TptKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
