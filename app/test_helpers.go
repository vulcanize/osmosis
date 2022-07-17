package app

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	dbm "github.com/cosmos/cosmos-sdk/db"
	"github.com/cosmos/cosmos-sdk/db/memdb"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// Setup initializes a new OsmosisApp. A Nop logger is set in SimApp.
func Setup(t testing.TB) *OsmosisApp {
	t.Helper()

	privVal := mock.NewPV()
	pubKey, err := privVal.GetPubKey(context.TODO())
	require.NoError(t, err)

	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// generate genesis account
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000))),
	}

	app := NewOsmosisApp(log.NewNopLogger(), memdb.NewDB(), nil, map[int64]bool{}, DefaultNodeHome, 5, MakeEncodingConfig(), simapp.EmptyAppOptions{}, GetWasmEnabledProposals(), EmptyWasmOpts)

	genesisState := NewDefaultGenesisState()
	genesisState = simapp.GenesisStateWithValSet(t, app.AppCodec(), genesisState, valSet, []authtypes.GenesisAccount{acc}, balance)
	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	if err != nil {
		panic(err)
	}

	app.InitChain(
		abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: simapp.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	return app
}

// SetupTestingAppWithBadgerDb initializes a new OsmosisApp intended for testing,
// with BadgerDB as a db.
func SetupTestingAppWithBadgerDb() (app *OsmosisApp, cleanupFn func()) {
	dir := "osmosis_testing"
	db, err := dbm.NewDB("osmosis_leveldb_testing", dbm.BadgerDBBackend, dir)
	if err != nil {
		panic(err)
	}
	app = NewOsmosisApp(log.NewNopLogger(), db, nil, map[int64]bool{}, DefaultNodeHome, 5, MakeEncodingConfig(), simapp.EmptyAppOptions{}, GetWasmEnabledProposals(), EmptyWasmOpts)

	genesisState := NewDefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	if err != nil {
		panic(err)
	}

	app.InitChain(
		abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: simapp.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	cleanupFn = func() {
		db.Close()
		err = os.RemoveAll(dir)
		if err != nil {
			panic(err)
		}
	}

	return app, cleanupFn
}
