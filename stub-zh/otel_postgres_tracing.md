## 用法

来源：

- [README.md](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/README.md)
- [otel_postgres_tracing/otel_postgres_tracing--0.1.1.sql](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_postgres_tracing/otel_postgres_tracing--0.1.1.sql)
- [otel_postgres_tracing/otel_postgres_tracing.c](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_postgres_tracing/otel_postgres_tracing.c)
- [otel_postgres_tracing/otel_postgres_tracing.control](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_postgres_tracing/otel_postgres_tracing.control)

`otel_postgres_tracing` 通过 PostgreSQL 钩子观测查询执行和工具语句，使用共享的 `otel_api` 提供方；生成的 span 需要由导出器保存或发送。

### 核心用法

```ini
shared_preload_libraries = 'otel_api,otel_postgres_tracing,otel_demo_exporter'
otel_demo_exporter.output_file = '/var/log/postgresql/otel_spans.jsonl'
```

### 运行边界

载入库时安装钩子，推荐使用全局预加载；会话加载只能尽力观测当前后端。版本化 SQL 没有用户 SQL 对象，因此 `CREATE EXTENSION` 仅供目录登记，并不是启用步骤。当前修订不要求它与 API 按特定顺序预加载。

项目仍是预发布状态，线协议和文件格式可能变化。原生 PostgreSQL 14 及以上支持 GUC/sqlcommenter 降级路径；协议头传播以及部分错误、并行进程行为要求配套的 PostgreSQL 补丁，不能假设普通服务器具有这些能力。要全局启用，应把库追加到现有预加载项并重启。轨迹可能包含敏感 SQL 和属性，应保护其访问权限。
