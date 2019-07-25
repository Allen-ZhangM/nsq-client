package models

import (
	"errors"
	"github.com/nsqio/go-nsq"
)

const (
	MsgTopicLog  = "bossLog"
	MsgTopicInfo = "bossInfo"
)

type ProducerConf struct {
	//nsqd的tcp地址及端口
	Addr string
	//用于验证异步发送消息,建议设置为每秒消息量,如果为0则不验证异步发送消息是否成功
	DoneChanSize int
	//对异步发送消息的错误处理,error:错误信息，string1：topic，string2：msg内容
	HandleAsyncErrFunc func(err error, topic string, msg []byte)
}

var (
	conf        *ProducerConf
	nsqProducer *nsq.Producer = nil
	doneChan    chan *nsq.ProducerTransaction
)

//初始化nsqd
//addr : nsqd的tcp地址和端口, ischeck : 检查异步发送消息是否成功
func InitNsqd(c *ProducerConf) error {
	if nsqProducer != nil {
		return nil
	}
	if c == nil {
		return errors.New("ProducerConf is nil")
	}
	conf = c
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(conf.Addr, config)
	if err != nil {
		return err
	}
	nsqProducer = p

	if conf.DoneChanSize > 0 {
		doneChan = make(chan *nsq.ProducerTransaction, conf.DoneChanSize)
		go checkMsg()
	}

	return nil
}

//异步发送消息
//参数：doneChan，如果消息发送完，chan里会返回一个err和一个args，args为用户输入并原样返回，可以用于消息重发
func PublishAsync(topic string, msg []byte) {
	nsqProducer.PublishAsync(topic, msg, doneChan, topic, msg)
}

//同步发送消息
func Publish(topic string, msg []byte) error {
	return nsqProducer.Publish(topic, msg)
}

//验证异步发送的消息
func checkMsg() {
	for {
		d := <-doneChan
		if d.Error != nil {
			conf.HandleAsyncErrFunc(d.Error, d.Args[0].(string), d.Args[1].([]byte))
		}
	}
}
