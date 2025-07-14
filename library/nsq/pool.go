package nsq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"middle/internal/consts"
	"middle/library/tool"
	"sync/atomic"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

// GetNsdGroup 根据topic来返回对应的生产者对象
func GetNsdGroup(ctx context.Context, topic string) ([]*nsq.Producer, error) {
	groupName := ""
	ok := false
	//只读map，没问题，直接访问
	if groupName, ok = consts.Topics[topic]; !ok {
		//直接崩掉，不能掩盖问题
		g.Log().Errorf(ctx, "nsq topic not defined, topic=%s", topic)
		panic(gerror.New("nsq topic not defined, topic=" + topic))
	}

	instanceKey := fmt.Sprintf("self-go-nsq-group.%s", groupName)

	result, _ := gcache.GetOrSetFuncLock(ctx, instanceKey, func(ctx context.Context) (interface{}, error) {
		//result := gins.GetOrSetFuncLock(instanceKey, func() interface{} {
		//因配置错误导致的，需要直接panic
		addrVars, _ := g.Cfg().Get(ctx, fmt.Sprintf("%s.%s", consts.NsqConfigNsdName, groupName))

		addrs := addrVars.Strings()
		if len(addrs) == 0 {
			panic(gerror.New("nsq config name error"))
		}

		ps := []*nsq.Producer{}
		ip, err := tool.IP.LocalIPv4s()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			cfg := nsq.NewConfig()
			cfg.ReadTimeout = time.Second * 30
			cfg.HeartbeatInterval = time.Second * 15
			cfg.Hostname = ip
			producer, err := nsq.NewProducer(addr, cfg)

			if err != nil {
				panic(gerror.Wrap(err, "nsq create producer error"))
			}
			ps = append(ps, producer)
		}

		return ps, nil
	}, 0)

	if result == nil {
		return nil, gerror.New("error get nsd")
	}

	ps := result.Array()
	nsds := []*nsq.Producer{}
	for _, p := range ps {
		nsds = append(nsds, p.(*nsq.Producer))
	}
	//if nsds, ok := ps.([]*nsq.Producer); ok {
	return nsds, nil
	//}

	return nil, gerror.New("error get nsd")
}

// NewNsqClient 实例化一个Client对象
func NewNsqClient() *Client {
	return &Client{}
}

// Client 定义消息发送对象
type Client struct {
	index uint64
}

// Send 发送消息，使用json编码
func (c *Client) Send(ctx context.Context, topic string, body interface{}, delay ...time.Duration) error {
	bytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	nsds, err := GetNsdGroup(ctx, topic)
	if err != nil {
		return err
	}
	index := atomic.AddUint64(&c.index, 1)
	client := nsds[int(index%uint64(len(nsds)))]
	if len(delay) > 0 {
		err = client.DeferredPublish(topic, delay[0], bytes)
	} else {
		err = client.Publish(topic, bytes)
	}
	return err
}

func (c *Client) SendBytes(ctx context.Context, topic string, body []byte, delay ...time.Duration) error {
	nsds, err := GetNsdGroup(ctx, topic)
	if err != nil {
		return err
	}
	index := atomic.AddUint64(&c.index, 1)
	client := nsds[int(index%uint64(len(nsds)))]
	if len(delay) > 0 {
		err = client.DeferredPublish(topic, delay[0], body)
	} else {
		err = client.Publish(topic, body)
	}
	return err
}
