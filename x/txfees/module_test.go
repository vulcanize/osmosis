package txfees_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	simapp "github.com/osmosis-labs/osmosis/v9/app"
)

func TestSetBaseDenomOnInitBlock(t *testing.T) {
	app := simapp.Setup(t)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	baseDenom, err := app.TxFeesKeeper.GetBaseDenom(ctx)
	require.Nil(t, err)
	require.NotEmpty(t, baseDenom)
}
