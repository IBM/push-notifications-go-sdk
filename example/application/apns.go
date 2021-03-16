package application

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
)

// GetApnsConfiguration gets apns configuration
func GetApnsConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	result, response, err := app.GetApnsConf(&pushservicev1.GetApnsConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.Certificate)
		fmt.Println(*result.IsSandBox)
		fmt.Println(*result.ValidUntil)
	}
}

// SaveApnsConfiguration saves apns configuration
func SaveApnsConfiguration(appID string, app *pushservicev1.PushServiceV1) {
	isSandBox := false
	password := ""

	fileDir, _ := os.Getwd()
	fileName := ""
	filePath := path.Join(fileDir, fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	certiificateData := ioutil.NopCloser(file)
	contentType := "multipart/form-data"

	result, response, err := app.SaveApnsConf(&pushservicev1.SaveApnsConfOptions{
		ApplicationID:          &appID,
		Password:               &password,
		IsSandBox:              &isSandBox,
		Certificate:            certiificateData,
		CertificateContentType: &contentType,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 200 {
		fmt.Println(*result.Certificate)
		fmt.Println(*result.IsSandBox)
		fmt.Println(*result.ValidUntil)
	}
}

// DeleteApnsConfiguration deletes apns configuration
func DeleteApnsConfiguration(appID string, app *pushservicev1.PushServiceV1) {

	response, err := app.DeleteApnsConf(&pushservicev1.DeleteApnsConfOptions{
		ApplicationID: &appID,
	})

	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode == 204 {
		fmt.Println("Deleted")
	}
}
