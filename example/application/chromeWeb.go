package application

import (
	"fmt"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
)

// GetChromewebConfiguration gets chrome web configuration
func GetChromewebConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	result, response, err := app.GetChromeWebConf(&pushservicev1.GetChromeWebConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.ApiKey)
		fmt.Println(*result.WebSiteURL)
	}
}

// SaveChromewebConfiguration saves chrome web configuration
func SaveChromewebConfiguration(appID string, app *pushservicev1.PushServiceV1) {
	apiKey := ""
	webSiteURL := ""
	result, response, err := app.SaveChromeWebConf(&pushservicev1.SaveChromeWebConfOptions{
		ApplicationID: &appID,
		ApiKey:        &apiKey,
		WebSiteURL:    &webSiteURL,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.ApiKey)
		fmt.Println(*result.WebSiteURL)
	}
}

// DeleteChromewebConfiguration delete chrome web configuration
func DeleteChromewebConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	response, err := app.DeleteChromeWebConf(&pushservicev1.DeleteChromeWebConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 204 {
		fmt.Println("Deleted")
	}
}

// GetWebPushServerKey gets web push server key
func GetWebPushServerKey(appID string, clientSecret string, app *pushservicev1.PushServiceV1) {

	result, response, err := app.GetWebpushServerKey(&pushservicev1.GetWebpushServerKeyOptions{
		ApplicationID: &appID,
		ClientSecret:  &clientSecret,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.WebpushServerKey)
	}
}
