## 用法

来源：

- [官方上游文档](https://github.com/no0p/pgsampler/blob/9feffce68a902dc9b35fd1ead5ef2c8095e25124/README.md)
- [官方工作进程与 GUC 实现](https://github.com/no0p/pgsampler/blob/9feffce68a902dc9b35fd1ead5ef2c8095e25124/pgsampler.c)
- [官方输出循环实现](https://github.com/no0p/pgsampler/blob/9feffce68a902dc9b35fd1ead5ef2c8095e25124/lib/control_loop.c)

`pgsampler` 是实验性 PostgreSQL 后台工作进程，定期读取集群活动与统计信息，并把样本写入 CSV 文件或发送到网络接收器。它没有有意义的 SQL 扩展对象；只能通过预加载配置和服务器重启启用。

### 配置与启动

将工作进程加入 `shared_preload_libraries` 前，应先选择输出模式：

```conf
shared_preload_libraries = 'pgsampler'
pgsampler.target_db = 'postgres'
pgsampler.output_mode = 'csv'
pgsampler.output_csv_dir = '/var/lib/postgresql/pgsampler'
pgsampler.cycle_db_seconds = 500
```

安装与服务器匹配的共享库后重启 PostgreSQL。使用网络输出时，将 `pgsampler.output_mode` 设为 `network`，配置 `pgsampler.output_network_host`，并设置 `pgsampler.token`；源码连接 TCP 端口 24831。版本化 SQL 文件没有创建面向用户的对象，因此不需要 `CREATE EXTENSION` 流程。

### 采集数据与间隔

工作进程读取活动、`pg_stat_statements`、数据库、表、索引、列、后台写进程、文件系统、复制、系统和 GUC 信息。主要间隔设置包括 `pgsampler.heartbeat_seconds`、`pgsampler.system_seconds`、`pgsampler.relation_seconds`、`pgsampler.bgwriter_seconds`、`pgsampler.guc_seconds`、`pgsampler.activity_seconds`、`pgsampler.replication_seconds` 和 `pgsampler.statements_seconds`。CSV 模式会在配置目录下追加文件；网络模式则将采样记录发送到接收器。

### 安全与可靠性边界

样本可能包含查询文本、模式细节、配置、复制状态和主机信息。CSV 目录应使用 PostgreSQL 所有者权限保护，并配备轮转、保留和磁盘空间监控。已复核源码中的网络协议没有内置 TLS，并在握手中使用 token，因此只能部署在可信网络中，或放在经过认证的加密隧道后；出站连接应限制为预期接收器。

上游称该工作进程为实验性项目，只建议用于要求不高的集群。它以后台工作进程身份访问共享内存和所有数据库，并针对 PostgreSQL 9.3 时代的接口开发。应在一次性集群中验证准确的服务器 ABI、统计开销、重启行为、接收器故障、文件错误和移除流程。从 `shared_preload_libraries` 中删除它同样需要重启。
