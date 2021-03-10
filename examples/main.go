package main

import (
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/push-notifications-go-sdk/examples/configureservice"
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

	configureservice.SaveGcmConfiguration(appID, app)

	configureservice.GetGcmConfiguration(appID, app)

	configureservice.GetGcmConfigurationPublic(appID, clientSecret, app)

	configureservice.SaveFirefoxConfiguration(appID, app)

	configureservice.GetFirefoxConfiguration(appID, app)

	configureservice.SaveChromewebConfiguration(appID, app)

	configureservice.GetChromewebConfiguration(appID, app)

	configureservice.GetWebPushServerKey(appID, clientSecret, app)

	configureservice.SaveChromeExtConfiguration(appID, app)

	configureservice.GetChromeExtConfiguration(appID, app)

	configureservice.GetChromeExtConfigurationPublic(appID, clientSecret, app)

	configureservice.SaveApnsConfiguration(appID, app)

	configureservice.GetApnsConfiguration(appID, app)

	configureservice.SaveSafariConfiguration(appID, app)

	configureservice.GetSafariConfiguration(appID, app)

	getSettings(appID, app)

	configureservice.DeleteGcmConfiguration(appID, app)

	configureservice.DeleteFirefoxConfiguration(appID, app)

	configureservice.DeleteChromewebConfiguration(appID, app)

	configureservice.DeleteChromeExtConfiguration(appID, app)

	configureservice.DeleteApnsConfiguration(appID, app)

	configureservice.DeleteSafariConfiguration(appID, app)

}
