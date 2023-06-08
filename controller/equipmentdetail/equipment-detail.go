package equipmentdetail

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/equipmentdetail"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"
	"data-platform-api-request-reads-cache-manager-rmq-kube/cache"
	rmqsessioncontroller "data-platform-api-request-reads-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type EquipmentDetailCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewEquipmentDetailCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *EquipmentDetailCtrl {
	return &EquipmentDetailCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *EquipmentDetailCtrl) EquipmentDetail(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractEquipmentDetailParam(msg)
	reqKey, err := getRequestKey(msg.Data())
	if err != nil {
		return xerrors.Errorf("reqKey error: %w", err)
	}
	sID, err := getSessionID(msg.Data())
	if err != nil {
		return xerrors.Errorf("session ID error: %w", err)
	}
	cacheResult := RedisCacheApiName{
		"redisCacheApiName": map[string]interface{}{},
	}
	defer func() {
		if err != nil {
			return
		}
		b, _ := json.Marshal(cacheResult)
		err = c.cache.Set(c.ctx, reqKey, b, 0)
		if err != nil {
			c.log.Error("cache set error: %w", err)
		}
	}()

	edRes, err := c.equipmentDetailRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("equipmentRequest error: %w", err)
	}
	bpRes, err := c.businessPartnerRequest(&params.Params, sID, reqKey, &cacheResult)
	err = c.addHeaderInfo(&params.Params, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("search header error: %w", err)
	}
	err = c.fin(params, edRes, bpRes, reqKey, "EquipmentDetail", &cacheResult)
	if err != nil {
		return err
	}
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())

	return nil
}

func (c *EquipmentDetailCtrl) equipmentDetailRequest(
	params *dpfm_api_input_reader.EquipmentDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.EquipmentRes, error) {
	defer recovery(c.log)
	edReq := equipmentdetail.CreateEquipmentReq(params, sID, c.log)
	res, err := c.request("data-platform-api-equipment-master-reads-queue", edReq, sID, reqKey, "EquipmentDetail", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("Equipment cache set error: %w", err)
	}
	edRes, err := apiresponses.CreateEquipmentRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Equipment response parse error: %w", err)
	}
	return edRes, nil
}

func (c *EquipmentDetailCtrl) businessPartnerRequest(
	params *dpfm_api_input_reader.EquipmentDetailParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.BusinessPartnerRes, error) {
	defer recovery(c.log)
	bpReq := equipmentdetail.CreateBusinessPartnerReq(params, sID, c.log)
	res, err := c.request("data-platform-api-business-partner-reads-queue", bpReq, sID, reqKey, "EquipmentDetail", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartner cache set error: %w", err)
	}
	bpRes, err := apiresponses.CreateBusinessPartnerRes(res)
	if err != nil {
		return nil, xerrors.Errorf("BusinessPartner response parse error: %w", err)
	}
	return bpRes, nil
}

func (c *EquipmentDetailCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
	resFunc := c.rmq.SessionRequest(queue, req, sID)
	res := resFunc()
	if res == nil {
		return nil, xerrors.Errorf("receive nil response")
	}
	// redisKey := strings.Join([]string{url, api}, "/")
	// err := c.cache.Set(c.ctx, redisKey, res.Raw(), 1*time.Hour)
	// if err != nil {
	// 	return nil, xerrors.Errorf("cache set error: %w", err)
	// }
	// 	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return res, nil
}

