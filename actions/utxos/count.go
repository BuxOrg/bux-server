package utxos

import (
	"net/http"

	"github.com/BuxOrg/bux"
	"github.com/BuxOrg/bux-server/actions"
	"github.com/julienschmidt/httprouter"
	apirouter "github.com/mrz1836/go-api-router"
)

// count will count all the utxos fulfilling the given conditions
func (a *Action) count(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	reqXPubID, _ := bux.GetXpubIDFromRequest(req)

	// Parse the params
	params := apirouter.GetParams(req)
	_, metadata, conditions, err := actions.GetQueryParameters(params)
	if err != nil {
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, err.Error())
		return
	}

	var dbConditions = map[string]interface{}{}
	if conditions != nil {
		dbConditions = *conditions
	}
	// force the xpub_id of the current user on query
	dbConditions["xpub_id"] = reqXPubID

	// Get a utxo using a xPub
	var count int64
	if count, err = a.Services.Bux.GetUtxosCount(
		req.Context(),
		metadata,
		&dbConditions,
	); err != nil {
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, err.Error())
		return
	}

	// Return response
	apirouter.ReturnResponse(w, req, http.StatusOK, count)
}
