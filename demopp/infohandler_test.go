package demopp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sabay-digital/sdk.golang.ssn.digital/ssnclient"
)

func TestInfohandler(t *testing.T) {
	Initialise(testSSNAPI, testSSNPAR, testIssuerSK, testSig01, testSig02, testDBHost, testDBName, testDBUsername, testDBPassword, "")
	srv := httptest.NewServer(Router())
	defer srv.Close()

	tt := []struct {
		name   string
		status int
		result string
	}{
		{"GET /v1/info", 200, "GARXYJB3ZUJ7DNQTDWYE5PWO356ZOXD26NQJZPPM7CDAQP5YPWAWHD36"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			apiResp := ssnclient.CashierInfoResponse{}

			req, err := http.NewRequest("GET", srv.URL+"/v1/info", nil)
			if err != nil {
				t.Fatalf("%v fails. Could not create request: %v", tc.name, err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("%v fails. Could not record response: %v", tc.name, err)
			}

			if res.StatusCode != tc.status {
				t.Fatalf("%v fails. Expected status code %v but got %v", tc.name, tc.status, res.StatusCode)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("%v fails. Error reading body: %v", tc.name, err)
			}

			jsonErr := json.Unmarshal(body, &apiResp)
			if err != nil {
				t.Fatalf("%v fails. Error unmarshalling json: %v", tc.name, jsonErr)
			}

			if tr := apiResp.Assets_issued[0].Asset_Issuer; tr != tc.result {
				t.Fatalf("%v fails. Expected %v but got %v", tc.name, tc.result, tr)
			}

		})
	}
}
