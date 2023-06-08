package update

import (
	"context"
	"data-platform-api-request-reads-cache-manager-rmq-kube/cache"
	rmqsessioncontroller "data-platform-api-request-reads-cache-manager-rmq-kube/rmq_session_controller"
	"encoding/json"

	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type Update struct {
	cache *cache.Cache
	rmq   *rmqsessioncontroller.RMQSessionCtrl
	ctx   context.Context
	log   *logger.Logger
}

func NewUpdateCtrl(ctx context.Context, c *cache.Cache, rmq *rmqsessioncontroller.RMQSessionCtrl, log *logger.Logger) *Update {
	return &Update{
		cache: c,
		rmq:   rmq,
		ctx:   ctx,
		log:   log,
	}
}

func (c *Update) Update(msg rabbitmq.RabbitmqMessage) error {
	keys, err := c.cache.GetAllKeys()
	if err != nil {
		return err
	}

	for _, key := range keys {
		m, err := c.cache.GetMap(c.ctx, key)
		if err != nil {
			return err
		}
		if _, ok := m["redisCacheApiName"]; !ok {
			continue
		}
		d := map[string]interface{}{}
		b, err := json.Marshal(m["redisCacheApiName"])
		if err != nil {
			c.log.Error(err)
			continue
		}
		err = json.Unmarshal(b, &d)
		if err != nil {
			c.log.Error(err)
			continue
		}

		for _, v := range d {
			b, err := json.Marshal(v)
			if err != nil {
				c.log.Error(err)
				continue
			}
			err = json.Unmarshal(b, &d)
			if err != nil {
				c.log.Error(err)
				continue
			}
			if _, ok := d["request"]; ok {
				b, err := json.Marshal(d["request"])
				if err != nil {
					c.log.Error(err)
					continue
				}
				err = json.Unmarshal(b, &d)
				if err != nil {
					c.log.Error(err)
					continue
				}
				c.rmq.Send("data-platform-api-request-reads-cache-manager-queue", d)
				break
			}
		}
	}

	return nil
}
