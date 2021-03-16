package application

import (
	"fmt"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
)

// GetChromeExtConfiguration gets chrome apps and extention configuration
func GetChromeExtConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	result, response, err := app.GetChromeAppExtConf(&pushservicev1.GetChromeAppExtConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.ApiKey)
		fmt.Println(*result.SenderID)
	}
}

// SaveChromeExtConfiguration saves chrome apps and extention configuration
func SaveChromeExtConfiguration(appID string, app *pushservicev1.PushServiceV1) {
	apiKey := ""
	senderID := ""

	result, response, err := app.SaveChromeAppExtConf(&pushservicev1.SaveChromeAppExtConfOptions{
		ApplicationID: &appID,
		ApiKey:        &apiKey,
		SenderID:      &senderID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.ApiKey)
		fmt.Println(*result.SenderID)
	}
}

// DeleteChromeExtConfiguration delete chrome apps and extention configuration
func DeleteChromeExtConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	response, err := app.DeleteChromeAppExtConf(&pushservicev1.DeleteChromeAppExtConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 204 {
		fmt.Println("Deleted")
	}
}

// GetChromeExtConfigurationPublic gets chrome apps and extention publiic configuration
func GetChromeExtConfigurationPublic(appID string, clientSecret string, app *pushservicev1.PushServiceV1) {

	result, response, err := app.GetChromeAppExtConfPublic(&pushservicev1.GetChromeAppExtConfPublicOptions{
		ApplicationID: &appID,
		ClientSecret:  &clientSecret,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.SenderID)
	}
}
