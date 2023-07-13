package cache

import (
	"data-platform-request-reads-cache-manager-rmq-kube/services"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	"strings"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

type Cache struct {
	rds       *redis.Client
	log       *logger.Logger
	keyPrefix *string
}

func NewCache(
	addr string,
	port interface{},
	log *logger.Logger,
	databaseNumber int,
	keyPrefix *string,
) *Cache {
	redisAddr := fmt.Sprintf("%s:%v", addr, port)
	opt := redis.Options{
		Addr: redisAddr,
		DB:   databaseNumber,
		// Password: "",
		//KeyPrefix: "",
	}

	cli := redis.NewClient(&opt)

	return &Cache{
		rds:       cli,
		log:       log,
		keyPrefix: keyPrefix,
	}
}

func (c *Cache) CreateKey(
	controller *beego.Controller,
	categories []string,
) string {
	userInfo := services.UserRequestParams(controller)

	category := strings.Join(categories, "/")

	redisKey := strings.Join([]string{
		*userInfo.UserID,
		category,
	}, "/")

	return redisKey
}

func (c *Cache) ConfirmCashDataExisting(
	redisKey string,
) ([]byte, error) {
	data, err := c.GetRaw(c.fixedRedisKey(redisKey))
	if err != nil {
		return nil, xerrors.Errorf("cache data does not exist: %w", err)
	}
	return data, nil
}

func (c *Cache) SetCache(
	redisKey string,
	data interface{},
) error {
	var expiration time.Duration

	b, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return c.Set(redisKey, b, expiration)
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) error {
	res := c.rds.Set(c.fixedRedisKey(key), value, expiration)
	if err := res.Err(); err != nil {
		return xerrors.Errorf("cache set error: %w", err)
	}
	c.log.Info("cache set to key %s", key)
	return nil
}

func (c *Cache) GetRaw(key string) ([]byte, error) {
	res := c.rds.Get(c.fixedRedisKey(key))
	if err := res.Err(); err != nil {
		return []byte{}, xerrors.Errorf("cache get error: %w", err)
	}
	d, err := res.Bytes()
	if err != nil {
		return []byte{}, xerrors.Errorf("cache get bytes error: %w", err)
	}
	return d, nil
}
func (c *Cache) GetSlice(key string) ([]map[string]interface{}, error) {
	d := []map[string]interface{}{}
	b, err := c.GetRaw(c.fixedRedisKey(key))
	if err != nil {
		return d, xerrors.Errorf("cache get raw error: %w", err)
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return d, xerrors.Errorf("cache data unmarshal error: %w", err)
	}
	return d, nil
}
func (c *Cache) GetMap(key string) (map[string]interface{}, error) {
	d := map[string]interface{}{}
	b, err := c.GetRaw(c.fixedRedisKey(key))
	if err != nil {
		return d, xerrors.Errorf("cache get raw error: %w", err)
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return d, xerrors.Errorf("cache data unmarshal error: %w", err)
	}
	return d, nil
}

func (c *Cache) GetAllKeys() ([]string, error) {
	keys, _, err := c.rds.Scan(0, "prefix:*", 0).Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (c *Cache) fixedRedisKey(
	redisKey string,
) string {
	fixedRedisKey := redisKey

	if c.keyPrefix != nil {
		fixedRedisKey = strings.Join([]string{
			fmt.Sprintf("%s:", *c.keyPrefix),
			redisKey,
		}, "")
	}

	return fixedRedisKey
}
