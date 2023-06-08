package productlist

import (
	"context"
	dpfm_api_input_reader "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-request-reads-cache-manager-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-request-reads-cache-manager-rmq-kube/api_requests/creator/productlist"
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

// CreateProductItemsReq
type ProductListCtrl struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewProductListCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *ProductListCtrl {
	return &ProductListCtrl{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *ProductListCtrl) ProductList(msg rabbitmq.RabbitmqMessage, l *logger.Logger) error {
	start := time.Now()
	params := extractProductListParam(msg)
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

	pRes, err := c.productRequest(&params.Params, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}
	pmdRes, err := c.productMasterDocRequest(&params.Params, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}
	pgRes, err := c.productGroupRequest(&params.Params, pRes, sID, reqKey, &cacheResult, l)
	if err != nil {
		return err
	}

	c.fin(params, pRes, pmdRes, pgRes, reqKey, "ProductList", &cacheResult)
	c.log.Info("Fin: %d ms\n", time.Since(start).Milliseconds())
	return nil
}

func (c *ProductListCtrl) productGroupRequest(
	params *dpfm_api_input_reader.ProductListParams,
	pmRes *apiresponses.ProductMasterRes,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.ProductGroupRes, error) {
	defer recovery(c.log)
	pgReq := productlist.CreateProductGroupReq(params, pmRes, sID, c.log)
	res, err := c.request("data-platform-api-product-group-reads-queue", pgReq, sID, reqKey, "ProductGroup", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("ProductGroup cache set error: %w", err)
	}
	pgRes, err := apiresponses.CreateProductGroupRes(res)
	if err != nil {
		return nil, xerrors.Errorf("ProductGroup response parse error: %w", err)
	}
	return pgRes, nil
}

func (c *ProductListCtrl) productRequest(
	params *dpfm_api_input_reader.ProductListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.ProductMasterRes, error) {
	defer recovery(c.log)
	pmReq := productlist.CreateProductRequest(params, sID, c.log)
	res, err := c.request("data-platform-api-product-master-reads-queue", pmReq, sID, reqKey, "Product", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	pmRes, err := apiresponses.CreateProductMasterRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}
	return pmRes, nil
}

func (c *ProductListCtrl) productMasterDocRequest(
	params *dpfm_api_input_reader.ProductListParams,
	sID string,
	reqKey string,
	setFlag *RedisCacheApiName,
	l *logger.Logger,
) (*apiresponses.ProductMasterDocRes, error) {
	defer recovery(c.log)
	pmReq := productlist.CreateProductMasterDocReq(params, sID, c.log)
	res, err := c.request("data-platform-api-product-master-doc-reads-queue", pmReq, sID, reqKey, "ProductMasterDoc", setFlag)
	if err != nil {
		return nil, xerrors.Errorf("product cache set error: %w", err)
	}
	pmRes, err := apiresponses.CreateProductMasterDocRes(res)
	if err != nil {
		return nil, xerrors.Errorf("Product master response parse error: %w", err)
	}
	return pmRes, nil
}

func (c *ProductListCtrl) request(queue string, req interface{}, sID string, url, api string, setFlag *RedisCacheApiName) (rabbitmq.RabbitmqMessage, error) {
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

func extractProductListParam(msg rabbitmq.RabbitmqMessage) *dpfm_api_input_reader.ProductList {
	data := dpfm_api_input_reader.ReadProductList(msg)
	return data
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

type RedisCacheApiName map[string]map[string]interface{}

func (c *ProductListCtrl) fin(
	params *dpfm_api_input_reader.ProductList,
	pmRes *apiresponses.ProductMasterRes,
	pmdRes *apiresponses.ProductMasterDocRes,
	pgRes *apiresponses.ProductGroupRes,
	url, api string, setFlag *RedisCacheApiName,
) error {
	pNames := make(map[string]*string)
	for _, v := range *pmRes.Message.ProductDescByBP {
		pNames[v.Product] = v.ProductDescription
	}
	pgNames := make(map[string]*string)
	for _, v := range *pgRes.Message.ProductGroupText {
		pgNames[v.ProductGroup] = v.ProductGroupName
	}

	docs := make(map[string]*dpfm_api_output_formatter.ProductImage)
	for _, v := range *pmdRes.Message.HeaderDoc {
		docs[v.Product] = &dpfm_api_output_formatter.ProductImage{
			BusinessPartnerID: v.DocIssuerBusinessPartner,
			DocID:             v.DocID,
			FileExtension:     v.FileExtension,
		}
	}

	productionOrders := make([]dpfm_api_output_formatter.Product, 0)
	for _, v := range *pmRes.Message.General {
		productionOrders = append(productionOrders, dpfm_api_output_formatter.Product{
			Product:             v.Product,
			ProductDescription:  pNames[v.Product],
			ProductGroup:        v.ProductGroup,
			ProductGroupName:    pgNames[*v.ProductGroup],
			BaseUnit:            v.BaseUnit,
			ValidityStartDate:   v.ValidityStartDate,
			IsMarkedForDeletion: v.IsMarkedForDeletion,
			Images: dpfm_api_output_formatter.Images{
				Product: docs[v.Product],
			},
		},
		)
	}

	data := dpfm_api_output_formatter.ProductList{
		Products: productionOrders,
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

func (c *ProductListCtrl) finEmptyProcess(
	params interface{},
	url, api string, setFlag *RedisCacheApiName,

) error {
	data := dpfm_api_output_formatter.DeliveryDocumentList{
		DeliveryDocuments: make([]dpfm_api_output_formatter.DeliveryDocument, 0),
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

func orderAsc[T any](d map[int]T) []T {
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

func orderDesc[T any](d map[int]T) []T {
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

func (c *ProductListCtrl) Log(args ...interface{}) {
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
