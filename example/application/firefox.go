package application

import (
	"fmt"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
)

// GetFirefoxConfiguration gets firefox configuration
func GetFirefoxConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	result, response, err := app.GetFirefoxWebConf(&pushservicev1.GetFirefoxWebConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.WebSiteURL)
	}
}

// SaveFirefoxConfiguration save firefox configuration
func SaveFirefoxConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	webSiteURL := ""
	result, response, err := app.SaveFirefoxWebConf(&pushservicev1.SaveFirefoxWebConfOptions{
		ApplicationID: &appID,
		WebSiteURL:    &webSiteURL,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.WebSiteURL)
	}
}

// DeleteFirefoxConfiguration deletes firefox configuration
func DeleteFirefoxConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	response, err := app.DeleteFirefoxWebConf(&pushservicev1.DeleteFirefoxWebConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 204 {
		fmt.Println("Deleted")
	}
}
