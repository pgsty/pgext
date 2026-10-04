## 用法

来源：

- [extension/sql/kafgres--0.2.0--0.3.0.sql](https://github.com/RayElg/kafgres/blob/8def1919c65f71217665cdc9b72df4c693811cc2/extension/sql/kafgres--0.2.0--0.3.0.sql)
- [.github/workflows/ci.yml](https://github.com/RayElg/kafgres/blob/8def1919c65f71217665cdc9b72df4c693811cc2/.github/workflows/ci.yml)
- [README 0.3.0](https://github.com/RayElg/kafgres/blob/8def1919c65f71217665cdc9b72df4c693811cc2/README.md)
- [Configuration 0.3.0](https://github.com/RayElg/kafgres/blob/8def1919c65f71217665cdc9b72df4c693811cc2/docs/configuration.md)

- [0.3.0 release](https://github.com/RayElg/kafgres/releases/tag/0.3.0)

`kafgres` 在 PostgreSQL 中提供 Kafka 协议代理。上游 0.3.0 构建并测试 PostgreSQL 13–18。需要超级用户安装、共享预加载及重启，许可为 Elastic License 2.0。

### 启用代理

```ini
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

CDC 还需要 `wal_level = logical`；部分 PostgreSQL 构建另外要求配置 `output_plugin_libraries` 白名单。0.3.0 支持带投影和过滤的 SQL CDC 映射。部署前应审阅映射与恢复流程；标签中的 CI 矩阵及发行产物覆盖 PostgreSQL 13–18。

### 持久性设置

0.3.0 默认开启 `kafgres.fsync_before_ack`，默认关闭 `kafgres.relaxed_produce_commit`。放宽前者可能在断电时丢失已经确认的分段记录；放宽后者可能在崩溃后丢失最新的幂等生产者状态，使重试产生重复记录。这些参数的作用范围小于事务性 SQL 生产，并不统一作用于表引擎。应保留严格默认值，直到明确接受对应的持久性取舍。

### 0.3.0 升级

安装匹配的 0.3.0 库并重启后，在 broker 数据库执行 `ALTER EXTENSION kafgres UPDATE`。SQL 升级收紧管理函数授权，并固定生产函数的 search_path。段存储引擎首次启动时会移除 0.2.0 的旧时间戳索引；旧段仍可通过扫描读取，新段生成当前格式的索引。
