### 启动agent

```
1.配置nsqd.cfg文件
2.配置supervisord.conf
3.初始化supervisor
supervisord -c supervisord.conf
4.启动
supervisorctl start nsq-test
```

#### nsqd.cfg

- lookupd-tcp-address : 全局唯一，用于nsqd节点服务发现
- tcp-address : 默认0.0.0.0:4150
- http-address : 默认0.0.0.0:4151
- broadcast-address : 广播地址，需要其他机器访问则设置为访问ip
- data-path : 存放消息数据的磁盘路径，默认为当前路径。必须要创建好文件夹

### agent接口

1.调用初始化方法InitNsq(ProducerConf)
2.调用发送消息方法PublishAsync(异步)和Publish

### 说明

在服务所在机器部署agent，即nsqd

nsqd 是一个守护进程，负责接收，排队，投递消息给客户端

它主要负责message的收发，队列的维护。nsqd会默认监听一个tcp端口(4150)和一个http端口(4151)以及一个可选的https端口

tcp端口用于传输数据

http端口可以调用接口执行一些管理操作

nsqd 具有以下功能或特性

- 对订阅了同一个topic，同一个channel的消费者使用负载均衡策略（不是轮询）
- 只要channel存在，即使没有该channel的消费者，也会将生产者的message缓存到队列中（注意消息的过期处理）
- 保证队列中的message至少会被消费一次，即使nsqd退出，也会将队列中的消息暂存磁盘上(结束进程等意外情况除外)
- 限定内存占用，能够配置nsqd中每个channel队列在内存中缓存的message数量，一旦超出，message将被缓存到磁盘中
- topic，channel一旦建立，将会一直存在，要及时在管理台或者用代码清除无效的topic和channel，避免资源的浪费

### nsqd接口

```
//查看nsqd节点状态
curl http://172.19.0.87:4151/stats
```