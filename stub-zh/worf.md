## 用法

来源：

- [README.md](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/README.md)
- [LICENSE](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/LICENSE)
- [src/Worf/Worf.csproj](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/src/Worf/Worf.csproj)
- [src/Worf/worf.control](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/src/Worf/worf.control)
- [src/Worf/Worf.cs](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/src/Worf/Worf.cs)
- [src/Worf/Node.cs](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/src/Worf/Node.cs)
- [.github/workflows/ci.yml](https://github.com/ankus-net/worf/blob/7299f69272cbde9ee01b0cd03e46fc7ddeb54597/.github/workflows/ci.yml)

`worf` 提供逐数据库逻辑复制，支持复制集、多上游与可配置的冲突处理。此 C#/.NET Native AOT 源码在 Ankus 项目和运行时代码中均声明版本 0.1.0，Ankus 根据控制文件与 SQL 属性生成匹配的 control 和带版本 SQL。审查时尚无标签发布。上游支持 PostgreSQL 15–19。

### 配置两端节点

在提供方与订阅方安装匹配文件，保留已有预加载库，配置逻辑 WAL 和足够的工作进程、槽及发送进程，然后重启 PostgreSQL：

```ini
wal_level = logical
shared_preload_libraries = 'worf'
max_worker_processes = 10
max_replication_slots = 10
max_wal_senders = 10
worf.conflict_resolution = 'error'
```

容量值仅为示例，应按实际节点及其他使用方调整。PostgreSQL 15.19、16.15、17.11、18.6、19 或更新版本，还须在提供方的输出插件允许列表中加入 worf，并保留其他需要的插件：

```ini
output_plugin_libraries = 'pgoutput, test_decoding, worf'
```

管理操作要求超级用户权限。每个参与数据库都须安装扩展，并在两端创建一致的模式、表、列类型及主键；订阅方约束不能比提供方更严格。

### 提供方与订阅方

在提供方执行：

```sql
CREATE EXTENSION worf;
SELECT worf.create_node(
  node_name := 'provider1',
  dsn := 'host=providerhost port=5432 dbname=app'
);
SELECT worf.replication_set_add_all_tables('default', ARRAY['public']);
```

在订阅方执行：

```sql
CREATE EXTENSION worf;
SELECT worf.create_node(
  node_name := 'subscriber1',
  dsn := 'host=subscriberhost port=5432 dbname=app'
);
SELECT worf.create_subscription(
  subscription_name := 'subscription1',
  provider_dsn := 'host=providerhost port=5432 dbname=app'
);
SELECT worf.wait_for_subscription_sync_complete('subscription1');
SELECT * FROM worf.show_subscription_status();
```

提供方须允许普通数据库复制连接，并保护连接凭据。固定的 `worf` 模式保存节点、订阅和复制集状态。`worf.replicate_ddl_command` 可将明确选择的 DDL 操作排入队列，普通 DDL 不会自动复制。`worf_create_subscriber` 通过完整基础备份克隆数据库，不会把复制范围限制到所选复制集。

### 冲突与维护边界

默认冲突策略为 `apply_remote`，可能覆盖本地行；示例主动选择 `error`，以便停止并审查。基于时间戳的策略和 `keep_local` 要求 `track_commit_timestamp`。多上游冲突处理不代表具备完整多主语义。应监控复制槽与 WAL 保留、同步状态及冲突日志。

UPDATE/DELETE 要求可用主键或副本标识，不支持 `REPLICA IDENTITY FULL`。额外唯一约束可能使冲突应用停止，且不支持可延迟唯一索引。应用期间不强制检查外键约束。UNLOGGED/TEMPORARY 表与大对象不复制，序列状态按周期传播而非逐次递增传播。表重同步会立即截断表，并在复制期间保持为空。应把模式变更和同步安排为维护操作，不应将其视为自动故障转移保证。
