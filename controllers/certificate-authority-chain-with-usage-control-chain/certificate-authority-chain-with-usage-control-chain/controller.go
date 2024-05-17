package controllersCertificateAuthorityChainWithUsageControlChain

import (
	apiInputReader "data-platform-request-reads-cache-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/business-partner/business-partner"
	apiModuleRuntimesRequestsCertificateAuthorityChainWithUsageControlChain "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-requests/certificate-authority-chain-with-usage-control-chain/certificate-authority-chain-with-usage-control-chain"
	apiModuleRuntimesResponsesBusinessPartner "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/business-partner"
	apiModuleRuntimesResponsesCertificateAuthorityChainWithUsageControlChain "data-platform-request-reads-cache-manager-rmq-kube/api-module-runtimes-responses/certificate-authority-chain-with-usage-control-chain"
	apiOutputFormatter "data-platform-request-reads-cache-manager-rmq-kube/api-output-formatter"
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type CertificateAuthorityChainWithUsageControlChainController struct {
	beego.Controller
	RedisCache   *cache.Cache
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

//const (
//	buyer  = "buyer"
//	seller = "seller"
//)

func (controller *CertificateAuthorityChainWithUsageControlChainController) Get() {
	controller.UserInfo = services.UserRequestParams(&controller.Controller)
	redisKeyCategory1 := "certificate-authority-chain-with-usage-control-chain"
	certificateAuthorityChain := controller.GetString("certificateAuthorityChain")
	usageControlChain := controller.GetString("usageControlChain")
	certificateObject := controller.GetString("certificateObject")
	certificateObjectLabel := controller.GetString("certificateObjectLabel")
	//	userType := controller.GetString(":userType")
	//	pBuyer, _ := controller.GetInt("buyer")
	//	pSeller, _ := controller.GetInt("seller")

	CertificateAuthorityChainWithUsageControlChain := apiInputReader.CertificateAuthorityChainWithUsageControlChainGlobal{}

	isMarkedForDeletion := false

	//	docType := "QRCODE"

	CertificateAuthorityChainWithUsageControlChain = apiInputReader.CertificateAuthorityChainWithUsageControlChainGlobal{
		CertificateAuthorityChainWithUsageControlChain: &apiInputReader.CertificateAuthorityChainWithUsageControlChain{
			CertificateAuthorityChain: certificateAuthorityChain,
			CertificateObject:         certificateObject,
			CertificateObjectLabel:    certificateObjectLabel,
			UsageControlChain:         usageControlChain,
			IsMarkedForDeletion:       &isMarkedForDeletion,
		},
	}

	controller.RedisKey = controller.RedisCache.CreateKey(
		&controller.Controller,
		[]string{
			redisKeyCategory1,
			certificateObject,
		},
	)

	cacheData, _ := controller.RedisCache.ConfirmCashDataExisting(controller.RedisKey)

	if cacheData != nil {
		var responseData apiOutputFormatter.CertificateAuthorityChainWithUsageControlChainGlobal

		err := json.Unmarshal(cacheData, &responseData)

		if err != nil {
			services.HandleError(
				&controller.Controller,
				err,
				nil,
			)
		}

		services.Respond(
			&controller.Controller,
			&responseData,
		)
	}

	if cacheData != nil {
		go func() {
			controller.request(CertificateAuthorityChainWithUsageControlChain)
		}()
	} else {
		controller.request(CertificateAuthorityChainWithUsageControlChain)
	}
}

func (
	controller *CertificateAuthorityChainWithUsageControlChainController,
) createCertificateAuthorityChainWithUsageControlChainRequestCertificateAuthorityChain(
	requestPram *apiInputReader.Request,
	input apiInputReader.CertificateAuthorityChainWithUsageControlChainGlobal,
) *apiModuleRuntimesResponsesCertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChainWithUsageControlChainRes {
	responseJsonData := apiModuleRuntimesResponsesCertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChainWithUsageControlChainRes{}
	responseBody := apiModuleRuntimesRequestsCertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChainReads(
		requestPram,
		input,
		&controller.Controller,
		"CertificateAuthorityChain",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("CertificateAuthorityChainReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *CertificateAuthorityChainWithUsageControlChainController,
) createCertificateAuthorityChainWithUsageControlChainRequestUsageControlChain(
	requestPram *apiInputReader.Request,
	input apiInputReader.CertificateAuthorityChainWithUsageControlChainGlobal,
) *apiModuleRuntimesResponsesCertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChainWithUsageControlChainRes {
	responseJsonData := apiModuleRuntimesResponsesCertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChainWithUsageControlChainRes{}
	responseBody := apiModuleRuntimesRequestsCertificateAuthorityChainWithUsageControlChain.UsageControlChainReads(
		requestPram,
		input,
		&controller.Controller,
		"UsageControlChain",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("UsageControlChainReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *CertificateAuthorityChainWithUsageControlChainController,
) createBusinessPartnerRequest(
	requestPram *apiInputReader.Request,
	certificateAuthorityChainRes *apiModuleRuntimesResponsesCertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChainWithUsageControlChainRes,
) *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes {
	input := make([]apiModuleRuntimesRequestsBusinessPartner.General, len(*certificateAuthorityChainRes.Message.CertificateAuthorityChain))

	for _, v := range *certificateAuthorityChainRes.Message.CertificateAuthorityChain {
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DataIssuer,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DataAuthorizer,
		})
		input = append(input, apiModuleRuntimesRequestsBusinessPartner.General{
			BusinessPartner: v.DataDistributor,
		})
	}

	responseJsonData := apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerReadsGeneralsByBusinessPartners(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("BusinessPartnerGeneralReads Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *CertificateAuthorityChainWithUsageControlChainController,
) request(
	input apiInputReader.CertificateAuthorityChainWithUsageControlChainGlobal,
) {
	defer services.Recover(controller.CustomLogger, &controller.Controller)

	certificateAuthorityChainRes := *controller.createCertificateAuthorityChainWithUsageControlChainRequestCertificateAuthorityChain(
		controller.UserInfo,
		input,
	)

	usageControlChainRes := *controller.createCertificateAuthorityChainWithUsageControlChainRequestUsageControlChain(
		controller.UserInfo,
		input,
	)

	businessPartnerRes := *controller.createBusinessPartnerRequest(
		controller.UserInfo,
		&certificateAuthorityChainRes,
	)

	controller.fin(
		input,
		&certificateAuthorityChainRes,
		&usageControlChainRes,
		&businessPartnerRes,
	)
}

func (
	controller *CertificateAuthorityChainWithUsageControlChainController,
) fin(
	input apiInputReader.CertificateAuthorityChainWithUsageControlChainGlobal,
	certificateAuthorityChainRes *apiModuleRuntimesResponsesCertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChainWithUsageControlChainRes,
	usageControlChainRes *apiModuleRuntimesResponsesCertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChainWithUsageControlChainRes,
	businessPartnerRes *apiModuleRuntimesResponsesBusinessPartner.BusinessPartnerRes,
) {
	businessPartnerMapper := services.BusinessPartnerNameMapper(
		businessPartnerRes,
	)

	data := apiOutputFormatter.CertificateAuthorityChainWithUsageControlChainGlobal{}

	for _, v := range *certificateAuthorityChainRes.Message.CertificateAuthorityChain {
		data.CertificateAuthorityChain = append(data.CertificateAuthorityChain,
			apiOutputFormatter.CertificateAuthorityChain{
				CertificateAuthorityChain: input.CertificateAuthorityChainWithUsageControlChain.CertificateAuthorityChain,
				CertificateObject:         input.CertificateAuthorityChainWithUsageControlChain.CertificateObject,
				CertificateObjectLabel:    input.CertificateAuthorityChainWithUsageControlChain.CertificateObjectLabel,
				DataIssuer:                v.DataIssuer,
				DataIssuerName:            businessPartnerMapper[v.DataIssuer].BusinessPartnerName,
				DataAuthorizer:            v.DataAuthorizer,
				DataAuthorizerName:        businessPartnerMapper[v.DataAuthorizer].BusinessPartnerName,
				DataDistributor:           v.DataDistributor,
				DataDistributorName:       businessPartnerMapper[v.DataDistributor].BusinessPartnerName,
			},
		)
	}

	for _, v := range *usageControlChainRes.Message.UsageControlChain {
		data.UsageControlChain = append(data.UsageControlChain,
			apiOutputFormatter.UsageControlChain{
				UsageControlChain:              v.UsageControlChain,
				UsageControlLess:               v.UsageControlLess,
				Perpetual:                      v.Perpetual,
				Rental:                         v.Rental,
				Duration:                       v.Duration,
				DurationUnit:                   v.DurationUnit,
				ValidityStartDate:              v.ValidityStartDate,
				ValidityStartTime:              v.ValidityStartTime,
				ValidityEndDate:                v.ValidityEndDate,
				ValidityEndTime:                v.ValidityEndTime,
				DeleteAfterValidityEnd:         v.DeleteAfterValidityEnd,
				ServiceLabelRestriction:        v.ServiceLabelRestriction,
				ApplicationRestriction:         v.ApplicationRestriction,
				PurposeRestriction:             v.PurposeRestriction,
				BusinessPartnerRoleRestriction: v.BusinessPartnerRoleRestriction,
				DataStateRestriction:           v.DataStateRestriction,
				NumberOfUsageRestriction:       v.NumberOfUsageRestriction,
				NumberOfActualUsage:            v.NumberOfActualUsage,
				IPAddressRestriction:           v.IPAddressRestriction,
				MACAddressRestriction:          v.MACAddressRestriction,
				ModifyIsAllowed:                v.ModifyIsAllowed,
				LocalLoggingIsAllowed:          v.LocalLoggingIsAllowed,
				RemoteNotificationIsAllowed:    v.RemoteNotificationIsAllowed,
				DistributeOnlyIfEncrypted:      v.DistributeOnlyIfEncrypted,
				AttachPolicyWhenDistribute:     v.AttachPolicyWhenDistribute,
				PostalCode:                     v.PostalCode,
				LocalSubRegion:                 v.LocalSubRegion,
				LocalRegion:                    v.LocalRegion,
				Country:                        v.Country,
				GlobalRegion:                   v.GlobalRegion,
				TimeZone:                       v.TimeZone,
				CreationDate:                   v.CreationDate,
				CreationTime:                   v.CreationTime,

				//				Images: apiOutputFormatter.Images{
				//					Product: img,
				//					QRCode:  qrcode,
				//DocumentImageCertificateAuthorityChainWithUsageControlChain: documentImage,
			},
		)
	}

	err := controller.RedisCache.SetCache(
		controller.RedisKey,
		data,
	)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
	}

	controller.Data["json"] = &data
	controller.ServeJSON()
}
