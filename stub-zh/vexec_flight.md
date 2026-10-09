## 用法

来源：

- [modules/vexec_flight/README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_flight/README.md)
- [modules/vexec_flight/src/flight.h](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_flight/src/flight.h)
- [modules/vexec_flight/src/vexec_flight.c](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_flight/src/vexec_flight.c)
- [modules/vexec_flight/Makefile](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/modules/vexec_flight/Makefile)
- [README.md](https://github.com/igor-suhorukov/pg_vexec/blob/634d868e85493895b736e6f5afdb669777bba433/README.md)

`vexec_flight` 在 `vexec` 启用向量化执行时提供 PostgreSQL 19 的 Arrow Flight SQL 接口。它是 C 版本为 0.1.0 的预加载库，没有 SQL control 文件。

### 核心用法

```ini
shared_preload_libraries = 'vexec,vexec_flight'
vexec.mode = auto
vexec_flight.listen_addresses = '127.0.0.1'
vexec_flight.port = 32010
```

### 运行边界

将两个库追加到现有预加载列表，配置监听地址并重启。地址为空时不监听，默认端口为 32010。语句要求 `vexec.mode` 为 auto 或 force。不存在 `CREATE EXTENSION vexec_flight`；批量导入需要可选的 `vexec` SQL 对象和 `vexec.ingest_stream`。

连接使用 PostgreSQL 角色、数据库权限和 pg_hba.conf。向可信主机之外开放凭据前，配置 TLS 证书和 hostssl 规则。`vexec_flight.max_sessions` 默认为 16，使用后台工作进程槽位；`vexec_flight.idle_session_timeout` 限制空闲会话。模块提供查询、预备语句、事务、取消、元数据和导入接口，但不支持 Substrait 计划或 DoExchange。依赖 nghttp2、protobuf-c 和 OpenSSL。
