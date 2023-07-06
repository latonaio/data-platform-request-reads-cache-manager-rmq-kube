package services

import (
	"bytes"
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	"data-platform-request-reads-cache-manager-rmq-kube/config"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	logger "github.com/latonaio/golang-logging-library-for-data-platform"
	"golang.org/x/xerrors"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	POST = "POST"
	GET  = "GET"
)

type ResponseData struct {
	StatusCode int    `json:"statusCode"`
	Name       string `json:"name"`
	Message    string `json:"message"`
	Data       struct {
		RuntimeSessionID *string `json:"runtimeSessionId"`
	} `json:"data"`
}

func UserRequestParams(
	controller *beego.Controller,
) *apiInputReader.Request {
	businessPartner, _ := controller.GetInt("businessPartner")
	language := controller.GetString("language")
	userId := controller.GetString("userId")

	return &apiInputReader.Request{
		Language:        &language,
		BusinessPartner: &businessPartner,
		UserID:          &userId,
	}
}

func Request(
	aPIServiceName string,
	aPIType string,
	body io.ReadCloser,
	controller *beego.Controller,
) []byte {
	conf := config.NewConf()
	nestjsURL := conf.REQUEST.RequestURL()

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

	if response.StatusCode >= 400 && response.StatusCode < 500 {
		HandleError(
			controller,
			responseBody,
			&response.StatusCode,
		)
		return nil
	}

	return responseBody
}

func HandleError(
	controller *beego.Controller,
	message interface{},
	statusCode *int,
) {
	l := logger.NewLogger()
	ctx := controller.Ctx

	responseData := ResponseData{}

	if statusCode == nil {
		ctx.Output.SetStatus(500)
	} else {
		ctx.Output.SetStatus(*statusCode)
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
}

func Respond(
	controller *beego.Controller,
	data interface{},
) {
	controller.Data["json"] = data
	controller.ServeJSON()
}
