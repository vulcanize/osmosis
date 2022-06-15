package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
)

func newSetSuperfluidAssetsHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func newRemoveSuperfluidAssetsHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
