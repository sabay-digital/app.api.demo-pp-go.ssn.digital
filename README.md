# app.api.demo-pp-go.ssn.digital

An example implementation of the SSN payment provider API written in Go

The current implementation covers the following endpoints from the [reference](https://api-reference.ssn.digital/?urls.primaryName=SSN%20Payment%20Provider%20API):

GET /status
GET /info

POST /authorize/{public_key}/{ssn_account}

POST /charge/auth/{payment_address}
POST /charge/onetime/{payment_address}