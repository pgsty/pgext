## 用法

来源：

- [README.md](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/README.md)
- [otel_demo_exporter/otel_demo_exporter--0.1.1.sql](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_demo_exporter/otel_demo_exporter--0.1.1.sql)
- [otel_demo_exporter/otel_demo_exporter.c](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_demo_exporter/otel_demo_exporter.c)
- [otel_demo_exporter/otel_demo_exporter.control](https://github.com/ringerc/postgres_otel_api/blob/7f1b649dd7008de651890996bfc6f671dfac3b2d/otel_demo_exporter/otel_demo_exporter.control)

`otel_demo_exporter` 是 `otel_api` span 的最小文件导出器，同步写入少量 JSON 字段，明确定位为开发示例，而非生产遥测流水线。

### 核心用法

```ini
shared_preload_libraries = 'otel_api,otel_postgres_tracing,otel_demo_exporter'
otel_demo_exporter.output_file = '/var/log/postgresql/otel_spans.jsonl'
```

### 运行边界

加载 API、观测生产者和此导出器即可启用；SQL 不包含用户对象，`CREATE EXTENSION` 仅供目录登记。将 `otel_demo_exporter.output_file` 指向服务器可写文件，空设置禁用输出，相对路径以数据目录为基准。修改后重载配置，后端会重新打开目标文件。

它没有批量发送、后台分发或 OTLP/HTTP/gRPC 传输；同步写入可能拖慢查询并耗尽磁盘，须设置访问控制、轮转和保留策略。旧的子目录 README 使用了已被替代的模块名，应以当前根 README 和源码中的 `otel_api` 加载路径为准。

项目仍是预发布状态，线协议和文件格式可能变化。原生 PostgreSQL 14 及以上支持 GUC/sqlcommenter 降级路径；协议头传播以及部分错误、并行进程行为要求配套的 PostgreSQL 补丁，不能假设普通服务器具有这些能力。要全局启用，应把库追加到现有预加载项并重启。轨迹可能包含敏感 SQL 和属性，应保护其访问权限。
