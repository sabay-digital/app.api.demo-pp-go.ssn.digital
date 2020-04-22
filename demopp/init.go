package demopp

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// DB Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stellar/go/keypair"
)

var (
	ssnAPI       string
	paResolver   string
	assetIssuer  string
	sig1         string
	sig2         string
	kp           keypair.KP
	slackWebhook string
	db           *gorm.DB
)

// PreAuthorization table model
type PreAuthorization struct {
	gorm.Model
	UserPubkey    string `gorm:"column:user_pk"`
	ServicePubkey string `gorm:"column:service_pk"`
	Currencies    string `gorm:"column:currencies"`
}

// Initialise ensures all package wide variables are correctly set at startup
func Initialise(apiURL, parURL, issuer, cashierSig1, cashierSig2, dbHost, dbName, dbUser, dbPassword, slackWH string) {
	ssnAPI = apiURL
	paResolver = parURL
	kp = keypair.MustParse(issuer)
	assetIssuer = kp.Address()
	sig1 = cashierSig1
	sig2 = cashierSig2

	slackWebhook = slackWH

	db, _ = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName))

	db.AutoMigrate(&PreAuthorization{})
}
