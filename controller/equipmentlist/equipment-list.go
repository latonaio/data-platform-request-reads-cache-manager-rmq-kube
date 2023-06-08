package equipmentlist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/equipmentlist"
	apiresponses "data-platform-api-request-reads-cache-manager-rmq-kube/api_responses"
	"data-platform-api-request-reads-cache-manager-rmq-kube/cache"
	rmqsessioncontroller "data-platform-api-request-reads-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type EquipmentListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewEquipmentListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *EquipmentListCtrl {
	return &EquipmentListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *EquipmentListCtrl) EquipmentList(msg rabbitmq.RabbitmqMessage) error {
	start := time.Now()
	params := extractEquipmentListParam(msg)
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
		b, _ := json.Marshal(cacheResult)
		err = c.cache.Set(c.ctx, reqKey, b, 0)
		if err != nil {
			c.log.Error("cache set error: %w", err)
		}
	}()

	eqRes, err := c.equipmentRequest(&params.Params, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("equipmentRequest error: %w", err)
	}
	if eqRes.Message.General == nil || len(*eqRes.Message.General) == 0 {
		c.finEmptyProcess(params, reqKey, "EquipmentList", &cacheResult)
		c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
		return nil
	}

	plRes, err := c.plantRequest(&params.Params, eqRes, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("plantRequest error: %w", err)
	}

	etRes, err := c.equipmentTypeRequest(&params.Params, eqRes, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("equipmentTypeRequest error: %w", err)
	}

	eqdRes, err := c.equipmentDocRequest(&params.Params, eqRes, sID, reqKey, &cacheResult)
	if err != nil {
		return xerrors.Errorf("equipmentDocRequest error: %w", err)
	}

	c.fin(params, eqRes, plRes, etRes, eqdRes, reqKey, "EquipmentList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *EquipmentListCtrl) equipmentDocRequest(
	params *dpfm_api_input_reader.EquipmentListParams,
	eqRes *apiresponses.EquipmentRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.EquipmentRes, error) {
	defer recovery(c.log)
	eqReq := equipmentlist.CreateProductMasterDocReq(params, eqRes, sID, c.log)
	res, err := c.request("data-platform-api-equipment-master-reads-queue", eqReq, sID, reqKey, "EquipmentList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("Equipment cache set error: %w", err)
	}
	eqdRes, err := apiresponses.CreateEquipmentRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Equipment response parse error: %w", err)
	}
	return eqdRes, nil
}

func (c *EquipmentListCtrl) equipmentRequest(
	params *dpfm_api_input_reader.EquipmentListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.EquipmentRes, error) {
	defer recovery(c.log)
	eqReq := equipmentlist.CreateEquipmentReq(params, sID, c.log)
	res, err := c.request("data-platform-api-equipment-master-reads-queue", eqReq, sID, reqKey, "EquipmentList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("Equipment cache set error: %w", err)
	}
	eqRes, err := apiresponses.CreateEquipmentRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Equipment response parse error: %w", err)
	}
	return eqRes, nil
}

func (c *EquipmentListCtrl) plantRequest(
	params *dpfm_api_input_reader.EquipmentListParams,
	eqRes *apiresponses.EquipmentRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.PlantRes, error) {
	defer recovery(c.log)
	plReq := equipmentlist.CreatePlantReq(params, eqRes, sID, c.log)
	res, err := c.request("data-platform-api-plant-reads-queue", plReq, sID, reqKey, "EquipmentList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("Plant cache set error: %w", err)
	}
	plRes, err := apiresponses.CreatePlantRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Plant response parse error: %w", err)
	}
	return plRes, nil
}

