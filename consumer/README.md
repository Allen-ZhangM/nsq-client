### agent接口

1.调用初始化方法InitNsq(ProducerConf)
2.调用发送消息方法PublishAsync(异步)和Publish

### 参数说明

可以使用命令查看参数：`nsqd --help`

- nohup  ...  >nsqd2.log 2>&1& : 守护进程方式启动，nsqd2.log为日志
- lookupd-tcp-address : 全局唯一，用于nsqd节点服务发现
- tcp-address : 默认4150
- http-address : 默认4151
- broadcast-address : 广播地址，需要外网访问则设置为外网访问ip，不需要则删掉配置参数
- data-path : 存放消息数据的磁盘路径，默认为当前路径。必须要创建好文件夹