package demopp

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssnclient"
)

type AddPreauthResponse struct {
	Currencies       []string
	Service_user_key string
	Service_key      string
	Service_name     string
	Redirect         string
}

func addPreauthHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract the URL encoded values from the request body
	in, err := ioutil.ReadAll(r.Body)
	ssn.Log(err, "addPreauthHandler: Read request body")
	req, err := url.ParseQuery(string(in))
	ssn.Log(err, "addPreauthHandler: Parse URL encoded values")

	// Verify Request Signature
	reqURL := "https://" + r.Host + r.RequestURI
	reqMesg := sha256.New()
	reqMesg.Write([]byte(reqURL))
	reqHash := hex.EncodeToString(reqMesg.Sum(nil))
	if reqHash == req.Get("hash") && req.Get("public_key") == ps.ByName("pk") {
		sigVerified, err := ssnclient.VerifySignature(req.Get("hash"), req.Get("signature"), req.Get("public_key"), ssnAPI)
		if ssn.Log(err, "addPreauthHandler: Verify signature") {
			// Error
			fmt.Println("VerifySignature function returned error")
		} else if !sigVerified {
			// Error
			fmt.Println("Signature invalid")
		} else {
			// Should look up trust lines to MK and compare with PP customers accounts
			ccy := make([]string, 0)
			trust, err := ssnclient.VerifyTrust(ps.ByName("mk"), "USD", assetIssuer, ssnAPI)
			ccy = append(ccy, "USD")
			if ssn.Log(err, "addPreauthHandler: Verify trust") {
				// Error
				fmt.Println("VerifyTrust function returned error")
			} else if !trust {
				// Error
				fmt.Println("Trust missing")
			} else {
				service, err := ssnclient.GetServiceName(ps.ByName("mk"), ssnAPI)
				ssn.Log(err, "addPreauthHandler: Get service name")

				// Execute template
				response := AddPreauthResponse{
					Currencies:       ccy,
					Service_user_key: ps.ByName("pk"),
					Service_key:      ps.ByName("mk"),
					Service_name:     service,
					Redirect:         req.Get("redirect"),
				}

				preAuthTemplate := template.Must(template.ParseFiles("templates/addpreauth.html"))
				preAuthTemplate.Execute(w, response)
			}
		}
	}
}
