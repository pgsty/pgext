## 用法

来源：

- [README.md](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/README.md)
- [otel_api/otel_api--0.1.1.sql](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_api/otel_api--0.1.1.sql)
- [otel_api/otel_api.c](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_api/otel_api.c)
- [otel_api/otel_api.control](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_api/otel_api.control)

`otel_api` 提供共享轨迹上下文状态，以及供观测模块和导出器调用的 C rendezvous API。本身不通过 OTLP 发送轨迹，也不会在没有观测模块时自动创建查询 span。

### 核心用法

```ini
shared_preload_libraries = 'otel_api,otel_postgres_tracing'
```

```sql
CREATE EXTENSION otel_api;
BEGIN;
SET LOCAL otel_api.traceparent = '00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01';
SELECT otel_current_traceparent();
COMMIT;
```

### 运行边界

API 须在后端启动时加载。`otel_api.traceparent` 与 `otel_api.tracestate` 通过 GUC 降级路径传递上下文；`otel_current_traceparent()` 是由 `CREATE EXTENSION` 安装的可选 SQL 检查函数，创建该 C 函数要求超级用户。可配合 `otel_postgres_tracing` 观测查询，再由导出器消费 span。

项目仍是预发布状态，线协议和文件格式可能变化。原生 PostgreSQL 14 及以上支持 GUC/sqlcommenter 降级路径；协议头传播以及部分错误、并行进程行为要求配套的 PostgreSQL 补丁，不能假设普通服务器具有这些能力。要全局启用，应把库追加到现有预加载项并重启。轨迹可能包含敏感 SQL 和属性，应保护其访问权限。
