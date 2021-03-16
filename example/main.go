package main

import (
	"log"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/push-notifications-go-sdk/example/application"
	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
)

func main() {

	appID := ""

	clientSecret := ""

	authenticator := &core.IamAuthenticator{
		ApiKey: "",
	}

	options := &pushservicev1.PushServiceV1Options{
		ServiceName:   "imfpush",
		Authenticator: authenticator,
		URL:           "",
	}

	app, err := pushservicev1.NewPushServiceV1(options)

	if err != nil {
		log.Fatal(err)
	}

	application.SaveGcmConfiguration(appID, app)

	application.GetGcmConfiguration(appID, app)

	application.GetGcmConfigurationPublic(appID, clientSecret, app)

	application.SaveFirefoxConfiguration(appID, app)

	application.GetFirefoxConfiguration(appID, app)

	application.SaveChromewebConfiguration(appID, app)

	application.GetChromewebConfiguration(appID, app)

	application.GetWebPushServerKey(appID, clientSecret, app)

	application.SaveChromeExtConfiguration(appID, app)

	application.GetChromeExtConfiguration(appID, app)

	application.GetChromeExtConfigurationPublic(appID, clientSecret, app)

	application.SaveApnsConfiguration(appID, app)

	application.GetApnsConfiguration(appID, app)

	application.SaveSafariConfiguration(appID, app)

	application.GetSafariConfiguration(appID, app)

	application.DeleteGcmConfiguration(appID, app)

	application.DeleteFirefoxConfiguration(appID, app)

	application.DeleteChromewebConfiguration(appID, app)

	application.DeleteChromeExtConfiguration(appID, app)

	application.DeleteApnsConfiguration(appID, app)

	application.DeleteSafariConfiguration(appID, app)

}
