package application

import (
	"fmt"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
)

// GetGcmConfiguration gets gcm configuration
func GetGcmConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	result, response, err := app.GetGCMConf(&pushservicev1.GetGCMConfOptions{
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

// SaveGcmConfiguration saves gcm configuration
func SaveGcmConfiguration(appID string, app *pushservicev1.PushServiceV1) {
	apiKey := ""
	senderID := ""
	result, response, err := app.SaveGCMConf(&pushservicev1.SaveGCMConfOptions{
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

// DeleteGcmConfiguration deletes gcm configuration
func DeleteGcmConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	response, err := app.DeleteGCMConf(&pushservicev1.DeleteGCMConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 204 {
		fmt.Println("Deleted")
	}
}

// GetGcmConfigurationPublic get gcm configuration public
func GetGcmConfigurationPublic(appID string, clientSecret string, app *pushservicev1.PushServiceV1) {

	result, response, err := app.GetGcmConfPublic(&pushservicev1.GetGcmConfPublicOptions{
		ApplicationID: &appID,
		ClientSecret:  &clientSecret,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		if response.StatusCode == 200 {
			fmt.Println(*result.SenderID)
		}
	}
}
