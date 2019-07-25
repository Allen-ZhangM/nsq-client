package consumer

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

type Conf struct {
	//主题
	Topic string
	//管道
	Channel string
	//nsqd的tcp地址及端口
	AddrNsqd string
}

type NSQHandler struct{}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func InitNsqConsumer(conf Conf) {
	config := nsq.NewConfig()
	//默认值为1，如果5个nsqd有相同的topic，且想用一个consumer接收，需要配置为5
	config.MaxInFlight = 2
	config.LookupdPollInterval = time.Second //设置重连时间

	//可以建立多个连接 topic， channel， config
	consumer, err := nsq.NewConsumer(conf.Topic, conf.Channel, config)
	if nil != err {
		log.Println("NewConsumer err:", err)
		return
	}

	//可以设置不同的handler
	consumer.AddHandler(&NSQHandler{})
	err = consumer.ConnectToNSQLookupd(conf.AddrNsqd)
	if nil != err {
		log.Println("ConnectToNSQD err:", err)
		return
	}

}
