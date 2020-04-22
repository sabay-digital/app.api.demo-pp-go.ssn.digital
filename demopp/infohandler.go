package demopp

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssnclient"
)

func infoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	assets := []ssnclient.CashierIssuedAsset{
		ssnclient.CashierIssuedAsset{
			Asset_code:   "USD",
			Asset_Issuer: assetIssuer,
		},
	}

	apiResp := ssnclient.CashierInfoResponse{
		Assets_issued:    assets,
		Payment_provider: "Sabay Demo Payment Provider",
		Payment_type:     "onetime, pre-authorized",
		Authorization:    "ssn-signature",
	}

	resp, err := json.Marshal(&apiResp)
	ssn.Log(err, "Info handler: Marshal response")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resp)
}
