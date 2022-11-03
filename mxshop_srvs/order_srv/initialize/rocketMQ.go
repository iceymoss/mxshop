package initialize

import (
	"fmt"
	"mxshop_srvs/order_srv/global"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go.uber.org/zap"
)

func InitMQ() {
	//test
	global.GroupInventory = producer.WithGroupName("mxshop-inventory")
	global.GroupOrder = producer.WithGroupName("mxshop-order")

	socket := fmt.Sprintf("%s:%d", global.ServerConfig.MqInfo.Host, global.ServerConfig.MqInfo.Port)
	fmt.Println(socket)
	global.MQInventory = InitMQNewProducer(global.GroupInventory, socket)
	global.MQOrder = InitMQNewProducer(global.GroupOrder, socket)
}

func InitMQNewProducer(Producer producer.Option, Socket string) rocketmq.Producer {
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{Socket}),
		Producer)
	if err != nil {
		zap.S().Errorf("初始化生产者失败%s", err)
	}
	if err = p.Start(); err != nil {
		zap.S().Errorf("启动生产者失败%s", err)
	}
	return p
}
