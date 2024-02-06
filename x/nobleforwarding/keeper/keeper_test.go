package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/settlus/chain/testutil/keeper"
	"github.com/settlus/chain/x/nobleforwarding/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NobleForwardingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