func (c *EquipmentDetailCtrl) addHeaderInfo(
	params *dpfm_api_input_reader.EquipmentDetailParams,
	url string, setFlag *RedisCacheApiName,
) error {
	// "101@gmail.com/equipment/list/user=DeliverToParty/headerCompleteDeliveryIsDefined=false/headerBillingStatus=false/headerDeliveryBlockStatus=false/headerIssuingBlockStatus=false/headerReceivingBlockStatus=false/headerBillingBlockStatus=false/EquipmentList"
	key := fmt.Sprintf(`%s/equipment/list/user=%s/headerCompleteDeliveryIsDefined=%v/headerBillingStatus=%v/headerDeliveryBlockStatus=%v/headerIssuingBlockStatus=%v/headerReceivingBlockStatus=%v/headerBillingBlockStatus=%v/EquipmentList`,
		params.UserID, params.User, false, false, false, false, false, false)
	api := "EquipmentDetailHeader"
	m, err := c.cache.GetMap(c.ctx, key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	eqList := dpfm_api_output_formatter.EquipmentList{}
	err = json.Unmarshal(b, &eqList)
	if err != nil {
		return err
	}

	idx := -1
	for i, v := range eqList.Equipments {
		if v.Equipment == params.Equipment {
			idx = i
			break
		}
	}

	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": key, "index": idx}
	return nil
}

func (c *EquipmentDetailCtrl) fin(
	params *dpfm_api_input_reader.EquipmentDetail,
	edRes *apiresponses.EquipmentRes,
	bpRes *apiresponses.BusinessPartnerRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	bpMapper := map[int]apiresponses.BPGeneral{}
	for _, v := range *bpRes.Message.Generals {
		bpMapper[v.BusinessPartner] = v
	}

	data := dpfm_api_output_formatter.EquipmentDetail{}

	for _, v := range *edRes.Message.General {
		data.EquipmentDetail = append(data.EquipmentDetail,
			dpfm_api_output_formatter.EquipmentDetailGeneral{
				EquipmentCategory:        *v.EquipmentCategory,
				TechnicalObjectType:      *v.TechnicalObjectType,
				GrossWeight:              *v.GrossWeight,
				NetWeight:                *v.NetWeight,
				WeightUnit:               *v.WeightUnit,
				SizeOrDimensionText:      *v.SizeOrDimensionText,
				OperationStartDate:       *v.OperationStartDate,
				OperationStartTime:       *v.OperationStartTime,
				OperationEndDate:         *v.OperationEndDate,
				OperationEndTime:         *v.OperationEndTime,
				AcquisitionDate:          *v.AcquisitionDate,
				BusinessPartnerName:      *bpMapper[*v.BusinessPartner].BusinessPartnerName,
				ManifacturerSerialNumber: *v.ManufacturerSerialNumber,
				MasterFixedAsset:         *v.MasterFixedAsset,
				FixedAsset:               *v.FixedAsset,
				ValidityEndDate:          *v.ValidityEndDate,
				IsMarkedForDeletion:      *v.IsMarkedForDeletion,
			},
		)
	}

	for _, v := range *edRes.Message.BusinessPartner {
		data.EquipmentBP = append(data.EquipmentBP,
			dpfm_api_output_formatter.EquipmentBP{
				EquipmentPartnerObjectNmbr: v.EquipmentPartnerObjectNmbr,
				BusinessPartner:            v.BusinessPartner,
				PartnerFunction:            v.PartnerFunction,
				//BP名称
				ValidityStartDate: v.ValidityStartDate,
				ValidityEndDate:   v.ValidityEndDate,
			},
		)
	}

	if params.ReqReceiveQueue != nil {
		c.rmq.Send(*params.ReqReceiveQueue, map[string]interface{}{
			"runtime_session_id": params.RuntimeSessionID,
			"responseData":       data,
		})
	}

	redisKey := strings.Join([]string{url, api}, "/")
	// redisKey := strings.Join([]string{url, api, params.User}, "/")
	b, _ := json.Marshal(data)
	err := c.cache.Set(c.ctx, redisKey, b, 1*time.Hour)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]string{
		"keyName": redisKey,
	}
	return nil
}

func getRequestKey(req interface{}) (string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", err
	}
	rawReqID, ok := m["ui_key_function_url"]
	if !ok {
		return "", xerrors.Errorf("keyName not included")
	}

	return fmt.Sprintf("%v", rawReqID), nil
}

func getSessionID(req interface{}) (string, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return "", err
	}
	rawSID, ok := m["runtime_session_id"]
	if !ok {
		return "", xerrors.Errorf("runtime_session_id not included")
	}

	return fmt.Sprintf("%v", rawSID), nil
}

func extractEquipmentDetailParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.EquipmentDetail {
	data := dpfm_api_input_reader.ReadEquipmentDetail(msg)
	return data
}

type RedisCacheApiName map[string]map[string]interface{}

func (c *EquipmentDetailCtrl) Log(args ...interface{}) {
	for _, v := range args {
		b, _ := json.Marshal(v)
		c.log.Error("%s", string(b))
	}
}

func recovery(l *logger.Logger) {
	if e := recover(); e != nil {
		l.Error("%+v", e)
		return
	}
}
