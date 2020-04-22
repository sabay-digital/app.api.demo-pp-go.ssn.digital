package demopp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"
)

func authHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract the URL encoded values from the request body
	in, err := ioutil.ReadAll(r.Body)
	ssn.Log(err, "onetimeChargeHandler: Read request body")
	req, err := url.ParseQuery(string(in))
	ssn.Log(err, "onetimeChargeHandler: Parse URL encoded values")

	var existingPreauth []PreAuthorization
	db.Where("user_pk = ? AND service_pk = ?", req.Get("service_user_key"), req.Get("service_key")).Find(&existingPreauth)
	fmt.Println(len(existingPreauth))
	fmt.Println(existingPreauth)

	if len(existingPreauth) == 0 {
		newPreAuth := PreAuthorization{
			UserPubkey:    req.Get("service_user_key"),
			ServicePubkey: req.Get("service_key"),
			Currencies:    fmt.Sprintf("%s", req.Get("currencies")),
		}
		db.Create(&newPreAuth)
	}

	if len(req.Get("redirect")) == 0 {
		out := ssn.PayloadItem{
			Value: req.Get("service_user_key"),
		}
		redirectTemplate := template.Must(template.ParseFiles("templates/authsuccess.html"))
		redirectTemplate.Execute(w, out)
	} else {
		resp := ssn.RedirectPayload{
			RedirectURL: req.Get("redirect"),
			Payload: []ssn.PayloadItem{
				ssn.PayloadItem{
					Key:   "pubkey",
					Value: req.Get("service_user_key"),
				},
			},
		}
		redirectTemplate := template.Must(template.ParseFiles("templates/redirect.html"))
		redirectTemplate.Execute(w, resp)
	}
}
