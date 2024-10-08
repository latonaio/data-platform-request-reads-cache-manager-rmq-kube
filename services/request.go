package services

import (
	"bytes"
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/config"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/google/uuid"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	POST = "POST"
	GET  = "GET"
)

type RequestWrapperController struct {
	Controller   *beego.Controller
	CustomLogger *logger.Logger
}

type ResponseData struct {
	StatusCode int    `json:"statusCode"`
	Name       string `json:"name"`
	Message    string `json:"message"`
	Data       struct {
		RuntimeSessionID *string `json:"runtimeSessionId"`
	} `json:"data"`
}

type AuthenticatorResponseData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func UserRequestParams(
	requestWrapperController RequestWrapperController,
) *apiInputReader.Request {
	businessPartner, _ := requestWrapperController.Controller.GetInt("businessPartner")
	language := requestWrapperController.Controller.GetString("language")
	userId := requestWrapperController.Controller.GetString("userId")

	runtimeSessionId := uuid.New().String()
	runtimeSessionId = strings.ReplaceAll(runtimeSessionId, "-", "")

	if requestWrapperController.CustomLogger != nil {
		requestWrapperController.CustomLogger.Info(
			"RuntimeSessionID: %v",
			runtimeSessionId,
		)
	}

	return &apiInputReader.Request{
		Language:         &language,
		BusinessPartner:  &businessPartner,
		UserID:           &userId,
		RuntimeSessionID: &runtimeSessionId,
	}
}

func Request(
	aPIServiceName string,
	aPIType string,
	body io.ReadCloser,
	controller *beego.Controller,
	requestPram *apiInputReader.Request,
) []byte {
	conf := config.NewConf()
	nestjsURL := conf.REQUEST.RequestURL()
	jwtToken := controller.Ctx.Input.Header("Authorization")

	method := POST
	requestUrl := fmt.Sprintf("%s/%s/%s", nestjsURL, aPIServiceName, aPIType)

	byteBody, err := ioutil.ReadAll(body)
	if err != nil {
		HandleError(
			controller,
			err,
			nil,
		)
	}

	req, err := http.NewRequest(
		method, requestUrl, ioutil.NopCloser(bytes.NewReader(byteBody)),
	)

	if err != nil {
		HandleError(
			controller,
			err,
			nil,
		)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if requestPram != nil && requestPram.RuntimeSessionID != nil {
		req.Header.Add("global-runtime-session-id", *requestPram.RuntimeSessionID)
	}

	req.Header.Set("Authorization", jwtToken)

	client := &http.Client{}

	response, err := client.Do(req)

	responseBody, err := ioutil.ReadAll(response.Body)

	err = response.Body.Close()
	if err != nil {
		HandleError(
			controller,
			err,
			nil,
		)
		return nil
	}

	//if response.StatusCode != 200 && response.StatusCode != 201 {
	//	HandleError(
	//		controller,
	//		responseBody,
	//		&response.StatusCode,
	//	)
	//	return nil
	//}

	return responseBody
}
func VerifyJwtToken(
	jwtToken string,
	ctx *context.Context,
) error {
	conf := config.NewConf()

	authenticatorURL := conf.AUTHENTICATOR.RequestURL()
	requestBody := []byte(`{}`)
	method := POST

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s/%s", authenticatorURL, "token/verify"),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return xerrors.Errorf("VerifyJwtToken newRequest error: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Set("Authorization", jwtToken)
	//req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbF9hZGRyZXNzIjoiMTAxQGdtYWlsLmNvbSIsImV4cCI6IjE2ODg3MzI3MzkifQ.uva4qc9KzA05z9bY5viS71_QYXC2DjrJByFRsh1Mf8IezeHbQs6KZ2W0PPUA0I0Z7jc-RvHsElgPX1oizjrteDaFpw9p62FaLGAyjPMADZwLd5wJLME9fwqARICwv2CTh-h1DG377ki6iSs0udSNAssVSzGjTf5kSIG7z-b25ZQM6EkydGSJx_MDaIm5KpBBxrw9XridKxowZewPaHrA4MkgUGWFicXJ3mEnGWxDTVS3oIoWRjFwjY31hfSult1eULp3ALTze0N8wkbcdBHVgqAq0fudGYhu5Y4YaTrYSKRJsiOmhVdY_2sSAbIt7ZlIWGRmU1Uge0e0ZXKFVO2v0xE4XAi6qKkS3a8bkBxI5fpVhbX8fHEg1uSIi4oIvYoGYEMyX0B4ci1LxuZqcizJaMtRuU0f05HGfQhd9iL7H2TqAJOkVKKgSRkr9jsMM5vD92-a9cYNsklrbiVeH2AGF-ruYkYqdTkhH1QeQyJinhaKNhene_9Dx2fPUpL0rFI2BmvZs6JxXchv3ldEKUqjfBqaPtkcgFxb92mqvTACPpQuQ4ESndX8c38FMIbTaWYgNoEuBOuCFN6FL9i-JRmRhju5zlXxDye4462jsXpzboUWz9KimnfZoiHVDKvWWds9yCz64-g8iLNMagVNoaoW7STQy8zMntFkq8TEuQZ26Is")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return xerrors.Errorf("VerifyJwtToken error: %w", err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return xerrors.Errorf("VerifyJwtToken error: %w", err)
	}

	err = response.Body.Close()
	if err != nil {
		return xerrors.Errorf("VerifyJwtToken error: %w", err)
	}

	authenticatorResponseData := AuthenticatorResponseData{}

	err = json.Unmarshal(responseBody, &authenticatorResponseData)
	if err != nil {
		return xerrors.Errorf("VerifyJwtToken error: %w", err)
	}

	jsonData, err := json.Marshal(authenticatorResponseData)
	if err != nil {
		return xerrors.Errorf("VerifyJwtToken marshal error: %w", err)
	}

	if response.StatusCode != 200 {
		ctx.Output.SetStatus(response.StatusCode)
		ctx.Output.Header("Content-Type", "application/json")
		_, err := ctx.ResponseWriter.Write(jsonData)
		if err != nil {
			return xerrors.Errorf("VerifyJwtToken write error: %w", err)
		}
		return xerrors.Errorf("VerifyJwtToken error: %w", err)
	}

	return nil
}

