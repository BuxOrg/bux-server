package utxos

import (
	"net/http"

	"github.com/BuxOrg/bux"
	"github.com/BuxOrg/bux-server/actions"
	"github.com/julienschmidt/httprouter"
	apirouter "github.com/mrz1836/go-api-router"
)

// search will fetch a list of utxos filtered on conditions and metadata
func (a *Action) search(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	reqXPubID, _ := bux.GetXpubIDFromRequest(req)

	// Parse the params
	params := apirouter.GetParams(req)
	queryParams, metadata, conditions, err := actions.GetQueryParameters(params)
	if err != nil {
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, err.Error())
		return
	}

	// Record a new transaction (get the hex from parameters)a
	var utxos []*bux.Utxo
	if utxos, err = a.Services.Bux.GetUtxosByXpubID(
		req.Context(),
		reqXPubID,
		metadata,
		conditions,
		queryParams,
	); err != nil {
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, err.Error())
		return
	}

	// Return response
	apirouter.ReturnResponse(w, req, http.StatusOK, bux.DisplayModels(utxos))
}
