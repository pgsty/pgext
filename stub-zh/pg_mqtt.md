## 用法

来源：

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/LICENSE)
- [extensions/pg_mqtt/pg_mqtt.control](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/pg_mqtt.control)
- [extensions/pg_mqtt/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/Cargo.toml)
- [extensions/pg_mqtt/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/src/lib.rs)
- [extensions/pg_mqtt/sql/pg_mqtt--0.3.2--0.4.0.sql](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/sql/pg_mqtt--0.3.2--0.4.0.sql)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/crates/pg_bgworker/src/supervision.rs)
- [extensions/pg_mqtt/src/protocol/codec.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/src/protocol/codec.rs)
- [extensions/pg_mqtt/src/protocol/handlers.rs](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/src/protocol/handlers.rs)
- [extensions/pg_mqtt/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/867872ab0ab5ec5952a2f3ff1b7904636b6d17b7/extensions/pg_mqtt/pgbrew.toml)

`pg_mqtt` 0.4.2 使用 `pgmqtt` 表保存消息并实现 MQTT 代理。所核对解码器只接受协议级别 5，拒绝其他版本；扩展 README 中旧的 3.1.1 声明已经过时。

### 核心用法

```ini
shared_preload_libraries = 'pg_mqtt'
pg_mqtt.database = 'postgres'
pg_mqtt.port = 1883
```

```sql
CREATE EXTENSION pg_mqtt;
SELECT pgmqtt.mqtt_publish('sensors/room1/temperature', '22.5');
SELECT * FROM pgmqtt.mqtt_messages('sensors/%', 10);
```

### 运行边界

控制文件要求超级用户安装。需要预加载并重启，在配置的工作数据库中创建扩展。`pg_mqtt.database`、`pg_mqtt.port` 和 `pg_mqtt.worker_count` 配置工作进程。源码默认在所有接口的 1883 端口监听，应隔离这一实验性服务。`mqtt_publish`、`mqtt_messages`、`mqtt_sessions` 及模式绑定函数提供数据接口；`mqtt_messages` 使用 SQL LIKE 模式；`mqtt_status` 只报告配置，不是监听端口健康检查。不能仅根据项目描述推断完整的 MQTT 安全或交付保证。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。

### 当前版本与升级

扩展版本 0.4.0 通过 `pg_mqtt.database` 选择后台工作进程使用的数据库。应在该库创建扩展；扩展尚不存在时，工作进程会等待，不再反复退出。调整需要重启的工作进程配置后须重启。 对扩展 0.3.0 及之后的版本，安装匹配文件后使用 ALTER EXTENSION UPDATE；更早版本仍需前述迁移。

### 工作进程监督

0.3.2 新增 `pgmqtt.worker_status()` 与仅超级用户可调用的 `pgmqtt.reset_workers()`。状态返回工作进程名称、状态、失败次数、重启次数、PID、进入当前状态的时间与最近一次失败信息。`pg_mqtt.max_worker_failures` 默认为连续失败 10 次；设为 0 表示无限重试。重启退避时间从 5 秒增长到 60 秒，达到限制后失败的工作进程保持空闲，直到重置。

安装配套共享库后重启 PostgreSQL，初始化新增共享内存，再在配置的工作数据库中更新扩展 SQL。重置前应先排查失败原因；重置不会修复故障本身。

```sql
ALTER EXTENSION pg_mqtt UPDATE;
SELECT * FROM pgmqtt.worker_status();
```

### 本次版本变化

源代码声明 PostgreSQL 14–18 功能开关。仓库发布版本为 0.7.1，各扩展保留各自的 SQL 版本；这不代表 Pigsty 软件包已构建。

0.4.0 通过非日志表 `pgmqtt.outbox` 向各工作进程的订阅者分发消息，默认轮询间隔为 50 毫秒。SQL `mqtt_publish` 在事务提交后才投递，回滚不投递；集群中其他长时间写事务也可能延迟分发。非日志队列不提供崩溃后的持久交付保证。升级脚本会创建队列，需先安装匹配的共享库再更新扩展 SQL。

### 0.4.1 类型化主题修复

0.4.1 修复类型化主题的底层表没有 `id` 列时 `mqtt_publish` 失败的问题：发布路径不再假定存在 `RETURNING id`。SQL 模式不变。应安装匹配的 0.4.1 文件并按扩展升级路径更新；上述事务发件箱与交付限制仍适用。

### 0.4.2 的空闲工作进程

仓库发布 0.7.2 包含扩展版本 0.4.2。只读轮询或未领取到作业的轮询不再仅因使用 SPI 桥接而分配事务 ID，实际写入行时仍会分配，因而空闲轮询不再产生空提交 WAL。 新的桥接请求会唤醒工作进程的 latch；无请求时主循环可休眠至多 100 ms。 安装匹配的共享库、重启预加载的工作进程并更新 SQL 扩展。既有权限、外部服务和 0.3.0 之前版本的迁移边界仍然适用。
