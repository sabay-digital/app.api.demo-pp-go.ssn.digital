package demopp

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssnclient"
)

func statusHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	statusReq, err := http.NewRequest("GET", ssnAPI, nil)
	ssn.Log(err, "Status handler: Create HTTP request")

	apiResp := ssnclient.CashierStatusResponse{}

	statusResp, err := http.DefaultClient.Do(statusReq)
	if ssn.Log(err, "Status handler: Send HTTP request") {
		apiResp.Status = 500
		apiResp.Title = "System is offline and cannot process payments"
	} else {
		apiResp.Status = statusResp.StatusCode
		apiResp.Title = "Ready"
	}

	resp, err := json.Marshal(&apiResp)
	ssn.Log(err, "Status handler: Marshal response")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiResp.Status)
	w.Write(resp)
}
