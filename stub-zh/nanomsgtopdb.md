## 用法

来源：

- [官方 README](https://github.com/sangli00/nanomsgtopdb/blob/87df402495454362ccb4ed6ab58c33e7683e7401/README.md)
- [后台 worker 实现](https://github.com/sangli00/nanomsgtopdb/blob/87df402495454362ccb4ed6ab58c33e7683e7401/nanomsgtopdb.c)
- [扩展 SQL](https://github.com/sangli00/nanomsgtopdb/blob/87df402495454362ccb4ed6ab58c33e7683e7401/nanomsgtopdb--1.0.sql)

`nanomsgtopdb` 1.0 是一个历史性的 PipelineDB 后台 worker 扩展。它通过 nanomsg TCP 接收文本数据报，并写入固定的 PipelineDB 流 `public.generic_stream`。它依赖旧 PipelineDB 内部接口，并不是面向当前 PostgreSQL 的通用摄取扩展。

### 配置与启动

安装原生 nanomsg 依赖，把模块加入 `shared_preload_libraries`，设置启动时 GUC，然后重启：

```conf
shared_preload_libraries = 'pipelinedb,nanomsgtopdb'
nanomsg_work.ip_address = '127.0.0.1'
nanomsg_work.port = 9999
nanomsg_work.database = 'pipeline'
nanomsg_work.isrec = true
nanomsg_work.nanomsg_works = 1
nanomsg_work.nanomsg_max_tuples_a_transaction = 1000
```

随后在配置的数据库中创建扩展：

```sql
CREATE EXTENSION nanomsgtopdb;

CREATE CONTINUOUS VIEW message_count AS
SELECT count(*) FROM generic_stream;
```

安装 SQL 会创建 `public.generic_stream(data text)`。主 worker 绑定 `tcp://IP:port`，经本地 IPC 端点转发消息，一个或多个写入 worker 再按事务批量写入收到的文本。`nanomsg_work.format` 虽然存在且默认为 `text`，但实现中的流元组描述符只支持文本。

### 运维与安全边界

全部 GUC 都是 `PGC_POSTMASTER`，修改后必须重启。仅当 `nanomsg_work.isrec=true` 时接收器才启动；`nanomsg_work.nanomsg_works` 允许 1–10 个 worker，事务批量上限允许 1–10,000 条元组。

监听器不提供认证、加密、消息确认或应用级验证。上游默认绑定地址为 `0.0.0.0`；应改为可信接口并使用防火墙限制。发送端与数据库写入之间没有原子性，恶意或格式异常文本也会进入数据库流。代码面向已经停止发展的 PipelineDB 及旧后台 worker API，因此在现代 PostgreSQL 上很可能存在源码不兼容；若没有维护中的移植版和端到端故障测试，不应投入生产。
