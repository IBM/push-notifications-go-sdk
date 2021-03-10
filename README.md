# IBM Cloud Mobile Services - Go server-side SDK for Push Notifications

The [IBM Cloud Push Notifications service](https://cloud.ibm.com/catalog/services/push-notifications) provides a unified push service to send real-time notifications to mobile and web applications. The Go SDK is used to manage Push Notifications service.

Ensure that you go through [IBM Cloud Push Notifications service documentation](https://cloud.ibm.com/docs/services/mobilepush?topic=mobile-pushnotification-gettingstartedtemplate#gettingstartedtemplate) before you start.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Import the SDK](#import-the-sdk)
- [Initialize SDK](#initialize-sdk)
- [Using the SDK](#using-the-sdk)
- [License](#license)

## Prerequisites

- An [IBM Cloud](https://cloud.ibm.com/registration) account.
- An [Push Notifications](https://cloud.ibm.com/docs/mobilepush) instance.
- Go version 1.15 or above.

## Installation

Install using the command.

```bash
go get -u github.com/IBM/push-notifications-go-sdk
```

## Import the SDK

To import the module

```go
import "github.com/IBM/push-notifications-go-sdk/pushservicev1"
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

## Initialize SDK

Initialize the sdk to connect with your App Configuration service instance.

```go
func init() {
  authenticator := &core.IamAuthenticator{
		ApiKey: "apikey",
	}

	options := &pushservicev1.PushServiceV1Options{
		ServiceName:   "imfpush",
		Authenticator: authenticator,
		URL:           "url",
	}

	app, err := pushservicev1.NewPushServiceV1(options)

	if err != nil {
		log.Fatal(err)
	}
}
```

- apikey : apikey of the Push notifications service. Get it from the service credentials section of the dashboard.
- url : url of the Push notifications Instance. URL instance can found from [here](https://cloud.ibm.com/apidocs/push-notifications#api-documentation-for-push-notifications)

## Using the SDK

Refer to the example directory

## License

This project is released under the Apache 2.0 license. The license's full text can be found in [LICENSE](/LICENSE)
