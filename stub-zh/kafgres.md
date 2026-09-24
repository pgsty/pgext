## 用法

来源：

- [README 0.1.0](https://github.com/RayElg/kafgres/blob/0.1.0/README.md)
- [Configuration 0.1.0](https://github.com/RayElg/kafgres/blob/0.1.0/docs/configuration.md)

`kafgres` 在 PostgreSQL 中提供 Kafka 协议代理。0.1.0 版使用上游 pgrx 0.16.1 源码，本包支持 PostgreSQL 16。需要超级用户安装、共享预加载及重启，许可为 Elastic License 2.0。

### 启用代理

```conf
shared_preload_libraries = 'kafgres'
kafgres.database = 'postgres'
kafgres.bind_host = '127.0.0.1'
kafgres.advertised_host = '127.0.0.1'
kafgres.port = 9092
```

```sql
CREATE EXTENSION kafgres;
SELECT kafgres_create_topic('demo', 1);
BEGIN;
SELECT kafgres_produce('demo', 'key', 'value');
COMMIT;
SELECT * FROM kafgres_partition_offsets('demo');
```

Kafka 客户端连接配置的代理端口，SQL 消息生产参与调用者的事务。将监听器暴露到受信任本地环境之外前，应配置 TLS、身份认证和访问控制。

### 存储与变更捕获

`kafgres.storage_engine` 默认为 segment，日志保存在独立文件中，需要使用扩展自己的复制与归档流程。依赖 segment 保留和恢复能力前，应配置 `kafgres.segment_archive_command` 并监控 `kafgres_archive_status()`。普通 PostgreSQL WAL/PITR 无法覆盖整个 segment 日志。table 引擎将日志保存在 PostgreSQL 表内；切换引擎不会迁移已有记录。

CDC 还需要 `wal_level = logical` 和适当的 `output_plugin_libraries` 白名单。部署前应审阅映射及恢复流程。本包不声明其他主版本支持：上游 0.1.0 的解码代码在 PostgreSQL 17 上编译失败。
