package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/walload/hbble/testutil/keeper"
	"github.com/walload/hbble/x/hbble/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HbbleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
