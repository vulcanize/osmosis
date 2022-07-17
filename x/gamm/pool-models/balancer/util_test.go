package balancer_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/db/memdb"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/osmosis-labs/osmosis/v9/x/gamm/pool-models/balancer"
	"github.com/osmosis-labs/osmosis/v9/x/gamm/types"
)

func createTestPool(t *testing.T, swapFee, exitFee sdk.Dec, poolAssets ...balancer.PoolAsset) types.PoolI {
	pool, err := balancer.NewBalancerPool(
		1,
		balancer.NewPoolParams(swapFee, exitFee, nil),
		poolAssets,
		"",
		time.Now(),
	)
	require.NoError(t, err)

	return &pool
}

func createTestContext(t *testing.T) sdk.Context {
	ms := store.NewCommitMultiStore(memdb.NewDB())
	return sdk.NewContext(ms, tmtypes.Header{}, false, log.NewNopLogger())
}
