### log-srv build

作为nsq的注册发现中心，

```
1.安装编译文件
nsq-build/bin目录下的文件放入/usr/local/bin下

```

### 部署nsqlookupd

用于nsqd服务注册与发现

```
1.配置nsqlookupd.cfg文件
2.配置supervisord.conf
3.初始化supervisor
supervisord -c supervisord.conf
4.启动
supervisorctl start nsqlookupd

```

- tcp-address : 用于让其他nsqd注册
- http_address : 用于通过接口发现其他nsqd

### 部署nsqadmin UI监控

用来汇集集群的实时统计，并执行不同的管理任务

```
1.配置nsqadmin.cfg文件
2.配置supervisord.conf
3.初始化supervisor
supervisord -c supervisord.conf
4.启动
supervisorctl start nsqadmin

```

主要修改参数

- lookupd-http-address: nsqlookupd地址
- http-address: 服务地址和端口，对外访问需开启外网端口


- 提供一个对topic和channel统一管理的操作界面以及各种实时监控数据的展示，界面设计的很简洁，操作也很简单
- 展示所有message的数量
- 能够在后台创建和删除topic和channel
- nsqadmin的所有功能都必须依赖于nsqlookupd，nsqadmin只是向nsqlookupd传递用户操作并展示来自nsqlookupd的数据

### 部署nsq_to_file

自动创建一个channl并把消息写入文件

``` 
1.启动

//注意修改启动命令的4个参数
nohup \
nsq_to_file \
--topic=bossLog \
--output-dir=/home/admin/log-srv/nsq_to_file_data \
--lookupd-http-address=127.0.0.1:4161 \
--datetime-format=%Y-%m-%d_%H \
>logs/nsq_to_file.log 2>&1&
```

- topic : topic
- output-dir : 输出消息到磁盘的目录，文件名称bossLog.ubuntu.2019-06-26_13.log，命名格式{topic}.{hostname}.{日期_小时}.log
- lookupd-http-address : 注意是http端口
- datetime-format : 日志文件命名格式，%Y-%m-%d_%H 为每小时一个日志文件