package main

import (
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/push-notifications-go-sdk/example/application"
	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
)

func getSettings(appID string, app *pushservicev1.PushServiceV1) {

	getSettingsOptions := &pushservicev1.GetSettingsOptions{
		ApplicationID: &appID,
	}

	result, response, err := app.GetSettings(getSettingsOptions)

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		if result.ApnsConf != nil {
			fmt.Println(*result.ApnsConf)
		}
		if result.GcmConf != nil {
			fmt.Println(*result.GcmConf)
		}
		if result.FirefoxWebConf != nil {
			fmt.Println(*result.FirefoxWebConf)
		}
		if result.ChromeWebConf != nil {
			fmt.Println(*result.ChromeWebConf)
		}
		if result.SafariWebConf != nil {
			fmt.Println(*result.SafariWebConf)
		}
	}
}

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

	getSettings(appID, app)

	application.DeleteGcmConfiguration(appID, app)

	application.DeleteFirefoxConfiguration(appID, app)

	application.DeleteChromewebConfiguration(appID, app)

	application.DeleteChromeExtConfiguration(appID, app)

	application.DeleteApnsConfiguration(appID, app)

	application.DeleteSafariConfiguration(appID, app)

}
