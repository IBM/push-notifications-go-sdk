package application

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
)

// GetSafariConfiguration gets safari configuration
func GetSafariConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	result, response, err := app.GetSafariWebConf(&pushservicev1.GetSafariWebConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.Certificate)
		fmt.Println(*result.UrlFormatString)
		fmt.Println(*result.WebsiteName)
		fmt.Println(result.WebsitePushID)
		fmt.Println(result.WebSiteURL)
	}
}

// SaveSafariConfiguration saves safari configuration
func SaveSafariConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	password := ""
	websiteName := ""
	websitePushID := ""
	webSiteURL := ""
	urlFormatString := ""

	fileDir, _ := os.Getwd()
	fileName := ""
	filePath := path.Join(fileDir, "", fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	certiificateData := ioutil.NopCloser(file)

	contentType := "multipart/form-data"

	result, response, err := app.SaveSafariWebConf(&pushservicev1.SaveSafariWebConfOptions{
		ApplicationID:          &appID,
		Password:               &password,
		WebsiteName:            &websiteName,
		WebsitePushID:          &websitePushID,
		WebSiteURL:             &webSiteURL,
		UrlFormatString:        &urlFormatString,
		Certificate:            certiificateData,
		CertificateContentType: &contentType,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.Certificate)
		fmt.Println(*result.UrlFormatString)
		fmt.Println(*result.WebsiteName)
		fmt.Println(result.WebsitePushID)
		fmt.Println(result.WebSiteURL)
	}
}

// DeleteSafariConfiguration deletes safari configuration
func DeleteSafariConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	response, err := app.DeleteSafariWebConf(&pushservicev1.DeleteSafariWebConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 204 {
		fmt.Println("Deleted")
	}
}
