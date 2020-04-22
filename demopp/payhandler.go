package demopp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssnclient"
)

func payHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract the URL encoded values from the request body
	in, err := ioutil.ReadAll(r.Body)
	ssn.Log(err, "payHandler: Read request body")
	req, err := url.ParseQuery(string(in))
	ssn.Log(err, "payHandler: Parse request body")

	// Step 3 - Verify Trust
	trusted, err := ssnclient.VerifyTrust(req.Get("payment_destination"), req.Get("asset_code"), assetIssuer, ssnAPI)
	if trusted && !ssn.Log(err, "payHandler: Verify trust") {
		// Step 5 - Build and sign transaction
		// Build Txn
		envelope, err := ssnclient.CreatePayment(assetIssuer, req.Get("payment_destination"), req.Get("amount"), req.Get("asset_code"), assetIssuer, req.Get("memo"), ssnAPI)
		ssn.Log(err, "payHandler: Create payment")

		// Sign Txn
		txn, err := ssnclient.SignTxnService(envelope, sig1)
		ssn.Log(err, "payHandler: Sign service 1")

		txn, err = ssnclient.SignTxnService(txn, sig2)
		ssn.Log(err, "payHandler: Sign service 2")

		// Submit Txn
		hash, err := ssnclient.SubmitTxn(txn, ssnAPI)
		ssn.Log(err, "payHandler: Submit txn")
		fmt.Println(hash)

		resp := ssn.RedirectPayload{
			RedirectURL: req.Get("redirect"),
			Payload: []ssn.PayloadItem{
				ssn.PayloadItem{
					Key:   "hash",
					Value: hash,
				},
			},
		}
		redirectTemplate := template.Must(template.ParseFiles("templates/redirect.html"))
		redirectTemplate.Execute(w, resp)
	} else {
		// Error
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong"))
	}
}