func HandleError(
	controller *beego.Controller,
	message interface{},
	statusCode *int,
) {
	l := logger.NewLogger()
	ctx := controller.Ctx

	responseData := ResponseData{
		Message: fmt.Sprintf("%v", message),
	}

	if statusCode == nil {
		ctx.Output.SetStatus(500)
		responseData.StatusCode = 500
	} else {
		ctx.Output.SetStatus(*statusCode)
		responseData.StatusCode = *statusCode
	}

	if msg, ok := message.([]byte); ok {
		err := json.Unmarshal(msg, &responseData)

		controller.Data["json"] = responseData
		controller.ServeJSON()

		if err != nil {
			l.Error(xerrors.Errorf("HandleError error: %w", err))
		}
	}

	if errMsg, ok := message.(error); ok {
		responseData = ResponseData{
			StatusCode: func() int {
				if statusCode != nil {
					return *statusCode
				}
				return 500
			}(),
			// todo エラーの種類をまとめておくこと
			Name:    "InternalServerError",
			Message: errMsg.Error(),
			Data: struct {
				RuntimeSessionID *string `json:"runtimeSessionId"`
			}{},
		}
	}

	controller.Data["json"] = responseData
	controller.ServeJSON()

	if statusCode != nil {
		controller.Abort(fmt.Sprintf("%d", &statusCode))
	} else {
		controller.Abort("500")
	}
}

func Respond(
	controller *beego.Controller,
	data interface{},
) {
	controller.Data["json"] = data
	controller.ServeJSON()
}
