/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pushservicev1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`PushServiceV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(pushServiceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(pushServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
				URL: "https://pushservicev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(pushServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PUSH_SERVICE_URL":       "https://pushservicev1/api",
				"PUSH_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1UsingExternalConfig(&pushservicev1.PushServiceV1Options{})
				Expect(pushServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := pushServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != pushServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(pushServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(pushServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1UsingExternalConfig(&pushservicev1.PushServiceV1Options{
					URL: "https://testService/api",
				})
				Expect(pushServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := pushServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != pushServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(pushServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(pushServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1UsingExternalConfig(&pushservicev1.PushServiceV1Options{})
				err := pushServiceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := pushServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != pushServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(pushServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(pushServiceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PUSH_SERVICE_URL":       "https://pushservicev1/api",
				"PUSH_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			pushServiceService, serviceErr := pushservicev1.NewPushServiceV1UsingExternalConfig(&pushservicev1.PushServiceV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(pushServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PUSH_SERVICE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			pushServiceService, serviceErr := pushservicev1.NewPushServiceV1UsingExternalConfig(&pushservicev1.PushServiceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(pushServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = pushservicev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		getSettingsPath := "/apps/testString/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSettings with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(pushservicev1.GetSettingsOptions)
				getSettingsOptionsModel.ApplicationID = core.StringPtr("testString")
				getSettingsOptionsModel.AppSecret = core.StringPtr("testString")
				getSettingsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
		getSettingsPath := "/apps/testString/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apnsConf": "ApnsConf", "gcmConf": "GcmConf", "chromeWebConf": "ChromeWebConf", "safariWebConf": "SafariWebConf", "firefoxWebConf": "FirefoxWebConf"}`)
				}))
			})
			It(`Invoke GetSettings successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(pushservicev1.GetSettingsOptions)
				getSettingsOptionsModel.ApplicationID = core.StringPtr("testString")
				getSettingsOptionsModel.AppSecret = core.StringPtr("testString")
				getSettingsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apnsConf": "ApnsConf", "gcmConf": "GcmConf", "chromeWebConf": "ChromeWebConf", "safariWebConf": "SafariWebConf", "firefoxWebConf": "FirefoxWebConf"}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(pushservicev1.GetSettingsOptions)
				getSettingsOptionsModel.ApplicationID = core.StringPtr("testString")
				getSettingsOptionsModel.AppSecret = core.StringPtr("testString")
				getSettingsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSettings with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(pushservicev1.GetSettingsOptions)
				getSettingsOptionsModel.ApplicationID = core.StringPtr("testString")
				getSettingsOptionsModel.AppSecret = core.StringPtr("testString")
				getSettingsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSettingsOptions model with no property values
				getSettingsOptionsModelNew := new(pushservicev1.GetSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.GetSettings(getSettingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetApnsConf(getApnsConfOptions *GetApnsConfOptions) - Operation response error`, func() {
		getApnsConfPath := "/apps/testString/settings/apnsConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApnsConfPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetApnsConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetApnsConfOptions model
				getApnsConfOptionsModel := new(pushservicev1.GetApnsConfOptions)
				getApnsConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getApnsConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getApnsConfOptionsModel.AppSecret = core.StringPtr("testString")
				getApnsConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.GetApnsConf(getApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.GetApnsConf(getApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetApnsConf(getApnsConfOptions *GetApnsConfOptions)`, func() {
		getApnsConfPath := "/apps/testString/settings/apnsConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApnsConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificate": "Certificate", "isSandBox": false, "validUntil": "ValidUntil"}`)
				}))
			})
			It(`Invoke GetApnsConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetApnsConfOptions model
				getApnsConfOptionsModel := new(pushservicev1.GetApnsConfOptions)
				getApnsConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getApnsConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getApnsConfOptionsModel.AppSecret = core.StringPtr("testString")
				getApnsConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.GetApnsConfWithContext(ctx, getApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.GetApnsConf(getApnsConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.GetApnsConfWithContext(ctx, getApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getApnsConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificate": "Certificate", "isSandBox": false, "validUntil": "ValidUntil"}`)
				}))
			})
			It(`Invoke GetApnsConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.GetApnsConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetApnsConfOptions model
				getApnsConfOptionsModel := new(pushservicev1.GetApnsConfOptions)
				getApnsConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getApnsConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getApnsConfOptionsModel.AppSecret = core.StringPtr("testString")
				getApnsConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.GetApnsConf(getApnsConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetApnsConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetApnsConfOptions model
				getApnsConfOptionsModel := new(pushservicev1.GetApnsConfOptions)
				getApnsConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getApnsConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getApnsConfOptionsModel.AppSecret = core.StringPtr("testString")
				getApnsConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.GetApnsConf(getApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetApnsConfOptions model with no property values
				getApnsConfOptionsModelNew := new(pushservicev1.GetApnsConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.GetApnsConf(getApnsConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SaveApnsConf(saveApnsConfOptions *SaveApnsConfOptions) - Operation response error`, func() {
		saveApnsConfPath := "/apps/testString/settings/apnsConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveApnsConfPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SaveApnsConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveApnsConfOptions model
				saveApnsConfOptionsModel := new(pushservicev1.SaveApnsConfOptions)
				saveApnsConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveApnsConfOptionsModel.Password = core.StringPtr("testString")
				saveApnsConfOptionsModel.IsSandBox = core.BoolPtr(true)
				saveApnsConfOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				saveApnsConfOptionsModel.CertificateContentType = core.StringPtr("testString")
				saveApnsConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveApnsConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveApnsConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.SaveApnsConf(saveApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.SaveApnsConf(saveApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SaveApnsConf(saveApnsConfOptions *SaveApnsConfOptions)`, func() {
		saveApnsConfPath := "/apps/testString/settings/apnsConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveApnsConfPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificate": "Certificate", "isSandBox": false, "validUntil": "ValidUntil"}`)
				}))
			})
			It(`Invoke SaveApnsConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the SaveApnsConfOptions model
				saveApnsConfOptionsModel := new(pushservicev1.SaveApnsConfOptions)
				saveApnsConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveApnsConfOptionsModel.Password = core.StringPtr("testString")
				saveApnsConfOptionsModel.IsSandBox = core.BoolPtr(true)
				saveApnsConfOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				saveApnsConfOptionsModel.CertificateContentType = core.StringPtr("testString")
				saveApnsConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveApnsConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveApnsConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.SaveApnsConfWithContext(ctx, saveApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.SaveApnsConf(saveApnsConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.SaveApnsConfWithContext(ctx, saveApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveApnsConfPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificate": "Certificate", "isSandBox": false, "validUntil": "ValidUntil"}`)
				}))
			})
			It(`Invoke SaveApnsConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.SaveApnsConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SaveApnsConfOptions model
				saveApnsConfOptionsModel := new(pushservicev1.SaveApnsConfOptions)
				saveApnsConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveApnsConfOptionsModel.Password = core.StringPtr("testString")
				saveApnsConfOptionsModel.IsSandBox = core.BoolPtr(true)
				saveApnsConfOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				saveApnsConfOptionsModel.CertificateContentType = core.StringPtr("testString")
				saveApnsConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveApnsConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveApnsConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.SaveApnsConf(saveApnsConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SaveApnsConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveApnsConfOptions model
				saveApnsConfOptionsModel := new(pushservicev1.SaveApnsConfOptions)
				saveApnsConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveApnsConfOptionsModel.Password = core.StringPtr("testString")
				saveApnsConfOptionsModel.IsSandBox = core.BoolPtr(true)
				saveApnsConfOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				saveApnsConfOptionsModel.CertificateContentType = core.StringPtr("testString")
				saveApnsConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveApnsConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveApnsConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.SaveApnsConf(saveApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SaveApnsConfOptions model with no property values
				saveApnsConfOptionsModelNew := new(pushservicev1.SaveApnsConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.SaveApnsConf(saveApnsConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteApnsConf(deleteApnsConfOptions *DeleteApnsConfOptions)`, func() {
		deleteApnsConfPath := "/apps/testString/settings/apnsConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteApnsConfPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteApnsConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := pushServiceService.DeleteApnsConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteApnsConfOptions model
				deleteApnsConfOptionsModel := new(pushservicev1.DeleteApnsConfOptions)
				deleteApnsConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteApnsConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteApnsConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteApnsConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = pushServiceService.DeleteApnsConf(deleteApnsConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteApnsConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the DeleteApnsConfOptions model
				deleteApnsConfOptionsModel := new(pushservicev1.DeleteApnsConfOptions)
				deleteApnsConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteApnsConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteApnsConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteApnsConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := pushServiceService.DeleteApnsConf(deleteApnsConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteApnsConfOptions model with no property values
				deleteApnsConfOptionsModelNew := new(pushservicev1.DeleteApnsConfOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = pushServiceService.DeleteApnsConf(deleteApnsConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGCMConf(getGCMConfOptions *GetGCMConfOptions) - Operation response error`, func() {
		getGcmConfPath := "/apps/testString/settings/gcmConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGcmConfPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGCMConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetGCMConfOptions model
				getGcmConfOptionsModel := new(pushservicev1.GetGCMConfOptions)
				getGcmConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getGcmConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getGcmConfOptionsModel.AppSecret = core.StringPtr("testString")
				getGcmConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.GetGCMConf(getGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.GetGCMConf(getGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetGCMConf(getGCMConfOptions *GetGCMConfOptions)`, func() {
		getGcmConfPath := "/apps/testString/settings/gcmConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGcmConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "senderId": "SenderID"}`)
				}))
			})
			It(`Invoke GetGCMConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetGCMConfOptions model
				getGcmConfOptionsModel := new(pushservicev1.GetGCMConfOptions)
				getGcmConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getGcmConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getGcmConfOptionsModel.AppSecret = core.StringPtr("testString")
				getGcmConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.GetGCMConfWithContext(ctx, getGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.GetGCMConf(getGcmConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.GetGCMConfWithContext(ctx, getGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGcmConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "senderId": "SenderID"}`)
				}))
			})
			It(`Invoke GetGCMConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.GetGCMConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGCMConfOptions model
				getGcmConfOptionsModel := new(pushservicev1.GetGCMConfOptions)
				getGcmConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getGcmConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getGcmConfOptionsModel.AppSecret = core.StringPtr("testString")
				getGcmConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.GetGCMConf(getGcmConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGCMConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetGCMConfOptions model
				getGcmConfOptionsModel := new(pushservicev1.GetGCMConfOptions)
				getGcmConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getGcmConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getGcmConfOptionsModel.AppSecret = core.StringPtr("testString")
				getGcmConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.GetGCMConf(getGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGCMConfOptions model with no property values
				getGcmConfOptionsModelNew := new(pushservicev1.GetGCMConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.GetGCMConf(getGcmConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SaveGCMConf(saveGCMConfOptions *SaveGCMConfOptions) - Operation response error`, func() {
		saveGcmConfPath := "/apps/testString/settings/gcmConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveGcmConfPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SaveGCMConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveGCMConfOptions model
				saveGcmConfOptionsModel := new(pushservicev1.SaveGCMConfOptions)
				saveGcmConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveGcmConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveGcmConfOptionsModel.SenderID = core.StringPtr("testString")
				saveGcmConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveGcmConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveGcmConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.SaveGCMConf(saveGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.SaveGCMConf(saveGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SaveGCMConf(saveGCMConfOptions *SaveGCMConfOptions)`, func() {
		saveGcmConfPath := "/apps/testString/settings/gcmConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveGcmConfPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "senderId": "SenderID"}`)
				}))
			})
			It(`Invoke SaveGCMConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the SaveGCMConfOptions model
				saveGcmConfOptionsModel := new(pushservicev1.SaveGCMConfOptions)
				saveGcmConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveGcmConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveGcmConfOptionsModel.SenderID = core.StringPtr("testString")
				saveGcmConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveGcmConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveGcmConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.SaveGCMConfWithContext(ctx, saveGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.SaveGCMConf(saveGcmConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.SaveGCMConfWithContext(ctx, saveGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveGcmConfPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "senderId": "SenderID"}`)
				}))
			})
			It(`Invoke SaveGCMConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.SaveGCMConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SaveGCMConfOptions model
				saveGcmConfOptionsModel := new(pushservicev1.SaveGCMConfOptions)
				saveGcmConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveGcmConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveGcmConfOptionsModel.SenderID = core.StringPtr("testString")
				saveGcmConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveGcmConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveGcmConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.SaveGCMConf(saveGcmConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SaveGCMConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveGCMConfOptions model
				saveGcmConfOptionsModel := new(pushservicev1.SaveGCMConfOptions)
				saveGcmConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveGcmConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveGcmConfOptionsModel.SenderID = core.StringPtr("testString")
				saveGcmConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveGcmConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveGcmConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.SaveGCMConf(saveGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SaveGCMConfOptions model with no property values
				saveGcmConfOptionsModelNew := new(pushservicev1.SaveGCMConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.SaveGCMConf(saveGcmConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteGCMConf(deleteGCMConfOptions *DeleteGCMConfOptions)`, func() {
		deleteGcmConfPath := "/apps/testString/settings/gcmConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteGcmConfPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteGCMConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := pushServiceService.DeleteGCMConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteGCMConfOptions model
				deleteGcmConfOptionsModel := new(pushservicev1.DeleteGCMConfOptions)
				deleteGcmConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteGcmConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteGcmConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteGcmConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = pushServiceService.DeleteGCMConf(deleteGcmConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteGCMConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the DeleteGCMConfOptions model
				deleteGcmConfOptionsModel := new(pushservicev1.DeleteGCMConfOptions)
				deleteGcmConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteGcmConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteGcmConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteGcmConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := pushServiceService.DeleteGCMConf(deleteGcmConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteGCMConfOptions model with no property values
				deleteGcmConfOptionsModelNew := new(pushservicev1.DeleteGCMConfOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = pushServiceService.DeleteGCMConf(deleteGcmConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWebpushServerKey(getWebpushServerKeyOptions *GetWebpushServerKeyOptions) - Operation response error`, func() {
		getWebpushServerKeyPath := "/apps/testString/settings/webpushServerKey"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWebpushServerKeyPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Clientsecret"]).ToNot(BeNil())
					Expect(req.Header["Clientsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWebpushServerKey with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetWebpushServerKeyOptions model
				getWebpushServerKeyOptionsModel := new(pushservicev1.GetWebpushServerKeyOptions)
				getWebpushServerKeyOptionsModel.ApplicationID = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.ClientSecret = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.GetWebpushServerKey(getWebpushServerKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.GetWebpushServerKey(getWebpushServerKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetWebpushServerKey(getWebpushServerKeyOptions *GetWebpushServerKeyOptions)`, func() {
		getWebpushServerKeyPath := "/apps/testString/settings/webpushServerKey"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWebpushServerKeyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Clientsecret"]).ToNot(BeNil())
					Expect(req.Header["Clientsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"webpushServerKey": "WebpushServerKey"}`)
				}))
			})
			It(`Invoke GetWebpushServerKey successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetWebpushServerKeyOptions model
				getWebpushServerKeyOptionsModel := new(pushservicev1.GetWebpushServerKeyOptions)
				getWebpushServerKeyOptionsModel.ApplicationID = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.ClientSecret = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.GetWebpushServerKeyWithContext(ctx, getWebpushServerKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.GetWebpushServerKey(getWebpushServerKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.GetWebpushServerKeyWithContext(ctx, getWebpushServerKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWebpushServerKeyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Clientsecret"]).ToNot(BeNil())
					Expect(req.Header["Clientsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"webpushServerKey": "WebpushServerKey"}`)
				}))
			})
			It(`Invoke GetWebpushServerKey successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.GetWebpushServerKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWebpushServerKeyOptions model
				getWebpushServerKeyOptionsModel := new(pushservicev1.GetWebpushServerKeyOptions)
				getWebpushServerKeyOptionsModel.ApplicationID = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.ClientSecret = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.GetWebpushServerKey(getWebpushServerKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetWebpushServerKey with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetWebpushServerKeyOptions model
				getWebpushServerKeyOptionsModel := new(pushservicev1.GetWebpushServerKeyOptions)
				getWebpushServerKeyOptionsModel.ApplicationID = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.ClientSecret = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getWebpushServerKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.GetWebpushServerKey(getWebpushServerKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWebpushServerKeyOptions model with no property values
				getWebpushServerKeyOptionsModelNew := new(pushservicev1.GetWebpushServerKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.GetWebpushServerKey(getWebpushServerKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSafariWebConf(getSafariWebConfOptions *GetSafariWebConfOptions) - Operation response error`, func() {
		getSafariWebConfPath := "/apps/testString/settings/safariWebConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSafariWebConfPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSafariWebConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetSafariWebConfOptions model
				getSafariWebConfOptionsModel := new(pushservicev1.GetSafariWebConfOptions)
				getSafariWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getSafariWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSafariWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getSafariWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.GetSafariWebConf(getSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.GetSafariWebConf(getSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSafariWebConf(getSafariWebConfOptions *GetSafariWebConfOptions)`, func() {
		getSafariWebConfPath := "/apps/testString/settings/safariWebConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSafariWebConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificate": "Certificate", "websiteName": "WebsiteName", "urlFormatString": "UrlFormatString", "websitePushID": {"anyKey": "anyValue"}, "webSiteUrl": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetSafariWebConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetSafariWebConfOptions model
				getSafariWebConfOptionsModel := new(pushservicev1.GetSafariWebConfOptions)
				getSafariWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getSafariWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSafariWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getSafariWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.GetSafariWebConfWithContext(ctx, getSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.GetSafariWebConf(getSafariWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.GetSafariWebConfWithContext(ctx, getSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSafariWebConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificate": "Certificate", "websiteName": "WebsiteName", "urlFormatString": "UrlFormatString", "websitePushID": {"anyKey": "anyValue"}, "webSiteUrl": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetSafariWebConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.GetSafariWebConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSafariWebConfOptions model
				getSafariWebConfOptionsModel := new(pushservicev1.GetSafariWebConfOptions)
				getSafariWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getSafariWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSafariWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getSafariWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.GetSafariWebConf(getSafariWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSafariWebConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetSafariWebConfOptions model
				getSafariWebConfOptionsModel := new(pushservicev1.GetSafariWebConfOptions)
				getSafariWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getSafariWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getSafariWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getSafariWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.GetSafariWebConf(getSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSafariWebConfOptions model with no property values
				getSafariWebConfOptionsModelNew := new(pushservicev1.GetSafariWebConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.GetSafariWebConf(getSafariWebConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SaveSafariWebConf(saveSafariWebConfOptions *SaveSafariWebConfOptions) - Operation response error`, func() {
		saveSafariWebConfPath := "/apps/testString/settings/safariWebConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveSafariWebConfPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SaveSafariWebConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveSafariWebConfOptions model
				saveSafariWebConfOptionsModel := new(pushservicev1.SaveSafariWebConfOptions)
				saveSafariWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Password = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.WebsiteName = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.UrlFormatString = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.WebsitePushID = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.CertificateContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.SaveSafariWebConf(saveSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.SaveSafariWebConf(saveSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SaveSafariWebConf(saveSafariWebConfOptions *SaveSafariWebConfOptions)`, func() {
		saveSafariWebConfPath := "/apps/testString/settings/safariWebConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveSafariWebConfPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificate": "Certificate", "websiteName": "WebsiteName", "urlFormatString": "UrlFormatString", "websitePushID": {"anyKey": "anyValue"}, "webSiteUrl": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke SaveSafariWebConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the SaveSafariWebConfOptions model
				saveSafariWebConfOptionsModel := new(pushservicev1.SaveSafariWebConfOptions)
				saveSafariWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Password = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.WebsiteName = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.UrlFormatString = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.WebsitePushID = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.CertificateContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.SaveSafariWebConfWithContext(ctx, saveSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.SaveSafariWebConf(saveSafariWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.SaveSafariWebConfWithContext(ctx, saveSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveSafariWebConfPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"certificate": "Certificate", "websiteName": "WebsiteName", "urlFormatString": "UrlFormatString", "websitePushID": {"anyKey": "anyValue"}, "webSiteUrl": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke SaveSafariWebConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.SaveSafariWebConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SaveSafariWebConfOptions model
				saveSafariWebConfOptionsModel := new(pushservicev1.SaveSafariWebConfOptions)
				saveSafariWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Password = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.WebsiteName = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.UrlFormatString = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.WebsitePushID = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.CertificateContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.SaveSafariWebConf(saveSafariWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SaveSafariWebConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveSafariWebConfOptions model
				saveSafariWebConfOptionsModel := new(pushservicev1.SaveSafariWebConfOptions)
				saveSafariWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Password = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Certificate = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.WebsiteName = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.UrlFormatString = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.WebsitePushID = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.CertificateContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon16x16 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon16x16ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon16x162x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon16x162xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon32x32 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon32x32ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon32x322x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon32x322xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon128x128 = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon128x128ContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Icon128x1282x = CreateMockReader("This is a mock file.")
				saveSafariWebConfOptionsModel.Icon128x1282xContentType = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveSafariWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.SaveSafariWebConf(saveSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SaveSafariWebConfOptions model with no property values
				saveSafariWebConfOptionsModelNew := new(pushservicev1.SaveSafariWebConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.SaveSafariWebConf(saveSafariWebConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteSafariWebConf(deleteSafariWebConfOptions *DeleteSafariWebConfOptions)`, func() {
		deleteSafariWebConfPath := "/apps/testString/settings/safariWebConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSafariWebConfPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSafariWebConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := pushServiceService.DeleteSafariWebConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSafariWebConfOptions model
				deleteSafariWebConfOptionsModel := new(pushservicev1.DeleteSafariWebConfOptions)
				deleteSafariWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteSafariWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteSafariWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteSafariWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = pushServiceService.DeleteSafariWebConf(deleteSafariWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSafariWebConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the DeleteSafariWebConfOptions model
				deleteSafariWebConfOptionsModel := new(pushservicev1.DeleteSafariWebConfOptions)
				deleteSafariWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteSafariWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteSafariWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteSafariWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := pushServiceService.DeleteSafariWebConf(deleteSafariWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSafariWebConfOptions model with no property values
				deleteSafariWebConfOptionsModelNew := new(pushservicev1.DeleteSafariWebConfOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = pushServiceService.DeleteSafariWebConf(deleteSafariWebConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetGcmConfPublic(getGcmConfPublicOptions *GetGcmConfPublicOptions) - Operation response error`, func() {
		getGcmConfPublicPath := "/apps/testString/settings/gcmConfPublic"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGcmConfPublicPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Clientsecret"]).ToNot(BeNil())
					Expect(req.Header["Clientsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetGcmConfPublic with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetGcmConfPublicOptions model
				getGcmConfPublicOptionsModel := new(pushservicev1.GetGcmConfPublicOptions)
				getGcmConfPublicOptionsModel.ApplicationID = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.ClientSecret = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.GetGcmConfPublic(getGcmConfPublicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.GetGcmConfPublic(getGcmConfPublicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetGcmConfPublic(getGcmConfPublicOptions *GetGcmConfPublicOptions)`, func() {
		getGcmConfPublicPath := "/apps/testString/settings/gcmConfPublic"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGcmConfPublicPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Clientsecret"]).ToNot(BeNil())
					Expect(req.Header["Clientsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"senderId": "SenderID"}`)
				}))
			})
			It(`Invoke GetGcmConfPublic successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetGcmConfPublicOptions model
				getGcmConfPublicOptionsModel := new(pushservicev1.GetGcmConfPublicOptions)
				getGcmConfPublicOptionsModel.ApplicationID = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.ClientSecret = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.GetGcmConfPublicWithContext(ctx, getGcmConfPublicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.GetGcmConfPublic(getGcmConfPublicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.GetGcmConfPublicWithContext(ctx, getGcmConfPublicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getGcmConfPublicPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Clientsecret"]).ToNot(BeNil())
					Expect(req.Header["Clientsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"senderId": "SenderID"}`)
				}))
			})
			It(`Invoke GetGcmConfPublic successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.GetGcmConfPublic(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetGcmConfPublicOptions model
				getGcmConfPublicOptionsModel := new(pushservicev1.GetGcmConfPublicOptions)
				getGcmConfPublicOptionsModel.ApplicationID = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.ClientSecret = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.GetGcmConfPublic(getGcmConfPublicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetGcmConfPublic with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetGcmConfPublicOptions model
				getGcmConfPublicOptionsModel := new(pushservicev1.GetGcmConfPublicOptions)
				getGcmConfPublicOptionsModel.ApplicationID = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.ClientSecret = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getGcmConfPublicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.GetGcmConfPublic(getGcmConfPublicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetGcmConfPublicOptions model with no property values
				getGcmConfPublicOptionsModelNew := new(pushservicev1.GetGcmConfPublicOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.GetGcmConfPublic(getGcmConfPublicOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetChromeWebConf(getChromeWebConfOptions *GetChromeWebConfOptions) - Operation response error`, func() {
		getChromeWebConfPath := "/apps/testString/settings/chromeWebConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChromeWebConfPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetChromeWebConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetChromeWebConfOptions model
				getChromeWebConfOptionsModel := new(pushservicev1.GetChromeWebConfOptions)
				getChromeWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getChromeWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.GetChromeWebConf(getChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.GetChromeWebConf(getChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetChromeWebConf(getChromeWebConfOptions *GetChromeWebConfOptions)`, func() {
		getChromeWebConfPath := "/apps/testString/settings/chromeWebConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChromeWebConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "webSiteUrl": "WebSiteURL"}`)
				}))
			})
			It(`Invoke GetChromeWebConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetChromeWebConfOptions model
				getChromeWebConfOptionsModel := new(pushservicev1.GetChromeWebConfOptions)
				getChromeWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getChromeWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.GetChromeWebConfWithContext(ctx, getChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.GetChromeWebConf(getChromeWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.GetChromeWebConfWithContext(ctx, getChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChromeWebConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "webSiteUrl": "WebSiteURL"}`)
				}))
			})
			It(`Invoke GetChromeWebConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.GetChromeWebConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetChromeWebConfOptions model
				getChromeWebConfOptionsModel := new(pushservicev1.GetChromeWebConfOptions)
				getChromeWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getChromeWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.GetChromeWebConf(getChromeWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetChromeWebConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetChromeWebConfOptions model
				getChromeWebConfOptionsModel := new(pushservicev1.GetChromeWebConfOptions)
				getChromeWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getChromeWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.GetChromeWebConf(getChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetChromeWebConfOptions model with no property values
				getChromeWebConfOptionsModelNew := new(pushservicev1.GetChromeWebConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.GetChromeWebConf(getChromeWebConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SaveChromeWebConf(saveChromeWebConfOptions *SaveChromeWebConfOptions) - Operation response error`, func() {
		saveChromeWebConfPath := "/apps/testString/settings/chromeWebConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveChromeWebConfPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SaveChromeWebConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveChromeWebConfOptions model
				saveChromeWebConfOptionsModel := new(pushservicev1.SaveChromeWebConfOptions)
				saveChromeWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.SaveChromeWebConf(saveChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.SaveChromeWebConf(saveChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SaveChromeWebConf(saveChromeWebConfOptions *SaveChromeWebConfOptions)`, func() {
		saveChromeWebConfPath := "/apps/testString/settings/chromeWebConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveChromeWebConfPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "webSiteUrl": "WebSiteURL"}`)
				}))
			})
			It(`Invoke SaveChromeWebConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the SaveChromeWebConfOptions model
				saveChromeWebConfOptionsModel := new(pushservicev1.SaveChromeWebConfOptions)
				saveChromeWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.SaveChromeWebConfWithContext(ctx, saveChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.SaveChromeWebConf(saveChromeWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.SaveChromeWebConfWithContext(ctx, saveChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveChromeWebConfPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "webSiteUrl": "WebSiteURL"}`)
				}))
			})
			It(`Invoke SaveChromeWebConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.SaveChromeWebConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SaveChromeWebConfOptions model
				saveChromeWebConfOptionsModel := new(pushservicev1.SaveChromeWebConfOptions)
				saveChromeWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.SaveChromeWebConf(saveChromeWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SaveChromeWebConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveChromeWebConfOptions model
				saveChromeWebConfOptionsModel := new(pushservicev1.SaveChromeWebConfOptions)
				saveChromeWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveChromeWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.SaveChromeWebConf(saveChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SaveChromeWebConfOptions model with no property values
				saveChromeWebConfOptionsModelNew := new(pushservicev1.SaveChromeWebConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.SaveChromeWebConf(saveChromeWebConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteChromeWebConf(deleteChromeWebConfOptions *DeleteChromeWebConfOptions)`, func() {
		deleteChromeWebConfPath := "/apps/testString/settings/chromeWebConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteChromeWebConfPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteChromeWebConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := pushServiceService.DeleteChromeWebConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteChromeWebConfOptions model
				deleteChromeWebConfOptionsModel := new(pushservicev1.DeleteChromeWebConfOptions)
				deleteChromeWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteChromeWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteChromeWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteChromeWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = pushServiceService.DeleteChromeWebConf(deleteChromeWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteChromeWebConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the DeleteChromeWebConfOptions model
				deleteChromeWebConfOptionsModel := new(pushservicev1.DeleteChromeWebConfOptions)
				deleteChromeWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteChromeWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteChromeWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteChromeWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := pushServiceService.DeleteChromeWebConf(deleteChromeWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteChromeWebConfOptions model with no property values
				deleteChromeWebConfOptionsModelNew := new(pushservicev1.DeleteChromeWebConfOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = pushServiceService.DeleteChromeWebConf(deleteChromeWebConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetFirefoxWebConf(getFirefoxWebConfOptions *GetFirefoxWebConfOptions) - Operation response error`, func() {
		getFirefoxWebConfPath := "/apps/testString/settings/firefoxWebConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFirefoxWebConfPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFirefoxWebConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetFirefoxWebConfOptions model
				getFirefoxWebConfOptionsModel := new(pushservicev1.GetFirefoxWebConfOptions)
				getFirefoxWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.GetFirefoxWebConf(getFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.GetFirefoxWebConf(getFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetFirefoxWebConf(getFirefoxWebConfOptions *GetFirefoxWebConfOptions)`, func() {
		getFirefoxWebConfPath := "/apps/testString/settings/firefoxWebConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFirefoxWebConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"webSiteUrl": "WebSiteURL"}`)
				}))
			})
			It(`Invoke GetFirefoxWebConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetFirefoxWebConfOptions model
				getFirefoxWebConfOptionsModel := new(pushservicev1.GetFirefoxWebConfOptions)
				getFirefoxWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.GetFirefoxWebConfWithContext(ctx, getFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.GetFirefoxWebConf(getFirefoxWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.GetFirefoxWebConfWithContext(ctx, getFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getFirefoxWebConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"webSiteUrl": "WebSiteURL"}`)
				}))
			})
			It(`Invoke GetFirefoxWebConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.GetFirefoxWebConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFirefoxWebConfOptions model
				getFirefoxWebConfOptionsModel := new(pushservicev1.GetFirefoxWebConfOptions)
				getFirefoxWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.GetFirefoxWebConf(getFirefoxWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetFirefoxWebConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetFirefoxWebConfOptions model
				getFirefoxWebConfOptionsModel := new(pushservicev1.GetFirefoxWebConfOptions)
				getFirefoxWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				getFirefoxWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.GetFirefoxWebConf(getFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetFirefoxWebConfOptions model with no property values
				getFirefoxWebConfOptionsModelNew := new(pushservicev1.GetFirefoxWebConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.GetFirefoxWebConf(getFirefoxWebConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SaveFirefoxWebConf(saveFirefoxWebConfOptions *SaveFirefoxWebConfOptions) - Operation response error`, func() {
		saveFirefoxWebConfPath := "/apps/testString/settings/firefoxWebConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveFirefoxWebConfPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SaveFirefoxWebConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveFirefoxWebConfOptions model
				saveFirefoxWebConfOptionsModel := new(pushservicev1.SaveFirefoxWebConfOptions)
				saveFirefoxWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.SaveFirefoxWebConf(saveFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.SaveFirefoxWebConf(saveFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SaveFirefoxWebConf(saveFirefoxWebConfOptions *SaveFirefoxWebConfOptions)`, func() {
		saveFirefoxWebConfPath := "/apps/testString/settings/firefoxWebConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveFirefoxWebConfPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"webSiteUrl": "WebSiteURL"}`)
				}))
			})
			It(`Invoke SaveFirefoxWebConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the SaveFirefoxWebConfOptions model
				saveFirefoxWebConfOptionsModel := new(pushservicev1.SaveFirefoxWebConfOptions)
				saveFirefoxWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.SaveFirefoxWebConfWithContext(ctx, saveFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.SaveFirefoxWebConf(saveFirefoxWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.SaveFirefoxWebConfWithContext(ctx, saveFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveFirefoxWebConfPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"webSiteUrl": "WebSiteURL"}`)
				}))
			})
			It(`Invoke SaveFirefoxWebConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.SaveFirefoxWebConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SaveFirefoxWebConfOptions model
				saveFirefoxWebConfOptionsModel := new(pushservicev1.SaveFirefoxWebConfOptions)
				saveFirefoxWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.SaveFirefoxWebConf(saveFirefoxWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SaveFirefoxWebConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveFirefoxWebConfOptions model
				saveFirefoxWebConfOptionsModel := new(pushservicev1.SaveFirefoxWebConfOptions)
				saveFirefoxWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.WebSiteURL = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveFirefoxWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.SaveFirefoxWebConf(saveFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SaveFirefoxWebConfOptions model with no property values
				saveFirefoxWebConfOptionsModelNew := new(pushservicev1.SaveFirefoxWebConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.SaveFirefoxWebConf(saveFirefoxWebConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteFirefoxWebConf(deleteFirefoxWebConfOptions *DeleteFirefoxWebConfOptions)`, func() {
		deleteFirefoxWebConfPath := "/apps/testString/settings/firefoxWebConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteFirefoxWebConfPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteFirefoxWebConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := pushServiceService.DeleteFirefoxWebConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteFirefoxWebConfOptions model
				deleteFirefoxWebConfOptionsModel := new(pushservicev1.DeleteFirefoxWebConfOptions)
				deleteFirefoxWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteFirefoxWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteFirefoxWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteFirefoxWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = pushServiceService.DeleteFirefoxWebConf(deleteFirefoxWebConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteFirefoxWebConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the DeleteFirefoxWebConfOptions model
				deleteFirefoxWebConfOptionsModel := new(pushservicev1.DeleteFirefoxWebConfOptions)
				deleteFirefoxWebConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteFirefoxWebConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteFirefoxWebConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteFirefoxWebConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := pushServiceService.DeleteFirefoxWebConf(deleteFirefoxWebConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteFirefoxWebConfOptions model with no property values
				deleteFirefoxWebConfOptionsModelNew := new(pushservicev1.DeleteFirefoxWebConfOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = pushServiceService.DeleteFirefoxWebConf(deleteFirefoxWebConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetChromeAppExtConf(getChromeAppExtConfOptions *GetChromeAppExtConfOptions) - Operation response error`, func() {
		getChromeAppExtConfPath := "/apps/testString/settings/chromeAppExtConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChromeAppExtConfPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetChromeAppExtConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetChromeAppExtConfOptions model
				getChromeAppExtConfOptionsModel := new(pushservicev1.GetChromeAppExtConfOptions)
				getChromeAppExtConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.AppSecret = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.GetChromeAppExtConf(getChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.GetChromeAppExtConf(getChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetChromeAppExtConf(getChromeAppExtConfOptions *GetChromeAppExtConfOptions)`, func() {
		getChromeAppExtConfPath := "/apps/testString/settings/chromeAppExtConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChromeAppExtConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "senderId": "SenderID"}`)
				}))
			})
			It(`Invoke GetChromeAppExtConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetChromeAppExtConfOptions model
				getChromeAppExtConfOptionsModel := new(pushservicev1.GetChromeAppExtConfOptions)
				getChromeAppExtConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.AppSecret = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.GetChromeAppExtConfWithContext(ctx, getChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.GetChromeAppExtConf(getChromeAppExtConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.GetChromeAppExtConfWithContext(ctx, getChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChromeAppExtConfPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "senderId": "SenderID"}`)
				}))
			})
			It(`Invoke GetChromeAppExtConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.GetChromeAppExtConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetChromeAppExtConfOptions model
				getChromeAppExtConfOptionsModel := new(pushservicev1.GetChromeAppExtConfOptions)
				getChromeAppExtConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.AppSecret = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.GetChromeAppExtConf(getChromeAppExtConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetChromeAppExtConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetChromeAppExtConfOptions model
				getChromeAppExtConfOptionsModel := new(pushservicev1.GetChromeAppExtConfOptions)
				getChromeAppExtConfOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.AppSecret = core.StringPtr("testString")
				getChromeAppExtConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.GetChromeAppExtConf(getChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetChromeAppExtConfOptions model with no property values
				getChromeAppExtConfOptionsModelNew := new(pushservicev1.GetChromeAppExtConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.GetChromeAppExtConf(getChromeAppExtConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SaveChromeAppExtConf(saveChromeAppExtConfOptions *SaveChromeAppExtConfOptions) - Operation response error`, func() {
		saveChromeAppExtConfPath := "/apps/testString/settings/chromeAppExtConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveChromeAppExtConfPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SaveChromeAppExtConf with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveChromeAppExtConfOptions model
				saveChromeAppExtConfOptionsModel := new(pushservicev1.SaveChromeAppExtConfOptions)
				saveChromeAppExtConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.SenderID = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.SaveChromeAppExtConf(saveChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.SaveChromeAppExtConf(saveChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SaveChromeAppExtConf(saveChromeAppExtConfOptions *SaveChromeAppExtConfOptions)`, func() {
		saveChromeAppExtConfPath := "/apps/testString/settings/chromeAppExtConf"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveChromeAppExtConfPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "senderId": "SenderID"}`)
				}))
			})
			It(`Invoke SaveChromeAppExtConf successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the SaveChromeAppExtConfOptions model
				saveChromeAppExtConfOptionsModel := new(pushservicev1.SaveChromeAppExtConfOptions)
				saveChromeAppExtConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.SenderID = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.SaveChromeAppExtConfWithContext(ctx, saveChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.SaveChromeAppExtConf(saveChromeAppExtConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.SaveChromeAppExtConfWithContext(ctx, saveChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(saveChromeAppExtConfPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"apiKey": "ApiKey", "senderId": "SenderID"}`)
				}))
			})
			It(`Invoke SaveChromeAppExtConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.SaveChromeAppExtConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SaveChromeAppExtConfOptions model
				saveChromeAppExtConfOptionsModel := new(pushservicev1.SaveChromeAppExtConfOptions)
				saveChromeAppExtConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.SenderID = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.SaveChromeAppExtConf(saveChromeAppExtConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SaveChromeAppExtConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the SaveChromeAppExtConfOptions model
				saveChromeAppExtConfOptionsModel := new(pushservicev1.SaveChromeAppExtConfOptions)
				saveChromeAppExtConfOptionsModel.ApplicationID = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.ApiKey = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.SenderID = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.AppSecret = core.StringPtr("testString")
				saveChromeAppExtConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.SaveChromeAppExtConf(saveChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SaveChromeAppExtConfOptions model with no property values
				saveChromeAppExtConfOptionsModelNew := new(pushservicev1.SaveChromeAppExtConfOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.SaveChromeAppExtConf(saveChromeAppExtConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteChromeAppExtConf(deleteChromeAppExtConfOptions *DeleteChromeAppExtConfOptions)`, func() {
		deleteChromeAppExtConfPath := "/apps/testString/settings/chromeAppExtConf"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteChromeAppExtConfPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteChromeAppExtConf successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := pushServiceService.DeleteChromeAppExtConf(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteChromeAppExtConfOptions model
				deleteChromeAppExtConfOptionsModel := new(pushservicev1.DeleteChromeAppExtConfOptions)
				deleteChromeAppExtConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteChromeAppExtConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteChromeAppExtConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteChromeAppExtConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = pushServiceService.DeleteChromeAppExtConf(deleteChromeAppExtConfOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteChromeAppExtConf with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the DeleteChromeAppExtConfOptions model
				deleteChromeAppExtConfOptionsModel := new(pushservicev1.DeleteChromeAppExtConfOptions)
				deleteChromeAppExtConfOptionsModel.ApplicationID = core.StringPtr("testString")
				deleteChromeAppExtConfOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteChromeAppExtConfOptionsModel.AppSecret = core.StringPtr("testString")
				deleteChromeAppExtConfOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := pushServiceService.DeleteChromeAppExtConf(deleteChromeAppExtConfOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteChromeAppExtConfOptions model with no property values
				deleteChromeAppExtConfOptionsModelNew := new(pushservicev1.DeleteChromeAppExtConfOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = pushServiceService.DeleteChromeAppExtConf(deleteChromeAppExtConfOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetChromeAppExtConfPublic(getChromeAppExtConfPublicOptions *GetChromeAppExtConfPublicOptions) - Operation response error`, func() {
		getChromeAppExtConfPublicPath := "/apps/testString/settings/chromeAppExtConfPublic"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChromeAppExtConfPublicPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Clientsecret"]).ToNot(BeNil())
					Expect(req.Header["Clientsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetChromeAppExtConfPublic with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetChromeAppExtConfPublicOptions model
				getChromeAppExtConfPublicOptionsModel := new(pushservicev1.GetChromeAppExtConfPublicOptions)
				getChromeAppExtConfPublicOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.ClientSecret = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.GetChromeAppExtConfPublic(getChromeAppExtConfPublicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.GetChromeAppExtConfPublic(getChromeAppExtConfPublicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetChromeAppExtConfPublic(getChromeAppExtConfPublicOptions *GetChromeAppExtConfPublicOptions)`, func() {
		getChromeAppExtConfPublicPath := "/apps/testString/settings/chromeAppExtConfPublic"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChromeAppExtConfPublicPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Clientsecret"]).ToNot(BeNil())
					Expect(req.Header["Clientsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"senderId": "SenderID"}`)
				}))
			})
			It(`Invoke GetChromeAppExtConfPublic successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the GetChromeAppExtConfPublicOptions model
				getChromeAppExtConfPublicOptionsModel := new(pushservicev1.GetChromeAppExtConfPublicOptions)
				getChromeAppExtConfPublicOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.ClientSecret = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.GetChromeAppExtConfPublicWithContext(ctx, getChromeAppExtConfPublicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.GetChromeAppExtConfPublic(getChromeAppExtConfPublicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.GetChromeAppExtConfPublicWithContext(ctx, getChromeAppExtConfPublicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getChromeAppExtConfPublicPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Clientsecret"]).ToNot(BeNil())
					Expect(req.Header["Clientsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"senderId": "SenderID"}`)
				}))
			})
			It(`Invoke GetChromeAppExtConfPublic successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.GetChromeAppExtConfPublic(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetChromeAppExtConfPublicOptions model
				getChromeAppExtConfPublicOptionsModel := new(pushservicev1.GetChromeAppExtConfPublicOptions)
				getChromeAppExtConfPublicOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.ClientSecret = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.GetChromeAppExtConfPublic(getChromeAppExtConfPublicOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetChromeAppExtConfPublic with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the GetChromeAppExtConfPublicOptions model
				getChromeAppExtConfPublicOptionsModel := new(pushservicev1.GetChromeAppExtConfPublicOptions)
				getChromeAppExtConfPublicOptionsModel.ApplicationID = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.ClientSecret = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getChromeAppExtConfPublicOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.GetChromeAppExtConfPublic(getChromeAppExtConfPublicOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetChromeAppExtConfPublicOptions model with no property values
				getChromeAppExtConfPublicOptionsModelNew := new(pushservicev1.GetChromeAppExtConfPublicOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.GetChromeAppExtConfPublic(getChromeAppExtConfPublicOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(pushServiceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(pushServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
				URL: "https://pushservicev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(pushServiceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PUSH_SERVICE_URL":       "https://pushservicev1/api",
				"PUSH_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1UsingExternalConfig(&pushservicev1.PushServiceV1Options{})
				Expect(pushServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := pushServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != pushServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(pushServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(pushServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1UsingExternalConfig(&pushservicev1.PushServiceV1Options{
					URL: "https://testService/api",
				})
				Expect(pushServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := pushServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != pushServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(pushServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(pushServiceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1UsingExternalConfig(&pushservicev1.PushServiceV1Options{})
				err := pushServiceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := pushServiceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != pushServiceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(pushServiceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(pushServiceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PUSH_SERVICE_URL":       "https://pushservicev1/api",
				"PUSH_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			pushServiceService, serviceErr := pushservicev1.NewPushServiceV1UsingExternalConfig(&pushservicev1.PushServiceV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(pushServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PUSH_SERVICE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			pushServiceService, serviceErr := pushservicev1.NewPushServiceV1UsingExternalConfig(&pushservicev1.PushServiceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(pushServiceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = pushservicev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`SendMessage(sendMessageOptions *SendMessageOptions) - Operation response error`, func() {
		sendMessagePath := "/apps/testString/messages"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendMessagePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SendMessage with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the Message model
				messageModel := new(pushservicev1.Message)
				messageModel.Alert = core.StringPtr("testString")
				messageModel.URL = core.StringPtr("testString")

				// Construct an instance of the Apns model
				apnsModel := new(pushservicev1.Apns)
				apnsModel.Badge = core.Int64Ptr(int64(38))
				apnsModel.InteractiveCategory = core.StringPtr("testString")
				apnsModel.Category = core.StringPtr("testString")
				apnsModel.IosActionKey = core.StringPtr("testString")
				apnsModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				apnsModel.Sound = core.StringPtr("testString")
				apnsModel.TitleLocKey = core.StringPtr("testString")
				apnsModel.LocKey = core.StringPtr("testString")
				apnsModel.LaunchImage = core.StringPtr("testString")
				apnsModel.TitleLocArgs = []string{"testString"}
				apnsModel.LocArgs = []string{"testString"}
				apnsModel.Title = core.StringPtr("testString")
				apnsModel.Subtitle = core.StringPtr("testString")
				apnsModel.AttachmentURL = core.StringPtr("testString")
				apnsModel.Type = core.StringPtr("DEFAULT")
				apnsModel.ApnsCollapseID = core.StringPtr("testString")
				apnsModel.ApnsThreadID = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Lights model
				lightsModel := new(pushservicev1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(38))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(pushservicev1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}

				// Construct an instance of the Gcm model
				gcmModel := new(pushservicev1.Gcm)
				gcmModel.CollapseKey = core.StringPtr("testString")
				gcmModel.InteractiveCategory = core.StringPtr("testString")
				gcmModel.Icon = core.StringPtr("testString")
				gcmModel.DelayWhileIdle = core.BoolPtr(true)
				gcmModel.Sync = core.BoolPtr(true)
				gcmModel.Visibility = core.StringPtr("testString")
				gcmModel.Redact = core.StringPtr("testString")
				gcmModel.ChannelID = core.StringPtr("testString")
				gcmModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				gcmModel.Priority = core.StringPtr("testString")
				gcmModel.Sound = core.StringPtr("testString")
				gcmModel.TimeToLive = core.Int64Ptr(int64(38))
				gcmModel.Lights = lightsModel
				gcmModel.AndroidTitle = core.StringPtr("testString")
				gcmModel.GroupID = core.StringPtr("testString")
				gcmModel.Style = styleModel
				gcmModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the FirefoxWeb model
				firefoxWebModel := new(pushservicev1.FirefoxWeb)
				firefoxWebModel.Title = core.StringPtr("testString")
				firefoxWebModel.IconURL = core.StringPtr("testString")
				firefoxWebModel.TimeToLive = core.Int64Ptr(int64(38))
				firefoxWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the ChromeWeb model
				chromeWebModel := new(pushservicev1.ChromeWeb)
				chromeWebModel.Title = core.StringPtr("testString")
				chromeWebModel.IconURL = core.StringPtr("testString")
				chromeWebModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the SafariWeb model
				safariWebModel := new(pushservicev1.SafariWeb)
				safariWebModel.Title = core.StringPtr("testString")
				safariWebModel.UrlArgs = []string{"testString"}
				safariWebModel.Action = core.StringPtr("testString")

				// Construct an instance of the ChromeAppExt model
				chromeAppExtModel := new(pushservicev1.ChromeAppExt)
				chromeAppExtModel.CollapseKey = core.StringPtr("testString")
				chromeAppExtModel.DelayWhileIdle = core.BoolPtr(true)
				chromeAppExtModel.Title = core.StringPtr("testString")
				chromeAppExtModel.IconURL = core.StringPtr("testString")
				chromeAppExtModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeAppExtModel.Payload = core.StringPtr("testString")

				// Construct an instance of the Settings model
				settingsModel := new(pushservicev1.Settings)
				settingsModel.Apns = apnsModel
				settingsModel.Gcm = gcmModel
				settingsModel.FirefoxWeb = firefoxWebModel
				settingsModel.ChromeWeb = chromeWebModel
				settingsModel.SafariWeb = safariWebModel
				settingsModel.ChromeAppExt = chromeAppExtModel

				// Construct an instance of the Target model
				targetModel := new(pushservicev1.Target)
				targetModel.DeviceIds = []string{"testString"}
				targetModel.UserIds = []string{"testString"}
				targetModel.Platforms = []string{"testString"}
				targetModel.TagNames = []string{"testString"}

				// Construct an instance of the SendMessageOptions model
				sendMessageOptionsModel := new(pushservicev1.SendMessageOptions)
				sendMessageOptionsModel.ApplicationID = core.StringPtr("testString")
				sendMessageOptionsModel.Message = messageModel
				sendMessageOptionsModel.Settings = settingsModel
				sendMessageOptionsModel.Validate = core.BoolPtr(true)
				sendMessageOptionsModel.Target = targetModel
				sendMessageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				sendMessageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.SendMessage(sendMessageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.SendMessage(sendMessageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SendMessage(sendMessageOptions *SendMessageOptions)`, func() {
		sendMessagePath := "/apps/testString/messages"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendMessagePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"message": {"message": {"alert": "Alert", "url": "URL"}, "settings": {"apns": {"badge": 5, "interactiveCategory": "InteractiveCategory", "category": "Category", "iosActionKey": "IosActionKey", "payload": {"anyKey": "anyValue"}, "sound": "Sound", "titleLocKey": "TitleLocKey", "locKey": "LocKey", "launchImage": "LaunchImage", "titleLocArgs": ["TitleLocArgs"], "locArgs": ["LocArgs"], "title": "Title", "subtitle": "Subtitle", "attachmentUrl": "AttachmentURL", "type": "DEFAULT", "apnsCollapseId": "ApnsCollapseID", "apnsThreadId": "ApnsThreadID", "apnsGroupSummaryArg": "ApnsGroupSummaryArg", "apnsGroupSummaryArgCount": 24}, "gcm": {"collapseKey": "CollapseKey", "interactiveCategory": "InteractiveCategory", "icon": "Icon", "delayWhileIdle": true, "sync": true, "visibility": "Visibility", "redact": "Redact", "channelId": "ChannelID", "payload": {"anyKey": "anyValue"}, "priority": "Priority", "sound": "Sound", "timeToLive": 10, "lights": {"ledArgb": "LedArgb", "ledOnMs": 7, "ledOffMs": "LedOffMs"}, "androidTitle": "AndroidTitle", "groupId": "GroupID", "style": {"type": "Type", "title": "Title", "url": "URL", "text": "Text", "lines": ["Lines"]}, "type": "DEFAULT"}, "firefoxWeb": {"title": "Title", "iconUrl": "IconURL", "timeToLive": 10, "payload": "Payload"}, "chromeWeb": {"title": "Title", "iconUrl": "IconURL", "timeToLive": 10, "payload": "Payload"}, "safariWeb": {"title": "Title", "urlArgs": ["UrlArgs"], "action": "Action"}, "chromeAppExt": {"collapseKey": "CollapseKey", "delayWhileIdle": true, "title": "Title", "iconUrl": "IconURL", "timeToLive": 10, "payload": "Payload"}}, "validate": true, "target": {"deviceIds": ["DeviceIds"], "userIds": ["UserIds"], "platforms": ["Platforms"], "tagNames": ["TagNames"]}}, "messageId": "MessageID"}`)
				}))
			})
			It(`Invoke SendMessage successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the Message model
				messageModel := new(pushservicev1.Message)
				messageModel.Alert = core.StringPtr("testString")
				messageModel.URL = core.StringPtr("testString")

				// Construct an instance of the Apns model
				apnsModel := new(pushservicev1.Apns)
				apnsModel.Badge = core.Int64Ptr(int64(38))
				apnsModel.InteractiveCategory = core.StringPtr("testString")
				apnsModel.Category = core.StringPtr("testString")
				apnsModel.IosActionKey = core.StringPtr("testString")
				apnsModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				apnsModel.Sound = core.StringPtr("testString")
				apnsModel.TitleLocKey = core.StringPtr("testString")
				apnsModel.LocKey = core.StringPtr("testString")
				apnsModel.LaunchImage = core.StringPtr("testString")
				apnsModel.TitleLocArgs = []string{"testString"}
				apnsModel.LocArgs = []string{"testString"}
				apnsModel.Title = core.StringPtr("testString")
				apnsModel.Subtitle = core.StringPtr("testString")
				apnsModel.AttachmentURL = core.StringPtr("testString")
				apnsModel.Type = core.StringPtr("DEFAULT")
				apnsModel.ApnsCollapseID = core.StringPtr("testString")
				apnsModel.ApnsThreadID = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Lights model
				lightsModel := new(pushservicev1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(38))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(pushservicev1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}

				// Construct an instance of the Gcm model
				gcmModel := new(pushservicev1.Gcm)
				gcmModel.CollapseKey = core.StringPtr("testString")
				gcmModel.InteractiveCategory = core.StringPtr("testString")
				gcmModel.Icon = core.StringPtr("testString")
				gcmModel.DelayWhileIdle = core.BoolPtr(true)
				gcmModel.Sync = core.BoolPtr(true)
				gcmModel.Visibility = core.StringPtr("testString")
				gcmModel.Redact = core.StringPtr("testString")
				gcmModel.ChannelID = core.StringPtr("testString")
				gcmModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				gcmModel.Priority = core.StringPtr("testString")
				gcmModel.Sound = core.StringPtr("testString")
				gcmModel.TimeToLive = core.Int64Ptr(int64(38))
				gcmModel.Lights = lightsModel
				gcmModel.AndroidTitle = core.StringPtr("testString")
				gcmModel.GroupID = core.StringPtr("testString")
				gcmModel.Style = styleModel
				gcmModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the FirefoxWeb model
				firefoxWebModel := new(pushservicev1.FirefoxWeb)
				firefoxWebModel.Title = core.StringPtr("testString")
				firefoxWebModel.IconURL = core.StringPtr("testString")
				firefoxWebModel.TimeToLive = core.Int64Ptr(int64(38))
				firefoxWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the ChromeWeb model
				chromeWebModel := new(pushservicev1.ChromeWeb)
				chromeWebModel.Title = core.StringPtr("testString")
				chromeWebModel.IconURL = core.StringPtr("testString")
				chromeWebModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the SafariWeb model
				safariWebModel := new(pushservicev1.SafariWeb)
				safariWebModel.Title = core.StringPtr("testString")
				safariWebModel.UrlArgs = []string{"testString"}
				safariWebModel.Action = core.StringPtr("testString")

				// Construct an instance of the ChromeAppExt model
				chromeAppExtModel := new(pushservicev1.ChromeAppExt)
				chromeAppExtModel.CollapseKey = core.StringPtr("testString")
				chromeAppExtModel.DelayWhileIdle = core.BoolPtr(true)
				chromeAppExtModel.Title = core.StringPtr("testString")
				chromeAppExtModel.IconURL = core.StringPtr("testString")
				chromeAppExtModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeAppExtModel.Payload = core.StringPtr("testString")

				// Construct an instance of the Settings model
				settingsModel := new(pushservicev1.Settings)
				settingsModel.Apns = apnsModel
				settingsModel.Gcm = gcmModel
				settingsModel.FirefoxWeb = firefoxWebModel
				settingsModel.ChromeWeb = chromeWebModel
				settingsModel.SafariWeb = safariWebModel
				settingsModel.ChromeAppExt = chromeAppExtModel

				// Construct an instance of the Target model
				targetModel := new(pushservicev1.Target)
				targetModel.DeviceIds = []string{"testString"}
				targetModel.UserIds = []string{"testString"}
				targetModel.Platforms = []string{"testString"}
				targetModel.TagNames = []string{"testString"}

				// Construct an instance of the SendMessageOptions model
				sendMessageOptionsModel := new(pushservicev1.SendMessageOptions)
				sendMessageOptionsModel.ApplicationID = core.StringPtr("testString")
				sendMessageOptionsModel.Message = messageModel
				sendMessageOptionsModel.Settings = settingsModel
				sendMessageOptionsModel.Validate = core.BoolPtr(true)
				sendMessageOptionsModel.Target = targetModel
				sendMessageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				sendMessageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.SendMessageWithContext(ctx, sendMessageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.SendMessage(sendMessageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.SendMessageWithContext(ctx, sendMessageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendMessagePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"message": {"message": {"alert": "Alert", "url": "URL"}, "settings": {"apns": {"badge": 5, "interactiveCategory": "InteractiveCategory", "category": "Category", "iosActionKey": "IosActionKey", "payload": {"anyKey": "anyValue"}, "sound": "Sound", "titleLocKey": "TitleLocKey", "locKey": "LocKey", "launchImage": "LaunchImage", "titleLocArgs": ["TitleLocArgs"], "locArgs": ["LocArgs"], "title": "Title", "subtitle": "Subtitle", "attachmentUrl": "AttachmentURL", "type": "DEFAULT", "apnsCollapseId": "ApnsCollapseID", "apnsThreadId": "ApnsThreadID", "apnsGroupSummaryArg": "ApnsGroupSummaryArg", "apnsGroupSummaryArgCount": 24}, "gcm": {"collapseKey": "CollapseKey", "interactiveCategory": "InteractiveCategory", "icon": "Icon", "delayWhileIdle": true, "sync": true, "visibility": "Visibility", "redact": "Redact", "channelId": "ChannelID", "payload": {"anyKey": "anyValue"}, "priority": "Priority", "sound": "Sound", "timeToLive": 10, "lights": {"ledArgb": "LedArgb", "ledOnMs": 7, "ledOffMs": "LedOffMs"}, "androidTitle": "AndroidTitle", "groupId": "GroupID", "style": {"type": "Type", "title": "Title", "url": "URL", "text": "Text", "lines": ["Lines"]}, "type": "DEFAULT"}, "firefoxWeb": {"title": "Title", "iconUrl": "IconURL", "timeToLive": 10, "payload": "Payload"}, "chromeWeb": {"title": "Title", "iconUrl": "IconURL", "timeToLive": 10, "payload": "Payload"}, "safariWeb": {"title": "Title", "urlArgs": ["UrlArgs"], "action": "Action"}, "chromeAppExt": {"collapseKey": "CollapseKey", "delayWhileIdle": true, "title": "Title", "iconUrl": "IconURL", "timeToLive": 10, "payload": "Payload"}}, "validate": true, "target": {"deviceIds": ["DeviceIds"], "userIds": ["UserIds"], "platforms": ["Platforms"], "tagNames": ["TagNames"]}}, "messageId": "MessageID"}`)
				}))
			})
			It(`Invoke SendMessage successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.SendMessage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Message model
				messageModel := new(pushservicev1.Message)
				messageModel.Alert = core.StringPtr("testString")
				messageModel.URL = core.StringPtr("testString")

				// Construct an instance of the Apns model
				apnsModel := new(pushservicev1.Apns)
				apnsModel.Badge = core.Int64Ptr(int64(38))
				apnsModel.InteractiveCategory = core.StringPtr("testString")
				apnsModel.Category = core.StringPtr("testString")
				apnsModel.IosActionKey = core.StringPtr("testString")
				apnsModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				apnsModel.Sound = core.StringPtr("testString")
				apnsModel.TitleLocKey = core.StringPtr("testString")
				apnsModel.LocKey = core.StringPtr("testString")
				apnsModel.LaunchImage = core.StringPtr("testString")
				apnsModel.TitleLocArgs = []string{"testString"}
				apnsModel.LocArgs = []string{"testString"}
				apnsModel.Title = core.StringPtr("testString")
				apnsModel.Subtitle = core.StringPtr("testString")
				apnsModel.AttachmentURL = core.StringPtr("testString")
				apnsModel.Type = core.StringPtr("DEFAULT")
				apnsModel.ApnsCollapseID = core.StringPtr("testString")
				apnsModel.ApnsThreadID = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Lights model
				lightsModel := new(pushservicev1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(38))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(pushservicev1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}

				// Construct an instance of the Gcm model
				gcmModel := new(pushservicev1.Gcm)
				gcmModel.CollapseKey = core.StringPtr("testString")
				gcmModel.InteractiveCategory = core.StringPtr("testString")
				gcmModel.Icon = core.StringPtr("testString")
				gcmModel.DelayWhileIdle = core.BoolPtr(true)
				gcmModel.Sync = core.BoolPtr(true)
				gcmModel.Visibility = core.StringPtr("testString")
				gcmModel.Redact = core.StringPtr("testString")
				gcmModel.ChannelID = core.StringPtr("testString")
				gcmModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				gcmModel.Priority = core.StringPtr("testString")
				gcmModel.Sound = core.StringPtr("testString")
				gcmModel.TimeToLive = core.Int64Ptr(int64(38))
				gcmModel.Lights = lightsModel
				gcmModel.AndroidTitle = core.StringPtr("testString")
				gcmModel.GroupID = core.StringPtr("testString")
				gcmModel.Style = styleModel
				gcmModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the FirefoxWeb model
				firefoxWebModel := new(pushservicev1.FirefoxWeb)
				firefoxWebModel.Title = core.StringPtr("testString")
				firefoxWebModel.IconURL = core.StringPtr("testString")
				firefoxWebModel.TimeToLive = core.Int64Ptr(int64(38))
				firefoxWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the ChromeWeb model
				chromeWebModel := new(pushservicev1.ChromeWeb)
				chromeWebModel.Title = core.StringPtr("testString")
				chromeWebModel.IconURL = core.StringPtr("testString")
				chromeWebModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the SafariWeb model
				safariWebModel := new(pushservicev1.SafariWeb)
				safariWebModel.Title = core.StringPtr("testString")
				safariWebModel.UrlArgs = []string{"testString"}
				safariWebModel.Action = core.StringPtr("testString")

				// Construct an instance of the ChromeAppExt model
				chromeAppExtModel := new(pushservicev1.ChromeAppExt)
				chromeAppExtModel.CollapseKey = core.StringPtr("testString")
				chromeAppExtModel.DelayWhileIdle = core.BoolPtr(true)
				chromeAppExtModel.Title = core.StringPtr("testString")
				chromeAppExtModel.IconURL = core.StringPtr("testString")
				chromeAppExtModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeAppExtModel.Payload = core.StringPtr("testString")

				// Construct an instance of the Settings model
				settingsModel := new(pushservicev1.Settings)
				settingsModel.Apns = apnsModel
				settingsModel.Gcm = gcmModel
				settingsModel.FirefoxWeb = firefoxWebModel
				settingsModel.ChromeWeb = chromeWebModel
				settingsModel.SafariWeb = safariWebModel
				settingsModel.ChromeAppExt = chromeAppExtModel

				// Construct an instance of the Target model
				targetModel := new(pushservicev1.Target)
				targetModel.DeviceIds = []string{"testString"}
				targetModel.UserIds = []string{"testString"}
				targetModel.Platforms = []string{"testString"}
				targetModel.TagNames = []string{"testString"}

				// Construct an instance of the SendMessageOptions model
				sendMessageOptionsModel := new(pushservicev1.SendMessageOptions)
				sendMessageOptionsModel.ApplicationID = core.StringPtr("testString")
				sendMessageOptionsModel.Message = messageModel
				sendMessageOptionsModel.Settings = settingsModel
				sendMessageOptionsModel.Validate = core.BoolPtr(true)
				sendMessageOptionsModel.Target = targetModel
				sendMessageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				sendMessageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.SendMessage(sendMessageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SendMessage with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the Message model
				messageModel := new(pushservicev1.Message)
				messageModel.Alert = core.StringPtr("testString")
				messageModel.URL = core.StringPtr("testString")

				// Construct an instance of the Apns model
				apnsModel := new(pushservicev1.Apns)
				apnsModel.Badge = core.Int64Ptr(int64(38))
				apnsModel.InteractiveCategory = core.StringPtr("testString")
				apnsModel.Category = core.StringPtr("testString")
				apnsModel.IosActionKey = core.StringPtr("testString")
				apnsModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				apnsModel.Sound = core.StringPtr("testString")
				apnsModel.TitleLocKey = core.StringPtr("testString")
				apnsModel.LocKey = core.StringPtr("testString")
				apnsModel.LaunchImage = core.StringPtr("testString")
				apnsModel.TitleLocArgs = []string{"testString"}
				apnsModel.LocArgs = []string{"testString"}
				apnsModel.Title = core.StringPtr("testString")
				apnsModel.Subtitle = core.StringPtr("testString")
				apnsModel.AttachmentURL = core.StringPtr("testString")
				apnsModel.Type = core.StringPtr("DEFAULT")
				apnsModel.ApnsCollapseID = core.StringPtr("testString")
				apnsModel.ApnsThreadID = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Lights model
				lightsModel := new(pushservicev1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(38))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(pushservicev1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}

				// Construct an instance of the Gcm model
				gcmModel := new(pushservicev1.Gcm)
				gcmModel.CollapseKey = core.StringPtr("testString")
				gcmModel.InteractiveCategory = core.StringPtr("testString")
				gcmModel.Icon = core.StringPtr("testString")
				gcmModel.DelayWhileIdle = core.BoolPtr(true)
				gcmModel.Sync = core.BoolPtr(true)
				gcmModel.Visibility = core.StringPtr("testString")
				gcmModel.Redact = core.StringPtr("testString")
				gcmModel.ChannelID = core.StringPtr("testString")
				gcmModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				gcmModel.Priority = core.StringPtr("testString")
				gcmModel.Sound = core.StringPtr("testString")
				gcmModel.TimeToLive = core.Int64Ptr(int64(38))
				gcmModel.Lights = lightsModel
				gcmModel.AndroidTitle = core.StringPtr("testString")
				gcmModel.GroupID = core.StringPtr("testString")
				gcmModel.Style = styleModel
				gcmModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the FirefoxWeb model
				firefoxWebModel := new(pushservicev1.FirefoxWeb)
				firefoxWebModel.Title = core.StringPtr("testString")
				firefoxWebModel.IconURL = core.StringPtr("testString")
				firefoxWebModel.TimeToLive = core.Int64Ptr(int64(38))
				firefoxWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the ChromeWeb model
				chromeWebModel := new(pushservicev1.ChromeWeb)
				chromeWebModel.Title = core.StringPtr("testString")
				chromeWebModel.IconURL = core.StringPtr("testString")
				chromeWebModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the SafariWeb model
				safariWebModel := new(pushservicev1.SafariWeb)
				safariWebModel.Title = core.StringPtr("testString")
				safariWebModel.UrlArgs = []string{"testString"}
				safariWebModel.Action = core.StringPtr("testString")

				// Construct an instance of the ChromeAppExt model
				chromeAppExtModel := new(pushservicev1.ChromeAppExt)
				chromeAppExtModel.CollapseKey = core.StringPtr("testString")
				chromeAppExtModel.DelayWhileIdle = core.BoolPtr(true)
				chromeAppExtModel.Title = core.StringPtr("testString")
				chromeAppExtModel.IconURL = core.StringPtr("testString")
				chromeAppExtModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeAppExtModel.Payload = core.StringPtr("testString")

				// Construct an instance of the Settings model
				settingsModel := new(pushservicev1.Settings)
				settingsModel.Apns = apnsModel
				settingsModel.Gcm = gcmModel
				settingsModel.FirefoxWeb = firefoxWebModel
				settingsModel.ChromeWeb = chromeWebModel
				settingsModel.SafariWeb = safariWebModel
				settingsModel.ChromeAppExt = chromeAppExtModel

				// Construct an instance of the Target model
				targetModel := new(pushservicev1.Target)
				targetModel.DeviceIds = []string{"testString"}
				targetModel.UserIds = []string{"testString"}
				targetModel.Platforms = []string{"testString"}
				targetModel.TagNames = []string{"testString"}

				// Construct an instance of the SendMessageOptions model
				sendMessageOptionsModel := new(pushservicev1.SendMessageOptions)
				sendMessageOptionsModel.ApplicationID = core.StringPtr("testString")
				sendMessageOptionsModel.Message = messageModel
				sendMessageOptionsModel.Settings = settingsModel
				sendMessageOptionsModel.Validate = core.BoolPtr(true)
				sendMessageOptionsModel.Target = targetModel
				sendMessageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				sendMessageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.SendMessage(sendMessageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SendMessageOptions model with no property values
				sendMessageOptionsModelNew := new(pushservicev1.SendMessageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.SendMessage(sendMessageOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SendMessagesInBulk(sendMessagesInBulkOptions *SendMessagesInBulkOptions) - Operation response error`, func() {
		sendMessagesInBulkPath := "/apps/testString/messages/bulk"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendMessagesInBulkPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SendMessagesInBulk with error: Operation response processing error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the Message model
				messageModel := new(pushservicev1.Message)
				messageModel.Alert = core.StringPtr("testString")
				messageModel.URL = core.StringPtr("testString")

				// Construct an instance of the Apns model
				apnsModel := new(pushservicev1.Apns)
				apnsModel.Badge = core.Int64Ptr(int64(38))
				apnsModel.InteractiveCategory = core.StringPtr("testString")
				apnsModel.Category = core.StringPtr("testString")
				apnsModel.IosActionKey = core.StringPtr("testString")
				apnsModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				apnsModel.Sound = core.StringPtr("testString")
				apnsModel.TitleLocKey = core.StringPtr("testString")
				apnsModel.LocKey = core.StringPtr("testString")
				apnsModel.LaunchImage = core.StringPtr("testString")
				apnsModel.TitleLocArgs = []string{"testString"}
				apnsModel.LocArgs = []string{"testString"}
				apnsModel.Title = core.StringPtr("testString")
				apnsModel.Subtitle = core.StringPtr("testString")
				apnsModel.AttachmentURL = core.StringPtr("testString")
				apnsModel.Type = core.StringPtr("DEFAULT")
				apnsModel.ApnsCollapseID = core.StringPtr("testString")
				apnsModel.ApnsThreadID = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Lights model
				lightsModel := new(pushservicev1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(38))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(pushservicev1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}

				// Construct an instance of the Gcm model
				gcmModel := new(pushservicev1.Gcm)
				gcmModel.CollapseKey = core.StringPtr("testString")
				gcmModel.InteractiveCategory = core.StringPtr("testString")
				gcmModel.Icon = core.StringPtr("testString")
				gcmModel.DelayWhileIdle = core.BoolPtr(true)
				gcmModel.Sync = core.BoolPtr(true)
				gcmModel.Visibility = core.StringPtr("testString")
				gcmModel.Redact = core.StringPtr("testString")
				gcmModel.ChannelID = core.StringPtr("testString")
				gcmModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				gcmModel.Priority = core.StringPtr("testString")
				gcmModel.Sound = core.StringPtr("testString")
				gcmModel.TimeToLive = core.Int64Ptr(int64(38))
				gcmModel.Lights = lightsModel
				gcmModel.AndroidTitle = core.StringPtr("testString")
				gcmModel.GroupID = core.StringPtr("testString")
				gcmModel.Style = styleModel
				gcmModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the FirefoxWeb model
				firefoxWebModel := new(pushservicev1.FirefoxWeb)
				firefoxWebModel.Title = core.StringPtr("testString")
				firefoxWebModel.IconURL = core.StringPtr("testString")
				firefoxWebModel.TimeToLive = core.Int64Ptr(int64(38))
				firefoxWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the ChromeWeb model
				chromeWebModel := new(pushservicev1.ChromeWeb)
				chromeWebModel.Title = core.StringPtr("testString")
				chromeWebModel.IconURL = core.StringPtr("testString")
				chromeWebModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the SafariWeb model
				safariWebModel := new(pushservicev1.SafariWeb)
				safariWebModel.Title = core.StringPtr("testString")
				safariWebModel.UrlArgs = []string{"testString"}
				safariWebModel.Action = core.StringPtr("testString")

				// Construct an instance of the ChromeAppExt model
				chromeAppExtModel := new(pushservicev1.ChromeAppExt)
				chromeAppExtModel.CollapseKey = core.StringPtr("testString")
				chromeAppExtModel.DelayWhileIdle = core.BoolPtr(true)
				chromeAppExtModel.Title = core.StringPtr("testString")
				chromeAppExtModel.IconURL = core.StringPtr("testString")
				chromeAppExtModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeAppExtModel.Payload = core.StringPtr("testString")

				// Construct an instance of the Settings model
				settingsModel := new(pushservicev1.Settings)
				settingsModel.Apns = apnsModel
				settingsModel.Gcm = gcmModel
				settingsModel.FirefoxWeb = firefoxWebModel
				settingsModel.ChromeWeb = chromeWebModel
				settingsModel.SafariWeb = safariWebModel
				settingsModel.ChromeAppExt = chromeAppExtModel

				// Construct an instance of the Target model
				targetModel := new(pushservicev1.Target)
				targetModel.DeviceIds = []string{"testString"}
				targetModel.UserIds = []string{"testString"}
				targetModel.Platforms = []string{"testString"}
				targetModel.TagNames = []string{"testString"}

				// Construct an instance of the SendMessageBody model
				sendMessageBodyModel := new(pushservicev1.SendMessageBody)
				sendMessageBodyModel.Message = messageModel
				sendMessageBodyModel.Settings = settingsModel
				sendMessageBodyModel.Validate = core.BoolPtr(true)
				sendMessageBodyModel.Target = targetModel

				// Construct an instance of the SendMessagesInBulkOptions model
				sendMessagesInBulkOptionsModel := new(pushservicev1.SendMessagesInBulkOptions)
				sendMessagesInBulkOptionsModel.ApplicationID = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.Body = []pushservicev1.SendMessageBody{*sendMessageBodyModel}
				sendMessagesInBulkOptionsModel.AcceptLanguage = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.AppSecret = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := pushServiceService.SendMessagesInBulk(sendMessagesInBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				pushServiceService.EnableRetries(0, 0)
				result, response, operationErr = pushServiceService.SendMessagesInBulk(sendMessagesInBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SendMessagesInBulk(sendMessagesInBulkOptions *SendMessagesInBulkOptions)`, func() {
		sendMessagesInBulkPath := "/apps/testString/messages/bulk"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendMessagesInBulkPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"messages": [{"createdTime": "2017-07-08T13:43:10.000Z", "messageId": "fnDWK8bG", "alert": "Sample message text", "href": "http://server/imfpush/v1/myapp/messages/fnDWK8bG"}]}`)
				}))
			})
			It(`Invoke SendMessagesInBulk successfully with retries`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())
				pushServiceService.EnableRetries(0, 0)

				// Construct an instance of the Message model
				messageModel := new(pushservicev1.Message)
				messageModel.Alert = core.StringPtr("testString")
				messageModel.URL = core.StringPtr("testString")

				// Construct an instance of the Apns model
				apnsModel := new(pushservicev1.Apns)
				apnsModel.Badge = core.Int64Ptr(int64(38))
				apnsModel.InteractiveCategory = core.StringPtr("testString")
				apnsModel.Category = core.StringPtr("testString")
				apnsModel.IosActionKey = core.StringPtr("testString")
				apnsModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				apnsModel.Sound = core.StringPtr("testString")
				apnsModel.TitleLocKey = core.StringPtr("testString")
				apnsModel.LocKey = core.StringPtr("testString")
				apnsModel.LaunchImage = core.StringPtr("testString")
				apnsModel.TitleLocArgs = []string{"testString"}
				apnsModel.LocArgs = []string{"testString"}
				apnsModel.Title = core.StringPtr("testString")
				apnsModel.Subtitle = core.StringPtr("testString")
				apnsModel.AttachmentURL = core.StringPtr("testString")
				apnsModel.Type = core.StringPtr("DEFAULT")
				apnsModel.ApnsCollapseID = core.StringPtr("testString")
				apnsModel.ApnsThreadID = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Lights model
				lightsModel := new(pushservicev1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(38))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(pushservicev1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}

				// Construct an instance of the Gcm model
				gcmModel := new(pushservicev1.Gcm)
				gcmModel.CollapseKey = core.StringPtr("testString")
				gcmModel.InteractiveCategory = core.StringPtr("testString")
				gcmModel.Icon = core.StringPtr("testString")
				gcmModel.DelayWhileIdle = core.BoolPtr(true)
				gcmModel.Sync = core.BoolPtr(true)
				gcmModel.Visibility = core.StringPtr("testString")
				gcmModel.Redact = core.StringPtr("testString")
				gcmModel.ChannelID = core.StringPtr("testString")
				gcmModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				gcmModel.Priority = core.StringPtr("testString")
				gcmModel.Sound = core.StringPtr("testString")
				gcmModel.TimeToLive = core.Int64Ptr(int64(38))
				gcmModel.Lights = lightsModel
				gcmModel.AndroidTitle = core.StringPtr("testString")
				gcmModel.GroupID = core.StringPtr("testString")
				gcmModel.Style = styleModel
				gcmModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the FirefoxWeb model
				firefoxWebModel := new(pushservicev1.FirefoxWeb)
				firefoxWebModel.Title = core.StringPtr("testString")
				firefoxWebModel.IconURL = core.StringPtr("testString")
				firefoxWebModel.TimeToLive = core.Int64Ptr(int64(38))
				firefoxWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the ChromeWeb model
				chromeWebModel := new(pushservicev1.ChromeWeb)
				chromeWebModel.Title = core.StringPtr("testString")
				chromeWebModel.IconURL = core.StringPtr("testString")
				chromeWebModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the SafariWeb model
				safariWebModel := new(pushservicev1.SafariWeb)
				safariWebModel.Title = core.StringPtr("testString")
				safariWebModel.UrlArgs = []string{"testString"}
				safariWebModel.Action = core.StringPtr("testString")

				// Construct an instance of the ChromeAppExt model
				chromeAppExtModel := new(pushservicev1.ChromeAppExt)
				chromeAppExtModel.CollapseKey = core.StringPtr("testString")
				chromeAppExtModel.DelayWhileIdle = core.BoolPtr(true)
				chromeAppExtModel.Title = core.StringPtr("testString")
				chromeAppExtModel.IconURL = core.StringPtr("testString")
				chromeAppExtModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeAppExtModel.Payload = core.StringPtr("testString")

				// Construct an instance of the Settings model
				settingsModel := new(pushservicev1.Settings)
				settingsModel.Apns = apnsModel
				settingsModel.Gcm = gcmModel
				settingsModel.FirefoxWeb = firefoxWebModel
				settingsModel.ChromeWeb = chromeWebModel
				settingsModel.SafariWeb = safariWebModel
				settingsModel.ChromeAppExt = chromeAppExtModel

				// Construct an instance of the Target model
				targetModel := new(pushservicev1.Target)
				targetModel.DeviceIds = []string{"testString"}
				targetModel.UserIds = []string{"testString"}
				targetModel.Platforms = []string{"testString"}
				targetModel.TagNames = []string{"testString"}

				// Construct an instance of the SendMessageBody model
				sendMessageBodyModel := new(pushservicev1.SendMessageBody)
				sendMessageBodyModel.Message = messageModel
				sendMessageBodyModel.Settings = settingsModel
				sendMessageBodyModel.Validate = core.BoolPtr(true)
				sendMessageBodyModel.Target = targetModel

				// Construct an instance of the SendMessagesInBulkOptions model
				sendMessagesInBulkOptionsModel := new(pushservicev1.SendMessagesInBulkOptions)
				sendMessagesInBulkOptionsModel.ApplicationID = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.Body = []pushservicev1.SendMessageBody{*sendMessageBodyModel}
				sendMessagesInBulkOptionsModel.AcceptLanguage = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.AppSecret = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := pushServiceService.SendMessagesInBulkWithContext(ctx, sendMessagesInBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				pushServiceService.DisableRetries()
				result, response, operationErr := pushServiceService.SendMessagesInBulk(sendMessagesInBulkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = pushServiceService.SendMessagesInBulkWithContext(ctx, sendMessagesInBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(sendMessagesInBulkPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Appsecret"]).ToNot(BeNil())
					Expect(req.Header["Appsecret"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"messages": [{"createdTime": "2017-07-08T13:43:10.000Z", "messageId": "fnDWK8bG", "alert": "Sample message text", "href": "http://server/imfpush/v1/myapp/messages/fnDWK8bG"}]}`)
				}))
			})
			It(`Invoke SendMessagesInBulk successfully`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := pushServiceService.SendMessagesInBulk(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Message model
				messageModel := new(pushservicev1.Message)
				messageModel.Alert = core.StringPtr("testString")
				messageModel.URL = core.StringPtr("testString")

				// Construct an instance of the Apns model
				apnsModel := new(pushservicev1.Apns)
				apnsModel.Badge = core.Int64Ptr(int64(38))
				apnsModel.InteractiveCategory = core.StringPtr("testString")
				apnsModel.Category = core.StringPtr("testString")
				apnsModel.IosActionKey = core.StringPtr("testString")
				apnsModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				apnsModel.Sound = core.StringPtr("testString")
				apnsModel.TitleLocKey = core.StringPtr("testString")
				apnsModel.LocKey = core.StringPtr("testString")
				apnsModel.LaunchImage = core.StringPtr("testString")
				apnsModel.TitleLocArgs = []string{"testString"}
				apnsModel.LocArgs = []string{"testString"}
				apnsModel.Title = core.StringPtr("testString")
				apnsModel.Subtitle = core.StringPtr("testString")
				apnsModel.AttachmentURL = core.StringPtr("testString")
				apnsModel.Type = core.StringPtr("DEFAULT")
				apnsModel.ApnsCollapseID = core.StringPtr("testString")
				apnsModel.ApnsThreadID = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Lights model
				lightsModel := new(pushservicev1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(38))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(pushservicev1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}

				// Construct an instance of the Gcm model
				gcmModel := new(pushservicev1.Gcm)
				gcmModel.CollapseKey = core.StringPtr("testString")
				gcmModel.InteractiveCategory = core.StringPtr("testString")
				gcmModel.Icon = core.StringPtr("testString")
				gcmModel.DelayWhileIdle = core.BoolPtr(true)
				gcmModel.Sync = core.BoolPtr(true)
				gcmModel.Visibility = core.StringPtr("testString")
				gcmModel.Redact = core.StringPtr("testString")
				gcmModel.ChannelID = core.StringPtr("testString")
				gcmModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				gcmModel.Priority = core.StringPtr("testString")
				gcmModel.Sound = core.StringPtr("testString")
				gcmModel.TimeToLive = core.Int64Ptr(int64(38))
				gcmModel.Lights = lightsModel
				gcmModel.AndroidTitle = core.StringPtr("testString")
				gcmModel.GroupID = core.StringPtr("testString")
				gcmModel.Style = styleModel
				gcmModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the FirefoxWeb model
				firefoxWebModel := new(pushservicev1.FirefoxWeb)
				firefoxWebModel.Title = core.StringPtr("testString")
				firefoxWebModel.IconURL = core.StringPtr("testString")
				firefoxWebModel.TimeToLive = core.Int64Ptr(int64(38))
				firefoxWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the ChromeWeb model
				chromeWebModel := new(pushservicev1.ChromeWeb)
				chromeWebModel.Title = core.StringPtr("testString")
				chromeWebModel.IconURL = core.StringPtr("testString")
				chromeWebModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the SafariWeb model
				safariWebModel := new(pushservicev1.SafariWeb)
				safariWebModel.Title = core.StringPtr("testString")
				safariWebModel.UrlArgs = []string{"testString"}
				safariWebModel.Action = core.StringPtr("testString")

				// Construct an instance of the ChromeAppExt model
				chromeAppExtModel := new(pushservicev1.ChromeAppExt)
				chromeAppExtModel.CollapseKey = core.StringPtr("testString")
				chromeAppExtModel.DelayWhileIdle = core.BoolPtr(true)
				chromeAppExtModel.Title = core.StringPtr("testString")
				chromeAppExtModel.IconURL = core.StringPtr("testString")
				chromeAppExtModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeAppExtModel.Payload = core.StringPtr("testString")

				// Construct an instance of the Settings model
				settingsModel := new(pushservicev1.Settings)
				settingsModel.Apns = apnsModel
				settingsModel.Gcm = gcmModel
				settingsModel.FirefoxWeb = firefoxWebModel
				settingsModel.ChromeWeb = chromeWebModel
				settingsModel.SafariWeb = safariWebModel
				settingsModel.ChromeAppExt = chromeAppExtModel

				// Construct an instance of the Target model
				targetModel := new(pushservicev1.Target)
				targetModel.DeviceIds = []string{"testString"}
				targetModel.UserIds = []string{"testString"}
				targetModel.Platforms = []string{"testString"}
				targetModel.TagNames = []string{"testString"}

				// Construct an instance of the SendMessageBody model
				sendMessageBodyModel := new(pushservicev1.SendMessageBody)
				sendMessageBodyModel.Message = messageModel
				sendMessageBodyModel.Settings = settingsModel
				sendMessageBodyModel.Validate = core.BoolPtr(true)
				sendMessageBodyModel.Target = targetModel

				// Construct an instance of the SendMessagesInBulkOptions model
				sendMessagesInBulkOptionsModel := new(pushservicev1.SendMessagesInBulkOptions)
				sendMessagesInBulkOptionsModel.ApplicationID = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.Body = []pushservicev1.SendMessageBody{*sendMessageBodyModel}
				sendMessagesInBulkOptionsModel.AcceptLanguage = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.AppSecret = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = pushServiceService.SendMessagesInBulk(sendMessagesInBulkOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SendMessagesInBulk with error: Operation validation and request error`, func() {
				pushServiceService, serviceErr := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(pushServiceService).ToNot(BeNil())

				// Construct an instance of the Message model
				messageModel := new(pushservicev1.Message)
				messageModel.Alert = core.StringPtr("testString")
				messageModel.URL = core.StringPtr("testString")

				// Construct an instance of the Apns model
				apnsModel := new(pushservicev1.Apns)
				apnsModel.Badge = core.Int64Ptr(int64(38))
				apnsModel.InteractiveCategory = core.StringPtr("testString")
				apnsModel.Category = core.StringPtr("testString")
				apnsModel.IosActionKey = core.StringPtr("testString")
				apnsModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				apnsModel.Sound = core.StringPtr("testString")
				apnsModel.TitleLocKey = core.StringPtr("testString")
				apnsModel.LocKey = core.StringPtr("testString")
				apnsModel.LaunchImage = core.StringPtr("testString")
				apnsModel.TitleLocArgs = []string{"testString"}
				apnsModel.LocArgs = []string{"testString"}
				apnsModel.Title = core.StringPtr("testString")
				apnsModel.Subtitle = core.StringPtr("testString")
				apnsModel.AttachmentURL = core.StringPtr("testString")
				apnsModel.Type = core.StringPtr("DEFAULT")
				apnsModel.ApnsCollapseID = core.StringPtr("testString")
				apnsModel.ApnsThreadID = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))

				// Construct an instance of the Lights model
				lightsModel := new(pushservicev1.Lights)
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(38))
				lightsModel.LedOffMs = core.StringPtr("testString")

				// Construct an instance of the Style model
				styleModel := new(pushservicev1.Style)
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}

				// Construct an instance of the Gcm model
				gcmModel := new(pushservicev1.Gcm)
				gcmModel.CollapseKey = core.StringPtr("testString")
				gcmModel.InteractiveCategory = core.StringPtr("testString")
				gcmModel.Icon = core.StringPtr("testString")
				gcmModel.DelayWhileIdle = core.BoolPtr(true)
				gcmModel.Sync = core.BoolPtr(true)
				gcmModel.Visibility = core.StringPtr("testString")
				gcmModel.Redact = core.StringPtr("testString")
				gcmModel.ChannelID = core.StringPtr("testString")
				gcmModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				gcmModel.Priority = core.StringPtr("testString")
				gcmModel.Sound = core.StringPtr("testString")
				gcmModel.TimeToLive = core.Int64Ptr(int64(38))
				gcmModel.Lights = lightsModel
				gcmModel.AndroidTitle = core.StringPtr("testString")
				gcmModel.GroupID = core.StringPtr("testString")
				gcmModel.Style = styleModel
				gcmModel.Type = core.StringPtr("DEFAULT")

				// Construct an instance of the FirefoxWeb model
				firefoxWebModel := new(pushservicev1.FirefoxWeb)
				firefoxWebModel.Title = core.StringPtr("testString")
				firefoxWebModel.IconURL = core.StringPtr("testString")
				firefoxWebModel.TimeToLive = core.Int64Ptr(int64(38))
				firefoxWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the ChromeWeb model
				chromeWebModel := new(pushservicev1.ChromeWeb)
				chromeWebModel.Title = core.StringPtr("testString")
				chromeWebModel.IconURL = core.StringPtr("testString")
				chromeWebModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeWebModel.Payload = core.StringPtr("testString")

				// Construct an instance of the SafariWeb model
				safariWebModel := new(pushservicev1.SafariWeb)
				safariWebModel.Title = core.StringPtr("testString")
				safariWebModel.UrlArgs = []string{"testString"}
				safariWebModel.Action = core.StringPtr("testString")

				// Construct an instance of the ChromeAppExt model
				chromeAppExtModel := new(pushservicev1.ChromeAppExt)
				chromeAppExtModel.CollapseKey = core.StringPtr("testString")
				chromeAppExtModel.DelayWhileIdle = core.BoolPtr(true)
				chromeAppExtModel.Title = core.StringPtr("testString")
				chromeAppExtModel.IconURL = core.StringPtr("testString")
				chromeAppExtModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeAppExtModel.Payload = core.StringPtr("testString")

				// Construct an instance of the Settings model
				settingsModel := new(pushservicev1.Settings)
				settingsModel.Apns = apnsModel
				settingsModel.Gcm = gcmModel
				settingsModel.FirefoxWeb = firefoxWebModel
				settingsModel.ChromeWeb = chromeWebModel
				settingsModel.SafariWeb = safariWebModel
				settingsModel.ChromeAppExt = chromeAppExtModel

				// Construct an instance of the Target model
				targetModel := new(pushservicev1.Target)
				targetModel.DeviceIds = []string{"testString"}
				targetModel.UserIds = []string{"testString"}
				targetModel.Platforms = []string{"testString"}
				targetModel.TagNames = []string{"testString"}

				// Construct an instance of the SendMessageBody model
				sendMessageBodyModel := new(pushservicev1.SendMessageBody)
				sendMessageBodyModel.Message = messageModel
				sendMessageBodyModel.Settings = settingsModel
				sendMessageBodyModel.Validate = core.BoolPtr(true)
				sendMessageBodyModel.Target = targetModel

				// Construct an instance of the SendMessagesInBulkOptions model
				sendMessagesInBulkOptionsModel := new(pushservicev1.SendMessagesInBulkOptions)
				sendMessagesInBulkOptionsModel.ApplicationID = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.Body = []pushservicev1.SendMessageBody{*sendMessageBodyModel}
				sendMessagesInBulkOptionsModel.AcceptLanguage = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.AppSecret = core.StringPtr("testString")
				sendMessagesInBulkOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := pushServiceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := pushServiceService.SendMessagesInBulk(sendMessagesInBulkOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SendMessagesInBulkOptions model with no property values
				sendMessagesInBulkOptionsModelNew := new(pushservicev1.SendMessagesInBulkOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = pushServiceService.SendMessagesInBulk(sendMessagesInBulkOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			pushServiceService, _ := pushservicev1.NewPushServiceV1(&pushservicev1.PushServiceV1Options{
				URL:           "http://pushservicev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewChromeWebPushCredendialsModel successfully`, func() {
				apiKey := "testString"
				webSiteURL := "testString"
				model, err := pushServiceService.NewChromeWebPushCredendialsModel(apiKey, webSiteURL)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewDeleteApnsConfOptions successfully`, func() {
				// Construct an instance of the DeleteApnsConfOptions model
				applicationID := "testString"
				deleteApnsConfOptionsModel := pushServiceService.NewDeleteApnsConfOptions(applicationID)
				deleteApnsConfOptionsModel.SetApplicationID("testString")
				deleteApnsConfOptionsModel.SetAcceptLanguage("testString")
				deleteApnsConfOptionsModel.SetAppSecret("testString")
				deleteApnsConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteApnsConfOptionsModel).ToNot(BeNil())
				Expect(deleteApnsConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteApnsConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteApnsConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(deleteApnsConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteChromeAppExtConfOptions successfully`, func() {
				// Construct an instance of the DeleteChromeAppExtConfOptions model
				applicationID := "testString"
				deleteChromeAppExtConfOptionsModel := pushServiceService.NewDeleteChromeAppExtConfOptions(applicationID)
				deleteChromeAppExtConfOptionsModel.SetApplicationID("testString")
				deleteChromeAppExtConfOptionsModel.SetAcceptLanguage("testString")
				deleteChromeAppExtConfOptionsModel.SetAppSecret("testString")
				deleteChromeAppExtConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteChromeAppExtConfOptionsModel).ToNot(BeNil())
				Expect(deleteChromeAppExtConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteChromeAppExtConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteChromeAppExtConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(deleteChromeAppExtConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteChromeWebConfOptions successfully`, func() {
				// Construct an instance of the DeleteChromeWebConfOptions model
				applicationID := "testString"
				deleteChromeWebConfOptionsModel := pushServiceService.NewDeleteChromeWebConfOptions(applicationID)
				deleteChromeWebConfOptionsModel.SetApplicationID("testString")
				deleteChromeWebConfOptionsModel.SetAcceptLanguage("testString")
				deleteChromeWebConfOptionsModel.SetAppSecret("testString")
				deleteChromeWebConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteChromeWebConfOptionsModel).ToNot(BeNil())
				Expect(deleteChromeWebConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteChromeWebConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteChromeWebConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(deleteChromeWebConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFirefoxWebConfOptions successfully`, func() {
				// Construct an instance of the DeleteFirefoxWebConfOptions model
				applicationID := "testString"
				deleteFirefoxWebConfOptionsModel := pushServiceService.NewDeleteFirefoxWebConfOptions(applicationID)
				deleteFirefoxWebConfOptionsModel.SetApplicationID("testString")
				deleteFirefoxWebConfOptionsModel.SetAcceptLanguage("testString")
				deleteFirefoxWebConfOptionsModel.SetAppSecret("testString")
				deleteFirefoxWebConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFirefoxWebConfOptionsModel).ToNot(BeNil())
				Expect(deleteFirefoxWebConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirefoxWebConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirefoxWebConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(deleteFirefoxWebConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteGCMConfOptions successfully`, func() {
				// Construct an instance of the DeleteGCMConfOptions model
				applicationID := "testString"
				deleteGcmConfOptionsModel := pushServiceService.NewDeleteGCMConfOptions(applicationID)
				deleteGcmConfOptionsModel.SetApplicationID("testString")
				deleteGcmConfOptionsModel.SetAcceptLanguage("testString")
				deleteGcmConfOptionsModel.SetAppSecret("testString")
				deleteGcmConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteGcmConfOptionsModel).ToNot(BeNil())
				Expect(deleteGcmConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteGcmConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteGcmConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(deleteGcmConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSafariWebConfOptions successfully`, func() {
				// Construct an instance of the DeleteSafariWebConfOptions model
				applicationID := "testString"
				deleteSafariWebConfOptionsModel := pushServiceService.NewDeleteSafariWebConfOptions(applicationID)
				deleteSafariWebConfOptionsModel.SetApplicationID("testString")
				deleteSafariWebConfOptionsModel.SetAcceptLanguage("testString")
				deleteSafariWebConfOptionsModel.SetAppSecret("testString")
				deleteSafariWebConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSafariWebConfOptionsModel).ToNot(BeNil())
				Expect(deleteSafariWebConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSafariWebConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteSafariWebConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(deleteSafariWebConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewFirefoxWebPushCredendialsModel successfully`, func() {
				webSiteURL := "testString"
				model, err := pushServiceService.NewFirefoxWebPushCredendialsModel(webSiteURL)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGCMCredendialsModel successfully`, func() {
				apiKey := "testString"
				senderID := "testString"
				model, err := pushServiceService.NewGCMCredendialsModel(apiKey, senderID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetApnsConfOptions successfully`, func() {
				// Construct an instance of the GetApnsConfOptions model
				applicationID := "testString"
				getApnsConfOptionsModel := pushServiceService.NewGetApnsConfOptions(applicationID)
				getApnsConfOptionsModel.SetApplicationID("testString")
				getApnsConfOptionsModel.SetAcceptLanguage("testString")
				getApnsConfOptionsModel.SetAppSecret("testString")
				getApnsConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getApnsConfOptionsModel).ToNot(BeNil())
				Expect(getApnsConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getApnsConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getApnsConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(getApnsConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetChromeAppExtConfOptions successfully`, func() {
				// Construct an instance of the GetChromeAppExtConfOptions model
				applicationID := "testString"
				getChromeAppExtConfOptionsModel := pushServiceService.NewGetChromeAppExtConfOptions(applicationID)
				getChromeAppExtConfOptionsModel.SetApplicationID("testString")
				getChromeAppExtConfOptionsModel.SetAcceptLanguage("testString")
				getChromeAppExtConfOptionsModel.SetAppSecret("testString")
				getChromeAppExtConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getChromeAppExtConfOptionsModel).ToNot(BeNil())
				Expect(getChromeAppExtConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getChromeAppExtConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getChromeAppExtConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(getChromeAppExtConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetChromeAppExtConfPublicOptions successfully`, func() {
				// Construct an instance of the GetChromeAppExtConfPublicOptions model
				applicationID := "testString"
				getChromeAppExtConfPublicOptionsModel := pushServiceService.NewGetChromeAppExtConfPublicOptions(applicationID)
				getChromeAppExtConfPublicOptionsModel.SetApplicationID("testString")
				getChromeAppExtConfPublicOptionsModel.SetClientSecret("testString")
				getChromeAppExtConfPublicOptionsModel.SetAcceptLanguage("testString")
				getChromeAppExtConfPublicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getChromeAppExtConfPublicOptionsModel).ToNot(BeNil())
				Expect(getChromeAppExtConfPublicOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getChromeAppExtConfPublicOptionsModel.ClientSecret).To(Equal(core.StringPtr("testString")))
				Expect(getChromeAppExtConfPublicOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getChromeAppExtConfPublicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetChromeWebConfOptions successfully`, func() {
				// Construct an instance of the GetChromeWebConfOptions model
				applicationID := "testString"
				getChromeWebConfOptionsModel := pushServiceService.NewGetChromeWebConfOptions(applicationID)
				getChromeWebConfOptionsModel.SetApplicationID("testString")
				getChromeWebConfOptionsModel.SetAcceptLanguage("testString")
				getChromeWebConfOptionsModel.SetAppSecret("testString")
				getChromeWebConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getChromeWebConfOptionsModel).ToNot(BeNil())
				Expect(getChromeWebConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getChromeWebConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getChromeWebConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(getChromeWebConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFirefoxWebConfOptions successfully`, func() {
				// Construct an instance of the GetFirefoxWebConfOptions model
				applicationID := "testString"
				getFirefoxWebConfOptionsModel := pushServiceService.NewGetFirefoxWebConfOptions(applicationID)
				getFirefoxWebConfOptionsModel.SetApplicationID("testString")
				getFirefoxWebConfOptionsModel.SetAcceptLanguage("testString")
				getFirefoxWebConfOptionsModel.SetAppSecret("testString")
				getFirefoxWebConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFirefoxWebConfOptionsModel).ToNot(BeNil())
				Expect(getFirefoxWebConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getFirefoxWebConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getFirefoxWebConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(getFirefoxWebConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGCMConfOptions successfully`, func() {
				// Construct an instance of the GetGCMConfOptions model
				applicationID := "testString"
				getGcmConfOptionsModel := pushServiceService.NewGetGCMConfOptions(applicationID)
				getGcmConfOptionsModel.SetApplicationID("testString")
				getGcmConfOptionsModel.SetAcceptLanguage("testString")
				getGcmConfOptionsModel.SetAppSecret("testString")
				getGcmConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGcmConfOptionsModel).ToNot(BeNil())
				Expect(getGcmConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getGcmConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getGcmConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(getGcmConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetGcmConfPublicOptions successfully`, func() {
				// Construct an instance of the GetGcmConfPublicOptions model
				applicationID := "testString"
				getGcmConfPublicOptionsModel := pushServiceService.NewGetGcmConfPublicOptions(applicationID)
				getGcmConfPublicOptionsModel.SetApplicationID("testString")
				getGcmConfPublicOptionsModel.SetClientSecret("testString")
				getGcmConfPublicOptionsModel.SetAcceptLanguage("testString")
				getGcmConfPublicOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getGcmConfPublicOptionsModel).ToNot(BeNil())
				Expect(getGcmConfPublicOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getGcmConfPublicOptionsModel.ClientSecret).To(Equal(core.StringPtr("testString")))
				Expect(getGcmConfPublicOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getGcmConfPublicOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSafariWebConfOptions successfully`, func() {
				// Construct an instance of the GetSafariWebConfOptions model
				applicationID := "testString"
				getSafariWebConfOptionsModel := pushServiceService.NewGetSafariWebConfOptions(applicationID)
				getSafariWebConfOptionsModel.SetApplicationID("testString")
				getSafariWebConfOptionsModel.SetAcceptLanguage("testString")
				getSafariWebConfOptionsModel.SetAppSecret("testString")
				getSafariWebConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSafariWebConfOptionsModel).ToNot(BeNil())
				Expect(getSafariWebConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getSafariWebConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getSafariWebConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(getSafariWebConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				applicationID := "testString"
				getSettingsOptionsModel := pushServiceService.NewGetSettingsOptions(applicationID)
				getSettingsOptionsModel.SetApplicationID("testString")
				getSettingsOptionsModel.SetAppSecret("testString")
				getSettingsOptionsModel.SetAcceptLanguage("testString")
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getSettingsOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(getSettingsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWebpushServerKeyOptions successfully`, func() {
				// Construct an instance of the GetWebpushServerKeyOptions model
				applicationID := "testString"
				getWebpushServerKeyOptionsModel := pushServiceService.NewGetWebpushServerKeyOptions(applicationID)
				getWebpushServerKeyOptionsModel.SetApplicationID("testString")
				getWebpushServerKeyOptionsModel.SetClientSecret("testString")
				getWebpushServerKeyOptionsModel.SetAcceptLanguage("testString")
				getWebpushServerKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWebpushServerKeyOptionsModel).ToNot(BeNil())
				Expect(getWebpushServerKeyOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(getWebpushServerKeyOptionsModel.ClientSecret).To(Equal(core.StringPtr("testString")))
				Expect(getWebpushServerKeyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getWebpushServerKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSaveApnsConfOptions successfully`, func() {
				// Construct an instance of the SaveApnsConfOptions model
				applicationID := "testString"
				password := "testString"
				isSandBox := true
				certificate := CreateMockReader("This is a mock file.")
				saveApnsConfOptionsModel := pushServiceService.NewSaveApnsConfOptions(applicationID, password, isSandBox, certificate)
				saveApnsConfOptionsModel.SetApplicationID("testString")
				saveApnsConfOptionsModel.SetPassword("testString")
				saveApnsConfOptionsModel.SetIsSandBox(true)
				saveApnsConfOptionsModel.SetCertificate(CreateMockReader("This is a mock file."))
				saveApnsConfOptionsModel.SetCertificateContentType("testString")
				saveApnsConfOptionsModel.SetAcceptLanguage("testString")
				saveApnsConfOptionsModel.SetAppSecret("testString")
				saveApnsConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(saveApnsConfOptionsModel).ToNot(BeNil())
				Expect(saveApnsConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(saveApnsConfOptionsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(saveApnsConfOptionsModel.IsSandBox).To(Equal(core.BoolPtr(true)))
				Expect(saveApnsConfOptionsModel.Certificate).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(saveApnsConfOptionsModel.CertificateContentType).To(Equal(core.StringPtr("testString")))
				Expect(saveApnsConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(saveApnsConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(saveApnsConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSaveChromeAppExtConfOptions successfully`, func() {
				// Construct an instance of the SaveChromeAppExtConfOptions model
				applicationID := "testString"
				saveChromeAppExtConfOptionsApiKey := "testString"
				saveChromeAppExtConfOptionsSenderID := "testString"
				saveChromeAppExtConfOptionsModel := pushServiceService.NewSaveChromeAppExtConfOptions(applicationID, saveChromeAppExtConfOptionsApiKey, saveChromeAppExtConfOptionsSenderID)
				saveChromeAppExtConfOptionsModel.SetApplicationID("testString")
				saveChromeAppExtConfOptionsModel.SetApiKey("testString")
				saveChromeAppExtConfOptionsModel.SetSenderID("testString")
				saveChromeAppExtConfOptionsModel.SetAcceptLanguage("testString")
				saveChromeAppExtConfOptionsModel.SetAppSecret("testString")
				saveChromeAppExtConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(saveChromeAppExtConfOptionsModel).ToNot(BeNil())
				Expect(saveChromeAppExtConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(saveChromeAppExtConfOptionsModel.ApiKey).To(Equal(core.StringPtr("testString")))
				Expect(saveChromeAppExtConfOptionsModel.SenderID).To(Equal(core.StringPtr("testString")))
				Expect(saveChromeAppExtConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(saveChromeAppExtConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(saveChromeAppExtConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSaveChromeWebConfOptions successfully`, func() {
				// Construct an instance of the SaveChromeWebConfOptions model
				applicationID := "testString"
				saveChromeWebConfOptionsApiKey := "testString"
				saveChromeWebConfOptionsWebSiteURL := "testString"
				saveChromeWebConfOptionsModel := pushServiceService.NewSaveChromeWebConfOptions(applicationID, saveChromeWebConfOptionsApiKey, saveChromeWebConfOptionsWebSiteURL)
				saveChromeWebConfOptionsModel.SetApplicationID("testString")
				saveChromeWebConfOptionsModel.SetApiKey("testString")
				saveChromeWebConfOptionsModel.SetWebSiteURL("testString")
				saveChromeWebConfOptionsModel.SetAcceptLanguage("testString")
				saveChromeWebConfOptionsModel.SetAppSecret("testString")
				saveChromeWebConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(saveChromeWebConfOptionsModel).ToNot(BeNil())
				Expect(saveChromeWebConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(saveChromeWebConfOptionsModel.ApiKey).To(Equal(core.StringPtr("testString")))
				Expect(saveChromeWebConfOptionsModel.WebSiteURL).To(Equal(core.StringPtr("testString")))
				Expect(saveChromeWebConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(saveChromeWebConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(saveChromeWebConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSaveFirefoxWebConfOptions successfully`, func() {
				// Construct an instance of the SaveFirefoxWebConfOptions model
				applicationID := "testString"
				saveFirefoxWebConfOptionsWebSiteURL := "testString"
				saveFirefoxWebConfOptionsModel := pushServiceService.NewSaveFirefoxWebConfOptions(applicationID, saveFirefoxWebConfOptionsWebSiteURL)
				saveFirefoxWebConfOptionsModel.SetApplicationID("testString")
				saveFirefoxWebConfOptionsModel.SetWebSiteURL("testString")
				saveFirefoxWebConfOptionsModel.SetAcceptLanguage("testString")
				saveFirefoxWebConfOptionsModel.SetAppSecret("testString")
				saveFirefoxWebConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(saveFirefoxWebConfOptionsModel).ToNot(BeNil())
				Expect(saveFirefoxWebConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(saveFirefoxWebConfOptionsModel.WebSiteURL).To(Equal(core.StringPtr("testString")))
				Expect(saveFirefoxWebConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(saveFirefoxWebConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(saveFirefoxWebConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSaveGCMConfOptions successfully`, func() {
				// Construct an instance of the SaveGCMConfOptions model
				applicationID := "testString"
				saveGcmConfOptionsApiKey := "testString"
				saveGcmConfOptionsSenderID := "testString"
				saveGcmConfOptionsModel := pushServiceService.NewSaveGCMConfOptions(applicationID, saveGcmConfOptionsApiKey, saveGcmConfOptionsSenderID)
				saveGcmConfOptionsModel.SetApplicationID("testString")
				saveGcmConfOptionsModel.SetApiKey("testString")
				saveGcmConfOptionsModel.SetSenderID("testString")
				saveGcmConfOptionsModel.SetAcceptLanguage("testString")
				saveGcmConfOptionsModel.SetAppSecret("testString")
				saveGcmConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(saveGcmConfOptionsModel).ToNot(BeNil())
				Expect(saveGcmConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(saveGcmConfOptionsModel.ApiKey).To(Equal(core.StringPtr("testString")))
				Expect(saveGcmConfOptionsModel.SenderID).To(Equal(core.StringPtr("testString")))
				Expect(saveGcmConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(saveGcmConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(saveGcmConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSaveSafariWebConfOptions successfully`, func() {
				// Construct an instance of the SaveSafariWebConfOptions model
				applicationID := "testString"
				password := "testString"
				certificate := CreateMockReader("This is a mock file.")
				websiteName := "testString"
				urlFormatString := "testString"
				websitePushID := "testString"
				webSiteURL := "testString"
				saveSafariWebConfOptionsModel := pushServiceService.NewSaveSafariWebConfOptions(applicationID, password, certificate, websiteName, urlFormatString, websitePushID, webSiteURL)
				saveSafariWebConfOptionsModel.SetApplicationID("testString")
				saveSafariWebConfOptionsModel.SetPassword("testString")
				saveSafariWebConfOptionsModel.SetCertificate(CreateMockReader("This is a mock file."))
				saveSafariWebConfOptionsModel.SetWebsiteName("testString")
				saveSafariWebConfOptionsModel.SetUrlFormatString("testString")
				saveSafariWebConfOptionsModel.SetWebsitePushID("testString")
				saveSafariWebConfOptionsModel.SetWebSiteURL("testString")
				saveSafariWebConfOptionsModel.SetCertificateContentType("testString")
				saveSafariWebConfOptionsModel.SetIcon16x16(CreateMockReader("This is a mock file."))
				saveSafariWebConfOptionsModel.SetIcon16x16ContentType("testString")
				saveSafariWebConfOptionsModel.SetIcon16x162x(CreateMockReader("This is a mock file."))
				saveSafariWebConfOptionsModel.SetIcon16x162xContentType("testString")
				saveSafariWebConfOptionsModel.SetIcon32x32(CreateMockReader("This is a mock file."))
				saveSafariWebConfOptionsModel.SetIcon32x32ContentType("testString")
				saveSafariWebConfOptionsModel.SetIcon32x322x(CreateMockReader("This is a mock file."))
				saveSafariWebConfOptionsModel.SetIcon32x322xContentType("testString")
				saveSafariWebConfOptionsModel.SetIcon128x128(CreateMockReader("This is a mock file."))
				saveSafariWebConfOptionsModel.SetIcon128x128ContentType("testString")
				saveSafariWebConfOptionsModel.SetIcon128x1282x(CreateMockReader("This is a mock file."))
				saveSafariWebConfOptionsModel.SetIcon128x1282xContentType("testString")
				saveSafariWebConfOptionsModel.SetAcceptLanguage("testString")
				saveSafariWebConfOptionsModel.SetAppSecret("testString")
				saveSafariWebConfOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(saveSafariWebConfOptionsModel).ToNot(BeNil())
				Expect(saveSafariWebConfOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.Certificate).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(saveSafariWebConfOptionsModel.WebsiteName).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.UrlFormatString).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.WebsitePushID).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.WebSiteURL).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.CertificateContentType).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.Icon16x16).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(saveSafariWebConfOptionsModel.Icon16x16ContentType).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.Icon16x162x).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(saveSafariWebConfOptionsModel.Icon16x162xContentType).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.Icon32x32).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(saveSafariWebConfOptionsModel.Icon32x32ContentType).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.Icon32x322x).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(saveSafariWebConfOptionsModel.Icon32x322xContentType).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.Icon128x128).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(saveSafariWebConfOptionsModel.Icon128x128ContentType).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.Icon128x1282x).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(saveSafariWebConfOptionsModel.Icon128x1282xContentType).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(saveSafariWebConfOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSendMessageOptions successfully`, func() {
				// Construct an instance of the Message model
				messageModel := new(pushservicev1.Message)
				Expect(messageModel).ToNot(BeNil())
				messageModel.Alert = core.StringPtr("testString")
				messageModel.URL = core.StringPtr("testString")
				Expect(messageModel.Alert).To(Equal(core.StringPtr("testString")))
				Expect(messageModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Apns model
				apnsModel := new(pushservicev1.Apns)
				Expect(apnsModel).ToNot(BeNil())
				apnsModel.Badge = core.Int64Ptr(int64(38))
				apnsModel.InteractiveCategory = core.StringPtr("testString")
				apnsModel.Category = core.StringPtr("testString")
				apnsModel.IosActionKey = core.StringPtr("testString")
				apnsModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				apnsModel.Sound = core.StringPtr("testString")
				apnsModel.TitleLocKey = core.StringPtr("testString")
				apnsModel.LocKey = core.StringPtr("testString")
				apnsModel.LaunchImage = core.StringPtr("testString")
				apnsModel.TitleLocArgs = []string{"testString"}
				apnsModel.LocArgs = []string{"testString"}
				apnsModel.Title = core.StringPtr("testString")
				apnsModel.Subtitle = core.StringPtr("testString")
				apnsModel.AttachmentURL = core.StringPtr("testString")
				apnsModel.Type = core.StringPtr("DEFAULT")
				apnsModel.ApnsCollapseID = core.StringPtr("testString")
				apnsModel.ApnsThreadID = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))
				Expect(apnsModel.Badge).To(Equal(core.Int64Ptr(int64(38))))
				Expect(apnsModel.InteractiveCategory).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.IosActionKey).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.Payload).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(apnsModel.Sound).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.TitleLocKey).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.LocKey).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.LaunchImage).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.TitleLocArgs).To(Equal([]string{"testString"}))
				Expect(apnsModel.LocArgs).To(Equal([]string{"testString"}))
				Expect(apnsModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.Subtitle).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.AttachmentURL).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.Type).To(Equal(core.StringPtr("DEFAULT")))
				Expect(apnsModel.ApnsCollapseID).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.ApnsThreadID).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.ApnsGroupSummaryArg).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.ApnsGroupSummaryArgCount).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the Lights model
				lightsModel := new(pushservicev1.Lights)
				Expect(lightsModel).ToNot(BeNil())
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(38))
				lightsModel.LedOffMs = core.StringPtr("testString")
				Expect(lightsModel.LedArgb).To(Equal(core.StringPtr("testString")))
				Expect(lightsModel.LedOnMs).To(Equal(core.Int64Ptr(int64(38))))
				Expect(lightsModel.LedOffMs).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Style model
				styleModel := new(pushservicev1.Style)
				Expect(styleModel).ToNot(BeNil())
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}
				Expect(styleModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.Lines).To(Equal([]string{"testString"}))

				// Construct an instance of the Gcm model
				gcmModel := new(pushservicev1.Gcm)
				Expect(gcmModel).ToNot(BeNil())
				gcmModel.CollapseKey = core.StringPtr("testString")
				gcmModel.InteractiveCategory = core.StringPtr("testString")
				gcmModel.Icon = core.StringPtr("testString")
				gcmModel.DelayWhileIdle = core.BoolPtr(true)
				gcmModel.Sync = core.BoolPtr(true)
				gcmModel.Visibility = core.StringPtr("testString")
				gcmModel.Redact = core.StringPtr("testString")
				gcmModel.ChannelID = core.StringPtr("testString")
				gcmModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				gcmModel.Priority = core.StringPtr("testString")
				gcmModel.Sound = core.StringPtr("testString")
				gcmModel.TimeToLive = core.Int64Ptr(int64(38))
				gcmModel.Lights = lightsModel
				gcmModel.AndroidTitle = core.StringPtr("testString")
				gcmModel.GroupID = core.StringPtr("testString")
				gcmModel.Style = styleModel
				gcmModel.Type = core.StringPtr("DEFAULT")
				Expect(gcmModel.CollapseKey).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.InteractiveCategory).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.Icon).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.DelayWhileIdle).To(Equal(core.BoolPtr(true)))
				Expect(gcmModel.Sync).To(Equal(core.BoolPtr(true)))
				Expect(gcmModel.Visibility).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.Redact).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.ChannelID).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.Payload).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(gcmModel.Priority).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.Sound).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.TimeToLive).To(Equal(core.Int64Ptr(int64(38))))
				Expect(gcmModel.Lights).To(Equal(lightsModel))
				Expect(gcmModel.AndroidTitle).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.Style).To(Equal(styleModel))
				Expect(gcmModel.Type).To(Equal(core.StringPtr("DEFAULT")))

				// Construct an instance of the FirefoxWeb model
				firefoxWebModel := new(pushservicev1.FirefoxWeb)
				Expect(firefoxWebModel).ToNot(BeNil())
				firefoxWebModel.Title = core.StringPtr("testString")
				firefoxWebModel.IconURL = core.StringPtr("testString")
				firefoxWebModel.TimeToLive = core.Int64Ptr(int64(38))
				firefoxWebModel.Payload = core.StringPtr("testString")
				Expect(firefoxWebModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(firefoxWebModel.IconURL).To(Equal(core.StringPtr("testString")))
				Expect(firefoxWebModel.TimeToLive).To(Equal(core.Int64Ptr(int64(38))))
				Expect(firefoxWebModel.Payload).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ChromeWeb model
				chromeWebModel := new(pushservicev1.ChromeWeb)
				Expect(chromeWebModel).ToNot(BeNil())
				chromeWebModel.Title = core.StringPtr("testString")
				chromeWebModel.IconURL = core.StringPtr("testString")
				chromeWebModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeWebModel.Payload = core.StringPtr("testString")
				Expect(chromeWebModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(chromeWebModel.IconURL).To(Equal(core.StringPtr("testString")))
				Expect(chromeWebModel.TimeToLive).To(Equal(core.Int64Ptr(int64(38))))
				Expect(chromeWebModel.Payload).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SafariWeb model
				safariWebModel := new(pushservicev1.SafariWeb)
				Expect(safariWebModel).ToNot(BeNil())
				safariWebModel.Title = core.StringPtr("testString")
				safariWebModel.UrlArgs = []string{"testString"}
				safariWebModel.Action = core.StringPtr("testString")
				Expect(safariWebModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(safariWebModel.UrlArgs).To(Equal([]string{"testString"}))
				Expect(safariWebModel.Action).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ChromeAppExt model
				chromeAppExtModel := new(pushservicev1.ChromeAppExt)
				Expect(chromeAppExtModel).ToNot(BeNil())
				chromeAppExtModel.CollapseKey = core.StringPtr("testString")
				chromeAppExtModel.DelayWhileIdle = core.BoolPtr(true)
				chromeAppExtModel.Title = core.StringPtr("testString")
				chromeAppExtModel.IconURL = core.StringPtr("testString")
				chromeAppExtModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeAppExtModel.Payload = core.StringPtr("testString")
				Expect(chromeAppExtModel.CollapseKey).To(Equal(core.StringPtr("testString")))
				Expect(chromeAppExtModel.DelayWhileIdle).To(Equal(core.BoolPtr(true)))
				Expect(chromeAppExtModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(chromeAppExtModel.IconURL).To(Equal(core.StringPtr("testString")))
				Expect(chromeAppExtModel.TimeToLive).To(Equal(core.Int64Ptr(int64(38))))
				Expect(chromeAppExtModel.Payload).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Settings model
				settingsModel := new(pushservicev1.Settings)
				Expect(settingsModel).ToNot(BeNil())
				settingsModel.Apns = apnsModel
				settingsModel.Gcm = gcmModel
				settingsModel.FirefoxWeb = firefoxWebModel
				settingsModel.ChromeWeb = chromeWebModel
				settingsModel.SafariWeb = safariWebModel
				settingsModel.ChromeAppExt = chromeAppExtModel
				Expect(settingsModel.Apns).To(Equal(apnsModel))
				Expect(settingsModel.Gcm).To(Equal(gcmModel))
				Expect(settingsModel.FirefoxWeb).To(Equal(firefoxWebModel))
				Expect(settingsModel.ChromeWeb).To(Equal(chromeWebModel))
				Expect(settingsModel.SafariWeb).To(Equal(safariWebModel))
				Expect(settingsModel.ChromeAppExt).To(Equal(chromeAppExtModel))

				// Construct an instance of the Target model
				targetModel := new(pushservicev1.Target)
				Expect(targetModel).ToNot(BeNil())
				targetModel.DeviceIds = []string{"testString"}
				targetModel.UserIds = []string{"testString"}
				targetModel.Platforms = []string{"testString"}
				targetModel.TagNames = []string{"testString"}
				Expect(targetModel.DeviceIds).To(Equal([]string{"testString"}))
				Expect(targetModel.UserIds).To(Equal([]string{"testString"}))
				Expect(targetModel.Platforms).To(Equal([]string{"testString"}))
				Expect(targetModel.TagNames).To(Equal([]string{"testString"}))

				// Construct an instance of the SendMessageOptions model
				applicationID := "testString"
				var sendMessageOptionsMessage *pushservicev1.Message = nil
				sendMessageOptionsModel := pushServiceService.NewSendMessageOptions(applicationID, sendMessageOptionsMessage)
				sendMessageOptionsModel.SetApplicationID("testString")
				sendMessageOptionsModel.SetMessage(messageModel)
				sendMessageOptionsModel.SetSettings(settingsModel)
				sendMessageOptionsModel.SetValidate(true)
				sendMessageOptionsModel.SetTarget(targetModel)
				sendMessageOptionsModel.SetAcceptLanguage("testString")
				sendMessageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(sendMessageOptionsModel).ToNot(BeNil())
				Expect(sendMessageOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(sendMessageOptionsModel.Message).To(Equal(messageModel))
				Expect(sendMessageOptionsModel.Settings).To(Equal(settingsModel))
				Expect(sendMessageOptionsModel.Validate).To(Equal(core.BoolPtr(true)))
				Expect(sendMessageOptionsModel.Target).To(Equal(targetModel))
				Expect(sendMessageOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(sendMessageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSendMessagesInBulkOptions successfully`, func() {
				// Construct an instance of the Message model
				messageModel := new(pushservicev1.Message)
				Expect(messageModel).ToNot(BeNil())
				messageModel.Alert = core.StringPtr("testString")
				messageModel.URL = core.StringPtr("testString")
				Expect(messageModel.Alert).To(Equal(core.StringPtr("testString")))
				Expect(messageModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Apns model
				apnsModel := new(pushservicev1.Apns)
				Expect(apnsModel).ToNot(BeNil())
				apnsModel.Badge = core.Int64Ptr(int64(38))
				apnsModel.InteractiveCategory = core.StringPtr("testString")
				apnsModel.Category = core.StringPtr("testString")
				apnsModel.IosActionKey = core.StringPtr("testString")
				apnsModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				apnsModel.Sound = core.StringPtr("testString")
				apnsModel.TitleLocKey = core.StringPtr("testString")
				apnsModel.LocKey = core.StringPtr("testString")
				apnsModel.LaunchImage = core.StringPtr("testString")
				apnsModel.TitleLocArgs = []string{"testString"}
				apnsModel.LocArgs = []string{"testString"}
				apnsModel.Title = core.StringPtr("testString")
				apnsModel.Subtitle = core.StringPtr("testString")
				apnsModel.AttachmentURL = core.StringPtr("testString")
				apnsModel.Type = core.StringPtr("DEFAULT")
				apnsModel.ApnsCollapseID = core.StringPtr("testString")
				apnsModel.ApnsThreadID = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArg = core.StringPtr("testString")
				apnsModel.ApnsGroupSummaryArgCount = core.Int64Ptr(int64(38))
				Expect(apnsModel.Badge).To(Equal(core.Int64Ptr(int64(38))))
				Expect(apnsModel.InteractiveCategory).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.IosActionKey).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.Payload).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(apnsModel.Sound).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.TitleLocKey).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.LocKey).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.LaunchImage).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.TitleLocArgs).To(Equal([]string{"testString"}))
				Expect(apnsModel.LocArgs).To(Equal([]string{"testString"}))
				Expect(apnsModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.Subtitle).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.AttachmentURL).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.Type).To(Equal(core.StringPtr("DEFAULT")))
				Expect(apnsModel.ApnsCollapseID).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.ApnsThreadID).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.ApnsGroupSummaryArg).To(Equal(core.StringPtr("testString")))
				Expect(apnsModel.ApnsGroupSummaryArgCount).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the Lights model
				lightsModel := new(pushservicev1.Lights)
				Expect(lightsModel).ToNot(BeNil())
				lightsModel.LedArgb = core.StringPtr("testString")
				lightsModel.LedOnMs = core.Int64Ptr(int64(38))
				lightsModel.LedOffMs = core.StringPtr("testString")
				Expect(lightsModel.LedArgb).To(Equal(core.StringPtr("testString")))
				Expect(lightsModel.LedOnMs).To(Equal(core.Int64Ptr(int64(38))))
				Expect(lightsModel.LedOffMs).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Style model
				styleModel := new(pushservicev1.Style)
				Expect(styleModel).ToNot(BeNil())
				styleModel.Type = core.StringPtr("testString")
				styleModel.Title = core.StringPtr("testString")
				styleModel.URL = core.StringPtr("testString")
				styleModel.Text = core.StringPtr("testString")
				styleModel.Lines = []string{"testString"}
				Expect(styleModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(styleModel.Lines).To(Equal([]string{"testString"}))

				// Construct an instance of the Gcm model
				gcmModel := new(pushservicev1.Gcm)
				Expect(gcmModel).ToNot(BeNil())
				gcmModel.CollapseKey = core.StringPtr("testString")
				gcmModel.InteractiveCategory = core.StringPtr("testString")
				gcmModel.Icon = core.StringPtr("testString")
				gcmModel.DelayWhileIdle = core.BoolPtr(true)
				gcmModel.Sync = core.BoolPtr(true)
				gcmModel.Visibility = core.StringPtr("testString")
				gcmModel.Redact = core.StringPtr("testString")
				gcmModel.ChannelID = core.StringPtr("testString")
				gcmModel.Payload = map[string]interface{}{"anyKey": "anyValue"}
				gcmModel.Priority = core.StringPtr("testString")
				gcmModel.Sound = core.StringPtr("testString")
				gcmModel.TimeToLive = core.Int64Ptr(int64(38))
				gcmModel.Lights = lightsModel
				gcmModel.AndroidTitle = core.StringPtr("testString")
				gcmModel.GroupID = core.StringPtr("testString")
				gcmModel.Style = styleModel
				gcmModel.Type = core.StringPtr("DEFAULT")
				Expect(gcmModel.CollapseKey).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.InteractiveCategory).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.Icon).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.DelayWhileIdle).To(Equal(core.BoolPtr(true)))
				Expect(gcmModel.Sync).To(Equal(core.BoolPtr(true)))
				Expect(gcmModel.Visibility).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.Redact).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.ChannelID).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.Payload).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(gcmModel.Priority).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.Sound).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.TimeToLive).To(Equal(core.Int64Ptr(int64(38))))
				Expect(gcmModel.Lights).To(Equal(lightsModel))
				Expect(gcmModel.AndroidTitle).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.GroupID).To(Equal(core.StringPtr("testString")))
				Expect(gcmModel.Style).To(Equal(styleModel))
				Expect(gcmModel.Type).To(Equal(core.StringPtr("DEFAULT")))

				// Construct an instance of the FirefoxWeb model
				firefoxWebModel := new(pushservicev1.FirefoxWeb)
				Expect(firefoxWebModel).ToNot(BeNil())
				firefoxWebModel.Title = core.StringPtr("testString")
				firefoxWebModel.IconURL = core.StringPtr("testString")
				firefoxWebModel.TimeToLive = core.Int64Ptr(int64(38))
				firefoxWebModel.Payload = core.StringPtr("testString")
				Expect(firefoxWebModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(firefoxWebModel.IconURL).To(Equal(core.StringPtr("testString")))
				Expect(firefoxWebModel.TimeToLive).To(Equal(core.Int64Ptr(int64(38))))
				Expect(firefoxWebModel.Payload).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ChromeWeb model
				chromeWebModel := new(pushservicev1.ChromeWeb)
				Expect(chromeWebModel).ToNot(BeNil())
				chromeWebModel.Title = core.StringPtr("testString")
				chromeWebModel.IconURL = core.StringPtr("testString")
				chromeWebModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeWebModel.Payload = core.StringPtr("testString")
				Expect(chromeWebModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(chromeWebModel.IconURL).To(Equal(core.StringPtr("testString")))
				Expect(chromeWebModel.TimeToLive).To(Equal(core.Int64Ptr(int64(38))))
				Expect(chromeWebModel.Payload).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SafariWeb model
				safariWebModel := new(pushservicev1.SafariWeb)
				Expect(safariWebModel).ToNot(BeNil())
				safariWebModel.Title = core.StringPtr("testString")
				safariWebModel.UrlArgs = []string{"testString"}
				safariWebModel.Action = core.StringPtr("testString")
				Expect(safariWebModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(safariWebModel.UrlArgs).To(Equal([]string{"testString"}))
				Expect(safariWebModel.Action).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ChromeAppExt model
				chromeAppExtModel := new(pushservicev1.ChromeAppExt)
				Expect(chromeAppExtModel).ToNot(BeNil())
				chromeAppExtModel.CollapseKey = core.StringPtr("testString")
				chromeAppExtModel.DelayWhileIdle = core.BoolPtr(true)
				chromeAppExtModel.Title = core.StringPtr("testString")
				chromeAppExtModel.IconURL = core.StringPtr("testString")
				chromeAppExtModel.TimeToLive = core.Int64Ptr(int64(38))
				chromeAppExtModel.Payload = core.StringPtr("testString")
				Expect(chromeAppExtModel.CollapseKey).To(Equal(core.StringPtr("testString")))
				Expect(chromeAppExtModel.DelayWhileIdle).To(Equal(core.BoolPtr(true)))
				Expect(chromeAppExtModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(chromeAppExtModel.IconURL).To(Equal(core.StringPtr("testString")))
				Expect(chromeAppExtModel.TimeToLive).To(Equal(core.Int64Ptr(int64(38))))
				Expect(chromeAppExtModel.Payload).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Settings model
				settingsModel := new(pushservicev1.Settings)
				Expect(settingsModel).ToNot(BeNil())
				settingsModel.Apns = apnsModel
				settingsModel.Gcm = gcmModel
				settingsModel.FirefoxWeb = firefoxWebModel
				settingsModel.ChromeWeb = chromeWebModel
				settingsModel.SafariWeb = safariWebModel
				settingsModel.ChromeAppExt = chromeAppExtModel
				Expect(settingsModel.Apns).To(Equal(apnsModel))
				Expect(settingsModel.Gcm).To(Equal(gcmModel))
				Expect(settingsModel.FirefoxWeb).To(Equal(firefoxWebModel))
				Expect(settingsModel.ChromeWeb).To(Equal(chromeWebModel))
				Expect(settingsModel.SafariWeb).To(Equal(safariWebModel))
				Expect(settingsModel.ChromeAppExt).To(Equal(chromeAppExtModel))

				// Construct an instance of the Target model
				targetModel := new(pushservicev1.Target)
				Expect(targetModel).ToNot(BeNil())
				targetModel.DeviceIds = []string{"testString"}
				targetModel.UserIds = []string{"testString"}
				targetModel.Platforms = []string{"testString"}
				targetModel.TagNames = []string{"testString"}
				Expect(targetModel.DeviceIds).To(Equal([]string{"testString"}))
				Expect(targetModel.UserIds).To(Equal([]string{"testString"}))
				Expect(targetModel.Platforms).To(Equal([]string{"testString"}))
				Expect(targetModel.TagNames).To(Equal([]string{"testString"}))

				// Construct an instance of the SendMessageBody model
				sendMessageBodyModel := new(pushservicev1.SendMessageBody)
				Expect(sendMessageBodyModel).ToNot(BeNil())
				sendMessageBodyModel.Message = messageModel
				sendMessageBodyModel.Settings = settingsModel
				sendMessageBodyModel.Validate = core.BoolPtr(true)
				sendMessageBodyModel.Target = targetModel
				Expect(sendMessageBodyModel.Message).To(Equal(messageModel))
				Expect(sendMessageBodyModel.Settings).To(Equal(settingsModel))
				Expect(sendMessageBodyModel.Validate).To(Equal(core.BoolPtr(true)))
				Expect(sendMessageBodyModel.Target).To(Equal(targetModel))

				// Construct an instance of the SendMessagesInBulkOptions model
				applicationID := "testString"
				body := []pushservicev1.SendMessageBody{}
				sendMessagesInBulkOptionsModel := pushServiceService.NewSendMessagesInBulkOptions(applicationID, body)
				sendMessagesInBulkOptionsModel.SetApplicationID("testString")
				sendMessagesInBulkOptionsModel.SetBody([]pushservicev1.SendMessageBody{*sendMessageBodyModel})
				sendMessagesInBulkOptionsModel.SetAcceptLanguage("testString")
				sendMessagesInBulkOptionsModel.SetAppSecret("testString")
				sendMessagesInBulkOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(sendMessagesInBulkOptionsModel).ToNot(BeNil())
				Expect(sendMessagesInBulkOptionsModel.ApplicationID).To(Equal(core.StringPtr("testString")))
				Expect(sendMessagesInBulkOptionsModel.Body).To(Equal([]pushservicev1.SendMessageBody{*sendMessageBodyModel}))
				Expect(sendMessagesInBulkOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(sendMessagesInBulkOptionsModel.AppSecret).To(Equal(core.StringPtr("testString")))
				Expect(sendMessagesInBulkOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSendMessageBody successfully`, func() {
				var message *pushservicev1.Message = nil
				_, err := pushServiceService.NewSendMessageBody(message)
				Expect(err).ToNot(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
