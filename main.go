package main

import (
	"net/http"
	"os"

	"github.com/sabay-digital/sdk.golang.ssn.digital/ssn"

	"git.sabay.com/payment-network/cashiers/app.web.cashier.demo-pp.ssn.digital/demopp"
)

func init() {
	demopp.Initialise(os.Getenv("SSN_API"),
		os.Getenv("SSN_PA_RESOLVER"),
		os.Getenv("ISSUER_SK"),
		os.Getenv("CASHIER_SIGNER01"),
		os.Getenv("CASHIER_SIGNER02"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SLACK_WEBHOOK_URL"),
	)
}

func main() {
	err := http.ListenAndServe(":3000", demopp.Router())
	ssn.Log(err, "Main: Server error")
}