func (c *EquipmentListCtrl) equipmentTypeRequest(
	params *dpfm_api_input_reader.EquipmentListParams,
	eqRes *apiresponses.EquipmentRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
) (*apiresponses.EquipmentTypeRes, error) {
	defer recovery(c.log)
	etReq := equipmentlist.CreateEquipmentTypeReq(params, eqRes, sID, c.log)
	res, err := c.request("data-platform-api-equipment-type-reads-queue", etReq, sID, reqKey, "EquipmentTypeList", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("Equipment cache set error: %w", err)
	}
	etRes, err := apiresponses.CreateEquipmentTypeRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Equipment response parse error: %w", err)
	}
	return etRes, nil
}

func (c *EquipmentListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func (c *EquipmentListCtrl) fin(
	params *dpfm_api_input_reader.EquipmentList,
	eqRes *apiresponses.EquipmentRes,
	plRes *apiresponses.PlantRes,
	etRes *apiresponses.EquipmentTypeRes,
	eqdRes *apiresponses.EquipmentRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	plantMapper := map[string]apiresponses.PlantGeneral{}
	for _, v := range *plRes.Message.Generals {
		plantMapper[v.Plant] = v
	}

	etype := map[string]*string{}
	for _, v := range *etRes.Message.EquipmentTypeText {
		etype[v.EquipmentType] = &v.EquipmentTypeName
	}

	imgMapper := map[int]apiresponses.EquipmentGeneralDoc{}
	for _, v := range *eqdRes.Message.GeneralDoc {
		imgMapper[v.Equipment] = v
	}

	data := dpfm_api_output_formatter.EquipmentList{}
	for _, v := range *eqRes.Message.General {
		img := &dpfm_api_output_formatter.ProductImage{}
		if i, ok := imgMapper[v.Equipment]; ok {
			img = &dpfm_api_output_formatter.ProductImage{
				BusinessPartnerID: *i.DocIssuerBusinessPartner,
				DocID:             i.DocID,
				FileExtension:     i.FileExtension,
			}
		}

		data.Equipments = append(data.Equipments,
			dpfm_api_output_formatter.Equipment{
				Equipment:           v.Equipment,
				EquipmentName:       *v.EquipmentName,
				EquipmentTypeName:   etype[*v.EquipmentType],
				PlantName:           plantMapper[v.MaintenancePlant].PlantName,
				ValidityStartDate:   v.ValidityStartDate,
				IsMarkedForDeletion: v.IsMarkedForDeletion,
				Images: dpfm_api_output_formatter.Images{
					Equipment: img,
				},
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
	err := c.cache.Set(c.ctx, redisKey, b, 0)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return nil
}

func (c *EquipmentListCtrl) finEmptyProcess(
	params interface{},
	url, api string, setFlag *RedisCacheApiName,

) error {
	data := dpfm_api_output_formatter.EquipmentList{
		Equipments: make([]dpfm_api_output_formatter.Equipment, 0),
	}

	redisKey := strings.Join([]string{url, api}, "/")
	b, _ := json.Marshal(data)
	err := c.cache.Set(c.ctx, redisKey, b, 0)
	if err != nil {
		return nil
	}
	(*setFlag)["redisCacheApiName"][api] = map[string]interface{}{"keyName": redisKey, "request": params}
	return nil
}

func extractEquipmentListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.EquipmentList {
	data := dpfm_api_input_reader.ReadEquipmentList(msg)
	return data
}

type RedisCacheApiName map[string]map[string]interface{}

func equipmentAsc[T any](d map[int]T) []T {
	ids := make([]int, 0, len(d))
	for i := range d {
		ids = append(ids, i)
	}
	sort.Ints(ids)
	sli := make([]T, 0, len(d))
	for _, i := range ids {
		sli = append(sli, d[i])
	}
	return sli
}

func equipmentDesc[T any](d map[int]T) []T {
	ids := make([]int, 0, len(d))
	for i := range d {
		ids = append(ids, i)
	}
	sort.Ints(ids)
	sli := make([]T, 0, len(d))
	for i := len(ids) - 1; i >= 0; i-- {
		sli = append(sli, d[ids[i]])
	}
	return sli
}

func (c *EquipmentListCtrl) Log(args ...interface{}) {
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
