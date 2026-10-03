## 用法

来源：

- [influx.control](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/influx.control)
- [README.md](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/README.md)
- [docs/getting-started.md](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/docs/getting-started.md)
- [docs/options.md](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/docs/options.md)
- [docs/procedures.md](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/docs/procedures.md)
- [influx--0.4.sql](https://github.com/timescale/pg_influx/blob/2b7017ea351e199f9f7b72822dd4a9f9c467024d/influx--0.4.sql)

`influx` 通过后台工作进程接收 InfluxDB 行协议数据报并写入 PostgreSQL 表。上游明确将其定位为教学实验，UDP 传输可能丢失数据。

### 核心用法

```sql
CREATE SCHEMA metrics;
CREATE EXTENSION influx WITH SCHEMA metrics;
SELECT metrics.worker_launch('8089');
CALL metrics.send_packet('cpu,host=demo usage=1.5 1574753954000000000', '8089');
```

### 运行边界

由超级用户安装。手动启动工作进程不需要服务器启动时预加载；如需自动启动，应配置 `shared_preload_libraries`、`influx.database`、`influx.schema`、`influx.role` 和 `influx.workers`，然后重启。角色默认为超级用户，采集时应选择受限角色。`worker_launch` 返回进程 PID，`send_packet` 用于发送测试数据，`_create` 控制自动建表。监听器绑定被动 UDP 地址，不能直接选择具体接口，应限制网络可达性。文档中的开发目标为 PostgreSQL 13。
