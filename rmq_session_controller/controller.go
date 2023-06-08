package rmqsessioncontroller

import (
	"data-platform-api-request-reads-cache-manager-rmq-kube/config"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
)

type RMQSessionCtrl struct {
	rmq    *rabbitmq.RabbitmqClient
	reqMap map[string]chan rabbitmq.RabbitmqMessage
	log    *logger.Logger
	mt     sync.Mutex
}

func NewRMQSessionCtrl(conf *config.Conf, log *logger.Logger) (*RMQSessionCtrl, error) {
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.SessionControlQueue(), "", nil, 0)
	if err != nil {
		return nil, err
	}
	c := &RMQSessionCtrl{
		rmq:    rmq,
		reqMap: make(map[string]chan rabbitmq.RabbitmqMessage),
		log:    log,
	}
	go c.watchResponse()
	return c, nil
}

func (c *RMQSessionCtrl) watchResponse() {
	iter, err := c.rmq.Iterator()
	if err != nil {
		c.log.Fatal(err)
	}
	for msg := range iter {
		msg.Success()
		d := msg.Data()
		id, ok := d["runtime_session_id"]
		if !ok {
			c.log.Error("unknown message received")
			continue
		}
		sID, ok := id.(string)
		if !ok {
			c.log.Error("unknown message received. runtime_session_id is %v", id)
			continue
		}
		c.mt.Lock()
		_, ok = c.reqMap[sID]
		c.mt.Unlock()
		if !ok {
			c.log.Error("unknown runtime_session_id %v", id)
			continue
		}
		c.reqMap[sID] <- msg
	}
}

func (c *RMQSessionCtrl) SessionRequest(sendQueue string, payload interface{}, sessionID string) func() rabbitmq.RabbitmqMessage {
	c.rmq.Send(sendQueue, payload)
	c.mt.Lock()
	c.reqMap[sessionID] = make(chan rabbitmq.RabbitmqMessage)
	c.mt.Unlock()
	return c.receiveResponse(sessionID)
}

func (c *RMQSessionCtrl) receiveResponse(sID string) func() rabbitmq.RabbitmqMessage {
	var msg rabbitmq.RabbitmqMessage = nil
	return func() rabbitmq.RabbitmqMessage {
		ticker := time.NewTicker(10 * time.Second)
		select {
		case msg = <-c.reqMap[sID]:
			c.mt.Lock()
			delete(c.reqMap, sID)
			c.mt.Unlock()
		case <-ticker.C:
			c.log.Error("could not get response of session id %v", sID)
		}
		return msg
	}
}

func (c *RMQSessionCtrl) Send(sendQueue string, payload interface{}) error {
	return c.rmq.Send(sendQueue, payload)
}
